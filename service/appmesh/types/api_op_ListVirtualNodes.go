// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListVirtualNodesInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListVirtualNodesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListVirtualNodesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListVirtualNodesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListVirtualNodesOutput struct {
	_ struct{} `type:"structure"`

	NextToken *string `locationName:"nextToken" type:"string"`

	// VirtualNodes is a required field
	VirtualNodes []VirtualNodeRef `locationName:"virtualNodes" type:"list" required:"true"`
}

// String returns the string representation
func (s ListVirtualNodesOutput) String() string {
	return awsutil.Prettify(s)
}