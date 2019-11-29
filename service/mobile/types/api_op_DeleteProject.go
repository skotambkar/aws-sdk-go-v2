// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request structure used to request a project be deleted.
type DeleteProjectInput struct {
	_ struct{} `type:"structure"`

	// Unique project identifier.
	//
	// ProjectId is a required field
	ProjectId *string `location:"uri" locationName:"projectId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteProjectInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteProjectInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteProjectInput"}

	if s.ProjectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result structure used in response to request to delete a project.
type DeleteProjectOutput struct {
	_ struct{} `type:"structure"`

	// Resources which were deleted.
	DeletedResources []Resource `locationName:"deletedResources" type:"list"`

	// Resources which were not deleted, due to a risk of losing potentially important
	// data or files.
	OrphanedResources []Resource `locationName:"orphanedResources" type:"list"`
}

// String returns the string representation
func (s DeleteProjectOutput) String() string {
	return awsutil.Prettify(s)
}