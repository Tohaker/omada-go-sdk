# SelectIdsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | Voucher Group ID. Voucher group can be created using &#39;Create Voucher Group&#39; interface, and Voucher Group ID can be obtained from &#39;Get Voucher Group list&#39; interface | 
**Ids** | Pointer to **[]string** | ID list of vouchers. Voucher can be created using &#39;Create Voucher Group&#39; interface, and Voucher ID can be obtained from &#39;Get Voucher Group Detail&#39; interface | [optional] 
**SearchKey** | Pointer to **string** | Fuzzy query parameters, support field: voucher code | [optional] 
**Status** | Pointer to **int32** | voucher status filter query parameters. It should be a value as follows: 0: Unused vouchers, 1: In use vouchers, 2: Expired vouchers | [optional] 
**Type** | **int32** | Select type. It should be a value as follows: 0: Represents selecting all vouchers in the voucher group, this selection does not pass parameter [ids]. 1: Parameter [ids] includes the IDs of vouchers in the voucher group to be selected. 2: Parameter [ids] includes the IDs of vouchers in the voucher group not to be selected | 

## Methods

### NewSelectIdsOpenApiVO

`func NewSelectIdsOpenApiVO(groupId string, type_ int32, ) *SelectIdsOpenApiVO`

NewSelectIdsOpenApiVO instantiates a new SelectIdsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectIdsOpenApiVOWithDefaults

`func NewSelectIdsOpenApiVOWithDefaults() *SelectIdsOpenApiVO`

NewSelectIdsOpenApiVOWithDefaults instantiates a new SelectIdsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *SelectIdsOpenApiVO) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *SelectIdsOpenApiVO) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *SelectIdsOpenApiVO) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetIds

`func (o *SelectIdsOpenApiVO) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *SelectIdsOpenApiVO) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *SelectIdsOpenApiVO) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *SelectIdsOpenApiVO) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetSearchKey

`func (o *SelectIdsOpenApiVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *SelectIdsOpenApiVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *SelectIdsOpenApiVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *SelectIdsOpenApiVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetStatus

`func (o *SelectIdsOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SelectIdsOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SelectIdsOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SelectIdsOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SelectIdsOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectIdsOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectIdsOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


