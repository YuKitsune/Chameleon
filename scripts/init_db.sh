#!/bin/bash
set -e

POSTGRES="psql --username ${POSTGRES_USER}"

echo "Creating database role: ${API_USER}"
$POSTGRES <<-EOSQL
CREATE USER ${API_USER} WITH ENCRYPTED PASSWORD '${API_PASSWORD}';
EOSQL

echo "Creating database: ${API_DB}"
$POSTGRES <<-EOSQL
CREATE DATABASE ${API_DB} OWNER '${API_USER}';
EOSQL
