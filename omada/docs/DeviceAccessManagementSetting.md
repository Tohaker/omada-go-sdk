# DeviceAccessManagementSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDiscovery** | **bool** | Control device app discovery, bind with device web https access, if appDiscvoery is enabled, webControlHttps must be enabled. | 
**WebControlHttp** | **bool** | Control device web http accessing. | 
**WebControlHttps** | **bool** | Control device web https accessing. | 

## Methods

### NewDeviceAccessManagementSetting

`func NewDeviceAccessManagementSetting(appDiscovery bool, webControlHttp bool, webControlHttps bool, ) *DeviceAccessManagementSetting`

NewDeviceAccessManagementSetting instantiates a new DeviceAccessManagementSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAccessManagementSettingWithDefaults

`func NewDeviceAccessManagementSettingWithDefaults() *DeviceAccessManagementSetting`

NewDeviceAccessManagementSettingWithDefaults instantiates a new DeviceAccessManagementSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDiscovery

`func (o *DeviceAccessManagementSetting) GetAppDiscovery() bool`

GetAppDiscovery returns the AppDiscovery field if non-nil, zero value otherwise.

### GetAppDiscoveryOk

`func (o *DeviceAccessManagementSetting) GetAppDiscoveryOk() (*bool, bool)`

GetAppDiscoveryOk returns a tuple with the AppDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDiscovery

`func (o *DeviceAccessManagementSetting) SetAppDiscovery(v bool)`

SetAppDiscovery sets AppDiscovery field to given value.


### GetWebControlHttp

`func (o *DeviceAccessManagementSetting) GetWebControlHttp() bool`

GetWebControlHttp returns the WebControlHttp field if non-nil, zero value otherwise.

### GetWebControlHttpOk

`func (o *DeviceAccessManagementSetting) GetWebControlHttpOk() (*bool, bool)`

GetWebControlHttpOk returns a tuple with the WebControlHttp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebControlHttp

`func (o *DeviceAccessManagementSetting) SetWebControlHttp(v bool)`

SetWebControlHttp sets WebControlHttp field to given value.


### GetWebControlHttps

`func (o *DeviceAccessManagementSetting) GetWebControlHttps() bool`

GetWebControlHttps returns the WebControlHttps field if non-nil, zero value otherwise.

### GetWebControlHttpsOk

`func (o *DeviceAccessManagementSetting) GetWebControlHttpsOk() (*bool, bool)`

GetWebControlHttpsOk returns a tuple with the WebControlHttps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebControlHttps

`func (o *DeviceAccessManagementSetting) SetWebControlHttps(v bool)`

SetWebControlHttps sets WebControlHttps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


