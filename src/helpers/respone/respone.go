package respone

import (
	"net/http"

	"github.com/labstack/echo"
)

// JSON returns a well formated response with a status code

func JSON(c echo.Context, message string, data interface{}) (interface{}, int) {

	result := struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Err     string      `json:"err"`
		Data    interface{} `json:"data"`
	}{
		Status:  true,
		Message: message,
		Data:    data,
		Err:     "",
	}
	return result, http.StatusOK
}

// ERROR returns a jsonified error response along with a status code.
func ERROR(c echo.Context, ms string, err error) (interface{}, int) {

	var er string
	if err != nil {
		er = err.Error()
	} else {
		er = ""
	}

	data := struct {
		Status  bool        `json:"status"`
		Message string      `json:"message"`
		Err     string      `json:"err"`
		Data    interface{} `json:"data"`
	}{
		Status:  false,
		Message: ms,
		Err:     er,
		Data:    nil,
	}
	return data, http.StatusBadRequest
}
