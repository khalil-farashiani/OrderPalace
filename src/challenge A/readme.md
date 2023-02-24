# Challenge A

This project consists of two applications, namely sender and receiver, which communicate with each other via Redis publish and subscribe mechanism. The applications are developed using the DDD (Domain-Driven Design) approach and are built with various technologies such as Mysql, mux gorrila, net/htttp, gorm, gomock, and faker.

## Technologies Used

- Redis - Publish and Subscribe messaging system
- Mysql - Database
- Mux gorrila - Router
- net/htttp - HTTP client and server
- Gorm - ORM
- Gomock - Mocking Library
- Faker - Creating fake data for unit testing

# Installation
for starting the project ensure that docker and docker compose are installed on your system.

finally just write these command on your system

`for the first time first step may get long time for downloading image file from registry`

- ./build
- ./run

`make sure that you are in the challenge A directory`

you can change the sample .env file for your own config database and redis