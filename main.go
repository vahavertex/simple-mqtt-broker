package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	// Импортируем ваш внутренний пакет
	"simple-mqtt-broker/broker"
)

func main() {
	// 1. Инициализируем брокер на стандартном порту 1883
	mqttBroker := broker.NewBroker("1883")

	// 2. Запускаем сервер
	mqttBroker.Start()
	log.Println("Приложение успешно запущено. Ожидание сигналов...")

	// 3. Ожидаем системный сигнал прерывания (Ctrl+C)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	// 4. Корректно завершаем работу модуля
	if err := mqttBroker.Stop(); err != nil {
		log.Fatalf("Ошибка при остановке брокера: %v", err)
	}
	log.Println("Приложение завершило работу.")
}
