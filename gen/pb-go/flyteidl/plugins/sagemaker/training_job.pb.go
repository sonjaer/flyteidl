// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/sagemaker/training_job.proto

package plugins

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type InputMode_Value int32

const (
	InputMode_FILE InputMode_Value = 0
	InputMode_PIPE InputMode_Value = 1
)

var InputMode_Value_name = map[int32]string{
	0: "FILE",
	1: "PIPE",
}

var InputMode_Value_value = map[string]int32{
	"FILE": 0,
	"PIPE": 1,
}

func (x InputMode_Value) String() string {
	return proto.EnumName(InputMode_Value_name, int32(x))
}

func (InputMode_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{0, 0}
}

type AlgorithmName_Value int32

const (
	AlgorithmName_CUSTOM  AlgorithmName_Value = 0
	AlgorithmName_XGBOOST AlgorithmName_Value = 1
)

var AlgorithmName_Value_name = map[int32]string{
	0: "CUSTOM",
	1: "XGBOOST",
}

var AlgorithmName_Value_value = map[string]int32{
	"CUSTOM":  0,
	"XGBOOST": 1,
}

func (x AlgorithmName_Value) String() string {
	return proto.EnumName(AlgorithmName_Value_name, int32(x))
}

func (AlgorithmName_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{1, 0}
}

type InputContentType_Value int32

const (
	InputContentType_TEXT_CSV InputContentType_Value = 0
)

var InputContentType_Value_name = map[int32]string{
	0: "TEXT_CSV",
}

var InputContentType_Value_value = map[string]int32{
	"TEXT_CSV": 0,
}

func (x InputContentType_Value) String() string {
	return proto.EnumName(InputContentType_Value_name, int32(x))
}

func (InputContentType_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{2, 0}
}

type DistributedProtocol_Value int32

const (
	DistributedProtocol_UNSPECIFIED DistributedProtocol_Value = 0
	DistributedProtocol_MPI         DistributedProtocol_Value = 1
)

var DistributedProtocol_Value_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "MPI",
}

var DistributedProtocol_Value_value = map[string]int32{
	"UNSPECIFIED": 0,
	"MPI":         1,
}

func (x DistributedProtocol_Value) String() string {
	return proto.EnumName(DistributedProtocol_Value_name, int32(x))
}

func (DistributedProtocol_Value) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{5, 0}
}

// The input mode that the algorithm supports. When using the File input mode, SageMaker downloads
// the training data from S3 to the provisioned ML storage Volume, and mounts the directory to docker
// volume for training container. When using Pipe input mode, Amazon SageMaker streams data directly
// from S3 to the container.
// See: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AlgorithmSpecification.html
// For the input modes that different SageMaker algorithms support, see:
// https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-algo-docker-registry-paths.html
type InputMode struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputMode) Reset()         { *m = InputMode{} }
func (m *InputMode) String() string { return proto.CompactTextString(m) }
func (*InputMode) ProtoMessage()    {}
func (*InputMode) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{0}
}

func (m *InputMode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputMode.Unmarshal(m, b)
}
func (m *InputMode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputMode.Marshal(b, m, deterministic)
}
func (m *InputMode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputMode.Merge(m, src)
}
func (m *InputMode) XXX_Size() int {
	return xxx_messageInfo_InputMode.Size(m)
}
func (m *InputMode) XXX_DiscardUnknown() {
	xxx_messageInfo_InputMode.DiscardUnknown(m)
}

var xxx_messageInfo_InputMode proto.InternalMessageInfo

// The algorithm name is used for deciding which pre-built image to point to.
// This is only required for use cases where SageMaker's built-in algorithm mode is used.
// While we currently only support a subset of the algorithms, more will be added to the list.
// See: https://docs.aws.amazon.com/sagemaker/latest/dg/algos.html
type AlgorithmName struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlgorithmName) Reset()         { *m = AlgorithmName{} }
func (m *AlgorithmName) String() string { return proto.CompactTextString(m) }
func (*AlgorithmName) ProtoMessage()    {}
func (*AlgorithmName) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{1}
}

func (m *AlgorithmName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlgorithmName.Unmarshal(m, b)
}
func (m *AlgorithmName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlgorithmName.Marshal(b, m, deterministic)
}
func (m *AlgorithmName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlgorithmName.Merge(m, src)
}
func (m *AlgorithmName) XXX_Size() int {
	return xxx_messageInfo_AlgorithmName.Size(m)
}
func (m *AlgorithmName) XXX_DiscardUnknown() {
	xxx_messageInfo_AlgorithmName.DiscardUnknown(m)
}

var xxx_messageInfo_AlgorithmName proto.InternalMessageInfo

// Specifies the type of file for input data. Different SageMaker built-in algorithms require different file types of input data
// See https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-training.html
// https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-algo-docker-registry-paths.html
type InputContentType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputContentType) Reset()         { *m = InputContentType{} }
func (m *InputContentType) String() string { return proto.CompactTextString(m) }
func (*InputContentType) ProtoMessage()    {}
func (*InputContentType) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{2}
}

func (m *InputContentType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputContentType.Unmarshal(m, b)
}
func (m *InputContentType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputContentType.Marshal(b, m, deterministic)
}
func (m *InputContentType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputContentType.Merge(m, src)
}
func (m *InputContentType) XXX_Size() int {
	return xxx_messageInfo_InputContentType.Size(m)
}
func (m *InputContentType) XXX_DiscardUnknown() {
	xxx_messageInfo_InputContentType.DiscardUnknown(m)
}

var xxx_messageInfo_InputContentType proto.InternalMessageInfo

// Specifies a metric that the training algorithm writes to stderr or stdout.
// This object is a pass-through.
// See this for details: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_MetricDefinition.html
type MetricDefinition struct {
	// User-defined name of the metric
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// SageMaker hyperparameter tuning parses your algorithm’s stdout and stderr streams to find algorithm metrics
	Regex                string   `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricDefinition) Reset()         { *m = MetricDefinition{} }
func (m *MetricDefinition) String() string { return proto.CompactTextString(m) }
func (*MetricDefinition) ProtoMessage()    {}
func (*MetricDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{3}
}

func (m *MetricDefinition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricDefinition.Unmarshal(m, b)
}
func (m *MetricDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricDefinition.Marshal(b, m, deterministic)
}
func (m *MetricDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricDefinition.Merge(m, src)
}
func (m *MetricDefinition) XXX_Size() int {
	return xxx_messageInfo_MetricDefinition.Size(m)
}
func (m *MetricDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_MetricDefinition proto.InternalMessageInfo

func (m *MetricDefinition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricDefinition) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

// Specifies the training algorithm to be used in the training job
// This object is mostly a pass-through, with a couple of exceptions include: (1) in Flyte, users don't need to specify
// TrainingImage; either use the built-in algorithm mode by using Flytekit's Simple Training Job and specifying an algorithm
// name and an algorithm version or (2) when users want to supply custom algorithms they should set algorithm_name field to
// CUSTOM. In this case, the value of the algorithm_version field has no effect
// For pass-through use cases: refer to this AWS official document for more details
// https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AlgorithmSpecification.html
type AlgorithmSpecification struct {
	// The input mode can be either PIPE or FILE
	InputMode InputMode_Value `protobuf:"varint,1,opt,name=input_mode,json=inputMode,proto3,enum=flyteidl.plugins.sagemaker.InputMode_Value" json:"input_mode,omitempty"`
	// The algorithm name is used for deciding which pre-built image to point to
	AlgorithmName AlgorithmName_Value `protobuf:"varint,2,opt,name=algorithm_name,json=algorithmName,proto3,enum=flyteidl.plugins.sagemaker.AlgorithmName_Value" json:"algorithm_name,omitempty"`
	// The algorithm version field is used for deciding which pre-built image to point to
	// This is only needed for use cases where SageMaker's built-in algorithm mode is chosen
	AlgorithmVersion string `protobuf:"bytes,3,opt,name=algorithm_version,json=algorithmVersion,proto3" json:"algorithm_version,omitempty"`
	// A list of metric definitions for SageMaker to evaluate/track on the progress of the training job
	// See this: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_AlgorithmSpecification.html
	// and this: https://docs.aws.amazon.com/sagemaker/latest/dg/automatic-model-tuning-define-metrics.html
	MetricDefinitions []*MetricDefinition `protobuf:"bytes,4,rep,name=metric_definitions,json=metricDefinitions,proto3" json:"metric_definitions,omitempty"`
	// The content type of the input
	// See https://docs.aws.amazon.com/sagemaker/latest/dg/cdf-training.html
	// https://docs.aws.amazon.com/sagemaker/latest/dg/sagemaker-algo-docker-registry-paths.html
	InputContentType     InputContentType_Value `protobuf:"varint,5,opt,name=input_content_type,json=inputContentType,proto3,enum=flyteidl.plugins.sagemaker.InputContentType_Value" json:"input_content_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AlgorithmSpecification) Reset()         { *m = AlgorithmSpecification{} }
func (m *AlgorithmSpecification) String() string { return proto.CompactTextString(m) }
func (*AlgorithmSpecification) ProtoMessage()    {}
func (*AlgorithmSpecification) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{4}
}

func (m *AlgorithmSpecification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlgorithmSpecification.Unmarshal(m, b)
}
func (m *AlgorithmSpecification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlgorithmSpecification.Marshal(b, m, deterministic)
}
func (m *AlgorithmSpecification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlgorithmSpecification.Merge(m, src)
}
func (m *AlgorithmSpecification) XXX_Size() int {
	return xxx_messageInfo_AlgorithmSpecification.Size(m)
}
func (m *AlgorithmSpecification) XXX_DiscardUnknown() {
	xxx_messageInfo_AlgorithmSpecification.DiscardUnknown(m)
}

var xxx_messageInfo_AlgorithmSpecification proto.InternalMessageInfo

func (m *AlgorithmSpecification) GetInputMode() InputMode_Value {
	if m != nil {
		return m.InputMode
	}
	return InputMode_FILE
}

func (m *AlgorithmSpecification) GetAlgorithmName() AlgorithmName_Value {
	if m != nil {
		return m.AlgorithmName
	}
	return AlgorithmName_CUSTOM
}

func (m *AlgorithmSpecification) GetAlgorithmVersion() string {
	if m != nil {
		return m.AlgorithmVersion
	}
	return ""
}

func (m *AlgorithmSpecification) GetMetricDefinitions() []*MetricDefinition {
	if m != nil {
		return m.MetricDefinitions
	}
	return nil
}

func (m *AlgorithmSpecification) GetInputContentType() InputContentType_Value {
	if m != nil {
		return m.InputContentType
	}
	return InputContentType_TEXT_CSV
}

type DistributedProtocol struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistributedProtocol) Reset()         { *m = DistributedProtocol{} }
func (m *DistributedProtocol) String() string { return proto.CompactTextString(m) }
func (*DistributedProtocol) ProtoMessage()    {}
func (*DistributedProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{5}
}

func (m *DistributedProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedProtocol.Unmarshal(m, b)
}
func (m *DistributedProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedProtocol.Marshal(b, m, deterministic)
}
func (m *DistributedProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedProtocol.Merge(m, src)
}
func (m *DistributedProtocol) XXX_Size() int {
	return xxx_messageInfo_DistributedProtocol.Size(m)
}
func (m *DistributedProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedProtocol proto.InternalMessageInfo

// TrainingJobResourceConfig is a pass-through, specifying the instance type to use for the training job, the
// number of instances to launch, and the size of the ML storage volume the user wants to provision
// Refer to SageMaker official doc for more details: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html
type TrainingJobResourceConfig struct {
	// The number of ML compute instances to use. For distributed training, provide a value greater than 1.
	InstanceCount int64 `protobuf:"varint,1,opt,name=instance_count,json=instanceCount,proto3" json:"instance_count,omitempty"`
	// The ML compute instance type
	InstanceType string `protobuf:"bytes,2,opt,name=instance_type,json=instanceType,proto3" json:"instance_type,omitempty"`
	// The size of the ML storage volume that you want to provision.
	VolumeSizeInGb       int64                `protobuf:"varint,3,opt,name=volume_size_in_gb,json=volumeSizeInGb,proto3" json:"volume_size_in_gb,omitempty"`
	DistributedProtocol  *DistributedProtocol `protobuf:"bytes,4,opt,name=distributed_protocol,json=distributedProtocol,proto3" json:"distributed_protocol,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TrainingJobResourceConfig) Reset()         { *m = TrainingJobResourceConfig{} }
func (m *TrainingJobResourceConfig) String() string { return proto.CompactTextString(m) }
func (*TrainingJobResourceConfig) ProtoMessage()    {}
func (*TrainingJobResourceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{6}
}

func (m *TrainingJobResourceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainingJobResourceConfig.Unmarshal(m, b)
}
func (m *TrainingJobResourceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainingJobResourceConfig.Marshal(b, m, deterministic)
}
func (m *TrainingJobResourceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainingJobResourceConfig.Merge(m, src)
}
func (m *TrainingJobResourceConfig) XXX_Size() int {
	return xxx_messageInfo_TrainingJobResourceConfig.Size(m)
}
func (m *TrainingJobResourceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainingJobResourceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TrainingJobResourceConfig proto.InternalMessageInfo

func (m *TrainingJobResourceConfig) GetInstanceCount() int64 {
	if m != nil {
		return m.InstanceCount
	}
	return 0
}

func (m *TrainingJobResourceConfig) GetInstanceType() string {
	if m != nil {
		return m.InstanceType
	}
	return ""
}

func (m *TrainingJobResourceConfig) GetVolumeSizeInGb() int64 {
	if m != nil {
		return m.VolumeSizeInGb
	}
	return 0
}

func (m *TrainingJobResourceConfig) GetDistributedProtocol() *DistributedProtocol {
	if m != nil {
		return m.DistributedProtocol
	}
	return nil
}

// The spec of a training job. This is mostly a pass-through object
// https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html
type TrainingJob struct {
	AlgorithmSpecification    *AlgorithmSpecification    `protobuf:"bytes,1,opt,name=algorithm_specification,json=algorithmSpecification,proto3" json:"algorithm_specification,omitempty"`
	TrainingJobResourceConfig *TrainingJobResourceConfig `protobuf:"bytes,2,opt,name=training_job_resource_config,json=trainingJobResourceConfig,proto3" json:"training_job_resource_config,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                   `json:"-"`
	XXX_unrecognized          []byte                     `json:"-"`
	XXX_sizecache             int32                      `json:"-"`
}

func (m *TrainingJob) Reset()         { *m = TrainingJob{} }
func (m *TrainingJob) String() string { return proto.CompactTextString(m) }
func (*TrainingJob) ProtoMessage()    {}
func (*TrainingJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a68f64d8fd9fe30, []int{7}
}

func (m *TrainingJob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TrainingJob.Unmarshal(m, b)
}
func (m *TrainingJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TrainingJob.Marshal(b, m, deterministic)
}
func (m *TrainingJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TrainingJob.Merge(m, src)
}
func (m *TrainingJob) XXX_Size() int {
	return xxx_messageInfo_TrainingJob.Size(m)
}
func (m *TrainingJob) XXX_DiscardUnknown() {
	xxx_messageInfo_TrainingJob.DiscardUnknown(m)
}

var xxx_messageInfo_TrainingJob proto.InternalMessageInfo

func (m *TrainingJob) GetAlgorithmSpecification() *AlgorithmSpecification {
	if m != nil {
		return m.AlgorithmSpecification
	}
	return nil
}

func (m *TrainingJob) GetTrainingJobResourceConfig() *TrainingJobResourceConfig {
	if m != nil {
		return m.TrainingJobResourceConfig
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.plugins.sagemaker.InputMode_Value", InputMode_Value_name, InputMode_Value_value)
	proto.RegisterEnum("flyteidl.plugins.sagemaker.AlgorithmName_Value", AlgorithmName_Value_name, AlgorithmName_Value_value)
	proto.RegisterEnum("flyteidl.plugins.sagemaker.InputContentType_Value", InputContentType_Value_name, InputContentType_Value_value)
	proto.RegisterEnum("flyteidl.plugins.sagemaker.DistributedProtocol_Value", DistributedProtocol_Value_name, DistributedProtocol_Value_value)
	proto.RegisterType((*InputMode)(nil), "flyteidl.plugins.sagemaker.InputMode")
	proto.RegisterType((*AlgorithmName)(nil), "flyteidl.plugins.sagemaker.AlgorithmName")
	proto.RegisterType((*InputContentType)(nil), "flyteidl.plugins.sagemaker.InputContentType")
	proto.RegisterType((*MetricDefinition)(nil), "flyteidl.plugins.sagemaker.MetricDefinition")
	proto.RegisterType((*AlgorithmSpecification)(nil), "flyteidl.plugins.sagemaker.AlgorithmSpecification")
	proto.RegisterType((*DistributedProtocol)(nil), "flyteidl.plugins.sagemaker.DistributedProtocol")
	proto.RegisterType((*TrainingJobResourceConfig)(nil), "flyteidl.plugins.sagemaker.TrainingJobResourceConfig")
	proto.RegisterType((*TrainingJob)(nil), "flyteidl.plugins.sagemaker.TrainingJob")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/sagemaker/training_job.proto", fileDescriptor_6a68f64d8fd9fe30)
}

var fileDescriptor_6a68f64d8fd9fe30 = []byte{
	// 658 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4b, 0x4f, 0xdb, 0x4a,
	0x14, 0x8e, 0x49, 0x78, 0x9d, 0x40, 0xae, 0x19, 0xb8, 0xdc, 0xc0, 0xad, 0xaa, 0xd4, 0x55, 0xa5,
	0x20, 0x4a, 0xac, 0x06, 0x21, 0x75, 0xd1, 0x4d, 0x09, 0x01, 0x19, 0x35, 0x10, 0x39, 0x21, 0x42,
	0xed, 0xc2, 0xf5, 0x63, 0x6c, 0xa6, 0xd8, 0x33, 0x96, 0x3d, 0x46, 0x0d, 0xbf, 0xa8, 0xbf, 0xaf,
	0x8b, 0xae, 0x2b, 0x8f, 0x63, 0x93, 0x46, 0x21, 0xdd, 0xd9, 0xdf, 0x9c, 0x39, 0x73, 0xbe, 0xc7,
	0x0c, 0x1c, 0xb9, 0xfe, 0x98, 0x63, 0xe2, 0xf8, 0x6a, 0xe8, 0x27, 0x1e, 0xa1, 0xb1, 0x1a, 0x9b,
	0x1e, 0x0e, 0xcc, 0x7b, 0x1c, 0xa9, 0x3c, 0x32, 0x09, 0x25, 0xd4, 0x33, 0xbe, 0x31, 0xab, 0x15,
	0x46, 0x8c, 0x33, 0xb4, 0x9f, 0x97, 0xb7, 0x26, 0xe5, 0xad, 0xa2, 0x7c, 0xff, 0xa5, 0xc7, 0x98,
	0xe7, 0x63, 0x55, 0x54, 0x5a, 0x89, 0xab, 0x3a, 0x49, 0x64, 0x72, 0xc2, 0x68, 0xb6, 0x57, 0x69,
	0xc2, 0xba, 0x46, 0xc3, 0x84, 0xf7, 0x98, 0x83, 0x95, 0xff, 0x61, 0x79, 0x64, 0xfa, 0x09, 0x46,
	0x6b, 0x50, 0x39, 0xd7, 0x3e, 0x75, 0xe5, 0x52, 0xfa, 0xd5, 0xd7, 0xfa, 0x5d, 0x59, 0x52, 0xde,
	0xc1, 0xe6, 0x47, 0xdf, 0x63, 0x11, 0xe1, 0x77, 0xc1, 0x95, 0x19, 0x60, 0xa5, 0x91, 0x57, 0x03,
	0xac, 0x74, 0x6e, 0x06, 0xc3, 0xeb, 0x9e, 0x5c, 0x42, 0x55, 0x58, 0xbd, 0xbd, 0x38, 0xbd, 0xbe,
	0x1e, 0x0c, 0x65, 0x49, 0x39, 0x00, 0x59, 0x34, 0xef, 0x30, 0xca, 0x31, 0xe5, 0xc3, 0x71, 0x88,
	0x95, 0x7f, 0xf3, 0x5d, 0x1b, 0xb0, 0x36, 0xec, 0xde, 0x0e, 0x8d, 0xce, 0x60, 0x24, 0x97, 0x94,
	0x0f, 0x20, 0xf7, 0x30, 0x8f, 0x88, 0x7d, 0x86, 0x5d, 0x42, 0x49, 0x3a, 0x21, 0x42, 0x50, 0xa1,
	0x66, 0x80, 0xeb, 0x52, 0x43, 0x6a, 0xae, 0xeb, 0xe2, 0x1b, 0xed, 0xc0, 0x72, 0x84, 0x3d, 0xfc,
	0xbd, 0xbe, 0x24, 0xc0, 0xec, 0x47, 0xf9, 0x51, 0x86, 0xdd, 0x62, 0xb8, 0x41, 0x88, 0x6d, 0xe2,
	0x12, 0x5b, 0xd0, 0x44, 0x97, 0x00, 0x24, 0x9d, 0xc1, 0x08, 0x98, 0x93, 0xb5, 0xaa, 0xb5, 0x0f,
	0x5b, 0xcf, 0x2b, 0xd6, 0x2a, 0xe4, 0x68, 0x89, 0x39, 0xf5, 0x75, 0x92, 0x03, 0x68, 0x04, 0x35,
	0x33, 0x3f, 0xc5, 0x10, 0xa3, 0x2d, 0x89, 0x7e, 0xea, 0xa2, 0x7e, 0x7f, 0x88, 0x36, 0xe9, 0xb9,
	0x69, 0x4e, 0x83, 0xe8, 0x10, 0xb6, 0x9e, 0xfa, 0x3e, 0xe0, 0x28, 0x26, 0x8c, 0xd6, 0xcb, 0x82,
	0xa0, 0x5c, 0x2c, 0x8c, 0x32, 0x1c, 0x7d, 0x01, 0x14, 0x08, 0xa5, 0x0c, 0xa7, 0x90, 0x2a, 0xae,
	0x57, 0x1a, 0xe5, 0x66, 0xb5, 0xfd, 0x76, 0xd1, 0x20, 0xb3, 0xfa, 0xea, 0x5b, 0xc1, 0x0c, 0x12,
	0xa3, 0xaf, 0x80, 0x32, 0xb5, 0xec, 0xcc, 0x32, 0x83, 0x8f, 0x43, 0x5c, 0x5f, 0x16, 0x2c, 0xdb,
	0x7f, 0x55, 0x6d, 0xca, 0xe7, 0x09, 0x51, 0x99, 0xcc, 0xfa, 0xff, 0x1e, 0xb6, 0xcf, 0x48, 0xcc,
	0x23, 0x62, 0x25, 0x1c, 0x3b, 0xfd, 0x34, 0x84, 0x36, 0xf3, 0x95, 0x57, 0x79, 0x2c, 0xfe, 0x81,
	0xea, 0xcd, 0xd5, 0xa0, 0xdf, 0xed, 0x68, 0xe7, 0x5a, 0xf7, 0x4c, 0x2e, 0xa1, 0x55, 0x28, 0xf7,
	0xfa, 0x9a, 0x2c, 0x29, 0xbf, 0x24, 0xd8, 0x1b, 0x4e, 0xd2, 0x7f, 0xc9, 0x2c, 0x1d, 0xc7, 0x2c,
	0x89, 0x6c, 0xdc, 0x61, 0xd4, 0x25, 0x1e, 0x7a, 0x03, 0x35, 0x42, 0x63, 0x6e, 0x52, 0x1b, 0x1b,
	0x36, 0x4b, 0x28, 0x17, 0x5e, 0x97, 0xf5, 0xcd, 0x1c, 0xed, 0xa4, 0x20, 0x7a, 0x0d, 0x05, 0x90,
	0x71, 0xcb, 0x72, 0xb4, 0x91, 0x83, 0xe9, 0x8c, 0xe8, 0x00, 0xb6, 0x1e, 0x98, 0x9f, 0x04, 0xd8,
	0x88, 0xc9, 0x23, 0x36, 0x08, 0x35, 0x3c, 0x4b, 0xf8, 0x51, 0xd6, 0x6b, 0xd9, 0xc2, 0x80, 0x3c,
	0x62, 0x8d, 0x5e, 0x58, 0xc8, 0x82, 0x1d, 0xe7, 0x89, 0x8e, 0x11, 0x4e, 0xf8, 0xd4, 0x2b, 0x0d,
	0xa9, 0x59, 0x5d, 0x1c, 0x8c, 0x39, 0x32, 0xe8, 0xdb, 0xce, 0x1c, 0x6d, 0x7e, 0x4a, 0x50, 0x9d,
	0x22, 0x8e, 0xee, 0xe1, 0xbf, 0xa7, 0xb8, 0xc4, 0xd3, 0x69, 0x17, 0x9c, 0xab, 0x8b, 0x9d, 0x9a,
	0x7f, 0x4f, 0xf4, 0x5d, 0x73, 0xfe, 0xfd, 0x79, 0x80, 0x17, 0xd3, 0x4f, 0x8e, 0x11, 0x4d, 0x64,
	0x4f, 0x13, 0xe2, 0x12, 0x4f, 0xe8, 0x57, 0x6d, 0x9f, 0x2c, 0x3a, 0xf1, 0x59, 0xd3, 0xf4, 0x3d,
	0xfe, 0xdc, 0xd2, 0xe9, 0xc9, 0xe7, 0x63, 0x8f, 0xf0, 0xbb, 0xc4, 0x6a, 0xd9, 0x2c, 0x50, 0xfd,
	0xb1, 0xcb, 0xd5, 0xe2, 0x55, 0xf4, 0x30, 0x55, 0x43, 0xeb, 0xc8, 0x63, 0xea, 0xec, 0x43, 0x69,
	0xad, 0x08, 0x07, 0x8e, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x4c, 0xbb, 0x48, 0x43, 0x05,
	0x00, 0x00,
}
