package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go/model"
	"go/utils/httpResp"
	"net/http"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// func AddPost(w http.ResponseWriter, r *http.Request) {
// 	var stud model.Post
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&stud); err != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
// 		return
// 	}

// 	defer r.Body.Close()
// 	saveErr := stud.Create()
// 	if saveErr != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
// 		return
// 	}

// 	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "Post added"})
// }

// func GetPost(w http.ResponseWriter, r *http.Request) {
// 	pid := mux.Vars(r)["id"]
// 	pidint := getIntPostID(pid)
// 	p := model.Post{Post_id: pidint}

// 	err := p.GetPost()

// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			httpResp.RespondWithError(w, http.StatusNotFound, "Post not found")
// 			fmt.Println("Post not found")
// 		default:
// 			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}
// 	httpResp.RespondWithJSON(w, http.StatusOK, p)

// }

// // convert string stdID to int
// func getIntPostID(userIdParam string) int {
// 	userId, _ := strconv.ParseInt(userIdParam, 10, 0)
// 	return int(userId)
// }

// // Delete Post
// func DeletePost(w http.ResponseWriter, r *http.Request) {
// 	pid := mux.Vars(r)["id"]
// 	pidint := getIntPostID(pid)
// 	p := model.Post{Post_id: pidint}

// 	err := p.DeletePost()

// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			httpResp.RespondWithError(w, http.StatusNotFound, "Post not found")
// 			fmt.Println("Post not found")
// 		default:
// 			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}
// 	httpResp.RespondWithJSON(w, http.StatusOK, "deleted the post")
// }

// // Update Post
// func UpdatePost(w http.ResponseWriter, r *http.Request) {
// 	pid := mux.Vars(r)["id"]
// 	pidint := getIntPostID(pid)
// 	p := model.Post{Post_id: pidint}

// 	err := p.UpdatePost()

// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			httpResp.RespondWithError(w, http.StatusNotFound, "Post not found")
// 			fmt.Println("Post not found")
// 		default:
// 			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}
// 	httpResp.RespondWithJSON(w, http.StatusCreated, p)
// }

// // Get all Post
// func GetAllPost(w http.ResponseWriter, r *http.Request) {
// 	posts, err := model.GetAllPost()
// 	if err != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	httpResp.RespondWithJSON(w, http.StatusOK, posts)
// }

// Sign Up
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "couldn't decode the request")
		fmt.Println("error in decoding the request", err)
		return
	}
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		return
	}
	fmt.Println("hash:", hash)
	fmt.Println("string(hash)", string(hash))
	addErr := user.CreateUser(hash)
	if addErr != nil {
		if e, ok := addErr.(*pq.Error); ok {
			if e.Code == "23505" {
				fmt.Print("duplicate key error")
				httpResp.RespondWithError(w, http.StatusNotAcceptable, "duplicate key error")
				return
			}
		} else {
			fmt.Println("error in inserting the data")
			httpResp.RespondWithError(w, http.StatusBadRequest, "error in inserting the data")
			return
		}
	}

	fmt.Println("successful")
	httpResp.RespondWithJSON(w, http.StatusCreated, "added successfully")
}

// Login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user model.User
	var err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
		fmt.Println("error in decoding the request")
		return
	}
	dbHash, checkErr := user.Check()
	if checkErr != nil {
		switch checkErr {
		case sql.ErrNoRows:
			fmt.Print("invalid login")
			httpResp.RespondWithError(w, http.StatusNotFound, "Invalid login, try again")

		default:
			fmt.Print("error in getting the data to compare", checkErr)
			httpResp.RespondWithError(w, http.StatusBadRequest, "error with the database")

		}
		return
	}
	pwdErr := bcrypt.CompareHashAndPassword([]byte(dbHash), []byte(user.Password))
	if pwdErr != nil {
		fmt.Println("invalid with password")
		httpResp.RespondWithError(w, http.StatusNotFound, "Invalid login, try again")
		return
	}
	fmt.Println("login successful")
	httpResp.RespondWithJSON(w, http.StatusAccepted, map[string]string{"message": "successful"})
}

// // saveurl

// // FormData represents the data structure for the form
// type FormData struct {
// 	ID       string `json:"id"`
// 	Title    string `json:"title"`
// 	ShortURL string `json:"shortUrl"`
// 	Details  string `json:"details"`
// }

// // FormHandler handles the form submission
// func FormHandler(w http.ResponseWriter, r *http.Request) {
// 	var formData FormData
// 	err := json.NewDecoder(r.Body).Decode(&formData)
// 	if err != nil {
// 		http.Error(w, "Couldn't decode the request", http.StatusBadRequest)
// 		return
// 	}

// 	// Process the form data here

// 	// Example: Print the form data
// 	fmt.Println("ID:", formData.ID)
// 	fmt.Println("Title:", formData.Title)
// 	fmt.Println("Short URL:", formData.ShortURL)
// 	fmt.Println("Details:", formData.Details)

//		// Send a response
//		response := "Form submitted successfully"
//		w.Header().Set("Content-Type", "text/plain")
//		w.WriteHeader(http.StatusOK)
//		fmt.Fprint(w, response)
//	}

// FormData represents the structure of the form data
// type FormData struct {
// 	ID       int    `json:"id"`
// 	Title    string `json:"title"`
// 	ShortURL string `json:"shortUrl"`
// 	Details  string `json:"details"`
// }

// var db *sql.DB

// // Initialize the database connection
// func init() {
// 	connStr := "postgres://username:password@localhost/dbname?sslmode=disable"
// 	var err error
// 	db, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// // GetFormData returns a single form data by ID
// func GetFormData(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id := params["id"]

// 	var formData FormData

// 	err := db.QueryRow("SELECT * FROM form_data WHERE id=$1", id).Scan(&formData.ID, &formData.Title, &formData.ShortURL, &formData.Details)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			http.NotFound(w, r)
// 			return
// 		}
// 		log.Fatal(err)
// 	}

// 	json.NewEncoder(w).Encode(formData)
// }

// // GetAllFormData returns all form data
// func GetAllFormData(w http.ResponseWriter, r *http.Request) {
// 	rows, err := db.Query("SELECT * FROM form_data")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	formDataList := []FormData{}

// 	for rows.Next() {
// 		var formData FormData
// 		err := rows.Scan(&formData.ID, &formData.Title, &formData.ShortURL, &formData.Details)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		formDataList = append(formDataList, formData)
// 	}

// 	json.NewEncoder(w).Encode(formDataList)
// }

// // CreateFormData adds a new form data
// func CreateFormData(w http.ResponseWriter, r *http.Request) {
// 	var formData FormData
// 	err := json.NewDecoder(r.Body).Decode(&formData)
// 	if err != nil {
// 		http.Error(w, "Couldn't decode the request", http.StatusBadRequest)
// 		return
// 	}

// 	stmt, err := db.Prepare("INSERT INTO form_data (title, short_url, details) VALUES ($1, $2, $3) RETURNING id")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = stmt.QueryRow(formData.Title, formData.ShortURL, formData.Details).Scan(&formData.ID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	json.NewEncoder(w).Encode(formData)
// }

// // UpdateFormData updates an existing form data
// func UpdateFormData(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id := params["id"]

// 	var formData FormData
// 	err := json.NewDecoder(r.Body).Decode(&formData)
// 	if err != nil {
// 		http.Error(w, "Couldn't decode the request", http.StatusBadRequest)
// 		return
// 	}

// 	stmt, err := db.Prepare("UPDATE form_data SET title=$1, short_url=$2, details=$3 WHERE id=$4")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = stmt.Exec(formData.Title, formData.ShortURL, formData.Details, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Fprintf(w, "Form data with ID %s has been updated", id)
// }

// // DeleteFormData deletes a form data by ID
// func DeleteFormData(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id := params["id"]

// 	stmt, err := db.Prepare("DELETE FROM form_data WHERE id=$1")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = stmt.Exec(id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Fprintf(w, "Form data with ID %s has been deleted", id)
// }

// type model struct {
// 	PostID   int    `json:"postID"`
// 	Title    string `json:"title"`
// 	ShortURL string `json:"shortURL"`
// 	Details  string `json:"details"`
// }

// func AddPost(w http.ResponseWriter, r *http.Request) {
// 	var post model.Post
// 	decoder := json.NewDecoder(r.Body)
// 	if err := decoder.Decode(&post); err != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
// 		return
// 	}

// 	defer r.Body.Close()
// 	saveErr := post.Create()
// 	if saveErr != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
// 		return
// 	}

// 	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"status": "Post added"})
// }

// func GetPost(w http.ResponseWriter, r *http.Request) {
// 	postID := mux.Vars(r)["id"]

// 	postIDInt := getIntPostID(postID)
// 	p := model.Post{PostID: postIDInt}

// 	err := p.GetPost()

// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			httpResp.RespondWithError(w, http.StatusNotFound, "Post not found")
// 			fmt.Println("Post not found")
// 		default:
// 			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}
// 	httpResp.RespondWithJSON(w, http.StatusOK, p)
// }

// // convert string postID to int
// func getIntPostID(postIDParam string) int {
// 	postID, _ := strconv.ParseInt(postIDParam, 10, 0)
// 	return int(postID)
// }

// func DeletePost(w http.ResponseWriter, r *http.Request) {
// 	postID := mux.Vars(r)["id"]
// 	postIDInt := getIntPostID(postID)
// 	p := model.Post{PostID: postIDInt}

// 	err := p.DeletePost()

// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			httpResp.RespondWithError(w, http.StatusNotFound, "Post not found")
// 			fmt.Println("Post not found")
// 		default:
// 			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}
// 	httpResp.RespondWithJSON(w, http.StatusOK, "Deleted the post")
// }

// func UpdatePost(w http.ResponseWriter, r *http.Request) {
// 	postID := mux.Vars(r)["id"]
// 	postIDInt := getIntPostID(postID)
// 	p := model.Post{PostID: postIDInt}

// 	err := p.UpdatePost()

// 	if err != nil {
// 		switch err {
// 		case sql.ErrNoRows:
// 			httpResp.RespondWithError(w, http.StatusNotFound, "Post not found")
// 			fmt.Println("Post not found")
// 		default:
// 			httpResp.RespondWithError(w, http.StatusInternalServerError, err.Error())
// 		}
// 		return
// 	}
// 	httpResp.RespondWithJSON(w, http.StatusCreated, p)
// }

// func GetAllPosts(w http.ResponseWriter, r *http.Request) {
// 	posts, err := model.GetAllPosts()
// 	if err != nil {
// 		httpResp.RespondWithError(w, http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	httpResp.RespondWithJSON(w, http.StatusOK, posts)
// }
