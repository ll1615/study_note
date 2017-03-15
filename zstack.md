# 架构

- 异步架构

  传统的Iaas采用同步架构，线程池为每个任务提供线程，线程只有在前一个任务完成之后才能为下一个任务服务。例如创建虚拟机，通常的执行路径为”identity service-->scheduler-->image service-->storage service-->network service-->hypervisor”，它的每一环节都要一定的耗时，大批量任务执行时，延迟效应就更明显，最后就可能出现任务失败。<br/>
  
  ZStack 99%的任务都是异步执行的，这是使它单个节点能够管理几十万的物理机、几百万的虚拟机和处理上万并发人物的关键。<br/>
  
  ZStack 的异步架构由三部分组成：异步消息、异步方法和异步HTTP调用。<br/>
  异步消息：<br/>&nbsp;&nbsp;
  ZStack使用 RabbitMQ 作为连接各种服务的消息总线。
