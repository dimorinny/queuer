package serializer

import (
	"time"

	"github.com/dimorinny/queuer/src/database"
)

type QueueSerializer struct {
	ID            int
	Title         string
	Description   string
	MaxPeoples    int
	Creator       *UserSerializer
	CurrentMember *MemberSerializer
	Created       time.Time
	IsActive      bool
}

func SerializeQueue(queue database.Queue) *QueueSerializer {
	return &QueueSerializer{
		ID:          queue.ID,
		Title:       queue.Title,
		Description: queue.Description,
		MaxPeoples:  queue.MaxPeoples,
		Creator:     SerializeUser(queue.Creator),
		Created:     queue.Created,
		IsActive:    queue.IsActive,
	}
}

func SerializeQueues(queues []database.Queue) []QueueSerializer {
	resultQueues := []QueueSerializer{}

	for _, val := range queues {
		resultQueues = append(resultQueues, QueueSerializer{
			ID:            val.ID,
			Title:         val.Title,
			Description:   val.Description,
			MaxPeoples:    val.MaxPeoples,
			Creator:       SerializeUser(val.Creator),
			Created:       val.Created,
			CurrentMember: SerializeMember(val.CurrentMember),
			IsActive:      val.IsActive,
		})
	}

	return resultQueues
}

type QueueDetailSerializer struct {
	QueueSerializer
	Members []MemberSerializer
}

func SerializeDetailQueue(queue database.Queue) *QueueDetailSerializer {
	members := []MemberSerializer{}

	for _, val := range queue.Members {
		members = append(members, *SerializeMember(val))
	}

	return &QueueDetailSerializer{
		QueueSerializer: QueueSerializer{
			ID:            queue.ID,
			Title:         queue.Title,
			Description:   queue.Description,
			MaxPeoples:    queue.MaxPeoples,
			Creator:       SerializeUser(queue.Creator),
			Created:       queue.Created,
			CurrentMember: SerializeMember(queue.CurrentMember),
			IsActive:      queue.IsActive,
		},
		Members: members,
	}
}

type UserSerializer struct {
	ID           int
	Email        string
	FirstName    string
	LastName     string
	IsSuperAdmin bool
}

func SerializeUser(user database.User) *UserSerializer {
	if user.ID != 0 {
		return &UserSerializer{
			ID:           user.ID,
			Email:        user.Email,
			FirstName:    user.FirstName,
			LastName:     user.LastName,
			IsSuperAdmin: user.IsSuperAdmin,
		}
	} else {
		return nil
	}
}

type MemberSerializer struct {
	ID               int
	SubscriptionTime time.Time
	User             *UserSerializer
}

func SerializeMember(member database.Member) *MemberSerializer {
	if member.ID != 0 {
		return &MemberSerializer{
			ID:               member.ID,
			SubscriptionTime: member.SubscriptionTime,
			User:             SerializeUser(member.User),
		}
	} else {
		return nil
	}
}
