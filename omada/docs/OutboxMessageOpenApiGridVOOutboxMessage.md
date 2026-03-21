# OutboxMessageOpenApiGridVOOutboxMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardStatus** | Pointer to **int32** | Sim card status. | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]OutboxMessage**](OutboxMessage.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewOutboxMessageOpenApiGridVOOutboxMessage

`func NewOutboxMessageOpenApiGridVOOutboxMessage() *OutboxMessageOpenApiGridVOOutboxMessage`

NewOutboxMessageOpenApiGridVOOutboxMessage instantiates a new OutboxMessageOpenApiGridVOOutboxMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboxMessageOpenApiGridVOOutboxMessageWithDefaults

`func NewOutboxMessageOpenApiGridVOOutboxMessageWithDefaults() *OutboxMessageOpenApiGridVOOutboxMessage`

NewOutboxMessageOpenApiGridVOOutboxMessageWithDefaults instantiates a new OutboxMessageOpenApiGridVOOutboxMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardStatus

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) GetCardStatus() int32`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) GetCardStatusOk() (*int32, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) SetCardStatus(v int32)`

SetCardStatus sets CardStatus field to given value.

### HasCardStatus

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) HasCardStatus() bool`

HasCardStatus returns a boolean if a field has been set.

### GetCurrentPage

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) GetData() []OutboxMessage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) GetDataOk() (*[]OutboxMessage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) SetData(v []OutboxMessage)`

SetData sets Data field to given value.

### HasData

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalRows

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *OutboxMessageOpenApiGridVOOutboxMessage) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


