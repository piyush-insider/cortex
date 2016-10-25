// Code generated by protoc-gen-go.
// source: cortex.proto
// DO NOT EDIT!

/*
Package cortex is a generated protocol buffer package.

It is generated from these files:
	cortex.proto

It has these top-level messages:
	Sample
	LabelPair
	TimeSeries
	LabelMatcher
	ReadRequest
	ReadResponse
	LabelValuesRequest
	LabelValuesResponse
*/
package cortex

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MatchType int32

const (
	MatchType_EQUAL          MatchType = 0
	MatchType_NOT_EQUAL      MatchType = 1
	MatchType_REGEX_MATCH    MatchType = 2
	MatchType_REGEX_NO_MATCH MatchType = 3
)

var MatchType_name = map[int32]string{
	0: "EQUAL",
	1: "NOT_EQUAL",
	2: "REGEX_MATCH",
	3: "REGEX_NO_MATCH",
}
var MatchType_value = map[string]int32{
	"EQUAL":          0,
	"NOT_EQUAL":      1,
	"REGEX_MATCH":    2,
	"REGEX_NO_MATCH": 3,
}

func (x MatchType) String() string {
	return proto.EnumName(MatchType_name, int32(x))
}
func (MatchType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Sample struct {
	Value       float64 `protobuf:"fixed64,1,opt,name=value" json:"value,omitempty"`
	TimestampMs int64   `protobuf:"varint,2,opt,name=timestamp_ms,json=timestampMs" json:"timestamp_ms,omitempty"`
}

func (m *Sample) Reset()                    { *m = Sample{} }
func (m *Sample) String() string            { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()               {}
func (*Sample) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LabelPair struct {
	Name  string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *LabelPair) Reset()                    { *m = LabelPair{} }
func (m *LabelPair) String() string            { return proto.CompactTextString(m) }
func (*LabelPair) ProtoMessage()               {}
func (*LabelPair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type TimeSeries struct {
	Labels []*LabelPair `protobuf:"bytes,1,rep,name=labels" json:"labels,omitempty"`
	// Sorted by time, oldest sample first.
	Samples []*Sample `protobuf:"bytes,2,rep,name=samples" json:"samples,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TimeSeries) GetLabels() []*LabelPair {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TimeSeries) GetSamples() []*Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type LabelMatcher struct {
	Type  MatchType `protobuf:"varint,1,opt,name=type,enum=cortex.MatchType" json:"type,omitempty"`
	Name  string    `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value string    `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *LabelMatcher) Reset()                    { *m = LabelMatcher{} }
func (m *LabelMatcher) String() string            { return proto.CompactTextString(m) }
func (*LabelMatcher) ProtoMessage()               {}
func (*LabelMatcher) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ReadRequest struct {
	StartTimestampMs int64           `protobuf:"varint,1,opt,name=start_timestamp_ms,json=startTimestampMs" json:"start_timestamp_ms,omitempty"`
	EndTimestampMs   int64           `protobuf:"varint,2,opt,name=end_timestamp_ms,json=endTimestampMs" json:"end_timestamp_ms,omitempty"`
	Matchers         []*LabelMatcher `protobuf:"bytes,3,rep,name=matchers" json:"matchers,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadRequest) GetMatchers() []*LabelMatcher {
	if m != nil {
		return m.Matchers
	}
	return nil
}

type ReadResponse struct {
	Timeseries []*TimeSeries `protobuf:"bytes,1,rep,name=timeseries" json:"timeseries,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ReadResponse) GetTimeseries() []*TimeSeries {
	if m != nil {
		return m.Timeseries
	}
	return nil
}

type LabelValuesRequest struct {
	LabelName string `protobuf:"bytes,1,opt,name=label_name,json=labelName" json:"label_name,omitempty"`
}

func (m *LabelValuesRequest) Reset()                    { *m = LabelValuesRequest{} }
func (m *LabelValuesRequest) String() string            { return proto.CompactTextString(m) }
func (*LabelValuesRequest) ProtoMessage()               {}
func (*LabelValuesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type LabelValuesResponse struct {
	LabelValues []string `protobuf:"bytes,1,rep,name=label_values,json=labelValues" json:"label_values,omitempty"`
}

func (m *LabelValuesResponse) Reset()                    { *m = LabelValuesResponse{} }
func (m *LabelValuesResponse) String() string            { return proto.CompactTextString(m) }
func (*LabelValuesResponse) ProtoMessage()               {}
func (*LabelValuesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func init() {
	proto.RegisterType((*Sample)(nil), "cortex.Sample")
	proto.RegisterType((*LabelPair)(nil), "cortex.LabelPair")
	proto.RegisterType((*TimeSeries)(nil), "cortex.TimeSeries")
	proto.RegisterType((*LabelMatcher)(nil), "cortex.LabelMatcher")
	proto.RegisterType((*ReadRequest)(nil), "cortex.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "cortex.ReadResponse")
	proto.RegisterType((*LabelValuesRequest)(nil), "cortex.LabelValuesRequest")
	proto.RegisterType((*LabelValuesResponse)(nil), "cortex.LabelValuesResponse")
	proto.RegisterEnum("cortex.MatchType", MatchType_name, MatchType_value)
}

func init() { proto.RegisterFile("cortex.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x52, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0x76, 0x2f, 0xed, 0x69, 0x26, 0xf1, 0x3c, 0xc7, 0x3e, 0xdc, 0x8b, 0x70, 0x0d, 0x08, 0x51,
	0xa4, 0x48, 0x8b, 0xe0, 0xeb, 0x29, 0x87, 0x22, 0xbd, 0xab, 0x6e, 0xa3, 0xf8, 0x16, 0xb6, 0xbd,
	0x01, 0x03, 0xd9, 0x24, 0x66, 0xb7, 0x62, 0xff, 0x12, 0xff, 0xdd, 0x92, 0xd9, 0xfc, 0x3a, 0xb8,
	0xb7, 0xcc, 0xcc, 0xb7, 0xdf, 0x7c, 0xdf, 0x37, 0x81, 0xf0, 0xb6, 0xac, 0x2d, 0xfd, 0x3b, 0xab,
	0xea, 0xd2, 0x96, 0x38, 0x75, 0x55, 0xb4, 0x82, 0xe9, 0xb5, 0xd2, 0x55, 0x4e, 0x78, 0x02, 0xc7,
	0x7f, 0x55, 0x7e, 0x47, 0x0b, 0xb1, 0x14, 0xb1, 0x90, 0xae, 0xc0, 0x53, 0x08, 0x6d, 0xa6, 0xc9,
	0x58, 0xa5, 0xab, 0x54, 0x9b, 0xc5, 0x64, 0x29, 0x62, 0x4f, 0x06, 0x7d, 0x6f, 0x63, 0xa2, 0xf7,
	0xe0, 0x5f, 0xaa, 0x1b, 0xca, 0xbf, 0xa9, 0xac, 0x46, 0x84, 0xa3, 0x42, 0x69, 0x47, 0xe2, 0x4b,
	0xfe, 0x1e, 0x98, 0x27, 0xdc, 0x74, 0x45, 0xa4, 0x00, 0x92, 0x4c, 0xd3, 0x35, 0xd5, 0x19, 0x19,
	0x7c, 0x0d, 0xd3, 0xbc, 0x21, 0x31, 0x0b, 0xb1, 0xf4, 0xe2, 0xe0, 0xfc, 0xf9, 0x59, 0x2b, 0xb7,
	0xa7, 0x96, 0x2d, 0x00, 0x63, 0x78, 0x6c, 0x58, 0x72, 0xa3, 0xa6, 0xc1, 0xce, 0x3a, 0xac, 0x73,
	0x22, 0xbb, 0x71, 0x94, 0x42, 0xc8, 0xcf, 0x37, 0xca, 0xde, 0xfe, 0xa6, 0x1a, 0x5f, 0xc1, 0x91,
	0xbd, 0xaf, 0x9c, 0xb8, 0xd9, 0xb0, 0x82, 0xc7, 0xc9, 0x7d, 0x45, 0x92, 0xc7, 0xbd, 0x87, 0xc9,
	0x21, 0x0f, 0xde, 0xd8, 0xc3, 0x7f, 0x01, 0x81, 0x24, 0xb5, 0x93, 0xf4, 0xe7, 0x8e, 0x8c, 0xc5,
	0xb7, 0x80, 0xc6, 0xaa, 0xda, 0xa6, 0x7b, 0x99, 0x09, 0xce, 0x6c, 0xce, 0x93, 0x64, 0x08, 0x0e,
	0x63, 0x98, 0x53, 0xb1, 0x4b, 0x0f, 0xe4, 0x3b, 0xa3, 0x62, 0x37, 0x46, 0xbe, 0x83, 0x27, 0xda,
	0x79, 0x30, 0x0b, 0x8f, 0x3d, 0x9f, 0xec, 0xe5, 0xd3, 0x1a, 0x94, 0x3d, 0x2a, 0xfa, 0x08, 0xa1,
	0x13, 0x66, 0xaa, 0xb2, 0x30, 0x84, 0xe7, 0x00, 0xbc, 0x87, 0xd3, 0x6e, 0x33, 0xc6, 0x8e, 0x63,
	0xb8, 0x83, 0x1c, 0xa1, 0xa2, 0x0b, 0x40, 0x66, 0xff, 0xd9, 0x78, 0x35, 0x9d, 0xc7, 0x97, 0x00,
	0x7c, 0x88, 0x74, 0x74, 0x67, 0x9f, 0x3b, 0x5b, 0xa5, 0x29, 0xfa, 0x00, 0x2f, 0xf6, 0x1e, 0xb5,
	0xfb, 0x4f, 0x21, 0x74, 0xaf, 0x38, 0x38, 0xa7, 0xc0, 0x97, 0x41, 0x3e, 0x40, 0xdf, 0x7c, 0x05,
	0xbf, 0xbf, 0x04, 0xfa, 0x70, 0xbc, 0xfe, 0xfe, 0x63, 0x75, 0x39, 0x7f, 0x84, 0x4f, 0xc1, 0xdf,
	0x5e, 0x25, 0xa9, 0x2b, 0x05, 0x3e, 0x83, 0x40, 0xae, 0x3f, 0xaf, 0x7f, 0xa5, 0x9b, 0x55, 0xf2,
	0xe9, 0xcb, 0x7c, 0x82, 0x08, 0x33, 0xd7, 0xd8, 0x5e, 0xb5, 0x3d, 0xef, 0x66, 0xca, 0x7f, 0xf9,
	0xc5, 0x43, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaa, 0xf2, 0xff, 0x3d, 0xf5, 0x02, 0x00, 0x00,
}