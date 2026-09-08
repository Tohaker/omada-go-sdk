# ClientHealthCategoryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to [**HealthDistributionVO**](HealthDistributionVO.md) |  | [optional] 
**Incident** | Pointer to [**IncidentVO**](IncidentVO.md) |  | [optional] 

## Methods

### NewClientHealthCategoryVO

`func NewClientHealthCategoryVO() *ClientHealthCategoryVO`

NewClientHealthCategoryVO instantiates a new ClientHealthCategoryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientHealthCategoryVOWithDefaults

`func NewClientHealthCategoryVOWithDefaults() *ClientHealthCategoryVO`

NewClientHealthCategoryVOWithDefaults instantiates a new ClientHealthCategoryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *ClientHealthCategoryVO) GetHealth() HealthDistributionVO`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ClientHealthCategoryVO) GetHealthOk() (*HealthDistributionVO, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ClientHealthCategoryVO) SetHealth(v HealthDistributionVO)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ClientHealthCategoryVO) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetIncident

`func (o *ClientHealthCategoryVO) GetIncident() IncidentVO`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *ClientHealthCategoryVO) GetIncidentOk() (*IncidentVO, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *ClientHealthCategoryVO) SetIncident(v IncidentVO)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *ClientHealthCategoryVO) HasIncident() bool`

HasIncident returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


