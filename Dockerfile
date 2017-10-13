FROM consul:latest

RUN mkdir /scripts

COPY handler /scripts/

COPY config.json /consul/config/

RUN chmod +x /scripts/handler
