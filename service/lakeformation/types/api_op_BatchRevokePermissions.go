// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchRevokePermissionsInput struct {
	_ struct{} `type:"structure"`

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string `min:"1" type:"string"`

	// A list of up to 20 entries for resource permissions to be revoked by batch
	// operation to the principal.
	//
	// Entries is a required field
	Entries []BatchPermissionsRequestEntry `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchRevokePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchRevokePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchRevokePermissionsInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.Entries == nil {
		invalidParams.Add(aws.NewErrParamRequired("Entries"))
	}
	if s.Entries != nil {
		for i, v := range s.Entries {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entries", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchRevokePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of failures to revoke permissions to the resources.
	Failures []BatchPermissionsFailureEntry `type:"list"`
}

// String returns the string representation
func (s BatchRevokePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}