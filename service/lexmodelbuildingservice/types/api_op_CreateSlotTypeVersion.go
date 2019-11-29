// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/enums"
)

type CreateSlotTypeVersionInput struct {
	_ struct{} `type:"structure"`

	// Checksum for the $LATEST version of the slot type that you want to publish.
	// If you specify a checksum and the $LATEST version of the slot type has a
	// different checksum, Amazon Lex returns a PreconditionFailedException exception
	// and doesn't publish the new version. If you don't specify a checksum, Amazon
	// Lex publishes the $LATEST version.
	Checksum *string `locationName:"checksum" type:"string"`

	// The name of the slot type that you want to create a new version for. The
	// name is case sensitive.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateSlotTypeVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSlotTypeVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSlotTypeVersionInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSlotTypeVersionOutput struct {
	_ struct{} `type:"structure"`

	// Checksum of the $LATEST version of the slot type.
	Checksum *string `locationName:"checksum" type:"string"`

	// The date that the slot type was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp"`

	// A description of the slot type.
	Description *string `locationName:"description" type:"string"`

	// A list of EnumerationValue objects that defines the values that the slot
	// type can take.
	EnumerationValues []EnumerationValue `locationName:"enumerationValues" min:"1" type:"list"`

	// The date that the slot type was updated. When you create a resource, the
	// creation date and last update date are the same.
	LastUpdatedDate *time.Time `locationName:"lastUpdatedDate" type:"timestamp"`

	// The name of the slot type.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The strategy that Amazon Lex uses to determine the value of the slot. For
	// more information, see PutSlotType.
	ValueSelectionStrategy enums.SlotValueSelectionStrategy `locationName:"valueSelectionStrategy" type:"string" enum:"true"`

	// The version assigned to the new slot type version.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s CreateSlotTypeVersionOutput) String() string {
	return awsutil.Prettify(s)
}