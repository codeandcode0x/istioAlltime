package com.msa.ticket.frontend;
import java.io.IOException;
import java.util.concurrent.TimeUnit;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import java.util.logging.Logger;
import com.msa.ticket.frontend.rpc.grpc.movie.server.MovieServer;
import com.msa.ticket.frontend.rpc.grpc.user.server.UserServer;

public class RpcServer {
    private static final Logger logger = Logger.getLogger(RpcServer.class.getName());

    private Server server;

    public void Start() throws IOException {
      /* The port on which the server should run */
      int port = 20154;
      server = ServerBuilder.forPort(port)
          .addService(new MovieServer())
          .addService(new UserServer())
          .build()
          .start();
      logger.info("Server started, listening on " + port);
      Runtime.getRuntime().addShutdownHook(new Thread() {
        @Override
        public void run() {
          // Use stderr here since the logger may have been reset by its JVM shutdown hook.
          System.err.println("*** shutting down gRPC server since JVM is shutting down");
          try {
            RpcServer.this.Stop();
          } catch (InterruptedException e) {
            e.printStackTrace(System.err);
          }
          System.err.println("*** server shut down");
        }
      });
    }
  
    private void Stop() throws InterruptedException {
      if (server != null) {
        server.shutdown().awaitTermination(30, TimeUnit.SECONDS);
      }
    }
  
    /**
     * Await termination on the main thread since the grpc library uses daemon threads.
     */
    public void BlockUntilShutdown() throws InterruptedException {
      if (server != null) {
        server.awaitTermination();
      }
    }  
}
