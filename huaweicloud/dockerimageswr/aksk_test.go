package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	vpc "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vpc/v2/model"
)

func RequestHandler(request http.Request) {
	fmt.Println(request)
}

func ResponseHandler(response http.Response) {
	fmt.Println(response)
}

func TestOfficialTemplate(t *testing.T) {
	client := vpc.NewVpcClient(
		vpc.VpcClientBuilder().
			WithEndpoint("{your endpoint}").
			WithCredential(
				basic.NewCredentialsBuilder().
					WithAk("{your ak string}").
					WithSk("{your sk string}").
					WithProjectId("{your project id}").
					Build()).
			WithHttpConfig(config.DefaultHttpConfig().
				WithIgnoreSSLVerification(true).
				WithHttpHandler(httphandler.
					NewHttpHandler().
					AddRequestHandler(RequestHandler).
					AddResponseHandler(ResponseHandler))).
			Build())

	limit := int32(1)
	request := &model.ListVpcsRequest{
		Limit: &limit,
	}
	response, err := client.ListVpcs(request)
	if err == nil {
		fmt.Printf("%+v\n\n", response.Vpcs)
	} else {
		fmt.Println(err)
	}
}

// User Name,Access Key Id,Secret Access Key
// "hw098403457",MIKVPONSDECXIBX01YDM,lskMAw3N6GJIF2XFTqpwwNbxQ9LVVRwE9SJHj9Hp
func TestMyAccount(t *testing.T) {
	reginstr := "cn-south-1"
	endpoint := "vpc." + reginstr + ".myhuaweicloud.com"
	client := vpc.NewVpcClient(
		vpc.VpcClientBuilder().
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

	limit := int32(1)
	request := &model.ListVpcsRequest{
		Limit: &limit,
	}
	response, err := client.ListVpcs(request)
	if err == nil {
		fmt.Printf("%+v\n\n", response.Vpcs)
	} else {
		fmt.Println(err)
	}
}
