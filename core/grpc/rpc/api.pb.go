// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: rpc/api.proto

package rpc

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

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// ------------ Request ------------
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets []*TargetInfo `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetTargets() []*TargetInfo {
	if x != nil {
		return x.Targets
	}
	return nil
}

type GetBetweenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Targets []*TargetInfo `protobuf:"bytes,1,rep,name=targets,proto3" json:"targets,omitempty"`
	Start   string        `protobuf:"bytes,2,opt,name=start,proto3" json:"start,omitempty"`
	End     string        `protobuf:"bytes,3,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *GetBetweenRequest) Reset() {
	*x = GetBetweenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBetweenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBetweenRequest) ProtoMessage() {}

func (x *GetBetweenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBetweenRequest.ProtoReflect.Descriptor instead.
func (*GetBetweenRequest) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetBetweenRequest) GetTargets() []*TargetInfo {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *GetBetweenRequest) GetStart() string {
	if x != nil {
		return x.Start
	}
	return ""
}

func (x *GetBetweenRequest) GetEnd() string {
	if x != nil {
		return x.End
	}
	return ""
}

type TargetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo    `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	Oj       []*OJAccount `protobuf:"bytes,2,rep,name=oj,proto3" json:"oj,omitempty"`
}

func (x *TargetInfo) Reset() {
	*x = TargetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetInfo) ProtoMessage() {}

func (x *TargetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetInfo.ProtoReflect.Descriptor instead.
func (*TargetInfo) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{4}
}

func (x *TargetInfo) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *TargetInfo) GetOj() []*OJAccount {
	if x != nil {
		return x.Oj
	}
	return nil
}

type OJAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OjName string   `protobuf:"bytes,1,opt,name=oj_name,json=ojName,proto3" json:"oj_name,omitempty"`
	Id     []string `protobuf:"bytes,2,rep,name=id,proto3" json:"id,omitempty"`
}

func (x *OJAccount) Reset() {
	*x = OJAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OJAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OJAccount) ProtoMessage() {}

func (x *OJAccount) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OJAccount.ProtoReflect.Descriptor instead.
func (*OJAccount) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{5}
}

func (x *OJAccount) GetOjName() string {
	if x != nil {
		return x.OjName
	}
	return ""
}

func (x *OJAccount) GetId() []string {
	if x != nil {
		return x.Id
	}
	return nil
}

// ------------ Response ------------
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScrapyTime string         `protobuf:"bytes,1,opt,name=scrapy_time,json=scrapyTime,proto3" json:"scrapy_time,omitempty"`
	StartTime  string         `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime    string         `protobuf:"bytes,3,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Data       []*CrawlerData `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{6}
}

func (x *Response) GetScrapyTime() string {
	if x != nil {
		return x.ScrapyTime
	}
	return ""
}

func (x *Response) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *Response) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

func (x *Response) GetData() []*CrawlerData {
	if x != nil {
		return x.Data
	}
	return nil
}

type CrawlerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo              `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	Data     map[string]*SolvedData `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CrawlerData) Reset() {
	*x = CrawlerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlerData) ProtoMessage() {}

func (x *CrawlerData) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlerData.ProtoReflect.Descriptor instead.
func (*CrawlerData) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{7}
}

func (x *CrawlerData) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *CrawlerData) GetData() map[string]*SolvedData {
	if x != nil {
		return x.Data
	}
	return nil
}

type SolvedData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *SolvedData_Data `protobuf:"bytes,1,opt,name=data,proto3,oneof" json:"data,omitempty"`
	Error *Error           `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *SolvedData) Reset() {
	*x = SolvedData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolvedData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolvedData) ProtoMessage() {}

func (x *SolvedData) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolvedData.ProtoReflect.Descriptor instead.
func (*SolvedData) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{8}
}

func (x *SolvedData) GetData() *SolvedData_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SolvedData) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type Statistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SolvedNum int32 `protobuf:"varint,1,opt,name=solved_num,json=solvedNum,proto3" json:"solved_num,omitempty"`
	SubmitNum int32 `protobuf:"varint,2,opt,name=submit_num,json=submitNum,proto3" json:"submit_num,omitempty"`
}

func (x *Statistics) Reset() {
	*x = Statistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Statistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Statistics) ProtoMessage() {}

func (x *Statistics) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Statistics.ProtoReflect.Descriptor instead.
func (*Statistics) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{9}
}

func (x *Statistics) GetSolvedNum() int32 {
	if x != nil {
		return x.SolvedNum
	}
	return 0
}

func (x *Statistics) GetSubmitNum() int32 {
	if x != nil {
		return x.SubmitNum
	}
	return 0
}

type SolvedInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemTitle string `protobuf:"bytes,1,opt,name=problem_title,json=problemTitle,proto3" json:"problem_title,omitempty"`
	StatusWord   string `protobuf:"bytes,2,opt,name=status_word,json=statusWord,proto3" json:"status_word,omitempty"`
	Status       bool   `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	SubmitTime   string `protobuf:"bytes,4,opt,name=submit_time,json=submitTime,proto3" json:"submit_time,omitempty"`
	RelatedUrl   string `protobuf:"bytes,5,opt,name=related_url,json=relatedUrl,proto3" json:"related_url,omitempty"`
}

func (x *SolvedInfo) Reset() {
	*x = SolvedInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolvedInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolvedInfo) ProtoMessage() {}

func (x *SolvedInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolvedInfo.ProtoReflect.Descriptor instead.
func (*SolvedInfo) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{10}
}

func (x *SolvedInfo) GetProblemTitle() string {
	if x != nil {
		return x.ProblemTitle
	}
	return ""
}

func (x *SolvedInfo) GetStatusWord() string {
	if x != nil {
		return x.StatusWord
	}
	return ""
}

func (x *SolvedInfo) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *SolvedInfo) GetSubmitTime() string {
	if x != nil {
		return x.SubmitTime
	}
	return ""
}

func (x *SolvedInfo) GetRelatedUrl() string {
	if x != nil {
		return x.RelatedUrl
	}
	return ""
}

type SolvedData_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statistics *Statistics   `protobuf:"bytes,1,opt,name=statistics,proto3" json:"statistics,omitempty"`
	Problems   []*SolvedInfo `protobuf:"bytes,2,rep,name=problems,proto3" json:"problems,omitempty"`
}

func (x *SolvedData_Data) Reset() {
	*x = SolvedData_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolvedData_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolvedData_Data) ProtoMessage() {}

func (x *SolvedData_Data) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolvedData_Data.ProtoReflect.Descriptor instead.
func (*SolvedData_Data) Descriptor() ([]byte, []int) {
	return file_rpc_api_proto_rawDescGZIP(), []int{8, 0}
}

func (x *SolvedData_Data) GetStatistics() *Statistics {
	if x != nil {
		return x.Statistics
	}
	return nil
}

func (x *SolvedData_Data) GetProblems() []*SolvedInfo {
	if x != nil {
		return x.Problems
	}
	return nil
}

var File_rpc_api_proto protoreflect.FileDescriptor

var file_rpc_api_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x22, 0x1e, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x19, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x34, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x22, 0x66, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x65, 0x74, 0x77,
	0x65, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x58, 0x0a,
	0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x02, 0x6f, 0x6a, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4f, 0x4a, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x02, 0x6f, 0x6a, 0x22, 0x34, 0x0a, 0x09, 0x4f, 0x4a, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6f, 0x6a, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x6a, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8b, 0x01,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63,
	0x72, 0x61, 0x70, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x63, 0x72, 0x61, 0x70, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e,
	0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xb3, 0x01, 0x0a, 0x0b,
	0x43, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x61, 0x77,
	0x6c, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x48, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xdb, 0x01, 0x0a, 0x0a, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12,
	0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x01, 0x52, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x1a, 0x64, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x73, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12,
	0x2b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x73, 0x42, 0x07, 0x0a, 0x05,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x4a, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x4e, 0x75, 0x6d, 0x22, 0xac, 0x01, 0x0a, 0x0a,
	0x53, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x57, 0x6f, 0x72, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6d,
	0x69, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x75, 0x62, 0x6d, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x55, 0x72, 0x6c, 0x32, 0xbd, 0x01, 0x0a, 0x03, 0x41,
	0x70, 0x69, 0x12, 0x27, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x0c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x57, 0x65, 0x65, 0x6b, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74,
	0x57, 0x65, 0x65, 0x6b, 0x12, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x42, 0x65, 0x74, 0x77, 0x65, 0x65,
	0x6e, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x65, 0x74, 0x77, 0x65,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x61, 0x63,
	0x6d, 0x77, 0x68, 0x75, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64,
	0x2d, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x72, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_api_proto_rawDescOnce sync.Once
	file_rpc_api_proto_rawDescData = file_rpc_api_proto_rawDesc
)

func file_rpc_api_proto_rawDescGZIP() []byte {
	file_rpc_api_proto_rawDescOnce.Do(func() {
		file_rpc_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_api_proto_rawDescData)
	})
	return file_rpc_api_proto_rawDescData
}

var file_rpc_api_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_rpc_api_proto_goTypes = []interface{}{
	(*UserInfo)(nil),          // 0: api.UserInfo
	(*Error)(nil),             // 1: api.Error
	(*Request)(nil),           // 2: api.Request
	(*GetBetweenRequest)(nil), // 3: api.GetBetweenRequest
	(*TargetInfo)(nil),        // 4: api.TargetInfo
	(*OJAccount)(nil),         // 5: api.OJAccount
	(*Response)(nil),          // 6: api.Response
	(*CrawlerData)(nil),       // 7: api.CrawlerData
	(*SolvedData)(nil),        // 8: api.SolvedData
	(*Statistics)(nil),        // 9: api.Statistics
	(*SolvedInfo)(nil),        // 10: api.SolvedInfo
	nil,                       // 11: api.CrawlerData.DataEntry
	(*SolvedData_Data)(nil),   // 12: api.SolvedData.Data
}
var file_rpc_api_proto_depIdxs = []int32{
	4,  // 0: api.Request.targets:type_name -> api.TargetInfo
	4,  // 1: api.GetBetweenRequest.targets:type_name -> api.TargetInfo
	0,  // 2: api.TargetInfo.user_info:type_name -> api.UserInfo
	5,  // 3: api.TargetInfo.oj:type_name -> api.OJAccount
	7,  // 4: api.Response.data:type_name -> api.CrawlerData
	0,  // 5: api.CrawlerData.user_info:type_name -> api.UserInfo
	11, // 6: api.CrawlerData.data:type_name -> api.CrawlerData.DataEntry
	12, // 7: api.SolvedData.data:type_name -> api.SolvedData.Data
	1,  // 8: api.SolvedData.error:type_name -> api.Error
	8,  // 9: api.CrawlerData.DataEntry.value:type_name -> api.SolvedData
	9,  // 10: api.SolvedData.Data.statistics:type_name -> api.Statistics
	10, // 11: api.SolvedData.Data.problems:type_name -> api.SolvedInfo
	2,  // 12: api.Api.GetAll:input_type -> api.Request
	2,  // 13: api.Api.GetWeek:input_type -> api.Request
	2,  // 14: api.Api.GetLastWeek:input_type -> api.Request
	3,  // 15: api.Api.GetBetween:input_type -> api.GetBetweenRequest
	6,  // 16: api.Api.GetAll:output_type -> api.Response
	6,  // 17: api.Api.GetWeek:output_type -> api.Response
	6,  // 18: api.Api.GetLastWeek:output_type -> api.Response
	6,  // 19: api.Api.GetBetween:output_type -> api.Response
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_rpc_api_proto_init() }
func file_rpc_api_proto_init() {
	if File_rpc_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_rpc_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_rpc_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_rpc_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBetweenRequest); i {
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
		file_rpc_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetInfo); i {
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
		file_rpc_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OJAccount); i {
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
		file_rpc_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_rpc_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlerData); i {
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
		file_rpc_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolvedData); i {
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
		file_rpc_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Statistics); i {
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
		file_rpc_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolvedInfo); i {
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
		file_rpc_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolvedData_Data); i {
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
	file_rpc_api_proto_msgTypes[8].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_api_proto_goTypes,
		DependencyIndexes: file_rpc_api_proto_depIdxs,
		MessageInfos:      file_rpc_api_proto_msgTypes,
	}.Build()
	File_rpc_api_proto = out.File
	file_rpc_api_proto_rawDesc = nil
	file_rpc_api_proto_goTypes = nil
	file_rpc_api_proto_depIdxs = nil
}
