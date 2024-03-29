This is the first Go Hello World program.

It's a basic Go Hello World application that connects to a Postgres database. You can start the Postgres database using the Docker Compose file present in the docker folder.

We're using the Echo Library to start the server. You can find the server code in the file server.go located in src/server.

Additionally, we're using GORM to create the database client. You can find the database client code in the file client.go located in src/database.

To start the application, run the following command from the main directory of the application: 
```sh
go run src/main.go
```

<img width="445" alt="image" src="https://github.com/saravanastar/go-hello-world/assets/14228896/ac53db46-7fe8-4149-a900-3132797e09fc">

HealthCheck

<img width="577" alt="image" src="https://github.com/saravanastar/go-hello-world/assets/14228896/0ef1f5cb-7267-4fcd-be60-7514a2040ad0">
