# \AlertAPI

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlerts**](AlertAPI.md#GetAlerts) | **Get** /alerts | 
[**PostAlerts**](AlertAPI.md#PostAlerts) | **Post** /alerts | 



## GetAlerts

> []GettableAlert GetAlerts(ctx).Active(active).Silenced(silenced).Inhibited(inhibited).Unprocessed(unprocessed).Filter(filter).Receiver(receiver).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/stream3715/am-client"
)

func main() {
	active := true // bool | Show active alerts (optional) (default to true)
	silenced := true // bool | Show silenced alerts (optional) (default to true)
	inhibited := true // bool | Show inhibited alerts (optional) (default to true)
	unprocessed := true // bool | Show unprocessed alerts (optional) (default to true)
	filter := []string{"Inner_example"} // []string | A list of matchers to filter alerts by (optional)
	receiver := "receiver_example" // string | A regex matching receivers to filter alerts by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertAPI.GetAlerts(context.Background()).Active(active).Silenced(silenced).Inhibited(inhibited).Unprocessed(unprocessed).Filter(filter).Receiver(receiver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.GetAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlerts`: []GettableAlert
	fmt.Fprintf(os.Stdout, "Response from `AlertAPI.GetAlerts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** | Show active alerts | [default to true]
 **silenced** | **bool** | Show silenced alerts | [default to true]
 **inhibited** | **bool** | Show inhibited alerts | [default to true]
 **unprocessed** | **bool** | Show unprocessed alerts | [default to true]
 **filter** | **[]string** | A list of matchers to filter alerts by | 
 **receiver** | **string** | A regex matching receivers to filter alerts by | 

### Return type

[**[]GettableAlert**](GettableAlert.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAlerts

> PostAlerts(ctx).Alerts(alerts).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/stream3715/am-client"
)

func main() {
	alerts := []openapiclient.PostableAlert{*openapiclient.NewPostableAlert(map[string]string{"key": "Inner_example"})} // []PostableAlert | The alerts to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AlertAPI.PostAlerts(context.Background()).Alerts(alerts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertAPI.PostAlerts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAlertsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alerts** | [**[]PostableAlert**](PostableAlert.md) | The alerts to create | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

