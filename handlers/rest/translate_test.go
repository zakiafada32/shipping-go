package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zakiafada32/shipping-go/handlers/rest"
)

func TestTranslateAPI(t *testing.T) {
	tt := []struct { // <1>
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "?word=hello",
			StatusCode:          200,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "?word=hello&language=german",
			StatusCode:          200,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint:   "?word=hello&language=japan",
			StatusCode: 404,
		},
	}

	handler := http.HandlerFunc(rest.TranslateHandler)

	for _, test := range tt { // <3>
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		handler.ServeHTTP(rr, req)

		if rr.Code != test.StatusCode {
			t.Errorf(`expected status %d but received %d`,
				test.StatusCode, rr.Code)
		}

		if rr.Code == 404 {
			continue
		}

		var resp rest.Resp
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		if err != nil {
			t.Errorf("unable to unmarshal response: %v", err)
		}

		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`expected language "%s" but received %s`,
				test.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`expected Translation "%s" but received "%s"`,
				test.ExpectedTranslation, resp.Translation)
		}

	}
}
