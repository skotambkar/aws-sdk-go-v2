// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateVariableInput struct {
	_ struct{} `type:"structure"`

	// The new default value of the variable.
	DefaultValue *string `locationName:"defaultValue" type:"string"`

	// The new description.
	Description *string `locationName:"description" type:"string"`

	// The name of the variable.
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// The variable type.
	VariableType *string `locationName:"variableType" type:"string"`
}

// String returns the string representation
func (s UpdateVariableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVariableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateVariableInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateVariableOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateVariableOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateVariable = "UpdateVariable"

// UpdateVariableRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Updates a variable.
//
//    // Example sending a request using UpdateVariableRequest.
//    req := client.UpdateVariableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/UpdateVariable
func (c *Client) UpdateVariableRequest(input *UpdateVariableInput) UpdateVariableRequest {
	op := &aws.Operation{
		Name:       opUpdateVariable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateVariableInput{}
	}

	req := c.newRequest(op, input, &UpdateVariableOutput{})
	return UpdateVariableRequest{Request: req, Input: input, Copy: c.UpdateVariableRequest}
}

// UpdateVariableRequest is the request type for the
// UpdateVariable API operation.
type UpdateVariableRequest struct {
	*aws.Request
	Input *UpdateVariableInput
	Copy  func(*UpdateVariableInput) UpdateVariableRequest
}

// Send marshals and sends the UpdateVariable API request.
func (r UpdateVariableRequest) Send(ctx context.Context) (*UpdateVariableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateVariableResponse{
		UpdateVariableOutput: r.Request.Data.(*UpdateVariableOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateVariableResponse is the response type for the
// UpdateVariable API operation.
type UpdateVariableResponse struct {
	*UpdateVariableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateVariable request.
func (r *UpdateVariableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
