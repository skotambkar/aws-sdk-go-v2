// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package neptune

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/neptune/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
)

const opRestoreDBClusterToPointInTime = "RestoreDBClusterToPointInTime"

// RestoreDBClusterToPointInTimeRequest returns a request value for making API operation for
// Amazon Neptune.
//
// Restores a DB cluster to an arbitrary point in time. Users can restore to
// any point in time before LatestRestorableTime for up to BackupRetentionPeriod
// days. The target DB cluster is created from the source DB cluster with the
// same configuration as the original DB cluster, except that the new DB cluster
// is created with the default DB security group.
//
// This action only restores the DB cluster, not the DB instances for that DB
// cluster. You must invoke the CreateDBInstance action to create DB instances
// for the restored DB cluster, specifying the identifier of the restored DB
// cluster in DBClusterIdentifier. You can create DB instances only after the
// RestoreDBClusterToPointInTime action has completed and the DB cluster is
// available.
//
//    // Example sending a request using RestoreDBClusterToPointInTimeRequest.
//    req := client.RestoreDBClusterToPointInTimeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/neptune-2014-10-31/RestoreDBClusterToPointInTime
func (c *Client) RestoreDBClusterToPointInTimeRequest(input *types.RestoreDBClusterToPointInTimeInput) RestoreDBClusterToPointInTimeRequest {
	op := &aws.Operation{
		Name:       opRestoreDBClusterToPointInTime,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RestoreDBClusterToPointInTimeInput{}
	}

	req := c.newRequest(op, input, &types.RestoreDBClusterToPointInTimeOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.RestoreDBClusterToPointInTimeMarshaler{Input: input}.GetNamedBuildHandler())

	return RestoreDBClusterToPointInTimeRequest{Request: req, Input: input, Copy: c.RestoreDBClusterToPointInTimeRequest}
}

// RestoreDBClusterToPointInTimeRequest is the request type for the
// RestoreDBClusterToPointInTime API operation.
type RestoreDBClusterToPointInTimeRequest struct {
	*aws.Request
	Input *types.RestoreDBClusterToPointInTimeInput
	Copy  func(*types.RestoreDBClusterToPointInTimeInput) RestoreDBClusterToPointInTimeRequest
}

// Send marshals and sends the RestoreDBClusterToPointInTime API request.
func (r RestoreDBClusterToPointInTimeRequest) Send(ctx context.Context) (*RestoreDBClusterToPointInTimeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBClusterToPointInTimeResponse{
		RestoreDBClusterToPointInTimeOutput: r.Request.Data.(*types.RestoreDBClusterToPointInTimeOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBClusterToPointInTimeResponse is the response type for the
// RestoreDBClusterToPointInTime API operation.
type RestoreDBClusterToPointInTimeResponse struct {
	*types.RestoreDBClusterToPointInTimeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBClusterToPointInTime request.
func (r *RestoreDBClusterToPointInTimeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
