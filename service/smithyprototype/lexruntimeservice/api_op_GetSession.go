// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexruntimeservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/rest"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/awslabs/smithy-go/middleware"
	smithyHTTP "github.com/awslabs/smithy-go/transport/http"
)

type GetSessionInput struct {
	_ struct{} `type:"structure"`

	// The alias in use for the bot that contains the session data.
	//
	// BotAlias is a required field
	BotAlias *string `location:"uri" locationName:"botAlias" type:"string" required:"true"`

	// The name of the bot that contains the session data.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" type:"string" required:"true"`

	// A string used to filter the intents returned in the recentIntentSummaryView
	// structure.
	//
	// When you specify a filter, only intents with their checkpointLabel field
	// set to that string are returned.
	CheckpointLabelFilter *string `location:"querystring" locationName:"checkpointLabelFilter" min:"1" type:"string"`

	// The ID of the client application user. Amazon Lex uses this to identify a
	// user's conversation with your bot.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s GetSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetSessionInput"}

	if s.BotAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotAlias"))
	}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}
	if s.CheckpointLabelFilter != nil && len(*s.CheckpointLabelFilter) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CheckpointLabelFilter", 1))
	}

	if s.UserId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("UserId", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSessionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BotAlias != nil {
		v := *s.BotAlias

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botAlias", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.BotName != nil {
		v := *s.BotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "botName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.UserId != nil {
		v := *s.UserId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "userId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CheckpointLabelFilter != nil {
		v := *s.CheckpointLabelFilter

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "checkpointLabelFilter", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetSessionOutput struct {
	_ struct{} `type:"structure"`

	// Describes the current state of the bot.
	DialogAction *DialogAction `locationName:"dialogAction" type:"structure"`

	// An array of information about the intents used in the session. The array
	// can contain a maximum of three summaries. If more than three intents are
	// used in the session, the recentIntentSummaryView operation contains information
	// about the last three intents used.
	//
	// If you set the checkpointLabelFilter parameter in the request, the array
	// contains only the intents with the specified label.
	RecentIntentSummaryView []IntentSummary `locationName:"recentIntentSummaryView" type:"list"`

	// Map of key/value pairs representing the session-specific context information.
	// It contains application information passed between Amazon Lex and a client
	// application.
	SessionAttributes map[string]string `locationName:"sessionAttributes" type:"map" sensitive:"true"`

	// A unique identifier for the session.
	SessionId *string `locationName:"sessionId" type:"string"`
}

// String returns the string representation
func (s GetSessionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetSessionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DialogAction != nil {
		v := s.DialogAction

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "dialogAction", v, metadata)
	}
	if s.RecentIntentSummaryView != nil {
		v := s.RecentIntentSummaryView

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "recentIntentSummaryView", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.SessionAttributes != nil {
		v := s.SessionAttributes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "sessionAttributes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.SessionId != nil {
		v := *s.SessionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sessionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetSession = "GetSession"

// GetSessionRequest returns a request value for making API operation for
// Amazon Lex Runtime Service.
//
// Returns session information for a specified bot, alias, and user ID.
//
//    // Example sending a request using GetSessionRequest.
//    req := client.GetSessionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/GetSession
func (c *Client) GetSessionRequest(input *GetSessionInput) GetSessionRequest {
	op := &aws.Operation{
		Name:       opGetSession,
		HTTPMethod: "GET",
		HTTPPath:   "/bot/{botName}/alias/{botAlias}/user/{userId}/session/",
	}

	if input == nil {
		input = &GetSessionInput{}
	}

	req := c.newRequest(op, input, &GetSessionOutput{})
	return GetSessionRequest{Request: req, Input: input, Copy: c.GetSessionRequest}
}

// GetSessionRequest is the request type for the
// GetSession API operation.
type GetSessionRequest struct {
	*aws.Request
	Input *GetSessionInput
	Copy  func(*GetSessionInput) GetSessionRequest
}

// Send marshals and sends the GetSession API request.
func (r GetSessionRequest) Send(ctx context.Context) (*GetSessionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSessionResponse{
		GetSessionOutput: r.Request.Data.(*GetSessionOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSessionResponse is the response type for the
// GetSession API operation.
type GetSessionResponse struct {
	*GetSessionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSession request.
func (r *GetSessionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

type getSessionSerializeMiddleware struct{}

// ID is the identifier for the middleware
func (g getSessionSerializeMiddleware) ID() string {
	return "GetSessionSerializeMiddleware"
}

// HandleSerialize will serialize the middleware input parameters to the provided input http request
func (g getSessionSerializeMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyHTTP.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*GetSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}

	restEncoder := rest.NewEncoder(request.Request)

	restEncoder.AddHeader("Content-Type").String("application/json")

	if err := serializeGetSessionInputAWSREST(input, restEncoder); err != nil {
		return middleware.SerializeOutput{}, metadata, err
	}

	return next.HandleSerialize(ctx, in)
}

// serializeDeleteSessionInputAWSREST marshals the top level members of GetSessionInput that have HTTP bindings
func serializeGetSessionInputAWSREST(v *GetSessionInput, encoder *rest.Encoder) error {
	if v == nil {
		return nil
	}

	if v.BotAlias != nil {
		if err := encoder.SetURI("botAlias").String(*v.BotAlias); err != nil {
			return err
		}
	}

	if v.BotName != nil {
		if err := encoder.SetURI("botName").String(*v.BotName); err != nil {
			return err
		}
	}

	if v.CheckpointLabelFilter != nil {
		encoder.SetQuery("checkpointLabelFilter").String(*v.CheckpointLabelFilter)
	}

	if v.UserId != nil {
		if err := encoder.SetURI("userId").String(*v.UserId); err != nil {
			return err
		}
	}

	return nil
}
