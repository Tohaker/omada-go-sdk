# OswStackRebootOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Macs** | Pointer to **[]string** | Selected Devices | [optional] 
**SelectAll** | Pointer to **bool** | Indicates whether to select the entire stack | [optional] 

## Methods

### NewOswStackRebootOpenApiVO

`func NewOswStackRebootOpenApiVO() *OswStackRebootOpenApiVO`

NewOswStackRebootOpenApiVO instantiates a new OswStackRebootOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackRebootOpenApiVOWithDefaults

`func NewOswStackRebootOpenApiVOWithDefaults() *OswStackRebootOpenApiVO`

NewOswStackRebootOpenApiVOWithDefaults instantiates a new OswStackRebootOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacs

`func (o *OswStackRebootOpenApiVO) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *OswStackRebootOpenApiVO) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *OswStackRebootOpenApiVO) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *OswStackRebootOpenApiVO) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetSelectAll

`func (o *OswStackRebootOpenApiVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *OswStackRebootOpenApiVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *OswStackRebootOpenApiVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.

### HasSelectAll

`func (o *OswStackRebootOpenApiVO) HasSelectAll() bool`

HasSelectAll returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


