package container

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"sync"
)

var (
	app = new(BaseContainer)
)

type Container interface {
	ParseConfig(component Component) (interface{}, error)
}

type BaseContainer struct {
	CompMap sync.Map
}

func (dc *BaseContainer) ParseConfig(comp Component) (interface{}, error) {
	var (
		compConfigValue interface{}
		err             error
	)
	// 通过配置读取参数
	compConfigKeys := strings.Split(comp.CfgKey(), ".")
	// todo 加载配置文件
	fmt.Println(compConfigKeys)
	var compConfig interface{}

	compConfigValue = reflect.New(reflect.TypeOf(comp.CfgType())).Interface()
	if err := mapstructure.WeakDecode(compConfig, compConfigValue); nil != err {
		err = errors.Wrap(err, "ParseConfig decode err")
	}
	return compConfigValue, err
}
