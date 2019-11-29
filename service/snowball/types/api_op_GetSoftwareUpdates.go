// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetSoftwareUpdatesInput struct {
	_ struct{} `type:"structure"`

	// The ID for a job that you want to get the software update file for, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	//
	// JobId is a required field
	JobId *string `min:"39" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSoftwareUpdatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSoftwareUpdatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSoftwareUpdatesInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if s.JobId != nil && len(*s.JobId) < 39 {
		invalidParams.Add(aws.NewErrParamMinLen("JobId", 39))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetSoftwareUpdatesOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 presigned URL for the update file associated with the specified
	// JobId value. The software update will be available for 2 days after this
	// request is made. To access an update after the 2 days have passed, you'll
	// have to make another call to GetSoftwareUpdates.
	UpdatesURI *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetSoftwareUpdatesOutput) String() string {
	return awsutil.Prettify(s)
}
