// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDescribeFpgaImageAttribute = "DescribeFpgaImageAttribute"

// DescribeFpgaImageAttributeRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Describes the specified attribute of the specified Amazon FPGA Image (AFI).
//
//    // Example sending a request using DescribeFpgaImageAttributeRequest.
//    req := client.DescribeFpgaImageAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DescribeFpgaImageAttribute
func (c *Client) DescribeFpgaImageAttributeRequest(input *types.DescribeFpgaImageAttributeInput) DescribeFpgaImageAttributeRequest {
	op := &aws.Operation{
		Name:       opDescribeFpgaImageAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DescribeFpgaImageAttributeInput{}
	}

	req := c.newRequest(op, input, &types.DescribeFpgaImageAttributeOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DescribeFpgaImageAttributeMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeFpgaImageAttributeRequest{Request: req, Input: input, Copy: c.DescribeFpgaImageAttributeRequest}
}

// DescribeFpgaImageAttributeRequest is the request type for the
// DescribeFpgaImageAttribute API operation.
type DescribeFpgaImageAttributeRequest struct {
	*aws.Request
	Input *types.DescribeFpgaImageAttributeInput
	Copy  func(*types.DescribeFpgaImageAttributeInput) DescribeFpgaImageAttributeRequest
}

// Send marshals and sends the DescribeFpgaImageAttribute API request.
func (r DescribeFpgaImageAttributeRequest) Send(ctx context.Context) (*DescribeFpgaImageAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFpgaImageAttributeResponse{
		DescribeFpgaImageAttributeOutput: r.Request.Data.(*types.DescribeFpgaImageAttributeOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeFpgaImageAttributeResponse is the response type for the
// DescribeFpgaImageAttribute API operation.
type DescribeFpgaImageAttributeResponse struct {
	*types.DescribeFpgaImageAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFpgaImageAttribute request.
func (r *DescribeFpgaImageAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
