# ApRFScanResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel2g** | Pointer to [**[]RFScanRadio2g**](RFScanRadio2g.md) | Channel 2g | [optional] 
**Channel5g** | Pointer to [**[]RFScanRadio5g**](RFScanRadio5g.md) | Channel 5g | [optional] 
**Channel5g2** | Pointer to [**[]RFScanRadio5g2**](RFScanRadio5g2.md) | Channel 5g2 | [optional] 
**Channel6g** | Pointer to [**[]RFScanRadio6g**](RFScanRadio6g.md) | Channel 6g | [optional] 
**CurrentChan2g** | Pointer to **string** | The 2G channel of AP, such as 1,6,11,13. It should be within the range of 1–13. | [optional] 
**CurrentChan5g** | Pointer to **string** | The 5G channel of AP, such as 36,161. It should be within the range of 36–161. | [optional] 
**CurrentChan5g2** | Pointer to **string** | The 5G2 channel of AP,such as 36,161. It should be within the range of 36–161. | [optional] 
**CurrentChan6g** | Pointer to **string** | The 6G channel of AP,such as 36,161. It should be within the range of 36–161. | [optional] 
**CurrentChanW2g** | Pointer to **int32** | The 2g channel bandwidth of the AP. It should be a value as follows: 2:20MHz, 3: 40MHz | [optional] 
**CurrentChanW5g** | Pointer to **int32** | The 5g channel bandwidth of the AP. It should be a value as follows: 2:20MHz, 3: 40MHz, 5: 80MHz2 menas 20MHz, 3 means 40MHz, 5 means 80MHz | [optional] 
**CurrentChanW5g2** | Pointer to **int32** | The 5g2 channel bandwidth of the AP. It should be a value as follows: 2:20MHz, 3: 40MHz, 5: 80MHz | [optional] 
**CurrentChanW6g** | Pointer to **int32** | The 6g channel bandwidth of the AP. It should be a value as follows: 2:20MHz, 3: 40MHz, 5: 80MHz | [optional] 
**Status** | Pointer to **int32** | Status should be a value as follows: 0: the scan result is displayed; 1: no scan result; 2: Scanning | [optional] 
**Status2g** | Pointer to **int32** | Status of 2.4GHz band should be a value as follows: 0: the scan result is displayed; 1: no scan result; 2: Scanning | [optional] 
**Status5g** | Pointer to **int32** | Status of 5GHz band should be a value as follows: 0: the scan result is displayed; 1: no scan result; 2: Scanning | [optional] 
**Status5g2** | Pointer to **int32** | Status of 5GHz-2 band should be a value as follows: 0: the scan result is displayed; 1: no scan result; 2: Scanning | [optional] 
**Status6g** | Pointer to **int32** | Status of 6GHz band should be a value as follows: 0: the scan result is displayed; 1: no scan result; 2: Scanning | [optional] 
**Time** | Pointer to **int64** | The scan time(13 bits), Unit: ms | [optional] 
**Time2g** | Pointer to **int64** | The scan time(13 bits) of 5GHz-2 band, Unit: ms | [optional] 
**Time5g** | Pointer to **int64** | The scan time(13 bits) of 5GHz band, Unit: ms | [optional] 
**Time6g** | Pointer to **int64** | The scan time(13 bits) of 6GHz band, Unit: ms | [optional] 

## Methods

### NewApRFScanResult

`func NewApRFScanResult() *ApRFScanResult`

NewApRFScanResult instantiates a new ApRFScanResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApRFScanResultWithDefaults

`func NewApRFScanResultWithDefaults() *ApRFScanResult`

NewApRFScanResultWithDefaults instantiates a new ApRFScanResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel2g

`func (o *ApRFScanResult) GetChannel2g() []RFScanRadio2g`

GetChannel2g returns the Channel2g field if non-nil, zero value otherwise.

### GetChannel2gOk

`func (o *ApRFScanResult) GetChannel2gOk() (*[]RFScanRadio2g, bool)`

GetChannel2gOk returns a tuple with the Channel2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel2g

`func (o *ApRFScanResult) SetChannel2g(v []RFScanRadio2g)`

SetChannel2g sets Channel2g field to given value.

### HasChannel2g

`func (o *ApRFScanResult) HasChannel2g() bool`

HasChannel2g returns a boolean if a field has been set.

### GetChannel5g

`func (o *ApRFScanResult) GetChannel5g() []RFScanRadio5g`

GetChannel5g returns the Channel5g field if non-nil, zero value otherwise.

### GetChannel5gOk

`func (o *ApRFScanResult) GetChannel5gOk() (*[]RFScanRadio5g, bool)`

GetChannel5gOk returns a tuple with the Channel5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel5g

`func (o *ApRFScanResult) SetChannel5g(v []RFScanRadio5g)`

SetChannel5g sets Channel5g field to given value.

### HasChannel5g

`func (o *ApRFScanResult) HasChannel5g() bool`

HasChannel5g returns a boolean if a field has been set.

### GetChannel5g2

`func (o *ApRFScanResult) GetChannel5g2() []RFScanRadio5g2`

GetChannel5g2 returns the Channel5g2 field if non-nil, zero value otherwise.

### GetChannel5g2Ok

`func (o *ApRFScanResult) GetChannel5g2Ok() (*[]RFScanRadio5g2, bool)`

GetChannel5g2Ok returns a tuple with the Channel5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel5g2

`func (o *ApRFScanResult) SetChannel5g2(v []RFScanRadio5g2)`

SetChannel5g2 sets Channel5g2 field to given value.

### HasChannel5g2

`func (o *ApRFScanResult) HasChannel5g2() bool`

HasChannel5g2 returns a boolean if a field has been set.

### GetChannel6g

`func (o *ApRFScanResult) GetChannel6g() []RFScanRadio6g`

GetChannel6g returns the Channel6g field if non-nil, zero value otherwise.

### GetChannel6gOk

`func (o *ApRFScanResult) GetChannel6gOk() (*[]RFScanRadio6g, bool)`

GetChannel6gOk returns a tuple with the Channel6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel6g

`func (o *ApRFScanResult) SetChannel6g(v []RFScanRadio6g)`

SetChannel6g sets Channel6g field to given value.

### HasChannel6g

`func (o *ApRFScanResult) HasChannel6g() bool`

HasChannel6g returns a boolean if a field has been set.

### GetCurrentChan2g

`func (o *ApRFScanResult) GetCurrentChan2g() string`

GetCurrentChan2g returns the CurrentChan2g field if non-nil, zero value otherwise.

### GetCurrentChan2gOk

`func (o *ApRFScanResult) GetCurrentChan2gOk() (*string, bool)`

GetCurrentChan2gOk returns a tuple with the CurrentChan2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChan2g

`func (o *ApRFScanResult) SetCurrentChan2g(v string)`

SetCurrentChan2g sets CurrentChan2g field to given value.

### HasCurrentChan2g

`func (o *ApRFScanResult) HasCurrentChan2g() bool`

HasCurrentChan2g returns a boolean if a field has been set.

### GetCurrentChan5g

`func (o *ApRFScanResult) GetCurrentChan5g() string`

GetCurrentChan5g returns the CurrentChan5g field if non-nil, zero value otherwise.

### GetCurrentChan5gOk

`func (o *ApRFScanResult) GetCurrentChan5gOk() (*string, bool)`

GetCurrentChan5gOk returns a tuple with the CurrentChan5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChan5g

`func (o *ApRFScanResult) SetCurrentChan5g(v string)`

SetCurrentChan5g sets CurrentChan5g field to given value.

### HasCurrentChan5g

`func (o *ApRFScanResult) HasCurrentChan5g() bool`

HasCurrentChan5g returns a boolean if a field has been set.

### GetCurrentChan5g2

`func (o *ApRFScanResult) GetCurrentChan5g2() string`

GetCurrentChan5g2 returns the CurrentChan5g2 field if non-nil, zero value otherwise.

### GetCurrentChan5g2Ok

`func (o *ApRFScanResult) GetCurrentChan5g2Ok() (*string, bool)`

GetCurrentChan5g2Ok returns a tuple with the CurrentChan5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChan5g2

`func (o *ApRFScanResult) SetCurrentChan5g2(v string)`

SetCurrentChan5g2 sets CurrentChan5g2 field to given value.

### HasCurrentChan5g2

`func (o *ApRFScanResult) HasCurrentChan5g2() bool`

HasCurrentChan5g2 returns a boolean if a field has been set.

### GetCurrentChan6g

`func (o *ApRFScanResult) GetCurrentChan6g() string`

GetCurrentChan6g returns the CurrentChan6g field if non-nil, zero value otherwise.

### GetCurrentChan6gOk

`func (o *ApRFScanResult) GetCurrentChan6gOk() (*string, bool)`

GetCurrentChan6gOk returns a tuple with the CurrentChan6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChan6g

`func (o *ApRFScanResult) SetCurrentChan6g(v string)`

SetCurrentChan6g sets CurrentChan6g field to given value.

### HasCurrentChan6g

`func (o *ApRFScanResult) HasCurrentChan6g() bool`

HasCurrentChan6g returns a boolean if a field has been set.

### GetCurrentChanW2g

`func (o *ApRFScanResult) GetCurrentChanW2g() int32`

GetCurrentChanW2g returns the CurrentChanW2g field if non-nil, zero value otherwise.

### GetCurrentChanW2gOk

`func (o *ApRFScanResult) GetCurrentChanW2gOk() (*int32, bool)`

GetCurrentChanW2gOk returns a tuple with the CurrentChanW2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChanW2g

`func (o *ApRFScanResult) SetCurrentChanW2g(v int32)`

SetCurrentChanW2g sets CurrentChanW2g field to given value.

### HasCurrentChanW2g

`func (o *ApRFScanResult) HasCurrentChanW2g() bool`

HasCurrentChanW2g returns a boolean if a field has been set.

### GetCurrentChanW5g

`func (o *ApRFScanResult) GetCurrentChanW5g() int32`

GetCurrentChanW5g returns the CurrentChanW5g field if non-nil, zero value otherwise.

### GetCurrentChanW5gOk

`func (o *ApRFScanResult) GetCurrentChanW5gOk() (*int32, bool)`

GetCurrentChanW5gOk returns a tuple with the CurrentChanW5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChanW5g

`func (o *ApRFScanResult) SetCurrentChanW5g(v int32)`

SetCurrentChanW5g sets CurrentChanW5g field to given value.

### HasCurrentChanW5g

`func (o *ApRFScanResult) HasCurrentChanW5g() bool`

HasCurrentChanW5g returns a boolean if a field has been set.

### GetCurrentChanW5g2

`func (o *ApRFScanResult) GetCurrentChanW5g2() int32`

GetCurrentChanW5g2 returns the CurrentChanW5g2 field if non-nil, zero value otherwise.

### GetCurrentChanW5g2Ok

`func (o *ApRFScanResult) GetCurrentChanW5g2Ok() (*int32, bool)`

GetCurrentChanW5g2Ok returns a tuple with the CurrentChanW5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChanW5g2

`func (o *ApRFScanResult) SetCurrentChanW5g2(v int32)`

SetCurrentChanW5g2 sets CurrentChanW5g2 field to given value.

### HasCurrentChanW5g2

`func (o *ApRFScanResult) HasCurrentChanW5g2() bool`

HasCurrentChanW5g2 returns a boolean if a field has been set.

### GetCurrentChanW6g

`func (o *ApRFScanResult) GetCurrentChanW6g() int32`

GetCurrentChanW6g returns the CurrentChanW6g field if non-nil, zero value otherwise.

### GetCurrentChanW6gOk

`func (o *ApRFScanResult) GetCurrentChanW6gOk() (*int32, bool)`

GetCurrentChanW6gOk returns a tuple with the CurrentChanW6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChanW6g

`func (o *ApRFScanResult) SetCurrentChanW6g(v int32)`

SetCurrentChanW6g sets CurrentChanW6g field to given value.

### HasCurrentChanW6g

`func (o *ApRFScanResult) HasCurrentChanW6g() bool`

HasCurrentChanW6g returns a boolean if a field has been set.

### GetStatus

`func (o *ApRFScanResult) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApRFScanResult) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApRFScanResult) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApRFScanResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatus2g

`func (o *ApRFScanResult) GetStatus2g() int32`

GetStatus2g returns the Status2g field if non-nil, zero value otherwise.

### GetStatus2gOk

`func (o *ApRFScanResult) GetStatus2gOk() (*int32, bool)`

GetStatus2gOk returns a tuple with the Status2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus2g

`func (o *ApRFScanResult) SetStatus2g(v int32)`

SetStatus2g sets Status2g field to given value.

### HasStatus2g

`func (o *ApRFScanResult) HasStatus2g() bool`

HasStatus2g returns a boolean if a field has been set.

### GetStatus5g

`func (o *ApRFScanResult) GetStatus5g() int32`

GetStatus5g returns the Status5g field if non-nil, zero value otherwise.

### GetStatus5gOk

`func (o *ApRFScanResult) GetStatus5gOk() (*int32, bool)`

GetStatus5gOk returns a tuple with the Status5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus5g

`func (o *ApRFScanResult) SetStatus5g(v int32)`

SetStatus5g sets Status5g field to given value.

### HasStatus5g

`func (o *ApRFScanResult) HasStatus5g() bool`

HasStatus5g returns a boolean if a field has been set.

### GetStatus5g2

`func (o *ApRFScanResult) GetStatus5g2() int32`

GetStatus5g2 returns the Status5g2 field if non-nil, zero value otherwise.

### GetStatus5g2Ok

`func (o *ApRFScanResult) GetStatus5g2Ok() (*int32, bool)`

GetStatus5g2Ok returns a tuple with the Status5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus5g2

`func (o *ApRFScanResult) SetStatus5g2(v int32)`

SetStatus5g2 sets Status5g2 field to given value.

### HasStatus5g2

`func (o *ApRFScanResult) HasStatus5g2() bool`

HasStatus5g2 returns a boolean if a field has been set.

### GetStatus6g

`func (o *ApRFScanResult) GetStatus6g() int32`

GetStatus6g returns the Status6g field if non-nil, zero value otherwise.

### GetStatus6gOk

`func (o *ApRFScanResult) GetStatus6gOk() (*int32, bool)`

GetStatus6gOk returns a tuple with the Status6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus6g

`func (o *ApRFScanResult) SetStatus6g(v int32)`

SetStatus6g sets Status6g field to given value.

### HasStatus6g

`func (o *ApRFScanResult) HasStatus6g() bool`

HasStatus6g returns a boolean if a field has been set.

### GetTime

`func (o *ApRFScanResult) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ApRFScanResult) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ApRFScanResult) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ApRFScanResult) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetTime2g

`func (o *ApRFScanResult) GetTime2g() int64`

GetTime2g returns the Time2g field if non-nil, zero value otherwise.

### GetTime2gOk

`func (o *ApRFScanResult) GetTime2gOk() (*int64, bool)`

GetTime2gOk returns a tuple with the Time2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime2g

`func (o *ApRFScanResult) SetTime2g(v int64)`

SetTime2g sets Time2g field to given value.

### HasTime2g

`func (o *ApRFScanResult) HasTime2g() bool`

HasTime2g returns a boolean if a field has been set.

### GetTime5g

`func (o *ApRFScanResult) GetTime5g() int64`

GetTime5g returns the Time5g field if non-nil, zero value otherwise.

### GetTime5gOk

`func (o *ApRFScanResult) GetTime5gOk() (*int64, bool)`

GetTime5gOk returns a tuple with the Time5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime5g

`func (o *ApRFScanResult) SetTime5g(v int64)`

SetTime5g sets Time5g field to given value.

### HasTime5g

`func (o *ApRFScanResult) HasTime5g() bool`

HasTime5g returns a boolean if a field has been set.

### GetTime6g

`func (o *ApRFScanResult) GetTime6g() int64`

GetTime6g returns the Time6g field if non-nil, zero value otherwise.

### GetTime6gOk

`func (o *ApRFScanResult) GetTime6gOk() (*int64, bool)`

GetTime6gOk returns a tuple with the Time6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime6g

`func (o *ApRFScanResult) SetTime6g(v int64)`

SetTime6g sets Time6g field to given value.

### HasTime6g

`func (o *ApRFScanResult) HasTime6g() bool`

HasTime6g returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


