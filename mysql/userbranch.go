package wymysql

import (
	"fmt"
	"github.com/aococo777/ltcommon/commonstruct"
	"math"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

/*
WY_user_branch:

获取直属下级列表 GetDirectUserbranchS(uuid int64) []commonstruct.BranchUsers
获取团队成员-不含自己 GetTeamuseridS(uuid int64) []int64
获取代理的所有下级 GetTeamIDList(uuid int64) []int64
获取直属下级个数 GetDirectuserNum(uuid int64) int64
获取用户branchinfo GetBranchinfo(uuid int64) (commonstruct.BranchUsers, error)
获取直属下级列表 GetDirectuserS(uuid int64) []commonstruct.BranchUsers
获取公司所有会员 GetCompanyuserS(uuid int64) []int64
创建新用户信息 CreateUserbranch(tx *gorm.DB, branch commonstruct.BranchUsers) error
获取系统所有用户 GetSysuuids() ([]commonstruct.BranchUsers, error)
获取团队某角色所有用户 GetTeamTotaluserS(uuid int64, roletype string) ([]commonstruct.Users, error)
获取团队某角色人数 GetTeamNum(uuid int64, roletype string) (int64, error)
获取团队某角色页码信息 GetTeamRoleSPageinfo(uuid int64, roletype string, pagecount int) (int64, int64)
获取团队某角色用户列表 GetTeamRoleS(uuid int64, roletype string, offset int, pagecount int) ([]commonstruct.Users, error)
*/

func GetTeamRoleS(uuids []int64, roletype string, limited string, sorttype string, offset int, pagecount int) ([]commonstruct.Users, error) {
	var list []commonstruct.Users

	var selectarg string

	switch limited {
	case "-1":
		selectarg = fmt.Sprintf("uuid in (?)")
	case "1,2":
		selectarg = fmt.Sprintf("uuid in (?) and limited in (1,2)")
	default:
		selectarg = fmt.Sprintf("uuid in (?) and limited = %v", limited)
	}
	var err error
	switch sorttype {
	case "createtime":
		err = WyMysql.Table(commonstruct.WY_user_base).Where(selectarg, uuids).
			Order("op_time desc, uuid").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&list).Error
	default:
		err = WyMysql.Table(commonstruct.WY_user_base).Where(selectarg, uuids).
			Order("last_logintime desc, uuid").Limit(pagecount).Offset((offset - 1) * pagecount).Find(&list).Error
	}
	return list, err
}

func GetTeamRoleSPageinfo(uuid int64, roletype string, pagecount int) (int64, int64) {
	var list commonstruct.Users
	var selectarg string
	switch roletype {
	case "all":
		selectarg = fmt.Sprintf("locate('%v',pre_list)>0", uuid)
	default:
		selectarg = fmt.Sprintf("locate('%v',pre_list)>0 and role_type = '%v'", uuid, roletype)
	}

	if err := WyMysql.Table(commonstruct.WY_user_branch).Select("count(*) as uuid").Where(selectarg).Find(&list).Error; err != nil {
		beego.Error("GetTeamRoleSPageinfo err", uuid, err)
	}
	return int64(list.Uuid), int64(math.Ceil(float64(list.Uuid) / float64(pagecount)))
}

func GetTeamNum(uuid int64, roletype string) (int64, error) {
	var info commonstruct.BranchUsers

	var selectarg string
	switch roletype {
	case "all":
		selectarg = fmt.Sprintf("locate('%v',pre_list)>0", uuid)
	default:
		selectarg = fmt.Sprintf("locate('%v',pre_list)>0 and role_type = '%v'", uuid, roletype)
	}

	err := WyMysql.Table(commonstruct.WY_user_branch).Select("count(*) as uuid").
		Where(selectarg).Find(&info).Error
	if err != nil {
		beego.Error("GetTeamNum err = ", err.Error())
	}
	return info.Uuid, err
}

func GetTeamTotaluserS(uuid int64, roletype string) ([]commonstruct.Users, error) {
	var list []commonstruct.Users
	var selectarg string
	switch roletype {
	case "all":
		selectarg = fmt.Sprintf("locate('%v',pre_list)>0", uuid)
	default:
		selectarg = fmt.Sprintf("locate('%v',pre_list)>0 and role_type = '%v'", uuid, roletype)
	}

	err := WyMysql.Table(commonstruct.WY_user_branch).Where(selectarg).Order("uuid").Find(&list).Error

	return list, err
}

// 获取系统所有用户
func GetSysuuids() ([]commonstruct.BranchUsers, error) {
	var info []commonstruct.BranchUsers
	if err := WyMysql.Table(commonstruct.WY_user_branch).Select("uuid,master_id").Find(&info).Error; err != nil {
		beego.Error("GetSysuuids err", err)
		return info, err
	}
	return info, nil
}

/***************************************
* 创建用户上下级信息
****************************************/
func CreateUserbranch(tx *gorm.DB, branch commonstruct.BranchUsers) error {
	return tx.Table(commonstruct.WY_user_branch).Create(&branch).Error
}

// 获取公司所有会员
func GetCompanyuserS(uuid int64) []int64 {
	var list []commonstruct.AgentOdds
	if err := WyMysql.Table(commonstruct.WY_user_branch).Select("uuid").Where("master_id = ? ", uuid).Find(&list).Error; err != nil {
		beego.Error("GetCompanyuserS err", uuid, err)
	}

	var ret []int64

	for _, user := range list {
		if user.Uuid != uuid {
			ret = append(ret, user.Uuid)
		}
	}
	return ret
}

func GetDirectuserS(uuid int64) []commonstruct.BranchUsers {
	var list []commonstruct.BranchUsers
	if err := WyMysql.Table(commonstruct.WY_user_branch).Where("pre_id = ?", uuid).Find(&list).Error; err != nil {
		beego.Error("GetDirectuserS err", uuid, err)
	}
	return list
}

// 获取用户branchinfo
func GetBranchinfo(uuid int64) (commonstruct.BranchUsers, error) {
	var list commonstruct.BranchUsers
	err := WyMysql.Table(commonstruct.WY_user_branch).Where("uuid = ?", uuid).Find(&list).Error
	if err != nil {
		beego.Error("GetBranchinfo err", uuid, err)
	}
	return list, err
}

// 获取直属下级个数
func GetDirectuserNum(uuid int64) int64 {
	var branchinfo commonstruct.BranchUsers
	if err := WyMysql.Table(commonstruct.WY_user_branch).Select("count(*) as uuid").Where("pre_id = ?", uuid).Find(&branchinfo).Error; err != nil {
		beego.Error("GetDirectuserNum err", uuid, err)
	}
	return branchinfo.Uuid
}

// 获取代理的所有下级
func GetTeamIDList(uuid int64) []int64 {
	var list []commonstruct.BranchUsers
	if err := WyMysql.Table(commonstruct.WY_user_branch).Select("uuid").Where(fmt.Sprintf("locate('%v',pre_list)>0", uuid)).Find(&list).Error; err != nil {
		beego.Error("GetTeamIDList err", uuid, err)
	}

	var ret []int64

	for _, user := range list {
		ret = append(ret, user.Uuid)
	}
	ret = append(ret, uuid)
	return ret
}

// 获取团队成员 - 不含自己
func GetTeamuseridS(uuid int64, roletype string, limited string) []int64 {

	var selectarg string
	switch roletype {
	case "all":
		selectarg = fmt.Sprintf("locate('%v',pre_list)>0", uuid)
	default:
		selectarg = fmt.Sprintf("locate('%v',pre_list)>0 and role_type = '%v'", uuid, roletype)
	}

	var list []commonstruct.AgentOdds
	if err := WyMysql.Table(commonstruct.WY_user_branch).Select("uuid").Where(selectarg).Find(&list).Error; err != nil {
		beego.Error("GetTeamIDList err", uuid, err)
	}

	var ret []int64
	for _, user := range list {
		if user.Uuid != uuid {

			switch limited {
			case "-1":
				ret = append(ret, user.Uuid)
			case "1,2":
				if userinfo, err := GetUserinfoByUuid(user.Uuid); err == nil {
					if userinfo.Limited == 1 || userinfo.Limited == 2 {
						ret = append(ret, user.Uuid)
					}
				}
			case "0", "3", "1", "2":
				limitednum, _ := strconv.ParseInt(limited, 10, 64)
				if userinfo, err := GetUserinfoByUuid(user.Uuid); err == nil {
					if userinfo.Limited == limitednum {
						ret = append(ret, user.Uuid)
					}
				}
			default:
				return nil
			}
		}
	}
	return ret
}

// 获取直属下级列表
func GetDirectUserbranchS(uuid int64) []commonstruct.BranchUsers {
	var SufUserS []commonstruct.BranchUsers
	if err := WyMysql.Table(commonstruct.WY_user_branch).Where("pre_id = ?", uuid).Find(&SufUserS).Error; err != nil {
		beego.Error("GetSufUserS err", uuid, err)
	}
	return SufUserS
}
