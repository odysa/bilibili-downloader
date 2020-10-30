package main

import (
	"github.com/odysa/bilibili-downloader/download"
	"github.com/odysa/bilibili-downloader/model"
)

func main() {
	download.Download(`http://upos-sz-mirrorks3.bilivideo.com/upgcxcode/05/17/246341705/246341705-1-32.flv?e=ig8euxZM2rNcNbNHhwdVhoMgnWdVhwdEto8g5X10ugNcXBlqNxHxNEVE5XREto8KqJZHUa6m5J0SqE85tZvEuENvNo8g2ENvNo8i8o859r1qXg8xNEVE5XREto8GuFGv2U7SuxI72X6fTr859r1qXg8gNEVE5XREto8z5JZC2X2gkX5L5F1eTX1jkXlsTXHeux_f2o859IB_&uipk=5&nbs=1&deadline=1604052261&gen=playurl&os=ks3bv&oi=1879056101&trid=fd2196c40b3c4c3fb3a76d713be3bd64u&platform=pc&upsig=bd234045da6071b7e8b6377dbb8620c9&uparams=e,uipk,nbs,deadline,gen,os,oi,trid,platform&mid=0&orderid=0,3&agrr=1&logo=80000000`, model.VideoPart{})
}
