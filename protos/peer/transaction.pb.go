// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/transaction.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import common "github.com/inklabsfoundation/inkchain/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TxValidationCode int32

const (
	TxValidationCode_VALID                        TxValidationCode = 0
	TxValidationCode_NIL_ENVELOPE                 TxValidationCode = 1
	TxValidationCode_BAD_PAYLOAD                  TxValidationCode = 2
	TxValidationCode_BAD_COMMON_HEADER            TxValidationCode = 3
	TxValidationCode_BAD_CREATOR_SIGNATURE        TxValidationCode = 4
	TxValidationCode_INVALID_ENDORSER_TRANSACTION TxValidationCode = 5
	TxValidationCode_INVALID_CONFIG_TRANSACTION   TxValidationCode = 6
	TxValidationCode_UNSUPPORTED_TX_PAYLOAD       TxValidationCode = 7
	TxValidationCode_BAD_PROPOSAL_TXID            TxValidationCode = 8
	TxValidationCode_DUPLICATE_TXID               TxValidationCode = 9
	TxValidationCode_ENDORSEMENT_POLICY_FAILURE   TxValidationCode = 10
	TxValidationCode_MVCC_READ_CONFLICT           TxValidationCode = 11
	TxValidationCode_PHANTOM_READ_CONFLICT        TxValidationCode = 12
	TxValidationCode_UNKNOWN_TX_TYPE              TxValidationCode = 13
	TxValidationCode_TARGET_CHAIN_NOT_FOUND       TxValidationCode = 14
	TxValidationCode_MARSHAL_TX_ERROR             TxValidationCode = 15
	TxValidationCode_NIL_TXACTION                 TxValidationCode = 16
	TxValidationCode_EXPIRED_CHAINCODE            TxValidationCode = 17
	TxValidationCode_CHAINCODE_VERSION_CONFLICT   TxValidationCode = 18
	TxValidationCode_BAD_HEADER_EXTENSION         TxValidationCode = 19
	TxValidationCode_BAD_CHANNEL_HEADER           TxValidationCode = 20
	TxValidationCode_BAD_RESPONSE_PAYLOAD         TxValidationCode = 21
	TxValidationCode_BAD_RWSET                    TxValidationCode = 22
	TxValidationCode_ILLEGAL_WRITESET             TxValidationCode = 23
	TxValidationCode_TRANSFER_CONFLICT            TxValidationCode = 100
	TxValidationCode_BAD_COUNTER                  TxValidationCode = 101
	TxValidationCode_BAD_SIGNATURE                TxValidationCode = 102
	TxValidationCode_BAD_BALANCE                  TxValidationCode = 103
	TxValidationCode_EXCEED_BALANCE               TxValidationCode = 104
	TxValidationCode_INVALID_OTHER_REASON         TxValidationCode = 255
)

var TxValidationCode_name = map[int32]string{
	0:   "VALID",
	1:   "NIL_ENVELOPE",
	2:   "BAD_PAYLOAD",
	3:   "BAD_COMMON_HEADER",
	4:   "BAD_CREATOR_SIGNATURE",
	5:   "INVALID_ENDORSER_TRANSACTION",
	6:   "INVALID_CONFIG_TRANSACTION",
	7:   "UNSUPPORTED_TX_PAYLOAD",
	8:   "BAD_PROPOSAL_TXID",
	9:   "DUPLICATE_TXID",
	10:  "ENDORSEMENT_POLICY_FAILURE",
	11:  "MVCC_READ_CONFLICT",
	12:  "PHANTOM_READ_CONFLICT",
	13:  "UNKNOWN_TX_TYPE",
	14:  "TARGET_CHAIN_NOT_FOUND",
	15:  "MARSHAL_TX_ERROR",
	16:  "NIL_TXACTION",
	17:  "EXPIRED_CHAINCODE",
	18:  "CHAINCODE_VERSION_CONFLICT",
	19:  "BAD_HEADER_EXTENSION",
	20:  "BAD_CHANNEL_HEADER",
	21:  "BAD_RESPONSE_PAYLOAD",
	22:  "BAD_RWSET",
	23:  "ILLEGAL_WRITESET",
	100: "TRANSFER_CONFLICT",
	101: "BAD_COUNTER",
	102: "BAD_SIGNATURE",
	103: "BAD_BALANCE",
	104: "EXCEED_BALANCE",
	255: "INVALID_OTHER_REASON",
}
var TxValidationCode_value = map[string]int32{
	"VALID":                        0,
	"NIL_ENVELOPE":                 1,
	"BAD_PAYLOAD":                  2,
	"BAD_COMMON_HEADER":            3,
	"BAD_CREATOR_SIGNATURE":        4,
	"INVALID_ENDORSER_TRANSACTION": 5,
	"INVALID_CONFIG_TRANSACTION":   6,
	"UNSUPPORTED_TX_PAYLOAD":       7,
	"BAD_PROPOSAL_TXID":            8,
	"DUPLICATE_TXID":               9,
	"ENDORSEMENT_POLICY_FAILURE":   10,
	"MVCC_READ_CONFLICT":           11,
	"PHANTOM_READ_CONFLICT":        12,
	"UNKNOWN_TX_TYPE":              13,
	"TARGET_CHAIN_NOT_FOUND":       14,
	"MARSHAL_TX_ERROR":             15,
	"NIL_TXACTION":                 16,
	"EXPIRED_CHAINCODE":            17,
	"CHAINCODE_VERSION_CONFLICT":   18,
	"BAD_HEADER_EXTENSION":         19,
	"BAD_CHANNEL_HEADER":           20,
	"BAD_RESPONSE_PAYLOAD":         21,
	"BAD_RWSET":                    22,
	"ILLEGAL_WRITESET":             23,
	"TRANSFER_CONFLICT":            100,
	"BAD_COUNTER":                  101,
	"BAD_SIGNATURE":                102,
	"BAD_BALANCE":                  103,
	"EXCEED_BALANCE":               104,
	"INVALID_OTHER_REASON":         255,
}

func (x TxValidationCode) String() string {
	return proto.EnumName(TxValidationCode_name, int32(x))
}
func (TxValidationCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

// This message is necessary to facilitate the verification of the signature
// (in the signature field) over the bytes of the transaction (in the
// transactionBytes field).
type SignedTransaction struct {
	// The bytes of the Transaction. NDD
	TransactionBytes []byte `protobuf:"bytes,1,opt,name=transaction_bytes,json=transactionBytes,proto3" json:"transaction_bytes,omitempty"`
	// Signature of the transactionBytes The public key of the signature is in
	// the header field of TransactionAction There might be multiple
	// TransactionAction, so multiple headers, but there should be same
	// transactor identity (cert) in all headers
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedTransaction) Reset()                    { *m = SignedTransaction{} }
func (m *SignedTransaction) String() string            { return proto.CompactTextString(m) }
func (*SignedTransaction) ProtoMessage()               {}
func (*SignedTransaction) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{0} }

func (m *SignedTransaction) GetTransactionBytes() []byte {
	if m != nil {
		return m.TransactionBytes
	}
	return nil
}

func (m *SignedTransaction) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// ProcessedTransaction wraps an Envelope that includes a transaction along with an indication
// of whether the transaction was validated or invalidated by committing peer.
// The use case is that GetTransactionByID API needs to retrieve the transaction Envelope
// from block storage, and return it to a client, and indicate whether the transaction
// was validated or invalidated by committing peer. So that the originally submitted
// transaction Envelope is not modified, the ProcessedTransaction wrapper is returned.
type ProcessedTransaction struct {
	// An Envelope which includes a processed transaction
	TransactionEnvelope *common.Envelope `protobuf:"bytes,1,opt,name=transactionEnvelope" json:"transactionEnvelope,omitempty"`
	// An indication of whether the transaction was validated or invalidated by committing peer
	ValidationCode int32  `protobuf:"varint,2,opt,name=validationCode" json:"validationCode,omitempty"`
	BlockHash      []byte `protobuf:"bytes,3,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (m *ProcessedTransaction) Reset()                    { *m = ProcessedTransaction{} }
func (m *ProcessedTransaction) String() string            { return proto.CompactTextString(m) }
func (*ProcessedTransaction) ProtoMessage()               {}
func (*ProcessedTransaction) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{1} }

func (m *ProcessedTransaction) GetTransactionEnvelope() *common.Envelope {
	if m != nil {
		return m.TransactionEnvelope
	}
	return nil
}

func (m *ProcessedTransaction) GetValidationCode() int32 {
	if m != nil {
		return m.ValidationCode
	}
	return 0
}

func (m *ProcessedTransaction) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

// The transaction to be sent to the ordering service. A transaction contains
// one or more TransactionAction. Each TransactionAction binds a proposal to
// potentially multiple actions. The transaction is atomic meaning that either
// all actions in the transaction will be committed or none will.  Note that
// while a Transaction might include more than one Header, the Header.creator
// field must be the same in each.
// A single client is free to issue a number of independent Proposal, each with
// their header (Header) and request payload (ChaincodeProposalPayload).  Each
// proposal is independently endorsed generating an action
// (ProposalResponsePayload) with one signature per Endorser. Any number of
// independent proposals (and their action) might be included in a transaction
// to ensure that they are treated atomically.
type Transaction struct {
	// The payload is an array of TransactionAction. An array is necessary to
	// accommodate multiple actions per transaction
	Actions []*TransactionAction `protobuf:"bytes,1,rep,name=actions" json:"actions,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{2} }

func (m *Transaction) GetActions() []*TransactionAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

// TransactionAction binds a proposal to its action.  The type field in the
// header dictates the type of action to be applied to the ledger.
type TransactionAction struct {
	// The header of the proposal action, which is the proposal header
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The payload of the action as defined by the type in the header For
	// chaincode, it's the bytes of ChaincodeActionPayload
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *TransactionAction) Reset()                    { *m = TransactionAction{} }
func (m *TransactionAction) String() string            { return proto.CompactTextString(m) }
func (*TransactionAction) ProtoMessage()               {}
func (*TransactionAction) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *TransactionAction) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TransactionAction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// ChaincodeActionPayload is the message to be used for the TransactionAction's
// payload when the Header's type is set to CHAINCODE.  It carries the
// chaincodeProposalPayload and an endorsed action to apply to the ledger.
type ChaincodeActionPayload struct {
	// This field contains the bytes of the ChaincodeProposalPayload message from
	// the original invocation (essentially the arguments) after the application
	// of the visibility function. The main visibility modes are "full" (the
	// entire ChaincodeProposalPayload message is included here), "hash" (only
	// the hash of the ChaincodeProposalPayload message is included) or
	// "nothing".  This field will be used to check the consistency of
	// ProposalResponsePayload.proposalHash.  For the CHAINCODE type,
	// ProposalResponsePayload.proposalHash is supposed to be H(ProposalHeader ||
	// f(ChaincodeProposalPayload)) where f is the visibility function.
	ChaincodeProposalPayload []byte `protobuf:"bytes,1,opt,name=chaincode_proposal_payload,json=chaincodeProposalPayload,proto3" json:"chaincode_proposal_payload,omitempty"`
	// The list of actions to apply to the ledger
	Action *ChaincodeEndorsedAction `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
}

func (m *ChaincodeActionPayload) Reset()                    { *m = ChaincodeActionPayload{} }
func (m *ChaincodeActionPayload) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeActionPayload) ProtoMessage()               {}
func (*ChaincodeActionPayload) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{4} }

func (m *ChaincodeActionPayload) GetChaincodeProposalPayload() []byte {
	if m != nil {
		return m.ChaincodeProposalPayload
	}
	return nil
}

func (m *ChaincodeActionPayload) GetAction() *ChaincodeEndorsedAction {
	if m != nil {
		return m.Action
	}
	return nil
}

// ChaincodeEndorsedAction carries information about the endorsement of a
// specific proposal
type ChaincodeEndorsedAction struct {
	// This is the bytes of the ProposalResponsePayload message signed by the
	// endorsers.  Recall that for the CHAINCODE type, the
	// ProposalResponsePayload's extenstion field carries a ChaincodeAction
	ProposalResponsePayload []byte `protobuf:"bytes,1,opt,name=proposal_response_payload,json=proposalResponsePayload,proto3" json:"proposal_response_payload,omitempty"`
	// The endorsement of the proposal, basically the endorser's signature over
	// proposalResponsePayload
	Endorsements []*Endorsement `protobuf:"bytes,2,rep,name=endorsements" json:"endorsements,omitempty"`
}

func (m *ChaincodeEndorsedAction) Reset()                    { *m = ChaincodeEndorsedAction{} }
func (m *ChaincodeEndorsedAction) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeEndorsedAction) ProtoMessage()               {}
func (*ChaincodeEndorsedAction) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *ChaincodeEndorsedAction) GetProposalResponsePayload() []byte {
	if m != nil {
		return m.ProposalResponsePayload
	}
	return nil
}

func (m *ChaincodeEndorsedAction) GetEndorsements() []*Endorsement {
	if m != nil {
		return m.Endorsements
	}
	return nil
}

func init() {
	proto.RegisterType((*SignedTransaction)(nil), "protos.SignedTransaction")
	proto.RegisterType((*ProcessedTransaction)(nil), "protos.ProcessedTransaction")
	proto.RegisterType((*Transaction)(nil), "protos.Transaction")
	proto.RegisterType((*TransactionAction)(nil), "protos.TransactionAction")
	proto.RegisterType((*ChaincodeActionPayload)(nil), "protos.ChaincodeActionPayload")
	proto.RegisterType((*ChaincodeEndorsedAction)(nil), "protos.ChaincodeEndorsedAction")
	proto.RegisterEnum("protos.TxValidationCode", TxValidationCode_name, TxValidationCode_value)
}

func init() { proto.RegisterFile("peer/transaction.proto", fileDescriptor11) }

var fileDescriptor11 = []byte{
	// 897 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xdd, 0x6e, 0xea, 0x46,
	0x17, 0xfd, 0x48, 0xbe, 0x24, 0xcd, 0x90, 0x9f, 0x61, 0x42, 0x08, 0x41, 0x69, 0xcf, 0x11, 0x17,
	0xd5, 0xd1, 0xa9, 0x14, 0xd4, 0x9c, 0x8b, 0x4a, 0x55, 0x6f, 0x06, 0x7b, 0x27, 0x58, 0x35, 0x33,
	0xd6, 0x78, 0x20, 0xa4, 0x17, 0x1d, 0x19, 0x98, 0x00, 0x0a, 0xd8, 0xc8, 0x26, 0x47, 0x3d, 0x2f,
	0xd1, 0x3e, 0x43, 0x5f, 0xa7, 0x2f, 0xd5, 0x6a, 0xfc, 0x03, 0x24, 0x69, 0x6f, 0x30, 0xb3, 0xd6,
	0xf2, 0xde, 0x6b, 0xff, 0xd8, 0x46, 0xb5, 0xa5, 0xd6, 0x71, 0x6b, 0x15, 0x07, 0x61, 0x12, 0x8c,
	0x56, 0xb3, 0x28, 0xbc, 0x5e, 0xc6, 0xd1, 0x2a, 0x22, 0xfb, 0xe9, 0x25, 0x69, 0xbc, 0x9b, 0x44,
	0xd1, 0x64, 0xae, 0x5b, 0xe9, 0x71, 0xf8, 0xfc, 0xd8, 0x5a, 0xcd, 0x16, 0x3a, 0x59, 0x05, 0x8b,
	0x65, 0x26, 0x6c, 0x5c, 0xa5, 0x01, 0x96, 0x71, 0xb4, 0x8c, 0x92, 0x60, 0xae, 0x62, 0x9d, 0x2c,
	0xa3, 0x30, 0xd1, 0x39, 0x7b, 0x36, 0x8a, 0x16, 0x8b, 0x28, 0x6c, 0x65, 0x97, 0x0c, 0x6c, 0xfe,
	0x8a, 0x2a, 0xfe, 0x6c, 0x12, 0xea, 0xb1, 0xdc, 0xa4, 0x25, 0xdf, 0xa1, 0xca, 0x96, 0x0b, 0x35,
	0xfc, 0xb2, 0xd2, 0x49, 0xbd, 0xf4, 0xbe, 0xf4, 0xe1, 0x48, 0xe0, 0x2d, 0xa2, 0x6d, 0x70, 0x72,
	0x85, 0x0e, 0x93, 0xd9, 0x24, 0x0c, 0x56, 0xcf, 0xb1, 0xae, 0xef, 0xa4, 0xa2, 0x0d, 0xd0, 0xfc,
	0xb3, 0x84, 0xaa, 0x5e, 0x1c, 0x8d, 0x74, 0x92, 0xbc, 0xcc, 0xd1, 0x46, 0x67, 0x5b, 0xa1, 0x20,
	0xfc, 0xac, 0xe7, 0xd1, 0x52, 0xa7, 0x59, 0xca, 0x37, 0xf8, 0x3a, 0x37, 0x59, 0xe0, 0xe2, 0xdf,
	0xc4, 0xe4, 0x5b, 0x74, 0xf2, 0x39, 0x98, 0xcf, 0xc6, 0x81, 0x41, 0xad, 0x68, 0x9c, 0xe5, 0xdf,
	0x13, 0xaf, 0x50, 0xf2, 0x35, 0x42, 0xc3, 0x79, 0x34, 0x7a, 0x52, 0xd3, 0x20, 0x99, 0xd6, 0x77,
	0x33, 0x8f, 0x29, 0xd2, 0x09, 0x92, 0x69, 0xb3, 0x8d, 0xca, 0xdb, 0xce, 0x3e, 0xa1, 0x83, 0xec,
	0x9f, 0xa9, 0x79, 0xf7, 0x43, 0xf9, 0xe6, 0x32, 0xeb, 0x55, 0x72, 0xbd, 0xa5, 0xa2, 0xe9, 0xaf,
	0x28, 0x94, 0x4d, 0x40, 0x95, 0x37, 0x2c, 0xa9, 0xa1, 0xfd, 0xa9, 0x0e, 0xc6, 0x3a, 0xce, 0x9b,
	0x97, 0x9f, 0x48, 0x1d, 0x1d, 0x2c, 0x83, 0x2f, 0xf3, 0x28, 0x18, 0xe7, 0x0d, 0x2b, 0x8e, 0xcd,
	0x3f, 0x4a, 0xa8, 0x66, 0x4d, 0x83, 0x59, 0x38, 0x8a, 0xc6, 0x3a, 0x8b, 0xe2, 0x65, 0x14, 0xf9,
	0x09, 0x35, 0x46, 0x05, 0xa3, 0xd6, 0x33, 0x2e, 0xe2, 0x64, 0x09, 0xea, 0x6b, 0x85, 0x97, 0x0b,
	0x8a, 0xbb, 0x7f, 0x40, 0xfb, 0x99, 0xb5, 0x34, 0x63, 0xf9, 0xe6, 0x5d, 0x51, 0xd3, 0x3a, 0x1b,
	0x84, 0xe3, 0x28, 0x4e, 0xf4, 0x38, 0xaf, 0x2c, 0x97, 0x37, 0x7f, 0x2f, 0xa1, 0x8b, 0xff, 0xd0,
	0x90, 0x1f, 0xd1, 0xe5, 0x9b, 0x65, 0x7b, 0xe5, 0xe8, 0xa2, 0x10, 0x88, 0x9c, 0xdf, 0x18, 0x3a,
	0xd2, 0x59, 0xb4, 0x85, 0x0e, 0x57, 0x49, 0x7d, 0x27, 0x6d, 0xf5, 0x59, 0x61, 0x0b, 0x36, 0x9c,
	0x78, 0x21, 0xfc, 0xf8, 0xd7, 0x1e, 0xc2, 0xf2, 0xb7, 0xfe, 0xcb, 0x09, 0x1f, 0xa2, 0xbd, 0x3e,
	0x75, 0x1d, 0x1b, 0xff, 0x8f, 0x60, 0x74, 0xc4, 0x1c, 0x57, 0x01, 0xeb, 0x83, 0xcb, 0x3d, 0xc0,
	0x25, 0x72, 0x8a, 0xca, 0x6d, 0x6a, 0x2b, 0x8f, 0x3e, 0xb8, 0x9c, 0xda, 0x78, 0x87, 0x9c, 0xa3,
	0x8a, 0x01, 0x2c, 0xde, 0xed, 0x72, 0xa6, 0x3a, 0x40, 0x6d, 0x10, 0x78, 0x97, 0x5c, 0xa2, 0xf3,
	0x14, 0x16, 0x40, 0x25, 0x17, 0xca, 0x77, 0xee, 0x18, 0x95, 0x3d, 0x01, 0xf8, 0xff, 0xe4, 0x3d,
	0xba, 0x72, 0x58, 0x9a, 0x41, 0x01, 0xb3, 0xb9, 0xf0, 0x41, 0x28, 0x29, 0x28, 0xf3, 0xa9, 0x25,
	0x1d, 0xce, 0xf0, 0x1e, 0xf9, 0x06, 0x35, 0x0a, 0x85, 0xc5, 0xd9, 0xad, 0x73, 0xf7, 0x82, 0xdf,
	0x27, 0x0d, 0x54, 0xeb, 0x31, 0xbf, 0xe7, 0x79, 0x5c, 0x48, 0xb0, 0x95, 0x1c, 0xac, 0xfd, 0x1c,
	0x14, 0x7e, 0x3c, 0xc1, 0x3d, 0xee, 0x53, 0x57, 0xc9, 0x81, 0x63, 0xe3, 0xaf, 0x08, 0x41, 0x27,
	0x76, 0xcf, 0x73, 0x1d, 0x8b, 0x4a, 0xc8, 0xb0, 0x43, 0x93, 0x26, 0x37, 0xd0, 0x05, 0x26, 0x95,
	0xc7, 0x5d, 0xc7, 0x7a, 0x50, 0xb7, 0xd4, 0x71, 0x8d, 0x51, 0x44, 0x6a, 0x88, 0x74, 0xfb, 0x96,
	0xa5, 0x04, 0xd0, 0xcc, 0x88, 0xeb, 0x58, 0x12, 0x97, 0x4d, 0x6d, 0x5e, 0x87, 0x32, 0xc9, 0xbb,
	0xaf, 0xa8, 0x23, 0x72, 0x86, 0x4e, 0x7b, 0xec, 0x67, 0xc6, 0xef, 0x99, 0x71, 0x25, 0x1f, 0x3c,
	0xc0, 0xc7, 0xc6, 0xae, 0xa4, 0xe2, 0x0e, 0xa4, 0xb2, 0x3a, 0xd4, 0x61, 0x8a, 0x71, 0xa9, 0x6e,
	0x79, 0x8f, 0xd9, 0xf8, 0x84, 0x54, 0x11, 0xee, 0x52, 0xe1, 0x77, 0x52, 0xa7, 0x0a, 0x84, 0xe0,
	0x02, 0x9f, 0x16, 0x7d, 0x97, 0x83, 0xbc, 0x64, 0x6c, 0xca, 0x82, 0x81, 0xe7, 0x08, 0xb0, 0xb3,
	0x20, 0x16, 0xb7, 0x01, 0x57, 0x4c, 0x09, 0xeb, 0xa3, 0xea, 0x83, 0xf0, 0x1d, 0xce, 0x36, 0x7e,
	0x08, 0xa9, 0xa3, 0xaa, 0xe9, 0x46, 0x36, 0x16, 0x05, 0x03, 0x09, 0xcc, 0x48, 0xf0, 0x99, 0x29,
	0x2e, 0x1d, 0x50, 0x87, 0x32, 0x06, 0x6e, 0x31, 0xb8, 0x6a, 0x71, 0x87, 0x00, 0xdf, 0xe3, 0xcc,
	0x87, 0x75, 0x67, 0xcf, 0xc9, 0x31, 0x3a, 0x4c, 0x99, 0x7b, 0x1f, 0x24, 0xae, 0x19, 0xe7, 0x8e,
	0xeb, 0xc2, 0x1d, 0x75, 0xd5, 0xbd, 0x70, 0x24, 0x18, 0xf4, 0xc2, 0xf8, 0x4c, 0x67, 0x75, 0x0b,
	0x62, 0xe3, 0x63, 0x5c, 0xac, 0x8d, 0xc5, 0x7b, 0x4c, 0x82, 0xc0, 0x9a, 0x54, 0xd0, 0xb1, 0x01,
	0x36, 0x7b, 0xf1, 0x58, 0x68, 0xda, 0xd4, 0xa5, 0xcc, 0x02, 0x3c, 0x31, 0x33, 0x83, 0x81, 0x05,
	0xb0, 0xc1, 0xa6, 0xe4, 0x12, 0x55, 0x8b, 0xd5, 0xe0, 0xb2, 0x03, 0xc2, 0x4c, 0xc0, 0xe7, 0x0c,
	0xff, 0x5d, 0x6a, 0x27, 0xe8, 0x63, 0x14, 0x4f, 0xae, 0x67, 0xe1, 0xd3, 0x3c, 0x18, 0x26, 0x8f,
	0xd1, 0x73, 0x98, 0x2d, 0xb5, 0x41, 0xd2, 0x67, 0xb9, 0x78, 0x1e, 0xcc, 0x9b, 0xbd, 0x4d, 0xb6,
	0x5e, 0x31, 0x5e, 0x30, 0x7a, 0x0a, 0x26, 0xfa, 0x97, 0xef, 0x27, 0xb3, 0xd5, 0xf4, 0x79, 0x68,
	0x5e, 0x98, 0xad, 0x37, 0x61, 0x5a, 0x45, 0x98, 0xec, 0x9b, 0x91, 0xb4, 0x4c, 0x98, 0x61, 0xf6,
	0x3d, 0xf9, 0xf4, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0x56, 0x08, 0x9c, 0x70, 0x06, 0x00,
	0x00,
}
