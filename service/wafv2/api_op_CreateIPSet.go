// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateIPSetInput struct {
	_ struct{} `type:"structure"`

	// Contains an array of strings that specify one or more IP addresses or blocks
	// of IP addresses in Classless Inter-Domain Routing (CIDR) notation. AWS WAF
	// supports all address ranges for IP versions IPv4 and IPv6.
	//
	// Examples:
	//
	//    * To configure AWS WAF to allow, block, or count requests that originated
	//    from the IP address 192.0.2.44, specify 192.0.2.44/32.
	//
	//    * To configure AWS WAF to allow, block, or count requests that originated
	//    from IP addresses from 192.0.2.0 to 192.0.2.255, specify 192.0.2.0/24.
	//
	//    * To configure AWS WAF to allow, block, or count requests that originated
	//    from the IP address 1111:0000:0000:0000:0000:0000:0000:0111, specify 1111:0000:0000:0000:0000:0000:0000:0111/128.
	//
	//    * To configure AWS WAF to allow, block, or count requests that originated
	//    from IP addresses 1111:0000:0000:0000:0000:0000:0000:0000 to 1111:0000:0000:0000:ffff:ffff:ffff:ffff,
	//    specify 1111:0000:0000:0000:0000:0000:0000:0000/64.
	//
	// For more information about CIDR notation, see the Wikipedia entry Classless
	// Inter-Domain Routing (https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).
	//
	// Addresses is a required field
	Addresses []string `type:"list" required:"true"`

	// A friendly description of the IP set. You cannot change the description of
	// an IP set after you create it.
	Description *string `min:"1" type:"string"`

	// Specify IPV4 or IPV6.
	//
	// IPAddressVersion is a required field
	IPAddressVersion IPAddressVersion `type:"string" required:"true" enum:"true"`

	// A friendly name of the IP set. You cannot change the name of an IPSet after
	// you create it.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB)
	// or an API Gateway stage.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//    * CLI - Specify the region when you use the CloudFront scope: --scope=CLOUDFRONT
	//    --region=us-east-1.
	//
	//    * API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// Scope is a required field
	Scope Scope `type:"string" required:"true" enum:"true"`

	// An array of key:value pairs to associate with the resource.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateIPSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateIPSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateIPSetInput"}

	if s.Addresses == nil {
		invalidParams.Add(aws.NewErrParamRequired("Addresses"))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if len(s.IPAddressVersion) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("IPAddressVersion"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.Scope) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Scope"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateIPSetOutput struct {
	_ struct{} `type:"structure"`

	// High-level information about an IPSet, returned by operations like create
	// and list. This provides information like the ID, that you can use to retrieve
	// and manage an IPSet, and the ARN, that you provide to the IPSetReferenceStatement
	// to use the address set in a Rule.
	Summary *IPSetSummary `type:"structure"`
}

// String returns the string representation
func (s CreateIPSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateIPSet = "CreateIPSet"

// CreateIPSetRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Creates an IPSet, which you use to identify web requests that originate from
// specific IP addresses or ranges of IP addresses. For example, if you're receiving
// a lot of requests from a ranges of IP addresses, you can configure AWS WAF
// to block them using an IPSet that lists those IP addresses.
//
//    // Example sending a request using CreateIPSetRequest.
//    req := client.CreateIPSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/CreateIPSet
func (c *Client) CreateIPSetRequest(input *CreateIPSetInput) CreateIPSetRequest {
	op := &aws.Operation{
		Name:       opCreateIPSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateIPSetInput{}
	}

	req := c.newRequest(op, input, &CreateIPSetOutput{})
	return CreateIPSetRequest{Request: req, Input: input, Copy: c.CreateIPSetRequest}
}

// CreateIPSetRequest is the request type for the
// CreateIPSet API operation.
type CreateIPSetRequest struct {
	*aws.Request
	Input *CreateIPSetInput
	Copy  func(*CreateIPSetInput) CreateIPSetRequest
}

// Send marshals and sends the CreateIPSet API request.
func (r CreateIPSetRequest) Send(ctx context.Context) (*CreateIPSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateIPSetResponse{
		CreateIPSetOutput: r.Request.Data.(*CreateIPSetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateIPSetResponse is the response type for the
// CreateIPSet API operation.
type CreateIPSetResponse struct {
	*CreateIPSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateIPSet request.
func (r *CreateIPSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
