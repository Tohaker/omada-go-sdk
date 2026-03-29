# MspDeviceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | CustomerId of device. | 
**Mac** | **string** | Mac of device. | 
**SiteId** | **string** | SiteId of device. | 

## Methods

### NewMspDeviceItem

`func NewMspDeviceItem(customerId string, mac string, siteId string, ) *MspDeviceItem`

NewMspDeviceItem instantiates a new MspDeviceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspDeviceItemWithDefaults

`func NewMspDeviceItemWithDefaults() *MspDeviceItem`

NewMspDeviceItemWithDefaults instantiates a new MspDeviceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerId

`func (o *MspDeviceItem) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *MspDeviceItem) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *MspDeviceItem) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.


### GetMac

`func (o *MspDeviceItem) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MspDeviceItem) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MspDeviceItem) SetMac(v string)`

SetMac sets Mac field to given value.


### GetSiteId

`func (o *MspDeviceItem) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MspDeviceItem) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MspDeviceItem) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


