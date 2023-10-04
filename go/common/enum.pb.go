// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: common/enum.proto

package common

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

type ProductCategory int32

const (
	ProductCategory_UnknownCat ProductCategory = 0
	ProductCategory_YouTrip    ProductCategory = 1
)

// Enum value maps for ProductCategory.
var (
	ProductCategory_name = map[int32]string{
		0: "UnknownCat",
		1: "YouTrip",
	}
	ProductCategory_value = map[string]int32{
		"UnknownCat": 0,
		"YouTrip":    1,
	}
)

func (x ProductCategory) Enum() *ProductCategory {
	p := new(ProductCategory)
	*p = x
	return p
}

func (x ProductCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_proto_enumTypes[0].Descriptor()
}

func (ProductCategory) Type() protoreflect.EnumType {
	return &file_common_enum_proto_enumTypes[0]
}

func (x ProductCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductCategory.Descriptor instead.
func (ProductCategory) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_proto_rawDescGZIP(), []int{0}
}

type CardType int32

const (
	CardType_UnKnownCardType CardType = 0
	CardType_PC              CardType = 1
	CardType_NPC             CardType = 2
)

// Enum value maps for CardType.
var (
	CardType_name = map[int32]string{
		0: "UnKnownCardType",
		1: "PC",
		2: "NPC",
	}
	CardType_value = map[string]int32{
		"UnKnownCardType": 0,
		"PC":              1,
		"NPC":             2,
	}
)

func (x CardType) Enum() *CardType {
	p := new(CardType)
	*p = x
	return p
}

func (x CardType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CardType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_proto_enumTypes[1].Descriptor()
}

func (CardType) Type() protoreflect.EnumType {
	return &file_common_enum_proto_enumTypes[1]
}

func (x CardType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CardType.Descriptor instead.
func (CardType) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_proto_rawDescGZIP(), []int{1}
}

type ProgramType int32

const (
	ProgramType_UnKnownProgramType ProgramType = 0
	ProgramType_FIS                ProgramType = 1
)

// Enum value maps for ProgramType.
var (
	ProgramType_name = map[int32]string{
		0: "UnKnownProgramType",
		1: "FIS",
	}
	ProgramType_value = map[string]int32{
		"UnKnownProgramType": 0,
		"FIS":                1,
	}
)

func (x ProgramType) Enum() *ProgramType {
	p := new(ProgramType)
	*p = x
	return p
}

func (x ProgramType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProgramType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_proto_enumTypes[2].Descriptor()
}

func (ProgramType) Type() protoreflect.EnumType {
	return &file_common_enum_proto_enumTypes[2]
}

func (x ProgramType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProgramType.Descriptor instead.
func (ProgramType) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_proto_rawDescGZIP(), []int{2}
}

type SMSStatus int32

const (
	SMSStatus_Unknown     SMSStatus = 0
	SMSStatus_Accepted    SMSStatus = 1
	SMSStatus_Queued      SMSStatus = 2
	SMSStatus_Sending     SMSStatus = 3
	SMSStatus_Sent        SMSStatus = 4
	SMSStatus_Receiving   SMSStatus = 5
	SMSStatus_Received    SMSStatus = 6
	SMSStatus_Delivered   SMSStatus = 7
	SMSStatus_Undelivered SMSStatus = 8
	SMSStatus_Failed      SMSStatus = 9
)

// Enum value maps for SMSStatus.
var (
	SMSStatus_name = map[int32]string{
		0: "Unknown",
		1: "Accepted",
		2: "Queued",
		3: "Sending",
		4: "Sent",
		5: "Receiving",
		6: "Received",
		7: "Delivered",
		8: "Undelivered",
		9: "Failed",
	}
	SMSStatus_value = map[string]int32{
		"Unknown":     0,
		"Accepted":    1,
		"Queued":      2,
		"Sending":     3,
		"Sent":        4,
		"Receiving":   5,
		"Received":    6,
		"Delivered":   7,
		"Undelivered": 8,
		"Failed":      9,
	}
)

func (x SMSStatus) Enum() *SMSStatus {
	p := new(SMSStatus)
	*p = x
	return p
}

func (x SMSStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SMSStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_proto_enumTypes[3].Descriptor()
}

func (SMSStatus) Type() protoreflect.EnumType {
	return &file_common_enum_proto_enumTypes[3]
}

func (x SMSStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SMSStatus.Descriptor instead.
func (SMSStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_proto_rawDescGZIP(), []int{3}
}

type TokenType int32

const (
	TokenType_UnknownTokenType TokenType = 0
	TokenType_Bearer           TokenType = 1
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "UnknownTokenType",
		1: "Bearer",
	}
	TokenType_value = map[string]int32{
		"UnknownTokenType": 0,
		"Bearer":           1,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_proto_enumTypes[4].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_common_enum_proto_enumTypes[4]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_proto_rawDescGZIP(), []int{4}
}

type ProductID int32

const (
	ProductID_UnknownProductID ProductID = 0
	ProductID_THYouTrip        ProductID = 1
	ProductID_SGYouTrip        ProductID = 2
)

// Enum value maps for ProductID.
var (
	ProductID_name = map[int32]string{
		0: "UnknownProductID",
		1: "THYouTrip",
		2: "SGYouTrip",
	}
	ProductID_value = map[string]int32{
		"UnknownProductID": 0,
		"THYouTrip":        1,
		"SGYouTrip":        2,
	}
)

func (x ProductID) Enum() *ProductID {
	p := new(ProductID)
	*p = x
	return p
}

func (x ProductID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductID) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_proto_enumTypes[5].Descriptor()
}

func (ProductID) Type() protoreflect.EnumType {
	return &file_common_enum_proto_enumTypes[5]
}

func (x ProductID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductID.Descriptor instead.
func (ProductID) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_proto_rawDescGZIP(), []int{5}
}

type TokenGroup int32

const (
	TokenGroup_UnknownTokenGroup TokenGroup = 0
	TokenGroup_OTPVerified       TokenGroup = 1
	TokenGroup_AppUnlocked       TokenGroup = 2
)

// Enum value maps for TokenGroup.
var (
	TokenGroup_name = map[int32]string{
		0: "UnknownTokenGroup",
		1: "OTPVerified",
		2: "AppUnlocked",
	}
	TokenGroup_value = map[string]int32{
		"UnknownTokenGroup": 0,
		"OTPVerified":       1,
		"AppUnlocked":       2,
	}
)

func (x TokenGroup) Enum() *TokenGroup {
	p := new(TokenGroup)
	*p = x
	return p
}

func (x TokenGroup) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenGroup) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_proto_enumTypes[6].Descriptor()
}

func (TokenGroup) Type() protoreflect.EnumType {
	return &file_common_enum_proto_enumTypes[6]
}

func (x TokenGroup) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenGroup.Descriptor instead.
func (TokenGroup) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_proto_rawDescGZIP(), []int{6}
}

type TicketType int32

const (
	TicketType_UnknownTicket   TicketType = 0
	TicketType_OrderCardTicket TicketType = 1
)

// Enum value maps for TicketType.
var (
	TicketType_name = map[int32]string{
		0: "UnknownTicket",
		1: "OrderCardTicket",
	}
	TicketType_value = map[string]int32{
		"UnknownTicket":   0,
		"OrderCardTicket": 1,
	}
)

func (x TicketType) Enum() *TicketType {
	p := new(TicketType)
	*p = x
	return p
}

func (x TicketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TicketType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_proto_enumTypes[7].Descriptor()
}

func (TicketType) Type() protoreflect.EnumType {
	return &file_common_enum_proto_enumTypes[7]
}

func (x TicketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TicketType.Descriptor instead.
func (TicketType) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_proto_rawDescGZIP(), []int{7}
}

type SortingDirection int32

const (
	SortingDirection_Ascending  SortingDirection = 0
	SortingDirection_Descending SortingDirection = 1
)

// Enum value maps for SortingDirection.
var (
	SortingDirection_name = map[int32]string{
		0: "Ascending",
		1: "Descending",
	}
	SortingDirection_value = map[string]int32{
		"Ascending":  0,
		"Descending": 1,
	}
)

func (x SortingDirection) Enum() *SortingDirection {
	p := new(SortingDirection)
	*p = x
	return p
}

func (x SortingDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortingDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_proto_enumTypes[8].Descriptor()
}

func (SortingDirection) Type() protoreflect.EnumType {
	return &file_common_enum_proto_enumTypes[8]
}

func (x SortingDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortingDirection.Descriptor instead.
func (SortingDirection) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_proto_rawDescGZIP(), []int{8}
}

type TerminationStatus int32

const (
	TerminationStatus_TerminationStatus_Succeeded TerminationStatus = 0
	TerminationStatus_TerminationStatus_Failed    TerminationStatus = 1
)

// Enum value maps for TerminationStatus.
var (
	TerminationStatus_name = map[int32]string{
		0: "TerminationStatus_Succeeded",
		1: "TerminationStatus_Failed",
	}
	TerminationStatus_value = map[string]int32{
		"TerminationStatus_Succeeded": 0,
		"TerminationStatus_Failed":    1,
	}
)

func (x TerminationStatus) Enum() *TerminationStatus {
	p := new(TerminationStatus)
	*p = x
	return p
}

func (x TerminationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TerminationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_proto_enumTypes[9].Descriptor()
}

func (TerminationStatus) Type() protoreflect.EnumType {
	return &file_common_enum_proto_enumTypes[9]
}

func (x TerminationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TerminationStatus.Descriptor instead.
func (TerminationStatus) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_proto_rawDescGZIP(), []int{9}
}

type FXTargetType int32

const (
	FXTargetType_UnknownFXTargetType FXTargetType = 0
	FXTargetType_SELL                FXTargetType = 1
	FXTargetType_BUY                 FXTargetType = 2
)

// Enum value maps for FXTargetType.
var (
	FXTargetType_name = map[int32]string{
		0: "UnknownFXTargetType",
		1: "SELL",
		2: "BUY",
	}
	FXTargetType_value = map[string]int32{
		"UnknownFXTargetType": 0,
		"SELL":                1,
		"BUY":                 2,
	}
)

func (x FXTargetType) Enum() *FXTargetType {
	p := new(FXTargetType)
	*p = x
	return p
}

func (x FXTargetType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FXTargetType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_enum_proto_enumTypes[10].Descriptor()
}

func (FXTargetType) Type() protoreflect.EnumType {
	return &file_common_enum_proto_enumTypes[10]
}

func (x FXTargetType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FXTargetType.Descriptor instead.
func (FXTargetType) EnumDescriptor() ([]byte, []int) {
	return file_common_enum_proto_rawDescGZIP(), []int{10}
}

var File_common_enum_proto protoreflect.FileDescriptor

var file_common_enum_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2a, 0x2e, 0x0a, 0x0f, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0e,
	0x0a, 0x0a, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x43, 0x61, 0x74, 0x10, 0x00, 0x12, 0x0b,
	0x0a, 0x07, 0x59, 0x6f, 0x75, 0x54, 0x72, 0x69, 0x70, 0x10, 0x01, 0x2a, 0x30, 0x0a, 0x08, 0x43,
	0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x6e, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02,
	0x50, 0x43, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x50, 0x43, 0x10, 0x02, 0x2a, 0x2e, 0x0a,
	0x0b, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12,
	0x55, 0x6e, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x54, 0x79,
	0x70, 0x65, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x46, 0x49, 0x53, 0x10, 0x01, 0x2a, 0x92, 0x01,
	0x0a, 0x09, 0x53, 0x4d, 0x53, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x75, 0x65, 0x75, 0x65, 0x64,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x03, 0x12,
	0x08, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x74, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x65, 0x63,
	0x65, 0x69, 0x76, 0x69, 0x6e, 0x67, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x64, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x65, 0x64, 0x10, 0x07, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x65, 0x64, 0x10, 0x08, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0x09, 0x2a, 0x2d, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x10, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x10,
	0x01, 0x2a, 0x3f, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x10, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x48, 0x59, 0x6f, 0x75, 0x54, 0x72, 0x69,
	0x70, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x47, 0x59, 0x6f, 0x75, 0x54, 0x72, 0x69, 0x70,
	0x10, 0x02, 0x2a, 0x45, 0x0a, 0x0a, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x15, 0x0a, 0x11, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x54, 0x50, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x55,
	0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x10, 0x02, 0x2a, 0x34, 0x0a, 0x0a, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x10, 0x01, 0x2a,
	0x31, 0x0a, 0x10, 0x53, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x44, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x10, 0x01, 0x2a, 0x52, 0x0a, 0x11, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x46, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x10, 0x01, 0x2a, 0x3a, 0x0a, 0x0c, 0x46, 0x58, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x46, 0x58, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x53, 0x45, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x55, 0x59,
	0x10, 0x02, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x79, 0x65, 0x72, 0x79, 0x61, 0x6e, 0x2f, 0x6c, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_enum_proto_rawDescOnce sync.Once
	file_common_enum_proto_rawDescData = file_common_enum_proto_rawDesc
)

func file_common_enum_proto_rawDescGZIP() []byte {
	file_common_enum_proto_rawDescOnce.Do(func() {
		file_common_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_enum_proto_rawDescData)
	})
	return file_common_enum_proto_rawDescData
}

var file_common_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 11)
var file_common_enum_proto_goTypes = []interface{}{
	(ProductCategory)(0),   // 0: common.ProductCategory
	(CardType)(0),          // 1: common.CardType
	(ProgramType)(0),       // 2: common.ProgramType
	(SMSStatus)(0),         // 3: common.SMSStatus
	(TokenType)(0),         // 4: common.TokenType
	(ProductID)(0),         // 5: common.ProductID
	(TokenGroup)(0),        // 6: common.TokenGroup
	(TicketType)(0),        // 7: common.TicketType
	(SortingDirection)(0),  // 8: common.SortingDirection
	(TerminationStatus)(0), // 9: common.TerminationStatus
	(FXTargetType)(0),      // 10: common.FXTargetType
}
var file_common_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_enum_proto_init() }
func file_common_enum_proto_init() {
	if File_common_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_enum_proto_rawDesc,
			NumEnums:      11,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_enum_proto_goTypes,
		DependencyIndexes: file_common_enum_proto_depIdxs,
		EnumInfos:         file_common_enum_proto_enumTypes,
	}.Build()
	File_common_enum_proto = out.File
	file_common_enum_proto_rawDesc = nil
	file_common_enum_proto_goTypes = nil
	file_common_enum_proto_depIdxs = nil
}