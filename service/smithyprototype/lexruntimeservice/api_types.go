// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexruntimeservice

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/json"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/internal/sdkio"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

var _ aws.Config
var _ = awsutil.Prettify

// Represents an option to be shown on the client platform (Facebook, Slack,
// etc.)
type Button struct {
	_ struct{} `type:"structure"`

	// Text that is visible to the user on the button.
	//
	// Text is a required field
	Text *string `locationName:"text" min:"1" type:"string" required:"true"`

	// The value sent to Amazon Lex when a user chooses the button. For example,
	// consider button text "NYC." When the user chooses the button, the value sent
	// can be "New York City."
	//
	// Value is a required field
	Value *string `locationName:"value" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s Button) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s Button) MarshalFields(e protocol.FieldEncoder) error {
	if s.Text != nil {
		v := *s.Text

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "text", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Value != nil {
		v := *s.Value

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "value", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Describes the next action that the bot should take in its interaction with
// the user and provides information about the context in which the action takes
// place. Use the DialogAction data type to set the interaction to a specific
// state, or to return the interaction to a previous state.
type DialogAction struct {
	_ struct{} `type:"structure"`

	// The fulfillment state of the intent. The possible values are:
	//
	//    * Failed - The Lambda function associated with the intent failed to fulfill
	//    the intent.
	//
	//    * Fulfilled - The intent has fulfilled by the Lambda function associated
	//    with the intent.
	//
	//    * ReadyForFulfillment - All of the information necessary for the intent
	//    is present and the intent ready to be fulfilled by the client application.
	FulfillmentState FulfillmentState `locationName:"fulfillmentState" type:"string" enum:"true"`

	// The name of the intent.
	IntentName *string `locationName:"intentName" type:"string"`

	// The message that should be shown to the user. If you don't specify a message,
	// Amazon Lex will use the message configured for the intent.
	Message *string `locationName:"message" min:"1" type:"string" sensitive:"true"`

	//    * PlainText - The message contains plain UTF-8 text.
	//
	//    * CustomPayload - The message is a custom format for the client.
	//
	//    * SSML - The message contains text formatted for voice output.
	//
	//    * Composite - The message contains an escaped JSON object containing one
	//    or more messages. For more information, see Message Groups (https://docs.aws.amazon.com/lex/latest/dg/howitworks-manage-prompts.html).
	MessageFormat MessageFormatType `locationName:"messageFormat" type:"string" enum:"true"`

	// The name of the slot that should be elicited from the user.
	SlotToElicit *string `locationName:"slotToElicit" type:"string"`

	// Map of the slots that have been gathered and their values.
	Slots map[string]string `locationName:"slots" type:"map" sensitive:"true"`

	// The next action that the bot should take in its interaction with the user.
	// The possible values are:
	//
	//    * ConfirmIntent - The next action is asking the user if the intent is
	//    complete and ready to be fulfilled. This is a yes/no question such as
	//    "Place the order?"
	//
	//    * Close - Indicates that the there will not be a response from the user.
	//    For example, the statement "Your order has been placed" does not require
	//    a response.
	//
	//    * Delegate - The next action is determined by Amazon Lex.
	//
	//    * ElicitIntent - The next action is to determine the intent that the user
	//    wants to fulfill.
	//
	//    * ElicitSlot - The next action is to elicit a slot value from the user.
	//
	// Type is a required field
	Type DialogActionType `locationName:"type" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s DialogAction) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DialogAction) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DialogAction"}
	if s.Message != nil && len(*s.Message) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Message", 1))
	}
	if len(s.Type) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Type"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DialogAction) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.FulfillmentState) > 0 {
		v := s.FulfillmentState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fulfillmentState", protocol.QuotedValue{ValueMarshaler: v}, metadata)
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
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

// Represents an option rendered to the user when a prompt is shown. It could
// be an image, a button, a link, or text.
type GenericAttachment struct {
	_ struct{} `type:"structure"`

	// The URL of an attachment to the response card.
	AttachmentLinkUrl *string `locationName:"attachmentLinkUrl" min:"1" type:"string"`

	// The list of options to show to the user.
	Buttons []Button `locationName:"buttons" type:"list"`

	// The URL of an image that is displayed to the user.
	ImageUrl *string `locationName:"imageUrl" min:"1" type:"string"`

	// The subtitle shown below the title.
	SubTitle *string `locationName:"subTitle" min:"1" type:"string"`

	// The title of the option.
	Title *string `locationName:"title" min:"1" type:"string"`
}

// String returns the string representation
func (s GenericAttachment) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GenericAttachment) MarshalFields(e protocol.FieldEncoder) error {
	if s.AttachmentLinkUrl != nil {
		v := *s.AttachmentLinkUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "attachmentLinkUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Buttons != nil {
		v := s.Buttons

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "buttons", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.ImageUrl != nil {
		v := *s.ImageUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "imageUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SubTitle != nil {
		v := *s.SubTitle

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "subTitle", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Title != nil {
		v := *s.Title

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "title", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Provides information about the state of an intent. You can use this information
// to get the current state of an intent so that you can process the intent,
// or so that you can return the intent to its previous state.
type IntentSummary struct {
	_ struct{} `type:"structure"`

	// A user-defined label that identifies a particular intent. You can use this
	// label to return to a previous intent.
	//
	// Use the checkpointLabelFilter parameter of the GetSessionRequest operation
	// to filter the intents returned by the operation to those with only the specified
	// label.
	CheckpointLabel *string `locationName:"checkpointLabel" min:"1" type:"string"`

	// The status of the intent after the user responds to the confirmation prompt.
	// If the user confirms the intent, Amazon Lex sets this field to Confirmed.
	// If the user denies the intent, Amazon Lex sets this value to Denied. The
	// possible values are:
	//
	//    * Confirmed - The user has responded "Yes" to the confirmation prompt,
	//    confirming that the intent is complete and that it is ready to be fulfilled.
	//
	//    * Denied - The user has responded "No" to the confirmation prompt.
	//
	//    * None - The user has never been prompted for confirmation; or, the user
	//    was prompted but did not confirm or deny the prompt.
	ConfirmationStatus ConfirmationStatus `locationName:"confirmationStatus" type:"string" enum:"true"`

	// The next action that the bot should take in its interaction with the user.
	// The possible values are:
	//
	//    * ConfirmIntent - The next action is asking the user if the intent is
	//    complete and ready to be fulfilled. This is a yes/no question such as
	//    "Place the order?"
	//
	//    * Close - Indicates that the there will not be a response from the user.
	//    For example, the statement "Your order has been placed" does not require
	//    a response.
	//
	//    * ElicitIntent - The next action is to determine the intent that the user
	//    wants to fulfill.
	//
	//    * ElicitSlot - The next action is to elicit a slot value from the user.
	//
	// DialogActionType is a required field
	DialogActionType DialogActionType `locationName:"dialogActionType" type:"string" required:"true" enum:"true"`

	// The fulfillment state of the intent. The possible values are:
	//
	//    * Failed - The Lambda function associated with the intent failed to fulfill
	//    the intent.
	//
	//    * Fulfilled - The intent has fulfilled by the Lambda function associated
	//    with the intent.
	//
	//    * ReadyForFulfillment - All of the information necessary for the intent
	//    is present and the intent ready to be fulfilled by the client application.
	FulfillmentState FulfillmentState `locationName:"fulfillmentState" type:"string" enum:"true"`

	// The name of the intent.
	IntentName *string `locationName:"intentName" type:"string"`

	// The next slot to elicit from the user. If there is not slot to elicit, the
	// field is blank.
	SlotToElicit *string `locationName:"slotToElicit" type:"string"`

	// Map of the slots that have been gathered and their values.
	Slots map[string]string `locationName:"slots" type:"map" sensitive:"true"`
}

// String returns the string representation
func (s IntentSummary) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *IntentSummary) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "IntentSummary"}
	if s.CheckpointLabel != nil && len(*s.CheckpointLabel) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CheckpointLabel", 1))
	}
	if len(s.DialogActionType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("DialogActionType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s IntentSummary) MarshalFields(e protocol.FieldEncoder) error {
	if s.CheckpointLabel != nil {
		v := *s.CheckpointLabel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "checkpointLabel", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.ConfirmationStatus) > 0 {
		v := s.ConfirmationStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "confirmationStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.DialogActionType) > 0 {
		v := s.DialogActionType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dialogActionType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.FulfillmentState) > 0 {
		v := s.FulfillmentState

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "fulfillmentState", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.IntentName != nil {
		v := *s.IntentName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "intentName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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

// If you configure a response card when creating your bots, Amazon Lex substitutes
// the session attributes and slot values that are available, and then returns
// it. The response card can also come from a Lambda function ( dialogCodeHook
// and fulfillmentActivity on an intent).
type ResponseCard struct {
	_ struct{} `type:"structure"`

	// The content type of the response.
	ContentType ContentType `locationName:"contentType" type:"string" enum:"true"`

	// An array of attachment objects representing options.
	GenericAttachments []GenericAttachment `locationName:"genericAttachments" type:"list"`

	// The version of the response card format.
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s ResponseCard) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ResponseCard) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.ContentType) > 0 {
		v := s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "contentType", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.GenericAttachments != nil {
		v := s.GenericAttachments

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "genericAttachments", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The sentiment expressed in an utterance.
//
// When the bot is configured to send utterances to Amazon Comprehend for sentiment
// analysis, this field structure contains the result of the analysis.
type SentimentResponse struct {
	_ struct{} `type:"structure"`

	// The inferred sentiment that Amazon Comprehend has the highest confidence
	// in.
	SentimentLabel *string `locationName:"sentimentLabel" type:"string"`

	// The likelihood that the sentiment was correctly inferred.
	SentimentScore *string `locationName:"sentimentScore" type:"string"`
}

// String returns the string representation
func (s SentimentResponse) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s SentimentResponse) MarshalFields(e protocol.FieldEncoder) error {
	if s.SentimentLabel != nil {
		v := *s.SentimentLabel

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sentimentLabel", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SentimentScore != nil {
		v := *s.SentimentScore

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sentimentScore", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

func serializeIntentSummaryListAWSJSON(v []IntentSummary, value json.Value) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	array := value.Array()

	for _, arrayValue := range v {
		av := array.Value()
		if err := serializeIntentSummaryAWSJSON(&arrayValue, av); err != nil {
			return err
		}
	}

	return nil
}

func serializeIntentSummaryAWSJSON(v *IntentSummary, value json.Value) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	object := value.Object()
	defer object.Close()

	if v.CheckpointLabel != nil {
		object.Key("checkpointLabel").String(*v.CheckpointLabel)
	}

	if len(v.ConfirmationStatus) > 0 {
		object.Key("confirmationStatus").String(string(v.ConfirmationStatus))
	}

	if len(v.DialogActionType) > 0 {
		object.Key("dialogActionType").String(string(v.DialogActionType))
	}

	if len(v.FulfillmentState) > 0 {
		object.Key("fulfillmentState").String(string(v.FulfillmentState))
	}

	if v.IntentName != nil {
		object.Key("intentName").String(*v.IntentName)
	}

	if v.SlotToElicit != nil {
		object.Key("slotToElicit").String(*v.SlotToElicit)
	}

	if v.Slots != nil {
		value := object.Key("slots")
		if err := serializeStringMapAWSJSON(v.Slots, value); err != nil {
			return err
		}
	}

	return nil
}

func serializeDialogActionAWSJSON(v *DialogAction, value json.Value) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	object := value.Object()
	defer object.Close()

	if len(v.FulfillmentState) > 0 {
		object.Key("fulfillmentState").String(string(v.FulfillmentState))
	}

	if v.IntentName != nil {
		object.Key("intentName").String(*v.IntentName)
	}

	if v.Message != nil {
		object.Key("message").String(*v.Message)
	}

	if len(v.MessageFormat) > 0 {
		object.Key("messageFormat").String(string(v.MessageFormat))
	}

	if v.SlotToElicit != nil {
		object.Key("slotToElicit").String(*v.SlotToElicit)
	}

	if v.Slots != nil {
		value := object.Key("slots")
		if err := serializeStringMapAWSJSON(v.Slots, value); err != nil {
			return err
		}
	}

	if len(v.Type) > 0 {
		object.Key("type").String(string(v.Type))
	}

	return nil
}

func serializeStringMapAWSJSON(v map[string]string, value json.Value) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	object := value.Object()
	defer object.Close()

	for k, kv := range v {
		object.Key(k).String(kv)
	}

	return nil
}

// awsjson_responseCardShapeDeserialize
func awsjson_responseCardShapeDeserialize(output *ResponseCard, decoder *json.Decoder, rb sdkio.RingBuffer) error {
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

		// location name : `contentType` key with value as enum
		if t == "contentType" {
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
					output.ContentType = v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected ContentType to be of type Enum, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
			}
		}

		// Todo: Add case for []GenericAttachment

		// location name : `version` key with value as `*string`
		if t == "version" {
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
					output.Version = &v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected Version to be of type *String, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
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

// awsjson_sentimentResponseShapeDeserialize
func awsjson_sentimentResponseShapeDeserialize(output *SentimentResponse, decoder *json.Decoder, rb sdkio.RingBuffer) error {
	startToken, err := decoder.Token()
	if err == io.EOF {
		// "Empty List"
		return nil
	}
	if err != nil {
		snapshot := make([]byte, 1024)
		rb.Read(snapshot)
		return aws.SerializationError{
			Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
				"Here's a snapshot: %s",
				snapshot, err),
		}
	}
	if t, ok := startToken.(json.Delim); !ok || t.String() != "{" {
		snapshot := make([]byte, 1024)
		rb.Read(snapshot)
		return aws.SerializationError{
			Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
				"expected `{` as start token; "+
				"Here's a snapshot: %s",
				snapshot),
		}
	}

	// decoder
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

		// location name : `sentimentScore` key with value as `*string`
		if t == "sentimentScore" {
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
					output.SentimentScore = &v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected Version to be of type *String, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
			}
		}

		// location name : `sentimentLabel` key with value as `*string`
		if t == "sentimentLabel" {
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
					output.SentimentLabel = &v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected Version to be of type *String, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
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
}

// awsjson_sessionAttributeMapShapeDeserialize
func awsjson_sessionAttributeMapShapeDeserialize(output *map[string]string, decoder *json.Decoder, rb *sdkio.RingBuffer) error {
	startToken, err := decoder.Token()
	if err == io.EOF {
		// "Empty List"
		return nil
	}
	if err != nil {
		snapshot := make([]byte, 1024)
		rb.Read(snapshot)
		return aws.SerializationError{
			Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
				"Here's a snapshot: %s",
				snapshot, err),
		}
	}
	if t, ok := startToken.(json.Delim); !ok || t.String() != "{" {
		snapshot := make([]byte, 1024)
		rb.Read(snapshot)
		return aws.SerializationError{
			Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
				"expected `{` as start token; "+
				"Here's a snapshot: %s",
				snapshot),
		}
	}

	// decoder
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			snapshot := make([]byte, 1024)
			rb.Read(snapshot)
			return aws.SerializationError{
				Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
					"Here's a snapshot: %s",
					snapshot, err),
			}
		}

		// based on struct Stage
		if key, ok := token.(string); ok {
			val, err := decoder.Token()
			if err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
						"Here's a snapshot: %s",
						snapshot, err),
				}
			}
			if v, ok := val.(string); ok {
				m := *output
				m[key] = v
			} else {
				return awserr.New(aws.ErrCodeSerialization,
					fmt.Sprintf("invalid json, expected string. Here's a snapshot: %s",
						getSnapshot(ringBuffer)), err)
			}
		} else {
			return awserr.New(aws.ErrCodeSerialization,
				fmt.Sprintf("Expected %v, to be of type string, got %T. "+
					"Here's a snapshot: %s", key, key,
					getSnapshot(ringBuffer)), err)
		}
	}

	// end of the map
	endToken, err := decoder.Token()
	if err != nil {
		return awserr.New(aws.ErrCodeSerialization,
			fmt.Sprintf("failed to decode response body with invalid JSON. Here's a snapshot: %s",
				getSnapshot(ringBuffer)), err)
	}
	if t, ok := endToken.(json.Delim); !ok || t.String() != "}" {
		return awserr.New(aws.ErrCodeSerialization,
			fmt.Sprintf("failed to decode response body with invalid JSON, expected `}` as end Token. "+
				"Here's a snapshot: %s",
				getSnapshot(ringBuffer)), err)
	}
	return nil
}

// awsjson_slotsMapShapeDeserialize
func awsjson_slotsMapShapeDeserialize(output *map[string]string, decoder *json.Decoder, rb *sdkio.RingBuffer) error {
	startToken, err := decoder.Token()
	if err == io.EOF {
		// "Empty List"
		return nil
	}
	if err != nil {
		snapshot := make([]byte, 1024)
		rb.Read(snapshot)
		return aws.SerializationError{
			Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
				"Here's a snapshot: %s",
				snapshot, err),
		}
	}
	if t, ok := startToken.(json.Delim); !ok || t.String() != "{" {
		snapshot := make([]byte, 1024)
		rb.Read(snapshot)
		return aws.SerializationError{
			Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
				"expected `{` as start token; "+
				"Here's a snapshot: %s",
				snapshot),
		}
	}

	// decoder
	for decoder.More() {
		token, err := decoder.Token()
		if err != nil {
			snapshot := make([]byte, 1024)
			rb.Read(snapshot)
			return aws.SerializationError{
				Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
					"Here's a snapshot: %s",
					snapshot, err),
			}
		}

		// based on struct Stage
		if key, ok := token.(string); ok {
			val, err := decoder.Token()
			if err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
						"Here's a snapshot: %s",
						snapshot, err),
				}
			}
			if v, ok := val.(string); ok {
				m := *output
				m[key] = v
			} else {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
						"Here's a snapshot: %s",
						snapshot, err),
				}
			}
		} else {
			snapshot := make([]byte, 1024)
			rb.Read(snapshot)
			return aws.SerializationError{
				Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
					"Here's a snapshot: %s",
					snapshot, err),
			}

			// end of the map
			endToken, err := decoder.Token()
			if err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
						"Here's a snapshot: %s",
						snapshot, err),
				}
				if t, ok := endToken.(json.Delim); !ok || t.String() != "}" {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
							"expected `}` as end token; "+
							"Here's a snapshot: %s",
							snapshot),
					}
				}
				return nil
			}
		}
	}
}

// awsjson_dialogActionShapeDeserialize
func awsjson_dialogActionShapeDeserialize(output *DialogAction, decoder *json.Decoder, rb *sdkio.RingBuffer) error {
	startToken, err := decoder.Token()
	if err == io.EOF {
		// "Empty List"
		return nil
	}
	if err != nil {
		snapshot := make([]byte, 1024)
		rb.Read(snapshot)
		return aws.SerializationError{
			Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
				"Here's a snapshot: %s",
				snapshot, err),
		}
	}
	if t, ok := startToken.(json.Delim); !ok || t.String() != "{" {
		snapshot := make([]byte, 1024)
		rb.Read(snapshot)
		return aws.SerializationError{
			Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
				"expected `{` as start token; "+
				"Here's a snapshot: %s",
				snapshot),
		}
	}

	for decoder.More() {

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

		if t == "type" {
			val, err := decoder.Token()
			if err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON.", err),
				}
				if v, ok := val.(string); ok {
					output.Type = v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected Type to be of type *String, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
			}
		}

		if t == "fulfillmentState" {
			val, err := decoder.Token()
			if err != nil {
				snapshot := make([]byte, 1024)
				rb.Read(snapshot)
				return aws.SerializationError{
					Err: fmt.Sprintf("failed to decode response body with invalid JSON.", err),
				}
				if v, ok := val.(string); ok {
					output.FulfillmentState = v
				} else {
					snapshot := make([]byte, 1024)
					rb.Read(snapshot)
					return aws.SerializationError{
						Err: fmt.Sprintf("expected Type to be of type *String, "+
							"Here's a snapshot: %s",
							snapshot, err),
					}
				}
			}
		}

	}

	// end of the map
	endToken, err := decoder.Token()
	if err != nil {
		snapshot := make([]byte, 1024)
		rb.Read(snapshot)
		return aws.SerializationError{
			Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
				"Here's a snapshot: %s",
				snapshot, err),
		}
	}
	if t, ok := endToken.(json.Delim); !ok || t.String() != "}" {
		snapshot := make([]byte, 1024)
		rb.Read(snapshot)
		return aws.SerializationError{
			Err: fmt.Sprintf("failed to decode response body with invalid JSON,"+
				"expected `}` as end token; "+
				"Here's a snapshot: %s",
				snapshot),
		}
	}
	return nil
}
