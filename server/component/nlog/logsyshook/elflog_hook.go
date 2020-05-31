package logsyshook

import (
	"encoding/json"
	"github.com/krilie/lico_alone/common/config"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/component/nlog"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type ElfLogHook struct {
	Key, Secret, Url string
	log              *nlog.NLog
}

func NewElfLogHook(cfg *config.Config, log *nlog.NLog) *ElfLogHook {
	return &ElfLogHook{
		Key:    cfg.ElfLog.Key,
		Secret: cfg.ElfLog.Secret,
		Url:    cfg.ElfLog.Url,
		log:    log,
	}
}

func (e *ElfLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (e *ElfLogHook) Fire(entry *logrus.Entry) error {
	var logData = &CreateLogReqModel{
		AppName:    entry.Data["ok"].(string),
		AppVersion: entry.Data,
		AppHost:    "",
		RemoteIp:   "",
		ModuleName: "",
		FuncName:   "",
		ClientId:   "",
		Time:       time.Time{},
		TraceId:    "",
		UserId:     "",
		Message:    "",
		TimeStamp:  time.Now().Unix(),
		Content:    "",
	}
	return e.PostLog(logData)
}

func (e *ElfLogHook) PostLog(logModel *CreateLogReqModel) error {
	jsonStr, err := json.Marshal(logModel)
	if err != nil {
		return errs.NewInternal().WithError(err)
	}
	jsonData := string(jsonStr)
	sign := pswd_util.Md5(jsonData + e.Secret)
	return e.postLogJson(e.Url, e.Key, sign, jsonData)
}

// post方法
func (e *ElfLogHook) postLogJson(url, key, sign, data string) error {
	method := "POST"
	payload := strings.NewReader(data)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		e.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	req.Header.Add("key", key)
	req.Header.Add("sign", sign)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		e.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		e.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	// code message detail code=2000=success
	var ret = &ElfLogReturn{}
	err = json.Unmarshal(body, ret)
	if err != nil {
		e.log.Error(err)
		return errs.NewInternal().WithError(err)
	}
	if ret.Code != 2000 {
		e.log.Errorf("%v %v %v %v", ret.Code, ret.Message, ret.Detail, string(body))
		return errs.NewInternal().WithMsg(ret.Message)
	}
	return nil
}

type ElfLogReturn struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

// key sign ...
type CreateLogReqModel struct {
	AppName    string    `json:"app_name"`
	AppVersion string    `json:"app_version"`
	AppHost    string    `json:"app_host"`
	RemoteIp   string    `json:"remote_ip"`
	ModuleName string    `json:"module_name"`
	FuncName   string    `json:"func_name"`
	ClientId   string    `json:"client_id"`
	Time       time.Time `json:"time"`
	TraceId    string    `json:"trace_id"`
	UserId     string    `json:"user_id"`
	Message    string    `json:"message"`
	TimeStamp  int64     `json:"time_stamp"` // unix 时间戳
	Content    string    `json:"content"`    // 所有内容的json形式
}
