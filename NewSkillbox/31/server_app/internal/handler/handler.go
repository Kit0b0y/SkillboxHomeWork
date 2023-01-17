package handler

import (
	"github.com/kit0b0y/skillboxHomeWork/NewSkillbox/30_new/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/create", h.CreateUser)
	router.Post("/make_friends", h.MakeFriends)
	router.Delete("/user", h.DeleteUser)
	router.Get("/friends/{user_id}", h.GetFriends)
	router.Put("/{user_id}", h.UpdateAge)

	return router
}
