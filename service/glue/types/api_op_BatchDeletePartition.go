// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type BatchDeletePartitionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Data Catalog where the partition to be deleted resides. If
	// none is provided, the AWS account ID is used by default.
	CatalogId *string `min:"1" type:"string"`

	// The name of the catalog database in which the table in question resides.
	//
	// DatabaseName is a required field
	DatabaseName *string `min:"1" type:"string" required:"true"`

	// A list of PartitionInput structures that define the partitions to be deleted.
	//
	// PartitionsToDelete is a required field
	PartitionsToDelete []PartitionValueList `type:"list" required:"true"`

	// The name of the table that contains the partitions to be deleted.
	//
	// TableName is a required field
	TableName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s BatchDeletePartitionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDeletePartitionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDeletePartitionInput"}
	if s.CatalogId != nil && len(*s.CatalogId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CatalogId", 1))
	}

	if s.DatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatabaseName"))
	}
	if s.DatabaseName != nil && len(*s.DatabaseName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DatabaseName", 1))
	}

	if s.PartitionsToDelete == nil {
		invalidParams.Add(aws.NewErrParamRequired("PartitionsToDelete"))
	}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 1))
	}
	if s.PartitionsToDelete != nil {
		for i, v := range s.PartitionsToDelete {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PartitionsToDelete", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type BatchDeletePartitionOutput struct {
	_ struct{} `type:"structure"`

	// The errors encountered when trying to delete the requested partitions.
	Errors []PartitionError `type:"list"`
}

// String returns the string representation
func (s BatchDeletePartitionOutput) String() string {
	return awsutil.Prettify(s)
}