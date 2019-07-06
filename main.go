package main

import (
	"log"
	"net/http"
	"os"

	"github.com/caarlos0/env"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/yasminromi/radix-hack/handler"
	"github.com/yasminromi/radix-hack/model"
	"github.com/yasminromi/radix-hack/service"
	"gopkg.in/olivere/elastic.v6"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("File .env not found, reading configuration from ENV")
	}

	var cfg model.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalln("Failed to parse ENV")
	}

	elasticClient, err := elastic.NewClient(elastic.SetURL(cfg.ElasticSearchUrl), elastic.SetSniff(false))
	if err != nil {
		log.Fatal("Error Creating Elastic Client: ", err)
	}

	log.Printf("Elastic Search Client Created")

	elasticService := &service.ElasticService{
		ElasticCLI: elasticClient,
	}

	chatHandler := &handler.Handler{
		Upgrader: websocket.Upgrader{},
		Service:  elasticService,
	}

	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", chatHandler.HandleConnections)
	http.HandleFunc("/createTicket", chatHandler.CreateViaPost)
	http.HandleFunc("/updateTicket", chatHandler.UpdateViaPost)

	go chatHandler.HandleMessages()

	log.Println("http server started on " + os.Getenv("PORT"))

	error := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if error != nil {
		log.Fatal("ListenAndServe: ", error)
	}

}
