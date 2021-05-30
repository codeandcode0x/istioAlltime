FROM golang:stretch as builder

WORKDIR /opt/app/

RUN rm -rf istioAlltime && \
	apt-get install git && \
	git clone https://github.com/codeandcode0x/istioAlltime.git


FROM mariadb:10.2.38 as prod

WORKDIR /app
COPY ./initdata.sh  /app
COPY --from=0 /opt/app/istioAlltime  /app/istioAlltime/

# run command
CMD ["/bin/bash", "/app/initdata.sh istioAlltime/projects/data/sql  ticket"]
