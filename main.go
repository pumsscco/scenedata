package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	. "scenedata/data"
)
func main() {	
	var files []string
	var gobFiles,_=filepath.Glob("/home/pluto/tmp/scenedata/*/*/GameObjs.gob")
	files=append(files,gobFiles...)
	var mstFiles,_=filepath.Glob("/home/pluto/tmp/scenedata/*/*/mstInfo.mst")
	files=append(files,mstFiles...)
	var npcFiles, _ = filepath.Glob("/home/pluto/tmp/scenedata/*/*/npcInfo.npc")
	files=append(files,npcFiles...)
	var pwyFiles, _ = filepath.Glob("/home/pluto/tmp/scenedata/*/*/pathway.pwy")
	files=append(files,pwyFiles...)
	var ch chan string=make(chan string,len(files))
    for _,f:=range files {
		if match,_:=regexp.MatchString("gob$",f); match {
			go GobInsert(f,ch)
		} else if match,_=regexp.MatchString("mst$",f); match {
			go MstInsert(f,ch)
		} else if match,_=regexp.MatchString("npc$",f); match {
			go NpcInsert(f,ch)
		} else if match,_=regexp.MatchString("pwy$",f); match {
			go PwyInsert(f,ch)
		}
	}
	for range files {
		msg:=<-ch
        fmt.Println(msg)
    }
}