package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mikey247/go-bms/pkg/routes"
)

func main()  {
	router:= mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/",router)
	fmt.Println("⚡running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}