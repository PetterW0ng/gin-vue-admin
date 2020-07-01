package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"time"
)

// @title    CreateTitUserTopicAnswer
// @description   create a TitUserTopicAnswer
// @param     userTopicAnswer               model.TitUserTopicAnswer
// @auth      weiqin
// @return    err             error

func CreateTitUserTopicAnswer(userTopicAnswer model.TitUserTopicAnswer) (err error) {
	err = global.GVA_DB.Create(&userTopicAnswer).Error
	return err
}

// @title    DeleteTitUserTopicAnswer
// @description   delete a TitUserTopicAnswer
// @auth      weiqin
// @param     userTopicAnswer               model.TitUserTopicAnswer
// @return                    error

func DeleteTitUserTopicAnswer(userTopicAnswer model.TitUserTopicAnswer) (err error) {
	err = global.GVA_DB.Delete(userTopicAnswer).Error
	return err
}

// @title    UpdateTitUserTopicAnswer
// @description   update a TitUserTopicAnswer
// @param     userTopicAnswer          *model.TitUserTopicAnswer
// @auth      weiqin
// @return                    error

func UpdateTitUserTopicAnswer(userTopicAnswer *model.TitUserTopicAnswer) (err error) {
	err = global.GVA_DB.Save(userTopicAnswer).Error
	return err
}

// @title    GetTitUserTopicAnswer
// @description   get the info of a TitUserTopicAnswer
// @auth     weiqin
// @param     id              uint
// @return                    error
// @return    TitUserTopicAnswer        TitUserTopicAnswer

func GetTitUserTopicAnswer(id uint) (err error, userTopicAnswer model.TitUserTopicAnswer) {
	err = global.GVA_DB.Where("id = ?", id).First(&userTopicAnswer).Error
	return
}

// @title    GetTitUserTopicAnswerInfoList
// @description   get TitUserTopicAnswer list by pagination, 分页获取用户列表
// @auth      weiqin
// @param     info            PageInfo
// @return                    error

func GetTitUserTopicAnswerInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB
	var userTopicAnswers []model.TitUserTopicAnswer
	err = db.Find(&userTopicAnswers).Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&userTopicAnswers).Error
	return err, userTopicAnswers, total
}

func batchAddUserTopicAnswer(userTopicAnswerRequest []model.TitUserTopicAnswer) (err error) {
	var buffer bytes.Buffer
	batchInsert := "INSERT INTO `tit_user_topic_answers` (`tit_user_id`,`tit_topic_id`,`topic_title`,`answer`,`score`, `order`, `topic_option_selected`, `topic_option_ids`, `batch_num`, `business_type`, `created_at`, `updated_at`) values "
	if _, err = buffer.WriteString(batchInsert); err != nil {
		return
	}
	db := global.GVA_DB
	now := time.Now().Format("2006-01-02 15:04:05")
	for i, e := range userTopicAnswerRequest {
		if i == len(userTopicAnswerRequest)-1 {
			buffer.WriteString(fmt.Sprintf("('%d','%d','%s','%s', '%d', '%d', '%s', '%s', '%d', '%d', '%s', '%s');", e.TitUserId, e.TitTopicId, e.TopicTitle, e.Answer, e.Score, e.Order, e.TopicOptionSelected, e.TopicOptionIds, e.BatchNum, e.BusinessType, now, now))
		} else {
			buffer.WriteString(fmt.Sprintf("('%d','%d','%s','%s', '%d', '%d', '%s', '%s', '%d', '%d', '%s', '%s'),", e.TitUserId, e.TitTopicId, e.TopicTitle, e.Answer, e.Score, e.Order, e.TopicOptionSelected, e.TopicOptionIds, e.BatchNum, e.BusinessType, now, now))
		}
	}
	err = db.Exec(buffer.String()).Error
	return
}

func SaveTitUserAnswer(userTopicAnswerRequest request.UserTopicAnswer, titUserId uint) (err error) {
	businessType := userTopicAnswerRequest.BusinessType
	_, topicOptions := QueryByBusinessType(businessType)
	_, titUser := GetTitUser(titUserId)

	// 查询出根据 IitUserId 获取新增的 batchNum
	var batchNum int
	if businessType == 3 {
		titUser.IndustryPerspectiveBatchNum += 1
		batchNum = titUser.IndustryPerspectiveBatchNum
	} else if businessType == 1 {
		titUser.JobInfoBatchNum += 1
		batchNum = titUser.JobInfoBatchNum
	} else {
		titUser.PerfessionKnowledgeBatchNum += 1
		batchNum = titUser.PerfessionKnowledgeBatchNum
	}
	// 根据用户提交的答题信息，解析出 TitUserTopicAnswer
	userTopicAnswers := userTopicAnswerRequest.Answers
	var userTopicAndAnswers = make([]model.TitUserTopicAnswer, len(userTopicAnswers))
	var topicOption model.TitTopic
	for i, item := range userTopicAnswers {
		topicId := item.TitTopicId
		selectedOptions := item.TitTopicOptionIds
		topicOptionIds, _ := json.Marshal(selectedOptions)
		for _, v := range topicOptions {
			if v.ID == uint(topicId) {
				topicOption = v
				break
			}
		}

		// 设置 options 选项 如：熟悉孤独症的治疗原则[option]熟悉孤独症的诊断及其评估
		var topicOptionSelected string
		var score int
		for i, v := range selectedOptions {
			for _, op := range topicOption.TitTopicOptions {
				if uint(v) == op.ID {
					if i == len(selectedOptions)-1 {
						topicOptionSelected += op.Title
					} else {
						topicOptionSelected += op.Title + "[option]"
					}
					score = op.Score
				}
			}
		}
		// topicTitle 加上 多选 和 选填 属性
		var multipleSelect string
		if topicOption.TopicType == 2 {
			multipleSelect = "【多选】"
		}
		topicTitle := multipleSelect + topicOption.Title
		userTopicAndAnswers[i] = model.TitUserTopicAnswer{TitUserId: int(titUserId), TitTopicId: topicId, TopicTitle: topicTitle, TopicOptionIds: string(topicOptionIds[1 : len(topicOptionIds)-1]), TopicOptionSelected: topicOptionSelected,
			Answer: item.Others, Score: score, BusinessType: businessType, Order: topicOption.Order, BatchNum: batchNum}
	}
	// 更新 TitUser 下 IndustryPerspectiveBatchNum、PerfessionKnowledgeBatchNum、JobInfoBatchNum 的 batchNum
	err = batchAddUserTopicAnswer(userTopicAndAnswers)
	if err != nil {
		global.GVA_LOG.Error("录入用户答题情况失败", err)
		return
	}
	// 更新 titUser 下关联的 basicId
	err = ModifyTitUser(&titUser)
	return err
}

func QueryTitUserTopicAnswer(userId uint, businessType, batchNum int) (err error, userTopicAnswers []model.TitUserTopicAnswer) {
	db := global.GVA_DB
	answer := model.TitUserTopicAnswer{BatchNum: batchNum, TitUserId: int(userId), BusinessType: businessType}
	err = db.Order("order").Where(&answer).Find(&userTopicAnswers).Error
	return
}

func GetTitTopicRelated(id int) (err error, titTopicRelated model.TitTopicRelated) {
	db := global.GVA_DB
	err = db.Where("id = ?", id).First(&titTopicRelated).Error
	return
}
