package main

import (
	"embed"
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed examples/*
var exampleFS embed.FS

//go:embed pdf.pdf
var pdf []byte

func servePDF(w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/pdf")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(pdf); err != nil {
		return fmt.Errorf("failed to write pdf: %w", err)
	}
	return nil
}

func main() {
	rtr := http.NewServeMux()

	rtr.HandleFunc("GET /", http.RedirectHandler("/examples/", http.StatusPermanentRedirect).ServeHTTP)
	rtr.HandleFunc("GET /examples/", http.FileServerFS(exampleFS).ServeHTTP)

	rtr.HandleFunc("GET /view.pdf", func(w http.ResponseWriter, r *http.Request) {
		if err := servePDF(w); err != nil {
			fmt.Println("error serving /view.pdf", err)
		}
	})
	rtr.HandleFunc("GET /download.pdf", func(w http.ResponseWriter, r *http.Request) {
		// Force download
		w.Header().Add("Content-Disposition", "attachment; filename=\"pdf.pdf\"")
		if err := servePDF(w); err != nil {
			fmt.Println("error serving /download.pdf", err)
		}
	})

	fmt.Println("Starting server at http://127.0.0.1:8000")
	if err := http.ListenAndServe("localhost:8000", rtr); err != nil {
		panic(fmt.Errorf("failed to serve: %w", err))
	}
}
