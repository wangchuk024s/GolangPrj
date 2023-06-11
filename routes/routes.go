package routes

import (
	"go/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", controller.SignupHandler).Methods("POST")

	router.HandleFunc("/login", controller.LoginHandler).Methods("POST")

	// router.HandleFunc("/post", controller.AddPost).Methods("POST")
	// router.HandleFunc("/post/{id}", controller.GetPost).Methods("GET")
	// router.HandleFunc("/post/{id}", controller.DeletePost).Methods("DELETE")
	// router.HandleFunc("/post/{id}", controller.UpdatePost).Methods("PUT")
	// router.HandleFunc("/posts", controller.GetAllPost).Methods("GET")

	// router.HandleFunc("/posts", controller.FormHandler).Methods("Post")

	// router.HandleFunc("/formdata/{id}", controller.GetFormData).Methods("GET")
	// router.HandleFunc("/formdata", controller.GetAllFormData).Methods("GET")
	// router.HandleFunc("/formdata", controller.CreateFormData).Methods("POST")
	// router.HandleFunc("/formdata/{id}", controller.UpdateFormData).Methods("PUT")
	// router.HandleFunc("/formdata/{id}", controller.DeleteFormData).Methods("DELETE")

	fhandler := http.FileServer(http.Dir("./view"))

	router.PathPrefix("/").Handler(fhandler)

	err := http.ListenAndServe(":4343", router)
	if err != nil {
		return
	}

}
func RegisterURLRoutes(r *mux.Router) {
	r.HandleFunc("/api/urls", controller.AddURLMapping).Methods("POST")
	r.HandleFunc("/api/urls/{id}", controller.DeleteURLMapping).Methods("DELETE")
	r.HandleFunc("/api/urls/{id}", controller.UpdateURLMapping).Methods("PUT")
	r.HandleFunc("/api/urls", controller.GetURLMappings).Methods("GET")
}
