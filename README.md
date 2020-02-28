# bookstore_users-api
A microservice to govern the user domain as specified in the tutorial <a href="https://www.udemy.com/course/golang-how-to-design-and-build-rest-microservices-in-go">Microservices in Go - Part 2</a>
## Getting Started

The microservice should run out of the box once you set the environment variables

```
* mysql_users_username=[username]
* mysql_users_password=[password]
* mysql_users_host=127.0.0.1:3306
* mysql_users_schema=users_db
```

### Running

With the environment valiables in place, run:
```
/bookstore_users-api/main.go
```

Check whether the service s running by pinging:

```
http://localhost:8080/ping
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests

TBD

## License

Same as: https://github.com/federicoleon/bookstore_users-api/
