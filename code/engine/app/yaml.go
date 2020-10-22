package app

import (
    "errors"
    "io"
    "io/ioutil"
    "os"

    "gopkg.in/yaml.v2"
)

// YamlConfig .
type YamlConfig struct {
    data map[interface{}]interface{}
}

// NewYamlConfig .
func NewYamlConfig(r io.Reader) (*YamlConfig, error) {
    d, err := ioutil.ReadAll(r)
    if err != nil {
        return nil, err
    }
    data := make(map[interface{}]interface{})
    err = yaml.Unmarshal(d, &data)
    if err != nil {
        return nil, err
    }
    return &YamlConfig{data: data}, nil
}

// NewYamlFromFile .
func NewYamlFromFile(filename string) (*YamlConfig, error) {
    fd, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer fd.Close()
    return NewYamlConfig(fd)
}

// GetConfigItem .
func GetConfigItem(yc *YamlConfig, item string) *YamlConfig {
    if v, ok := yc.data[item]; ok {
        if d, ok := v.(map[interface{}]interface{}); ok {
            return &YamlConfig{data: d}
        }
        return nil
    }
    return nil
}

// Int .
func (yc *YamlConfig) Int(key string) (int, error) {
    if v, ok := yc.data[key].(int); ok {
        return v, nil
    }
    return 0, errors.New("not int value")
}

// Int64 .
func (yc *YamlConfig) Int64(key string) (int64, error) {
    if v, ok := yc.data[key].(int64); ok {
        return v, nil
    }
    if v, ok := yc.data[key].(int); ok {
        return int64(v), nil
    }
    return 0, errors.New("not int64 value")
}

// String .
func (yc *YamlConfig) String(key string) (string, error) {
    if v, ok := yc.data[key].(string); ok {
        return v, nil
    }
    return "", errors.New("not string value")
}

// Bool .
func (yc *YamlConfig) Bool(key string) (bool, error) {
    if v, ok := yc.data[key].(bool); ok {
        return v, nil
    }
    return false, errors.New("not bool value")
}

// Float .
func (yc *YamlConfig) Float(key string) (float64, error) {
    if v, ok := yc.data[key].(float64); ok {
        return v, nil
    }
    return 0.0, errors.New("not float64 value")
}

// Strings .
func (yc *YamlConfig) Strings(key string) []string {
    var rst []string
    if v, ok := yc.data[key].([]interface{}); ok {
        for _, i := range v {
            if s, ok := i.(string); ok {
                rst = append(rst, s)
            } else {
                return rst[0:0:0]
            }
        }
        return rst
    }
    return rst[0:0:0]
}

// Diy .
func (yc *YamlConfig) Diy(key string) (interface{}, error) {
    if v, ok := yc.data[key]; ok {
        return v, nil
    }
    return nil, errors.New("key not exist")
}

// DefaultInt .
func (yc *YamlConfig) DefaultInt(key string, v int) int {
    val, err := yc.Int(key)
    if err != nil {
        return v
    }
    return val
}

// DefaultString .
func (yc *YamlConfig) DefaultString(key string, v string) string {
    val, err := yc.String(key)
    if err != nil {
        return v
    }
    return val
}

// DefaultInt64 .
func (yc *YamlConfig) DefaultInt64(key string, v int64) int64 {
    val, err := yc.Int64(key)
    if err != nil {
        return v
    }
    return val
}

// DefaultBool .
func (yc *YamlConfig) DefaultBool(key string, v bool) bool {
    val, err := yc.Bool(key)
    if err != nil {
        return v
    }
    return val
}

// DefaultFloat .
func (yc *YamlConfig) DefaultFloat(key string, v float64) float64 {
    val, err := yc.Float(key)
    if err != nil {
        return v
    }
    return val
}

// DefaultStrings .
func (yc *YamlConfig) DefaultStrings(key string, v []string) []string {
    val := yc.Strings(key)
    if val == nil {
        return v
    }
    return val
}