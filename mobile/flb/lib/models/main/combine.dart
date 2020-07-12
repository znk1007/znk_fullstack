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
  //币种
  final String coinType;
  //价格
  final String price;
  //图片路径
  final String path;

  ZNKCombineItem({
    this.identifier,
    this.title,
    this.coinType,
    this.price,
    this.path,
  });
}
