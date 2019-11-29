// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/enums"
)

type GetDataSetInput struct {
	_ struct{} `type:"structure"`

	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDataSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDataSetInput"}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDataSetOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies an AWS resource.
	Arn *string `type:"string"`

	// The type of file your data is stored in. Currently, the supported asset type
	// is S3_SNAPSHOT.
	AssetType enums.AssetType `type:"string" enum:"true"`

	// Dates and times in AWS Data Exchange are recorded in ISO 8601 format.
	CreatedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// A description of a resource.
	Description *string `type:"string"`

	// A unique identifier.
	Id *string `type:"string"`

	// The name of the model.
	Name *string `type:"string"`

	// A property that defines the data set as OWNED by the account (for providers)
	// or ENTITLED to the account (for subscribers). When an owned data set is published
	// in a product, AWS Data Exchange creates a copy of the data set. Subscribers
	// can access that copy of the data set as an entitled data set.
	Origin enums.Origin `type:"string" enum:"true"`

	OriginDetails *OriginDetails `type:"structure"`

	// A unique identifier.
	SourceId *string `type:"string"`

	Tags map[string]string `type:"map"`

	// Dates and times in AWS Data Exchange are recorded in ISO 8601 format.
	UpdatedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s GetDataSetOutput) String() string {
	return awsutil.Prettify(s)
}