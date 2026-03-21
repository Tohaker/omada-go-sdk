# LongestUptimeClientInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Client MAC. | [optional] 
**Name** | Pointer to **string** | Client name. | [optional] 
**TotalDuration** | Pointer to **int64** | Client uptime, unit is second. | [optional] 
**Type** | Pointer to **string** | Client type. | [optional] 

## Methods

### NewLongestUptimeClientInfo

`func NewLongestUptimeClientInfo() *LongestUptimeClientInfo`

NewLongestUptimeClientInfo instantiates a new LongestUptimeClientInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLongestUptimeClientInfoWithDefaults

`func NewLongestUptimeClientInfoWithDefaults() *LongestUptimeClientInfo`

NewLongestUptimeClientInfoWithDefaults instantiates a new LongestUptimeClientInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *LongestUptimeClientInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *LongestUptimeClientInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *LongestUptimeClientInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *LongestUptimeClientInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *LongestUptimeClientInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LongestUptimeClientInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LongestUptimeClientInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LongestUptimeClientInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotalDuration

`func (o *LongestUptimeClientInfo) GetTotalDuration() int64`

GetTotalDuration returns the TotalDuration field if non-nil, zero value otherwise.

### GetTotalDurationOk

`func (o *LongestUptimeClientInfo) GetTotalDurationOk() (*int64, bool)`

GetTotalDurationOk returns a tuple with the TotalDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDuration

`func (o *LongestUptimeClientInfo) SetTotalDuration(v int64)`

SetTotalDuration sets TotalDuration field to given value.

### HasTotalDuration

`func (o *LongestUptimeClientInfo) HasTotalDuration() bool`

HasTotalDuration returns a boolean if a field has been set.

### GetType

`func (o *LongestUptimeClientInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LongestUptimeClientInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LongestUptimeClientInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LongestUptimeClientInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


