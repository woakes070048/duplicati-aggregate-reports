// Package about displays the About page.
package about

import (
	"net/http"

	"../../lib/flight"
	"github.com/blue-jay/core/router"
)

// Load the routes.
func Load() {
	router.Get("/about", Index)
}

// Index displays the About page.
func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	//c.View.New("about/index").Render(w, r)
	c.View.New("clientstatus/index").Render(w, r)
}
