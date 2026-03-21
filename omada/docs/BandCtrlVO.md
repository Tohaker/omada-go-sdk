# BandCtrlVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EgressEnable** | **bool** | Egress enable status | 
**EgressLimit** | Pointer to **int32** | EgressLimit | [optional] 
**EgressUnit** | Pointer to **int32** | EgressUnit should be a value as follows: 1: Kbps and 2: Mbps | [optional] 
**IngressEnable** | **bool** | Ingress enable status | 
**IngressLimit** | Pointer to **int32** | IngressLimit | [optional] 
**IngressUnit** | Pointer to **int32** | IngressUnit should be a value as follows: 1: Kbps and 2: Mbps | [optional] 

## Methods

### NewBandCtrlVO

`func NewBandCtrlVO(egressEnable bool, ingressEnable bool, ) *BandCtrlVO`

NewBandCtrlVO instantiates a new BandCtrlVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandCtrlVOWithDefaults

`func NewBandCtrlVOWithDefaults() *BandCtrlVO`

NewBandCtrlVOWithDefaults instantiates a new BandCtrlVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEgressEnable

`func (o *BandCtrlVO) GetEgressEnable() bool`

GetEgressEnable returns the EgressEnable field if non-nil, zero value otherwise.

### GetEgressEnableOk

`func (o *BandCtrlVO) GetEgressEnableOk() (*bool, bool)`

GetEgressEnableOk returns a tuple with the EgressEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressEnable

`func (o *BandCtrlVO) SetEgressEnable(v bool)`

SetEgressEnable sets EgressEnable field to given value.


### GetEgressLimit

`func (o *BandCtrlVO) GetEgressLimit() int32`

GetEgressLimit returns the EgressLimit field if non-nil, zero value otherwise.

### GetEgressLimitOk

`func (o *BandCtrlVO) GetEgressLimitOk() (*int32, bool)`

GetEgressLimitOk returns a tuple with the EgressLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressLimit

`func (o *BandCtrlVO) SetEgressLimit(v int32)`

SetEgressLimit sets EgressLimit field to given value.

### HasEgressLimit

`func (o *BandCtrlVO) HasEgressLimit() bool`

HasEgressLimit returns a boolean if a field has been set.

### GetEgressUnit

`func (o *BandCtrlVO) GetEgressUnit() int32`

GetEgressUnit returns the EgressUnit field if non-nil, zero value otherwise.

### GetEgressUnitOk

`func (o *BandCtrlVO) GetEgressUnitOk() (*int32, bool)`

GetEgressUnitOk returns a tuple with the EgressUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressUnit

`func (o *BandCtrlVO) SetEgressUnit(v int32)`

SetEgressUnit sets EgressUnit field to given value.

### HasEgressUnit

`func (o *BandCtrlVO) HasEgressUnit() bool`

HasEgressUnit returns a boolean if a field has been set.

### GetIngressEnable

`func (o *BandCtrlVO) GetIngressEnable() bool`

GetIngressEnable returns the IngressEnable field if non-nil, zero value otherwise.

### GetIngressEnableOk

`func (o *BandCtrlVO) GetIngressEnableOk() (*bool, bool)`

GetIngressEnableOk returns a tuple with the IngressEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressEnable

`func (o *BandCtrlVO) SetIngressEnable(v bool)`

SetIngressEnable sets IngressEnable field to given value.


### GetIngressLimit

`func (o *BandCtrlVO) GetIngressLimit() int32`

GetIngressLimit returns the IngressLimit field if non-nil, zero value otherwise.

### GetIngressLimitOk

`func (o *BandCtrlVO) GetIngressLimitOk() (*int32, bool)`

GetIngressLimitOk returns a tuple with the IngressLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressLimit

`func (o *BandCtrlVO) SetIngressLimit(v int32)`

SetIngressLimit sets IngressLimit field to given value.

### HasIngressLimit

`func (o *BandCtrlVO) HasIngressLimit() bool`

HasIngressLimit returns a boolean if a field has been set.

### GetIngressUnit

`func (o *BandCtrlVO) GetIngressUnit() int32`

GetIngressUnit returns the IngressUnit field if non-nil, zero value otherwise.

### GetIngressUnitOk

`func (o *BandCtrlVO) GetIngressUnitOk() (*int32, bool)`

GetIngressUnitOk returns a tuple with the IngressUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressUnit

`func (o *BandCtrlVO) SetIngressUnit(v int32)`

SetIngressUnit sets IngressUnit field to given value.

### HasIngressUnit

`func (o *BandCtrlVO) HasIngressUnit() bool`

HasIngressUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


