package controller

import (
	"encoding/json"
	"exampleAPI/models"
	"exampleAPI/repositories"
	"fmt"
	"gopkg.in/mgo.v2"
	"net/http"
)
type UserController struct {
	UserRepo *repositories.UserRepository
}

func (c *UserController) NewUserController() *UserController {
	dbConnect, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	userRepo := repositories.NewUserRepository(dbConnect)
	return &UserController{
		userRepo,
	}
}

func HelloUser(w http.ResponseWriter, r *http.Request)  {

	mess := models.Hello{"suck mah dick"}
	j, err := json.Marshal(mess)
	if err != nil {
		fmt.Print("error")
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func ShowHang(w http.ResponseWriter, r *http.Request)  {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("error while parsing json")
	}
	replyMess := "Hello " + user.Name + " VNCH"
	j, err2 := json.Marshal(replyMess)
	if err2 != nil {
		fmt.Println("error")
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func CreateNewUser(w http.ResponseWriter, r *http.Request)  {
	var user models.User
	var replyMess string
	var userController *UserController

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Error")
	}
	userController = userController.NewUserController()
	err = userController.UserRepo.CreateUser(&user)
	if err != nil{
		fmt.Println("Cannot create new user")
	}
	replyMess = "Create new user successfully"
	j, err := json.Marshal(replyMess)
	if err != nil {
		fmt.Println("Error")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}