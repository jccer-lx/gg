package models

import (
	"fmt"
	"github.com/lvxin0315/gg/helper"
)

const (
	PlaneDirectionUp    = "up"
	PlaneDirectionDown  = "down"
	PlaneDirectionLeft  = "left"
	PlaneDirectionRight = "right"
)

//坐标结构
type PlaneCol struct {
	X string `json:"x"`
	Y int    `json:"y"`
}

//飞机结构
type Plane struct {
	//机头
	head *PlaneCol
	//机身
	body []*PlaneCol
	//健康情况
	health bool
	//所有坐标
	planeColList []*PlaneCol
	//机头方向
	direction string
	//最大长度
	maxSpecLen int
}

//造个飞机
func (p *Plane) AddPlane(head *PlaneCol, body []*PlaneCol, direction string, maxSpecLen int) {
	p.head = head
	p.body = body
	p.direction = direction
	p.maxSpecLen = maxSpecLen
	p.health = true
	p.planeColList = append(append(p.planeColList, head), body...)
}

//获取所有col
func (p *Plane) GetPlaneColList() []*PlaneCol {
	return p.planeColList
}

//用户的飞机结构
type UserPlane struct {
	//用户标识
	userKey string
	//飞机
	planeList []*Plane
	//所有健康飞机坐标
	healthPlaneColList []*PlaneCol
	//所有飞机坐标
	allPlaneColList []*PlaneCol
	//准备就绪
	Ok bool
}

//检查飞机坐标是否重叠
func (up *UserPlane) CheckCol(p *Plane) error {
	for _, upc := range up.allPlaneColList {
		for _, pc := range p.planeColList {
			if upc.X == pc.X && upc.Y == pc.Y {
				return fmt.Errorf("有重复")
			}
		}
	}
	return nil
}

//保存飞机
func (up *UserPlane) Save(p *Plane) error {
	err := up.CheckCol(p)
	if err != nil {
		return err
	}
	up.planeList = append(up.planeList, p)
	//填充所有col
	up.healthPlaneColList = append(up.healthPlaneColList, p.planeColList...)
	up.allPlaneColList = append(up.allPlaneColList, p.planeColList...)
	return nil
}

//所有飞机坐标
func (up *UserPlane) GetAllPlaneColList() []*PlaneCol {
	return up.allPlaneColList
}
