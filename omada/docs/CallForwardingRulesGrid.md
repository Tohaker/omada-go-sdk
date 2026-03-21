# CallForwardingRulesGrid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BookList** | Pointer to [**[]BookOpenApiVO**](BookOpenApiVO.md) | Contacts. | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]CallForwardingRule**](CallForwardingRule.md) |  | [optional] 
**DeviceList** | Pointer to [**[]DeviceOpenApiVO**](DeviceOpenApiVO.md) | Telephony devices. | [optional] 
**NumberList** | Pointer to [**[]NumberOpenApiVO**](NumberOpenApiVO.md) | Telephone numbers. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewCallForwardingRulesGrid

`func NewCallForwardingRulesGrid() *CallForwardingRulesGrid`

NewCallForwardingRulesGrid instantiates a new CallForwardingRulesGrid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallForwardingRulesGridWithDefaults

`func NewCallForwardingRulesGridWithDefaults() *CallForwardingRulesGrid`

NewCallForwardingRulesGridWithDefaults instantiates a new CallForwardingRulesGrid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookList

`func (o *CallForwardingRulesGrid) GetBookList() []BookOpenApiVO`

GetBookList returns the BookList field if non-nil, zero value otherwise.

### GetBookListOk

`func (o *CallForwardingRulesGrid) GetBookListOk() (*[]BookOpenApiVO, bool)`

GetBookListOk returns a tuple with the BookList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookList

`func (o *CallForwardingRulesGrid) SetBookList(v []BookOpenApiVO)`

SetBookList sets BookList field to given value.

### HasBookList

`func (o *CallForwardingRulesGrid) HasBookList() bool`

HasBookList returns a boolean if a field has been set.

### GetCurrentPage

`func (o *CallForwardingRulesGrid) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *CallForwardingRulesGrid) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *CallForwardingRulesGrid) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *CallForwardingRulesGrid) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *CallForwardingRulesGrid) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *CallForwardingRulesGrid) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *CallForwardingRulesGrid) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *CallForwardingRulesGrid) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *CallForwardingRulesGrid) GetData() []CallForwardingRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CallForwardingRulesGrid) GetDataOk() (*[]CallForwardingRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CallForwardingRulesGrid) SetData(v []CallForwardingRule)`

SetData sets Data field to given value.

### HasData

`func (o *CallForwardingRulesGrid) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDeviceList

`func (o *CallForwardingRulesGrid) GetDeviceList() []DeviceOpenApiVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *CallForwardingRulesGrid) GetDeviceListOk() (*[]DeviceOpenApiVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *CallForwardingRulesGrid) SetDeviceList(v []DeviceOpenApiVO)`

SetDeviceList sets DeviceList field to given value.

### HasDeviceList

`func (o *CallForwardingRulesGrid) HasDeviceList() bool`

HasDeviceList returns a boolean if a field has been set.

### GetNumberList

`func (o *CallForwardingRulesGrid) GetNumberList() []NumberOpenApiVO`

GetNumberList returns the NumberList field if non-nil, zero value otherwise.

### GetNumberListOk

`func (o *CallForwardingRulesGrid) GetNumberListOk() (*[]NumberOpenApiVO, bool)`

GetNumberListOk returns a tuple with the NumberList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberList

`func (o *CallForwardingRulesGrid) SetNumberList(v []NumberOpenApiVO)`

SetNumberList sets NumberList field to given value.

### HasNumberList

`func (o *CallForwardingRulesGrid) HasNumberList() bool`

HasNumberList returns a boolean if a field has been set.

### GetTotalRows

`func (o *CallForwardingRulesGrid) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *CallForwardingRulesGrid) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *CallForwardingRulesGrid) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *CallForwardingRulesGrid) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


