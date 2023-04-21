package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	swr "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/swr/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/swr/v2/model"
)

func RequestHandlerSWR(request http.Request) {
	fmt.Println(request)
}

func ResponseHandlerSWR(response http.Response) {
	fmt.Println(response)
}

// 华为云sdk 最佳实践 基准测试
// User Name,Access Key Id,Secret Access Key
// "hw098403457",MIKVPONSDECXIBX01YDM,lskMAw3N6GJIF2XFTqpwwNbxQ9LVVRwE9SJHj9Hp
func TestMyAccountSWR(t *testing.T) {
	reginstr := "cn-south-1"
	endpoint := "swr-api." + reginstr + ".myhuaweicloud.com"
	client := swr.NewSwrClient(
		swr.SwrClientBuilder().
			WithEndpoint(endpoint).
			WithCredential(
				basic.NewCredentialsBuilder().
					WithAk("MIKVPONSDECXIBX01YDM").
					WithSk("lskMAw3N6GJIF2XFTqpwwNbxQ9LVVRwE9SJHj9Hp").
					WithProjectId("978e85f4697c4fcd80cadf6bc7061b2a").
					Build()).
			WithHttpConfig(config.DefaultHttpConfig().
				WithIgnoreSSLVerification(true).
				WithHttpHandler(httphandler.
					NewHttpHandler().
					AddRequestHandler(RequestHandler).
					AddResponseHandler(ResponseHandler))).
			Build())

	// limit := int32(1)
	request := &model.ListImageAutoSyncReposDetailsRequest{}
	response, err := client.ListImageAutoSyncReposDetails(request)
	if err == nil {
		fmt.Printf("%+v\n\n", response.Body)
	} else {
		fmt.Println(err)
	}
}
