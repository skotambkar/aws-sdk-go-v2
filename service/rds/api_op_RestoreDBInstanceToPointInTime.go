// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const opRestoreDBInstanceToPointInTime = "RestoreDBInstanceToPointInTime"

// RestoreDBInstanceToPointInTimeRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Restores a DB instance to an arbitrary point in time. You can restore to
// any point in time before the time identified by the LatestRestorableTime
// property. You can restore to a point up to the number of days specified by
// the BackupRetentionPeriod property.
//
// The target database is created with most of the original configuration, but
// in a system-selected Availability Zone, with the default security group,
// the default subnet group, and the default DB parameter group. By default,
// the new DB instance is created as a single-AZ deployment except when the
// instance is a SQL Server instance that has an option group that is associated
// with mirroring; in this case, the instance becomes a mirrored deployment
// and not a single-AZ deployment.
//
// This command doesn't apply to Aurora MySQL and Aurora PostgreSQL. For Aurora,
// use RestoreDBClusterToPointInTime.
//
//    // Example sending a request using RestoreDBInstanceToPointInTimeRequest.
//    req := client.RestoreDBInstanceToPointInTimeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/RestoreDBInstanceToPointInTime
func (c *Client) RestoreDBInstanceToPointInTimeRequest(input *types.RestoreDBInstanceToPointInTimeInput) RestoreDBInstanceToPointInTimeRequest {
	op := &aws.Operation{
		Name:       opRestoreDBInstanceToPointInTime,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.RestoreDBInstanceToPointInTimeInput{}
	}

	req := c.newRequest(op, input, &types.RestoreDBInstanceToPointInTimeOutput{})
	return RestoreDBInstanceToPointInTimeRequest{Request: req, Input: input, Copy: c.RestoreDBInstanceToPointInTimeRequest}
}

// RestoreDBInstanceToPointInTimeRequest is the request type for the
// RestoreDBInstanceToPointInTime API operation.
type RestoreDBInstanceToPointInTimeRequest struct {
	*aws.Request
	Input *types.RestoreDBInstanceToPointInTimeInput
	Copy  func(*types.RestoreDBInstanceToPointInTimeInput) RestoreDBInstanceToPointInTimeRequest
}

// Send marshals and sends the RestoreDBInstanceToPointInTime API request.
func (r RestoreDBInstanceToPointInTimeRequest) Send(ctx context.Context) (*RestoreDBInstanceToPointInTimeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreDBInstanceToPointInTimeResponse{
		RestoreDBInstanceToPointInTimeOutput: r.Request.Data.(*types.RestoreDBInstanceToPointInTimeOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreDBInstanceToPointInTimeResponse is the response type for the
// RestoreDBInstanceToPointInTime API operation.
type RestoreDBInstanceToPointInTimeResponse struct {
	*types.RestoreDBInstanceToPointInTimeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreDBInstanceToPointInTime request.
func (r *RestoreDBInstanceToPointInTimeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
