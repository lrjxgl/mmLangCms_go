package indexIndex

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

/*@@IndexIndex@@*/
func IndexIndex(c echo.Context) (err error) {

	reData := make(map[string]interface{})
	reData["error"] = 0
	reData["message"] = "success"

	reJson := make(map[string]interface{})
	reJson["error"] = 0
	reJson["message"] = "success"
	reJson["data"] = reData
	return c.JSON(http.StatusOK, reJson)

}
