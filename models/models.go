package models

import "gorm.io/gorm"

type Fact struct {
	gorm.Model
	Question string `json:"question" gorm:"text;not null;default:null`
	Answer   string `json:"answer" gorm:"text;not null;default:null`
}

type Post struct {
	gorm.Model
	Image     string `json:"image" gorm:"text;default:null`
	Thumbnail string `json:"thumbnail" gorm:"text;not null;default:null`
	Slug      string `json:"slug" gorm:"text;default:null`
	Title     string `json:"title" gorm:"text;not null;default:null`
	Subtitle  string `json:"subtitle" gorm:"text;not null;default:null`
	Content   string `json:"content" gorm:"text;not null;default:null`
	Author    int    `json:"author" gorm:"number;not null;default:null`
	Rate      int    `json:"rate" gorm:"number;not null;default:null`
}

type Comment struct {
	gorm.Model
	PostID  int    `json:"postid" gorm:"number;default:null`
	Content string `json:"content" gorm:"text;not null;default:null`
	Author  int    `json:"author" gorm:"number;not null;default:null`
	Likes   int    `json:"likes" gorm:"number;not null;default:null`
	Image   string `json:"image" gorm:"text;default:null`
}

/*
	class NotificationTypes extends Enum {
	    public const COMMENT_RECEIVED = 1;
	    public const POST_UNLOCKED = 2;
	    public const PAYMENT_RECEIVED = 3;
	    public const LIKE_RECEIVED = 4;
	    public const TIP_POST_RECEIVED = 5;
	    public const MESSAGE_UNLOCKED = 6;
	    public const SUBSCRIPTION = 7;
	    public const MEDIA_APPROVED = 8;
	}
*/
type Notification struct {
	gorm.Model
	UserID           int    `json:"user_id" gorm:"number;default:null`
	UserIdFrom       int    `json:"user_id_form" gorm:"number;not null;default:null`
	Readed           bool   `json:"readed" gorm:"bool;default:false`
	Content          string `json:"content" gorm:"text;not null`
	NotificationType int    `json:"type" gorm:"number;not null;default:1`
	ExtraData        string `json:"extra_data" gorm:"text;default:null`
	Channel          string `json:"channel" gorm:"text;default:email`
}

type Log struct {
	gorm.Model
	UserID      int    `json:"user_id" gorm:"number;default:null`
	Controller  string `json:"controller" gorm:"text;default:null`
	Function    string `json:"function" gorm:"text;default:null`
	Result      string `json:"result" gorm:"text;default:null`
	Browser     string `json:"browser" gorm:"text;default:null`
	Os          string `json:"os" gorm:"text;default:null`
	App         string `json:"app" gorm:"text;default:null`
	Ip          string `json:"ip" gorm:"text;default:null`
	Table       string `json:"table" gorm:"text;default:null`
	TableId     int    `json:"table_id" gorm:"text;default:null`
	Action      string `json:"action" gorm:"text;default:null`
	Micro       string `json:"micro" gorm:"text;default:null`
	Description string `json:"description" gorm:"text;default:null`
	Time        string `json:"time" gorm:"text;not null`
}

type Emails struct {
	gorm.Model
	UserID      int    `json:"user_id" gorm:"number;default:0`
	EmailTo     string `json:"email_to" gorm:"text;not null;default:'dddd'`
	EmailFrom   string `json:"email_from" gorm:"text;default:null`
	Enviornment string `json:"enviornment" gorm:"text;;default:null`
	Subject     string `json:"suject" gorm:"text;default:null`
	Body        string `json:"body" gorm:"text;default:null`
	Result      string `json:"result" gorm:"text;default:email`
	Time        string `json:"time" gorm:"text;not null`
}
