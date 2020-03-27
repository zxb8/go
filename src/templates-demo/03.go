package main

import (
	"log"
	"html/template"
	"net/http"
)

type(
	BooksPage struct{
		BooksListing []Book
	}

	Book struct{
		Isbn string
		Author string
		Title string
		PubDate string
		Price Currency
	}
	Currency float64
)

func main(){
	http.HandleFunc("/books",booksHandler1)
	http.Handle("/",http.FileServer(http.Dir("client")))
	log.Fatal(http.ListenAndServe(":8090",nil))
}


func booksHandler1(w http.ResponseWriter,r *http.Request){
	booksPage := &BooksPage{
		[]Book{
			{"9030","Author 1","Book title1","07/21/2001",9.43},
			{"9030","Author 1","Book title1","07/21/2001",9.43},
		},
	}
	tmpl := template.Must(template.ParseFiles("assets/index.html"))
	err :=tmpl.Execute(w,booksPage)
	if err != nil{
		panic(err)
	}
}