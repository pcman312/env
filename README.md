# env
[![Build Status](https://travis-ci.org/pcman312/env.svg?branch=master)](https://travis-ci.org/pcman312/env)
[![Coverage Status](https://coveralls.io/repos/github/pcman312/env/badge.svg?branch=travis)](https://coveralls.io/github/pcman312/env?branch=travis)

# Why should I use this library?
Are you using your environment for config values in your Go program? Are you tired of manually parsing each of those fields into structs? Then this library is for you! It can take your environment and parse it into a struct that you give it. Think of this like the built in JSON parser, except parsing from the environment.

I created this because I was unsatisfied with other environment parsing libraries out there and wanted to remove a lot of boilerplate code in my config parsing. After implementing this, the size of my parse function went from 100 lines to 27 lines and the tests went from 873 lines to 206 lines. The only thing remaining in that function is some business-specific logic around a few of the fields which can't easily be encapsulated in a generic library. I no longer have to worry about whether a given field is parsed correctly or if it is within custom min/max boundaries.

# How do I use this?
1. Create a struct that you want to parse from the environment. Each of the fields in the struct you want to have parsed need to have the `env` struct tag set on it.
Example:
```go
type Config struct {
  StringField string `env:"strfield"`
  IntField int       `env:"intfield"`
}
```
2. Create a pointer to an instance of that object and pass it to the `env.Parse` function.
Example:
```go
c := &Config{}
err := env.Parse(c)
if err != nil {
  fmt.Printf("Error parsing config: %s\n", err)
}
```
3. ???
4. Profit

# What fields does it support?
It currently supports these fields:
- bool
- string
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- float32
- float64
- time.Duration
- *url.URL

It also supports slices of each of these types:
- []bool
- []string
- []int
- []int8
- []int16
- []int32
- []int64
- []uint
- []uint8
- []uint16
- []uint32
- []uint64
- []float32
- []float64
- []time.Duration
- []*url.URL

# What struct tags are available?
- `env` - the name of the environment variable to parse
- `default` - the default value of the environment variable if it's not found
- `required` - is the field required? Must be either "true" or "false" or it will error
- `min` - minimum allowed value in the field. Only applies to numeric fields. Other fields will ignore this tag
- `max` - maximum allowed value in the field. Only applies to numeric fields. Other fields will ignore this tag

**Note:** `min` and `max` are both inclusive. For instance, if you specify `min:"5" max:"10"` the values of `5` and `10` will be allowed, but `4` and `11` will not.

# Citations
This is heavily influenced by https://github.com/caarlos0/env and can be thought of as a fork and expansion on that library, however this does not match exactly 1:1 with that library. I decided against maintaining a direct fork for two big reasons: 1) I intended to make some significant structural changes and additions that were not going to be pulled into his main library and 2) Maintaining a fork in github has its own set of problems making maintaining it more difficult.
