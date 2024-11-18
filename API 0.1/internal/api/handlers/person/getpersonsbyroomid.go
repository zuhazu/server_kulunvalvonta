package person

import (
	"context"
	"encoding/json"
	service "goapi/internal/api/service/person"
	"log"
	"net/http"
	"time"
)

// Otetaan url:sta id ja haetaan sen perusteella sql tietokannasta henkilöt joilla on kyseinen room id
func GetPersonsByRoomIdHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.PersonService) {

	id := r.PathValue("id")

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	data, err := ds.ReadPersonsByRoomId(id, ctx)
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
