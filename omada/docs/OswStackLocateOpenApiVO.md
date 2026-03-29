# OswStackLocateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocateEnable** | **bool** | Indicates whether locate is enabled | 
**Macs** | Pointer to **[]string** | Selected Device | [optional] 
**SelectAll** | Pointer to **bool** | Indicates whether to select the entire stack | [optional] 
**StandardPorts** | Pointer to **[]string** | Standard port should be in the format of unit/slot/portId. e.g. 1/0/1 . When parameter[standardPorts] is not empty, paramter[selectAll] and [macs] are not needed | [optional] 

## Methods

### NewOswStackLocateOpenApiVO

`func NewOswStackLocateOpenApiVO(locateEnable bool, ) *OswStackLocateOpenApiVO`

NewOswStackLocateOpenApiVO instantiates a new OswStackLocateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackLocateOpenApiVOWithDefaults

`func NewOswStackLocateOpenApiVOWithDefaults() *OswStackLocateOpenApiVO`

NewOswStackLocateOpenApiVOWithDefaults instantiates a new OswStackLocateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocateEnable

`func (o *OswStackLocateOpenApiVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *OswStackLocateOpenApiVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *OswStackLocateOpenApiVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.


### GetMacs

`func (o *OswStackLocateOpenApiVO) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *OswStackLocateOpenApiVO) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *OswStackLocateOpenApiVO) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *OswStackLocateOpenApiVO) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetSelectAll

`func (o *OswStackLocateOpenApiVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *OswStackLocateOpenApiVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *OswStackLocateOpenApiVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.

### HasSelectAll

`func (o *OswStackLocateOpenApiVO) HasSelectAll() bool`

HasSelectAll returns a boolean if a field has been set.

### GetStandardPorts

`func (o *OswStackLocateOpenApiVO) GetStandardPorts() []string`

GetStandardPorts returns the StandardPorts field if non-nil, zero value otherwise.

### GetStandardPortsOk

`func (o *OswStackLocateOpenApiVO) GetStandardPortsOk() (*[]string, bool)`

GetStandardPortsOk returns a tuple with the StandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPorts

`func (o *OswStackLocateOpenApiVO) SetStandardPorts(v []string)`

SetStandardPorts sets StandardPorts field to given value.

### HasStandardPorts

`func (o *OswStackLocateOpenApiVO) HasStandardPorts() bool`

HasStandardPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


