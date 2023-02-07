package main

import (
	"github.com/MrMelon54/ktane-mod-ideas/src-api/structure"
	"github.com/MrMelon54/ktane-mod-ideas/src-api/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"strings"
	"xorm.io/xorm"
)

func main() {
	log.Println("[Main] Starting up KTaNE Mod Ideas")

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[Main] Loading database")
	dbEnv := os.Getenv("DB")
	var engine *xorm.Engine
	if strings.HasPrefix(dbEnv, "sqlite:") {
		engine, err = xorm.NewEngine("sqlite3", strings.TrimPrefix(dbEnv, "sqlite:"))
	} else if strings.HasPrefix(dbEnv, "mysql:") {
		engine, err = xorm.NewEngine("mysql", strings.TrimPrefix(dbEnv, "mysql:"))
	} else {
		log.Fatalln("[Main] Only mysql and sqlite are supported")
	}
	if err != nil {
		log.Fatalf("Unable to load database (\"%s\")\n", dbEnv)
	}
	check(engine.Sync(&structure.Idea{}, &structure.User{}, &structure.Vote{}))

	s := web.New(engine).SetupWeb()
	err = s.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Printf("[Http] The HTTP server shutdown successfully\n")
		} else {
			log.Printf("[Http] Error trying to host the HTTP server: %s\n", err.Error())
		}
	}
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
