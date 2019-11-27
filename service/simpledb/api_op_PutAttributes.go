// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package simpledb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/simpledb/types"
)

const opPutAttributes = "PutAttributes"

// PutAttributesRequest returns a request value for making API operation for
// Amazon SimpleDB.
//
// The PutAttributes operation creates or replaces attributes in an item. The
// client may specify new attributes using a combination of the Attribute.X.Name
// and Attribute.X.Value parameters. The client specifies the first attribute
// by the parameters Attribute.0.Name and Attribute.0.Value, the second attribute
// by the parameters Attribute.1.Name and Attribute.1.Value, and so on.
//
// Attributes are uniquely identified in an item by their name/value combination.
// For example, a single item can have the attributes { "first_name", "first_value"
// } and { "first_name", second_value" }. However, it cannot have two attribute
// instances where both the Attribute.X.Name and Attribute.X.Value are the same.
//
// Optionally, the requestor can supply the Replace parameter for each individual
// attribute. Setting this value to true causes the new attribute value to replace
// the existing attribute value(s). For example, if an item has the attributes
// { 'a', '1' }, { 'b', '2'} and { 'b', '3' } and the requestor calls PutAttributes
// using the attributes { 'b', '4' } with the Replace parameter set to true,
// the final attributes of the item are changed to { 'a', '1' } and { 'b', '4'
// }, which replaces the previous values of the 'b' attribute with the new value.
//   Using
//    PutAttributes
//  to replace attribute values that do not exist will not result in an error
//  response.
// You cannot specify an empty string as an attribute name.
//
// Because Amazon SimpleDB makes multiple copies of client data and uses an
// eventual consistency update model, an immediate GetAttributes or Select operation
// (read) immediately after a PutAttributes or DeleteAttributes operation (write)
// might not return the updated data.
//
// The following limitations are enforced for this operation:
//    * 256 total attribute name-value pairs per item
//
//    * One billion attributes per domain
//
//    * 10 GB of total user data storage per domain
//
//    // Example sending a request using PutAttributesRequest.
//    req := client.PutAttributesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) PutAttributesRequest(input *types.PutAttributesInput) PutAttributesRequest {
	op := &aws.Operation{
		Name:       opPutAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.PutAttributesInput{}
	}

	req := c.newRequest(op, input, &types.PutAttributesOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return PutAttributesRequest{Request: req, Input: input, Copy: c.PutAttributesRequest}
}

// PutAttributesRequest is the request type for the
// PutAttributes API operation.
type PutAttributesRequest struct {
	*aws.Request
	Input *types.PutAttributesInput
	Copy  func(*types.PutAttributesInput) PutAttributesRequest
}

// Send marshals and sends the PutAttributes API request.
func (r PutAttributesRequest) Send(ctx context.Context) (*PutAttributesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutAttributesResponse{
		PutAttributesOutput: r.Request.Data.(*types.PutAttributesOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutAttributesResponse is the response type for the
// PutAttributes API operation.
type PutAttributesResponse struct {
	*types.PutAttributesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutAttributes request.
func (r *PutAttributesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
