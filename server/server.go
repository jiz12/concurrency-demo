package server

import (
	"net/http"
)

//定义 Server接口
type Server interface {

	Route(pattern string,handler http.HandlerFunc)

	Run(address string)error

}

//定义 HttpServer接口实现
type HttpServer struct {

	Name string
	Srv  *http.Server

}

//实现 Server接口中Route方法
func (httpServer *HttpServer)Route(pattern string,handler http.HandlerFunc){
	http.HandleFunc(pattern,handler)
}

//实现 Server接口中Run方法
func (httpServer *HttpServer) Run()error{
	return httpServer.Srv.ListenAndServe()
}

// 返回httpServer实例
func NewServer(name string,address string)*HttpServer{

	return &HttpServer{
		Name:name,
		Srv:&http.Server{Addr: address},
	}

}


////启动 HTTP server
//func StartHttpServer(server *http.Server) error {
//	http.HandleFunc("/hello", HelloServer2)
//	fmt.Println("http server start")
//	err := srv.ListenAndServe()
//	return err
//}
//
//func HelloServer2(w http.ResponseWriter, r *http.Request) {
//	io.WriteString(w, "hello, world!\n")
//}