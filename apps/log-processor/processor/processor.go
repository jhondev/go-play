package processor

import (
	"bufio"
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"logprocessor.37Widgets.co/stats"
)

// sensor types
const (
	Thermometer = "thermometer"
	Humidity    = "humidity"
)

// quality levels
const (
	Discard      = "discard"
	OK           = "OK"
	Precise      = "precise"
	VeryPrecise  = "very precise"
	UltraPrecise = "ultra precise"
)

type Sensor struct {
	_type string
	name  string
	cfg   *SensorCfg
}

type SensorCfg struct {
	ruleFunc func(total float64, counter int, readings []float64, ref float64) string
}

// Having a map with sensor configs allow us
// to add new sensors and quality calculation rules easier
var sensorCfgs = map[string]SensorCfg{
	Thermometer: {ruleFunc: PrecisionRule},
	Humidity:    {ruleFunc: DiscardRule},
}

func PrecisionRule(total float64, counter int, readings []float64, ref float64) string {
	mean := total / float64(counter)
	diff := math.Abs(ref - mean)
	sdev := stats.StandardDeviation(readings, counter, &mean)

	if diff <= 0.5 && sdev < 3 {
		return UltraPrecise
	}
	if diff <= 0.5 && sdev < 5 {
		return VeryPrecise
	}
	return Precise
}

func DiscardRule(_ float64, _ int, readings []float64, ref float64) string {
	minDiff := ref * 0.01
	for _, r := range readings {
		diff := math.Abs(ref - r)
		if diff > minDiff {
			return Discard
		}
	}
	return OK
}

// For better performance we make calculations while we are reading the file
// this could change depending on the scenario and if a more decoupled code is wanted
// the calculation rules are decoupled by injection though, given certain flexibility
// to add more sensors
func ProcessLog(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	refs, err := ReadRefLine(scanner.Text())
	if err != nil {
		return nil, err
	}
	var sensor *Sensor
	var rtotal float64
	var rcounter int
	var readings []float64
	qualityCtrl := make(map[string]string)

	for scanner.Scan() {
		s, reading, err := ReadLine(scanner.Text())
		if err != nil {
			return nil, err
		}
		// first read after ref, first sensor
		if sensor == nil {
			sensor = s
			continue
		}
		// if match means a new sensor data set
		if s != nil && s.cfg != nil {
			// apply rule quality calculation for current sensor
			result := sensor.cfg.ruleFunc(rtotal, rcounter, readings, refs[sensor._type])
			qualityCtrl[sensor.name] = result

			// reset for new sensor
			sensor = s
			readings = nil
			rtotal = 0
			rcounter = 0

			continue
		}

		// aggregate readings for current sensor data set
		readings = append(readings, *reading)
		rtotal += *reading
		rcounter++
	}
	// apply rule for last sensor data set
	result := sensor.cfg.ruleFunc(rtotal, rcounter, readings, refs[sensor._type])
	qualityCtrl[sensor.name] = result

	return qualityCtrl, nil
}

// format: reference <temperature> <humidity>
func ReadRefLine(ln string) (map[string]float64, error) {
	fields := strings.Fields(ln)
	if len(fields) != 3 {
		return nil, errors.New("Incorrect reference line format")
	}

	refs := make(map[string]float64)

	temperature, err := strconv.ParseFloat(fields[1], 64)
	if err != nil {
		return nil, err
	}
	refs[Thermometer] = temperature

	humidity, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return nil, err
	}
	refs[Humidity] = humidity

	return refs, nil
}

func ReadLine(ln string) (*Sensor, *float64, error) {
	fields := strings.Fields(ln)
	len := len(fields)
	if len != 2 && len != 3 {
		return nil, nil, errors.New("Incorrect line format")
	}
	sensorCfg := sensorCfgs[fields[0]]

	// for format: <sensorType> <sensorName>
	if sensorCfg.ruleFunc != nil {
		return &Sensor{_type: fields[0], name: fields[1], cfg: &sensorCfg}, nil, nil
	}

	// for format: <date> <sensorName> <reading>
	const layout = "2006-01-02T15:04"
	_, err := time.Parse(layout, fields[0])
	if err != nil {
		return nil, nil, err
	}
	reading, err := strconv.ParseFloat(fields[2], 64)
	if err != nil {
		return nil, nil, err
	}

	return &Sensor{name: fields[1]}, &reading, nil
}
