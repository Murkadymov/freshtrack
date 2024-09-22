package logger

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"time"
)

func NewMiddlewareLogger(log *slog.Logger, next echo.HandlerFunc) echo.HandlerFunc {

	log = slog.With(
		slog.String("copmonent", "middleware/logger"),
	)
	log.Info("middleware logger enabled")

	return func(c echo.Context) error {

		stats := log.With(

			slog.String("method", c.Request().Method),
			slog.String("path", c.Request().URL.Path),
			slog.String("user-agent", c.Request().UserAgent()),
		)
		t1 := time.Now()

		err := next(c)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err,
			)
		}

		defer func() {
			stats.Info(
				"request completed",
				slog.String("time taken: ", time.Since(t1).String()),
			)
		}()

		return nil
	}

}
