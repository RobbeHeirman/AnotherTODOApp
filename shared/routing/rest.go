package routing

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONSerializable interface {
	json.Marshaler
	json.Unmarshaler
}
type RestPostFunc[T JSONSerializable, U JSONSerializable] func(*T, http.ResponseWriter, *http.Request) U

func RestPostHandler[T JSONSerializable, U JSONSerializable](restFunc RestPostFunc[T, U]) http.HandlerFunc {
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
		}

		result := restFunc(&data, writer, request)
		if err := json.NewEncoder(writer).Encode(result); err != nil {
			log.Println("Error:", err)
			http.Error(writer, "Internal server error", http.StatusInternalServerError)
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusCreated)
	}
}
