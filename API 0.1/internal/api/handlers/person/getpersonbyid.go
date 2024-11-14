package person

import (
	"context"
	"encoding/json"
	service "goapi/internal/api/service/person"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetPersonByIDHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.PersonService) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		// * This is a User Error: format of id is invalid, response in JSON and with a 400 status code
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Missconfigured ID."}`))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	data, err := ds.ReadOnePerson(id, ctx)
	if err != nil {
		logger.Println("Could not read one:", err, id)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
	if data == nil {
		// * This is a User Error, response in JSON and with a 404 status code
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Resource not found."}`))
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Println("Error encoding data:", err, data)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
}
