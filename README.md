DB:
Start docker db: docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
Connect to smth: docker exec -it container_id /bin/bash
Use docker db_postgres: psql -U postgres