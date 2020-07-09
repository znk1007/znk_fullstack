class ZNKHelp {
  /* 安全字符串 */
  static String safeString(dynamic source) {
    return source != null ? source.toString() : '';
  }
}