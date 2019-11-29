// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/enums"
)

type ListMembersInput struct {
	_ struct{} `type:"structure"`

	// An optional Boolean value. If provided, the request is limited either to
	// members that the current AWS account owns (true) or that other AWS accounts
	// own (false). If omitted, all members are listed.
	IsOwned *bool `location:"querystring" locationName:"isOwned" type:"boolean"`

	// The maximum number of members to return in the request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The optional name of the member to list.
	Name *string `location:"querystring" locationName:"name" type:"string"`

	// The unique identifier of the network for which to list members.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// An optional status specifier. If provided, only members currently in this
	// status are listed.
	Status enums.MemberStatus `location:"querystring" locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s ListMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListMembersInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListMembersOutput struct {
	_ struct{} `type:"structure"`

	// An array of MemberSummary objects. Each object contains details about a network
	// member.
	Members []MemberSummary `type:"list"`

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListMembersOutput) String() string {
	return awsutil.Prettify(s)
}
