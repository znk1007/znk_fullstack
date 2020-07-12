import 'package:flb/viewmodels/main/prodlist.dart';
import 'package:flb/views/base/base.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class ZNKProdListView extends StatelessWidget {
  ZNKProdListView({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ZNKBaseView<ZNKProdListViewModel>(
      model: ZNKProdListViewModel(
        api: Provider.of(context),
      ),
      onReady: (model) => model.fetch(),
      builder: (context, model, child) => Container(
        child: ListView.builder(
          itemCount: 1,
          itemBuilder: (BuildContext context, int index) {
            return;
          },
        ),
      ),
    );
  }
}
