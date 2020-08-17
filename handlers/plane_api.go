package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/helper"
	"github.com/lvxin0315/gg/models"
	"github.com/lvxin0315/gg/params"
)

const planeSpecLen = 14
const maxPlane = 3

//绘制表格
var colList []map[string]interface{}

//用户与飞机数据存储
var userWithPlaneList = make(map[string]*models.UserPlane)

type PlanePointCoordinate struct {
	X string `json:"x" validate:"required"`
	Y int    `json:"y" validate:"min=0"`
}

func (p *PlanePointCoordinate) NewParams() params.GGParams {
	return new(PlanePointCoordinate)
}

type ShowPlanePointCoordinate struct {
	Direction string `json:"direction" validate:"required"`
	PlanePointCoordinate
}

func (p *ShowPlanePointCoordinate) NewParams() params.GGParams {
	return new(ShowPlanePointCoordinate)
}

func init() {
	params.InitParams("github.com/lvxin0315/gg/handlers.PlaneCoordinateApi", &ShowPlanePointCoordinate{})
	params.InitParams("github.com/lvxin0315/gg/handlers.SavePlaneApi", &ShowPlanePointCoordinate{})
	//A:0 -> N:13
	for i := 0; i < planeSpecLen; i++ {
		rowList := make(map[string]interface{})
		planeRowList := make(map[string]*models.PlaneCol)
		rowList["id"] = i
		for j := 0; j < planeSpecLen; j++ {
			rowList[helper.Num2Letter(j)] = ""
			planeRowList[helper.Num2Letter(j)] = new(models.PlaneCol)
		}
		colList = append(colList, rowList)
	}
}

//显示表格
func PlaneListApi(c *gin.Context) {
	if len(userWithPlaneList) > maxPlane {
		setGGError(c, fmt.Errorf("超过最大玩家数"))
		return
	}
	//token 作为用户标识
	if userWithPlaneList[getGGToken(c)] == nil {
		userWithPlaneList[getGGToken(c)] = new(models.UserPlane)
	}
	//获取已经存在的飞机点位
	userWithPlaneList[getGGToken(c)].GetAllPlaneColList()
	output := ggOutput(c)
	output.Data = colList
}

//显示表格
func GetUserAllPlaneListApi(c *gin.Context) {
	output := ggOutput(c)
	output.Data = userWithPlaneList[getGGToken(c)].GetAllPlaneColList()
}

//飞机顶点布局计算
func PlaneCoordinateApi(c *gin.Context) {
	ppc := ggParams(c).(*ShowPlanePointCoordinate)
	ppcX := helper.Letter2Num(ppc.X)
	ppcY := ppc.Y
	ppcList, err := pointHandler(ppcX, ppcY, ppc.Direction)
	if err != nil {
		setGGError(c, err)
		return
	}
	planeModel := new(models.Plane)
	planeModel.AddPlane(&models.PlaneCol{
		X: ppc.X,
		Y: ppc.Y,
	}, ppcList, ppc.Direction, planeSpecLen)
	//判断和已经确定的部分是否有重复
	err = userWithPlaneList[getGGToken(c)].CheckCol(planeModel)
	if err != nil {
		setGGError(c, err)
		return
	}
	output := ggOutput(c)
	output.Data = planeModel.GetPlaneColList()
}

//头向上飞机
func showUpPlaneFunc(ppcX, ppcY, x, y int) (*models.PlaneCol, error) {
	return checkPoint(x+ppcX, y+ppcY)
}

//头向下飞机
func showDownPlaneFunc(ppcX, ppcY, x, y int) (*models.PlaneCol, error) {
	return checkPoint(x+ppcX, -y+ppcY)
}

//头向左飞机
func showLeftPlaneFunc(ppcX, ppcY, x, y int) (*models.PlaneCol, error) {
	return checkPoint(y+ppcX, x+ppcY)
}

//头向右飞机
func showRightPlaneFunc(ppcX, ppcY, x, y int) (*models.PlaneCol, error) {
	return checkPoint(-y+ppcX, x+ppcY)
}

//验证每个点都是正数
func checkPoint(x, y int) (*models.PlaneCol, error) {
	if x < 0 {
		return nil, fmt.Errorf("左侧不足")
	}
	if x >= planeSpecLen {
		return nil, fmt.Errorf("右侧不足")
	}
	if y < 0 {
		return nil, fmt.Errorf("顶部不足")
	}
	if y >= planeSpecLen {
		return nil, fmt.Errorf("底部不足")
	}
	pl := new(models.PlaneCol)
	pl.X = helper.Num2Letter(x)
	pl.Y = y
	return pl, nil
}

//坐标计算(无顶点)
func pointHandler(ppcX, ppcY int, direction string) ([]*models.PlaneCol, error) {
	//先回执除顶点外，向上飞机其他部分坐标
	x := 0
	y := 0
	pointList := [][2]int{
		//翅膀（5格）
		{x - 2, y + 1}, {x - 1, y + 1}, {x, y + 1}, {x + 1, y + 1}, {x + 2, y + 1},
		//腰子（2格）
		{x, y + 2}, {x, y + 3},
		//尾部（3格）
		{x - 1, y + 4}, {x, y + 4}, {x + 1, y + 4},
	}
	var ppcList []*models.PlaneCol
	//预设函数
	var planeColFunc func(ppcX, ppcY, x, y int) (*models.PlaneCol, error)
	//根据方向设置
	switch direction {
	case models.PlaneDirectionLeft:
		planeColFunc = showLeftPlaneFunc
	case models.PlaneDirectionDown:
		planeColFunc = showDownPlaneFunc
	case models.PlaneDirectionRight:
		planeColFunc = showRightPlaneFunc
	default:
		planeColFunc = showUpPlaneFunc
	}
	for _, p := range pointList {
		pc, err := planeColFunc(ppcX, ppcY, p[0], p[1])
		if err != nil {
			return nil, err
		}
		ppcList = append(ppcList, pc)
	}
	return ppcList, nil
}

//保存飞机
func SavePlaneApi(c *gin.Context) {
	ppc := ggParams(c).(*ShowPlanePointCoordinate)
	ppcX := helper.Letter2Num(ppc.X)
	ppcY := ppc.Y
	ppcList, err := pointHandler(ppcX, ppcY, ppc.Direction)
	if err != nil {
		setGGError(c, err)
		return
	}
	//保存飞机
	planeModel := new(models.Plane)
	planeModel.AddPlane(&models.PlaneCol{
		X: ppc.X,
		Y: ppc.Y,
	}, ppcList, ppc.Direction, planeSpecLen)
	err = userWithPlaneList[getGGToken(c)].Save(planeModel)
	if err != nil {
		setGGError(c, err)
		return
	}
	output := ggOutput(c)
	output.Data = planeModel.GetPlaneColList()
}
