/*
VPC API

Testing DefaultApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package vpc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mulesoft-anypoint/anypoint-client-go/vpc"
)

func Test_vpc_DefaultApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DefaultApiService OrganizationsOrgIdVpcsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdVpcsGet(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdVpcsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdVpcsPost(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdVpcsVpcIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var vpcId string

		httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdVpcsVpcIdDelete(context.Background(), orgId, vpcId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdVpcsVpcIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var vpcId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdVpcsVpcIdGet(context.Background(), orgId, vpcId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DefaultApiService OrganizationsOrgIdVpcsVpcIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var vpcId string

		resp, httpRes, err := apiClient.DefaultApi.OrganizationsOrgIdVpcsVpcIdPut(context.Background(), orgId, vpcId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}