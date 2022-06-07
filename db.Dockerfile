# start with base image
FROM mysql:5.7

# import data into container
# All scripts in docker-entrypoint-initdb.d/ are automatically executed during container startup
COPY ./db/ /docker-entrypoint-initdb.d/