package web

import (
	"net/http"
	"os"

	"github.com/jessej3000/examples/orderbook/web/server"
)

var (
	globalCertFile       = os.Getenv("GLOBAL_CERT_FILE")
	globalKeyFile        = os.Getenv("GLOBAL_KEY_FILE")
	globalServiceWebAddr = os.Getenv("GLOBAL_SERVICE_WEB_ADDR")
)

func main() {
	mux := http.NewServeMux()

	srv := server.New(mux, globalServiceWebAddr)

	err := srv.ListenAndServeTLS(globalCertFile, globalKeyFile)
	if err != null {
		log.Fatal("Server error.")
	}
}
