// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/swf/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
)

const opDescribeDomain = "DescribeDomain"

// DescribeDomainRequest returns a request value for making API operation for
// Amazon Simple Workflow Service.
//
// Returns information about the specified domain, including description and
// status.
//
// Access Control
//
// You can use IAM policies to control this action's access to Amazon SWF resources
// as follows:
//
//    * Use a Resource element with the domain name to limit the action to only
//    specified domains.
//
//    * Use an Action element to allow or deny permission to call this action.
//
//    * You cannot use an IAM policy to constrain this action's parameters.
//
// If the caller doesn't have sufficient permissions to invoke the action, or
// the parameter values fall outside the specified constraints, the action fails.
// The associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
//
//    // Example sending a request using DescribeDomainRequest.
//    req := client.DescribeDomainRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeDomainRequest(input *types.DescribeDomainInput) DescribeDomainRequest {
	op := &aws.Operation{
		Name:       opDescribeDomain,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeDomainInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDomainOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DescribeDomainMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeDomainRequest{Request: req, Input: input, Copy: c.DescribeDomainRequest}
}

// DescribeDomainRequest is the request type for the
// DescribeDomain API operation.
type DescribeDomainRequest struct {
	*aws.Request
	Input *types.DescribeDomainInput
	Copy  func(*types.DescribeDomainInput) DescribeDomainRequest
}

// Send marshals and sends the DescribeDomain API request.
func (r DescribeDomainRequest) Send(ctx context.Context) (*DescribeDomainResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDomainResponse{
		DescribeDomainOutput: r.Request.Data.(*types.DescribeDomainOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDomainResponse is the response type for the
// DescribeDomain API operation.
type DescribeDomainResponse struct {
	*types.DescribeDomainOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDomain request.
func (r *DescribeDomainResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
