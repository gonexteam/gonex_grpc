// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var user_pb = require('./user_pb.js');

function serialize_gonex_user_GetUserRequest(arg) {
  if (!(arg instanceof user_pb.GetUserRequest)) {
    throw new Error('Expected argument of type gonex_user.GetUserRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gonex_user_GetUserRequest(buffer_arg) {
  return user_pb.GetUserRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gonex_user_LoginRequest(arg) {
  if (!(arg instanceof user_pb.LoginRequest)) {
    throw new Error('Expected argument of type gonex_user.LoginRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gonex_user_LoginRequest(buffer_arg) {
  return user_pb.LoginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gonex_user_RegisterRequest(arg) {
  if (!(arg instanceof user_pb.RegisterRequest)) {
    throw new Error('Expected argument of type gonex_user.RegisterRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gonex_user_RegisterRequest(buffer_arg) {
  return user_pb.RegisterRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gonex_user_UserResponse(arg) {
  if (!(arg instanceof user_pb.UserResponse)) {
    throw new Error('Expected argument of type gonex_user.UserResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gonex_user_UserResponse(buffer_arg) {
  return user_pb.UserResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gonex_user_VerifyRequest(arg) {
  if (!(arg instanceof user_pb.VerifyRequest)) {
    throw new Error('Expected argument of type gonex_user.VerifyRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gonex_user_VerifyRequest(buffer_arg) {
  return user_pb.VerifyRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_gonex_user_VerifyResponse(arg) {
  if (!(arg instanceof user_pb.VerifyResponse)) {
    throw new Error('Expected argument of type gonex_user.VerifyResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_gonex_user_VerifyResponse(buffer_arg) {
  return user_pb.VerifyResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var UserSvcService = exports.UserSvcService = {
  register: {
    path: '/gonex_user.UserSvc/register',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.RegisterRequest,
    responseType: user_pb.UserResponse,
    requestSerialize: serialize_gonex_user_RegisterRequest,
    requestDeserialize: deserialize_gonex_user_RegisterRequest,
    responseSerialize: serialize_gonex_user_UserResponse,
    responseDeserialize: deserialize_gonex_user_UserResponse,
  },
  login: {
    path: '/gonex_user.UserSvc/login',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.LoginRequest,
    responseType: user_pb.UserResponse,
    requestSerialize: serialize_gonex_user_LoginRequest,
    requestDeserialize: deserialize_gonex_user_LoginRequest,
    responseSerialize: serialize_gonex_user_UserResponse,
    responseDeserialize: deserialize_gonex_user_UserResponse,
  },
  verify: {
    path: '/gonex_user.UserSvc/verify',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.VerifyRequest,
    responseType: user_pb.VerifyResponse,
    requestSerialize: serialize_gonex_user_VerifyRequest,
    requestDeserialize: deserialize_gonex_user_VerifyRequest,
    responseSerialize: serialize_gonex_user_VerifyResponse,
    responseDeserialize: deserialize_gonex_user_VerifyResponse,
  },
  getUser: {
    path: '/gonex_user.UserSvc/getUser',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.GetUserRequest,
    responseType: user_pb.VerifyResponse,
    requestSerialize: serialize_gonex_user_GetUserRequest,
    requestDeserialize: deserialize_gonex_user_GetUserRequest,
    responseSerialize: serialize_gonex_user_VerifyResponse,
    responseDeserialize: deserialize_gonex_user_VerifyResponse,
  },
};

exports.UserSvcClient = grpc.makeGenericClientConstructor(UserSvcService);
