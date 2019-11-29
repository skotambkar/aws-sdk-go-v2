// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
)

const opCreateHyperParameterTuningJob = "CreateHyperParameterTuningJob"

// CreateHyperParameterTuningJobRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Starts a hyperparameter tuning job. A hyperparameter tuning job finds the
// best version of a model by running many training jobs on your dataset using
// the algorithm you choose and values for hyperparameters within ranges that
// you specify. It then chooses the hyperparameter values that result in a model
// that performs the best, as measured by an objective metric that you choose.
//
//    // Example sending a request using CreateHyperParameterTuningJobRequest.
//    req := client.CreateHyperParameterTuningJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/CreateHyperParameterTuningJob
func (c *Client) CreateHyperParameterTuningJobRequest(input *types.CreateHyperParameterTuningJobInput) CreateHyperParameterTuningJobRequest {
	op := &aws.Operation{
		Name:       opCreateHyperParameterTuningJob,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateHyperParameterTuningJobInput{}
	}

	req := c.newRequest(op, input, &types.CreateHyperParameterTuningJobOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateHyperParameterTuningJobMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateHyperParameterTuningJobRequest{Request: req, Input: input, Copy: c.CreateHyperParameterTuningJobRequest}
}

// CreateHyperParameterTuningJobRequest is the request type for the
// CreateHyperParameterTuningJob API operation.
type CreateHyperParameterTuningJobRequest struct {
	*aws.Request
	Input *types.CreateHyperParameterTuningJobInput
	Copy  func(*types.CreateHyperParameterTuningJobInput) CreateHyperParameterTuningJobRequest
}

// Send marshals and sends the CreateHyperParameterTuningJob API request.
func (r CreateHyperParameterTuningJobRequest) Send(ctx context.Context) (*CreateHyperParameterTuningJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateHyperParameterTuningJobResponse{
		CreateHyperParameterTuningJobOutput: r.Request.Data.(*types.CreateHyperParameterTuningJobOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateHyperParameterTuningJobResponse is the response type for the
// CreateHyperParameterTuningJob API operation.
type CreateHyperParameterTuningJobResponse struct {
	*types.CreateHyperParameterTuningJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateHyperParameterTuningJob request.
func (r *CreateHyperParameterTuningJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
