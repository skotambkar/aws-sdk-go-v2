// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/enums"
)

type PutBotInput struct {
	_ struct{} `type:"structure"`

	// When Amazon Lex can't understand the user's input in context, it tries to
	// elicit the information a few times. After that, Amazon Lex sends the message
	// defined in abortStatement to the user, and then aborts the conversation.
	// To set the number of retries, use the valueElicitationPrompt field for the
	// slot type.
	//
	// For example, in a pizza ordering bot, Amazon Lex might ask a user "What type
	// of crust would you like?" If the user's response is not one of the expected
	// responses (for example, "thin crust, "deep dish," etc.), Amazon Lex tries
	// to elicit a correct response a few more times.
	//
	// For example, in a pizza ordering application, OrderPizza might be one of
	// the intents. This intent might require the CrustType slot. You specify the
	// valueElicitationPrompt field when you create the CrustType slot.
	AbortStatement *Statement `locationName:"abortStatement" type:"structure"`

	// Identifies a specific revision of the $LATEST version.
	//
	// When you create a new bot, leave the checksum field blank. If you specify
	// a checksum you get a BadRequestException exception.
	//
	// When you want to update a bot, set the checksum field to the checksum of
	// the most recent revision of the $LATEST version. If you don't specify the
	// checksum field, or if the checksum does not match the $LATEST version, you
	// get a PreconditionFailedException exception.
	Checksum *string `locationName:"checksum" type:"string"`

	// For each Amazon Lex bot created with the Amazon Lex Model Building Service,
	// you must specify whether your use of Amazon Lex is related to a website,
	// program, or other application that is directed or targeted, in whole or in
	// part, to children under age 13 and subject to the Children's Online Privacy
	// Protection Act (COPPA) by specifying true or false in the childDirected field.
	// By specifying true in the childDirected field, you confirm that your use
	// of Amazon Lex is related to a website, program, or other application that
	// is directed or targeted, in whole or in part, to children under age 13 and
	// subject to COPPA. By specifying false in the childDirected field, you confirm
	// that your use of Amazon Lex is not related to a website, program, or other
	// application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to COPPA. You may not specify a default value for
	// the childDirected field that does not accurately reflect whether your use
	// of Amazon Lex is related to a website, program, or other application that
	// is directed or targeted, in whole or in part, to children under age 13 and
	// subject to COPPA.
	//
	// If your use of Amazon Lex relates to a website, program, or other application
	// that is directed in whole or in part, to children under age 13, you must
	// obtain any required verifiable parental consent under COPPA. For information
	// regarding the use of Amazon Lex in connection with websites, programs, or
	// other applications that are directed or targeted, in whole or in part, to
	// children under age 13, see the Amazon Lex FAQ. (https://aws.amazon.com/lex/faqs#data-security)
	//
	// ChildDirected is a required field
	ChildDirected *bool `locationName:"childDirected" type:"boolean" required:"true"`

	// When Amazon Lex doesn't understand the user's intent, it uses this message
	// to get clarification. To specify how many times Amazon Lex should repeate
	// the clarification prompt, use the maxAttempts field. If Amazon Lex still
	// doesn't understand, it sends the message in the abortStatement field.
	//
	// When you create a clarification prompt, make sure that it suggests the correct
	// response from the user. for example, for a bot that orders pizza and drinks,
	// you might create this clarification prompt: "What would you like to do? You
	// can say 'Order a pizza' or 'Order a drink.'"
	ClarificationPrompt *Prompt `locationName:"clarificationPrompt" type:"structure"`

	CreateVersion *bool `locationName:"createVersion" type:"boolean"`

	// A description of the bot.
	Description *string `locationName:"description" type:"string"`

	// The maximum time in seconds that Amazon Lex retains the data gathered in
	// a conversation.
	//
	// A user interaction session remains active for the amount of time specified.
	// If no conversation occurs during this time, the session expires and Amazon
	// Lex deletes any data provided before the timeout.
	//
	// For example, suppose that a user chooses the OrderPizza intent, but gets
	// sidetracked halfway through placing an order. If the user doesn't complete
	// the order within the specified time, Amazon Lex discards the slot information
	// that it gathered, and the user must start over.
	//
	// If you don't include the idleSessionTTLInSeconds element in a PutBot operation
	// request, Amazon Lex uses the default value. This is also true if the request
	// replaces an existing bot.
	//
	// The default is 300 seconds (5 minutes).
	IdleSessionTTLInSeconds *int64 `locationName:"idleSessionTTLInSeconds" min:"60" type:"integer"`

	// An array of Intent objects. Each intent represents a command that a user
	// can express. For example, a pizza ordering bot might support an OrderPizza
	// intent. For more information, see how-it-works.
	Intents []Intent `locationName:"intents" type:"list"`

	// Specifies the target locale for the bot. Any intent used in the bot must
	// be compatible with the locale of the bot.
	//
	// The default is en-US.
	//
	// Locale is a required field
	Locale enums.Locale `locationName:"locale" type:"string" required:"true" enum:"true"`

	// The name of the bot. The name is not case sensitive.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" min:"2" type:"string" required:"true"`

	// If you set the processBehavior element to BUILD, Amazon Lex builds the bot
	// so that it can be run. If you set the element to SAVE Amazon Lex saves the
	// bot, but doesn't build it.
	//
	// If you don't specify this value, the default value is BUILD.
	ProcessBehavior enums.ProcessBehavior `locationName:"processBehavior" type:"string" enum:"true"`

	// The Amazon Polly voice ID that you want Amazon Lex to use for voice interactions
	// with the user. The locale configured for the voice must match the locale
	// of the bot. For more information, see Available Voices (http://docs.aws.amazon.com/polly/latest/dg/voicelist.html)
	// in the Amazon Polly Developer Guide.
	VoiceId *string `locationName:"voiceId" type:"string"`
}

// String returns the string representation
func (s PutBotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutBotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutBotInput"}

	if s.ChildDirected == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChildDirected"))
	}
	if s.IdleSessionTTLInSeconds != nil && *s.IdleSessionTTLInSeconds < 60 {
		invalidParams.Add(aws.NewErrParamMinValue("IdleSessionTTLInSeconds", 60))
	}
	if len(s.Locale) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Locale"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 2))
	}
	if s.AbortStatement != nil {
		if err := s.AbortStatement.Validate(); err != nil {
			invalidParams.AddNested("AbortStatement", err.(aws.ErrInvalidParams))
		}
	}
	if s.ClarificationPrompt != nil {
		if err := s.ClarificationPrompt.Validate(); err != nil {
			invalidParams.AddNested("ClarificationPrompt", err.(aws.ErrInvalidParams))
		}
	}
	if s.Intents != nil {
		for i, v := range s.Intents {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Intents", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutBotOutput struct {
	_ struct{} `type:"structure"`

	// The message that Amazon Lex uses to abort a conversation. For more information,
	// see PutBot.
	AbortStatement *Statement `locationName:"abortStatement" type:"structure"`

	// Checksum of the bot that you created.
	Checksum *string `locationName:"checksum" type:"string"`

	// For each Amazon Lex bot created with the Amazon Lex Model Building Service,
	// you must specify whether your use of Amazon Lex is related to a website,
	// program, or other application that is directed or targeted, in whole or in
	// part, to children under age 13 and subject to the Children's Online Privacy
	// Protection Act (COPPA) by specifying true or false in the childDirected field.
	// By specifying true in the childDirected field, you confirm that your use
	// of Amazon Lex is related to a website, program, or other application that
	// is directed or targeted, in whole or in part, to children under age 13 and
	// subject to COPPA. By specifying false in the childDirected field, you confirm
	// that your use of Amazon Lex is not related to a website, program, or other
	// application that is directed or targeted, in whole or in part, to children
	// under age 13 and subject to COPPA. You may not specify a default value for
	// the childDirected field that does not accurately reflect whether your use
	// of Amazon Lex is related to a website, program, or other application that
	// is directed or targeted, in whole or in part, to children under age 13 and
	// subject to COPPA.
	//
	// If your use of Amazon Lex relates to a website, program, or other application
	// that is directed in whole or in part, to children under age 13, you must
	// obtain any required verifiable parental consent under COPPA. For information
	// regarding the use of Amazon Lex in connection with websites, programs, or
	// other applications that are directed or targeted, in whole or in part, to
	// children under age 13, see the Amazon Lex FAQ. (https://aws.amazon.com/lex/faqs#data-security)
	ChildDirected *bool `locationName:"childDirected" type:"boolean"`

	// The prompts that Amazon Lex uses when it doesn't understand the user's intent.
	// For more information, see PutBot.
	ClarificationPrompt *Prompt `locationName:"clarificationPrompt" type:"structure"`

	CreateVersion *bool `locationName:"createVersion" type:"boolean"`

	// The date that the bot was created.
	CreatedDate *time.Time `locationName:"createdDate" type:"timestamp"`

	// A description of the bot.
	Description *string `locationName:"description" type:"string"`

	// If status is FAILED, Amazon Lex provides the reason that it failed to build
	// the bot.
	FailureReason *string `locationName:"failureReason" type:"string"`

	// The maximum length of time that Amazon Lex retains the data gathered in a
	// conversation. For more information, see PutBot.
	IdleSessionTTLInSeconds *int64 `locationName:"idleSessionTTLInSeconds" min:"60" type:"integer"`

	// An array of Intent objects. For more information, see PutBot.
	Intents []Intent `locationName:"intents" type:"list"`

	// The date that the bot was updated. When you create a resource, the creation
	// date and last updated date are the same.
	LastUpdatedDate *time.Time `locationName:"lastUpdatedDate" type:"timestamp"`

	// The target locale for the bot.
	Locale enums.Locale `locationName:"locale" type:"string" enum:"true"`

	// The name of the bot.
	Name *string `locationName:"name" min:"2" type:"string"`

	// When you send a request to create a bot with processBehavior set to BUILD,
	// Amazon Lex sets the status response element to BUILDING. After Amazon Lex
	// builds the bot, it sets status to READY. If Amazon Lex can't build the bot,
	// Amazon Lex sets status to FAILED. Amazon Lex returns the reason for the failure
	// in the failureReason response element.
	//
	// When you set processBehaviorto SAVE, Amazon Lex sets the status code to NOT
	// BUILT.
	Status enums.Status `locationName:"status" type:"string" enum:"true"`

	// The version of the bot. For a new bot, the version is always $LATEST.
	Version *string `locationName:"version" min:"1" type:"string"`

	// The Amazon Polly voice ID that Amazon Lex uses for voice interaction with
	// the user. For more information, see PutBot.
	VoiceId *string `locationName:"voiceId" type:"string"`
}

// String returns the string representation
func (s PutBotOutput) String() string {
	return awsutil.Prettify(s)
}