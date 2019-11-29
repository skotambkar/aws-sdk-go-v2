// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetDatabasesInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog from which to retrieve Databases. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The maximum number of databases to return in one response.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetDatabasesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDatabasesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDatabasesInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetDatabasesOutput struct {
	_ struct{} `type:"structure"`

	// A list of Database objects from the specified catalog.
	//
	// DatabaseList is a required field
	DatabaseList []Database `type:"list" required:"true"`

	// A continuation token for paginating the returned list of tokens, returned
	// if the current segment of the list is not the last.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetDatabasesOutput) String() string {
	return awsutil.Prettify(s)
}
