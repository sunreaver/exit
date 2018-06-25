# exit
优雅的退出守护进程

`RegistExiter` 注册监控退出信号

返回值`exiter` :当有退出信号到达，会写入此信号
返回值`delay` :业务方收到退出信号exiter后，执行业务缓存等逻辑，完成后关闭此delay信号，才会正式退出进程


`UnRegistExiter` 取消监控退出信号

参数`exiter` :业务方传入RegistExiter返回的exiter来取消此信号的监控
