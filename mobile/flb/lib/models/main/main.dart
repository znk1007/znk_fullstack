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

//广告
class ZNKBannerModel {
  final String identifier;
  final String path;
  ZNKBannerModel({this.identifier, this.path});
}

//集合
class ZNKNav {
  final String identifier;
  final String title;
  final String path;
  ZNKNav({this.identifier, this.title, this.path});
}

//便捷入口
class ZNKMagic {
  //唯一标识
  final String identifier;
  //icon路径
  final String path;

  ZNKMagic({this.identifier, this.path});
}

class ZNKSeckill {
  //标题
  final String title;
  //时间
  final String time;
  //项目
  final List<ZNKSeckillItem> items;
  ZNKSeckill({this.title, this.time, this.items});
}

class ZNKSeckillItem {
  //唯一标识
  final String identifier;
  //标题
  final String title;
  //原价
  final String orgPrice;
  //秒杀价
  final String newPrice;

  ZNKSeckillItem({this.identifier, this.title, this.orgPrice, this.newPrice});
}

class ZNKCombine {
  //标题
  final String title;
  //项目集合
  final List<ZNKCombineItem> items;

  ZNKCombine({this.title, this.items});
}

class ZNKCombineItem {
  //唯一标识
  final String identifier;
  //标题
  final String title;
  //价格
  final String price;

  ZNKCombineItem({this.identifier, this.title, this.price});
}

//产品列表
class ZNKProd {
  //段标题
  final String sectionTitle;
  //大标题
  final String title;
  //详情
  final String detail;
  //图片路径
  final String path;
  //标签集合
  final List<String> tags;
  //原价
  final String orgPrice;
  //折扣价
  final String newPrice;
  //已卖
  final String solt;
  //库存
  final String stock;
  //商品列表
  final List<ZNKProd> prods;

  ZNKProd(
      {this.sectionTitle,
      this.title,
      this.detail,
      this.path,
      this.tags,
      this.orgPrice,
      this.newPrice,
      this.solt,
      this.stock,
      this.prods});
}
