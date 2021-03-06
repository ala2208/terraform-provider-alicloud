package cms

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetMyGroups invokes the cms.GetMyGroups API synchronously
// api document: https://help.aliyun.com/api/cms/getmygroups.html
func (client *Client) GetMyGroups(request *GetMyGroupsRequest) (response *GetMyGroupsResponse, err error) {
	response = CreateGetMyGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// GetMyGroupsWithChan invokes the cms.GetMyGroups API asynchronously
// api document: https://help.aliyun.com/api/cms/getmygroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMyGroupsWithChan(request *GetMyGroupsRequest) (<-chan *GetMyGroupsResponse, <-chan error) {
	responseChan := make(chan *GetMyGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMyGroups(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetMyGroupsWithCallback invokes the cms.GetMyGroups API asynchronously
// api document: https://help.aliyun.com/api/cms/getmygroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMyGroupsWithCallback(request *GetMyGroupsRequest, callback func(response *GetMyGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMyGroupsResponse
		var err error
		defer close(result)
		response, err = client.GetMyGroups(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// GetMyGroupsRequest is the request struct for api GetMyGroups
type GetMyGroupsRequest struct {
	*requests.RpcRequest
	InstanceId          string           `position:"Query" name:"InstanceId"`
	GroupId             requests.Integer `position:"Query" name:"GroupId"`
	Type                string           `position:"Query" name:"Type"`
	SelectContactGroups requests.Boolean `position:"Query" name:"SelectContactGroups"`
	BindUrl             string           `position:"Query" name:"BindUrl"`
	GroupName           string           `position:"Query" name:"GroupName"`
}

// GetMyGroupsResponse is the response struct for api GetMyGroups
type GetMyGroupsResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Group        Group  `json:"Group" xml:"Group"`
}

// CreateGetMyGroupsRequest creates a request to invoke GetMyGroups API
func CreateGetMyGroupsRequest() (request *GetMyGroupsRequest) {
	request = &GetMyGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2018-03-08", "GetMyGroups", "cms", "openAPI")
	return
}

// CreateGetMyGroupsResponse creates a response to parse from GetMyGroups response
func CreateGetMyGroupsResponse() (response *GetMyGroupsResponse) {
	response = &GetMyGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
