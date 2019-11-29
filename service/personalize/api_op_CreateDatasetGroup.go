// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
)

const opCreateDatasetGroup = "CreateDatasetGroup"

// CreateDatasetGroupRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Creates an empty dataset group. A dataset group contains related datasets
// that supply data for training a model. A dataset group can contain at most
// three datasets, one for each type of dataset:
//
//    * Interactions
//
//    * Items
//
//    * Users
//
// To train a model (create a solution), a dataset group that contains an Interactions
// dataset is required. Call CreateDataset to add a dataset to the group.
//
// A dataset group can be in one of the following states:
//
//    * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or- CREATE FAILED
//
//    * DELETE PENDING
//
// To get the status of the dataset group, call DescribeDatasetGroup. If the
// status shows as CREATE FAILED, the response includes a failureReason key,
// which describes why the creation failed.
//
// You must wait until the status of the dataset group is ACTIVE before adding
// a dataset to the group.
//
// You can specify an AWS Key Management Service (KMS) key to encrypt the datasets
// in the group. If you specify a KMS key, you must also include an AWS Identity
// and Access Management (IAM) role that has permission to access the key.
//
// APIs that require a dataset group ARN in the request
//
//    * CreateDataset
//
//    * CreateEventTracker
//
//    * CreateSolution
//
// Related APIs
//
//    * ListDatasetGroups
//
//    * DescribeDatasetGroup
//
//    * DeleteDatasetGroup
//
//    // Example sending a request using CreateDatasetGroupRequest.
//    req := client.CreateDatasetGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/CreateDatasetGroup
func (c *Client) CreateDatasetGroupRequest(input *types.CreateDatasetGroupInput) CreateDatasetGroupRequest {
	op := &aws.Operation{
		Name:       opCreateDatasetGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateDatasetGroupInput{}
	}

	req := c.newRequest(op, input, &types.CreateDatasetGroupOutput{})
	return CreateDatasetGroupRequest{Request: req, Input: input, Copy: c.CreateDatasetGroupRequest}
}

// CreateDatasetGroupRequest is the request type for the
// CreateDatasetGroup API operation.
type CreateDatasetGroupRequest struct {
	*aws.Request
	Input *types.CreateDatasetGroupInput
	Copy  func(*types.CreateDatasetGroupInput) CreateDatasetGroupRequest
}

// Send marshals and sends the CreateDatasetGroup API request.
func (r CreateDatasetGroupRequest) Send(ctx context.Context) (*CreateDatasetGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDatasetGroupResponse{
		CreateDatasetGroupOutput: r.Request.Data.(*types.CreateDatasetGroupOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDatasetGroupResponse is the response type for the
// CreateDatasetGroup API operation.
type CreateDatasetGroupResponse struct {
	*types.CreateDatasetGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDatasetGroup request.
func (r *CreateDatasetGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
