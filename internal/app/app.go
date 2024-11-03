package app

import (
	"context"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/knoxie-s/notification-service/internal/api"
	"github.com/knoxie-s/notification-service/internal/closer"
	"github.com/knoxie-s/notification-service/internal/config"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type App struct {
	serviceProvider *serviceProvider
	httpServer      *http.Server
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		err := a.runHTTPServer(ctx)
		if err != nil {
			log.Fatalf("failed to run HTTP server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runWorker(ctx)
		if err != nil {
			log.Fatalf("failed to run worker: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	initFuncs := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initHTTPServer,
	}

	for _, f := range initFuncs {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	err := config.Load(configPath)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(ctx context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := http.NewServeMux()

	api.RegisterRoutes(mux, a.serviceProvider.NotificationAPI(ctx))

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: mux,
	}

	return nil
}

func (a *App) runHTTPServer(ctx context.Context) error {
	log.Printf("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())

	go func() {
		gracefulShutdown(ctx, nil)

		log.Println("Shutting down HTTP server...")

		if err := a.httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server shutdown error: %v", err)
		}
	}()

	err := a.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (a *App) runWorker(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go gracefulShutdown(ctx, cancel)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping notification worker due to shutdown signal.")
			return nil
		case <-ticker.C:
			log.Println("Starting notification worker cycle")
		}
	}
}

func gracefulShutdown(ctx context.Context, cancel context.CancelFunc) {
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-waitSignal():
		log.Println("terminating: via signal")
	}

	if cancel != nil {
		cancel()
	}
}

func waitSignal() chan os.Signal {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	return sigterm
}
