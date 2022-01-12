CREATE TABLE requests_received (
    request_id STRING,
    ts TIMESTAMP(3),
    request STRING,
    client_id STRING,
    client_ip STRING,
    controller_id STRING,
    controller_host STRING,
    conversation_id STRING,
    proctime AS PROCTIME(),   -- generates processing-time attribute using computed column
    WATERMARK FOR ts AS ts - INTERVAL '5' SECOND  -- defines watermark on ts column, marks ts as event-time attribute
) WITH (
    'connector' = 'kafka',  -- using kafka connector
    'topic' = 'requests_received',  -- kafka topic
    'scan.startup.mode' = 'earliest-offset',  -- reading from the beginning
    'properties.bootstrap.servers' = 'kafka2:9094',  -- kafka broker address
    'format' = 'json'  -- the data format is json
);