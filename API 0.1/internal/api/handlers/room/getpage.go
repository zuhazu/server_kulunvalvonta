package room

import (
	"context"
	page_service "goapi/internal/api/service/page"
	service "goapi/internal/api/service/room"
	"log"
	"net/http"
	"time"
)

// * User sends a POST request to /data with a JSON payload in the request body *
// * curl -X POST http://127.0.0.1:8080/data -i -u admin:password -H "Content-Type: application/json" -d '{"room_id": "room1", "room_name": "Room 1", "room_description": "Description of Room 1"}'
func GetPageHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.RoomService) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	data, err := ds.GetPersonsByRoomID(r.URL.Query().Get("room"), ctx)
	logger.Print(r.URL.Query().Get("room"))
	if err != nil || data == nil {
		logger.Println("Could not get data:", err, data)
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(page_service.GetPageModel(data)))
	// 	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	// 	if err != nil {
	// 		if err == err.(*strconv.NumError) {
	// 			// * There was no page specified, so we default to 0 *
	// 			page = 0
	// 		} else {
	// 			// * Invalid page specified, return a 400 status code *
	// 			w.WriteHeader(http.StatusBadRequest)
	// 			w.Write([]byte(`{"error": "Invalid page specified."}`))
	// 			return
	// 		}
	// 	}

	// 	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	// 	defer cancel()

	// 	data, err := ds.ReadMany(page, 10, ctx)
	// 	if err != nil {
	// 		logger.Println("Could not get data:", err, data)
	// 		http.Error(w, "Internal Server error.", http.StatusInternalServerError)
	// 		return
	// 	}
	// 	if len(data) == 0 {
	// 		w.WriteHeader(http.StatusNotFound)
	// 		w.Write([]byte(`{"error": "Resource not found."}`))
	// 		return
	// 	}

	// 	// * Return the data to the user as JSON with a 200 OK status code
	// 	w.WriteHeader(http.StatusOK)
	// 	if err := json.NewEncoder(w).Encode(data); err != nil {
	// 		logger.Println("Error encoding data:", err, data)
	// 		http.Error(w, "Internal Server error.", http.StatusInternalServerError)
	// 		return
	// 	}

}
