// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type StartGameSessionPlacementInput struct {
	_ struct{} `type:"structure"`

	// Set of information on each player to create a player session for.
	DesiredPlayerSessions []DesiredPlayerSession `type:"list"`

	// Set of custom properties for a game session, formatted as key:value pairs.
	// These properties are passed to a game server process in the GameSession object
	// with a request to start a new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	GameProperties []GameProperty `type:"list"`

	// Set of custom game session properties, formatted as a single string value.
	// This data is passed to a game server process in the GameSession object with
	// a request to start a new game session (see Start a Game Session (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-server-api.html#gamelift-sdk-server-startsession)).
	GameSessionData *string `min:"1" type:"string"`

	// Descriptive label that is associated with a game session. Session names do
	// not need to be unique.
	GameSessionName *string `min:"1" type:"string"`

	// Name of the queue to use to place the new game session.
	//
	// GameSessionQueueName is a required field
	GameSessionQueueName *string `min:"1" type:"string" required:"true"`

	// Maximum number of players that can be connected simultaneously to the game
	// session.
	//
	// MaximumPlayerSessionCount is a required field
	MaximumPlayerSessionCount *int64 `type:"integer" required:"true"`

	// Unique identifier to assign to the new game session placement. This value
	// is developer-defined. The value must be unique across all regions and cannot
	// be reused unless you are resubmitting a canceled or timed-out placement request.
	//
	// PlacementId is a required field
	PlacementId *string `min:"1" type:"string" required:"true"`

	// Set of values, expressed in milliseconds, indicating the amount of latency
	// that a player experiences when connected to AWS regions. This information
	// is used to try to place the new game session where it can offer the best
	// possible gameplay experience for the players.
	PlayerLatencies []PlayerLatency `type:"list"`
}

// String returns the string representation
func (s StartGameSessionPlacementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartGameSessionPlacementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartGameSessionPlacementInput"}
	if s.GameSessionData != nil && len(*s.GameSessionData) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionData", 1))
	}
	if s.GameSessionName != nil && len(*s.GameSessionName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionName", 1))
	}

	if s.GameSessionQueueName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GameSessionQueueName"))
	}
	if s.GameSessionQueueName != nil && len(*s.GameSessionQueueName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GameSessionQueueName", 1))
	}

	if s.MaximumPlayerSessionCount == nil {
		invalidParams.Add(aws.NewErrParamRequired("MaximumPlayerSessionCount"))
	}

	if s.PlacementId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlacementId"))
	}
	if s.PlacementId != nil && len(*s.PlacementId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlacementId", 1))
	}
	if s.DesiredPlayerSessions != nil {
		for i, v := range s.DesiredPlayerSessions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "DesiredPlayerSessions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.GameProperties != nil {
		for i, v := range s.GameProperties {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GameProperties", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.PlayerLatencies != nil {
		for i, v := range s.PlayerLatencies {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "PlayerLatencies", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type StartGameSessionPlacementOutput struct {
	_ struct{} `type:"structure"`

	// Object that describes the newly created game session placement. This object
	// includes all the information provided in the request, as well as start/end
	// time stamps and placement status.
	GameSessionPlacement *GameSessionPlacement `type:"structure"`
}

// String returns the string representation
func (s StartGameSessionPlacementOutput) String() string {
	return awsutil.Prettify(s)
}