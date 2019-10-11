package routers

import (
	"exampleAPI/controller"
	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) *mux.Router {
	userRouter := mux.NewRouter()
	userRouter.HandleFunc("/hello", controller.HelloUser).Methods("GET")
	userRouter.HandleFunc("/newuser", controller.ShowHang).Methods("POST")
	return userRouter
}