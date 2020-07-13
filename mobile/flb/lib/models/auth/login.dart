class ZNKLogin {
  //账号
  final String account;
  //密码
  final String password;

  ZNKLogin({this.account, this.password})
      : assert(account != null && account.length > 0),
        assert(password != null && password.length > 0);
}
