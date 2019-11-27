// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteAttributesInput struct {
	_ struct{} `type:"structure"`

	// A list of Attributes. Similar to columns on a spreadsheet, attributes represent
	// categories of data that can be assigned to items.
	Attributes []DeletableAttribute `locationNameList:"Attribute" type:"list" flattened:"true"`

	// The name of the domain in which to perform the operation.
	//
	// DomainName is a required field
	DomainName *string `type:"string" required:"true"`

	// The update condition which, if specified, determines whether the specified
	// attributes will be deleted or not. The update condition must be satisfied
	// in order for this request to be processed and the attributes to be deleted.
	Expected *UpdateCondition `type:"structure"`

	// The name of the item. Similar to rows on a spreadsheet, items represent individual
	// objects that contain one or more value-attribute pairs.
	//
	// ItemName is a required field
	ItemName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAttributesInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}

	if s.ItemName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ItemName"))
	}
	if s.Attributes != nil {
		for i, v := range s.Attributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteAttributesOutput) String() string {
	return awsutil.Prettify(s)
}
