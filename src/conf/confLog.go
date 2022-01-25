package config

import (
	"fmt"
	"io"

	//github.com/chriszhangmq/file-rotatelogs v1.88.0 ：只有该版本最好用,有点缺陷：会显示压缩、删除失败。因为有5个hook对象执行，同一时间删除文件，但该文件已经被压缩了，所以出现压缩、删除失败（但不影响最终结果）
	//github.com/chriszhangmq/file-rotatelogs v1.88.0 ：屏蔽（不显示）v1.80.0的错误，可以使用。在6个钩子函数的情况下，防止重复压缩、删除（替代方案。使用多路输出代替钩子函数）
	rotatelogs "github.com/chriszhangmq/file-rotatelogs"
	"github.com/chriszhangmq/lfshook"
	"github.com/chriszhangmq/loglumber"
	//别人的（稍微好用点）
	//rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	log "github.com/lwydyby/logrus"
	"os"
	"strings"
	"time"
)

/**
经测试：已经可以使用
*/

const LogPath = "./log/"
const FileSuffix = ".log"
const FileName = "log"
const logLevel = "debug"
const dataFormat = "2006-01-02"
const TimeFormat = "2006-01-02 15:04:05"
const LogTimeFormat = "2006-01-02 15-04-05"
const LogTimeFormat1 = "2006-01-02T15-04-05"
const LogTimeFormat2 = "2006-01-02"
const LogDateFormat = "2006-01-02"
const TimeZone = "Asia/Shanghai"

func InitLog() {
	log.Info()
	//设置日志输出
	log.SetOutput(os.Stdout)
	//设置日志级别
	var loglevel log.Level
	err := loglevel.UnmarshalText([]byte(logLevel))
	if err != nil {
		log.Panicf("Failed to set log level：%v", err)
	}
	log.SetLevel(loglevel)

	//方式1（推荐使用）：按天、压缩，并发差1K
	//log.SetFormatter(&MineFormatter{})
	//mw := io.MultiWriter(os.Stdout, getLogger())
	//log.SetOutput(mw)

	//方式2（自己改的）：并发2W、可按照天数切割、压缩
	//问题：不建议使用（不用钩子函数）：这个容易导致bug，出现还未压缩文件时，直接把源文件删除地情况.
	//原因：6个钩子函数，同时执行6个定时任务，就会导致误删除文件
	//解决方式（还没处理这个问题，后续可用该方式改进并发量）：将定时任务移动到代码层面，只执行一个定时任务。
	//NewSimpleLogger(log.StandardLogger(), LogPath)

	//推荐使用该方法：多路输出日志（并发量1w）
	log.SetFormatter(&MineFormatter{})
	mw := io.MultiWriter(os.Stdout, writer1(LogPath))
	log.SetOutput(mw)
}

//(这是最好用的一个案例，但是并发量不行，只有1k)
//问题:有个bug,当并发上去之后,执行分割时(按大小或者天数),就会出现本应放在分割之前的日志数据,被放到分割之后的新文件内.
//    即:分割时间节点之前的数据,被写入新的文件洪
//原因:这是因为多协程执行日志写入的时候,刚好遇到了文件分割操作,此时旧的日志数据还未被写入旧的文件内,旧的文件就被关闭了,
//    执行分割压缩操作,这就导致了旧数据被写入新的文件内,出现日志数据混乱的错误
//解决方式:可以关闭这个组件的日志分割功能,只是用日志保存功能,然后使用centos自带的logrotate工具来分割日志.

//参考：https://github.com/natefinch/lumberjack
//参考：https://www.bbsmax.com/A/kjdw0kQ6zN/
func getLogger() *lumberjack.Logger {
	logger := &lumberjack.Logger{
		// 单个日志文件的大小, 单位是 MB（默认1PB：永不按照大小分割）（超过就进行日志分割、压缩）
		LogMaxSize: 10,

		// 日志文件保存的份数（默认为0：保存所有文件）（一般不用这个配置）
		//LogMaxSaveQuantity: 3,

		// 保留过期文件的时间,单位：天（默认为0：永不删除文件）
		LogMaxSaveDay: 4,

		//日志划分时间（默认为0:必须搭配LogMaxSaveDay才能使用）
		LogSplitDay: 1,

		// 是否需要压缩滚动日志, 使用的 gzip 压缩  (默认为false:关闭)
		Compress: true,

		//使用本地时间(默认false:使用UTC时间)
		LocalTime: true,

		//日志保存路径
		LogPathName: LogPath,

		//日志名称
		LogFileName: FileName,

		//日志后缀
		LogFileSuffix: FileSuffix,

		//日志中的时间格式
		LogFileTimeFormat: TimeFormat,
	}
	logger.Init()
	return logger
}

//日志文件
func NewSimpleLogger(logger *log.Logger, logPath string) {
	//（自己改的，不好用）所有级别的日志，都写入同一份文件
	lfHook1 := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer1(logPath),
		log.InfoLevel:  writer1(logPath),
		log.WarnLevel:  writer1(logPath),
		log.ErrorLevel: writer1(logPath),
		log.FatalLevel: writer1(logPath),
		log.PanicLevel: writer1(logPath),
	}, &MineFormatter{})
	logger.AddHook(lfHook1)

	//所有级别的日志，都写入同一份文件
	//lfHook3 := lfshook.NewHook(lfshook.WriterMap{
	//	log.DebugLevel: writer3(logPath + FileName),
	//	log.InfoLevel:  writer3(logPath + FileName),
	//	log.WarnLevel:  writer3(logPath + FileName),
	//	log.ErrorLevel: writer3(logPath + FileName),
	//	log.FatalLevel: writer3(logPath + FileName),
	//	log.PanicLevel: writer3(logPath + FileName),
	//}, &MineFormatter{})
	//logger.AddHook(lfHook3)

	// 为不同级别设置不同的输出目的
	//lfHook2 := lfshook.NewHook(lfshook.WriterMap{
	//	log.DebugLevel: writer2(logPath, "debug", save),
	//	log.InfoLevel:  writer2(logPath, "info", save),
	//	log.WarnLevel:  writer2(logPath, "warn", save),
	//	log.ErrorLevel: writer2(logPath, "error", save),
	//	log.FatalLevel: writer2(logPath, "fatal", save),
	//	log.PanicLevel: writer2(logPath, "panic", save),
	//}, &MineFormatter{})
	//logger.AddHook(lfHook2)
}

//这是自己的改的（好用）：github.com/chriszhangmq/file-rotatelogs
//目前功能：按天切割、按大小切割，并发量2W（实际部署选择v1.88.0，这个版本屏蔽了压缩失败的错误，在6个钩子函数的情况下，防止重复压缩、删除）
//可优化的位置：
//（1）getWriterNolock()函数，else if !sizeRotation && rl.rotationTime > 0，判断是否需要按照天来分割，这里的代码可以优化一下，减少代码的判断时间。
//（2）增加RotateCount参数：按照文件数量保留、删除
//文件设置：数据写入设定
func writer1(logPath string) *rotatelogs.RotateLogs {
	logFullPath := logPath
	var cstSh, _ = time.LoadLocation(TimeZone)
	//文件名：必须含有时间字段，且采用时间格式 ：%Y%m%d%H%M ，否则无法实现分割功能
	fileSuffix := time.Now().In(cstSh).Format(LogTimeFormat2) + FileSuffix
	fmt.Println("Create log file：" + logFullPath + "-" + fileSuffix)
	logier, err := rotatelogs.New(
		rotatelogs.WithFilePath(LogPath),
		rotatelogs.WithFileName(FileName),
		rotatelogs.WithMaxAge(2),       // 文件最大保存时间：2天
		rotatelogs.WithRotationTime(1), // 日志切割时间间隔：1天割一次log
		rotatelogs.WithLocation(time.Local),
		rotatelogs.WithRotationSize(200), //20MB
		rotatelogs.WithCompressFile(true),
		rotatelogs.WithCronTime("0 0 1 * * ?"),
	)
	if err != nil {
		fmt.Println("ERROR: rotatelogs.New--------------")
		panic(err)
	}
	logier.Init()
	return logier
}

//使用这个包 ： github.com/lestrrat/go-file-rotatelogs
//func writer2(logPath string, level string, save uint) *rotatelogs.RotateLogs {
//	logFullPath := path.Join(logPath, level)
//	var cstSh, _ = time.LoadLocation(TimeZone)
//	//文件名：必须含有时间字段，且采用时间格式 ：%Y%m%d%H%M ，否则无法实现分割功能
//	fileSuffix := time.Now().In(cstSh).Format("%Y%m%d%H%M") + FileSuffix
//	fmt.Println("Create log file：" + logFullPath + "-" + fileSuffix)
//	logier, err := rotatelogs.New(
//		logFullPath+"-"+fileSuffix,
//		rotatelogs.WithLinkName(logFullPath),      						// 生成软链，指向最新日志文件
//		rotatelogs.WithMaxAge(time.Duration(1) * time.Minute),   		// 文件最大保存时间：1min
//		rotatelogs.WithRotationTime(time.Duration(1) * time.Minute), 	// 日志切割时间间隔：1min分割一次log
//	)
//	if err != nil {
//		fmt.Println("ERROR: rotatelogs.New--------------")
//		panic(err)
//	}
//	return logier
//}

//func writer3(logPath string) *rotatelogs.RotateLogs {
//	logFullPath := logPath
//	var cstSh, _ = time.LoadLocation(TimeZone)
//	//文件名：必须含有时间字段，且采用时间格式 ：%Y%m%d%H%M ，否则无法实现分割功能
//	fileSuffix := time.Now().In(cstSh).Format("%Y%m%d%H%M") + FileSuffix
//	fmt.Println("Create log file：" + logFullPath + "-" + fileSuffix)
//	logier, err := rotatelogs.New(
//		logFullPath+"-"+fileSuffix,
//		rotatelogs.WithLinkName(logFullPath),      						// 生成软链，指向最新日志文件
//		rotatelogs.WithMaxAge(time.Duration(1) * time.Minute),   		// 文件最大保存时间：1min
//		rotatelogs.WithRotationTime(time.Duration(1) * time.Minute), 	// 日志切割时间间隔：1min分割一次log
//	)
//	if err != nil {
//		fmt.Println("ERROR: rotatelogs.New--------------")
//		panic(err)
//	}
//	return logier
//}

//日志数据记录格式
type MineFormatter struct {
}

func (s *MineFormatter) Format(entry *log.Entry) ([]byte, error) {
	msg := fmt.Sprintf("[%s] [%s] %s\n", time.Now().Local().Format(TimeFormat), strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}
