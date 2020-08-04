# noodlenet
## websocket通信
写是狗屎 是毛线 是他娘的锤锤

## 说明
关于广播，大多的解决方案是使用go+for发送。而这里是利用close chan会通知给所有监听的特性处理，更好的利用go语言的特性。 

## 更新
- 2020年7月18日 超时机制
- 2020年7月30日 使用proto

## 演示
#### 演示1
Echo服务器
#### 演示2
简单的聊天工具，示例图片如下：

![示例01](https://github.com/Jourmey/noodlenet/blob/master/_example/2.broadcast/view/example01.jpg)
![示例02](https://github.com/Jourmey/noodlenet/blob/master/_example/2.broadcast/view/example02.jpg)