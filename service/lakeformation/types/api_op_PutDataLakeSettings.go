// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutDataLakeSettingsInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string `min:"1" type:"string"`

	// A list of AWS Lake Formation principals.
	//
	// DataLakeSettings is a required field
	DataLakeSettings *DataLakeSettings `type:"structure" required:"true"`
}

// String returns the string representation
func (s PutDataLakeSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutDataLakeSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutDataLakeSettingsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DataLakeSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataLakeSettings"))
	}
	if s.DataLakeSettings != nil {
		if err := s.DataLakeSettings.Validate(); err != nil {
			invalidParams.AddNested("DataLakeSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutDataLakeSettingsOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutDataLakeSettingsOutput) String() string {
	return awsutil.Prettify(s)
}
