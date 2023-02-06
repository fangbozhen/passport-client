package passport_client

import "time"

type User struct {
	Name      string          `json:"Name"`
	ZjuId     string          `json:"ZjuId"`
	LoginType string          `json:"LoginType"` // @see enum LoginType
	QscUser   *UserProfileQsc `json:"QscUser,omitempty"`
}

type UserProfileQsc struct {
	ZjuId      string    `json:"zjuid" bson:"ZjuId"`
	QscId      string    `json:"qscid" bson:"QscId"`
	Password   string    `json:"-" bson:"Password"`
	Name       string    `json:"name" bson:"Name"`
	Gender     string    `json:"gender" bson:"Gender"`
	Department string    `json:"department" bson:"Department"`
	Direction  string    `json:"direction" bson:"Direction"`
	Position   string    `json:"position" bson:"Position"`
	Status     string    `json:"status" bson:"Status"`
	Phone      string    `json:"phone" bson:"Phone"`
	Email      string    `json:"email" bson:"Email"`
	Note       string    `json:"note" bson:"Note"`
	Birthday   time.Time `json:"birthday,omitempty" bson:"Birthday"`
	JoinTime   time.Time `json:"jointime,omitempty" bson:"JoinTime"`
}
