package main

import (
	"backend/api/jobs"
	"backend/db"
	"backend/db/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func initDB() *pg.DB {
	conn := db.Connect()

	var jobs []models.Job

	b, err := ioutil.ReadFile("jobs.json")
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(b, &jobs)

	if err != nil {
		fmt.Print(err)
	}

	//FIXME сделать проверку при добавлении
	for _, job := range jobs {

		err = job.Insert(conn)
		if err != nil {
			fmt.Print(err)
		}

	}

	return conn
}

func main() {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn := initDB()

	// Echo instance
	apiPublic := echo.New()

	apiPublic.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	apiPublic.Use(middleware.Logger())
	apiPublic.Use(middleware.Recover())

	// Routes
	apiPublic.GET("/", hello)

	apiPublic.GET("/job/all", jobs.GetAllVacancy(conn))

	apiPublic.GET("/job/full/:uuid", jobs.GetWithUUID(conn))

	// apiPublic.POST("/place/create", places.Create(conn))
	// apiPublic.POST("/place/import", places.Import(conn))
	// apiPublic.POST("/place/update", places.Update(conn))
	// apiPublic.GET("/place/all", places.GetAllPlaces(conn))
	// apiPublic.GET("/place/name/:name/find", places.FindLikeName(conn))

	// apiPublic.POST("/types/import", types.Import(conn))
	// apiPublic.GET("/types/all", types.GetAll(conn))

	// apiPublic.POST("/item/create", items.Create(conn))
	// apiPublic.POST("/item/import", items.Import(conn))
	// apiPublic.POST("/item/update", items.Update(conn))
	// apiPublic.GET("/item/all", items.GetAllItems(conn))
	// apiPublic.GET("/item/current/:place_uuid/all", items.GetFromCurretPlace(conn))
	// apiPublic.GET("/item/serial/:serial_number/find", items.FindLikeSerialNumber(conn))
	// apiPublic.GET("/item/internal/:internal_number/find", items.FindLikeInternalNumber(conn))
	// apiPublic.GET("/item/uuid/:uuid/find", items.FindLikeUUID(conn))

	// apiPublic.POST("/migration/create", migration.Create(conn))
	// apiPublic.GET("/migration/all", migration.GetAllMigrations(conn))

	// Start server
	//FIXME: remove port
	apiPublic.Logger.Fatal(apiPublic.Start(":8008"))

}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
