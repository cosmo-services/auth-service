package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"main/internal/config"
	"time"

	"github.com/nats-io/nats.go"
)

type Nats struct {
	*nats.Conn
	JS nats.JetStreamContext
}

func NewNats(logger Logger, env config.Env) *Nats {
	connStr := fmt.Sprintf("nats://%s:%s", env.NatsHost, env.NatsPort)
	logger.Info(fmt.Sprintf("Connecting to NATS at %s", connStr))

	nc, err := nats.Connect(connStr,
		nats.Timeout(10*time.Second),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
		nats.ReconnectWait(2*time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			logger.Warn(fmt.Sprintf("Disconnected from NATS: %v", err))
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			logger.Info("Reconnected to NATS")
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			logger.Info("NATS connection closed")
		}),
	)

	if err != nil {
		logger.Fatal("Failed to connect to NATS: ", err)
	}

	if !nc.IsConnected() {
		logger.Fatal("NATS connection is not established")
	}
	logger.Info("Successfully connected to NATS")

	js, err := nc.JetStream(
		nats.MaxWait(10*time.Second),
		nats.PublishAsyncMaxPending(256),
	)
	if err != nil {
		logger.Fatal("Failed to create JetStream context: ", err)
	}
	logger.Info("Successfully initialized JetStream")

	return &Nats{
		Conn: nc,
		JS:   js,
	}
}

func (n *Nats) Close() error {
	if n.Conn != nil && !n.Conn.IsClosed() {
		n.Conn.Drain()
		n.Conn.Close()
	}
	return nil
}

func (n *Nats) PublishJSON(ctx context.Context, subject string, data interface{}) (*nats.PubAck, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}
	ack, err := n.JS.Publish(subject, jsonData)
	if err != nil {
		return nil, fmt.Errorf("failed to publish to NATS: %w", err)
	}

	return ack, nil
}

func (n *Nats) Subscribe(subject string, handler func(msg *nats.Msg) error) (*nats.Subscription, error) {
	return n.JS.Subscribe(subject, func(msg *nats.Msg) {
		if err := handler(msg); err != nil {
			msg.NakWithDelay(time.Minute)
		} else {
			msg.Ack()
		}
	}, nats.ManualAck())
}

func (n *Nats) SubscribePull(subject, consumer, stream string) (*nats.Subscription, error) {
	return n.JS.PullSubscribe(subject, consumer, nats.BindStream(stream))
}
