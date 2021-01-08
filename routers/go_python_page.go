package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"task5/models"
	"task5/pkg/util"
)
//初始化rocket
func init ()  {
	util.StartRocketMQ()
}

func goPythonPost (c * gin.Context) {
	//获取查询的timePoint(起始时间), hour(持续时长), serviceType(预测物)
	hasDone := c.PostForm("hasDone")
	fmt.Println("python端数据抓取和模型计算完毕，开始生产消息" + hasDone)
	c.JSON(http.StatusBadRequest,gin.H{
		"msg": "接受消息成功！",
		"code": http.StatusOK,
	})
	//从数据库拿预测数据，因为此时的模型环境还没有配置好，此时只是模拟这个过程
	mlmets := models.ListMlmets()
	for _, mlmet := range(mlmets) {
		fmt.Println(mlmet)
		data, _ := json.Marshal(&mlmet)
		fmt.Printf("%s", data)
		fmt.Println()
		util.SendMessageToRocketChan(data, 1,"test")
	}
	//消息发送完毕
}