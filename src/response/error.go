package response

import (
	"net/http"

	"github.com/labstack/echo"
)

const (
	statusError = "Error"
)

const (
	codeRegisterParamsError        = 11
	codeRegisterAlreadyExistsError = 12
	codeAuthError                  = 13
	codeInvalidHeader              = 14
	codeInvalidToken               = 15
	codeExpiredToken               = 16
	codeAdminRequired              = 17
	codeQueueNotFound              = 18
	codeQueueCreateError           = 19
	codeQueueRemoveError           = 20
	codeAlreadyInQueue             = 21
	codeNotFoundInQueue            = 22
)

func AuthErrorHandler(c *echo.Context) {
	c.JSON(http.StatusForbidden, GenerateErrorResponse(codeAuthError, "Auth error"))
}

func HeaderInvalidHandler(c *echo.Context) {
	c.JSON(http.StatusBadRequest, GenerateErrorResponse(codeInvalidHeader,
		"Token header not found or has not valid format"))
}

func TokenInvalidHandler(c *echo.Context) {
	c.JSON(http.StatusForbidden, GenerateErrorResponse(codeInvalidToken, "Token not valid"))
}

func TokenExpireHandler(c *echo.Context) {
	c.JSON(http.StatusForbidden, GenerateErrorResponse(codeExpiredToken, "Token expired"))
}

// TODO: auth required

func AdminRequiredHandler(c *echo.Context) {
	c.JSON(http.StatusForbidden, GenerateErrorResponse(codeAdminRequired, "Admin required error"))
}

func QueueNotFoundHandler(c *echo.Context) {
	c.JSON(http.StatusNotFound, GenerateErrorResponse(codeQueueNotFound, "Queue not found"))
}

func QueueCreateError(c *echo.Context) {
	c.JSON(http.StatusBadRequest, GenerateErrorResponse(codeQueueCreateError, "Create queue error"))
}

func QueueParamsCreateError(c *echo.Context) {
	c.JSON(http.StatusBadRequest, GenerateErrorResponse(codeQueueCreateError, "Invalid params"))
}

func QueueRemoveError(c *echo.Context) {
	c.JSON(http.StatusBadRequest, GenerateErrorResponse(codeQueueRemoveError, "Remove queue error"))
}

func QueueRemoveNotPermitted(c *echo.Context) {
	c.JSON(http.StatusForbidden, GenerateErrorResponse(codeQueueRemoveError, "Remove queue not permitted"))
}

func RegisterParamsError(c *echo.Context) {
	c.JSON(http.StatusBadRequest, GenerateErrorResponse(codeRegisterParamsError, "Invalid params"))
}

func RegisterAlreadyExistsError(c *echo.Context) {
	c.JSON(http.StatusBadRequest, GenerateErrorResponse(codeRegisterAlreadyExistsError, "User with this email already exists"))
}

func AlreadyInQueueError(c *echo.Context) {
	c.JSON(http.StatusBadRequest, GenerateErrorResponse(codeAlreadyInQueue, "Already in queue"))
}

func NotFoundInQueueError(c *echo.Context) {
	c.JSON(http.StatusNotFound, GenerateErrorResponse(codeNotFoundInQueue, "Not found in queue"))
}
