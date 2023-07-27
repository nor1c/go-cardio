package authcontroller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nor1c/GoJWTMux/configs"
	"github.com/nor1c/GoJWTMux/helpers"
	"github.com/nor1c/GoJWTMux/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		log.Fatal("Failed to decode json")
	}
	defer r.Body.Close()

	var user models.User
	if err := models.DB.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		response := map[string]string{"message": "User not found!"}
		helpers.ResponseJSON(w, http.StatusNotFound, response)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Username or password invalid!"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)
	claims := &configs.JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(configs.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": "Something went wrong when trying to signing your auth token!", "err": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Path:     "/",
		Name:     "token",
		Value:    tokenString,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Authenticated!", "token": tokenString}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": "Failed to decode json!"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	// hash user password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashedPassword)

	// insert
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": "Failed to create a new account!"}
		helpers.ResponseJSON(w, http.StatusExpectationFailed, response)
		return
	}

	response := map[string]string{"message": "Account created!"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Path:   "/",
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	})

	response := map[string]string{"message": "Session cleared!"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}
