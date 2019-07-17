
class Validators {
  static final RegExp _accountRegExp = RegExp(
    r'^[0-9A-Za-z\u4e00-\u9fa5]$',
    multiLine: true,
  );

  static final RegExp _passwordRegExp = RegExp(
    r'(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$',
  );
  // 账号是否符合规则
  static isValidAccount(String account) {
    print('match account: $account');
    print('match account res: ${_passwordRegExp.hasMatch(account)}');
    return !_accountRegExp.hasMatch(account);
  }
  // 密码是否符合规则
  static isValidPassword(String password) {
    print('match password: $password');
    print('match password res: ${_passwordRegExp.hasMatch(password)}');
    return _passwordRegExp.hasMatch(password);
  }

}