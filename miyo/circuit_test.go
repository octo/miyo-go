package miyo

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCircuitAllResponse(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/circuit-all.json")
	if err != nil {
		t.Fatal(err)
	}

	var got circuitAllResponse
	if err := json.Unmarshal([]byte(data), &got); err != nil {
		t.Fatal(err)
	}

	want := circuitAllResponse{
		Status: "success",
		ID:     0,
		Params: struct {
			Circuits map[string]Circuit "json:\"circuits\""
		}{
			Circuits: map[string]Circuit{
				"{a48b7871-4d8b-45fc-90d6-9225f2535926}": {
					ID:   "{a48b7871-4d8b-45fc-90d6-9225f2535926}",
					Name: "Rasen",
					Params: CircuitParams{
						AutomaticMode:   true,
						BorderBottom:    "40",
						BorderTop:       "60",
						ConsiderCharge:  true,
						ConsiderWeather: true,
						Day0:            "06:30-09:00;19:00-22:00",
						Day1:            "06:30-09:00;19:00-22:00",
						Day2:            "06:30-09:00;19:00-22:00",
						Day3:            "06:30-09:00;19:00-22:00",
						Day4:            "06:30-09:00;19:00-22:00",
						Day5:            "06:30-09:00;19:00-22:00",
						Day6:            "06:30-09:00;19:00-22:00",
						SoilType:        1,
					},
					SensorValve: SensorValve{
						Valve:   "{f223afe9-f8b9-46ae-8dcc-a868e96f2d2b}",
						Channel: 1,
					},
					Valves: map[string]Valve{
						"0": {
							ID:      "{f223afe9-f8b9-46ae-8dcc-a868e96f2d2b}",
							Channel: 1,
							Data: Device{
								Channel:    3276902,
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
					State: CircuitState{
						AutomaticMode:       true,
						IrrigationNextEnd:   1648710000,
						IrrigationNextStart: 1648701000,
					},
					Sensor: "{a6563d0a-28d2-432c-800f-838496a807de}",
					SensorData: Device{
						Channel:    3539041,
						ID:         "{a6563d0a-28d2-432c-800f-838496a807de}",
						IPv6:       "fe80::211:7d00:30:9ff6%zmd0",
						LastUpdate: 1648642654,
						Type:       "moistureOutdoor",
						State: DeviceState{
							ChargingLess:  true,
							LastResetType: -1,
							Moisture:      100,
							RSSI:          -200,
							SunWithinWeek: true,
						},
					},
				},
				"{b85746b6-5ebf-4038-9137-5dc4d2fa715a}": {
					ID:   "{b85746b6-5ebf-4038-9137-5dc4d2fa715a}",
					Name: "Rosen",
					Params: CircuitParams{
						AutomaticMode:   true,
						BorderBottom:    "50",
						BorderTop:       "70",
						ConsiderCharge:  true,
						ConsiderWeather: true,
						Day0:            "20:00-22:00",
						Day1:            "20:00-22:00",
						Day2:            "20:00-22:00",
						Day3:            "20:00-22:00",
						Day4:            "20:00-22:00",
						Day5:            "20:00-22:00",
						Day6:            "20:00-22:00",
						IrrigationType:  1,
						PlantType:       2,
						SoilType:        3,
					},
					SensorValve: SensorValve{
						Valve:   "{e00cc4d9-c0fa-41c3-aa8e-7ba088dc0f77}",
						Channel: 1,
					},
					Valves: map[string]Valve{
						"0": {
							ID:      "{e00cc4d9-c0fa-41c3-aa8e-7ba088dc0f77}",
							Channel: 1,
							Data: Device{
								Channel:    3145829,
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
						},
					},
					State: CircuitState{
						AutomaticMode:       true,
						IrrigationNextEnd:   1648756800,
						IrrigationNextStart: 1648749600,
					},
					Sensor: "{364795e9-df24-4b35-a5ab-53598fe38a13}",
					SensorData: Device{
						Channel:    3538995,
						ID:         "{364795e9-df24-4b35-a5ab-53598fe38a13}",
						IPv6:       "fe80::211:7d00:30:7a7a%zmd0",
						LastUpdate: 1648642654,
						Type:       "moistureOutdoor",
						State: DeviceState{
							ChargingLess:  true,
							LastResetType: -1,
							Moisture:      100,
							RSSI:          -200,
							SunWithinWeek: true,
						},
					},
				},
			},
		},
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("parsed response differs (-want/+got):\n%s", diff)
	}
}
