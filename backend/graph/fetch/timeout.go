package fetch

import (
	"net/http"
	"time"
)

func init() {
	http.DefaultClient.Timeout = time.Second * 30
}
