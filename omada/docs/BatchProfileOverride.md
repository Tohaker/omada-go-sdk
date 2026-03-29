# BatchProfileOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortList** | **[]int32** | Port ID List. | 
**ProfileOverrideEnable** | **bool** | Profile Override Enable. | 

## Methods

### NewBatchProfileOverride

`func NewBatchProfileOverride(portList []int32, profileOverrideEnable bool, ) *BatchProfileOverride`

NewBatchProfileOverride instantiates a new BatchProfileOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchProfileOverrideWithDefaults

`func NewBatchProfileOverrideWithDefaults() *BatchProfileOverride`

NewBatchProfileOverrideWithDefaults instantiates a new BatchProfileOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortList

`func (o *BatchProfileOverride) GetPortList() []int32`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *BatchProfileOverride) GetPortListOk() (*[]int32, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *BatchProfileOverride) SetPortList(v []int32)`

SetPortList sets PortList field to given value.


### GetProfileOverrideEnable

`func (o *BatchProfileOverride) GetProfileOverrideEnable() bool`

GetProfileOverrideEnable returns the ProfileOverrideEnable field if non-nil, zero value otherwise.

### GetProfileOverrideEnableOk

`func (o *BatchProfileOverride) GetProfileOverrideEnableOk() (*bool, bool)`

GetProfileOverrideEnableOk returns a tuple with the ProfileOverrideEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileOverrideEnable

`func (o *BatchProfileOverride) SetProfileOverrideEnable(v bool)`

SetProfileOverrideEnable sets ProfileOverrideEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


