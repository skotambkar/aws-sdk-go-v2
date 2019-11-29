// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/enums"
)

// Represents the input for a request action.
type UpdateGameSessionInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for the game session to update.
	//
	// GameSessionId is a required field
	GameSessionId *string `min:"1" type:"string" required:"true"`

	// Maximum number of players that can be connected simultaneously to the game
	// session.
	MaximumPlayerSessionCount *int64 `type:"integer"`

	// Descriptive label that is associated with a game session. Session names do
	// not need to be unique.
	Name *string `min:"1" type:"string"`

	// Policy determining whether or not the game session accepts new players.
	PlayerSessionCreationPolicy enums.PlayerSessionCreationPolicy `type:"string" enum:"true"`

	// Game session protection policy to apply to this game session only.
	//
	//    * NoProtection -- The game session can be terminated during a scale-down
	//    event.
	//
	//    * FullProtection -- If the game session is in an ACTIVE status, it cannot
	//    be terminated during a scale-down event.
	ProtectionPolicy enums.ProtectionPolicy `type:"string" enum:"true"`
}

// String returns the string representation
func (s UpdateGameSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateGameSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateGameSessionInput"}

	if s.GameSessionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameSessionId"))
	}
	if s.GameSessionId != nil && len(*s.GameSessionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionId", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type UpdateGameSessionOutput struct {
	_ struct{} `type:"structure"`

	// Object that contains the updated game session metadata.
	GameSession *GameSession `type:"structure"`
}

// String returns the string representation
func (s UpdateGameSessionOutput) String() string {
	return awsutil.Prettify(s)
}