## DemonstrationServiceL0

# Запуск приложения в docker
docker compose build
docker compose up -d

# Вывод данных доступен по 
http://localhost:8081/id
Где id значения от 1

# База данных 
demo_servis_database

POSTGRES_DB: demo_service_database
POSTGRES_USER: Pavel
POSTGRES_PASSWORD: qwerty
POSTGRES_HOST: db
POSTGRES_PORT: 5432

# nats-streaming
clusterID: nats-cluster
ports: 4222
       8222
