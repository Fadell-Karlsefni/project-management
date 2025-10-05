package utils

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status       string      `json:"status"`
	ResponseCode int         `json:"response_code"`
	Messege      string      `json:"messege,omitempty"`
	Data         interface{} `json:"data,omitempty"`
	Error        string      `json:"error,omitempty"`
}

func Success(c *fiber.Ctx, messege string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Status:       "Success",
		ResponseCode: fiber.StatusOK,
		Messege:      messege,
		Data:         data,
	})
}

func Created(c *fiber.Ctx, messege string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Status:       "Created",
		ResponseCode: fiber.StatusCreated,
		Messege:      messege,
		Data:         data,
	})
}

func BadRequest(c *fiber.Ctx, messege string, err string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Status:       "Error Bad Request",
		ResponseCode: fiber.StatusBadRequest,
		Messege:      messege,
		Error: err,
	})
}

func NotFound(c *fiber.Ctx, messege string, err string) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Status:       "Error Not Found",
		ResponseCode: fiber.StatusNotFound,
		Messege:      messege,
		Error: err,
	})
}

func Unauthorized(c *fiber.Ctx, messege string, err string) error {
	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Status:       "Error Unauthorized",
		ResponseCode: fiber.StatusUnauthorized,
		Messege:      messege,
		Error: err,
	})
}
