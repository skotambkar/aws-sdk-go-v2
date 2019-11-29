// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePresignedNotebookInstanceUrlInput struct {
	_ struct{} `type:"structure"`

	// The name of the notebook instance.
	//
	// NotebookInstanceName is a required field
	NotebookInstanceName *string `type:"string" required:"true"`

	// The duration of the session, in seconds. The default is 12 hours.
	SessionExpirationDurationInSeconds *int64 `min:"1800" type:"integer"`
}

// String returns the string representation
func (s CreatePresignedNotebookInstanceUrlInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePresignedNotebookInstanceUrlInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePresignedNotebookInstanceUrlInput"}

	if s.NotebookInstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotebookInstanceName"))
	}
	if s.SessionExpirationDurationInSeconds != nil && *s.SessionExpirationDurationInSeconds < 1800 {
		invalidParams.Add(aws.NewErrParamMinValue("SessionExpirationDurationInSeconds", 1800))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePresignedNotebookInstanceUrlOutput struct {
	_ struct{} `type:"structure"`

	// A JSON object that contains the URL string.
	AuthorizedUrl *string `type:"string"`
}

// String returns the string representation
func (s CreatePresignedNotebookInstanceUrlOutput) String() string {
	return awsutil.Prettify(s)
}