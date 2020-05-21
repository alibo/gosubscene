package handler

import (
	"github.com/alibo/gosubscene/pkg/subscene"
	"github.com/gofiber/fiber"
	"net/http"
)

func List(c *fiber.Ctx) {
	result, err := subscene.ListSubtitles(c.Params("name"))

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

	c.JSON(Response{
		Status: http.StatusOK,
		Error:  "",
		Data:   result,
	})
}
