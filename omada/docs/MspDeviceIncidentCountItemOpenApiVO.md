# MspDeviceIncidentCountItemOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | Customer ID | 
**Mac** | **string** | Device MAC address for which the incident count is requested. | 

## Methods

### NewMspDeviceIncidentCountItemOpenApiVO

`func NewMspDeviceIncidentCountItemOpenApiVO(customerId string, mac string, ) *MspDeviceIncidentCountItemOpenApiVO`

NewMspDeviceIncidentCountItemOpenApiVO instantiates a new MspDeviceIncidentCountItemOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspDeviceIncidentCountItemOpenApiVOWithDefaults

`func NewMspDeviceIncidentCountItemOpenApiVOWithDefaults() *MspDeviceIncidentCountItemOpenApiVO`

NewMspDeviceIncidentCountItemOpenApiVOWithDefaults instantiates a new MspDeviceIncidentCountItemOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *MspDeviceIncidentCountItemOpenApiVO) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *MspDeviceIncidentCountItemOpenApiVO) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *MspDeviceIncidentCountItemOpenApiVO) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetMac

`func (o *MspDeviceIncidentCountItemOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MspDeviceIncidentCountItemOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MspDeviceIncidentCountItemOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


