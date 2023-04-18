package main

import (
	"backend/controllers"
	"backend/database"
	"backend/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"


	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"

)

const JWTSecret = "secret"

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
}
