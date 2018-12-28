package main

import (
	"fmt"
	"github.com/Jeffail/tunny"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

func main() {
	//sStr := "acbdbaab"
	//tStr := "aabd"
	var left int
	fmt.Println(left == 0)
	os.Exit(0)
	numCPUs := runtime.NumCPU()

	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} {
		var result []byte

		// TODO: Something CPU heavy with payload

		return result
	})
	defer pool.Close()

	http.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		input, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
		defer r.Body.Close()

		// Funnel this work into our pool. This call is synchronous and will
		// block until the job is completed.
		result := pool.Process(input)

		w.Write(result.([]byte))
	})

	http.ListenAndServe(":8080", nil)
}
