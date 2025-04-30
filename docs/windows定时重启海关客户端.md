# 重启脚本代码

```bat
@echo off
taskkill /f /im "SW.ClientApp.exe"
timeout /t 8 /nobreak >nul
"C:\Users\高新\Desktop\中国国际贸易单一窗口客户端.appref-ms"
```

# 设置重启任务

```CMD
# 设置start和kill任务 (在3点直接设置重启任务实效，可能是因为有那个弹窗，导致无法再次启动，因此在网络重启前，先kill，重启完后再start)
schtasks /create /tn "KillCustomsClientTask" /tr "C:\Users\高新\Desktop\scripts\kill.bat" /sc daily /st 01:00
schtasks /create /tn "StartCustomsClientTask" /tr "C:\Users\高新\Desktop\scripts\start.bat" /sc daily /st 03:00

# 删除任务
schtasks /delete /tn "CustomsClientTask" /f
```
