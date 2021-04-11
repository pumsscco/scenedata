//场景物体信息再次解析
package data

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Gob struct {
    Scene,Sect,Id,Path,Model,Texture string
	InitCoor,Coor2,TripleInt,Func string
	Kind int
	UnkFloat float32
    Attr string
}
func GobDecode(file string) (gobs []Gob) {
	scene:=filepath.Base(filepath.Dir(filepath.Dir(file)))
    sect:=filepath.Base(filepath.Dir(file))
	fp,_:=os.Open(file)
	defer fp.Close()
	recs:=GetInt(fp)
	var kinds []int
    for i:=0;i<recs;i++ {
        kinds=append(kinds,GetInt(fp))
    }
	for j:=0; j<recs; j++ {
		var gob Gob
		gob.Scene,gob.Sect=scene,sect
		gob.Kind=kinds[j]
		gob.Id=GetStr(fp,GetInt(fp))
		gob.Id=gob.Id[:len(gob.Id)-1]
		gob.Path=GetStr(fp,GetInt(fp))
		gob.Path=gob.Path[:len(gob.Path)-1]
		gob.Model=GetStr(fp,GetInt(fp))
		gob.Model=gob.Model[:len(gob.Model)-1]
		gob.Texture=GetStr(fp,GetInt(fp))
		if gob.Texture!="" {
			gob.Texture=gob.Texture[:len(gob.Texture)-1]
		}
		gob.InitCoor=fmt.Sprintf("%5.3f, %5.3f, %5.3f",GetFloat(fp),GetFloat(fp),GetFloat(fp))
		gob.Coor2=fmt.Sprintf("%5.3f, %5.3f, %5.3f",GetFloat(fp),GetFloat(fp),GetFloat(fp))
		gob.Func=GetStr(fp,GetInt(fp))
		gob.Func=gob.Func[:len(gob.Func)-1]
		gob.TripleInt=fmt.Sprintf("%d, %d, %d",GetInt(fp),GetInt(fp),GetInt(fp))
		gob.UnkFloat=GetFloat(fp)
		GetFlag(fp,8)
		var (
			field,s1 string
			i1,i2 int
			f1 float32
		)
		field=GetStr(fp,GetInt(fp))
		i1,i2=GetInt(fp),GetInt(fp)
		field=fmt.Sprintf("%s: %d, %d;\n",field,i1,i2)
		gob.Attr+=field
		for {
			offset,_:=fp.Seek(0,1)
			field=GetStr(fp,GetInt(fp))
			if match, _ := regexp.MatchString("^(PAL4|CSTR)", field); !match {
				fp.Seek(offset,0)
				break
			} else if match,_=regexp.MatchString("(name|func|launchEff)$",field); match {
				s1,i2=GetStr(fp,GetInt(fp)),GetInt(fp)
				field=fmt.Sprintf("%s: %s, %d;\n",field,s1,i2)
				gob.Attr+=field
			} else if match,_=regexp.MatchString("(time|scale|X|Y|Z|amplitude|frequency|maxval|minval|effDuration|effTime|effHurt|tMax|tMin|thrust|directDist|suction)$",field); match {
				f1,i2=GetFloat(fp),GetInt(fp)
				field=fmt.Sprintf("%s: %5.3f, %d;\n",field,f1,i2)
				gob.Attr+=field
			} else  {
				i1,i2=GetInt(fp),GetInt(fp)
				field=fmt.Sprintf("%s: %d, %d;\n",field,i1,i2)
				gob.Attr+=field
			}
		}
		gobs=append(gobs,gob)
	}
	return
}
func GobInsert(file string,ch chan string) {
	start:=time.Now().Format(Layout)
	gobs:=GobDecode(file)
	if len(gobs)>0 {	
		insertSql := `insert into GameObject(
			scene,section,kind,gob_id,path,model,texture,init_coor,coor2,func,triple_int,unk_float,attr
			) values `
		var vals []interface{}
		for _, gob := range gobs {
			insertSql += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"
			vals = append(vals, 
				gob.Scene,gob.Sect,gob.Kind,gob.Id,gob.Path,gob.Model,gob.Texture,
				gob.InitCoor,gob.Coor2,gob.Func,gob.TripleInt,gob.UnkFloat,gob.Attr,
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
