package response

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/mall/pkg/e"
)

const (
	SUCCESS = 1
	FAIL    = 0
)

type body struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
	Errs    any    `json:"errs"`
}

func Build(c *fiber.Ctx, err error, data any) error {
	if err == nil {
		c.SendStatus(http.StatusOK)
		return c.JSON(&body{
			Code:    SUCCESS,
			Success: true,
			Status:  http.StatusOK,
			Msg:     "success",
			Data:    data,
			Errs:    nil,
		})
	}

	var myErr *e.Error

	if !errors.As(err, &myErr) {
		c.SendStatus(http.StatusInternalServerError)
		return c.JSON(&body{
			Code:    FAIL,
			Success: false,
			Status:  http.StatusInternalServerError,
			Msg:     "Unknown Error",
			Data:    nil,
			Errs:    nil,
		})
	}

	c.SendStatus(myErr.Code)
	return c.JSON(&body{
		Code:    FAIL,
		Success: false,
		Status:  myErr.Code,
		Msg:     myErr.Msg,
		Data:    data,
		Errs:    myErr.Err,
	})
}
