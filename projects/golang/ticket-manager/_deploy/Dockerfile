# docker build -t roandocker/ticket-manager:0.0.1 Dockerfile .

FROM golang:stretch as builder

WORKDIR /opt/app/

RUN apt-get install git && \
	git clone https://github.com/codeandcode0x/istioAlltime.git && \
	cd istioAlltime/projects/golang/ticket-manager/ && \
    CGO_ENABLED=0 GOOS=linux go build -o /opt/app/ticket-backend main.go && \
    ls -all /opt/app/


FROM ubuntu:16.04 as prod

WORKDIR /root/

COPY --from=0 /opt/app/ticket-backend .

RUN ls -all /root/ && chmod 777 ticket-backend

CMD ["./ticket-backend"]