package main

import (
	"backend/controllers"
	"backend/database"
	"backend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

const JWTSecret = "secret"

func createTestJWT(userId uint) string {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    strconv.Itoa(int(userId)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWTSecret))
	if err != nil {
		return ""
	}

	return tokenString
}

func TestRegisterNewCode(t *testing.T) {
	// initialize test database
	database.Connect()
	defer database.DB.Migrator().DropTable(&models.User{}, &models.UnregisteredCodes{})

	// create test user with admin role
	testUser := models.User{
		Email: "test@email.com",
		Role:  "admin",
	}
	database.DB.Create(&testUser)

	// create JWT token for test user
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Issuer:    strconv.Itoa(int(testUser.ID)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(JWTSecret))

	// create test request
	reqData := map[string]string{"code": "testcode"}
	reqBytes, _ := json.Marshal(reqData)
	req := httptest.NewRequest(http.MethodPost, "/api/register-code", bytes.NewReader(reqBytes))
	req.Header.Set("Content-Type", "application/json")
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: tokenString,
	}
	req.AddCookie(cookie)

	// create test fiber app
	app := fiber.New()

	// create test context
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	// set user value
	ctx.Locals("jwt", tokenString)

	// set request body
	ctx.Request().SetBody(reqBytes)

	// call handler function
	if err := controllers.RegisterNewCode(ctx); err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	// check response status code
	if ctx.Response().StatusCode() != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, ctx.Response().StatusCode())
	}
}

func TestGetUser(t *testing.T) {
	// initialize test database
	database.Connect()
	defer database.DB.Migrator().DropTable(&models.User{})

	// create test user
	testUser := models.User{
		Id:        3,
		FirstName: "Test User",
		Email:     "testuser@test.com",
		Role:      "user",
	}
	database.DB.Create(&testUser)

	// create JWT token for test user
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Issuer:    strconv.Itoa(int(testUser.ID)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(JWTSecret))

	// create test request
	reqData := map[string]string{"id": "3"}
	reqBytes, _ := json.Marshal(reqData)
	req := httptest.NewRequest(http.MethodPost, "/api/get-user", bytes.NewReader(reqBytes))
	req.Header.Set("Content-Type", "application/json")
	cookie := &http.Cookie{
		Name:  "jwt",
		Value: tokenString,
	}
	req.AddCookie(cookie)

	// create test fiber app
	app := fiber.New()

	// create test context
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	// set user value
	ctx.Locals("jwt", tokenString)

	// set request body
	ctx.Request().SetBody(reqBytes)

	// call handler function
	if err := controllers.GetUser(ctx); err != nil {
		t.Errorf("expected no error but got %v", err)
	}

	// call handler function
	err := controllers.GetUser(ctx)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// check response status code
	if ctx.Response().StatusCode() != http.StatusOK {
		t.Errorf("expected status code %d but got %d", http.StatusOK, ctx.Response().StatusCode())
	}
	// Delete test user from database
	database.DB.Delete(&testUser)
}
func createFormFileUploadRequest(t *testing.T, app *fiber.App, file *os.File, userID uint) *http.Request {
	// Create a new multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create a new form file field
	part, err := writer.CreateFormFile("profile_picture", filepath.Base(file.Name()))
	if err != nil {
		t.Fatal(err)
	}

	// Copy the file contents into the form file field
	_, err = io.Copy(part, file)
	if err != nil {
		t.Fatal(err)
	}

	// Add the user ID to the URL path
	path := fmt.Sprintf("/api/users/%d/profile_picture", userID)

	// Create a new HTTP POST request with the multipart form data
	req, err := http.NewRequest("POST", path, body)
	if err != nil {
		t.Fatal(err)
	}

	// Set the content type header to multipart/form-data
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		t.Fatal(err)
	}

	return req
}

func TestUpdateProfilePicture(t *testing.T) {
	app := fiber.New()
	database.Connect()
	defer database.DB.Migrator().DropTable(&models.User{}, &models.UnregisteredCodes{})

	// Create a temporary directory to store uploaded files
	dir, err := ioutil.TempDir("", "test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(dir)

	// Set up test user
	user := models.User{
		Email:     "test@test.com",
		FirstName: "Test",
		LastName:  "User",
	}
	database.DB.Create(&user)

	// Create a test JWT for the user
	tokenString := createTestJWT(user.Id)

	// Set up test file
	file, err := os.Open("profile_pictures/gator-test.jpg")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	// Create a test multipart request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("profile_picture", "test_image.jpg")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		t.Fatalf("Failed to copy file to form file: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("Failed to close multipart writer: %v", err)
	}

	// Create a new request
	req := httptest.NewRequest(http.MethodPost, "/api/update-picture/"+strconv.Itoa(int(user.Id)), body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Cookie", "jwt="+tokenString)

	// Make the request to update the user's profile picture
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check that the user's profile picture was updated in the database
	var updatedUser models.User
	database.DB.First(&updatedUser, user.Id)
	if updatedUser.ProfilePicture == "" {
		t.Errorf("Expected profile picture to be set but got empty string")
	}
}

func TestGetProfilePicture(t *testing.T) {
	// Initialize database and create test user
	database.Connect()
	defer database.DB.Migrator().DropTable(&models.User{})

	user := models.User{
		Email:          "testuser@example.com",
		ProfilePicture: "profile_pictures/gator-test.jpg",
	}
	database.DB.Create(&user)

	// Create test Fiber context
	app := fiber.New()
	req := httptest.NewRequest(http.MethodGet, "/api/profile-picture/"+strconv.Itoa(int(user.Id)), nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	// Check if response status code is OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	// Check if response body contains the expected file content type
	expectedContentType := "image/jpg"
	actualContentType := resp.Header.Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected content type %s but got %s", expectedContentType, actualContentType)
	}

	// Check if response body contains the expected file content
	expectedContent, err := ioutil.ReadFile(user.ProfilePicture)
	if err != nil {
		t.Error(err)
	}
	actualContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(expectedContent, actualContent) {
		t.Error("Response body does not match expected content")
	}

	// Delete test user from database
	database.DB.Delete(&user)
}

func TestUpdateName(t *testing.T) {
	// Initialize a test database
	// initialize test database
	database.Connect()
	defer database.DB.Migrator().DropTable(&models.User{})

	// Initialize the app with the test database
	app := fiber.New()

	// Create a test user
	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
	}
	database.DB.Create(&user)

	// Create a test JWT token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    strconv.Itoa(int(user.Id)),
	}).SignedString([]byte(JWTSecret))
	if err != nil {
		t.Fatalf("failed to create test JWT token: %v", err)
	}

	// Define the test cases
	testCases := []struct {
		name       string
		reqBody    string
		statusCode int
		respBody   string
	}{
		{
			name: "update name successfully",
			reqBody: `{
				"first_name": "Jane",
				"last_name": "Doe"
			}`,
			statusCode: fiber.StatusOK,
			respBody:   `{"message":"success"}`,
		},
		{
			name:       "unauthenticated request",
			reqBody:    `{}`,
			statusCode: fiber.StatusUnauthorized,
			respBody:   `{"message":"Unauthenticated"}`,
		},
		{
			name: "invalid request body",
			reqBody: `{
				"first_name": 123,
				"last_name": "Doe"
			}`,
			statusCode: fiber.StatusBadRequest,
			respBody:   `{"error":"cannot parse request body"}`,
		},
	}

	// Run the test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create the request
			req := httptest.NewRequest(fiber.MethodPost, "/api/update-name", bytes.NewBufferString(tc.reqBody))
			req.Header.Set("Cookie", fmt.Sprintf("jwt=%s", token))

			fmt.Println("Checked.")

			// Make the request to update the user's name
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("failed to make request: %v", err)
			}
			fmt.Println("Checked.")
			// Check the response status code
			if resp.StatusCode != tc.statusCode {
				t.Errorf("expected status code %d but got %d", tc.statusCode, resp.StatusCode)
			}

			// Check the response body
			respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("failed to read response body: %v", err)
			}
			if string(respBody) != tc.respBody {
				t.Errorf("expected response body %s but got %s", tc.respBody, string(respBody))
			}
		})
	}
}


func TestUpdateBio(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %s", err)
	}
	defer db.Close()

	// Set up the mock database response
	rows := sqlmock.NewRows([]string{"id", "first_name", "email", "bio_text"}).AddRow(1, "John Doe", "johndoe@example.com", "old bio text")
	mock.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+)").WillReturnRows(rows)

	token := createTestJWT(1)

	// Set up the request and response objects
	app := fiber.New()
	req := httptest.NewRequest(http.MethodPost, "/api/update-bio", strings.NewReader(`{"id": "1",bio_text": "new bio text"}`))
	req.Header.Set("Cookie", fmt.Sprintf("jwt=%s", token))
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("failed to send request: %s", err)
	}
	controllers.UpdateBio(app.AcquireCtx(&fasthttp.RequestCtx{}))
	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		//t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}

	// Check the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %s", err)
	}
	expectedBody := `{"message":"success"}`
	if string(body) != expectedBody {
		t.Errorf("unexpected response body: got %s, want %s", string(body), expectedBody)
	}
	// Check the database update
	expectedQuery := "UPDATE `users` SET `bio_text`=? WHERE `id` = ?"
	expectedArgs := []driver.Value{"new bio text", 1}
	errEx := mock.ExpectationsWereMet()
	if errEx != nil {
		t.Errorf("unexpected query/args: want %v, %v", expectedQuery, expectedArgs)
	}

}

func TestIsLoggedIn(t *testing.T) {
	// Create a mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %s", err)
	}
	defer db.Close()

	// Set up the mock database response
	rows := sqlmock.NewRows([]string{"id", "email"}).AddRow(1, "test@example.com")
	mock.ExpectQuery("SELECT (.+) FROM users WHERE id = (.+)").WillReturnRows(rows)

	// Set up the request and response objects
	app := fiber.New()
	req := httptest.NewRequest(http.MethodGet, "/api/is-logged-in", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: "mock_token"})
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("failed to send request: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		//t.Errorf("unexpected status code: got %d, want %d", resp.StatusCode, http.StatusOK)
	}
	//Check the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %s", err)
	} else {
		fmt.Println("Response was successfully read.")
	}
	expectedBody := `{"status":true}`
	assert.Equal(t, expectedBody, string(body))

}
