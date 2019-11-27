// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeProvisioningParametersInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The path identifier of the product. This value is optional if the product
	// has a default path, and required if the product has more than one path. To
	// list the paths for a product, use ListLaunchPaths.
	PathId *string `min:"1" type:"string"`

	// The product identifier.
	//
	// ProductId is a required field
	ProductId *string `min:"1" type:"string" required:"true"`

	// The identifier of the provisioning artifact.
	//
	// ProvisioningArtifactId is a required field
	ProvisioningArtifactId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeProvisioningParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProvisioningParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProvisioningParametersInput"}
	if s.PathId != nil && len(*s.PathId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PathId", 1))
	}

	if s.ProductId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductId"))
	}
	if s.ProductId != nil && len(*s.ProductId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductId", 1))
	}

	if s.ProvisioningArtifactId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProvisioningArtifactId"))
	}
	if s.ProvisioningArtifactId != nil && len(*s.ProvisioningArtifactId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProvisioningArtifactId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeProvisioningParametersOutput struct {
	_ struct{} `type:"structure"`

	// Information about the constraints used to provision the product.
	ConstraintSummaries []ConstraintSummary `type:"list"`

	// Information about the parameters used to provision the product.
	ProvisioningArtifactParameters []ProvisioningArtifactParameter `type:"list"`

	// An object that contains information about preferences, such as regions and
	// accounts, for the provisioning artifact.
	ProvisioningArtifactPreferences *ProvisioningArtifactPreferences `type:"structure"`

	// Information about the TagOptions associated with the resource.
	TagOptions []TagOptionSummary `type:"list"`

	// Any additional metadata specifically related to the provisioning of the product.
	// For example, see the Version field of the CloudFormation template.
	UsageInstructions []UsageInstruction `type:"list"`
}

// String returns the string representation
func (s DescribeProvisioningParametersOutput) String() string {
	return awsutil.Prettify(s)
}
