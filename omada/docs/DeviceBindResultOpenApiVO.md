# DeviceBindResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | The error code for failed operation. | [optional] 
**Mac** | Pointer to **string** | The mac address of device, like AA-BB-CC-DD-EE-FF. | [optional] 
**Msg** | Pointer to **string** | The message for failed operation. | [optional] 
**SiteId** | Pointer to **string** | The ID of target site. | [optional] 

## Methods

### NewDeviceBindResultOpenApiVO

`func NewDeviceBindResultOpenApiVO() *DeviceBindResultOpenApiVO`

NewDeviceBindResultOpenApiVO instantiates a new DeviceBindResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceBindResultOpenApiVOWithDefaults

`func NewDeviceBindResultOpenApiVOWithDefaults() *DeviceBindResultOpenApiVO`

NewDeviceBindResultOpenApiVOWithDefaults instantiates a new DeviceBindResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *DeviceBindResultOpenApiVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *DeviceBindResultOpenApiVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *DeviceBindResultOpenApiVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *DeviceBindResultOpenApiVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMac

`func (o *DeviceBindResultOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceBindResultOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceBindResultOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DeviceBindResultOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMsg

`func (o *DeviceBindResultOpenApiVO) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *DeviceBindResultOpenApiVO) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *DeviceBindResultOpenApiVO) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *DeviceBindResultOpenApiVO) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetSiteId

`func (o *DeviceBindResultOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DeviceBindResultOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DeviceBindResultOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DeviceBindResultOpenApiVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


