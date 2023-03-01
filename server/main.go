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
	//HandleFunc registers the handler function for the given pattern in the server
	//inside the handle function we need to provide the route information = path
	//then we need to provide the function we need to call when we see /users
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")    // to get a user
	r.HandleFunc("/users", CreateUser).Methods("POST")     // post because we are putting information in
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT") // put because we are changing information
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/users/{id}/instagram", GetSocial).Methods("GET")
	r.HandleFunc("/users/code", AddUserSocial).Methods("POST")

	r.HandleFunc("/check-is-prime", isPrimeHandler).Methods("GET")

	//r.HandleFunc("/home", HomeHandler)
	//r.HandleFunc("/{id}", LoginHandler)
	//r.HandleFunc("/dashboard", LoginHandler)
	//r.HandleFunc("/update-profile", LoginHandler)
	//r.HandleFunc("/udate-socials", LoginHandler)
	r.HandleFunc("/update-profile-picture", UpdateProfilePicture).Methods("POST")
	//r.HandleFunc("/add-code", LoginHandler)

	//ListenAndServe(address, handler http.Handler) listens on the TCP network address
	//then calls serve with handler to handle requests on incoming connections.
	log.Fatal(http.ListenAndServe(":9000", r)) // port: 9000, router: r
	//if we use nil, instead of r, we get the default multiplexer ServerMux. An HTTP multiplexer matches
	//the URL of each incoming request with a list of registered patterns and calls the handler for the
	//pattern that most closely matches the URL.
}

func main() {
	//calling the router
	InitialMigration()
	initRouter()

}

//short variable declarations (:=) can only be used inside functions in go

//slices in golang just like arrays but without a fixed sized

//parsing templates
//include: "html/template"
//
//var tpl *template.Template ---- global variable, pointer to our template
//in main:
//func ParseFiles(filenames...string) (*Template, error)
//tpl, _ = template.ParseFiles("templates/index.html") ------ if the file is not in the main directory you need to use /<name of lower directory>/index.html, if it's higher directory use ../index.html
//	or you can use tpl, _ = tpl.ParseFiles("templates/index.html")
//	for multiple files, use the wild character "*" to parse the files:
//	tpl, _ = template.ParseGlob("templates/*.html"), it is saying, parse anything that ends with .html
//then call handler function to handle the index
//r.HandleFunc("/", indexHandler) ----- register path to a handler
//r.ListenAnd........ fire up the server
//
//to serve that parsed template to the writer:
//func indexHandler(w http.ResponseWriter, r *http.Request)
//func (t *Template) Execute(wr io.Writer, data interface{}) error
//tpl.Execute(w, nil)---nil if do not want to pass in any data
//
//parsing templates with data
//{{.}}		renders the root element
//{{.Name}}		renders the Name field in a nested element
//{{if .Done}} {{else}} {{end}}		defines an if/else statement
//{{range.List}} {{.}} {{end}}
