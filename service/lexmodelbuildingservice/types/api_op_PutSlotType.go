// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/enums"
)

type PutSlotTypeInput struct {
	_ struct{} `type:"structure"`

	// Identifies a specific revision of the $LATEST version.
	//
	// When you create a new slot type, leave the checksum field blank. If you specify
	// a checksum you get a BadRequestException exception.
	//
	// When you want to update a slot type, set the checksum field to the checksum
	// of the most recent revision of the $LATEST version. If you don't specify
	// the checksum field, or if the checksum does not match the $LATEST version,
	// you get a PreconditionFailedException exception.
	Checksum *string `locationName:"checksum" type:"string"`

	CreateVersion *bool `locationName:"createVersion" type:"boolean"`

	// A description of the slot type.
	Description *string `locationName:"description" type:"string"`

	// A list of EnumerationValue objects that defines the values that the slot
	// type can take. Each value can have a list of synonyms, which are additional
	// values that help train the machine learning model about the values that it
	// resolves for a slot.
	//
	// When Amazon Lex resolves a slot value, it generates a resolution list that
	// contains up to five possible values for the slot. If you are using a Lambda
	// function, this resolution list is passed to the function. If you are not
	// using a Lambda function you can choose to return the value that the user
	// entered or the first value in the resolution list as the slot value. The
	// valueSelectionStrategy field indicates the option to use.
	EnumerationValues []EnumerationValue `locationName:"enumerationValues" min:"1" type:"list"`

	// The name of the slot type. The name is not case sensitive.
	//
	// The name can't match a built-in slot type name, or a built-in slot type name
	// with "AMAZON." removed. For example, because there is a built-in slot type
	// called AMAZON.DATE, you can't create a custom slot type called DATE.
	//
	// For a list of built-in slot types, see Slot Type Reference (https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/slot-type-reference)
	// in the Alexa Skills Kit.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"1" type:"string" required:"true"`

	// Determines the slot resolution strategy that Amazon Lex uses to return slot
	// type values. The field can be set to one of the following values:
	//
	//    * ORIGINAL_VALUE - Returns the value entered by the user, if the user
	//    value is similar to the slot value.
	//
	//    * TOP_RESOLUTION - If there is a resolution list for the slot, return
	//    the first value in the resolution list as the slot type value. If there
	//    is no resolution list, null is returned.
	//
	// If you don't specify the valueSelectionStrategy, the default is ORIGINAL_VALUE.
	ValueSelectionStrategy enums.SlotValueSelectionStrategy `locationName:"valueSelectionStrategy" type:"string" enum:"true"`
}

// String returns the string representation
func (s PutSlotTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutSlotTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutSlotTypeInput"}
	if s.EnumerationValues != nil && len(s.EnumerationValues) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EnumerationValues", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.EnumerationValues != nil {
		for i, v := range s.EnumerationValues {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "EnumerationValues", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutSlotTypeOutput struct {
	_ struct{} `type:"structure"`

	// Checksum of the $LATEST version of the slot type.
	Checksum *string `locationName:"checksum" type:"string"`

	CreateVersion *bool `locationName:"createVersion" type:"boolean"`

	// The date that the slot type was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp"`

	// A description of the slot type.
	Description *string `locationName:"description" type:"string"`

	// A list of EnumerationValue objects that defines the values that the slot
	// type can take.
	EnumerationValues []EnumerationValue `locationName:"enumerationValues" min:"1" type:"list"`

	// The date that the slot type was updated. When you create a slot type, the
	// creation date and last update date are the same.
	LastUpdatedDate *time.Time `locationName:"lastUpdatedDate" type:"timestamp"`

	// The name of the slot type.
	Name *string `locationName:"name" min:"1" type:"string"`

	// The slot resolution strategy that Amazon Lex uses to determine the value
	// of the slot. For more information, see PutSlotType.
	ValueSelectionStrategy enums.SlotValueSelectionStrategy `locationName:"valueSelectionStrategy" type:"string" enum:"true"`

	// The version of the slot type. For a new slot type, the version is always
	// $LATEST.
	Version *string `locationName:"version" min:"1" type:"string"`
}

// String returns the string representation
func (s PutSlotTypeOutput) String() string {
	return awsutil.Prettify(s)
}