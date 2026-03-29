# SSIDStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientsNum** | Pointer to **int32** | Client number. | [optional] 
**Ssid** | Pointer to **string** | SSID name. | [optional] 

## Methods

### NewSSIDStat

`func NewSSIDStat() *SSIDStat`

NewSSIDStat instantiates a new SSIDStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSIDStatWithDefaults

`func NewSSIDStatWithDefaults() *SSIDStat`

NewSSIDStatWithDefaults instantiates a new SSIDStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientsNum

`func (o *SSIDStat) GetClientsNum() int32`

GetClientsNum returns the ClientsNum field if non-nil, zero value otherwise.

### GetClientsNumOk

`func (o *SSIDStat) GetClientsNumOk() (*int32, bool)`

GetClientsNumOk returns a tuple with the ClientsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsNum

`func (o *SSIDStat) SetClientsNum(v int32)`

SetClientsNum sets ClientsNum field to given value.

### HasClientsNum

`func (o *SSIDStat) HasClientsNum() bool`

HasClientsNum returns a boolean if a field has been set.

### GetSsid

`func (o *SSIDStat) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *SSIDStat) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *SSIDStat) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *SSIDStat) HasSsid() bool`

HasSsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


