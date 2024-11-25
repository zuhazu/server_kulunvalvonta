package person_test

import (
	"encoding/json"
	"goapi/internal/api/handlers/person"
	service "goapi/internal/api/service/person"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetByIDInvalidID(t *testing.T) {
	mockPersonService := &service.MockPersonServiceSuccessful{}
	req, err := http.NewRequest("GET", "/api/person/invalid", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("id", "invalid") // * Required for routing *
	rr := httptest.NewRecorder()

	person.GetPersonByIDHandler(rr, req, log.Default(), mockPersonService)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	if strings.TrimSpace(rr.Body.String()) != `{"error": "Missconfigured ID."}` {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), `{"error": "Missconfigured ID."}`)
	}
}

func TestGetByIdInternalError(t *testing.T) {
	mockDataService := &service.MockPersonServiceError{}
	req, err := http.NewRequest("GET", "/api/person/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("id", "1") // * Required for routing *

	rr := httptest.NewRecorder()

	person.GetPersonByIDHandler(rr, req, log.Default(), mockDataService)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	if strings.TrimSpace(rr.Body.String()) != `Internal server error.` {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), `Internal server error.`)
	}
}

func TestGetByIdNotFound(t *testing.T) {
	mockPersonService := &service.MockPersonServiceNotFound{}
	req, err := http.NewRequest("GET", "/api/person/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("id", "1") // * Required for routing *

	rr := httptest.NewRecorder()

	person.GetPersonByIDHandler(rr, req, log.Default(), mockPersonService)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	if strings.TrimSpace(rr.Body.String()) != `{"error": "Resource not found."}` {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), `{"error": "Resource not found."}`)
	}
}

func TestGetByIdSuccessful(t *testing.T) {
	mockPersonService := &service.MockPersonServiceSuccessful{}
	req, err := http.NewRequest("GET", "/api/person/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.SetPathValue("id", "1") // * Required for routing *

	rr := httptest.NewRecorder()

	person.GetPersonByIDHandler(rr, req, log.Default(), mockPersonService)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	person, _ := mockPersonService.ReadOnePerson(1, nil)
	expected, _ := json.Marshal(person)

	if strings.TrimSpace(rr.Body.String()) != string(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), string(expected))
	}
}
