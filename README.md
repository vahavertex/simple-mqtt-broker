# simple-mqtt-broker
Simple mqtt broker for testing

Архитектура проекта будет выглядеть так:

simple-mqtt-broker/
├── go.mod
├── go.sum
├── main.go            # Точка входа (обработка сигналов ОС)
└── broker/            # Ваш новый внутренний пакет
    └── broker.go      # Логика инициализации и запуска MQTT

### 3. Запуск и проверка

Запустите сервер:

```bash
go run main.go
```

Для проверки отправки и получения сообщений можно использовать стандартные утилиты `mosquitto_sub` и `mosquitto_pub`:

* **Подписка на топик (в отдельном терминале):**

  ```bash
  mosquitto_sub -h localhost -p 1883 -t "test/topic"
  ```

* **Публикация сообщения:**

  ```bash
  mosquitto_pub -h localhost -p 1883 -t "test/topic" -m "Hello from Go Broker"
  ```
