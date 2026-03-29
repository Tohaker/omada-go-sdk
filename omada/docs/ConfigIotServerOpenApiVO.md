# ConfigIotServerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | This parameter becomes mandatory when the authentication method is set to \&quot;Use Token\&quot;.&lt;br/&gt;Note:The parameter [clientId] should be 1 ~ 128 characters. | 
**Authentication** | **int32** | The parameter [authentication] should be a value as follows:[0:Use Token]. | 
**BlePeriodicTelemetry** | **bool** | Whether to enable the BLE Periodic Telemetry. When disabled no periodic packets will be uploaded. | 
**ClientId** | **string** | This parameter becomes mandatory when the authentication method is set to \&quot;Use Token\&quot;.&lt;br /&gt;Note:The parameter [clientId] should be 1 ~ 128 characters. | 
**CountOnly** | Pointer to **bool** | A switch that controls whether the AP device exclusively reports the count of IoT devices, and it cannot be null when blePeriodicTelemetry is enabled. | [optional] 
**DeviceClasses** | **[]int32** | Supports protocol-based filtering during IoT data reporting processes.&lt;br /&gt;The device class list should contain the value as follows: [0:minew; 1:iBeacon; 2:Eddystone]. | 
**Enable** | **bool** | Whether to enable the IoT Transport Stream setting. | 
**Filters** | Pointer to **map[string][]string** | The keys in the [filters] map represent the filter types, while the values correspond to the specific filtering criteria or values associated with each filter type.&lt;br /&gt;Note:&lt;br /&gt;Filter type &#x3D; 0, The Company Identifier must conform to a 4-digit or 6-digit hexadecimal encoding. It is only applicable to ibeacon devices&lt;br /&gt;Filter type &#x3D; 1, The Vendor should not exceed 255 bytes in length.&lt;br /&gt;Filter type &#x3D; 2, The Local Name should not exceed 120 bytes in length. It is only applicable to minew devices.&lt;br /&gt;Filter type &#x3D; 3, The Service UUID must conform to a 4-digit hexadecimal encoding. It is only applicable to minew and eddystone devices.&lt;br /&gt;Filter type &#x3D; 4, The Mac Oui must conform to a 6-digit hexadecimal encoding.&lt;br /&gt;Filter type &#x3D; 5, The iBeacon UUID must conform to a 32-digit hexadecimal encoding. It is only applicable to iBeacon devices.&lt;br /&gt;Filter type &#x3D; 6, The UID must conform to a 20-digit or 32-digit hexadecimal encoding. It is only applicable to eddystone devices.&lt;br /&gt;Filter type &#x3D; 7, The URL should not include a scheme. It is only applicable to eddystone devices.&lt;br /&gt; | [optional] 
**FiltersType** | Pointer to **[]int32** | User-defined settings to manage AP device filtering rules for IoT devices.&lt;br /&gt;The parameter [filtersType] should contain the value as follows:[0:Company Identifier; 1:Vendor; 2:Local Name; 3:Service UUID; 4:Mac Oui; 5:iBeacon UUID; 6:UID; 7:URL]. | [optional] 
**Name** | **string** | IoT Transport Stream setting name. | 
**RawData** | **bool** | Whether to enable the BLE Data Forwarding. When enabled, the AP directly reports the Bluetooth packet rawData to the server. | 
**ReportInterval** | Pointer to **int32** | Data reporting interval configuration for AP devices in IoT systems.The parameter [reportInterval] should be within the range of 1–3600 in seconds, and it cannot be null when blePeriodicTelemetry is enabled. | [optional] 
**RssiFormat** | **int32** | The signal strength reporting format currently supports five types: [0:Average; 1:Max; 2:Last; 3:Smooth; 4:Bulk]. | 
**ServerType** | **int32** | The server type should be a value as follows: [0: http]. | 
**ServerUrl** | **string** | If the service type is http, the server URL must start with http://. | 

## Methods

### NewConfigIotServerOpenApiVO

`func NewConfigIotServerOpenApiVO(accessToken string, authentication int32, blePeriodicTelemetry bool, clientId string, deviceClasses []int32, enable bool, name string, rawData bool, rssiFormat int32, serverType int32, serverUrl string, ) *ConfigIotServerOpenApiVO`

NewConfigIotServerOpenApiVO instantiates a new ConfigIotServerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigIotServerOpenApiVOWithDefaults

`func NewConfigIotServerOpenApiVOWithDefaults() *ConfigIotServerOpenApiVO`

NewConfigIotServerOpenApiVOWithDefaults instantiates a new ConfigIotServerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *ConfigIotServerOpenApiVO) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ConfigIotServerOpenApiVO) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ConfigIotServerOpenApiVO) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAuthentication

`func (o *ConfigIotServerOpenApiVO) GetAuthentication() int32`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *ConfigIotServerOpenApiVO) GetAuthenticationOk() (*int32, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *ConfigIotServerOpenApiVO) SetAuthentication(v int32)`

SetAuthentication sets Authentication field to given value.


### GetBlePeriodicTelemetry

`func (o *ConfigIotServerOpenApiVO) GetBlePeriodicTelemetry() bool`

GetBlePeriodicTelemetry returns the BlePeriodicTelemetry field if non-nil, zero value otherwise.

### GetBlePeriodicTelemetryOk

`func (o *ConfigIotServerOpenApiVO) GetBlePeriodicTelemetryOk() (*bool, bool)`

GetBlePeriodicTelemetryOk returns a tuple with the BlePeriodicTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlePeriodicTelemetry

`func (o *ConfigIotServerOpenApiVO) SetBlePeriodicTelemetry(v bool)`

SetBlePeriodicTelemetry sets BlePeriodicTelemetry field to given value.


### GetClientId

`func (o *ConfigIotServerOpenApiVO) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ConfigIotServerOpenApiVO) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ConfigIotServerOpenApiVO) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCountOnly

`func (o *ConfigIotServerOpenApiVO) GetCountOnly() bool`

GetCountOnly returns the CountOnly field if non-nil, zero value otherwise.

### GetCountOnlyOk

`func (o *ConfigIotServerOpenApiVO) GetCountOnlyOk() (*bool, bool)`

GetCountOnlyOk returns a tuple with the CountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOnly

`func (o *ConfigIotServerOpenApiVO) SetCountOnly(v bool)`

SetCountOnly sets CountOnly field to given value.

### HasCountOnly

`func (o *ConfigIotServerOpenApiVO) HasCountOnly() bool`

HasCountOnly returns a boolean if a field has been set.

### GetDeviceClasses

`func (o *ConfigIotServerOpenApiVO) GetDeviceClasses() []int32`

GetDeviceClasses returns the DeviceClasses field if non-nil, zero value otherwise.

### GetDeviceClassesOk

`func (o *ConfigIotServerOpenApiVO) GetDeviceClassesOk() (*[]int32, bool)`

GetDeviceClassesOk returns a tuple with the DeviceClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClasses

`func (o *ConfigIotServerOpenApiVO) SetDeviceClasses(v []int32)`

SetDeviceClasses sets DeviceClasses field to given value.


### GetEnable

`func (o *ConfigIotServerOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ConfigIotServerOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ConfigIotServerOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetFilters

`func (o *ConfigIotServerOpenApiVO) GetFilters() map[string][]string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ConfigIotServerOpenApiVO) GetFiltersOk() (*map[string][]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ConfigIotServerOpenApiVO) SetFilters(v map[string][]string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ConfigIotServerOpenApiVO) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFiltersType

`func (o *ConfigIotServerOpenApiVO) GetFiltersType() []int32`

GetFiltersType returns the FiltersType field if non-nil, zero value otherwise.

### GetFiltersTypeOk

`func (o *ConfigIotServerOpenApiVO) GetFiltersTypeOk() (*[]int32, bool)`

GetFiltersTypeOk returns a tuple with the FiltersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersType

`func (o *ConfigIotServerOpenApiVO) SetFiltersType(v []int32)`

SetFiltersType sets FiltersType field to given value.

### HasFiltersType

`func (o *ConfigIotServerOpenApiVO) HasFiltersType() bool`

HasFiltersType returns a boolean if a field has been set.

### GetName

`func (o *ConfigIotServerOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigIotServerOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigIotServerOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRawData

`func (o *ConfigIotServerOpenApiVO) GetRawData() bool`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *ConfigIotServerOpenApiVO) GetRawDataOk() (*bool, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *ConfigIotServerOpenApiVO) SetRawData(v bool)`

SetRawData sets RawData field to given value.


### GetReportInterval

`func (o *ConfigIotServerOpenApiVO) GetReportInterval() int32`

GetReportInterval returns the ReportInterval field if non-nil, zero value otherwise.

### GetReportIntervalOk

`func (o *ConfigIotServerOpenApiVO) GetReportIntervalOk() (*int32, bool)`

GetReportIntervalOk returns a tuple with the ReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInterval

`func (o *ConfigIotServerOpenApiVO) SetReportInterval(v int32)`

SetReportInterval sets ReportInterval field to given value.

### HasReportInterval

`func (o *ConfigIotServerOpenApiVO) HasReportInterval() bool`

HasReportInterval returns a boolean if a field has been set.

### GetRssiFormat

`func (o *ConfigIotServerOpenApiVO) GetRssiFormat() int32`

GetRssiFormat returns the RssiFormat field if non-nil, zero value otherwise.

### GetRssiFormatOk

`func (o *ConfigIotServerOpenApiVO) GetRssiFormatOk() (*int32, bool)`

GetRssiFormatOk returns a tuple with the RssiFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiFormat

`func (o *ConfigIotServerOpenApiVO) SetRssiFormat(v int32)`

SetRssiFormat sets RssiFormat field to given value.


### GetServerType

`func (o *ConfigIotServerOpenApiVO) GetServerType() int32`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ConfigIotServerOpenApiVO) GetServerTypeOk() (*int32, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ConfigIotServerOpenApiVO) SetServerType(v int32)`

SetServerType sets ServerType field to given value.


### GetServerUrl

`func (o *ConfigIotServerOpenApiVO) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *ConfigIotServerOpenApiVO) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *ConfigIotServerOpenApiVO) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


