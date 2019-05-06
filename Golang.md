golang
编译型语言，强类型，并发性能优秀，可面向过程也可面向对象，库函数支持充分，开发快？吗

new make 区别：
make是对内建类型的切片，channel，map；new是对所有对象的；
new 相当于 &T{} ,分配内存，也就是建一个空的T，并把对应内存地址返回，内部对象都是nil
make 还带初始化的功能，map，slice，chan 数据部分都是指向块空间的，只不过空间是大小为0
make和new都是在堆上分配内存

init：
任何包中的init方法都会在该包第一次加载是被执行一次，来初始化一些包在运行中必要的数据，一个包中可以有多个init，但不建议这么做，太乱

自有类型：
数组

slice 
是数组的一种上层封装，结构有三个属性：指向数组的指针，长度，容量；没有结尾指针，通过长度控制；
reslice重新申请空间，只能扩大不能缩小，规则是如果原容量小于1024，则直接翻倍；如果大于1024每次扩大1/4（1.25倍）；内部核心增长容量函数扩容cap是可以传入的，如果传入大于double，这直接扩容cap；如果小于double且原来容量页小于1024，则double；如果小于double，原容量大于1024则每次扩1/4，直到扩到大于cap；扩容后地址可能不变（在原数组后申请空间），也可能变化，申请全新空间，复制老数组数据到新空间
append追加数据，如果len要超过cap，就要扩容了，这是地址可能变化了
copy复制一个切片到另一个切片，返回复制的长度 先判断fm 和 to哪个长，返回的是短的长度


map指针

chan 指针
hchan核心结构，count，size，elemtype/size，sendq，recvq，sendx，recvx，close，lock，buf；
根据size=0时，没有缓冲，可做同步信号用。size大于0时，buf有长度，则变异步
chansend写channel，如果有缓冲区，且没满，则直接写入buf，并尝试唤起阻塞的recv；如果缓存区满了，则构造sudog结构，写入sendq中
chanrecv堵channel，如果有缓冲区，且不是空，则直接读buf，并尝试唤起阻塞的send；如果缓存区空了，则构造sudog结构，写入secvq中
close关闭channel，关闭的channel不能写，不能关，可以读，读到的是空，不会被阻塞；close时会循环唤醒sendq和recvq的goroutine


gc
1.3是开始，1.5是飞跃，逐步提升中，2s到100um

goroutine（PMG模型）
Sched维护p、m队列的管理器，m是线程，p是处理器用来执行g，存这归它执行的g列表，g就是需要运行的goroutine信息；GOMAXPROCS设置的是p的个数，最大可以是256，其实和cpu无关；
四个结构支撑了goroutine的调度；首先程序启动时会执行runtime schedinit，初始化sched，会根据GOMAXPROCS生成相应数量的P，处于空闲状态；然后newproc创建主goroutine，就是main里的逻辑；然后新创建一个m，用来监控sysmon；之后就正常运行业务逻辑了
每个go xxx()都会调用newproc创建一个新的g，然后把g挂到所属m的p列表里
golang没有语言级的创建m的方法，什么时候该创建，有内核线程决定，当g多，p有空闲是，就会创建m；创建的过程是newm得到新线程，然后把空闲p绑定到m上acquirep，与之相反的是，当m空闲时需要与p解绑releasep；
m和p绑定后，执行schedule，先从p中获得可执行的g，如果没有，会去全局等待队列（sched中？）中取g，如果还没有，就去其他的m的p中取一半挂在自己的p中；还没有，就解绑m、p，m:sleep,p空闲；
如果发现g数量上升，会再次唤醒m重组p，继续执行g；
执行g的过程可能会出现挂起，如channel阻塞、定时器、网络等都会gopark，这时要让出p（处于syscall，开头的sysmon会检测出来），换一个空闲的m去执行剩余阻塞的g，而这个g会在执行完后，放入全局等待队列，待执行





channel

sort
map
reflect
net
sync
unsafe
io、string
pporf
interface