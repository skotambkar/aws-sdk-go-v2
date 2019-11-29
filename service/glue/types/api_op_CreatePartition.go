// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePartitionInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID of the catalog in which the partition is to be created.
	CatalogId *string `min:"1" type:"string"`

	// The name of the metadata database in which the partition is to be created.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// A PartitionInput structure defining the partition to be created.
	//
	// PartitionInput is a required field
	PartitionInput *PartitionInput `type:"structure" required:"true"`

	// The name of the metadata table in which the partition is to be created.
	//
	// TableName is a required field
	TableName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePartitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePartitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePartitionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.PartitionInput == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartitionInput"))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 1))
	}
	if s.PartitionInput != nil {
		if err := s.PartitionInput.Validate(); err != nil {
			invalidParams.AddNested("PartitionInput", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePartitionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreatePartitionOutput) String() string {
	return awsutil.Prettify(s)
}
