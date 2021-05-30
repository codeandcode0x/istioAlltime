FROM mariadb:10.2.38

WORKDIR /app
COPY ./initdata.sh  /app

# run command
CMD ["/bin/bash", "/app/initdata.sh"]
