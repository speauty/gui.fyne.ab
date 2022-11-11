package gui

type Cfg struct {
	AppName     string `json:"app_name,omitempty" yaml:"app_name" xml:"app_name" mapstructure:"app_name"`                     // 应用名称(主题窗口)
	AppIconPath string `json:"app_icon_path,omitempty" yaml:"app_icon_path" xml:"app_icon_path" mapstructure:"app_icon_path"` // 应用ICON资源路径
}
