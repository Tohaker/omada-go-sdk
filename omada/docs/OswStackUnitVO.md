# OswStackUnitVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | **string** | Device mac | 
**PortList** | [**[]OswStandPortVO**](OswStandPortVO.md) | Port List | 
**Unit** | **int32** | Unit | 

## Methods

### NewOswStackUnitVO

`func NewOswStackUnitVO(mac string, portList []OswStandPortVO, unit int32, ) *OswStackUnitVO`

NewOswStackUnitVO instantiates a new OswStackUnitVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackUnitVOWithDefaults

`func NewOswStackUnitVOWithDefaults() *OswStackUnitVO`

NewOswStackUnitVOWithDefaults instantiates a new OswStackUnitVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *OswStackUnitVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswStackUnitVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswStackUnitVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetPortList

`func (o *OswStackUnitVO) GetPortList() []OswStandPortVO`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *OswStackUnitVO) GetPortListOk() (*[]OswStandPortVO, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *OswStackUnitVO) SetPortList(v []OswStandPortVO)`

SetPortList sets PortList field to given value.


### GetUnit

`func (o *OswStackUnitVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswStackUnitVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswStackUnitVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


