# OswNetworkBaseOpenApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Network ID | 
**Status** | Pointer to **bool** | Enable status of the network vlan. | [optional] 
**Vlan** | **int32** | VLAN ID. | 

## Methods

### NewOswNetworkBaseOpenApi

`func NewOswNetworkBaseOpenApi(id string, vlan int32, ) *OswNetworkBaseOpenApi`

NewOswNetworkBaseOpenApi instantiates a new OswNetworkBaseOpenApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswNetworkBaseOpenApiWithDefaults

`func NewOswNetworkBaseOpenApiWithDefaults() *OswNetworkBaseOpenApi`

NewOswNetworkBaseOpenApiWithDefaults instantiates a new OswNetworkBaseOpenApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OswNetworkBaseOpenApi) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswNetworkBaseOpenApi) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswNetworkBaseOpenApi) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *OswNetworkBaseOpenApi) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswNetworkBaseOpenApi) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswNetworkBaseOpenApi) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswNetworkBaseOpenApi) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVlan

`func (o *OswNetworkBaseOpenApi) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *OswNetworkBaseOpenApi) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *OswNetworkBaseOpenApi) SetVlan(v int32)`

SetVlan sets Vlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


