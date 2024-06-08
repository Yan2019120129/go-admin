package iris

import (
	// add iris adapter
	_ "ProgrammerYan/go-admin/adapter/iris"
	"ProgrammerYan/go-admin/modules/config"
	"ProgrammerYan/go-admin/modules/language"
	"ProgrammerYan/go-admin/plugins/admin/modules/table"

	// add mysql driver
	_ "ProgrammerYan/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "ProgrammerYan/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "ProgrammerYan/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "ProgrammerYan/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	"github.com/GoAdminGroup/themes/adminlte"

	"ProgrammerYan/go-admin/template"
	"ProgrammerYan/go-admin/template/chartjs"

	"net/http"
	"os"

	"ProgrammerYan/go-admin/engine"
	"ProgrammerYan/go-admin/plugins/admin"
	"ProgrammerYan/go-admin/plugins/example"
	"ProgrammerYan/go-admin/tests/tables"
	"github.com/kataras/iris/v12"
)

func internalHandler() http.Handler {
	app := iris.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)
	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	if err := app.Build(); err != nil {
		panic(err)
	}

	return app.Router
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	app := iris.New()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(gens)

	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfig(&config.Config{
		Databases: dbs,
		UrlPrefix: "admin",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language:    language.EN,
		IndexUrl:    "/",
		Debug:       true,
		ColorScheme: adminlte.ColorschemeSkinBlack,
	}).
		AddPlugins(adminPlugin, examplePlugin).Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	if err := app.Build(); err != nil {
		panic(err)
	}

	return app.Router
}
