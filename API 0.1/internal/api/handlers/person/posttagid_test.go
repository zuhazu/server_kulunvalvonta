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

func TestPostTagIdInvalidRequestBody(t *testing.T) {

	req, err := http.NewRequest("PUT", "/api/tagx", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Body = io.NopCloser(strings.NewReader(`Plain text, not JSON`))
	rr := httptest.NewRecorder()
	person.PostTagHandler(rr, req, log.Default(), &service.MockPersonServiceSuccessful{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `{"error": "Invalid request data. Please check your input."}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPostTagIdErrorCreatingData(t *testing.T) {

	req, err := http.NewRequest("PUT", "/api/tag", nil)
	if err != nil {
		t.Fatal(err)
	}

	dataJSON, _ := json.Marshal(models.Person{
		ID:     30,
		TagID:  "14286886888",
		RoomID: "78",
	})

	req.Body = io.NopCloser(strings.NewReader(string(dataJSON)))
	rr := httptest.NewRecorder()

	person.PostTagHandler(rr, req, log.Default(), &service.MockPersonServiceError{})

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `Access denied` // * This message is passed from the MockDataServiceError
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestPostTagIdSuccessful(t *testing.T) {

	req, err := http.NewRequest("PUT", "/api/tag", nil)
	if err != nil {
		t.Fatal(err)
	}

	dataJSON, _ := json.Marshal(models.Person{
		ID:     1,
		TagID:  "142",
		RoomID: "8294",
	})

	// * Create new reader with the JSON payload
	req.Body = io.NopCloser(strings.NewReader(string(dataJSON)))

	rr := httptest.NewRecorder()

	// * Call the handler
	person.PostTagHandler(rr, req, log.Default(), &service.MockPersonServiceSuccessful{})

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// * Check the response body
	expected := `success`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
