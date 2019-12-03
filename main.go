package main

import (
	"flag"
	"github.com/hurtuh/indriver/domain"
	"github.com/hurtuh/indriver/handlers"
	"github.com/hurtuh/indriver/mysql_repository"
	"github.com/hurtuh/indriver/server"
	"github.com/hurtuh/indriver/service"
	"log"
	"time"
)

func main() {

	serverPort := flag.String("server port", ":8080", "working port")
	cache := flag.Bool("cache", false, "use cache")
	dbHost := flag.String("host", "127.0.0.1", "db host")
	dbPort := flag.String("port", "3306", "db port")
	dbUser := flag.String("user", "inDriver", "db user")
	dbPass := flag.String("pass", "IsCool", "db pass")
	dbName := flag.String("name", "interview", "db name")
	flag.Parse()

	var repo domain.Repository
	var err error

	if !*cache {
		repo, err = mysql_repository.NewMysqlRepo(*dbHost, *dbPort, *dbUser, *dbPass, *dbName)
		if err != nil {
			log.Fatal(err)
		}
	}

	logic := service.NewLogic(repo)
	handlersFuncs := handlers.NewHandlers(logic)
	serv := server.InitServer(handlersFuncs)

	log.Printf("Start server on %s port and %s time", *serverPort, time.Now().Format("2006-01-02 15:04:05"))
	err = serv.StartServer(*serverPort)
	if err != nil {
		log.Fatal(err)
	}
}
