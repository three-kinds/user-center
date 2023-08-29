package test_utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
)

func PrepareTestController(handler gin.HandlerFunc, routerPath string, requestPath string, body gin.H, accessToken string) (int, map[string]any) {
	content, _ := json.Marshal(body)
	request, _ := http.NewRequest("POST", requestPath, strings.NewReader(string(content)))
	if accessToken != "" {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}

	router := gin.New()
	router.POST(routerPath, handler)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	statusCode := w.Code
	var responseJson = map[string]any{}
	_ = json.Unmarshal([]byte(w.Body.String()), &responseJson)
	return statusCode, responseJson
}
