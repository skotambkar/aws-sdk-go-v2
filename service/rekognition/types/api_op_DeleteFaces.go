// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFacesInput struct {
	_ struct{} `type:"structure"`

	// Collection from which to remove the specific faces.
	//
	// CollectionId is a required field
	CollectionId *string `min:"1" type:"string" required:"true"`

	// An array of face IDs to delete.
	//
	// FaceIds is a required field
	FaceIds []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s DeleteFacesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFacesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFacesInput"}

	if s.CollectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CollectionId"))
	}
	if s.CollectionId != nil && len(*s.CollectionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CollectionId", 1))
	}

	if s.FaceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("FaceIds"))
	}
	if s.FaceIds != nil && len(s.FaceIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FaceIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFacesOutput struct {
	_ struct{} `type:"structure"`

	// An array of strings (face IDs) of the faces that were deleted.
	DeletedFaces []string `min:"1" type:"list"`
}

// String returns the string representation
func (s DeleteFacesOutput) String() string {
	return awsutil.Prettify(s)
}
