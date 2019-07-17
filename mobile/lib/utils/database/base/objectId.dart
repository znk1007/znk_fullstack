import 'dart:convert';
import 'dart:math';

const _lowerOrderTwoBytes = 0x0000FFFF;
const _lowerOrderThreeBytes = 0x00FFFFFF;
const int _encodedLen = 20;
const int _rawLen = 12;
///随机数
int _globalCounter = new Random.secure().nextInt(_lowerOrderThreeBytes);
///下一个随机数
int _nextCounter() {
  _globalCounter = (_globalCounter + 1) & _lowerOrderThreeBytes;
  return _globalCounter;
}
///模拟机器码
int _createMachineId() {
  return new Random.secure().nextInt(_lowerOrderThreeBytes);
}
/// 模拟进程id
int _createProcessId() {
  return new Random.secure().nextInt(_lowerOrderTwoBytes);
}
/// 生成机器码
int _generatedMachineId = _createMachineId();
/// 生成进程id
int _generatedProcessId = _createProcessId();
///生成整型
int _makeInt(int b3, int b2, int b1, int b0) => 
  (b3 << 24) | (b2 << 16) | (b1 << 8) | b0;

int _int3(int x) => (x & 0xFF000000) >> 24;

int _int2(int x) => (x & 0x00FF0000) >> 16;

int _int1(int x) => (x & 0x0000FF00) >> 8;

int _int0(int x) => (x & 0x000000FF);

class _ObjIdGen {
  /// 编码字节
  List<int> _dec;
  /// ascii编码
  AsciiCodec _codec;
  /// 编码字符串
  final _encodingStr = '0123456789abcdefghijklmnopqrstuv';  
  /**
   * [
   * 48:0, 49:1, 50:2, 51:3, 52:4, 53:5, 54:6, 55:7, 56:8, 57:9
   * 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 
   * 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 
   * ]
   *  */
  void _init() {
    _codec = AsciiCodec();
    _dec = new List<int>(256);
    int _decLen = _dec.length;
    
    for (var i = 0; i < _decLen; i++) {
      _dec[i] = 0xFF;
    }
    for (var i = 0; i < _encodingStr.length; i++) {
      Runes(_encodingStr);
      var idxs = _codec.encode(_encodingStr[i]);
      int idx = idxs.first;
      if (idx < _decLen) {
        _dec[idx] = i;
      }
    }
  }

  _ObjIdGen._() {
    _init();
  }
  static _ObjIdGen _gen;
  /// 单例模式
  static _ObjIdGen shared() {
    if (_gen == null) {
      _gen = _ObjIdGen._();
    }
    return _gen;
  }
  ///转字符串
  String string(List<int> bytes) {
    return _codec.decode(bytes);
  }
  /// 转字节
  List<int> bytes(String src) {
    return _codec.encode(src);
  }

  /// 编码
  void encode(List<int> dst, List<int> src) {
    if (dst.length != _encodedLen || src.length != _rawLen) {
      return;
    }
    enc(int byte) {
      return _codec.encode(_encodingStr[byte]).first;
    }

    dst[19] = enc((src[11]<<4)&0x1F);
    dst[18] = enc((src[11]>>1)&0x1F);
    dst[17] = enc((src[11]>>6)&0x1F|(src[10]<<2)&0x1F);
    dst[16] = enc(src[10]>>3);
    dst[15] = enc(src[9]&0x1F);
    dst[14] = enc((src[9]>>5)|(src[8]<<3)&0x1F);
    dst[13] = enc((src[8]>>2)&0x1F);
    dst[12] = enc((src[8]>>7)|(src[7]<<1)&0x1F);
    dst[11] = enc((src[7]>>4)&0x1F|(src[6]<<4)&0x1F);
    dst[10] = enc((src[6]>>1)&0x1F);
    dst[9]  = enc((src[6]>>6)&0x1F|(src[5]<<2)&0x1F);
    dst[8]  = enc((src[5]>>3));
    dst[7]  = enc(src[4]&0x1F);
    dst[6]  = enc((src[4]>>5)&0x1F|(src[3]<<3)&0x1F);
    dst[5]  = enc((src[3]>>2)&0x1F);
    dst[4]  = enc((src[3]>>7)&0x1F|(src[2]<<1)&0x1F);
    dst[3]  = enc((src[2]>>4)&0x1F|(src[1]<<4)&0x1F);
    dst[2]  = enc((src[1]>>1)&0x1F);
    dst[1]  = enc((src[1]>>6)&0x1F|(src[0]<<2)&0x1F);
    dst[0]  = enc((src[0]>>3));
  }

  /// 解码
  void decode(List<int> dst, List<int> src) {
    try {
      dst[11] = (_dec[src[17]]<<6)&0xFF | (_dec[src[18]]<<1)&0xFF | (_dec[src[19]]>>4)&0xFF;
      dst[10] = (_dec[src[16]]<<3)&0xFF | (_dec[src[17]]>>2)&0xFF;
      dst[9] = (_dec[src[14]]<<5)&0xFF | (_dec[src[15]])&0xFF;
      dst[8] = (_dec[src[12]]<<7)&0xFF | (_dec[src[13]]<<2)&0xFF | (_dec[src[14]]>>3)&0xFF;
      dst[7] = (_dec[src[11]]<<4)&0xFF | (_dec[src[12]]>>1)&0xFF;
      dst[6] = (_dec[src[9]]<<6)&0xFF | (_dec[src[10]]<<1)&0xFF | (_dec[src[11]]>>4)&0xFF;
      dst[5] = (_dec[src[8]]<<3)&0xFF | (_dec[src[9]]>>2)&0xFF;
      dst[4] = (_dec[src[6]]<<5)&0xFF | (_dec[src[7]])&0xFF;
      dst[3] = (_dec[src[4]]<<7)&0xFF | (_dec[src[5]]<<2)&0xFF | (_dec[src[6]]>>3)&0xFF;
      dst[2] = (_dec[src[3]]<<4)&0xFF | (_dec[src[4]]>>1)&0xFF;
      dst[1] = (_dec[src[1]]<<6)&0xFF | (_dec[src[2]]<<1)&0xFF | (_dec[src[3]]>>4)&0xFF;
      dst[0] = (_dec[src[0]]<<3)&0xFF | (_dec[src[1]]>>2)&0xFF;
    } catch (e) {
      print('decode e: ${e}');
    }
  }

}

class ObjectId implements Comparable<ObjectId> {

  List<int> _bytes = new List<int>(_rawLen);
  /// ObjId字节
  List<int> get bytes => _bytes;
  /// 时间戳
  int get timestamp => _timestamp;
  int _timestamp;
  /// 日期
  DateTime get date =>
    new DateTime.fromMillisecondsSinceEpoch(_timestamp * 1000);
  /// 设备id
  int get machineId => _machineId;
  int _machineId;
  // 进程id
  int get processId => _processId;
  int _processId;
  /// 随机数
  int get counter => _counter;
  int _counter;
  /// ObjId字符串
  String get string => _toString();

  /// id编解码器
  _ObjIdGen _gen = _ObjIdGen.shared();
  /// 工厂模式
  factory ObjectId.fromBytes(List<int> bytes) {
    return new ObjectId._byBytes(bytes);
  }

  /// 工厂模式
  factory ObjectId({DateTime time}) {
    DateTime theTime = time;
    if (theTime == null) {
      theTime = DateTime.now();
    }
    int timestamp = theTime.millisecondsSinceEpoch ~/ 1000;
    int machineId = _generatedMachineId;
    int processId = _generatedProcessId;
    int counter = _nextCounter();
    return new ObjectId._internal(timestamp, machineId, processId, counter);
  }
  /// 从字符生成
  factory ObjectId.fromString(String src) {    
    return new ObjectId._string(src);
  }
  /// 初始化
  ObjectId._string(String src) {
    if (src.length != _encodedLen) {
      throw 'invalid string';
    }
    List<int> dst = _gen.bytes(src);
    this._gen.decode(_bytes, dst);
    this._fromBytes(bytes);
  }

  /// 初始化
  ObjectId._byBytes(List<int> bytes) {
    this._fromBytes(bytes);
  }

  ///初始化
  ObjectId._internal(this._timestamp, this._machineId, this._processId, this._counter) {
    this._bytes[0] = _int3(this._timestamp);
    this._bytes[1] = _int2(this._timestamp);
    this._bytes[2] = _int1(this._timestamp);
    this._bytes[3] = _int0(this._timestamp);

    this._bytes[4] = _int2(this._machineId);
    this._bytes[5] = _int1(this._machineId);
    this._bytes[6] = _int0(this._machineId);

    this._bytes[7] = _int1(this._processId);
    this._bytes[8] = _int0(this._processId);

    this._bytes[9] = _int2(this._counter);
    this._bytes[10] = _int1(this._counter);
    this._bytes[11] = _int0(this._counter);    
  }
  /// 转字符串
  String _toString() {
    List<int> dst = new List<int>(_encodedLen);
    _gen.encode(dst, _bytes);
    return _gen.string(dst);
  }
  /// 字节处理
  void _fromBytes(List<int> bytes) {
    // print(object)
    if (bytes.length != _rawLen) {
      throw 'invalid bytes';
    }
    this._bytes = bytes;
    this._timestamp = _makeInt(bytes[0], bytes[1], bytes[2], bytes[3]);
    this._machineId = _makeInt(0, bytes[4], bytes[5], bytes[6]);
    this._processId = _makeInt(0, 0, bytes[7], bytes[8]);
    this._counter = _makeInt(0, bytes[9], bytes[10], bytes[11]);
    print('_timestamp: ${_timestamp}');
  }

  @override
  // TODO: implement hashCode
  int get hashCode {
    int result = _timestamp;
    result = 31 * result + _machineId;
    result = 31 * result + _processId;
    result = 31 * result + _counter;
    return result;
  }

  @override
  int compareTo(ObjectId other) {
    // TODO: implement compareTo
    if (this._timestamp == other._timestamp && 
      this._machineId == other._machineId &&
      this._processId == other._processId &&
      this._counter == other._counter) {
      return 1;
    }
    return 0;
  }

}

