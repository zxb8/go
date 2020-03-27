package main

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
	"webapp/app"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "Assigning the port ")
	flag.Parse()
	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}
	envPort := os.Getenv("PORT")
	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	s := app.NewServer()
	s.Init(port)
	s.Start()

}
