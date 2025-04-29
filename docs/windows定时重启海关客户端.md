# 重启脚本代码

```bat
@echo off
taskkill /f /im "SW.ClientApp.exe"
timeout /t 8 /nobreak >nul
"C:\Users\高新\Desktop\中国国际贸易单一窗口客户端.appref-ms"
```

# 设置重启任务

```CMD
# 设置重启任务
schtasks /create /tn "CustomsClientTask" /tr "C:\Users\高新\Desktop\restart.bat" /sc daily /st 03:00

# 删除任务
schtasks /delete /tn "CustomsClientTask" /f
```
