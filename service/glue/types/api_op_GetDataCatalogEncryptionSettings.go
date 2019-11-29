// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDataCatalogEncryptionSettingsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog to retrieve the security configuration for. If
	// none is provided, the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetDataCatalogEncryptionSettingsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDataCatalogEncryptionSettingsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDataCatalogEncryptionSettingsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDataCatalogEncryptionSettingsOutput struct {
	_ struct{} `type:"structure"`

	// The requested security configuration.
	DataCatalogEncryptionSettings *DataCatalogEncryptionSettings `type:"structure"`
}

// String returns the string representation
func (s GetDataCatalogEncryptionSettingsOutput) String() string {
	return awsutil.Prettify(s)
}