# ClientTrafficFilterOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to **string** | Network name | [optional] 
**VlanId** | Pointer to **string** | VLAN ID. The VLAN ID range is 1–4096 or Untag. | [optional] 

## Methods

### NewClientTrafficFilterOpenApiVO

`func NewClientTrafficFilterOpenApiVO() *ClientTrafficFilterOpenApiVO`

NewClientTrafficFilterOpenApiVO instantiates a new ClientTrafficFilterOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientTrafficFilterOpenApiVOWithDefaults

`func NewClientTrafficFilterOpenApiVOWithDefaults() *ClientTrafficFilterOpenApiVO`

NewClientTrafficFilterOpenApiVOWithDefaults instantiates a new ClientTrafficFilterOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *ClientTrafficFilterOpenApiVO) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClientTrafficFilterOpenApiVO) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClientTrafficFilterOpenApiVO) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ClientTrafficFilterOpenApiVO) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetVlanId

`func (o *ClientTrafficFilterOpenApiVO) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *ClientTrafficFilterOpenApiVO) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *ClientTrafficFilterOpenApiVO) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *ClientTrafficFilterOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


