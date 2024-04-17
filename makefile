# Define variables for docker commands
FRONTEND_DIR=./frontend
BACKEND_DIR=./backend
DOCKER_COMPOSE_FILE=docker-compose.yml
FRONTEND_IMAGE_NAME=vue-frontend
BACKEND_IMAGE_NAME=go-backend

# Build the frontend and backend Docker images
build:
	docker build -t $(FRONTEND_IMAGE_NAME) $(FRONTEND_DIR)
	docker build -t $(BACKEND_IMAGE_NAME) $(BACKEND_DIR)

# Run docker-compose up to start all services
up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up

# Run docker-compose up in detached mode
up-detached:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop and remove containers, networks, images, and volumes
down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Build and run the frontend service only
frontend: 
	docker build -t $(FRONTEND_IMAGE_NAME) $(FRONTEND_DIR)
	docker run -p 5173:5173 $(FRONTEND_IMAGE_NAME)

# Build and run the backend service only
backend: 
	docker build -t $(BACKEND_IMAGE_NAME) $(BACKEND_DIR)
	docker run -p 8081:8081 $(BACKEND_IMAGE_NAME)

# Remove all unused containers, networks, images (both dangling and unreferenced), and optionally, volumes
clean:
	docker system prune -a

# Access the command line of the running frontend container
frontend-shell:
	docker exec -it $(shell docker ps -qf "ancestor=$(FRONTEND_IMAGE_NAME)") sh

# Access the command line of the running backend container
backend-shell:
	docker exec -it $(shell docker ps -qf "ancestor=$(BACKEND_IMAGE_NAME)") sh

.PHONY: build up up-detached down frontend backend clean frontend-shell backend-shell
