# WireUpLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplex** | Pointer to **int32** | Duplex | [optional] 
**IsFiberOptic** | Pointer to **bool** | Fiber Optic | [optional] 
**LinkSpeed** | Pointer to **int32** | LinkSpeed | [optional] 
**Port** | Pointer to [**WiredPortDTO**](WiredPortDTO.md) |  | [optional] 
**UpLinkPort** | Pointer to [**WiredPortDTO**](WiredPortDTO.md) |  | [optional] 

## Methods

### NewWireUpLink

`func NewWireUpLink() *WireUpLink`

NewWireUpLink instantiates a new WireUpLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireUpLinkWithDefaults

`func NewWireUpLinkWithDefaults() *WireUpLink`

NewWireUpLinkWithDefaults instantiates a new WireUpLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *WireUpLink) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *WireUpLink) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *WireUpLink) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *WireUpLink) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetIsFiberOptic

`func (o *WireUpLink) GetIsFiberOptic() bool`

GetIsFiberOptic returns the IsFiberOptic field if non-nil, zero value otherwise.

### GetIsFiberOpticOk

`func (o *WireUpLink) GetIsFiberOpticOk() (*bool, bool)`

GetIsFiberOpticOk returns a tuple with the IsFiberOptic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFiberOptic

`func (o *WireUpLink) SetIsFiberOptic(v bool)`

SetIsFiberOptic sets IsFiberOptic field to given value.

### HasIsFiberOptic

`func (o *WireUpLink) HasIsFiberOptic() bool`

HasIsFiberOptic returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *WireUpLink) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *WireUpLink) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *WireUpLink) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *WireUpLink) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetPort

`func (o *WireUpLink) GetPort() WiredPortDTO`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WireUpLink) GetPortOk() (*WiredPortDTO, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WireUpLink) SetPort(v WiredPortDTO)`

SetPort sets Port field to given value.

### HasPort

`func (o *WireUpLink) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUpLinkPort

`func (o *WireUpLink) GetUpLinkPort() WiredPortDTO`

GetUpLinkPort returns the UpLinkPort field if non-nil, zero value otherwise.

### GetUpLinkPortOk

`func (o *WireUpLink) GetUpLinkPortOk() (*WiredPortDTO, bool)`

GetUpLinkPortOk returns a tuple with the UpLinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLinkPort

`func (o *WireUpLink) SetUpLinkPort(v WiredPortDTO)`

SetUpLinkPort sets UpLinkPort field to given value.

### HasUpLinkPort

`func (o *WireUpLink) HasUpLinkPort() bool`

HasUpLinkPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


