# OverviewStatisticsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **int32** | Number of access incidents (category code 11) | [optional] 
**All** | Pointer to **int32** | Total incident count across all categories | [optional] 
**AllLevel** | Pointer to **int32** | Total incident count across all levels | [optional] 
**AllStatus** | Pointer to **int32** | Total incident count across all statuses | [optional] 
**AnomalyCategoryEventCounts** | Pointer to [**[]AnomalyCategoryEventCountVO**](AnomalyCategoryEventCountVO.md) | Anomaly-code level event count information | [optional] 
**Authentication** | Pointer to **int32** | Number of authentication incidents (category code 12) | [optional] 
**Critical** | Pointer to **int32** | Number of critical-level incidents | [optional] 
**DeviceStatus** | Pointer to **int32** | Number of device status incidents (category code 18) | [optional] 
**Error** | Pointer to **int32** | Number of error-level incidents | [optional] 
**Ignored** | Pointer to **int32** | Number of ignored incidents | [optional] 
**Info** | Pointer to **int32** | Number of info-level incidents | [optional] 
**Link** | Pointer to **int32** | Number of link incidents (category code 16) | [optional] 
**Ongoing** | Pointer to **int32** | Number of ongoing incidents | [optional] 
**Others** | Pointer to **int32** | Deprecated events since v6.3 | [optional] 
**Resolved** | Pointer to **int32** | Number of resolved incidents | [optional] 
**Roaming** | Pointer to **int32** | Number of roaming incidents (category code 13) | [optional] 
**Security** | Pointer to **int32** | Number of security incidents (category code 19) | [optional] 
**Unresolved** | Pointer to **int32** | Number of unresolved incidents | [optional] 
**WanAndServices** | Pointer to **int32** | Number of WAN and services incidents (category code 17) | [optional] 
**Warning** | Pointer to **int32** | Number of warning-level incidents | [optional] 
**WiredNetwork** | Pointer to **int32** | Number of wired network incidents (category code 15) | [optional] 
**WirelessNetwork** | Pointer to **int32** | Number of wireless network incidents (category code 14) | [optional] 

## Methods

### NewOverviewStatisticsOpenApiVO

`func NewOverviewStatisticsOpenApiVO() *OverviewStatisticsOpenApiVO`

NewOverviewStatisticsOpenApiVO instantiates a new OverviewStatisticsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverviewStatisticsOpenApiVOWithDefaults

`func NewOverviewStatisticsOpenApiVOWithDefaults() *OverviewStatisticsOpenApiVO`

NewOverviewStatisticsOpenApiVOWithDefaults instantiates a new OverviewStatisticsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *OverviewStatisticsOpenApiVO) GetAccess() int32`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *OverviewStatisticsOpenApiVO) GetAccessOk() (*int32, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *OverviewStatisticsOpenApiVO) SetAccess(v int32)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *OverviewStatisticsOpenApiVO) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAll

`func (o *OverviewStatisticsOpenApiVO) GetAll() int32`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *OverviewStatisticsOpenApiVO) GetAllOk() (*int32, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *OverviewStatisticsOpenApiVO) SetAll(v int32)`

SetAll sets All field to given value.

### HasAll

`func (o *OverviewStatisticsOpenApiVO) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetAllLevel

`func (o *OverviewStatisticsOpenApiVO) GetAllLevel() int32`

GetAllLevel returns the AllLevel field if non-nil, zero value otherwise.

### GetAllLevelOk

`func (o *OverviewStatisticsOpenApiVO) GetAllLevelOk() (*int32, bool)`

GetAllLevelOk returns a tuple with the AllLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLevel

`func (o *OverviewStatisticsOpenApiVO) SetAllLevel(v int32)`

SetAllLevel sets AllLevel field to given value.

### HasAllLevel

`func (o *OverviewStatisticsOpenApiVO) HasAllLevel() bool`

HasAllLevel returns a boolean if a field has been set.

### GetAllStatus

`func (o *OverviewStatisticsOpenApiVO) GetAllStatus() int32`

GetAllStatus returns the AllStatus field if non-nil, zero value otherwise.

### GetAllStatusOk

`func (o *OverviewStatisticsOpenApiVO) GetAllStatusOk() (*int32, bool)`

GetAllStatusOk returns a tuple with the AllStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllStatus

`func (o *OverviewStatisticsOpenApiVO) SetAllStatus(v int32)`

SetAllStatus sets AllStatus field to given value.

### HasAllStatus

`func (o *OverviewStatisticsOpenApiVO) HasAllStatus() bool`

HasAllStatus returns a boolean if a field has been set.

### GetAnomalyCategoryEventCounts

`func (o *OverviewStatisticsOpenApiVO) GetAnomalyCategoryEventCounts() []AnomalyCategoryEventCountVO`

GetAnomalyCategoryEventCounts returns the AnomalyCategoryEventCounts field if non-nil, zero value otherwise.

### GetAnomalyCategoryEventCountsOk

`func (o *OverviewStatisticsOpenApiVO) GetAnomalyCategoryEventCountsOk() (*[]AnomalyCategoryEventCountVO, bool)`

GetAnomalyCategoryEventCountsOk returns a tuple with the AnomalyCategoryEventCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyCategoryEventCounts

`func (o *OverviewStatisticsOpenApiVO) SetAnomalyCategoryEventCounts(v []AnomalyCategoryEventCountVO)`

SetAnomalyCategoryEventCounts sets AnomalyCategoryEventCounts field to given value.

### HasAnomalyCategoryEventCounts

`func (o *OverviewStatisticsOpenApiVO) HasAnomalyCategoryEventCounts() bool`

HasAnomalyCategoryEventCounts returns a boolean if a field has been set.

### GetAuthentication

`func (o *OverviewStatisticsOpenApiVO) GetAuthentication() int32`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *OverviewStatisticsOpenApiVO) GetAuthenticationOk() (*int32, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *OverviewStatisticsOpenApiVO) SetAuthentication(v int32)`

SetAuthentication sets Authentication field to given value.

### HasAuthentication

`func (o *OverviewStatisticsOpenApiVO) HasAuthentication() bool`

HasAuthentication returns a boolean if a field has been set.

### GetCritical

`func (o *OverviewStatisticsOpenApiVO) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *OverviewStatisticsOpenApiVO) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *OverviewStatisticsOpenApiVO) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *OverviewStatisticsOpenApiVO) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *OverviewStatisticsOpenApiVO) GetDeviceStatus() int32`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *OverviewStatisticsOpenApiVO) GetDeviceStatusOk() (*int32, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *OverviewStatisticsOpenApiVO) SetDeviceStatus(v int32)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *OverviewStatisticsOpenApiVO) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetError

`func (o *OverviewStatisticsOpenApiVO) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OverviewStatisticsOpenApiVO) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OverviewStatisticsOpenApiVO) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *OverviewStatisticsOpenApiVO) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIgnored

`func (o *OverviewStatisticsOpenApiVO) GetIgnored() int32`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *OverviewStatisticsOpenApiVO) GetIgnoredOk() (*int32, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *OverviewStatisticsOpenApiVO) SetIgnored(v int32)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *OverviewStatisticsOpenApiVO) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### GetInfo

`func (o *OverviewStatisticsOpenApiVO) GetInfo() int32`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OverviewStatisticsOpenApiVO) GetInfoOk() (*int32, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OverviewStatisticsOpenApiVO) SetInfo(v int32)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OverviewStatisticsOpenApiVO) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetLink

`func (o *OverviewStatisticsOpenApiVO) GetLink() int32`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *OverviewStatisticsOpenApiVO) GetLinkOk() (*int32, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *OverviewStatisticsOpenApiVO) SetLink(v int32)`

SetLink sets Link field to given value.

### HasLink

`func (o *OverviewStatisticsOpenApiVO) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetOngoing

`func (o *OverviewStatisticsOpenApiVO) GetOngoing() int32`

GetOngoing returns the Ongoing field if non-nil, zero value otherwise.

### GetOngoingOk

`func (o *OverviewStatisticsOpenApiVO) GetOngoingOk() (*int32, bool)`

GetOngoingOk returns a tuple with the Ongoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoing

`func (o *OverviewStatisticsOpenApiVO) SetOngoing(v int32)`

SetOngoing sets Ongoing field to given value.

### HasOngoing

`func (o *OverviewStatisticsOpenApiVO) HasOngoing() bool`

HasOngoing returns a boolean if a field has been set.

### GetOthers

`func (o *OverviewStatisticsOpenApiVO) GetOthers() int32`

GetOthers returns the Others field if non-nil, zero value otherwise.

### GetOthersOk

`func (o *OverviewStatisticsOpenApiVO) GetOthersOk() (*int32, bool)`

GetOthersOk returns a tuple with the Others field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthers

`func (o *OverviewStatisticsOpenApiVO) SetOthers(v int32)`

SetOthers sets Others field to given value.

### HasOthers

`func (o *OverviewStatisticsOpenApiVO) HasOthers() bool`

HasOthers returns a boolean if a field has been set.

### GetResolved

`func (o *OverviewStatisticsOpenApiVO) GetResolved() int32`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *OverviewStatisticsOpenApiVO) GetResolvedOk() (*int32, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *OverviewStatisticsOpenApiVO) SetResolved(v int32)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *OverviewStatisticsOpenApiVO) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetRoaming

`func (o *OverviewStatisticsOpenApiVO) GetRoaming() int32`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *OverviewStatisticsOpenApiVO) GetRoamingOk() (*int32, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *OverviewStatisticsOpenApiVO) SetRoaming(v int32)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *OverviewStatisticsOpenApiVO) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.

### GetSecurity

`func (o *OverviewStatisticsOpenApiVO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *OverviewStatisticsOpenApiVO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *OverviewStatisticsOpenApiVO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *OverviewStatisticsOpenApiVO) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetUnresolved

`func (o *OverviewStatisticsOpenApiVO) GetUnresolved() int32`

GetUnresolved returns the Unresolved field if non-nil, zero value otherwise.

### GetUnresolvedOk

`func (o *OverviewStatisticsOpenApiVO) GetUnresolvedOk() (*int32, bool)`

GetUnresolvedOk returns a tuple with the Unresolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnresolved

`func (o *OverviewStatisticsOpenApiVO) SetUnresolved(v int32)`

SetUnresolved sets Unresolved field to given value.

### HasUnresolved

`func (o *OverviewStatisticsOpenApiVO) HasUnresolved() bool`

HasUnresolved returns a boolean if a field has been set.

### GetWanAndServices

`func (o *OverviewStatisticsOpenApiVO) GetWanAndServices() int32`

GetWanAndServices returns the WanAndServices field if non-nil, zero value otherwise.

### GetWanAndServicesOk

`func (o *OverviewStatisticsOpenApiVO) GetWanAndServicesOk() (*int32, bool)`

GetWanAndServicesOk returns a tuple with the WanAndServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanAndServices

`func (o *OverviewStatisticsOpenApiVO) SetWanAndServices(v int32)`

SetWanAndServices sets WanAndServices field to given value.

### HasWanAndServices

`func (o *OverviewStatisticsOpenApiVO) HasWanAndServices() bool`

HasWanAndServices returns a boolean if a field has been set.

### GetWarning

`func (o *OverviewStatisticsOpenApiVO) GetWarning() int32`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *OverviewStatisticsOpenApiVO) GetWarningOk() (*int32, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *OverviewStatisticsOpenApiVO) SetWarning(v int32)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *OverviewStatisticsOpenApiVO) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetWiredNetwork

`func (o *OverviewStatisticsOpenApiVO) GetWiredNetwork() int32`

GetWiredNetwork returns the WiredNetwork field if non-nil, zero value otherwise.

### GetWiredNetworkOk

`func (o *OverviewStatisticsOpenApiVO) GetWiredNetworkOk() (*int32, bool)`

GetWiredNetworkOk returns a tuple with the WiredNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredNetwork

`func (o *OverviewStatisticsOpenApiVO) SetWiredNetwork(v int32)`

SetWiredNetwork sets WiredNetwork field to given value.

### HasWiredNetwork

`func (o *OverviewStatisticsOpenApiVO) HasWiredNetwork() bool`

HasWiredNetwork returns a boolean if a field has been set.

### GetWirelessNetwork

`func (o *OverviewStatisticsOpenApiVO) GetWirelessNetwork() int32`

GetWirelessNetwork returns the WirelessNetwork field if non-nil, zero value otherwise.

### GetWirelessNetworkOk

`func (o *OverviewStatisticsOpenApiVO) GetWirelessNetworkOk() (*int32, bool)`

GetWirelessNetworkOk returns a tuple with the WirelessNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessNetwork

`func (o *OverviewStatisticsOpenApiVO) SetWirelessNetwork(v int32)`

SetWirelessNetwork sets WirelessNetwork field to given value.

### HasWirelessNetwork

`func (o *OverviewStatisticsOpenApiVO) HasWirelessNetwork() bool`

HasWirelessNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


