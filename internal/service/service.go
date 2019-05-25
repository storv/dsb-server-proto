package service

import (
	"context"
	dsb_server "dsb-server-admin/internal/client/dsb-server"
	"dsb-server-admin/internal/config"
	"dsb-server-admin/internal/dao"
	"dsb-server-admin/internal/utils"
	"encoding/json"
	"fmt"
	"github.com/bilibili/kratos/pkg/conf/paladin"
	concurrent "github.com/fanliao/go-concurrentMap"
	"github.com/prometheus/common/log"
	admin "gitlab.duoshengbu.com/be/dsb-server-proto/work/sg/admin/v1"
	server "gitlab.duoshengbu.com/be/dsb-server-proto/work/sg/server/v1"
	"sync"
)

// Service service.
type Service struct {
	ac        *paladin.Map
	dao       *dao.Dao
	ZkService *utils.ZKService
}

var websocketServerNodeOnLine = make([]string, 10)
var dsbServerClientsConnMap = make(map[string]server.DsbApiClient)

// New new a service and return.
func New() (s *Service) {
	var ac = new(paladin.TOML)
	if err := paladin.Watch("application.toml", ac); err != nil {
		panic(err)
	}

	var rc struct {
		Server *config.AppConfig
	}
	if err := paladin.Get("application.toml").UnmarshalTOML(&rc); err != nil {
		panic(err)
	}
	s = &Service{
		ac:        ac,
		dao:       dao.New(),
		ZkService: utils.NewZkService(rc.Server),
	}
	// watch websocket-server node
	s.ZkService.ChildrenW(s.ZkService.ZkPushServerPath, func(strings []string) {
		websocketServerNodeOnLine = strings
		log.Infof("watch [%s] changed as value is %v \n", s.ZkService.ZkPushServerPath, strings)
	}, func(e error) {
		log.Errorf("error on get watching changes %v n", e)
	})

	return s
}

// SayHelloURL bm demo func.
func (s *Service) SendMessage(ctx context.Context, req *admin.DsbAdminMessageReq) (reply *admin.DsbAdminMessageResp, err error) {
	var wg sync.WaitGroup
	var ipUsers = concurrent.NewConcurrentMap(100, float32(0.7), 1024)
	for _, v := range req.UserIdList {
		wg.Add(1)
		go func(userId int64) {
			defer wg.Done()
			path := s.ZkService.GetClientNodePath(userId)
			data, err := s.ZkService.GetData(path)
			if err != nil {
				log.Errorf(" error when  load user %s in path [%s]", userId, path)
				return
			}

			var dataMap map[string]int
			if json.Unmarshal(data, &dataMap) != nil {
				log.Errorf(" error when  load user %s in path [%s]", userId, path)
				return
			}
			for ip, _ := range dataMap {
				ipUsers.PutIfAbsent(ip, make([]int64, 100))
				ips, _ := ipUsers.Get(ip)
				ips = append(ips.([]int64), v)
				ipUsers.Put(ip, ips)
			}

		}(v)
	}
	wg.Wait()
	ipIt := ipUsers.Iterator()
	for ipIt.HasNext() {
		ip, uList, _ := ipIt.Next()
		if conn := s.getDsbServerClientConn(ip.(string)); conn != nil {
			_, err := conn.SendMessage(context.Background(),
				&server.DsbMessageReq{
					Message:    req.Message,
					UserIdList: uList.([]int64),
				})
			if err != nil {
				log.Errorf("send message to %s error,uList is %v,error is %v \n", ip, uList, err)
			}
		}
	}

	reply = &admin.DsbAdminMessageResp{
		Clients: 0,
		Status:  0,
	}
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
	s.dao.Close()
}

func (s *Service) getDsbServerClientConn(address string) server.DsbApiClient {
	clientConn := dsbServerClientsConnMap[address]
	if clientConn != nil {
		return clientConn
	}
	// 获取zk信息
	data, err := s.ZkService.GetData(s.ZkService.GetIpNodePath(address))
	if err != nil {
		return nil
	}
	var dataMap map[string]string
	if json.Unmarshal(data, &dataMap) != nil {
		return nil
	}

	gPort := dataMap["grpc.port"]
	if len(gPort) == 0 {
		return nil
	}

	host := fmt.Sprintf("%s:%s", address, gPort)
	cConn := dsb_server.NewClient(host)
	// 检测状态
	if cConn == nil {
		return nil
	}
	// 生成客户端连接
	dsbServerClientsConnMap[address] = cConn
	return cConn

}
