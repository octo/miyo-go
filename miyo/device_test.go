package miyo

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDeviceAllResponse(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/device-all.json")
	if err != nil {
		t.Fatal(err)
	}

	var got deviceAllResponse
	if err := json.Unmarshal([]byte(data), &got); err != nil {
		t.Fatal(err)
	}

	want := deviceAllResponse{
		Status: "success",
		ID:     0,
		Params: struct {
			Devices map[string]Device "json:\"devices\""
		}{
			Devices: map[string]Device{
				"{364795e9-df24-4b35-a5ab-53598fe38a13};1": {
					Channel:    1,
					ID:         "{364795e9-df24-4b35-a5ab-53598fe38a13}",
					Type:       "moistureOutdoor",
					IPv6:       "fe80::211:7d00:30:7a7a%zmd0",
					LastUpdate: 1648642654,
					State: DeviceState{
						ChargingLess:  true,
						LastResetType: -1,
						Moisture:      100,
						RSSI:          -200,
						SunWithinWeek: true,
					},
				},
				"{a6563d0a-28d2-432c-800f-838496a807de};1": {
					Channel:    1,
					ID:         "{a6563d0a-28d2-432c-800f-838496a807de}",
					Type:       "moistureOutdoor",
					IPv6:       "fe80::211:7d00:30:9ff6%zmd0",
					LastUpdate: 1648642654,
					State: DeviceState{
						ChargingLess:  true,
						LastResetType: -1,
						Moisture:      100,
						RSSI:          -200,
						SunWithinWeek: true,
					},
				},
				"{e00cc4d9-c0fa-41c3-aa8e-7ba088dc0f77};1": {
					Channel:    1,
					ID:         "{e00cc4d9-c0fa-41c3-aa8e-7ba088dc0f77}",
					Type:       "valve",
					IPv6:       "fe80::211:7d00:30:9d3d%zmd0",
					LastUpdate: 1648642654,
					State: DeviceState{
						ChargingLess:      true,
						LastResetType:     -1,
						RSSI:              -200,
						SunWithinWeek:     true,
						ValveInitialClose: true,
					},
				},
				"{f223afe9-f8b9-46ae-8dcc-a868e96f2d2b};1": {
					Channel:    1,
					ID:         "{f223afe9-f8b9-46ae-8dcc-a868e96f2d2b}",
					Type:       "valve",
					IPv6:       "fe80::211:7d00:30:9d58%zmd0",
					LastUpdate: 1648642654,
					State: DeviceState{
						ChargingLess:      true,
						LastResetType:     -1,
						RSSI:              -200,
						SunWithinWeek:     true,
						ValveInitialClose: true,
					},
				},
			},
		},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parsed response differs (-want/+got):\n%s", diff)
	}
}
