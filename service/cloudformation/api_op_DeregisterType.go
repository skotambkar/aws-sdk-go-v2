// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

const opDeregisterType = "DeregisterType"

// DeregisterTypeRequest returns a request value for making API operation for
// AWS CloudFormation.
//
// Removes a type or type version from active use in the CloudFormation registry.
// If a type or type version is deregistered, it cannot be used in CloudFormation
// operations.
//
// To deregister a type, you must individually deregister all registered versions
// of that type. If a type has only a single registered version, deregistering
// that version results in the type itself being deregistered.
//
// You cannot deregister the default version of a type, unless it is the only
// registered version of that type, in which case the type itself is deregistered
// as well.
//
//    // Example sending a request using DeregisterTypeRequest.
//    req := client.DeregisterTypeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15/DeregisterType
func (c *Client) DeregisterTypeRequest(input *types.DeregisterTypeInput) DeregisterTypeRequest {
	op := &aws.Operation{
		Name:       opDeregisterType,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeregisterTypeInput{}
	}

	req := c.newRequest(op, input, &types.DeregisterTypeOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeregisterTypeMarshaler{Input: input}.GetNamedBuildHandler())

	return DeregisterTypeRequest{Request: req, Input: input, Copy: c.DeregisterTypeRequest}
}

// DeregisterTypeRequest is the request type for the
// DeregisterType API operation.
type DeregisterTypeRequest struct {
	*aws.Request
	Input *types.DeregisterTypeInput
	Copy  func(*types.DeregisterTypeInput) DeregisterTypeRequest
}

// Send marshals and sends the DeregisterType API request.
func (r DeregisterTypeRequest) Send(ctx context.Context) (*DeregisterTypeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeregisterTypeResponse{
		DeregisterTypeOutput: r.Request.Data.(*types.DeregisterTypeOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeregisterTypeResponse is the response type for the
// DeregisterType API operation.
type DeregisterTypeResponse struct {
	*types.DeregisterTypeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeregisterType request.
func (r *DeregisterTypeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
