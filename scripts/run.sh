#!/bin/sh

# Clear out any cached files and images from each individual rerun deployment
docker system prune -a -f

# run docker compose
docker-compose build --no-cache
docker-compose up -d 