package main

import (
	"encoding/json"
	"fmt"
	"github.com/genghongjie/go-code/middleware/kafka/ksmonitor_producer/md_kafka"
	"github.com/labstack/gommon/log"
	"math/rand"
	"strconv"
	"time"
)

type Event struct {
	ItemApp string `json:"itemapp"`
	//分组标识
	ItemGroup string `json:"itemgroup"`
	//类型
	ItemType string `json:"itemtype"`
	//ID
	ItemId     string `json:"itemid"`
	Date       string `json:"date"`
	Time       string `json:"time"`
	StatusFrom string `json:"statusfrom"`
	//最新状态
	StatusTo string `json:"statusto"`
	StatusFn string `json:"statusfn"`
	StatusTn string `json:"statustn"`
}

func main() {
	clientKfaka := md_kafka.KafkaEngine{}
	err := clientKfaka.GetEngine(3, md_kafka.KafakaEngineConfig{Topic: "Topic_Event", BootstrapServers: "192.168.102.57:29092"})
	if err != nil {
		log.Fatal(err)
	}

	dataarray := make([]Event, 0)

	//JQ: 金桥 WGQ：外高桥 BJ:北京 XN:成都 BD:深圳
	//同步、异步网关 Switch ok
	GroupNamesSwitch := []string{
		////集中交易
		"JQ_JZJY_SPXA_G1", //金桥 异步网关
		"WGQ_JZJY_SPXA_G1",
		"BJ_JZJY_SPXA_G1",
		"XN_JZJY_SPXA_G1",
		"BD_JZJY_SPXA_G1",

		"JQ_JZJY_SPX_G1", //金桥 同步网关
		"WGQ_JZJY_SPX_G1",
		"BJ_JZJY_SPX_G1",
		"XN_JZJY_SPX_G1",
		"BD_JZJY_SPX_G1",

		////融资融券
		"JQ_RZRQ_SPXA_G1", //金桥 异步网关
		"WGQ_RZRQ_SPXA_G1",
		"XN_RZRQ_SPXA_G1",
		"BD_RZRQ_SPXA_G1",

		"JQ_RZRQ_SPX_G1", //金桥 同步网关
		"WGQ_RZRQ_SPX_G1",
		"XN_RZRQ_SPX_G1",
		"BD_RZRQ_SPX_G1",

		/////个股期权
		"JQ_GGQQ_SPXA_G1", //金桥 异步网关
	}
	//Cache ok
	GroupNamesCache := []string{
		////集中交易
		"JQ_JZJY_CACHSVR_G1",
		"WGQ_JZJY_CACHSVR_G1",
		"BJ_JZJY_CACHSVR_G1",
		"XN_JZJY_CACHSVR_G1",
		"BD_JZJY_CACHSVR_G1",
	}

	//cmds ok
	GroupNamesCMDS := []string{
		////集中交易
		"JQ_JZJY_CMDS_G1",
		"WGQ_JZJY_CMDS_G1",
		"BJ_JZJY_CMDS_G1",
		"XN_JZJY_CMDS_G1",
		"BD_JZJY_CMDS_G1",
	}

	//DRTP ok
	GroupNamesDRTP := []string{
		////集中交易
		"JQ_JZJY_DRTP_G1_L1",
		"JQ_JZJY_DRTP_G2_L1",
		"JQ_JZJY_DRTP_G3_L1",
		"JQ_JZJY_DRTP_G4_L1",
		"JQ_JZJY_DRTP_G5_L1",

		////融资融券
		"JQ_RZRQ_DRTP_G1_L1",

		/////个股期权
		"JQ_GGQQ_DRTP4_G1_L1", // 一级通讯平台
		"JQ_GGQQ_DRTP4_G2_L2", // 二级通讯平台

	}
	//KSMBCC ok
	GroupNamesKSAPP := []string{
		////集中交易
		"JQ_JZJY_KSMBCC_G1",
		"JQ_JZJY_KSMBCC_G2",
		"JQ_JZJY_KSMBCC_G3",
		"JQ_JZJY_KSMBCC_G4",
		"JQ_JZJY_KSMBCC_G5",

		////融资融券
		"JQ_RZRQ_KSMBCC_G1",
		"JQ_RZRQ_KSMBCC_G2",

		/////个股期权
		"JQ_GGQQ_KSMBCC_G1", // 主核心
		"JQ_GGQQ_KSMBCC_G2", // 备核心
	}

	//三方存管
	GroupNamesBankFront := []string{
		////集中交易
		"JQ_JZJY_BANK_G1",
		"JQ_JZJY_BANKBCC_G1",

		////融资融券
		"JQ_RZRQ_BANK_G1",
		"JQ_RZRQ_BANKBCC_G1",

		////个股期权
		"JQ_GGQQ_BANK_G1",
		"JQ_GGQQ_BANKBCC_G1",
	}
	//中登 上海 深圳
	GroupNamesMIOProc := []string{
		////集中交易
		"JQ_JZJY_MIOProc_G1",
		"BD_JZJY_MIOProc_G1",
	}
	//报盘 ok
	GroupNamesBaoPan := []string{
		////集中交易
		//上交所
		"SHA_J",            //交易所状态
		"SHB_J",            //交易所状态
		"JQ_JZJY_G4BCC_G1", //金仕达状态

		//深交所
		"SZA_J",            //交易所状态
		"SZB_J",            //交易所状态
		"JQ_JZJY_V5BCC_G1", //金仕达状态

		//北交所
		"BJA_J",

		////融资融券
		//上交所
		"SHA_R",            //交易所状态
		"SHB_R",            //交易所状态
		"JQ_RZRQ_G4BCC_G1", //金仕达状态

		//深交所
		"SZA_R",            //交易所状态
		"SZB_R",            //交易所状态
		"JQ_RZRQ_V5BCC_G1", //金仕达状态

		//北交所
		"BJA_R",

		////个股期权

		//上交所
		"SHA_G",            //上交所 交易所状态
		"JQ_GGQQ_G4BCC_G1", //上交所 金仕达状态

		//深交所
		"SZA_G",            //深交所 交易所状态
		"JQ_GGQQ_V5BCC_G1", //上交所 金仕达状态

	}

	//其它
	//报盘 ok
	GroupNamesOther := []string{

		////个股期权
		//其他组件一
		"JQ_GGQQ_QQZJ_G1",
		//其他组件二
		"JQ_GGQQ_QQZJ_G2",
	}
	for i := 0; i < 20; i++ {

		intR := rand.Intn(200)
		for j := 0; j < 3; j++ {
			alert := "正常"
			switch j {
			case 0:
				alert = "正常"
			case 1:
				alert = "报警"
			case 2:
				alert = "报警"

			}
			if len(GroupNamesSwitch) > i {
				in := Event{
					ItemApp:    "J",
					ItemGroup:  GroupNamesSwitch[i],
					ItemType:   "Switch",
					ItemId:     "10.0.0." + strconv.Itoa(intR),
					Date:       "20221224",
					Time:       "11:49:00",
					StatusFrom: "0",
					StatusTo:   strconv.Itoa(j),
					StatusFn:   "正常",
					StatusTn:   alert,
				}
				dataarray = append(dataarray, in)

			}
			if len(GroupNamesCache) > i {
				inCache := Event{
					ItemApp:   "J",
					ItemGroup: GroupNamesCache[i],
					ItemType:  "Switch",
					ItemId:    "CacheServer_10.0.0." + strconv.Itoa(intR),
					Date:      "20221224",
					Time:      "11:49:00",

					StatusFrom: "0",
					StatusTo:   strconv.Itoa(j),
					StatusFn:   "正常",
					StatusTn:   alert,
				}
				dataarray = append(dataarray, inCache)

			}
			if len(GroupNamesCMDS) > i {
				inCmds := Event{
					ItemApp:   "J",
					ItemGroup: GroupNamesCMDS[i],
					ItemType:  "CMDS",
					ItemId:    "CMDS_10.0.0." + strconv.Itoa(intR),
					Date:      "20221224",
					Time:      "11:49:00",

					StatusFrom: "0",
					StatusTo:   strconv.Itoa(j),
					StatusFn:   "正常",
					StatusTn:   alert,
				}
				dataarray = append(dataarray, inCmds)

			}

			if len(GroupNamesDRTP) > i {
				inDRTP := Event{
					ItemApp:    "J",
					ItemGroup:  GroupNamesDRTP[i],
					ItemType:   "DRTP",
					ItemId:     "DRTP_10.0.0." + strconv.Itoa(intR),
					Date:       "20221224",
					Time:       "11:49:00",
					StatusFrom: "0",
					StatusTo:   strconv.Itoa(j),
					StatusFn:   "正常",
					StatusTn:   alert,
				}
				dataarray = append(dataarray, inDRTP)

			}
			if len(GroupNamesKSAPP) > i {
				inKsmbcc := Event{
					ItemApp:    "J",
					ItemGroup:  GroupNamesKSAPP[i],
					ItemType:   "KSMBCC",
					ItemId:     "ksmbcc.0.0." + strconv.Itoa(intR),
					Date:       "20221224",
					Time:       "11:49:00",
					StatusFrom: "0",
					StatusTo:   strconv.Itoa(j),
					StatusFn:   "正常",
					StatusTn:   alert,
				}
				dataarray = append(dataarray, inKsmbcc)
			}

			if len(GroupNamesBankFront) > i {
				inBankFront := Event{
					ItemApp:    "J",
					ItemGroup:  GroupNamesBankFront[i],
					ItemType:   "BankFront",
					ItemId:     "BankFront.0.0." + strconv.Itoa(intR),
					Date:       "20221224",
					Time:       "11:49:00",
					StatusFrom: "0",
					StatusTo:   strconv.Itoa(j),
					StatusFn:   "正常",
					StatusTn:   alert,
				}
				dataarray = append(dataarray, inBankFront)
			}
			if len(GroupNamesMIOProc) > i {
				inMIOProc := Event{
					ItemApp:    "J",
					ItemGroup:  GroupNamesMIOProc[i],
					ItemType:   "MIOProc",
					ItemId:     "MIOProc.0.0." + strconv.Itoa(intR),
					Date:       "20221224",
					Time:       "11:49:00",
					StatusFrom: "0",
					StatusTo:   strconv.Itoa(j),
					StatusFn:   "正常",
					StatusTn:   alert,
				}
				dataarray = append(dataarray, inMIOProc)
			}
			if len(GroupNamesBaoPan) > i {
				inBaoPan := Event{
					ItemApp:    "J",
					ItemGroup:  GroupNamesBaoPan[i],
					ItemType:   "BaoPan",
					ItemId:     "BaoPan.0.0." + strconv.Itoa(intR),
					Date:       "20221224",
					Time:       "11:49:00",
					StatusFrom: "0",
					StatusTo:   strconv.Itoa(j),
					StatusFn:   "正常",
					StatusTn:   alert,
				}
				dataarray = append(dataarray, inBaoPan)
			}
			if len(GroupNamesOther) > i {
				inBaoPan := Event{
					ItemApp:    "J",
					ItemGroup:  GroupNamesOther[i],
					ItemType:   "Other",
					ItemId:     "Other.0.0." + strconv.Itoa(intR),
					Date:       "20221224",
					Time:       "11:49:00",
					StatusFrom: "0",
					StatusTo:   strconv.Itoa(j),
					StatusFn:   "正常",
					StatusTn:   alert,
				}
				dataarray = append(dataarray, inBaoPan)
			}

		}

	}

	for {

		num := rand.Intn(len(dataarray))
		e := dataarray[num]
		e.Date = time.Now().Format("2006-01-02")
		e.Time = time.Now().Format("15:04:05")
		b, _ := json.Marshal(e)
		fmt.Println(e.ItemId)

		//intR := rand.Intn(5000)
		//fmt.Println(string(b))
		time.Sleep(time.Duration(100) * time.Millisecond)
		//time.Sleep(time.Duration(30) * time.Second)
		clientKfaka.Produce("", b)
	}

}
