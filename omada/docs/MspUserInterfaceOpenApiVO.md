# MspUserInterfaceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateLabelingEnable** | Pointer to **bool** | Enable private label. This field applies to the Omada Pro Controller only. | [optional] 
**PrivateLabelingFileId** | Pointer to **string** | The file id of private label picture. This field applies to the Omada Pro Controller only. | [optional] 
**PrivateLabelingFileName** | Pointer to **string** | The file name of private label picture. This field applies to the Omada Pro Controller only. | [optional] 
**PrivateLabelingUrl** | Pointer to **string** | The url of private label. This field applies to the Omada Pro Controller only. | [optional] 
**RefreshRate** | **int32** | It should be a value as follows: 0: 15 seconds; 1: 1 minute; 2: 2 minutes; 3: 5 minutes; 4: never refresh | 

## Methods

### NewMspUserInterfaceOpenApiVO

`func NewMspUserInterfaceOpenApiVO(refreshRate int32, ) *MspUserInterfaceOpenApiVO`

NewMspUserInterfaceOpenApiVO instantiates a new MspUserInterfaceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspUserInterfaceOpenApiVOWithDefaults

`func NewMspUserInterfaceOpenApiVOWithDefaults() *MspUserInterfaceOpenApiVO`

NewMspUserInterfaceOpenApiVOWithDefaults instantiates a new MspUserInterfaceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateLabelingEnable

`func (o *MspUserInterfaceOpenApiVO) GetPrivateLabelingEnable() bool`

GetPrivateLabelingEnable returns the PrivateLabelingEnable field if non-nil, zero value otherwise.

### GetPrivateLabelingEnableOk

`func (o *MspUserInterfaceOpenApiVO) GetPrivateLabelingEnableOk() (*bool, bool)`

GetPrivateLabelingEnableOk returns a tuple with the PrivateLabelingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingEnable

`func (o *MspUserInterfaceOpenApiVO) SetPrivateLabelingEnable(v bool)`

SetPrivateLabelingEnable sets PrivateLabelingEnable field to given value.

### HasPrivateLabelingEnable

`func (o *MspUserInterfaceOpenApiVO) HasPrivateLabelingEnable() bool`

HasPrivateLabelingEnable returns a boolean if a field has been set.

### GetPrivateLabelingFileId

`func (o *MspUserInterfaceOpenApiVO) GetPrivateLabelingFileId() string`

GetPrivateLabelingFileId returns the PrivateLabelingFileId field if non-nil, zero value otherwise.

### GetPrivateLabelingFileIdOk

`func (o *MspUserInterfaceOpenApiVO) GetPrivateLabelingFileIdOk() (*string, bool)`

GetPrivateLabelingFileIdOk returns a tuple with the PrivateLabelingFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingFileId

`func (o *MspUserInterfaceOpenApiVO) SetPrivateLabelingFileId(v string)`

SetPrivateLabelingFileId sets PrivateLabelingFileId field to given value.

### HasPrivateLabelingFileId

`func (o *MspUserInterfaceOpenApiVO) HasPrivateLabelingFileId() bool`

HasPrivateLabelingFileId returns a boolean if a field has been set.

### GetPrivateLabelingFileName

`func (o *MspUserInterfaceOpenApiVO) GetPrivateLabelingFileName() string`

GetPrivateLabelingFileName returns the PrivateLabelingFileName field if non-nil, zero value otherwise.

### GetPrivateLabelingFileNameOk

`func (o *MspUserInterfaceOpenApiVO) GetPrivateLabelingFileNameOk() (*string, bool)`

GetPrivateLabelingFileNameOk returns a tuple with the PrivateLabelingFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingFileName

`func (o *MspUserInterfaceOpenApiVO) SetPrivateLabelingFileName(v string)`

SetPrivateLabelingFileName sets PrivateLabelingFileName field to given value.

### HasPrivateLabelingFileName

`func (o *MspUserInterfaceOpenApiVO) HasPrivateLabelingFileName() bool`

HasPrivateLabelingFileName returns a boolean if a field has been set.

### GetPrivateLabelingUrl

`func (o *MspUserInterfaceOpenApiVO) GetPrivateLabelingUrl() string`

GetPrivateLabelingUrl returns the PrivateLabelingUrl field if non-nil, zero value otherwise.

### GetPrivateLabelingUrlOk

`func (o *MspUserInterfaceOpenApiVO) GetPrivateLabelingUrlOk() (*string, bool)`

GetPrivateLabelingUrlOk returns a tuple with the PrivateLabelingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateLabelingUrl

`func (o *MspUserInterfaceOpenApiVO) SetPrivateLabelingUrl(v string)`

SetPrivateLabelingUrl sets PrivateLabelingUrl field to given value.

### HasPrivateLabelingUrl

`func (o *MspUserInterfaceOpenApiVO) HasPrivateLabelingUrl() bool`

HasPrivateLabelingUrl returns a boolean if a field has been set.

### GetRefreshRate

`func (o *MspUserInterfaceOpenApiVO) GetRefreshRate() int32`

GetRefreshRate returns the RefreshRate field if non-nil, zero value otherwise.

### GetRefreshRateOk

`func (o *MspUserInterfaceOpenApiVO) GetRefreshRateOk() (*int32, bool)`

GetRefreshRateOk returns a tuple with the RefreshRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshRate

`func (o *MspUserInterfaceOpenApiVO) SetRefreshRate(v int32)`

SetRefreshRate sets RefreshRate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


