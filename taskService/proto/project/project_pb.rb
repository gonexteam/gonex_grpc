# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: project/project.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("project/project.proto", :syntax => :proto3) do
    add_message "gonex_project.CreateProjectRequest" do
      optional :user_id, :string, 1
      optional :name, :string, 2
    end
    add_message "gonex_project.CreateTagRequest" do
      optional :user_id, :string, 1
      optional :name, :string, 2
      optional :project_id, :string, 3
    end
    add_message "gonex_project.GetProjectRequest" do
      optional :user_id, :string, 1
      optional :project_id, :string, 2
    end
    add_message "gonex_project.ProjectResponse" do
      optional :id, :string, 1
      optional :name, :string, 2
      repeated :tags, :message, 3, "gonex_project.TagResponse"
    end
    add_message "gonex_project.TagResponse" do
      optional :id, :string, 1
      optional :name, :string, 2
      optional :project_id, :string, 3
    end
  end
end

module GonexProject
  CreateProjectRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gonex_project.CreateProjectRequest").msgclass
  CreateTagRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gonex_project.CreateTagRequest").msgclass
  GetProjectRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gonex_project.GetProjectRequest").msgclass
  ProjectResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gonex_project.ProjectResponse").msgclass
  TagResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("gonex_project.TagResponse").msgclass
end
