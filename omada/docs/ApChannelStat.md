# ApChannelStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApNum** | Pointer to **int32** | AP number of the channel | [optional] 
**Channel** | Pointer to **int32** | This entry indicates the channel index. For example, 36 means the 36 / 2160MHz channel. | [optional] 
**ChannelUtilization** | Pointer to **float64** | The average utilization of the channel by APs in percentage. Null means no data. | [optional] 
**ClientNum** | Pointer to **int32** | Client number of the channel | [optional] 

## Methods

### NewApChannelStat

`func NewApChannelStat() *ApChannelStat`

NewApChannelStat instantiates a new ApChannelStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApChannelStatWithDefaults

`func NewApChannelStatWithDefaults() *ApChannelStat`

NewApChannelStatWithDefaults instantiates a new ApChannelStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApNum

`func (o *ApChannelStat) GetApNum() int32`

GetApNum returns the ApNum field if non-nil, zero value otherwise.

### GetApNumOk

`func (o *ApChannelStat) GetApNumOk() (*int32, bool)`

GetApNumOk returns a tuple with the ApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApNum

`func (o *ApChannelStat) SetApNum(v int32)`

SetApNum sets ApNum field to given value.

### HasApNum

`func (o *ApChannelStat) HasApNum() bool`

HasApNum returns a boolean if a field has been set.

### GetChannel

`func (o *ApChannelStat) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApChannelStat) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApChannelStat) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApChannelStat) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetChannelUtilization

`func (o *ApChannelStat) GetChannelUtilization() float64`

GetChannelUtilization returns the ChannelUtilization field if non-nil, zero value otherwise.

### GetChannelUtilizationOk

`func (o *ApChannelStat) GetChannelUtilizationOk() (*float64, bool)`

GetChannelUtilizationOk returns a tuple with the ChannelUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelUtilization

`func (o *ApChannelStat) SetChannelUtilization(v float64)`

SetChannelUtilization sets ChannelUtilization field to given value.

### HasChannelUtilization

`func (o *ApChannelStat) HasChannelUtilization() bool`

HasChannelUtilization returns a boolean if a field has been set.

### GetClientNum

`func (o *ApChannelStat) GetClientNum() int32`

GetClientNum returns the ClientNum field if non-nil, zero value otherwise.

### GetClientNumOk

`func (o *ApChannelStat) GetClientNumOk() (*int32, bool)`

GetClientNumOk returns a tuple with the ClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNum

`func (o *ApChannelStat) SetClientNum(v int32)`

SetClientNum sets ClientNum field to given value.

### HasClientNum

`func (o *ApChannelStat) HasClientNum() bool`

HasClientNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


