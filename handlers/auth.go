package handlers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/oshorefueled/taxexpress/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type authHandler struct {
	handlerHelper
}

func authRoute (r chi.Router) {
	h := authHandler{}
	r.Get("/register", h.authView)
	r.Post("/login", h.login)
	r.Post("/register", h.register)
}

func (h authHandler) authView (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to register auth"))
}

func (h authHandler) login (w http.ResponseWriter, r *http.Request) {
	var admin models.Admin
	_ = json.NewDecoder(r.Body).Decode(&admin)
	plainPassword := []byte(admin.Password)
	err := admin.GetAdminUser()
	if err != nil {
		h.RespondWithJSON(w, 400, err.Error())
	}
	if comparePasswords(admin.Password, plainPassword) {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": map[string]interface{}{
				"username": admin.Username,
				"token": admin.Token,
			},
		})
	} else {
		h.RespondWithJSON(w, 400, "wrong username or password")
	}
}

func (h authHandler) register (w http.ResponseWriter, r *http.Request) {
	var admin models.Admin
	_ = json.NewDecoder(r.Body).Decode(&admin)
	admin.Token = generateToken(admin.Email)
	admin.Password = hashAndSalt([]byte(admin.Password))
	_, _, err := admin.CreateAdminUser()
	if err != nil {
		h.RespondWithError(w, 400, err.Error())
	} else {
		h.RespondWithJSON(w, 200, map[string]interface{}{
			"status": "success",
			"data": map[string]string{
				"message": "admin created successfully",
			},
		})
	}
}

func changePassword () {
	// todo
}

func updateProfile () {
	// todo
}

func generateToken (email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

func tokenFromHash (hash string) string {
	return base64.StdEncoding.EncodeToString([]byte(hash))
}

func hashPassword () {
	// todo
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}