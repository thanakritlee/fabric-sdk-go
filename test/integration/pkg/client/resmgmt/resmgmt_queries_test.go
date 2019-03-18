/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package resmgmt

import (
	"testing"

	"github.com/thanakritlee/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/thanakritlee/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/thanakritlee/fabric-sdk-go/pkg/fabsdk"
)

func TestResMgmtClientQueries(t *testing.T) {

	// Using shared SDK instance to increase test speed.
	sdk := mainSDK
	testSetup := mainTestSetup
	chaincodeID := mainChaincodeID

	//prepare contexts
	org1AdminClientContext := sdk.Context(fabsdk.WithUser(org1AdminUser), fabsdk.WithOrg(org1Name))

	// Resource management client
	client, err := resmgmt.New(org1AdminClientContext)
	if err != nil {
		t.Fatalf("Failed to create new resource management client: %s", err)
	}

	// Our target for queries will be primary peer on this channel
	target := testSetup.Targets[0]

	testQueryConfigFromOrderer(t, testSetup.ChannelID, client)

	testInstalledChaincodes(t, chaincodeID, target, client)

	testInstantiatedChaincodes(t, testSetup.ChannelID, chaincodeID, target, client)

	testQueryChannels(t, testSetup.ChannelID, target, client)

}

func testInstantiatedChaincodes(t *testing.T, channelID string, ccID string, target string, client *resmgmt.Client) {

	chaincodeQueryResponse, err := client.QueryInstantiatedChaincodes(channelID, resmgmt.WithTargetEndpoints(target), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		t.Fatalf("QueryInstantiatedChaincodes return error: %s", err)
	}

	found := false
	for _, chaincode := range chaincodeQueryResponse.Chaincodes {
		t.Logf("**InstantiatedCC: %s", chaincode)
		if chaincode.Name == ccID {
			found = true
		}
	}

	if !found {
		t.Fatalf("QueryInstantiatedChaincodes failed to find instantiated %s chaincode", ccID)
	}
}

func testInstalledChaincodes(t *testing.T, ccID string, target string, client *resmgmt.Client) {

	chaincodeQueryResponse, err := client.QueryInstalledChaincodes(resmgmt.WithTargetEndpoints(target), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		t.Fatalf("QueryInstalledChaincodes return error: %s", err)
	}

	found := false
	for _, chaincode := range chaincodeQueryResponse.Chaincodes {
		t.Logf("**InstalledCC: %s", chaincode)
		if chaincode.Name == ccID {
			found = true
		}
	}

	if !found {
		t.Fatalf("QueryInstalledChaincodes failed to find installed %s chaincode", ccID)
	}
}

func testQueryChannels(t *testing.T, channelID string, target string, client *resmgmt.Client) {

	channelQueryResponse, err := client.QueryChannels(resmgmt.WithTargetEndpoints(target), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		t.Fatalf("QueryChannels return error: %s", err)
	}

	found := false
	for _, channel := range channelQueryResponse.Channels {
		t.Logf("**Channel: %s", channel)
		if channel.ChannelId == channelID {
			found = true
		}
	}

	if !found {
		t.Fatalf("QueryChannels failed, peer did not join '%s' channel", channelID)
	}

}

func testQueryConfigFromOrderer(t *testing.T, channelID string, client *resmgmt.Client) {
	expected := "orderer.example.com:7050"
	channelCfg, err := client.QueryConfigFromOrderer(channelID, resmgmt.WithOrdererEndpoint("orderer.example.com"), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		t.Fatalf("QueryConfig return error: %s", err)
	}
	if !contains(channelCfg.Orderers(), expected) {
		t.Fatalf("Expected orderer %s, got %s", expected, channelCfg.Orderers())
	}

	_, err = client.QueryConfigFromOrderer(channelID, resmgmt.WithOrdererEndpoint("non-existent"), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err == nil {
		t.Fatal("QueryConfig should have failed for invalid orderer")
	}

}

func contains(list []string, value string) bool {
	for _, e := range list {
		if e == value {
			return true
		}
	}
	return false
}
