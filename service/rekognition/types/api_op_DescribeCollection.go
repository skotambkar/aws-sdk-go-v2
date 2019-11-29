// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeCollectionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the collection to describe.
	//
	// CollectionId is a required field
	CollectionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeCollectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeCollectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeCollectionInput"}

	if s.CollectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CollectionId"))
	}
	if s.CollectionId != nil && len(*s.CollectionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CollectionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeCollectionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the collection.
	CollectionARN *string `type:"string"`

	// The number of milliseconds since the Unix epoch time until the creation of
	// the collection. The Unix epoch time is 00:00:00 Coordinated Universal Time
	// (UTC), Thursday, 1 January 1970.
	CreationTimestamp *time.Time `type:"timestamp"`

	// The number of faces that are indexed into the collection. To index faces
	// into a collection, use IndexFaces.
	FaceCount *int64 `type:"long"`

	// The version of the face model that's used by the collection for face detection.
	//
	// For more information, see Model Versioning in the Amazon Rekognition Developer
	// Guide.
	FaceModelVersion *string `type:"string"`
}

// String returns the string representation
func (s DescribeCollectionOutput) String() string {
	return awsutil.Prettify(s)
}