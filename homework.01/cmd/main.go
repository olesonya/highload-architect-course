package main

import (
	"context"

	logger "github.com/sirupsen/logrus"

	"github.com/olesonya/highload-architect-course/homework.01/internal/app"
)

func main() {
	ctx := context.Background()

	app, err := app.NewApp(ctx)
	if err != nil {
		logger.Fatalf("app.NewApp(ctx): %v", err)
	}

	app.Run()
}
