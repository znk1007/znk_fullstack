//1.搜索，2.消息，3.幻灯片，4.导航栏，5.魔方栏，6.公告，7.秒杀，8.广告栏，9.产品列表
enum ZNKMainModule {
  search,
  msessage,
  slide,
  nav,
  magic,
  notify,
  seckill,
  ads,
  prod,
}

class ZNKMainModel {
  //模块名称
  final ZNKMainModule module;
  //是否显示
  final bool show;

  ZNKMainModel({this.module, this.show});
}
