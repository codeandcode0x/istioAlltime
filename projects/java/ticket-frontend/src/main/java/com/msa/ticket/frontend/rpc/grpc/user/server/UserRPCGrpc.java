package com.msa.ticket.frontend.rpc.grpc.user.server;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.38.0)",
    comments = "Source: user.proto")
public final class UserRPCGrpc {

  private UserRPCGrpc() {}

  public static final String SERVICE_NAME = "user.UserRPC";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest,
      com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply> getGetAllUsersMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAllUsers",
      requestType = com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest.class,
      responseType = com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest,
      com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply> getGetAllUsersMethod() {
    io.grpc.MethodDescriptor<com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest, com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply> getGetAllUsersMethod;
    if ((getGetAllUsersMethod = UserRPCGrpc.getGetAllUsersMethod) == null) {
      synchronized (UserRPCGrpc.class) {
        if ((getGetAllUsersMethod = UserRPCGrpc.getGetAllUsersMethod) == null) {
          UserRPCGrpc.getGetAllUsersMethod = getGetAllUsersMethod =
              io.grpc.MethodDescriptor.<com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest, com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAllUsers"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply.getDefaultInstance()))
              .setSchemaDescriptor(new UserRPCMethodDescriptorSupplier("GetAllUsers"))
              .build();
        }
      }
    }
    return getGetAllUsersMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static UserRPCStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserRPCStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserRPCStub>() {
        @java.lang.Override
        public UserRPCStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserRPCStub(channel, callOptions);
        }
      };
    return UserRPCStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static UserRPCBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserRPCBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserRPCBlockingStub>() {
        @java.lang.Override
        public UserRPCBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserRPCBlockingStub(channel, callOptions);
        }
      };
    return UserRPCBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static UserRPCFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<UserRPCFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<UserRPCFutureStub>() {
        @java.lang.Override
        public UserRPCFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new UserRPCFutureStub(channel, callOptions);
        }
      };
    return UserRPCFutureStub.newStub(factory, channel);
  }

  /**
   */
  public static abstract class UserRPCImplBase implements io.grpc.BindableService {

    /**
     */
    public void getAllUsers(com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest request,
        io.grpc.stub.StreamObserver<com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAllUsersMethod(), responseObserver);
    }

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
          .addMethod(
            getGetAllUsersMethod(),
            io.grpc.stub.ServerCalls.asyncUnaryCall(
              new MethodHandlers<
                com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest,
                com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply>(
                  this, METHODID_GET_ALL_USERS)))
          .build();
    }
  }

  /**
   */
  public static final class UserRPCStub extends io.grpc.stub.AbstractAsyncStub<UserRPCStub> {
    private UserRPCStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserRPCStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserRPCStub(channel, callOptions);
    }

    /**
     */
    public void getAllUsers(com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest request,
        io.grpc.stub.StreamObserver<com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAllUsersMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   */
  public static final class UserRPCBlockingStub extends io.grpc.stub.AbstractBlockingStub<UserRPCBlockingStub> {
    private UserRPCBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserRPCBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserRPCBlockingStub(channel, callOptions);
    }

    /**
     */
    public com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply getAllUsers(com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAllUsersMethod(), getCallOptions(), request);
    }
  }

  /**
   */
  public static final class UserRPCFutureStub extends io.grpc.stub.AbstractFutureStub<UserRPCFutureStub> {
    private UserRPCFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected UserRPCFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new UserRPCFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply> getAllUsers(
        com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAllUsersMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_GET_ALL_USERS = 0;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final UserRPCImplBase serviceImpl;
    private final int methodId;

    MethodHandlers(UserRPCImplBase serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_GET_ALL_USERS:
          serviceImpl.getAllUsers((com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgRequest) request,
              (io.grpc.stub.StreamObserver<com.msa.ticket.frontend.rpc.grpc.user.server.UserMsgReply>) responseObserver);
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

  private static abstract class UserRPCBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    UserRPCBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return com.msa.ticket.frontend.rpc.grpc.user.server.User.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("UserRPC");
    }
  }

  private static final class UserRPCFileDescriptorSupplier
      extends UserRPCBaseDescriptorSupplier {
    UserRPCFileDescriptorSupplier() {}
  }

  private static final class UserRPCMethodDescriptorSupplier
      extends UserRPCBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final String methodName;

    UserRPCMethodDescriptorSupplier(String methodName) {
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
      synchronized (UserRPCGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new UserRPCFileDescriptorSupplier())
              .addMethod(getGetAllUsersMethod())
              .build();
        }
      }
    }
    return result;
  }
}
