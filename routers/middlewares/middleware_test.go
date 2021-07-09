package middlewares

import (
	"fmt"
	"testing"
)

func TestGinLogger(t *testing.T) {
	authToken := "3333"
	_, _ = fmt.Sscanf(authToken, "Bearer %s", &authToken)
	fmt.Println(authToken)
}
