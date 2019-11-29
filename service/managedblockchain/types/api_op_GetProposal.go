// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetProposalInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the network for which the proposal is made.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`

	// The unique identifier of the proposal.
	//
	// ProposalId is a required field
	ProposalId *string `location:"uri" locationName:"proposalId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetProposalInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetProposalInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetProposalInput"}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if s.ProposalId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProposalId"))
	}
	if s.ProposalId != nil && len(*s.ProposalId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProposalId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetProposalOutput struct {
	_ struct{} `type:"structure"`

	// Information about a proposal.
	Proposal *Proposal `type:"structure"`
}

// String returns the string representation
func (s GetProposalOutput) String() string {
	return awsutil.Prettify(s)
}