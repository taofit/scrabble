package internal

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestMove(t *testing.T) {
	app := fiber.New()
	var game = NewGame()
	SetupRoutes(app, game)

	tests := []struct {
		description  string
		route        string
		method       string
		requestBody  map[string]interface{}
		expectedCode int
	}{
		{
			description: "Get HTTP status 200",
			route:       "/move",
			method:      http.MethodPost,
			requestBody: map[string]interface{}{
				"Word": map[string]interface{}{
					"Word": "m#e",
					"PointStart": map[string]int{
						"x": 3,
						"y": 9,
					},
					"PointEnd": map[string]int{
						"x": 3,
						"y": 11,
					},
				},
				"Player": 1,
			},
			expectedCode: 400,
		},
		{
			description:  "Get HTTP status 404, route does not exist",
			route:        "/notFound",
			method:       http.MethodGet,
			requestBody:  nil,
			expectedCode: 404,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			rBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest(test.method, test.route, bytes.NewReader(rBody))
			req.Header.Add(`Content-Type`, `application/json`)
			resp, _ := app.Test(req)
			body, _ := ioutil.ReadAll(resp.Body)
			bodyMsg := string(body)

			if test.expectedCode != resp.StatusCode {
				t.Errorf("%d failed, expected %d, message: %s", resp.StatusCode, test.expectedCode, bodyMsg)
			} else {
				t.Logf("passed, Expected %d, got %d, message: %s", test.expectedCode, resp.StatusCode, bodyMsg)
			}
		})
	}
}
