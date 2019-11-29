// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/enums"
)

type CreateIngestionInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID of the dataset used in the ingestion.
	//
	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`

	// An ID for the ingestion.
	//
	// IngestionId is a required field
	IngestionId *string `location:"uri" locationName:"IngestionId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateIngestionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateIngestionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateIngestionInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if s.IngestionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IngestionId"))
	}
	if s.IngestionId != nil && len(*s.IngestionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IngestionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateIngestionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the data ingestion.
	Arn *string `type:"string"`

	// An ID for the ingestion.
	IngestionId *string `min:"1" type:"string"`

	// The ingestion status.
	IngestionStatus enums.IngestionStatus `type:"string" enum:"true"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s CreateIngestionOutput) String() string {
	return awsutil.Prettify(s)
}
