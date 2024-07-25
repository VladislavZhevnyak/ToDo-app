1. docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
2. docker exec -it container_id /bin/bash
3. psql -U postgres