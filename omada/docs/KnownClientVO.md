# KnownClientVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Block** | Pointer to **bool** |  | [optional] 
**BlockDisable** | Pointer to **bool** |  | [optional] 
**Download** | Pointer to **int64** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**Guest** | Pointer to **bool** |  | [optional] 
**LastSeen** | Pointer to **int64** |  | [optional] 
**LockToAp** | Pointer to **bool** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**Manager** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Upload** | Pointer to **int64** |  | [optional] 
**Vid** | Pointer to **int32** |  | [optional] 
**Wireless** | Pointer to **bool** |  | [optional] 

## Methods

### NewKnownClientVO

`func NewKnownClientVO() *KnownClientVO`

NewKnownClientVO instantiates a new KnownClientVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnownClientVOWithDefaults

`func NewKnownClientVOWithDefaults() *KnownClientVO`

NewKnownClientVOWithDefaults instantiates a new KnownClientVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlock

`func (o *KnownClientVO) GetBlock() bool`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *KnownClientVO) GetBlockOk() (*bool, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *KnownClientVO) SetBlock(v bool)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *KnownClientVO) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetBlockDisable

`func (o *KnownClientVO) GetBlockDisable() bool`

GetBlockDisable returns the BlockDisable field if non-nil, zero value otherwise.

### GetBlockDisableOk

`func (o *KnownClientVO) GetBlockDisableOk() (*bool, bool)`

GetBlockDisableOk returns a tuple with the BlockDisable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDisable

`func (o *KnownClientVO) SetBlockDisable(v bool)`

SetBlockDisable sets BlockDisable field to given value.

### HasBlockDisable

`func (o *KnownClientVO) HasBlockDisable() bool`

HasBlockDisable returns a boolean if a field has been set.

### GetDownload

`func (o *KnownClientVO) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *KnownClientVO) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *KnownClientVO) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *KnownClientVO) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetDuration

`func (o *KnownClientVO) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *KnownClientVO) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *KnownClientVO) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *KnownClientVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetGuest

`func (o *KnownClientVO) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *KnownClientVO) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *KnownClientVO) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *KnownClientVO) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetLastSeen

`func (o *KnownClientVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *KnownClientVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *KnownClientVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *KnownClientVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLockToAp

`func (o *KnownClientVO) GetLockToAp() bool`

GetLockToAp returns the LockToAp field if non-nil, zero value otherwise.

### GetLockToApOk

`func (o *KnownClientVO) GetLockToApOk() (*bool, bool)`

GetLockToApOk returns a tuple with the LockToAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockToAp

`func (o *KnownClientVO) SetLockToAp(v bool)`

SetLockToAp sets LockToAp field to given value.

### HasLockToAp

`func (o *KnownClientVO) HasLockToAp() bool`

HasLockToAp returns a boolean if a field has been set.

### GetMac

`func (o *KnownClientVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *KnownClientVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *KnownClientVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *KnownClientVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManager

`func (o *KnownClientVO) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *KnownClientVO) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *KnownClientVO) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *KnownClientVO) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetName

`func (o *KnownClientVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KnownClientVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KnownClientVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KnownClientVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpload

`func (o *KnownClientVO) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *KnownClientVO) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *KnownClientVO) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *KnownClientVO) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetVid

`func (o *KnownClientVO) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *KnownClientVO) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *KnownClientVO) SetVid(v int32)`

SetVid sets Vid field to given value.

### HasVid

`func (o *KnownClientVO) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetWireless

`func (o *KnownClientVO) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *KnownClientVO) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *KnownClientVO) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *KnownClientVO) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


