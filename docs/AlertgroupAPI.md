# \AlertgroupAPI

All URIs are relative to *http://localhost/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAlertGroups**](AlertgroupAPI.md#GetAlertGroups) | **Get** /alerts/groups | 



## GetAlertGroups

> []AlertGroup GetAlertGroups(ctx).Active(active).Silenced(silenced).Inhibited(inhibited).Filter(filter).Receiver(receiver).Execute()





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
	filter := []string{"Inner_example"} // []string | A list of matchers to filter alerts by (optional)
	receiver := "receiver_example" // string | A regex matching receivers to filter alerts by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AlertgroupAPI.GetAlertGroups(context.Background()).Active(active).Silenced(silenced).Inhibited(inhibited).Filter(filter).Receiver(receiver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AlertgroupAPI.GetAlertGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlertGroups`: []AlertGroup
	fmt.Fprintf(os.Stdout, "Response from `AlertgroupAPI.GetAlertGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **bool** | Show active alerts | [default to true]
 **silenced** | **bool** | Show silenced alerts | [default to true]
 **inhibited** | **bool** | Show inhibited alerts | [default to true]
 **filter** | **[]string** | A list of matchers to filter alerts by | 
 **receiver** | **string** | A regex matching receivers to filter alerts by | 

### Return type

[**[]AlertGroup**](AlertGroup.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

