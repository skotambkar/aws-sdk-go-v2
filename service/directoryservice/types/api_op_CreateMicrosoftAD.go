// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/enums"
)

// Creates an AWS Managed Microsoft AD directory.
type CreateMicrosoftADInput struct {
	_ struct{} `type:"structure"`

	// A textual description for the directory. This label will appear on the AWS
	// console Directory Details page after the directory is created.
	Description *string `type:"string"`

	// AWS Managed Microsoft AD is available in two editions: Standard and Enterprise.
	// Enterprise is the default.
	Edition enums.DirectoryEdition `type:"string" enum:"true"`

	// The fully qualified domain name for the directory, such as corp.example.com.
	// This name will resolve inside your VPC only. It does not need to be publicly
	// resolvable.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`

	// The password for the default administrative user named Admin.
	//
	// Password is a required field
	Password *string `type:"string" required:"true" sensitive:"true"`

	// The NetBIOS name for your domain. A short identifier for your domain, such
	// as CORP. If you don't specify a NetBIOS name, it will default to the first
	// part of your directory DNS. For example, CORP for the directory DNS corp.example.com.
	ShortName *string `type:"string"`

	// The tags to be assigned to the AWS Managed Microsoft AD directory.
	Tags []Tag `type:"list"`

	// Contains VPC information for the CreateDirectory or CreateMicrosoftAD operation.
	//
	// VpcSettings is a required field
	VpcSettings *DirectoryVpcSettings `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateMicrosoftADInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMicrosoftADInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMicrosoftADInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}

	if s.VpcSettings == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcSettings"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.VpcSettings != nil {
		if err := s.VpcSettings.Validate(); err != nil {
			invalidParams.AddNested("VpcSettings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result of a CreateMicrosoftAD request.
type CreateMicrosoftADOutput struct {
	_ struct{} `type:"structure"`

	// The identifier of the directory that was created.
	DirectoryId *string `type:"string"`
}

// String returns the string representation
func (s CreateMicrosoftADOutput) String() string {
	return awsutil.Prettify(s)
}