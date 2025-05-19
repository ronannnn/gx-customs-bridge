# 重启脚本代码

该程序是ClickOnce程序，要启动的话，必须要`start "" "path\to\appref-ms"`，否则CMD的定时任务运行不起来。
```bat
@echo off
taskkill /f /im "SW.ClientApp.exe"
timeout /t 8 /nobreak >nul
start "" "C:\Users\高新\Desktop\中国国际贸易单一窗口客户端.appref-ms"
```

# 设置重启任务

```CMD
# 设置start和kill任务 (在3点直接设置重启任务实效，可能是因为有那个弹窗，导致无法再次启动，因此在网络重启前，先kill，重启完后再start)
schtasks /create /tn "KillCustomsClientTask" /tr "C:\Users\高新\Desktop\scripts\kill.bat" /sc daily /st 00:00
schtasks /create /tn "StartCustomsClientTask" /tr "C:\Users\高新\Desktop\scripts\start.bat" /sc daily /st 05:00

# 删除任务
schtasks /delete /tn "CustomsClientTask" /f
```
