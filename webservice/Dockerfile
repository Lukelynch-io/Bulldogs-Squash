FROM postgres:16

COPY testdata/setup_testcontainer.sql /docker-entrypoint-initdb.d/
RUN chown postgres:postgres /docker-entrypoint-initdb.d/setup_testcontainer.sql