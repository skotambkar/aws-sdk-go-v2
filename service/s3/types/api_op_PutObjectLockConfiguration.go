// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/s3/enums"
)

type PutObjectLockConfigurationInput struct {
	_ struct{} `type:"structure" payload:"ObjectLockConfiguration"`

	// The bucket whose Object Lock configuration you want to create or replace.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// The Object Lock configuration that you want to apply to the specified bucket.
	ObjectLockConfiguration *ObjectLockConfiguration `locationName:"ObjectLockConfiguration" type:"structure" xmlURI:"http://s3.amazonaws.com/doc/2006-03-01/"`

	// Confirms that the requester knows that she or he will be charged for the
	// request. Bucket owners need not specify this parameter in their requests.
	// Documentation on downloading objects from requester pays buckets can be found
	// at http://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html
	RequestPayer enums.RequestPayer `location:"header" locationName:"x-amz-request-payer" type:"string" enum:"true"`

	// A token to allow Object Lock to be enabled for an existing bucket.
	Token *string `location:"header" locationName:"x-amz-bucket-object-lock-token" type:"string"`
}

// String returns the string representation
func (s PutObjectLockConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutObjectLockConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutObjectLockConfigurationInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *PutObjectLockConfigurationInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type PutObjectLockConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged enums.RequestCharged `location:"header" locationName:"x-amz-request-charged" type:"string" enum:"true"`
}

// String returns the string representation
func (s PutObjectLockConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}