// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/s3/enums"
)

type DeleteObjectsInput struct {
	_ struct{} `type:"structure" payload:"Delete"`

	// The bucket name containing the objects to delete.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Specifies whether you want to delete this object even if it has a Governance-type
	// Object Lock in place. You must have sufficient permissions to perform this
	// operation.
	BypassGovernanceRetention *bool `location:"header" locationName:"x-amz-bypass-governance-retention" type:"boolean"`

	// Container for the request.
	//
	// Delete is a required field
	Delete *Delete `locationName:"Delete" type:"structure" required:"true" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// The concatenation of the authentication device's serial number, a space,
	// and the value that is displayed on your authentication device. Required to
	// permanently delete a versioned object if versioning is configured with MFA
	// Delete enabled.
	MFA *string `location:"header" locationName:"x-amz-mfa" type:"string"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer enums.RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`
}

// String returns the string representation
func (s DeleteObjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteObjectsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteObjectsInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Delete == nil {
		invalidParams.Add(aws.NewErrParamRequired("Delete"))
	}
	if s.Delete != nil {
		if err := s.Delete.Validate(); err != nil {
			invalidParams.AddNested("Delete", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeleteObjectsInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type DeleteObjectsOutput struct {
	_ struct{} `type:"structure"`

	// Container element for a successful delete. It identifies the object that
	// was successfully deleted.
	Deleted []DeletedObject `type:"list" flattened:"true"`

	// Container for a failed delete operation that describes the object that Amazon
	// S3 attempted to delete and the error it encountered.
	Errors []Error `locationName:"Error" type:"list" flattened:"true"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged enums.RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s DeleteObjectsOutput) String() string {
	return awsutil.Prettify(s)
}
