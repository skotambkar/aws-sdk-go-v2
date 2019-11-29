// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/enums"
)

type ListObjectParentsInput struct {
	_ struct{} `type:"structure"`

	// Represents the manner and timing in which the successful write or update
	// of an object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel enums.ConsistencyLevel `location:"header" locationName:"x-amz-consistency-level" type:"string" enum:"true"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// the object resides. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// When set to True, returns all ListObjectParentsResponse$ParentLinks. There
	// could be multiple links between a parent-child pair.
	IncludeAllLinksToEachParent *bool `type:"boolean"`

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token.
	NextToken *string `type:"string"`

	// The reference that identifies the object for which parent objects are being
	// listed.
	//
	// ObjectReference is a required field
	ObjectReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s ListObjectParentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListObjectParentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListObjectParentsInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ObjectReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ObjectReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListObjectParentsOutput struct {
	_ struct{} `type:"structure"`

	// The pagination token.
	NextToken *string `type:"string"`

	// Returns a list of parent reference and LinkName Tuples.
	ParentLinks []ObjectIdentifierAndLinkNameTuple `type:"list"`

	// The parent structure, which is a map with key as the ObjectIdentifier and
	// LinkName as the value.
	Parents map[string]string `type:"map"`
}

// String returns the string representation
func (s ListObjectParentsOutput) String() string {
	return awsutil.Prettify(s)
}