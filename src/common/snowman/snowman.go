// Package snowman 雪花ID生成中心
package snowman

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"sync"
)

var (
	api  *Snow
	once sync.Once
)

// NewSnowApi 雪花服务接口
func NewSnowApi() *Snow {
	once.Do(func() {
		// 自定义雪花参数
		snowflake.Epoch = 1663862400000 // 应用纪元时间戳（2022-09-23 00:00:00）
		snowflake.NodeBits = 10
		snowflake.StepBits = 8

		api = &Snow{nodeIdx: 1}
		api.initNode()
		fmt.Println(fmt.Sprintf("服务启动: 雪花服务初始化成功, 节点: %d", api.nodeIdx))
	})
	return api
}

type Snow struct {
	nodeIdx    int64
	nodeClient *snowflake.Node
}

// GetIdInt64 获取int64类型的id
func (sm *Snow) GetIdInt64() int64 {
	return sm.nodeClient.Generate().Int64()
}

// GetIdStr 获取string类型的id
func (sm *Snow) GetIdStr() string {
	return sm.nodeClient.Generate().String()
}

// initNode 初始化节点, 当前固定节点索引, 不提供参数配置
func (sm *Snow) initNode() {
	tmpNodeClient, err := snowflake.NewNode(sm.nodeIdx)
	if err != nil {
		fmt.Println(fmt.Sprintf("雪花节点[%d]初始化失败, 原因: %s", sm.nodeIdx, err.Error()))
		panic(err)
	}
	sm.nodeClient = tmpNodeClient
}
