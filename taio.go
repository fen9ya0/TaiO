package main

import (
	"fmt"
	"github.com/binganao/Taio/service/scan"
)

func main() {
	//r := gin.Default()
	//
	//v1 := r.Group("api/v1")
	//{
	//	v1.GET("test", test.Test)
	//	v1.GET("job", job.AddJob)
	//}
	//
	//r.Run()

	fmt.Println(scan.MasScan(""))
}
