package application

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"v1/config"
	"v1/db"
	"v1/entity"
	"v1/repository"
	"v1/service"
)

type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

func Get() (*Application, error) {
	cfg := config.Get()

	dbc, err := db.Get(context.Background(), cfg.DSN)
	if err != nil {
		fmt.Print(err)
	}

	app := &Application{
		DB:  dbc,
		Cfg: cfg,
	}

	r := repository.NewRepository(app.DB)
	systemMouseRepository := repository.NewSystemMouseRepository(app.DB)
	s := service.NewBrainlinkService(r, systemMouseRepository)

	buildMapService := service.NewBuilderMap(r, systemMouseRepository)
	p, conn := getConn("0.0.0.0", 1234)

	go func() {
		http.HandleFunc("/", buildMapService.Handle)
		http.ListenAndServe(":5001", nil)
	}()

	go func() {
		http.HandleFunc("/event", buildMapService.HandleEvent)
		http.ListenAndServe(":5001", nil)
	}()

	listEEG(conn, p, s)

	return app, nil
}

func listEEG(conn *net.UDPConn, p []byte, s service.IBrainlinkService) {
	for {
		n, _, err := conn.ReadFromUDP(p)

		print(n)
		if err != nil {
			panic(err)
		}

		var model entity.EegDto
		err = json.Unmarshal(p[:n], &model)
		if err != nil {
			panic(err)
		}
		s.Add(context.Background(), &model)
		fmt.Println(model)
	}
}

func sendMap(conn *net.UDPConn, s service.IBrainlinkService) {
	m, err := s.List(context.Background())
	if err != nil {
		//todo: err
		print(err)
	}
	b, err := json.Marshal(m)
	if err != nil {
		//todo: err
		print(err)
	}
	_, err = conn.Write(b)
	if err != nil {
		//todo: err
		print(err)
	}
}
func getMap(s service.IBrainlinkService) ([]byte, error) {
	m, err := s.List(context.Background())
	if err != nil {
		//todo: err
		print(err)
	}
	return json.Marshal(m)
}

func getConn(ip string, port int) ([]byte, *net.UDPConn) {
	p := make([]byte, 2048) // буфер для получения данных
	addr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(ip),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}
	return p, conn
}

func (a *Application) Stop() error {
	a.DB.Close()
	//defer a.Ext.C.Close()

	return nil
}
