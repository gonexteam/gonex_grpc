// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: protos/api/api.proto

package api

import (
	project "gonex.net/protos/project"
	task "gonex.net/protos/task"
	user "gonex.net/protos/user"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protos_api_api_proto protoreflect.FileDescriptor

var file_protos_api_api_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x61, 0x70, 0x69,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdd, 0x06, 0x0a, 0x03, 0x41, 0x50, 0x49, 0x12,
	0x61, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64, 0x65,
	0x6d, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x76,
	0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x3a,
	0x01, 0x2a, 0x12, 0x58, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x17, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x71, 0x0a, 0x0d,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x2e,
	0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12,
	0x72, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1f, 0x2e,
	0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12, 0x1c, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x5f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5f, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0xef, 0x01, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x12, 0x1b, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa6,
	0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x9f, 0x01, 0x12, 0x3e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x7d, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x61, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x7b, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x7d, 0x5a, 0x24, 0x12, 0x22, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x5a, 0x37,
	0x12, 0x35, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x7b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x6f, 0x6b, 0x65, 0x72, 0x36, 0x36, 0x36, 0x2f, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x64, 0x65, 0x6d, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_protos_api_api_proto_goTypes = []interface{}{
	(*user.RegisterRequest)(nil),         // 0: demo_user.RegisterRequest
	(*user.LoginRequest)(nil),            // 1: demo_user.LoginRequest
	(*project.CreateProjectRequest)(nil), // 2: demo_project.CreateProjectRequest
	(*project.GetProjectRequest)(nil),    // 3: demo_project.GetProjectRequest
	(*task.CreateTaskRequest)(nil),       // 4: demo_task.CreateTaskRequest
	(*task.UpdateTaskRequest)(nil),       // 5: demo_task.UpdateTaskRequest
	(*task.ListTasksRequest)(nil),        // 6: demo_task.ListTasksRequest
	(*user.UserResponse)(nil),            // 7: demo_user.UserResponse
	(*project.ProjectResponse)(nil),      // 8: demo_project.ProjectResponse
	(*task.TaskResponse)(nil),            // 9: demo_task.TaskResponse
	(*task.ListTasksResponse)(nil),       // 10: demo_task.ListTasksResponse
}
var file_protos_api_api_proto_depIdxs = []int32{
	0,  // 0: demo_api.API.RegisterUser:input_type -> demo_user.RegisterRequest
	1,  // 1: demo_api.API.LoginUser:input_type -> demo_user.LoginRequest
	2,  // 2: demo_api.API.CreateProject:input_type -> demo_project.CreateProjectRequest
	3,  // 3: demo_api.API.GetProject:input_type -> demo_project.GetProjectRequest
	4,  // 4: demo_api.API.CreateTask:input_type -> demo_task.CreateTaskRequest
	5,  // 5: demo_api.API.UpdateTask:input_type -> demo_task.UpdateTaskRequest
	6,  // 6: demo_api.API.ListTasks:input_type -> demo_task.ListTasksRequest
	7,  // 7: demo_api.API.RegisterUser:output_type -> demo_user.UserResponse
	7,  // 8: demo_api.API.LoginUser:output_type -> demo_user.UserResponse
	8,  // 9: demo_api.API.CreateProject:output_type -> demo_project.ProjectResponse
	8,  // 10: demo_api.API.GetProject:output_type -> demo_project.ProjectResponse
	9,  // 11: demo_api.API.CreateTask:output_type -> demo_task.TaskResponse
	9,  // 12: demo_api.API.UpdateTask:output_type -> demo_task.TaskResponse
	10, // 13: demo_api.API.ListTasks:output_type -> demo_task.ListTasksResponse
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_protos_api_api_proto_init() }
func file_protos_api_api_proto_init() {
	if File_protos_api_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_api_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_api_api_proto_goTypes,
		DependencyIndexes: file_protos_api_api_proto_depIdxs,
	}.Build()
	File_protos_api_api_proto = out.File
	file_protos_api_api_proto_rawDesc = nil
	file_protos_api_api_proto_goTypes = nil
	file_protos_api_api_proto_depIdxs = nil
}
