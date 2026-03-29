# MspClientDetailInfoSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientDataTrendEnable** | Pointer to **bool** | Whether client data trend record is enabled. When enabled, client trend statistics and charts will be retained, which will take up lots of storage space. | [optional] 
**ClientHealthEnable** | Pointer to **bool** | Whether client health is enabled. When enabled, client health data will be recorded, which may consume a significant amount of storage space. | [optional] 
**ClientHistoryEnable** | Pointer to **bool** | Whether client history is enabled. When enabled, client history, client logs will be recorded. This will occupy much storage space. | [optional] 

## Methods

### NewMspClientDetailInfoSettingVO

`func NewMspClientDetailInfoSettingVO() *MspClientDetailInfoSettingVO`

NewMspClientDetailInfoSettingVO instantiates a new MspClientDetailInfoSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspClientDetailInfoSettingVOWithDefaults

`func NewMspClientDetailInfoSettingVOWithDefaults() *MspClientDetailInfoSettingVO`

NewMspClientDetailInfoSettingVOWithDefaults instantiates a new MspClientDetailInfoSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientDataTrendEnable

`func (o *MspClientDetailInfoSettingVO) GetClientDataTrendEnable() bool`

GetClientDataTrendEnable returns the ClientDataTrendEnable field if non-nil, zero value otherwise.

### GetClientDataTrendEnableOk

`func (o *MspClientDetailInfoSettingVO) GetClientDataTrendEnableOk() (*bool, bool)`

GetClientDataTrendEnableOk returns a tuple with the ClientDataTrendEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDataTrendEnable

`func (o *MspClientDetailInfoSettingVO) SetClientDataTrendEnable(v bool)`

SetClientDataTrendEnable sets ClientDataTrendEnable field to given value.

### HasClientDataTrendEnable

`func (o *MspClientDetailInfoSettingVO) HasClientDataTrendEnable() bool`

HasClientDataTrendEnable returns a boolean if a field has been set.

### GetClientHealthEnable

`func (o *MspClientDetailInfoSettingVO) GetClientHealthEnable() bool`

GetClientHealthEnable returns the ClientHealthEnable field if non-nil, zero value otherwise.

### GetClientHealthEnableOk

`func (o *MspClientDetailInfoSettingVO) GetClientHealthEnableOk() (*bool, bool)`

GetClientHealthEnableOk returns a tuple with the ClientHealthEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHealthEnable

`func (o *MspClientDetailInfoSettingVO) SetClientHealthEnable(v bool)`

SetClientHealthEnable sets ClientHealthEnable field to given value.

### HasClientHealthEnable

`func (o *MspClientDetailInfoSettingVO) HasClientHealthEnable() bool`

HasClientHealthEnable returns a boolean if a field has been set.

### GetClientHistoryEnable

`func (o *MspClientDetailInfoSettingVO) GetClientHistoryEnable() bool`

GetClientHistoryEnable returns the ClientHistoryEnable field if non-nil, zero value otherwise.

### GetClientHistoryEnableOk

`func (o *MspClientDetailInfoSettingVO) GetClientHistoryEnableOk() (*bool, bool)`

GetClientHistoryEnableOk returns a tuple with the ClientHistoryEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHistoryEnable

`func (o *MspClientDetailInfoSettingVO) SetClientHistoryEnable(v bool)`

SetClientHistoryEnable sets ClientHistoryEnable field to given value.

### HasClientHistoryEnable

`func (o *MspClientDetailInfoSettingVO) HasClientHistoryEnable() bool`

HasClientHistoryEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


