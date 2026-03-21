# OswNetworkDetailInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Network&#39;s ID | [optional] 
**Isolation** | Pointer to **bool** | Whether network isolated. | [optional] 
**Name** | Pointer to **string** | Network&#39;s name | [optional] 
**PortLagNetworkInfo** | Pointer to [**OswPortAndLagNetworkVO**](OswPortAndLagNetworkVO.md) |  | [optional] 
**Vlan** | Pointer to **int32** | VLAN ID | [optional] 

## Methods

### NewOswNetworkDetailInfoOpenApiVO

`func NewOswNetworkDetailInfoOpenApiVO() *OswNetworkDetailInfoOpenApiVO`

NewOswNetworkDetailInfoOpenApiVO instantiates a new OswNetworkDetailInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswNetworkDetailInfoOpenApiVOWithDefaults

`func NewOswNetworkDetailInfoOpenApiVOWithDefaults() *OswNetworkDetailInfoOpenApiVO`

NewOswNetworkDetailInfoOpenApiVOWithDefaults instantiates a new OswNetworkDetailInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OswNetworkDetailInfoOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswNetworkDetailInfoOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswNetworkDetailInfoOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswNetworkDetailInfoOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsolation

`func (o *OswNetworkDetailInfoOpenApiVO) GetIsolation() bool`

GetIsolation returns the Isolation field if non-nil, zero value otherwise.

### GetIsolationOk

`func (o *OswNetworkDetailInfoOpenApiVO) GetIsolationOk() (*bool, bool)`

GetIsolationOk returns a tuple with the Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolation

`func (o *OswNetworkDetailInfoOpenApiVO) SetIsolation(v bool)`

SetIsolation sets Isolation field to given value.

### HasIsolation

`func (o *OswNetworkDetailInfoOpenApiVO) HasIsolation() bool`

HasIsolation returns a boolean if a field has been set.

### GetName

`func (o *OswNetworkDetailInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswNetworkDetailInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswNetworkDetailInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswNetworkDetailInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortLagNetworkInfo

`func (o *OswNetworkDetailInfoOpenApiVO) GetPortLagNetworkInfo() OswPortAndLagNetworkVO`

GetPortLagNetworkInfo returns the PortLagNetworkInfo field if non-nil, zero value otherwise.

### GetPortLagNetworkInfoOk

`func (o *OswNetworkDetailInfoOpenApiVO) GetPortLagNetworkInfoOk() (*OswPortAndLagNetworkVO, bool)`

GetPortLagNetworkInfoOk returns a tuple with the PortLagNetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLagNetworkInfo

`func (o *OswNetworkDetailInfoOpenApiVO) SetPortLagNetworkInfo(v OswPortAndLagNetworkVO)`

SetPortLagNetworkInfo sets PortLagNetworkInfo field to given value.

### HasPortLagNetworkInfo

`func (o *OswNetworkDetailInfoOpenApiVO) HasPortLagNetworkInfo() bool`

HasPortLagNetworkInfo returns a boolean if a field has been set.

### GetVlan

`func (o *OswNetworkDetailInfoOpenApiVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *OswNetworkDetailInfoOpenApiVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *OswNetworkDetailInfoOpenApiVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *OswNetworkDetailInfoOpenApiVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


