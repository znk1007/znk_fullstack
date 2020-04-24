package usernet

// //updatePswSrv 更新密码服务
// type updatePswSrv struct {
// 	req     *userproto.UserUpdatePswReq
// 	resChan chan updatePswRes
// 	doing   map[string]bool
// 	token   *usermiddleware.Token
// 	pool    *userpayload.WorkerPool
// }

// //newUpdatePswSrv 初始化更新密码服务
// func newUpdatePswSrv() *updatePswSrv {
// 	srv := &updatePswSrv{
// 		resChan: make(chan updatePswRes),
// 		doing:   make(map[string]bool),
// 		token:   usermiddleware.NewToken(60 * 5),
// 		pool:    userpayload.NewWorkerPool(100),
// 	}
// 	srv.pool.Run()
// 	return srv
// }

// //write 写入数据
// func (up *updatePswSrv) write(req *userproto.UserUpdatePswReq) {
// 	up.pool.WriteHandler(func(j chan userpayload.Job) {
// 		up.req = req
// 		j <- up
// 	})
// }

// //read 读取数据
// func (up *updatePswSrv) read(ctx context.Context) (res *userproto.UserUpdatePswRes, err error) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			return nil, ctx.Err()
// 		case res := <-up.resChan:
// 			return res.res, res.err
// 		}
// 	}
// }

// /*
// 用户ID：userID，
// 会话ID：sessionID，
// 时间戳：timestamp，
// 设备ID：deviceID，
// 设备名：deviceName，
// 平台类型：platform，
// 应用标识：appkey，
// 旧密码：password，
// 新密码：newPsw
// */

// //handlUpdatePsw 处理更新密码
// func (up *updatePswSrv) handlUpdatePsw() {
// 	req := up.req
// 	//账号
// 	acc := req.GetAccount()
// 	//校验账号是否为空
// 	if len(acc) == 0 {
// 		log.Info().Msg("miss `account` or account cannot be empty")
// 		up.makeUpdatePswToken("", http.StatusBadRequest, errors.New("miss `account` or account cannot be empty"))
// 		return
// 	}
// 	//token
// 	tkstr := req.GetToken()
// 	if len(tkstr) == 0 {
// 		log.Info().Msg("token cannot be empty")
// 		up.makeUpdatePswToken(acc, http.StatusBadRequest, errors.New("token cannot be empty"))
// 		return
// 	}
// 	//正在执行中
// 	if up.doing[acc] {
// 		log.Info().Msg("account is operating, please do it later")
// 		up.makeUpdatePswToken(acc, http.StatusBadRequest, errors.New("account is operating, please do it later"))
// 		return
// 	}
// 	//解析token
// 	tk := up.token
// 	err := tk.Parse(tkstr)
// 	if err != nil {
// 		log.Info().Msg(err.Error())
// 		up.makeUpdatePswToken(acc, http.StatusBadRequest, err)
// 		return
// 	}
// 	//通用请求校验
// 	code, err := usermiddleware.CommonRequestVerify(acc, tk)
// 	if err != nil {
// 		log.Info().Msg(err.Error())
// 		up.makeUpdatePswToken(acc, code, err)
// 		return
// 	}
// 	//更新密码
// 	err = usermodel.SetUserPassword(acc, tk.UserID, tk.Password)
// 	if err != nil {
// 		log.Info().Msg(err.Error())
// 		up.makeUpdatePswToken(acc, http.StatusBadRequest, err)
// 		return
// 	}
// 	up.makeUpdatePswToken(acc, http.StatusOK, nil)
// }

// /*
// 状态码：code，
// 反馈消息：message，
// 时间戳：timestamp
// */
// //makeUpdatePswToken 生成更新密码响应token
// func (up *updatePswSrv) makeUpdatePswToken(acc string, code int, err error) {
// 	msg := ""
// 	if err != nil {
// 		msg = err.Error()
// 	}
// 	resmap := map[string]interface{}{
// 		"code":      code,
// 		"message":   msg,
// 		"timestamp": time.Now().String(),
// 	}
// 	var tk string
// 	tk, err = up.token.Generate(resmap)

// 	res := updatePswRes{
// 		res: &userproto.UserUpdatePswRes{
// 			Account: acc,
// 			Token:   tk,
// 		},
// 		err: err,
// 	}
// 	up.resChan <- res
// }

// func (up *updatePswSrv) Do() {
// 	go up.handlUpdatePsw()
// }
