package person_test

import (
	"encoding/json"
	"goapi/internal/api/handlers/person"
	"goapi/internal/api/repository/models"
	service "goapi/internal/api/service/person"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdatePersonInvalidRequestBody(t *testing.T) {

	req, err := http.NewRequest("PUT", "/api/personX", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Body = io.NopCloser(strings.NewReader(`Plain text, not JSON`))
	rr := httptest.NewRecorder()
	person.PutPersonHandler(rr, req, log.Default(), &service.MockPersonServiceSuccessful{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error": "Invalid request data. Please check your input."}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdatePersonErrorCreatingData(t *testing.T) {

	req, err := http.NewRequest("PUT", "/api/person", nil)
	if err != nil {
		t.Fatal(err)
	}

	dataJSON, _ := json.Marshal(models.Person{
		ID:         8,
		PersonID:   "1221",
		TagID:      "afddfa",
		PersonName: "jumppa",
		RoomID:     "jenninhuone",
	})

	req.Body = io.NopCloser(strings.NewReader(string(dataJSON)))
	rr := httptest.NewRecorder()

	person.PutPersonHandler(rr, req, log.Default(), &service.MockPersonServiceError{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error": "Error updating data."}` // * This message is passed from the MockDataServiceError
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestUpdatePersonTagIdSuccessful(t *testing.T) {

	req, err := http.NewRequest("PUT", "/api/tag", nil)
	if err != nil {
		t.Fatal(err)
	}

	dataJSON, _ := json.Marshal(models.Person{
		ID:         8,
		PersonID:   "1221",
		TagID:      "afddfa",
		PersonName: "jumppa",
		RoomID:     "jenninhuone",
	})

	// * Create new reader with the JSON payload
	req.Body = io.NopCloser(strings.NewReader(string(dataJSON)))

	rr := httptest.NewRecorder()

	// * Call the handler
	person.PutPersonHandler(rr, req, log.Default(), &service.MockPersonServiceSuccessful{})

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// * Check the response body
	expected := `{"id":8,"person_id":"1221","tag_id":"afddfa","person_name":"jumppa","room_id":"jenninhuone"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
