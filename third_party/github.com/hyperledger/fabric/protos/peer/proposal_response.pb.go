/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/proposal_response.proto

package peer // import "github.com/thanakritlee/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A ProposalResponse is returned from an endorser to the proposal submitter.
// The idea is that this message contains the endorser's response to the
// request of a client to perform an action over a chaincode (or more
// generically on the ledger); the response might be success/error (conveyed in
// the Response field) together with a description of the action and a
// signature over it by that endorser.  If a sufficient number of distinct
// endorsers agree on the same action and produce signature to that effect, a
// transaction can be generated and sent for ordering.
type ProposalResponse struct {
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Timestamp is the time that the message
	// was created as  defined by the sender
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// A response message indicating whether the
	// endorsement of the action was successful
	Response *Response `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
	// The payload of response. It is the bytes of ProposalResponsePayload
	Payload []byte `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	// The endorsement of the proposal, basically
	// the endorser's signature over the payload
	Endorsement          *Endorsement `protobuf:"bytes,6,opt,name=endorsement,proto3" json:"endorsement,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ProposalResponse) Reset()         { *m = ProposalResponse{} }
func (m *ProposalResponse) String() string { return proto.CompactTextString(m) }
func (*ProposalResponse) ProtoMessage()    {}
func (*ProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_proposal_response_22a755721b685f40, []int{0}
}
func (m *ProposalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalResponse.Unmarshal(m, b)
}
func (m *ProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalResponse.Marshal(b, m, deterministic)
}
func (dst *ProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalResponse.Merge(dst, src)
}
func (m *ProposalResponse) XXX_Size() int {
	return xxx_messageInfo_ProposalResponse.Size(m)
}
func (m *ProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalResponse proto.InternalMessageInfo

func (m *ProposalResponse) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *ProposalResponse) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ProposalResponse) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ProposalResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ProposalResponse) GetEndorsement() *Endorsement {
	if m != nil {
		return m.Endorsement
	}
	return nil
}

// A response with a representation similar to an HTTP response that can
// be used within another message.
type Response struct {
	// A status code that should follow the HTTP status codes.
	Status int32 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	// A message associated with the response code.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// A payload that can be used to include metadata with this response.
	Payload              []byte   `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_proposal_response_22a755721b685f40, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Response) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// ProposalResponsePayload is the payload of a proposal response.  This message
// is the "bridge" between the client's request and the endorser's action in
// response to that request. Concretely, for chaincodes, it contains a hashed
// representation of the proposal (proposalHash) and a representation of the
// chaincode state changes and events inside the extension field.
type ProposalResponsePayload struct {
	// Hash of the proposal that triggered this response. The hash is used to
	// link a response with its proposal, both for bookeeping purposes on an
	// asynchronous system and for security reasons (accountability,
	// non-repudiation). The hash usually covers the entire Proposal message
	// (byte-by-byte). However this implies that the hash can only be verified
	// if the entire proposal message is available when ProposalResponsePayload is
	// included in a transaction or stored in the ledger. For confidentiality
	// reasons, with chaincodes it might be undesirable to store the proposal
	// payload in the ledger.  If the type is CHAINCODE, this is handled by
	// separating the proposal's header and
	// the payload: the header is always hashed in its entirety whereas the
	// payload can either be hashed fully, or only its hash may be hashed, or
	// nothing from the payload can be hashed. The PayloadVisibility field in the
	// Header's extension controls to which extent the proposal payload is
	// "visible" in the sense that was just explained.
	ProposalHash []byte `protobuf:"bytes,1,opt,name=proposal_hash,json=proposalHash,proto3" json:"proposal_hash,omitempty"`
	// Extension should be unmarshaled to a type-specific message. The type of
	// the extension in any proposal response depends on the type of the proposal
	// that the client selected when the proposal was initially sent out.  In
	// particular, this information is stored in the type field of a Header.  For
	// chaincode, it's a ChaincodeAction message
	Extension            []byte   `protobuf:"bytes,2,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProposalResponsePayload) Reset()         { *m = ProposalResponsePayload{} }
func (m *ProposalResponsePayload) String() string { return proto.CompactTextString(m) }
func (*ProposalResponsePayload) ProtoMessage()    {}
func (*ProposalResponsePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_proposal_response_22a755721b685f40, []int{2}
}
func (m *ProposalResponsePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProposalResponsePayload.Unmarshal(m, b)
}
func (m *ProposalResponsePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProposalResponsePayload.Marshal(b, m, deterministic)
}
func (dst *ProposalResponsePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposalResponsePayload.Merge(dst, src)
}
func (m *ProposalResponsePayload) XXX_Size() int {
	return xxx_messageInfo_ProposalResponsePayload.Size(m)
}
func (m *ProposalResponsePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposalResponsePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ProposalResponsePayload proto.InternalMessageInfo

func (m *ProposalResponsePayload) GetProposalHash() []byte {
	if m != nil {
		return m.ProposalHash
	}
	return nil
}

func (m *ProposalResponsePayload) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

// An endorsement is a signature of an endorser over a proposal response.  By
// producing an endorsement message, an endorser implicitly "approves" that
// proposal response and the actions contained therein. When enough
// endorsements have been collected, a transaction can be generated out of a
// set of proposal responses.  Note that this message only contains an identity
// and a signature but no signed payload. This is intentional because
// endorsements are supposed to be collected in a transaction, and they are all
// expected to endorse a single proposal response/action (many endorsements
// over a single proposal response)
type Endorsement struct {
	// Identity of the endorser (e.g. its certificate)
	Endorser []byte `protobuf:"bytes,1,opt,name=endorser,proto3" json:"endorser,omitempty"`
	// Signature of the payload included in ProposalResponse concatenated with
	// the endorser's certificate; ie, sign(ProposalResponse.payload + endorser)
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endorsement) Reset()         { *m = Endorsement{} }
func (m *Endorsement) String() string { return proto.CompactTextString(m) }
func (*Endorsement) ProtoMessage()    {}
func (*Endorsement) Descriptor() ([]byte, []int) {
	return fileDescriptor_proposal_response_22a755721b685f40, []int{3}
}
func (m *Endorsement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endorsement.Unmarshal(m, b)
}
func (m *Endorsement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endorsement.Marshal(b, m, deterministic)
}
func (dst *Endorsement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endorsement.Merge(dst, src)
}
func (m *Endorsement) XXX_Size() int {
	return xxx_messageInfo_Endorsement.Size(m)
}
func (m *Endorsement) XXX_DiscardUnknown() {
	xxx_messageInfo_Endorsement.DiscardUnknown(m)
}

var xxx_messageInfo_Endorsement proto.InternalMessageInfo

func (m *Endorsement) GetEndorser() []byte {
	if m != nil {
		return m.Endorser
	}
	return nil
}

func (m *Endorsement) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposalResponse)(nil), "sdk.protos.ProposalResponse")
	proto.RegisterType((*Response)(nil), "sdk.protos.Response")
	proto.RegisterType((*ProposalResponsePayload)(nil), "sdk.protos.ProposalResponsePayload")
	proto.RegisterType((*Endorsement)(nil), "sdk.protos.Endorsement")
}

func init() {
	proto.RegisterFile("peer/proposal_response.proto", fileDescriptor_proposal_response_22a755721b685f40)
}

var fileDescriptor_proposal_response_22a755721b685f40 = []byte{
	// 367 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xe9, 0xfe, 0xff, 0xcd, 0x2d, 0x9b, 0x30, 0x2a, 0x68, 0x19, 0x03, 0x47, 0x7d, 0x99,
	0x20, 0x29, 0x28, 0x82, 0xcf, 0x03, 0xd1, 0xc7, 0x11, 0xc4, 0x07, 0x11, 0x24, 0xdd, 0xee, 0xd2,
	0x62, 0xdb, 0x84, 0xdc, 0x54, 0xdc, 0x07, 0xf6, 0x7b, 0x48, 0xd3, 0xa6, 0xab, 0xe2, 0xd3, 0x38,
	0x77, 0x27, 0xbf, 0x7b, 0xcf, 0xed, 0x25, 0x73, 0x05, 0xa0, 0x23, 0xa5, 0xa5, 0x92, 0xc8, 0xb3,
	0x37, 0x0d, 0xa8, 0x64, 0x81, 0x40, 0x95, 0x96, 0x46, 0xfa, 0x03, 0xfb, 0x83, 0xb3, 0x73, 0x21,
	0xa5, 0xc8, 0x20, 0xb2, 0x32, 0x2e, 0x77, 0x91, 0x49, 0x73, 0x40, 0xc3, 0x73, 0x55, 0x1b, 0xc3,
	0x2f, 0x8f, 0x4c, 0xd7, 0x0d, 0x84, 0x35, 0x0c, 0x3f, 0x20, 0x47, 0x1f, 0xa0, 0x31, 0x95, 0x45,
	0xe0, 0x2d, 0xbc, 0x65, 0x9f, 0x39, 0xe9, 0xdf, 0x91, 0x51, 0x4b, 0x08, 0x7a, 0x0b, 0x6f, 0x39,
	0xbe, 0x9e, 0xd1, 0xba, 0x07, 0x75, 0x3d, 0xe8, 0x93, 0x73, 0xb0, 0x83, 0xd9, 0xbf, 0x22, 0x43,
	0x37, 0x63, 0xf0, 0xdf, 0x3e, 0x9c, 0xd6, 0x2f, 0x90, 0xba, 0xbe, 0xac, 0x75, 0x54, 0x13, 0x28,
	0xbe, 0xcf, 0x24, 0xdf, 0x06, 0xfd, 0x85, 0xb7, 0x9c, 0x30, 0x27, 0xfd, 0x5b, 0x32, 0x86, 0x62,
	0x2b, 0x35, 0x42, 0x0e, 0x85, 0x09, 0x06, 0x16, 0x75, 0xe2, 0x50, 0xf7, 0x87, 0xbf, 0x58, 0xd7,
	0x17, 0x3e, 0x93, 0x61, 0x1b, 0xef, 0x94, 0x0c, 0xd0, 0x70, 0x53, 0x62, 0x93, 0xae, 0x51, 0x55,
	0xd3, 0x1c, 0x10, 0xb9, 0x00, 0x1b, 0x6d, 0xc4, 0x9c, 0xec, 0x8e, 0xf3, 0xef, 0xc7, 0x38, 0xe1,
	0x2b, 0x39, 0xfb, 0xbd, 0xbe, 0x75, 0x33, 0xe9, 0x05, 0x39, 0x6e, 0x3f, 0x4f, 0xc2, 0x31, 0xb1,
	0xdd, 0x26, 0x6c, 0xe2, 0x8a, 0x8f, 0x1c, 0x13, 0x7f, 0x4e, 0x46, 0xf0, 0x69, 0xa0, 0xb0, 0xcb,
	0xee, 0x59, 0xc3, 0xa1, 0x10, 0x3e, 0x90, 0x71, 0x27, 0x91, 0x3f, 0x23, 0xc3, 0x26, 0x93, 0x6e,
	0x60, 0xad, 0xae, 0x40, 0x98, 0x8a, 0x82, 0x9b, 0x52, 0x83, 0x03, 0xb5, 0x85, 0x55, 0x42, 0x42,
	0xa9, 0x05, 0x4d, 0xf6, 0x0a, 0x74, 0x06, 0x5b, 0x01, 0x9a, 0xee, 0x78, 0xac, 0xd3, 0x8d, 0x5b,
	0x5c, 0x75, 0x4d, 0xab, 0x3f, 0xa2, 0x6c, 0xde, 0xb9, 0x80, 0x97, 0x4b, 0x91, 0x9a, 0xa4, 0x8c,
	0xe9, 0x46, 0xe6, 0x51, 0x87, 0x11, 0xd5, 0x8c, 0xfa, 0xba, 0x30, 0xaa, 0x18, 0x71, 0x7d, 0x79,
	0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x52, 0x0b, 0x35, 0xa0, 0x02, 0x00, 0x00,
}
