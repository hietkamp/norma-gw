FROM maven:3.6-jdk-8-slim AS builder

FROM flink:1.14.2-scala_2.11

# Copy sql-client script
COPY bin/* /opt/sql-client/
RUN chmod +x /opt/sql-client/start-sql-client.sh
RUN mkdir -p /opt/sql-client/lib

# Download connector libraries
RUN wget -P /opt/sql-client/lib https://repo.maven.apache.org/maven2/org/apache/flink/flink-sql-connector-elasticsearch7_2.12/1.14.2/flink-sql-connector-elasticsearch7_2.12-1.14.2.jar; \
    wget -P /opt/sql-client/lib/ https://repo.maven.apache.org/maven2/org/apache/flink/flink-sql-connector-kafka_2.12/1.14.2/flink-sql-connector-kafka_2.12-1.14.2.jar;

# Copy configuration
COPY conf/* /opt/flink/conf/

WORKDIR /opt/sql-client
ENV SQL_CLIENT_HOME /opt/sql-client

COPY docker-entrypoint.sh /
RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]