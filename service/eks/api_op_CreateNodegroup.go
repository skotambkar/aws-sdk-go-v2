// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package eks

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/eks/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/eks/types"
)

const opCreateNodegroup = "CreateNodegroup"

// CreateNodegroupRequest returns a request value for making API operation for
// Amazon Elastic Kubernetes Service.
//
// Creates a managed worker node group for an Amazon EKS cluster. You can only
// create a node group for your cluster that is equal to the current Kubernetes
// version for the cluster. All node groups are created with the latest AMI
// release version for the respective minor Kubernetes version of the cluster.
//
// An Amazon EKS managed node group is an Amazon EC2 Auto Scaling group and
// associated Amazon EC2 instances that are managed by AWS for an Amazon EKS
// cluster. Each node group uses a version of the Amazon EKS-optimized Amazon
// Linux 2 AMI. For more information, see Managed Node Groups (https://docs.aws.amazon.com/eks/latest/userguide/managed-node-groups.html)
// in the Amazon EKS User Guide.
//
//    // Example sending a request using CreateNodegroupRequest.
//    req := client.CreateNodegroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/eks-2017-11-01/CreateNodegroup
func (c *Client) CreateNodegroupRequest(input *types.CreateNodegroupInput) CreateNodegroupRequest {
	op := &aws.Operation{
		Name:       opCreateNodegroup,
		HTTPMethod: "POST",
		HTTPPath:   "/clusters/{name}/node-groups",
	}

	if input == nil {
		input = &types.CreateNodegroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateNodegroupOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.CreateNodegroupMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateNodegroupRequest{Request: req, Input: input, Copy: c.CreateNodegroupRequest}
}

// CreateNodegroupRequest is the request type for the
// CreateNodegroup API operation.
type CreateNodegroupRequest struct {
	*aws.Request
	Input *types.CreateNodegroupInput
	Copy  func(*types.CreateNodegroupInput) CreateNodegroupRequest
}

// Send marshals and sends the CreateNodegroup API request.
func (r CreateNodegroupRequest) Send(ctx context.Context) (*CreateNodegroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNodegroupResponse{
		CreateNodegroupOutput: r.Request.Data.(*types.CreateNodegroupOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNodegroupResponse is the response type for the
// CreateNodegroup API operation.
type CreateNodegroupResponse struct {
	*types.CreateNodegroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNodegroup request.
func (r *CreateNodegroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
