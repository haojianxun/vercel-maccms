package service

import (
	"encoding/json"
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"goapi/pkg/logger"
	"goapi/pkg/redis"
	"reflect"
)

func SaveTable(tableName, name string, value interface{}) {
	bytes, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	data := cmap.New().Items()
	data[name] = string(bytes)
	_, _ = redis.Client.HMSet(tableName, data)
}

func GetTable(tableName, name string, structName interface{}) (interface{}, error) {
	// 从 Redis 获取数据
	data, err := redis.Client.HMGet(tableName, name).Result()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	// 检查返回的数据是否为空
	if len(data) == 0 || data[0] == nil {
		return nil, fmt.Errorf("未找到键：%s，字段：%s 的数据", tableName, name)
	}

	// 将数据转换为字符串类型
	value, ok := data[0].(string)
	if !ok {
		return nil, fmt.Errorf("值的数据类型不是字符串：%T", data[0])
	}
	// 根据传入的结构体名称，动态创建反射类型
	structType := reflect.TypeOf(structName)
	// 使用反射创建切片并反序列化数据
	slicePtr := reflect.New(structType).Interface()
	if err := json.Unmarshal([]byte(value), slicePtr); err != nil {
		panic(err)
		return nil, err
	}
	return slicePtr, nil
}
