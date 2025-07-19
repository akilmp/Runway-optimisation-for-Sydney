package perf

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

type Config struct {
	Runway string
	Temp   float64
	Alt    int
	Weight int
	Wind   int
	Slope  float64
}

type Result struct {
	V1          int
	Vr          int
	V2          int
	FieldLength int
	FuelDelta   float64
}

func CalcTakeoff(cfg Config) (Result, error) {
	f, err := os.Open("internal/perf/data/737-800-takeoff.csv")
	if err != nil {
		return Result{}, err
	}
	defer f.Close()
	r := csv.NewReader(f)
	_, _ = r.Read() // skip header
	for {
		rec, err := r.Read()
		if err != nil {
			break
		}
		rwy := rec[0]
		if rwy != cfg.Runway {
			continue
		}
		temp, _ := strconv.Atoi(rec[1])
		if temp != int(cfg.Temp) {
			continue
		}
		v1, _ := strconv.Atoi(rec[2])
		vr, _ := strconv.Atoi(rec[3])
		v2, _ := strconv.Atoi(rec[4])
		field, _ := strconv.Atoi(rec[5])
		fuel, _ := strconv.ParseFloat(rec[6], 64)
		return Result{v1, vr, v2, field, fuel}, nil
	}
	return Result{}, errors.New("no match")
}
