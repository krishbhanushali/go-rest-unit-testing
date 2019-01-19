package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type entry struct {
	ID           int
	FirstName    string
	LastName     string
	EmailAddress string
	PhoneNumber  string
}

type exception struct {
	message string
	success bool
}

func main() {
	var router *mux.Router
	router = mux.NewRouter().StrictSlash(true)
	apiRouter := router.PathPrefix("/api").Subrouter()                        // /api will give access to all the API endpoints
	apiRouter.PathPrefix("/entries").HandlerFunc(getEntries).Methods("GET")   // /api/entries returns listing all the entries
	apiRouter.PathPrefix("/entry").HandlerFunc(getEntryByID).Methods("GET")   // GET /api/entry?id=1 returns the entry with id 1.
	apiRouter.PathPrefix("/entry").HandlerFunc(createEntry).Methods("POST")   // POST /api/entry creates an entry
	apiRouter.PathPrefix("/entry").HandlerFunc(updateEntry).Methods("PUT")    // PUT /api/entry updates an entry
	apiRouter.PathPrefix("/entry").HandlerFunc(deleteEntry).Methods("DELETE") // DELETE /api/entry deletes an entry

	fmt.Println("Listening on port :12345")
	http.ListenAndServe(":12345", router)

}

// Get All Entries
// URL : /entries
// Parameters: none
// Method: GET
// Output: JSON Encoded Entries object if found else JSON Encoded Exception.
func getEntries(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/online_address_book?charset=utf8&parseTime=True&loc=Local")
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

func getEntryByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Entry By ID")
}

func createEntry(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create entry")
}

func updateEntry(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update entry")
}

func deleteEntry(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete entry")
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
