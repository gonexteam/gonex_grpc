# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: project.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='project.proto',
  package='demo_project',
  syntax='proto3',
  serialized_options=b'Z4github.com/Joker666/microservice-demo/protos/project',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\rproject.proto\x12\x0c\x64\x65mo_project\"5\n\x14\x43reateProjectRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"E\n\x10\x43reateTagRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x12\n\nproject_id\x18\x03 \x01(\t\"8\n\x11GetProjectRequest\x12\x0f\n\x07user_id\x18\x01 \x01(\t\x12\x12\n\nproject_id\x18\x02 \x01(\t\"T\n\x0fProjectResponse\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\'\n\x04tags\x18\x03 \x03(\x0b\x32\x19.demo_project.TagResponse\";\n\x0bTagResponse\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x12\n\nproject_id\x18\x03 \x01(\t2\xf6\x01\n\nProjectSvc\x12R\n\rcreateProject\x12\".demo_project.CreateProjectRequest\x1a\x1d.demo_project.ProjectResponse\x12\x46\n\tcreateTag\x12\x1e.demo_project.CreateTagRequest\x1a\x19.demo_project.TagResponse\x12L\n\ngetProject\x12\x1f.demo_project.GetProjectRequest\x1a\x1d.demo_project.ProjectResponseB6Z4github.com/Joker666/microservice-demo/protos/projectb\x06proto3'
)




_CREATEPROJECTREQUEST = _descriptor.Descriptor(
  name='CreateProjectRequest',
  full_name='demo_project.CreateProjectRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='demo_project.CreateProjectRequest.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='demo_project.CreateProjectRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=31,
  serialized_end=84,
)


_CREATETAGREQUEST = _descriptor.Descriptor(
  name='CreateTagRequest',
  full_name='demo_project.CreateTagRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='demo_project.CreateTagRequest.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='demo_project.CreateTagRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='demo_project.CreateTagRequest.project_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=86,
  serialized_end=155,
)


_GETPROJECTREQUEST = _descriptor.Descriptor(
  name='GetProjectRequest',
  full_name='demo_project.GetProjectRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='user_id', full_name='demo_project.GetProjectRequest.user_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='demo_project.GetProjectRequest.project_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=157,
  serialized_end=213,
)


_PROJECTRESPONSE = _descriptor.Descriptor(
  name='ProjectResponse',
  full_name='demo_project.ProjectResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='demo_project.ProjectResponse.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='demo_project.ProjectResponse.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='tags', full_name='demo_project.ProjectResponse.tags', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=215,
  serialized_end=299,
)


_TAGRESPONSE = _descriptor.Descriptor(
  name='TagResponse',
  full_name='demo_project.TagResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='demo_project.TagResponse.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='demo_project.TagResponse.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='demo_project.TagResponse.project_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=301,
  serialized_end=360,
)

_PROJECTRESPONSE.fields_by_name['tags'].message_type = _TAGRESPONSE
DESCRIPTOR.message_types_by_name['CreateProjectRequest'] = _CREATEPROJECTREQUEST
DESCRIPTOR.message_types_by_name['CreateTagRequest'] = _CREATETAGREQUEST
DESCRIPTOR.message_types_by_name['GetProjectRequest'] = _GETPROJECTREQUEST
DESCRIPTOR.message_types_by_name['ProjectResponse'] = _PROJECTRESPONSE
DESCRIPTOR.message_types_by_name['TagResponse'] = _TAGRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateProjectRequest = _reflection.GeneratedProtocolMessageType('CreateProjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPROJECTREQUEST,
  '__module__' : 'project_pb2'
  # @@protoc_insertion_point(class_scope:demo_project.CreateProjectRequest)
  })
_sym_db.RegisterMessage(CreateProjectRequest)

CreateTagRequest = _reflection.GeneratedProtocolMessageType('CreateTagRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATETAGREQUEST,
  '__module__' : 'project_pb2'
  # @@protoc_insertion_point(class_scope:demo_project.CreateTagRequest)
  })
_sym_db.RegisterMessage(CreateTagRequest)

GetProjectRequest = _reflection.GeneratedProtocolMessageType('GetProjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETPROJECTREQUEST,
  '__module__' : 'project_pb2'
  # @@protoc_insertion_point(class_scope:demo_project.GetProjectRequest)
  })
_sym_db.RegisterMessage(GetProjectRequest)

ProjectResponse = _reflection.GeneratedProtocolMessageType('ProjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _PROJECTRESPONSE,
  '__module__' : 'project_pb2'
  # @@protoc_insertion_point(class_scope:demo_project.ProjectResponse)
  })
_sym_db.RegisterMessage(ProjectResponse)

TagResponse = _reflection.GeneratedProtocolMessageType('TagResponse', (_message.Message,), {
  'DESCRIPTOR' : _TAGRESPONSE,
  '__module__' : 'project_pb2'
  # @@protoc_insertion_point(class_scope:demo_project.TagResponse)
  })
_sym_db.RegisterMessage(TagResponse)


DESCRIPTOR._options = None

_PROJECTSVC = _descriptor.ServiceDescriptor(
  name='ProjectSvc',
  full_name='demo_project.ProjectSvc',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=363,
  serialized_end=609,
  methods=[
  _descriptor.MethodDescriptor(
    name='createProject',
    full_name='demo_project.ProjectSvc.createProject',
    index=0,
    containing_service=None,
    input_type=_CREATEPROJECTREQUEST,
    output_type=_PROJECTRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='createTag',
    full_name='demo_project.ProjectSvc.createTag',
    index=1,
    containing_service=None,
    input_type=_CREATETAGREQUEST,
    output_type=_TAGRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='getProject',
    full_name='demo_project.ProjectSvc.getProject',
    index=2,
    containing_service=None,
    input_type=_GETPROJECTREQUEST,
    output_type=_PROJECTRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_PROJECTSVC)

DESCRIPTOR.services_by_name['ProjectSvc'] = _PROJECTSVC

# @@protoc_insertion_point(module_scope)
