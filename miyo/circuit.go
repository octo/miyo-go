package miyo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Circuit struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Params      CircuitParams    `json:"params"`
	SensorValve SensorValve      `json:"sensorValve"`
	Valves      map[string]Valve `json:"valves"`
	State       CircuitState     `json:"stateTypes"`
	Sensor      string           `json:"sensor"`
	SensorData  Device           `json:"sensorData"`
}

// Status returns a human readable status message of the circuit.
func (c Circuit) Status() string {
	if !c.SensorData.State.Reachable {
		return "unreachable"
	}

	status := fmt.Sprintf("moisture %d", c.SensorData.State.Moisture)
	switch {
	case c.State.Irrigation:
		status += "; irrigation active"
	case c.SensorData.State.IrrigationNecessary:
		status += "; very dry"
	case c.SensorData.State.IrrigationPossible:
		status += "; dry"
	}

	return status
}

type SensorValve struct {
	Valve   string `json:"valve"`
	Channel int    `json:"channel"`
}

type Valve struct {
	ID      string `json:"valve"`
	Data    Device `json:"valveData"`
	Channel int    `json:"channel"`
}

type CircuitParams struct {
	AutomaticMode           bool     `json:"automaticMode"`
	BorderBottom            string   `json:"borderBottom"`
	BorderTop               string   `json:"borderTop"`
	ConsiderCharge          bool     `json:"considerCharge"`
	ConsiderMower           bool     `json:"considerMower"`
	ConsiderWeather         bool     `json:"considerWeather"`
	Day0                    string   `json:"day0"`
	Day1                    string   `json:"day1"`
	Day2                    string   `json:"day2"`
	Day3                    string   `json:"day3"`
	Day4                    string   `json:"day4"`
	Day5                    string   `json:"day5"`
	Day6                    string   `json:"day6"`
	IrrigationDelayForecast bool     `json:"irrigationDelayForecast"`
	IrrigationType          int      `json:"irrigationType"`
	LocationType            int      `json:"locationType"`
	PlantType               int      `json:"plantType"`
	SoilType                SoilType `json:"soilType"`
	TemperatureOffset       int      `json:"temperatureOffset"`
	ValveStaggering         bool     `json:"valveStaggering"`
}

type CircuitState struct {
	AutomaticMode        bool
	ExternBlock          bool
	Irrigation           bool
	IrrigationNextEnd    int
	IrrigationNextStart  int
	ValveStaggeringIndex int
	WinterMode           bool
}

// UnmarshalJSON parses the JSON-encoded data and stores the result in s.
func (s *CircuitState) UnmarshalJSON(d []byte) error {
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
		case "automaticMode":
			err = json.Unmarshal(kv.Value, &s.AutomaticMode)
		case "externBlock":
			err = json.Unmarshal(kv.Value, &s.ExternBlock)
		case "irrigation":
			err = json.Unmarshal(kv.Value, &s.Irrigation)
		case "irrigationNextEnd":
			err = json.Unmarshal(kv.Value, &s.IrrigationNextEnd)
		case "irrigationNextStart":
			err = json.Unmarshal(kv.Value, &s.IrrigationNextStart)
		case "valveStaggeringIndex":
			err = json.Unmarshal(kv.Value, &s.ValveStaggeringIndex)
		case "winterMode":
			err = json.Unmarshal(kv.Value, &s.WinterMode)

		default:
			err = errors.New("unknown key")
		}
		if err != nil {
			return fmt.Errorf("unmarshalling %s: %w", kv.Key, err)
		}
	}

	return nil
}

type SoilType int

const (
	SoilType_Loamy SoilType = iota
	SoilType_Sandy
	SoilType_LoamySandy
	SoilType_Unknown
)

func (s SoilType) String() string {
	names := map[SoilType]string{
		SoilType_Loamy:      "loamy",
		SoilType_Sandy:      "sandy",
		SoilType_LoamySandy: "loamy_sandy",
		SoilType_Unknown:    "unknown",
	}
	if name, ok := names[s]; ok {
		return name
	}
	return fmt.Sprintf("SoilType#%d", s)
}

// CircuitAll returns status information for all devices.
func (c *Conn) CircuitAll(ctx context.Context) ([]Circuit, error) {
	url := "http://" + c.host + "/api/circuit/all?apiKey=" + c.apiKey
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var car circuitAllResponse
	if err := json.NewDecoder(res.Body).Decode(&car); err != nil {
		return nil, err
	}

	if car.Status != "success" {
		return nil, fmt.Errorf("/api/circuit/all: %s", car.Error)
	}

	var ret []Circuit
	for _, c := range car.Params.Circuits {
		ret = append(ret, c)
	}

	return ret, nil
}

type circuitAllResponse struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Error  string `json:"error"`
	Params struct {
		Circuits map[string]Circuit `json:"circuits"`
	} `json:"params"`
}
