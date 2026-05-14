package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"simple-mqtt-broker/broker"
)

func main() {
	// Инициализируем структурированный логгер для вывода в формате JSON или Text
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// 1. Создаем контекст, который автоматически отменится при Ctrl+C (SIGINT) или SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 2. Определяем конфигурацию (в будущем можно брать из os.Getenv)
	cfg := broker.Config{
		Port: "1883",
		ID:   "main-tcp-listener",
	}

	mqttBroker := broker.NewBroker(cfg)

	// Kanal для отслеживания завершения работы брокера
	done := make(chan struct{})

	// 3. Запускаем брокер в горутине, передавая ему контекст
	go func() {
		if err := mqttBroker.Start(ctx); err != nil {
			slog.Error("Broker stopped with error", "error", err)
		}
		close(done)
	}()

	slog.Info("Application initialized. Press Ctrl+C to exit.")

	// Блокируем main до тех пор, пока Start() не завершит работу (после отмены ctx)
	<-done

	// Дополнительный таймаут для гарантии завершения всех фоновых процессов
	_, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	slog.Info("Application successfully stopped")
}
