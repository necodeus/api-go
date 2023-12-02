package common_api

import (
	"fmt"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"message\": \"Hello, CommonIndex!\"}"))

	time := time.Now()
	fmt.Println("common/index! Now time: ", time)
}
