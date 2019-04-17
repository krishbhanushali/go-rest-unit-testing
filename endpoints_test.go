package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetEntries(t *testing.T) {
	req, err := http.NewRequest("GET", "/entries", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEntries)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
<<<<<<< HEAD
	expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"},{"id":2,"first_name":"Rick","last_name":"Parker","email_address":"rick.parker@gmail.com","phone_number":"1234567890"},{"id":3,"first_name":"Kelly","last_name":"Franco","email_address":"kelly.franco@gmail.com","phone_number":"1112223333"}]`
=======
	expected := `[{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"6264648201"},{"id":2,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"},{"id":6,"first_name":"Juheeshwari","last_name":"Rajmohan","email_address":"lrrajmohan16@gmail.com","phone_number":"1111111111"}]`
>>>>>>> 888598d616a32df55200163f984bb0e6130786ed
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetEntryByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/entry", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEntryByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":1,"first_name":"Krish","last_name":"Bhanushali","email_address":"krishsb2405@gmail.com","phone_number":"0987654321"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetEntryByIDNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/entry", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "123")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetEntryByID)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestCreateEntry(t *testing.T) {

<<<<<<< HEAD
	var jsonStr = []byte(`{"id":11,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`)
=======
	var jsonStr = []byte(`{"id":4,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`)
>>>>>>> 888598d616a32df55200163f984bb0e6130786ed

	req, err := http.NewRequest("POST", "/entry", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateEntry)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
<<<<<<< HEAD
	expected := `{"id":11,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
=======
	expected := `{"id":4,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
>>>>>>> 888598d616a32df55200163f984bb0e6130786ed
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestEditEntry(t *testing.T) {

<<<<<<< HEAD
	var jsonStr = []byte(`{"id":11,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`)
=======
	var jsonStr = []byte(`{"id":4,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`)
>>>>>>> 888598d616a32df55200163f984bb0e6130786ed

	req, err := http.NewRequest("PUT", "/entry", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateEntry)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
<<<<<<< HEAD
	expected := `{"id":11,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
=======
	expected := `{"id":4,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
>>>>>>> 888598d616a32df55200163f984bb0e6130786ed
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestDeleteEntry(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/entry", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
<<<<<<< HEAD
	q.Add("id", "11")
=======
	q.Add("id", "4")
>>>>>>> 888598d616a32df55200163f984bb0e6130786ed
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteEntry)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
<<<<<<< HEAD
	expected := `{"id":11,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
=======
	expected := `{"id":4,"first_name":"xyz change","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
>>>>>>> 888598d616a32df55200163f984bb0e6130786ed
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
