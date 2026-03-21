# UsbInfoListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsbList** | Pointer to [**[]UsbInfo**](UsbInfo.md) | USB info list. | [optional] 

## Methods

### NewUsbInfoListResponse

`func NewUsbInfoListResponse() *UsbInfoListResponse`

NewUsbInfoListResponse instantiates a new UsbInfoListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsbInfoListResponseWithDefaults

`func NewUsbInfoListResponseWithDefaults() *UsbInfoListResponse`

NewUsbInfoListResponseWithDefaults instantiates a new UsbInfoListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsbList

`func (o *UsbInfoListResponse) GetUsbList() []UsbInfo`

GetUsbList returns the UsbList field if non-nil, zero value otherwise.

### GetUsbListOk

`func (o *UsbInfoListResponse) GetUsbListOk() (*[]UsbInfo, bool)`

GetUsbListOk returns a tuple with the UsbList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbList

`func (o *UsbInfoListResponse) SetUsbList(v []UsbInfo)`

SetUsbList sets UsbList field to given value.

### HasUsbList

`func (o *UsbInfoListResponse) HasUsbList() bool`

HasUsbList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


