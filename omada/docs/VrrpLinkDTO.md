# VrrpLinkDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to **bool** | Whether The Link Is Blocked By STP Or Not | [optional] 
**DownLink** | Pointer to [**WiredPortV3DTO**](WiredPortV3DTO.md) |  | [optional] 
**Duplex** | Pointer to **int32** | Duplex | [optional] 
**LinkSpeed** | Pointer to **int32** | LinkSpeed | [optional] 
**UpLink** | Pointer to [**WiredPortV3DTO**](WiredPortV3DTO.md) |  | [optional] 

## Methods

### NewVrrpLinkDTO

`func NewVrrpLinkDTO() *VrrpLinkDTO`

NewVrrpLinkDTO instantiates a new VrrpLinkDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrrpLinkDTOWithDefaults

`func NewVrrpLinkDTOWithDefaults() *VrrpLinkDTO`

NewVrrpLinkDTOWithDefaults instantiates a new VrrpLinkDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *VrrpLinkDTO) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *VrrpLinkDTO) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *VrrpLinkDTO) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *VrrpLinkDTO) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetDownLink

`func (o *VrrpLinkDTO) GetDownLink() WiredPortV3DTO`

GetDownLink returns the DownLink field if non-nil, zero value otherwise.

### GetDownLinkOk

`func (o *VrrpLinkDTO) GetDownLinkOk() (*WiredPortV3DTO, bool)`

GetDownLinkOk returns a tuple with the DownLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLink

`func (o *VrrpLinkDTO) SetDownLink(v WiredPortV3DTO)`

SetDownLink sets DownLink field to given value.

### HasDownLink

`func (o *VrrpLinkDTO) HasDownLink() bool`

HasDownLink returns a boolean if a field has been set.

### GetDuplex

`func (o *VrrpLinkDTO) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *VrrpLinkDTO) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *VrrpLinkDTO) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *VrrpLinkDTO) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *VrrpLinkDTO) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *VrrpLinkDTO) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *VrrpLinkDTO) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *VrrpLinkDTO) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetUpLink

`func (o *VrrpLinkDTO) GetUpLink() WiredPortV3DTO`

GetUpLink returns the UpLink field if non-nil, zero value otherwise.

### GetUpLinkOk

`func (o *VrrpLinkDTO) GetUpLinkOk() (*WiredPortV3DTO, bool)`

GetUpLinkOk returns a tuple with the UpLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLink

`func (o *VrrpLinkDTO) SetUpLink(v WiredPortV3DTO)`

SetUpLink sets UpLink field to given value.

### HasUpLink

`func (o *VrrpLinkDTO) HasUpLink() bool`

HasUpLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


