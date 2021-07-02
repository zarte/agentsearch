# 使用说明
common.data为可用代理列表
## 获取可用代理
执行
main.exe ./
./为配置文件config.ini所在文件，配置文件需无bom头否则首行需插入任意一行无用配置信息。
## 现有代理过滤
在./目录下新建proxlist.data文件，格式内容为上一步获取到的common.data.
## 注意点
common.data每次执行不会自动覆盖，需要手动删除防止重复。