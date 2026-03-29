# BatchConfigIotTransmitPowerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Macs** | **[]string** | Device mac list. If selectAll is true, mac in macs is excluded. | 
**Override** | **bool** | Whether to override the configuration. | 
**SearchKey** | Pointer to **string** | Look for a specific piece of data. | [optional] 
**SelectAll** | **bool** | Whether to select all device. The selected devices are the ones filtered out based on the searchKey. | 
**TransmitPower** | Pointer to **int32** | Broadcast transmission power.&lt;br /&gt;The parameter [transmitPower] should be a value as follows:[-20, -18, -15, -12, -10, -9, -6, -5, -3, 0, 1, 2, 3, 4, 5, 14, 15, 16, 17, 18, 19, 20].(0 by default) | [optional] 

## Methods

### NewBatchConfigIotTransmitPowerOpenApiVO

`func NewBatchConfigIotTransmitPowerOpenApiVO(macs []string, override bool, selectAll bool, ) *BatchConfigIotTransmitPowerOpenApiVO`

NewBatchConfigIotTransmitPowerOpenApiVO instantiates a new BatchConfigIotTransmitPowerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchConfigIotTransmitPowerOpenApiVOWithDefaults

`func NewBatchConfigIotTransmitPowerOpenApiVOWithDefaults() *BatchConfigIotTransmitPowerOpenApiVO`

NewBatchConfigIotTransmitPowerOpenApiVOWithDefaults instantiates a new BatchConfigIotTransmitPowerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacs

`func (o *BatchConfigIotTransmitPowerOpenApiVO) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *BatchConfigIotTransmitPowerOpenApiVO) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *BatchConfigIotTransmitPowerOpenApiVO) SetMacs(v []string)`

SetMacs sets Macs field to given value.


### GetOverride

`func (o *BatchConfigIotTransmitPowerOpenApiVO) GetOverride() bool`

GetOverride returns the Override field if non-nil, zero value otherwise.

### GetOverrideOk

`func (o *BatchConfigIotTransmitPowerOpenApiVO) GetOverrideOk() (*bool, bool)`

GetOverrideOk returns a tuple with the Override field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverride

`func (o *BatchConfigIotTransmitPowerOpenApiVO) SetOverride(v bool)`

SetOverride sets Override field to given value.


### GetSearchKey

`func (o *BatchConfigIotTransmitPowerOpenApiVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *BatchConfigIotTransmitPowerOpenApiVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *BatchConfigIotTransmitPowerOpenApiVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *BatchConfigIotTransmitPowerOpenApiVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectAll

`func (o *BatchConfigIotTransmitPowerOpenApiVO) GetSelectAll() bool`

GetSelectAll returns the SelectAll field if non-nil, zero value otherwise.

### GetSelectAllOk

`func (o *BatchConfigIotTransmitPowerOpenApiVO) GetSelectAllOk() (*bool, bool)`

GetSelectAllOk returns a tuple with the SelectAll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectAll

`func (o *BatchConfigIotTransmitPowerOpenApiVO) SetSelectAll(v bool)`

SetSelectAll sets SelectAll field to given value.


### GetTransmitPower

`func (o *BatchConfigIotTransmitPowerOpenApiVO) GetTransmitPower() int32`

GetTransmitPower returns the TransmitPower field if non-nil, zero value otherwise.

### GetTransmitPowerOk

`func (o *BatchConfigIotTransmitPowerOpenApiVO) GetTransmitPowerOk() (*int32, bool)`

GetTransmitPowerOk returns a tuple with the TransmitPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitPower

`func (o *BatchConfigIotTransmitPowerOpenApiVO) SetTransmitPower(v int32)`

SetTransmitPower sets TransmitPower field to given value.

### HasTransmitPower

`func (o *BatchConfigIotTransmitPowerOpenApiVO) HasTransmitPower() bool`

HasTransmitPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


