package response

import (
	"net/http"

	"github.com/labstack/echo"
)

const (
	statusOk = "Ok"
)

const (
	codeOk = 1
)

func QueuesResponse(c *echo.Context, queues interface{}) {
	c.JSON(http.StatusOK, GenerateSuccessResponse(queues))
}

func QueueResponse(c *echo.Context, queue interface{}) {
	c.JSON(http.StatusOK, GenerateSuccessResponse(queue))
}

func LoginResponseHandler(c *echo.Context, identity interface{}, accessToken string, refreshToken string) {
	c.JSON(http.StatusOK, GenerateSuccessResponse(map[string]interface{}{
		"AccessToken":  accessToken,
		"RefreshToken": refreshToken,
	}))
}

func RefreshResponseHandler(c *echo.Context, identity interface{}, accessToken string, refreshToken string) {
	c.JSON(http.StatusOK, GenerateSuccessResponse(map[string]interface{}{
		"AccessToken":  accessToken,
		"RefreshToken": refreshToken,
	}))
}

func EmptyResponseHandler(c *echo.Context) {
	c.JSON(http.StatusOK, GenerateSuccessResponse(map[string]interface{}{}))
}
