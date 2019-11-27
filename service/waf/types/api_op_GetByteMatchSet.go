// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetByteMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The ByteMatchSetId of the ByteMatchSet that you want to get. ByteMatchSetId
	// is returned by CreateByteMatchSet and by ListByteMatchSets.
	//
	// ByteMatchSetId is a required field
	ByteMatchSetId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetByteMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetByteMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetByteMatchSetInput"}

	if s.ByteMatchSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ByteMatchSetId"))
	}
	if s.ByteMatchSetId != nil && len(*s.ByteMatchSetId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ByteMatchSetId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetByteMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// Information about the ByteMatchSet that you specified in the GetByteMatchSet
	// request. For more information, see the following topics:
	//
	//    * ByteMatchSet: Contains ByteMatchSetId, ByteMatchTuples, and Name
	//
	//    * ByteMatchTuples: Contains an array of ByteMatchTuple objects. Each ByteMatchTuple
	//    object contains FieldToMatch, PositionalConstraint, TargetString, and
	//    TextTransformation
	//
	//    * FieldToMatch: Contains Data and Type
	ByteMatchSet *ByteMatchSet `type:"structure"`
}

// String returns the string representation
func (s GetByteMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}
