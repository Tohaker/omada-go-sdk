# IotServerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokenIotServer** | Pointer to **string** | This parameter becomes mandatory when the authentication method is set to \&quot;Use Token\&quot;.&lt;br/&gt;Note:The parameter [clientId] should be 1 ~ 128 characters. | [optional] 
**AuthenticationIotServer** | **int32** | The parameter [authentication] should be a value as follows:[0:Use Token]. | 
**BlePeriodicTelemetryIotServer** | **bool** | Whether to enable the BLE Periodic Telemetry. When disabled no periodic packets will be uploaded. | 
**ClientIdIotServer** | Pointer to **string** | This parameter becomes mandatory when the authentication method is set to \&quot;Use Token\&quot;.&lt;br /&gt;Note:The parameter [clientId] should be 1 ~ 128 characters. | [optional] 
**CountOnlyIotServer** | Pointer to **bool** | A switch that controls whether the AP device exclusively reports the count of IoT devices. | [optional] 
**DeviceClassesIotServer** | **[]int32** | Supports protocol-based filtering during IoT data reporting processes.&lt;br /&gt;The device class list should contain the value as follows: [0:minew; 1:iBeacon; 2:Eddystone]. | 
**EnableIotServer** | **bool** | Whether to enable the IoT Transport Stream setting. | 
**FiltersTypeIotServer** | Pointer to **[]int32** | User-defined settings to manage AP device filtering rules for IoT devices.&lt;br /&gt;The parameter [filtersType] should contain the value as follows:[0:Company Identifier; 1:Vendor; 2:Local Name; 3:Service UUID; 4:Mac Oui; 5:iBeacon UUID; 6:UID; 7:URL]. | [optional] 
**FiltersIotServer** | Pointer to **map[string][]string** | The keys in the [filters] map represent the filter types, while the values correspond to the specific filtering criteria or values associated with each filter type.&lt;br /&gt;Note:&lt;br /&gt;Filter type &#x3D; 0, The Company Identifier must conform to a 4-digit or 6-digit hexadecimal encoding. It is only applicable to ibeacon devices&lt;br /&gt;Filter type &#x3D; 1, The Vendor should not exceed 255 bytes in length.&lt;br /&gt;Filter type &#x3D; 2, The Local Name should not exceed 120 bytes in length. It is only applicable to minew devices.&lt;br /&gt;Filter type &#x3D; 3, The Service UUID must conform to a 4-digit hexadecimal encoding. It is only applicable to minew and eddystone devices.&lt;br /&gt;Filter type &#x3D; 4, The Mac Oui must conform to a 6-digit hexadecimal encoding.&lt;br /&gt;Filter type &#x3D; 5, The iBeacon UUID must conform to a 32-digit hexadecimal encoding. It is only applicable to iBeacon devices.&lt;br /&gt;Filter type &#x3D; 6, The UID must conform to a 20-digit or 32-digit hexadecimal encoding. It is only applicable to eddystone devices.&lt;br /&gt;Filter type &#x3D; 7, The URL should not include a scheme. It is only applicable to eddystone devices.&lt;br /&gt; | [optional] 
**Id** | Pointer to **string** | The IoT Transport Stream entry ID. | [optional] 
**NameIotServer** | **string** | IoT Transport Stream setting name. | 
**RawDataIotServer** | **bool** | Whether to enable the BLE Data Forwarding. When enabled, the AP directly reports the Bluetooth packet rawData to the server. | 
**ReportIntervalIotServer** | Pointer to **int32** | Data reporting interval configuration for AP devices in IoT systems. | [optional] 
**RssiFormatIotServer** | **int32** | The signal strength reporting format currently supports five types: [0:Average; 1:Max; 2:Last; 3:Smooth; 4:Bulk]. | 
**ServerTypeIotServer** | **int32** | The server type should be a value as follows: [0: http]. | 
**ServerUrlIotServer** | **string** | If the service type is http, the server URL must start with http://. | 

## Methods

### NewIotServerOpenApiVO

`func NewIotServerOpenApiVO(authenticationIotServer int32, blePeriodicTelemetryIotServer bool, deviceClassesIotServer []int32, enableIotServer bool, nameIotServer string, rawDataIotServer bool, rssiFormatIotServer int32, serverTypeIotServer int32, serverUrlIotServer string, ) *IotServerOpenApiVO`

NewIotServerOpenApiVO instantiates a new IotServerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIotServerOpenApiVOWithDefaults

`func NewIotServerOpenApiVOWithDefaults() *IotServerOpenApiVO`

NewIotServerOpenApiVOWithDefaults instantiates a new IotServerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokenIotServer

`func (o *IotServerOpenApiVO) GetAccessTokenIotServer() string`

GetAccessTokenIotServer returns the AccessTokenIotServer field if non-nil, zero value otherwise.

### GetAccessTokenIotServerOk

`func (o *IotServerOpenApiVO) GetAccessTokenIotServerOk() (*string, bool)`

GetAccessTokenIotServerOk returns a tuple with the AccessTokenIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenIotServer

`func (o *IotServerOpenApiVO) SetAccessTokenIotServer(v string)`

SetAccessTokenIotServer sets AccessTokenIotServer field to given value.

### HasAccessTokenIotServer

`func (o *IotServerOpenApiVO) HasAccessTokenIotServer() bool`

HasAccessTokenIotServer returns a boolean if a field has been set.

### GetAuthenticationIotServer

`func (o *IotServerOpenApiVO) GetAuthenticationIotServer() int32`

GetAuthenticationIotServer returns the AuthenticationIotServer field if non-nil, zero value otherwise.

### GetAuthenticationIotServerOk

`func (o *IotServerOpenApiVO) GetAuthenticationIotServerOk() (*int32, bool)`

GetAuthenticationIotServerOk returns a tuple with the AuthenticationIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationIotServer

`func (o *IotServerOpenApiVO) SetAuthenticationIotServer(v int32)`

SetAuthenticationIotServer sets AuthenticationIotServer field to given value.


### GetBlePeriodicTelemetryIotServer

`func (o *IotServerOpenApiVO) GetBlePeriodicTelemetryIotServer() bool`

GetBlePeriodicTelemetryIotServer returns the BlePeriodicTelemetryIotServer field if non-nil, zero value otherwise.

### GetBlePeriodicTelemetryIotServerOk

`func (o *IotServerOpenApiVO) GetBlePeriodicTelemetryIotServerOk() (*bool, bool)`

GetBlePeriodicTelemetryIotServerOk returns a tuple with the BlePeriodicTelemetryIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlePeriodicTelemetryIotServer

`func (o *IotServerOpenApiVO) SetBlePeriodicTelemetryIotServer(v bool)`

SetBlePeriodicTelemetryIotServer sets BlePeriodicTelemetryIotServer field to given value.


### GetClientIdIotServer

`func (o *IotServerOpenApiVO) GetClientIdIotServer() string`

GetClientIdIotServer returns the ClientIdIotServer field if non-nil, zero value otherwise.

### GetClientIdIotServerOk

`func (o *IotServerOpenApiVO) GetClientIdIotServerOk() (*string, bool)`

GetClientIdIotServerOk returns a tuple with the ClientIdIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIdIotServer

`func (o *IotServerOpenApiVO) SetClientIdIotServer(v string)`

SetClientIdIotServer sets ClientIdIotServer field to given value.

### HasClientIdIotServer

`func (o *IotServerOpenApiVO) HasClientIdIotServer() bool`

HasClientIdIotServer returns a boolean if a field has been set.

### GetCountOnlyIotServer

`func (o *IotServerOpenApiVO) GetCountOnlyIotServer() bool`

GetCountOnlyIotServer returns the CountOnlyIotServer field if non-nil, zero value otherwise.

### GetCountOnlyIotServerOk

`func (o *IotServerOpenApiVO) GetCountOnlyIotServerOk() (*bool, bool)`

GetCountOnlyIotServerOk returns a tuple with the CountOnlyIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountOnlyIotServer

`func (o *IotServerOpenApiVO) SetCountOnlyIotServer(v bool)`

SetCountOnlyIotServer sets CountOnlyIotServer field to given value.

### HasCountOnlyIotServer

`func (o *IotServerOpenApiVO) HasCountOnlyIotServer() bool`

HasCountOnlyIotServer returns a boolean if a field has been set.

### GetDeviceClassesIotServer

`func (o *IotServerOpenApiVO) GetDeviceClassesIotServer() []int32`

GetDeviceClassesIotServer returns the DeviceClassesIotServer field if non-nil, zero value otherwise.

### GetDeviceClassesIotServerOk

`func (o *IotServerOpenApiVO) GetDeviceClassesIotServerOk() (*[]int32, bool)`

GetDeviceClassesIotServerOk returns a tuple with the DeviceClassesIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassesIotServer

`func (o *IotServerOpenApiVO) SetDeviceClassesIotServer(v []int32)`

SetDeviceClassesIotServer sets DeviceClassesIotServer field to given value.


### GetEnableIotServer

`func (o *IotServerOpenApiVO) GetEnableIotServer() bool`

GetEnableIotServer returns the EnableIotServer field if non-nil, zero value otherwise.

### GetEnableIotServerOk

`func (o *IotServerOpenApiVO) GetEnableIotServerOk() (*bool, bool)`

GetEnableIotServerOk returns a tuple with the EnableIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableIotServer

`func (o *IotServerOpenApiVO) SetEnableIotServer(v bool)`

SetEnableIotServer sets EnableIotServer field to given value.


### GetFiltersTypeIotServer

`func (o *IotServerOpenApiVO) GetFiltersTypeIotServer() []int32`

GetFiltersTypeIotServer returns the FiltersTypeIotServer field if non-nil, zero value otherwise.

### GetFiltersTypeIotServerOk

`func (o *IotServerOpenApiVO) GetFiltersTypeIotServerOk() (*[]int32, bool)`

GetFiltersTypeIotServerOk returns a tuple with the FiltersTypeIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersTypeIotServer

`func (o *IotServerOpenApiVO) SetFiltersTypeIotServer(v []int32)`

SetFiltersTypeIotServer sets FiltersTypeIotServer field to given value.

### HasFiltersTypeIotServer

`func (o *IotServerOpenApiVO) HasFiltersTypeIotServer() bool`

HasFiltersTypeIotServer returns a boolean if a field has been set.

### GetFiltersIotServer

`func (o *IotServerOpenApiVO) GetFiltersIotServer() map[string][]string`

GetFiltersIotServer returns the FiltersIotServer field if non-nil, zero value otherwise.

### GetFiltersIotServerOk

`func (o *IotServerOpenApiVO) GetFiltersIotServerOk() (*map[string][]string, bool)`

GetFiltersIotServerOk returns a tuple with the FiltersIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersIotServer

`func (o *IotServerOpenApiVO) SetFiltersIotServer(v map[string][]string)`

SetFiltersIotServer sets FiltersIotServer field to given value.

### HasFiltersIotServer

`func (o *IotServerOpenApiVO) HasFiltersIotServer() bool`

HasFiltersIotServer returns a boolean if a field has been set.

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

### GetNameIotServer

`func (o *IotServerOpenApiVO) GetNameIotServer() string`

GetNameIotServer returns the NameIotServer field if non-nil, zero value otherwise.

### GetNameIotServerOk

`func (o *IotServerOpenApiVO) GetNameIotServerOk() (*string, bool)`

GetNameIotServerOk returns a tuple with the NameIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameIotServer

`func (o *IotServerOpenApiVO) SetNameIotServer(v string)`

SetNameIotServer sets NameIotServer field to given value.


### GetRawDataIotServer

`func (o *IotServerOpenApiVO) GetRawDataIotServer() bool`

GetRawDataIotServer returns the RawDataIotServer field if non-nil, zero value otherwise.

### GetRawDataIotServerOk

`func (o *IotServerOpenApiVO) GetRawDataIotServerOk() (*bool, bool)`

GetRawDataIotServerOk returns a tuple with the RawDataIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawDataIotServer

`func (o *IotServerOpenApiVO) SetRawDataIotServer(v bool)`

SetRawDataIotServer sets RawDataIotServer field to given value.


### GetReportIntervalIotServer

`func (o *IotServerOpenApiVO) GetReportIntervalIotServer() int32`

GetReportIntervalIotServer returns the ReportIntervalIotServer field if non-nil, zero value otherwise.

### GetReportIntervalIotServerOk

`func (o *IotServerOpenApiVO) GetReportIntervalIotServerOk() (*int32, bool)`

GetReportIntervalIotServerOk returns a tuple with the ReportIntervalIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportIntervalIotServer

`func (o *IotServerOpenApiVO) SetReportIntervalIotServer(v int32)`

SetReportIntervalIotServer sets ReportIntervalIotServer field to given value.

### HasReportIntervalIotServer

`func (o *IotServerOpenApiVO) HasReportIntervalIotServer() bool`

HasReportIntervalIotServer returns a boolean if a field has been set.

### GetRssiFormatIotServer

`func (o *IotServerOpenApiVO) GetRssiFormatIotServer() int32`

GetRssiFormatIotServer returns the RssiFormatIotServer field if non-nil, zero value otherwise.

### GetRssiFormatIotServerOk

`func (o *IotServerOpenApiVO) GetRssiFormatIotServerOk() (*int32, bool)`

GetRssiFormatIotServerOk returns a tuple with the RssiFormatIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssiFormatIotServer

`func (o *IotServerOpenApiVO) SetRssiFormatIotServer(v int32)`

SetRssiFormatIotServer sets RssiFormatIotServer field to given value.


### GetServerTypeIotServer

`func (o *IotServerOpenApiVO) GetServerTypeIotServer() int32`

GetServerTypeIotServer returns the ServerTypeIotServer field if non-nil, zero value otherwise.

### GetServerTypeIotServerOk

`func (o *IotServerOpenApiVO) GetServerTypeIotServerOk() (*int32, bool)`

GetServerTypeIotServerOk returns a tuple with the ServerTypeIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTypeIotServer

`func (o *IotServerOpenApiVO) SetServerTypeIotServer(v int32)`

SetServerTypeIotServer sets ServerTypeIotServer field to given value.


### GetServerUrlIotServer

`func (o *IotServerOpenApiVO) GetServerUrlIotServer() string`

GetServerUrlIotServer returns the ServerUrlIotServer field if non-nil, zero value otherwise.

### GetServerUrlIotServerOk

`func (o *IotServerOpenApiVO) GetServerUrlIotServerOk() (*string, bool)`

GetServerUrlIotServerOk returns a tuple with the ServerUrlIotServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrlIotServer

`func (o *IotServerOpenApiVO) SetServerUrlIotServer(v string)`

SetServerUrlIotServer sets ServerUrlIotServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


