// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateServiceActionInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// A map that defines the self-service action.
	Definition map[string]string `min:"1" type:"map"`

	// The self-service action description.
	Description *string `type:"string"`

	// The self-service action identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The self-service action name.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateServiceActionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateServiceActionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateServiceActionInput"}
	if s.Definition != nil && len(s.Definition) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Definition", 1))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateServiceActionOutput struct {
	_ struct{} `type:"structure"`

	// Detailed information about the self-service action.
	ServiceActionDetail *ServiceActionDetail `type:"structure"`
}

// String returns the string representation
func (s UpdateServiceActionOutput) String() string {
	return awsutil.Prettify(s)
}