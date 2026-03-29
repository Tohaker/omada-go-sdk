# OswLinkCapVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplex** | Pointer to **int32** | Duplex should be a value as follows: 0: Auto; 1: Half; 2: Full | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed | [optional] 

## Methods

### NewOswLinkCapVO

`func NewOswLinkCapVO() *OswLinkCapVO`

NewOswLinkCapVO instantiates a new OswLinkCapVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswLinkCapVOWithDefaults

`func NewOswLinkCapVOWithDefaults() *OswLinkCapVO`

NewOswLinkCapVOWithDefaults instantiates a new OswLinkCapVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *OswLinkCapVO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *OswLinkCapVO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *OswLinkCapVO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *OswLinkCapVO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *OswLinkCapVO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *OswLinkCapVO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *OswLinkCapVO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *OswLinkCapVO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


