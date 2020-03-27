package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Author *Author `json"author"`
}

type Author struct{
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`

}

//init books var as a slice Book struct

var books []Book


func getBooks(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)

	for _,item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_  =json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books,book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter,r *http.Request){

}

func deleteBook(w http.ResponseWriter,r *http.Request){

}

func main(){
	r :=mux.NewRouter()

	books = append(books,Book{ID:"1",Isbn:"448743",Title:"Book One",Author:&Author{FirstName:"Johnn",LastName:"doe"}})
    books = append(books,Book{ID:"2",Isbn:"811111",Title:"测试",Author:&Author{FirstName:"Johnn",LastName:"doe"}})
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080",r))
}
