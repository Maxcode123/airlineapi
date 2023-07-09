APP_IMAGE=ticket-service-app

include .envrc

build-service:
	docker build --tag ${APP_IMAGE} --build-arg API_PORT=${API_PORT} ticket_service 

run-service:
	docker run --network host -it -p 127.0.0.1:${API_PORT}:${API_PORT} ${APP_IMAGE}

build-proxy:

run-proxy: