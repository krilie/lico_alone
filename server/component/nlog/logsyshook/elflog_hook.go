package logsyshook

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/krilie/lico_alone/common/config"
	context2 "github.com/krilie/lico_alone/common/context"
	"github.com/krilie/lico_alone/common/errs"
	"github.com/krilie/lico_alone/common/utils/pswd_util"
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
			TimestampFormat:  time.RFC3339Nano,
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
	return e.PostLog(e.GetJsonContent(entry))
}

func (e *ElfLogHook) PostLog(logData string) error {
	timeStamp := time.Now().Unix()
	sign := pswd_util.Md5(logData + fmt.Sprintf("%v", timeStamp) + e.Secret)
	e.logChannel <- &logChannelReq{
		Url:       e.Url,
		Key:       e.Key,
		Sign:      sign,
		TimeStamp: timeStamp,
		Data:      logData,
	}
	return nil
}

// post方法
func (e *ElfLogHook) postLogJson(url, key, sign, data string, timeStamp int64) error {
	method := "POST"
	payload := strings.NewReader(data)
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		println(err.Error())
		return errs.NewInternal().WithError(err)
	}
	req.Header.Add("key", key)
	req.Header.Add("sign", sign)
	req.Header.Add("time_stamp", fmt.Sprintf("%v", timeStamp))
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
						err := e.postLogJson(log.Url, log.Key, log.Sign, log.Data, log.TimeStamp)
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

type logChannelReq struct {
	Url       string
	Key       string
	Sign      string
	TimeStamp int64
	Data      string
}
