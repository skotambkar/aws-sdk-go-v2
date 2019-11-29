// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetObjectTaggingInput struct {
	_ struct{} `type:"structure"`

	// The bucket name containing the object for which to get the tagging information.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`

	// Object key for which to get the tagging information.
	//
	// Key is a required field
	Key *string `location:"uri" locationName:"Key" min:"1" type:"string" required:"true"`

	// The versionId of the object for which to get the tagging information.
	VersionId *string `location:"querystring" locationName:"versionId" type:"string"`
}

// String returns the string representation
func (s GetObjectTaggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetObjectTaggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetObjectTaggingInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if s.Key == nil {
		invalidParams.Add(aws.NewErrParamRequired("Key"))
	}
	if s.Key != nil && len(*s.Key) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Key", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetObjectTaggingInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type GetObjectTaggingOutput struct {
	_ struct{} `type:"structure"`

	// Contains the tag set.
	//
	// TagSet is a required field
	TagSet []Tag `locationNameList:"Tag" type:"list" required:"true"`

	// The versionId of the object for which you got the tagging information.
	VersionId *string `location:"header" locationName:"x-amz-version-id" type:"string"`
}

// String returns the string representation
func (s GetObjectTaggingOutput) String() string {
	return awsutil.Prettify(s)
}
