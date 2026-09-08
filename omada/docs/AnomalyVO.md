# AnomalyVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyCode** | Pointer to **string** | For the values of Anomaly event code, refer to section 5.7.2.1 of the Open API Access | [optional] 
**Category** | Pointer to **int32** | Anomaly category. 11: Network, 12: Device, 13: Client, etc. | [optional] 
**Causes** | Pointer to [**map[string]AnomalyCauseVO**](AnomalyCauseVO.md) | Root causes map. Key is cause code, value is cause detail. | [optional] 
**Clients** | Pointer to [**map[string]ClientObjectDTO**](ClientObjectDTO.md) | Client objects map. Key is MAC address, value is client info. | [optional] 
**ContentParams** | Pointer to **map[string]string** | Content parameter map for rendering the event content template. | [optional] 
**Detail** | Pointer to **map[string]map[string]interface{}** | Detailed chart data for the event content display (line chart, timeline, list, bar chart, protocol replay, port POE, etc.). | [optional] 
**Devices** | Pointer to [**map[string]DeviceObjectDTO**](DeviceObjectDTO.md) | Device objects map. Key is MAC address, value is device info. | [optional] 
**EndTime** | Pointer to **int64** | Continuous event end time in milliseconds (Unix timestamp). | [optional] 
**ImpactedClients** | Pointer to **[]string** | Set of MAC addresses of all clients impacted by this incident. | [optional] 
**ImpactedDevices** | Pointer to **[]string** | Set of MAC addresses of all devices impacted by this incident. | [optional] 
**IncidentId** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **int32** | Event severity level. 0: Critical, 1: Error, 2: Warning, 3: Info. | [optional] 
**Object** | Pointer to **[]string** | List of device/client names where the anomaly occurred. | [optional] 
**ObjectType** | Pointer to **string** | Object type of the anomaly target: gateway, switch, ap, wiredClient, wirelessClient. | [optional] 
**StartTime** | Pointer to **int64** | Continuous event start time in milliseconds (Unix timestamp). | [optional] 
**Status** | Pointer to **int32** | Event status. 0: Unresolved, 1: Resolved, 2: Ignored. | [optional] 
**Time** | Pointer to **int64** | Event creation time in milliseconds (Unix timestamp). | [optional] 
**TitleParams** | Pointer to **map[string]string** | Title parameter map for rendering the event title template. | [optional] 

## Methods

### NewAnomalyVO

`func NewAnomalyVO() *AnomalyVO`

NewAnomalyVO instantiates a new AnomalyVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyVOWithDefaults

`func NewAnomalyVOWithDefaults() *AnomalyVO`

NewAnomalyVOWithDefaults instantiates a new AnomalyVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyCode

`func (o *AnomalyVO) GetAnomalyCode() string`

GetAnomalyCode returns the AnomalyCode field if non-nil, zero value otherwise.

### GetAnomalyCodeOk

`func (o *AnomalyVO) GetAnomalyCodeOk() (*string, bool)`

GetAnomalyCodeOk returns a tuple with the AnomalyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyCode

`func (o *AnomalyVO) SetAnomalyCode(v string)`

SetAnomalyCode sets AnomalyCode field to given value.

### HasAnomalyCode

`func (o *AnomalyVO) HasAnomalyCode() bool`

HasAnomalyCode returns a boolean if a field has been set.

### GetCategory

`func (o *AnomalyVO) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AnomalyVO) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AnomalyVO) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AnomalyVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCauses

`func (o *AnomalyVO) GetCauses() map[string]AnomalyCauseVO`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *AnomalyVO) GetCausesOk() (*map[string]AnomalyCauseVO, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *AnomalyVO) SetCauses(v map[string]AnomalyCauseVO)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *AnomalyVO) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetClients

`func (o *AnomalyVO) GetClients() map[string]ClientObjectDTO`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *AnomalyVO) GetClientsOk() (*map[string]ClientObjectDTO, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *AnomalyVO) SetClients(v map[string]ClientObjectDTO)`

SetClients sets Clients field to given value.

### HasClients

`func (o *AnomalyVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetContentParams

`func (o *AnomalyVO) GetContentParams() map[string]string`

GetContentParams returns the ContentParams field if non-nil, zero value otherwise.

### GetContentParamsOk

`func (o *AnomalyVO) GetContentParamsOk() (*map[string]string, bool)`

GetContentParamsOk returns a tuple with the ContentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentParams

`func (o *AnomalyVO) SetContentParams(v map[string]string)`

SetContentParams sets ContentParams field to given value.

### HasContentParams

`func (o *AnomalyVO) HasContentParams() bool`

HasContentParams returns a boolean if a field has been set.

### GetDetail

`func (o *AnomalyVO) GetDetail() map[string]map[string]interface{}`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AnomalyVO) GetDetailOk() (*map[string]map[string]interface{}, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AnomalyVO) SetDetail(v map[string]map[string]interface{})`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AnomalyVO) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDevices

`func (o *AnomalyVO) GetDevices() map[string]DeviceObjectDTO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *AnomalyVO) GetDevicesOk() (*map[string]DeviceObjectDTO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *AnomalyVO) SetDevices(v map[string]DeviceObjectDTO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *AnomalyVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetEndTime

`func (o *AnomalyVO) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *AnomalyVO) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *AnomalyVO) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *AnomalyVO) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetImpactedClients

`func (o *AnomalyVO) GetImpactedClients() []string`

GetImpactedClients returns the ImpactedClients field if non-nil, zero value otherwise.

### GetImpactedClientsOk

`func (o *AnomalyVO) GetImpactedClientsOk() (*[]string, bool)`

GetImpactedClientsOk returns a tuple with the ImpactedClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedClients

`func (o *AnomalyVO) SetImpactedClients(v []string)`

SetImpactedClients sets ImpactedClients field to given value.

### HasImpactedClients

`func (o *AnomalyVO) HasImpactedClients() bool`

HasImpactedClients returns a boolean if a field has been set.

### GetImpactedDevices

`func (o *AnomalyVO) GetImpactedDevices() []string`

GetImpactedDevices returns the ImpactedDevices field if non-nil, zero value otherwise.

### GetImpactedDevicesOk

`func (o *AnomalyVO) GetImpactedDevicesOk() (*[]string, bool)`

GetImpactedDevicesOk returns a tuple with the ImpactedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactedDevices

`func (o *AnomalyVO) SetImpactedDevices(v []string)`

SetImpactedDevices sets ImpactedDevices field to given value.

### HasImpactedDevices

`func (o *AnomalyVO) HasImpactedDevices() bool`

HasImpactedDevices returns a boolean if a field has been set.

### GetIncidentId

`func (o *AnomalyVO) GetIncidentId() string`

GetIncidentId returns the IncidentId field if non-nil, zero value otherwise.

### GetIncidentIdOk

`func (o *AnomalyVO) GetIncidentIdOk() (*string, bool)`

GetIncidentIdOk returns a tuple with the IncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentId

`func (o *AnomalyVO) SetIncidentId(v string)`

SetIncidentId sets IncidentId field to given value.

### HasIncidentId

`func (o *AnomalyVO) HasIncidentId() bool`

HasIncidentId returns a boolean if a field has been set.

### GetLevel

`func (o *AnomalyVO) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AnomalyVO) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AnomalyVO) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AnomalyVO) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetObject

`func (o *AnomalyVO) GetObject() []string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AnomalyVO) GetObjectOk() (*[]string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AnomalyVO) SetObject(v []string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AnomalyVO) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetObjectType

`func (o *AnomalyVO) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AnomalyVO) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AnomalyVO) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AnomalyVO) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetStartTime

`func (o *AnomalyVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *AnomalyVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *AnomalyVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *AnomalyVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *AnomalyVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnomalyVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnomalyVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AnomalyVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *AnomalyVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AnomalyVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AnomalyVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *AnomalyVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTitleParams

`func (o *AnomalyVO) GetTitleParams() map[string]string`

GetTitleParams returns the TitleParams field if non-nil, zero value otherwise.

### GetTitleParamsOk

`func (o *AnomalyVO) GetTitleParamsOk() (*map[string]string, bool)`

GetTitleParamsOk returns a tuple with the TitleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleParams

`func (o *AnomalyVO) SetTitleParams(v map[string]string)`

SetTitleParams sets TitleParams field to given value.

### HasTitleParams

`func (o *AnomalyVO) HasTitleParams() bool`

HasTitleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


