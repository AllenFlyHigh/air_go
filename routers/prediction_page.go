package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"task5/models"
)

func serviceGet (c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", nil)
}

func servicePost (c * gin.Context) {
	//获取查询的timePoint(起始时间), hour(持续时长), serviceType(预测物)
	timePoint := c.PostForm("time")
	serviceType := c.PostForm("serviceType")
	hour := c.PostForm("hour")
	_, err:= strconv.Atoi(hour)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"msg": "输入格式不正确",
			"code": http.StatusBadRequest,
		})
	} else {
		mlouts := models.ListMlouts(timePoint, hour, serviceType)
		c.HTML(http.StatusOK, "prediction_list.html", gin.H{
			"code": http.StatusOK,
			"data": mlouts,
		})
	}
}