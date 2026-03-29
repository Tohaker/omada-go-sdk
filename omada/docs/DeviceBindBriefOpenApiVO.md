# DeviceBindBriefOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | **string** | The mac address of device, like AA-BB-CC-DD-EE-FF. | 
**SiteId** | **string** | The ID of target site. | 

## Methods

### NewDeviceBindBriefOpenApiVO

`func NewDeviceBindBriefOpenApiVO(mac string, siteId string, ) *DeviceBindBriefOpenApiVO`

NewDeviceBindBriefOpenApiVO instantiates a new DeviceBindBriefOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceBindBriefOpenApiVOWithDefaults

`func NewDeviceBindBriefOpenApiVOWithDefaults() *DeviceBindBriefOpenApiVO`

NewDeviceBindBriefOpenApiVOWithDefaults instantiates a new DeviceBindBriefOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *DeviceBindBriefOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceBindBriefOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceBindBriefOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetSiteId

`func (o *DeviceBindBriefOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DeviceBindBriefOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DeviceBindBriefOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


