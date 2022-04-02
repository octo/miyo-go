package miyo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type deviceAllResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Params struct {
		Devices map[string]Device `json:"devices"`
	} `json:"params"`
}

// Device represents a moisture sensor or a valve.
type Device struct {
	Channel    int         `json:"channel"`
	ID         string      `json:"id"`
	Type       string      `json:"deviceTypeId"`
	Firmware   string      `json:"firmware"`
	IPv6       string      `json:"ipv6"`
	LastUpdate int         `json:"lastUpdate"`
	State      DeviceState `json:"stateTypes"`
}

func (d Device) Status() string {
	if !d.State.Reachable {
		return "unreachable"
	}

	switch d.Type {
	case "valve":
		if d.State.ValveStatus {
			return "open"
		}
		return "closed"
	case "moistureOutdoor":
		return fmt.Sprintf("moisture %d", d.State.Moisture)
	default:
		return "unknown type"
	}
}

// Devices returns status information for all devices.
func (c *Conn) Devices(ctx context.Context) ([]Device, error) {
	url := "http://" + c.host + "/api/device/all?apiKey=" + c.apiKey
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var dar deviceAllResponse
	if err := json.NewDecoder(res.Body).Decode(&dar); err != nil {
		return nil, err
	}

	var devs []Device
	for _, d := range dar.Params.Devices {
		devs = append(devs, d)
	}

	return devs, nil
}

// DeviceState holds the state of a device.
type DeviceState struct {
	// Valve is to be closed
	ValveInitialClose bool
	// Valve is open
	ValveStatus bool
	// Valve is to be opened
	OpenValve bool
	// Unix Timestamp from last watering start
	LastIrrigationStart int
	// Unix Timestamp of the last end of irrigation
	LastIrrigationEnd int
	// Duration of the last irrigation
	LastIrrigationDuration int
	// Signal strength of the device
	RSSI int
	// Device is accessible from the cube
	Reachable bool
	// Solar voltage of the device
	SolarVoltage int
	// Current always true
	SunWithinWeek bool
	// Device has little battery
	LowPower bool
	// Installation of an update is possible
	OTAUPossible bool
	// Update progress
	OTAUProgress int
	// Update status of the device
	OTAUStatus string
	// Winter mode activated
	WinterMode bool
	// Loading time per day within the last week
	ChargingDurationDay int
	// Device charging
	Charging bool
	// Device does not charge enough
	ChargingLess bool
	// Time of the last reset of the device
	LastResetTime int
	// Type of the last reset of the device
	LastResetType int
	// Sensor humidity %
	Moisture int
	// Brightness of the sensor in lux
	Brightness int
	// Temperature of the sensor (near ground) in °C
	Temperature int
	// Frequency of the humidity sensor
	Frequency int
	// Irrigation needed (soil very dry)
	IrrigationNecessary bool
	// Irrigation possible (soil dry)
	IrrigationPossible bool
	// Temperature offset of the sensor
	TemperatureOffset int
}

// UnmarshalJSON parses the JSON-encoded data and stores the result in s.
func (s *DeviceState) UnmarshalJSON(d []byte) error {
	// {
	//  "0":{"type":"moisture","value":100},
	//  "1":{"type":"brightness","value":0},
	//  ...

	var parsed map[string]struct {
		Key   string          `json:"type"`
		Value json.RawMessage `json:"value"`
	}
	if err := json.Unmarshal(d, &parsed); err != nil {
		return err
	}

	for _, kv := range parsed {
		var err error
		switch kv.Key {
		case "valveInitialClose":
			err = json.Unmarshal(kv.Value, &s.ValveInitialClose)
		case "valveStatus":
			err = json.Unmarshal(kv.Value, &s.ValveStatus)
		case "openValve":
			err = json.Unmarshal(kv.Value, &s.OpenValve)
		case "lastIrrigationStart":
			err = json.Unmarshal(kv.Value, &s.LastIrrigationStart)
		case "lastIrrigationEnd":
			err = json.Unmarshal(kv.Value, &s.LastIrrigationEnd)
		case "lastIrrigationDuration":
			err = json.Unmarshal(kv.Value, &s.LastIrrigationDuration)
		case "rssi":
			err = json.Unmarshal(kv.Value, &s.RSSI)
		case "reachable":
			err = json.Unmarshal(kv.Value, &s.Reachable)
		case "solarVoltage":
			err = json.Unmarshal(kv.Value, &s.SolarVoltage)
		case "sunWithinWeek":
			err = json.Unmarshal(kv.Value, &s.SunWithinWeek)
		case "lowPower":
			err = json.Unmarshal(kv.Value, &s.LowPower)
		case "otauPossible":
			err = json.Unmarshal(kv.Value, &s.OTAUPossible)
		case "otauProgress":
			err = json.Unmarshal(kv.Value, &s.OTAUProgress)
		case "otauStatus":
			err = json.Unmarshal(kv.Value, &s.OTAUStatus)
		case "winterMode":
			err = json.Unmarshal(kv.Value, &s.WinterMode)
		case "chargingDurationDay":
			err = json.Unmarshal(kv.Value, &s.ChargingDurationDay)
		case "charging":
			err = json.Unmarshal(kv.Value, &s.Charging)
		case "chargingLess":
			err = json.Unmarshal(kv.Value, &s.ChargingLess)
		case "lastResetTime":
			err = json.Unmarshal(kv.Value, &s.LastResetTime)
		case "lastResetType":
			err = json.Unmarshal(kv.Value, &s.LastResetType)
		case "moisture":
			err = json.Unmarshal(kv.Value, &s.Moisture)
		case "brightness":
			err = json.Unmarshal(kv.Value, &s.Brightness)
		case "temperature":
			err = json.Unmarshal(kv.Value, &s.Temperature)
		case "frequency":
			err = json.Unmarshal(kv.Value, &s.Frequency)
		case "irrigationNecessary":
			err = json.Unmarshal(kv.Value, &s.IrrigationNecessary)
		case "irrigationPossible":
			err = json.Unmarshal(kv.Value, &s.IrrigationPossible)
		case "temperatureOffset":
			err = json.Unmarshal(kv.Value, &s.TemperatureOffset)
		default:
			err = errors.New("unknown key")
		}
		if err != nil {
			return fmt.Errorf("unmarshalling %s: %w", kv.Key, err)
		}
	}

	return nil
}
