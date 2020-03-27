package main

import (
	//"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error{
	f := p.Title + ".txt"
	return ioutil.WriteFile(f,p.Body,0600)
}

func load(title string)(*Page,error){
	f := title +".txt"
	body,err := ioutil.ReadFile(f)
	if err != nil{
		return nil,err
	}
	return &Page{Title:title,Body:body},nil
}

func view(w http.ResponseWriter,r *http.Request){
	title := r.URL.Path[len("/test/"):]
	//fmt.Println(r.URL.Path)  //  /test/hh
	//fmt.Println(title)
	p,_ := load(title)

	t ,_ :=template.ParseFiles("test.html")
	t.Execute(w,p)
}

func edit(w http.ResponseWriter,r *http.Request){
	title := r.URL.Path[len("/edit/"):]
	p,_ :=load(title)
	t,_:=template.ParseFiles("edit.html")
	t.Execute(w,p)
}

func main(){
	p := &Page{Title:"Test",Body:[]byte("Welcome to the Test App")}
	p.save()

	http.HandleFunc("/test/",view)
	http.HandleFunc("/edit",edit)
	http.ListenAndServe(":8080",nil)
}
