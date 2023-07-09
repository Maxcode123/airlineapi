APP_IMAGE=ticket-service-app
PROXY_IMAGE=ticket-service-app-proxy

include .envrc

build-service:
	docker build --tag ${APP_IMAGE} ticket_service 

run-service:
	docker run --network host -it -p 127.0.0.1:${API_PORT}:${API_PORT} ${APP_IMAGE}

build-proxy:
	docker build --tag ${PROXY_IMAGE} proxy

run-proxy:
	docker run --network host -it -p 127.0.0.1:${PROXY_PORT}:${PROXY_PORT} ${PROXY_IMAGE}
