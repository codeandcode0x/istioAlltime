# docker build -t roandocker/ticket-frontend:0.0.1  --no-cache -f _docker/Dockerfile .
# docker build --target builder -t roandocker/ticket-manager:0.0.1 .

FROM maven:3.6.3-openjdk-8 as builder

WORKDIR /opt/app/


RUN apt-get install git && \
	git clone https://github.com/codeandcode0x/istioAlltime.git && \
	cd istioAlltime/projects/java/ticket-fronted/ && \
	mvn clean install && \
	cp target/ticket.fronted-0.0.1-SNAPSHOT.jar  /opt/app/ticket-fronted.jar


FROM openjdk:8-jdk-oraclelinux7 as prod

WORKDIR /root/

COPY --from=0 /opt/app/ticket-fronted.jar /root

RUN ls -all
RUN chmod 777 ticket-fronted.jar

CMD ["java -jar /root/ticket-fronted.jar"]