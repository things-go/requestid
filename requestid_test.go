package requestid

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func TestRequestID(t *testing.T) {
	tests := map[string]struct {
		requestIDHeader  string
		request          func() *http.Request
		expectedResponse string
	}{
		"Retrieves Request Id from default header": {
			"X-Request-Id",
			func() *http.Request {
				req := httptest.NewRequest("GET", "/", nil)
				req.Header.Add("X-Request-Id", "req-123456")

				return req
			},
			"RequestId: req-123456",
		},
		"Retrieves Request Id from custom header": {
			"X-Trace-Id",
			func() *http.Request {
				req := httptest.NewRequest("GET", "/", nil)
				req.Header.Add("X-Trace-Id", "trace:abc123")

				return req
			},
			"RequestId: trace:abc123",
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()

		r := gin.New()

		r.Use(RequestId(WithRequestIdHeader(test.requestIDHeader)))
		r.GET("/", func(c *gin.Context) {
			requestID := GetRequestId(c)
			response := fmt.Sprintf("RequestId: %s", requestID)
			_, _ = w.WriteString(response)
		})
		r.ServeHTTP(w, test.request())

		if w.Body.String() != test.expectedResponse {
			t.Fatalf("RequestId was not the expected value")
		}
	}
}

func TestFirstRequest(t *testing.T) {
	var gotResponse string

	w := httptest.NewRecorder()

	r := gin.New()
	r.Use(RequestId(WithNextRequestId(NextRequestId)))
	r.GET("/", func(c *gin.Context) {
		requestId := GetRequestId(c)
		gotResponse = fmt.Sprintf("RequestId: %s", requestId)
		_, _ = w.WriteString(gotResponse)
	})

	req := httptest.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	if w.Body.String() != gotResponse {
		t.Fatalf("RequestId was not the expected value")
	}
}
