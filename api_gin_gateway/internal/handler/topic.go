package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wendaxitong/api_gin_gateway/internal/service"
	"wendaxitong/api_gin_gateway/pkg/codeMsg"
	"wendaxitong/api_gin_gateway/pkg/util"
)

type Topic struct {
	Id          uint64 `json:"id"`
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
	ParentId    uint64 `json:"parent_id"`
	RootId      uint64 `json:"root_id"`
	Choose      string `json:"choose"`
}

// CreateTopic 创建话题
func CreateTopic(c *gin.Context) {
	var topic Topic
	userName := c.MustGet("userName").(string)
	if userName == "" {
		c.JSON(http.StatusOK, "userName为空，参数传递错误")
		c.Abort()
		return
	}
	err := c.ShouldBind(&topic)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: err.Error(),
			Data:    "null",
		})
		c.Abort()
		return
	}
	fmt.Println(topic)
	if topic.Content == "" || topic.ContentType == "" {
		c.JSON(http.StatusOK, gin.H{"error": "内容和内容类型均不能为空"})
		c.Abort()
		return
	}

	request := &service.TopicRequest{
		TopicInfo: &service.TopicModel{
			UserName:    userName,
			Content:     topic.Content,
			ContentType: topic.ContentType,
		},
	}
	response, err := GrpcUerServiceClient.CreateTopic(context.Background(), request)
	if err != nil {
		fmt.Println("GrpcUerServiceClient.CreateTopic Error:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage,
		Data:    "",
	})

}

// DeleteTopic 删除某个话题
func DeleteTopic(c *gin.Context) {
	var topic Topic
	userName := c.MustGet("userName").(string)
	if userName == "" {
		c.JSON(http.StatusOK, "userName为空，参数传递错误")
		c.Abort()
		return
	}
	err := c.ShouldBind(&topic)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: err.Error(),
			Data:    "null",
		})
		c.Abort()
		return
	}

	request := &service.TopicRequest{
		TopicInfo: &service.TopicModel{
			Id:       topic.Id,
			UserName: userName,
		},
	}
	response, err := GrpcUerServiceClient.DeleteTopic(context.Background(), request)
	if err != nil {
		fmt.Println("GrpcUerServiceClient.DeleteTopic Error:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage,
		Data:    "",
	})
}

// CommentTopic 评论话题或者回复评论
func CommentTopic(c *gin.Context) {
	var topic Topic
	userName := c.MustGet("userName").(string)
	if userName == "" {
		c.JSON(http.StatusOK, "userName为空，参数传递错误")
		c.Abort()
		return
	}
	err := c.ShouldBind(&topic)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: err.Error(),
			Data:    "null",
		})
		c.Abort()
		return
	}
	if topic.ParentId == 0 || topic.RootId == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "RootId和ParentId均不能为0"})
		c.Abort()
		return
	}
	if topic.Content == "" || topic.ContentType == "" {
		c.JSON(http.StatusOK, gin.H{"error": "内容和内容类型均不能为空"})
		c.Abort()
		return
	}

	request := &service.TopicRequest{
		TopicInfo: &service.TopicModel{
			UserName:    userName,
			RootId:      topic.RootId,
			ParentId:    topic.ParentId,
			Content:     topic.Content,
			ContentType: topic.ContentType,
		},
	}
	response, err := GrpcUerServiceClient.CommentTopic(context.Background(), request)
	if err != nil {
		fmt.Println("GrpcUerServiceClient.CommentTopic Error:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage,
		Data:    "",
	})

}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	var topic Topic
	userName := c.MustGet("userName").(string)
	if userName == "" {
		c.JSON(http.StatusOK, "userName为空，参数传递错误")
		c.Abort()
		return
	}
	err := c.ShouldBind(&topic)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: err.Error(),
			Data:    "null",
		})
		c.Abort()
		return
	}

	request := &service.TopicRequest{
		TopicInfo: &service.TopicModel{
			Id:       topic.Id,
			UserName: userName,
		},
	}
	response, err := GrpcUerServiceClient.DeleteComment(context.Background(), request)
	if err != nil {
		fmt.Println("GrpcUerServiceClient.DeleteComment Error:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage,
		Data:    "",
	})
}

// LikesTopicOrComment 点赞/取消点赞话题或评论
func LikesTopicOrComment(c *gin.Context) {
	var topic Topic
	userName := c.MustGet("userName").(string)
	if userName == "" {
		c.JSON(http.StatusOK, "userName为空，参数传递错误")
		c.Abort()
		return
	}
	err := c.ShouldBind(&topic)
	if err != nil {
		c.JSON(http.StatusOK, util.JsonData{
			Code:    codeMsg.Failed,
			Message: err.Error(),
			Data:    "null",
		})
		c.Abort()
		return
	}

	request := &service.TopicRequest{
		TopicInfo: &service.TopicModel{
			Id:       topic.Id,
			UserName: userName,
			Remarks:  topic.Choose,
		},
	}
	response, err := GrpcUerServiceClient.LikesTopicOrComment(context.Background(), request)
	if err != nil {
		fmt.Println("GrpcUerServiceClient.LikesTopicOrComment Error:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage,
		Data:    "",
	})
}

// QueryTopicInfo 查询用户所有话题信息
func QueryTopicInfo(c *gin.Context) {
	var topic service.TopicRequest
	err := c.ShouldBind(&topic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		c.Abort()
		return
	}

	request := &topic
	response, err := GrpcUerServiceClient.QueryTopicInfo(context.Background(), request)
	if err != nil {
		fmt.Println("GrpcUerServiceClient.QueryTopicInfo Error:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage,
		Data:    response.TopicInfo,
	})

}

// QueryCommentInfo 查询用户所有评论（回复）信息
func QueryCommentInfo(c *gin.Context) {
	var topic service.TopicRequest
	err := c.ShouldBind(&topic)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		c.Abort()
		return
	}

	request := &topic
	response, err := GrpcUerServiceClient.QueryCommentInfo(context.Background(), request)
	if err != nil {
		fmt.Println("GrpcUerServiceClient.QueryCommentInfo Error:", err.Error())
		c.JSON(http.StatusOK, err.Error())
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, util.JsonData{
		Code:    response.StatusCode,
		Message: response.StatusMessage,
		Data:    response.TopicInfo,
	})

}
