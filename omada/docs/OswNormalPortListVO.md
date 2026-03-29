# OswNormalPortListVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Switch MAC Address | [optional] 
**PortList** | Pointer to **[]int32** | Switch port list. | [optional] 

## Methods

### NewOswNormalPortListVO

`func NewOswNormalPortListVO() *OswNormalPortListVO`

NewOswNormalPortListVO instantiates a new OswNormalPortListVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswNormalPortListVOWithDefaults

`func NewOswNormalPortListVOWithDefaults() *OswNormalPortListVO`

NewOswNormalPortListVOWithDefaults instantiates a new OswNormalPortListVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *OswNormalPortListVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswNormalPortListVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswNormalPortListVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswNormalPortListVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPortList

`func (o *OswNormalPortListVO) GetPortList() []int32`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *OswNormalPortListVO) GetPortListOk() (*[]int32, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *OswNormalPortListVO) SetPortList(v []int32)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *OswNormalPortListVO) HasPortList() bool`

HasPortList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


