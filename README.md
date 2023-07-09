# airlineapi

Weekend project to practice some Go programming.</br>

The application consists of one RESTful API written in Go (ticket_service directory)
 and another one that acts as a proxy (proxy directory), written in Python.</br>

The service provides airline tickets from https://duffel.com/.

## Requirements
1. [Go](https://go.dev/)
2. [Python](https://www.python.org/)
3. [Make](https://www.gnu.org/software/make/)
4. [Docker](https://www.docker.com/)
5. [direnv](https://direnv.net/)

## Configuration
Configuration is handled through .env files (see .env.example in proxy and ticket_service).</br>

## Build the application
A .envrc file (used by direnv) is needed if you want to build the application (or just set the env values manually in the Makefile). Change directory to `src` and build
 the application with `make build`.

## Run the application
Run the application with `make run`.

## Use the application
Once the application is running you can access the swagger page at http://localhost:${PROXY_PORT}/docs. You can also open up your browser at `http://localhost:24081/tickets?origin=ATH&destination=AUH&date=2023-08-15` (mind that the date should
be a future date) and watch the response.


## Shutdown the application
Shutdown the application  with `make shutdown`