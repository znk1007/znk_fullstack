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
