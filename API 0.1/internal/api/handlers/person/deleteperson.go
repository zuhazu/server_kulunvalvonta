package person

import (
	"context"
	"goapi/internal/api/repository/models"
	service "goapi/internal/api/service/person"
	"log"
	"net/http"
	"strconv"
	"time"
)

func DeletePersonHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.PersonService) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		// * This is a User Error: format of id is invalid, response in JSON and with a 400 status code
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Missconfigured ID."}`))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	aff, err := ds.DeletePerson(&models.Person{ID: id}, ctx)
	if err != nil {
		logger.Println("Could not delete data:", err, id)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	// * Check if the data was found and deleted
	if aff == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Resource not found."}`))
		return
	}

	// * This is a Success, response in JSON and with a 204 status code when data was successfully deleted
	w.WriteHeader(http.StatusNoContent)
}
