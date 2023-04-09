package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/votoznotna/go-rest-api/internal/comment"
	"github.com/votoznotna/go-rest-api/internal/db"
	transportHttp "github.com/votoznotna/go-rest-api/internal/transport/http"
)

func Run() error {
	fmt.Println("starting up application")
	db, err := db.NewDatabase()
	if err != nil {
		log.Error("failed to setup connection to the database")
		return err
	}

	if err := db.MigrateDB(); err != nil {
		fmt.Println("failed to setup database")
		return err
	}

	fmt.Println("successfully migrated the database")

	cmtService := comment.NewService(db)
	// cmtService.PostComment(
	// 	context.Background(),
	// 	comment.Comment{
	// 		ID: "9a31bf83-28dc-4b8d-bf70-7d347a24ff2e",
	// 		Slug: "manual-test",
	// 		Author: "Elliot",
	// 		Body: "Hello World",
	// 	},
	// )
	// fmt.Println(cmtService.GetComment(context.Background(), "9a31bf83-28dc-4b8d-bf70-7d347a24ff2e"))

	handler := transportHttp.NewHandler(cmtService)

	if err := handler.Serve(); err != nil {
		log.Error("failed to gracefully serve our application")
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	fmt.Println("successfully connected and pinged database")

	return nil
}

func main() {
	fmt.Println("Go REST API Closure")
	if err := Run(); err != nil {
		log.Error(err)
		log.Fatal("Error starting up our REST API")
	}

}
