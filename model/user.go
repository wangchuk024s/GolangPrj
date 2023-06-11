package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go/dataStore/postgres"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) CreateUser(a []byte) error {
	const queryCreateUser = "INSERT INTO users (email,username,password) VALUES($1,$2,$3);"
	_, err := postgres.Db.Exec(queryCreateUser, u.Email, u.Username, a)
	fmt.Print(err)
	return err
}

func (u *User) Check() ([]byte, error) {
	var dbHash []byte
	const queryCheck = "SELECT * FROM users WHERE email = $1;"
	err := postgres.Db.QueryRow(queryCheck, u.Email).Scan(&u.Email, &u.Username, &dbHash)
	return dbHash, err
}

// type Post struct {
// 	id       int    `json:"id`
// 	Title    string `json:"title"`
// 	ShortURL string `json:"shortURL"`
// 	Details  string `json:"details"`
// }

// FormData represents the structure of the form data
type FormData struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	ShortURL string `json:"shortUrl"`
	Details  string `json:"details"`
}

var db *sql.DB

// Initialize the database connection
func init() {
	connStr := "postgres://username:password@localhost/dbname?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

// GetFormData returns a single form data by ID
func GetFormData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var formData FormData

	err := db.QueryRow("SELECT * FROM form_data WHERE id=$1", id).Scan(&formData.ID, &formData.Title, &formData.ShortURL, &formData.Details)
	if err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(formData)
}

// GetAllFormData returns all form data
func GetAllFormData(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM form_data")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	formDataList := []FormData{}

	for rows.Next() {
		var formData FormData
		err := rows.Scan(&formData.ID, &formData.Title, &formData.ShortURL, &formData.Details)
		if err != nil {
			log.Fatal(err)
		}
		formDataList = append(formDataList, formData)
	}

	json.NewEncoder(w).Encode(formDataList)
}

// CreateFormData adds a new form data
func CreateFormData(w http.ResponseWriter, r *http.Request) {
	var formData FormData
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		http.Error(w, "Couldn't decode the request", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO form_data (title, short_url, details) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		log.Fatal(err)
	}

	err = stmt.QueryRow(formData.Title, formData.ShortURL, formData.Details).Scan(&formData.ID)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(formData)
}

// UpdateFormData updates an existing form data
func UpdateFormData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var formData FormData
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		http.Error(w, "Couldn't decode the request", http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("UPDATE form_data SET title=$1, short_url=$2, details=$3 WHERE id=$4")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(formData.Title, formData.ShortURL, formData.Details, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Form data with ID %s has been updated", id)
}

// DeleteFormData deletes a form data by ID
func DeleteFormData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	stmt, err := db.Prepare("DELETE FROM form_data WHERE id=$1")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Form data with ID %s has been deleted", id)
}
