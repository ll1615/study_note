# L2~L4网络
  - L2: 物理层 ＋ 数据链路层, 属交换网络（MAC地址识别） 
  - L3: 物理层 ＋ 数据链路层 ＋ 网络层, 属路由网络（IP＋MAC地址识别）
  - L4: 就包含了传输层了, 应加入端口以及协议号进行识别了
  - [OSI参考模型](http://baike.baidu.com/link?url=3rLRQ-0KFNgz381PFGnxU9O4rmPgDBsK0Z-IFmzhYBkvqKiUSVDTJ4cKgsQmociSB3rh54WwejsKyaUCIjdCZkDEYR93Av-hMB796SlkWMEOZwDZyjRWm6ygOY8PfikF)

# [hypervisor](http://www.ibm.com/developerworks/cn/linux/l-hypervisor/)
 - hypervisor 之于操作系统类似于操作系统之于进程
 - hypervisor 可以划分为两大类。首先是类型 1，这种 hypervisor 是直接运行在物理硬件之上的。其次是类型 2，这种 hypervisor 运行在另一个操作系统（运行在物理硬件之上）中。类型 1 hypervisor 的一个例子是基于内核的虚拟机（KVM(kernel-based virtual machine) —— 它本身是一个基于操作系统的 hypervisor）。类型 2 hypervisor 包括 QEMU 和 WINE。
