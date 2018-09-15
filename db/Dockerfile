FROM postgres:9.6-alpine

ENV POSTGRES_DB shopee
# RUN [[ -e "pg_dump.backup" ]] && pg_restore pg_dump.backup > pg_dump.sql
# Preload database on init
COPY ./db/init.sql . /docker-entrypoint-initdb.d/