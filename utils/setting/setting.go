package setting

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func init() {
	initJson()
	initDB()
	initServer()
	fmt.Println(Server)
	fmt.Println(DBConfig.URL)
}

var jsonData map[string]interface{}

// 初始化json，返回的map传入jsonData
func initJson() {
	bytes, err := ioutil.ReadFile("./config/app.json")
	if err != nil {
		log.Fatal("LoadConfigError:", err.Error())
	}

	configStr := string(bytes)
	fmt.Println(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		log.Fatal("JsonUnmarshalError:", err.Error())
	}
}

// input jsonData["database"] >> DBconfig
// 解析json节点，输出到对应的map
func setFiled(in map[string]interface{}, key string, out interface{}) {
	dbBytes, _ := json.Marshal(in[key])
	err := json.Unmarshal(dbBytes, out)
	if err != nil {
		log.Fatal("JsonUnmarshalDBError:", err.Error())
	}
}

// get DBconfig
type dBConfig struct {
	Dialect      string `json:"dialect"`
	Database     string `json:"database"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Charset      string `json:"charset"`
	URL          string `json:"url"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
	TablePrefix  string `json:"table_prefix"`
}

var DBConfig dBConfig

func initDB() {
	setFiled(jsonData, "database", &DBConfig)
	DBConfig.URL = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.Database, DBConfig.Charset)

}

// get server config
type server struct {
	Port         string `json:"port"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	PageSize     int    `json:"page_size"`
	JwtSecret    string `json:"jwt_secret"`
}

var Server server

func initServer() {
	setFiled(jsonData, "server", &Server)
}
