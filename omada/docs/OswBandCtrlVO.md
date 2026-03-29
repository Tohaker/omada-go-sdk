# OswBandCtrlVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EgressEnable** | **bool** | Indicates whether egress is enabled | 
**EgressLimit** | Pointer to **int32** | Egress Limit, when the user-configured value is much larger than the actual maximum value, the device will process it at the maximum value | [optional] 
**EgressUnit** | Pointer to **int32** | Egress Unit should be a value as follows: 1: Kbps; 2: Mbps | [optional] 
**IngressEnable** | **bool** | Indicates whether egress is enabled | 
**IngressLimit** | Pointer to **int32** | Ingress Limit, when the user-configured value is much larger than the actual maximum value, the device will process it at the maximum value | [optional] 
**IngressUnit** | Pointer to **int32** | Ingress Unit should be a value as follows: 1: Kbps; 2: Mbps | [optional] 
**LimitRange** | Pointer to [**[]OswBandCtrlLimitRangeVO**](OswBandCtrlLimitRangeVO.md) |  | [optional] 
**UnitSupport** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewOswBandCtrlVO

`func NewOswBandCtrlVO(egressEnable bool, ingressEnable bool, ) *OswBandCtrlVO`

NewOswBandCtrlVO instantiates a new OswBandCtrlVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswBandCtrlVOWithDefaults

`func NewOswBandCtrlVOWithDefaults() *OswBandCtrlVO`

NewOswBandCtrlVOWithDefaults instantiates a new OswBandCtrlVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEgressEnable

`func (o *OswBandCtrlVO) GetEgressEnable() bool`

GetEgressEnable returns the EgressEnable field if non-nil, zero value otherwise.

### GetEgressEnableOk

`func (o *OswBandCtrlVO) GetEgressEnableOk() (*bool, bool)`

GetEgressEnableOk returns a tuple with the EgressEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressEnable

`func (o *OswBandCtrlVO) SetEgressEnable(v bool)`

SetEgressEnable sets EgressEnable field to given value.


### GetEgressLimit

`func (o *OswBandCtrlVO) GetEgressLimit() int32`

GetEgressLimit returns the EgressLimit field if non-nil, zero value otherwise.

### GetEgressLimitOk

`func (o *OswBandCtrlVO) GetEgressLimitOk() (*int32, bool)`

GetEgressLimitOk returns a tuple with the EgressLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressLimit

`func (o *OswBandCtrlVO) SetEgressLimit(v int32)`

SetEgressLimit sets EgressLimit field to given value.

### HasEgressLimit

`func (o *OswBandCtrlVO) HasEgressLimit() bool`

HasEgressLimit returns a boolean if a field has been set.

### GetEgressUnit

`func (o *OswBandCtrlVO) GetEgressUnit() int32`

GetEgressUnit returns the EgressUnit field if non-nil, zero value otherwise.

### GetEgressUnitOk

`func (o *OswBandCtrlVO) GetEgressUnitOk() (*int32, bool)`

GetEgressUnitOk returns a tuple with the EgressUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressUnit

`func (o *OswBandCtrlVO) SetEgressUnit(v int32)`

SetEgressUnit sets EgressUnit field to given value.

### HasEgressUnit

`func (o *OswBandCtrlVO) HasEgressUnit() bool`

HasEgressUnit returns a boolean if a field has been set.

### GetIngressEnable

`func (o *OswBandCtrlVO) GetIngressEnable() bool`

GetIngressEnable returns the IngressEnable field if non-nil, zero value otherwise.

### GetIngressEnableOk

`func (o *OswBandCtrlVO) GetIngressEnableOk() (*bool, bool)`

GetIngressEnableOk returns a tuple with the IngressEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressEnable

`func (o *OswBandCtrlVO) SetIngressEnable(v bool)`

SetIngressEnable sets IngressEnable field to given value.


### GetIngressLimit

`func (o *OswBandCtrlVO) GetIngressLimit() int32`

GetIngressLimit returns the IngressLimit field if non-nil, zero value otherwise.

### GetIngressLimitOk

`func (o *OswBandCtrlVO) GetIngressLimitOk() (*int32, bool)`

GetIngressLimitOk returns a tuple with the IngressLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressLimit

`func (o *OswBandCtrlVO) SetIngressLimit(v int32)`

SetIngressLimit sets IngressLimit field to given value.

### HasIngressLimit

`func (o *OswBandCtrlVO) HasIngressLimit() bool`

HasIngressLimit returns a boolean if a field has been set.

### GetIngressUnit

`func (o *OswBandCtrlVO) GetIngressUnit() int32`

GetIngressUnit returns the IngressUnit field if non-nil, zero value otherwise.

### GetIngressUnitOk

`func (o *OswBandCtrlVO) GetIngressUnitOk() (*int32, bool)`

GetIngressUnitOk returns a tuple with the IngressUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressUnit

`func (o *OswBandCtrlVO) SetIngressUnit(v int32)`

SetIngressUnit sets IngressUnit field to given value.

### HasIngressUnit

`func (o *OswBandCtrlVO) HasIngressUnit() bool`

HasIngressUnit returns a boolean if a field has been set.

### GetLimitRange

`func (o *OswBandCtrlVO) GetLimitRange() []OswBandCtrlLimitRangeVO`

GetLimitRange returns the LimitRange field if non-nil, zero value otherwise.

### GetLimitRangeOk

`func (o *OswBandCtrlVO) GetLimitRangeOk() (*[]OswBandCtrlLimitRangeVO, bool)`

GetLimitRangeOk returns a tuple with the LimitRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitRange

`func (o *OswBandCtrlVO) SetLimitRange(v []OswBandCtrlLimitRangeVO)`

SetLimitRange sets LimitRange field to given value.

### HasLimitRange

`func (o *OswBandCtrlVO) HasLimitRange() bool`

HasLimitRange returns a boolean if a field has been set.

### GetUnitSupport

`func (o *OswBandCtrlVO) GetUnitSupport() []int32`

GetUnitSupport returns the UnitSupport field if non-nil, zero value otherwise.

### GetUnitSupportOk

`func (o *OswBandCtrlVO) GetUnitSupportOk() (*[]int32, bool)`

GetUnitSupportOk returns a tuple with the UnitSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitSupport

`func (o *OswBandCtrlVO) SetUnitSupport(v []int32)`

SetUnitSupport sets UnitSupport field to given value.

### HasUnitSupport

`func (o *OswBandCtrlVO) HasUnitSupport() bool`

HasUnitSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


