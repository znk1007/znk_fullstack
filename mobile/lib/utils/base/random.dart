import 'dart:math' as math;

const _preList = [
  "130","131","132","133","134","135","136","137","138","139",
  "147",
  "150","151","152","153","155","156","157","158","159",
  "166",
  "171","176","177",
  "186","187","188",
  "198"
];

int _min = 0;
int _max = 10;

class RandomManager {
  // 随机生成手机号
  static randomPhone() {
    String prefix = _preList[math.Random.secure().nextInt(_preList.length)];
    String subfix = '';
    for (var i = 0; i < 8; i++) {
      subfix = subfix + (_min + (math.Random.secure().nextInt(_max - _min))).toString();
    }
    return prefix + subfix;
  }
}