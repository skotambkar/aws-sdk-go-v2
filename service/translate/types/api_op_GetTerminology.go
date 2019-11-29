// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/translate/enums"
)

type GetTerminologyInput struct {
	_ struct{} `type:"structure"`

	// The name of the custom terminology being retrieved.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The data format of the custom terminology being retrieved, either CSV or
	// TMX.
	//
	// TerminologyDataFormat is a required field
	TerminologyDataFormat enums.TerminologyDataFormat `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s GetTerminologyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTerminologyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTerminologyInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.TerminologyDataFormat) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("TerminologyDataFormat"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTerminologyOutput struct {
	_ struct{} `type:"structure"`

	// The data location of the custom terminology being retrieved. The custom terminology
	// file is returned in a presigned url that has a 30 minute expiration.
	TerminologyDataLocation *TerminologyDataLocation `type:"structure"`

	// The properties of the custom terminology being retrieved.
	TerminologyProperties *TerminologyProperties `type:"structure"`
}

// String returns the string representation
func (s GetTerminologyOutput) String() string {
	return awsutil.Prettify(s)
}