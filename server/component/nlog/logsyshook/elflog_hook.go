package logsyshook

import (
	"context"
	"encoding/json"
	"fmt"
	context_enum "github.com/krilie/lico_alone/common/com-model/context-enum"
	"github.com/krilie/lico_alone/common/config"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
	"github.com/krilie/lico_alone/common/utils/time_util"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type ElfLogHook struct {
	Key, Secret, Url string
	jsonFormatter    *logrus.JSONFormatter
	logChannel       chan *logChannelReq
	waitChannel      sync.WaitGroup
	onceStart        sync.Once
	onceStop         sync.Once
}

func NewElfLogHook(cfg *config.Config) *ElfLogHook {
	var elflog = &ElfLogHook{
		Key:    cfg.ElfLog.Key,
		Secret: cfg.ElfLog.Secret,
		Url:    cfg.ElfLog.Url,
		jsonFormatter: &logrus.JSONFormatter{
			TimestampFormat:  time_util.DefaultFormat,
			DisableTimestamp: false,
			DataKey:          "",
			FieldMap:         nil,
			CallerPrettyfier: nil,
			PrettyPrint:      false,
		},
		logChannel:  make(chan *logChannelReq, 600),
		waitChannel: sync.WaitGroup{},
		onceStart:   sync.Once{},
		onceStop:    sync.Once{},
	}
	elflog.StartPushLog(context2.NewContext())
	return elflog
}

func (e *ElfLogHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func GetLogStrValOrDefault(fields logrus.Fields, key, defVal string) string {
	val, ok := fields[key]
	if ok {
		return val.(string)
	} else {
		return defVal
	}
}

func (e *ElfLogHook) GetJsonContent(entry *logrus.Entry) string {
	format, err := e.jsonFormatter.Format(entry)
	if err != nil {
		return err.Error()
	}
	return string(format)
}

func (e *ElfLogHook) Fire(entry *logrus.Entry) error {
	var logData = &CreateLogReqModel{
		AppName:    GetLogStrValOrDefault(entry.Data, context_enum.AppName.Str(), ""),
		AppVersion: GetLogStrValOrDefault(entry.Data, context_enum.AppVersion.Str(), ""),
		AppHost:    GetLogStrValOrDefault(entry.Data, context_enum.AppHost.Str(), ""),
		RemoteIp:   GetLogStrValOrDefault(entry.Data, context_enum.RemoteIp.Str(), ""),
		ModuleName: GetLogStrValOrDefault(entry.Data, context_enum.Module.Str(), ""),
		FuncName:   GetLogStrValOrDefault(entry.Data, context_enum.Function.Str(), ""),
		ClientId:   GetLogStrValOrDefault(entry.Data, context_enum.ClientId.Str(), ""),
		Time:       entry.Time,
		TraceId:    GetLogStrValOrDefault(entry.Data, context_enum.TraceId.Str(), ""),
		UserId:     GetLogStrValOrDefault(entry.Data, context_enum.UserId.Str(), ""),
		Message:    entry.Message,
		TimeStamp:  time.Now().Unix(),
		Content:    e.GetJsonContent(entry),
		Level:      int(entry.Level), // 0-6
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
	e.logChannel <- &logChannelReq{
		Url:  e.Url,
		Key:  e.Key,
		Sign: sign,
		Data: jsonData,
	}
	return nil
}

// post方法
func (e *ElfLogHook) postLogJson(url, key, sign, data string) error {
	method := "POST"
	payload := strings.NewReader(data)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		println(err.Error())
		return errs.NewInternal().WithError(err)
	}
	req.Header.Add("key", key)
	req.Header.Add("sign", sign)
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		println(err.Error())
		return errs.NewInternal().WithError(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		println(err.Error())
		return errs.NewInternal().WithError(err)
	}
	// code message detail code=2000=success
	var ret = &ElfLogReturn{}
	err = json.Unmarshal(body, ret)
	if err != nil {
		println(err.Error())
		return errs.NewInternal().WithError(err)
	}
	if ret.Code != 2000 {
		fmt.Printf("code:%v message:%v detail:%v body:%v", ret.Code, ret.Message, ret.Detail, string(body))
		return errs.NewInternal().WithMsg(ret.Message)
	}
	return nil
}

func (e *ElfLogHook) StopPushLogWorker(ctx context.Context) {
	e.onceStop.Do(func() {
		close(e.logChannel)
		e.waitChannel.Wait()
		println("log channel closed.")
	})
}

func (e *ElfLogHook) StartPushLog(ctx context.Context) {
	e.onceStart.Do(func() {
		e.waitChannel.Add(1)
		go func() {
			defer e.waitChannel.Done()
			for {
				log, ok := <-e.logChannel
				if !ok {
					break
				} else {
					func() {
						defer func() {
							if pan := recover(); pan != nil {
								fmt.Printf("%v", pan)
							}
						}()
						err := e.postLogJson(log.Url, log.Key, log.Sign, log.Data)
						if err != nil {
							println(err.Error())
						}
					}()
				}
			}
		}()
	})
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
	TimeStamp  int64     `json:"time_stamp"` // unix 时间戳 防非法请求
	Content    string    `json:"content"`    // 所有内容的json形式
	Level      int       `json:"level"`      // level
}

type logChannelReq struct {
	Url  string
	Key  string
	Sign string
	Data string
}
