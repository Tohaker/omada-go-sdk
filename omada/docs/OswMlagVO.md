# OswMlagVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevicesNum** | Pointer to **int32** | Number of devices | [optional] 
**GroupId** | Pointer to **string** | M-LAG group ID | [optional] 
**Id** | Pointer to **string** | M-LAG ID | [optional] 
**LocateEnable** | Pointer to **bool** | Indicates whether the locate function is enabled | [optional] 
**Members** | Pointer to [**[]OswMlagMemberVO**](OswMlagMemberVO.md) | M-LAG Group member list | [optional] 
**Name** | Pointer to **string** | M-LAG group Name | [optional] 
**Status** | Pointer to **int32** | M-LAG group status should be a value as follows: 0: NORMAL; 1: PEER ERROR; 2: DAD ERROR; 3: ABNORMAL; 4: CONFIGURATION UNSYNCHRONIZED; 5: MLAG MEMBERS INCOMPATIBLE; 6: MLAG MEMBERS MISMATCH. | [optional] 
**Version** | Pointer to **string** | Version | [optional] 

## Methods

### NewOswMlagVO

`func NewOswMlagVO() *OswMlagVO`

NewOswMlagVO instantiates a new OswMlagVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswMlagVOWithDefaults

`func NewOswMlagVOWithDefaults() *OswMlagVO`

NewOswMlagVOWithDefaults instantiates a new OswMlagVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevicesNum

`func (o *OswMlagVO) GetDevicesNum() int32`

GetDevicesNum returns the DevicesNum field if non-nil, zero value otherwise.

### GetDevicesNumOk

`func (o *OswMlagVO) GetDevicesNumOk() (*int32, bool)`

GetDevicesNumOk returns a tuple with the DevicesNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicesNum

`func (o *OswMlagVO) SetDevicesNum(v int32)`

SetDevicesNum sets DevicesNum field to given value.

### HasDevicesNum

`func (o *OswMlagVO) HasDevicesNum() bool`

HasDevicesNum returns a boolean if a field has been set.

### GetGroupId

`func (o *OswMlagVO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *OswMlagVO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *OswMlagVO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *OswMlagVO) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetId

`func (o *OswMlagVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswMlagVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswMlagVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswMlagVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocateEnable

`func (o *OswMlagVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *OswMlagVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *OswMlagVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *OswMlagVO) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetMembers

`func (o *OswMlagVO) GetMembers() []OswMlagMemberVO`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *OswMlagVO) GetMembersOk() (*[]OswMlagMemberVO, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *OswMlagVO) SetMembers(v []OswMlagMemberVO)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *OswMlagVO) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *OswMlagVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswMlagVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswMlagVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswMlagVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OswMlagVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswMlagVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswMlagVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswMlagVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *OswMlagVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OswMlagVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OswMlagVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OswMlagVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


