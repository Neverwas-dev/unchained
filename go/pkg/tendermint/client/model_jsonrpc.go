/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// JSONRPC struct for JSONRPC
type JSONRPC struct {
	Id *int32 `json:"id,omitempty"`
	Jsonrpc *string `json:"jsonrpc,omitempty"`
}

// NewJSONRPC instantiates a new JSONRPC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONRPC() *JSONRPC {
	this := JSONRPC{}
	return &this
}

// NewJSONRPCWithDefaults instantiates a new JSONRPC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONRPCWithDefaults() *JSONRPC {
	this := JSONRPC{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *JSONRPC) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONRPC) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *JSONRPC) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *JSONRPC) SetId(v int32) {
	o.Id = &v
}

// GetJsonrpc returns the Jsonrpc field value if set, zero value otherwise.
func (o *JSONRPC) GetJsonrpc() string {
	if o == nil || o.Jsonrpc == nil {
		var ret string
		return ret
	}
	return *o.Jsonrpc
}

// GetJsonrpcOk returns a tuple with the Jsonrpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONRPC) GetJsonrpcOk() (*string, bool) {
	if o == nil || o.Jsonrpc == nil {
		return nil, false
	}
	return o.Jsonrpc, true
}

// HasJsonrpc returns a boolean if a field has been set.
func (o *JSONRPC) HasJsonrpc() bool {
	if o != nil && o.Jsonrpc != nil {
		return true
	}

	return false
}

// SetJsonrpc gets a reference to the given string and assigns it to the Jsonrpc field.
func (o *JSONRPC) SetJsonrpc(v string) {
	o.Jsonrpc = &v
}

func (o JSONRPC) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Jsonrpc != nil {
		toSerialize["jsonrpc"] = o.Jsonrpc
	}
	return json.Marshal(toSerialize)
}

type NullableJSONRPC struct {
	value *JSONRPC
	isSet bool
}

func (v NullableJSONRPC) Get() *JSONRPC {
	return v.value
}

func (v *NullableJSONRPC) Set(val *JSONRPC) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONRPC) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONRPC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONRPC(val *JSONRPC) *NullableJSONRPC {
	return &NullableJSONRPC{value: val, isSet: true}
}

func (v NullableJSONRPC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONRPC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


