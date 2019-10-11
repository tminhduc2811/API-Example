package main

import (
	"exampleAPI/routers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main()  {
	router := mux.NewRouter().StrictSlash(false)
	router = routers.SetupRoutes(router)
	fmt.Println("Listening at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
