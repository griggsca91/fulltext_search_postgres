#!/bin/bash


# reset the passwor via this command
# docker exec -it fulltext_search_postgres-elasticsearch01-1 /usr/share/elasticsearch/bin/elasticsearch-reset-password -u elastic

# Get the certificate with this command
docker cp fulltext_search_postgres-elasticsearch01-1:/usr/share/elasticsearch/config/certs/http_ca.crt ..

