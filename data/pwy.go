//场景路径信息解析
package data

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)
type Pwy struct {
	Scene, Sect, Id,Func  string
	CoorNum             int
	CoorList                         string
}

func PwyDecode(file string) (pwys []Pwy) {
	scene := filepath.Base(filepath.Dir(filepath.Dir(file)))
	sect := filepath.Base(filepath.Dir(file))
	fp, _ := os.Open(file)
	defer fp.Close()
	recs := GetInt(fp)
	for j := 0; j < recs; j++ {
		var pwy Pwy
		pwy.Scene, pwy.Sect = scene, sect
		pwy.Id = GetStr(fp, GetInt(fp))
		pwy.Id = pwy.Id[:len(pwy.Id)-1]
		pwy.Func = GetStr(fp, GetInt(fp))
		pwy.Func = pwy.Func[:len(pwy.Func)-1]
		GetFlag(fp, 4)
		pwy.CoorNum=GetInt(fp)
		for k:=0;k<pwy.CoorNum;k++ {
			tmpCoor := fmt.Sprintf("%5.3f, %5.3f, %5.3f; ", GetFloat(fp), GetFloat(fp), GetFloat(fp))
			pwy.CoorList+=tmpCoor
		}
		pwys = append(pwys, pwy)
	}
	return
}
func PwyInsert(file string,ch chan string) {
	start:=time.Now().Format(Layout)
	pwys:=PwyDecode(file)
	if len(pwys)>0 {	
		insertSql := "insert into PathWay(scene,section,pwy_id,func,coor_num,coor_list) values "
		var vals []interface{}
		for _, pwy := range pwys {
			insertSql += "(?, ?, ?, ?, ?, ?),"
			vals = append(vals, 
				pwy.Scene,pwy.Sect,pwy.Id,pwy.Func,pwy.CoorNum,pwy.CoorList,
			)
		}
		insertSql = insertSql[:len(insertSql)-1]
		_,err:=Db.Exec(insertSql,vals...)
		if err != nil {
			ch<-fmt.Sprintf("文件名：%s；无法创建新记录，错误：%v\n",file,err)
		}
	}
    end:=time.Now().Format(Layout)
    ch<-fmt.Sprintf("文件：%v；插入启动时间：%v；插入结束时间：%v",file,start,end)
}
