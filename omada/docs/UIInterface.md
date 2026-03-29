# UIInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MacFormat** | Pointer to **int32** |  | [optional] 
**PrivateLabelingEnable** | Pointer to **bool** | Enable customer labeling of controller, this configuration applies to Omada Pro Controller only | [optional] 
**PrivateLabelingFileId** | Pointer to **string** | The file id of labeling image, this configuration applies to Omada Pro Controller only | [optional] 
**PrivateLabelingFileName** | Pointer to **string** | The file name of labeling image, this configuration applies to Omada Pro Controller only | [optional] 
**PrivateLabelingUrl** | Pointer to **string** | The redirection url of labeling image, this configuration applies to Omada Pro Controller only | [optional] 
**RefreshRate** | **int32** | Refresh rate should be a value as follows: 0: 15 seconds;1: 1 minute; 2: 2 minutes; 3: 5 minutes; 4: never refresh. | 
**TimeZone** | Pointer to **int32** | Time zone should be a value as follows: 0: Site&#39;s; 1: Browser&#39;s; 2: Controller&#39;s; 3: UTC | [optional] 

## Methods

### NewUIInterface

`func NewUIInterface(refreshRate int32, ) *UIInterface`

NewUIInterface instantiates a new UIInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUIInterfaceWithDefaults

`func NewUIInterfaceWithDefaults() *UIInterface`

NewUIInterfaceWithDefaults instantiates a new UIInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacFormat

`func (o *UIInterface) GetMacFormat() int32`

GetMacFormat returns the MacFormat field if non-nil, zero value otherwise.

### GetMacFormatOk

`func (o *UIInterface) GetMacFormatOk() (*int32, bool)`

GetMacFormatOk returns a tuple with the MacFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFormat

`func (o *UIInterface) SetMacFormat(v int32)`

SetMacFormat sets MacFormat field to given value.

### HasMacFormat

`func (o *UIInterface) HasMacFormat() bool`

HasMacFormat returns a boolean if a field has been set.

### GetPrivateLabelingEnable

`func (o *UIInterface) GetPrivateLabelingEnable() bool`

GetPrivateLabelingEnable returns the PrivateLabelingEnable field if non-nil, zero value otherwise.

### GetPrivateLabelingEnableOk

`func (o *UIInterface) GetPrivateLabelingEnableOk() (*bool, bool)`

GetPrivateLabelingEnableOk returns a tuple with the PrivateLabelingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingEnable

`func (o *UIInterface) SetPrivateLabelingEnable(v bool)`

SetPrivateLabelingEnable sets PrivateLabelingEnable field to given value.

### HasPrivateLabelingEnable

`func (o *UIInterface) HasPrivateLabelingEnable() bool`

HasPrivateLabelingEnable returns a boolean if a field has been set.

### GetPrivateLabelingFileId

`func (o *UIInterface) GetPrivateLabelingFileId() string`

GetPrivateLabelingFileId returns the PrivateLabelingFileId field if non-nil, zero value otherwise.

### GetPrivateLabelingFileIdOk

`func (o *UIInterface) GetPrivateLabelingFileIdOk() (*string, bool)`

GetPrivateLabelingFileIdOk returns a tuple with the PrivateLabelingFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingFileId

`func (o *UIInterface) SetPrivateLabelingFileId(v string)`

SetPrivateLabelingFileId sets PrivateLabelingFileId field to given value.

### HasPrivateLabelingFileId

`func (o *UIInterface) HasPrivateLabelingFileId() bool`

HasPrivateLabelingFileId returns a boolean if a field has been set.

### GetPrivateLabelingFileName

`func (o *UIInterface) GetPrivateLabelingFileName() string`

GetPrivateLabelingFileName returns the PrivateLabelingFileName field if non-nil, zero value otherwise.

### GetPrivateLabelingFileNameOk

`func (o *UIInterface) GetPrivateLabelingFileNameOk() (*string, bool)`

GetPrivateLabelingFileNameOk returns a tuple with the PrivateLabelingFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingFileName

`func (o *UIInterface) SetPrivateLabelingFileName(v string)`

SetPrivateLabelingFileName sets PrivateLabelingFileName field to given value.

### HasPrivateLabelingFileName

`func (o *UIInterface) HasPrivateLabelingFileName() bool`

HasPrivateLabelingFileName returns a boolean if a field has been set.

### GetPrivateLabelingUrl

`func (o *UIInterface) GetPrivateLabelingUrl() string`

GetPrivateLabelingUrl returns the PrivateLabelingUrl field if non-nil, zero value otherwise.

### GetPrivateLabelingUrlOk

`func (o *UIInterface) GetPrivateLabelingUrlOk() (*string, bool)`

GetPrivateLabelingUrlOk returns a tuple with the PrivateLabelingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingUrl

`func (o *UIInterface) SetPrivateLabelingUrl(v string)`

SetPrivateLabelingUrl sets PrivateLabelingUrl field to given value.

### HasPrivateLabelingUrl

`func (o *UIInterface) HasPrivateLabelingUrl() bool`

HasPrivateLabelingUrl returns a boolean if a field has been set.

### GetRefreshRate

`func (o *UIInterface) GetRefreshRate() int32`

GetRefreshRate returns the RefreshRate field if non-nil, zero value otherwise.

### GetRefreshRateOk

`func (o *UIInterface) GetRefreshRateOk() (*int32, bool)`

GetRefreshRateOk returns a tuple with the RefreshRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshRate

`func (o *UIInterface) SetRefreshRate(v int32)`

SetRefreshRate sets RefreshRate field to given value.


### GetTimeZone

`func (o *UIInterface) GetTimeZone() int32`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UIInterface) GetTimeZoneOk() (*int32, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UIInterface) SetTimeZone(v int32)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UIInterface) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


