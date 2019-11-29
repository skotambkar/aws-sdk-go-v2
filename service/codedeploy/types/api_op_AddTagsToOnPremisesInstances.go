// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of, and adds tags to, an on-premises instance operation.
type AddTagsToOnPremisesInstancesInput struct {
	_ struct{} `type:"structure"`

	// The names of the on-premises instances to which to add tags.
	//
	// InstanceNames is a required field
	InstanceNames []string `locationName:"instanceNames" type:"list" required:"true"`

	// The tag key-value pairs to add to the on-premises instances.
	//
	// Keys and values are both required. Keys cannot be null or empty strings.
	// Value-only tags are not allowed.
	//
	// Tags is a required field
	Tags []Tag `locationName:"tags" type:"list" required:"true"`
}

// String returns the string representation
func (s AddTagsToOnPremisesInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddTagsToOnPremisesInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddTagsToOnPremisesInstancesInput"}

	if s.InstanceNames == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceNames"))
	}

	if s.Tags == nil {
		invalidParams.Add(aws.NewErrParamRequired("Tags"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddTagsToOnPremisesInstancesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddTagsToOnPremisesInstancesOutput) String() string {
	return awsutil.Prettify(s)
}
