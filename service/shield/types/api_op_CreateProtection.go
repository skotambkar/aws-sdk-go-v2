// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateProtectionInput struct {
	_ struct{} `type:"structure"`

	// Friendly name for the Protection you are creating.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The ARN (Amazon Resource Name) of the resource to be protected.
	//
	// The ARN should be in one of the following formats:
	//
	//    * For an Application Load Balancer: arn:aws:elasticloadbalancing:region:account-id:loadbalancer/app/load-balancer-name/load-balancer-id
	//
	//    * For an Elastic Load Balancer (Classic Load Balancer): arn:aws:elasticloadbalancing:region:account-id:loadbalancer/load-balancer-name
	//
	//    * For an AWS CloudFront distribution: arn:aws:cloudfront::account-id:distribution/distribution-id
	//
	//    * For an AWS Global Accelerator accelerator: arn:aws:globalaccelerator::account-id:accelerator/accelerator-id
	//
	//    * For Amazon Route 53: arn:aws:route53:::hostedzone/hosted-zone-id
	//
	//    * For an Elastic IP address: arn:aws:ec2:region:account-id:eip-allocation/allocation-id
	//
	// ResourceArn is a required field
	ResourceArn *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateProtectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateProtectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateProtectionInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.ResourceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceArn"))
	}
	if s.ResourceArn != nil && len(*s.ResourceArn) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceArn", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateProtectionOutput struct {
	_ struct{} `type:"structure"`

	// The unique identifier (ID) for the Protection object that is created.
	ProtectionId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateProtectionOutput) String() string {
	return awsutil.Prettify(s)
}