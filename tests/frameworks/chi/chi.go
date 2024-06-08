package chi

import (
	// add chi adapter
	_ "ProgrammerYan/go-admin/adapter/chi"
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

	"net/http"
	"os"

	"ProgrammerYan/go-admin/engine"
	"ProgrammerYan/go-admin/plugins/admin"
	"ProgrammerYan/go-admin/plugins/example"
	"ProgrammerYan/go-admin/template"
	"ProgrammerYan/go-admin/template/chartjs"
	"ProgrammerYan/go-admin/tests/tables"
	"github.com/go-chi/chi"
)

func internalHandler() http.Handler {
	r := chi.NewRouter()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	adminPlugin.AddGenerator("user", tables.GetUserTable)
	examplePlugin := example.NewExample()
	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddPlugins(adminPlugin, examplePlugin).Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return r
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	r := chi.NewRouter()

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(gens)
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
		AddPlugins(adminPlugin).Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return r
}
