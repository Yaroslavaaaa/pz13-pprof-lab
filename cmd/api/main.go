package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // автоматически регистрирует на http.DefaultServeMux
	"runtime"

	"example.com/pz13pprof-lab/internal/work"
)

func enableLocks() {
	runtime.SetBlockProfileRate(1)     // включить Block profile
	runtime.SetMutexProfileFraction(1) // включить Mutex profile
}

func main() {
	enableLocks()
	// Используем DefaultServeMux вместо создания нового
	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		defer work.TimeIt("FibFast(38)")()
		n := 38 // достаточно тяжело для CPU
		res := work.FibFast(n)
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte((fmtInt(res))))
	})

	log.Println("Server on :8080; pprof on /debug/pprof/")
	log.Fatal(http.ListenAndServe(":8080", nil)) // nil = используем DefaultServeMux
}

func fmtInt(v int) string { return fmt.Sprintf("%d\n", v) }
