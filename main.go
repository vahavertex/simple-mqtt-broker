package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func main() {
	// 1. Создаем экземпляр сервера (в v2 используется функция New)
	b := mqtt.New(nil)

	// 2. Создаем TCP-слушатель на стандартном MQTT порту 1883
	tcp := listeners.NewTCP(listeners.Config{
		ID:      "t1",
		Address: ":1883",
	})

	// 3. Добавляем слушатель в брокер
	err := b.AddListener(tcp)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Запускаем брокер в отдельной горутине (метод Serve блокирует поток)
	go func() {
		err := b.Serve()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("MQTT-брокер успешно запущен на порту :1883")

	// 5. Ожидаем системный сигнал для корректного завершения (Graceful Shutdown)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// 6. Закрываем брокер
	log.Println("Остановка брокера...")
	b.Close()
}
