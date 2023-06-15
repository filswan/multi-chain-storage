export default {
  route: {
    dashboard: '首页',
    Deal: '我的文件',
    Miner: '存储提供者管理',
    myProfile: '个人信息',
    browse: '浏览任务',
    Tools: '工具',
    My_FS3: 'FS3',
    Stats: '统计',
    myAccount: '设置',
    Upload_files: '上传文件',
    Search_file: '搜索文件',
    documentation: '文档',
    metaSpace: 'Bucket存储',
    About: '关于',
    Pricing: '定价',
    Resources: '资源',
    Login: '登录',
    ApiKey: 'API密钥',
    report: '审计报告'
  },
  navbar: {
    sidebar_header: 'Multi-Chain Storage',
    section: '部分',
    myprofile: '个人信息',
    support: '支持',
    hi: '您好，',
    logout: '登出',
    log_in: '登录',
    sign_up: '注册',
    dashboard: '首页',
    minerDetail: '存储提供者详情',
    deal: '详情',
    miner: '存储提供者管理',
    deal_details: '任务详情',
    My_Profile: '个人信息',
    Browse_tasks: '浏览任务',
    create_tasks: '创建任务',
    Upload_files: '上传文件',
    Tools: '工具',
    Stats: '统计',
    copy: '© 2021 FilSwan Canada',
    ScoreChange: '历史评分',
    ScoreChange_legend: '评分',
    PowerChange: '历史算力',
    MiningData: '存储/数据封装统计',
    selectDate: {
      sevenDays: '7天',
      thirtyDays: '30天',
      oneYearL: '1年',
      roundClock: '24时'
    },
    Power: '有效算力',
    PowerGrowth: '有效算力增量',
    BillingHistory: '账单详情',
    language: '语言',
    language_cn: '简体中文',
    swan_share: 'swan分配'
  },
  // 首页
  dashboard: {
    text01: '存储提供者总数',
    text02: '活跃存储提供者',
    form_name: 'Filecoin存储提供者',
    form_swith: '接受离线交易',
    form_swith_in: '网上交易',
    search: '搜索存储节点ID...',
    formTable01: '地区',
    formTable02: '存储节点ID',
    formTable03: '价格 (GiB/区块时间)',
    formTable04: '价格 (32GiB/年)',
    formTable05: '真实数据价格(GiB/区块时间)',
    formTable06: '规格',
    formTable07: '离线交易',
    formTable08: '评分',
    formTable09: '上次更新',
    formTable10: '最大规格',
    formTable11: '最小规格',
    formTable12: '状态',
    formTable13: '联系方式',
    formTable14: '节点算力',
    formNotData: '空',
    ON: '开',
    OFF: '关'
  },
  deal: {
    upload_title: '创建您的离线任务',
    upload_form_name: '任务名称:',
    upload_form_desc: '说明:',
    upload_form_tags: '标签:',
    description: '任务描述:',
    curated_dataset: '数据集:',
    upload_form_type: '类型:',
    upload_form_file: 'CSV文件:',
    upload_form_file_tip: '文件最大为 25 GB',
    upload_form_file_tip01: '文件最小为 1 MB',
    upload_form_minerId: '存储节点ID:',
    upload_form_dataset: '数据集:',
    upload_form_minerId_tips: '存储节点ID不能为空',
    upload_form_miner: '目标存储提供者:',
    upload_form_btn: '附加文件',
    Submit: '提交',
    Cancel: '取消',
    save: '保存',
    saveSetting: '保存设置',
    form_name: '任务列表',
    form_select_title: '任务状态:',
    form_select_title01: '交易状态:',
    form_search: '高级搜索',
    task_search: '任务名称搜索...',
    task_dataset: '数据集搜索...',
    task_dataset_csv: '数据集CSV',
    formTable01: '任务名称',
    formTable02: '存储节点ID',
    formTable03: '任务状态',
    formTable04: '类型',
    formTable05: '纪录',
    formTable06: '描述',
    formTable07: '标签',
    formTable08: '创建日期（EST）',
    formTable09: '操作',
    formTable10: '客户名称',
    // formTable08: 'GENERATION TIME',
    formNotData: '空',
    detailTable01: '交易CID',
    detailTable02: '文件路径',
    detailTable03: '源URL',
    detailTable04: '状态',
    detailTable05: '起始',
    detailTable06: '上次更新',
    detailTable07: '创建日期',
    detailTable08: '纪录',
    detailTable09: '数据CID',
    backto: '返回',
    title_top_name: '任务名称',
    title_top_miner: '任务存储提供者',
    title_top_price: '价格',
    title_top_created: '创建于',
    title_top_Last_Update: '上次更新',
    delete: '删除',
    edit: '编辑',
    tags_tips: '最多输入5个最能描述任务的标记。 存储提供者将使用这些标签找到他们最感兴趣和最有经验的任务。',
    table_miner_title: '已分配任务列表',
    Clear_all: '全部清除',
    table_th_02: '分配日期（EST）',
    table_th_03: '指定存储提供者',
    table_th_04: '交易',
    table_th_05: '投标',
    table_th_06: '操作',
    table_th_07: 'MD5',
    table_th_08: 'MD5',
    module_deal_tips: '是否确实要将状态修改为 ',
    Validate: '验证',
    Tips: '提示',
    deal_not_found: '未找到离线交易。',
    proposal_not_null: '描述您的建议不能为空！',
    please_tip01: '请输入任务名称',
    please_tip02: '请输入目标存储提供者',
    file_error: '这个文件已经付款了。'
  },
  miner: {
    search_place: '搜索任务...',
    title_Type: '任务类型',
    title_Tags: '标签',
    title_Status: '状态',
    title_hasMiner: '已分配存储提供者',
    title_CreateDate: '创建日期',
    title_location: '具体位置',
    Reset_filters: '重置筛选器',
    bids_title: '发布人信息',
    bids_title01: '对此任务出价',
    tab_left: '作为客户',
    tab_right: '作为存储提供者',
    onle_one: '只有一个存储提供者',
    Miner: '存储提供者',
    Assign_Miner: '指派存储提供者'
  },
  stats: {
    network_overview: '网络概述'
  },
  my_profile: {
    left_menu_01: '简况',
    left_menu_02: '密码',
    left_menu_03: '支付',
    left_menu_04: '开发人员设置',
    left_menu_05: '工具',
    contact_Email: '电子邮件',
    Change_password: '修改密码',
    wrong_password: '密码错误',
    Reenter_new_password: '重新输入新密码',
    apiKey_your_title: 'API密钥',
    apiKey_your_Domain: '自定义域',
    apiKey_your_Domain_label: '域地址',
    apiKey_tips_01: 'API密钥提供对您的swan帐户的完全访问，因此请确保它们的安全。',
    apiKey_tips_02: '如何使用api密钥。',
    apiKey_tips_03: '保持API密钥安全的技巧。',
    apiKey_btn_01: '创建API密钥',
    apiKey_btn_02: '撤销',
    apiKey_btn_03: '创建自定义域',
    apiKey_btn_04: '删除',
    table_apiKey_th_01: '名称',
    table_apiKey_th_02: '密钥',
    table_apiKey_th_03: '访问令牌',
    table_apiKey_th_04: '创建日期',
    table_apiKey_th_05: '状态',
    create_api_title: '创建密钥',
    create_api_title01: '已成功创建API令牌',
    create_api_title02: '重要',
    create_api_tips: '名称不能为空！',
    create_api_tips01: '此令牌是您访问Filswan平台服务的私人机密。',
    create_api_tips02: '密钥只显示一次。在继续之前，您需要把它复制到安全的地方。',
    create_api_tips03: '您的新API密钥是：',
    create_api_tips04: '您的新访问令牌是：',
    create_api_tips05: '当您生成密钥时，您将看到一个带有Swan API密钥、Swan API Secret和JWT的模式。',
    create_api_tips06: '您的“Swan API密钥”充当我们restapi的公钥，您的“Swan Secret API密钥”充当您公钥的密码。JWT是两者的编码混合。一定要保密！',
    create_api_tips07: '为了增强客户的安全性，这些密钥在Swan一侧加密，并且只显示一次，所以请将它们写下来。如果丢失了这些值，则需要撤消密钥并创建一个新的密钥。',
    create_api_tips08: '关闭页面后，您将无法再次查看您的访问令牌。',
    Check_dataCap: '检查DataCap',
    form_contact: '联系方式',
    form_username: '用户名',
    form_ProfileHeadline: '简介标题',
    form_Summary: '摘要',
    list_title: '存储提供者名单',
    list_button: '管理',
    table_child_01: '存储节点ID',
    table_child_02: '提供离线交易',
    table_child_03: '摘要',
    pass_current: '当前密码',
    pass_new: '新密码',
    pass_confirm: '确认密码',
    Your_swan_miners: '您的swan存储提供者',
    swan_miners_add: '添加',
    swan_miners_Delete: '删除',
    swan_miners_Refresh: '刷新',
    miner_add_h1: '签名验证',
    miner_add_title: '目前仅对 有效算力 ≥ 32 GiB 的存储提供者开放自主认证',
    miner_add_OwnerAddress: 'Owner地址',
    miner_add_Contact: '联系方式',
    miner_add_contact_place: '邮箱/手机号...',
    miner_add_Message: '信息',
    miner_add_SignCode: '签名代码',
    miner_add_Signature: '签名',
    miner_add_sign_place: '请复制上面的代码，在filecoin钱包中进行签名，将签名内容粘贴到此处',
    miner_add_ms: '特别提示：账户名称不得使用敏感词、违禁词、名人姓名等，如有违规将进行删除和禁用签名处理，谢谢配合。',
    module_tips01: '是否要删除存储提供者 ',
    Executing: '执行...',
    module_tips02: '请选择存储提供者。',
    module_tips03: '是否要刷新存储提供者 ',
    module_tips04: '是否要删除API密钥？',
    module_tips05: '是否要更改信息 ',
    Success: '操作成功',
    Choose_Picture: '选择图片'
  },
  components: {
    choose_name: '为任务选择一个名称',
    choose_description: '任务描述',
    curated_dataset: '数据集',
    new_tag: '新标签',
    tag_description: '最多输入5个最能描述任务的标记。 存储提供者将使用这些标签找到他们最感兴趣和最有经验的任务。',
    type_Regular: '常规',
    type_Verified: '真实',
    Upload_CSV: '上传CSV',
    upload_desc: '您可以在',
    upload_desc01: '下载CSV模版。',
    upload_here: '这里',
    yes: '是',
    no: '否',
    Task_expected: '任务预计完成时间',
    day: '天',
    interval: '(最少: 180 - 最多: 540)',
    accepted: '从接受任务的那一刻算起。',
    upload_re_upload: '请重新上传.csv文件，包括deal cid和start epoch。',
    upload_Download: '下载示例。',
    Open_bid: '开放竞标',
    Estimated_budget: '预计费用',
    Estimated_budget_min: '最少',
    Estimated_budget_max: '最多',
    e_g_name: '例如Task001'
  },
  // 登录
  login: {
    login_loginphone: '手机登录',
    login_loginmail: '邮箱登录',
    login_phonenum: '请输入您的手机号',
    login_phonepw: '请输入您的密码',
    login_phonelogin: '登录',
    login_phonereg: '注册',
    login_phoneforget: '找回密码',
    login_mailnum: '请输入您的邮箱账号',
    login_mailpw: '请输入您的密码',
    login_maillogin: '登录',
    login_mailreg: '注册',
    login_mailforget: '找回密码',
    login_validate_title: '安全验证',
    login_validate_phone: '手机号',
    login_validate_phonecode: '请输入手机验证码',
    login_validate_phone_tips_true: '请输入正确的验证码',
    login_validate_google: '谷歌验证',
    login_validate_googlecode: '请输入谷歌验证码',
    login_validate_google_tips_true: '请输入正确的验证码',
    login_validate_btn: '确定',
    login_verify_phone_tips_true: '请输入正确的手机号',
    login_verify_phone_tips_empty: '手机号不能为空',
    login_verify_phonePassword_tips_true: '请输入正确的密码',
    login_verify_mail_tips_true: '请输入正确的邮箱',
    login_verify_mail_tips_empty: '邮箱不能为空',
    login_verify_mailPassword_tips_true: '请输入正确的密码',
    login_getVerifyCodeWord: '发送验证码',
    login_getVerifyCodeWord_time: 's后获取',
    login_verify_account_tips_no: '用户不存在',
    login_verify_pw_tips_empty: '密码不能为空',
    login_verify_code_tips_empty: '请输入正确的验证码',
    login_verify_pw_tips_rule: '密码至少为8-16个字符, 至少包含1个大写字母，1个小写字母和1个数字',
    login_validate_mailcode: '请输入邮箱验证码',
    login_validate_mail_tips_true: '请输入正确的验证码',
    login_validate_mail: '邮箱',
    login_verify_zh_tips_empty: '账号不能为空',
    login_h1: '没有账号，立即注册',
    login_h2: '欢迎登录'
  },
  // 注册（账号激活）
  accountActivation: {
    accountActivation_title: '账号激活',
    accountActivation_tips_1: '验证邮件已发送至',
    accountActivation_tips_2: '请登录邮箱点击激活链接，完成账号激活。',
    accountActivation_resend: '重新发送邮件',
    getVerifyCodeWord_time: 's后获取'
  },
  // 注册（激活验证）
  activationSuccess: {
    activationSuccess_success_title: '注册验证成功',
    activationSuccess_success_tips: '后跳至登录',
    activationSuccess_success_jumptime_unit: '秒',
    activationSuccess_fail_title: '激活链接已失效',
    activationSuccess_fail_tips: '请尝试重新登录获取激活邮件',
    activationSuccess_fail_tips01: '重新获取激活邮件'
  },
  // 注册
  register: {
    register_registerphone: '手机注册',
    register_registermail: '邮箱注册',
    register_phonenum: '请输入您的手机号',
    register_phonecheckcode: '请输入手机验证码',
    register_phonepw: '请输入登录密码',
    register_phonepwconfirm: '请确认登录密码',
    register_agree: '我已阅读并同意',
    register_agreelaw: '用户服务条款',
    register_phonereg: '注册',
    register_phonelogin: '登录',
    register_mailnum: '请输入您的邮箱',
    register_mailpw: '请输入登录密码',
    register_mailpwconfirm: '请确认登录密码',
    register_mailreg: '注册',
    register_maillogin: '登录',
    register_regsuccess: '注册成功',
    register_regsuccess_tips: '后跳至登录',
    register_verify_phone_tips_true: '请输入正确的手机号',
    register_verify_phone_tips_empty: '手机号不能为空',
    register_verify_verifyCode_tips_true: '请输入正确的验证码',
    register_verify_phonePassword_tips_true: '请输入正确的密码',
    register_verify_phonePasswordVerify_tips_true: '请输入正确的密码',
    register_verify_term_tips_true: '请勾选服务条款',
    register_verify_mail_tips_true: '请输入正确的邮箱',
    register_verify_mail_tips_empty: '邮箱不能为空',
    register_verify_mailPassword_tips_true: '请输入正确的密码',
    register_verify_mailPasswordVerify_tips_true: '请输入正确的密码',
    register_getVerifyCodeWord: '获取验证码',
    register_getVerifyCodeWord_time: 's后获取',
    register_verify_account_tips_exist: '用户已存在',
    register_verify_pw_tips_empty: '密码不能为空',
    register_verify_pw_tips_rule: '密码至少为8-16个字符, 至少包含1个大写字母，1个小写字母和1个数字',
    register_verify_placehold_tips_rule: '请输入密码',
    register_verify_pw_tips_reenter: '请再次输入密码',
    register_verify_pw_tips_same: '两次密码输入需保持一致',
    register_jumptime_unit: '秒',
    register_h1: '已有账号，立即登录',
    register_h2: '邀请者邀请码（选填）'
  },
  // 找回密码
  forgetPassword: {
    forgetPassword_title: '找回密码',
    forgetPassword_phonenum: '请输入您的手机号',
    forgetPassword_phonecheckcode: '请输入手机验证码',
    forgetPassword_phonepw: '请输入登录密码',
    forgetPassword_phonepwconfirm: '请确认登录密码',
    forgetPassword_phonebtn: '确认',
    forgetPassword_mailget: '用邮箱找回密码',
    forgetPassword_mailbtn: '确认',
    forgetPassword_phoneget: '用手机找回密码',
    forgetPassword_mailnum: '请输入您的邮箱',
    forgetPassword_verify_phone_tips_true: '请输入正确的手机号',
    forgetPassword_verify_phone_tips_empty: '手机号不能为空',
    forgetPassword_verify_phonePassword_tips_true: '请输入正确的密码',
    forgetPassword_verify_phonePasswordVerify_tips_true: '请输入正确的密码',
    forgetPassword_verify_mail_tips_true: '请输入正确的邮箱',
    forgetPassword_verify_mail_tips_empty: '邮箱不能为空',
    forgetPassword_getVerifyCodeWord: '获取验证码',
    forgetPassword_getVerifyCodeWord_time: 's后获取',
    forgetPassword_verify_verifyCode_tips_true: '请输入正确的验证码',
    forgetPassword_verify_account_tips_no: '用户不存在',
    forgetPassword_verify_pw_tips_empty: '密码不能为空',
    forgetPassword_verify_pw_tips_rule: '密码至少为8-16个字符, 至少包含1个大写字母，1个小写字母和1个数字',
    forgetPassword_verify_pw_tips_same: '两次密码输入需保持一致',
    forgetPassword_success_title: '密码重置成功！',
    forgetPassword_success_btn: '登录',
    forgetPassword_h1: '已有账号，立即登录',
    forgetPassword_reenter: '再次输入密码',
    forgetPassword_submit: '提交',
    forgetPassword_phone: '手机找回',
    forgetPassword_mail: '邮箱找回'
  },
  // 找回密码（重置邮件）
  mailForget: {
    mailForget_title: '账号密码重置',
    mailForget_tips_1: '验证邮件已发送至',
    mailForget_tips_2: '，请登录邮箱点击重置密码链接，完成密码重置。',
    mailForget_resend: '重新发送邮件',
    getVerifyCodeWord: '重新发送邮件',
    getVerifyCodeWord_time: 's后获取'
  },
  // 找回密码（重置密码）
  mailResetPassword: {
    mailResetPassword_title: '重置密码',
    mailResetPassword_pw: '请输入登录密码',
    mailResetPassword_new_pw: '请输入新登录密码',
    mailResetPassword_pwconfirm: '请确认登录密码',
    mailResetPassword_btn: '确认',
    mailResetPassword_verify_pw_tips_true: '请输入正确的密码',
    mailResetPassword_verify_pwconfirm_tips_true: '请输入正确的密码',
    mailResetPassword_verify_pw_tips_rule: '密码至少为8-16个字符, 至少包含1个大写字母，1个小写字母和1个数字',
    mailResetPassword_verify_pw_tips_same: '两次密码输入需保持一致',
    mailResetPassword_link_tips: '重置密码链接有误'
  },
  // 找回密码（重置密码成功）
  mailForgetSuccess: {
    mailForgetSuccess_title: '密码重置成功',
    mailForgetSuccess_tips: '后跳至登录',
    mailForgetSuccess_jumptime_unit: '秒'
  },
  fs3: {
    title: '通用存储服务',
    access_key: 'Access Key',
    secret: 'Secret',
    URL: 'URL',
    Entry_btn: '前往FS3',
    copy: '复制',
    copied: '已复制',
    connected_to: '连接到',
    Connect_Wallet: '连接钱包',
    Disconnect: '登出'
  },
  fs3Login: {
    toptip_01: '目前支持',
    toptip_02: '；额外的链条即将到来！',
    toptip_03: '连接到您的MetaMask钱包',
    toptip_04: '如果您想访问MCS Polygon Mainnet，请单击',
    toptip_04_1: 'Testnet，请单击',
    toptip_04_main: '如果您想访问MCS',
    toptip_Network: '或',
    title: '连接您的钱包。',
    MetaMask: 'MetaMask',
    MetaMask_tip: '请将您的钱包连接到Mumbai Testnet或BSC TestNet。',
    Account: '账户',
    Connected_MetaMask: '连接到MetaMask',
    Connected_Email: '连接到Email',
    Connected_Email_Address: '使用您的邮箱地址连接',
    View_explorer: '在浏览器查看',
    Copy_Address: '复制钱包地址',
    Disconnect_mailbox: '断开与邮箱的连接',
    Change_Address: '更改您的邮箱地址',
    Copied: '已复制',
    Connect_to_MetaMask: '连接到MetaMask',
    Connect_cont_Wallet: '连接钱包',
    Connect_cont_tip: '连接您的钱包地址以创建多链存储帐户。',
    Connect_text: '跨区块链分散Web3存储',
    Connect_text_desc: '在一个位置轻松地将单个文件或Bucket存储上传到IPFS和Filecoin。获取您的CID，然后只需单击一次即可进入OpenSea。',
    Connect_text_btn: '使用多链存储享受旅程',
    Connected: '已连接',
    Connect_form_label: '请在此处输入您的电子邮件地址',
    Connect_form_label_change: '在此处输入您的新电子邮件地址',
    Connect_checkbox: '我已阅读并同意 ',
    Connect_checkbox_1: '隐私政策',
    Connect_checkbox_2: ' 以及接收时事通讯。',
    Connect_form_btn: '提交',
    Connect_form_btn_change: '更换',
    Connect_email_title: '成功',
    Connect_email_desc: '查看您的电子邮件并点击链接激活您的帐户奖励。',
    Connect_email_again: '没有收到任何电子邮件？再试一次。',
    Connect_email_jump: '单击此处跳转到页面。',
    Connect_StartFree: '入门',
    Connect_BuyNow: '立即购买',
    Connect_TutorialVideo: '教程视频',
    back_home: '返回主页',
    Skip: '跳过',
    Connect_tip: '确认断开与邮箱的连接？'
  },
  uploadFile: {
    title: '文件列表',
    contract_id: '合约ID',
    contract_id01: '合约ID',
    cid: 'CID',
    status: 'PIN状态',
    status_tooltip: '文件或数据在FilSwan IPFS节点上的存储状态。',
    status_button_tooltip: '点击unpin后即不可修改。',
    file_name: '文件名称',
    file_name_tooltip: '请点击文件名以查阅更多交易详情和DAO签名详情。',
    file_name_tip: '请选择文件',
    file_size: '文件大小',
    file_status: '状态',
    upload_title: '任务信息',
    task_name: '任务名称',
    w3ss_id: '存储提供者',
    w3ss_ids: '存储提供者',
    w3ss_id_tooltip: '为Filecoin网络提供存储容量的存储提供者。',
    w3ss_id_nothing: '排队中',
    w3ss_id_nothing_tooltip: '正在整合本次上传的数据，以备存储到Filecoin。Filecoin交易会在上传后48小时内完成。',
    create_time: '创建时间',
    upload_time: '上传时间',
    deal_id: '交易ID',
    payment: '支付',
    MINT: '铸造',
    mint_view: '查看',
    Canceled: '已取消',
    pay: '支付',
    failed: '失败',
    paid: '已支付',
    refund: '退款',
    download_link: '下载链接',
    uploadDray: '拖曳文件拖到此处',
    uploadDray_or: '或',
    uploadDray_text: '浏览文件',
    upload: '上传文件',
    upload_form_file_tip: '上传文件不能为空',
    search_y: '按CID搜索',
    search_input: '输入值',
    deal_title: '选定批次的 Filecoin 交易（待定交易的 deal_id = 0）：',
    Live_deal: '实时交易',
    live_deal_btn: '查看交易',
    live_deal_btn_no: '没有链接：错误',
    Upload_More_Files: '上传文件',
    topTip: '请点击存储供应商ID以查阅交易详情和DAO详情',
    search_title: '按文件名称搜索',
    data_tooltip: '文件或数据块的内容标识符。',
    payment_tip: '请等待资金锁定流程完成。',
    payment_tip_deal: '请耐心等待该过程完成。',
    payment_tip_deal01: '您已上传成功。',
    payment_tip_deal02: '请耐心等待支付完成。',
    payment_tip_deal03: '您已成功创建集合。',
    COMPLETED: '转账成功',
    SUCCESS: '您已转账成功，请查看转账历史记录。',
    CLOSE: '关闭',
    Fail: '转账失败',
    FailTIP: '您的交易失败了。有关更多细节，请查看交易历史记录。',
    waiting: '等待中...',
    waitingTIP: '请稍后在我的文件中确认文件和支付状态。有关更多细节，请查看交易历史纪录。',
    until: '请等待任务分配给存储提供者。',
    OK: '确定',
    Deal_Detail: '交易详情',
    detail_tip01: '这意味着您的交易还没有被存储服务商接受。',
    no_fund_locked: '没有基金锁定。',
    Successfully_unlocked_funds: '成功解锁基金。',
    Successfully_signed: '等待解锁资金。',
    Successfully_signed_all: '成功签名，',
    Waiting_for_signature: '等待签名',
    detail_IPFSDownload: 'IPFS下载',
    detail_Network: '网络',
    detail_Locked_funds: '锁定资金',
    detail_Storage_Price: '存储价格',
    detail_ProposalCID: '订单CID',
    detail_MessageCID: '消息CID',
    detail_PieceCID: '分片CID',
    detail_Client_Address: '客户地址',
    detail_Verified_Deal: '真实数据',
    detail_Storage_Price_Per_Epoch: '每Epoch存储价格',
    detail_Signature_Type: '签名类型',
    detail_Retrieval_Filecoin: '从Filecoin网络取回',
    detail_Retrieval_Filecoin_tooltip: '该交易在发布区块链上之前不能取回。本命令末尾的‘output-file’是您想要保存的文件名，您也可以为这个变量添加一个路径。',
    detail_DAO_Signatures: 'DAO签名',
    detail_DAO_Signatures_tooltip: '签名是用来解锁资金给存储提供者。',
    detail_DAO_RKH_Address: 'DAO地址',
    detail_Time: '日期',
    view_deal_logs: '查看交易日志',
    deal_logs: '交易日志',
    no_logs: '无日志',
    uploadFile_title: '请上传您的文件，系统将为您计算预计存储费用。',
    upload_funds: '资金不足？点击此处获取更多测试网代币。',
    Duration: '存储时间',
    Duration_tip: '请输入存储时间',
    Duration_tooltip: '存储时间是指您希望文件在Filecoin网络上存储的期限。',
    Storage_copy: '存储份数',
    Storage_copy_tooltip: '将文件存储到Filecoin网络的份数。您的文件将会被分配到不同的存储提供者。',
    Estimated_Storage_Cost: '预计存储费用',
    Estimated_Storage_Cost_tooltip: '预计存储费用是根据文件大小，存储时间，和平均数据存储价格计算的。',
    Free_Storage_Capacity: '免费储存空间配额',
    Free_Storage_Capacity_tooltip: '每月享有10GB免费储存空间配额。',
    Select_Lock_Funds_Plan: '选择资金锁定计划',
    Select_Lock_Funds_Plan_tooltip: '锁定的资金越多，你的文件就会越早存储在Filecoin网络上。超额支付的资金将在Filecoin交易上链后自动返回。',
    latest_exchange_rate: 'USDC兑FIL的最近汇率是',
    Low: '低',
    Average: '中等',
    High: '高',
    File_uploading: '文件上传',
    File_uploading_tooltip: '您的文件还在上传到IPFS的过程中。请保持此窗口打开，直到上传完成。',
    file_uploaded: '该文件已经上传。',
    file_uploaded_tip: '我们目前只支持上传一次文件。请选择另一个要上传的文件。',
    file_uploaded_tip01: '谢谢您的理解。',
    file_uploaded_tip02: '这个文件你不用付钱，它会加到你的文件里的。因为我们目前只支持上传一次文件。',
    update_time: '更新时间:',
    nft_title: '创建您的',
    nft_Name: '名称',
    nft_Amount: '数量',
    nft_Description: '描述',
    nft_External: '外部链接',
    nft_Seller: '卖方费用基点',
    nft_Fee_Recipient: '费用接收人',
    nft_Image: '头像',
    nft_IPFSURL: 'IPFS链接',
    nft_list_empty: '此列表为空',
    Payment_Transaction_Hash: '支付交易哈希',
    Back: '关闭',
    Minting: '铸造中…',
    Mint_NFT: '铸造',
    Mint_Create_NFT: '创建新的NFT集合',
    View_Your_NFT: '查看您的NFT',
    View_Your_NFT_tips: '您的NFT已经成功铸造！',
    View_Your_NFT_tips01: '您可以在这里查看该交易：',
    View_Your_NFT_OpenSea: '点击到OpenSea查看您的NFT',
    View_Your_NFT_Note: '注意：OpenSea需要一些时间加载您的NFT。',
    goTo: '前往',
    goTopage: '页',
    xhr_tip: '上传文件已更改，请重新上传。',
    xhr_error500: '创建交易失败，请联系管理员。',
    filter_minted: '已铸造',
    filter_no_minted: '未铸造',
    filter_status_Unpaid: '未支付',
    filter_status_Paid: '已支付',
    filter_status_Refunding: '可退款',
    filter_status_Refundable: '可退款',
    filter_status_Refunded: '已退款',
    filter_status_Unlocked: '解锁',
    filter_status_Success: '完成',
    filter_status_Pending: '待支付',
    filter_status_Processing: '处理中',
    filter_status_Free: '免费',
    filter_status_Failed: '失败',
    bucket_Uploading: '正在上传...',
    bucket_Upload: '上传',
    Start_uploading: '开始上传',
    Upload_failed: '上传失败:',
    Upload_view: '查看',
    Upload_path: '路径',
    mint_pending: '正在从链上扫描collection信息，请等待!'
  },
  billing: {
    search_placeholder: '搜索方式',
    search_option_filename: '文件名',
    search_option_transaction: '交易',
    search_btn: '搜索',
    clear_btn: '清除',
    TRANSACTION: '交易哈希',
    AMOUNT: '锁定金额',
    UNLOCKAMOUNT: '解锁金额',
    TOKEN: '代币',
    FILENAME: '文件名称',
    PAYLOADCID: '数据CID',
    deals_price: '价格',
    WALLET: '钱包',
    NETWORK: '网络',
    PAYMENTDATE: '付款日期',
    UNLOCKDATE: '解锁日期',
    Deadline: '截止日期',
    download_module_title: '下载数据（',
    download_module_title_kh: '）',
    download_module_title_file: '文件详情',
    download_module_title_billing: '账单详情',
    download_module_title_tooltip: '下载时间依据数据大小而调整',
    download_module_btn: '下载',
    verify: '验证:',
    verify_result: '验证成功',
    verify_tip: '将滑块拖动到最右侧',
    end_date: '结束日期',
    time_to: '至',
    data_download: '导出',
    data_export: '格式',
    bill_overview: '计费概述',
    bill_btn_pay: '升级为付费计划',
    bill_btn_paid: '您已成为付费用户',
    bill_cont1_title: '账单详情',
    bill_cont1_desc: '查看过去和当前的发票',
    bill_cont2_title: '定价',
    bill_cont2_desc: '查看定价和常见问题解答',
    have_faq: '常见问题?',
    bill_tip: '事务处理可能会延迟。请在确认成功后刷新此页面。请不要重复付款，因为这可能会造成经济损失。'
  },
  network: {
    title: 'All Chains'
  },
  footer: {
    User_Start_Guide: '用户入门指南',
    Developer_Quick_Start: '开发人员快速入门',
    FAQ: '常见问题解答',
    Get_Help: '获取帮助',
    copy: ' MULTICHAIN.STORAGE. ALL RIGHTS RESERVED.'
  },
  metaSpace: {
    home_title: '多链存储',
    home_Introduction: '介绍',
    home_Introduction_cont: '多链存储是一个基于智能合约的跨链存储网关，与Oracle技术和Filecoin网络集成。它通过桥接多个区块链网络，加速了去中心化存储的大规模采用。',
    home_Our_Features: '我们的特点',
    home_Our_Features_cont01: '拖放文件以固定到IPFS并同时存储在Filecoin网络上',
    home_Our_Features_cont02: '通过可靠的边缘网络实现全球快速检索',
    home_Our_Features_cont03: 'JS、Python和Golang SDK的开发者支持',
    empty_prompt: '您还没有Bucket存储空间！',
    empty_prompt_detail: '开始上传文件。',
    search_bucket: '搜索Bucket存储',
    bucket_name: 'Bucket名称',
    bucket_subname: '子Bucket名称',
    folder_name: '文件夹名称',
    add_bucket: '添加Bucket',
    add_subbucket: '添加子Bucket',
    create_folder: '创建文件夹',
    create_bucket: '创建存储桶',
    Submit: '提交',
    Cancel: '取消',
    Delete: '删除',
    Close: '关闭',
    list_bucket: 'Bucket列表',
    list_bucket_tip: '活跃Bucket存储总量/创建的Bucket存储总数',
    table_name: '名称',
    table_type: '类型',
    table_subdomain: '子域',
    table_space: '空间',
    table_size: '大小',
    table_createon: '创建时间',
    table_cid: 'CID',
    table_cid_tip: '文件或存储桶的内容ID',
    table_status: '状态',
    table_backup_status: '备份状态',
    table_LastModified: '上次修改时间',
    table_action: '操作',
    delete_title: '删除Bucket',
    delete_title_detail: '是否永久删除？',
    delete_desc: '此Bucket项目将被永久删除。此操作无法反转。',
    Upload_btn: '上传',
    Upload_File: '上传文件',
    Upload_Folder: '上传文件夹',
    Browse_Folders: '浏览文件夹',
    Browse_Folders_title: '选择文件夹的存储位置：',
    Browse_Folders_title01: '名称',
    Browse_Folders_mcs_desc: '*该文件夹将存储在MCS系统中，以后可以编辑但不能共享。可以查看和共享文件夹中的文件。',
    Browse_Folders_desc: '*该文件夹将永久存储在IPFS系统中，可以稍后共享，但无法编辑。',
    upload_desc: '将文件拖放到此处',
    uploadDray_or: '或',
    uploadDray_text: '浏览文件',
    detail_title: 'Bucket详细信息',
    detail_folder_title: '文件夹详细信息',
    detail_BucketName: 'Bucket名称:',
    detail_folderName: '文件夹名称:',
    detail_DateCreated: '创建日期:',
    detail_LastModified: '上次修改时间:',
    detail_CurrentFiles: '当前文件:',
    detail_CurrentSize: '当前大小:',
    detail_RemainingServiceDays: '剩余服务天数:',
    detail_BackupInfo: '备份信息:',
    detail_StorageProvider: '存储提供程序 ',
    detail_PieceCID: '块CID:',
    ob_detail_share: '分享',
    ob_detail_title: '对象详细信息',
    ob_detail_ObjectName: '对象名称:',
    ob_detail_DateUploaded: '上传日期:',
    ob_detail_ObjectSize: '对象大小:',
    ob_detail_ObjectIPFSLink: '对象IPFS链接:',
    ob_detail_ObjectCID: '对象CID:',
    Pay: '支付',
    pay_title: '你没有足够的信用卡。',
    pay_tip: '*通过购买新的Bucket存储，您将使用当前钱包地址增加“可用Bucket存储”。',
    pay_tip1: '*Bucket存储的最大数量为33。',
    pay_body_BucketName: '命名您的Bucket',
    pay_body_Bucket: 'Bucket',
    pay_body_Bucket_desc: '每个存储桶有30GB的存储空间',
    pay_body_PriceBucket: '价格/桶',
    pay_body_PriceBucket_desc: '每个存储桶的服务从存储桶创建之日开始。',
    pay_body_TotalPayment: '付款总额',
    pay_load: '请稍候',
    pay_success: '付款成功！',
    pay_success_desc: '感谢您的支付。已添加桶信用。',
    try_again: '再试一次',
    Name_bucket: '命名此存储桶',
    active_bucket: '你必须先激活你的存储桶'
  },
  apiKey: {
    day: '天',
    label_Expiration: '到期（天）',
    label_tip01: '* 输入0表示API密钥永久处于活跃状态',
    label_tip02: '* 最大数量为365',
    Add_bucket: '添加此存储桶',
    apikey_cont: 'APIKey将被永久删除。此操作无法撤消。'
  },
  comment: {
    Rate_Product: '评价我们的产品',
    Tell_Comment: '告诉我们您的意见',
    theme: '主题'
  },
  pricing: {
    'title_1': 'Choose the plan that works for you',
    'title_2': 'Start storing your data in web3.',
    'title_3': 'Stay tuned and contact us for more service details. ',
    'title_4': 'In the meantime, we invite you to contact us with your proposals, feedback, or any inquiries.',
    'btn': 'Contact us',
    'bid_title_1': 'Free',
    'bid_title_2': 'Pro',
    'bid_desc_1': 'Get started with free essentials',
    'bid_desc_2': 'Empower your storage with higher limits',
    'get_plan': 'Get a Plan',
    'coming_soon': 'Coming soon...',
    'supported_wallet': 'Supported wallet',
    'prepay_using': 'Prepay using',
    'volume_title': 'Looking for more storage or volume?',
    'volume_title_1': 'How to create a Discord ticket?',
    'volume_p': "Our enterprise solutions and tailored plans are here to help you reach your goals. Contact us now via Discord and let's explore how we can unlock your storage potential together.",
    'volume_create': 'Gif of how to create a ticket in a discord',
    'question_title': 'Frequently asked questions',
    'collapse_1_title': 'Does MultiChain Storage provide a free version?',
    'collapse_1_desc': 'Yes, our free plan provides essential features and suits any individual to get started with decentralized storage',
    'collapse_2_title': 'How many buckets and folders can I create?',
    'collapse_2_desc': 'You can create unlimited buckets and folders as soon as you do not go over your storage capacity',
    'collapse_3_title': 'Do you provide with SDK?',
    'collapse_3_desc': 'Yes! We provide SDKs for Python and JavaScript. You can check them here (Link)',
    'collapse_4_title': 'What if I run out of limits in Pro plan?',
    'collapse_4_desc': 'No worries! Our team always open to communicate and customize a special plan for your needs. Create a ticket in our Discord community and we will get back to you shortly after.'
  }
}
