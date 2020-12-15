package constants

import (
	"errors"
	"github.com/jinzhu/gorm"
)

const (
	ERR_PARAMETER           = "参数错误"
	ERR_USER_INFO           = "用户信息错误"
	ERR_GET_USER_INFO       = "用户信息获取失败"
	ERR_USER_REPEAT_OPERATE = "重复操作"
	ERR_DATA_IS_NULL        = "数据不存在"
	ERR_WX_INFO             = "微信信息获取异常，请稍后再试"

	ERR_LOVECAR_ALREADY_BIND      = "该车辆已被其他车主绑定"
	ERR_LOVECAR_NUM_MAX           = "该车辆共同用车人已达上限"
	ERR_LOVECAR_REPEAT_BIND       = "您已存在与该车辆的绑定关系"
	ERR_LOVECAR_REPEAT_COMPLANINT = "当前车架有进行中的申诉，有问题请联系客服！"
	ERR_USER_HB_INSUFFICIENT      = "H币余额不足"
	ERR_SERVICE                   = "服务器异常，请稍后再试"
	ERR_WX_INFO_INVALID           = "微信信息已失效请,请返回上一页重新操作"
	ERR_ASK_CATE_NOT_EXIST        = "问答分类不存在或已下线"
	ERR_ASK_ADOPTED               = "不允许采纳"
	ERR_ASK_CATEGORY              = "问答分类不存在或已下架"
	ERR_ASK_LIKE                  = "收藏失败"
	ERR_ASK_REPLY_ZAN             = "点赞失败"
	ERR_ASK_REPLY_CANCEL_ZAN      = "取消点赞失败"
	ERR_ASK_QUESTION_CATE         = "问答分类最多选择2个"
	ERR_ASK_QUESTION_KM           = "行驶里程最小数为1，最大数为9999999"
	ERR_ASK_QUESTION_NOT_FIND     = "内容已下线"
	ERR_DEALER_NOT_EXIST          = "经销商不存在，请重新选择"
	ERR_APPOINT_ISNOT_FINISH      = "预约维保未完成"
	ERR_APPOINT_NOT_CANCEL        = "当前预约单不可以取消"
	ERR_CAR_STYLE_NOT_EXIST       = "所选车系不存在"
	ERR_APPEALING                 = "正在申诉中"
	ERR_CANCLE_APPOINT_MAX        = "操作过于频繁，请明日再试"
	ERR_TEST_DRIVE_EVERY_DAY      = "每天仅能提交2次预约单"
	ERR_APP_LOGIN_UNIQUE          = "同一个设备仅允许3个账号登录"
	ERR_REPORTED                  = "当前内容已被举报，不可重复举报"
	ERR_NDMS_CANCLE_APPOINT       = "当前预约单不可以取消"
	ERR_RECOMMEND_EQUAL_PH        = "所填手机号和推荐人手机号相同，请重新填写"
	ERR_RECOMMEND_EXPIRED         = "推荐信息已过期，请联系推荐人重新推荐"
	ERR_RECOMMEND_EXIST           = "您已预约过该店，请勿重复预约"
	ERR_RECOMMEND_LEAD_INVALID    = "同一人邀请最多只能预约三家店"
)

func RecordErrFormat(err error, val string) error {
	if gorm.IsRecordNotFoundError(err) {
		if val != "" {
			return errors.New(val)
		}
		return errors.New(ERR_DATA_IS_NULL)
	}
	return err
}
