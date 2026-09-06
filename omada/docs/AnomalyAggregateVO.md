# AnomalyAggregateVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyCode** | Pointer to **string** | For the values of Anomaly event code, refer to section 5.7.2.1 of the Open API Access | [optional] 
**Category** | Pointer to **int32** | Anomaly category. 11: Network, 12: Device, 13: Client, etc. | [optional] 
**Causes** | Pointer to [**[]CauseVO**](CauseVO.md) | Root causes and advices for this anomaly. | [optional] 
**Clients** | Pointer to [**map[string]ClientObjectDTO**](ClientObjectDTO.md) | Client objects map. Key is MAC address, value is client info. | [optional] 
**ContentParams** | Pointer to **map[string]string** | Content parameter map for rendering the event content template. | [optional] 
**Count** | Pointer to **int32** | Number of occurrences of this anomaly. | [optional] 
**Devices** | Pointer to [**map[string]DeviceObjectDTO**](DeviceObjectDTO.md) | Device objects map. Key is MAC address, value is device info. | [optional] 
**FirstTime** | Pointer to **int64** | First occurrence time in milliseconds (Unix timestamp). | [optional] 
**InfluencingClients** | Pointer to **[]string** | List of MAC addresses of influencing/impacted clients. | [optional] 
**InfluencingDevices** | Pointer to **[]string** | List of MAC addresses of influencing/impacted devices. | [optional] 
**LastTime** | Pointer to **int64** | Last occurrence time in milliseconds (Unix timestamp). | [optional] 
**Level** | Pointer to **int32** | Event severity level. 0: Critical, 1: Error, 2: Warning, 3: Info. | [optional] 
**Object** | Pointer to **[]string** | List of device/client names where the anomaly occurred. | [optional] 
**ObjectType** | Pointer to **string** | Object type of the anomaly target: gateway, switch, ap, wiredClient, wirelessClient. | [optional] 
**Status** | Pointer to **int32** | Event status. 0: Unresolved, 1: Resolved, 2: Ignored. | [optional] 
**TitleParams** | Pointer to **map[string]string** | Title parameter map for rendering the event title template. | [optional] 

## Methods

### NewAnomalyAggregateVO

`func NewAnomalyAggregateVO() *AnomalyAggregateVO`

NewAnomalyAggregateVO instantiates a new AnomalyAggregateVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyAggregateVOWithDefaults

`func NewAnomalyAggregateVOWithDefaults() *AnomalyAggregateVO`

NewAnomalyAggregateVOWithDefaults instantiates a new AnomalyAggregateVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyCode

`func (o *AnomalyAggregateVO) GetAnomalyCode() string`

GetAnomalyCode returns the AnomalyCode field if non-nil, zero value otherwise.

### GetAnomalyCodeOk

`func (o *AnomalyAggregateVO) GetAnomalyCodeOk() (*string, bool)`

GetAnomalyCodeOk returns a tuple with the AnomalyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyCode

`func (o *AnomalyAggregateVO) SetAnomalyCode(v string)`

SetAnomalyCode sets AnomalyCode field to given value.

### HasAnomalyCode

`func (o *AnomalyAggregateVO) HasAnomalyCode() bool`

HasAnomalyCode returns a boolean if a field has been set.

### GetCategory

`func (o *AnomalyAggregateVO) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AnomalyAggregateVO) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AnomalyAggregateVO) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AnomalyAggregateVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetCauses

`func (o *AnomalyAggregateVO) GetCauses() []CauseVO`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *AnomalyAggregateVO) GetCausesOk() (*[]CauseVO, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *AnomalyAggregateVO) SetCauses(v []CauseVO)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *AnomalyAggregateVO) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetClients

`func (o *AnomalyAggregateVO) GetClients() map[string]ClientObjectDTO`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *AnomalyAggregateVO) GetClientsOk() (*map[string]ClientObjectDTO, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *AnomalyAggregateVO) SetClients(v map[string]ClientObjectDTO)`

SetClients sets Clients field to given value.

### HasClients

`func (o *AnomalyAggregateVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetContentParams

`func (o *AnomalyAggregateVO) GetContentParams() map[string]string`

GetContentParams returns the ContentParams field if non-nil, zero value otherwise.

### GetContentParamsOk

`func (o *AnomalyAggregateVO) GetContentParamsOk() (*map[string]string, bool)`

GetContentParamsOk returns a tuple with the ContentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentParams

`func (o *AnomalyAggregateVO) SetContentParams(v map[string]string)`

SetContentParams sets ContentParams field to given value.

### HasContentParams

`func (o *AnomalyAggregateVO) HasContentParams() bool`

HasContentParams returns a boolean if a field has been set.

### GetCount

`func (o *AnomalyAggregateVO) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AnomalyAggregateVO) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AnomalyAggregateVO) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AnomalyAggregateVO) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetDevices

`func (o *AnomalyAggregateVO) GetDevices() map[string]DeviceObjectDTO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *AnomalyAggregateVO) GetDevicesOk() (*map[string]DeviceObjectDTO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *AnomalyAggregateVO) SetDevices(v map[string]DeviceObjectDTO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *AnomalyAggregateVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetFirstTime

`func (o *AnomalyAggregateVO) GetFirstTime() int64`

GetFirstTime returns the FirstTime field if non-nil, zero value otherwise.

### GetFirstTimeOk

`func (o *AnomalyAggregateVO) GetFirstTimeOk() (*int64, bool)`

GetFirstTimeOk returns a tuple with the FirstTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTime

`func (o *AnomalyAggregateVO) SetFirstTime(v int64)`

SetFirstTime sets FirstTime field to given value.

### HasFirstTime

`func (o *AnomalyAggregateVO) HasFirstTime() bool`

HasFirstTime returns a boolean if a field has been set.

### GetInfluencingClients

`func (o *AnomalyAggregateVO) GetInfluencingClients() []string`

GetInfluencingClients returns the InfluencingClients field if non-nil, zero value otherwise.

### GetInfluencingClientsOk

`func (o *AnomalyAggregateVO) GetInfluencingClientsOk() (*[]string, bool)`

GetInfluencingClientsOk returns a tuple with the InfluencingClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluencingClients

`func (o *AnomalyAggregateVO) SetInfluencingClients(v []string)`

SetInfluencingClients sets InfluencingClients field to given value.

### HasInfluencingClients

`func (o *AnomalyAggregateVO) HasInfluencingClients() bool`

HasInfluencingClients returns a boolean if a field has been set.

### GetInfluencingDevices

`func (o *AnomalyAggregateVO) GetInfluencingDevices() []string`

GetInfluencingDevices returns the InfluencingDevices field if non-nil, zero value otherwise.

### GetInfluencingDevicesOk

`func (o *AnomalyAggregateVO) GetInfluencingDevicesOk() (*[]string, bool)`

GetInfluencingDevicesOk returns a tuple with the InfluencingDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluencingDevices

`func (o *AnomalyAggregateVO) SetInfluencingDevices(v []string)`

SetInfluencingDevices sets InfluencingDevices field to given value.

### HasInfluencingDevices

`func (o *AnomalyAggregateVO) HasInfluencingDevices() bool`

HasInfluencingDevices returns a boolean if a field has been set.

### GetLastTime

`func (o *AnomalyAggregateVO) GetLastTime() int64`

GetLastTime returns the LastTime field if non-nil, zero value otherwise.

### GetLastTimeOk

`func (o *AnomalyAggregateVO) GetLastTimeOk() (*int64, bool)`

GetLastTimeOk returns a tuple with the LastTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTime

`func (o *AnomalyAggregateVO) SetLastTime(v int64)`

SetLastTime sets LastTime field to given value.

### HasLastTime

`func (o *AnomalyAggregateVO) HasLastTime() bool`

HasLastTime returns a boolean if a field has been set.

### GetLevel

`func (o *AnomalyAggregateVO) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AnomalyAggregateVO) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AnomalyAggregateVO) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *AnomalyAggregateVO) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetObject

`func (o *AnomalyAggregateVO) GetObject() []string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AnomalyAggregateVO) GetObjectOk() (*[]string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AnomalyAggregateVO) SetObject(v []string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AnomalyAggregateVO) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetObjectType

`func (o *AnomalyAggregateVO) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AnomalyAggregateVO) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AnomalyAggregateVO) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AnomalyAggregateVO) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetStatus

`func (o *AnomalyAggregateVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnomalyAggregateVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnomalyAggregateVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AnomalyAggregateVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitleParams

`func (o *AnomalyAggregateVO) GetTitleParams() map[string]string`

GetTitleParams returns the TitleParams field if non-nil, zero value otherwise.

### GetTitleParamsOk

`func (o *AnomalyAggregateVO) GetTitleParamsOk() (*map[string]string, bool)`

GetTitleParamsOk returns a tuple with the TitleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleParams

`func (o *AnomalyAggregateVO) SetTitleParams(v map[string]string)`

SetTitleParams sets TitleParams field to given value.

### HasTitleParams

`func (o *AnomalyAggregateVO) HasTitleParams() bool`

HasTitleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


