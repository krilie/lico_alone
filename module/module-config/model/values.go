package model

type ConfigItems string

func (item ConfigItems) Val() string {
	return string(item)
}

const (
	ConfigItemsIsInitData        ConfigItems = "common.is_init_data"       // 是否初始化数据 bool类型
	ConfigItemsNotificationEmail ConfigItems = "common.notification_email" // 系统消息的通知邮件
	ConfigItemsIcpInfo           ConfigItems = "icp信息"                     // icp信息
	ConfigItemsAMapKey           ConfigItems = "AMapKey"                   // AMap key 高德地图
	ConfigItemsAboutApp          ConfigItems = "关于本站"                      // 关于本站信息 可以有格式
)
