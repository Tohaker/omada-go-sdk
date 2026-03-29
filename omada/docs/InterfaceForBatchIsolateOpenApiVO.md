# InterfaceForBatchIsolateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Lan network Id. | [optional] 
**Isolation** | Pointer to **bool** | Whether isolate network. | [optional] 
**Name** | Pointer to **string** | Lan network name. | [optional] 
**Vlan** | Pointer to **int32** | vlan Id. | [optional] 
**Vlans** | Pointer to **string** | multiple vlan Ids | [optional] 

## Methods

### NewInterfaceForBatchIsolateOpenApiVO

`func NewInterfaceForBatchIsolateOpenApiVO() *InterfaceForBatchIsolateOpenApiVO`

NewInterfaceForBatchIsolateOpenApiVO instantiates a new InterfaceForBatchIsolateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceForBatchIsolateOpenApiVOWithDefaults

`func NewInterfaceForBatchIsolateOpenApiVOWithDefaults() *InterfaceForBatchIsolateOpenApiVO`

NewInterfaceForBatchIsolateOpenApiVOWithDefaults instantiates a new InterfaceForBatchIsolateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InterfaceForBatchIsolateOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterfaceForBatchIsolateOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterfaceForBatchIsolateOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InterfaceForBatchIsolateOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsolation

`func (o *InterfaceForBatchIsolateOpenApiVO) GetIsolation() bool`

GetIsolation returns the Isolation field if non-nil, zero value otherwise.

### GetIsolationOk

`func (o *InterfaceForBatchIsolateOpenApiVO) GetIsolationOk() (*bool, bool)`

GetIsolationOk returns a tuple with the Isolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolation

`func (o *InterfaceForBatchIsolateOpenApiVO) SetIsolation(v bool)`

SetIsolation sets Isolation field to given value.

### HasIsolation

`func (o *InterfaceForBatchIsolateOpenApiVO) HasIsolation() bool`

HasIsolation returns a boolean if a field has been set.

### GetName

`func (o *InterfaceForBatchIsolateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterfaceForBatchIsolateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterfaceForBatchIsolateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InterfaceForBatchIsolateOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVlan

`func (o *InterfaceForBatchIsolateOpenApiVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *InterfaceForBatchIsolateOpenApiVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *InterfaceForBatchIsolateOpenApiVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *InterfaceForBatchIsolateOpenApiVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetVlans

`func (o *InterfaceForBatchIsolateOpenApiVO) GetVlans() string`

GetVlans returns the Vlans field if non-nil, zero value otherwise.

### GetVlansOk

`func (o *InterfaceForBatchIsolateOpenApiVO) GetVlansOk() (*string, bool)`

GetVlansOk returns a tuple with the Vlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlans

`func (o *InterfaceForBatchIsolateOpenApiVO) SetVlans(v string)`

SetVlans sets Vlans field to given value.

### HasVlans

`func (o *InterfaceForBatchIsolateOpenApiVO) HasVlans() bool`

HasVlans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


