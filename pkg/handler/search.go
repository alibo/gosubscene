package handler

import (
	"fmt"
	"github.com/alibo/gosubscene/pkg/subscene"
	"github.com/gofiber/fiber"
	"net/http"
)

func Search(c *fiber.Ctx) {
	result, err := subscene.Search(c.Query("q"))

	if err != nil {
		c.Status(http.StatusInternalServerError).JSON(Response{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
			Data:   nil,
		})
		return
	}

	if !result.Found {
		c.Status(http.StatusNotFound).JSON(Response{
			Status: http.StatusNotFound,
			Error:  fmt.Errorf("%q is not found", c.Query("q")).Error(),
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
