// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package migrationhub

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/migrationhub/types"
)

const opAssociateCreatedArtifact = "AssociateCreatedArtifact"

// AssociateCreatedArtifactRequest returns a request value for making API operation for
// AWS Migration Hub.
//
// Associates a created artifact of an AWS cloud resource, the target receiving
// the migration, with the migration task performed by a migration tool. This
// API has the following traits:
//
//    * Migration tools can call the AssociateCreatedArtifact operation to indicate
//    which AWS artifact is associated with a migration task.
//
//    * The created artifact name must be provided in ARN (Amazon Resource Name)
//    format which will contain information about type and region; for example:
//    arn:aws:ec2:us-east-1:488216288981:image/ami-6d0ba87b.
//
//    * Examples of the AWS resource behind the created artifact are, AMI's,
//    EC2 instance, or DMS endpoint, etc.
//
//    // Example sending a request using AssociateCreatedArtifactRequest.
//    req := client.AssociateCreatedArtifactRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/AWSMigrationHub-2017-05-31/AssociateCreatedArtifact
func (c *Client) AssociateCreatedArtifactRequest(input *types.AssociateCreatedArtifactInput) AssociateCreatedArtifactRequest {
	op := &aws.Operation{
		Name:       opAssociateCreatedArtifact,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssociateCreatedArtifactInput{}
	}

	req := c.newRequest(op, input, &types.AssociateCreatedArtifactOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.AssociateCreatedArtifactMarshaler{Input: input}.GetNamedBuildHandler())

	return AssociateCreatedArtifactRequest{Request: req, Input: input, Copy: c.AssociateCreatedArtifactRequest}
}

// AssociateCreatedArtifactRequest is the request type for the
// AssociateCreatedArtifact API operation.
type AssociateCreatedArtifactRequest struct {
	*aws.Request
	Input *types.AssociateCreatedArtifactInput
	Copy  func(*types.AssociateCreatedArtifactInput) AssociateCreatedArtifactRequest
}

// Send marshals and sends the AssociateCreatedArtifact API request.
func (r AssociateCreatedArtifactRequest) Send(ctx context.Context) (*AssociateCreatedArtifactResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateCreatedArtifactResponse{
		AssociateCreatedArtifactOutput: r.Request.Data.(*types.AssociateCreatedArtifactOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateCreatedArtifactResponse is the response type for the
// AssociateCreatedArtifact API operation.
type AssociateCreatedArtifactResponse struct {
	*types.AssociateCreatedArtifactOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateCreatedArtifact request.
func (r *AssociateCreatedArtifactResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
