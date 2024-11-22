//This is the testing file of handler.go

package handler_test

import (
	"net/http"
	"net/http/httptest"
	"newsAPI/cmd/api-server/internal/handler"
	"testing"
)

func Test_PostNews(t *testing.T) {

	testCases := []struct {
		name           string
		expectedStatus int
	}{

		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			//Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			//Act about POST
			handler.postNews(w, r)

			//Asserting response errors
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_GetAllNews(t *testing.T) {

	testCases := []struct {
		name           string
		expectedStatus int
	}{

		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			//Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			//Act about POST
			handler.GetNewsByID()(w, r)

			//Asserting response errors
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func Test_UpdateNewsByID(t *testing.T) {

	testCases := []struct {
		name           string
		expectedStatus int
	}{

		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			//Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			//Act about POST
			handler.UpdateNewsByID()(w, r)

			//Asserting response errors
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

func DeleteNewsByID(t *testing.T) {

	testCases := []struct {
		name           string
		expectedStatus int
	}{

		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			//Arrange
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			//Act about POST
			handler.DeleteNewsByID()(w, r)

			//Asserting response errors
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected :%d, got: %d", tc.expectedStatus, w.Result().StatusCode)
			}

		})
	}
}
