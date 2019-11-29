// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutInventoryInput struct {
	_ struct{} `type:"structure"`

	// An instance ID where you want to add or update inventory items.
	//
	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// The inventory items that you want to add or update on instances.
	//
	// Items is a required field
	Items []InventoryItem `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s PutInventoryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutInventoryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutInventoryInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if s.Items == nil {
		invalidParams.Add(aws.NewErrParamRequired("Items"))
	}
	if s.Items != nil && len(s.Items) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Items", 1))
	}
	if s.Items != nil {
		for i, v := range s.Items {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Items", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutInventoryOutput struct {
	_ struct{} `type:"structure"`

	// Information about the request.
	Message *string `type:"string"`
}

// String returns the string representation
func (s PutInventoryOutput) String() string {
	return awsutil.Prettify(s)
}