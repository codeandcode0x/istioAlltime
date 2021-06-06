package com.msa.ticket.frontend;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;
import org.springframework.http.HttpStatus;
import org.springframework.http.client.SimpleClientHttpRequestFactory;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.client.RestTemplate;

import java.io.IOException;

import javax.servlet.RequestDispatcher;
import javax.servlet.http.HttpServletRequest;
import java.util.logging.Logger;

@SpringBootApplication
//@EnableScheduling
public class Application {
	private static final Logger logger = Logger.getLogger(Application.class.getName());
	public static void main(String[] args) {
		SpringApplication.run(Application.class, args);
		try {
			RpcServerStart();
		} catch (InterruptedException e) {
			logger.info(e.getMessage());
		} catch (IOException e ) {
			logger.info(e.getMessage());
		}
	}

	private static void RpcServerStart() throws IOException, InterruptedException {
		final RpcServer Server = new RpcServer();
		Server.Start();
		Server.BlockUntilShutdown();
	}

	@Bean
	RestTemplate restTemplate(){
		SimpleClientHttpRequestFactory requestFactory = new SimpleClientHttpRequestFactory();
		requestFactory.setConnectTimeout(1000);
		requestFactory.setReadTimeout(1000);

		RestTemplate restTemplate = new RestTemplate(requestFactory);
		return restTemplate;
	}

	@RequestMapping("/error")
	public String handleError(HttpServletRequest request) {
		Object status = request.getAttribute(RequestDispatcher.ERROR_STATUS_CODE);
		if (status != null) {
			Integer statusCode = Integer.valueOf(status.toString());

			if(statusCode == HttpStatus.NOT_FOUND.value()) {
				return "error-404";
			} else if(statusCode == HttpStatus.INTERNAL_SERVER_ERROR.value()) {
				return "error-500";
			}
		}
		return "error";
	}

}
