# BatchUpdateApLagOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMacList** | **[]string** | AP mac list | 
**Enable** | Pointer to **bool** | Whether the device enables LAG. The following situations cause port aggregation to be ineffective: 1. When a specific Uplink Port is selected. 2. When PoE Out is enabled on target ports. 3. When Custom VLAN configurations is configured on target ports. 4. When Status is disabled on target ports. | [optional] 
**Mode** | Pointer to **int32** | LAG mode. Mode should be a value as follows: 0：SRC MAC + DST MAC; 1：DST MAC; 2：SRC MAC. | [optional] 

## Methods

### NewBatchUpdateApLagOpenApiVO

`func NewBatchUpdateApLagOpenApiVO(apMacList []string, ) *BatchUpdateApLagOpenApiVO`

NewBatchUpdateApLagOpenApiVO instantiates a new BatchUpdateApLagOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchUpdateApLagOpenApiVOWithDefaults

`func NewBatchUpdateApLagOpenApiVOWithDefaults() *BatchUpdateApLagOpenApiVO`

NewBatchUpdateApLagOpenApiVOWithDefaults instantiates a new BatchUpdateApLagOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMacList

`func (o *BatchUpdateApLagOpenApiVO) GetApMacList() []string`

GetApMacList returns the ApMacList field if non-nil, zero value otherwise.

### GetApMacListOk

`func (o *BatchUpdateApLagOpenApiVO) GetApMacListOk() (*[]string, bool)`

GetApMacListOk returns a tuple with the ApMacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMacList

`func (o *BatchUpdateApLagOpenApiVO) SetApMacList(v []string)`

SetApMacList sets ApMacList field to given value.


### GetEnable

`func (o *BatchUpdateApLagOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *BatchUpdateApLagOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *BatchUpdateApLagOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *BatchUpdateApLagOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetMode

`func (o *BatchUpdateApLagOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BatchUpdateApLagOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BatchUpdateApLagOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *BatchUpdateApLagOpenApiVO) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


