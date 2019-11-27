// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to update an XssMatchSet.
type UpdateXssMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// An array of XssMatchSetUpdate objects that you want to insert into or delete
	// from an XssMatchSet. For more information, see the applicable data types:
	//
	//    * XssMatchSetUpdate: Contains Action and XssMatchTuple
	//
	//    * XssMatchTuple: Contains FieldToMatch and TextTransformation
	//
	//    * FieldToMatch: Contains Data and Type
	//
	// Updates is a required field
	Updates []XssMatchSetUpdate `min:"1" type:"list" required:"true"`

	// The XssMatchSetId of the XssMatchSet that you want to update. XssMatchSetId
	// is returned by CreateXssMatchSet and by ListXssMatchSets.
	//
	// XssMatchSetId is a required field
	XssMatchSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateXssMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateXssMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateXssMatchSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.Updates == nil {
		invalidParams.Add(aws.NewErrParamRequired("Updates"))
	}
	if s.Updates != nil && len(s.Updates) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Updates", 1))
	}

	if s.XssMatchSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("XssMatchSetId"))
	}
	if s.XssMatchSetId != nil && len(*s.XssMatchSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("XssMatchSetId", 1))
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

// The response to an UpdateXssMatchSets request.
type UpdateXssMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the UpdateXssMatchSet request. You
	// can also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateXssMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}
