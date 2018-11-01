package main

import (
	"io/ioutil"
	"net/http"
	"runtime"

	"github.com/Jeffail/tunny"
)

func main() {
	numCPUs := runtime.NumCPU()

	pool := tunny.NewFunc(numCPUs, func(payload interface{}) interface{} { //创建一个工作池
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
		/*result, err := pool.ProcessTimed(input, time.Second*5)
		if err == tunny.ErrJobTimedOut {
			http.Error(w, "Request timed out", http.StatusRequestTimeout)
		}*/

		w.Write(result.([]byte))
	})

	http.ListenAndServe(":8080", nil)
}
