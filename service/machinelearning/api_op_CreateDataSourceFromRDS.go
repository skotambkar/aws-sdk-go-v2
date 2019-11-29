// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package machinelearning

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
)

const opCreateDataSourceFromRDS = "CreateDataSourceFromRDS"

// CreateDataSourceFromRDSRequest returns a request value for making API operation for
// Amazon Machine Learning.
//
// Creates a DataSource object from an Amazon Relational Database Service (http://aws.amazon.com/rds/)
// (Amazon RDS). A DataSource references data that can be used to perform CreateMLModel,
// CreateEvaluation, or CreateBatchPrediction operations.
//
// CreateDataSourceFromRDS is an asynchronous operation. In response to CreateDataSourceFromRDS,
// Amazon Machine Learning (Amazon ML) immediately returns and sets the DataSource
// status to PENDING. After the DataSource is created and ready for use, Amazon
// ML sets the Status parameter to COMPLETED. DataSource in the COMPLETED or
// PENDING state can be used only to perform >CreateMLModel>, CreateEvaluation,
// or CreateBatchPrediction operations.
//
// If Amazon ML cannot accept the input source, it sets the Status parameter
// to FAILED and includes an error message in the Message attribute of the GetDataSource
// operation response.
//
//    // Example sending a request using CreateDataSourceFromRDSRequest.
//    req := client.CreateDataSourceFromRDSRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) CreateDataSourceFromRDSRequest(input *types.CreateDataSourceFromRDSInput) CreateDataSourceFromRDSRequest {
	op := &aws.Operation{
		Name:       opCreateDataSourceFromRDS,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDataSourceFromRDSInput{}
	}

	req := c.newRequest(op, input, &types.CreateDataSourceFromRDSOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateDataSourceFromRDSMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateDataSourceFromRDSRequest{Request: req, Input: input, Copy: c.CreateDataSourceFromRDSRequest}
}

// CreateDataSourceFromRDSRequest is the request type for the
// CreateDataSourceFromRDS API operation.
type CreateDataSourceFromRDSRequest struct {
	*aws.Request
	Input *types.CreateDataSourceFromRDSInput
	Copy  func(*types.CreateDataSourceFromRDSInput) CreateDataSourceFromRDSRequest
}

// Send marshals and sends the CreateDataSourceFromRDS API request.
func (r CreateDataSourceFromRDSRequest) Send(ctx context.Context) (*CreateDataSourceFromRDSResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDataSourceFromRDSResponse{
		CreateDataSourceFromRDSOutput: r.Request.Data.(*types.CreateDataSourceFromRDSOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDataSourceFromRDSResponse is the response type for the
// CreateDataSourceFromRDS API operation.
type CreateDataSourceFromRDSResponse struct {
	*types.CreateDataSourceFromRDSOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDataSourceFromRDS request.
func (r *CreateDataSourceFromRDSResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
