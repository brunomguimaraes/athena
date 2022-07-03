package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSON(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.WriteHeader(statusCode)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		fmt.Fprintf(writer, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
