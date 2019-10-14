// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/domain_category.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A category generated automatically by crawling a domain. If a campaign uses
// the DynamicSearchAdsSetting, then domain categories will be generated for
// the domain. The categories can be targeted using WebpageConditionInfo.
// See: https://support.google.com/google-ads/answer/2471185
type DomainCategory struct {
	// The resource name of the domain category.
	// Domain category resource names have the form:
	//
	//
	// `customers/{customer_id}/domainCategories/{campaign_id}~{category_base64}~{language_code}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign this category is recommended for.
	Campaign *wrappers.StringValue `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Recommended category for the website domain. e.g. if you have a website
	// about electronics, the categories could be "cameras", "televisions", etc.
	Category *wrappers.StringValue `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	// The language code specifying the language of the website. e.g. "en" for
	// English. The language can be specified in the DynamicSearchAdsSetting
	// required for dynamic search ads. This is the language of the pages from
	// your website that you want Google Ads to find, create ads for,
	// and match searches with.
	LanguageCode *wrappers.StringValue `protobuf:"bytes,4,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// The domain for the website. The domain can be specified in the
	// DynamicSearchAdsSetting required for dynamic search ads.
	Domain *wrappers.StringValue `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`
	// Fraction of pages on your site that this category matches.
	CoverageFraction *wrappers.DoubleValue `protobuf:"bytes,6,opt,name=coverage_fraction,json=coverageFraction,proto3" json:"coverage_fraction,omitempty"`
	// The position of this category in the set of categories. Lower numbers
	// indicate a better match for the domain. null indicates not recommended.
	CategoryRank *wrappers.Int64Value `protobuf:"bytes,7,opt,name=category_rank,json=categoryRank,proto3" json:"category_rank,omitempty"`
	// Indicates whether this category has sub-categories.
	HasChildren *wrappers.BoolValue `protobuf:"bytes,8,opt,name=has_children,json=hasChildren,proto3" json:"has_children,omitempty"`
	// The recommended cost per click for the category.
	RecommendedCpcBidMicros *wrappers.Int64Value `protobuf:"bytes,9,opt,name=recommended_cpc_bid_micros,json=recommendedCpcBidMicros,proto3" json:"recommended_cpc_bid_micros,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}             `json:"-"`
	XXX_unrecognized        []byte               `json:"-"`
	XXX_sizecache           int32                `json:"-"`
}

func (m *DomainCategory) Reset()         { *m = DomainCategory{} }
func (m *DomainCategory) String() string { return proto.CompactTextString(m) }
func (*DomainCategory) ProtoMessage()    {}
func (*DomainCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_domain_category_695d83200adb3248, []int{0}
}
func (m *DomainCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DomainCategory.Unmarshal(m, b)
}
func (m *DomainCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DomainCategory.Marshal(b, m, deterministic)
}
func (dst *DomainCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DomainCategory.Merge(dst, src)
}
func (m *DomainCategory) XXX_Size() int {
	return xxx_messageInfo_DomainCategory.Size(m)
}
func (m *DomainCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_DomainCategory.DiscardUnknown(m)
}

var xxx_messageInfo_DomainCategory proto.InternalMessageInfo

func (m *DomainCategory) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *DomainCategory) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *DomainCategory) GetCategory() *wrappers.StringValue {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *DomainCategory) GetLanguageCode() *wrappers.StringValue {
	if m != nil {
		return m.LanguageCode
	}
	return nil
}

func (m *DomainCategory) GetDomain() *wrappers.StringValue {
	if m != nil {
		return m.Domain
	}
	return nil
}

func (m *DomainCategory) GetCoverageFraction() *wrappers.DoubleValue {
	if m != nil {
		return m.CoverageFraction
	}
	return nil
}

func (m *DomainCategory) GetCategoryRank() *wrappers.Int64Value {
	if m != nil {
		return m.CategoryRank
	}
	return nil
}

func (m *DomainCategory) GetHasChildren() *wrappers.BoolValue {
	if m != nil {
		return m.HasChildren
	}
	return nil
}

func (m *DomainCategory) GetRecommendedCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.RecommendedCpcBidMicros
	}
	return nil
}

func init() {
	proto.RegisterType((*DomainCategory)(nil), "google.ads.googleads.v1.resources.DomainCategory")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/domain_category.proto", fileDescriptor_domain_category_695d83200adb3248)
}

var fileDescriptor_domain_category_695d83200adb3248 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xd1, 0x6a, 0x14, 0x3d,
	0x1c, 0xc5, 0xd9, 0x6d, 0xbf, 0xfd, 0xda, 0x74, 0x57, 0x74, 0xbc, 0x70, 0x58, 0x8b, 0xb4, 0x4a,
	0xa1, 0x57, 0x19, 0x46, 0x8b, 0xca, 0x88, 0xe0, 0xec, 0x16, 0x4b, 0x05, 0xa5, 0xac, 0xb0, 0x88,
	0x2c, 0x0c, 0xd9, 0x24, 0xcd, 0x86, 0xce, 0xe4, 0x3f, 0x24, 0x33, 0x2b, 0xde, 0xf9, 0x2c, 0x5e,
	0xfa, 0x24, 0xe2, 0xa3, 0xf8, 0x14, 0x32, 0x93, 0x64, 0x50, 0x8a, 0x76, 0xef, 0xfe, 0x4c, 0xce,
	0xef, 0x9c, 0x93, 0x4c, 0x82, 0x9e, 0x09, 0x00, 0x91, 0xf3, 0x88, 0x30, 0x13, 0xd9, 0xb1, 0x99,
	0xd6, 0x71, 0xa4, 0xb9, 0x81, 0x5a, 0x53, 0x6e, 0x22, 0x06, 0x05, 0x91, 0x2a, 0xa3, 0xa4, 0xe2,
	0x02, 0xf4, 0x67, 0x5c, 0x6a, 0xa8, 0x20, 0x38, 0xb4, 0x6a, 0x4c, 0x98, 0xc1, 0x1d, 0x88, 0xd7,
	0x31, 0xee, 0xc0, 0xf1, 0x03, 0xe7, 0xdd, 0x02, 0xcb, 0xfa, 0x32, 0xfa, 0xa4, 0x49, 0x59, 0x72,
	0x6d, 0xac, 0xc5, 0x78, 0xdf, 0x67, 0x97, 0x32, 0x22, 0x4a, 0x41, 0x45, 0x2a, 0x09, 0xca, 0xad,
	0x3e, 0xfc, 0xbe, 0x8d, 0x6e, 0x9d, 0xb6, 0xd1, 0x53, 0x97, 0x1c, 0x3c, 0x42, 0x23, 0xef, 0x9e,
	0x29, 0x52, 0xf0, 0xb0, 0x77, 0xd0, 0x3b, 0xde, 0x9d, 0x0d, 0xfd, 0xc7, 0x77, 0xa4, 0xe0, 0xc1,
	0x73, 0xb4, 0x43, 0x49, 0x51, 0x12, 0x29, 0x54, 0xd8, 0x3f, 0xe8, 0x1d, 0xef, 0x3d, 0xde, 0x77,
	0x05, 0xb1, 0x2f, 0x82, 0xdf, 0x57, 0x5a, 0x2a, 0x31, 0x27, 0x79, 0xcd, 0x67, 0x9d, 0xda, 0x92,
	0x36, 0x2a, 0xdc, 0xda, 0x8c, 0x74, 0xc5, 0x52, 0x34, 0xca, 0x89, 0x12, 0x35, 0x11, 0x3c, 0xa3,
	0xc0, 0x78, 0xb8, 0xbd, 0x01, 0x3e, 0xf4, 0xc8, 0x14, 0x18, 0x0f, 0x4e, 0xd0, 0xc0, 0x1e, 0x74,
	0xf8, 0xdf, 0x06, 0xac, 0xd3, 0x06, 0xe7, 0xe8, 0x0e, 0x85, 0x35, 0xd7, 0x4d, 0xf0, 0xa5, 0x26,
	0xb4, 0x39, 0xc0, 0x70, 0xf0, 0x17, 0x83, 0x53, 0xa8, 0x97, 0x39, 0xb7, 0x06, 0xb7, 0x3d, 0xf6,
	0xda, 0x51, 0xc1, 0x2b, 0x34, 0xf2, 0xfb, 0xc9, 0x34, 0x51, 0x57, 0xe1, 0xff, 0xad, 0xcd, 0xfd,
	0x6b, 0x36, 0xe7, 0xaa, 0x7a, 0x7a, 0xe2, 0xb6, 0xe0, 0x89, 0x19, 0x51, 0x57, 0xc1, 0x4b, 0x34,
	0x5c, 0x11, 0x93, 0xd1, 0x95, 0xcc, 0x99, 0xe6, 0x2a, 0xdc, 0x69, 0x0d, 0xc6, 0xd7, 0x0c, 0x26,
	0x00, 0xb9, 0xe5, 0xf7, 0x56, 0xc4, 0x4c, 0x9d, 0x3c, 0xf8, 0x80, 0xc6, 0x9a, 0x53, 0x28, 0x0a,
	0xae, 0x18, 0x67, 0x19, 0x2d, 0x69, 0xb6, 0x94, 0x2c, 0x2b, 0x24, 0xd5, 0x60, 0xc2, 0xdd, 0x9b,
	0xdb, 0xdc, 0xfb, 0x0d, 0x9f, 0x96, 0x74, 0x22, 0xd9, 0xdb, 0x96, 0x9d, 0x7c, 0xe9, 0xa3, 0x23,
	0x0a, 0x05, 0xbe, 0xf1, 0xca, 0x4e, 0xee, 0xfe, 0x79, 0xe3, 0x2e, 0x9a, 0x94, 0x8b, 0xde, 0xc7,
	0x37, 0x8e, 0x14, 0xd0, 0xfc, 0x33, 0x0c, 0x5a, 0x44, 0x82, 0xab, 0xb6, 0x83, 0x7f, 0x35, 0xa5,
	0x34, 0xff, 0x78, 0x44, 0x2f, 0xba, 0xe9, 0x6b, 0x7f, 0xeb, 0x2c, 0x4d, 0xbf, 0xf5, 0x0f, 0xcf,
	0xac, 0x65, 0xca, 0x0c, 0xb6, 0x63, 0x33, 0xcd, 0x63, 0x3c, 0xf3, 0xca, 0x1f, 0x5e, 0xb3, 0x48,
	0x99, 0x59, 0x74, 0x9a, 0xc5, 0x3c, 0x5e, 0x74, 0x9a, 0x9f, 0xfd, 0x23, 0xbb, 0x90, 0x24, 0x29,
	0x33, 0x49, 0xd2, 0xa9, 0x92, 0x64, 0x1e, 0x27, 0x49, 0xa7, 0x5b, 0x0e, 0xda, 0xb2, 0x4f, 0x7e,
	0x05, 0x00, 0x00, 0xff, 0xff, 0xce, 0x80, 0xc4, 0x0b, 0xf0, 0x03, 0x00, 0x00,
}