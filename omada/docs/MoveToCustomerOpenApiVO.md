# MoveToCustomerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Target site id. | [optional] 
**Customer** | **string** | Target customer id. | 
**DeviceMac** | Pointer to **string** | Mac adress of the device. If [deviceMac] is not empty, it takes effect with a higher priority than [deviceMacs]. | [optional] 
**DeviceMacs** | Pointer to **[]string** | Mac adress of the device. | [optional] 

## Methods

### NewMoveToCustomerOpenApiVO

`func NewMoveToCustomerOpenApiVO(customer string, ) *MoveToCustomerOpenApiVO`

NewMoveToCustomerOpenApiVO instantiates a new MoveToCustomerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveToCustomerOpenApiVOWithDefaults

`func NewMoveToCustomerOpenApiVOWithDefaults() *MoveToCustomerOpenApiVO`

NewMoveToCustomerOpenApiVOWithDefaults instantiates a new MoveToCustomerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *MoveToCustomerOpenApiVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MoveToCustomerOpenApiVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MoveToCustomerOpenApiVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *MoveToCustomerOpenApiVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetCustomer

`func (o *MoveToCustomerOpenApiVO) GetCustomer() string`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *MoveToCustomerOpenApiVO) GetCustomerOk() (*string, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *MoveToCustomerOpenApiVO) SetCustomer(v string)`

SetCustomer sets Customer field to given value.


### GetDeviceMac

`func (o *MoveToCustomerOpenApiVO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *MoveToCustomerOpenApiVO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *MoveToCustomerOpenApiVO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *MoveToCustomerOpenApiVO) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceMacs

`func (o *MoveToCustomerOpenApiVO) GetDeviceMacs() []string`

GetDeviceMacs returns the DeviceMacs field if non-nil, zero value otherwise.

### GetDeviceMacsOk

`func (o *MoveToCustomerOpenApiVO) GetDeviceMacsOk() (*[]string, bool)`

GetDeviceMacsOk returns a tuple with the DeviceMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMacs

`func (o *MoveToCustomerOpenApiVO) SetDeviceMacs(v []string)`

SetDeviceMacs sets DeviceMacs field to given value.

### HasDeviceMacs

`func (o *MoveToCustomerOpenApiVO) HasDeviceMacs() bool`

HasDeviceMacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


