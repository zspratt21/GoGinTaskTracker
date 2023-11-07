FROM amazon/dynamodb-local

WORKDIR /home/dynamodblocal

USER root

RUN mkdir -p /home/dynamodblocal/data

RUN chown -R dynamodblocal:dynamodblocal /home/dynamodblocal/data

USER dynamodblocal
