# ApplicationStatVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowApps** | Pointer to **int32** |  | [optional] 
**BlockApps** | Pointer to **int32** |  | [optional] 

## Methods

### NewApplicationStatVO

`func NewApplicationStatVO() *ApplicationStatVO`

NewApplicationStatVO instantiates a new ApplicationStatVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationStatVOWithDefaults

`func NewApplicationStatVOWithDefaults() *ApplicationStatVO`

NewApplicationStatVOWithDefaults instantiates a new ApplicationStatVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowApps

`func (o *ApplicationStatVO) GetAllowApps() int32`

GetAllowApps returns the AllowApps field if non-nil, zero value otherwise.

### GetAllowAppsOk

`func (o *ApplicationStatVO) GetAllowAppsOk() (*int32, bool)`

GetAllowAppsOk returns a tuple with the AllowApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowApps

`func (o *ApplicationStatVO) SetAllowApps(v int32)`

SetAllowApps sets AllowApps field to given value.

### HasAllowApps

`func (o *ApplicationStatVO) HasAllowApps() bool`

HasAllowApps returns a boolean if a field has been set.

### GetBlockApps

`func (o *ApplicationStatVO) GetBlockApps() int32`

GetBlockApps returns the BlockApps field if non-nil, zero value otherwise.

### GetBlockAppsOk

`func (o *ApplicationStatVO) GetBlockAppsOk() (*int32, bool)`

GetBlockAppsOk returns a tuple with the BlockApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockApps

`func (o *ApplicationStatVO) SetBlockApps(v int32)`

SetBlockApps sets BlockApps field to given value.

### HasBlockApps

`func (o *ApplicationStatVO) HasBlockApps() bool`

HasBlockApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


