// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DetachObjectInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) that is associated with the Directory where
	// objects reside. For more information, see arns.
	//
	// DirectoryArn is a required field
	DirectoryArn *string `location:"header" locationName:"x-amz-data-partition" type:"string" required:"true"`

	// The link name associated with the object that needs to be detached.
	//
	// LinkName is a required field
	LinkName *string `min:"1" type:"string" required:"true"`

	// The parent reference from which the object with the specified link name is
	// detached.
	//
	// ParentReference is a required field
	ParentReference *ObjectReference `type:"structure" required:"true"`
}

// String returns the string representation
func (s DetachObjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachObjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachObjectInput"}

	if s.DirectoryArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryArn"))
	}

	if s.LinkName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LinkName"))
	}
	if s.LinkName != nil && len(*s.LinkName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LinkName", 1))
	}

	if s.ParentReference == nil {
		invalidParams.Add(aws.NewErrParamRequired("ParentReference"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachObjectOutput struct {
	_ struct{} `type:"structure"`

	// The ObjectIdentifier that was detached from the object.
	DetachedObjectIdentifier *string `type:"string"`
}

// String returns the string representation
func (s DetachObjectOutput) String() string {
	return awsutil.Prettify(s)
}
