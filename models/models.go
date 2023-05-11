package models

import "gorm.io/gorm"

type Fact struct {
    gorm.Model
    Question string `json:"question" gorm:"text;not null;default:null`
    Answer   string `json:"answer" gorm:"text;not null;default:null`
}


type Post struct {
    gorm.Model
    Image string `json:"image" gorm:"text;default:null`
    Thumbnail string `json:"thumbnail" gorm:"text;not null;default:null`
    Slug   string `json:"slug" gorm:"text;default:null`
    Title   string `json:"title" gorm:"text;not null;default:null`
    Subtitle   string `json:"subtitle" gorm:"text;not null;default:null`
    Content   string `json:"content" gorm:"text;not null;default:null`
    Author   int `json:"author" gorm:"number;not null;default:null`
    Rate   int `json:"rate" gorm:"number;not null;default:null`
}



type Comment struct {
    gorm.Model
    PostID int `json:"postid" gorm:"number;default:null`
    Content string `json:"content" gorm:"text;not null;default:null`
    Author   int `json:"author" gorm:"number;not null;default:null`
    Likes   int `json:"likes" gorm:"number;not null;default:null`
    Image string `json:"image" gorm:"text;default:null`
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
    UserID int `json:"user_id" gorm:"number;default:null`
    UserIdFrom int `json:"user_id_form" gorm:"number;not null;default:null`
    Readed bool `json:"readed" gorm:"bool;default:false`
    Content   string `json:"content" gorm:"text;not null`
    NotificationType   int `json:"type" gorm:"number;not null;default:1`
    ExtraData string `json:"extra_data" gorm:"text;default:null`
    Channel string `json:"channel" gorm:"text;default:email`
}

type Log struct {
    gorm.Model
    UserID int `json:"user_id" gorm:"number;default:null`
    Controller   string `json:"ccontroller" gorm:"text;not null`
    Function   string `json:"function" gorm:"text;not null`
    Result   string `json:"result" gorm:"text;not null`
    Time   string `json:"time" gorm:"text;not null`
}



