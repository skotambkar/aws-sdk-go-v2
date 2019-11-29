// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateSizeConstraintSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// The SizeConstraintSetId of the SizeConstraintSet that you want to update.
	// SizeConstraintSetId is returned by CreateSizeConstraintSet and by ListSizeConstraintSets.
	//
	// SizeConstraintSetId is a required field
	SizeConstraintSetId *string `min:"1" type:"string" required:"true"`

	// An array of SizeConstraintSetUpdate objects that you want to insert into
	// or delete from a SizeConstraintSet. For more information, see the applicable
	// data types:
	//
	//    * SizeConstraintSetUpdate: Contains Action and SizeConstraint
	//
	//    * SizeConstraint: Contains FieldToMatch, TextTransformation, ComparisonOperator,
	//    and Size
	//
	//    * FieldToMatch: Contains Data and Type
	//
	// Updates is a required field
	Updates []SizeConstraintSetUpdate `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateSizeConstraintSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSizeConstraintSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSizeConstraintSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.SizeConstraintSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SizeConstraintSetId"))
	}
	if s.SizeConstraintSetId != nil && len(*s.SizeConstraintSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SizeConstraintSetId", 1))
	}

	if s.Updates == nil {
		invalidParams.Add(aws.NewErrParamRequired("Updates"))
	}
	if s.Updates != nil && len(s.Updates) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Updates", 1))
	}
	if s.Updates != nil {
		for i, v := range s.Updates {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Updates", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateSizeConstraintSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the UpdateSizeConstraintSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateSizeConstraintSetOutput) String() string {
	return awsutil.Prettify(s)
}
