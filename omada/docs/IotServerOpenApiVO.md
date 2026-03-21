# IotServerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** | This parameter becomes mandatory when the authentication method is set to \&quot;Use Token\&quot;.&lt;br/&gt;Note:The parameter [clientId] should be 1 ~ 128 characters. | [optional] 
**Authentication** | **int32** | The parameter [authentication] should be a value as follows:[0:Use Token]. | 
**BlePeriodicTelemetry** | **bool** | Whether to enable the BLE Periodic Telemetry. When disabled no periodic packets will be uploaded. | 
**ClientId** | Pointer to **string** | This parameter becomes mandatory when the authentication method is set to \&quot;Use Token\&quot;.&lt;br /&gt;Note:The parameter [clientId] should be 1 ~ 128 characters. | [optional] 
**CountOnly** | Pointer to **bool** | A switch that controls whether the AP device exclusively reports the count of IoT devices. | [optional] 
**DeviceClasses** | **[]int32** | Supports protocol-based filtering during IoT data reporting processes.&lt;br /&gt;The device class list should contain the value as follows: [0:minew; 1:iBeacon; 2:Eddystone]. | 
**Enable** | **bool** | Whether to enable the IoT Transport Stream setting. | 
**Filters** | Pointer to **map[string][]string** | The keys in the [filters] map represent the filter types, while the values correspond to the specific filtering criteria or values associated with each filter type.&lt;br /&gt;Note:&lt;br /&gt;Filter type &#x3D; 0, The Company Identifier must conform to a 4-digit or 6-digit hexadecimal encoding. It is only applicable to ibeacon devices&lt;br /&gt;Filter type &#x3D; 1, The Vendor should not exceed 255 bytes in length.&lt;br /&gt;Filter type &#x3D; 2, The Local Name should not exceed 120 bytes in length. It is only applicable to minew devices.&lt;br /&gt;Filter type &#x3D; 3, The Service UUID must conform to a 4-digit hexadecimal encoding. It is only applicable to minew and eddystone devices.&lt;br /&gt;Filter type &#x3D; 4, The Mac Oui must conform to a 6-digit hexadecimal encoding.&lt;br /&gt;Filter type &#x3D; 5, The iBeacon UUID must conform to a 32-digit hexadecimal encoding. It is only applicable to iBeacon devices.&lt;br /&gt;Filter type &#x3D; 6, The UID must conform to a 20-digit or 32-digit hexadecimal encoding. It is only applicable to eddystone devices.&lt;br /&gt;Filter type &#x3D; 7, The URL should not include a scheme. It is only applicable to eddystone devices.&lt;br /&gt; | [optional] 
**FiltersType** | Pointer to **[]int32** | User-defined settings to manage AP device filtering rules for IoT devices.&lt;br /&gt;The parameter [filtersType] should contain the value as follows:[0:Company Identifier; 1:Vendor; 2:Local Name; 3:Service UUID; 4:Mac Oui; 5:iBeacon UUID; 6:UID; 7:URL]. | [optional] 
**Id** | Pointer to **string** | The IoT Transport Stream entry ID. | [optional] 
**Name** | **string** | IoT Transport Stream setting name. | 
**RawData** | **bool** | Whether to enable the BLE Data Forwarding. When enabled, the AP directly reports the Bluetooth packet rawData to the server. | 
**ReportInterval** | Pointer to **int32** | Data reporting interval configuration for AP devices in IoT systems. | [optional] 
**RssiFormat** | **int32** | The signal strength reporting format currently supports five types: [0:Average; 1:Max; 2:Last; 3:Smooth; 4:Bulk]. | 
**ServerType** | **int32** | The server type should be a value as follows: [0: http]. | 
**ServerUrl** | **string** | If the service type is http, the server URL must start with http://. | 

## Methods

### NewIotServerOpenApiVO

`func NewIotServerOpenApiVO(authentication int32, blePeriodicTelemetry bool, deviceClasses []int32, enable bool, name string, rawData bool, rssiFormat int32, serverType int32, serverUrl string, ) *IotServerOpenApiVO`

NewIotServerOpenApiVO instantiates a new IotServerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIotServerOpenApiVOWithDefaults

`func NewIotServerOpenApiVOWithDefaults() *IotServerOpenApiVO`

NewIotServerOpenApiVOWithDefaults instantiates a new IotServerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *IotServerOpenApiVO) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *IotServerOpenApiVO) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *IotServerOpenApiVO) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *IotServerOpenApiVO) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetAuthentication

`func (o *IotServerOpenApiVO) GetAuthentication() int32`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *IotServerOpenApiVO) GetAuthenticationOk() (*int32, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *IotServerOpenApiVO) SetAuthentication(v int32)`

SetAuthentication sets Authentication field to given value.


### GetBlePeriodicTelemetry

`func (o *IotServerOpenApiVO) GetBlePeriodicTelemetry() bool`

GetBlePeriodicTelemetry returns the BlePeriodicTelemetry field if non-nil, zero value otherwise.

### GetBlePeriodicTelemetryOk

`func (o *IotServerOpenApiVO) GetBlePeriodicTelemetryOk() (*bool, bool)`

GetBlePeriodicTelemetryOk returns a tuple with the BlePeriodicTelemetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlePeriodicTelemetry

`func (o *IotServerOpenApiVO) SetBlePeriodicTelemetry(v bool)`

SetBlePeriodicTelemetry sets BlePeriodicTelemetry field to given value.


### GetClientId

`func (o *IotServerOpenApiVO) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IotServerOpenApiVO) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IotServerOpenApiVO) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IotServerOpenApiVO) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCountOnly

`func (o *IotServerOpenApiVO) GetCountOnly() bool`

GetCountOnly returns the CountOnly field if non-nil, zero value otherwise.

### GetCountOnlyOk

`func (o *IotServerOpenApiVO) GetCountOnlyOk() (*bool, bool)`

GetCountOnlyOk returns a tuple with the CountOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOnly

`func (o *IotServerOpenApiVO) SetCountOnly(v bool)`

SetCountOnly sets CountOnly field to given value.

### HasCountOnly

`func (o *IotServerOpenApiVO) HasCountOnly() bool`

HasCountOnly returns a boolean if a field has been set.

### GetDeviceClasses

`func (o *IotServerOpenApiVO) GetDeviceClasses() []int32`

GetDeviceClasses returns the DeviceClasses field if non-nil, zero value otherwise.

### GetDeviceClassesOk

`func (o *IotServerOpenApiVO) GetDeviceClassesOk() (*[]int32, bool)`

GetDeviceClassesOk returns a tuple with the DeviceClasses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClasses

`func (o *IotServerOpenApiVO) SetDeviceClasses(v []int32)`

SetDeviceClasses sets DeviceClasses field to given value.


### GetEnable

`func (o *IotServerOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IotServerOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IotServerOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetFilters

`func (o *IotServerOpenApiVO) GetFilters() map[string][]string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *IotServerOpenApiVO) GetFiltersOk() (*map[string][]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *IotServerOpenApiVO) SetFilters(v map[string][]string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *IotServerOpenApiVO) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetFiltersType

`func (o *IotServerOpenApiVO) GetFiltersType() []int32`

GetFiltersType returns the FiltersType field if non-nil, zero value otherwise.

### GetFiltersTypeOk

`func (o *IotServerOpenApiVO) GetFiltersTypeOk() (*[]int32, bool)`

GetFiltersTypeOk returns a tuple with the FiltersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersType

`func (o *IotServerOpenApiVO) SetFiltersType(v []int32)`

SetFiltersType sets FiltersType field to given value.

### HasFiltersType

`func (o *IotServerOpenApiVO) HasFiltersType() bool`

HasFiltersType returns a boolean if a field has been set.

### GetId

`func (o *IotServerOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IotServerOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IotServerOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IotServerOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IotServerOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IotServerOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IotServerOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRawData

`func (o *IotServerOpenApiVO) GetRawData() bool`

GetRawData returns the RawData field if non-nil, zero value otherwise.

### GetRawDataOk

`func (o *IotServerOpenApiVO) GetRawDataOk() (*bool, bool)`

GetRawDataOk returns a tuple with the RawData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawData

`func (o *IotServerOpenApiVO) SetRawData(v bool)`

SetRawData sets RawData field to given value.


### GetReportInterval

`func (o *IotServerOpenApiVO) GetReportInterval() int32`

GetReportInterval returns the ReportInterval field if non-nil, zero value otherwise.

### GetReportIntervalOk

`func (o *IotServerOpenApiVO) GetReportIntervalOk() (*int32, bool)`

GetReportIntervalOk returns a tuple with the ReportInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportInterval

`func (o *IotServerOpenApiVO) SetReportInterval(v int32)`

SetReportInterval sets ReportInterval field to given value.

### HasReportInterval

`func (o *IotServerOpenApiVO) HasReportInterval() bool`

HasReportInterval returns a boolean if a field has been set.

### GetRssiFormat

`func (o *IotServerOpenApiVO) GetRssiFormat() int32`

GetRssiFormat returns the RssiFormat field if non-nil, zero value otherwise.

### GetRssiFormatOk

`func (o *IotServerOpenApiVO) GetRssiFormatOk() (*int32, bool)`

GetRssiFormatOk returns a tuple with the RssiFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiFormat

`func (o *IotServerOpenApiVO) SetRssiFormat(v int32)`

SetRssiFormat sets RssiFormat field to given value.


### GetServerType

`func (o *IotServerOpenApiVO) GetServerType() int32`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *IotServerOpenApiVO) GetServerTypeOk() (*int32, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *IotServerOpenApiVO) SetServerType(v int32)`

SetServerType sets ServerType field to given value.


### GetServerUrl

`func (o *IotServerOpenApiVO) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *IotServerOpenApiVO) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *IotServerOpenApiVO) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


