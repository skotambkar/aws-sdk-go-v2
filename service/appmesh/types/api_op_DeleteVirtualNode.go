// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteVirtualNodeInput struct {
	_ struct{} `type:"structure"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// VirtualNodeName is a required field
	VirtualNodeName *string `location:"uri" locationName:"virtualNodeName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVirtualNodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVirtualNodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteVirtualNodeInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if s.VirtualNodeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualNodeName"))
	}
	if s.VirtualNodeName != nil && len(*s.VirtualNodeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualNodeName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteVirtualNodeOutput struct {
	_ struct{} `type:"structure" payload:"VirtualNode"`

	// An object that represents a virtual node returned by a describe operation.
	//
	// VirtualNode is a required field
	VirtualNode *VirtualNodeData `locationName:"virtualNode" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteVirtualNodeOutput) String() string {
	return awsutil.Prettify(s)
}
