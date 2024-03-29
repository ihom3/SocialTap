package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


// checking the status of the get users handler
func TestGetUsersStatusCode(t *testing.T) {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	server := httptest.NewServer(http.HandlerFunc(GetUsers))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", resp.StatusCode)
	}
}

// testing the getting user by code with id = sticker_code as parameter
func TestGetUserByCode(t *testing.T) {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	//_, err := DB.Exec("INSERT INTO users (first_name, last_name) VALUES (?,?)", "apple", "Doe")
	user := User{
		FirstName:      "Ian",
		UserEmail:      "ian.n.",
		LastName:       "B",
		StickerCode:    "hello",
		BioText:        "hello world",
		ProfilePicture: "1",
		SocialList: Socials{
			Facebook: SocialType{
				SocialName: "",
				Status:     false,
				URL:        "",
			},
			Snapchat: SocialType{
				SocialName: "",
				Status:     false,
				URL:        "",
			},
			Instagram: SocialType{
				SocialName: "",
				Status:     false,
				URL:        "",
			},
		},
	}

	DB.Create(&user)

	req, err := http.NewRequest("GET", "/users/?id=hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserEnhanced)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"user_email":"ian.n.","first_name":"Ian","last_name":"B","sticker_code":"hello","bio_text":"hello world","profile_picture":"1","social_list":{"facebook":{"name":"","status":false,"url":""},"snapchat":{"name":"","status":false,"url":""},"instagram":{"name":"","status":false,"url":""}}}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}


// testing the adding a user request handler
func TestCreateUser(t *testing.T) {
	// Connect to the test database
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// Create a new user
	user := User{
		FirstName:      "John",
		LastName:       "Doe",
		UserEmail:      "johndoe@example.com",
		StickerCode:    "abcd1234",
		BioText:        "Hello, world!",
		ProfilePicture: "1",
		SocialList: Socials{
			Facebook: SocialType{
				SocialName: "Facebook",
				Status:     true,
				URL:        "/johndoe",
			},
			Instagram: SocialType{
				SocialName: "Instagram",
				Status:     true,
				URL:        "/johndoe",
			},
			Snapchat: SocialType{
				SocialName: "Snapchat",
				Status:     true,
				URL:        "/johndoe",
			},
		},
	}
	DB.Create(&user)

	// Encode the user as JSON and make a POST request to the handler
	requestBody, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Query the database to check if the user was created

	//err = DB.QueryRow("SELECT COUNT(*) FROM users WHERE user_email = ?", user.UserEmail).Scan(&count)
	result := DB.First(&user).RowsAffected // returns count of records found
	count := result

	if err != nil {
		t.Fatal(err)
	}

	if count != 1 {
		t.Errorf("Expected 1 row to be affected, but got %d", count)
	}
}

// testing the proper migration of the tables into the database
func TestInitialMigration(t *testing.T) {
	// Connect to the test database
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("	 Cannot connect to DB")
	}

	// Perform the initial database migration
	InitialMigration()

	// Check if the User table was created
	if !DB.Migrator().HasTable(&User{}) {
		t.Errorf("	User table was not created")
	}

	// Check if the UnregisteredCodes table was created
	if !DB.Migrator().HasTable(&UnregisteredCodes{}) {
		t.Errorf("	UnregisteredCodes table was not created")
	}

	// Check if the Socials table was created
	if !DB.Migrator().HasTable(&Socials{}) {
		t.Errorf("	Socials table was not created")
	}
	fmt.Println("	Everything was properly migrated.")
	fmt.Println()
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", DeleteUser)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("	Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Assert response body
	result := strings.TrimSpace(rr.Body.String())
	expected := "\"The user is deleted successfully.\""
	if result != expected {
		t.Errorf("	Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	} else {
		fmt.Println("	The user got deleted successfully.")
	}
	fmt.Println()
}

// testing the picture upload

func TestUpdateProfilePicture(t *testing.T) {
	// Create a new HTTP request with a test file
	file, err := os.Open("test-files/apple-test.jpg")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("profile_picture", filepath.Base(file.Name()))
	if err != nil {
		t.Fatal(err)
	}
	io.Copy(part, file)
	writer.Close()

	req, err := http.NewRequest("POST", "/update_profile_picture", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()

	// Call the handler function
	UpdateProfilePicture(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check that the file was uploaded correctly
	_, err = os.Stat("profile_pictures")
	if os.IsNotExist(err) {
		t.Errorf("directory profile_pictures not created")
	}

	files, err := ioutil.ReadDir("profile_pictures")
	if err != nil {
		t.Errorf("error reading directory: %v", err)
	}

	if len(files) != 1 {
		t.Errorf("expected 1 file in directory, got %v", len(files))
	}

	// Clean up temporary files and directories
	os.RemoveAll("profile_pictures")
}
func TestDeleteCode(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/unregistered/testcode5", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/unregistered/{sticker_code}", DeleteCode)

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("	Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Assert response body
	result := strings.TrimSpace(rr.Body.String())
	expected := "\"The code is deleted successfully.\""
	if result != expected {
		t.Errorf("	Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	} else {
		fmt.Println("	The code got deleted successfully.")
	}
	fmt.Println()
}

func TestAddCode(t *testing.T) {
	// create a new HTTP request with a valid sticker code
	reqBody := `{"sticker_code": "ABC123"}`
	req, err := http.NewRequest("POST", "/unregistered", strings.NewReader(reqBody))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// call the Dashboard handler function with the request and response recorder
	handler := http.HandlerFunc(AddCode)
	handler.ServeHTTP(rr, req)

	// check the HTTP method of the request
	if req.Method != http.MethodPost {
		t.Errorf("AddCode was not called with the correct HTTP method: got %v, want %v", req.Method, http.MethodPost)
	}
}

func TestDashboard(t *testing.T) {
	// create a new HTTP request with a valid sticker code
	req, err := http.NewRequest("GET", "/dashboard/testcode5", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// call the Dashboard handler function with the request and response recorder
	handler := http.HandlerFunc(Dashboard)
	handler.ServeHTTP(rr, req)

	// check the HTTP method of the request
	if req.Method != http.MethodGet {
		t.Errorf("Dashboard was not called with the correct HTTP method: got %v, want %v", req.Method, http.MethodGet)
	}
}

func TestIDRoute(t *testing.T) {
	// create a new HTTP request with a valid sticker code
	req, err := http.NewRequest("GET", "/testcode2", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// create a response recorder to capture the response
	rr := httptest.NewRecorder()

	// call the Dashboard handler function with the request and response recorder
	handler := http.HandlerFunc(IDRoute)
	handler.ServeHTTP(rr, req)

	// check the HTTP method of the request
	if req.Method != http.MethodGet {
		t.Errorf("IDRoute was not called with the correct HTTP method: got %v, want %v", req.Method, http.MethodGet)
	}
}

func TestGetUserNameByCode(t *testing.T) {
	// Initialize a test database
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening test database: %v", err)
	}
func TestGetUserNameByCode(t *testing.T) {
	// Initialize a test database
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		t.Fatalf("Error opening test database: %v", err)
	}
	//defer db.Close()

	// Migrate the test database schema
	db.AutoMigrate(&User{})

	// Insert a test user into the database
	testUser := User{
		FirstName:      "Ian",
		UserEmail:      "ian.n.",
		LastName:       "B",
		StickerCode:    "hello",
		BioText:        "hello world",
		ProfilePicture: "1",
		SocialList: Socials{
			Facebook: SocialType{
				SocialName: "",
				Status:     false,
				URL:        "",
			},
			Snapchat: SocialType{
				SocialName: "",
				Status:     false,
				URL:        "",
			},
			Instagram: SocialType{
				SocialName: "",
				Status:     false,
				URL:        "",
			},
		},
	}
	db.Create(&testUser)

	// Initialize a new request with the test sticker code
	req, err := http.NewRequest("GET", "/user/hello", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	// Initialize a new response recorder
	recorder := httptest.NewRecorder()

	// Call the handler function with the test request and response recorder
	GetUserNameByCode(recorder, req)

	// Check the response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	// Check the response body
	expected := `{"first_name":"Ian","last_name":"B"}`
	if recorder.Body.String() != expected {
		t.Errorf("Expected response body '%s', but got '%s'", expected, recorder.Body.String())
	}
	// Debugging statements
	// t.Logf("Response status code: %d", recorder.Code)
	// t.Logf("Response body: %s", recorder.Body.String())

	

}

