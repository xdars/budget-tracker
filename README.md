Slow work-in-progress budget tracker built on microservices using gRPC, for fun

The way the connection to postgres is made implies that `auth-service(or whatever service that is)/.env` file has to look like this:

```POSTGRES_URL="host=localhost port=54320 user=postgres password=my_password dbname=name sslmode=disable"```