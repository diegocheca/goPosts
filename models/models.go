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
