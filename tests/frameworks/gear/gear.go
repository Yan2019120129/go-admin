package gear

import (
	// add gin adapter
	ada "ProgrammerYan/go-admin/adapter/gear"
	"github.com/teambition/gear"

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
	"ProgrammerYan/go-admin/modules/config"
	"ProgrammerYan/go-admin/modules/language"
	"ProgrammerYan/go-admin/plugins/admin/modules/table"
	"ProgrammerYan/go-admin/template"
	"ProgrammerYan/go-admin/template/chartjs"
	"ProgrammerYan/go-admin/tests/tables"
)

func internalHandler() http.Handler {
	app := gear.New()

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddGenerators(tables.Generators).
		AddGenerator("user", tables.GetUserTable).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	app := gear.New()

	eng := engine.Default()

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
		AddAdapter(new(ada.Gear)).
		AddGenerators(gens).
		Use(app); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	return app
}
