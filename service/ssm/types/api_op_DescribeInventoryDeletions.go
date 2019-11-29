// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeInventoryDeletionsInput struct {
	_ struct{} `type:"structure"`

	// Specify the delete inventory ID for which you want information. This ID was
	// returned by the DeleteInventory action.
	DeletionId *string `type:"string"`

	// The maximum number of items to return for this call. The call also returns
	// a token that you can specify in a subsequent call to get the next set of
	// results.
	MaxResults *int64 `min:"1" type:"integer"`

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInventoryDeletionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInventoryDeletionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInventoryDeletionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeInventoryDeletionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of status items for deleted inventory.
	InventoryDeletions []InventoryDeletionStatusItem `type:"list"`

	// The token for the next set of items to return. Use this token to get the
	// next set of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeInventoryDeletionsOutput) String() string {
	return awsutil.Prettify(s)
}