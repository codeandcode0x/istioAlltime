package com.msa.ticket.frontend.rpc.grpc.movie.server;
import io.grpc.stub.StreamObserver;
import java.util.logging.Logger;

import com.msa.ticket.frontend.pojo.MovieDatas;

import org.springframework.web.client.RestTemplate;
// import com.msa.ticket.frontend.pojo.MovieDatas;

/**
 * Server that manages startup/shutdown of a {@code Greeter} server.
 */
public class MovieServer extends MovieRPCGrpc.MovieRPCImplBase {
  private static final Logger logger = Logger.getLogger(MovieRPCGrpc.class.getName());

  @Override
    public void getAllMovies(MovieMsgRequest request,StreamObserver<MovieMsgReply> responseObserver) {
      String ticketHost = "http://localhost:8080";
      String url = ticketHost+"/api/movies";
      RestTemplate restTemplate = new RestTemplate();
      MovieDatas movies = restTemplate.getForObject(url, MovieDatas.class);
      logger.info(movies.toString());
      
      MovieMsgReply reply = MovieMsgReply.newBuilder().setMessage(movies.toString()).build();

      responseObserver.onNext(reply);
      responseObserver.onCompleted();
    }

}
