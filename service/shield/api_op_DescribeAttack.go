// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
)

const opDescribeAttack = "DescribeAttack"

// DescribeAttackRequest returns a request value for making API operation for
// AWS Shield.
//
// Describes the details of a DDoS attack.
//
//    // Example sending a request using DescribeAttackRequest.
//    req := client.DescribeAttackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/DescribeAttack
func (c *Client) DescribeAttackRequest(input *types.DescribeAttackInput) DescribeAttackRequest {
	op := &aws.Operation{
		Name:       opDescribeAttack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeAttackInput{}
	}

	req := c.newRequest(op, input, &types.DescribeAttackOutput{})
	return DescribeAttackRequest{Request: req, Input: input, Copy: c.DescribeAttackRequest}
}

// DescribeAttackRequest is the request type for the
// DescribeAttack API operation.
type DescribeAttackRequest struct {
	*aws.Request
	Input *types.DescribeAttackInput
	Copy  func(*types.DescribeAttackInput) DescribeAttackRequest
}

// Send marshals and sends the DescribeAttack API request.
func (r DescribeAttackRequest) Send(ctx context.Context) (*DescribeAttackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeAttackResponse{
		DescribeAttackOutput: r.Request.Data.(*types.DescribeAttackOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeAttackResponse is the response type for the
// DescribeAttack API operation.
type DescribeAttackResponse struct {
	*types.DescribeAttackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeAttack request.
func (r *DescribeAttackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
