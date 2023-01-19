package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"fmt"
	"github.com/cliclitv/go-clicli/db"
	"github.com/julienschmidt/httprouter"
)

func AddComment(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	req, _ := ioutil.ReadAll(r.Body)
	body := &db.Comment{}

	if err := json.Unmarshal(req, body); err != nil {
		sendMsg(w, 401, "参数解析失败")
		return
	}

	if _, err := db.AddComment(body.Rate, body.Content, body.Pid, body.Uid); err != nil {
		sendMsg(w, 500,fmt.Sprintf("%s", err))
		return
	} else {
		sendMsg(w, 500, "添加成功了")
	}

}

// func UpdateVideo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

// 	id := p.ByName("id")
// 	vid, _ := strconv.Atoi(id)
// 	req, _ := ioutil.ReadAll(r.Body)
// 	body := &def.Video{}

// 	if err := json.Unmarshal(req, body); err != nil {
// 		sendMsg(w, 401, "参数解析失败")
// 		return
// 	}

// 	if resp, err := db.UpdateVideo(vid, body.Oid, body.Title, body.Content, body.Pid, body.Uid); err != nil {
// 		sendMsg(w, 401, "数据库错误")
// 		return
// 	} else {
// 		sendVideoResponse(w, resp, 200)
// 	}

// }

func GetComments(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	pid, _ := strconv.Atoi(r.URL.Query().Get("pid"))
	uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	pageSize, _ := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if pageSize > 300 {
		sendMsg(w, 401, "pageSize太大了")
		return
	}
	resp, err := db.Getcomments(pid, uid, page, pageSize)
	if err != nil {
		sendMsg(w, 500,fmt.Sprintf("%s", err))
		return
	} else {
		res := &db.Comments{Comments: resp}
		sendCommentsResponse(w, res, 200)
	}
}
