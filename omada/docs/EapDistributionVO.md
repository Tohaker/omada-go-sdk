# EapDistributionVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eaps** | Pointer to [**[]EapClientVO**](EapClientVO.md) |  | [optional] 
**Others** | Pointer to [**EapClientVO**](EapClientVO.md) |  | [optional] 
**TotalEapClients** | Pointer to **int32** |  | [optional] 
**TotalEapDistribution** | Pointer to **float64** |  | [optional] 

## Methods

### NewEapDistributionVO

`func NewEapDistributionVO() *EapDistributionVO`

NewEapDistributionVO instantiates a new EapDistributionVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEapDistributionVOWithDefaults

`func NewEapDistributionVOWithDefaults() *EapDistributionVO`

NewEapDistributionVOWithDefaults instantiates a new EapDistributionVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEaps

`func (o *EapDistributionVO) GetEaps() []EapClientVO`

GetEaps returns the Eaps field if non-nil, zero value otherwise.

### GetEapsOk

`func (o *EapDistributionVO) GetEapsOk() (*[]EapClientVO, bool)`

GetEapsOk returns a tuple with the Eaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaps

`func (o *EapDistributionVO) SetEaps(v []EapClientVO)`

SetEaps sets Eaps field to given value.

### HasEaps

`func (o *EapDistributionVO) HasEaps() bool`

HasEaps returns a boolean if a field has been set.

### GetOthers

`func (o *EapDistributionVO) GetOthers() EapClientVO`

GetOthers returns the Others field if non-nil, zero value otherwise.

### GetOthersOk

`func (o *EapDistributionVO) GetOthersOk() (*EapClientVO, bool)`

GetOthersOk returns a tuple with the Others field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOthers

`func (o *EapDistributionVO) SetOthers(v EapClientVO)`

SetOthers sets Others field to given value.

### HasOthers

`func (o *EapDistributionVO) HasOthers() bool`

HasOthers returns a boolean if a field has been set.

### GetTotalEapClients

`func (o *EapDistributionVO) GetTotalEapClients() int32`

GetTotalEapClients returns the TotalEapClients field if non-nil, zero value otherwise.

### GetTotalEapClientsOk

`func (o *EapDistributionVO) GetTotalEapClientsOk() (*int32, bool)`

GetTotalEapClientsOk returns a tuple with the TotalEapClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEapClients

`func (o *EapDistributionVO) SetTotalEapClients(v int32)`

SetTotalEapClients sets TotalEapClients field to given value.

### HasTotalEapClients

`func (o *EapDistributionVO) HasTotalEapClients() bool`

HasTotalEapClients returns a boolean if a field has been set.

### GetTotalEapDistribution

`func (o *EapDistributionVO) GetTotalEapDistribution() float64`

GetTotalEapDistribution returns the TotalEapDistribution field if non-nil, zero value otherwise.

### GetTotalEapDistributionOk

`func (o *EapDistributionVO) GetTotalEapDistributionOk() (*float64, bool)`

GetTotalEapDistributionOk returns a tuple with the TotalEapDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEapDistribution

`func (o *EapDistributionVO) SetTotalEapDistribution(v float64)`

SetTotalEapDistribution sets TotalEapDistribution field to given value.

### HasTotalEapDistribution

`func (o *EapDistributionVO) HasTotalEapDistribution() bool`

HasTotalEapDistribution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


