// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/enums"
)

// The request to update an asset.
type UpdateAssetInput struct {
	_ struct{} `type:"structure"`

	// AssetId is a required field
	AssetId *string `location:"uri" locationName:"AssetId" type:"string" required:"true"`

	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`

	// The name of the asset. When importing from Amazon S3, the S3 object key is
	// used as the asset name. When exporting to Amazon S3, the asset name is used
	// as default target S3 object key.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// RevisionId is a required field
	RevisionId *string `location:"uri" locationName:"RevisionId" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateAssetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAssetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAssetInput"}

	if s.AssetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssetId"))
	}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.RevisionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RevisionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateAssetOutput struct {
	_ struct{} `type:"structure"`

	// An Amazon Resource Name (ARN) that uniquely identifies an AWS resource.
	Arn *string `type:"string"`

	AssetDetails *AssetDetails `type:"structure"`

	// The type of file your data is stored in. Currently, the supported asset type
	// is S3_SNAPSHOT.
	AssetType enums.AssetType `type:"string" enum:"true"`

	// Dates and times in AWS Data Exchange are recorded in ISO 8601 format.
	CreatedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`

	// A unique identifier.
	DataSetId *string `type:"string"`

	// A unique identifier.
	Id *string `type:"string"`

	// The name of the asset. When importing from Amazon S3, the S3 object key is
	// used as the asset name. When exporting to Amazon S3, the asset name is used
	// as default target S3 object key.
	Name *string `type:"string"`

	// A unique identifier.
	RevisionId *string `type:"string"`

	// A unique identifier.
	SourceId *string `type:"string"`

	// Dates and times in AWS Data Exchange are recorded in ISO 8601 format.
	UpdatedAt *time.Time `type:"timestamp" timestampFormat:"iso8601"`
}

// String returns the string representation
func (s UpdateAssetOutput) String() string {
	return awsutil.Prettify(s)
}
