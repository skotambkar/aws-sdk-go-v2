// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
)

const opCreateQueryLoggingConfig = "CreateQueryLoggingConfig"

// CreateQueryLoggingConfigRequest returns a request value for making API operation for
// Amazon Route 53.
//
// Creates a configuration for DNS query logging. After you create a query logging
// configuration, Amazon Route 53 begins to publish log data to an Amazon CloudWatch
// Logs log group.
//
// DNS query logs contain information about the queries that Route 53 receives
// for a specified public hosted zone, such as the following:
//
//    * Route 53 edge location that responded to the DNS query
//
//    * Domain or subdomain that was requested
//
//    * DNS record type, such as A or AAAA
//
//    * DNS response code, such as NoError or ServFail
//
// Log Group and Resource Policy
//
// Before you create a query logging configuration, perform the following operations.
//
// If you create a query logging configuration using the Route 53 console, Route
// 53 performs these operations automatically.
//
// Create a CloudWatch Logs log group, and make note of the ARN, which you specify
// when you create a query logging configuration. Note the following:
//
//    * You must create the log group in the us-east-1 region.
//
//    * You must use the same AWS account to create the log group and the hosted
//    zone that you want to configure query logging for.
//
//    * When you create log groups for query logging, we recommend that you
//    use a consistent prefix, for example: /aws/route53/hosted zone name In
//    the next step, you'll create a resource policy, which controls access
//    to one or more log groups and the associated AWS resources, such as Route
//    53 hosted zones. There's a limit on the number of resource policies that
//    you can create, so we recommend that you use a consistent prefix so you
//    can use the same resource policy for all the log groups that you create
//    for query logging.
//
// Create a CloudWatch Logs resource policy, and give it the permissions that
// Route 53 needs to create log streams and to send query logs to log streams.
// For the value of Resource, specify the ARN for the log group that you created
// in the previous step. To use the same resource policy for all the CloudWatch
// Logs log groups that you created for query logging configurations, replace
// the hosted zone name with *, for example:
//
// arn:aws:logs:us-east-1:123412341234:log-group:/aws/route53/*
//
// You can't use the CloudWatch console to create or edit a resource policy.
// You must use the CloudWatch API, one of the AWS SDKs, or the AWS CLI.
//
// Log Streams and Edge Locations
//
// When Route 53 finishes creating the configuration for DNS query logging,
// it does the following:
//
//    * Creates a log stream for an edge location the first time that the edge
//    location responds to DNS queries for the specified hosted zone. That log
//    stream is used to log all queries that Route 53 responds to for that edge
//    location.
//
//    * Begins to send query logs to the applicable log stream.
//
// The name of each log stream is in the following format:
//
// hosted zone ID/edge location code
//
// The edge location code is a three-letter code and an arbitrarily assigned
// number, for example, DFW3. The three-letter code typically corresponds with
// the International Air Transport Association airport code for an airport near
// the edge location. (These abbreviations might change in the future.) For
// a list of edge locations, see "The Route 53 Global Network" on the Route
// 53 Product Details (http://aws.amazon.com/route53/details/) page.
//
// Queries That Are Logged
//
// Query logs contain only the queries that DNS resolvers forward to Route 53.
// If a DNS resolver has already cached the response to a query (such as the
// IP address for a load balancer for example.com), the resolver will continue
// to return the cached response. It doesn't forward another query to Route
// 53 until the TTL for the corresponding resource record set expires. Depending
// on how many DNS queries are submitted for a resource record set, and depending
// on the TTL for that resource record set, query logs might contain information
// about only one query out of every several thousand queries that are submitted
// to DNS. For more information about how DNS works, see Routing Internet Traffic
// to Your Website or Web Application (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/welcome-dns-service.html)
// in the Amazon Route 53 Developer Guide.
//
// Log File Format
//
// For a list of the values in each query log and the format of each value,
// see Logging DNS Queries (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html)
// in the Amazon Route 53 Developer Guide.
//
// Pricing
//
// For information about charges for query logs, see Amazon CloudWatch Pricing
// (http://aws.amazon.com/cloudwatch/pricing/).
//
// How to Stop Logging
//
// If you want Route 53 to stop sending query logs to CloudWatch Logs, delete
// the query logging configuration. For more information, see DeleteQueryLoggingConfig
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_DeleteQueryLoggingConfig.html).
//
//    // Example sending a request using CreateQueryLoggingConfigRequest.
//    req := client.CreateQueryLoggingConfigRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/route53-2013-04-01/CreateQueryLoggingConfig
func (c *Client) CreateQueryLoggingConfigRequest(input *types.CreateQueryLoggingConfigInput) CreateQueryLoggingConfigRequest {
	op := &aws.Operation{
		Name:       opCreateQueryLoggingConfig,
		HTTPMethod: "POST",
		HTTPPath:   "/2013-04-01/queryloggingconfig",
	}

	if input == nil {
		input = &types.CreateQueryLoggingConfigInput{}
	}

	req := c.newRequest(op, input, &types.CreateQueryLoggingConfigOutput{})
	return CreateQueryLoggingConfigRequest{Request: req, Input: input, Copy: c.CreateQueryLoggingConfigRequest}
}

// CreateQueryLoggingConfigRequest is the request type for the
// CreateQueryLoggingConfig API operation.
type CreateQueryLoggingConfigRequest struct {
	*aws.Request
	Input *types.CreateQueryLoggingConfigInput
	Copy  func(*types.CreateQueryLoggingConfigInput) CreateQueryLoggingConfigRequest
}

// Send marshals and sends the CreateQueryLoggingConfig API request.
func (r CreateQueryLoggingConfigRequest) Send(ctx context.Context) (*CreateQueryLoggingConfigResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateQueryLoggingConfigResponse{
		CreateQueryLoggingConfigOutput: r.Request.Data.(*types.CreateQueryLoggingConfigOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateQueryLoggingConfigResponse is the response type for the
// CreateQueryLoggingConfig API operation.
type CreateQueryLoggingConfigResponse struct {
	*types.CreateQueryLoggingConfigOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateQueryLoggingConfig request.
func (r *CreateQueryLoggingConfigResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
