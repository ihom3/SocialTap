//To initiate this file:
//If there is no go module:
//		On the termial type: go mod init example.com/SocialTap
//		Next: go mod tidy
//Else:
//		go get -u github.com/gorilla/mux
//		go get -u gorm.io/gorm
//		go get -u gorm.io/driver/mysql
//To build and run:
//		go build
//		.\SocialTap.exe

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// method to initialize the router
func initRouter() {
	r := mux.NewRouter()
	//inside the handle function we need to provide the route information = path
	//then we need to provide the function we need to call when we see /users
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")    // to get a user
	r.HandleFunc("/users", CreateUser).Methods("POST")     // post because we are putting information in
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT") // put because we are changing information
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	//ListenAndServe(address, handler http.Handler) listens on the TCP network address
	//then calls serve with handler to handle requests on incoming connections.
	log.Fatal(http.ListenAndServe(":9000", r)) // port: 9000, router: r
}

func main() {
	//calling the router
	InitialMigration()
	initRouter()

}
