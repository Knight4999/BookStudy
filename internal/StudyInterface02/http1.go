package studyinterface02

import (
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars) Strings() string {
	return fmt.Sprintf("$%.2f", d)
}

type Databases map[string]dollars

func (db Databases) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
