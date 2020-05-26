FROM rabbitmq:3.5.3-management

COPY rabbitmq-cluster /usr/local/bin/
COPY entrypoint.sh /

RUN chmod +x /usr/local/bin/rabbitmq-cluster
RUN chmod +x entrypoint.sh

EXPOSE 5672 15672 25672 4369 9100 9101 9102 9103 9104 9105
ENTRYPOINT ["/entrypoint.sh"]
CMD ["rabbitmq-cluster"]
