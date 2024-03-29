package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	user_app "github.com/kit0b0y/skillboxHomeWork/NewSkillbox/30_new"
	"github.com/rs/zerolog/log"
)

// CreateUser handles the request. Creates user.
// Reads JSON. If successful, returns id and username,
// status 201
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user user_app.RequestCreate
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("POST: New user %v", string(content)))

	err = json.Unmarshal(content, &user)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	message, err := h.services.User.CreateUser(user)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	message = fmt.Sprintf("%v %v", message, user.Name)
	newMessageResponse(w, http.StatusCreated, message)
}

// MakeFriends handles the request. Makes two users friends.
// Reads JSON. If successful, returns a message that users have become friends,
// status 200
func (h *Handler) MakeFriends(w http.ResponseWriter, r *http.Request) {
	var friends user_app.RequestMakeFriend
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("POST: Make friends %v", string(content)))

	err = json.Unmarshal(content, &friends)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	message, err := h.services.User.MakeFriends(friends.SourceID, friends.TargetID)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newMessageResponse(w, http.StatusOK, message)
}

// MakeFriends handles the request. Deletes a user.
// Reads JSON. If successful, it returns a message that the user has been deleted,
// status 200
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user user_app.RequestDeleteUser
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info().Msg(fmt.Sprintf("DELETE: %v", string(content)))

	err = json.Unmarshal(content, &user)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	message, err := h.services.User.DeleteUser(user.TargetID)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newMessageResponse(w, http.StatusOK, message)
}

// GetFriends handles the request. Collects a list of friends.
// Reads URL. If successful, it returns the list of the user's friends,
// status 200
func (h *Handler) GetFriends(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user_id")

	log.Info().Msg(fmt.Sprintf("GET: friends %v", string(id)))

	friends, err := h.services.User.GetFriends(id)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	message := strings.Join(friends, ";")
	log.Info().Msg("Sending: " + message)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

// UpdateAge handles the request. Updates the user's age.
// Reads URL and JSON. If successful, it returns a message that the user's age has been updated,
// status 200
func (h *Handler) UpdateAge(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user_id")

	log.Info().Msg(fmt.Sprintf("PUT: update age id %v", string(id)))

	var age user_app.RequestAge
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(content, &age)
	if err != nil {
		newMessageResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	message, err := h.services.User.UpdateAge(id, age.NewAge)
	if err != nil {
		newMessageResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	newMessageResponse(w, http.StatusOK, message)
}
