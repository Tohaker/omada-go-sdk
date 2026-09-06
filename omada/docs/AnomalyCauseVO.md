# AnomalyCauseVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advices** | Pointer to [**map[string]AnomalyAdviceVO**](AnomalyAdviceVO.md) |  | [optional] 
**Clients** | Pointer to [**map[string]ClientObjectDTO**](ClientObjectDTO.md) |  | [optional] 
**ContentParams** | Pointer to **map[string]string** |  | [optional] 
**Detail** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Devices** | Pointer to [**map[string]DeviceObjectDTO**](DeviceObjectDTO.md) |  | [optional] 
**TitleParams** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAnomalyCauseVO

`func NewAnomalyCauseVO() *AnomalyCauseVO`

NewAnomalyCauseVO instantiates a new AnomalyCauseVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyCauseVOWithDefaults

`func NewAnomalyCauseVOWithDefaults() *AnomalyCauseVO`

NewAnomalyCauseVOWithDefaults instantiates a new AnomalyCauseVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvices

`func (o *AnomalyCauseVO) GetAdvices() map[string]AnomalyAdviceVO`

GetAdvices returns the Advices field if non-nil, zero value otherwise.

### GetAdvicesOk

`func (o *AnomalyCauseVO) GetAdvicesOk() (*map[string]AnomalyAdviceVO, bool)`

GetAdvicesOk returns a tuple with the Advices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvices

`func (o *AnomalyCauseVO) SetAdvices(v map[string]AnomalyAdviceVO)`

SetAdvices sets Advices field to given value.

### HasAdvices

`func (o *AnomalyCauseVO) HasAdvices() bool`

HasAdvices returns a boolean if a field has been set.

### GetClients

`func (o *AnomalyCauseVO) GetClients() map[string]ClientObjectDTO`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *AnomalyCauseVO) GetClientsOk() (*map[string]ClientObjectDTO, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *AnomalyCauseVO) SetClients(v map[string]ClientObjectDTO)`

SetClients sets Clients field to given value.

### HasClients

`func (o *AnomalyCauseVO) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetContentParams

`func (o *AnomalyCauseVO) GetContentParams() map[string]string`

GetContentParams returns the ContentParams field if non-nil, zero value otherwise.

### GetContentParamsOk

`func (o *AnomalyCauseVO) GetContentParamsOk() (*map[string]string, bool)`

GetContentParamsOk returns a tuple with the ContentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentParams

`func (o *AnomalyCauseVO) SetContentParams(v map[string]string)`

SetContentParams sets ContentParams field to given value.

### HasContentParams

`func (o *AnomalyCauseVO) HasContentParams() bool`

HasContentParams returns a boolean if a field has been set.

### GetDetail

`func (o *AnomalyCauseVO) GetDetail() map[string]map[string]interface{}`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *AnomalyCauseVO) GetDetailOk() (*map[string]map[string]interface{}, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *AnomalyCauseVO) SetDetail(v map[string]map[string]interface{})`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *AnomalyCauseVO) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetDevices

`func (o *AnomalyCauseVO) GetDevices() map[string]DeviceObjectDTO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *AnomalyCauseVO) GetDevicesOk() (*map[string]DeviceObjectDTO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *AnomalyCauseVO) SetDevices(v map[string]DeviceObjectDTO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *AnomalyCauseVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetTitleParams

`func (o *AnomalyCauseVO) GetTitleParams() map[string]string`

GetTitleParams returns the TitleParams field if non-nil, zero value otherwise.

### GetTitleParamsOk

`func (o *AnomalyCauseVO) GetTitleParamsOk() (*map[string]string, bool)`

GetTitleParamsOk returns a tuple with the TitleParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleParams

`func (o *AnomalyCauseVO) SetTitleParams(v map[string]string)`

SetTitleParams sets TitleParams field to given value.

### HasTitleParams

`func (o *AnomalyCauseVO) HasTitleParams() bool`

HasTitleParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


