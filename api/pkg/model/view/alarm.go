package view

import (
	"github.com/clickvisual/clickvisual/api/pkg/model/db"
)

type ReqAlarmCreate struct {
	Name             string                    `json:"alarmName" form:"alarmName"` // 告警名称
	Desc             string                    `json:"desc" form:"desc"`           // 描述说明
	Interval         int                       `json:"interval" form:"interval"`   // 告警频率
	Unit             int                       `json:"unit" form:"unit"`           // 0 m 1 s 2 h 3 d 4 w 5 y
	Status           int                       `json:"status" form:"status"`
	AlertRule        string                    `json:"alertRule" form:"alertRule"` // prometheus alert rule
	View             string                    `json:"view" form:"view"`           // 数据转换视图
	NoDataOp         int                       `json:"noDataOp" form:"noDataOp"`
	Tags             map[string]string         `json:"tags" form:"tags"` //
	ChannelIds       []int                     `json:"channelIds" form:"channelIds"`
	Filters          []ReqAlarmFilterCreate    `json:"filters" form:"filters"`
	Conditions       []ReqAlarmConditionCreate `json:"conditions" form:"conditions"`
	Level            int                       `json:"level" form:"level"`
	DutyOfficers     []int                     `json:"dutyOfficers" form:"dutyOfficers"`
	IsDisableResolve int                       `json:"isDisableResolve" form:"isDisableResolve"`
}

func (r *ReqAlarmCreate) ConvertV2() {
	if len(r.Conditions) == 0 {
		return
	}
	if len(r.Filters) == 1 {
		r.Filters[0].Conditions = r.Conditions
	}
}

type AlarmFilterItem struct {
	*db.AlarmFilter
	Exp string
}

type ReqAlarmFilterCreate struct {
	Tid            int                       `json:"tid" form:"tid" binding:"required"`
	When           string                    `json:"when" form:"when" binding:"required"` // 执行条件
	SetOperatorTyp int                       `json:"typ" form:"typ"`                      // 0 default 1 INNER 2 LEFT OUTER 3 RIGHT OUTER 4 FULL OUTER 5 CROSS
	SetOperatorExp string                    `json:"exp" form:"exp"`                      // 操作
	Mode           int                       `json:"mode" form:"mode"`
	Conditions     []ReqAlarmConditionCreate `json:"conditions" form:"conditions"`
}

type ReqAlarmConditionCreate struct {
	SetOperatorTyp int `json:"typ" form:"typ"`                      // 0 when 1 and  2 or
	SetOperatorExp int `json:"exp" form:"exp"`                      // 0 avg 1 min 2 max 3 sum 4 count
	Cond           int `json:"cond" form:"cond"`                    // 0 above 1 below 2 outside range 3 within range
	Val1           int `json:"val1" form:"val1" binding:"required"` // 基准值/最小值
	Val2           int `json:"val2" form:"val2"`                    // 最大值
}

type (
	RespAlarmInfo struct {
		Filters     []RespAlarmInfoFilter          `json:"filters"`
		RelatedList []*db.RespAlarmListRelatedInfo `json:"relatedList"`

		Ctime int64 `json:"ctime"`
		Utime int64 `json:"utime"`

		db.Alarm

		Uid              int    `gorm:"-" json:"uid"`
		OaId             int64  `gorm:"column:oa_id;type:bigint(20);NOT NULL" json:"oaId"`                                // oa_id
		Username         string `gorm:"column:username;type:varchar(128);NOT NULL;index:uix_user,unique" json:"username"` // 用户名
		Nickname         string `gorm:"column:nickname;type:varchar(128);NOT NULL;index:uix_user,unique" json:"nickname"` // 昵称
		Secret           string `gorm:"column:secret;type:varchar(256);NOT NULL" json:"secret"`                           // 实例名称
		Phone            string `gorm:"column:phone;type:varchar(64);NOT NULL" json:"phone"`                              // phone
		Email            string `gorm:"column:email;type:varchar(64);NOT NULL" json:"email"`                              // email
		Avatar           string `gorm:"column:avatar;type:varchar(256);NOT NULL" json:"avatar"`                           // avatar
		Hash             string `gorm:"column:hash;type:varchar(256);NOT NULL" json:"hash"`                               // hash
		WebUrl           string `gorm:"column:web_url;type:varchar(256);NOT NULL" json:"webUrl"`                          // webUrl
		Oauth            string `gorm:"column:oauth;type:varchar(256);NOT NULL" json:"oauth"`                             // oauth
		State            string `gorm:"column:state;type:varchar(256);NOT NULL" json:"state"`                             // state
		OauthId          string `gorm:"column:oauth_id;type:varchar(256);NOT NULL" json:"oauthId"`                        // oauthId
		Password         string `gorm:"column:password;type:varchar(256);NOT NULL" json:"password"`                       // password
		CurrentAuthority string `gorm:"column:current_authority;type:varchar(256);NOT NULL" json:"currentAuthority"`      // currentAuthority
		Access           string `gorm:"column:access;type:varchar(256);NOT NULL" json:"access"`                           // access

		// Deprecated:
		Table db.BaseTable `json:"table"`
		// Deprecated:
		Instance db.BaseInstance `json:"instance"`
		// Deprecated: Conditions
		Conditions []*db.AlarmCondition `json:"conditions"`
	}

	RespAlarmInfoFilter struct {
		*db.AlarmFilter
		TableName  string               `json:"tableName"`
		Conditions []*db.AlarmCondition `json:"conditions"`
	}
)

type (
	ReqAlarmHistoryList struct {
		AlarmId   int `json:"alarmId" form:"alarmId"`
		StartTime int `json:"startTime" form:"startTime"`
		EndTime   int `json:"endTime" form:"endTime"` // 0 m 1 s 2 h 3 d 4 w 5 y
		db.ReqPage
	}

	RespAlarmHistoryList struct {
		Total int64              `json:"total"`
		Succ  int64              `json:"succ"`
		List  []*db.AlarmHistory `json:"list"`
	}
)

type (
	RespAlarmList struct {
		*db.Alarm
		RelatedList []*db.RespAlarmListRelatedInfo `json:"relatedList"`

		// Deprecated:
		TableName string `json:"tableName"`
		// Deprecated:
		TableDesc string `json:"tableDesc"`
		// Deprecated:
		Tid int `json:"tid"`
		// Deprecated:
		DatabaseName string `json:"databaseName"`
		// Deprecated:
		DatabaseDesc string `json:"databaseDesc"`
		// Deprecated:
		Did int `json:"did"`
		// Deprecated:
		InstanceName string `json:"instanceName"`
		// Deprecated:
		InstanceDesc string `json:"instanceDesc"`
		// Deprecated:
		Iid int `json:"iid"`
	}

	At struct {
		AtMobiles []string `json:"atMobiles"`
		IsAtAll   bool     `json:"isAtAll"`
	}

	DingTalkMarkdown struct {
		MsgType  string    `json:"msgtype"`
		At       *At       `json:"at"`
		Markdown *Markdown `json:"markdown"`
	}
	WeComMarkdown struct {
		MsgType       string         `json:"msgtype"`
		MentionedList *MentionedList `json:"mentionedList"`
		Markdown      *Markdown      `json:"markdown"`
	}

	MentionedList struct {
		UserIdList []string `json:"userIdList"`
		MobileList []string `json:"mobileList"`
	}

	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	}
)
