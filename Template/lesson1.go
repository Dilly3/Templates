package main

import (
	"fmt"

	"net/http"
	"text/template"
)

func Println(s interface{}) {
	fmt.Println(s)
}
func Home(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, "/homepage.html")
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is my About page")
}
func AddBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is my Add blog page")
}
func RenderPage(w http.ResponseWriter, addr string) {
	parsedTemplate, err := template.ParseFiles(fmt.Sprintf("./tmpl%s", addr))

	if err != nil {
		Println(err)
	}
	parsedTemplate.Execute(w, nil)
}
func main() {
	type PortID = string
	const PORT_NUMBER PortID = ":9080"
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello World")
	// 	if err != nil {
	// 		log.Fatal(error.Error(err))
	// 	}
	// 	fmt.Printf("%v Bytes\n", n)
	// })
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/addblog", AddBlog)
	fmt.Print("Server Running\n")
	http.ListenAndServe(PORT_NUMBER, nil)
}
