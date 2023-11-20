# delivery-check
Демонстрационный сервис, отображающий данные о заказе
## Возможности
* Генерация новых валидных и невалидных данных и их отправка в канал nats-streaming
* Получение данных из канала nats-streaming, проверка их валидности, запись в базу данных в случае валидности данных
* Получение информации о заказе из базы данных по uid заказа
## Технологии
* [Gin Framework](https://github.com/gin-gonic/gin)
* [Nats-Streaming](https://hub.docker.com/_/nats-streaming)
## Цели
* Добавить кеширование данных и восстановление кеша
* Добавить автотестирование
* Воспользоваться утилитами WRK и Vegeta
* Создать простейший интерфейс для сервиса
