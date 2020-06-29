package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/gg/etc"
	"github.com/sirupsen/logrus"
)

//记录访问次数
var PongTime = 0

func Pong(c *gin.Context) {
	PongTime++
	logrus.SetLevel(logrus.DebugLevel)
	d, _ := c.GetRawData()
	logrus.Println("我是test4的：", string(d))
	//试试配置文件
	//tryConfig()
	//试试连接池
	//tryDBClient()
	//试试模型操作-查询100条
	//tryDBSelect()
	//试试memDB
	//if PongTime > 1 {
	//	tryGetDataMemDB()
	//} else {
	//	tryMemDB()
	//}
	//返回值：pong
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//试试连接池
//func tryDBClient() {
//	logrus.Debug("tryDBClient:")
//	t := 20
//	for {
//		databases.NewDB().Exec("SELECT * FROM mall_article")
//		t--
//		//试试模型操作-插入数据
//		articleModel := new(models.MallArticle)
//		articleModel.Author = fmt.Sprintf("Author%d", t)
//		articleModel.Title = fmt.Sprintf("Title%d", t)
//		articleModel.ShareTitle = fmt.Sprintf("ShareTitle%d", t)
//		err := databases.NewDB().Model(models.MallArticle{}).Save(articleModel).Error
//		if err != nil {
//			logrus.Error(err)
//		}
//
//		if t == 0 {
//			break
//		}
//	}
//}

//试试查询
//func tryDBSelect() {
//	logrus.Debug("tryDBSelect:")
//	var articleModelList []*models.MallArticle
//	err := databases.NewDB().Model(&models.MallArticle{}).Where("id > ?", 10).Limit(100).Scan(&articleModelList).Error
//	if err != nil {
//		logrus.Error(err)
//	}
//	logrus.Println(articleModelList[0].Author)
//}

//试试配置文件
func tryConfig() {
	logrus.Debug("tryConfig:")
	logrus.Println("etc.Config.APPName:", etc.Config.APPName)
	//logrus.Println("etc.Config.Contacts[0].Name:", etc.Config.Contacts[0].Name)
	logrus.Println("etc.Config.DB.Host:", etc.Config.DB.Host)
	logrus.Println("etc.Config.DB.MaxIdleConns:", etc.Config.DB.MaxIdleConns)
	logrus.Println("etc.Config.DB.MaxOpenConns:", etc.Config.DB.MaxOpenConns)
}

//试试memdb
//func tryMemDB() {
//	logrus.Debug("tryMemDB:")
//	var articleModelList []*models.MallArticle
//	err := databases.NewDB().Model(&models.MallArticle{}).Limit(10000).Scan(&articleModelList).Error
//	if err != nil {
//		logrus.Error(err)
//	}
//	//保存到memDB
//	memberTable, err := databases.NewMemBD().CreateTableSchema("MallArticle", reflect.TypeOf(&models.MallArticle{}))
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//	for _, articleModel := range articleModelList {
//		_, err = memberTable.Insert(articleModel)
//		if err != nil {
//			logrus.Error(err)
//			break
//		}
//	}
//}

//试试在memDB 读取数据
//func tryGetDataMemDB() {
//	logrus.Debug("tryGetDataMemDB:")
//	table, err := databases.NewMemBD().GetTableSchema("MallArticle")
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//	dataList, err := table.Select(10, 10)
//	if err != nil {
//		logrus.Error(err)
//		return
//	}
//	for _, article := range dataList {
//		logrus.Debug("article.Author", article.(*models.MallArticle).Author)
//	}
//}
