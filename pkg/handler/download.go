package handler

import (
	"github.com/alibo/gosubscene/pkg/subscene"
	"github.com/gofiber/fiber"
	"io"
	"net/http"
	"strconv"
)

func Download(c *fiber.Ctx) {
	token := c.Query("token")

	err := subscene.Download(token, func(contentDisposition, contentType string, size int, reader io.Reader) {
		c.Set("Content-Disposition", contentDisposition)
		c.Set("Content-Length", strconv.Itoa(size))
		c.Set("Content-Type", contentType)
		c.Fasthttp.SetBodyStream(reader, size)
	})

	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "Not Found" {
			status = http.StatusNotFound
		}

		c.Status(status).JSON(Response{
			Status: status,
			Error:  err.Error(),
			Data:   nil,
		})
		return
	}
}
