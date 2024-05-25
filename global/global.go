package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/ohwin/core/config"
	"github.com/ohwin/core/tools"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	Redis   *RDB
	Config  config.ServerConfig
	Mq      tools.MQ
	Enforce *casbin.SyncedCachedEnforcer
)

const (
	MQ_NORMAL_EXCHANGE = "mq_normal_exchange" // 普通交换机
	MQ_NORMAL_QUEUE    = "mq_normal_queue"    // 普通队列
	MQ_DLX_EXCHANGE    = "mq_dlx_exchange"    // 死信交换机
	MQ_DLX_QUEUE       = "mq_dlx_queue"       // 死信队列
)
