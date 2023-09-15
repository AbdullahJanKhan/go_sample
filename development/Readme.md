# Delopment Dependency

Store your development dependencies in this directory, e.g. docker files, docker compose file, sample configurations file.

# Using Env File with Postgres image:

You will need to update the docker-compose file with the following code in postgres
```docker
 postgres:
    container_name: postgres
    image: postgres:latest     
    env_file:
      - './.env'
    ports:
      - "5432:5432"
```
The Env file must look like
```env
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=db

```
