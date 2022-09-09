package repository

import (
	"strings"
	"time"
	"wendaxitong/user/internal/service"
	"wendaxitong/user/pkg/codeMsg"
)

type TopicInfo struct {
	Id          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserName    string    `gorm:"not null;unique" json:"user_name"`
	Content     string    `json:"content"`
	ContentType string    `json:"content_type"`
	ParentId    uint64    `gorm:"not null" json:"parent_id"`
	RootId      uint64    `gorm:"not null" json:"root_id"`
	NumLikes    uint64    `json:"num_likes"`
	LikesNames  string    `json:"likes_name"`
	Remarks     string    `json:"remarks"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

/**
***
*** 用户互动****************************************************************************
***
 */

// CreateTopic 创建话题
func (t *TopicInfo) CreateTopic(request *service.TopicRequest) codeMsg.CodeMessage {
	var u UserInfo
	isExist, u := u.CheckUserExist(request.TopicInfo.UserName)
	if !isExist {
		return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorUserNotExist, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorUserNotExist)} // 该用户不存在
	}
	var topic TopicInfo
	topic = TopicInfo{
		UserName:    request.TopicInfo.UserName,
		Content:     request.TopicInfo.Content,
		ContentType: request.TopicInfo.ContentType,
		ParentId:    0,
		RootId:      0,
	}

	tx := DB.Begin() //开启事务
	err := tx.Create(&topic).Error
	if err != nil {
		tx.Rollback() // 遇到错误时回滚事务

		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "创建话题失败"} // 创建话题失败
	}
	err = UpdateValueByName("user_name", request.TopicInfo.UserName, UserInfo{}, "num_topic", u.NumTopic+1)
	if err != nil {
		tx.Rollback()                                                                   // 遇到错误时回滚事务
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "创建话题失败"} // 创建话题失败
	}
	tx.Commit() // 提交事务
	return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "创建成功"}
}

// DeleteTopic 删除话题
func (t *TopicInfo) DeleteTopic(request *service.TopicRequest) codeMsg.CodeMessage {
	var u UserInfo

	err := DB.Where("id = ?", request.TopicInfo.Id).First(t).Error
	if err != nil {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "查询失败：" + err.Error()}
	}
	if t.UserName != request.TopicInfo.UserName {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "权限错误"}
	}

	err = DB.Where("user_name = ?", request.TopicInfo.UserName).First(&u).Error
	if err != nil {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "查询失败：" + err.Error()}
	}

	var parentId uint64 = t.ParentId

	if parentId != 0 {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "id错误，将要删除的不是话题"}
	}

	tx := DB.Begin() //开启事务

	// 删除话题的所有评论
	var allChildren []TopicInfo
	err = DB.Where("root_id = ?", t.Id).Debug().Find(&allChildren).Error // 查找所有父评论为当前评论的子评论
	if err != nil {
		tx.Rollback() // 遇到错误时回滚事务
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "查询失败：" + err.Error()}
	}
	for i := 0; i < len(allChildren); i++ {
		err = tx.Unscoped().Delete(&allChildren[i]).Error
		if err != nil {
			tx.Rollback() // 遇到错误时回滚事务
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "删除失败：" + err.Error()}
		}
	}

	err = UpdateValueByName("user_name", request.TopicInfo.UserName, UserInfo{}, "num_topic", u.NumTopic-1)
	if err != nil {
		tx.Rollback() // 遇到错误时回滚事务
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "删除失败：" + err.Error()}
	}

	err = tx.Unscoped().Delete(t).Error
	if err != nil {
		tx.Rollback() // 遇到错误时回滚事务
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "删除失败：" + err.Error()}
	}
	tx.Commit() // 提交事务
	return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "删除成功"}
}

// CommentTopic 评论话题或者回复评论
func (t *TopicInfo) CommentTopic(request *service.TopicRequest) codeMsg.CodeMessage {
	var u UserInfo
	isExist, u := u.CheckUserExist(request.TopicInfo.UserName)
	if !isExist {
		return codeMsg.CodeMessage{StatusCode: codeMsg.ErrorUserNotExist, StatusMessage: codeMsg.GetErrorMsg(codeMsg.ErrorUserNotExist)} // 该用户不存在
	}

	var count int64
	err := DB.Where("id = ?", request.TopicInfo.ParentId).First(&TopicInfo{}).Count(&count).Error
	if err != nil {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "查询失败：" + err.Error()}
	}
	if count == 0 {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "ParentId错误，该条评论不存在"}
	}
	err = DB.Where("id = ?", request.TopicInfo.RootId).First(&TopicInfo{}).Count(&count).Error
	if err != nil {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "RootId错误，该话题不存在"}
	}
	if count == 0 {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "id错误，该条评论不存在"}
	}

	var topic TopicInfo
	topic = TopicInfo{
		UserName:    request.TopicInfo.UserName,
		Content:     request.TopicInfo.Content,
		ContentType: request.TopicInfo.ContentType,
		ParentId:    request.TopicInfo.ParentId,
		RootId:      request.TopicInfo.RootId,
	}

	tx := DB.Begin() //开启事务
	err = tx.Create(&topic).Error
	if err != nil {
		tx.Rollback() // 遇到错误时回滚事务

		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "评论失败"} // 创建话题失败
	}
	tx.Commit() // 提交事务
	return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "评论成功"}
}

// DeleteComment 删除评论
func (t *TopicInfo) DeleteComment(request *service.TopicRequest) codeMsg.CodeMessage {
	var topic TopicInfo
	var count int64
	err := DB.Where("id = ?", request.TopicInfo.Id).First(&topic).Count(&count).Error
	if err != nil {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "查询失败：" + err.Error()}
	}

	if count == 0 {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "id错误，该条评论不存在"}
	}

	if topic.UserName != request.TopicInfo.UserName {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "权限错误"}
	}

	if topic.RootId == 0 {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "id错误，该内容是话题"}
	}

	tx := DB.Begin() //开启事务

	// 删除话题的所有子评论
	var allChildren []TopicInfo
	err = DB.Where("parent_id = ?", request.TopicInfo.Id).Debug().Find(&allChildren).Error // 查找所有父评论为当前评论的子评论
	if err != nil {
		tx.Rollback() // 遇到错误时回滚事务
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "查询失败：" + err.Error()}
	}
	for i := 0; i < len(allChildren); i++ {
		request1 := &service.TopicRequest{
			TopicInfo: &service.TopicModel{
				Id:       allChildren[i].Id,
				UserName: allChildren[i].UserName,
			},
		}
		msg := t.DeleteComment(request1)
		if msg.StatusCode == codeMsg.Failed {
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "删除失败：" + err.Error()}
		}
		err = tx.Unscoped().Delete(&allChildren[i]).Error
		if err != nil {
			tx.Rollback() // 遇到错误时回滚事务
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "删除失败：" + err.Error()}
		}
	}

	err = tx.Unscoped().Delete(topic).Error
	if err != nil {
		tx.Rollback() // 遇到错误时回滚事务
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "删除失败：" + err.Error()}
	}
	tx.Commit() // 提交事务
	return codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "删除成功"}
}

// LikesTopicOrComment 点赞/取消点赞话题或评论
func (t *TopicInfo) LikesTopicOrComment(request *service.TopicRequest) codeMsg.CodeMessage {
	var topic TopicInfo
	var count int64
	err := DB.Where("id = ?", request.TopicInfo.Id).First(&topic).Count(&count).Error
	if err != nil {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "查询失败：" + err.Error()}
	}
	if count == 0 {
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "id错误，该条话题或评论不存在"}
	}

	// 获取点赞用户姓名
	namesStr := strings.Trim(topic.LikesNames, ",")
	namesList := strings.Split(namesStr, ",")

	if request.TopicInfo.Remarks == "1" { // 点赞
		for i := 0; i < len(namesList); i++ {
			if namesList[i] == request.TopicInfo.UserName {
				return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "已赞"}
			}
		}
		tx := DB.Begin() //开启事务
		err = DB.Model(&TopicInfo{}).Where("id = ?", request.TopicInfo.Id).Update("num_likes", topic.NumLikes+1).Error
		if err != nil {
			tx.Rollback() // 遇到错误时回滚事务
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "点赞失败"}
		}
		newLikesNameStr := namesStr + "," + request.TopicInfo.UserName
		newLikesNameStr = strings.Trim(newLikesNameStr, ",")
		err = DB.Model(&TopicInfo{}).Where("id = ?", request.TopicInfo.Id).Update("likes_names", newLikesNameStr).Error
		if err != nil {
			tx.Rollback() // 遇到错误时回滚事务
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "点赞失败"}
		}
		tx.Commit() // 提交事务
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "点赞成功"}

	} else { // 取消点赞
		isLike := false
		var likesNameMap = make(map[string]interface{}, 10)
		for i := 0; i < len(namesList); i++ {
			likesNameMap[namesList[i]] = ""
			if namesList[i] == request.TopicInfo.UserName {
				isLike = true
			}
		}
		if !isLike {
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "未点赞"}
		}
		tx := DB.Begin() //开启事务
		err = DB.Model(&TopicInfo{}).Where("id = ?", request.TopicInfo.Id).Update("num_likes", topic.NumLikes-1).Error
		if err != nil {
			tx.Rollback() // 遇到错误时回滚事务
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "点赞失败"}
		}

		delete(likesNameMap, request.TopicInfo.UserName)
		var newLikesNameStr string
		for k, _ := range likesNameMap {
			newLikesNameStr = newLikesNameStr + "," + k
		}
		newLikesNameStr = strings.Trim(newLikesNameStr, ",")

		err = DB.Model(&TopicInfo{}).Where("id = ?", request.TopicInfo.Id).Update("likes_names", newLikesNameStr).Error
		if err != nil {
			tx.Rollback() // 遇到错误时回滚事务
			return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "取消点赞失败"}
		}
		tx.Commit() // 提交事务
		return codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "取消点赞成功"}
	}
}

// QueryTopicInfo 查询用户所有话题信息
func (t *TopicInfo) QueryTopicInfo(request *service.TopicRequest) ([]TopicInfo, codeMsg.CodeMessage) {
	var topics []TopicInfo
	err := DB.Where("user_name = ? and parent_id = ?", request.TopicInfo.UserName, 0).Debug().Find(&topics).Error
	if err != nil {
		return nil, codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "查询失败：" + err.Error()}
	}
	return topics, codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "查询成功"}

}

// QueryCommentInfo 查询用户所有评论（回复）信息

func (t *TopicInfo) QueryCommentInfo(request *service.TopicRequest) ([]TopicInfo, codeMsg.CodeMessage) {
	var topics []TopicInfo
	err := DB.Where("user_name = ? and parent_id <> ?", request.TopicInfo.UserName, 0).Debug().Find(&topics).Error
	if err != nil {
		return nil, codeMsg.CodeMessage{StatusCode: codeMsg.Failed, StatusMessage: "查询失败：" + err.Error()}
	}
	return topics, codeMsg.CodeMessage{StatusCode: codeMsg.SUCCESS, StatusMessage: "查询成功"}
}
