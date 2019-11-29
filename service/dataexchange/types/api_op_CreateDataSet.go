// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/enums"
)

// A request to create a data set that contains one or more revisions.
type CreateDataSetInput struct {
	_ struct{} `type:"structure"`

	// The type of file your data is stored in. Currently, the supported asset type
	// is S3_SNAPSHOT.
	//
	// AssetType is a required field
	AssetType enums.AssetType `type:"string" required:"true" enum:"true"`

	// A description for the data set. This value can be up to 16,348 characters
	// long.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`

	// The name of the data set.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// A data set tag is an optional label that you can assign to a data set when
	// you create it. Each tag consists of a key and an optional value, both of
	// which you define. When you use tagging, you can also use tag-based access
	// control in IAM policies to control access to these data sets and revisions.
	Tags map[string]string `type:"map"`
}

// String returns the string representation
func (s CreateDataSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDataSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDataSetInput"}
	if len(s.AssetType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AssetType"))
	}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDataSetOutput struct {
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
func (s CreateDataSetOutput) String() string {
	return awsutil.Prettify(s)
}
