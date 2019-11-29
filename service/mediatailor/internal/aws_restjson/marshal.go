// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_restjson service, marshal.go contains the generated operation marshalers and methods on it.

package aws_restjson

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/private/protocol/rest"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
)

var _ bytes.Buffer
var _ awserr.Error

// DeletePlaybackConfigurationMarshaler defines marshaler for DeletePlaybackConfiguration operation
type DeletePlaybackConfigurationMarshaler struct {
	Input *types.DeletePlaybackConfigurationInput
}

// GetPlaybackConfigurationMarshaler defines marshaler for GetPlaybackConfiguration operation
type GetPlaybackConfigurationMarshaler struct {
	Input *types.GetPlaybackConfigurationInput
}

// ListPlaybackConfigurationsMarshaler defines marshaler for ListPlaybackConfigurations operation
type ListPlaybackConfigurationsMarshaler struct {
	Input *types.ListPlaybackConfigurationsInput
}

// ListTagsForResourceMarshaler defines marshaler for ListTagsForResource operation
type ListTagsForResourceMarshaler struct {
	Input *types.ListTagsForResourceInput
}

// PutPlaybackConfigurationMarshaler defines marshaler for PutPlaybackConfiguration operation
type PutPlaybackConfigurationMarshaler struct {
	Input *types.PutPlaybackConfigurationInput
}

// TagResourceMarshaler defines marshaler for TagResource operation
type TagResourceMarshaler struct {
	Input *types.TagResourceInput
}

// UntagResourceMarshaler defines marshaler for UntagResource operation
type UntagResourceMarshaler struct {
	Input *types.UntagResourceInput
}

func MarshalDeletePlaybackConfigurationInputShapeAWSREST(v *types.DeletePlaybackConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDeletePlaybackConfigurationInputShapeAWSJSON(v *types.DeletePlaybackConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetPlaybackConfigurationInputShapeAWSREST(v *types.GetPlaybackConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetPlaybackConfigurationInputShapeAWSJSON(v *types.GetPlaybackConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListPlaybackConfigurationsInputShapeAWSREST(v *types.ListPlaybackConfigurationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListPlaybackConfigurationsInputShapeAWSJSON(v *types.ListPlaybackConfigurationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListTagsForResourceInputShapeAWSREST(v *types.ListTagsForResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListTagsForResourceInputShapeAWSJSON(v *types.ListTagsForResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalPutPlaybackConfigurationInputShapeAWSREST(v *types.PutPlaybackConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalPutPlaybackConfigurationInputShapeAWSJSON(v *types.PutPlaybackConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalTagResourceInputShapeAWSREST(v *types.TagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalTagResourceInputShapeAWSJSON(v *types.TagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUntagResourceInputShapeAWSREST(v *types.UntagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUntagResourceInputShapeAWSJSON(v *types.UntagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}
func (m DeletePlaybackConfigurationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeletePlaybackConfigurationInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDeletePlaybackConfigurationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeletePlaybackConfigurationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeletePlaybackConfiguration.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetPlaybackConfigurationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetPlaybackConfigurationInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetPlaybackConfigurationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetPlaybackConfigurationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetPlaybackConfiguration.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListPlaybackConfigurationsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListPlaybackConfigurationsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListPlaybackConfigurationsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListPlaybackConfigurationsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListPlaybackConfigurations.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListTagsForResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListTagsForResourceInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListTagsForResourceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListTagsForResourceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListTagsForResource.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m PutPlaybackConfigurationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalPutPlaybackConfigurationInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalPutPlaybackConfigurationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m PutPlaybackConfigurationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "PutPlaybackConfiguration.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m TagResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalTagResourceInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalTagResourceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m TagResourceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "TagResource.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UntagResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUntagResourceInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUntagResourceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UntagResourceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UntagResource.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
