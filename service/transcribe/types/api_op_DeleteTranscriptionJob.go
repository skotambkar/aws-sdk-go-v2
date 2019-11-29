// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteTranscriptionJobInput struct {
	_ struct{} `type:"structure"`

	// The name of the transcription job to be deleted.
	//
	// TranscriptionJobName is a required field
	TranscriptionJobName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTranscriptionJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTranscriptionJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTranscriptionJobInput"}

	if s.TranscriptionJobName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TranscriptionJobName"))
	}
	if s.TranscriptionJobName != nil && len(*s.TranscriptionJobName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TranscriptionJobName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTranscriptionJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTranscriptionJobOutput) String() string {
	return awsutil.Prettify(s)
}
