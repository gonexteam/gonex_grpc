# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: user/user.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("user/user.proto", :syntax => :proto3) do
    add_message "gonex_user.VerifyRequest" do
      optional :token, :string, 1
    end
    add_message "gonex_user.GetUserRequest" do
      optional :user_id, :string, 1
    end
    add_message "gonex_user.LoginRequest" do
      optional :email, :string, 1
      optional :password, :string, 2
    end
    add_message "gonex_user.RegisterRequest" do
      optional :name, :string, 1
      optional :email, :string, 2
      optional :password, :string, 3
    end
    add_message "gonex_user.UserResponse" do
      optional :id, :string, 1
      optional :name, :string, 2
      optional :email, :string, 3
      optional :token, :string, 4
    end
    add_message "gonex_user.VerifyResponse" do
      optional :id, :string, 1
      optional :name, :string, 2
      optional :email, :string, 3
    end
  end
end

module GonexUser
  VerifyRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gonex_user.VerifyRequest").msgclass
  GetUserRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gonex_user.GetUserRequest").msgclass
  LoginRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gonex_user.LoginRequest").msgclass
  RegisterRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gonex_user.RegisterRequest").msgclass
  UserResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gonex_user.UserResponse").msgclass
  VerifyResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gonex_user.VerifyResponse").msgclass
end
