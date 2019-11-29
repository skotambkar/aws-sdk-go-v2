// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const opPutItem = "PutItem"

// PutItemRequest returns a request value for making API operation for
// Amazon DynamoDB.
//
// Creates a new item, or replaces an old item with a new item. If an item that
// has the same primary key as the new item already exists in the specified
// table, the new item completely replaces the existing item. You can perform
// a conditional put operation (add a new item if one with the specified primary
// key doesn't exist), or replace an existing item if it has certain attribute
// values. You can return the item's attribute values in the same operation,
// using the ReturnValues parameter.
//
// This topic provides general information about the PutItem API.
//
// For information on how to call the PutItem API using the AWS SDK in specific
// languages, see the following:
//
//    * PutItem in the AWS Command Line Interface (http://docs.aws.amazon.com/goto/aws-cli/dynamodb-2012-08-10/PutItem)
//
//    * PutItem in the AWS SDK for .NET (http://docs.aws.amazon.com/goto/DotNetSDKV3/dynamodb-2012-08-10/PutItem)
//
//    * PutItem in the AWS SDK for C++ (http://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/PutItem)
//
//    * PutItem in the AWS SDK for Go (http://docs.aws.amazon.com/goto/SdkForGoV1/dynamodb-2012-08-10/PutItem)
//
//    * PutItem in the AWS SDK for Java (http://docs.aws.amazon.com/goto/SdkForJava/dynamodb-2012-08-10/PutItem)
//
//    * PutItem in the AWS SDK for JavaScript (http://docs.aws.amazon.com/goto/AWSJavaScriptSDK/dynamodb-2012-08-10/PutItem)
//
//    * PutItem in the AWS SDK for PHP V3 (http://docs.aws.amazon.com/goto/SdkForPHPV3/dynamodb-2012-08-10/PutItem)
//
//    * PutItem in the AWS SDK for Python (http://docs.aws.amazon.com/goto/boto3/dynamodb-2012-08-10/PutItem)
//
//    * PutItem in the AWS SDK for Ruby V2 (http://docs.aws.amazon.com/goto/SdkForRubyV2/dynamodb-2012-08-10/PutItem)
//
// When you add an item, the primary key attributes are the only required attributes.
// Attribute values cannot be null. String and Binary type attributes must have
// lengths greater than zero. Set type attributes cannot be empty. Requests
// with empty values will be rejected with a ValidationException exception.
//
// To prevent a new item from replacing an existing item, use a conditional
// expression that contains the attribute_not_exists function with the name
// of the attribute being used as the partition key for the table. Since every
// record must contain that attribute, the attribute_not_exists function will
// only succeed if no matching item exists.
//
// For more information about PutItem, see Working with Items (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithItems.html)
// in the Amazon DynamoDB Developer Guide.
//
//    // Example sending a request using PutItemRequest.
//    req := client.PutItemRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dynamodb-2012-08-10/PutItem
func (c *Client) PutItemRequest(input *types.PutItemInput) PutItemRequest {
	op := &aws.Operation{
		Name:       opPutItem,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutItemInput{}
	}

	req := c.newRequest(op, input, &types.PutItemOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.PutItemMarshaler{Input: input}.GetNamedBuildHandler())

	return PutItemRequest{Request: req, Input: input, Copy: c.PutItemRequest}
}

// PutItemRequest is the request type for the
// PutItem API operation.
type PutItemRequest struct {
	*aws.Request
	Input *types.PutItemInput
	Copy  func(*types.PutItemInput) PutItemRequest
}

// Send marshals and sends the PutItem API request.
func (r PutItemRequest) Send(ctx context.Context) (*PutItemResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutItemResponse{
		PutItemOutput: r.Request.Data.(*types.PutItemOutput),
		response:      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutItemResponse is the response type for the
// PutItem API operation.
type PutItemResponse struct {
	*types.PutItemOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutItem request.
func (r *PutItemResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
