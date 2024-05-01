
# Microservice with golang

Make Microservice Backend
- Broker-service: send task to right Microservice
- Mail-service: send-Mail
- logger-service: seve some infomation which broker need to save in mongoDB
- Authentication-service: check use already registered?
- Lister: Listen remaining task in RabbitMQ
- front-end: Test Microservice from browser


## Installation

Start back-end microservice

```bash
  cd project
  make up
```

Test microservice in front-end

Local machine:
```bash
  cd front-end
  make frontend
```
or container:
```bash
  cd project
  make build_front
```
    
## 🛠 Skills
* Golang web framework with CHI
* API: http, rgc, gRPC
* Database: MongoDB, PostgresDB, RabbitMQ
* Mail: MailHog

