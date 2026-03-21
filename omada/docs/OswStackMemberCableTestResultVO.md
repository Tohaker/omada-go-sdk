# OswStackMemberCableTestResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]OswCableTestResultVO**](OswCableTestResultVO.md) | Test results. | [optional] 
**Mac** | Pointer to **string** | Device Mac. | [optional] 
**Unit** | Pointer to **int32** | unit | [optional] 

## Methods

### NewOswStackMemberCableTestResultVO

`func NewOswStackMemberCableTestResultVO() *OswStackMemberCableTestResultVO`

NewOswStackMemberCableTestResultVO instantiates a new OswStackMemberCableTestResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackMemberCableTestResultVOWithDefaults

`func NewOswStackMemberCableTestResultVOWithDefaults() *OswStackMemberCableTestResultVO`

NewOswStackMemberCableTestResultVOWithDefaults instantiates a new OswStackMemberCableTestResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OswStackMemberCableTestResultVO) GetData() []OswCableTestResultVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OswStackMemberCableTestResultVO) GetDataOk() (*[]OswCableTestResultVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OswStackMemberCableTestResultVO) SetData(v []OswCableTestResultVO)`

SetData sets Data field to given value.

### HasData

`func (o *OswStackMemberCableTestResultVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMac

`func (o *OswStackMemberCableTestResultVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswStackMemberCableTestResultVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswStackMemberCableTestResultVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswStackMemberCableTestResultVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetUnit

`func (o *OswStackMemberCableTestResultVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswStackMemberCableTestResultVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswStackMemberCableTestResultVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswStackMemberCableTestResultVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


