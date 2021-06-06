package com.msa.ticket.frontend.rpc.grpc.user.server;
import io.grpc.stub.StreamObserver;
import java.util.logging.Logger;

import com.msa.ticket.frontend.pojo.UserDatas;
import com.msa.ticket.frontend.rpc.grpc.user.server.UserRPCGrpc.UserRPCImplBase;

import org.springframework.web.client.RestTemplate;

/**
 * Server that manages startup/shutdown of a {@code Greeter} server.
 */
public class UserServer extends UserRPCImplBase {
  private static final Logger logger = Logger.getLogger(UserRPCGrpc.class.getName());

  @Override
    public void getAllUsers(UserMsgRequest request,StreamObserver<UserMsgReply> responseObserver) {
      String ticketHost = "http://localhost:8080";
      String url = ticketHost+"/api/users";
      RestTemplate restTemplate = new RestTemplate();
      UserDatas users = restTemplate.getForObject(url, UserDatas.class);
      logger.info(users.toString());
      
      UserMsgReply reply = UserMsgReply.newBuilder().setMessage(users.toString()).build();

      responseObserver.onNext(reply);
      responseObserver.onCompleted();
    }

}
