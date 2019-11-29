// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The request representing the confirmation for a password reset.
type ConfirmForgotPasswordInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Pinpoint analytics metadata for collecting metrics for ConfirmForgotPassword
	// calls.
	AnalyticsMetadata *AnalyticsMetadataType `type:"structure"`

	// The app client ID of the app associated with the user pool.
	//
	// ClientId is a required field
	ClientId *string `min:"1" type:"string" required:"true" sensitive:"true"`

	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers.
	//
	// You create custom workflows by assigning AWS Lambda functions to user pool
	// triggers. When you use the ConfirmForgotPassword API action, Amazon Cognito
	// invokes the function that is assigned to the post confirmation trigger. When
	// Amazon Cognito invokes this function, it passes a JSON payload, which the
	// function receives as input. This payload contains a clientMetadata attribute,
	// which provides the data that you assigned to the ClientMetadata parameter
	// in your ConfirmForgotPassword request. In your function code in AWS Lambda,
	// you can process the clientMetadata value to enhance your workflow for your
	// specific needs.
	//
	// For more information, see Customizing User Pool Workflows with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide.
	//
	// Take the following limitations into consideration when you use the ClientMetadata
	// parameter:
	//
	//    * Amazon Cognito does not store the ClientMetadata value. This data is
	//    available only to AWS Lambda triggers that are assigned to a user pool
	//    to support custom workflows. If your user pool configuration does not
	//    include triggers, the ClientMetadata parameter serves no purpose.
	//
	//    * Amazon Cognito does not validate the ClientMetadata value.
	//
	//    * Amazon Cognito does not encrypt the the ClientMetadata value, so don't
	//    use it to provide sensitive information.
	ClientMetadata map[string]string `type:"map"`

	// The confirmation code sent by a user's request to retrieve a forgotten password.
	// For more information, see
	//
	// ConfirmationCode is a required field
	ConfirmationCode *string `min:"1" type:"string" required:"true"`

	// The password sent by a user's request to retrieve a forgotten password.
	//
	// Password is a required field
	Password *string `min:"6" type:"string" required:"true" sensitive:"true"`

	// A keyed-hash message authentication code (HMAC) calculated using the secret
	// key of a user pool client and username plus the client ID in the message.
	SecretHash *string `min:"1" type:"string" sensitive:"true"`

	// Contextual data such as the user's device fingerprint, IP address, or location
	// used for evaluating the risk of an unexpected event by Amazon Cognito advanced
	// security.
	UserContextData *UserContextDataType `type:"structure"`

	// The user name of the user for whom you want to enter a code to retrieve a
	// forgotten password.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true" sensitive:"true"`
}

// String returns the string representation
func (s ConfirmForgotPasswordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfirmForgotPasswordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConfirmForgotPasswordInput"}

	if s.ClientId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientId"))
	}
	if s.ClientId != nil && len(*s.ClientId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientId", 1))
	}

	if s.ConfirmationCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConfirmationCode"))
	}
	if s.ConfirmationCode != nil && len(*s.ConfirmationCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ConfirmationCode", 1))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}
	if s.Password != nil && len(*s.Password) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("Password", 6))
	}
	if s.SecretHash != nil && len(*s.SecretHash) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecretHash", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The response from the server that results from a user's request to retrieve
// a forgotten password.
type ConfirmForgotPasswordOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ConfirmForgotPasswordOutput) String() string {
	return awsutil.Prettify(s)
}