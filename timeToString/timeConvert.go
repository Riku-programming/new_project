package TimeToString

import (
	"time"
)

//  timeLayoutがもしかしたらどう言うふうに指定していいかわからないかも

func TimeToString(timeLayout string, datetime time.Time) string {
	str := datetime.Format(timeLayout)
	return str
}