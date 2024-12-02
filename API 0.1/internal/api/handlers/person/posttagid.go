package person

import (
	"context"
	"encoding/json"
	service "goapi/internal/api/service/person"
	"log"
	"net/http"
	"time"
)

// Tämä handleri huolehtii RFID-tagin käytöstä.
// Jos henkilön roomId = -1, asetetaan se newRoomId:n mukaiseksi
// Jos henkilöllä on muu kuin -1 roomId, asetetaan se -1:ksi
func PostTagHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.PersonService) {

	var input struct {
		TagID     string `json:"tag_id"`
		NewRoomID string `json:"room_id"`
	}

	// Muutetaan pyynnön body input muotoon joka määritelty yllä
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request data. Please check your input."}`))
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// Päivitetään roomId personille
	message, err := ds.UpdateRoomIDByTagID(input.TagID, input.NewRoomID, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Print(err.Error())
		w.Write([]byte(message))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
