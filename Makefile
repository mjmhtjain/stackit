.PHONY: build run rerun clean

# Build the application
build:
	docker-compose build

# Run the application
run:
	docker-compose up -d

# Rerun the application (stop, remove, rebuild and run)
rerun:
	docker-compose down
	docker-compose build --no-cache
	docker-compose up

# Clean up containers, images, and volumes
clean:
	docker-compose down -v
	docker system prune -f

# Stop the application
stop:
	docker-compose down

# View logs
logs:
	docker-compose logs -f 