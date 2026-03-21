# SelectVoucherGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupIds** | Pointer to **[]string** | ID list of voucher groups. Voucher group can be created using &#39;Create Voucher Group&#39; interface, and Voucher Group ID can be obtained from &#39;Get Voucher Group list&#39; interface | [optional] 
**SearchKey** | Pointer to **string** | Fuzzy query parameters, support field: voucher group name, voucher code | [optional] 
**TimeStart** | Pointer to **int64** | End timestamp filter query parameters, unit: MS | [optional] 
**Type** | **int32** | Select type. It should be a value as follows: 0: Represents selecting all voucher groups, this selection does not pass parameter [groupIds]. 1: Parameter [groupIds] includes the IDs of the voucher groups to be selected. 2: Parameter [groupIds] includes the IDs of the voucher groups not to be selected | 

## Methods

### NewSelectVoucherGroupOpenApiVO

`func NewSelectVoucherGroupOpenApiVO(type_ int32, ) *SelectVoucherGroupOpenApiVO`

NewSelectVoucherGroupOpenApiVO instantiates a new SelectVoucherGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectVoucherGroupOpenApiVOWithDefaults

`func NewSelectVoucherGroupOpenApiVOWithDefaults() *SelectVoucherGroupOpenApiVO`

NewSelectVoucherGroupOpenApiVOWithDefaults instantiates a new SelectVoucherGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupIds

`func (o *SelectVoucherGroupOpenApiVO) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *SelectVoucherGroupOpenApiVO) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *SelectVoucherGroupOpenApiVO) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *SelectVoucherGroupOpenApiVO) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### GetSearchKey

`func (o *SelectVoucherGroupOpenApiVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *SelectVoucherGroupOpenApiVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *SelectVoucherGroupOpenApiVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *SelectVoucherGroupOpenApiVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetTimeStart

`func (o *SelectVoucherGroupOpenApiVO) GetTimeStart() int64`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *SelectVoucherGroupOpenApiVO) GetTimeStartOk() (*int64, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *SelectVoucherGroupOpenApiVO) SetTimeStart(v int64)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *SelectVoucherGroupOpenApiVO) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetType

`func (o *SelectVoucherGroupOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectVoucherGroupOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectVoucherGroupOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


