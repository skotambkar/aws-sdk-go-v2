// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opModifyInstanceMetadataOptions = "ModifyInstanceMetadataOptions"

// ModifyInstanceMetadataOptionsRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modify the instance metadata parameters on a running or stopped instance.
// When you modify the parameters on a stopped instance, they are applied when
// the instance is started. When you modify the parameters on a running instance,
// the API responds with a state of “pending”. After the parameter modifications
// are successfully applied to the instance, the state of the modifications
// changes from “pending” to “applied” in subsequent describe-instances
// API calls. For more information, see Instance Metadata and User Data (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html).
//
//    // Example sending a request using ModifyInstanceMetadataOptionsRequest.
//    req := client.ModifyInstanceMetadataOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyInstanceMetadataOptions
func (c *Client) ModifyInstanceMetadataOptionsRequest(input *types.ModifyInstanceMetadataOptionsInput) ModifyInstanceMetadataOptionsRequest {
	op := &aws.Operation{
		Name:       opModifyInstanceMetadataOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ModifyInstanceMetadataOptionsInput{}
	}

	req := c.newRequest(op, input, &types.ModifyInstanceMetadataOptionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.ModifyInstanceMetadataOptionsMarshaler{Input: input}.GetNamedBuildHandler())

	return ModifyInstanceMetadataOptionsRequest{Request: req, Input: input, Copy: c.ModifyInstanceMetadataOptionsRequest}
}

// ModifyInstanceMetadataOptionsRequest is the request type for the
// ModifyInstanceMetadataOptions API operation.
type ModifyInstanceMetadataOptionsRequest struct {
	*aws.Request
	Input *types.ModifyInstanceMetadataOptionsInput
	Copy  func(*types.ModifyInstanceMetadataOptionsInput) ModifyInstanceMetadataOptionsRequest
}

// Send marshals and sends the ModifyInstanceMetadataOptions API request.
func (r ModifyInstanceMetadataOptionsRequest) Send(ctx context.Context) (*ModifyInstanceMetadataOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyInstanceMetadataOptionsResponse{
		ModifyInstanceMetadataOptionsOutput: r.Request.Data.(*types.ModifyInstanceMetadataOptionsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyInstanceMetadataOptionsResponse is the response type for the
// ModifyInstanceMetadataOptions API operation.
type ModifyInstanceMetadataOptionsResponse struct {
	*types.ModifyInstanceMetadataOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyInstanceMetadataOptions request.
func (r *ModifyInstanceMetadataOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
