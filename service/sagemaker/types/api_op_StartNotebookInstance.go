// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartNotebookInstanceInput struct {
	_ struct{} `type:"structure"`

	// The name of the notebook instance to start.
	//
	// NotebookInstanceName is a required field
	NotebookInstanceName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s StartNotebookInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartNotebookInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartNotebookInstanceInput"}

	if s.NotebookInstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("NotebookInstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartNotebookInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartNotebookInstanceOutput) String() string {
	return awsutil.Prettify(s)
}