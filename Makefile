# Windows版

# 项目根目录
DirProjectRoot = E:\MonkeyCode\github.com\speauty\gui.fyne.ab
# 可执行文件目录
DirProjectBin = ${DirProjectRoot}\bin
# 源码目录
DirProjectSrc = ${DirProjectRoot}\src
#链接参数
ParamsLDFlags = -s -w -H windowsgui
#应用名称
AppName = gui.fyne.ab.exe

run:
	cd ${DirProjectSrc} && go build -ldflags="${ParamsLDFlags}" -o ${DirProjectBin}\${AppName} gui.fyne.ab && cd ${DirProjectBin} && ${AppName}

build:
	cd ${DirProjectSrc} && go build -ldflags="${ParamsLDFlags}" -o ${DirProjectBin}\${AppName} gui.fyne.ab && exit

kill:
	 taskkill /f /t /im ${AppName}

# @todo 环境存在问题
release:
	fyne-cross windows -arch=* -developer=speauty -env=GOPROXY="https://goproxy.cn" -env=ldflags="-s -w -H windowsgui"

compress: # 必须安装upx服务, 否则该指令无法使用
	cd ${DirProjectBin} && upx -9 *.exe

# 生成syso文件
genSyso:
	windres -o gui.fyne.ab.syso gui.fyne.ab.rc