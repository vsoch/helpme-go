# HelpMe

This is eventually going to work and be a thing, for now it's learning GoLang for VanessaSaurus!

## Usage

You can either use GoLang locally, or build a Docker image with it included. For this example we will use Docker to avoid installation of additional dependencies.

```
docker build -t vanessa/helpme .
```

For the database, a more robust application might use docker-compose, but we are just going to start a postgres container, map it's data folder to the host, and then connect to it.

```
# Database
mkdir -p /tmp/data
docker run --name helpme-postgres --env POSTGRES_PASSWORD=helpme \
                                   --env POSTGRES_USER=helpme  \
                                   --env POSTGRES_DB=db --publish 5432:5432 \
                                      -v /tmp/data:/var/run/postgresql \
                                      -d postgres
```

Now we can run the Docker container (note this isn't mapped correctly yet)

```
# Run (Docker)
docker run -p 80:80 -v /tmp/data:/tmp/data vanessa/helpme -db-connect "host=/tmp/data user=helpme dbname=db sslmode=disable"
```

and this does work - we can compile the code with `go get` to make a binary `helpme` under the `GOROOT` and run the helpme app natively:

```
#     (Local)
helpme -db-connect "host=/tmp/data user=helpme dbname=db sslmode=disable"
```
