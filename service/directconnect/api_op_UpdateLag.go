// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
)

const opUpdateLag = "UpdateLag"

// UpdateLagRequest returns a request value for making API operation for
// AWS Direct Connect.
//
// Updates the attributes of the specified link aggregation group (LAG).
//
// You can update the following attributes:
//
//    * The name of the LAG.
//
//    * The value for the minimum number of connections that must be operational
//    for the LAG itself to be operational.
//
// When you create a LAG, the default value for the minimum number of operational
// connections is zero (0). If you update this value and the number of operational
// connections falls below the specified value, the LAG automatically goes down
// to avoid over-utilization of the remaining connections. Adjust this value
// with care, as it could force the LAG down if it is set higher than the current
// number of operational connections.
//
//    // Example sending a request using UpdateLagRequest.
//    req := client.UpdateLagRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/directconnect-2012-10-25/UpdateLag
func (c *Client) UpdateLagRequest(input *types.UpdateLagInput) UpdateLagRequest {
	op := &aws.Operation{
		Name:       opUpdateLag,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateLagInput{}
	}

	req := c.newRequest(op, input, &types.UpdateLagOutput{})
	return UpdateLagRequest{Request: req, Input: input, Copy: c.UpdateLagRequest}
}

// UpdateLagRequest is the request type for the
// UpdateLag API operation.
type UpdateLagRequest struct {
	*aws.Request
	Input *types.UpdateLagInput
	Copy  func(*types.UpdateLagInput) UpdateLagRequest
}

// Send marshals and sends the UpdateLag API request.
func (r UpdateLagRequest) Send(ctx context.Context) (*UpdateLagResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateLagResponse{
		UpdateLagOutput: r.Request.Data.(*types.UpdateLagOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateLagResponse is the response type for the
// UpdateLag API operation.
type UpdateLagResponse struct {
	*types.UpdateLagOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateLag request.
func (r *UpdateLagResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
