// package strip provides the same functionality as http.StripPrefix
// and can be used in martini instance level and request context level.
package strip

import (
	"gopkg.in/macaron.v1"
	"net/http"
	"strings"
)

// strip Prefix for every incoming http request
func Prefix(prefix string) macaron.Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if prefix == "" {
			return
		}
		if p := strings.TrimPrefix(r.URL.Path, prefix); len(p) < len(r.URL.Path) {
			r.URL.Path = p
		} else {
			http.NotFound(w, r)
		}
	}
}
