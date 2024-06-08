package example

import (
	"github.com/Yan2019120129/go-admin/context"
	"github.com/Yan2019120129/go-admin/modules/auth"
	"github.com/Yan2019120129/go-admin/modules/db"
	"github.com/Yan2019120129/go-admin/modules/service"
)

func (e *Example) initRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), e.TestHandler)

	return app
}
