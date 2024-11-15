package person

import (
	"context"
	"encoding/json"
	service "goapi/internal/api/service/person"
	"log"
	"net/http"
	"time"
)

func PostTagHandler(w http.ResponseWriter, r *http.Request, logger *log.Logger, ds service.PersonService) {

	var input struct {
		TagID     string `json:"tag_id"`
		NewRoomID string `json:"new_room_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request data. Please check your input."}`))
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	message, err := ds.UpdateRoomIDByTagID(input.TagID, input.NewRoomID, ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Epäonnistui päivittäminen: ` + err.Error() + `"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "` + message + `"}`))
}
