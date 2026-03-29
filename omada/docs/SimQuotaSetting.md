# SimQuotaSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardStatus** | Pointer to **int32** | Sim card status. | [optional] 
**DataSetting** | Pointer to [**QuotaDataSettingOpenApiVO**](QuotaDataSettingOpenApiVO.md) |  | [optional] 
**Resource** | Pointer to **int32** | Sim quota setting creation resource,such as: 0: new created, 1: from template, 2: override. | [optional] 
**SimCard** | Pointer to **int32** | When device supports Dual-SIM card, using parameter [simCard] to point which card to configure. 1: SIM1; 2:SIM2. | [optional] 
**SmsSetting** | Pointer to [**QuotaSmsSettingOpenApiVO**](QuotaSmsSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewSimQuotaSetting

`func NewSimQuotaSetting() *SimQuotaSetting`

NewSimQuotaSetting instantiates a new SimQuotaSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimQuotaSettingWithDefaults

`func NewSimQuotaSettingWithDefaults() *SimQuotaSetting`

NewSimQuotaSettingWithDefaults instantiates a new SimQuotaSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardStatus

`func (o *SimQuotaSetting) GetCardStatus() int32`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *SimQuotaSetting) GetCardStatusOk() (*int32, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *SimQuotaSetting) SetCardStatus(v int32)`

SetCardStatus sets CardStatus field to given value.

### HasCardStatus

`func (o *SimQuotaSetting) HasCardStatus() bool`

HasCardStatus returns a boolean if a field has been set.

### GetDataSetting

`func (o *SimQuotaSetting) GetDataSetting() QuotaDataSettingOpenApiVO`

GetDataSetting returns the DataSetting field if non-nil, zero value otherwise.

### GetDataSettingOk

`func (o *SimQuotaSetting) GetDataSettingOk() (*QuotaDataSettingOpenApiVO, bool)`

GetDataSettingOk returns a tuple with the DataSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSetting

`func (o *SimQuotaSetting) SetDataSetting(v QuotaDataSettingOpenApiVO)`

SetDataSetting sets DataSetting field to given value.

### HasDataSetting

`func (o *SimQuotaSetting) HasDataSetting() bool`

HasDataSetting returns a boolean if a field has been set.

### GetResource

`func (o *SimQuotaSetting) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *SimQuotaSetting) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *SimQuotaSetting) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *SimQuotaSetting) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSimCard

`func (o *SimQuotaSetting) GetSimCard() int32`

GetSimCard returns the SimCard field if non-nil, zero value otherwise.

### GetSimCardOk

`func (o *SimQuotaSetting) GetSimCardOk() (*int32, bool)`

GetSimCardOk returns a tuple with the SimCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCard

`func (o *SimQuotaSetting) SetSimCard(v int32)`

SetSimCard sets SimCard field to given value.

### HasSimCard

`func (o *SimQuotaSetting) HasSimCard() bool`

HasSimCard returns a boolean if a field has been set.

### GetSmsSetting

`func (o *SimQuotaSetting) GetSmsSetting() QuotaSmsSettingOpenApiVO`

GetSmsSetting returns the SmsSetting field if non-nil, zero value otherwise.

### GetSmsSettingOk

`func (o *SimQuotaSetting) GetSmsSettingOk() (*QuotaSmsSettingOpenApiVO, bool)`

GetSmsSettingOk returns a tuple with the SmsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsSetting

`func (o *SimQuotaSetting) SetSmsSetting(v QuotaSmsSettingOpenApiVO)`

SetSmsSetting sets SmsSetting field to given value.

### HasSmsSetting

`func (o *SimQuotaSetting) HasSmsSetting() bool`

HasSmsSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


