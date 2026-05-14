package broker

import (
	"fmt"
	"log"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/listeners"
)

// MQTTBroker инкапсулирует сервер mochi-mqtt
type MQTTBroker struct {
	server *mqtt.Server
}

// NewBroker инициализирует новый экземпляр брокера на указанном порту
func NewBroker(port string) *MQTTBroker {
	b := mqtt.New(nil)

	// Конфигурируем TCP-слушатель
	tcp := listeners.NewTCP(listeners.Config{
		ID:      "t1",
		Address: fmt.Sprintf("0.0.0.0:%s", port),
	})

	err := b.AddListener(tcp)
	if err != nil {
		log.Fatalf("Ошибка добавления слушателя: %v", err)
	}

	return &MQTTBroker{
		server: b,
	}
}

// Start запускает брокер в фоновом режиме (в горутине)
func (mb *MQTTBroker) Start() {
	go func() {
		log.Println("Запуск MQTT-брокера...")
		err := mb.server.Serve()
		if err != nil {
			log.Fatalf("Ошибка во время работы сервера: %v", err)
		}
	}()
}

// Stop корректно останавливает брокер
func (mb *MQTTBroker) Stop() error {
	log.Println("Остановка MQTT-брокера...")
	return mb.server.Close()
}
