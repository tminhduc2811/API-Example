package controller

import (
	"encoding/json"
	"exampleAPI/models"
	"fmt"
	"net/http"
)

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