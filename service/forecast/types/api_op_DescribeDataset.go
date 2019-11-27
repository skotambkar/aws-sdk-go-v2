// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/forecast/enums"
)

type DescribeDatasetInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset.
	//
	// DatasetArn is a required field
	DatasetArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDatasetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDatasetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDatasetInput"}

	if s.DatasetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDatasetOutput struct {
	_ struct{} `type:"structure"`

	// When the dataset was created.
	CreationTime *time.Time `type:"timestamp"`

	// The frequency of data collection.
	//
	// Valid intervals are Y (Year), M (Month), W (Week), D (Day), H (Hour), 30min
	// (30 minutes), 15min (15 minutes), 10min (10 minutes), 5min (5 minutes), and
	// 1min (1 minute). For example, "M" indicates every month and "30min" indicates
	// every 30 minutes.
	DataFrequency *string `type:"string"`

	// The Amazon Resource Name (ARN) of the dataset.
	DatasetArn *string `type:"string"`

	// The name of the dataset.
	DatasetName *string `min:"1" type:"string"`

	// The dataset type.
	DatasetType enums.DatasetType `type:"string" enum:"true"`

	// The dataset domain.
	Domain enums.Domain `type:"string" enum:"true"`

	// An AWS Key Management Service (KMS) key and the AWS Identity and Access Management
	// (IAM) role that Amazon Forecast can assume to access the key.
	EncryptionConfig *EncryptionConfig `type:"structure"`

	// When the dataset is created, LastModificationTime is the same as CreationTime.
	// After a CreateDatasetImportJob operation is called, LastModificationTime
	// is when the import job finished or failed. While data is being imported to
	// the dataset, LastModificationTime is the current query time.
	LastModificationTime *time.Time `type:"timestamp"`

	// An array of SchemaAttribute objects that specify the dataset fields. Each
	// SchemaAttribute specifies the name and data type of a field.
	Schema *Schema `type:"structure"`

	// The status of the dataset. States include:
	//
	//    * ACTIVE
	//
	//    * CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	//    * DELETE_PENDING, DELETE_IN_PROGRESS, DELETE_FAILED
	//
	//    * UPDATE_PENDING, UPDATE_IN_PROGRESS, UPDATE_FAILED
	//
	// The UPDATE states apply while data is imported to the dataset from a call
	// to the CreateDatasetImportJob operation. During this time, the status reflects
	// the status of the dataset import job. For example, when the import job status
	// is CREATE_IN_PROGRESS, the status of the dataset is UPDATE_IN_PROGRESS.
	//
	// The Status of the dataset must be ACTIVE before you can import training data.
	Status *string `type:"string"`
}

// String returns the string representation
func (s DescribeDatasetOutput) String() string {
	return awsutil.Prettify(s)
}
