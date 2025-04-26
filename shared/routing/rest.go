package routing

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type RestError struct {
	Code    int
	Message string
}

func (e *RestError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func RestPostHandleFunc[T any, U any](restFunc func(*T) (U, error)) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodPost {
			http.Error(writer, "Invalid HTTP method", http.StatusMethodNotAllowed)
			return
		}

		var data T
		decoder := json.NewDecoder(request.Body)
		if err := decoder.Decode(&data); err != nil {
			log.Println("Error:", err)
			http.Error(writer, "Invalid JSON", http.StatusBadRequest)
			return
		}
		result, err := restFunc(&data)
		if err != nil {
			var restErr *RestError
			if errors.As(err, &restErr) {
				http.Error(writer, restErr.Message, restErr.Code)
			} else {
				log.Println("Error:", err)
				http.Error(writer, "Internal error", http.StatusInternalServerError)
			}
			return
		}

		if err := json.NewEncoder(writer).Encode(result); err != nil {
			log.Println("Error:", err)
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
			return
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
	}
}
