package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello test!"))
}

func helloName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + GetParam(r, "name")))
}

func helloNames(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello " + GetParam(r, "first-name") + GetParam(r, "middle-name") + GetParam(r, "last-name")))
}

func BenchmarkRootMatch(b *testing.B) {
	// Create route
	r := New("/")
	r.Add("/", http.HandlerFunc(hello))

	req, _ := http.NewRequest("GET", "http://test.com/", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Match(req)
	}
}

func BenchmarkPathMatch(b *testing.B) {
	// Create route
	r := New("/")
	r.Add("/some/path/to/match", http.HandlerFunc(hello))

	req, _ := http.NewRequest("GET", "http://test.com/some/path/to/match", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Match(req)
	}
}

func BenchmarkParamMatch(b *testing.B) {
	// Create route
	r := New("/")
	r.Add("/hello/:name", http.HandlerFunc(helloName))

	req, _ := http.NewRequest("GET", "http://test.com/hello/joe", nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Match(req)
	}
}

func BenchmarkMultiParamMatch(b *testing.B) {
	// Create route
	r := New("/")
	r.Add("/hello/:first-name/:middle-name/:last-name", http.HandlerFunc(helloNames))

	req, _ := http.NewRequest("GET", "http://test.com/hello/joe/x/smith", nil)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Match(req)
	}
}

func BenchmarkRootDispatch(b *testing.B) {
	// Create route
	r := New("/")
	r.Add("/", http.HandlerFunc(hello))
	d := Route(r)

	req, _ := http.NewRequest("GET", "http://test.com", nil)
	res := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.ServeHTTP(res, req)
	}
}

func BenchmarkParamDispatch(b *testing.B) {
	// Create route
	r := New("/")
	r.Add("/hello/:name", http.HandlerFunc(helloName))
	d := Route(r)

	req, _ := http.NewRequest("GET", "http://test.com/hello/joe", nil)
	res := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.ServeHTTP(res, req)
	}
}

func BenchmarkMultiParamDispatch(b *testing.B) {
	// Create route
	r := New("/")
	r.Add("/hello/:first-name/:middle-name/:last-name", http.HandlerFunc(helloNames))
	d := Route(r)

	req, _ := http.NewRequest("GET", "http://test.com/hello/joe/x/smith", nil)
	res := httptest.NewRecorder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		d.ServeHTTP(res, req)
	}
}
