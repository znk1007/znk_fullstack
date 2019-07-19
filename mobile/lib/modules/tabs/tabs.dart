import 'package:flutter/material.dart';
import 'package:znk/modules/tabs/chat/chat.dart';
import 'package:znk/modules/tabs/owner/owner_page.dart';
import 'package:znk/modules/tabs/schedule/schedule.dart';
import 'package:znk/core/user/index.dart';
class Tabs extends StatefulWidget {

  UserRepository _userRepository;
  
  Tabs({Key key, @required UserRepository userRepository}):
    assert(userRepository != null),
    _userRepository = userRepository,
    super(key: key);
  
  _TabsState createState() => _TabsState();
}

class _TabsState extends State<Tabs> with WidgetsBindingObserver {

  // 当前位置
  int _currentIndex = 2;
  List<Widget> _tabsList = [];
 
  @override
  void initState() {
    super.initState();
    _tabsList.add(Chat(userRepository: widget._userRepository));
    _tabsList.add(Schedule(userRepository: widget._userRepository));
    _tabsList.add(Owner(userRepository: widget._userRepository));
    WidgetsBinding.instance.addObserver(this);
  }

  @override
  void dispose() {
    WidgetsBinding.instance.removeObserver(this);
    super.dispose();
  }


  @override
  void didChangeAppLifecycleState(AppLifecycleState state) {
    print('current state: $state');
    if (state == AppLifecycleState.resumed) {
      widget._userRepository.updateOnline(context, true);
    } else {
      widget._userRepository.updateOnline(context, false);
    }
  }


  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: this._currentIndex > this._tabsList.length ? this._tabsList[0] : this._tabsList[this._currentIndex],
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: this._currentIndex,
        onTap: (idx) {
          setState(() {
            this._currentIndex = idx;
          });
        },
        items: [
          BottomNavigationBarItem(
            icon:  Image.asset('lib/images/iOS/tabs/tab_msg_n.png'),
            activeIcon: Image.asset('lib/images/iOS/tabs/tab_msg_s.png'),
            title: Text('消息')
          ),
          BottomNavigationBarItem(
            icon:  Image.asset('lib/images/iOS/tabs/tab_schedule_n.png'),
            activeIcon: Image.asset('lib/images/iOS/tabs/tab_schedule_s.png'),
            title: Text('日程')
          ),
          BottomNavigationBarItem(
            icon:  Image.asset('lib/images/iOS/tabs/tab_owner_n.png'),
            activeIcon: Image.asset('lib/images/iOS/tabs/tab_owner_s.png'),
            title: Text('我的')
          ),
        ],
      ),
    );
  }
}