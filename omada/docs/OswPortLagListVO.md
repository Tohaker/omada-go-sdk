# OswPortLagListVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagList** | Pointer to **[]int32** | LAG List | [optional] 
**Mac** | **string** | Switch Mac | 
**PortList** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | Port List | [optional] 
**StackId** | Pointer to **string** | Stack ID | [optional] 
**Unit** | Pointer to **int32** | Unit | [optional] 

## Methods

### NewOswPortLagListVO

`func NewOswPortLagListVO(mac string, ) *OswPortLagListVO`

NewOswPortLagListVO instantiates a new OswPortLagListVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortLagListVOWithDefaults

`func NewOswPortLagListVOWithDefaults() *OswPortLagListVO`

NewOswPortLagListVOWithDefaults instantiates a new OswPortLagListVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagList

`func (o *OswPortLagListVO) GetLagList() []int32`

GetLagList returns the LagList field if non-nil, zero value otherwise.

### GetLagListOk

`func (o *OswPortLagListVO) GetLagListOk() (*[]int32, bool)`

GetLagListOk returns a tuple with the LagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagList

`func (o *OswPortLagListVO) SetLagList(v []int32)`

SetLagList sets LagList field to given value.

### HasLagList

`func (o *OswPortLagListVO) HasLagList() bool`

HasLagList returns a boolean if a field has been set.

### GetMac

`func (o *OswPortLagListVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswPortLagListVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswPortLagListVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetPortList

`func (o *OswPortLagListVO) GetPortList() []OswStandPortVO`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *OswPortLagListVO) GetPortListOk() (*[]OswStandPortVO, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *OswPortLagListVO) SetPortList(v []OswStandPortVO)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *OswPortLagListVO) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetStackId

`func (o *OswPortLagListVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswPortLagListVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswPortLagListVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswPortLagListVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetUnit

`func (o *OswPortLagListVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswPortLagListVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswPortLagListVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswPortLagListVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


