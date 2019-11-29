// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

const opUpdateNodegroupConfig = "UpdateNodegroupConfig"

// UpdateNodegroupConfigRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Updates an Amazon EKS managed node group configuration. Your node group continues
// to function during the update. The response output includes an update ID
// that you can use to track the status of your node group update with the DescribeUpdate
// API operation. Currently you can update the Kubernetes labels for a node
// group or the scaling configuration.
//
//    // Example sending a request using UpdateNodegroupConfigRequest.
//    req := client.UpdateNodegroupConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/UpdateNodegroupConfig
func (c *Client) UpdateNodegroupConfigRequest(input *types.UpdateNodegroupConfigInput) UpdateNodegroupConfigRequest {
	op := &aws.Operation{
		Name:       opUpdateNodegroupConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/clusters/{name}/node-groups/{nodegroupName}/update-config",
	}

	if input == nil {
		input = &types.UpdateNodegroupConfigInput{}
	}

	req := c.newRequest(op, input, &types.UpdateNodegroupConfigOutput{})
	return UpdateNodegroupConfigRequest{Request: req, Input: input, Copy: c.UpdateNodegroupConfigRequest}
}

// UpdateNodegroupConfigRequest is the request type for the
// UpdateNodegroupConfig API operation.
type UpdateNodegroupConfigRequest struct {
	*aws.Request
	Input *types.UpdateNodegroupConfigInput
	Copy  func(*types.UpdateNodegroupConfigInput) UpdateNodegroupConfigRequest
}

// Send marshals and sends the UpdateNodegroupConfig API request.
func (r UpdateNodegroupConfigRequest) Send(ctx context.Context) (*UpdateNodegroupConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateNodegroupConfigResponse{
		UpdateNodegroupConfigOutput: r.Request.Data.(*types.UpdateNodegroupConfigOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateNodegroupConfigResponse is the response type for the
// UpdateNodegroupConfig API operation.
type UpdateNodegroupConfigResponse struct {
	*types.UpdateNodegroupConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateNodegroupConfig request.
func (r *UpdateNodegroupConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
