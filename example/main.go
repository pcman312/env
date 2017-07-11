package main

import (
	"fmt"
	"net/url"
	"time"

	"encoding/json"

	"github.com/pcman312/env"
)

type SampleConfig struct {
	MyBool bool `env:"mybool"`

	MyString string `env:"mystring"`

	MyInt      int           `env:"myint" min:"0" max:"128"`
	MyInt8     int8          `env:"myint8"`
	MyInt16    int16         `env:"myint16"`
	MyInt32    int32         `env:"myint32"`
	MyInt64    int64         `env:"myint64"`
	MyDuration time.Duration `env:"myduration" min:"1m" max:"1h" default:"1h"`

	MyUint   uint   `env:"myuint" min:"128" max:"256"`
	MyUint8  uint8  `env:"myuint8"`
	MyUint16 uint16 `env:"myuint16"`
	MyUint32 uint32 `env:"myuint32"`
	MyUint64 uint64 `env:"myuint64"`

	MyFloat32 float32 `env:"myfloat32" min:"5.01" max:"6.0"`
	MyFloat64 float64 `env:"myfloat64"`

	MyURL *url.URL `env:"myurl"`

	////////////////////////////////////
	// Slices
	MyBoolArr []bool   `env:"myboolarr"`
	MyStrArr  []string `env:"mystrarr"`

	MyIntArr      []int           `env:"myintarr" min:"5" max:"10"`
	MyInt8Arr     []int8          `env:"myint8arr"`
	MyInt16Arr    []int16         `env:"myint16arr"`
	MyInt32Arr    []int32         `env:"myint32arr" delimiter:" " default:"1 2 3 4 5"`
	MyInt64Arr    []int64         `env:"myint64arr"`
	MyDurationArr []time.Duration `env:"mydurationarr" default:"5m, 1h5m"`

	MyUintArr   []uint   `env:"myuintarr"`
	MyUInt8Arr  []uint8  `env:"myuint8arr"`
	MyUInt16Arr []uint16 `env:"myuint16arr"`
	MyUInt32Arr []uint32 `env:"myuint32arr"`
	MyUInt64Arr []uint64 `env:"myuint64arr"`

	MyFloat32Arr []float32 `env:"myfloat32arr" min:"5" max:"10"`
	MyFloat64Arr []float64 `env:"myfloat64arr"`

	MyURLArr []*url.URL `env:"myurlarr" delimiter:" " default:"http://www.google.com http://www.reddit.com"`

	MySkippedStr      string `env:""`
	MyOtherSkippedStr string `env:"-"`
}

func main() {
	s := &SampleConfig{}

	err := env.Parse(s)
	output, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("%s\n", output)
	fmt.Printf("Err: %s\n", err)
}
