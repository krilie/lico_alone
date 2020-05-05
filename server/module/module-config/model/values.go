package model

type ConfigItems string

func (item ConfigItems) Value() string {
	return string(item)
}

const (
	ConfigItemsIsInitData        ConfigItems = "common.is_init_data"       // 是否初始化数据 bool类型
	ConfigItemsNotificationEmail ConfigItems = "common.notification_email" // 系统消息的通知邮件
)
