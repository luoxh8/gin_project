基于gin框架的一个生产就绪模板。

# 部署
使用了supervisor作为进程监控。



依赖安装命令：sh dependencies.sh即可把相关依赖装上，如果出现依赖无法安装



请自行去github搜索下下载，放到相应的文件夹下面



supervisor配置文件的名字：gin_project.conf



ubuntu下，请使用```ln -s /root/gin_project/gin_project.conf /ect/supervisor/conf.d/gin_project.conf``` 



**注意**，其中**/root/gin_project/gin_project.conf**替换成自己的配置文件的位置。



# 项目框架以及依赖

[gin](https://github.com/gin-gonic/gin)

[orm](https://github.com/astaxie/beego/tree/develop/orm)

[websocket](https://github.com/gorilla/websocket)

[redis](https://github.com/go-redis/redis)

[sessions](https://github.com/gin-contrib/sessions)

