// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
)

const opPutAnomalyDetector = "PutAnomalyDetector"

// PutAnomalyDetectorRequest returns a request value for making API operation for
// Amazon CloudWatch.
//
// Creates an anomaly detection model for a CloudWatch metric. You can use the
// model to display a band of expected normal values when the metric is graphed.
//
// For more information, see CloudWatch Anomaly Detection (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Anomaly_Detection.html).
//
//    // Example sending a request using PutAnomalyDetectorRequest.
//    req := client.PutAnomalyDetectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/monitoring-2010-08-01/PutAnomalyDetector
func (c *Client) PutAnomalyDetectorRequest(input *types.PutAnomalyDetectorInput) PutAnomalyDetectorRequest {
	op := &aws.Operation{
		Name:       opPutAnomalyDetector,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutAnomalyDetectorInput{}
	}

	req := c.newRequest(op, input, &types.PutAnomalyDetectorOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.PutAnomalyDetectorMarshaler{Input: input}.GetNamedBuildHandler())

	return PutAnomalyDetectorRequest{Request: req, Input: input, Copy: c.PutAnomalyDetectorRequest}
}

// PutAnomalyDetectorRequest is the request type for the
// PutAnomalyDetector API operation.
type PutAnomalyDetectorRequest struct {
	*aws.Request
	Input *types.PutAnomalyDetectorInput
	Copy  func(*types.PutAnomalyDetectorInput) PutAnomalyDetectorRequest
}

// Send marshals and sends the PutAnomalyDetector API request.
func (r PutAnomalyDetectorRequest) Send(ctx context.Context) (*PutAnomalyDetectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutAnomalyDetectorResponse{
		PutAnomalyDetectorOutput: r.Request.Data.(*types.PutAnomalyDetectorOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutAnomalyDetectorResponse is the response type for the
// PutAnomalyDetector API operation.
type PutAnomalyDetectorResponse struct {
	*types.PutAnomalyDetectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutAnomalyDetector request.
func (r *PutAnomalyDetectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
