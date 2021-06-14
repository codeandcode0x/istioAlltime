# docker build -t roandocker/ticket-frontend:0.0.1  --no-cache -f _docker/Dockerfile .
# docker build --target builder -t roandocker/ticket-manager:0.0.1 .

FROM maven:3.6.3-openjdk-8 as builder

WORKDIR /opt/app/


RUN apt-get install git && \
	git clone https://github.com/codeandcode0x/istioAlltime.git && \
	cd istioAlltime/projects/java/ticket-frontend/ && \
	mvn clean install && \
	cp target/ticket.frontend-0.0.1-SNAPSHOT.jar  /opt/app/ticket-frontend.jar


FROM openjdk:8-jdk-oraclelinux7 as prod

WORKDIR /root/

COPY --from=0 /opt/app/ticket-frontend.jar /root

RUN chmod 777 ticket-frontend.jar

# health check
# HEALTHCHECK --interval=5s --timeout=3s --retries=3 \
#   CMD curl -fs http://localhost:8090 || exit 1

CMD ["/bin/sh", "-c", "java -jar /root/ticket-frontend.jar"]