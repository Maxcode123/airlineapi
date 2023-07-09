APP_IMAGE=ticket-service-app
PROXY_IMAGE=ticket-service-app-proxy

include .envrc

build: build-service build-proxy

run: run-service run-proxy

build-service:
	docker build --tag ${APP_IMAGE} ticket_service 

run-service:
	docker run --network host --name app -dp 127.0.0.1:${API_PORT}:${API_PORT} ${APP_IMAGE}

build-proxy:
	docker build --tag ${PROXY_IMAGE} proxy

run-proxy:
	docker run --network host --name proxy -dp 127.0.0.1:${PROXY_PORT}:${PROXY_PORT} ${PROXY_IMAGE}

shutdown:
	docker stop app proxy
	docker rm app proxy