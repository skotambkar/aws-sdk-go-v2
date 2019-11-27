// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/forecast/enums"
)

type CreateDatasetGroupInput struct {
	_ struct{} `type:"structure"`

	// An array of Amazon Resource Names (ARNs) of the datasets that you want to
	// include in the dataset group.
	DatasetArns []string `type:"list"`

	// A name for the dataset group.
	//
	// DatasetGroupName is a required field
	DatasetGroupName *string `min:"1" type:"string" required:"true"`

	// The domain associated with the dataset group. The Domain and DatasetType
	// that you choose determine the fields that must be present in the training
	// data that you import to the dataset. For example, if you choose the RETAIL
	// domain and TARGET_TIME_SERIES as the DatasetType, Amazon Forecast requires
	// item_id, timestamp, and demand fields to be present in your data. For more
	// information, see howitworks-datasets-groups.
	//
	// Domain is a required field
	Domain enums.Domain `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateDatasetGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDatasetGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDatasetGroupInput"}

	if s.DatasetGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetGroupName"))
	}
	if s.DatasetGroupName != nil && len(*s.DatasetGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatasetGroupName", 1))
	}
	if len(s.Domain) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Domain"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDatasetGroupOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string `type:"string"`
}

// String returns the string representation
func (s CreateDatasetGroupOutput) String() string {
	return awsutil.Prettify(s)
}
