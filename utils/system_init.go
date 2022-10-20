package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigFile("app")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))

}

func InitMySQL() {
	newLogger := logger.New(
		//自定义日志模板，打印sql语句
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
		//// log.LstdFlags 设置初始值：相当于 log.Ldate|log.Ltime
		////log.Llongfile 显示完整的文件名和行数 除了这之外还有 Lmicroseconds
		//// Llongfile Lshortfile LUTC
		//// log.new 有三个参数，第一个输出位置，第二个为日志输出前缀，第三个设置logger的属性
		//logger = log.New(file, "crm_", log.LstdFlags|log.Llongfile)
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{Logger: newLogger})
	//user := models.UserBasic{}
	//DB.Find(&user)
	//fmt.Println(user)
}

func InitRedis() {
	Red := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"), // no password set
		DB:           viper.GetInt("redis.DB"),          // use default DB
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	pong, err := Red.Ping().Result()
	if err != nil {
		fmt.Println("init redis failed...", err)
	} else {
		fmt.Println("init redis success", pong)

	}
}
