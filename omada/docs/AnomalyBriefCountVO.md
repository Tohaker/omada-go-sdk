# AnomalyBriefCountVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyCode** | Pointer to **string** | For the values of Anomaly event code, refer to section 5.7.2.1 of the Open API Access | [optional] 
**Category** | Pointer to **int32** | Anomaly category. 11: Network, 12: Device, 13: Client, etc. | [optional] 
**Causes** | Pointer to [**map[string]AnomalyCauseVO**](AnomalyCauseVO.md) | Root cause information map | [optional] 
**Clients** | Pointer to [**map[string]ClientObjectDTO**](ClientObjectDTO.md) | Client objects map, key is MAC address | [optional] 
**ContentParams** | Pointer to **map[string]string** | Content parameters for anomaly message | [optional] 
**Count** | Pointer to **int32** | Total occurrence count | [optional] 
**Devices** | Pointer to [**map[string]DeviceObjectDTO**](DeviceObjectDTO.md) | Device objects map, key is MAC address | [optional] 
**IncidentId** | Pointer to **string** | Incident ID | [optional] 
**InfluencingClients** | Pointer to **[]string** | List of influenced client MAC addresses | [optional] 
**InfluencingDevices** | Pointer to **[]string** | List of influenced device MAC addresses | [optional] 
**LastTime** | Pointer to **int64** | Last occurrence timestamp (ms) | [optional] 
**Level** | Pointer to **int32** | Anomaly severity level | [optional] 
**Macs** | Pointer to **[]string** | List of affected device MAC addresses | [optional] 
**Object** | Pointer to **[]string** | List of affected device MAC addresses | [optional] 
**Status** | Pointer to **int32** | Anomaly incident status (0&#x3D;unresolved, 1&#x3D;resolved, 2&#x3D;ignored, 3&#x3D;ongoing) | [optional] 
**TitleParams** | Pointer to **map[string]string** | Title parameters for anomaly message | [optional] 

## Methods

### NewAnomalyBriefCountVO

`func NewAnomalyBriefCountVO() *AnomalyBriefCountVO`

NewAnomalyBriefCountVO instantiates a new AnomalyBriefCountVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyBriefCountVOWithDefaults

`func NewAnomalyBriefCountVOWithDefaults() *AnomalyBriefCountVO`

NewAnomalyBriefCountVOWithDefaults instantiates a new AnomalyBriefCountVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyCode

`func (o *AnomalyBriefCountVO) GetAnomalyCode() string`

GetAnomalyCode returns the AnomalyCode field if non-nil, zero value otherwise.

### GetAnomalyCodeOk

`func (o *AnomalyBriefCountVO) GetAnomalyCodeOk() (*string, bool)`

GetAnomalyCodeOk returns a tuple with the AnomalyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyCode

`func (o *AnomalyBriefCountVO) SetAnomalyCode(v string)`

SetAnomalyCode sets AnomalyCode field to given value.

### HasAnomalyCode

`func (o *AnomalyBriefCountVO) HasAnomalyCode() bool`

HasAnomalyCode returns a boolean if a field has been set.

### GetCategory

`func (o *AnomalyBriefCountVO) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AnomalyBriefCountVO) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AnomalyBriefCountVO) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AnomalyBriefCountVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCauses

`func (o *AnomalyBriefCountVO) GetCauses() map[string]AnomalyCauseVO`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *AnomalyBriefCountVO) GetCausesOk() (*map[string]AnomalyCauseVO, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *AnomalyBriefCountVO) SetCauses(v map[string]AnomalyCauseVO)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *AnomalyBriefCountVO) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetClients

`func (o *AnomalyBriefCountVO) GetClients() map[string]ClientObjectDTO`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *AnomalyBriefCountVO) GetClientsOk() (*map[string]ClientObjectDTO, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *AnomalyBriefCountVO) SetClients(v map[string]ClientObjectDTO)`

SetClients sets Clients field to given value.

### HasClients

`func (o *AnomalyBriefCountVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetContentParams

`func (o *AnomalyBriefCountVO) GetContentParams() map[string]string`

GetContentParams returns the ContentParams field if non-nil, zero value otherwise.

### GetContentParamsOk

`func (o *AnomalyBriefCountVO) GetContentParamsOk() (*map[string]string, bool)`

GetContentParamsOk returns a tuple with the ContentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentParams

`func (o *AnomalyBriefCountVO) SetContentParams(v map[string]string)`

SetContentParams sets ContentParams field to given value.

### HasContentParams

`func (o *AnomalyBriefCountVO) HasContentParams() bool`

HasContentParams returns a boolean if a field has been set.

### GetCount

`func (o *AnomalyBriefCountVO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AnomalyBriefCountVO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AnomalyBriefCountVO) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AnomalyBriefCountVO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDevices

`func (o *AnomalyBriefCountVO) GetDevices() map[string]DeviceObjectDTO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *AnomalyBriefCountVO) GetDevicesOk() (*map[string]DeviceObjectDTO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *AnomalyBriefCountVO) SetDevices(v map[string]DeviceObjectDTO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *AnomalyBriefCountVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetIncidentId

`func (o *AnomalyBriefCountVO) GetIncidentId() string`

GetIncidentId returns the IncidentId field if non-nil, zero value otherwise.

### GetIncidentIdOk

`func (o *AnomalyBriefCountVO) GetIncidentIdOk() (*string, bool)`

GetIncidentIdOk returns a tuple with the IncidentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentId

`func (o *AnomalyBriefCountVO) SetIncidentId(v string)`

SetIncidentId sets IncidentId field to given value.

### HasIncidentId

`func (o *AnomalyBriefCountVO) HasIncidentId() bool`

HasIncidentId returns a boolean if a field has been set.

### GetInfluencingClients

`func (o *AnomalyBriefCountVO) GetInfluencingClients() []string`

GetInfluencingClients returns the InfluencingClients field if non-nil, zero value otherwise.

### GetInfluencingClientsOk

`func (o *AnomalyBriefCountVO) GetInfluencingClientsOk() (*[]string, bool)`

GetInfluencingClientsOk returns a tuple with the InfluencingClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluencingClients

`func (o *AnomalyBriefCountVO) SetInfluencingClients(v []string)`

SetInfluencingClients sets InfluencingClients field to given value.

### HasInfluencingClients

`func (o *AnomalyBriefCountVO) HasInfluencingClients() bool`

HasInfluencingClients returns a boolean if a field has been set.

### GetInfluencingDevices

`func (o *AnomalyBriefCountVO) GetInfluencingDevices() []string`

GetInfluencingDevices returns the InfluencingDevices field if non-nil, zero value otherwise.

### GetInfluencingDevicesOk

`func (o *AnomalyBriefCountVO) GetInfluencingDevicesOk() (*[]string, bool)`

GetInfluencingDevicesOk returns a tuple with the InfluencingDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluencingDevices

`func (o *AnomalyBriefCountVO) SetInfluencingDevices(v []string)`

SetInfluencingDevices sets InfluencingDevices field to given value.

### HasInfluencingDevices

`func (o *AnomalyBriefCountVO) HasInfluencingDevices() bool`

HasInfluencingDevices returns a boolean if a field has been set.

### GetLastTime

`func (o *AnomalyBriefCountVO) GetLastTime() int64`

GetLastTime returns the LastTime field if non-nil, zero value otherwise.

### GetLastTimeOk

`func (o *AnomalyBriefCountVO) GetLastTimeOk() (*int64, bool)`

GetLastTimeOk returns a tuple with the LastTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTime

`func (o *AnomalyBriefCountVO) SetLastTime(v int64)`

SetLastTime sets LastTime field to given value.

### HasLastTime

`func (o *AnomalyBriefCountVO) HasLastTime() bool`

HasLastTime returns a boolean if a field has been set.

### GetLevel

`func (o *AnomalyBriefCountVO) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AnomalyBriefCountVO) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AnomalyBriefCountVO) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AnomalyBriefCountVO) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMacs

`func (o *AnomalyBriefCountVO) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *AnomalyBriefCountVO) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *AnomalyBriefCountVO) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *AnomalyBriefCountVO) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetObject

`func (o *AnomalyBriefCountVO) GetObject() []string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AnomalyBriefCountVO) GetObjectOk() (*[]string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AnomalyBriefCountVO) SetObject(v []string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AnomalyBriefCountVO) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetStatus

`func (o *AnomalyBriefCountVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnomalyBriefCountVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnomalyBriefCountVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AnomalyBriefCountVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitleParams

`func (o *AnomalyBriefCountVO) GetTitleParams() map[string]string`

GetTitleParams returns the TitleParams field if non-nil, zero value otherwise.

### GetTitleParamsOk

`func (o *AnomalyBriefCountVO) GetTitleParamsOk() (*map[string]string, bool)`

GetTitleParamsOk returns a tuple with the TitleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleParams

`func (o *AnomalyBriefCountVO) SetTitleParams(v map[string]string)`

SetTitleParams sets TitleParams field to given value.

### HasTitleParams

`func (o *AnomalyBriefCountVO) HasTitleParams() bool`

HasTitleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


