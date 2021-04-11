//场景敌人信息解析
package data

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)
type Mst struct {
    Scene,Sect,Id,Model string
	InitCoor,Coor2,Coor3 string
	Fix1,MstNum int
    MstList,Attr string
}
func MstDecode(file string) (msts []Mst) {
	scene:=filepath.Base(filepath.Dir(filepath.Dir(file)))
    sect:=filepath.Base(filepath.Dir(file))
	fp,_:=os.Open(file)
	defer fp.Close()
	recs:=GetInt(fp)
	for j:=0; j<recs; j++ {
		var mst Mst
		mst.Scene,mst.Sect=scene,sect
		mst.Id=GetStr(fp,GetInt(fp))
		mst.Id=mst.Id[:len(mst.Id)-1]
		mst.Model=GetStr(fp,GetInt(fp))
		mst.Model=mst.Model[:len(mst.Model)-1]
		mst.InitCoor=fmt.Sprintf("%5.3f, %5.3f, %5.3f",GetFloat(fp),GetFloat(fp),GetFloat(fp))
		mst.Coor2=fmt.Sprintf("%5.3f, %5.3f, %5.3f",GetFloat(fp),GetFloat(fp),GetFloat(fp))
		mst.Fix1=GetInt(fp)
		mst.Coor3=fmt.Sprintf("%5.3f, %5.3f, %5.3f",GetFloat(fp),GetFloat(fp),GetFloat(fp))
		mst.MstNum=GetInt(fp)
		for k:=0;k<mst.MstNum;k++ {
			mst.MstList+=fmt.Sprintf("%d, ",GetInt(fp))
		}
		if mst.MstList!="" {
			mst.MstList=mst.MstList[:len(mst.MstList)-2]
		}
		GetFlag(fp,8)
		var (
			field,s1 string
			i1,i2 int
			f1 float32
		)
		field=GetStr(fp,GetInt(fp))
		i1,i2=GetInt(fp),GetInt(fp)
		if i2==1 {
			field=fmt.Sprintf("%s: %d, %d;\n",field,i1,i2)
			mst.Attr+=field
			for {
				offset,_:=fp.Seek(0,1)
				flag:=GetFlag(fp,8)
				if flag!="0100000001000000" {
					fp.Seek(offset,0)
				} else {
					break
				}
				field=GetStr(fp,GetInt(fp))
				if match,_:=regexp.MatchString("(GroupName|GroupPathWayID)$",field);match {
					s1,i2=GetStr(fp,GetInt(fp)),GetInt(fp)
					field=fmt.Sprintf("%s: %s, %d;\n",field,s1,i2)
					mst.Attr+=field
				} else if match,_=regexp.MatchString("(ObjectExpectTime|ObjectWaitTime|max_force|max_speed|max_range)$",field);match {
					f1,i2=GetFloat(fp),GetInt(fp)
					field=fmt.Sprintf("%s: %5.3f, %d;\n",field,f1,i2)
					mst.Attr+=field
				} else  {
					i1,i2=GetInt(fp),GetInt(fp)
					field=fmt.Sprintf("%s: %d, %d;\n",field,i1,i2)
					mst.Attr+=field
				}
			}
		} else if i2==0 {
			field=fmt.Sprintf("%s: %d, %d;\n",field,i1,i2)
			mst.Attr+=field
			GetFlag(fp,8)
		}
		for {
			offset,_:=fp.Seek(0,1)
			field=GetStr(fp,GetInt(fp))
			if match,_:=regexp.MatchString("^(MSTINFO|AAB_MST_ESSBEH)",field);!match {
				fp.Seek(offset,0)
				break
			}
			if match,_:=regexp.MatchString("max_pursuit_speed$",field);match {
				f1,i2=GetFloat(fp),GetInt(fp)
				field=fmt.Sprintf("%s: %5.3f, %d;\n",field,f1,i2)
				mst.Attr+=field
			} else if match,_=regexp.MatchString("script_func$",field);match {
				s1,i2=GetStr(fp,GetInt(fp)),GetInt(fp)
				if i2!=0 {
					s2:=GetStr(fp,GetInt(fp))
					field=fmt.Sprintf("%s: %s, %d, %s;\n",field,s1,i2,s2)
					GetFlag(fp,8)
				} else {
					field=fmt.Sprintf("%s: %s, %d;\n",field,s1,i2)
				}
				mst.Attr+=field
			} else {
				i1,i2=GetInt(fp),GetInt(fp)
				field=fmt.Sprintf("%s: %d, %d;\n",field,i1,i2)
				mst.Attr+=field
			}
		}
		msts=append(msts,mst)
	}
	return
}
func MstInsert(file string,ch chan string) {
	start:=time.Now().Format(Layout)
	msts:=MstDecode(file)
	if len(msts)>0 {	
		insertSql := `insert into MstInfo(
			scene,section,mst_id,model,init_coor,coor2,fix1,coor3,mst_num,mst_list,attr
			) values `
		var vals []interface{}
		for _, mst := range msts {
			insertSql += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"
			vals = append(vals, 
				mst.Scene,mst.Sect,mst.Id,mst.Model,mst.InitCoor,mst.Coor2,
				mst.Fix1,mst.Coor3,mst.MstNum,mst.MstList,mst.Attr,
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