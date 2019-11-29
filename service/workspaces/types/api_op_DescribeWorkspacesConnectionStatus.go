// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeWorkspacesConnectionStatusInput struct {
	_ struct{} `type:"structure"`

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string `min:"1" type:"string"`

	// The identifiers of the WorkSpaces. You can specify up to 25 WorkSpaces.
	WorkspaceIds []string `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeWorkspacesConnectionStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWorkspacesConnectionStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWorkspacesConnectionStatusInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.WorkspaceIds != nil && len(s.WorkspaceIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkspaceIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeWorkspacesConnectionStatusOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next set of results, or null if no more
	// results are available.
	NextToken *string `min:"1" type:"string"`

	// Information about the connection status of the WorkSpace.
	WorkspacesConnectionStatus []WorkspaceConnectionStatus `type:"list"`
}

// String returns the string representation
func (s DescribeWorkspacesConnectionStatusOutput) String() string {
	return awsutil.Prettify(s)
}
