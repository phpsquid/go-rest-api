package user

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
)

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{userID}", GetAUser)
	router.Post("/{userID}", CreateUser)
	router.Get("/", GetAllUsers)
	return router
}

func GetAUser(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(chi.URLParam(r, "userID"))

	user := User{
		ID: userID,
		Username: "Billy",
	}
	render.JSON(w, r, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "Created user",
	}
	render.JSON(w, r, response)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{
			ID: 1,
			Username: "Billy",
		},
		{
			ID: 2,
			Username: "Bob",
		},
	}
	render.JSON(w, r, users)
}