// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ecs/enums"
)

type DescribeServicesInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN)the cluster that hosts the
	// service to describe. If you do not specify a cluster, the default cluster
	// is assumed. This parameter is required if the service or services you are
	// describing were launched in any cluster other than the default cluster.
	Cluster *string `locationName:"cluster" type:"string"`

	// Specifies whether you want to see the resource tags for the service. If TAGS
	// is specified, the tags are included in the response. If this field is omitted,
	// tags are not included in the response.
	Include []enums.ServiceField `locationName:"include" type:"list"`

	// A list of services to describe. You may specify up to 10 services to describe
	// in a single operation.
	//
	// Services is a required field
	Services []string `locationName:"services" type:"list" required:"true"`
}

// String returns the string representation
func (s DescribeServicesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeServicesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeServicesInput"}

	if s.Services == nil {
		invalidParams.Add(aws.NewErrParamRequired("Services"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeServicesOutput struct {
	_ struct{} `type:"structure"`

	// Any failures associated with the call.
	Failures []Failure `locationName:"failures" type:"list"`

	// The list of services described.
	Services []Service `locationName:"services" type:"list"`
}

// String returns the string representation
func (s DescribeServicesOutput) String() string {
	return awsutil.Prettify(s)
}