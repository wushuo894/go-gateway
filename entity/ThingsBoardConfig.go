package entity

type ThingsBoardConfig struct {
	/**
	主机IP
	*/
	Host string `json:"host"`
	/**
	端口
	*/
	Port int32 `json:"port"`
	/**
	客户端ID
	*/
	ClientId string `json:"client_id"`
	/**
	用户名
	*/
	UserName string `json:"user_name"`
	/**
	密码
	*/
	Password string `json:"password"`
	/**
	超时时间
	*/
	Timeout int64 `json:"timeout"`
}
