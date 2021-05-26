## Golang with RabbitMQ
---
* Get RabbitMQ:
```
docker run -d -p 15672:15672 -p 5672:5672 -p 15671:15671 -p 5671:5671 -p 4369:4369 rabbitmq:3-management
```
* Publish message:
```
go run main.go
```
* Receive message:
```
go run consumer.go
```
### And finally:
---
<img src="./utils/go.png"
     alt="Gopher"
     style="width:100%" />
---
