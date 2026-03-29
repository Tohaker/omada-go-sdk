# ClassRuleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassType** | **int32** | The Class value selected in the Qos Class configuration should be a value as follows: 1: Class 1, 2: Class 2, 3: Class 3. | 
**Dscp** | **string** | The DSCP value selected in the DSCP configuration should be a value as follows: any: any; 8: IP precedence 1; 16: IP precedence 2; 24: IP precedence 3; 32: IP precedence 4; 40: IP precedence 5; 48: IP precedence 6; 56: IP precedence 7; 10: AF Class 1 (Low Drop); 12: AF Class 1 (Medium Drop); 14: AF Class 1 (High Drop); 18: AF Class 2 (Low Drop); 20: AF Class 2 (Medium Drop); 22: AF Class 2 (High Drop); 26: AF Class 3 (Low Drop); 28: AF Class 3 (Medium Drop); 30: AF Class 3 (High Drop); 34: AF Class 4 (Low Drop); 36: AF Class 4 (Medium Drop); 38: AF Class 4 (High Drop); 46: EF Class. | 
**Enable** | **bool** | The status of class rule. valid values are true or false. | 
**IpVersion** | **int32** | The IP Version of class rule should be a value as follows: 0: IPv4; 1: IPv6. | 
**LocalIp** | **string** | The ID of IP Group or IPv6 Group selected in the Local Address configuration. The ID can be obtained from &#39;Get group profile list&#39; interface. | 
**RemoteIp** | **string** | The ID of IP Group or IPv6 Group selected in the Remote Address configuration. The ID can be obtained from &#39;Get group profile list&#39; interface. | 
**ServiceType** | **string** | The ID of Gateway Qos Service selected in the Service Name configuration. The ID can be obtained from &#39;Get all Gateway QoS Service&#39;s ID and name info&#39; interface. | 

## Methods

### NewClassRuleOpenApiVO

`func NewClassRuleOpenApiVO(classType int32, dscp string, enable bool, ipVersion int32, localIp string, remoteIp string, serviceType string, ) *ClassRuleOpenApiVO`

NewClassRuleOpenApiVO instantiates a new ClassRuleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClassRuleOpenApiVOWithDefaults

`func NewClassRuleOpenApiVOWithDefaults() *ClassRuleOpenApiVO`

NewClassRuleOpenApiVOWithDefaults instantiates a new ClassRuleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassType

`func (o *ClassRuleOpenApiVO) GetClassType() int32`

GetClassType returns the ClassType field if non-nil, zero value otherwise.

### GetClassTypeOk

`func (o *ClassRuleOpenApiVO) GetClassTypeOk() (*int32, bool)`

GetClassTypeOk returns a tuple with the ClassType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassType

`func (o *ClassRuleOpenApiVO) SetClassType(v int32)`

SetClassType sets ClassType field to given value.


### GetDscp

`func (o *ClassRuleOpenApiVO) GetDscp() string`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *ClassRuleOpenApiVO) GetDscpOk() (*string, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *ClassRuleOpenApiVO) SetDscp(v string)`

SetDscp sets Dscp field to given value.


### GetEnable

`func (o *ClassRuleOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ClassRuleOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ClassRuleOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetIpVersion

`func (o *ClassRuleOpenApiVO) GetIpVersion() int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *ClassRuleOpenApiVO) GetIpVersionOk() (*int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *ClassRuleOpenApiVO) SetIpVersion(v int32)`

SetIpVersion sets IpVersion field to given value.


### GetLocalIp

`func (o *ClassRuleOpenApiVO) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *ClassRuleOpenApiVO) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *ClassRuleOpenApiVO) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.


### GetRemoteIp

`func (o *ClassRuleOpenApiVO) GetRemoteIp() string`

GetRemoteIp returns the RemoteIp field if non-nil, zero value otherwise.

### GetRemoteIpOk

`func (o *ClassRuleOpenApiVO) GetRemoteIpOk() (*string, bool)`

GetRemoteIpOk returns a tuple with the RemoteIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIp

`func (o *ClassRuleOpenApiVO) SetRemoteIp(v string)`

SetRemoteIp sets RemoteIp field to given value.


### GetServiceType

`func (o *ClassRuleOpenApiVO) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ClassRuleOpenApiVO) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ClassRuleOpenApiVO) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


