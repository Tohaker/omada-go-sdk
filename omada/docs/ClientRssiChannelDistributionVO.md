# ClientRssiChannelDistributionVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distribution2G** | Pointer to [**ClientsRssiDistributionVO**](ClientsRssiDistributionVO.md) |  | [optional] 
**Distribution5G** | Pointer to [**ClientsRssiDistributionVO**](ClientsRssiDistributionVO.md) |  | [optional] 
**Distribution6G** | Pointer to [**ClientsRssiDistributionVO**](ClientsRssiDistributionVO.md) |  | [optional] 
**From55To45** | Pointer to **int32** |  | [optional] 
**From65To55** | Pointer to **int32** |  | [optional] 
**From71To65** | Pointer to **int32** |  | [optional] 
**LessThan72** | Pointer to **int32** |  | [optional] 
**MoreThan45** | Pointer to **int32** |  | [optional] 

## Methods

### NewClientRssiChannelDistributionVO

`func NewClientRssiChannelDistributionVO() *ClientRssiChannelDistributionVO`

NewClientRssiChannelDistributionVO instantiates a new ClientRssiChannelDistributionVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientRssiChannelDistributionVOWithDefaults

`func NewClientRssiChannelDistributionVOWithDefaults() *ClientRssiChannelDistributionVO`

NewClientRssiChannelDistributionVOWithDefaults instantiates a new ClientRssiChannelDistributionVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistribution2G

`func (o *ClientRssiChannelDistributionVO) GetDistribution2G() ClientsRssiDistributionVO`

GetDistribution2G returns the Distribution2G field if non-nil, zero value otherwise.

### GetDistribution2GOk

`func (o *ClientRssiChannelDistributionVO) GetDistribution2GOk() (*ClientsRssiDistributionVO, bool)`

GetDistribution2GOk returns a tuple with the Distribution2G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution2G

`func (o *ClientRssiChannelDistributionVO) SetDistribution2G(v ClientsRssiDistributionVO)`

SetDistribution2G sets Distribution2G field to given value.

### HasDistribution2G

`func (o *ClientRssiChannelDistributionVO) HasDistribution2G() bool`

HasDistribution2G returns a boolean if a field has been set.

### GetDistribution5G

`func (o *ClientRssiChannelDistributionVO) GetDistribution5G() ClientsRssiDistributionVO`

GetDistribution5G returns the Distribution5G field if non-nil, zero value otherwise.

### GetDistribution5GOk

`func (o *ClientRssiChannelDistributionVO) GetDistribution5GOk() (*ClientsRssiDistributionVO, bool)`

GetDistribution5GOk returns a tuple with the Distribution5G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution5G

`func (o *ClientRssiChannelDistributionVO) SetDistribution5G(v ClientsRssiDistributionVO)`

SetDistribution5G sets Distribution5G field to given value.

### HasDistribution5G

`func (o *ClientRssiChannelDistributionVO) HasDistribution5G() bool`

HasDistribution5G returns a boolean if a field has been set.

### GetDistribution6G

`func (o *ClientRssiChannelDistributionVO) GetDistribution6G() ClientsRssiDistributionVO`

GetDistribution6G returns the Distribution6G field if non-nil, zero value otherwise.

### GetDistribution6GOk

`func (o *ClientRssiChannelDistributionVO) GetDistribution6GOk() (*ClientsRssiDistributionVO, bool)`

GetDistribution6GOk returns a tuple with the Distribution6G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistribution6G

`func (o *ClientRssiChannelDistributionVO) SetDistribution6G(v ClientsRssiDistributionVO)`

SetDistribution6G sets Distribution6G field to given value.

### HasDistribution6G

`func (o *ClientRssiChannelDistributionVO) HasDistribution6G() bool`

HasDistribution6G returns a boolean if a field has been set.

### GetFrom55To45

`func (o *ClientRssiChannelDistributionVO) GetFrom55To45() int32`

GetFrom55To45 returns the From55To45 field if non-nil, zero value otherwise.

### GetFrom55To45Ok

`func (o *ClientRssiChannelDistributionVO) GetFrom55To45Ok() (*int32, bool)`

GetFrom55To45Ok returns a tuple with the From55To45 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom55To45

`func (o *ClientRssiChannelDistributionVO) SetFrom55To45(v int32)`

SetFrom55To45 sets From55To45 field to given value.

### HasFrom55To45

`func (o *ClientRssiChannelDistributionVO) HasFrom55To45() bool`

HasFrom55To45 returns a boolean if a field has been set.

### GetFrom65To55

`func (o *ClientRssiChannelDistributionVO) GetFrom65To55() int32`

GetFrom65To55 returns the From65To55 field if non-nil, zero value otherwise.

### GetFrom65To55Ok

`func (o *ClientRssiChannelDistributionVO) GetFrom65To55Ok() (*int32, bool)`

GetFrom65To55Ok returns a tuple with the From65To55 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom65To55

`func (o *ClientRssiChannelDistributionVO) SetFrom65To55(v int32)`

SetFrom65To55 sets From65To55 field to given value.

### HasFrom65To55

`func (o *ClientRssiChannelDistributionVO) HasFrom65To55() bool`

HasFrom65To55 returns a boolean if a field has been set.

### GetFrom71To65

`func (o *ClientRssiChannelDistributionVO) GetFrom71To65() int32`

GetFrom71To65 returns the From71To65 field if non-nil, zero value otherwise.

### GetFrom71To65Ok

`func (o *ClientRssiChannelDistributionVO) GetFrom71To65Ok() (*int32, bool)`

GetFrom71To65Ok returns a tuple with the From71To65 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom71To65

`func (o *ClientRssiChannelDistributionVO) SetFrom71To65(v int32)`

SetFrom71To65 sets From71To65 field to given value.

### HasFrom71To65

`func (o *ClientRssiChannelDistributionVO) HasFrom71To65() bool`

HasFrom71To65 returns a boolean if a field has been set.

### GetLessThan72

`func (o *ClientRssiChannelDistributionVO) GetLessThan72() int32`

GetLessThan72 returns the LessThan72 field if non-nil, zero value otherwise.

### GetLessThan72Ok

`func (o *ClientRssiChannelDistributionVO) GetLessThan72Ok() (*int32, bool)`

GetLessThan72Ok returns a tuple with the LessThan72 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLessThan72

`func (o *ClientRssiChannelDistributionVO) SetLessThan72(v int32)`

SetLessThan72 sets LessThan72 field to given value.

### HasLessThan72

`func (o *ClientRssiChannelDistributionVO) HasLessThan72() bool`

HasLessThan72 returns a boolean if a field has been set.

### GetMoreThan45

`func (o *ClientRssiChannelDistributionVO) GetMoreThan45() int32`

GetMoreThan45 returns the MoreThan45 field if non-nil, zero value otherwise.

### GetMoreThan45Ok

`func (o *ClientRssiChannelDistributionVO) GetMoreThan45Ok() (*int32, bool)`

GetMoreThan45Ok returns a tuple with the MoreThan45 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreThan45

`func (o *ClientRssiChannelDistributionVO) SetMoreThan45(v int32)`

SetMoreThan45 sets MoreThan45 field to given value.

### HasMoreThan45

`func (o *ClientRssiChannelDistributionVO) HasMoreThan45() bool`

HasMoreThan45 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


