# golang-destributed
Golang distributed application with RabbitMq and web-socket. 

* [amqp] - https://github.com/krishanthisera/amqp.git
* [rabbitmq] - https://hub.docker.com/_/rabbitmq
```sh
$ docker run -d --hostname my-rabbit --name some-rabbit -p 8080:15672 rabbitmq:3-management
$ docker ps
```
