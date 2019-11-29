// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
)

const opDeleteLoadBalancerListeners = "DeleteLoadBalancerListeners"

// DeleteLoadBalancerListenersRequest returns a request value for making API operation for
// Elastic Load Balancing.
//
// Deletes the specified listeners from the specified load balancer.
//
//    // Example sending a request using DeleteLoadBalancerListenersRequest.
//    req := client.DeleteLoadBalancerListenersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticloadbalancing-2012-06-01/DeleteLoadBalancerListeners
func (c *Client) DeleteLoadBalancerListenersRequest(input *types.DeleteLoadBalancerListenersInput) DeleteLoadBalancerListenersRequest {
	op := &aws.Operation{
		Name:       opDeleteLoadBalancerListeners,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteLoadBalancerListenersInput{}
	}

	req := c.newRequest(op, input, &types.DeleteLoadBalancerListenersOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteLoadBalancerListenersMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteLoadBalancerListenersRequest{Request: req, Input: input, Copy: c.DeleteLoadBalancerListenersRequest}
}

// DeleteLoadBalancerListenersRequest is the request type for the
// DeleteLoadBalancerListeners API operation.
type DeleteLoadBalancerListenersRequest struct {
	*aws.Request
	Input *types.DeleteLoadBalancerListenersInput
	Copy  func(*types.DeleteLoadBalancerListenersInput) DeleteLoadBalancerListenersRequest
}

// Send marshals and sends the DeleteLoadBalancerListeners API request.
func (r DeleteLoadBalancerListenersRequest) Send(ctx context.Context) (*DeleteLoadBalancerListenersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLoadBalancerListenersResponse{
		DeleteLoadBalancerListenersOutput: r.Request.Data.(*types.DeleteLoadBalancerListenersOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLoadBalancerListenersResponse is the response type for the
// DeleteLoadBalancerListeners API operation.
type DeleteLoadBalancerListenersResponse struct {
	*types.DeleteLoadBalancerListenersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLoadBalancerListeners request.
func (r *DeleteLoadBalancerListenersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
