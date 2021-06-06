package com.msa.ticket.frontend.rpc.grpc.movie.server;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.38.0)",
    comments = "Source: movie.proto")
public final class MovieRPCGrpc {

  private MovieRPCGrpc() {}

  public static final String SERVICE_NAME = "movie.MovieRPC";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest,
      com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply> getGetAllMoviesMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAllMovies",
      requestType = com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest.class,
      responseType = com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest,
      com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply> getGetAllMoviesMethod() {
    io.grpc.MethodDescriptor<com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest, com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply> getGetAllMoviesMethod;
    if ((getGetAllMoviesMethod = MovieRPCGrpc.getGetAllMoviesMethod) == null) {
      synchronized (MovieRPCGrpc.class) {
        if ((getGetAllMoviesMethod = MovieRPCGrpc.getGetAllMoviesMethod) == null) {
          MovieRPCGrpc.getGetAllMoviesMethod = getGetAllMoviesMethod =
              io.grpc.MethodDescriptor.<com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest, com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAllMovies"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply.getDefaultInstance()))
              .setSchemaDescriptor(new MovieRPCMethodDescriptorSupplier("GetAllMovies"))
              .build();
        }
      }
    }
    return getGetAllMoviesMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static MovieRPCStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MovieRPCStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MovieRPCStub>() {
        @java.lang.Override
        public MovieRPCStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MovieRPCStub(channel, callOptions);
        }
      };
    return MovieRPCStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static MovieRPCBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MovieRPCBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MovieRPCBlockingStub>() {
        @java.lang.Override
        public MovieRPCBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MovieRPCBlockingStub(channel, callOptions);
        }
      };
    return MovieRPCBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static MovieRPCFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<MovieRPCFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<MovieRPCFutureStub>() {
        @java.lang.Override
        public MovieRPCFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new MovieRPCFutureStub(channel, callOptions);
        }
      };
    return MovieRPCFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class MovieRPCImplBase implements io.grpc.BindableService {

    /**
     */
    public void getAllMovies(com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest request,
        io.grpc.stub.StreamObserver<com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAllMoviesMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetAllMoviesMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest,
                com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply>(
                  this, METHODID_GET_ALL_MOVIES)))
          .build();
    }
  }

  /**
   */
  public static final class MovieRPCStub extends io.grpc.stub.AbstractAsyncStub<MovieRPCStub> {
    private MovieRPCStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MovieRPCStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MovieRPCStub(channel, callOptions);
    }

    /**
     */
    public void getAllMovies(com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest request,
        io.grpc.stub.StreamObserver<com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAllMoviesMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class MovieRPCBlockingStub extends io.grpc.stub.AbstractBlockingStub<MovieRPCBlockingStub> {
    private MovieRPCBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MovieRPCBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MovieRPCBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply getAllMovies(com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAllMoviesMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class MovieRPCFutureStub extends io.grpc.stub.AbstractFutureStub<MovieRPCFutureStub> {
    private MovieRPCFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected MovieRPCFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new MovieRPCFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply> getAllMovies(
        com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAllMoviesMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_ALL_MOVIES = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final MovieRPCImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(MovieRPCImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_ALL_MOVIES:
          serviceImpl.getAllMovies((com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgRequest) request,
              (io.grpc.stub.StreamObserver<com.msa.ticket.frontend.rpc.grpc.movie.server.MovieMsgReply>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  private static abstract class MovieRPCBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    MovieRPCBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.msa.ticket.frontend.rpc.grpc.movie.server.Movie.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("MovieRPC");
    }
  }

  private static final class MovieRPCFileDescriptorSupplier
      extends MovieRPCBaseDescriptorSupplier {
    MovieRPCFileDescriptorSupplier() {}
  }

  private static final class MovieRPCMethodDescriptorSupplier
      extends MovieRPCBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    MovieRPCMethodDescriptorSupplier(String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (MovieRPCGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new MovieRPCFileDescriptorSupplier())
              .addMethod(getGetAllMoviesMethod())
              .build();
        }
      }
    }
    return result;
  }
}
