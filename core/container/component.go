package container

import (
	"github.com/pkg/errors"
	"reflect"
)

const (
	useNotIsPointer      = "Use(Component) not is pointer"
	useParseConfigErr    = "Use(Component) Parse config Err, compKey: %v, err: %v"
	useHavingRegisterErr = "Use(Component) Having Register err, compKey: %v"
	useInitErr           = "Use(Component) Init err, compKey: %v, err: %v"
)

// 组件标准
type Component interface {
	Init(...interface{}) error
	CfgKey() string
	CfgType() interface{}
	CfgUpdate(interface{})
}

// 组件加载
func Use(c Component, options ...interface{}) error {
	comp := reflect.TypeOf(c)
	if comp.Kind() == reflect.Ptr {
		return errors.New(useNotIsPointer)
	}
	compKey := c.CfgKey()
	compType := c.CfgType()

	// 	无组件配置类型，不需要配置, 有组件配置类型，先检查传入参数，然后再从配置文件中读取
	if nil != compType {
		var isHaveConfig bool
		for _, option := range options {
			// 配置通过参数传进来
			if reflect.TypeOf(option) == reflect.TypeOf(compType) {
				isHaveConfig = true
				break
			}
		}
		if !isHaveConfig {
			// 通过配置读取参数
			compConfigValue, err := app.ParseConfig(c)
			if nil != err {
				return errors.Wrapf(err, useParseConfigErr, compKey, err)
			}
			options = append(append([]interface{}{}, compConfigValue), options...)
		}
	}

	// 检查组建时候加载
	if _, ok := app.CompMap.LoadOrStore(compKey, c); ok {
		return errors.Errorf(useHavingRegisterErr, compKey)
	}

	// 初始化组建
	if err := c.Init(options...); err != nil {
		return errors.Wrapf(err, useInitErr, compKey, err)
	}
	return nil
}

func IsUsed(c string) bool {
	_, ok := app.CompMap.Load(c)
	return ok
}

func UsedList() []string {
	var r []string
	app.CompMap.Range(func(k, v interface{}) bool {
		ks := k.(string)
		r = append(r, ks)
		return true
	})
	return r
}
