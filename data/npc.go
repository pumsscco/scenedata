//场景NPC信息解析
package data

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)
type Npc struct {
	Scene, Sect, Id, Model, Func string
	InitCoor, Coor2,TwoInt              string
	UnkFloat                    float32
	Attr                         string
}

func NpcDecode(file string) (npcs []Npc) {
	scene := filepath.Base(filepath.Dir(filepath.Dir(file)))
	sect := filepath.Base(filepath.Dir(file))
	fp, _ := os.Open(file)
	defer fp.Close()
	recs := GetInt(fp)
	for j := 0; j < recs; j++ {
		var npc Npc
		npc.Scene, npc.Sect = scene, sect
		npc.Id = GetStr(fp, GetInt(fp))
		npc.Id = npc.Id[:len(npc.Id)-1]
		npc.Model = GetStr(fp, GetInt(fp))
		npc.Model = npc.Model[:len(npc.Model)-1]
		npc.Func = GetStr(fp, GetInt(fp))
		npc.Func = npc.Model[:len(npc.Model)-1]
		npc.InitCoor = fmt.Sprintf("%5.3f, %5.3f, %5.3f", GetFloat(fp), GetFloat(fp), GetFloat(fp))
		npc.Coor2 = fmt.Sprintf("%5.3f, %5.3f, %5.3f", GetFloat(fp), GetFloat(fp), GetFloat(fp))
		npc.TwoInt = fmt.Sprintf("%d, %d", GetInt(fp), GetInt(fp))
		npc.UnkFloat = GetFloat(fp)
		GetFlag(fp, 8)
		var (
			field, s1 string
			i1, i2    int
			f1        float32
		)
		field = GetStr(fp, GetInt(fp))
		i1, i2 = GetInt(fp), GetInt(fp)
		if i2 == 1 {
			field = fmt.Sprintf("%s: %d, %d;\n", field, i1, i2)
			npc.Attr += field
			for {
				offset, _ := fp.Seek(0, 1)
				flag := GetFlag(fp, 8)
				if flag != "0100000001000000" {
					fp.Seek(offset, 0)
				} else {
					break
				}
				field = GetStr(fp, GetInt(fp))
				if match, _ := regexp.MatchString("(GroupName|GroupPathWayID)$", field); match {
					s1, i2 = GetStr(fp, GetInt(fp)), GetInt(fp)
					field = fmt.Sprintf("%s: %s, %d;\n", field, s1, i2)
					npc.Attr += field
				} else if match, _ = regexp.MatchString("(ObjectExpectTime|ObjectWaitTime|max_force|max_speed|max_range)$", field); match {
					f1, i2 = GetFloat(fp), GetInt(fp)
					field = fmt.Sprintf("%s: %5.3f, %d;\n", field, f1, i2)
					npc.Attr += field
				} else {
					i1, i2 = GetInt(fp), GetInt(fp)
					field = fmt.Sprintf("%s: %d, %d;\n", field, i1, i2)
					npc.Attr += field
				}
			}
		} else if i2 == 0 {
			field = fmt.Sprintf("%s: %d, %d;\n", field, i1, i2)
			npc.Attr += field
			GetFlag(fp, 8)
		}
		for {
			offset, _ := fp.Seek(0, 1)
			field = GetStr(fp, GetInt(fp))
			if match, _ := regexp.MatchString("^NPCINFO", field); !match {
				fp.Seek(offset, 0)
				break
			}
			if match, _ := regexp.MatchString("Attr_defaultAct$", field); match {
				s1, i2 = GetStr(fp, GetInt(fp)), GetInt(fp)
				field = fmt.Sprintf("%s: %s, %d;\n", field, s1, i2)
				npc.Attr += field
			} else {
				i1, i2 = GetInt(fp), GetInt(fp)
				field = fmt.Sprintf("%s: %d, %d;\n", field, i1, i2)
				npc.Attr += field
			}
		}
		npcs = append(npcs, npc)
	}
	return
}
func NpcInsert(file string,ch chan string) {
	start:=time.Now().Format(Layout)
	npcs:=NpcDecode(file)
	if len(npcs)>0 {	
		insertSql := "insert into NPCInfo(scene,section,npc_id,model,func,init_coor,coor2,two_int,unk_float,attr) values "
		var vals []interface{}
		for _, npc := range npcs {
			insertSql += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"
			vals = append(vals, 
				npc.Scene,npc.Sect,npc.Id,npc.Model,npc.Func,
				npc.InitCoor,npc.Coor2,npc.TwoInt,npc.UnkFloat,npc.Attr,
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

