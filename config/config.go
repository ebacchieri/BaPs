package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"log"
	"os"
	"reflect"
	"strconv"
	"github.com/gucooing/protocol/mx"
)

type Config struct {
	IsLite           bool       `json:"IsLite"`
	LogLevel         string     `json:"LogLevel"`
	Language         string     `json:"Language"`
	ResourcesPath    string     `json:"ResourcesPath"`
	DataPath         string     `json:"DataPath"`
	GucooingApiKey   string     `json:"GucooingApiKey"`
	AutoRegistration bool       `json:"AutoRegistration"`
	Tutorial         bool       `json:"Tutorial"`
	OtherAddr        *OtherAddr `json:"OtherAddr"`
	HttpNet          *HttpNet   `json:"HttpNet"`
	GateWay          *GateWay   `json:"GateWay"`
	DB               *DB        `json:"DB"`
	Irc              *Irc       `json:"Irc"`
	RankDB           *DB        `json:"RankDB"`
	Mail             *Mail      `json:"Mail"`
	Bot              *Bot       `json:"Bot"`
}

type OtherAddr struct {
	ServerInfoUrl     string `json:"ServerInfoUrl"`
	ManagementDataUrl string `json:"ManagementDataUrl"`
}

type GateWay struct {
	MaxPlayerNum       int64           `json:"MaxPlayerNum"`
	MaxCachePlayerTime int             `json:"MaxCachePlayerTime"`
	BlackCmd           map[string]bool `json:"BlackCmd"`
	IsLogMsgPlayer     bool            `json:"IsLogMsgPlayer"`
}

type HttpNet struct {
	InnerIp   string `json:"InnerIp"`
	InnerPort string `json:"InnerPort"`
	OuterAddr string `json:"OuterAddr"`
	Tls       bool   `json:"Tls"`
	CertFile  string `json:"CertFile"`
	KeyFile   string `json:"KeyFile"`
	Encoding  bool   `json:"Encoding"`
}

type DB struct {
	DbType string `json:"dbType"`
	Dsn    string `json:"dsn"`
}

type Irc struct {
	HostAddress string `json:"HostAddress"`
	Port        int32  `json:"Port"`
	Password    string `json:"Password"`
}

type Mail struct {
	Enable   bool   `json:"Enable"`
	Username string `json:"Username"`
	Password string `json:"Password"`
	Host     string `json:"Host"`
	Port     int    `json:"Port"`
}

type Bot struct {
	Addr     string `json:"Addr"`
	LoginNum int64  `json:"LoginNum"`
	CycLogin bool   `json:"CycLogin"`
}

var CONF *Config = nil

func SetDefaultConfig() {
	log.Printf("config不存在,使用默认配置\n")
	CONF = DefaultConfig
	CONF.check()
}

func GetConfig() *Config {
	if CONF == nil {
		SetDefaultConfig()
	}
	return CONF
}

func (c *Config) String() string {
	str, _ := sonic.MarshalString(c)
	return str
}

func GetIsLite() bool {
	return GetConfig().IsLite
}

func GetResourcesPath() string {
	return GetConfig().ResourcesPath
}

func GetDataPath() string {
	return GetConfig().DataPath
}

func GetGucooingApiKey() string {
	return GetConfig().GucooingApiKey
}

func GetAutoRegistration() bool {
	return GetConfig().AutoRegistration
}

func GetTutorial() bool {
	return GetConfig().Tutorial
}

func GetOtherAddr() *OtherAddr {
	return GetConfig().OtherAddr
}

func (x *OtherAddr) GetServerInfoUrl() string {
	return x.ServerInfoUrl
}

func (x *OtherAddr) GetManagementDataUrl() string {
	switch x.ManagementDataUrl {
	case "local":
		return fmt.Sprintf("%s/prod/index.json", GetHttpNet().GetOuterAddr())
	}
	return x.ManagementDataUrl
}

func GetHttpNet() *HttpNet {
	return GetConfig().HttpNet
}

func (x *HttpNet) GetOuterAddr() string {
	return x.OuterAddr
}

func GetGateWay() *GateWay {
	return GetConfig().GateWay
}

func GetIsLogMsgPlayer() bool {
	return GetConfig().GateWay.IsLogMsgPlayer
}

func GetBlackCmd() map[string]bool {
	return GetConfig().GateWay.BlackCmd
}

func GetRankDB() *DB {
	return GetConfig().RankDB
}

func GetIrc() *Irc {
	return GetConfig().Irc
}

func GetMail() *Mail {
	return GetConfig().Mail
}

func GetBot() *Bot {
	return GetConfig().Bot
}

var FileNotExist = errors.New("config file not found")

var DefaultConfig = &Config{
	IsLite:           false,
	LogLevel:         "Info",
	Language:         "",
	ResourcesPath:    "./resources",
	DataPath:         "./data",
	GucooingApiKey:   mx.GetMxToken(0, 32), // rand secure key, relies on 0 not being used as seed
	AutoRegistration: true,
	Tutorial:         false,
	OtherAddr: &OtherAddr{
		ServerInfoUrl:     "https://yostar-serverinfo.bluearchiveyostar.com",
		ManagementDataUrl: "https://prod-noticeindex.bluearchiveyostar.com/prod/index.json",
	},
	HttpNet: &HttpNet{
		InnerIp:   "0.0.0.0",
		InnerPort: "5000",
		OuterAddr: "http://127.0.0.1:5000",
		Tls:       false,
		CertFile:  "./config/cert.pem",
		KeyFile:   "./config/key.pem",
		Encoding:  true,
	},
	GateWay: &GateWay{
		MaxPlayerNum:       0,
		MaxCachePlayerTime: 10,
		BlackCmd:           make(map[string]bool),
		IsLogMsgPlayer:     false,
	},
	DB: &DB{
		DbType: "sqlite",
		Dsn:    "./sqlite/BaPs.db",
	},
	RankDB: &DB{
		DbType: "sqlite",
		Dsn:    "./sqlite/Rank.db",
	},
	Irc: &Irc{
		HostAddress: "127.0.0.1",
		Port:        16666,
		Password:    "mx123",
	},
	Mail: &Mail{
		Enable:   false,
		Username: "gucooing@BaPs.com",
		Password: "gucooing",
		Host:     "BaPs.com",
		Port:     587,
	},
}

func LoadConfig(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("配置文件读取失败将使用默认配置\n")
		CONF = DefaultConfig
	} else {
		defer func() {
			_ = f.Close()
		}()
		CONF = new(Config)
		d := json.NewDecoder(f)
		if err := d.Decode(CONF); err != nil {
			return err
		}
	}
	//log.Printf("env:%s\n\n", os.Environ())
	overrideWithEnv(reflect.ValueOf(CONF).Elem(), "Config")
	CONF.check()
	return nil
}

func overrideWithEnv(val reflect.Value, nestKey string) {
	if val.Kind() != reflect.Struct {
		return
	}
	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		if !field.CanSet() {
			continue
		}
		jsonTag := fieldType.Tag.Get("json")
		if jsonTag == "" || jsonTag == "-" {
			continue
		}
		envKey := nestKey
		if envKey != "" {
			envKey += "."
		}
		envKey += jsonTag
		if field.Kind() == reflect.Struct {
			overrideWithEnv(field, envKey)
			continue
		}
		if field.Kind() == reflect.Ptr && field.Type().Elem().Kind() == reflect.Struct {
			if field.IsNil() {
				field.Set(reflect.New(field.Type().Elem()))
			}
			overrideWithEnv(field.Elem(), envKey)
			continue
		}
		envValue, exists := os.LookupEnv(envKey)
		if !exists {
			continue
		}
		setFieldValue(field, envValue, envKey)
	}
}

func setFieldValue(field reflect.Value, envValue string, envKey string) {
	target := field
	if field.Kind() == reflect.Ptr {
		if field.IsNil() {
			field.Set(reflect.New(field.Type().Elem()))
		}
		target = field.Elem()
	}
	switch target.Kind() {
	case reflect.String:
		target.SetString(envValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if intVal, err := strconv.ParseInt(envValue, 10, 64); err == nil {
			target.SetInt(intVal)
		} else {
			log.Printf("环境变量 %s 的值 %s 无法转换为整数: %v", envKey, envValue, err)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if uintVal, err := strconv.ParseUint(envValue, 10, 64); err == nil {
			target.SetUint(uintVal)
		} else {
			log.Printf("环境变量 %s 的值 %s 无法转换为无符号整数: %v", envKey, envValue, err)
		}
	case reflect.Bool:
		if boolVal, err := strconv.ParseBool(envValue); err == nil {
			target.SetBool(boolVal)
		} else {
			log.Printf("环境变量 %s 的值 %s 无法转换为布尔值: %v", envKey, envValue, err)
		}
	case reflect.Float32, reflect.Float64:
		if floatVal, err := strconv.ParseFloat(envValue, 64); err == nil {
			target.SetFloat(floatVal)
		} else {
			log.Printf("环境变量 %s 的值 %s 无法转换为浮点数: %v", envKey, envValue, err)
		}
	case reflect.Struct:
		log.Printf("警告: 环境变量 %s 尝试设置结构体字段", envKey)
	default:
		log.Printf("不支持的类型: %s 字段类型 %v", envKey, target.Kind())
	}
}
