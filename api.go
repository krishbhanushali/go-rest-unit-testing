package main

import (
	"bytes"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var host = "http://localhost"
var port = "12345"
var connectionString = "root:1234@tcp(127.0.0.1:3306)/online_address_book?charset=utf8&parseTime=True&loc=Local"

func main() {

	var router *mux.Router
	router = mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api").Subrouter()                                               // /api will give access to all the API endpoints
	apiRouter.PathPrefix("/entries").HandlerFunc(GetEntries).Methods("GET")                          // /api/entries returns listing all the entries
	apiRouter.PathPrefix("/entry").HandlerFunc(GetEntryByID).Methods("GET")                          // GET /api/entry?id=1 returns the entry with id 1.
	apiRouter.PathPrefix("/entry").HandlerFunc(CreateEntry).Methods("POST")                          // POST /api/entry creates an entry
	apiRouter.PathPrefix("/entry").HandlerFunc(UpdateEntry).Methods("PUT")                           // PUT /api/entry updates an entry
	apiRouter.PathPrefix("/entry").HandlerFunc(DeleteEntry).Methods("DELETE")                        // DELETE /api/entry deletes an entry
	apiRouter.PathPrefix("/upload-entries-CSV").HandlerFunc(UploadEntriesThroughCSV).Methods("POST") // POST /api/upload-entries-CSV imports CSV into the database
	apiRouter.PathPrefix("/download-entries-CSV").HandlerFunc(DownloadEntriesToCSV).Methods("GET")   //GET /api/download-entries-CSV exports CSV from the database
	fmt.Println("Listening on port :12345")
	http.ListenAndServe(":"+port, router)

}

// GetEntries : Get All Entries
// URL : /entries
// Parameters: none
// Method: GET
// Output: JSON Encoded Entries object if found else JSON Encoded Exception.
func GetEntries(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}
	var entries []entry
	rows, err := db.Query("SELECT * from address_book;")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong.")
		return
	}
	defer rows.Close()
	for rows.Next() {
		var eachEntry entry
		var id int
		var firstName sql.NullString
		var lastName sql.NullString
		var emailAddress sql.NullString
		var phoneNumber sql.NullString

		err = rows.Scan(&id, &firstName, &lastName, &emailAddress, &phoneNumber)
		eachEntry.ID = id
		eachEntry.FirstName = firstName.String
		eachEntry.LastName = lastName.String
		eachEntry.EmailAddress = emailAddress.String
		eachEntry.PhoneNumber = phoneNumber.String
		entries = append(entries, eachEntry)
	}
	respondWithJSON(w, http.StatusOK, entries)
}

// GetEntryByID - Get Entry By ID
// URL : /entries?id=1
// Parameters: int id
// Method: GET
// Output: JSON Encoded Address Book Entry object if found else JSON Encoded Exception.
func GetEntryByID(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}
	id := r.URL.Query().Get("id")
	var firstName sql.NullString
	var lastName sql.NullString
	var emailAddress sql.NullString
	var phoneNumber sql.NullString
	err = db.QueryRow("SELECT first_name, last_name, email_address, phone_number from address_book where id=?", id).Scan(&firstName, &lastName, &emailAddress, &phoneNumber)
	switch {
	case err == sql.ErrNoRows:
		respondWithError(w, http.StatusBadRequest, "No entry found with the id="+id)
		return
	case err != nil:
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	default:
		var eachEntry entry
		eachEntry.ID, _ = strconv.Atoi(id)
		eachEntry.FirstName = firstName.String
		eachEntry.LastName = lastName.String
		eachEntry.EmailAddress = emailAddress.String
		eachEntry.PhoneNumber = phoneNumber.String
		respondWithJSON(w, http.StatusOK, eachEntry)
	}

}

// CreateEntry - Create Entry
// URL : /entry
// Method: POST
// Body:
/*
 * {
 *	"first_name":"John",
 *	"last_name":"Doe",
 *	"email_address":"john.doe@gmail.com",
 *	"phone_number":"1234567890",
 }
*/
// Output: JSON Encoded Address Book Entry object if created else JSON Encoded Exception.
func CreateEntry(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var entry entry
	err = decoder.Decode(&entry)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	}
	statement, err := db.Prepare("insert into address_book (first_name, last_name, email_address, phone_number) values(?,?,?,?)")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	}
	defer statement.Close()
	res, err := statement.Exec(entry.FirstName, entry.LastName, entry.EmailAddress, entry.PhoneNumber)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "There was problem entering the entry.")
		return
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		id, _ := res.LastInsertId()
		entry.ID = int(id)
		respondWithJSON(w, http.StatusOK, entry)
	}

}

// UpdateEntry - Update Entry
// URL : /entry
// Method: PUT
// Body:
/*
 * {
 *	"id":1,
 *	"first_name":"Krish",
 *	"last_name":"Bhanushali",
 *	"email_address":"krishsb2405@gmail.com",
 *	"phone_number":"7798775575",
 }
*/
// Output: JSON Encoded Address Book Entry object if updated else JSON Encoded Exception.
func UpdateEntry(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}
	decoder := json.NewDecoder(r.Body)
	var entry entry
	err = decoder.Decode(&entry)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	}

	statement, err := db.Prepare("update address_book set first_name=?, last_name=?, email_address=?, phone_number=? where id=?")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	}
	defer statement.Close()
	res, err := statement.Exec(entry.FirstName, entry.LastName, entry.EmailAddress, entry.PhoneNumber, entry.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "There was problem entering the entry.")
		return
	}
	if rowsAffected, _ := res.RowsAffected(); rowsAffected == 1 {
		respondWithJSON(w, http.StatusOK, entry)
	}
}

// DeleteEntry -  Delete Entry By ID
// URL : /entries?id=1
// Parameters: int id
// Method: DELETE
// Output: JSON Encoded Address Book Entry object if found & deleted else JSON Encoded Exception.
func DeleteEntry(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", connectionString)
	defer db.Close()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not connect to the database")
		return
	}

	id := r.URL.Query().Get("id")
	var firstName sql.NullString
	var lastName sql.NullString
	var emailAddress sql.NullString
	var phoneNumber sql.NullString
	err = db.QueryRow("SELECT first_name, last_name, email_address, phone_number from address_book where id=?", id).Scan(&firstName, &lastName, &emailAddress, &phoneNumber)
	switch {
	case err == sql.ErrNoRows:
		respondWithError(w, http.StatusBadRequest, "No entry found with the id="+id)
		return
	case err != nil:
		respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
		return
	default:

		res, err := db.Exec("DELETE from address_book where id=?", id)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
			return
		}
		count, err := res.RowsAffected()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Some problem occurred.")
			return
		}
		if count == 1 {
			var eachEntry entry
			eachEntry.ID, _ = strconv.Atoi(id)
			eachEntry.FirstName = firstName.String
			eachEntry.LastName = lastName.String
			eachEntry.EmailAddress = emailAddress.String
			eachEntry.PhoneNumber = phoneNumber.String

			respondWithJSON(w, http.StatusOK, eachEntry)
			return
		}

	}
}

//UploadEntriesThroughCSV - Reads CSV, Parses the CSV and creates all the entries in the database
func UploadEntriesThroughCSV(w http.ResponseWriter, r *http.Request) {
	//var buf bytes.Buffer
	file, _, err := r.FormFile("csvFile")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong while opening the CSV.")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 5
	csvData, err := reader.ReadAll()
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong while parsing the CSV.")
		return
	}

	var entry entry

	for _, eachEntry := range csvData {
		if eachEntry[1] != "first_name" {
			entry.FirstName = eachEntry[1]
		}
		if eachEntry[2] != "last_name" {
			entry.LastName = eachEntry[2]
		}
		if eachEntry[3] != "email_address" {
			entry.EmailAddress = eachEntry[3]
		}
		if eachEntry[4] != "phone_number" {
			entry.PhoneNumber = eachEntry[4]
		}
		if entry.FirstName != "" && entry.LastName != "" && entry.EmailAddress != "" && entry.PhoneNumber != "" {
			jsonString, err := json.Marshal(entry)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, "Something went wrong while parsing the CSV.")
				return
			}
			req, err := http.NewRequest("POST", host+":"+port+"/api/entry", bytes.NewBuffer(jsonString))
			if err != nil {
				respondWithError(w, http.StatusBadRequest, "Something went wrong while requesting to the Creation endpoint.")
				return
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				respondWithError(w, http.StatusBadRequest, "Something went wrong while requesting to the Creation endpoint.")
				return
			}
			defer resp.Body.Close()
			if resp.Status == strconv.Itoa(http.StatusBadRequest) {
				respondWithError(w, http.StatusBadRequest, "Something went wrong while inserting.")
				return
			} else if resp.Status == strconv.Itoa(http.StatusInternalServerError) {
				respondWithError(w, http.StatusInternalServerError, "Something went wrong while inserting.")
				return
			}
		}
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"success": "Upload successful"})
}

//DownloadEntriesToCSV - GetAllEntries, creates a CSV and downloads the CSV.
func DownloadEntriesToCSV(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get(host + ":" + port + "/api/entries")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Somehow host could not be reached.")
		return
	}
	data, _ := ioutil.ReadAll(response.Body)
	var entries []entry
	err = json.Unmarshal(data, &entries)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to unmarshal data.")
		return
	}
	b := &bytes.Buffer{}
	t := time.Now().Unix()
	fileName := "address-book-" + strconv.FormatInt(t, 10) + ".csv"
	writer := csv.NewWriter(b)
	heading := []string{"id", "first_name", "last_name", "email_address", "phone_number"}
	writer.Write(heading)
	for _, eachEntry := range entries {
		var record []string
		record = append(record, strconv.Itoa(eachEntry.ID))
		record = append(record, eachEntry.FirstName)
		record = append(record, eachEntry.LastName)
		record = append(record, eachEntry.EmailAddress)
		record = append(record, eachEntry.PhoneNumber)
		writer.Write(record)
	}
	writer.Flush()
	w.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv
	w.Header().Set("Content-Disposition", "attachment;filename="+fileName)
	w.WriteHeader(http.StatusOK)
	w.Write(b.Bytes())
	return
}

// RespondWithError is called on an error to return info regarding error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

// Called for responses to encode and send json data
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//encode payload to json
	response, _ := json.Marshal(payload)

	// set headers and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
