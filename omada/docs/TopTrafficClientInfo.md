# TopTrafficClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Client MAC. | [optional] 
**Name** | Pointer to **string** | Client name. | [optional] 
**Traffic** | Pointer to **int64** | Client traffic, unit is Byte. | [optional] 
**TrafficPercent** | Pointer to **int32** | Percentage of this client&#39;s traffic to total traffic | [optional] 
**Type** | Pointer to **string** | Client type. | [optional] 

## Methods

### NewTopTrafficClientInfo

`func NewTopTrafficClientInfo() *TopTrafficClientInfo`

NewTopTrafficClientInfo instantiates a new TopTrafficClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopTrafficClientInfoWithDefaults

`func NewTopTrafficClientInfoWithDefaults() *TopTrafficClientInfo`

NewTopTrafficClientInfoWithDefaults instantiates a new TopTrafficClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *TopTrafficClientInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopTrafficClientInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopTrafficClientInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopTrafficClientInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *TopTrafficClientInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopTrafficClientInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopTrafficClientInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopTrafficClientInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTraffic

`func (o *TopTrafficClientInfo) GetTraffic() int64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *TopTrafficClientInfo) GetTrafficOk() (*int64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *TopTrafficClientInfo) SetTraffic(v int64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *TopTrafficClientInfo) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetTrafficPercent

`func (o *TopTrafficClientInfo) GetTrafficPercent() int32`

GetTrafficPercent returns the TrafficPercent field if non-nil, zero value otherwise.

### GetTrafficPercentOk

`func (o *TopTrafficClientInfo) GetTrafficPercentOk() (*int32, bool)`

GetTrafficPercentOk returns a tuple with the TrafficPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPercent

`func (o *TopTrafficClientInfo) SetTrafficPercent(v int32)`

SetTrafficPercent sets TrafficPercent field to given value.

### HasTrafficPercent

`func (o *TopTrafficClientInfo) HasTrafficPercent() bool`

HasTrafficPercent returns a boolean if a field has been set.

### GetType

`func (o *TopTrafficClientInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopTrafficClientInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopTrafficClientInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopTrafficClientInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


