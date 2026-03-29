# SdWanLinkedSpoke

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac1** | Pointer to **string** | The device MAC of one of the two linked-spokes | [optional] 
**DeviceMac2** | Pointer to **string** | Another device MAC of the two linked-spokes | [optional] 
**Status** | Pointer to **int32** | The connection status of SD-WAN tunnel between spokes | [optional] 

## Methods

### NewSdWanLinkedSpoke

`func NewSdWanLinkedSpoke() *SdWanLinkedSpoke`

NewSdWanLinkedSpoke instantiates a new SdWanLinkedSpoke object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdWanLinkedSpokeWithDefaults

`func NewSdWanLinkedSpokeWithDefaults() *SdWanLinkedSpoke`

NewSdWanLinkedSpokeWithDefaults instantiates a new SdWanLinkedSpoke object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac1

`func (o *SdWanLinkedSpoke) GetDeviceMac1() string`

GetDeviceMac1 returns the DeviceMac1 field if non-nil, zero value otherwise.

### GetDeviceMac1Ok

`func (o *SdWanLinkedSpoke) GetDeviceMac1Ok() (*string, bool)`

GetDeviceMac1Ok returns a tuple with the DeviceMac1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac1

`func (o *SdWanLinkedSpoke) SetDeviceMac1(v string)`

SetDeviceMac1 sets DeviceMac1 field to given value.

### HasDeviceMac1

`func (o *SdWanLinkedSpoke) HasDeviceMac1() bool`

HasDeviceMac1 returns a boolean if a field has been set.

### GetDeviceMac2

`func (o *SdWanLinkedSpoke) GetDeviceMac2() string`

GetDeviceMac2 returns the DeviceMac2 field if non-nil, zero value otherwise.

### GetDeviceMac2Ok

`func (o *SdWanLinkedSpoke) GetDeviceMac2Ok() (*string, bool)`

GetDeviceMac2Ok returns a tuple with the DeviceMac2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac2

`func (o *SdWanLinkedSpoke) SetDeviceMac2(v string)`

SetDeviceMac2 sets DeviceMac2 field to given value.

### HasDeviceMac2

`func (o *SdWanLinkedSpoke) HasDeviceMac2() bool`

HasDeviceMac2 returns a boolean if a field has been set.

### GetStatus

`func (o *SdWanLinkedSpoke) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SdWanLinkedSpoke) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SdWanLinkedSpoke) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SdWanLinkedSpoke) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


