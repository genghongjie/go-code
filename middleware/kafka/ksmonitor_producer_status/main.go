package main

import (
	"encoding/json"
	"fmt"
	"github.com/genghongjie/go-code/middleware/kafka/ksmonitor_producer_status/md_kafka"
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
		"JZJY_JQ_SPX",
		"JZJY_WGQ_SPX",
		"JZJY_BJ_SPX",
		"JZJY_XN_SPX",
		"JZJY_BD_SPX",
	}
	//Cache ok
	GroupNamesCache := []string{
		"JZJY_JQ_CACHSERVER",
		"JZJY_WGQ_CACHSERVER",
		"JZJY_BJ_CACHSERVER",
		"JZJY_XN_CACHSERVER",
		"JZJY_BD_CACHSERVER",
	}

	//cmds
	GroupNamesCMDS := []string{
		"JZJY_SH_CMDS",
		"JZJY_WGQ_CMDS",
		"JZJY_BJ_CMDS",
		"JZJY_XN_CMDS",
		"JZJY_BD_CMDS",
	}

	//DRTP
	GroupNamesDRTP := []string{
		"JQ_JZJY_DRTP_G1_L1",
		"JQ_JZJY_DRTP_G2_L1",
		"JQ_JZJY_DRTP_G3_L1",
		"JQ_JZJY_DRTP_G4_L1",
		"JQ_JZJY_DRTP_G5_L1",
	}
	//KSMBCC
	GroupNamesKSAPP := []string{
		"JQ_JZJY_KSMBCC_G1",
		"JQ_JZJY_KSMBCC_G2",
		"JQ_JZJY_KSMBCC_G3",
		"JQ_JZJY_KSMBCC_G4",
		"JQ_JZJY_KSMBCC_G5",
	}

	//三方存管
	GroupNamesBankFront := []string{
		"JQ_JZJY_BankFront_G1",
	}
	//中登 上海 深圳
	GroupNamesMIOProc := []string{
		"JQ_JZJY_MIOProc_G1",
		"BD_JZJY_MIOProc_G1",
	}
	//报盘 ok
	GroupNamesBaoPan := []string{
		"JZJY_SH_BP",
		"JZJY_BD_BP",
		"JZJY_BJ_BP",
	}
	for i := 0; i < 5; i++ {

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
			in := Event{
				ItemApp:   "J",
				ItemGroup: GroupNamesSwitch[i],
				ItemType:  "Switch",
				ItemId:    "10.0.0." + strconv.Itoa(intR),
				Date:      "20221224",
				Time:      "11:49:00",

				StatusTo: strconv.Itoa(j),

				StatusTn: alert,
			}
			inCache := Event{
				ItemApp:   "J",
				ItemGroup: GroupNamesCache[i],
				ItemType:  "Switch",
				ItemId:    "CacheServer_10.0.0." + strconv.Itoa(intR),
				Date:      "20221224",
				Time:      "11:49:00",

				StatusTo: strconv.Itoa(j),

				StatusTn: alert,
			}
			inCmds := Event{
				ItemApp:   "J",
				ItemGroup: GroupNamesCMDS[i],
				ItemType:  "CMDS",
				ItemId:    "CMDS_10.0.0." + strconv.Itoa(intR),
				Date:      "20221224",
				Time:      "11:49:00",

				StatusTo: strconv.Itoa(j),

				StatusTn: alert,
			}
			inDRTP := Event{
				ItemApp:   "J",
				ItemGroup: GroupNamesDRTP[i],
				ItemType:  "DRTP",
				ItemId:    "DRTP_10.0.0." + strconv.Itoa(intR),
				Date:      "20221224",
				Time:      "11:49:00",

				StatusTo: strconv.Itoa(j),

				StatusTn: alert,
			}
			inKsmbcc := Event{
				ItemApp:   "J",
				ItemGroup: GroupNamesKSAPP[i],
				ItemType:  "KSMBCC",
				ItemId:    "ksmbcc.0.0." + strconv.Itoa(intR),
				Date:      "20221224",
				Time:      "11:49:00",

				StatusTo: strconv.Itoa(j),

				StatusTn: alert,
			}

			if len(GroupNamesBankFront) > i {
				inBankFront := Event{
					ItemApp:   "J",
					ItemGroup: GroupNamesBankFront[i],
					ItemType:  "BankFront",
					ItemId:    "BankFront.0.0." + strconv.Itoa(intR),
					Date:      "20221224",
					Time:      "11:49:00",

					StatusTo: strconv.Itoa(j),

					StatusTn: alert,
				}
				dataarray = append(dataarray, inBankFront)
			}
			if len(GroupNamesMIOProc) > i {
				inMIOProc := Event{
					ItemApp:   "J",
					ItemGroup: GroupNamesMIOProc[i],
					ItemType:  "MIOProc",
					ItemId:    "MIOProc.0.0." + strconv.Itoa(intR),
					Date:      "20221224",
					Time:      "11:49:00",

					StatusTo: strconv.Itoa(j),

					StatusTn: alert,
				}
				dataarray = append(dataarray, inMIOProc)
			}
			if len(GroupNamesBaoPan) > i {
				inBaoPan := Event{
					ItemApp:   "J",
					ItemGroup: GroupNamesBaoPan[i],
					ItemType:  "BaoPan",
					ItemId:    "BaoPan.0.0." + strconv.Itoa(intR),
					Date:      "20221224",
					Time:      "11:49:00",

					StatusTo: strconv.Itoa(j),

					StatusTn: alert,
				}
				dataarray = append(dataarray, inBaoPan)
			}

			dataarray = append(dataarray, in)
			dataarray = append(dataarray, inCache)
			dataarray = append(dataarray, inCmds)
			dataarray = append(dataarray, inDRTP)
			dataarray = append(dataarray, inKsmbcc)
		}

	}

	for {

		num := rand.Intn(len(dataarray))
		b, _ := json.Marshal(dataarray[num])
		fmt.Println(num)

		intR := rand.Intn(3000)
		fmt.Println(string(b))
		time.Sleep(time.Duration(intR) * time.Millisecond)
		clientKfaka.Produce("", b)
	}

}
