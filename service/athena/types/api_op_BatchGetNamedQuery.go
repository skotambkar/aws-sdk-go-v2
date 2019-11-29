// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchGetNamedQueryInput struct {
	_ struct{} `type:"structure"`

	// An array of query IDs.
	//
	// NamedQueryIds is a required field
	NamedQueryIds []string `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchGetNamedQueryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchGetNamedQueryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchGetNamedQueryInput"}

	if s.NamedQueryIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("NamedQueryIds"))
	}
	if s.NamedQueryIds != nil && len(s.NamedQueryIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NamedQueryIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchGetNamedQueryOutput struct {
	_ struct{} `type:"structure"`

	// Information about the named query IDs submitted.
	NamedQueries []NamedQuery `type:"list"`

	// Information about provided query IDs.
	UnprocessedNamedQueryIds []UnprocessedNamedQueryId `type:"list"`
}

// String returns the string representation
func (s BatchGetNamedQueryOutput) String() string {
	return awsutil.Prettify(s)
}