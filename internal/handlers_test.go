package internal

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestMove(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		expectedCode int
	}{
		{
			description:  "Get HTTP status 200",
			route:        "/move",
			expectedCode: 400,
		},
		{
			description:  "Get HTTP status 404, route does not exist",
			route:        "/notFound",
			expectedCode: 404,
		},
	}
	app := fiber.New()
	app.Post("/move", func(c *fiber.Ctx) error {
		word := Word{
			Word: "ild",
			PointStart: Point{
				X: 2,
				Y: 3,
			},
			PointEnd: Point{
				X: 2,
				Y: 5,
			},
		}
		w, err := json.Marshal(word)
		if err != nil {
			panic(err)
		}
		return c.SendString(string(w))
	})

	for _, test := range tests {
		req := httptest.NewRequest("Post", test.route, nil)
		resp, _ := app.Test(req, 1)
		fmt.Println(resp.StatusCode, test.expectedCode)
		if test.expectedCode != resp.StatusCode {
			t.Errorf("%d failed, expected %d", resp.StatusCode, test.expectedCode)
		} else {
			t.Logf("passed, Expected %d, got %d", test.expectedCode, resp.StatusCode)
		}

	}
}
