package request

type Resigter struct {
	Name     string `form:"name" json:"name" binding:"requried"`
	Mobile   string `form:"mobile" json:"mobile" binding:"requried"`
	Password string `form:"password" json:"password" binding:"requried"`
}

func (resigter Resigter) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Name.required":     "用户名不能为空",
		"Mobile.required":   "手机号码不能为空",
		"Password.required": "用户密码不能为空",
	}
}
