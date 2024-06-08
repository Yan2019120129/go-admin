package example

import (
	c "ProgrammerYan/go-admin/modules/config"
	"ProgrammerYan/go-admin/modules/service"
	"ProgrammerYan/go-admin/plugins"
)

type Example struct {
	*plugins.Base
}

func NewExample() *Example {
	return &Example{
		Base: &plugins.Base{PlugName: "example"},
	}
}

func (e *Example) InitPlugin(srv service.List) {
	e.InitBase(srv, "example")
	e.App = e.initRouter(c.Prefix(), srv)
}
