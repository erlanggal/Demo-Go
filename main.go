package main

import (
	"net/http"
	"fmt"
)

type page struct {

}

func (p page)  ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Hello World")
}

func main()  {
	var p page;

	err:= http.ListenAndServe("localhost:3000", p)

	checkError(err)
}

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}