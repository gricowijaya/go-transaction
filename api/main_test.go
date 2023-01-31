package api

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// set the gin into test mode if we run the test
func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
