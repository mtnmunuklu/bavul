// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.3
// source: vuln.proto

package pb

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

type CVE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CveId       string   `protobuf:"bytes,2,opt,name=cve_id,json=cveId,proto3" json:"cve_id,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Severity    string   `protobuf:"bytes,4,opt,name=severity,proto3" json:"severity,omitempty"`
	Product     string   `protobuf:"bytes,5,opt,name=product,proto3" json:"product,omitempty"`
	Vendor      string   `protobuf:"bytes,6,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Links       []string `protobuf:"bytes,7,rep,name=links,proto3" json:"links,omitempty"`
	Published   string   `protobuf:"bytes,8,opt,name=published,proto3" json:"published,omitempty"`
	Modified    string   `protobuf:"bytes,9,opt,name=modified,proto3" json:"modified,omitempty"`
}

func (x *CVE) Reset() {
	*x = CVE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vuln_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CVE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CVE) ProtoMessage() {}

func (x *CVE) ProtoReflect() protoreflect.Message {
	mi := &file_vuln_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CVE.ProtoReflect.Descriptor instead.
func (*CVE) Descriptor() ([]byte, []int) {
	return file_vuln_proto_rawDescGZIP(), []int{0}
}

func (x *CVE) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CVE) GetCveId() string {
	if x != nil {
		return x.CveId
	}
	return ""
}

func (x *CVE) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CVE) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *CVE) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *CVE) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *CVE) GetLinks() []string {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *CVE) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *CVE) GetModified() string {
	if x != nil {
		return x.Modified
	}
	return ""
}

type AddCVERequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CveId       string   `protobuf:"bytes,1,opt,name=cve_id,json=cveId,proto3" json:"cve_id,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Severity    string   `protobuf:"bytes,3,opt,name=severity,proto3" json:"severity,omitempty"`
	Product     string   `protobuf:"bytes,4,opt,name=product,proto3" json:"product,omitempty"`
	Vendor      string   `protobuf:"bytes,5,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Links       []string `protobuf:"bytes,6,rep,name=links,proto3" json:"links,omitempty"`
	Published   string   `protobuf:"bytes,7,opt,name=published,proto3" json:"published,omitempty"`
	Modified    string   `protobuf:"bytes,8,opt,name=modified,proto3" json:"modified,omitempty"`
}

func (x *AddCVERequest) Reset() {
	*x = AddCVERequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vuln_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCVERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCVERequest) ProtoMessage() {}

func (x *AddCVERequest) ProtoReflect() protoreflect.Message {
	mi := &file_vuln_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCVERequest.ProtoReflect.Descriptor instead.
func (*AddCVERequest) Descriptor() ([]byte, []int) {
	return file_vuln_proto_rawDescGZIP(), []int{1}
}

func (x *AddCVERequest) GetCveId() string {
	if x != nil {
		return x.CveId
	}
	return ""
}

func (x *AddCVERequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddCVERequest) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *AddCVERequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *AddCVERequest) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *AddCVERequest) GetLinks() []string {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *AddCVERequest) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *AddCVERequest) GetModified() string {
	if x != nil {
		return x.Modified
	}
	return ""
}

type SearchCVERequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CveId     string `protobuf:"bytes,1,opt,name=cve_id,json=cveId,proto3" json:"cve_id,omitempty"`
	Severity  string `protobuf:"bytes,2,opt,name=severity,proto3" json:"severity,omitempty"`
	Product   string `protobuf:"bytes,3,opt,name=product,proto3" json:"product,omitempty"`
	Vendor    string `protobuf:"bytes,4,opt,name=vendor,proto3" json:"vendor,omitempty"`
	StartDate string `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate   string `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
}

func (x *SearchCVERequest) Reset() {
	*x = SearchCVERequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vuln_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchCVERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchCVERequest) ProtoMessage() {}

func (x *SearchCVERequest) ProtoReflect() protoreflect.Message {
	mi := &file_vuln_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchCVERequest.ProtoReflect.Descriptor instead.
func (*SearchCVERequest) Descriptor() ([]byte, []int) {
	return file_vuln_proto_rawDescGZIP(), []int{2}
}

func (x *SearchCVERequest) GetCveId() string {
	if x != nil {
		return x.CveId
	}
	return ""
}

func (x *SearchCVERequest) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *SearchCVERequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *SearchCVERequest) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *SearchCVERequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *SearchCVERequest) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type GetAllCVEsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllCVEsRequest) Reset() {
	*x = GetAllCVEsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vuln_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCVEsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCVEsRequest) ProtoMessage() {}

func (x *GetAllCVEsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vuln_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCVEsRequest.ProtoReflect.Descriptor instead.
func (*GetAllCVEsRequest) Descriptor() ([]byte, []int) {
	return file_vuln_proto_rawDescGZIP(), []int{3}
}

type DeleteCVERequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CveId string `protobuf:"bytes,1,opt,name=cve_id,json=cveId,proto3" json:"cve_id,omitempty"`
}

func (x *DeleteCVERequest) Reset() {
	*x = DeleteCVERequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vuln_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCVERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCVERequest) ProtoMessage() {}

func (x *DeleteCVERequest) ProtoReflect() protoreflect.Message {
	mi := &file_vuln_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCVERequest.ProtoReflect.Descriptor instead.
func (*DeleteCVERequest) Descriptor() ([]byte, []int) {
	return file_vuln_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteCVERequest) GetCveId() string {
	if x != nil {
		return x.CveId
	}
	return ""
}

type DeleteCVEResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CveId string `protobuf:"bytes,1,opt,name=cve_id,json=cveId,proto3" json:"cve_id,omitempty"`
}

func (x *DeleteCVEResponse) Reset() {
	*x = DeleteCVEResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vuln_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCVEResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCVEResponse) ProtoMessage() {}

func (x *DeleteCVEResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vuln_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCVEResponse.ProtoReflect.Descriptor instead.
func (*DeleteCVEResponse) Descriptor() ([]byte, []int) {
	return file_vuln_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCVEResponse) GetCveId() string {
	if x != nil {
		return x.CveId
	}
	return ""
}

type UpdateCVERequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CveId       string   `protobuf:"bytes,1,opt,name=cve_id,json=cveId,proto3" json:"cve_id,omitempty"`
	Description string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Severity    string   `protobuf:"bytes,3,opt,name=severity,proto3" json:"severity,omitempty"`
	Product     string   `protobuf:"bytes,4,opt,name=product,proto3" json:"product,omitempty"`
	Vendor      string   `protobuf:"bytes,5,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Links       []string `protobuf:"bytes,6,rep,name=links,proto3" json:"links,omitempty"`
	Published   string   `protobuf:"bytes,7,opt,name=published,proto3" json:"published,omitempty"`
	Modified    string   `protobuf:"bytes,8,opt,name=modified,proto3" json:"modified,omitempty"`
}

func (x *UpdateCVERequest) Reset() {
	*x = UpdateCVERequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vuln_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCVERequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCVERequest) ProtoMessage() {}

func (x *UpdateCVERequest) ProtoReflect() protoreflect.Message {
	mi := &file_vuln_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCVERequest.ProtoReflect.Descriptor instead.
func (*UpdateCVERequest) Descriptor() ([]byte, []int) {
	return file_vuln_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCVERequest) GetCveId() string {
	if x != nil {
		return x.CveId
	}
	return ""
}

func (x *UpdateCVERequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateCVERequest) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *UpdateCVERequest) GetProduct() string {
	if x != nil {
		return x.Product
	}
	return ""
}

func (x *UpdateCVERequest) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *UpdateCVERequest) GetLinks() []string {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *UpdateCVERequest) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *UpdateCVERequest) GetModified() string {
	if x != nil {
		return x.Modified
	}
	return ""
}

type FetchNVDFeedsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiKey string `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *FetchNVDFeedsRequest) Reset() {
	*x = FetchNVDFeedsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vuln_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchNVDFeedsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchNVDFeedsRequest) ProtoMessage() {}

func (x *FetchNVDFeedsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vuln_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchNVDFeedsRequest.ProtoReflect.Descriptor instead.
func (*FetchNVDFeedsRequest) Descriptor() ([]byte, []int) {
	return file_vuln_proto_rawDescGZIP(), []int{7}
}

func (x *FetchNVDFeedsRequest) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

var File_vuln_proto protoreflect.FileDescriptor

var file_vuln_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x76, 0x75, 0x6c, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0xec, 0x01, 0x0a, 0x03, 0x43, 0x56, 0x45, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15,
	0x0a, 0x06, 0x63, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x63, 0x76, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76,
	0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0xe6, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x43, 0x56, 0x45,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x76, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x76, 0x65, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c,
	0x69, 0x6e, 0x6b, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0xb1,
	0x01, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x56, 0x45, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x76, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65,
	0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x56, 0x45, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x29, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x56, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63,
	0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x76, 0x65,
	0x49, 0x64, 0x22, 0x2a, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x56, 0x45, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x76, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x76, 0x65, 0x49, 0x64, 0x22, 0xe9,
	0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x56, 0x45, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x76, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x73, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x2f, 0x0a, 0x14, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x4e, 0x56, 0x44, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x32, 0xf5, 0x02, 0x0a, 0x0b,
	0x56, 0x75, 0x6c, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x41,
	0x64, 0x64, 0x43, 0x56, 0x45, 0x12, 0x17, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x41, 0x64, 0x64, 0x43, 0x56, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x56, 0x45, 0x12, 0x38, 0x0a,
	0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x56, 0x45, 0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x56, 0x45, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x43, 0x56, 0x45, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x56, 0x45, 0x73, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x56, 0x45, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x56,
	0x45, 0x30, 0x01, 0x12, 0x44, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x56, 0x45,
	0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x56, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x56,
	0x45, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x56, 0x45, 0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x56, 0x45, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x56,
	0x45, 0x12, 0x40, 0x0a, 0x0d, 0x46, 0x65, 0x74, 0x63, 0x68, 0x4e, 0x56, 0x44, 0x46, 0x65, 0x65,
	0x64, 0x73, 0x12, 0x1e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x46, 0x65,
	0x74, 0x63, 0x68, 0x4e, 0x56, 0x44, 0x46, 0x65, 0x65, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x43, 0x56,
	0x45, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_vuln_proto_rawDescOnce sync.Once
	file_vuln_proto_rawDescData = file_vuln_proto_rawDesc
)

func file_vuln_proto_rawDescGZIP() []byte {
	file_vuln_proto_rawDescOnce.Do(func() {
		file_vuln_proto_rawDescData = protoimpl.X.CompressGZIP(file_vuln_proto_rawDescData)
	})
	return file_vuln_proto_rawDescData
}

var file_vuln_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_vuln_proto_goTypes = []interface{}{
	(*CVE)(nil),                  // 0: messages.CVE
	(*AddCVERequest)(nil),        // 1: messages.AddCVERequest
	(*SearchCVERequest)(nil),     // 2: messages.SearchCVERequest
	(*GetAllCVEsRequest)(nil),    // 3: messages.GetAllCVEsRequest
	(*DeleteCVERequest)(nil),     // 4: messages.DeleteCVERequest
	(*DeleteCVEResponse)(nil),    // 5: messages.DeleteCVEResponse
	(*UpdateCVERequest)(nil),     // 6: messages.UpdateCVERequest
	(*FetchNVDFeedsRequest)(nil), // 7: messages.FetchNVDFeedsRequest
}
var file_vuln_proto_depIdxs = []int32{
	1, // 0: messages.VulnService.AddCVE:input_type -> messages.AddCVERequest
	2, // 1: messages.VulnService.SearchCVE:input_type -> messages.SearchCVERequest
	3, // 2: messages.VulnService.GetAllCVEs:input_type -> messages.GetAllCVEsRequest
	4, // 3: messages.VulnService.DeleteCVE:input_type -> messages.DeleteCVERequest
	6, // 4: messages.VulnService.UpdateCVE:input_type -> messages.UpdateCVERequest
	7, // 5: messages.VulnService.FetchNVDFeeds:input_type -> messages.FetchNVDFeedsRequest
	0, // 6: messages.VulnService.AddCVE:output_type -> messages.CVE
	0, // 7: messages.VulnService.SearchCVE:output_type -> messages.CVE
	0, // 8: messages.VulnService.GetAllCVEs:output_type -> messages.CVE
	5, // 9: messages.VulnService.DeleteCVE:output_type -> messages.DeleteCVEResponse
	0, // 10: messages.VulnService.UpdateCVE:output_type -> messages.CVE
	0, // 11: messages.VulnService.FetchNVDFeeds:output_type -> messages.CVE
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vuln_proto_init() }
func file_vuln_proto_init() {
	if File_vuln_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vuln_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CVE); i {
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
		file_vuln_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCVERequest); i {
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
		file_vuln_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchCVERequest); i {
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
		file_vuln_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCVEsRequest); i {
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
		file_vuln_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCVERequest); i {
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
		file_vuln_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCVEResponse); i {
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
		file_vuln_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCVERequest); i {
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
		file_vuln_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchNVDFeedsRequest); i {
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
			RawDescriptor: file_vuln_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vuln_proto_goTypes,
		DependencyIndexes: file_vuln_proto_depIdxs,
		MessageInfos:      file_vuln_proto_msgTypes,
	}.Build()
	File_vuln_proto = out.File
	file_vuln_proto_rawDesc = nil
	file_vuln_proto_goTypes = nil
	file_vuln_proto_depIdxs = nil
}
