/*
	{
		"created_on": "12 june 2017",
		"aim": "To develop REST APIs for iOS, Android and Web apps",
		"code_by": "Rishikesh Agrawani"
	}
*/

package main

import (
	"fmt"
	"net/http"
)


func Home(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, htmlTextHome)
}

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

//HTML TEXTS
var htmlTextHome=	
