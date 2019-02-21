package kong

import (
	"os"
)

func kongURL() string {
	v := os.Getenv("KONG_URL")
	if v == "" {
		v = "http://localhost:8001"
	}
	return v
}
