// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateVirtualServiceInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `locationName:"clientToken" type:"string" idempotencyToken:"true"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	// An object that represents the specification of a virtual service.
	//
	// Spec is a required field
	Spec *VirtualServiceSpec `locationName:"spec" type:"structure" required:"true"`

	Tags []TagRef `locationName:"tags" type:"list"`

	// VirtualServiceName is a required field
	VirtualServiceName *string `locationName:"virtualServiceName" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVirtualServiceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVirtualServiceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVirtualServiceInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}

	if s.Spec == nil {
		invalidParams.Add(aws.NewErrParamRequired("Spec"))
	}

	if s.VirtualServiceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualServiceName"))
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			invalidParams.AddNested("Spec", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateVirtualServiceOutput struct {
	_ struct{} `type:"structure" payload:"VirtualService"`

	// An object that represents a virtual service returned by a describe operation.
	//
	// VirtualService is a required field
	VirtualService *VirtualServiceData `locationName:"virtualService" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateVirtualServiceOutput) String() string {
	return awsutil.Prettify(s)
}