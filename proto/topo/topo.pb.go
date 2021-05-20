// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: topo.proto

package topo

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

type Node_Type int32

const (
	Node_Unknown     Node_Type = 0
	Node_Host        Node_Type = 1
	Node_AristaCEOS  Node_Type = 2
	Node_JuniperCEVO Node_Type = 3
	Node_CiscoCXR    Node_Type = 4
	Node_Quagga      Node_Type = 5
	Node_FRR         Node_Type = 6
	Node_JuniperVMX  Node_Type = 7
	Node_CiscoCSR    Node_Type = 8
	Node_NokiaSRL    Node_Type = 9
	Node_IxiaTG      Node_Type = 10
)

// Enum value maps for Node_Type.
var (
	Node_Type_name = map[int32]string{
		0:  "Unknown",
		1:  "Host",
		2:  "AristaCEOS",
		3:  "JuniperCEVO",
		4:  "CiscoCXR",
		5:  "Quagga",
		6:  "FRR",
		7:  "JuniperVMX",
		8:  "CiscoCSR",
		9:  "NokiaSRL",
		10: "IxiaTG",
	}
	Node_Type_value = map[string]int32{
		"Unknown":     0,
		"Host":        1,
		"AristaCEOS":  2,
		"JuniperCEVO": 3,
		"CiscoCXR":    4,
		"Quagga":      5,
		"FRR":         6,
		"JuniperVMX":  7,
		"CiscoCSR":    8,
		"NokiaSRL":    9,
		"IxiaTG":      10,
	}
)

func (x Node_Type) Enum() *Node_Type {
	p := new(Node_Type)
	*p = x
	return p
}

func (x Node_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Node_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_topo_proto_enumTypes[0].Descriptor()
}

func (Node_Type) Type() protoreflect.EnumType {
	return &file_topo_proto_enumTypes[0]
}

func (x Node_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Node_Type.Descriptor instead.
func (Node_Type) EnumDescriptor() ([]byte, []int) {
	return file_topo_proto_rawDescGZIP(), []int{1, 0}
}

// Topology message defines what nodes and links will be created inside the mesh.
type Topology struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`   // Name of the topology - will be linked to the cluster name
	Nodes []*Node `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty"` // List of nodes in the topology
	Links []*Link `protobuf:"bytes,3,rep,name=links,proto3" json:"links,omitempty"` // connections between Nodes.
}

func (x *Topology) Reset() {
	*x = Topology{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topology) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topology) ProtoMessage() {}

func (x *Topology) ProtoReflect() protoreflect.Message {
	mi := &file_topo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topology.ProtoReflect.Descriptor instead.
func (*Topology) Descriptor() ([]byte, []int) {
	return file_topo_proto_rawDescGZIP(), []int{0}
}

func (x *Topology) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Topology) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *Topology) GetLinks() []*Link {
	if x != nil {
		return x.Links
	}
	return nil
}

// Node is a single container inside the topology
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`                                                                                                       // Name of the node in the topology. Must be unique.
	Type        Node_Type           `protobuf:"varint,2,opt,name=type,proto3,enum=topo.Node_Type" json:"type,omitempty"`                                                                                  // Type of node to create.
	Labels      map[string]string   `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`           // Metadata labels describing the node.
	Config      *Config             `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`                                                                                                   // Pod specfic configuration of the node.
	Services    map[uint32]*Service `protobuf:"bytes,6,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`      // Map of services to enable on the node.
	Constraints map[string]string   `protobuf:"bytes,7,rep,name=constraints,proto3" json:"constraints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Any k8s constraints required by node.
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_topo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_topo_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetType() Node_Type {
	if x != nil {
		return x.Type
	}
	return Node_Unknown
}

func (x *Node) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Node) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Node) GetServices() map[uint32]*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Node) GetConstraints() map[string]string {
	if x != nil {
		return x.Constraints
	}
	return nil
}

// Link is single link between nodes in the topology.
// Interfaces must start eth1 - eth0 is the default k8s interface.
type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ANode string `protobuf:"bytes,1,opt,name=a_node,json=aNode,proto3" json:"a_node,omitempty"`
	AInt  string `protobuf:"bytes,2,opt,name=a_int,json=aInt,proto3" json:"a_int,omitempty"`
	ZNode string `protobuf:"bytes,3,opt,name=z_node,json=zNode,proto3" json:"z_node,omitempty"`
	ZInt  string `protobuf:"bytes,4,opt,name=z_int,json=zInt,proto3" json:"z_int,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_topo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_topo_proto_rawDescGZIP(), []int{2}
}

func (x *Link) GetANode() string {
	if x != nil {
		return x.ANode
	}
	return ""
}

func (x *Link) GetAInt() string {
	if x != nil {
		return x.AInt
	}
	return ""
}

func (x *Link) GetZNode() string {
	if x != nil {
		return x.ZNode
	}
	return ""
}

func (x *Link) GetZInt() string {
	if x != nil {
		return x.ZInt
	}
	return ""
}

// Config is the k8s pod specific configuration for a node.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command      []string          `protobuf:"bytes,1,rep,name=command,proto3" json:"command,omitempty"`                                                                                 // Command to pass into pod.
	Args         []string          `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`                                                                                       // Command args to pass into the pod.
	Image        string            `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`                                                                                     // Docker image to use with pod.
	Env          map[string]string `protobuf:"bytes,4,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Map of environment variables to pass into the pod.
	EntryCommand string            `protobuf:"bytes,5,opt,name=entry_command,json=entryCommand,proto3" json:"entry_command,omitempty"`                                                   // Specific entry point command for accessing the pod.
	ConfigPath   string            `protobuf:"bytes,6,opt,name=config_path,json=configPath,proto3" json:"config_path,omitempty"`                                                         // Mount point for configuration inside the pod.
	ConfigFile   string            `protobuf:"bytes,7,opt,name=config_file,json=configFile,proto3" json:"config_file,omitempty"`                                                         // Default configuration file name for the pod.
	Sleep        uint32            `protobuf:"varint,8,opt,name=sleep,proto3" json:"sleep,omitempty"`                                                                                    // Sleeptime before starting the pod.
	// Types that are assignable to ConfigData:
	//	*Config_Data
	//	*Config_File
	ConfigData isConfig_ConfigData `protobuf_oneof:"config_data"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_topo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_topo_proto_rawDescGZIP(), []int{3}
}

func (x *Config) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *Config) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *Config) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Config) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *Config) GetEntryCommand() string {
	if x != nil {
		return x.EntryCommand
	}
	return ""
}

func (x *Config) GetConfigPath() string {
	if x != nil {
		return x.ConfigPath
	}
	return ""
}

func (x *Config) GetConfigFile() string {
	if x != nil {
		return x.ConfigFile
	}
	return ""
}

func (x *Config) GetSleep() uint32 {
	if x != nil {
		return x.Sleep
	}
	return 0
}

func (m *Config) GetConfigData() isConfig_ConfigData {
	if m != nil {
		return m.ConfigData
	}
	return nil
}

func (x *Config) GetData() []byte {
	if x, ok := x.GetConfigData().(*Config_Data); ok {
		return x.Data
	}
	return nil
}

func (x *Config) GetFile() string {
	if x, ok := x.GetConfigData().(*Config_File); ok {
		return x.File
	}
	return ""
}

type isConfig_ConfigData interface {
	isConfig_ConfigData()
}

type Config_Data struct {
	Data []byte `protobuf:"bytes,101,opt,name=data,proto3,oneof"` // Byte data for the startup configuration file.
}

type Config_File struct {
	File string `protobuf:"bytes,102,opt,name=file,proto3,oneof"` // Local file to read for the configuration file.
}

func (*Config_Data) isConfig_ConfigData() {}

func (*Config_File) isConfig_ConfigData() {}

// Service is k8s Service to expose to the cluster
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`        // Name of the service (optional)
	Inside  uint32 `protobuf:"varint,2,opt,name=inside,proto3" json:"inside,omitempty"`   // Inside port to map
	Outside uint32 `protobuf:"varint,3,opt,name=outside,proto3" json:"outside,omitempty"` // Outside port to map (0 = autoassign from cluster)
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_topo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_topo_proto_rawDescGZIP(), []int{4}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetInside() uint32 {
	if x != nil {
		return x.Inside
	}
	return 0
}

func (x *Service) GetOutside() uint32 {
	if x != nil {
		return x.Outside
	}
	return 0
}

var File_topo_proto protoreflect.FileDescriptor

var file_topo_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f,
	0x70, 0x6f, 0x22, 0x62, 0x0a, 0x08, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x22, 0xed, 0x04, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x34,
	0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x6f, 0x70, 0x6f,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69,
	0x6e, 0x74, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4a,
	0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x23, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x43, 0x6f,
	0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x99, 0x01, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x72,
	0x69, 0x73, 0x74, 0x61, 0x43, 0x45, 0x4f, 0x53, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x4a, 0x75,
	0x6e, 0x69, 0x70, 0x65, 0x72, 0x43, 0x45, 0x56, 0x4f, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43,
	0x69, 0x73, 0x63, 0x6f, 0x43, 0x58, 0x52, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x75, 0x61,
	0x67, 0x67, 0x61, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x52, 0x52, 0x10, 0x06, 0x12, 0x0e,
	0x0a, 0x0a, 0x4a, 0x75, 0x6e, 0x69, 0x70, 0x65, 0x72, 0x56, 0x4d, 0x58, 0x10, 0x07, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x69, 0x73, 0x63, 0x6f, 0x43, 0x53, 0x52, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08,
	0x4e, 0x6f, 0x6b, 0x69, 0x61, 0x53, 0x52, 0x4c, 0x10, 0x09, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x78,
	0x69, 0x61, 0x54, 0x47, 0x10, 0x0a, 0x22, 0x5e, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x15,
	0x0a, 0x06, 0x61, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x61, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x49, 0x6e, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x7a, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x7a, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x13, 0x0a, 0x05, 0x7a, 0x5f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x7a, 0x49, 0x6e, 0x74, 0x22, 0xe5, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x27, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x6f, 0x70, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x6c, 0x65, 0x65, 0x70, 0x12, 0x14, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x14, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x66, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x0d, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4f,
	0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x69, 0x6e, 0x73, 0x69, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69,
	0x6e, 0x73, 0x69, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x42,
	0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6b, 0x6e, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x6f, 0x70, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_topo_proto_rawDescOnce sync.Once
	file_topo_proto_rawDescData = file_topo_proto_rawDesc
)

func file_topo_proto_rawDescGZIP() []byte {
	file_topo_proto_rawDescOnce.Do(func() {
		file_topo_proto_rawDescData = protoimpl.X.CompressGZIP(file_topo_proto_rawDescData)
	})
	return file_topo_proto_rawDescData
}

var file_topo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_topo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_topo_proto_goTypes = []interface{}{
	(Node_Type)(0),   // 0: topo.Node.Type
	(*Topology)(nil), // 1: topo.Topology
	(*Node)(nil),     // 2: topo.Node
	(*Link)(nil),     // 3: topo.Link
	(*Config)(nil),   // 4: topo.Config
	(*Service)(nil),  // 5: topo.Service
	nil,              // 6: topo.Node.LabelsEntry
	nil,              // 7: topo.Node.ServicesEntry
	nil,              // 8: topo.Node.ConstraintsEntry
	nil,              // 9: topo.Config.EnvEntry
}
var file_topo_proto_depIdxs = []int32{
	2, // 0: topo.Topology.nodes:type_name -> topo.Node
	3, // 1: topo.Topology.links:type_name -> topo.Link
	0, // 2: topo.Node.type:type_name -> topo.Node.Type
	6, // 3: topo.Node.labels:type_name -> topo.Node.LabelsEntry
	4, // 4: topo.Node.config:type_name -> topo.Config
	7, // 5: topo.Node.services:type_name -> topo.Node.ServicesEntry
	8, // 6: topo.Node.constraints:type_name -> topo.Node.ConstraintsEntry
	9, // 7: topo.Config.env:type_name -> topo.Config.EnvEntry
	5, // 8: topo.Node.ServicesEntry.value:type_name -> topo.Service
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_topo_proto_init() }
func file_topo_proto_init() {
	if File_topo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_topo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Topology); i {
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
		file_topo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_topo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
		file_topo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_topo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
	file_topo_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Config_Data)(nil),
		(*Config_File)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_topo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_topo_proto_goTypes,
		DependencyIndexes: file_topo_proto_depIdxs,
		EnumInfos:         file_topo_proto_enumTypes,
		MessageInfos:      file_topo_proto_msgTypes,
	}.Build()
	File_topo_proto = out.File
	file_topo_proto_rawDesc = nil
	file_topo_proto_goTypes = nil
	file_topo_proto_depIdxs = nil
}
