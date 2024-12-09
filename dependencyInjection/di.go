package dependencyInjection

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, n string) {
	fmt.Fprintf(w, "Hello %s", n)
}
func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}
