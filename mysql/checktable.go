package wymysql

import (
	"fmt"
	"github.com/aococo777ltcommon/commonstruct"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

/*
动态创建的表:
wy_user_navi_
每500个用户一张表，且永久有效

wy_user_portclassset_
每200个用户一张表，且永久有效

wy_tmp_lotteryorder
主表保留最近两个月的,备用表按月分表

wy_tmp_user_lotdata_data
主表保留最近两个月的,备用表按月分表

wy_tmp_user_lotdata_item
主表保留最近两个月的,备用表按月分表

*/

func GetUsernaviTablename(uuid int64) string {
	numtblname := fmt.Sprintf("%v%v", commonstruct.WY_user_navi_, uuid/500)
	return numtblname
}

func CheckUsernaviTablename(numtblname string) error {
	numsql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v (
							  uuid bigint NOT NULL,
							  navi_id bigint NOT NULL DEFAULT '0',
							  group_id bigint DEFAULT '0',
							  value bigint DEFAULT '0',
							  PRIMARY KEY (uuid,navi_id)
							) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;`, numtblname)

	err := WyMysql.Exec(numsql).Error
	if err != nil {
		beego.Error("CreateTable num err ", numtblname, err)
	}
	return err
}

func GetUserPortclasssetTablename(uuid int64) string {
	numtblname := fmt.Sprintf("wy_user_portclassset_%v", uuid/200)
	return numtblname
}

func CheckUserportclasssetTable(numtblname string) error {

	numsql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v (
		  uuid bigint NOT NULL,
		  room_id bigint NOT NULL DEFAULT '0',
		  port_id bigint NOT NULL DEFAULT '0',
		  odds_a decimal(20,4) DEFAULT '0',
		  tuishui_a decimal(20,4) DEFAULT '0',
		  odds_b decimal(20,4) DEFAULT '0',
		  tuishui_b decimal(20,4) DEFAULT '0',
		  odds_c decimal(20,4) DEFAULT '0',
		  tuishui_c decimal(20,4) DEFAULT '0',
		  odds_d decimal(20,4) DEFAULT '0',
		  tuishui_d decimal(20,4) DEFAULT '0',
		  default_odds decimal(20,4) DEFAULT '0',
		  min_amount bigint DEFAULT '0',
		  max_amount bigint DEFAULT '0',
		  max_amount_expect bigint DEFAULT '0',
		  transfer_amount bigint DEFAULT '0',
		  is_transfer bigint DEFAULT '0',
		  warning_amount bigint DEFAULT '0',
		  warning_loopamount bigint DEFAULT '0',
		  port_switch bigint DEFAULT '0',
		  des_odds decimal(20,4) DEFAULT '0',
		  leiji_amount bigint DEFAULT '0',
		  min_odds decimal(20,4) DEFAULT '0',
		  des_type varchar(20) DEFAULT NULL,
		  PRIMARY KEY (uuid,room_id,port_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;`, numtblname)

	err := WyMysql.Exec(numsql).Error
	if err != nil {
		beego.Error("CreateTable num err ", numtblname, err)
	}
	return err
}

func GetLotteryorderTablename(date int64) string {
	numtblname := fmt.Sprintf("wy_tmp_lotteryorder_%v", date/100)
	return numtblname
}

func CheckLotteryorderTable(tblname string) error {

	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v (
			  orderid bigint NOT NULL AUTO_INCREMENT,
			  exp_orderid varchar(50) DEFAULT NULL,
			  parent_id bigint DEFAULT '0',
			  uuid bigint DEFAULT NULL,
			  account varchar(50) DEFAULT NULL,
  			  role_type varchar(16) DEFAULT NULL,
			  roomid int DEFAULT NULL,
			  roomeng varchar(20) DEFAULT NULL,
			  roomcn varchar(40) DEFAULT NULL,
			  pan varchar(10) DEFAULT NULL,
			  level int DEFAULT NULL,
			  pre_id bigint DEFAULT NULL,
			  master_id bigint DEFAULT NULL,
			  buhuo_type bigint DEFAULT '1',
			  res_id bigint DEFAULT '0',
			  expect varchar(40) DEFAULT NULL,
			  lottery_dalei varchar(40) DEFAULT NULL,
			  settle_type varchar(40)  DEFAULT NULL,
			  game_dalei varchar(40)  DEFAULT NULL,
			  game_xiaolei varchar(40)  DEFAULT NULL,
			  port_id int DEFAULT NULL,
			  item_id int DEFAULT NULL,
			  iteminfo varchar(40)  DEFAULT NULL,
			  official int DEFAULT NULL,
			  amount decimal(20,2) DEFAULT NULL,
			  is_zidongbuhuo int DEFAULT '0',
			  is_shoudongbuhuo int DEFAULT '0',
			  touzhufangshi varchar(10) DEFAULT NULL,
			  warning_flag int DEFAULT '0',
			  zhancheng_info text ,
			  numberinfo varchar(40) DEFAULT NULL,
			  base_odds decimal(20,4) DEFAULT '0',
			  dynimic_odds decimal(20,4) DEFAULT '0',
			  changlong_odds decimal(20,4) DEFAULT '0',
			  manager_odds decimal(20,4) DEFAULT '0',
			  order_num varchar(100) DEFAULT NULL,
			  win_items varchar(200) DEFAULT NULL,
			  orderinfo varchar(400)  NOT NULL,
			  optime bigint DEFAULT NULL,
			  state int DEFAULT '0',
			  settle_time bigint DEFAULT NULL,
			  opencode varchar(100)  DEFAULT NULL,
			  is_win int DEFAULT NULL,
			  win_list varchar(50)  DEFAULT NULL,
			  winodds decimal(20,4) DEFAULT NULL,
			  wager decimal(20,4) DEFAULT NULL,
			  shui_per decimal(20,4) DEFAULT NULL,
			  shui_amount decimal(20,4) DEFAULT NULL,
			  is_stats bigint DEFAULT '0',
			  expinfo varchar(40)  DEFAULT NULL,
			  PRIMARY KEY (orderid),
			  KEY room_expect_item (roomid,expect,item_id),
			  KEY uuid_roomid_optime_rettype (uuid,roomid,optime,is_win),
			  KEY index_settle (settle_time) USING BTREE
			) ENGINE=InnoDB AUTO_INCREMENT=11532600 DEFAULT CHARSET=utf8mb3;`, tblname)

	err := WyMysql.Exec(sql).Error
	if err != nil {
		beego.Error("CreateTable num err ", tblname, err)
	}
	return err
}

func GetLotdateTablename(date int64) string {
	numtblname := fmt.Sprintf("wy_tmp_user_lotdata_data_%v", date/100)
	return numtblname
}

func CheckLotdateTable(tblname string) error {

	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v (
			  uuid bigint NOT NULL,
			  pre_id bigint DEFAULT NULL,
			  master_id bigint DEFAULT NULL,
			  role_type varchar(20) DEFAULT NULL,
			  date bigint NOT NULL,
			  room_id bigint NOT NULL DEFAULT '0',
			  room_eng varchar(20) DEFAULT NULL,
			  room_cn varchar(40) DEFAULT NULL,
			  order_num bigint DEFAULT '0',
			  order_amount decimal(20,2) DEFAULT '0',
			  settled_num bigint DEFAULT '0',
			  settled_amount decimal(20,2) DEFAULT '0',
			  valid_amount decimal(20,2) DEFAULT '0',
			  revoke_amount decimal(20,2) DEFAULT '0',
			  wager decimal(20,4) DEFAULT '0' COMMENT '下注派彩',
			  tuishui decimal(20,4) DEFAULT '0' COMMENT '下注退水',
			  yingshouxiaxian decimal(20,4) DEFAULT '0',
			  shizhanhuoliang decimal(20,4) DEFAULT '0',
			  shizhanshuying decimal(20,4) DEFAULT '0',
			  shizhanjieguo decimal(20,4) DEFAULT '0',
			  shizhantuishui decimal(20,4) DEFAULT '0',
			  profit_tuishui decimal(20,4) DEFAULT '0' COMMENT '代理赚水',
			  profit_wager decimal(20,4) DEFAULT '0' COMMENT '代理赚赔',
			  shizhanpeicha decimal(20,4) DEFAULT '0',
			  yingkuijieguo decimal(20,4) DEFAULT '0',
			  shangjiaohuoliang decimal(20,4) DEFAULT '0',
			  shangjijiaoshou decimal(20,4) DEFAULT '0',
			  shouhuo decimal(20,4) DEFAULT '0',
			  shouhuoyingkui decimal(20,4) DEFAULT '0',
			  shoudongbuchu decimal(20,4) DEFAULT '0',
			  buchu decimal(20,4) DEFAULT '0',
			  buchuyingkui decimal(20,4) DEFAULT '0',
			  settle_shoudongbuchu decimal(20,4) DEFAULT '0',
			  settle_shoudongbuchuyingkui decimal(20,4) DEFAULT '0',
			  shizhanbuhuo decimal(20,4) DEFAULT '0',
			  shizhanbuhuoyingkui decimal(20,4) DEFAULT '0',
			  PRIMARY KEY (uuid,date,room_id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;`, tblname)

	err := WyMysql.Exec(sql).Error
	if err != nil {
		beego.Error("CreateTable num err ", tblname, err)
	}
	return err
}

func GetLotitemTablename(date int64) string {
	numtblname := fmt.Sprintf("wy_tmp_user_lotdata_item_%v", date/100)
	return numtblname
}

func CheckLotitemTable(tblname string) error {

	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %v (
			  uuid bigint NOT NULL,
			  pre_id bigint DEFAULT NULL,
			  master_id bigint DEFAULT NULL,
			  role_type varchar(20) DEFAULT NULL,
			  date bigint NOT NULL,
			  room_id bigint NOT NULL DEFAULT '0',
			  room_eng varchar(20) DEFAULT NULL,
			  room_cn varchar(40) DEFAULT NULL,
			  pan varchar(10) NOT NULL,
			  expect varchar(40) NOT NULL,
			  port_id bigint NOT NULL,
			  item_id bigint NOT NULL,
			  order_num bigint DEFAULT '0',
			  order_amount decimal(20,2) DEFAULT '0',
			  zhanchenghuoliang decimal(20,4) DEFAULT '0',
			  shizhanhuoliang decimal(20,4) DEFAULT '0',
			  xiajibuhuo decimal(20,4) DEFAULT '0',
			  zidongbuchu decimal(20,4) DEFAULT '0',
			  shoudongbuchu decimal(20,4) DEFAULT '0',
			  settled_num decimal(20,4) DEFAULT '0',
			  settled_amount decimal(20,2) DEFAULT '0',
			  valid_amount decimal(20,2) DEFAULT '0',
			  revoke_amount decimal(20,2) DEFAULT '0',
			  wager decimal(20,4) DEFAULT '0',
			  tuishui decimal(20,4) DEFAULT '0',
			  yingshouxiaxian decimal(20,4) DEFAULT '0',
			  settle_shizhanhuoliang decimal(20,4) DEFAULT '0',
			  shizhanshuying decimal(20,4) DEFAULT '0',
			  shizhanjieguo decimal(20,4) DEFAULT '0',
			  shizhantuishui decimal(20,4) DEFAULT '0',
			  profit_tuishui decimal(20,4) DEFAULT '0',
			  profit_wager decimal(20,4) DEFAULT '0',
			  shizhanpeicha decimal(20,4) DEFAULT '0',
			  yingkuijieguo decimal(20,4) DEFAULT '0',
			  shangjiaohuoliang decimal(20,4) DEFAULT '0',
			  shangjijiaoshou decimal(20,4) DEFAULT '0',
			  shouhuo decimal(20,4) DEFAULT '0',
			  shouhuoyingkui decimal(20,4) DEFAULT '0',
			  buchu decimal(20,4) DEFAULT '0',
			  buchuyingkui decimal(20,4) DEFAULT '0',
			  settle_shoudongbuchu decimal(20,4) DEFAULT '0',
			  settle_shoudongbuchuyingkui decimal(20,4) DEFAULT '0',
			  shizhanbuhuo decimal(20,4) DEFAULT '0',
			  shizhanbuhuoyingkui decimal(20,4) DEFAULT '0',
			  PRIMARY KEY (uuid,date,room_id,pan,expect,port_id,item_id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;`, tblname)

	err := WyMysql.Exec(sql).Error
	if err != nil {
		beego.Error("CreateTable num err ", tblname, err)
	}
	return err
}
