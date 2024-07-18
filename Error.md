## 进度

2和14的概念看的有点点懵，先实战？

3已看完，就是配置ubuntu

> （四）、二对应的就是3.4 的测试网络

4的话好像目前是卡在下载fabric的源码，不知道为什么gitee登录不上去（网络没问题 -- 已搞定

## skill

```bash
# 在图形化界面中进入根文件夹
nautilus /

# 删除
rm -f /var/tmp/environment.swp

#修改权限
sudo chown -R xuan:xuan /usr/local/

# 配置文件夹
sudo vim ~/.bashrc

# 查看某个变量
echo $GOPATH

# 重启
sudo reboot
```

## 问题

### E:无法定位软件包 open-vm-tools-desktop

解决办法就是sudo apt-get update

### 使用sudo apt-get update报错显示域名无法解析，

如“错误:4 http://security.ubuntu.com/ubuntu jammy-security InRelease暂时不能解析域名“security.ubuntu.com”

分析原因：没有联网，设置处的网络适配器改成桥接模式即可联网，也可以通过判断右上角的网络状态查看是否联网成功，联网后再执行命令就可以正常update

搭梯子的话就开网络代理，填本机ip，端口号填小猫的7890

![image-20240714003833135](https://cdn.jsdelivr.net/gh/kixuan/PicGo/image-20240714003833135.png)

【install前都要看看网络是否连接成功，尤其是出现不能解析域名的情况】

### 实现复制

[将Windows复制的文字或文件粘贴到VMware_windows文件拖拽至vmware中-CSDN博客](https://blog.csdn.net/qq_44786250/article/details/125683954)

### 配置了环境变量但还是报错 configtxgen 无法找到命令

已重启试过🤔

很诡异的一个报错，一直显示configtxgen:找不到命令

路径什么的都是对的

后面才发现是bin文件夹里面都没有这个工具

进到可视化根目录发现莫名其妙多了个farci-simples文件夹（实际上应为faric-samples）把这个里面的东西全放进faric-samples就ok了，应该是创建文件夹的时候不小心打错字了oO

![image-20240718144102011](https://cdn.jsdelivr.net/gh/kixuan/PicGo/image-20240718144102011.png)

### docker pull一直超时

方法二可行(๑•̀ㅂ•́)
و✧：[从Docker Hub 拉取镜像一直失败超时？这些解决方案帮你解决烦恼_docker拉取镜像超时-CSDN博客](https://blog.csdn.net/weixin_50160384/article/details/139861337)

### Shell脚本执行报错 ：Synatx error: “(” unexpected (expecting “fi”)

![image-20240717135803720](https://cdn.jsdelivr.net/gh/kixuan/PicGo/image-20240717135803720.png)

解决办法：https://blog.csdn.net/qq_40907977/article/details/102603098

### git错误

git checkout release-2.3 fatal: detected dubious ownership in repository at '
/usr/local/gocode/src/github.com/hyperledger/fabric' To add an exception for this directory, call:    git config
--global --add safe.directory /usr/local/gocode/src/github.com/hyperledger/fabric

这个错误消息表明，Git 发现你尝试操作的仓库目录的所有权有问题，这可能是由于以下原因之一：

1. 你正在以不同于目录所有者的用户身份运行 Git。
2. 你正在使用的仓库目录是通过 `sudo` 或其他用户创建的，但你现在尝试以普通用户身份访问它。

方法 1: 将目录添加到安全目录列表

```Bash
git config --global --add safe.directory /usr/local/gocode/src/github.com/hyperledger/fabric
```

这会将目录 `/usr/local/gocode/src/``github.com/hyperledger/fabric` 添加到全局的安全目录列表中，从而避免出现该错误。

方法 2: 修改目录所有权

```Bash
sudo chown -R xuan:xuan /usr/local/gocode/src/github.com/hyperledger/fabric
```

方法 3: 使用 `sudo` 运行 Git 命令

```Bash
sudo git checkout release-2.3
```

不过，这不是推荐的解决方案，除非你确实需要以超级用户身份进行 Git 操作。

### 无法mkdir，显示权限不够

> 不需要使用root用户去创建别的用户的文件夹!!
>
> 主要看看该目录的所有者属于谁:/home/user, sudo chown user:user /home/user

E505: "/etc/profile" 只读 (请加 ! 强制执行) 29,0-1 底端：也是权限问题

### 创建通道报错

服了怎么又出bug：在创建通道的时候报错 > <

![image-20240717154919666](https://cdn.jsdelivr.net/gh/kixuan/PicGo/image-20240717154919666.png)

查看csdn给出的解决办法都是修改`createChannel.sh`文件，把channelID改成--channel-id，但是我的已经是--channel-id了，于是
osnadmin --help，发现创建通道的命令是：

`channel join --channelID=CHANNELID --config-block=CONFIG-BLOCK`
Join an Ordering Service Node (OSN) to a channel. If the channel does not yet exist, it will be created.

翻译一下：

- `--channelID=CHANNELID`: 要加入的通道 ID。

- `--config-block=CONFIG-BLOCK`: 包含通道配置区块的文件路径。

那就修改✌

![image-20240717155508804](https://cdn.jsdelivr.net/gh/kixuan/PicGo/image-20240717155508804.png)

### 部署Go链码时的报错

Error: failed to normalize chaincode path: failed to determine module root: exec: "go": executable file not found in
$PATH

其实这里就是go的环境变量没配置好或者go没下载好

但是不记得动了什么东西，ubuntu突然中文乱码而且搞了半个小时还没搞好，7.17 先keep到这里

7.18 continue 不知道什么原因，已经配置好go的环境变量了，但是go version还是会显示没有

这个时候只要再`source  ~/.bashrc`就可以了，奇怪捏🤔

### firefox访问很慢

特别是外网：

1.
用处不大：[ubuntu下firefox有时打不开个别网页解决办法_ubuntu火狐unable to connect-CSDN博客](https://blog.csdn.net/lssyg2011/article/details/104058116?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-4-104058116-blog-131341796.235^v43^pc_blog_bottom_relevance_base5&spm=1001.2101.3001.4242.3&utm_relevant_index=5)
2.
修改DNS【妙！找个时间学学DNS，好像项目高性能方法也可以用到它】：[解决ubuntu系统火狐浏览器一些网页打不开的问题（重新配置DNS）_ubuntu20配置火狐浏览器的dns-CSDN博客](https://blog.csdn.net/weixin_44494462/article/details/104411177?spm=1001.2101.3001.6650.6&utm_medium=distribute.pc_relevant.none-task-blog-2~default~BlogCommendFromBaidu~Rate-6-104411177-blog-104058116.235^v43^pc_blog_bottom_relevance_base5&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2~default~BlogCommendFromBaidu~Rate-6-104411177-blog-104058116.235^v43^pc_blog_bottom_relevance_base5&utm_relevant_index=11)

### 中文乱码

在网上找了很多方法都不行，然后第二天早起一起来突发奇想设置一下编码就可以了，之前都是简体中文，改成UFT8

![image-20240718143859436](https://cdn.jsdelivr.net/gh/kixuan/PicGo/image-20240718143859436.png)

tree命令中文乱码

两种办法：https://blog.csdn.net/cxrsdn/article/details/100006348
