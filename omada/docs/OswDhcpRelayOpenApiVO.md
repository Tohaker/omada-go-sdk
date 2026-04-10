# OswDhcpRelayOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | Pointer to **string** | Address IP, like 192.168.0.1 | [optional] 
**ServerAddrs** | Pointer to **[]string** | Server Address IP List, like 192.168.0.1 | [optional] 
**VrfId** | Pointer to **string** | VRF ID | [optional] 

## Methods

### NewOswDhcpRelayOpenApiVO

`func NewOswDhcpRelayOpenApiVO() *OswDhcpRelayOpenApiVO`

NewOswDhcpRelayOpenApiVO instantiates a new OswDhcpRelayOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswDhcpRelayOpenApiVOWithDefaults

`func NewOswDhcpRelayOpenApiVOWithDefaults() *OswDhcpRelayOpenApiVO`

NewOswDhcpRelayOpenApiVOWithDefaults instantiates a new OswDhcpRelayOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *OswDhcpRelayOpenApiVO) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *OswDhcpRelayOpenApiVO) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *OswDhcpRelayOpenApiVO) SetAddr(v string)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *OswDhcpRelayOpenApiVO) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetServerAddrs

`func (o *OswDhcpRelayOpenApiVO) GetServerAddrs() []string`

GetServerAddrs returns the ServerAddrs field if non-nil, zero value otherwise.

### GetServerAddrsOk

`func (o *OswDhcpRelayOpenApiVO) GetServerAddrsOk() (*[]string, bool)`

GetServerAddrsOk returns a tuple with the ServerAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddrs

`func (o *OswDhcpRelayOpenApiVO) SetServerAddrs(v []string)`

SetServerAddrs sets ServerAddrs field to given value.

### HasServerAddrs

`func (o *OswDhcpRelayOpenApiVO) HasServerAddrs() bool`

HasServerAddrs returns a boolean if a field has been set.

### GetVrfId

`func (o *OswDhcpRelayOpenApiVO) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *OswDhcpRelayOpenApiVO) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *OswDhcpRelayOpenApiVO) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *OswDhcpRelayOpenApiVO) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


