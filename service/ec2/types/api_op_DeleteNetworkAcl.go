// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteNetworkAclInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the network ACL.
	//
	// NetworkAclId is a required field
	NetworkAclId *string `locationName:"networkAclId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNetworkAclInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNetworkAclInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteNetworkAclInput"}

	if s.NetworkAclId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkAclId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteNetworkAclOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteNetworkAclOutput) String() string {
	return awsutil.Prettify(s)
}
