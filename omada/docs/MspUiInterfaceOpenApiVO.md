# MspUiInterfaceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControllerNotification** | Pointer to **bool** | Enable controller upgrade notification | [optional] 
**FixedMenu** | **bool** |  | 
**Language** | Pointer to **int32** | It should be a value as follows: 1: English; 4: German; 7: French; 8: Spanish; 10: Italian; 12: Portuguese; 13: Russian; 15: Turkish; 17: Japanese; 18: Traditional Chinese; 21: Korean | [optional] 
**PrivateLabelingEnable** | Pointer to **bool** | Enable private label | [optional] 
**PrivateLabelingFileId** | Pointer to **string** | The file id of private label picture | [optional] 
**PrivateLabelingFileName** | Pointer to **string** | The file name of private label picture | [optional] 
**PrivateLabelingUrl** | Pointer to **string** | The url of private label | [optional] 
**RefreshBtnEnable** | **bool** | Enable refresh button | 
**RefreshRate** | **int32** | It should be a value as follows: 0: 15 seconds; 1: 1 minute; 2: 2 minutes; 3: 5 minutes; 4: never refresh | 
**ShowPDevices** | Pointer to **bool** | Show pending devices | [optional] 
**Theme** | Pointer to **int32** | It should be a value as follows: 0: default; 1: dark | [optional] 
**Use24hour** | **bool** |  | 
**WebsocketEnable** | **bool** | Enable websocket connection | 

## Methods

### NewMspUiInterfaceOpenApiVO

`func NewMspUiInterfaceOpenApiVO(fixedMenu bool, refreshBtnEnable bool, refreshRate int32, use24hour bool, websocketEnable bool, ) *MspUiInterfaceOpenApiVO`

NewMspUiInterfaceOpenApiVO instantiates a new MspUiInterfaceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspUiInterfaceOpenApiVOWithDefaults

`func NewMspUiInterfaceOpenApiVOWithDefaults() *MspUiInterfaceOpenApiVO`

NewMspUiInterfaceOpenApiVOWithDefaults instantiates a new MspUiInterfaceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControllerNotification

`func (o *MspUiInterfaceOpenApiVO) GetControllerNotification() bool`

GetControllerNotification returns the ControllerNotification field if non-nil, zero value otherwise.

### GetControllerNotificationOk

`func (o *MspUiInterfaceOpenApiVO) GetControllerNotificationOk() (*bool, bool)`

GetControllerNotificationOk returns a tuple with the ControllerNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControllerNotification

`func (o *MspUiInterfaceOpenApiVO) SetControllerNotification(v bool)`

SetControllerNotification sets ControllerNotification field to given value.

### HasControllerNotification

`func (o *MspUiInterfaceOpenApiVO) HasControllerNotification() bool`

HasControllerNotification returns a boolean if a field has been set.

### GetFixedMenu

`func (o *MspUiInterfaceOpenApiVO) GetFixedMenu() bool`

GetFixedMenu returns the FixedMenu field if non-nil, zero value otherwise.

### GetFixedMenuOk

`func (o *MspUiInterfaceOpenApiVO) GetFixedMenuOk() (*bool, bool)`

GetFixedMenuOk returns a tuple with the FixedMenu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedMenu

`func (o *MspUiInterfaceOpenApiVO) SetFixedMenu(v bool)`

SetFixedMenu sets FixedMenu field to given value.


### GetLanguage

`func (o *MspUiInterfaceOpenApiVO) GetLanguage() int32`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *MspUiInterfaceOpenApiVO) GetLanguageOk() (*int32, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *MspUiInterfaceOpenApiVO) SetLanguage(v int32)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *MspUiInterfaceOpenApiVO) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetPrivateLabelingEnable

`func (o *MspUiInterfaceOpenApiVO) GetPrivateLabelingEnable() bool`

GetPrivateLabelingEnable returns the PrivateLabelingEnable field if non-nil, zero value otherwise.

### GetPrivateLabelingEnableOk

`func (o *MspUiInterfaceOpenApiVO) GetPrivateLabelingEnableOk() (*bool, bool)`

GetPrivateLabelingEnableOk returns a tuple with the PrivateLabelingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingEnable

`func (o *MspUiInterfaceOpenApiVO) SetPrivateLabelingEnable(v bool)`

SetPrivateLabelingEnable sets PrivateLabelingEnable field to given value.

### HasPrivateLabelingEnable

`func (o *MspUiInterfaceOpenApiVO) HasPrivateLabelingEnable() bool`

HasPrivateLabelingEnable returns a boolean if a field has been set.

### GetPrivateLabelingFileId

`func (o *MspUiInterfaceOpenApiVO) GetPrivateLabelingFileId() string`

GetPrivateLabelingFileId returns the PrivateLabelingFileId field if non-nil, zero value otherwise.

### GetPrivateLabelingFileIdOk

`func (o *MspUiInterfaceOpenApiVO) GetPrivateLabelingFileIdOk() (*string, bool)`

GetPrivateLabelingFileIdOk returns a tuple with the PrivateLabelingFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingFileId

`func (o *MspUiInterfaceOpenApiVO) SetPrivateLabelingFileId(v string)`

SetPrivateLabelingFileId sets PrivateLabelingFileId field to given value.

### HasPrivateLabelingFileId

`func (o *MspUiInterfaceOpenApiVO) HasPrivateLabelingFileId() bool`

HasPrivateLabelingFileId returns a boolean if a field has been set.

### GetPrivateLabelingFileName

`func (o *MspUiInterfaceOpenApiVO) GetPrivateLabelingFileName() string`

GetPrivateLabelingFileName returns the PrivateLabelingFileName field if non-nil, zero value otherwise.

### GetPrivateLabelingFileNameOk

`func (o *MspUiInterfaceOpenApiVO) GetPrivateLabelingFileNameOk() (*string, bool)`

GetPrivateLabelingFileNameOk returns a tuple with the PrivateLabelingFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingFileName

`func (o *MspUiInterfaceOpenApiVO) SetPrivateLabelingFileName(v string)`

SetPrivateLabelingFileName sets PrivateLabelingFileName field to given value.

### HasPrivateLabelingFileName

`func (o *MspUiInterfaceOpenApiVO) HasPrivateLabelingFileName() bool`

HasPrivateLabelingFileName returns a boolean if a field has been set.

### GetPrivateLabelingUrl

`func (o *MspUiInterfaceOpenApiVO) GetPrivateLabelingUrl() string`

GetPrivateLabelingUrl returns the PrivateLabelingUrl field if non-nil, zero value otherwise.

### GetPrivateLabelingUrlOk

`func (o *MspUiInterfaceOpenApiVO) GetPrivateLabelingUrlOk() (*string, bool)`

GetPrivateLabelingUrlOk returns a tuple with the PrivateLabelingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingUrl

`func (o *MspUiInterfaceOpenApiVO) SetPrivateLabelingUrl(v string)`

SetPrivateLabelingUrl sets PrivateLabelingUrl field to given value.

### HasPrivateLabelingUrl

`func (o *MspUiInterfaceOpenApiVO) HasPrivateLabelingUrl() bool`

HasPrivateLabelingUrl returns a boolean if a field has been set.

### GetRefreshBtnEnable

`func (o *MspUiInterfaceOpenApiVO) GetRefreshBtnEnable() bool`

GetRefreshBtnEnable returns the RefreshBtnEnable field if non-nil, zero value otherwise.

### GetRefreshBtnEnableOk

`func (o *MspUiInterfaceOpenApiVO) GetRefreshBtnEnableOk() (*bool, bool)`

GetRefreshBtnEnableOk returns a tuple with the RefreshBtnEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshBtnEnable

`func (o *MspUiInterfaceOpenApiVO) SetRefreshBtnEnable(v bool)`

SetRefreshBtnEnable sets RefreshBtnEnable field to given value.


### GetRefreshRate

`func (o *MspUiInterfaceOpenApiVO) GetRefreshRate() int32`

GetRefreshRate returns the RefreshRate field if non-nil, zero value otherwise.

### GetRefreshRateOk

`func (o *MspUiInterfaceOpenApiVO) GetRefreshRateOk() (*int32, bool)`

GetRefreshRateOk returns a tuple with the RefreshRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshRate

`func (o *MspUiInterfaceOpenApiVO) SetRefreshRate(v int32)`

SetRefreshRate sets RefreshRate field to given value.


### GetShowPDevices

`func (o *MspUiInterfaceOpenApiVO) GetShowPDevices() bool`

GetShowPDevices returns the ShowPDevices field if non-nil, zero value otherwise.

### GetShowPDevicesOk

`func (o *MspUiInterfaceOpenApiVO) GetShowPDevicesOk() (*bool, bool)`

GetShowPDevicesOk returns a tuple with the ShowPDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPDevices

`func (o *MspUiInterfaceOpenApiVO) SetShowPDevices(v bool)`

SetShowPDevices sets ShowPDevices field to given value.

### HasShowPDevices

`func (o *MspUiInterfaceOpenApiVO) HasShowPDevices() bool`

HasShowPDevices returns a boolean if a field has been set.

### GetTheme

`func (o *MspUiInterfaceOpenApiVO) GetTheme() int32`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *MspUiInterfaceOpenApiVO) GetThemeOk() (*int32, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *MspUiInterfaceOpenApiVO) SetTheme(v int32)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *MspUiInterfaceOpenApiVO) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetUse24hour

`func (o *MspUiInterfaceOpenApiVO) GetUse24hour() bool`

GetUse24hour returns the Use24hour field if non-nil, zero value otherwise.

### GetUse24hourOk

`func (o *MspUiInterfaceOpenApiVO) GetUse24hourOk() (*bool, bool)`

GetUse24hourOk returns a tuple with the Use24hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse24hour

`func (o *MspUiInterfaceOpenApiVO) SetUse24hour(v bool)`

SetUse24hour sets Use24hour field to given value.


### GetWebsocketEnable

`func (o *MspUiInterfaceOpenApiVO) GetWebsocketEnable() bool`

GetWebsocketEnable returns the WebsocketEnable field if non-nil, zero value otherwise.

### GetWebsocketEnableOk

`func (o *MspUiInterfaceOpenApiVO) GetWebsocketEnableOk() (*bool, bool)`

GetWebsocketEnableOk returns a tuple with the WebsocketEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsocketEnable

`func (o *MspUiInterfaceOpenApiVO) SetWebsocketEnable(v bool)`

SetWebsocketEnable sets WebsocketEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


