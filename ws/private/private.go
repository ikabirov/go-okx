package private

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/ikabirov/go-okx/common"
	"github.com/ikabirov/go-okx/ws"
)

type Private struct {
	Auth common.Auth
	C    *ws.Client
}

// new Private
func NewPrivate(auth common.Auth) *Private {
	private := &Private{
		Auth: auth,
		C:    ws.DefaultClientPrivate,
	}
	if auth.Simulated {
		private.C = ws.DefaultClientPrivateSimulated
	}
	return private
}

// subscribe
func (p *Private) Subscribe(ctx context.Context, args any, handler ws.Handler, handlerError ws.HandlerError) error {
	subscribe := ws.NewOperateSubscribe(args, handler, handlerError)
	return p.C.Operate(ctx, subscribe, p.login)
}

// loging private
func (p *Private) login(conn *websocket.Conn) error {
	args := ws.NewArgsLoginFromAuth(p.Auth)
	login := ws.NewOperateLogin(args)
	return p.C.MessageOperate(conn, login)
}
