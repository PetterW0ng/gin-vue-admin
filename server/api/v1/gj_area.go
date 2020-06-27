package v1

import (
	"fmt"
	"gin-vue-admin/global/response"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/service"
	"github.com/gin-gonic/gin"
)

func QueryAllProvince(c *gin.Context) {
	err, provinces := service.QueryAllProvince()
	if err == nil {
		response.OkWithData(gin.H{"provinces": provinces}, c)
	} else {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	}
}

func AreaList(c *gin.Context) {
	var areaTree []resp.AreaTree
	allArea := service.GetAreaAll()
	var areaMap = make(map[string]resp.AreaTree, len(areaTree))
	for _, item := range allArea {
		areaMap[item.Id] = resp.AreaTree{Id: item.Id, Name: item.Name, Pid: item.Pid}
		if item.Pid == "" {
			areaTree = append(areaTree, resp.AreaTree{Id: item.Id, Name: item.Name, Pid: "0", Children: nil})
		}
	}

	for _, a := range allArea {
		if areaT, ok := areaMap[a.Pid]; ok {
			areaT.Children = append(areaT.Children, areaMap[a.Id])
			areaMap[a.Pid] = areaT
			if areaT.Pid != "" {
				ppArea := areaMap[areaT.Pid]
				for i, v := range ppArea.Children {
					if v.Id == areaT.Id {
						ppArea.Children[i] = areaT
					}
				}
			}
		}
	}
	for i, v := range areaTree {
		v.Children = append(v.Children, areaMap[v.Id].Children...)
		areaTree[i] = v
	}
	response.OkWithData(gin.H{"provinces": areaTree}, c)
}

func FindByPid(c *gin.Context) {
	pid := c.Param("pid")
	err, children := service.FindCityOrArea(pid)
	if err == nil {
		response.OkWithData(gin.H{"children": children}, c)
	} else {
		response.FailWithMessage(fmt.Sprintf("查询失败，%v", err), c)
	}
}
