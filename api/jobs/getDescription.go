package jobs

import (
	"backend/db/models"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
)

func GetWithUUID(conn *pg.DB) func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		job := &models.Job{}
		job.UUID = ctx.Param("uuid")

		err := job.FindWithUUID(conn)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest,
				struct{ Error string }{err.Error()})
		}
		return ctx.JSON(http.StatusOK, job)
	}
}
