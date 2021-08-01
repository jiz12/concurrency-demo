package main

import (
	"concurrency-demo/server"
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	c := context.Background()

	c, cancel := context.WithCancel(c)

	group, errC := errgroup.WithContext(c)

	/*初始化服务器对象*/
	server := server.NewServer("server A",":9090")

	/*启动服务器*/
	group.Go(func() error {

		server.Route("/hello", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w,"你好世界。")
		})

		return server.Run()

	})

	/*关闭server*/
	group.Go(func() error {
		<-errC.Done() //
		fmt.Println("http server stop")
		return server.Srv.Shutdown(errC)
	})

	chanel := make(chan os.Signal, 1)
	signal.Notify(chanel)

	group.Go(func() error {
		for {
			select {
			case <-errC.Done():
				return errC.Err()
			case <-chanel:
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")

}