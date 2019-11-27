// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecast

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
)

const opCreateDataset = "CreateDataset"

// CreateDatasetRequest returns a request value for making API operation for
// Amazon Forecast Service.
//
// Creates an Amazon Forecast dataset. The information about the dataset that
// you provide helps Forecast understand how to consume the data for model training.
// This includes the following:
//
//    * DataFrequency - How frequently your historical time-series data is collected.
//    Amazon Forecast uses this information when training the model and generating
//    a forecast.
//
//    * Domain and DatasetType - Each dataset has an associated dataset domain
//    and a type within the domain. Amazon Forecast provides a list of predefined
//    domains and types within each domain. For each unique dataset domain and
//    type within the domain, Amazon Forecast requires your data to include
//    a minimum set of predefined fields.
//
//    * Schema - A schema specifies the fields of the dataset, including the
//    field name and data type.
//
// After creating a dataset, you import your training data into the dataset
// and add the dataset to a dataset group. You then use the dataset group to
// create a predictor. For more information, see howitworks-datasets-groups.
//
// To get a list of all your datasets, use the ListDatasets operation.
//
// The Status of a dataset must be ACTIVE before you can import training data.
// Use the DescribeDataset operation to get the status.
//
//    // Example sending a request using CreateDatasetRequest.
//    req := client.CreateDatasetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/forecast-2018-06-26/CreateDataset
func (c *Client) CreateDatasetRequest(input *types.CreateDatasetInput) CreateDatasetRequest {
	op := &aws.Operation{
		Name:       opCreateDataset,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDatasetInput{}
	}

	req := c.newRequest(op, input, &types.CreateDatasetOutput{})
	return CreateDatasetRequest{Request: req, Input: input, Copy: c.CreateDatasetRequest}
}

// CreateDatasetRequest is the request type for the
// CreateDataset API operation.
type CreateDatasetRequest struct {
	*aws.Request
	Input *types.CreateDatasetInput
	Copy  func(*types.CreateDatasetInput) CreateDatasetRequest
}

// Send marshals and sends the CreateDataset API request.
func (r CreateDatasetRequest) Send(ctx context.Context) (*CreateDatasetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDatasetResponse{
		CreateDatasetOutput: r.Request.Data.(*types.CreateDatasetOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDatasetResponse is the response type for the
// CreateDataset API operation.
type CreateDatasetResponse struct {
	*types.CreateDatasetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDataset request.
func (r *CreateDatasetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
