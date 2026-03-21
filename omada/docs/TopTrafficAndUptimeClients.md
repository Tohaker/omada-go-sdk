# TopTrafficAndUptimeClients

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveClients** | Pointer to [**[]TopTrafficClientInfo**](TopTrafficClientInfo.md) | Top traffic clients. | [optional] 
**LongestUptimeClients** | Pointer to [**[]LongestUptimeClientInfo**](LongestUptimeClientInfo.md) | Longest uptime clients. | [optional] 

## Methods

### NewTopTrafficAndUptimeClients

`func NewTopTrafficAndUptimeClients() *TopTrafficAndUptimeClients`

NewTopTrafficAndUptimeClients instantiates a new TopTrafficAndUptimeClients object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopTrafficAndUptimeClientsWithDefaults

`func NewTopTrafficAndUptimeClientsWithDefaults() *TopTrafficAndUptimeClients`

NewTopTrafficAndUptimeClientsWithDefaults instantiates a new TopTrafficAndUptimeClients object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveClients

`func (o *TopTrafficAndUptimeClients) GetActiveClients() []TopTrafficClientInfo`

GetActiveClients returns the ActiveClients field if non-nil, zero value otherwise.

### GetActiveClientsOk

`func (o *TopTrafficAndUptimeClients) GetActiveClientsOk() (*[]TopTrafficClientInfo, bool)`

GetActiveClientsOk returns a tuple with the ActiveClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveClients

`func (o *TopTrafficAndUptimeClients) SetActiveClients(v []TopTrafficClientInfo)`

SetActiveClients sets ActiveClients field to given value.

### HasActiveClients

`func (o *TopTrafficAndUptimeClients) HasActiveClients() bool`

HasActiveClients returns a boolean if a field has been set.

### GetLongestUptimeClients

`func (o *TopTrafficAndUptimeClients) GetLongestUptimeClients() []LongestUptimeClientInfo`

GetLongestUptimeClients returns the LongestUptimeClients field if non-nil, zero value otherwise.

### GetLongestUptimeClientsOk

`func (o *TopTrafficAndUptimeClients) GetLongestUptimeClientsOk() (*[]LongestUptimeClientInfo, bool)`

GetLongestUptimeClientsOk returns a tuple with the LongestUptimeClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongestUptimeClients

`func (o *TopTrafficAndUptimeClients) SetLongestUptimeClients(v []LongestUptimeClientInfo)`

SetLongestUptimeClients sets LongestUptimeClients field to given value.

### HasLongestUptimeClients

`func (o *TopTrafficAndUptimeClients) HasLongestUptimeClients() bool`

HasLongestUptimeClients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


