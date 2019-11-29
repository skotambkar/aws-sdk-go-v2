// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opCreateEndpoint = "CreateEndpoint"

// CreateEndpointRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Creates an endpoint using the endpoint configuration specified in the request.
// Amazon SageMaker uses the endpoint to provision resources and deploy models.
// You create the endpoint configuration with the CreateEndpointConfig (https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html)
// API.
//
// Use this API only for hosting models using Amazon SageMaker hosting services.
//
// You must not delete an EndpointConfig in use by an endpoint that is live
// or while the UpdateEndpoint or CreateEndpoint operations are being performed
// on the endpoint. To update an endpoint, you must create a new EndpointConfig.
//
// The endpoint name must be unique within an AWS Region in your AWS account.
//
// When it receives the request, Amazon SageMaker creates the endpoint, launches
// the resources (ML compute instances), and deploys the model(s) on them.
//
// When Amazon SageMaker receives the request, it sets the endpoint status to
// Creating. After it creates the endpoint, it sets the status to InService.
// Amazon SageMaker can then process incoming requests for inferences. To check
// the status of an endpoint, use the DescribeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/dg/API_DescribeEndpoint.html)
// API.
//
// For an example, see Exercise 1: Using the K-Means Algorithm Provided by Amazon
// SageMaker (https://docs.aws.amazon.com/sagemaker/latest/dg/ex1.html).
//
// If any of the models hosted at this endpoint get model data from an Amazon
// S3 location, Amazon SageMaker uses AWS Security Token Service to download
// model artifacts from the S3 path you provided. AWS STS is activated in your
// IAM user account by default. If you previously deactivated AWS STS for a
// region, you need to reactivate AWS STS for that region. For more information,
// see Activating and Deactivating AWS STS i an AWS Region (IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
// in the AWS Identity and Access Management User Guide.
//
//    // Example sending a request using CreateEndpointRequest.
//    req := client.CreateEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreateEndpoint
func (c *Client) CreateEndpointRequest(input *types.CreateEndpointInput) CreateEndpointRequest {
	op := &aws.Operation{
		Name:       opCreateEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateEndpointInput{}
	}

	req := c.newRequest(op, input, &types.CreateEndpointOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateEndpointMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateEndpointRequest{Request: req, Input: input, Copy: c.CreateEndpointRequest}
}

// CreateEndpointRequest is the request type for the
// CreateEndpoint API operation.
type CreateEndpointRequest struct {
	*aws.Request
	Input *types.CreateEndpointInput
	Copy  func(*types.CreateEndpointInput) CreateEndpointRequest
}

// Send marshals and sends the CreateEndpoint API request.
func (r CreateEndpointRequest) Send(ctx context.Context) (*CreateEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateEndpointResponse{
		CreateEndpointOutput: r.Request.Data.(*types.CreateEndpointOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateEndpointResponse is the response type for the
// CreateEndpoint API operation.
type CreateEndpointResponse struct {
	*types.CreateEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateEndpoint request.
func (r *CreateEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
