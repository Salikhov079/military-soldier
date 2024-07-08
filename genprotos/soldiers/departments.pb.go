// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: departments.proto

package soldiers

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateDeportment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommandersId string `protobuf:"bytes,1,opt,name=commanders_id,json=commandersId,proto3" json:"commanders_id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateDeportment) Reset() {
	*x = CreateDeportment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_departments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDeportment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDeportment) ProtoMessage() {}

func (x *CreateDeportment) ProtoReflect() protoreflect.Message {
	mi := &file_departments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDeportment.ProtoReflect.Descriptor instead.
func (*CreateDeportment) Descriptor() ([]byte, []int) {
	return file_departments_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDeportment) GetCommandersId() string {
	if x != nil {
		return x.CommandersId
	}
	return ""
}

func (x *CreateDeportment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateDeportment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CommandersId string `protobuf:"bytes,2,opt,name=commanders_id,json=commandersId,proto3" json:"commanders_id,omitempty"`
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateDeportment) Reset() {
	*x = UpdateDeportment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_departments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateDeportment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeportment) ProtoMessage() {}

func (x *UpdateDeportment) ProtoReflect() protoreflect.Message {
	mi := &file_departments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeportment.ProtoReflect.Descriptor instead.
func (*UpdateDeportment) Descriptor() ([]byte, []int) {
	return file_departments_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateDeportment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateDeportment) GetCommandersId() string {
	if x != nil {
		return x.CommandersId
	}
	return ""
}

func (x *UpdateDeportment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAllDepartmentFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAllDepartmentFilter) Reset() {
	*x = GetAllDepartmentFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_departments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllDepartmentFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllDepartmentFilter) ProtoMessage() {}

func (x *GetAllDepartmentFilter) ProtoReflect() protoreflect.Message {
	mi := &file_departments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllDepartmentFilter.ProtoReflect.Descriptor instead.
func (*GetAllDepartmentFilter) Descriptor() ([]byte, []int) {
	return file_departments_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllDepartmentFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Department struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Commander *Commander `protobuf:"bytes,3,opt,name=commander,proto3" json:"commander,omitempty"`
}

func (x *Department) Reset() {
	*x = Department{}
	if protoimpl.UnsafeEnabled {
		mi := &file_departments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Department) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Department) ProtoMessage() {}

func (x *Department) ProtoReflect() protoreflect.Message {
	mi := &file_departments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Department.ProtoReflect.Descriptor instead.
func (*Department) Descriptor() ([]byte, []int) {
	return file_departments_proto_rawDescGZIP(), []int{3}
}

func (x *Department) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Department) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Department) GetCommander() *Commander {
	if x != nil {
		return x.Commander
	}
	return nil
}

type AllDepartments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Departments []*Department `protobuf:"bytes,1,rep,name=departments,proto3" json:"departments,omitempty"`
}

func (x *AllDepartments) Reset() {
	*x = AllDepartments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_departments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllDepartments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllDepartments) ProtoMessage() {}

func (x *AllDepartments) ProtoReflect() protoreflect.Message {
	mi := &file_departments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllDepartments.ProtoReflect.Descriptor instead.
func (*AllDepartments) Descriptor() ([]byte, []int) {
	return file_departments_proto_rawDescGZIP(), []int{4}
}

func (x *AllDepartments) GetDepartments() []*Department {
	if x != nil {
		return x.Departments
	}
	return nil
}

var File_departments_proto protoreflect.FileDescriptor

var file_departments_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x1a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x65, 0x72, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x5b, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72,
	0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x65, 0x0a, 0x0a, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x09,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65,
	0x72, 0x22, 0x4b, 0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x9d,
	0x02, 0x0a, 0x11, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17,
	0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x2c,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x72, 0x73,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x17, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3e,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x1a, 0x1b, 0x2e, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x41, 0x6c, 0x6c, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x14,
	0x5a, 0x12, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x64,
	0x69, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_departments_proto_rawDescOnce sync.Once
	file_departments_proto_rawDescData = file_departments_proto_rawDesc
)

func file_departments_proto_rawDescGZIP() []byte {
	file_departments_proto_rawDescOnce.Do(func() {
		file_departments_proto_rawDescData = protoimpl.X.CompressGZIP(file_departments_proto_rawDescData)
	})
	return file_departments_proto_rawDescData
}

var file_departments_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_departments_proto_goTypes = []interface{}{
	(*CreateDeportment)(nil),       // 0: departments.CreateDeportment
	(*UpdateDeportment)(nil),       // 1: departments.UpdateDeportment
	(*GetAllDepartmentFilter)(nil), // 2: departments.GetAllDepartmentFilter
	(*Department)(nil),             // 3: departments.Department
	(*AllDepartments)(nil),         // 4: departments.AllDepartments
	(*Commander)(nil),              // 5: commanders.Commander
	(*ById)(nil),                   // 6: commanders.ById
	(*Void)(nil),                   // 7: commanders.Void
}
var file_departments_proto_depIdxs = []int32{
	5, // 0: departments.Department.commander:type_name -> commanders.Commander
	3, // 1: departments.AllDepartments.departments:type_name -> departments.Department
	3, // 2: departments.DepartmentService.Create:input_type -> departments.Department
	3, // 3: departments.DepartmentService.Update:input_type -> departments.Department
	6, // 4: departments.DepartmentService.Delete:input_type -> commanders.ById
	6, // 5: departments.DepartmentService.Get:input_type -> commanders.ById
	3, // 6: departments.DepartmentService.GetAll:input_type -> departments.Department
	7, // 7: departments.DepartmentService.Create:output_type -> commanders.Void
	7, // 8: departments.DepartmentService.Update:output_type -> commanders.Void
	7, // 9: departments.DepartmentService.Delete:output_type -> commanders.Void
	3, // 10: departments.DepartmentService.Get:output_type -> departments.Department
	4, // 11: departments.DepartmentService.GetAll:output_type -> departments.AllDepartments
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_departments_proto_init() }
func file_departments_proto_init() {
	if File_departments_proto != nil {
		return
	}
	file_commanders_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_departments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDeportment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_departments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateDeportment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_departments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllDepartmentFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_departments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Department); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_departments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllDepartments); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_departments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_departments_proto_goTypes,
		DependencyIndexes: file_departments_proto_depIdxs,
		MessageInfos:      file_departments_proto_msgTypes,
	}.Build()
	File_departments_proto = out.File
	file_departments_proto_rawDesc = nil
	file_departments_proto_goTypes = nil
	file_departments_proto_depIdxs = nil
}
