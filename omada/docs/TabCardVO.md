# TabCardVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopK** | Pointer to **int32** | topK: 5, 10, 20 | [optional] 
**Type** | Pointer to **string** | card type enum: overviewSummary,clientConnectionTrend,internet,network,alertSummary,gatewaySummary,ispLoad,switchStatus,switchAlertReboot,topSwitchCpuMemory,topSwitchByTrafficAndPoePower,poePowerTrend,wirelessTraffic,apStatus,topApByTrafficAndClient,topApByInterference,topApByRtAndDrop,topApByCpuAndMemory,topSsidByTraffic,clientsOverview, clientsAssociationActivities,clientsWithOnboardingTimes,topClient,appCategories,topApplicationByTraffic | [optional] 

## Methods

### NewTabCardVO

`func NewTabCardVO() *TabCardVO`

NewTabCardVO instantiates a new TabCardVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTabCardVOWithDefaults

`func NewTabCardVOWithDefaults() *TabCardVO`

NewTabCardVOWithDefaults instantiates a new TabCardVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopK

`func (o *TabCardVO) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *TabCardVO) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *TabCardVO) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *TabCardVO) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetType

`func (o *TabCardVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TabCardVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TabCardVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TabCardVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


