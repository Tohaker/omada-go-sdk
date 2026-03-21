# OswUpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplex** | Pointer to **int32** | Duplex | [optional] 
**LinkSpeed** | Pointer to **int32** | LinkSpeed | [optional] 
**Port** | Pointer to [**WiredPortDTO**](WiredPortDTO.md) |  | [optional] 
**UpLinkPort** | Pointer to [**WiredPortDTO**](WiredPortDTO.md) |  | [optional] 

## Methods

### NewOswUpInfo

`func NewOswUpInfo() *OswUpInfo`

NewOswUpInfo instantiates a new OswUpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswUpInfoWithDefaults

`func NewOswUpInfoWithDefaults() *OswUpInfo`

NewOswUpInfoWithDefaults instantiates a new OswUpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *OswUpInfo) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswUpInfo) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswUpInfo) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswUpInfo) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswUpInfo) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswUpInfo) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswUpInfo) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswUpInfo) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetPort

`func (o *OswUpInfo) GetPort() WiredPortDTO`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswUpInfo) GetPortOk() (*WiredPortDTO, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswUpInfo) SetPort(v WiredPortDTO)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswUpInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUpLinkPort

`func (o *OswUpInfo) GetUpLinkPort() WiredPortDTO`

GetUpLinkPort returns the UpLinkPort field if non-nil, zero value otherwise.

### GetUpLinkPortOk

`func (o *OswUpInfo) GetUpLinkPortOk() (*WiredPortDTO, bool)`

GetUpLinkPortOk returns a tuple with the UpLinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLinkPort

`func (o *OswUpInfo) SetUpLinkPort(v WiredPortDTO)`

SetUpLinkPort sets UpLinkPort field to given value.

### HasUpLinkPort

`func (o *OswUpInfo) HasUpLinkPort() bool`

HasUpLinkPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


