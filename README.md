# DOCKER-COMPOSE COMMANDS

// поднимаем редис
docker compose up -d redis

// проверяем
docker exec -it roomtime-redis redis-cli ping
