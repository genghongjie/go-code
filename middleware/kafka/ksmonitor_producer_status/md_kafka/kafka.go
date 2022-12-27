package md_kafka

import (
	"github.com/labstack/gommon/log"
	"sync"

	"github.com/Shopify/sarama"
)

const (
	KAFKA_TYPE = "kafka"
)

var kafkaCache = make(map[int]*KafkaEngine)

type KafkaEngine struct {
	Id     int //唯一标志
	Client *sarama.SyncProducer
	Config *KafakaEngineConfig // 连接信息

	Out   *map[string]interface{}
	syncR *sync.RWMutex
	Error error
}

type KafkaData struct {
	TypeId int //类型ID 对应本项目中 元数据ID
	Data   map[string]interface{}
}

var KafkaOutDataCh = make(chan KafkaData, 100)

//配置参数参考官方	https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md
type KafakaEngineConfig struct {
	//Alias for metadata.broker.list: Initial list of brokers as a CSV list of broker host or host:port. The application may also use rd_kafka_brokers_add() to add brokers during runtime.
	//Type: string
	BootstrapServers string `json:"bootstrap_servers"` //localhost:29092
	//Client group id string. All clients sharing the same group.id belong to the same group.
	//Type: string
	//GroupId string `json:"group_id"`
	//smallest, earliest, beginning, largest, latest, end, error
	//AutoOffetReset string `json:"auto_offet_reset"` //earliest
	Topic   string `json:"topic"` //topic 多个使用逗号分割
	RunMode string `json:"run_mode"`
}

func (config *KafakaEngineConfig) Compare(otherConfig KafakaEngineConfig) bool {
	if config.BootstrapServers != otherConfig.BootstrapServers {
		return false
	}
	//if config.GroupId != otherConfig.GroupId {
	//	return false
	//}
	//if config.AutoOffetReset != otherConfig.AutoOffetReset {
	//	return false
	//}
	if config.Topic != otherConfig.Topic {
		return false
	}
	return true
}

//1 初始化数据源连接
func (engine *KafkaEngine) GetEngine(id int, config KafakaEngineConfig) error {
	engine.Config = &config
	//if engine.Config.GroupId == "" {
	//	engine.Config.GroupId = "showcase-exporter"
	//}
	//if engine.Config.AutoOffetReset == "" {
	//	engine.Config.AutoOffetReset = "earliest"
	//}
	engine.syncR = new(sync.RWMutex)

	if v, ok := kafkaCache[id]; ok {
		//比对新旧config
		//一致 直接返回，不一致关闭连接重新创建
		if v.Config.Compare(config) {

			engine.Client = v.Client
			engine.Out = v.Out
			engine.syncR = v.syncR
			return v.Error
		}
		log.Infof("%s 数据源配置信息变更，重新创建连接....  Url:%s,Id:%d", KAFKA_TYPE, engine.Config.BootstrapServers, engine.Id)
		//否 关闭连接 重新init并返回
		v.Close()
	}

	if engine.Client == nil {
		engine.open()
	}
	if engine.Error != nil {
		return engine.Error
	}
	kafkaCache[id] = engine
	return nil
}

//3 关闭连接 优雅退出时调用
func (engine *KafkaEngine) Close() {

	kc := *engine.Client
	_ = kc.Close()
	//移除
	delete(kafkaCache, engine.Id)
}

func (engine *KafkaEngine) open() {
	engine.syncR = &sync.RWMutex{}
	engineConfig := *engine.Config

	config := sarama.NewConfig()

	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{engineConfig.BootstrapServers}, config)

	if err != nil {
		engine.Error = err
		log.Errorf("%s-创建连接失败 Host[%s] 检查连接状态异常。%s", KAFKA_TYPE, engine.Config.BootstrapServers, err)
		return
	}

	log.Infof("%s- Host[%s] 成功创建连接。", KAFKA_TYPE, engine.Config.BootstrapServers)
	engine.Client = &client
}

func (engine *KafkaEngine) reConnect() {
	engine.Error = nil
	engine.open()
}

func (engine *KafkaEngine) Produce(key string, data []byte) {
	//可选的chan Event，可以用来监听发送的结果
	msg := &sarama.ProducerMessage{}
	//msg.Key = sarama.StringEncoder(key)
	msg.Topic = engine.Config.Topic
	msg.Value = sarama.StringEncoder(data)

	//发送消息
	kafkaClient := *engine.Client
	pid, offset, err := kafkaClient.SendMessage(msg)

	if err != nil {
		log.Errorf("kafka produce error,%s", err)
		return
	}
	//监听结果
	log.Infof("kafka produce pid:%v offset:%v\n,", pid, offset)
}
