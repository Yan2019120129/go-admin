package gin

import (
	// add gin adapter

	ada "github.com/Yan2019120129/go-admin/adapter/gin"
	// add mysql driver
	_ "github.com/Yan2019120129/go-admin/modules/db/drivers/mysql"
	// add postgresql driver
	_ "github.com/Yan2019120129/go-admin/modules/db/drivers/postgres"
	// add sqlite driver
	_ "github.com/Yan2019120129/go-admin/modules/db/drivers/sqlite"
	// add mssql driver
	_ "github.com/Yan2019120129/go-admin/modules/db/drivers/mssql"
	// add adminlte ui theme
	"github.com/GoAdminGroup/themes/adminlte"

	"net/http"
	"os"

	"github.com/Yan2019120129/go-admin/engine"
	"github.com/Yan2019120129/go-admin/modules/config"
	"github.com/Yan2019120129/go-admin/modules/language"
	"github.com/Yan2019120129/go-admin/plugins/admin/modules/table"
	"github.com/Yan2019120129/go-admin/template"
	"github.com/Yan2019120129/go-admin/template/chartjs"
	"github.com/Yan2019120129/go-admin/tests/tables"
	"github.com/gin-gonic/gin"
)

func internalHandler() http.Handler {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	eng := engine.Default()

	template.AddComp(chartjs.NewChart())

	if err := eng.AddConfigFromJSON(os.Args[len(os.Args)-1]).
		AddGenerators(tables.Generators).
		AddGenerator("user", tables.GetUserTable).
		Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	r.Static("/uploads", "./uploads")

	return r
}

func NewHandler(dbs config.DatabaseList, gens table.GeneratorList) http.Handler {
	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

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
		AddAdapter(new(ada.Gin)).
		AddGenerators(gens).
		Use(r); err != nil {
		panic(err)
	}

	eng.HTML("GET", "/admin", tables.GetContent)

	r.Static("/uploads", "./uploads")

	return r
}
