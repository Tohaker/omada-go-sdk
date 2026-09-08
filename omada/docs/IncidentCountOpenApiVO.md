# IncidentCountOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **int32** | Number of access incidents in this time bucket | [optional] 
**All** | Pointer to **int32** | Total incident count across all levels in this time bucket | [optional] 
**Authentication** | Pointer to **int32** | Number of authentication incidents in this time bucket | [optional] 
**Critical** | Pointer to **int32** | Number of critical incidents in this time bucket | [optional] 
**DeviceStatus** | Pointer to **int32** | Number of device status incidents in this time bucket | [optional] 
**Error** | Pointer to **int32** | Number of error incidents in this time bucket | [optional] 
**Info** | Pointer to **int32** | Number of info incidents in this time bucket | [optional] 
**Link** | Pointer to **int32** | Number of link incidents in this time bucket | [optional] 
**Roaming** | Pointer to **int32** | Number of roaming incidents in this time bucket | [optional] 
**Security** | Pointer to **int32** | Number of security incidents in this time bucket | [optional] 
**Time** | Pointer to **int64** | Timestamp in seconds representing this time bucket | [optional] 
**WanAndServices** | Pointer to **int32** | Number of WAN and services incidents in this time bucket | [optional] 
**Warning** | Pointer to **int32** | Number of warning incidents in this time bucket | [optional] 
**WiredNetwork** | Pointer to **int32** | Number of wired network incidents in this time bucket | [optional] 
**WirelessNetwork** | Pointer to **int32** | Number of wireless network incidents in this time bucket | [optional] 

## Methods

### NewIncidentCountOpenApiVO

`func NewIncidentCountOpenApiVO() *IncidentCountOpenApiVO`

NewIncidentCountOpenApiVO instantiates a new IncidentCountOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentCountOpenApiVOWithDefaults

`func NewIncidentCountOpenApiVOWithDefaults() *IncidentCountOpenApiVO`

NewIncidentCountOpenApiVOWithDefaults instantiates a new IncidentCountOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *IncidentCountOpenApiVO) GetAccess() int32`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *IncidentCountOpenApiVO) GetAccessOk() (*int32, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *IncidentCountOpenApiVO) SetAccess(v int32)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *IncidentCountOpenApiVO) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAll

`func (o *IncidentCountOpenApiVO) GetAll() int32`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *IncidentCountOpenApiVO) GetAllOk() (*int32, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *IncidentCountOpenApiVO) SetAll(v int32)`

SetAll sets All field to given value.

### HasAll

`func (o *IncidentCountOpenApiVO) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAuthentication

`func (o *IncidentCountOpenApiVO) GetAuthentication() int32`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *IncidentCountOpenApiVO) GetAuthenticationOk() (*int32, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *IncidentCountOpenApiVO) SetAuthentication(v int32)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *IncidentCountOpenApiVO) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetCritical

`func (o *IncidentCountOpenApiVO) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *IncidentCountOpenApiVO) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *IncidentCountOpenApiVO) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *IncidentCountOpenApiVO) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *IncidentCountOpenApiVO) GetDeviceStatus() int32`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *IncidentCountOpenApiVO) GetDeviceStatusOk() (*int32, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *IncidentCountOpenApiVO) SetDeviceStatus(v int32)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *IncidentCountOpenApiVO) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetError

`func (o *IncidentCountOpenApiVO) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IncidentCountOpenApiVO) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IncidentCountOpenApiVO) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *IncidentCountOpenApiVO) HasError() bool`

HasError returns a boolean if a field has been set.

### GetInfo

`func (o *IncidentCountOpenApiVO) GetInfo() int32`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *IncidentCountOpenApiVO) GetInfoOk() (*int32, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *IncidentCountOpenApiVO) SetInfo(v int32)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *IncidentCountOpenApiVO) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLink

`func (o *IncidentCountOpenApiVO) GetLink() int32`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *IncidentCountOpenApiVO) GetLinkOk() (*int32, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *IncidentCountOpenApiVO) SetLink(v int32)`

SetLink sets Link field to given value.

### HasLink

`func (o *IncidentCountOpenApiVO) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetRoaming

`func (o *IncidentCountOpenApiVO) GetRoaming() int32`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *IncidentCountOpenApiVO) GetRoamingOk() (*int32, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *IncidentCountOpenApiVO) SetRoaming(v int32)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *IncidentCountOpenApiVO) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.

### GetSecurity

`func (o *IncidentCountOpenApiVO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *IncidentCountOpenApiVO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *IncidentCountOpenApiVO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *IncidentCountOpenApiVO) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetTime

`func (o *IncidentCountOpenApiVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *IncidentCountOpenApiVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *IncidentCountOpenApiVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *IncidentCountOpenApiVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetWanAndServices

`func (o *IncidentCountOpenApiVO) GetWanAndServices() int32`

GetWanAndServices returns the WanAndServices field if non-nil, zero value otherwise.

### GetWanAndServicesOk

`func (o *IncidentCountOpenApiVO) GetWanAndServicesOk() (*int32, bool)`

GetWanAndServicesOk returns a tuple with the WanAndServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanAndServices

`func (o *IncidentCountOpenApiVO) SetWanAndServices(v int32)`

SetWanAndServices sets WanAndServices field to given value.

### HasWanAndServices

`func (o *IncidentCountOpenApiVO) HasWanAndServices() bool`

HasWanAndServices returns a boolean if a field has been set.

### GetWarning

`func (o *IncidentCountOpenApiVO) GetWarning() int32`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *IncidentCountOpenApiVO) GetWarningOk() (*int32, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *IncidentCountOpenApiVO) SetWarning(v int32)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *IncidentCountOpenApiVO) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetWiredNetwork

`func (o *IncidentCountOpenApiVO) GetWiredNetwork() int32`

GetWiredNetwork returns the WiredNetwork field if non-nil, zero value otherwise.

### GetWiredNetworkOk

`func (o *IncidentCountOpenApiVO) GetWiredNetworkOk() (*int32, bool)`

GetWiredNetworkOk returns a tuple with the WiredNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredNetwork

`func (o *IncidentCountOpenApiVO) SetWiredNetwork(v int32)`

SetWiredNetwork sets WiredNetwork field to given value.

### HasWiredNetwork

`func (o *IncidentCountOpenApiVO) HasWiredNetwork() bool`

HasWiredNetwork returns a boolean if a field has been set.

### GetWirelessNetwork

`func (o *IncidentCountOpenApiVO) GetWirelessNetwork() int32`

GetWirelessNetwork returns the WirelessNetwork field if non-nil, zero value otherwise.

### GetWirelessNetworkOk

`func (o *IncidentCountOpenApiVO) GetWirelessNetworkOk() (*int32, bool)`

GetWirelessNetworkOk returns a tuple with the WirelessNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessNetwork

`func (o *IncidentCountOpenApiVO) SetWirelessNetwork(v int32)`

SetWirelessNetwork sets WirelessNetwork field to given value.

### HasWirelessNetwork

`func (o *IncidentCountOpenApiVO) HasWirelessNetwork() bool`

HasWirelessNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


