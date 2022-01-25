package middleware

//
///**
// * Created by Chris on 2021/9/15.
// */
//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"github.com/labstack/echo"
//	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
//	log "github.com/lwydyby/logrus"
//	"github.com/rifflock/lfshook"
//	"github.com/sirupsen/logrus"
//	"os"
//	"path"
//	"strconv"
//	"time"
//)
//
//// 日志记录到文件
//func LoggerToFile() echo.MiddlewareFunc {
//	fmt.Println("11111111111111111111111111111")
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		fmt.Println("2222222222222222222222222222")
//		logFilePath := conf.LoggerConfig.LogFilePath
//		logFileName := conf.LoggerConfig.LogFileName
//
//		// 日志文件
//		fileName := path.Join(logFilePath, logFileName)
//
//		// 写入文件
//		src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
//		if err != nil {
//			fmt.Println("err", err)
//		}
//
//		// 实例化
//		logger := logrus.New()
//
//		// 设置输出
//		logger.Out = src
//
//		// 设置日志级别
//		logger.SetLevel(logrus.DebugLevel)
//
//		// 设置 rotatelogs
//		logWriter, err := rotatelogs.New(
//			// 分割后的文件名称
//			fileName + ".%Y%m%d.log",
//
//			// 生成软链，指向最新日志文件
//			rotatelogs.WithLinkName(fileName),
//
//			// 设置最大保存时间(7天)
//			rotatelogs.WithMaxAge(1*time.Second),
//
//			// 设置日志切割时间间隔(1天)
//			rotatelogs.WithRotationTime(20*time.Second),
//		)
//
//		writeMap := lfshook.WriterMap{
//			logrus.InfoLevel:  logWriter,
//			logrus.FatalLevel: logWriter,
//			logrus.DebugLevel: logWriter,
//			logrus.WarnLevel:  logWriter,
//			logrus.ErrorLevel: logWriter,
//			logrus.PanicLevel: logWriter,
//		}
//
//		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
//			TimestampFormat:"2006-01-02 15:04:05",
//		})
//
//		// 新增 Hook
//		logger.AddHook(lfHook)
//		return func(c echo.Context) (err error) {
//			fmt.Println("3333333333333333333333333333333")
//			req := c.Request()
//			res := c.Response()
//			start := time.Now()
//			err = next(c)
//			stop := time.Now()
//			path1 := req.URL.Path
//			if path1 == "" {
//				path1 = "/"
//			}
//			var e string
//			if err != nil {
//				b, _ := json.Marshal(err.Error())
//				b = b[1 : len(b)-1]
//				e = string(b)
//			}
//			cl := req.Header.Get(echo.HeaderContentLength)
//			if cl == "" {
//				cl = "0"
//			}
//			log.Infof("remote_ip:%v; host:%v; uri:%v; method:%v; path1:%v; protocol:%v; referer:%v; user_agent:%v; status:%v; error:%v; "+
//				"latency_human:%v; bytes_in:%v; bytes_out:%v",
//				c.RealIP(), req.Host, req.RequestURI, req.Method, path1, req.Proto, req.Referer(), req.UserAgent(), res.Status, e,
//				stop.Sub(start).String(), cl, strconv.FormatInt(res.Size, 10))
//			// 日志格式
//			logger.WithFields(logrus.Fields{
//				"status_code"  : c.RealIP(),
//				"latency_time" : req.Method,
//				"client_ip"    : res.Status,
//				"req_method"   : stop.Sub(start).String(),
//				"req_uri"      : req.Host,
//			}).Info()
//			return
//		}
//	}
//}
//
//
//// 日志记录到 MongoDB
//func LoggerToMongo() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}
//
//// 日志记录到 ES
//func LoggerToES() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}
//
//// 日志记录到 MQ
//func LoggerToMQ() gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//	}
//}
//
