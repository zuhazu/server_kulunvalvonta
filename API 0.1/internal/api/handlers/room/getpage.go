package room

import (
	"context"
	page_service "goapi/internal/api/service/page"
	service "goapi/internal/api/service/room"
	"log"
	"net/http"
	"time"
)

// Tämä handleri huolehtii nettisivun palauttamisesta
func GetPageHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.RoomService) {

	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	//Haetaan personit roomId:n perusteella
	data, err := ds.GetPersonsByRoomID(r.URL.Query().Get("room"), ctx)
	logger.Print(r.URL.Query().Get("room"))
	if err != nil {
		logger.Println("Could not get data:", err, data)
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	// Kirjoitetaan internetsivu oikealla datalla vastaukseen
	w.Write([]byte(page_service.GetPageModel(data)))
}
