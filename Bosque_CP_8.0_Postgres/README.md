Bosque - Postgres
=================

This is one bit of "Floresta Sintática" in a SQL format and a small how-to get it running in a postgres database.

We had to made small changes to the original SQL file, that can be downloaded from [here](http://www.linguateca.pt/Floresta/ficheiros/gz/Bosque_CP_8.0.sql.gz), to get it working.

Donwload the ready-to-use Docker image
--------------------------------------

1. ```docker pull belimawr/florestasintatica```
2. ```docker run --name pg-floresta -e POSTGRES_PASSWORD=123mudar -d belimawr/florestasintatica```

How to load "bosque" on postgres using the "originais" SQL files (we just removed some escaped apostrophe characters)
---------------------------------------------------------------------------------------------------------------------

1. Build the Dockerfile: ```docker build -t postgres:floresta .```
2. Start a container: ```docker run --name pg-floresta -e POSTGRES_PASSWORD=123mudar -d postgres:floresta```
3. Create the database: ```docker exec pg-floresta createdb -U postgres -E ISO_8859_1 -T template0 --locale=pt_BR.ISO-8859-1 floresta```
4. Copy the SQL files to the container: ```docker cp ./*.sql pg-floresta:/tmp/```
5. Create all tables: ```docker exec pg-floresta psql floresta -U postgres -f /tmp/create_tables.sql```
6. Insert all data: ```docker exec pg-floresta psql floresta -U postgres -f /tmp/bosque.sql```
7. Wait a wee bit, it takes a while
8. Have fun!

How to load "bosque" using our dump (a sigle SQL file)
------------------------------------------------------

1. Build the Dockerfile: ```docker build -t postgres:floresta .```
2. Start a container: ```docker run --name pg-floresta -e POSTGRES_PASSWORD=123mudar -d postgres:floresta```
3. Create the database: ```docker exec pg-floresta createdb -U postgres -E ISO_8859_1 -T template0 --locale=pt_BR.ISO-8859-1 floresta```
4. Run our SQL dump script: ```docker exec pg-floresta psql floresta -U postgres -f /tmp/bosque_cp8_postgres_dump.sql```