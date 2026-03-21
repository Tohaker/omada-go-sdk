# OswMlagCccVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | M-LAG group member mac | [optional] 
**MlagCccResult** | Pointer to [**[]MlagCccResultVO**](MlagCccResultVO.md) | M-LAG group members configuration check result | [optional] 

## Methods

### NewOswMlagCccVO

`func NewOswMlagCccVO() *OswMlagCccVO`

NewOswMlagCccVO instantiates a new OswMlagCccVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswMlagCccVOWithDefaults

`func NewOswMlagCccVOWithDefaults() *OswMlagCccVO`

NewOswMlagCccVOWithDefaults instantiates a new OswMlagCccVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *OswMlagCccVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswMlagCccVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswMlagCccVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswMlagCccVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMlagCccResult

`func (o *OswMlagCccVO) GetMlagCccResult() []MlagCccResultVO`

GetMlagCccResult returns the MlagCccResult field if non-nil, zero value otherwise.

### GetMlagCccResultOk

`func (o *OswMlagCccVO) GetMlagCccResultOk() (*[]MlagCccResultVO, bool)`

GetMlagCccResultOk returns a tuple with the MlagCccResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagCccResult

`func (o *OswMlagCccVO) SetMlagCccResult(v []MlagCccResultVO)`

SetMlagCccResult sets MlagCccResult field to given value.

### HasMlagCccResult

`func (o *OswMlagCccVO) HasMlagCccResult() bool`

HasMlagCccResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


