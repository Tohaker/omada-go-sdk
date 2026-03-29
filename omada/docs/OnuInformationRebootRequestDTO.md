# OnuInformationRebootRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RebootRequestList** | [**[]SingleOnuRebootRequestDTO**](SingleOnuRebootRequestDTO.md) | Reboot request list | 

## Methods

### NewOnuInformationRebootRequestDTO

`func NewOnuInformationRebootRequestDTO(rebootRequestList []SingleOnuRebootRequestDTO, ) *OnuInformationRebootRequestDTO`

NewOnuInformationRebootRequestDTO instantiates a new OnuInformationRebootRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuInformationRebootRequestDTOWithDefaults

`func NewOnuInformationRebootRequestDTOWithDefaults() *OnuInformationRebootRequestDTO`

NewOnuInformationRebootRequestDTOWithDefaults instantiates a new OnuInformationRebootRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRebootRequestList

`func (o *OnuInformationRebootRequestDTO) GetRebootRequestList() []SingleOnuRebootRequestDTO`

GetRebootRequestList returns the RebootRequestList field if non-nil, zero value otherwise.

### GetRebootRequestListOk

`func (o *OnuInformationRebootRequestDTO) GetRebootRequestListOk() (*[]SingleOnuRebootRequestDTO, bool)`

GetRebootRequestListOk returns a tuple with the RebootRequestList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootRequestList

`func (o *OnuInformationRebootRequestDTO) SetRebootRequestList(v []SingleOnuRebootRequestDTO)`

SetRebootRequestList sets RebootRequestList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


