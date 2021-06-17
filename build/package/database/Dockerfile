FROM postgres:latest

# Custom initialization script
COPY ../../../scripts/init_db.sh   /docker-entrypoint-initdb.d/init_db.sh
RUN chmod a+r /docker-entrypoint-initdb.d/*