// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAttributesInput struct {
	_ struct{} `type:"structure"`

	// The names of the attributes.
	AttributeNames []string `locationNameList:"AttributeName" type:"list" flattened:"true"`

	// Determines whether or not strong consistency should be enforced when data
	// is read from SimpleDB. If
	//    true
	// , any data previously written to SimpleDB will be returned. Otherwise, results
	// will be consistent eventually, and the client may not see data that was written
	// immediately before your read.
	ConsistentRead *bool `type:"boolean"`

	// The name of the domain in which to perform the operation.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// The name of the item.
	//
	// ItemName is a required field
	ItemName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAttributesInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.ItemName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ItemName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAttributesOutput struct {
	_ struct{} `type:"structure"`

	// The list of attributes returned by the operation.
	Attributes []Attribute `locationNameList:"Attribute" type:"list" flattened:"true"`
}

// String returns the string representation
func (s GetAttributesOutput) String() string {
	return awsutil.Prettify(s)
}