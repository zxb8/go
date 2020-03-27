package app

import (
	"log"
	"net/http"
	"webapp/router"
)

type Server struct{
	port string
}

func NewServer() Server{
	return Server{}
}

func (s *Server) Init(port string){
	log.Println("Initializing server....")
	s.port = ":" + port
}

func (s *Server) Start(){
	log.Println("Starting server om port "+ s.port)

	r := router.NewRouter()
	r.Init()

	http.ListenAndServe(s.port,r.Router )
}
