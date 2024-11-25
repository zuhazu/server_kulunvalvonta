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

func TestPostInvalidRequestBody(t *testing.T) {

	req, err := http.NewRequest("POST", "/api/person", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Body = io.NopCloser(strings.NewReader(`Plain text, not JSON`))
	rr := httptest.NewRecorder()
	person.PostPersonHandler(rr, req, log.Default(), &service.MockPersonServiceSuccessful{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error": "Invalid request data. Please check your input."}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPostErrorCreatingData(t *testing.T) {

	req, err := http.NewRequest("POST", "/api/person", nil)
	if err != nil {
		t.Fatal(err)
	}

	dataJSON, _ := json.Marshal(models.Person{
		ID:         1,
		PersonID:   "121",
		TagID:      "221",
		PersonName: "Pekka Puupää",
		RoomID:     "036",
	})

	req.Body = io.NopCloser(strings.NewReader(string(dataJSON)))
	rr := httptest.NewRecorder()

	person.PostPersonHandler(rr, req, log.Default(), &service.MockPersonServiceError{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error": "Error creating data."}` // * This message is passed from the MockDataServiceError
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPostSuccessful(t *testing.T) {

	req, err := http.NewRequest("POST", "/api/person", nil)
	if err != nil {
		t.Fatal(err)
	}

	dataJSON, _ := json.Marshal(models.Person{
		ID:         1,
		PersonID:   "121",
		TagID:      "221",
		PersonName: "Pekka Puupää",
		RoomID:     "036",
	})

	// * Create new reader with the JSON payload
	req.Body = io.NopCloser(strings.NewReader(string(dataJSON)))

	rr := httptest.NewRecorder()

	// * Call the handler
	person.PostPersonHandler(rr, req, log.Default(), &service.MockPersonServiceSuccessful{})

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// * Check the response body
	expected := `{"id":1,"person_id":"121","tag_id":"221","person_name":"Pekka Puupää","room_id":"036"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
