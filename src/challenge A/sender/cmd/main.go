package main

import (
	"context"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/adapters/broker/redis"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	redisDSN     = "REDIS_DSN"
	portErrorMsg = "port should be pass as argument"
)

func main() {
	appSetting := GetAppConfig()
	r := redis.InitRedis(appSetting.BrokerDSN)

	routes := getRoutes(*r)
	server := &http.Server{
		Addr:    ":" + appSetting.PortServer,
		Handler: routes,
	}

	go func() {
		log.Printf("Starting server on %s", appSetting.PortServer)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")

	// Create a context with a timeout of 5 seconds to allow existing requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server stopped")
}

func GetAppConfig() config.AppConfig {
	if len(os.Args) < 2 {
		log.Fatalf(portErrorMsg)
	}
	return config.AppConfig{
		InProduction: false,
		BrokerDSN:    os.Getenv(redisDSN),
		PortServer:   os.Args[1],
	}
}
