// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/datasync/enums"
)

// CreateLocationS3Request
type CreateLocationS3Input struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the Amazon S3 bucket.
	//
	// S3BucketArn is a required field
	S3BucketArn *string `type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management
	// (IAM) role that is used to access an Amazon S3 bucket.
	//
	// For detailed information about using such a role, see Creating a Location
	// for Amazon S3 in the AWS DataSync User Guide.
	//
	// S3Config is a required field
	S3Config *S3Config `type:"structure" required:"true"`

	// The Amazon S3 storage class that you want to store your files in when this
	// location is used as a task destination. For more information about S3 storage
	// classes, see Amazon S3 Storage Classes (https://aws.amazon.com/s3/storage-classes/)
	// in the Amazon Simple Storage Service Developer Guide. Some storage classes
	// have behaviors that can affect your S3 storage cost. For detailed information,
	// see using-storage-classes.
	S3StorageClass enums.S3StorageClass `type:"string" enum:"true"`

	// A subdirectory in the Amazon S3 bucket. This subdirectory in Amazon S3 is
	// used to read data from the S3 source location or write data to the S3 destination.
	Subdirectory *string `type:"string"`

	// The key-value pair that represents the tag that you want to add to the location.
	// The value can be an empty string. We recommend using tags to name your resources.
	Tags []TagListEntry `type:"list"`
}

// String returns the string representation
func (s CreateLocationS3Input) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateLocationS3Input) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateLocationS3Input"}

	if s.S3BucketArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3BucketArn"))
	}

	if s.S3Config == nil {
		invalidParams.Add(aws.NewErrParamRequired("S3Config"))
	}
	if s.S3Config != nil {
		if err := s.S3Config.Validate(); err != nil {
			invalidParams.AddNested("S3Config", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// CreateLocationS3Response
type CreateLocationS3Output struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the source Amazon S3 bucket location that
	// is created.
	LocationArn *string `type:"string"`
}

// String returns the string representation
func (s CreateLocationS3Output) String() string {
	return awsutil.Prettify(s)
}