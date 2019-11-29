// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_restjson service, marshal.go contains the generated operation marshalers and methods on it.

package aws_restjson

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/private/protocol/rest"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice/types"
)

var _ bytes.Buffer
var _ awserr.Error

// CreateConfigurationSetMarshaler defines marshaler for CreateConfigurationSet operation
type CreateConfigurationSetMarshaler struct {
	Input *types.CreateConfigurationSetInput
}

// CreateConfigurationSetEventDestinationMarshaler defines marshaler for CreateConfigurationSetEventDestination operation
type CreateConfigurationSetEventDestinationMarshaler struct {
	Input *types.CreateConfigurationSetEventDestinationInput
}

// DeleteConfigurationSetMarshaler defines marshaler for DeleteConfigurationSet operation
type DeleteConfigurationSetMarshaler struct {
	Input *types.DeleteConfigurationSetInput
}

// DeleteConfigurationSetEventDestinationMarshaler defines marshaler for DeleteConfigurationSetEventDestination operation
type DeleteConfigurationSetEventDestinationMarshaler struct {
	Input *types.DeleteConfigurationSetEventDestinationInput
}

// GetConfigurationSetEventDestinationsMarshaler defines marshaler for GetConfigurationSetEventDestinations operation
type GetConfigurationSetEventDestinationsMarshaler struct {
	Input *types.GetConfigurationSetEventDestinationsInput
}

// ListConfigurationSetsMarshaler defines marshaler for ListConfigurationSets operation
type ListConfigurationSetsMarshaler struct {
	Input *types.ListConfigurationSetsInput
}

// SendVoiceMessageMarshaler defines marshaler for SendVoiceMessage operation
type SendVoiceMessageMarshaler struct {
	Input *types.SendVoiceMessageInput
}

// UpdateConfigurationSetEventDestinationMarshaler defines marshaler for UpdateConfigurationSetEventDestination operation
type UpdateConfigurationSetEventDestinationMarshaler struct {
	Input *types.UpdateConfigurationSetEventDestinationInput
}

func MarshalCreateConfigurationSetInputShapeAWSREST(v *types.CreateConfigurationSetInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalCreateConfigurationSetInputShapeAWSJSON(v *types.CreateConfigurationSetInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalCreateConfigurationSetEventDestinationInputShapeAWSREST(v *types.CreateConfigurationSetEventDestinationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalCreateConfigurationSetEventDestinationInputShapeAWSJSON(v *types.CreateConfigurationSetEventDestinationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDeleteConfigurationSetInputShapeAWSREST(v *types.DeleteConfigurationSetInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDeleteConfigurationSetInputShapeAWSJSON(v *types.DeleteConfigurationSetInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDeleteConfigurationSetEventDestinationInputShapeAWSREST(v *types.DeleteConfigurationSetEventDestinationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDeleteConfigurationSetEventDestinationInputShapeAWSJSON(v *types.DeleteConfigurationSetEventDestinationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetConfigurationSetEventDestinationsInputShapeAWSREST(v *types.GetConfigurationSetEventDestinationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetConfigurationSetEventDestinationsInputShapeAWSJSON(v *types.GetConfigurationSetEventDestinationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListConfigurationSetsInputShapeAWSREST(v *types.ListConfigurationSetsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListConfigurationSetsInputShapeAWSJSON(v *types.ListConfigurationSetsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalSendVoiceMessageInputShapeAWSREST(v *types.SendVoiceMessageInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalSendVoiceMessageInputShapeAWSJSON(v *types.SendVoiceMessageInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUpdateConfigurationSetEventDestinationInputShapeAWSREST(v *types.UpdateConfigurationSetEventDestinationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateConfigurationSetEventDestinationInputShapeAWSJSON(v *types.UpdateConfigurationSetEventDestinationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}
func (m CreateConfigurationSetMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateConfigurationSetInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalCreateConfigurationSetInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateConfigurationSetMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateConfigurationSet.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateConfigurationSetEventDestinationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateConfigurationSetEventDestinationInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalCreateConfigurationSetEventDestinationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateConfigurationSetEventDestinationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateConfigurationSetEventDestination.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteConfigurationSetMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteConfigurationSetInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDeleteConfigurationSetInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteConfigurationSetMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteConfigurationSet.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteConfigurationSetEventDestinationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteConfigurationSetEventDestinationInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDeleteConfigurationSetEventDestinationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteConfigurationSetEventDestinationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteConfigurationSetEventDestination.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetConfigurationSetEventDestinationsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetConfigurationSetEventDestinationsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetConfigurationSetEventDestinationsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetConfigurationSetEventDestinationsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetConfigurationSetEventDestinations.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListConfigurationSetsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListConfigurationSetsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListConfigurationSetsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListConfigurationSetsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListConfigurationSets.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SendVoiceMessageMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSendVoiceMessageInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalSendVoiceMessageInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SendVoiceMessageMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SendVoiceMessage.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateConfigurationSetEventDestinationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateConfigurationSetEventDestinationInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateConfigurationSetEventDestinationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateConfigurationSetEventDestinationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateConfigurationSetEventDestination.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
