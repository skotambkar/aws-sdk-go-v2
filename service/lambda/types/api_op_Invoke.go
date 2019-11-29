// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/lambda/enums"
)

type InvokeInput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// Up to 3583 bytes of base64-encoded data about the invoking client to pass
	// to the function in the context object.
	ClientContext *string `location:"header" locationName:"X-Amz-Client-Context" type:"string"`

	// The name of the Lambda function, version, or alias.
	//
	// Name formats
	//
	//    * Function name - my-function (name-only), my-function:v1 (with alias).
	//
	//    * Function ARN - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//    * Partial ARN - 123456789012:function:my-function.
	//
	// You can append a version number or alias to any of the formats. The length
	// constraint applies only to the full ARN. If you specify only the function
	// name, it is limited to 64 characters in length.
	//
	// FunctionName is a required field
	FunctionName *string `location:"uri" locationName:"FunctionName" min:"1" type:"string" required:"true"`

	// Choose from the following options.
	//
	//    * RequestResponse (default) - Invoke the function synchronously. Keep
	//    the connection open until the function returns a response or times out.
	//    The API response includes the function response and additional data.
	//
	//    * Event - Invoke the function asynchronously. Send events that fail multiple
	//    times to the function's dead-letter queue (if it's configured). The API
	//    response only includes a status code.
	//
	//    * DryRun - Validate parameter values and verify that the user or role
	//    has permission to invoke the function.
	InvocationType enums.InvocationType `location:"header" locationName:"X-Amz-Invocation-Type" type:"string" enum:"true"`

	// Set to Tail to include the execution log in the response.
	LogType enums.LogType `location:"header" locationName:"X-Amz-Log-Type" type:"string" enum:"true"`

	// The JSON that you want to provide to your Lambda function as input.
	Payload []byte `type:"blob" sensitive:"true"`

	// Specify a version or alias to invoke a published version of the function.
	Qualifier *string `location:"querystring" locationName:"Qualifier" min:"1" type:"string"`
}

// String returns the string representation
func (s InvokeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InvokeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InvokeInput"}

	if s.FunctionName == nil {
		invalidParams.Add(aws.NewErrParamRequired("FunctionName"))
	}
	if s.FunctionName != nil && len(*s.FunctionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FunctionName", 1))
	}
	if s.Qualifier != nil && len(*s.Qualifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Qualifier", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type InvokeOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The version of the function that executed. When you invoke a function with
	// an alias, this indicates which version the alias resolved to.
	ExecutedVersion *string `location:"header" locationName:"X-Amz-Executed-Version" min:"1" type:"string"`

	// If present, indicates that an error occurred during function execution. Details
	// about the error are included in the response payload.
	//
	//    * Handled - The runtime caught an error thrown by the function and formatted
	//    it into a JSON document.
	//
	//    * Unhandled - The runtime didn't handle the error. For example, the function
	//    ran out of memory or timed out.
	FunctionError *string `location:"header" locationName:"X-Amz-Function-Error" type:"string"`

	// The last 4 KB of the execution log, which is base64 encoded.
	LogResult *string `location:"header" locationName:"X-Amz-Log-Result" type:"string"`

	// The response from the function, or an error object.
	Payload []byte `type:"blob" sensitive:"true"`

	// The HTTP status code is in the 200 range for a successful request. For the
	// RequestResponse invocation type, this status code is 200. For the Event invocation
	// type, this status code is 202. For the DryRun invocation type, the status
	// code is 204.
	StatusCode *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s InvokeOutput) String() string {
	return awsutil.Prettify(s)
}