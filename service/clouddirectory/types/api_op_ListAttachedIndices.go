// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/enums"
)

type ListAttachedIndicesInput struct {
	_ struct{} `type:"structure"`

	// The consistency level to use for this operation.
	ConsistencyLevel enums.ConsistencyLevel `location:"header" locationName:"x-amz-consistency-level" type:"string" enum:"true"`

	// The ARN of the directory.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// The maximum number of results to retrieve.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// A reference to the object that has indices attached.
	//
	// TargetReference is a required field
	TargetReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s ListAttachedIndicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAttachedIndicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAttachedIndicesInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.TargetReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListAttachedIndicesOutput struct {
	_ struct{} `type:"structure"`

	// The indices attached to the specified object.
	IndexAttachments []IndexAttachment `type:"list"`

	// The pagination token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAttachedIndicesOutput) String() string {
	return awsutil.Prettify(s)
}