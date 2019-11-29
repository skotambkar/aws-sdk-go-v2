// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteAssetInput struct {
	_ struct{} `type:"structure"`

	// AssetId is a required field
	AssetId *string `location:"uri" locationName:"AssetId" type:"string" required:"true"`

	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`

	// RevisionId is a required field
	RevisionId *string `location:"uri" locationName:"RevisionId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAssetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAssetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAssetInput"}

	if s.AssetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssetId"))
	}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if s.RevisionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RevisionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAssetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAssetOutput) String() string {
	return awsutil.Prettify(s)
}
