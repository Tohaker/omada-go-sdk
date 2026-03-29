# MoveToCustomerVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Target site ID | [optional] 
**Customer** | **string** | Target customer ID | 
**DeviceMacs** | **[]string** | DeviceMacs should contain 1 entry | 

## Methods

### NewMoveToCustomerVO

`func NewMoveToCustomerVO(customer string, deviceMacs []string, ) *MoveToCustomerVO`

NewMoveToCustomerVO instantiates a new MoveToCustomerVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveToCustomerVOWithDefaults

`func NewMoveToCustomerVOWithDefaults() *MoveToCustomerVO`

NewMoveToCustomerVOWithDefaults instantiates a new MoveToCustomerVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *MoveToCustomerVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MoveToCustomerVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MoveToCustomerVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *MoveToCustomerVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetCustomer

`func (o *MoveToCustomerVO) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *MoveToCustomerVO) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *MoveToCustomerVO) SetCustomer(v string)`

SetCustomer sets Customer field to given value.


### GetDeviceMacs

`func (o *MoveToCustomerVO) GetDeviceMacs() []string`

GetDeviceMacs returns the DeviceMacs field if non-nil, zero value otherwise.

### GetDeviceMacsOk

`func (o *MoveToCustomerVO) GetDeviceMacsOk() (*[]string, bool)`

GetDeviceMacsOk returns a tuple with the DeviceMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMacs

`func (o *MoveToCustomerVO) SetDeviceMacs(v []string)`

SetDeviceMacs sets DeviceMacs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


