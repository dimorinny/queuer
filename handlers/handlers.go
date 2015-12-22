package handlers

import (
	"strconv"

	"time"

	"github.com/dimorinny/queuer/auth"
	"github.com/dimorinny/queuer/database"
	"github.com/dimorinny/queuer/response"
	"github.com/dimorinny/queuer/serializer"
	"github.com/labstack/echo"
)

func Queues(c *echo.Context) error {
	queues := []database.Queue{}
	database.Db.Not("is_deleted", true).Preload("Creator").Preload("CurrentMember.User").Order("Created").Find(&queues)
	response.QueuesResponse(c, serializer.SerializeQueues(queues))
	return nil
}

func MyQueues(c *echo.Context) error {
	//	user := c.Get(auth.Config.IdentityKey).(database.User)
	//	SELECT DISTINCT queues.* FROM members
	//	LEFT JOIN queues ON queue_id = queues.id
	//	WHERE user_id = 1

	//	members := []database.Member{}
	//	indexes := []int{}
	//	database.Db.Select("id").Where(&database.Member{UserID: user.ID}).Find(&members)
	//
	//	for _, val := range members {
	//		indexes = append(indexes, val.ID)
	//	}
	//
	//	queues := []database.Queue{}
	//	database.Db.Not("is_deleted", true).Preload("Creator").Preload("CurrentMember.User").Order("Created").Where("")Find(&queues)

	//	response.QueuesResponse(c, queues)
	return nil
}

func Queue(c *echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	queue := database.Queue{}
	err := database.Db.Not("is_deleted", true).Preload("Creator").Preload("Members.User").First(&queue, id).Error

	if err != nil {
		response.QueueNotFoundHandler(c)
		return nil
	}

	if queue.CurrentMemberID != 0 {
		member := database.Member{}

		if database.Db.First(&member, queue.CurrentMemberID).Error != nil {
			queue.CurrentMember = member
		}
	}

	response.QueueResponse(c, serializer.SerializeDetailQueue(queue))
	return nil
}

func CreateQueue(c *echo.Context) error {
	title := c.Form("title")
	description := c.Form("description")
	maxPeoples, err := strconv.Atoi(c.Form("maxPeoples"))

	user := c.Get(auth.Config.IdentityKey).(database.User)

	if title == "" || description == "" || err != nil {
		response.QueueParamsCreateError(c)
		return nil
	}

	queue := database.Queue{
		Title:       title,
		Description: description,
		MaxPeoples:  maxPeoples,
		CreatorID:   user.ID,
		Created:     time.Now(),
		IsActive:    true,
		IsDeleted:   false,
	}

	if database.Db.Create(&queue).Error != nil {
		response.QueueCreateError(c)
		return nil
	}

	queue.Creator = user
	response.QueueResponse(c, serializer.SerializeQueue(queue))
	return nil
}

func NextMember(c *echo.Context) error {
	user := c.Get(auth.Config.IdentityKey).(database.User)
	id, _ := strconv.Atoi(c.Param("id"))

	// Crunch, because Update nobody return error
	queue := database.Queue{}
	if database.Db.Not("is_deleted", true).First(&queue, id).Select("id").Error != nil {
		response.QueueNotFoundHandler(c)
		return nil
	}

	if user.ID != queue.CreatorID && !user.IsSuperAdmin {
//		response.QueueRemoveNotPermitted(c)
		return nil
	}

	return nil
}

func DeleteMember(c *echo.Context) error {
	//	queueID, _ := strconv.Atoi(c.Param("queueID"))
	//	memberID, _ := strconv.Atoi(c.Param("memberID"))
	return nil
}

func DeleteQueue(c *echo.Context) error {
	user := c.Get(auth.Config.IdentityKey).(database.User)
	id, _ := strconv.Atoi(c.Param("id"))

	// Crunch, because Update nobody return error
	queue := database.Queue{}
	if database.Db.Not("is_deleted", true).First(&queue, id).Select("id").Error != nil {
		response.QueueNotFoundHandler(c)
		return nil
	}

	if user.ID != queue.CreatorID && !user.IsSuperAdmin {
		response.QueueRemoveNotPermitted(c)
		return nil
	}

	if database.Db.Model(database.Queue{ID: id}).UpdateColumn("is_deleted", true).Error != nil {
		response.QueueRemoveError(c)
		return nil
	}

	response.EmptyResponseHandler(c)
	return nil
}

func ActiveQueue(c *echo.Context) error {
	//	id, _ := strconv.Atoi(c.Param("id"))
	return nil
}

func JoinQueue(c *echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := c.Get(auth.Config.IdentityKey).(database.User)

	// Crunch
	queue := database.Queue{}
	if database.Db.Not("is_deleted", true).First(&queue, id).Select("id").Error != nil {
		response.QueueNotFoundHandler(c)
		return nil
	}

	member := database.Member{
		SubscriptionTime: time.Now(),
		QueueID:          id,
		UserID:           user.ID,
	}

	if err := database.Db.Create(&member).Error; err != nil {
		response.AlreadyInQueueError(c)
		return nil
	}

	response.EmptyResponseHandler(c)
	return nil
}

func LeaveQueue(c *echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := c.Get(auth.Config.IdentityKey).(database.User)

	// Crunch
	queue := database.Queue{}
	if database.Db.Not("is_deleted", true).First(&queue, id).Select("id").Error != nil {
		response.QueueNotFoundHandler(c)
		return nil
	}

	member := database.Member{}
	if err := database.Db.Where(database.Member{UserID: user.ID, QueueID: id}).First(&member).Select("id").Error; err != nil {
		response.NotFoundInQueueError(c)
		return err
	}

	// Check is current user

	if err := database.Db.Where("queue_id = ? AND user_id = ?", id, user.ID).Delete(&database.Member{}).Error; err != nil {
		return err
	}

	response.EmptyResponseHandler(c)
	return nil
}
