# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: user/user.proto for package 'gonex_user'

require 'grpc'
require 'user/user_pb'

module GonexUser
  module UserSvc
    class Service

      include ::GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gonex_user.UserSvc'

      rpc :register, ::GonexUser::RegisterRequest, ::GonexUser::UserResponse
      rpc :login, ::GonexUser::LoginRequest, ::GonexUser::UserResponse
      rpc :verify, ::GonexUser::VerifyRequest, ::GonexUser::VerifyResponse
      rpc :getUser, ::GonexUser::GetUserRequest, ::GonexUser::VerifyResponse
    end

    Stub = Service.rpc_stub_class
  end
end
