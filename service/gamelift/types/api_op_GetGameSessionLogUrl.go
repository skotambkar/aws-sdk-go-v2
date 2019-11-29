// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type GetGameSessionLogUrlInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for the game session to get logs for.
	//
	// GameSessionId is a required field
	GameSessionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetGameSessionLogUrlInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetGameSessionLogUrlInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetGameSessionLogUrlInput"}

	if s.GameSessionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameSessionId"))
	}
	if s.GameSessionId != nil && len(*s.GameSessionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type GetGameSessionLogUrlOutput struct {
	_ struct{} `type:"structure"`

	// Location of the requested game session logs, available for download. This
	// URL is valid for 15 minutes, after which S3 will reject any download request
	// using this URL. You can request a new URL any time within the 14-day period
	// that the logs are retained.
	PreSignedUrl *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetGameSessionLogUrlOutput) String() string {
	return awsutil.Prettify(s)
}
