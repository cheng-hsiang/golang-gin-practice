url :https://medium.com/alberthg-docker-notes/docker%E7%AD%86%E8%A8%98-%E9%80%B2%E5%85%A5container-%E5%BB%BA%E7%AB%8B%E4%B8%A6%E6%93%8D%E4%BD%9C-postgresql-container-d221ba39aaec


docker create --name my-postgres -p 5432:5432 -e POSTGRES_PASSWORD=admin postgres
docker exec -it my-postgres psql -U postgres -c "create role mo with login password 'password';"
docker exec -it my-postgres psql -U postgres -c "create database ec_api_v1 owner mo"
