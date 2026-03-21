# OsgPortBandCtrlVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EgressEnable** | **bool** |  | 
**EgressLimit** | Pointer to **int32** |  | [optional] 
**IngressEnable** | **bool** |  | 
**IngressLimit** | Pointer to **int32** |  | [optional] 

## Methods

### NewOsgPortBandCtrlVO

`func NewOsgPortBandCtrlVO(egressEnable bool, ingressEnable bool, ) *OsgPortBandCtrlVO`

NewOsgPortBandCtrlVO instantiates a new OsgPortBandCtrlVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgPortBandCtrlVOWithDefaults

`func NewOsgPortBandCtrlVOWithDefaults() *OsgPortBandCtrlVO`

NewOsgPortBandCtrlVOWithDefaults instantiates a new OsgPortBandCtrlVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEgressEnable

`func (o *OsgPortBandCtrlVO) GetEgressEnable() bool`

GetEgressEnable returns the EgressEnable field if non-nil, zero value otherwise.

### GetEgressEnableOk

`func (o *OsgPortBandCtrlVO) GetEgressEnableOk() (*bool, bool)`

GetEgressEnableOk returns a tuple with the EgressEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressEnable

`func (o *OsgPortBandCtrlVO) SetEgressEnable(v bool)`

SetEgressEnable sets EgressEnable field to given value.


### GetEgressLimit

`func (o *OsgPortBandCtrlVO) GetEgressLimit() int32`

GetEgressLimit returns the EgressLimit field if non-nil, zero value otherwise.

### GetEgressLimitOk

`func (o *OsgPortBandCtrlVO) GetEgressLimitOk() (*int32, bool)`

GetEgressLimitOk returns a tuple with the EgressLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressLimit

`func (o *OsgPortBandCtrlVO) SetEgressLimit(v int32)`

SetEgressLimit sets EgressLimit field to given value.

### HasEgressLimit

`func (o *OsgPortBandCtrlVO) HasEgressLimit() bool`

HasEgressLimit returns a boolean if a field has been set.

### GetIngressEnable

`func (o *OsgPortBandCtrlVO) GetIngressEnable() bool`

GetIngressEnable returns the IngressEnable field if non-nil, zero value otherwise.

### GetIngressEnableOk

`func (o *OsgPortBandCtrlVO) GetIngressEnableOk() (*bool, bool)`

GetIngressEnableOk returns a tuple with the IngressEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressEnable

`func (o *OsgPortBandCtrlVO) SetIngressEnable(v bool)`

SetIngressEnable sets IngressEnable field to given value.


### GetIngressLimit

`func (o *OsgPortBandCtrlVO) GetIngressLimit() int32`

GetIngressLimit returns the IngressLimit field if non-nil, zero value otherwise.

### GetIngressLimitOk

`func (o *OsgPortBandCtrlVO) GetIngressLimitOk() (*int32, bool)`

GetIngressLimitOk returns a tuple with the IngressLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressLimit

`func (o *OsgPortBandCtrlVO) SetIngressLimit(v int32)`

SetIngressLimit sets IngressLimit field to given value.

### HasIngressLimit

`func (o *OsgPortBandCtrlVO) HasIngressLimit() bool`

HasIngressLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


