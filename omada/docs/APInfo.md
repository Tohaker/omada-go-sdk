# APInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WiredUpLink** | Pointer to [**WiredUpLinkInfo**](WiredUpLinkInfo.md) |  | [optional] 
**WirelessUpLink** | Pointer to [**WirelessUpLinkInfo**](WirelessUpLinkInfo.md) |  | [optional] 

## Methods

### NewAPInfo

`func NewAPInfo() *APInfo`

NewAPInfo instantiates a new APInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPInfoWithDefaults

`func NewAPInfoWithDefaults() *APInfo`

NewAPInfoWithDefaults instantiates a new APInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWiredUpLink

`func (o *APInfo) GetWiredUpLink() WiredUpLinkInfo`

GetWiredUpLink returns the WiredUpLink field if non-nil, zero value otherwise.

### GetWiredUpLinkOk

`func (o *APInfo) GetWiredUpLinkOk() (*WiredUpLinkInfo, bool)`

GetWiredUpLinkOk returns a tuple with the WiredUpLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredUpLink

`func (o *APInfo) SetWiredUpLink(v WiredUpLinkInfo)`

SetWiredUpLink sets WiredUpLink field to given value.

### HasWiredUpLink

`func (o *APInfo) HasWiredUpLink() bool`

HasWiredUpLink returns a boolean if a field has been set.

### GetWirelessUpLink

`func (o *APInfo) GetWirelessUpLink() WirelessUpLinkInfo`

GetWirelessUpLink returns the WirelessUpLink field if non-nil, zero value otherwise.

### GetWirelessUpLinkOk

`func (o *APInfo) GetWirelessUpLinkOk() (*WirelessUpLinkInfo, bool)`

GetWirelessUpLinkOk returns a tuple with the WirelessUpLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessUpLink

`func (o *APInfo) SetWirelessUpLink(v WirelessUpLinkInfo)`

SetWirelessUpLink sets WirelessUpLink field to given value.

### HasWirelessUpLink

`func (o *APInfo) HasWirelessUpLink() bool`

HasWirelessUpLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


