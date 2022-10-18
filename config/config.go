package config

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"gopkg.in/ini.v1"
)

type app struct {
	Host string
	Port string
}

type mysqlDB struct {
	Host string
	Port string
	User string
	Pass string
	DB   string
}

type redisHost struct {
	Host string
	Port string
	Pass string
	DB   string
}

var (
	App           app
	MysqlDB       mysqlDB
	RedisHost     redisHost
	PrivateKey    *rsa.PrivateKey
	PublicKey     *rsa.PublicKey
	RefreshSecret string
)

func init() {
	iniPath := "config/config.ini"
	if args := os.Args; len(args) > 1 {
		iniPath = args[1]
	}

	iniFile, err := ini.Load(iniPath)
	if err != nil {
		log.Fatalf("load %s error: %s \n", iniPath, err.Error())
		os.Exit(1)
	}

	app := iniFile.Section("app")
	App.Host = app.Key("Host").String()
	App.Port = app.Key("Port").String()

	database := iniFile.Section("mysql")
	MysqlDB.Host = database.Key("MysqlHost").String()
	MysqlDB.Port = database.Key("MysqlPort").String()
	MysqlDB.User = database.Key("MysqlUser").String()
	MysqlDB.Pass = database.Key("MysqlPass").String()
	MysqlDB.DB = database.Key("MysqlDB").String()

	redis := iniFile.Section("redis")
	RedisHost.Host = redis.Key("Host").String()
	RedisHost.Port = redis.Key("Port").String()
	RedisHost.Pass = redis.Key("Pass").String()
	RedisHost.DB = redis.Key("Db").String()

	// load rsa keys
	rsaSection := iniFile.Section("rsa")
	privKeyFile := rsaSection.Key("PRIV_KEY_FILE").String()
	priv, err := ioutil.ReadFile(privKeyFile)
	if err != nil {
		panic(err)
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)
	if err != nil {
		panic(err)
	}
	PrivateKey = privKey

	pubKeyFile := rsaSection.Key("PUB_KEY_FILE").String()
	pub, err := ioutil.ReadFile(pubKeyFile)
	if err != nil {
		panic(err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)
	if err != nil {
		panic(err)
	}
	PublicKey = pubKey

	RefreshSecret = rsaSection.Key("REFRESH_SECRET").String()
}
