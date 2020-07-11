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
