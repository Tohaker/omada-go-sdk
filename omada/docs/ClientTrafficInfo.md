# ClientTrafficInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients2gNum** | Pointer to **int32** | The number of wireless clients in the 2G band. | [optional] 
**Clients2gTraffic** | Pointer to **int64** | Total traffic of wireless clients in the 2G band. | [optional] 
**Clients5gNum** | Pointer to **int32** | The number of wireless clients in the 5G band. | [optional] 
**Clients5gTraffic** | Pointer to **int64** | Total traffic of wireless clients in the 5G band. | [optional] 
**Clients6gNum** | Pointer to **int32** | The number of wireless clients in the 6G band. | [optional] 
**Clients6gTraffic** | Pointer to **int64** | Total traffic of wireless clients in the 6G band. | [optional] 
**TotalTraffic** | Pointer to **int64** | Total traffic of all clients. | [optional] 
**WiredClientsNum** | Pointer to **int32** | The number of all wired clients. | [optional] 
**WiredClientsTraffic** | Pointer to **int64** | Total traffic of all wired clients. | [optional] 

## Methods

### NewClientTrafficInfo

`func NewClientTrafficInfo() *ClientTrafficInfo`

NewClientTrafficInfo instantiates a new ClientTrafficInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientTrafficInfoWithDefaults

`func NewClientTrafficInfoWithDefaults() *ClientTrafficInfo`

NewClientTrafficInfoWithDefaults instantiates a new ClientTrafficInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients2gNum

`func (o *ClientTrafficInfo) GetClients2gNum() int32`

GetClients2gNum returns the Clients2gNum field if non-nil, zero value otherwise.

### GetClients2gNumOk

`func (o *ClientTrafficInfo) GetClients2gNumOk() (*int32, bool)`

GetClients2gNumOk returns a tuple with the Clients2gNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients2gNum

`func (o *ClientTrafficInfo) SetClients2gNum(v int32)`

SetClients2gNum sets Clients2gNum field to given value.

### HasClients2gNum

`func (o *ClientTrafficInfo) HasClients2gNum() bool`

HasClients2gNum returns a boolean if a field has been set.

### GetClients2gTraffic

`func (o *ClientTrafficInfo) GetClients2gTraffic() int64`

GetClients2gTraffic returns the Clients2gTraffic field if non-nil, zero value otherwise.

### GetClients2gTrafficOk

`func (o *ClientTrafficInfo) GetClients2gTrafficOk() (*int64, bool)`

GetClients2gTrafficOk returns a tuple with the Clients2gTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients2gTraffic

`func (o *ClientTrafficInfo) SetClients2gTraffic(v int64)`

SetClients2gTraffic sets Clients2gTraffic field to given value.

### HasClients2gTraffic

`func (o *ClientTrafficInfo) HasClients2gTraffic() bool`

HasClients2gTraffic returns a boolean if a field has been set.

### GetClients5gNum

`func (o *ClientTrafficInfo) GetClients5gNum() int32`

GetClients5gNum returns the Clients5gNum field if non-nil, zero value otherwise.

### GetClients5gNumOk

`func (o *ClientTrafficInfo) GetClients5gNumOk() (*int32, bool)`

GetClients5gNumOk returns a tuple with the Clients5gNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients5gNum

`func (o *ClientTrafficInfo) SetClients5gNum(v int32)`

SetClients5gNum sets Clients5gNum field to given value.

### HasClients5gNum

`func (o *ClientTrafficInfo) HasClients5gNum() bool`

HasClients5gNum returns a boolean if a field has been set.

### GetClients5gTraffic

`func (o *ClientTrafficInfo) GetClients5gTraffic() int64`

GetClients5gTraffic returns the Clients5gTraffic field if non-nil, zero value otherwise.

### GetClients5gTrafficOk

`func (o *ClientTrafficInfo) GetClients5gTrafficOk() (*int64, bool)`

GetClients5gTrafficOk returns a tuple with the Clients5gTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients5gTraffic

`func (o *ClientTrafficInfo) SetClients5gTraffic(v int64)`

SetClients5gTraffic sets Clients5gTraffic field to given value.

### HasClients5gTraffic

`func (o *ClientTrafficInfo) HasClients5gTraffic() bool`

HasClients5gTraffic returns a boolean if a field has been set.

### GetClients6gNum

`func (o *ClientTrafficInfo) GetClients6gNum() int32`

GetClients6gNum returns the Clients6gNum field if non-nil, zero value otherwise.

### GetClients6gNumOk

`func (o *ClientTrafficInfo) GetClients6gNumOk() (*int32, bool)`

GetClients6gNumOk returns a tuple with the Clients6gNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients6gNum

`func (o *ClientTrafficInfo) SetClients6gNum(v int32)`

SetClients6gNum sets Clients6gNum field to given value.

### HasClients6gNum

`func (o *ClientTrafficInfo) HasClients6gNum() bool`

HasClients6gNum returns a boolean if a field has been set.

### GetClients6gTraffic

`func (o *ClientTrafficInfo) GetClients6gTraffic() int64`

GetClients6gTraffic returns the Clients6gTraffic field if non-nil, zero value otherwise.

### GetClients6gTrafficOk

`func (o *ClientTrafficInfo) GetClients6gTrafficOk() (*int64, bool)`

GetClients6gTrafficOk returns a tuple with the Clients6gTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients6gTraffic

`func (o *ClientTrafficInfo) SetClients6gTraffic(v int64)`

SetClients6gTraffic sets Clients6gTraffic field to given value.

### HasClients6gTraffic

`func (o *ClientTrafficInfo) HasClients6gTraffic() bool`

HasClients6gTraffic returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *ClientTrafficInfo) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *ClientTrafficInfo) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *ClientTrafficInfo) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *ClientTrafficInfo) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetWiredClientsNum

`func (o *ClientTrafficInfo) GetWiredClientsNum() int32`

GetWiredClientsNum returns the WiredClientsNum field if non-nil, zero value otherwise.

### GetWiredClientsNumOk

`func (o *ClientTrafficInfo) GetWiredClientsNumOk() (*int32, bool)`

GetWiredClientsNumOk returns a tuple with the WiredClientsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredClientsNum

`func (o *ClientTrafficInfo) SetWiredClientsNum(v int32)`

SetWiredClientsNum sets WiredClientsNum field to given value.

### HasWiredClientsNum

`func (o *ClientTrafficInfo) HasWiredClientsNum() bool`

HasWiredClientsNum returns a boolean if a field has been set.

### GetWiredClientsTraffic

`func (o *ClientTrafficInfo) GetWiredClientsTraffic() int64`

GetWiredClientsTraffic returns the WiredClientsTraffic field if non-nil, zero value otherwise.

### GetWiredClientsTrafficOk

`func (o *ClientTrafficInfo) GetWiredClientsTrafficOk() (*int64, bool)`

GetWiredClientsTrafficOk returns a tuple with the WiredClientsTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredClientsTraffic

`func (o *ClientTrafficInfo) SetWiredClientsTraffic(v int64)`

SetWiredClientsTraffic sets WiredClientsTraffic field to given value.

### HasWiredClientsTraffic

`func (o *ClientTrafficInfo) HasWiredClientsTraffic() bool`

HasWiredClientsTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


