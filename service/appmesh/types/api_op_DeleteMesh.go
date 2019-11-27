// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteMeshInput struct {
	_ struct{} `type:"structure"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMeshInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMeshInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMeshInput"}

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

type DeleteMeshOutput struct {
	_ struct{} `type:"structure" payload:"Mesh"`

	// An object representing a service mesh returned by a describe operation.
	//
	// Mesh is a required field
	Mesh *MeshData `locationName:"mesh" type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteMeshOutput) String() string {
	return awsutil.Prettify(s)
}
