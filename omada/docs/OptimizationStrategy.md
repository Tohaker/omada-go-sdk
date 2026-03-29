# OptimizationStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptimizationStrategy** | **int32** | Strategy of WLAN Optimization. 0: Global Optimization. 1: Optimization Adjustment, when WLAN Optimization is performed for the first time, that is, when API [get RF Planning Deploy History] returns ‘true’, this option is equivalent to Global Optimization. | 

## Methods

### NewOptimizationStrategy

`func NewOptimizationStrategy(optimizationStrategy int32, ) *OptimizationStrategy`

NewOptimizationStrategy instantiates a new OptimizationStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptimizationStrategyWithDefaults

`func NewOptimizationStrategyWithDefaults() *OptimizationStrategy`

NewOptimizationStrategyWithDefaults instantiates a new OptimizationStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptimizationStrategy

`func (o *OptimizationStrategy) GetOptimizationStrategy() int32`

GetOptimizationStrategy returns the OptimizationStrategy field if non-nil, zero value otherwise.

### GetOptimizationStrategyOk

`func (o *OptimizationStrategy) GetOptimizationStrategyOk() (*int32, bool)`

GetOptimizationStrategyOk returns a tuple with the OptimizationStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizationStrategy

`func (o *OptimizationStrategy) SetOptimizationStrategy(v int32)`

SetOptimizationStrategy sets OptimizationStrategy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


