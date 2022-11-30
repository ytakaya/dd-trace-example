package main

import (
	"net/http"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(tracer.WithDebugMode(true))
	defer tracer.Stop()

	mux := httptrace.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		childSpan, _ := tracer.SpanFromContext(r.Context())
		childSpan.SetOperationName("http.request")
		defer childSpan.Finish()

		w.Write([]byte("Hello World!\n"))
	})
	http.ListenAndServe(":8080", mux)
}
