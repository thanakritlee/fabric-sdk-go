/*
Notice: This file has been modified for Hyperledger Fabric SDK Go usage.
Please review third_party pinning scripts and patches for more details.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/proposal.proto

package peer // import "github.com/thanakritlee/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import token "github.com/thanakritlee/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/token"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This structure is necessary to sign the proposal which contains the header
// and the payload. Without this structure, we would have to concatenate the
// header and the payload to verify the signature, which could be expensive
// with large payload
//
// When an endorser receives a SignedProposal message, it should verify the
// signature over the proposal bytes. This verification requires the following
// steps:
// 1. Verification of the validity of the certificate that was used to produce
//    the signature.  The certificate will be available once proposalBytes has
//    been unmarshalled to a Proposal message, and Proposal.header has been
//    unmarshalled to a Header message. While this unmarshalling-before-verifying
//    might not be ideal, it is unavoidable because i) the signature needs to also
//    protect the signing certificate; ii) it is desirable that Header is created
//    once by the client and never changed (for the sake of accountability and
//    non-repudiation). Note also that it is actually impossible to conclusively
//    verify the validity of the certificate included in a Proposal, because the
//    proposal needs to first be endorsed and ordered with respect to certificate
//    expiration transactions. Still, it is useful to pre-filter expired
//    certificates at this stage.
// 2. Verification that the certificate is trusted (signed by a trusted CA) and
//    that it is allowed to transact with us (with respect to some ACLs);
// 3. Verification that the signature on proposalBytes is valid;
// 4. Detect replay attacks;
type SignedProposal struct {
	// The bytes of Proposal
	ProposalBytes []byte `protobuf:"bytes,1,opt,name=proposal_bytes,json=proposalBytes,proto3" json:"proposal_bytes,omitempty"`
	// Signaure over proposalBytes; this signature is to be verified against
	// the creator identity contained in the header of the Proposal message
	// marshaled as proposalBytes
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedProposal) Reset()         { *m = SignedProposal{} }
func (m *SignedProposal) String() string { return proto.CompactTextString(m) }
func (*SignedProposal) ProtoMessage()    {}
func (*SignedProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_proposal_2be65988745cf952, []int{0}
}
func (m *SignedProposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedProposal.Unmarshal(m, b)
}
func (m *SignedProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedProposal.Marshal(b, m, deterministic)
}
func (dst *SignedProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedProposal.Merge(dst, src)
}
func (m *SignedProposal) XXX_Size() int {
	return xxx_messageInfo_SignedProposal.Size(m)
}
func (m *SignedProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedProposal.DiscardUnknown(m)
}

var xxx_messageInfo_SignedProposal proto.InternalMessageInfo

func (m *SignedProposal) GetProposalBytes() []byte {
	if m != nil {
		return m.ProposalBytes
	}
	return nil
}

func (m *SignedProposal) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// A Proposal is sent to an endorser for endorsement.  The proposal contains:
// 1. A header which should be unmarshaled to a Header message.  Note that
//    Header is both the header of a Proposal and of a Transaction, in that i)
//    both headers should be unmarshaled to this message; and ii) it is used to
//    compute cryptographic hashes and signatures.  The header has fields common
//    to all proposals/transactions.  In addition it has a type field for
//    additional customization. An example of this is the ChaincodeHeaderExtension
//    message used to extend the Header for type CHAINCODE.
// 2. A payload whose type depends on the header's type field.
// 3. An extension whose type depends on the header's type field.
//
// Let us see an example. For type CHAINCODE (see the Header message),
// we have the following:
// 1. The header is a Header message whose extensions field is a
//    ChaincodeHeaderExtension message.
// 2. The payload is a ChaincodeProposalPayload message.
// 3. The extension is a ChaincodeAction that might be used to ask the
//    endorsers to endorse a specific ChaincodeAction, thus emulating the
//    submitting peer model.
type Proposal struct {
	// The header of the proposal. It is the bytes of the Header
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The payload of the proposal as defined by the type in the proposal
	// header.
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// Optional extensions to the proposal. Its content depends on the Header's
	// type field.  For the type CHAINCODE, it might be the bytes of a
	// ChaincodeAction message.
	Extension            []byte   `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_proposal_2be65988745cf952, []int{1}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Proposal.Unmarshal(m, b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
}
func (dst *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(dst, src)
}
func (m *Proposal) XXX_Size() int {
	return xxx_messageInfo_Proposal.Size(m)
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Proposal) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Proposal) GetExtension() []byte {
	if m != nil {
		return m.Extension
	}
	return nil
}

// ChaincodeHeaderExtension is the Header's extentions message to be used when
// the Header's type is CHAINCODE.  This extensions is used to specify which
// chaincode to invoke and what should appear on the ledger.
type ChaincodeHeaderExtension struct {
	// The PayloadVisibility field controls to what extent the Proposal's payload
	// (recall that for the type CHAINCODE, it is ChaincodeProposalPayload
	// message) field will be visible in the final transaction and in the ledger.
	// Ideally, it would be configurable, supporting at least 3 main visibility
	// modes:
	// 1. all bytes of the payload are visible;
	// 2. only a hash of the payload is visible;
	// 3. nothing is visible.
	// Notice that the visibility function may be potentially part of the ESCC.
	// In that case it overrides PayloadVisibility field.  Finally notice that
	// this field impacts the content of ProposalResponsePayload.proposalHash.
	PayloadVisibility []byte `protobuf:"bytes,1,opt,name=payload_visibility,json=payloadVisibility,proto3" json:"payload_visibility,omitempty"`
	// The ID of the chaincode to target.
	ChaincodeId          *ChaincodeID `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ChaincodeHeaderExtension) Reset()         { *m = ChaincodeHeaderExtension{} }
func (m *ChaincodeHeaderExtension) String() string { return proto.CompactTextString(m) }
func (*ChaincodeHeaderExtension) ProtoMessage()    {}
func (*ChaincodeHeaderExtension) Descriptor() ([]byte, []int) {
	return fileDescriptor_proposal_2be65988745cf952, []int{2}
}
func (m *ChaincodeHeaderExtension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeHeaderExtension.Unmarshal(m, b)
}
func (m *ChaincodeHeaderExtension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeHeaderExtension.Marshal(b, m, deterministic)
}
func (dst *ChaincodeHeaderExtension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeHeaderExtension.Merge(dst, src)
}
func (m *ChaincodeHeaderExtension) XXX_Size() int {
	return xxx_messageInfo_ChaincodeHeaderExtension.Size(m)
}
func (m *ChaincodeHeaderExtension) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeHeaderExtension.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeHeaderExtension proto.InternalMessageInfo

func (m *ChaincodeHeaderExtension) GetPayloadVisibility() []byte {
	if m != nil {
		return m.PayloadVisibility
	}
	return nil
}

func (m *ChaincodeHeaderExtension) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

// ChaincodeProposalPayload is the Proposal's payload message to be used when
// the Header's type is CHAINCODE.  It contains the arguments for this
// invocation.
type ChaincodeProposalPayload struct {
	// Input contains the arguments for this invocation. If this invocation
	// deploys a new chaincode, ESCC/VSCC are part of this field.
	// This is usually a marshaled ChaincodeInvocationSpec
	Input []byte `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// TransientMap contains data (e.g. cryptographic material) that might be used
	// to implement some form of application-level confidentiality. The contents
	// of this field are supposed to always be omitted from the transaction and
	// excluded from the ledger.
	TransientMap         map[string][]byte `protobuf:"bytes,2,rep,name=TransientMap,proto3" json:"TransientMap,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ChaincodeProposalPayload) Reset()         { *m = ChaincodeProposalPayload{} }
func (m *ChaincodeProposalPayload) String() string { return proto.CompactTextString(m) }
func (*ChaincodeProposalPayload) ProtoMessage()    {}
func (*ChaincodeProposalPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_proposal_2be65988745cf952, []int{3}
}
func (m *ChaincodeProposalPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeProposalPayload.Unmarshal(m, b)
}
func (m *ChaincodeProposalPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeProposalPayload.Marshal(b, m, deterministic)
}
func (dst *ChaincodeProposalPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeProposalPayload.Merge(dst, src)
}
func (m *ChaincodeProposalPayload) XXX_Size() int {
	return xxx_messageInfo_ChaincodeProposalPayload.Size(m)
}
func (m *ChaincodeProposalPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeProposalPayload.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeProposalPayload proto.InternalMessageInfo

func (m *ChaincodeProposalPayload) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *ChaincodeProposalPayload) GetTransientMap() map[string][]byte {
	if m != nil {
		return m.TransientMap
	}
	return nil
}

// ChaincodeAction contains the actions the events generated by the execution
// of the chaincode.
type ChaincodeAction struct {
	// This field contains the read set and the write set produced by the
	// chaincode executing this invocation.
	Results []byte `protobuf:"bytes,1,opt,name=results,proto3" json:"results,omitempty"`
	// This field contains the events generated by the chaincode executing this
	// invocation.
	Events []byte `protobuf:"bytes,2,opt,name=events,proto3" json:"events,omitempty"`
	// This field contains the result of executing this invocation.
	Response *Response `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	// This field contains the ChaincodeID of executing this invocation. Endorser
	// will set it with the ChaincodeID called by endorser while simulating proposal.
	// Committer will validate the version matching with latest chaincode version.
	// Adding ChaincodeID to keep version opens up the possibility of multiple
	// ChaincodeAction per transaction.
	ChaincodeId *ChaincodeID `protobuf:"bytes,4,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	// This field contains the token expectation generated by the chaincode
	// executing this invocation
	TokenExpectation     *token.TokenExpectation `protobuf:"bytes,5,opt,name=token_expectation,json=tokenExpectation,proto3" json:"token_expectation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ChaincodeAction) Reset()         { *m = ChaincodeAction{} }
func (m *ChaincodeAction) String() string { return proto.CompactTextString(m) }
func (*ChaincodeAction) ProtoMessage()    {}
func (*ChaincodeAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_proposal_2be65988745cf952, []int{4}
}
func (m *ChaincodeAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChaincodeAction.Unmarshal(m, b)
}
func (m *ChaincodeAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChaincodeAction.Marshal(b, m, deterministic)
}
func (dst *ChaincodeAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChaincodeAction.Merge(dst, src)
}
func (m *ChaincodeAction) XXX_Size() int {
	return xxx_messageInfo_ChaincodeAction.Size(m)
}
func (m *ChaincodeAction) XXX_DiscardUnknown() {
	xxx_messageInfo_ChaincodeAction.DiscardUnknown(m)
}

var xxx_messageInfo_ChaincodeAction proto.InternalMessageInfo

func (m *ChaincodeAction) GetResults() []byte {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ChaincodeAction) GetEvents() []byte {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *ChaincodeAction) GetResponse() *Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *ChaincodeAction) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeAction) GetTokenExpectation() *token.TokenExpectation {
	if m != nil {
		return m.TokenExpectation
	}
	return nil
}

func init() {
	proto.RegisterType((*SignedProposal)(nil), "sdk.protos.SignedProposal")
	proto.RegisterType((*Proposal)(nil), "sdk.protos.Proposal")
	proto.RegisterType((*ChaincodeHeaderExtension)(nil), "sdk.protos.ChaincodeHeaderExtension")
	proto.RegisterType((*ChaincodeProposalPayload)(nil), "sdk.protos.ChaincodeProposalPayload")
	proto.RegisterMapType((map[string][]byte)(nil), "sdk.protos.ChaincodeProposalPayload.TransientMapEntry")
	proto.RegisterType((*ChaincodeAction)(nil), "sdk.protos.ChaincodeAction")
}

func init() { proto.RegisterFile("peer/proposal.proto", fileDescriptor_proposal_2be65988745cf952) }

var fileDescriptor_proposal_2be65988745cf952 = []byte{
	// 487 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0x57, 0x5a, 0x36, 0x36, 0xb7, 0x6c, 0xad, 0x37, 0xa1, 0xa8, 0xda, 0x61, 0x8a, 0x84, 0x34,
	0x24, 0x48, 0xa4, 0x22, 0x21, 0xc4, 0x05, 0x51, 0xa8, 0xc4, 0x0e, 0x48, 0x53, 0x18, 0x3b, 0xec,
	0x52, 0x9c, 0xe4, 0x91, 0x5a, 0x0d, 0xb6, 0x65, 0x3b, 0xd5, 0x72, 0xe4, 0xe3, 0xf1, 0x6d, 0xf8,
	0x08, 0xc8, 0xb1, 0x9d, 0x76, 0xed, 0x85, 0x53, 0xf2, 0xde, 0xef, 0xfd, 0x7e, 0xef, 0xaf, 0xd1,
	0x99, 0x00, 0x90, 0x89, 0x90, 0x5c, 0x70, 0x45, 0xaa, 0x58, 0x48, 0xae, 0x39, 0x3e, 0x6c, 0x3f,
	0x6a, 0x72, 0xde, 0x82, 0xf9, 0x92, 0x50, 0x96, 0xf3, 0x02, 0x2c, 0x3a, 0xb9, 0x78, 0x44, 0x59,
	0x48, 0x50, 0x82, 0x33, 0xe5, 0xd1, 0x50, 0xf3, 0x15, 0xb0, 0x04, 0x1e, 0x04, 0xe4, 0x9a, 0x68,
	0xca, 0x99, 0xb2, 0x48, 0xf4, 0x1d, 0x9d, 0x7c, 0xa3, 0x25, 0x83, 0xe2, 0xc6, 0x51, 0xf1, 0x0b,
	0x74, 0xd2, 0xc9, 0x64, 0x8d, 0x06, 0x15, 0x06, 0x97, 0xc1, 0xd5, 0x30, 0x7d, 0xe6, 0xbd, 0x33,
	0xe3, 0xc4, 0x17, 0xe8, 0x58, 0xd1, 0x92, 0x11, 0x5d, 0x4b, 0x08, 0x7b, 0x6d, 0xc4, 0xc6, 0x11,
	0xdd, 0xa3, 0xa3, 0x4e, 0xf0, 0x39, 0x3a, 0x5c, 0x02, 0x29, 0x40, 0x3a, 0x21, 0x67, 0xe1, 0x10,
	0x3d, 0x15, 0xa4, 0xa9, 0x38, 0x29, 0x1c, 0xdf, 0x9b, 0x46, 0x1b, 0x1e, 0x34, 0x30, 0x45, 0x39,
	0x0b, 0xfb, 0x56, 0xbb, 0x73, 0x44, 0xbf, 0x03, 0x14, 0x7e, 0xf2, 0xed, 0x7f, 0x69, 0xb5, 0xe6,
	0x1e, 0xc4, 0xaf, 0x11, 0x76, 0x2a, 0x8b, 0x35, 0x55, 0x34, 0xa3, 0x15, 0xd5, 0x8d, 0x4b, 0x3c,
	0x76, 0xc8, 0x5d, 0x07, 0xe0, 0xb7, 0x68, 0xd8, 0x4d, 0x72, 0x41, 0x6d, 0x21, 0x83, 0xe9, 0x99,
	0x1d, 0x8e, 0x8a, 0xbb, 0x34, 0xd7, 0x9f, 0xd3, 0x41, 0x17, 0x78, 0x5d, 0x44, 0x7f, 0xb6, 0x6b,
	0xf0, 0x9d, 0xde, 0xb8, 0xf2, 0xcf, 0xd1, 0x01, 0x65, 0xa2, 0xd6, 0x2e, 0xad, 0x35, 0xf0, 0x1d,
	0x1a, 0xde, 0x4a, 0xc2, 0x14, 0x05, 0xa6, 0xbf, 0x12, 0x11, 0xf6, 0x2e, 0xfb, 0x57, 0x83, 0xe9,
	0x74, 0x2f, 0xd5, 0x8e, 0x5a, 0xbc, 0x4d, 0x9a, 0x33, 0x2d, 0x9b, 0xf4, 0x91, 0xce, 0xe4, 0x03,
	0x1a, 0xef, 0x85, 0xe0, 0x11, 0xea, 0xaf, 0xc0, 0xf6, 0x7d, 0x9c, 0x9a, 0x5f, 0x53, 0xd4, 0x9a,
	0x54, 0xb5, 0xdf, 0x95, 0x35, 0xde, 0xf7, 0xde, 0x05, 0xd1, 0xdf, 0x00, 0x9d, 0x76, 0xd9, 0x3f,
	0xe6, 0xe6, 0x3a, 0xcc, 0x6e, 0x24, 0xa8, 0xba, 0xd2, 0x7e, 0xfb, 0xde, 0x34, 0xdb, 0x84, 0x35,
	0x30, 0xad, 0x9c, 0x90, 0xb3, 0xf0, 0x2b, 0x74, 0xe4, 0x8f, 0xae, 0x5d, 0xd9, 0x60, 0x3a, 0xf2,
	0xad, 0xa5, 0xce, 0x9f, 0x76, 0x11, 0x7b, 0x73, 0x7f, 0xf2, 0x7f, 0x73, 0xc7, 0x73, 0x34, 0x6e,
	0x4f, 0x79, 0xb1, 0x75, 0xca, 0xe1, 0x41, 0x4b, 0x0e, 0x3d, 0xf9, 0xd6, 0x04, 0xcc, 0x37, 0x78,
	0x3a, 0xd2, 0x3b, 0x9e, 0xd9, 0x0f, 0x14, 0x71, 0x59, 0xc6, 0xcb, 0x46, 0x80, 0xac, 0xa0, 0x28,
	0x41, 0xc6, 0x3f, 0x49, 0x26, 0x69, 0xee, 0x35, 0xcc, 0x6b, 0x9a, 0x9d, 0x6e, 0x56, 0x91, 0xaf,
	0x48, 0x09, 0xf7, 0x2f, 0x4b, 0xaa, 0x97, 0x75, 0x16, 0xe7, 0xfc, 0x57, 0xb2, 0xc5, 0x4d, 0x2c,
	0x37, 0xb1, 0xdc, 0xc4, 0x70, 0x33, 0xfb, 0x5a, 0xdf, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xaa,
	0xf2, 0x0a, 0xf6, 0xcb, 0x03, 0x00, 0x00,
}
