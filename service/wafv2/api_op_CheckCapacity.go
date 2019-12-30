// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CheckCapacityInput struct {
	_ struct{} `type:"structure"`

	// An array of Rule that you're configuring to use in a rule group or web ACL.
	//
	// Rules is a required field
	Rules []Rule `type:"list" required:"true"`

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
}

// String returns the string representation
func (s CheckCapacityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CheckCapacityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CheckCapacityInput"}

	if s.Rules == nil {
		invalidParams.Add(aws.NewErrParamRequired("Rules"))
	}
	if len(s.Scope) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Scope"))
	}
	if s.Rules != nil {
		for i, v := range s.Rules {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Rules", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CheckCapacityOutput struct {
	_ struct{} `type:"structure"`

	// The capacity required by the rules and scope.
	Capacity *int64 `type:"long"`
}

// String returns the string representation
func (s CheckCapacityOutput) String() string {
	return awsutil.Prettify(s)
}

const opCheckCapacity = "CheckCapacity"

// CheckCapacityRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Returns the web ACL capacity unit (WCU) requirements for a specified scope
// and set of rules. You can use this to check the capacity requirements for
// the rules you want to use in a RuleGroup or WebACL.
//
// AWS WAF uses WCUs to calculate and control the operating resources that are
// used to run your rules, rule groups, and web ACLs. AWS WAF calculates capacity
// differently for each rule type, to reflect the relative cost of each rule.
// Simple rules that cost little to run use fewer WCUs than more complex rules
// that use more processing power. Rule group capacity is fixed at creation,
// which helps users plan their web ACL WCU usage when they use a rule group.
// The WCU limit for web ACLs is 1,500.
//
//    // Example sending a request using CheckCapacityRequest.
//    req := client.CheckCapacityRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/CheckCapacity
func (c *Client) CheckCapacityRequest(input *CheckCapacityInput) CheckCapacityRequest {
	op := &aws.Operation{
		Name:       opCheckCapacity,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CheckCapacityInput{}
	}

	req := c.newRequest(op, input, &CheckCapacityOutput{})
	return CheckCapacityRequest{Request: req, Input: input, Copy: c.CheckCapacityRequest}
}

// CheckCapacityRequest is the request type for the
// CheckCapacity API operation.
type CheckCapacityRequest struct {
	*aws.Request
	Input *CheckCapacityInput
	Copy  func(*CheckCapacityInput) CheckCapacityRequest
}

// Send marshals and sends the CheckCapacity API request.
func (r CheckCapacityRequest) Send(ctx context.Context) (*CheckCapacityResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CheckCapacityResponse{
		CheckCapacityOutput: r.Request.Data.(*CheckCapacityOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CheckCapacityResponse is the response type for the
// CheckCapacity API operation.
type CheckCapacityResponse struct {
	*CheckCapacityOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CheckCapacity request.
func (r *CheckCapacityResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
