# CallBlockingProfileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncomingCallsBlocking** | Pointer to [**IncomingCallsBlockingVO**](IncomingCallsBlockingVO.md) |  | [optional] 
**IncomingCallsBlockingEnable** | Pointer to **bool** | Enable incoming calls blocking or not. | [optional] 
**OmadacId** | Pointer to **string** | Omadac ID | [optional] 
**OutgoingCallsBlocking** | Pointer to [**OutgoingCallsBlockingVO**](OutgoingCallsBlockingVO.md) |  | [optional] 
**OutgoingCallsBlockingEnable** | Pointer to **bool** | Enable outgoing calls blocking or not. | [optional] 
**ProfileId** | Pointer to **string** | Call blocking profile ID | [optional] 
**ProfileName** | Pointer to **string** | Call blocking profile name | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 

## Methods

### NewCallBlockingProfileEntity

`func NewCallBlockingProfileEntity() *CallBlockingProfileEntity`

NewCallBlockingProfileEntity instantiates a new CallBlockingProfileEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallBlockingProfileEntityWithDefaults

`func NewCallBlockingProfileEntityWithDefaults() *CallBlockingProfileEntity`

NewCallBlockingProfileEntityWithDefaults instantiates a new CallBlockingProfileEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncomingCallsBlocking

`func (o *CallBlockingProfileEntity) GetIncomingCallsBlocking() IncomingCallsBlockingVO`

GetIncomingCallsBlocking returns the IncomingCallsBlocking field if non-nil, zero value otherwise.

### GetIncomingCallsBlockingOk

`func (o *CallBlockingProfileEntity) GetIncomingCallsBlockingOk() (*IncomingCallsBlockingVO, bool)`

GetIncomingCallsBlockingOk returns a tuple with the IncomingCallsBlocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingCallsBlocking

`func (o *CallBlockingProfileEntity) SetIncomingCallsBlocking(v IncomingCallsBlockingVO)`

SetIncomingCallsBlocking sets IncomingCallsBlocking field to given value.

### HasIncomingCallsBlocking

`func (o *CallBlockingProfileEntity) HasIncomingCallsBlocking() bool`

HasIncomingCallsBlocking returns a boolean if a field has been set.

### GetIncomingCallsBlockingEnable

`func (o *CallBlockingProfileEntity) GetIncomingCallsBlockingEnable() bool`

GetIncomingCallsBlockingEnable returns the IncomingCallsBlockingEnable field if non-nil, zero value otherwise.

### GetIncomingCallsBlockingEnableOk

`func (o *CallBlockingProfileEntity) GetIncomingCallsBlockingEnableOk() (*bool, bool)`

GetIncomingCallsBlockingEnableOk returns a tuple with the IncomingCallsBlockingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingCallsBlockingEnable

`func (o *CallBlockingProfileEntity) SetIncomingCallsBlockingEnable(v bool)`

SetIncomingCallsBlockingEnable sets IncomingCallsBlockingEnable field to given value.

### HasIncomingCallsBlockingEnable

`func (o *CallBlockingProfileEntity) HasIncomingCallsBlockingEnable() bool`

HasIncomingCallsBlockingEnable returns a boolean if a field has been set.

### GetOmadacId

`func (o *CallBlockingProfileEntity) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *CallBlockingProfileEntity) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *CallBlockingProfileEntity) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *CallBlockingProfileEntity) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetOutgoingCallsBlocking

`func (o *CallBlockingProfileEntity) GetOutgoingCallsBlocking() OutgoingCallsBlockingVO`

GetOutgoingCallsBlocking returns the OutgoingCallsBlocking field if non-nil, zero value otherwise.

### GetOutgoingCallsBlockingOk

`func (o *CallBlockingProfileEntity) GetOutgoingCallsBlockingOk() (*OutgoingCallsBlockingVO, bool)`

GetOutgoingCallsBlockingOk returns a tuple with the OutgoingCallsBlocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingCallsBlocking

`func (o *CallBlockingProfileEntity) SetOutgoingCallsBlocking(v OutgoingCallsBlockingVO)`

SetOutgoingCallsBlocking sets OutgoingCallsBlocking field to given value.

### HasOutgoingCallsBlocking

`func (o *CallBlockingProfileEntity) HasOutgoingCallsBlocking() bool`

HasOutgoingCallsBlocking returns a boolean if a field has been set.

### GetOutgoingCallsBlockingEnable

`func (o *CallBlockingProfileEntity) GetOutgoingCallsBlockingEnable() bool`

GetOutgoingCallsBlockingEnable returns the OutgoingCallsBlockingEnable field if non-nil, zero value otherwise.

### GetOutgoingCallsBlockingEnableOk

`func (o *CallBlockingProfileEntity) GetOutgoingCallsBlockingEnableOk() (*bool, bool)`

GetOutgoingCallsBlockingEnableOk returns a tuple with the OutgoingCallsBlockingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingCallsBlockingEnable

`func (o *CallBlockingProfileEntity) SetOutgoingCallsBlockingEnable(v bool)`

SetOutgoingCallsBlockingEnable sets OutgoingCallsBlockingEnable field to given value.

### HasOutgoingCallsBlockingEnable

`func (o *CallBlockingProfileEntity) HasOutgoingCallsBlockingEnable() bool`

HasOutgoingCallsBlockingEnable returns a boolean if a field has been set.

### GetProfileId

`func (o *CallBlockingProfileEntity) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CallBlockingProfileEntity) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CallBlockingProfileEntity) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *CallBlockingProfileEntity) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProfileName

`func (o *CallBlockingProfileEntity) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *CallBlockingProfileEntity) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *CallBlockingProfileEntity) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.

### HasProfileName

`func (o *CallBlockingProfileEntity) HasProfileName() bool`

HasProfileName returns a boolean if a field has been set.

### GetSiteId

`func (o *CallBlockingProfileEntity) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CallBlockingProfileEntity) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CallBlockingProfileEntity) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CallBlockingProfileEntity) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


