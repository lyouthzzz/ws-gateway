// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.5
// source: app/ws-gateway/internal/config/config.proto

package config

import (
	client "github.com/lyouthzzz/ws-gateway/pkg/client"
	server "github.com/lyouthzzz/ws-gateway/pkg/server"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Bootstrap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server   *Server   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Client   *Client   `protobuf:"bytes,2,opt,name=client,proto3" json:"client,omitempty"`
	Data     *Data     `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Registry *Registry `protobuf:"bytes,4,opt,name=registry,proto3" json:"registry,omitempty"`
}

func (x *Bootstrap) Reset() {
	*x = Bootstrap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bootstrap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bootstrap) ProtoMessage() {}

func (x *Bootstrap) ProtoReflect() protoreflect.Message {
	mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bootstrap.ProtoReflect.Descriptor instead.
func (*Bootstrap) Descriptor() ([]byte, []int) {
	return file_app_ws_gateway_internal_config_config_proto_rawDescGZIP(), []int{0}
}

func (x *Bootstrap) GetServer() *Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Bootstrap) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *Bootstrap) GetData() *Data {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Bootstrap) GetRegistry() *Registry {
	if x != nil {
		return x.Registry
	}
	return nil
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Http *server.Config `protobuf:"bytes,1,opt,name=http,proto3" json:"http,omitempty"`
	Grpc *server.Config `protobuf:"bytes,2,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_app_ws_gateway_internal_config_config_proto_rawDescGZIP(), []int{1}
}

func (x *Server) GetHttp() *server.Config {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Server) GetGrpc() *server.Config {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WsAPI *client.Config `protobuf:"bytes,1,opt,name=wsAPI,proto3" json:"wsAPI,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_app_ws_gateway_internal_config_config_proto_rawDescGZIP(), []int{2}
}

func (x *Client) GetWsAPI() *client.Config {
	if x != nil {
		return x.WsAPI
	}
	return nil
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_app_ws_gateway_internal_config_config_proto_rawDescGZIP(), []int{3}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Etcd *Registry_ETCD `protobuf:"bytes,1,opt,name=etcd,proto3" json:"etcd,omitempty"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_app_ws_gateway_internal_config_config_proto_rawDescGZIP(), []int{4}
}

func (x *Registry) GetEtcd() *Registry_ETCD {
	if x != nil {
		return x.Etcd
	}
	return nil
}

type Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_app_ws_gateway_internal_config_config_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network      string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr         string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	ReadTimeout  *durationpb.Duration `protobuf:"bytes,3,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`
	WriteTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_app_ws_gateway_internal_config_config_proto_rawDescGZIP(), []int{3, 1}
}

func (x *Data_Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

type Registry_ETCD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr        string               `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	DialTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=dialTimeout,proto3" json:"dialTimeout,omitempty"`
}

func (x *Registry_ETCD) Reset() {
	*x = Registry_ETCD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry_ETCD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry_ETCD) ProtoMessage() {}

func (x *Registry_ETCD) ProtoReflect() protoreflect.Message {
	mi := &file_app_ws_gateway_internal_config_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry_ETCD.ProtoReflect.Descriptor instead.
func (*Registry_ETCD) Descriptor() ([]byte, []int) {
	return file_app_ws_gateway_internal_config_config_proto_rawDescGZIP(), []int{4, 0}
}

func (x *Registry_ETCD) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Registry_ETCD) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

var File_app_ws_gateway_internal_config_config_proto protoreflect.FileDescriptor

var file_app_ws_gateway_internal_config_config_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x70, 0x2f, 0x77, 0x73, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x01, 0x0a, 0x09, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61,
	0x70, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2a, 0x0a,
	0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x30, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x22, 0x50, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x22, 0x0a, 0x04, 0x68,
	0x74, 0x74, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12,
	0x22, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x04, 0x67,
	0x72, 0x70, 0x63, 0x22, 0x2e, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x05, 0x77, 0x73, 0x41, 0x50, 0x49, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x77, 0x73,
	0x41, 0x50, 0x49, 0x22, 0xaf, 0x02, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x0a, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x1a, 0x3a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a,
	0xb3, 0x01, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x92, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x12, 0x2d, 0x0a, 0x04, 0x65, 0x74, 0x63, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x54, 0x43, 0x44, 0x52, 0x04, 0x65, 0x74, 0x63,
	0x64, 0x1a, 0x57, 0x0a, 0x04, 0x45, 0x54, 0x43, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x3b, 0x0a,
	0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64,
	0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x27, 0x5a, 0x25, 0x61, 0x70,
	0x70, 0x2f, 0x77, 0x73, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3b, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_ws_gateway_internal_config_config_proto_rawDescOnce sync.Once
	file_app_ws_gateway_internal_config_config_proto_rawDescData = file_app_ws_gateway_internal_config_config_proto_rawDesc
)

func file_app_ws_gateway_internal_config_config_proto_rawDescGZIP() []byte {
	file_app_ws_gateway_internal_config_config_proto_rawDescOnce.Do(func() {
		file_app_ws_gateway_internal_config_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_ws_gateway_internal_config_config_proto_rawDescData)
	})
	return file_app_ws_gateway_internal_config_config_proto_rawDescData
}

var file_app_ws_gateway_internal_config_config_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_app_ws_gateway_internal_config_config_proto_goTypes = []interface{}{
	(*Bootstrap)(nil),           // 0: kratos.api.Bootstrap
	(*Server)(nil),              // 1: kratos.api.Server
	(*Client)(nil),              // 2: kratos.api.Client
	(*Data)(nil),                // 3: kratos.api.Data
	(*Registry)(nil),            // 4: kratos.api.Registry
	(*Data_Database)(nil),       // 5: kratos.api.Data.Database
	(*Data_Redis)(nil),          // 6: kratos.api.Data.Redis
	(*Registry_ETCD)(nil),       // 7: kratos.api.Registry.ETCD
	(*server.Config)(nil),       // 8: server.Config
	(*client.Config)(nil),       // 9: client.Config
	(*durationpb.Duration)(nil), // 10: google.protobuf.Duration
}
var file_app_ws_gateway_internal_config_config_proto_depIdxs = []int32{
	1,  // 0: kratos.api.Bootstrap.server:type_name -> kratos.api.Server
	2,  // 1: kratos.api.Bootstrap.client:type_name -> kratos.api.Client
	3,  // 2: kratos.api.Bootstrap.data:type_name -> kratos.api.Data
	4,  // 3: kratos.api.Bootstrap.registry:type_name -> kratos.api.Registry
	8,  // 4: kratos.api.Server.http:type_name -> server.Config
	8,  // 5: kratos.api.Server.grpc:type_name -> server.Config
	9,  // 6: kratos.api.Client.wsAPI:type_name -> client.Config
	5,  // 7: kratos.api.Data.database:type_name -> kratos.api.Data.Database
	7,  // 8: kratos.api.Registry.etcd:type_name -> kratos.api.Registry.ETCD
	10, // 9: kratos.api.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	10, // 10: kratos.api.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	10, // 11: kratos.api.Registry.ETCD.dialTimeout:type_name -> google.protobuf.Duration
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_app_ws_gateway_internal_config_config_proto_init() }
func file_app_ws_gateway_internal_config_config_proto_init() {
	if File_app_ws_gateway_internal_config_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_ws_gateway_internal_config_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bootstrap); i {
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
		file_app_ws_gateway_internal_config_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_app_ws_gateway_internal_config_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_app_ws_gateway_internal_config_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_app_ws_gateway_internal_config_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
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
		file_app_ws_gateway_internal_config_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Database); i {
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
		file_app_ws_gateway_internal_config_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Redis); i {
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
		file_app_ws_gateway_internal_config_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry_ETCD); i {
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
			RawDescriptor: file_app_ws_gateway_internal_config_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_ws_gateway_internal_config_config_proto_goTypes,
		DependencyIndexes: file_app_ws_gateway_internal_config_config_proto_depIdxs,
		MessageInfos:      file_app_ws_gateway_internal_config_config_proto_msgTypes,
	}.Build()
	File_app_ws_gateway_internal_config_config_proto = out.File
	file_app_ws_gateway_internal_config_config_proto_rawDesc = nil
	file_app_ws_gateway_internal_config_config_proto_goTypes = nil
	file_app_ws_gateway_internal_config_config_proto_depIdxs = nil
}
