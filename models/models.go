package models

var ISDEBUG = false

type Config struct {
	EjudgeConAddress string `yaml:"ejudge_con_address"`
	DbPath           string `yaml:"db_path"`
	LogPath          string `yaml:"log_path"`
	Port             string `yaml:"port"`
	IsDebug          bool   `yaml:"is_debug"`
	AdminLogin       string `yaml:"admin_login"`
	AdminPassword    string `yaml:"admin_password"`
}

type User struct {
	EjId           int    `gorm:"primaryKey;autoIncrement:false" json:"ejId"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	ProfilePicture string `json:"profilePicture"`
	Status         string `json:"status"`
}

// Group по факту это отдельная система. например, группа лкш2023, группа контестов 10и
type Group struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type BasicContest struct {
	Name           string `json:"name"`
	Url            string `json:"url"`
	ContestPicture string `json:"contestPicture"`
	Comment        string `json:"comment"`
	StatementsUrl  string `json:"statementsUrl"`
}

type Contest struct {
	ID int `gorm:"primaryKey"` // делаем сами
	BasicContest
}

type Admin struct {
	ID          int `gorm:"primaryKey"`
	Description string
}

// тут тема какая. Если модер в группе хост, то эта таблица для него неактуальна. А если нет, то редачить может только те, где роль админ. Все.

type ModeratorContestId struct {
	ModeratorContest string `gorm:"primaryKey;autoIncrement:false"`
	Role             Role   `gorm:"foreignKey:AdminRefer"`
}

type GroupContestId struct {
	GroupContest string `gorm:"primaryKey;autoIncrement:false"`
	Belongs      bool
}

type UserAndContest struct {
	UserId    int  `json:"userId"`
	ContestId int  `json:"contestId"`
	Role      Role `json:"role"`
}

type SimpleModerator struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Moderators struct {
	ID int `gorm:"primaryKey"`
	SimpleModerator
}

type ModeratorGroup struct {
	ModeratorGroupId string `gorm:"primaryKey;autoIncrement:false"`
	IsHost           bool
}

type Role int

const (
	NoAdmin Role = iota + 1
	YesAdmin
)

func (r Role) String() string {
	return []string{"не администратор", "администратор"}[r-1]
}
