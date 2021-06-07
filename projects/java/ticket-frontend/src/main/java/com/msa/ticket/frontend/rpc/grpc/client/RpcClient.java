package com.msa.ticket.frontend.rpc.grpc.client;

import java.util.concurrent.TimeUnit;
import java.util.logging.Level;
import java.util.logging.Logger;

import com.msa.ticket.frontend.rpc.grpc.movie.server.MovieRPCGrpc;
import com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply ;
import com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest;

import com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply;
import com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest;
import com.msa.ticket.frontend.rpc.grpc.user.server.UserRPCGrpc;

import io.grpc.Channel;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;

public class RpcClient {
  private static final Logger logger = Logger.getLogger(RpcClient.class.getName());

  private final MovieRPCGrpc.MovieRPCBlockingStub movieStub;
  private final UserRPCGrpc.UserRPCBlockingStub userStub;

  /** Construct client for accessing HelloWorld server using the existing channel. */
  public RpcClient(Channel channel) {
    movieStub = MovieRPCGrpc.newBlockingStub(channel);
    userStub = UserRPCGrpc.newBlockingStub(channel);
  }

  public void getMovies(Integer count) {
    logger.info("list count:  " + count + " ...");
    MovieMsgRequest request = MovieMsgRequest.newBuilder().setCount(count).build();
    MovieMsgReply response;
    try {
      response = movieStub.getAllMovies(request);
    } catch (StatusRuntimeException e) {
      logger.log(Level.WARNING, "RPC failed: {0}", e.getStatus());
      return;
    }
    logger.info("movies: " + response.getMessage());
  }


  public void getUsers(Integer count) {
    logger.info("list count:  " + count + " ...");
    UserMsgRequest request = UserMsgRequest.newBuilder().setCount(count).build();
    UserMsgReply response;
    try {
      response = userStub.getAllUsers(request);
    } catch (StatusRuntimeException e) {
      logger.log(Level.WARNING, "RPC failed: {0}", e.getStatus());
      return;
    }
    logger.info("users: " + response.getMessage());
  }

  /**
   * Greet server. If provided, the first element of {@code args} is the name to use in the
   * greeting. The second argument is the target server.
   */
  public static void main(String[] args) throws Exception {
    Integer count = 3;
    // Access a service running on the local machine on port 50051
    String target = "localhost:20153";
    // Allow passing in the count and target strings as command line arguments
    if (args.length > 0) {
      if ("--help".equals(args[0])) {
        System.err.println("Usage: [name [target]]");
        System.err.println("");
        System.err.println("  name    The name you wish to be greeted by. Defaults to " + count);
        System.err.println("  target  The server to connect to. Defaults to " + target);
        System.exit(1);
      }
    //   count = args[0];
    }
    if (args.length > 1) {
      target = args[1];
    }

    // Create a communication channel to the server, known as a Channel. Channels are thread-safe
    // and reusable. It is common to create channels at the beginning of your application and reuse
    // them until the application shuts down.
    ManagedChannel channel = ManagedChannelBuilder.forTarget(target)
        // Channels are secure by default (via SSL/TLS). For the example we disable TLS to avoid
        // needing certificates.
        .usePlaintext()
        .build();
    try {
      RpcClient client = new RpcClient(channel);
      client.getMovies(count);
      client.getUsers(count);
    } finally {
      // ManagedChannels use resources like threads and TCP connections. To prevent leaking these
      // resources the channel should be shut down when it will no longer be used. If it may be used
      // again leave it running.
      channel.shutdownNow().awaitTermination(5, TimeUnit.SECONDS);
    }
  }
}
