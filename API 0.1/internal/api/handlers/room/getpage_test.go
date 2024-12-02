package room_test

import (
	"goapi/internal/api/handlers/room"
	service "goapi/internal/api/service/room"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetByIDInvalidID(t *testing.T) {
	mockRoomService := &service.MockRoomServiceNotFound{}
	req, err := http.NewRequest("GET", "/main?room=invalid", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("room", "invalid") // * Required for routing *
	rr := httptest.NewRecorder()

	room.GetPageHandler(rr, req, log.Default(), mockRoomService)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	if !strings.HasPrefix(strings.TrimSpace(rr.Body.String()), "Page not found") {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), "Page not found")
	}
}
