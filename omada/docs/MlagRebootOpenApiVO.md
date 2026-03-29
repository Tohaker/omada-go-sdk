# MlagRebootOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Macs** | Pointer to **[]string** | Selected Devices Mac List. When selectAll is false, it cannot be empty. | [optional] 
**SelectAll** | **bool** | Indicates whether to select the entire M-LAG group | 

## Methods

### NewMlagRebootOpenApiVO

`func NewMlagRebootOpenApiVO(selectAll bool, ) *MlagRebootOpenApiVO`

NewMlagRebootOpenApiVO instantiates a new MlagRebootOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlagRebootOpenApiVOWithDefaults

`func NewMlagRebootOpenApiVOWithDefaults() *MlagRebootOpenApiVO`

NewMlagRebootOpenApiVOWithDefaults instantiates a new MlagRebootOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacs

`func (o *MlagRebootOpenApiVO) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *MlagRebootOpenApiVO) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *MlagRebootOpenApiVO) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *MlagRebootOpenApiVO) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetSelectAll

`func (o *MlagRebootOpenApiVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *MlagRebootOpenApiVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *MlagRebootOpenApiVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


