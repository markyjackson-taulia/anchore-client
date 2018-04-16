// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddPolicy adds a new policy

Adds a new policy bundle to the system
*/
func (a *Client) AddPolicy(params *AddPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*AddPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "add_policy",
		Method:             "POST",
		PathPattern:        "/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddPolicyOK), nil

}

/*
DeletePolicy deletes policy

Delete the specified policy
*/
func (a *Client) DeletePolicy(params *DeletePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_policy",
		Method:             "DELETE",
		PathPattern:        "/policies/{policyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePolicyOK), nil

}

/*
GetPolicy gets specific policy

Get the policy bundle content
*/
func (a *Client) GetPolicy(params *GetPolicyParams, authInfo runtime.ClientAuthInfoWriter) (*GetPolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_policy",
		Method:             "GET",
		PathPattern:        "/policies/{policyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPolicyOK), nil

}

/*
ListPolicies lists policies

List all saved policy bundles
*/
func (a *Client) ListPolicies(params *ListPoliciesParams, authInfo runtime.ClientAuthInfoWriter) (*ListPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "list_policies",
		Method:             "GET",
		PathPattern:        "/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListPoliciesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListPoliciesOK), nil

}

/*
UpdatePolicy updates policy

Update/replace and existing policy
*/
func (a *Client) UpdatePolicy(params *UpdatePolicyParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePolicyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePolicyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update_policy",
		Method:             "PUT",
		PathPattern:        "/policies/{policyId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdatePolicyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdatePolicyOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}