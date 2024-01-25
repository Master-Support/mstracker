#!/bin/bash
psql -v ON_ERROR_STOP=1 --username appms --dbname mstracker <<-EOSQL
    CREATE DATABASE mstracker;
EOSQL
