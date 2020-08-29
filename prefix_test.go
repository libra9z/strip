package strip

import (
	macaron "gopkg.in/macaron.v1"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStripPrefix(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/foo/bar", nil)
	Prefix("/foo").(func(http.ResponseWriter, *http.Request))(w, r)
	if r.URL.Path != "/bar" {
		t.Fatalf("Strip Prefix Failed")
	}
}

func TestInMartini(t *testing.T) {
	m := macaron.New()
	m.Use(Prefix("/foo"))
	m.Use(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bar" {
			t.Fatalf("Strip Prefix Failed")
		}
	})
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/foo/bar", nil)
	m.ServeHTTP(w, r)
}

func TestInRequestContext(t *testing.T) {
	m := macaron.Classic()
	m.Get("/foo/bar", Prefix("/foo"), func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/bar" {
			t.Fatalf("Strip Prefix Failed")
		}
	})
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/foo/bar", nil)
	m.ServeHTTP(w, r)
}
