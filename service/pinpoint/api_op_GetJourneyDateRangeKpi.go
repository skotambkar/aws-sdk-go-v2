// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opGetJourneyDateRangeKpi = "GetJourneyDateRangeKpi"

// GetJourneyDateRangeKpiRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Retrieves (queries) pre-aggregated data for a standard engagement metric
// that applies to a journey.
//
//    // Example sending a request using GetJourneyDateRangeKpiRequest.
//    req := client.GetJourneyDateRangeKpiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/GetJourneyDateRangeKpi
func (c *Client) GetJourneyDateRangeKpiRequest(input *types.GetJourneyDateRangeKpiInput) GetJourneyDateRangeKpiRequest {
	op := &aws.Operation{
		Name:       opGetJourneyDateRangeKpi,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/apps/{application-id}/journeys/{journey-id}/kpis/daterange/{kpi-name}",
	}

	if input == nil {
		input = &types.GetJourneyDateRangeKpiInput{}
	}

	req := c.newRequest(op, input, &types.GetJourneyDateRangeKpiOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetJourneyDateRangeKpiMarshaler{Input: input}.GetNamedBuildHandler())

	return GetJourneyDateRangeKpiRequest{Request: req, Input: input, Copy: c.GetJourneyDateRangeKpiRequest}
}

// GetJourneyDateRangeKpiRequest is the request type for the
// GetJourneyDateRangeKpi API operation.
type GetJourneyDateRangeKpiRequest struct {
	*aws.Request
	Input *types.GetJourneyDateRangeKpiInput
	Copy  func(*types.GetJourneyDateRangeKpiInput) GetJourneyDateRangeKpiRequest
}

// Send marshals and sends the GetJourneyDateRangeKpi API request.
func (r GetJourneyDateRangeKpiRequest) Send(ctx context.Context) (*GetJourneyDateRangeKpiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetJourneyDateRangeKpiResponse{
		GetJourneyDateRangeKpiOutput: r.Request.Data.(*types.GetJourneyDateRangeKpiOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetJourneyDateRangeKpiResponse is the response type for the
// GetJourneyDateRangeKpi API operation.
type GetJourneyDateRangeKpiResponse struct {
	*types.GetJourneyDateRangeKpiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetJourneyDateRangeKpi request.
func (r *GetJourneyDateRangeKpiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
