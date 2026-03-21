# MlagLocateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocateEnable** | **bool** | Indicates whether locate is enabled | 
**Macs** | Pointer to **[]string** | Selected Devices Mac List. When selectAll is false, it cannot be empty. | [optional] 
**SelectAll** | **bool** | Indicates whether to select the entire M-LAG group | 

## Methods

### NewMlagLocateOpenApiVO

`func NewMlagLocateOpenApiVO(locateEnable bool, selectAll bool, ) *MlagLocateOpenApiVO`

NewMlagLocateOpenApiVO instantiates a new MlagLocateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlagLocateOpenApiVOWithDefaults

`func NewMlagLocateOpenApiVOWithDefaults() *MlagLocateOpenApiVO`

NewMlagLocateOpenApiVOWithDefaults instantiates a new MlagLocateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocateEnable

`func (o *MlagLocateOpenApiVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *MlagLocateOpenApiVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *MlagLocateOpenApiVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.


### GetMacs

`func (o *MlagLocateOpenApiVO) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *MlagLocateOpenApiVO) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *MlagLocateOpenApiVO) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *MlagLocateOpenApiVO) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetSelectAll

`func (o *MlagLocateOpenApiVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *MlagLocateOpenApiVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *MlagLocateOpenApiVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


