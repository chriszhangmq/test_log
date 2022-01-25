package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/lwydyby/logrus"
	config "github.com/test_log/src/conf"
	"github.com/test_log/src/create_file"
	"math/rand"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

/**
此代码最好用：已经通过测试，可以正常使用
日志写入文件测试代码
*/
func main() {
	//创建日志保存的文件夹
	create_file.InitCreateDirectory()
	//初始化日志
	config.InitLog()
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})
	//设置随机数
	rand.Seed(time.Now().UnixNano())
	//单线程测试
	//go func() {
	//	index := int64(1)
	//	pid := goID()
	//	for {
	//		log.Info("[ ", pid, " ]=== ", index, " ==================")
	//		log.Info("[ ", pid, " ]=== ", index, "========log info")
	//		log.Debug("[ ", pid, " ]=== ", index, "========log debug")
	//		err := errors.New("   记录新错误    ")
	//		log.Errorf("========log error, %s",err)
	//		index++
	//		//随机休眠
	//		randNum := rand.Intn(50)
	//		time.Sleep(time.Duration(randNum) * time.Millisecond)
	//		fmt.Println("============================================================",time.Now().Format(config.TimeFormat), )
	//	}
	//}()

	//多线程测试（并发100不会出现0点时数据错乱，200会出现14条数据写入错乱，并发量2w/s）
	for itor := 0; itor < 200; itor++ {
		go func() {
			index := int64(1)
			pid := goID()
			for {
				//测试大量数据写入
				log.Info("[ ", pid, " ]=== ", index, " ============ddsfgsfhfghgjkhkjhkjlhjklhjkljhklhjklhjklhjklhjklhjkllllllllllllllllllllll43cdfg3aw4sgty4a3wsg3t43taze3gt34sety3aswtyw33333rsx3ggtyx3yx3ys3======who have touched their lives.Love begins with a smile,grows with a kiss and ends with a tear.The brightest future will always be based on a forgotten past, you can’t go on well in lifeuntil you let go of your past failures and heartaches.When you were born,you were crying and everyone around you was smiling.Live your life so that when you die,you're the one who is smiling and everyone around you is crying.Please send this message to those people who mean something to you,to those who have touched your life in one way or another,to those who make you smile when you really need it,to those that make you see the brighter side of things when you are really down,to those who you want to let them know that you appreciate their friendship.And if you don’t, don’t worry,nothing bad will happen to you,you will just miss out on the opportunity to brighten someone’s day with this message.")
				log.Info("[ ", pid, " ]=== ", index, "========log info===========================3awescdfrtgcsxtacgpsadfgjnvaso9ugjpsaghasdghisdhgihsdohgidfhgdhgiudfhgusfghsfhgsfdhgsfdhgiuhgiusdghusdfghsoidghugThere are moments in life when you miss someone so much that you just want to pick them from your dreams and hug them for real! Dream what you want to dream;go where you want to go;be what you want to be,because you have only one life and one chance to do all the things you want to do.May you have enough happiness to make you sweet,enough trials to make you strong,enough sorrow to keep you human,enough hope to make you happy? Always put yourself in others’shoes.If you feel that it hurts you,it probably hurts the other person, too.The happiest of people don’t necessarily have the best of everything;they just make the most of everything that comes along their way.Happiness lies for those who cry,those who hurt, those who have searched,and those who have tried,for only they can appreciate the importance of people")
				log.Debug("[ ", pid, " ]=== ", index, "========log debug========================sfdjghsjghoiwerjgojogwgajrgiajdfgjsdjgfoasdjfoasdjfioasdjfaoifojaiodfoapfjpajfojdfpajfjadjfioajiofjoadifjoafjpofjafiToday I will pluck grapes of wisdom from the tallest and fullest vines in the vineyard,for these were planted by the wisest of my profession who have come before me,generation upon generation.Today I will savor the taste of grapes from these vines and verily I will swallow the seed of success buried in each and new life will sprout within me.The career I have chosen is laden with opportunity yet it is fraught with heartbreak and despair and the bodies of those who have failed, were they piled one atop another, would cast a shadow down upon all the pyramids of the earth.Yet I will not fail, as the others, for in my hands I now hold the charts which will guide through perilous waters to shores which only yesterday seemed but a dream.Failure no longer will be my payment for struggle. Just as nature made no provision for my body to tolerate pain neither has it made any provision for my life to suffer failure. Failure, like pain, is alien to my life. In the past I accepted it as I accepted pain. Now I reject it and I am prepared for wisdom and principles which will guide me out of the shadows into the sunlight of wealth, position, and happiness far beyond my most extravagant dreams until even the golden apples in the Garden of Hesperides will seem no more than my just reward.")
				log.Trace("[ ", pid, " ]=== ", index, "In truth, experience teaches thoroughly yet her course of instruction devours men's years so the value of her lessons diminishes with the time necessary to acquire her special wisdom. The end finds it wasted on dead men. Furthermore, experience is comparable to fashion; an action that proved successful today will be unworkable and impractical tomorrow.Only principles endure and these I now possess, for the laws that will lead me to greatness are contained in the words of these scrolls. What they will teach me is more to prevent failure than to gain success, for what is success other than a state of mind? Which two, among a thouand wise men, will define success in the same words; yet failure is always described but one way. Failure is man's inability to reach his goals in life, whatever they may be.In truth, the only difference between those who have failed and those who have successed lies in the difference of their habits. Good habits are the key to all success. Bad habits are the unlocked door to failure. Thus, the first law I will obey, which precedeth all others is --I will form good habits and become their slave.As a child I was slave to my impulses; now I am slave to my habits, as are all grown men. I have surrendered my free will to the years of accumulated habits and the past deeds of my life have already marked out a path which threatens to imprison my future. My actions are ruled by appetite, passion, prejudice, greed, love, fear, environment, habit, and the worst of these tyrants is habit. Therefore, if I must be a slave to habit let me be a slave to good habits. My bad habits must be destroyed and new furrows And how will I accomplish this difficult feat? Through these scrolls, it will be done, for each scroll contains a principle which will drive a bad habit from my life and replace it with one which will bring me closer to success. For it is another of nature's laws that only a habit can subdue another habit. So, in order for these written words to perform their chosen task, I must discipline myself with the first of my new habits which is as follows:")
				err := errors.New("   记录新错误    ")
				log.Errorf("========log error, %s", err)
				index++
				//随机休眠
				randNum := rand.Intn(50)
				time.Sleep(time.Duration(randNum) * time.Millisecond)
				fmt.Println("============================================================", time.Now().Format(config.TimeFormat))
			}
		}()
	}
	e.Start(":1323")
	//for{
	//	fmt.Println(time.Now())
	//	time.Sleep(time.Second)
	//}

}

func goID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
