package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var handler http.Handler
var server gin.ResponseWriter


func server() {
	http.ListenAndServe(":3000", http.HandleFunc("/", (w , r ) {

	})
}
}
