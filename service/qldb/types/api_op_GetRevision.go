// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRevisionInput struct {
	_ struct{} `type:"structure"`

	// The block location of the document revision to be verified. An address is
	// an Amazon Ion structure that has two fields: strandId and sequenceNo.
	//
	// For example: {strandId:"BlFTjlSXze9BIh1KOszcE3",sequenceNo:14}
	//
	// BlockAddress is a required field
	BlockAddress *ValueHolder `type:"structure" required:"true" sensitive:"true"`

	// The latest block location covered by the digest for which to request a proof.
	// An address is an Amazon Ion structure that has two fields: strandId and sequenceNo.
	//
	// For example: {strandId:"BlFTjlSXze9BIh1KOszcE3",sequenceNo:49}
	DigestTipAddress *ValueHolder `type:"structure" sensitive:"true"`

	// The unique ID of the document to be verified.
	//
	// DocumentId is a required field
	DocumentId *string `min:"22" type:"string" required:"true"`

	// The name of the ledger.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRevisionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRevisionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRevisionInput"}

	if s.BlockAddress == nil {
		invalidParams.Add(aws.NewErrParamRequired("BlockAddress"))
	}

	if s.DocumentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DocumentId"))
	}
	if s.DocumentId != nil && len(*s.DocumentId) < 22 {
		invalidParams.Add(aws.NewErrParamMinLen("DocumentId", 22))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.BlockAddress != nil {
		if err := s.BlockAddress.Validate(); err != nil {
			invalidParams.AddNested("BlockAddress", err.(aws.ErrInvalidParams))
		}
	}
	if s.DigestTipAddress != nil {
		if err := s.DigestTipAddress.Validate(); err != nil {
			invalidParams.AddNested("DigestTipAddress", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRevisionOutput struct {
	_ struct{} `type:"structure"`

	// The proof object in Amazon Ion format returned by a GetRevision request.
	// A proof contains the list of hash values that are required to recalculate
	// the specified digest using a Merkle tree, starting with the specified document
	// revision.
	Proof *ValueHolder `type:"structure" sensitive:"true"`

	// The document revision data object in Amazon Ion format.
	//
	// Revision is a required field
	Revision *ValueHolder `type:"structure" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s GetRevisionOutput) String() string {
	return awsutil.Prettify(s)
}
