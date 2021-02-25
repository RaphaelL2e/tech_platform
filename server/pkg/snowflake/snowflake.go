package snowflake

import (
	"strconv"

	"github.com/leeyf888/go-tools/snowflake"
	"github.com/sirupsen/logrus"
)

func GetUniqueId() string {
	sf, err := snowflake.NewWorker(1)
	if err != nil {
		logrus.Errorln(err)
		return GetUniqueId()
	}
	return strconv.FormatInt(sf.GetId(), 10)
}
