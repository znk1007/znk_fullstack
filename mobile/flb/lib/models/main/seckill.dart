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
  //图片地址
  final String path;
  //币种
  final String coinType;

  ZNKSeckillItem({
    this.identifier,
    this.path = '',
    this.coinType,
    this.title,
    this.orgPrice,
    this.newPrice,
  });
}
