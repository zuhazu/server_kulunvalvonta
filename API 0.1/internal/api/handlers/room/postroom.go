package room

import (
	"context"
	"encoding/json"
	"goapi/internal/api/repository/models"
	service "goapi/internal/api/service/room"
	"log"
	"net/http"
	"time"
)

// * User sends a POST request to /data with a JSON payload in the request body *
// * curl -X POST http://127.0.0.1:8080/data -i -u admin:password -H "Content-Type: application/json" -d '{"room_id": "room1", "room_name": "Room 1", "room_description": "Description of Room 1"}'
func PostRoomHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.RoomService) {
	var data models.Room

	// * Decode the JSON payload from the request body into the data struct
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {

		// * This is a User Error: format of body is invalid, response in JSON and with a 400 status code
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request data. Please check your input."}`))
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// * Try to create the room data in the database
	if err := ds.CreateRoom(&data, ctx); err != nil {
		switch err.(type) {
		case service.RoomError:
			// * If the error is a RoomError, handle it as a client error
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "` + err.Error() + `"}`))
			return
		default:
			// * If it is not a RoomError, handle it as a server error
			logger.Println("Error creating room:", err, data)
			http.Error(w, "Internal server error.", http.StatusInternalServerError)
			return
		}
	}

	// * Return the room data to the user as JSON with a 201 Created status code
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Println("Error encoding room data:", err, data)
		http.Error(w, "Internal server error.", http.StatusInternalServerError)
		return
	}
}
