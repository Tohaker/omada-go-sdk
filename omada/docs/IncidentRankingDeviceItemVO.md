# IncidentRankingDeviceItemVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incidents** | Pointer to **int32** | Number of incidents involving this device | [optional] 
**Mac** | Pointer to **string** | MAC address of the device | [optional] 
**Model** | Pointer to **string** | Device model (e.g. EAP225) | [optional] 
**ModelVersion** | Pointer to **string** | Device model version (e.g. 3.0) | [optional] 
**Name** | Pointer to **string** | Display name of the device | [optional] 
**Type** | Pointer to **string** | Device type (e.g. ap, gateway, switch) | [optional] 

## Methods

### NewIncidentRankingDeviceItemVO

`func NewIncidentRankingDeviceItemVO() *IncidentRankingDeviceItemVO`

NewIncidentRankingDeviceItemVO instantiates a new IncidentRankingDeviceItemVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentRankingDeviceItemVOWithDefaults

`func NewIncidentRankingDeviceItemVOWithDefaults() *IncidentRankingDeviceItemVO`

NewIncidentRankingDeviceItemVOWithDefaults instantiates a new IncidentRankingDeviceItemVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidents

`func (o *IncidentRankingDeviceItemVO) GetIncidents() int32`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *IncidentRankingDeviceItemVO) GetIncidentsOk() (*int32, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *IncidentRankingDeviceItemVO) SetIncidents(v int32)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *IncidentRankingDeviceItemVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetMac

`func (o *IncidentRankingDeviceItemVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *IncidentRankingDeviceItemVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *IncidentRankingDeviceItemVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *IncidentRankingDeviceItemVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *IncidentRankingDeviceItemVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *IncidentRankingDeviceItemVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *IncidentRankingDeviceItemVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *IncidentRankingDeviceItemVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *IncidentRankingDeviceItemVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *IncidentRankingDeviceItemVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *IncidentRankingDeviceItemVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *IncidentRankingDeviceItemVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *IncidentRankingDeviceItemVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncidentRankingDeviceItemVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncidentRankingDeviceItemVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IncidentRankingDeviceItemVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IncidentRankingDeviceItemVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentRankingDeviceItemVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentRankingDeviceItemVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentRankingDeviceItemVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


