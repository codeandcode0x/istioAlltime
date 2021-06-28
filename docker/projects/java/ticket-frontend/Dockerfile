# docker build -t roandocker/ticket-frontend:0.0.1  --no-cache -f _docker/Dockerfile .
# docker build --target builder -t roandocker/ticket-manager:0.0.1 .

FROM maven:3.6.3-openjdk-8 as builder

WORKDIR /opt/app/

COPY docker/tools/java-tracing-agent/opentracing-specialagent-1.7.5-SNAPSHOT.jar /opt/app/opentracing-specialagent-1.7.5-SNAPSHOT.jar

RUN apt-get install git && \
	git clone https://github.com/codeandcode0x/istioAlltime.git && \
	cd istioAlltime/projects/java/ticket-frontend/ && \
	mvn clean install && \
	cp target/ticket.frontend-0.0.1-SNAPSHOT.jar  /opt/app/ticket-frontend.jar


FROM openjdk:8-jdk-oraclelinux7 as prod

WORKDIR /root/

COPY --from=0 /opt/app/opentracing-specialagent-1.7.5-SNAPSHOT.jar /root
COPY --from=0 /opt/app/ticket-frontend.jar /root

RUN chmod 777 ticket-frontend.jar && \
	chmod 777 opentracing-specialagent-1.7.5-SNAPSHOT.jar

# health check
# HEALTHCHECK --interval=5s --timeout=3s --retries=3 \
#   CMD curl -fs http://localhost:8090 || exit 1

CMD ["/bin/sh", "-c", "java -javaagent:/root/opentracing-specialagent-1.7.5-SNAPSHOT.jar -DJAEGER_SERVICE_NAME=\${SERVICE_NAME:-'UNKNOWN_SERVICE'} -DJAEGER_SAMPLER_TYPE=\${TRACE_SAMPLER_TYPE:-'probabilistic'} -DJAEGER_SAMPLER_PARAM=\${TRACE_SAMPLER_PARAM:-'0.01'} -DJAEGER_ENDPOINT=\${TRACE_ENDPOINT:-'http://jaeger-collector:14268/api/traces'} \${TRACE_DSA_EXTENDS:-}  -jar /root/ticket-frontend.jar"]