package main

import (
	"encoding/json"
	"fmt"
	"github.com/genghongjie/go-code/middleware/kafka/ksmonitor_producer/md_kafka"
	"github.com/labstack/gommon/log"
	"math/rand"
	"time"
)

type Status struct {
	//ItemApp string `json:"itemapp"`
	//分组标识
	ItemGroup string `json:"itemgroup"`
	//类型
	//ItemKey string `json:"itemkey"`

	//公共
	Last_flow_per_sec    int `json:"last_flow_per_sec"`     //平均流量
	Last_maxFlow_per_sec int `json:"last_max_flow_per_sec"` //峰值流量

	//DRTPRouter 时时展示值
	Last_commsvr_load int `json:"last_commsvr_load"` //负载
	Last_rqst_queues  int `json:"last_rqst_queues"`  //请求队列1
	Last_resp_queues  int `json:"last_resp_queues"`  //回复队列2
	//Last_flow_per_sec    int `json:"last_flow_per_sec"` //平均流量
	//Last_maxFlow_per_sec int `json:"last_max_flow_per_sec"` //峰值流量

	//CMDS 时时展示数据
	Last_load int `json:"last_load"` //负载
	//Last_flow_per_sec    int `json:"last_flow_per_sec"` //平均流量

	//接入层 spxA 网关
	Last_ontheway_num int `json:"last_ontheway_num"` //在途数
	//Last_flow_per_sec    int `json:"last_flow_per_sec"` //平均流量
	//Last_maxFlow_per_sec int `json:"last_max_flow_per_sec"` //峰值流量

	//BaoPan
	Last_pbu_send_num       int `json:"last_pbu_send_num"`       //委托
	Last_pbu_ack_num        int `json:"last_pbu_ack_num"`        //应答
	Last_pbu_commission_num int `json:"last_pbu_commission_num"` //回报

	//CacherServer

	LastCache    float64 `json:"last_cache"`     //cache处理数 占比
	LastNonCache float64 `json:"last_non_cache"` //非cache处理数 占比

	//KSMBCC 待提供
	Last_bu_load          int `json:"last_bu_load"`          //bu数
	Last_conf_bu_num      int `json:"last_conf_bu_num"`      //bu总数
	Last_queue_pndg_rqsts int `json:"last_queue_pndg_rqsts"` //ksmbcc的队列中待处理请求数
}

func main() {
	clientKfaka := md_kafka.KafkaEngine{}
	err := clientKfaka.GetEngine(3, md_kafka.KafakaEngineConfig{Topic: "Topic_Status", BootstrapServers: "192.168.102.57:29092"})
	if err != nil {
		log.Fatal(err)
	}

	dataarray := make([]Status, 0)

	//JQ: 金桥 WGQ：外高桥 BJ:北京 XN:成都 BD:深圳
	//同步、异步网关 Switch ok
	GroupNamesSwitch := []string{
		"JQ_JZJY_SPX_SPXA",
		"WGQ_JZJY_SPX_SPXA",
		"BJ_JZJY_SPX_SPXA",
		"XN_JZJY_SPX_SPXA",
		"BD_JZJY_SPX_SPXA",
	}
	//Cache ok
	GroupNamesCache := []string{
		"JQ_JZJY_CACHSVR_G1",
		"WGQ_JZJY_CACHSVR_G1",
		"BJ_JZJY_CACHSVR_G1",
		"XN_JZJY_CACHSVR_G1",
		"BD_JZJY_CACHSVR_G1",
	}

	//cmds ok
	GroupNamesCMDS := []string{
		"JQ_JZJY_CMDS_G1",
		"WGQ_JZJY_CMDS_G1",
		"BJ_JZJY_CMDS_G1",
		"XN_JZJY_CMDS_G1",
		"BD_JZJY_CMDS_G1",
	}

	//DRTP
	GroupNamesDRTP := []string{
		"JQ_JZJY_DRTP_G1_L1",
		"JQ_JZJY_DRTP_G2_L1",
		"JQ_JZJY_DRTP_G3_L1",
		"JQ_JZJY_DRTP_G4_L1",
		"JQ_JZJY_DRTP_G5_L1",
		"JQ_JZJY_DRTP_G6_L1",
	}
	////KSMBCC
	GroupNamesKSAPP := []string{
		"JQ_JZJY_KSMBCC_G1",
		"JQ_JZJY_KSMBCC_G2",
		"JQ_JZJY_KSMBCC_G3",
		"JQ_JZJY_KSMBCC_G4",
	}
	//
	////三方存管
	//GroupNamesBankFront := []string{
	//	"JQ_JZJY_BankFront_G1",
	//}
	////中登 上海 深圳
	//GroupNamesMIOProc := []string{
	//	"JQ_JZJY_MIOProc_G1",
	//	"BD_JZJY_MIOProc_G1",
	//}
	//报盘 ok
	GroupNamesBaoPan := []string{
		"SH_J_BAOPAN",
		"BD_J_BAOPAN",
		"BJ_J_BAOPAN",
	}
	for i := 0; i < 6; i++ {

		if len(GroupNamesSwitch) > i {
			in := Status{
				ItemGroup:               GroupNamesSwitch[i],
				Last_flow_per_sec:       rand.Intn(100000),
				Last_maxFlow_per_sec:    rand.Intn(100000),
				Last_commsvr_load:       rand.Intn(100000),
				Last_rqst_queues:        rand.Intn(100000),
				Last_resp_queues:        rand.Intn(100000),
				Last_load:               rand.Intn(100000),
				Last_ontheway_num:       rand.Intn(100000),
				Last_pbu_send_num:       rand.Intn(100000),
				Last_pbu_ack_num:        rand.Intn(100000),
				Last_pbu_commission_num: rand.Intn(100000),
				LastCache:               float64(rand.Intn(100)) / 100,
				LastNonCache:            float64(rand.Intn(100)) / 100,
			}
			dataarray = append(dataarray, in)

		}
		if len(GroupNamesCache) > i {
			inCache := Status{
				ItemGroup:               GroupNamesCache[i],
				Last_flow_per_sec:       rand.Intn(100000),
				Last_maxFlow_per_sec:    rand.Intn(100000),
				Last_commsvr_load:       rand.Intn(100000),
				Last_rqst_queues:        rand.Intn(100000),
				Last_resp_queues:        rand.Intn(100000),
				Last_load:               rand.Intn(100000),
				Last_ontheway_num:       rand.Intn(100000),
				Last_pbu_send_num:       rand.Intn(100000),
				Last_pbu_ack_num:        rand.Intn(100000),
				Last_pbu_commission_num: rand.Intn(100000),
				LastCache:               float64(rand.Intn(10000)) / 10000,
				LastNonCache:            float64(rand.Intn(10000)) / 10000,
			}
			dataarray = append(dataarray, inCache)

		}
		if len(GroupNamesCMDS) > i {
			inCmds := Status{
				ItemGroup:               GroupNamesCMDS[i],
				Last_flow_per_sec:       rand.Intn(100000),
				Last_maxFlow_per_sec:    rand.Intn(100000),
				Last_commsvr_load:       rand.Intn(100000),
				Last_rqst_queues:        rand.Intn(100000),
				Last_resp_queues:        rand.Intn(100000),
				Last_load:               rand.Intn(100000),
				Last_ontheway_num:       rand.Intn(100000),
				Last_pbu_send_num:       rand.Intn(100000),
				Last_pbu_ack_num:        rand.Intn(100000),
				Last_pbu_commission_num: rand.Intn(100000),
				LastCache:               float64(rand.Intn(100)) / 100,
				LastNonCache:            float64(rand.Intn(100)) / 100,
			}
			dataarray = append(dataarray, inCmds)

		}

		if len(GroupNamesDRTP) > i {
			inDRTP := Status{
				ItemGroup:               GroupNamesDRTP[i],
				Last_flow_per_sec:       rand.Intn(100000),
				Last_maxFlow_per_sec:    rand.Intn(100000),
				Last_commsvr_load:       rand.Intn(100000),
				Last_rqst_queues:        rand.Intn(100000),
				Last_resp_queues:        rand.Intn(100000),
				Last_load:               rand.Intn(100000),
				Last_ontheway_num:       rand.Intn(100000),
				Last_pbu_send_num:       rand.Intn(100000),
				Last_pbu_ack_num:        rand.Intn(100000),
				Last_pbu_commission_num: rand.Intn(100000),
				LastCache:               float64(rand.Intn(100)) / 100,
				LastNonCache:            float64(rand.Intn(100)) / 100,
			}
			dataarray = append(dataarray, inDRTP)

		}
		if len(GroupNamesKSAPP) > i {
			inKsmbcc := Status{
				ItemGroup:               GroupNamesKSAPP[i],
				Last_flow_per_sec:       rand.Intn(100000),
				Last_maxFlow_per_sec:    rand.Intn(100000),
				Last_commsvr_load:       rand.Intn(100000),
				Last_rqst_queues:        rand.Intn(100000),
				Last_resp_queues:        rand.Intn(100000),
				Last_load:               rand.Intn(100000),
				Last_ontheway_num:       rand.Intn(100000),
				Last_pbu_send_num:       rand.Intn(100000),
				Last_pbu_ack_num:        rand.Intn(100000),
				Last_pbu_commission_num: rand.Intn(100000),
				LastCache:               float64(rand.Intn(100)) / 100,
				LastNonCache:            float64(rand.Intn(100)) / 100,
				Last_bu_load:            rand.Intn(50000),
				Last_conf_bu_num:        50000,
				Last_queue_pndg_rqsts:   rand.Intn(100000),
			}
			dataarray = append(dataarray, inKsmbcc)
		}
		//
		//if len(GroupNamesBankFront) > i {
		//	inBankFront := Status{
		//		Last_flow_per_sec:       rand.Intn(100000),
		//		Last_maxFlow_per_sec:    rand.Intn(100000),
		//		Last_commsvr_load:       rand.Intn(100000),
		//		Last_rqst_queues:        rand.Intn(100000),
		//		Last_resp_queues:        rand.Intn(100000),
		//		Last_load:               rand.Intn(100000),
		//		Last_ontheway_num:       rand.Intn(100000),
		//		Last_pbu_send_num:       rand.Intn(100000),
		//		Last_pbu_ack_num:        rand.Intn(100000),
		//		Last_pbu_commission_num: rand.Intn(100000),
		//		LastCache:               float64(rand.Intn(100)) / 100,
		//		LastNonCache:            float64(rand.Intn(100)) / 100,
		//	}
		//	dataarray = append(dataarray, inBankFront)
		//}
		//if len(GroupNamesMIOProc) > i {
		//	inMIOProc := Status{
		//
		//		ItemGroup:               GroupNamesMIOProc[i],
		//		Last_flow_per_sec:       rand.Intn(100000),
		//		Last_maxFlow_per_sec:    rand.Intn(100000),
		//		Last_commsvr_load:       rand.Intn(100000),
		//		Last_rqst_queues:        rand.Intn(100000),
		//		Last_resp_queues:        rand.Intn(100000),
		//		Last_load:               rand.Intn(100000),
		//		Last_ontheway_num:       rand.Intn(100000),
		//		Last_pbu_send_num:       rand.Intn(100000),
		//		Last_pbu_ack_num:        rand.Intn(100000),
		//		Last_pbu_commission_num: rand.Intn(100000),
		//		LastCache:               float64(rand.Intn(100)) / 100,
		//		LastNonCache:            float64(rand.Intn(100)) / 100,
		//	}
		//	dataarray = append(dataarray, inMIOProc)
		//}
		if len(GroupNamesBaoPan) > i {
			inBaoPan := Status{
				ItemGroup:               GroupNamesBaoPan[i],
				Last_flow_per_sec:       rand.Intn(100000),
				Last_maxFlow_per_sec:    rand.Intn(100000),
				Last_commsvr_load:       rand.Intn(100000),
				Last_rqst_queues:        rand.Intn(100000),
				Last_resp_queues:        rand.Intn(100000),
				Last_load:               rand.Intn(100000),
				Last_ontheway_num:       rand.Intn(100000),
				Last_pbu_send_num:       rand.Intn(100000),
				Last_pbu_ack_num:        rand.Intn(100000),
				Last_pbu_commission_num: rand.Intn(100000),
				LastCache:               float64(rand.Intn(100)) / 100,
				LastNonCache:            float64(rand.Intn(100)) / 100,
			}
			dataarray = append(dataarray, inBaoPan)
		}

	}

	for {

		num := rand.Intn(len(dataarray))
		item := dataarray[num]
		item.Last_flow_per_sec = rand.Intn(10000)
		item.Last_maxFlow_per_sec = rand.Intn(10000)
		item.Last_commsvr_load = rand.Intn(10000)
		item.Last_rqst_queues = rand.Intn(10000)
		item.Last_resp_queues = rand.Intn(10000)
		item.Last_load = rand.Intn(10000)
		item.Last_ontheway_num = rand.Intn(10000)
		item.Last_pbu_send_num = rand.Intn(10000)
		item.Last_pbu_ack_num = rand.Intn(10000)
		item.Last_pbu_commission_num = rand.Intn(10000)
		item.LastCache = float64(rand.Intn(10000)) / 10000
		item.LastNonCache = float64(rand.Intn(10000)) / 10000
		b, _ := json.Marshal(item)
		fmt.Println(num)

		//intR := rand.Intn(500)
		//fmt.Println(string(b))
		time.Sleep(time.Duration(10) * time.Millisecond)
		clientKfaka.Produce("", b)
	}

}
