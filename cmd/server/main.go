package main

import (
	"context"
	"net/http"

	"go-kafka/internal/config"
	"go-kafka/internal/factory"
	"go-kafka/internal/infrastructure/db"
	"go-kafka/internal/infrastructure/kafka"
	"go-kafka/internal/routes"
	"go-kafka/pkg/graceful"
	"go-kafka/pkg/logger"

	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
)

func program(state overseer.State) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	graceful.SetupGracefulShutdown(cancel)

	cfg := config.LoadConfig()
	logger.InitLogFile(cfg.App.LogFile)

	dbConn, err := db.ConnectDB(cfg.Database)
	if err != nil {
		logger.Errorf("❌ Error connecting to DB: %v", err)
	}

	kafkaClient := kafka.NewKafkaClient(cfg.Kafka)

	go func() {
		err = kafkaClient.ConsumeAvro()
		if err != nil {
			logger.Fatalf("Kafka consumer error: %v", err)
		}
	}()

	// Factory
	f := factory.NewFactory(dbConn, kafkaClient)

	// Init routes using factory
	r := routes.InitRouter(f)

	logger.Infof("✅ Server started at :%d", cfg.App.Port)
	http.Serve(state.Listener, r)

	ctx.Done()

	db.CloseDB()

	logger.Infof("🛑 Server stopped")
}

func main() {
	debug := config.GetAppEnv() == "development"

	overseer.Run(overseer.Config{
		Program:       program,
		Address:       ":" + config.GetAppPort(),
		Fetcher:       &fetcher.File{Path: config.GetAppBinFile(), Interval: 3},
		RestartSignal: graceful.RestartSignal,
		Debug:         debug,
	})
}
