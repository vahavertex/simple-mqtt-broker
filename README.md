# simple-mqtt-broker
Simple mqtt broker for testing

Архитектура проекта будет выглядеть так:

simple-mqtt-broker/
├── go.mod
├── go.sum
├── main.go            # Точка входа (обработка сигналов ОС)
└── broker/            # Ваш новый внутренний пакет
    └── broker.go      # Логика инициализации и запуска MQTT