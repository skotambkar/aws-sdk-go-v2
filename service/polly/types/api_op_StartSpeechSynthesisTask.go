// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/polly/enums"
)

type StartSpeechSynthesisTaskInput struct {
	_ struct{} `type:"structure"`

	// Specifies the engine (standard or neural) for Amazon Polly to use when processing
	// input text for speech synthesis. Using a voice that is not supported for
	// the engine selected will result in an error.
	Engine enums.Engine `type:"string" enum:"true"`

	// Optional language code for the Speech Synthesis request. This is only necessary
	// if using a bilingual voice, such as Aditi, which can be used for either Indian
	// English (en-IN) or Hindi (hi-IN).
	//
	// If a bilingual voice is used and no language code is specified, Amazon Polly
	// will use the default language of the bilingual voice. The default language
	// for any voice is the one returned by the DescribeVoices (https://docs.aws.amazon.com/polly/latest/dg/API_DescribeVoices.html)
	// operation for the LanguageCode parameter. For example, if no language code
	// is specified, Aditi will use Indian English rather than Hindi.
	LanguageCode enums.LanguageCode `type:"string" enum:"true"`

	// List of one or more pronunciation lexicon names you want the service to apply
	// during synthesis. Lexicons are applied only if the language of the lexicon
	// is the same as the language of the voice.
	LexiconNames []string `type:"list"`

	// The format in which the returned output will be encoded. For audio stream,
	// this will be mp3, ogg_vorbis, or pcm. For speech marks, this will be json.
	//
	// OutputFormat is a required field
	OutputFormat enums.OutputFormat `type:"string" required:"true" enum:"true"`

	// Amazon S3 bucket name to which the output file will be saved.
	//
	// OutputS3BucketName is a required field
	OutputS3BucketName *string `type:"string" required:"true"`

	// The Amazon S3 key prefix for the output speech file.
	OutputS3KeyPrefix *string `type:"string"`

	// The audio frequency specified in Hz.
	//
	// The valid values for mp3 and ogg_vorbis are "8000", "16000", "22050", and
	// "24000". The default value for standard voices is "22050". The default value
	// for neural voices is "24000".
	//
	// Valid values for pcm are "8000" and "16000" The default value is "16000".
	SampleRate *string `type:"string"`

	// ARN for the SNS topic optionally used for providing status notification for
	// a speech synthesis task.
	SnsTopicArn *string `type:"string"`

	// The type of speech marks returned for the input text.
	SpeechMarkTypes []enums.SpeechMarkType `type:"list"`

	// The input text to synthesize. If you specify ssml as the TextType, follow
	// the SSML format for the input text.
	//
	// Text is a required field
	Text *string `type:"string" required:"true"`

	// Specifies whether the input text is plain text or SSML. The default value
	// is plain text.
	TextType enums.TextType `type:"string" enum:"true"`

	// Voice ID to use for the synthesis.
	//
	// VoiceId is a required field
	VoiceId enums.VoiceId `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s StartSpeechSynthesisTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSpeechSynthesisTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartSpeechSynthesisTaskInput"}
	if len(s.OutputFormat) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("OutputFormat"))
	}

	if s.OutputS3BucketName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputS3BucketName"))
	}

	if s.Text == nil {
		invalidParams.Add(aws.NewErrParamRequired("Text"))
	}
	if len(s.VoiceId) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("VoiceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartSpeechSynthesisTaskOutput struct {
	_ struct{} `type:"structure"`

	// SynthesisTask object that provides information and attributes about a newly
	// submitted speech synthesis task.
	SynthesisTask *SynthesisTask `type:"structure"`
}

// String returns the string representation
func (s StartSpeechSynthesisTaskOutput) String() string {
	return awsutil.Prettify(s)
}
