// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexruntimeservice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/rest"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/internal/sdkio"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/awslabs/smithy-go/middleware"
	smithyHTTP "github.com/awslabs/smithy-go/transport/http"
)

type PostTextInput struct {
	_ struct{} `type:"structure"`

	// The alias of the Amazon Lex bot.
	//
	// BotAlias is a required field
	BotAlias *string `location:"uri" locationName:"botAlias" type:"string" required:"true"`

	// The name of the Amazon Lex bot.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" type:"string" required:"true"`

	// The text that the user entered (Amazon Lex interprets this text).
	//
	// InputText is a required field
	InputText *string `locationName:"inputText" min:"1" type:"string" required:"true" sensitive:"true"`

	// Request-specific information passed between Amazon Lex and a client application.
	//
	// The namespace x-amz-lex: is reserved for special attributes. Don't create
	// any request attributes with the prefix x-amz-lex:.
	//
	// For more information, see Setting Request Attributes (https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-request-attribs).
	RequestAttributes map[string]string `locationName:"requestAttributes" type:"map" sensitive:"true"`

	// Application-specific information passed between Amazon Lex and a client application.
	//
	// For more information, see Setting Session Attributes (https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html#context-mgmt-session-attribs).
	SessionAttributes map[string]string `locationName:"sessionAttributes" type:"map" sensitive:"true"`

	// The ID of the client application user. Amazon Lex uses this to identify a
	// user's conversation with your bot. At runtime, each request must contain
	// the userID field.
	//
	// To decide the user ID to use for your application, consider the following
	// factors.
	//
	//    * The userID field must not contain any personally identifiable information
	//    of the user, for example, name, personal identification numbers, or other
	//    end user personal information.
	//
	//    * If you want a user to start a conversation on one device and continue
	//    on another device, use a user-specific identifier.
	//
	//    * If you want the same user to be able to have two independent conversations
	//    on two different devices, choose a device-specific identifier.
	//
	//    * A user can't have two independent conversations with two different versions
	//    of the same bot. For example, a user can't have a conversation with the
	//    PROD and BETA versions of the same bot. If you anticipate that a user
	//    will need to have conversation with two different versions, for example,
	//    while testing, include the bot alias in the user ID to separate the two
	//    conversations.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PostTextInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostTextInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PostTextInput"}

	if s.BotAlias == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotAlias"))
	}

	if s.BotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BotName"))
	}

	if s.InputText == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputText"))
	}
	if s.InputText != nil && len(*s.InputText) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("InputText", 1))
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
func (s PostTextInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.InputText != nil {
		v := *s.InputText

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "inputText", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestAttributes != nil {
		v := s.RequestAttributes

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "requestAttributes", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

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
	return nil
}

type PostTextOutput struct {
	_ struct{} `type:"structure"`

	// Identifies the current state of the user interaction. Amazon Lex returns
	// one of the following values as dialogState. The client can optionally use
	// this information to customize the user interface.
	//
	//    * ElicitIntent - Amazon Lex wants to elicit user intent. For example,
	//    a user might utter an intent ("I want to order a pizza"). If Amazon Lex
	//    cannot infer the user intent from this utterance, it will return this
	//    dialogState.
	//
	//    * ConfirmIntent - Amazon Lex is expecting a "yes" or "no" response. For
	//    example, Amazon Lex wants user confirmation before fulfilling an intent.
	//    Instead of a simple "yes" or "no," a user might respond with additional
	//    information. For example, "yes, but make it thick crust pizza" or "no,
	//    I want to order a drink". Amazon Lex can process such additional information
	//    (in these examples, update the crust type slot value, or change intent
	//    from OrderPizza to OrderDrink).
	//
	//    * ElicitSlot - Amazon Lex is expecting a slot value for the current intent.
	//    For example, suppose that in the response Amazon Lex sends this message:
	//    "What size pizza would you like?". A user might reply with the slot value
	//    (e.g., "medium"). The user might also provide additional information in
	//    the response (e.g., "medium thick crust pizza"). Amazon Lex can process
	//    such additional information appropriately.
	//
	//    * Fulfilled - Conveys that the Lambda function configured for the intent
	//    has successfully fulfilled the intent.
	//
	//    * ReadyForFulfillment - Conveys that the client has to fulfill the intent.
	//
	//    * Failed - Conveys that the conversation with the user failed. This can
	//    happen for various reasons including that the user did not provide an
	//    appropriate response to prompts from the service (you can configure how
	//    many times Amazon Lex can prompt a user for specific information), or
	//    the Lambda function failed to fulfill the intent.
	DialogState DialogState `locationName:"dialogState" type:"string" enum:"true"`

	// The current user intent that Amazon Lex is aware of.
	IntentName *string `locationName:"intentName" type:"string"`

	// The message to convey to the user. The message can come from the bot's configuration
	// or from a Lambda function.
	//
	// If the intent is not configured with a Lambda function, or if the Lambda
	// function returned Delegate as the dialogAction.type its response, Amazon
	// Lex decides on the next course of action and selects an appropriate message
	// from the bot's configuration based on the current interaction context. For
	// example, if Amazon Lex isn't able to understand user input, it uses a clarification
	// prompt message.
	//
	// When you create an intent you can assign messages to groups. When messages
	// are assigned to groups Amazon Lex returns one message from each group in
	// the response. The message field is an escaped JSON string containing the
	// messages. For more information about the structure of the JSON string returned,
	// see msg-prompts-formats.
	//
	// If the Lambda function returns a message, Amazon Lex passes it to the client
	// in its response.
	Message *string `locationName:"message" min:"1" type:"string" sensitive:"true"`

	// The format of the response message. One of the following values:
	//
	//    * PlainText - The message contains plain UTF-8 text.
	//
	//    * CustomPayload - The message is a custom format defined by the Lambda
	//    function.
	//
	//    * SSML - The message contains text formatted for voice output.
	//
	//    * Composite - The message contains an escaped JSON object containing one
	//    or more messages from the groups that messages were assigned to when the
	//    intent was created.
	MessageFormat MessageFormatType `locationName:"messageFormat" type:"string" enum:"true"`

	// Represents the options that the user has to respond to the current prompt.
	// Response Card can come from the bot configuration (in the Amazon Lex console,
	// choose the settings button next to a slot) or from a code hook (Lambda function).
	ResponseCard *ResponseCard `locationName:"responseCard" type:"structure"`

	// The sentiment expressed in and utterance.
	//
	// When the bot is configured to send utterances to Amazon Comprehend for sentiment
	// analysis, this field contains the result of the analysis.
	SentimentResponse *SentimentResponse `locationName:"sentimentResponse" type:"structure"`

	// A map of key-value pairs representing the session-specific context information.
	SessionAttributes map[string]string `locationName:"sessionAttributes" type:"map" sensitive:"true"`

	// A unique identifier for the session.
	SessionId *string `locationName:"sessionId" type:"string"`

	// If the dialogState value is ElicitSlot, returns the name of the slot for
	// which Amazon Lex is eliciting a value.
	SlotToElicit *string `locationName:"slotToElicit" type:"string"`

	// The intent slots that Amazon Lex detected from the user input in the conversation.
	//
	// Amazon Lex creates a resolution list containing likely values for a slot.
	// The value that it returns is determined by the valueSelectionStrategy selected
	// when the slot type was created or updated. If valueSelectionStrategy is set
	// to ORIGINAL_VALUE, the value provided by the user is returned, if the user
	// value is similar to the slot values. If valueSelectionStrategy is set to
	// TOP_RESOLUTION Amazon Lex returns the first value in the resolution list
	// or, if there is no resolution list, null. If you don't specify a valueSelectionStrategy,
	// the default is ORIGINAL_VALUE.
	Slots map[string]string `locationName:"slots" type:"map" sensitive:"true"`
}

// String returns the string representation
func (s PostTextOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PostTextOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.DialogState) > 0 {
		v := s.DialogState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dialogState", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.IntentName != nil {
		v := *s.IntentName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "intentName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Message != nil {
		v := *s.Message

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "message", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.MessageFormat) > 0 {
		v := s.MessageFormat

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "messageFormat", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.ResponseCard != nil {
		v := s.ResponseCard

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "responseCard", v, metadata)
	}
	if s.SentimentResponse != nil {
		v := s.SentimentResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "sentimentResponse", v, metadata)
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
	if s.SlotToElicit != nil {
		v := *s.SlotToElicit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "slotToElicit", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Slots != nil {
		v := s.Slots

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "slots", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	return nil
}

const opPostText = "PostText"

// PostTextRequest returns a request value for making API operation for
// Amazon Lex Runtime Service.
//
// Sends user input to Amazon Lex. Client applications can use this API to send
// requests to Amazon Lex at runtime. Amazon Lex then interprets the user input
// using the machine learning model it built for the bot.
//
// In response, Amazon Lex returns the next message to convey to the user an
// optional responseCard to display. Consider the following example messages:
//
//    * For a user input "I would like a pizza", Amazon Lex might return a response
//    with a message eliciting slot data (for example, PizzaSize): "What size
//    pizza would you like?"
//
//    * After the user provides all of the pizza order information, Amazon Lex
//    might return a response with a message to obtain user confirmation "Proceed
//    with the pizza order?".
//
//    * After the user replies to a confirmation prompt with a "yes", Amazon
//    Lex might return a conclusion statement: "Thank you, your cheese pizza
//    has been ordered.".
//
// Not all Amazon Lex messages require a user response. For example, a conclusion
// statement does not require a response. Some messages require only a "yes"
// or "no" user response. In addition to the message, Amazon Lex provides additional
// context about the message in the response that you might use to enhance client
// behavior, for example, to display the appropriate client user interface.
// These are the slotToElicit, dialogState, intentName, and slots fields in
// the response. Consider the following examples:
//
//    * If the message is to elicit slot data, Amazon Lex returns the following
//    context information: dialogState set to ElicitSlot intentName set to the
//    intent name in the current context slotToElicit set to the slot name for
//    which the message is eliciting information slots set to a map of slots,
//    configured for the intent, with currently known values
//
//    * If the message is a confirmation prompt, the dialogState is set to ConfirmIntent
//    and SlotToElicit is set to null.
//
//    * If the message is a clarification prompt (configured for the intent)
//    that indicates that user intent is not understood, the dialogState is
//    set to ElicitIntent and slotToElicit is set to null.
//
// In addition, Amazon Lex also returns your application-specific sessionAttributes.
// For more information, see Managing Conversation Context (https://docs.aws.amazon.com/lex/latest/dg/context-mgmt.html).
//
//    // Example sending a request using PostTextRequest.
//    req := client.PostTextRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostText
func (c *Client) PostTextRequest(input *PostTextInput) PostTextRequest {
	op := &aws.Operation{
		Name:       opPostText,
		HTTPMethod: "POST",
		HTTPPath:   "/bot/{botName}/alias/{botAlias}/user/{userId}/text",
	}

	if input == nil {
		input = &PostTextInput{}
	}

	req := c.newRequest(op, input, &PostTextOutput{})
	return PostTextRequest{Request: req, Input: input, Copy: c.PostTextRequest}
}

// PostTextRequest is the request type for the
// PostText API operation.
type PostTextRequest struct {
	*aws.Request
	Input *PostTextInput
	Copy  func(*PostTextInput) PostTextRequest
}

// Send marshals and sends the PostText API request.
func (r PostTextRequest) Send(ctx context.Context) (*PostTextResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PostTextResponse{
		PostTextOutput: r.Request.Data.(*PostTextOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PostTextResponse is the response type for the
// PostText API operation.
type PostTextResponse struct {
	*PostTextOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PostText request.
func (r *PostTextResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}

type postTextSerializeMiddleware struct{}

func (p postTextSerializeMiddleware) ID() string {
	return "PostTextSerializeMiddleware"
}

func (p postTextSerializeMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyHTTP.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*PostTextInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}

	restEncoder := rest.NewEncoder(request.Request)

	restEncoder.AddHeader("Content-Type").String("application/json")

	if err := serializePostTextInputAWSREST(input, restEncoder); err != nil {
		return out, metadata, err
	}

	jsonEncoder := json.NewEncoder()
	if err := serializePostTextInputAWSJSON(input, jsonEncoder.Value); err != nil {
		return out, metadata, err
	}

	request.Stream = bytes.NewReader(jsonEncoder.Bytes())

	return next.HandleSerialize(ctx, in)
}

func serializePostTextInputAWSJSON(v *PostTextInput, value json.Value) error {
	object := value.Object()
	defer object.Close()

	if v == nil {
		return nil
	}

	if v.InputText != nil {
		object.Key("inputText").String(*v.InputText)
	}

	if v.RequestAttributes != nil {
		value := object.Key("requestAttributes")
		if err := serializeStringMapAWSJSON(v.RequestAttributes, value); err != nil {
			return err
		}
	}

	if v.SessionAttributes != nil {
		value := object.Key("sessionAttributes")
		if err := serializeStringMapAWSJSON(v.SessionAttributes, value); err != nil {
			return err
		}
	}

	return nil
}

func serializePostTextInputAWSREST(v *PostTextInput, encoder *rest.Encoder) error {
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

	if v.UserId != nil {
		if err := encoder.SetURI("userId").String(*v.UserId); err != nil {
			return err
		}
	}

	return nil
}

// awsrestjson_postTextDeserializeMiddleware
type awsrestjson_postTextDeserializeMiddleware struct{}

// ID is the middleware identifier
func (p awsrestjson_postTextDeserializeMiddleware) ID() string {
	return "awsrestjson_postTextDeserializeMiddleware"
}

// HandleSerialize
func (p awsrestjson_postTextDeserializeMiddleware) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	response, ok := out.RawResponse.(*smithyHTTP.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", out.RawResponse)
	}

	output, ok := out.Result.(*PostTextOutput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown output parameters type %T", out.Result)
	}

	buff := make([]byte, 1024)
	ringBuffer := sdkio.NewRingBuffer(buff)

	// wrap a TeeReader to read from response body & write on snapshot
	body := io.TeeReader(response.Body, ringBuffer)
	defer r.HTTPResponse.Body.Close()
	decoder := json.NewDecoder(body)

	if err := awsjson_postTextOutputDeserialize(output, decoder, ringBuffer); err != nil {
		return out, metadata, err
	}

	return out, metadata, err
}

// awsjson_postTextOutputDeserialize
func awsjson_postTextOutputDeserialize(output *PostTextOutput, decoder *json.Decoder, rb *sdkio.RingBuffer) error {
	if output == nil {
		return nil
	}

	startToken, err := decoder.Token()
	if err == io.EOF {
		// "Empty Response"
		return nil
	}
	if err != nil {
		return err
	}

	if t, ok := startToken.(json.Delim); !ok {
		if t.String() != "{" {
			snapshot := make([]byte, 1024)
			rb.Read(snapshot)
			return aws.SerializationError{
				Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
					"expected `{` as start token; "+
					"Here's a snapshot: %s",
					snapshot),
			}
		}
	}

	for decoder.More() {
		// fetch token for key
		t, err := decoder.Token()
		if err != nil {
			snapshot := make([]byte, 1024)
			rb.Read(snapshot)
			return aws.SerializationError{
				Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
					"Here's a snapshot: %s",
					snapshot, err),
			}
		}

		// location name : `dialogState` key with value as `string`
		if t == "dialogState" {
			val, err := decoder.Token()
			if err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON. "+
						"Here's a snapshot: %s",
						snapshot, err),
				}
				if v, ok := val.(string); ok {
					output.DialogState = v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected DialogState to be of type String, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
			}
		}

		// location name: `intentName` key with value as `string`
		if t == "intentName" {
			val, err := decoder.Token()
			if err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON. "+
						"Here's a snapshot: %s",
						snapshot, err),
				}
				if v, ok := val.(string); ok {
					output.IntentName = &v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected IntentName to be of type *String, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
			}
		}

		// Todo : handle sensitive fields
		if t == "message" {
			val, err := decoder.Token()
			if err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON.", err),
				}
				if v, ok := val.(string); ok {
					output.Message = &v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected Message to be of type *String, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
			}
		}

		if t == "messageFormat" {
			val, err := decoder.Token()
			if err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON.", err),
				}
				if v, ok := val.(string); ok {
					output.MessageFormat = v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected MessageFormat to be of type *String, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
			}
		}

		if t == "responseCard" {
			v := &ResponseCard{}
			if err = awsjson_responseCardShapeDeserialize(v, decoder, rb); err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("expected ResponseCard to be of type *ResponseCard, "+
						"Here's a snapshot: %s",
						snapshot, err),
				}
			} else {
				output.ResponseCard = v
			}

		}

		if t == "sentimentResponse" {
			v := &SentimentResponse{}
			if err = awsjson_sentimentResponseShapeDeserialize(v, decoder, rb); err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("expected SentimentResponse to be of type *SentimentResponse, "+
						"Here's a snapshot: %s",
						snapshot, err),
				}
			} else {
				output.SentimentResponse = v
			}
		}

		if t == "sessionAttributes" {
			// create []string as modeled for the member shape
			v := make(map[string]string, 0)
			if err = awsjson_sessionAttributeMapShapeDeserialize(v, decoder, rb); err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("expected SessionAttributes to be of type map[string]string, "+
						"Here's a snapshot: %s",
						snapshot, err),
				}
			} else {
				output.SessionAttributes = v
			}
		}

		if t == "sessionId" {
			val, err := decoder.Token()
			if err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON.", err),
				}
				if v, ok := val.(string); ok {
					output.SessionId = &v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected SessionId to be of type *String, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
			}
		}

		if t == "slotToElicit" {
			val, err := decoder.Token()
			if err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON.", err),
				}
				if v, ok := val.(string); ok {
					output.SlotToElicit = &v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected SlotToElicit to be of type *String, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
			}
		}

		if t == "slots" {
			// create []string as modeled for the member shape
			v := make(map[string]string, 0)
			if err = awsjson_slotsMapShapeDeserialize(v, decoder, rb); err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("expected SessionAttributes to be of type map[string]string, "+
						"Here's a snapshot: %s",
						snapshot, err),
				}
			} else {
				output.Slots = v
			}
		}

	}

	// end of the json response body
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t.String() != "}" {
		snapshot := make([]byte, 1024)
		rb.Read(snapshot)
		return aws.SerializationError{
			Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
				"expected `}` as end token; "+
				"Here's a snapshot: %s",
				snapshot, err),
		}
	}

	return nil
}
