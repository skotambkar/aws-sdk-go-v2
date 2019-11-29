// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteDataSetInput struct {
	_ struct{} `type:"structure"`

	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDataSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDataSetInput"}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDataSetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDataSetOutput) String() string {
	return awsutil.Prettify(s)
}