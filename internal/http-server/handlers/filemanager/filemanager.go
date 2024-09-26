package filemanager

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

type FileManager struct {
}

func NewFileManager() *FileManager {
	return &FileManager{}
}

const uploadPath = "E:/Projects/freshtrack/uploaded-files"

func (f *FileManager) UploadFile(log *slog.Logger) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {

		log = log.With("component", "filemanager/uploadLogger")

		file, err := c.FormFile("file")
		if err != nil {
			return c.String(http.StatusBadRequest, "Failed to get file")
		}

		src, err := file.Open()
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				error.Error(err),
			)
		}
		defer src.Close()

		if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
			os.Mkdir(uploadPath, os.ModePerm)
		}

		dst, err := os.Create(filepath.Join(uploadPath, file.Filename))
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				echo.Map{"error": error.Error(err), "filename": file.Filename},
			)
		}
		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			return c.String(http.StatusInternalServerError, "Failed to save file")
		}

		return c.JSON(
			http.StatusCreated,
			fmt.Sprintf("file has been created with the name: %s", dst.Name()),
		)
	})
}
