// Copyright 2018 NetApp, Inc. All Rights Reserved.

package solidfire

import (
	"strings"
	"testing"

	drivers "github.com/netapp/trident/storage_drivers"
	"github.com/netapp/trident/storage_drivers/solidfire/api"
)

const (
	configEndpoint  = "https://admin:solidfire@10.63.171.151/json-rpc/7.0"
	minimumEndpoint = "https://admin:solidfire@10.63.171.151/json-rpc/" + sfMinimumAPIVersion
	newerEndpoint   = "https://admin:solidfire@10.63.171.151/json-rpc/9.0"
)

func getDriver() *SANStorageDriver {
	return &SANStorageDriver{
		Config: drivers.SolidfireStorageDriverConfig{
			CommonStorageDriverConfig: &drivers.CommonStorageDriverConfig{
				Version:           1,
				StorageDriverName: "solidfire-san",
			},
			SolidfireStorageDriverConfigDefaults: drivers.SolidfireStorageDriverConfigDefaults{
				CommonStorageDriverConfigDefaults: drivers.CommonStorageDriverConfigDefaults{
					Size: "100GiB",
				},
			},
			TenantName:     "test",
			EndPoint:       configEndpoint,
			SVIP:           "10.63.171.153:3260",
			InitiatorIFace: "default",
			Types: &[]api.VolType{
				{
					Type: "Bronze", QOS: api.QoS{
						MinIOPS:   1000,
						MaxIOPS:   2000,
						BurstIOPS: 4000,
					}},
				{
					Type: "Silver",
					QOS: api.QoS{
						MinIOPS:   4000,
						MaxIOPS:   6000,
						BurstIOPS: 8000,
					},
				},
				{
					Type: "Gold",
					QOS: api.QoS{
						MinIOPS:   6000,
						MaxIOPS:   8000,
						BurstIOPS: 10000,
					},
				},
			},
		},
	}
}

func TestGetExternalConfig(t *testing.T) {
	driver := getDriver()
	newConfig := driver.GetExternalConfig().(drivers.SolidfireStorageDriverConfig)
	if newConfig.EndPoint == driver.Config.EndPoint {
		t.Errorf("EndPoints are equal; expected different. Got: %s", newConfig.EndPoint)
	}
	if strings.Contains(newConfig.EndPoint, "admin") {
		t.Error("Username not removed from external config endpoint.")
	}
	if strings.Contains(newConfig.EndPoint, "solidfire") {
		t.Error("Password not removed from external config endpoint.")
	}
	t.Log("External config endpoint:  ", newConfig.EndPoint)
	if !strings.Contains(driver.Config.EndPoint, "admin") {
		t.Error("Username removed from main config endpoint.")
	}
	if !strings.Contains(driver.Config.EndPoint, "solidfire") {
		t.Error("Password removed from main config endpoint.")
	}
	t.Log("Main config endpoint:  ", driver.Config.EndPoint)
}

func TestUpgradeOlderEndpointAPIVersion(t *testing.T) {
	driver := getDriver()
	endpoint, err := driver.getEndpoint(&driver.Config)
	if err != nil {
		t.Errorf("Received error from getEndpoint: %v", err)
	}
	if endpoint != minimumEndpoint {
		t.Error("Client endpoint not changed to minimum version.")
	}
}

func TestNoUpgradeNewerEndpointAPIVersion(t *testing.T) {
	driver := getDriver()
	driver.Config.EndPoint = newerEndpoint
	endpoint, err := driver.getEndpoint(&driver.Config)
	if err != nil {
		t.Errorf("Received error from getEndpoint: %v", err)
	}
	if endpoint != newerEndpoint {
		t.Error("Client endpoint changed to minimum version.")
	}
}

func TestGetEndpointCredentials(t *testing.T) {

	// good path, should work
	driver := getDriver()
	driver.Config.EndPoint = newerEndpoint
	username, password, err := driver.getEndpointCredentials(driver.Config)
	if err != nil {
		t.Errorf("Received error from getEndpointCredentials: %v", err)
	}

	if username == "" {
		t.Errorf("Received empty username from getEndpointCredentials: %v", err)
	}
	if password == "" {
		t.Errorf("Received empty password from getEndpointCredentials: %v", err)
	}

	if username != "admin" {
		t.Errorf("Received unexpected username %v from getEndpointCredentials: %v", username, err)
	}
	if password != "solidfire" {
		t.Errorf("Received unexpected password %v from getEndpointCredentials: %v", password, err)
	}

	// invalid URL, should error
	badEndpoint := "://admin:solidfire@10.63.171.151/json-rpc/9.0"
	driver.Config.EndPoint = badEndpoint
	username, password, err = driver.getEndpointCredentials(driver.Config)
	if err == nil {
		t.Errorf("Should have received error from getEndpointCredentials: %v", err)
	}
	if err != nil && err.Error() != "could not determine credentials" {
		t.Errorf("Received error from getEndpointCredentials: %v", err)
	}
	if username != "" {
		t.Errorf("Received unexpected username %v from getEndpointCredentials: %v", username, err)
	}
	if password != "" {
		t.Errorf("Received unexpected password %v from getEndpointCredentials: %v", password, err)
	}

	// no username, password (wouldn't work, but we should handle)
	badEndpoint2 := "https://10.63.171.151/json-rpc/9.0"
	driver.Config.EndPoint = badEndpoint2
	username, password, err = driver.getEndpointCredentials(driver.Config)
	if username != "" {
		t.Errorf("Received unexpected username %v from getEndpointCredentials: %v", username, err)
	}
	if password != "" {
		t.Errorf("Received unexpected password %v from getEndpointCredentials: %v", password, err)
	}
}
