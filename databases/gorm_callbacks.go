package databases

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//plugin:gg_after_create
func ggAfterCreate(scope *gorm.Scope) {
	logrus.Debug("ggAfterCreate:")
	for _, v := range scope.Fields() {
		logrus.Debug("field:", v.Name, "value:", v.Field.String())
	}
	logrus.Debug("SQL:", scope.SQL)
	for _, v := range scope.SQLVars {
		logrus.Debug("SQLVars:", v)
	}
	logrus.Debug("CombinedConditionSql:", scope.CombinedConditionSql())
	logrus.Debug("TableName:", scope.TableName())
}

//plugin:gg_before_query_destination
func ggBeforeQueryDestination(scope *gorm.Scope) {
	logrus.Debug("ggBeforeQueryDestination:")
	logrus.Debug("TableName:", scope.TableName())
	for _, v := range scope.SQLVars {
		logrus.Debug("SQLVars:", v)
	}
	logrus.Debug("SQL:", scope.SQL)
}
