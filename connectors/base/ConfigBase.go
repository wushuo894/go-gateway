package base

// ConfigBase 设备配置
type ConfigBase struct {
	/**
	设备名称
	*/
	DeviceName string `json:"device_name"`
	/**
	设备类型
	*/
	DeviceType string `json:"device_type"`
	/**
	设备Id
	*/
	DeviceId string `json:"device_id"`
	/**
	文件
	*/
	FileName string `json:"file_name"`
	/**
	设备连接
	*/
	Connector *ConnectorBase `json:"connector"`
}
