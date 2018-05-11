package route

import(
	"net/http"
	"github.com/gorilla/mux"
	"strings"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	if strings.TrimRight(r.RequestURI, "/") != r.RequestURI {
		http.NotFound(w, r)
	} else {
		http.FileServer(http.Dir("./src/static")).ServeHTTP(w, r)
	}
}

func NewRouter() *mux.Router{
	r :=mux.NewRouter()

	r.PathPrefix("/").HandlerFunc(StaticHandler)

	return r;
}