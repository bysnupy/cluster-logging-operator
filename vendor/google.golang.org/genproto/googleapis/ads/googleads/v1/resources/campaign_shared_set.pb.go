// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/campaign_shared_set.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// CampaignSharedSets are used for managing the shared sets associated with a
// campaign.
type CampaignSharedSet struct {
	// The resource name of the campaign shared set.
	// Campaign shared set resource names have the form:
	//
	// `customers/{customer_id}/campaignSharedSets/{campaign_id}~{shared_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign to which the campaign shared set belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The shared set associated with the campaign. This may be a negative keyword
	// shared set of another customer. This customer should be a manager of the
	// other customer, otherwise the campaign shared set will exist but have no
	// serving effect. Only negative keyword shared sets can be associated with
	// Shopping campaigns. Only negative placement shared sets can be associated
	// with Display mobile app campaigns.
	SharedSet *wrappers.StringValue `protobuf:"bytes,4,opt,name=shared_set,json=sharedSet,proto3" json:"shared_set,omitempty"`
	// The status of this campaign shared set. Read only.
	Status               enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus `protobuf:"varint,2,opt,name=status,proto3,enum=google.ads.googleads.v1.enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                  `json:"-"`
	XXX_unrecognized     []byte                                                    `json:"-"`
	XXX_sizecache        int32                                                     `json:"-"`
}

func (m *CampaignSharedSet) Reset()         { *m = CampaignSharedSet{} }
func (m *CampaignSharedSet) String() string { return proto.CompactTextString(m) }
func (*CampaignSharedSet) ProtoMessage()    {}
func (*CampaignSharedSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c77b63cdc30c4b2, []int{0}
}

func (m *CampaignSharedSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignSharedSet.Unmarshal(m, b)
}
func (m *CampaignSharedSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignSharedSet.Marshal(b, m, deterministic)
}
func (m *CampaignSharedSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignSharedSet.Merge(m, src)
}
func (m *CampaignSharedSet) XXX_Size() int {
	return xxx_messageInfo_CampaignSharedSet.Size(m)
}
func (m *CampaignSharedSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignSharedSet.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignSharedSet proto.InternalMessageInfo

func (m *CampaignSharedSet) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignSharedSet) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignSharedSet) GetSharedSet() *wrappers.StringValue {
	if m != nil {
		return m.SharedSet
	}
	return nil
}

func (m *CampaignSharedSet) GetStatus() enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus {
	if m != nil {
		return m.Status
	}
	return enums.CampaignSharedSetStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignSharedSet)(nil), "google.ads.googleads.v1.resources.CampaignSharedSet")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/campaign_shared_set.proto", fileDescriptor_5c77b63cdc30c4b2)
}

var fileDescriptor_5c77b63cdc30c4b2 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6a, 0xe3, 0x30,
	0x18, 0xc4, 0xce, 0x12, 0x36, 0xda, 0x1f, 0x58, 0x1f, 0x16, 0x13, 0xc2, 0x92, 0x6c, 0x09, 0xe4,
	0x24, 0xe3, 0xf4, 0x52, 0x14, 0x28, 0x38, 0xa5, 0x04, 0x7a, 0x28, 0xc1, 0x06, 0x1f, 0x8a, 0xc1,
	0x28, 0xb1, 0xaa, 0xba, 0xc4, 0x92, 0x91, 0xe4, 0xf4, 0x01, 0xfa, 0x18, 0xbd, 0xf5, 0xd8, 0x47,
	0xe9, 0xa3, 0xf4, 0x29, 0x4a, 0x6c, 0x4b, 0x3d, 0xa4, 0x69, 0x7b, 0x1b, 0x7b, 0x66, 0xbe, 0xf9,
	0x46, 0x12, 0x98, 0x51, 0xce, 0xe9, 0x86, 0x78, 0x38, 0x93, 0x5e, 0x03, 0x77, 0x68, 0xeb, 0x7b,
	0x82, 0x48, 0x5e, 0x89, 0x35, 0x91, 0xde, 0x1a, 0x17, 0x25, 0xce, 0x29, 0x4b, 0xe5, 0x0d, 0x16,
	0x24, 0x4b, 0x25, 0x51, 0xb0, 0x14, 0x5c, 0x71, 0x67, 0xd4, 0x38, 0x20, 0xce, 0x24, 0x34, 0x66,
	0xb8, 0xf5, 0xa1, 0x31, 0xf7, 0x4f, 0x0f, 0xcd, 0x27, 0xac, 0x2a, 0xde, 0x9d, 0x9d, 0x4a, 0x85,
	0x55, 0x25, 0x9b, 0x88, 0xfe, 0x40, 0xfb, 0xcb, 0xdc, 0xc3, 0x8c, 0x71, 0x85, 0x55, 0xce, 0x99,
	0x66, 0xff, 0xb5, 0x6c, 0xfd, 0xb5, 0xaa, 0xae, 0xbd, 0x3b, 0x81, 0xcb, 0x92, 0x88, 0x96, 0xff,
	0xff, 0x60, 0x83, 0x3f, 0x67, 0x6d, 0x44, 0x54, 0x27, 0x44, 0x44, 0x39, 0x47, 0xe0, 0x97, 0x5e,
	0x30, 0x65, 0xb8, 0x20, 0xae, 0x35, 0xb4, 0x26, 0xbd, 0xf0, 0xa7, 0xfe, 0x79, 0x89, 0x0b, 0xe2,
	0x9c, 0x80, 0xef, 0x7a, 0x39, 0xb7, 0x33, 0xb4, 0x26, 0x3f, 0xa6, 0x83, 0xb6, 0x23, 0xd4, 0x69,
	0x30, 0x52, 0x22, 0x67, 0x34, 0xc6, 0x9b, 0x8a, 0x84, 0x46, 0xed, 0xcc, 0x00, 0x78, 0x6b, 0xe3,
	0x7e, 0xfb, 0x82, 0xb7, 0x27, 0xcd, 0x6e, 0xb7, 0xa0, 0xdb, 0xf4, 0x77, 0xed, 0xa1, 0x35, 0xf9,
	0x3d, 0x0d, 0xe1, 0xa1, 0x33, 0xae, 0x0f, 0x10, 0xee, 0xb5, 0x8b, 0x6a, 0xf7, 0x39, 0xab, 0x8a,
	0x43, 0x5c, 0xd8, 0x26, 0xcc, 0xef, 0x6d, 0x30, 0x5e, 0xf3, 0x02, 0x7e, 0x7a, 0x8b, 0xf3, 0xbf,
	0x7b, 0xa3, 0x96, 0xbb, 0x1e, 0x4b, 0xeb, 0xea, 0xa2, 0x35, 0x53, 0xbe, 0xc1, 0x8c, 0x42, 0x2e,
	0xa8, 0x47, 0x09, 0xab, 0x5b, 0xea, 0xfb, 0x2e, 0x73, 0xf9, 0xc1, 0xf3, 0x9a, 0x19, 0xf4, 0x68,
	0x77, 0x16, 0x41, 0xf0, 0x64, 0x8f, 0x16, 0xcd, 0xc8, 0x20, 0x93, 0xb0, 0x81, 0x3b, 0x14, 0xfb,
	0x30, 0xd4, 0xca, 0x67, 0xad, 0x49, 0x82, 0x4c, 0x26, 0x46, 0x93, 0xc4, 0x7e, 0x62, 0x34, 0x2f,
	0xf6, 0xb8, 0x21, 0x10, 0x0a, 0x32, 0x89, 0x90, 0x51, 0x21, 0x14, 0xfb, 0x08, 0x19, 0xdd, 0xaa,
	0x5b, 0x2f, 0x7b, 0xfc, 0x1a, 0x00, 0x00, 0xff, 0xff, 0x85, 0xca, 0xc7, 0x69, 0x0a, 0x03, 0x00,
	0x00,
}