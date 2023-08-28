package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/three-kinds/user-center/initializers"
	"net/http"
	"testing"
	"time"
)

func TestMainFunc(t *testing.T) {
	go main()
	time.Sleep(time.Second)

	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/api/health", initializers.Config.ServerPort))
	assert.Equal(t, nil, err)
	assert.Equal(t, 200, resp.StatusCode)

	defer func() {
		_ = resp.Body.Close()
	}()

	jr := map[string]string{}
	err = json.NewDecoder(resp.Body).Decode(&jr)
	assert.Equal(t, nil, err)
	assert.Equal(t, "OK", jr["status"])
}
