# UpdateSsidRateControlOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CckRatesDisable** | Pointer to **bool** | Whether to disable 2G CCK Rates. If this field is true, Parameter [lowerDensity2g] can not enter the following values: [1, 2, 5.5, 11]. | [optional] 
**ClientRatesRequire2g** | Pointer to **bool** | Whether to require clients to use rates at or above the specified value of 2.4GHz Data Rate Control. | [optional] 
**ClientRatesRequire5g** | Pointer to **bool** | Whether to require clients to use rates at or above the specified value of 5GHz Data Rate Control. | [optional] 
**ClientRatesRequire6g** | Pointer to **bool** | Whether to require clients to use rates at or above the specified value of 6GHz Data Rate Control. Note: This field will no longer be supported since Omada Controller V5.14.30. | [optional] 
**HigherDensity2g** | Pointer to **int32** | 2.4GHz Data Rate Control higher density value(Unit: Mbps); It should be a value as follows: [54]. | [optional] 
**HigherDensity5g** | Pointer to **int32** | 5GHz Data Rate Control higher density value(Unit: Mbps); It should be a value as follows: [54]. | [optional] 
**HigherDensity6g** | Pointer to **int32** | 6GHz Data Rate Control higher density value(Unit: Mbps); It should be a value as follows: [54].Note: This field will no longer be supported since Omada Controller V5.14.30. | [optional] 
**LowerDensity2g** | Pointer to **float32** | 2.4GHz Data Rate Control lower density value(Unit: Mbps); It should be a value as follows: [1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54]. | [optional] 
**LowerDensity5g** | Pointer to **int32** | 5GHz Data Rate Control lower density value(Unit: Mbps); It should be a value as follows: [6, 9, 12, 18, 24, 36, 48, 54]. | [optional] 
**LowerDensity6g** | Pointer to **int32** | 6GHz Data Rate Control lower density value(Unit: Mbps); It should be a value as follows: [6, 9, 12, 18, 24, 36, 48, 54].Note: This field will no longer be supported since Omada Controller V5.14.30. | [optional] 
**ManageRateControl2g** | Pointer to **float32** | 2.4GHz Manage Rate Control lower density value(Unit: Mbps); It should be a value as follows: [1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54]. The higher density value is fixed at 54. | [optional] 
**ManageRateControl2gEnable** | Pointer to **bool** | Whether to enable 2GHz Manage Rate Control. | [optional] 
**ManageRateControl5g** | Pointer to **int32** | 5GHz Manage Rate Control lower density value(Unit: Mbps); It should be a value as follows: [6, 9, 12, 18, 24, 36, 48, 54]. The higher density value is fixed at 54. | [optional] 
**ManageRateControl5gEnable** | Pointer to **bool** | Whether to enable 5GHz Manage Rate Control. | [optional] 
**Rate2gCtrlEnable** | **bool** | Whether to enable 2.4GHz Data Rate Control | 
**Rate5gCtrlEnable** | **bool** | Whether to enable 5GHz Data Rate Control. | 
**Rate6gCtrlEnable** | Pointer to **bool** | Whether to enable 6GHz Data Rate Control. Note: This field will no longer be supported since Omada Controller V5.14.30. | [optional] 
**SendBeacons2g** | Pointer to **bool** | Whether to enable send beacons at 1Mbps of 2.4GHz Data Rate Control. | [optional] 
**SendBeacons5g** | Pointer to **bool** | Whether to enable send beacons at 6Mbps of 5GHz Data Rate Control. | [optional] 
**SendBeacons6g** | Pointer to **bool** | Whether to enable send beacons at 6Mbps of 6GHz Data Rate Control. Note: This field will no longer be supported since Omada Controller V5.14.30. | [optional] 

## Methods

### NewUpdateSsidRateControlOpenApiVO

`func NewUpdateSsidRateControlOpenApiVO(rate2gCtrlEnable bool, rate5gCtrlEnable bool, ) *UpdateSsidRateControlOpenApiVO`

NewUpdateSsidRateControlOpenApiVO instantiates a new UpdateSsidRateControlOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSsidRateControlOpenApiVOWithDefaults

`func NewUpdateSsidRateControlOpenApiVOWithDefaults() *UpdateSsidRateControlOpenApiVO`

NewUpdateSsidRateControlOpenApiVOWithDefaults instantiates a new UpdateSsidRateControlOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCckRatesDisable

`func (o *UpdateSsidRateControlOpenApiVO) GetCckRatesDisable() bool`

GetCckRatesDisable returns the CckRatesDisable field if non-nil, zero value otherwise.

### GetCckRatesDisableOk

`func (o *UpdateSsidRateControlOpenApiVO) GetCckRatesDisableOk() (*bool, bool)`

GetCckRatesDisableOk returns a tuple with the CckRatesDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCckRatesDisable

`func (o *UpdateSsidRateControlOpenApiVO) SetCckRatesDisable(v bool)`

SetCckRatesDisable sets CckRatesDisable field to given value.

### HasCckRatesDisable

`func (o *UpdateSsidRateControlOpenApiVO) HasCckRatesDisable() bool`

HasCckRatesDisable returns a boolean if a field has been set.

### GetClientRatesRequire2g

`func (o *UpdateSsidRateControlOpenApiVO) GetClientRatesRequire2g() bool`

GetClientRatesRequire2g returns the ClientRatesRequire2g field if non-nil, zero value otherwise.

### GetClientRatesRequire2gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetClientRatesRequire2gOk() (*bool, bool)`

GetClientRatesRequire2gOk returns a tuple with the ClientRatesRequire2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRatesRequire2g

`func (o *UpdateSsidRateControlOpenApiVO) SetClientRatesRequire2g(v bool)`

SetClientRatesRequire2g sets ClientRatesRequire2g field to given value.

### HasClientRatesRequire2g

`func (o *UpdateSsidRateControlOpenApiVO) HasClientRatesRequire2g() bool`

HasClientRatesRequire2g returns a boolean if a field has been set.

### GetClientRatesRequire5g

`func (o *UpdateSsidRateControlOpenApiVO) GetClientRatesRequire5g() bool`

GetClientRatesRequire5g returns the ClientRatesRequire5g field if non-nil, zero value otherwise.

### GetClientRatesRequire5gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetClientRatesRequire5gOk() (*bool, bool)`

GetClientRatesRequire5gOk returns a tuple with the ClientRatesRequire5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRatesRequire5g

`func (o *UpdateSsidRateControlOpenApiVO) SetClientRatesRequire5g(v bool)`

SetClientRatesRequire5g sets ClientRatesRequire5g field to given value.

### HasClientRatesRequire5g

`func (o *UpdateSsidRateControlOpenApiVO) HasClientRatesRequire5g() bool`

HasClientRatesRequire5g returns a boolean if a field has been set.

### GetClientRatesRequire6g

`func (o *UpdateSsidRateControlOpenApiVO) GetClientRatesRequire6g() bool`

GetClientRatesRequire6g returns the ClientRatesRequire6g field if non-nil, zero value otherwise.

### GetClientRatesRequire6gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetClientRatesRequire6gOk() (*bool, bool)`

GetClientRatesRequire6gOk returns a tuple with the ClientRatesRequire6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientRatesRequire6g

`func (o *UpdateSsidRateControlOpenApiVO) SetClientRatesRequire6g(v bool)`

SetClientRatesRequire6g sets ClientRatesRequire6g field to given value.

### HasClientRatesRequire6g

`func (o *UpdateSsidRateControlOpenApiVO) HasClientRatesRequire6g() bool`

HasClientRatesRequire6g returns a boolean if a field has been set.

### GetHigherDensity2g

`func (o *UpdateSsidRateControlOpenApiVO) GetHigherDensity2g() int32`

GetHigherDensity2g returns the HigherDensity2g field if non-nil, zero value otherwise.

### GetHigherDensity2gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetHigherDensity2gOk() (*int32, bool)`

GetHigherDensity2gOk returns a tuple with the HigherDensity2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigherDensity2g

`func (o *UpdateSsidRateControlOpenApiVO) SetHigherDensity2g(v int32)`

SetHigherDensity2g sets HigherDensity2g field to given value.

### HasHigherDensity2g

`func (o *UpdateSsidRateControlOpenApiVO) HasHigherDensity2g() bool`

HasHigherDensity2g returns a boolean if a field has been set.

### GetHigherDensity5g

`func (o *UpdateSsidRateControlOpenApiVO) GetHigherDensity5g() int32`

GetHigherDensity5g returns the HigherDensity5g field if non-nil, zero value otherwise.

### GetHigherDensity5gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetHigherDensity5gOk() (*int32, bool)`

GetHigherDensity5gOk returns a tuple with the HigherDensity5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigherDensity5g

`func (o *UpdateSsidRateControlOpenApiVO) SetHigherDensity5g(v int32)`

SetHigherDensity5g sets HigherDensity5g field to given value.

### HasHigherDensity5g

`func (o *UpdateSsidRateControlOpenApiVO) HasHigherDensity5g() bool`

HasHigherDensity5g returns a boolean if a field has been set.

### GetHigherDensity6g

`func (o *UpdateSsidRateControlOpenApiVO) GetHigherDensity6g() int32`

GetHigherDensity6g returns the HigherDensity6g field if non-nil, zero value otherwise.

### GetHigherDensity6gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetHigherDensity6gOk() (*int32, bool)`

GetHigherDensity6gOk returns a tuple with the HigherDensity6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigherDensity6g

`func (o *UpdateSsidRateControlOpenApiVO) SetHigherDensity6g(v int32)`

SetHigherDensity6g sets HigherDensity6g field to given value.

### HasHigherDensity6g

`func (o *UpdateSsidRateControlOpenApiVO) HasHigherDensity6g() bool`

HasHigherDensity6g returns a boolean if a field has been set.

### GetLowerDensity2g

`func (o *UpdateSsidRateControlOpenApiVO) GetLowerDensity2g() float32`

GetLowerDensity2g returns the LowerDensity2g field if non-nil, zero value otherwise.

### GetLowerDensity2gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetLowerDensity2gOk() (*float32, bool)`

GetLowerDensity2gOk returns a tuple with the LowerDensity2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerDensity2g

`func (o *UpdateSsidRateControlOpenApiVO) SetLowerDensity2g(v float32)`

SetLowerDensity2g sets LowerDensity2g field to given value.

### HasLowerDensity2g

`func (o *UpdateSsidRateControlOpenApiVO) HasLowerDensity2g() bool`

HasLowerDensity2g returns a boolean if a field has been set.

### GetLowerDensity5g

`func (o *UpdateSsidRateControlOpenApiVO) GetLowerDensity5g() int32`

GetLowerDensity5g returns the LowerDensity5g field if non-nil, zero value otherwise.

### GetLowerDensity5gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetLowerDensity5gOk() (*int32, bool)`

GetLowerDensity5gOk returns a tuple with the LowerDensity5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerDensity5g

`func (o *UpdateSsidRateControlOpenApiVO) SetLowerDensity5g(v int32)`

SetLowerDensity5g sets LowerDensity5g field to given value.

### HasLowerDensity5g

`func (o *UpdateSsidRateControlOpenApiVO) HasLowerDensity5g() bool`

HasLowerDensity5g returns a boolean if a field has been set.

### GetLowerDensity6g

`func (o *UpdateSsidRateControlOpenApiVO) GetLowerDensity6g() int32`

GetLowerDensity6g returns the LowerDensity6g field if non-nil, zero value otherwise.

### GetLowerDensity6gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetLowerDensity6gOk() (*int32, bool)`

GetLowerDensity6gOk returns a tuple with the LowerDensity6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerDensity6g

`func (o *UpdateSsidRateControlOpenApiVO) SetLowerDensity6g(v int32)`

SetLowerDensity6g sets LowerDensity6g field to given value.

### HasLowerDensity6g

`func (o *UpdateSsidRateControlOpenApiVO) HasLowerDensity6g() bool`

HasLowerDensity6g returns a boolean if a field has been set.

### GetManageRateControl2g

`func (o *UpdateSsidRateControlOpenApiVO) GetManageRateControl2g() float32`

GetManageRateControl2g returns the ManageRateControl2g field if non-nil, zero value otherwise.

### GetManageRateControl2gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetManageRateControl2gOk() (*float32, bool)`

GetManageRateControl2gOk returns a tuple with the ManageRateControl2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageRateControl2g

`func (o *UpdateSsidRateControlOpenApiVO) SetManageRateControl2g(v float32)`

SetManageRateControl2g sets ManageRateControl2g field to given value.

### HasManageRateControl2g

`func (o *UpdateSsidRateControlOpenApiVO) HasManageRateControl2g() bool`

HasManageRateControl2g returns a boolean if a field has been set.

### GetManageRateControl2gEnable

`func (o *UpdateSsidRateControlOpenApiVO) GetManageRateControl2gEnable() bool`

GetManageRateControl2gEnable returns the ManageRateControl2gEnable field if non-nil, zero value otherwise.

### GetManageRateControl2gEnableOk

`func (o *UpdateSsidRateControlOpenApiVO) GetManageRateControl2gEnableOk() (*bool, bool)`

GetManageRateControl2gEnableOk returns a tuple with the ManageRateControl2gEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageRateControl2gEnable

`func (o *UpdateSsidRateControlOpenApiVO) SetManageRateControl2gEnable(v bool)`

SetManageRateControl2gEnable sets ManageRateControl2gEnable field to given value.

### HasManageRateControl2gEnable

`func (o *UpdateSsidRateControlOpenApiVO) HasManageRateControl2gEnable() bool`

HasManageRateControl2gEnable returns a boolean if a field has been set.

### GetManageRateControl5g

`func (o *UpdateSsidRateControlOpenApiVO) GetManageRateControl5g() int32`

GetManageRateControl5g returns the ManageRateControl5g field if non-nil, zero value otherwise.

### GetManageRateControl5gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetManageRateControl5gOk() (*int32, bool)`

GetManageRateControl5gOk returns a tuple with the ManageRateControl5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageRateControl5g

`func (o *UpdateSsidRateControlOpenApiVO) SetManageRateControl5g(v int32)`

SetManageRateControl5g sets ManageRateControl5g field to given value.

### HasManageRateControl5g

`func (o *UpdateSsidRateControlOpenApiVO) HasManageRateControl5g() bool`

HasManageRateControl5g returns a boolean if a field has been set.

### GetManageRateControl5gEnable

`func (o *UpdateSsidRateControlOpenApiVO) GetManageRateControl5gEnable() bool`

GetManageRateControl5gEnable returns the ManageRateControl5gEnable field if non-nil, zero value otherwise.

### GetManageRateControl5gEnableOk

`func (o *UpdateSsidRateControlOpenApiVO) GetManageRateControl5gEnableOk() (*bool, bool)`

GetManageRateControl5gEnableOk returns a tuple with the ManageRateControl5gEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageRateControl5gEnable

`func (o *UpdateSsidRateControlOpenApiVO) SetManageRateControl5gEnable(v bool)`

SetManageRateControl5gEnable sets ManageRateControl5gEnable field to given value.

### HasManageRateControl5gEnable

`func (o *UpdateSsidRateControlOpenApiVO) HasManageRateControl5gEnable() bool`

HasManageRateControl5gEnable returns a boolean if a field has been set.

### GetRate2gCtrlEnable

`func (o *UpdateSsidRateControlOpenApiVO) GetRate2gCtrlEnable() bool`

GetRate2gCtrlEnable returns the Rate2gCtrlEnable field if non-nil, zero value otherwise.

### GetRate2gCtrlEnableOk

`func (o *UpdateSsidRateControlOpenApiVO) GetRate2gCtrlEnableOk() (*bool, bool)`

GetRate2gCtrlEnableOk returns a tuple with the Rate2gCtrlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate2gCtrlEnable

`func (o *UpdateSsidRateControlOpenApiVO) SetRate2gCtrlEnable(v bool)`

SetRate2gCtrlEnable sets Rate2gCtrlEnable field to given value.


### GetRate5gCtrlEnable

`func (o *UpdateSsidRateControlOpenApiVO) GetRate5gCtrlEnable() bool`

GetRate5gCtrlEnable returns the Rate5gCtrlEnable field if non-nil, zero value otherwise.

### GetRate5gCtrlEnableOk

`func (o *UpdateSsidRateControlOpenApiVO) GetRate5gCtrlEnableOk() (*bool, bool)`

GetRate5gCtrlEnableOk returns a tuple with the Rate5gCtrlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate5gCtrlEnable

`func (o *UpdateSsidRateControlOpenApiVO) SetRate5gCtrlEnable(v bool)`

SetRate5gCtrlEnable sets Rate5gCtrlEnable field to given value.


### GetRate6gCtrlEnable

`func (o *UpdateSsidRateControlOpenApiVO) GetRate6gCtrlEnable() bool`

GetRate6gCtrlEnable returns the Rate6gCtrlEnable field if non-nil, zero value otherwise.

### GetRate6gCtrlEnableOk

`func (o *UpdateSsidRateControlOpenApiVO) GetRate6gCtrlEnableOk() (*bool, bool)`

GetRate6gCtrlEnableOk returns a tuple with the Rate6gCtrlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate6gCtrlEnable

`func (o *UpdateSsidRateControlOpenApiVO) SetRate6gCtrlEnable(v bool)`

SetRate6gCtrlEnable sets Rate6gCtrlEnable field to given value.

### HasRate6gCtrlEnable

`func (o *UpdateSsidRateControlOpenApiVO) HasRate6gCtrlEnable() bool`

HasRate6gCtrlEnable returns a boolean if a field has been set.

### GetSendBeacons2g

`func (o *UpdateSsidRateControlOpenApiVO) GetSendBeacons2g() bool`

GetSendBeacons2g returns the SendBeacons2g field if non-nil, zero value otherwise.

### GetSendBeacons2gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetSendBeacons2gOk() (*bool, bool)`

GetSendBeacons2gOk returns a tuple with the SendBeacons2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBeacons2g

`func (o *UpdateSsidRateControlOpenApiVO) SetSendBeacons2g(v bool)`

SetSendBeacons2g sets SendBeacons2g field to given value.

### HasSendBeacons2g

`func (o *UpdateSsidRateControlOpenApiVO) HasSendBeacons2g() bool`

HasSendBeacons2g returns a boolean if a field has been set.

### GetSendBeacons5g

`func (o *UpdateSsidRateControlOpenApiVO) GetSendBeacons5g() bool`

GetSendBeacons5g returns the SendBeacons5g field if non-nil, zero value otherwise.

### GetSendBeacons5gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetSendBeacons5gOk() (*bool, bool)`

GetSendBeacons5gOk returns a tuple with the SendBeacons5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBeacons5g

`func (o *UpdateSsidRateControlOpenApiVO) SetSendBeacons5g(v bool)`

SetSendBeacons5g sets SendBeacons5g field to given value.

### HasSendBeacons5g

`func (o *UpdateSsidRateControlOpenApiVO) HasSendBeacons5g() bool`

HasSendBeacons5g returns a boolean if a field has been set.

### GetSendBeacons6g

`func (o *UpdateSsidRateControlOpenApiVO) GetSendBeacons6g() bool`

GetSendBeacons6g returns the SendBeacons6g field if non-nil, zero value otherwise.

### GetSendBeacons6gOk

`func (o *UpdateSsidRateControlOpenApiVO) GetSendBeacons6gOk() (*bool, bool)`

GetSendBeacons6gOk returns a tuple with the SendBeacons6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBeacons6g

`func (o *UpdateSsidRateControlOpenApiVO) SetSendBeacons6g(v bool)`

SetSendBeacons6g sets SendBeacons6g field to given value.

### HasSendBeacons6g

`func (o *UpdateSsidRateControlOpenApiVO) HasSendBeacons6g() bool`

HasSendBeacons6g returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


