# Interference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inf** | Pointer to **int32** | Interference should be within the range of -96 ~ -48 | [optional] 
**InfData** | Pointer to [**InterferenceDataEntity**](InterferenceDataEntity.md) |  | [optional] 
**InfType** | Pointer to **int32** | Interference Type should be a value as follows: 0: invalid parameter; 1: MWO; 2: CW; 3: WLAN; 4: FHSS. | [optional] 

## Methods

### NewInterference

`func NewInterference() *Interference`

NewInterference instantiates a new Interference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterferenceWithDefaults

`func NewInterferenceWithDefaults() *Interference`

NewInterferenceWithDefaults instantiates a new Interference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInf

`func (o *Interference) GetInf() int32`

GetInf returns the Inf field if non-nil, zero value otherwise.

### GetInfOk

`func (o *Interference) GetInfOk() (*int32, bool)`

GetInfOk returns a tuple with the Inf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInf

`func (o *Interference) SetInf(v int32)`

SetInf sets Inf field to given value.

### HasInf

`func (o *Interference) HasInf() bool`

HasInf returns a boolean if a field has been set.

### GetInfData

`func (o *Interference) GetInfData() InterferenceDataEntity`

GetInfData returns the InfData field if non-nil, zero value otherwise.

### GetInfDataOk

`func (o *Interference) GetInfDataOk() (*InterferenceDataEntity, bool)`

GetInfDataOk returns a tuple with the InfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfData

`func (o *Interference) SetInfData(v InterferenceDataEntity)`

SetInfData sets InfData field to given value.

### HasInfData

`func (o *Interference) HasInfData() bool`

HasInfData returns a boolean if a field has been set.

### GetInfType

`func (o *Interference) GetInfType() int32`

GetInfType returns the InfType field if non-nil, zero value otherwise.

### GetInfTypeOk

`func (o *Interference) GetInfTypeOk() (*int32, bool)`

GetInfTypeOk returns a tuple with the InfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfType

`func (o *Interference) SetInfType(v int32)`

SetInfType sets InfType field to given value.

### HasInfType

`func (o *Interference) HasInfType() bool`

HasInfType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


