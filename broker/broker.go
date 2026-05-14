package broker

import (
	"context"
	"fmt"
	"log/slog" // Используем современный структурированный логгер Go 1.21+

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

// Config хранит настройки для инициализации брокера
type Config struct {
	Port string
	ID   string
}

// MQTTBroker инкапсулирует сервер mochi-mqtt
type MQTTBroker struct {
	server *mqtt.Server
	cfg    Config
}

// NewBroker инициализирует новый экземпляр брокера с конфигурацией
func NewBroker(cfg Config) *MQTTBroker {
	b := mqtt.New(nil)
	_ = b.AddHook(new(auth.AllowHook), nil)

	tcp := listeners.NewTCP(listeners.Config{
		ID:      cfg.ID,
		Address: fmt.Sprintf("0.0.0.0:%s", cfg.Port),
	})

	err := b.AddListener(tcp)
	if err != nil {
		// Вместо Fatalf возвращаем панику на этапе инициализации,
		// чтобы вызывающий код мог обработать ошибку
		panic(fmt.Sprintf("failed to add listener: %v", err))
	}

	return &MQTTBroker{
		server: b,
		cfg:    cfg,
	}
}

// Start запускает брокер и блокирует поток до завершения Context
func (mb *MQTTBroker) Start(ctx context.Context) error {
	errChan := make(chan error, 1)

	// Запуск самого сервера mochi-mqtt в фоновом режиме
	go func() {
		slog.Info("Starting MQTT broker", "port", mb.cfg.Port, "id", mb.cfg.ID)
		if err := mb.server.Serve(); err != nil {
			errChan <- err
		}
	}()

	// Ожидаем либо ошибку старта, либо сигнал отмены Context
	select {
	case err := <-errChan:
		return fmt.Errorf("broker server error: %w", err)
	case <-ctx.Done():
		slog.Info("Shutting down MQTT broker via context signal")
		return mb.server.Close()
	}
}
