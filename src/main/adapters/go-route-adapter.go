package adapters

import (
	"encoding/json"
	"github.com/douglasdennys/clean-architecture-go/src/adapters/controllers/ports"
	"github.com/douglasdennys/clean-architecture-go/src/adapters/controllers/usecases"
	"io/ioutil"
	"net/http"
)

func AdaptRouter(controller usecases.Controller) http.HandlerFunc {
	return func (w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var httpRequest ports.HttpRequest
		converter, _ := ioutil.ReadAll(req.Body)
		httpRequest.Body = converter
		response := controller.Handle(httpRequest)

		w.WriteHeader(response.StatusCode)
		_ = json.NewEncoder(w).Encode(response.Body)
		return
	}
}
