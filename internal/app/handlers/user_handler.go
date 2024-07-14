package handlers

import (
	"context"
	"encoding/json"
	"github.com/alkosmas92/platform-go-challenge/internal/app/models"
	"github.com/alkosmas92/platform-go-challenge/internal/app/services"
	"github.com/alkosmas92/platform-go-challenge/internal/app/utils"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Service services.UserService
	Logger  *logrus.Logger
}

func NewUserHandler(service services.UserService, logger *logrus.Logger) *UserHandler {
	return &UserHandler{Service: service, Logger: logger}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.Logger.Error("failed to decode request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		h.Logger.Error("failed to hash password: ", err)
		http.Error(w, "failed to hash password", http.StatusInternalServerError)
		return
	}

	newUser := models.NewUser(user.Username, string(hashedPassword), user.FirstName, user.LastName)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	err = h.Service.RegisterUser(ctx, newUser)
	if err != nil {
		h.Logger.Error("failed to register user: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Logger.Info("user registered: ", user.Username)
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		h.Logger.Error("failed to decode request body: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	user, err := h.Service.AuthenticateUser(ctx, credentials.Username, credentials.Password)
	if err != nil {
		h.Logger.Error("failed to authenticate user: ", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		h.Logger.Error("invalid password: ", err)
		http.Error(w, "invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.UserID, user.Username)
	if err != nil {
		h.Logger.Error("failed to generate JWT: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.Logger.Info("user logged in: ", user.Username)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
