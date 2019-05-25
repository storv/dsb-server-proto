package utils

import (
	"dsb-server-admin/internal/config"
	"fmt"
	"github.com/bilibili/kratos/pkg/log"
	"github.com/samuel/go-zookeeper/zk"
	"strconv"
	"strings"
	"time"
)

type ZKService struct {
	conn             *zk.Conn
	Config           *config.AppConfig
	ZkPushServerPath string
	ZkServerDataPath string
	timeStr          string
}

//供config server分配pushsever节点使用
type ServerNodeInfo struct {
	IDC string `json:"idc"`
	Ip  string `json:"ip_port"`
}

type AppInfo struct {
	AppId        string
	ConnectCount int32
}

type PushServerInfo struct {
	MaxConnectCount int32
	ConnectCount    int32
	AppInfos        map[string]AppInfo
}

func NewZkService(appConfig *config.AppConfig) *ZKService {
	zKService := &ZKService{
		Config: appConfig,
	}
	zKService.ConnectToZk(appConfig, time.Second*5)
	return zKService
}

func (zkService *ZKService) StateLoop(event <-chan zk.Event) {
	for {
		select {
		case en := <-event:
			if en.State == zk.StateConnected {
				log.Info("zk state change to %v", zk.StateConnected)
			}
			if en.State == zk.StateExpired {
				log.Info("zk state change to %v", zk.StateExpired)
			}
		}
	}
}

// 连接到zookeeper注册中心
// 参数port
// 参数time.Duration是值的是重连时间周期
func (zkService *ZKService) ConnectToZk(config *config.AppConfig, recvTimeout time.Duration) (*zk.Conn, <-chan zk.Event) {
	zkService.timeStr = strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
	connArray := strings.Split(config.ZkConnInfo, ",")
	//连接 zookeeper
	conn, eventChan, err := zk.Connect(connArray, recvTimeout) //*10)
	zkService.conn = conn
	if err != nil {
		log.Info("Connect is error  : %s\n", err.Error())
		return nil, eventChan
	} else {
		//循环创建相关path
		zkService.ZkPushServerPath = config.ZkPushServerPath
		zkService.ZkServerDataPath = config.ZkServerDataPath
		zkService.CreatePath(zkService.ZkPushServerPath)
		zkService.CreatePath(zkService.ZkServerDataPath)
		//监听管道做出改变
		go zkService.StateLoop(eventChan)
		return conn, eventChan
	}

}

//创建永久path
// 如果path不存在循环创建子path
//
func (zkService *ZKService) CreatePath(path string) (string, error) {
	exist, _, error := zkService.conn.Exists(path)
	if error != nil {
		log.Info("Path is error  : %s\n", error.Error())
		return path, error
	}

	// 如果path存在直接返回
	if exist {
		log.Info("Path is exist : %s\n", path)
		return path, nil
	}

	log.Info("Path is : %s\n", path)
	if !exist {

		//切分path为数组
		paths := strings.Split(path, "/")
		temp := ""
		// 循环创建path
		for i, v := range paths {

			if paths[i] == "" {
				continue
			}
			temp += "/" + v

			exist, _, error := zkService.conn.Exists(temp)
			if error != nil {
				log.Info("Sub Path is exist  error  : %s\n\r", error.Error())
				break
			}
			log.Info("Sub path is: %s and paths is: %v and temp path is:%s and paths array size is:%d\n\r", v, paths, temp, len(paths))
			if exist {
				continue
			}
			if !exist {
				_, err := zkService.conn.Create(temp, nil, 0, zk.WorldACL(zk.PermAll))
				if err != nil {
					log.Info("Create Path is error  : %s, path is: %s \n\r", error.Error(), temp)
					break
				} else {
					log.Info("Create Path is ok , path is: %s \n\r", temp)
				}
			}
		}
	}
	return path, nil
}

// 获取指定path路径下的数据
func (zkService *ZKService) GetData(path string) ([]byte, error) {
	data, _, error := zkService.conn.Get(path)
	if error != nil {
		return []byte{}, error
	} else {
		return data, nil
	}
}

func (zkService *ZKService) GetClientNodePath(userId int64) (clientPath string) {
	clientPath = zkService.ZkServerDataPath + "/" + fmt.Sprintf("%d", userId)
	return
}

func (zkService *ZKService) GetIpNodePath(ip string) (ipPath string) {
	ipPath = zkService.ZkPushServerPath + "/" + ip
	return
}

func (zkService *ZKService) getRootPath() (clientPath string) {
	return zkService.ZkPushServerPath
}

func (zkService *ZKService) Close() {
	zkService.conn.Close()
}

func (zkService *ZKService) mirror(path string) (chan []string, chan error) {
	snapshots := make(chan []string)
	errors := make(chan error)

	go func() {
		for {
			snapshot, _, events, err := zkService.conn.ChildrenW(path)
			if err != nil {
				errors <- err
				return
			}
			snapshots <- snapshot
			evt := <-events
			if evt.Err != nil {
				errors <- evt.Err
				return
			}
		}
	}()

	return snapshots, errors
}

func (zkService *ZKService) ChildrenW(path string, onChange func([]string), onError func(error)) {
	if exist, _, _ := zkService.conn.Exists(path); !exist {
		log.Error("watching path %s is not exist.", path)
		return
	}
	snapshots, errors := zkService.mirror(path)
	go func() {
		for {
			select {
			case snapshot := <-snapshots:
				onChange(snapshot)
			case err := <-errors:
				onError(err)
			}
		}
	}()
}
