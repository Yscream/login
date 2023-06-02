package main

import (
	"log"
	"net/http"

	"github.com/Yscream/login/pkg/config"
	"github.com/Yscream/login/pkg/repository"
	"github.com/Yscream/login/pkg/router"
	"github.com/Yscream/login/pkg/service"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	conf := config.UnmarshalYAML("./cmd")

	db, err := sqlx.Connect("pgx", conf.DB.URL())
	if err != nil {
		log.Fatal("Couldn't connect to db: ", err)
	}

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	r := router.NewRouter(svc)

	log.Fatal(http.ListenAndServe(":8033", r))
}
