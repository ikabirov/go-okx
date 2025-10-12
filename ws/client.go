package ws

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	EndpointPublic           = "wss://ws.okx.com:8443/ws/v5/public"
	EndpointPrivate          = "wss://ws.okx.com:8443/ws/v5/private"
	EndpointBusiness         = "wss://ws.okx.com:8443/ws/v5/business"
	EndpointPublicSimulated  = "wss://wspap.okx.com:8443/ws/v5/public?brokerId=9999"
	EndpointPrivateSimulated = "wss://wspap.okx.com:8443/ws/v5/private?brokerId=9999"

	PingTimeout  = 20 * time.Second
	PingDeadline = 10 * time.Second
)

var (
	DefaultClientPublic           = NewClient(EndpointPublic)
	DefaultClientPrivate          = NewClient(EndpointPrivate)
	DefaultClientBusiness         = NewClient(EndpointBusiness)
	DefaultClientPublicSimulated  = NewClient(EndpointPublicSimulated)
	DefaultClientPrivateSimulated = NewClient(EndpointPrivateSimulated)

	PingMessage = []byte("ping")
)

type OperateCallback func(*websocket.Conn) error

type Client struct {
	Endpoint string
	Dialer   *websocket.Dialer
}

// new Client
func NewClient(endpoint string) *Client {
	return &Client{
		Endpoint: endpoint,
	}
}

// operate
func (c *Client) Operate(ctx context.Context, operate *Operate, callback OperateCallback) error {
	conn, _, err := c.dial(ctx)
	if err != nil {
		return err
	}

	if callback != nil {
		if err := callback(conn); err != nil {
			return err
		}
	}

	if err := c.MessageOperate(conn, operate); err != nil {
		return err
	}

	if operate.Handler != nil {
		ticker := time.NewTicker(PingTimeout)
		go c.keepAlive(ctx, conn, ticker)
		go c.messageLoop(ctx, conn, operate)
	}

	return nil
}

// message operate
func (c *Client) MessageOperate(conn *websocket.Conn, operate *Operate) error {
	if operate.Request == nil {
		return nil
	}
	if err := conn.WriteJSON(operate.Request); err != nil {
		return err
	}
	if err := conn.ReadJSON(&operate.Response); err != nil {
		return err
	}
	return operate.Response.Error()
}

// loop websocket message
func (c *Client) messageLoop(ctx context.Context, conn *websocket.Conn, operate *Operate) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				operate.HandlerError(err)
				return
			}
			operate.Handler(message)
		}
	}
}

// keep websocket alive
func (c *Client) keepAlive(ctx context.Context, conn *websocket.Conn, ticker *time.Ticker) {
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			deadline := time.Now().Add(PingDeadline)
			if err := conn.WriteControl(websocket.PingMessage, PingMessage, deadline); err != nil {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

// dial endpoint
func (c *Client) dial(ctx context.Context) (*websocket.Conn, *http.Response, error) {
	if c.Dialer == nil {
		c.Dialer = websocket.DefaultDialer
	}
	return c.Dialer.DialContext(ctx, c.Endpoint, nil)
}
