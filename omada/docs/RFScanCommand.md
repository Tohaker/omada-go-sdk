# RFScanCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RadioIdList** | Pointer to **[]int32** | Parameter [radioIdList] indicates the set of frequency bands for radio frequency scan. The value range of each element in the list should be between 0 and 3. 0: 2.4GHz, 1: 5GHz, 2: 5GHz-2, 3: 6GHz. This parameter takes effect only when the device supports selecting bands for radio frequency scan; If the device does not support selecting bands, or if this parameter is not passed or is an empty list, the device will perform radio frequency scan for all its frequency bands. | [optional] 

## Methods

### NewRFScanCommand

`func NewRFScanCommand() *RFScanCommand`

NewRFScanCommand instantiates a new RFScanCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRFScanCommandWithDefaults

`func NewRFScanCommandWithDefaults() *RFScanCommand`

NewRFScanCommandWithDefaults instantiates a new RFScanCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRadioIdList

`func (o *RFScanCommand) GetRadioIdList() []int32`

GetRadioIdList returns the RadioIdList field if non-nil, zero value otherwise.

### GetRadioIdListOk

`func (o *RFScanCommand) GetRadioIdListOk() (*[]int32, bool)`

GetRadioIdListOk returns a tuple with the RadioIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioIdList

`func (o *RFScanCommand) SetRadioIdList(v []int32)`

SetRadioIdList sets RadioIdList field to given value.

### HasRadioIdList

`func (o *RFScanCommand) HasRadioIdList() bool`

HasRadioIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


