# IncidentRankingClientItemVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incidents** | Pointer to **int32** | Number of incidents involving this client | [optional] 
**Mac** | Pointer to **string** | MAC address of the client | [optional] 
**Manager** | Pointer to **bool** | Whether this client is a managed client | [optional] 
**Name** | Pointer to **string** | Display name of the client | [optional] 
**Type** | Pointer to **string** | Device type of the client (e.g. iphone, android, pc) | [optional] 

## Methods

### NewIncidentRankingClientItemVO

`func NewIncidentRankingClientItemVO() *IncidentRankingClientItemVO`

NewIncidentRankingClientItemVO instantiates a new IncidentRankingClientItemVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentRankingClientItemVOWithDefaults

`func NewIncidentRankingClientItemVOWithDefaults() *IncidentRankingClientItemVO`

NewIncidentRankingClientItemVOWithDefaults instantiates a new IncidentRankingClientItemVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncidents

`func (o *IncidentRankingClientItemVO) GetIncidents() int32`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *IncidentRankingClientItemVO) GetIncidentsOk() (*int32, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *IncidentRankingClientItemVO) SetIncidents(v int32)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *IncidentRankingClientItemVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetMac

`func (o *IncidentRankingClientItemVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *IncidentRankingClientItemVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *IncidentRankingClientItemVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *IncidentRankingClientItemVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManager

`func (o *IncidentRankingClientItemVO) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *IncidentRankingClientItemVO) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *IncidentRankingClientItemVO) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *IncidentRankingClientItemVO) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetName

`func (o *IncidentRankingClientItemVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IncidentRankingClientItemVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IncidentRankingClientItemVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IncidentRankingClientItemVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *IncidentRankingClientItemVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IncidentRankingClientItemVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IncidentRankingClientItemVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IncidentRankingClientItemVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


