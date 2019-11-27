// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateProductInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The tags to add to the product.
	AddTags []Tag `type:"list"`

	// The updated description of the product.
	Description *string `type:"string"`

	// The updated distributor of the product.
	Distributor *string `type:"string"`

	// The product identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The updated product name.
	Name *string `type:"string"`

	// The updated owner of the product.
	Owner *string `type:"string"`

	// The tags to remove from the product.
	RemoveTags []string `type:"list"`

	// The updated support description for the product.
	SupportDescription *string `type:"string"`

	// The updated support email for the product.
	SupportEmail *string `type:"string"`

	// The updated support URL for the product.
	SupportUrl *string `type:"string"`
}

// String returns the string representation
func (s UpdateProductInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateProductInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateProductInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if s.AddTags != nil {
		for i, v := range s.AddTags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AddTags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateProductOutput struct {
	_ struct{} `type:"structure"`

	// Information about the product view.
	ProductViewDetail *ProductViewDetail `type:"structure"`

	// Information about the tags associated with the product.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s UpdateProductOutput) String() string {
	return awsutil.Prettify(s)
}
