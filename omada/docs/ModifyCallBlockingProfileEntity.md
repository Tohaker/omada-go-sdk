# ModifyCallBlockingProfileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingCallsBlocking** | Pointer to [**IncomingCallsBlockingVO**](IncomingCallsBlockingVO.md) |  | [optional] 
**IncomingCallsBlockingEnable** | Pointer to **bool** | Enable incoming calls blocking or not. The default value is false. | [optional] 
**OutgoingCallsBlocking** | Pointer to [**OutgoingCallsBlockingVO**](OutgoingCallsBlockingVO.md) |  | [optional] 
**OutgoingCallsBlockingEnable** | Pointer to **bool** | Enable outgoing calls blocking or not. The default value is false. | [optional] 
**ProfileName** | Pointer to **string** | Call blocking profile name | [optional] 

## Methods

### NewModifyCallBlockingProfileEntity

`func NewModifyCallBlockingProfileEntity() *ModifyCallBlockingProfileEntity`

NewModifyCallBlockingProfileEntity instantiates a new ModifyCallBlockingProfileEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifyCallBlockingProfileEntityWithDefaults

`func NewModifyCallBlockingProfileEntityWithDefaults() *ModifyCallBlockingProfileEntity`

NewModifyCallBlockingProfileEntityWithDefaults instantiates a new ModifyCallBlockingProfileEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingCallsBlocking

`func (o *ModifyCallBlockingProfileEntity) GetIncomingCallsBlocking() IncomingCallsBlockingVO`

GetIncomingCallsBlocking returns the IncomingCallsBlocking field if non-nil, zero value otherwise.

### GetIncomingCallsBlockingOk

`func (o *ModifyCallBlockingProfileEntity) GetIncomingCallsBlockingOk() (*IncomingCallsBlockingVO, bool)`

GetIncomingCallsBlockingOk returns a tuple with the IncomingCallsBlocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingCallsBlocking

`func (o *ModifyCallBlockingProfileEntity) SetIncomingCallsBlocking(v IncomingCallsBlockingVO)`

SetIncomingCallsBlocking sets IncomingCallsBlocking field to given value.

### HasIncomingCallsBlocking

`func (o *ModifyCallBlockingProfileEntity) HasIncomingCallsBlocking() bool`

HasIncomingCallsBlocking returns a boolean if a field has been set.

### GetIncomingCallsBlockingEnable

`func (o *ModifyCallBlockingProfileEntity) GetIncomingCallsBlockingEnable() bool`

GetIncomingCallsBlockingEnable returns the IncomingCallsBlockingEnable field if non-nil, zero value otherwise.

### GetIncomingCallsBlockingEnableOk

`func (o *ModifyCallBlockingProfileEntity) GetIncomingCallsBlockingEnableOk() (*bool, bool)`

GetIncomingCallsBlockingEnableOk returns a tuple with the IncomingCallsBlockingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingCallsBlockingEnable

`func (o *ModifyCallBlockingProfileEntity) SetIncomingCallsBlockingEnable(v bool)`

SetIncomingCallsBlockingEnable sets IncomingCallsBlockingEnable field to given value.

### HasIncomingCallsBlockingEnable

`func (o *ModifyCallBlockingProfileEntity) HasIncomingCallsBlockingEnable() bool`

HasIncomingCallsBlockingEnable returns a boolean if a field has been set.

### GetOutgoingCallsBlocking

`func (o *ModifyCallBlockingProfileEntity) GetOutgoingCallsBlocking() OutgoingCallsBlockingVO`

GetOutgoingCallsBlocking returns the OutgoingCallsBlocking field if non-nil, zero value otherwise.

### GetOutgoingCallsBlockingOk

`func (o *ModifyCallBlockingProfileEntity) GetOutgoingCallsBlockingOk() (*OutgoingCallsBlockingVO, bool)`

GetOutgoingCallsBlockingOk returns a tuple with the OutgoingCallsBlocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingCallsBlocking

`func (o *ModifyCallBlockingProfileEntity) SetOutgoingCallsBlocking(v OutgoingCallsBlockingVO)`

SetOutgoingCallsBlocking sets OutgoingCallsBlocking field to given value.

### HasOutgoingCallsBlocking

`func (o *ModifyCallBlockingProfileEntity) HasOutgoingCallsBlocking() bool`

HasOutgoingCallsBlocking returns a boolean if a field has been set.

### GetOutgoingCallsBlockingEnable

`func (o *ModifyCallBlockingProfileEntity) GetOutgoingCallsBlockingEnable() bool`

GetOutgoingCallsBlockingEnable returns the OutgoingCallsBlockingEnable field if non-nil, zero value otherwise.

### GetOutgoingCallsBlockingEnableOk

`func (o *ModifyCallBlockingProfileEntity) GetOutgoingCallsBlockingEnableOk() (*bool, bool)`

GetOutgoingCallsBlockingEnableOk returns a tuple with the OutgoingCallsBlockingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingCallsBlockingEnable

`func (o *ModifyCallBlockingProfileEntity) SetOutgoingCallsBlockingEnable(v bool)`

SetOutgoingCallsBlockingEnable sets OutgoingCallsBlockingEnable field to given value.

### HasOutgoingCallsBlockingEnable

`func (o *ModifyCallBlockingProfileEntity) HasOutgoingCallsBlockingEnable() bool`

HasOutgoingCallsBlockingEnable returns a boolean if a field has been set.

### GetProfileName

`func (o *ModifyCallBlockingProfileEntity) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *ModifyCallBlockingProfileEntity) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *ModifyCallBlockingProfileEntity) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *ModifyCallBlockingProfileEntity) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


