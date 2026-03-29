# MoveSiteInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** | Error code | [optional] 
**Mac** | Pointer to **string** | Mac address | [optional] 
**MoveSiteStatus** | Pointer to **int32** | Move site status.MoveSiteStatus should be a value as follows:0: Device is in MoveSite;1: MoveSite Success;2: MoveSite Failed | [optional] 
**Name** | Pointer to **string** | Name | [optional] 

## Methods

### NewMoveSiteInfoVO

`func NewMoveSiteInfoVO() *MoveSiteInfoVO`

NewMoveSiteInfoVO instantiates a new MoveSiteInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveSiteInfoVOWithDefaults

`func NewMoveSiteInfoVOWithDefaults() *MoveSiteInfoVO`

NewMoveSiteInfoVOWithDefaults instantiates a new MoveSiteInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *MoveSiteInfoVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MoveSiteInfoVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MoveSiteInfoVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *MoveSiteInfoVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMac

`func (o *MoveSiteInfoVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MoveSiteInfoVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MoveSiteInfoVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *MoveSiteInfoVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMoveSiteStatus

`func (o *MoveSiteInfoVO) GetMoveSiteStatus() int32`

GetMoveSiteStatus returns the MoveSiteStatus field if non-nil, zero value otherwise.

### GetMoveSiteStatusOk

`func (o *MoveSiteInfoVO) GetMoveSiteStatusOk() (*int32, bool)`

GetMoveSiteStatusOk returns a tuple with the MoveSiteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveSiteStatus

`func (o *MoveSiteInfoVO) SetMoveSiteStatus(v int32)`

SetMoveSiteStatus sets MoveSiteStatus field to given value.

### HasMoveSiteStatus

`func (o *MoveSiteInfoVO) HasMoveSiteStatus() bool`

HasMoveSiteStatus returns a boolean if a field has been set.

### GetName

`func (o *MoveSiteInfoVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MoveSiteInfoVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MoveSiteInfoVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MoveSiteInfoVO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


