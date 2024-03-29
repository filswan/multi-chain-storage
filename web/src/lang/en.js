export default {
  route: {
    dashboard: 'Dashboard',
    Deal: 'Onchain Storage',
    Miner: 'Storage Provider Management',
    myProfile: 'My Profile',
    browse: 'Browse Tasks',
    Tools: 'Tools',
    My_FS3: 'FS3',
    Stats: 'Statistics',
    myAccount: 'Setting',
    Upload_files: 'Upload File',
    Search_file: 'Search File',
    documentation: 'Document',
    metaSpace: 'Bucket Storage',
    About: 'About',
    Pricing: 'Pricing',
    Resources: 'Resources',
    Login: 'Login',
    ApiKey: 'API Key',
    report: 'Audit Report'
  },
  navbar: {
    sidebar_header: 'Multi-Chain Storage',
    section: 'Section',
    myprofile: 'My Profile',
    support: 'Support',
    hi: 'Hi, ',
    logout: 'Logout',
    log_in: 'Log In',
    sign_up: 'Sign Up',
    dashboard: 'DASHBOARD',
    minerDetail: 'STORAGE PROVIDER DETAIL',
    deal: 'Details',
    miner: 'STORAGE PROVIDER MANAGEMENT',
    deal_details: 'TASK DETAIL',
    My_Profile: 'MY PROFILE',
    Browse_tasks: 'BROWSE TASKS',
    create_tasks: 'CREATE TASK',
    Upload_files: 'UPLOAD FILE',
    Tools: 'TOOLS',
    Stats: 'STATISTICS',
    copy: '© 2021 Swan Canada',
    ScoreChange: 'Score Change',
    ScoreChange_legend: 'Score',
    PowerChange: 'Power Change',
    MiningData: 'Storage Sealing Data',
    selectDate: {
      sevenDays: '7days',
      thirtyDays: '30days',
      oneYearL: '1year',
      roundClock: '24hours'
    },
    Power: 'Adj.Storage Power',
    PowerGrowth: 'Adj.Storage Power Growth',
    BillingHistory: 'Billing History',
    language: 'Language',
    language_cn: 'Chinese',
    swan_share: 'swan share'
  },
  // 首页
  dashboard: {
    text01: 'Total Storage Provider',
    text02: 'Active Storage Provider',
    form_name: 'Filecoin Storage Provider',
    form_swith: 'Accept Offline Deal',
    form_swith_in: 'Online Deal',
    search: 'W3SS ID Search...',
    formTable01: 'LOCATION',
    formTable02: 'W3SS ID',
    formTable03: 'PRICE (GiB/epoch)',
    formTable04: 'PRICE (32GiB/year)',
    formTable05: 'PRICE VERIFIED (GiB/epoch)',
    formTable06: 'SIZE',
    formTable07: 'OFFLINE DEALS',
    formTable08: 'SCORE',
    formTable09: 'LAST UPDATE',
    formTable10: 'MAX PIECE SIZE',
    formTable11: 'MIN PIECE SIZE',
    formTable12: 'STATUS',
    formTable13: 'CONTACT INFO',
    formTable14: 'ADJUSTED POWER',
    formNotData: 'No Data',
    ON: 'ON',
    OFF: 'OFF'
  },
  deal: {
    upload_title: 'Create Your Task for Offline Sealing',
    upload_form_name: 'Task Name:',
    upload_form_desc: 'Description:',
    upload_form_tags: 'Tags:',
    description: 'Task Description:',
    curated_dataset: 'Curated Dataset:',
    upload_form_type: 'Type:',
    upload_form_file: 'CSV File:',
    upload_form_file_tip: 'The maximum size of the file is 25 GB',
    upload_form_file_tip01: 'The minimum value of the file is 1 MB',
    upload_form_minerId: 'W3SS ID:',
    upload_form_dataset: 'Curated Dataset',
    upload_form_minerId_tips: 'W3SS ID cannot be null',
    upload_form_miner: 'Target Storage Provider:',
    upload_form_btn: 'Attach File',
    Submit: 'Submit',
    Cancel: 'Cancel',
    save: 'Save',
    saveSetting: 'Save Settings',
    form_name: 'Tasks List',
    form_select_title: 'Task Status:',
    form_select_title01: 'Deal Status:',
    form_search: 'Advanced Search',
    task_search: 'Task Name Search...',
    task_dataset: 'Dataset Search...',
    task_dataset_csv: 'Dataset CSV',
    formTable01: 'TASK NAME',
    formTable02: 'W3SS ID',
    formTable03: 'TASK STATUS',
    formTable04: 'TYPE',
    formTable05: 'NOTE',
    formTable06: 'DESCRIPTION',
    formTable07: 'TAGS',
    formTable08: 'DATE CREATE (EST)',
    formTable09: 'ACTION',
    formTable10: 'CLIENT NAME',
    // formTable08: 'GENERATION TIME',
    formNotData: 'No Data',
    detailTable01: 'DEAL CID',
    detailTable02: 'file_path',
    detailTable03: 'SOURCE URL',
    detailTable04: 'STATUS',
    detailTable05: 'START EPOCH',
    detailTable06: 'LAST UPDATED',
    detailTable07: 'DATE CREATED',
    detailTable08: 'NOTE',
    detailTable09: 'DATA CID',
    backto: 'Go Back',
    title_top_name: 'Task Name',
    title_top_miner: 'Task Storage Provider',
    title_top_price: 'Price',
    title_top_created: 'Created At',
    title_top_Last_Update: 'Last Update',
    delete: 'delete',
    edit: 'Edit',
    tags_tips: 'Enter up to 5 tags that best describe your task. Web3 Storage Provider (W3SP) will use these tags to find task they are most interested and experienced in.',
    table_miner_title: 'Assigned Tasks List',
    Clear_all: 'Clear All',
    table_th_02: 'DATE ASSIGNED (EST)',
    table_th_03: 'ASSIGNED STORAGE PROVIDER',
    table_th_04: 'DEALS',
    table_th_05: 'BIDS',
    table_th_06: 'ACTION',
    table_th_07: 'MD5',
    table_th_08: 'MD5_ORIGIN',
    module_deal_tips: 'Are you sure you want to modify the status to ',
    Validate: 'Validate',
    Tips: 'Tips',
    deal_not_found: 'Offline deal not found.',
    proposal_not_null: 'Describe your proposal cannot be null!',
    please_tip01: 'Please enter a task name',
    please_tip02: 'Please enter W3SS ID',
    file_error: 'This file has been paid.'
  },
  miner: {
    search_place: 'Search for Tasks...',
    title_Type: 'Task Type',
    title_Tags: 'Tags',
    title_Status: 'Status',
    title_hasMiner: 'Has Storage Provider',
    title_CreateDate: 'Create Date',
    title_location: 'Specific Location',
    Reset_filters: 'Reset filters',
    bids_title: 'About the Poster',
    bids_title01: 'Place a Bid on this Task',
    tab_left: 'As Client',
    tab_right: 'As Storage Provider',
    onle_one: 'Only one Storage Provider',
    Miner: 'Storage Provider',
    Assign_Miner: 'Assign Storage Provider'
  },
  stats: {
    network_overview: 'Network overview'
  },
  my_profile: {
    left_menu_01: 'Profile',
    left_menu_02: 'Password',
    left_menu_03: 'Payment',
    left_menu_04: 'Developer Settings',
    left_menu_05: 'Tools',
    contact_Email: 'Email',
    Change_password: 'Change Password',
    wrong_password: 'wrong password',
    Reenter_new_password: 'Reenter New Password',
    apiKey_your_title: 'API keys',
    apiKey_your_Domain: 'Custom Domain',
    apiKey_your_Domain_label: 'Domain address',
    apiKey_tips_01: 'API keys provide full access to your swan account, so keep them safe.',
    apiKey_tips_02: 'How to use API key.',
    apiKey_tips_03: 'Tips on keeping API keys secure.',
    apiKey_btn_01: 'Create API Key',
    apiKey_btn_02: 'Revoke',
    apiKey_btn_03: 'Create Custom Domain',
    apiKey_btn_04: 'Remove',
    table_apiKey_th_01: 'NAME',
    table_apiKey_th_02: 'KEY',
    table_apiKey_th_03: 'ACCESS TOKEN',
    table_apiKey_th_04: 'DATE CREATED',
    table_apiKey_th_05: 'STATUS',
    create_api_title: 'Create an API key',
    create_api_title01: 'API token successfully created',
    create_api_title02: 'Important',
    create_api_tips: 'Name cannot be null!',
    create_api_tips01: 'This token is your private secret to access Swan platform services.',
    create_api_tips02: "The secret key is only displayed once. You'll need to copy it somewhere safe before continuing.",
    create_api_tips03: 'Your new API Key is：',
    create_api_tips04: 'Your new Access Token is：',
    create_api_tips05: 'When you generate your keys, you will see a modal with the Swan  API Key, Swan API Secret, and the JWT.',
    create_api_tips06: 'Your "Swan API Key" acts as your public key for our REST API, and your "Swan Secret API Key" acts as the password for your public key. The JWT is an encoded mix of the two. Be sure to keep your secret key private!',
    create_api_tips07: "For added customer security, these keys are encrypted on Swan's side and will only ever be displayed once, so write them down. If you lose these values you'll need to revoke the key and create a new one.",
    create_api_tips08: 'You will not be able to view your access token again when you close the page.',
    Check_dataCap: 'Check DataCap',
    form_contact: 'Contact Info',
    form_username: 'User Name',
    form_ProfileHeadline: 'Profile Headline',
    form_Summary: 'Summary',
    list_title: 'Storage Provider List',
    list_button: 'Manage',
    table_child_01: 'W3StorageSystemID',
    table_child_02: 'Offline Deal Available',
    table_child_03: 'Summary',
    pass_current: 'Current Password',
    pass_new: 'New Password',
    pass_confirm: 'Confirm Password',
    Your_swan_miners: 'Your Swan Storage Providers',
    swan_miners_add: 'Add',
    swan_miners_Delete: 'Delete',
    swan_miners_Refresh: 'Refresh',
    miner_add_h1: 'Signature verification',
    miner_add_title: 'Only open claim for accounts with storage power ≥ 32 GiB',
    miner_add_OwnerAddress: 'Owner address',
    miner_add_Contact: 'Contact',
    miner_add_contact_place: 'Email, Cell phone...',
    miner_add_Message: 'Message',
    miner_add_SignCode: 'Sign code',
    miner_add_Signature: 'Signature',
    miner_add_sign_place: 'Please copy the above sign code, sign it with Filecoin wallet, and enter the signature',
    miner_add_ms: ' Special notes: sensitive words, prohibited words, celebrity names, etc. are not allowed in account names. If there is any violation, the signature will be deleted and prohibited. Thank you for your cooperation. ',
    module_tips01: 'Do you want to delete Storage Provider ',
    Executing: 'Executing...',
    module_tips02: 'Please select a Storage Provider.',
    module_tips03: 'Do you want to refresh Storage Provider ',
    module_tips04: 'Do you want to delete API key?',
    module_tips05: 'Do you want to change the information for ',
    Success: 'Success',
    Choose_Picture: 'Choose Picture'
  },
  components: {
    choose_name: 'Choose a name for your task',
    choose_description: 'Task Description',
    curated_dataset: 'Curated Dataset',
    new_tag: 'New Tag',
    tag_description: 'Enter up to 5 tags that best describe your task. Web3 Storage Provider (W3SP) will use these tags to find task they are most interested and experienced in.',
    type_Regular: 'Regular',
    type_Verified: 'Verified',
    Upload_CSV: 'Upload CSV',
    upload_desc: 'You can download CSV template from',
    upload_desc01: '',
    upload_here: ' here',
    yes: 'Yes',
    no: 'No',
    Task_expected: 'Task Expected Complete in',
    day: 'days',
    interval: '(min: 180 - max: 540)',
    accepted: 'Counting from the moment the task is accepted.',
    upload_re_upload: 'Please re-upload the .csv file including deal cid and start-epoch.',
    upload_Download: 'Download Example.',
    Open_bid: 'Open bid',
    Estimated_budget: 'Estimated budget',
    Estimated_budget_min: 'Min',
    Estimated_budget_max: 'Max',
    e_g_name: 'e.g.Task001'
  },
  login: {
    login_loginphone: 'Phone',
    login_loginmail: 'Email',
    login_phonenum: 'Mobile number',
    login_phonepw: 'Password',
    login_phonelogin: 'Log In ',
    login_phonereg: 'Sign Up',
    login_phoneforget: 'Forgot Password?',
    login_mailnum: 'Email',
    login_mailpw: 'Password',
    login_maillogin: 'Log in ',
    login_mailreg: 'Sign Up',
    login_mailforget: 'Forgot password?',
    login_validate_title: 'Security Verification',
    login_validate_phone: 'Mobile number ',
    login_validate_phonecode: 'Verification Code',
    login_validate_phone_tips_true: 'Please enter valid verification code！',
    login_validate_google: 'Google Authentication Code',
    login_validate_googlecode: 'Google Authentication Code',
    login_validate_google_tips_true: 'Please enter valid verification code！',
    login_validate_btn: 'OK',
    login_verify_phone_tips_true: 'Please enter valid phone number',
    login_verify_phone_tips_empty: 'The phone number shall not be null！',
    login_verify_phonePassword_tips_true: 'Please enter correct password',
    login_verify_mail_tips_true: 'Please enter valid email ',
    login_verify_mail_tips_empty: 'Please enter an email address.',
    login_verify_mailPassword_tips_true: 'Please enter correct password',
    login_getVerifyCodeWord: 'Send code',
    login_getVerifyCodeWord_time: 's Send',
    login_verify_account_tips_no: 'Cannot Find the Account',
    login_verify_pw_tips_empty: 'Please enter a password.',
    login_verify_code_tips_empty: 'Please enter valid verification code',
    login_verify_pw_tips_rule: 'The password must contain 8-16 bytes, and at least contains 1 uppercase letter, 1 lowercase letter and 1 number.',
    login_validate_mailcode: 'Email Verification Code',
    login_validate_mail_tips_true: 'Please enter valid verification code！',
    login_validate_mail: 'Email',
    login_verify_zh_tips_empty: 'Account number cannot be empty',
    login_h1: 'New customer? Sign up here',
    login_h2: 'Welcome'

  },
  accountActivation: {
    accountActivation_title: 'Enable Your Account',
    accountActivation_tips_1: 'Authentication mail has sent to ',
    accountActivation_tips_2: 'Please go to email to active your account.',
    accountActivation_resend: 'Resend',
    getVerifyCodeWord_time: 's Send'
  },
  activationSuccess: {
    activationSuccess_success_title: 'Verified!',
    activationSuccess_success_tips: 'Skip to log in',
    activationSuccess_success_jumptime_unit: 's',
    activationSuccess_fail_title: 'The link is invalid',
    activationSuccess_fail_tips: 'Please log in again to get active links',
    activationSuccess_fail_tips01: 'Retrieve activation mail'
  },
  register: {
    register_registerphone: 'Phone',
    register_registermail: 'Email',
    register_phonenum: 'Mobile number',
    register_phonecheckcode: 'Verification code',
    register_phonepw: 'Password',
    register_phonepwconfirm: 'Confirm Password',
    register_agree: 'I agree to the Btcsoon',
    register_agreelaw: 'Terms of User',
    register_phonereg: 'Sign Up',
    register_phonelogin: 'Log in ',
    register_mailnum: 'Email Address',
    register_mailpw: 'Password',
    register_mailpwconfirm: 'Confirm Password',
    register_mailreg: 'Sign Up',
    register_maillogin: 'Log In ',
    register_regsuccess: 'Successful!',
    register_regsuccess_tips: 'back to log in ',
    register_verify_phone_tips_true: 'Please Enter Valid Phone Number',
    register_verify_phone_tips_empty: 'The Phone Number Shall Not Be Null！',
    register_verify_verifyCode_tips_true: 'Please Enter Valid Verification Code！',
    register_verify_phonePassword_tips_true: 'Password',
    register_verify_phonePasswordVerify_tips_true: 'Please Enter Valid Password',
    register_verify_term_tips_true: 'Please Check Terms of Users',
    register_verify_mail_tips_true: 'Please Enter Valid Email ',
    register_verify_mail_tips_empty: 'Please enter an email address.',
    register_verify_mailPassword_tips_true: 'Please enter valid password',
    register_verify_mailPasswordVerify_tips_true: 'Please enter valid password',
    register_getVerifyCodeWord: 'Send code',
    register_getVerifyCodeWord_time: 's Send',
    register_verify_account_tips_exist: 'The Account Is Existed!',
    register_verify_pw_tips_empty: 'Please enter a password.',
    register_verify_pw_tips_rule: 'The password must contain 8-16 bytes, and at least contains 1 uppercase letter, 1 lowercase letter and 1 number.',
    register_verify_placehold_tips_rule: 'Password',
    register_verify_pw_tips_reenter: 'Reenter password',
    register_verify_pw_tips_same: 'Two passwords do not match',
    register_jumptime_unit: 's',
    register_h1: 'Already have an account? Log in here',
    register_h2: 'Reference code (optional)'
  },
  forgetPassword: {
    forgetPassword_title: 'Forgot Password?',
    forgetPassword_phonenum: 'Mobile number',
    forgetPassword_phonecheckcode: 'Verification Code',
    forgetPassword_phonepw: 'Login Password',
    forgetPassword_phonepwconfirm: 'Confirm Your Password',
    forgetPassword_phonebtn: 'OK',
    forgetPassword_mailget: 'Find Password via Email',
    forgetPassword_mailbtn: 'OK',
    forgetPassword_phoneget: 'Find Password via Phone',
    forgetPassword_mailnum: 'Email Address',
    forgetPassword_verify_phone_tips_true: 'Please Enter valid Phone Number',
    forgetPassword_verify_phone_tips_empty: 'The Phone Number Shall Not Be Null！',
    forgetPassword_verify_phonePassword_tips_true: 'Please Enter Correct Password',
    forgetPassword_verify_phonePasswordVerify_tips_true: 'Please Enter Correct Password',
    forgetPassword_verify_mail_tips_true: 'Please Enter valid Email ',
    forgetPassword_verify_mail_tips_empty: 'Please enter an email address.',
    forgetPassword_getVerifyCodeWord: 'Send code',
    forgetPassword_getVerifyCodeWord_time: 's Send',
    forgetPassword_verify_verifyCode_tips_true: 'Verification Code',
    forgetPassword_verify_account_tips_no: 'The Account Is Not Existed!',
    forgetPassword_verify_pw_tips_empty: 'Please enter a password.',
    forgetPassword_verify_pw_tips_rule: 'The password must contain 8-16 bytes, and at least contains 1 uppercase letter, 1 lowercase letter and 1 number.',
    forgetPassword_verify_pw_tips_same: 'Two Passwords Do Not Match',
    forgetPassword_success_title: 'Success',
    forgetPassword_success_btn: 'Log In',
    forgetPassword_h1: 'Already have an account? Log in here',
    forgetPassword_reenter: 'Reenter password',
    forgetPassword_submit: 'Submit',
    forgetPassword_phone: 'Phone',
    forgetPassword_mail: 'Email'
  },
  mailForget: {
    mailForget_title: 'Reset ',
    mailForget_tips_1: 'Authentication Mail Has Sent To',
    mailForget_tips_2: 'Please Log In The Mail and Complete Setting',
    mailForget_resend: 'Send code',
    getVerifyCodeWord: 'Send code again',
    getVerifyCodeWord_time: 's Send'
  },
  mailResetPassword: {
    mailResetPassword_title: 'Reset',
    mailResetPassword_pw: 'Old Password',
    mailResetPassword_new_pw: 'New Password',
    mailResetPassword_pwconfirm: 'Confirm Password',
    mailResetPassword_btn: 'OK',
    mailResetPassword_verify_pw_tips_true: 'Please enter correct password',
    mailResetPassword_verify_pwconfirm_tips_true: 'Please enter correct password',
    mailResetPassword_verify_pw_tips_rule: 'The password must contain 8-16 bytes, and at least contains 1 uppercase letter, 1 lowercase letter and 1 number.',
    mailResetPassword_verify_pw_tips_same: 'Two Passwords Do Not Match',
    mailResetPassword_link_tips: 'The Link Is Incorrect'
  },
  mailForgetSuccess: {
    mailForgetSuccess_title: 'Success',
    mailForgetSuccess_tips: 'Skip to log in',
    mailForgetSuccess_jumptime_unit: 's'
  },
  fs3: {
    title: 'General Storage Service',
    access_key: 'Access Key',
    secret: 'Secret',
    URL: 'URL',
    Entry_btn: 'Entry FS3',
    copy: 'Copy',
    copied: 'Copied',
    connected_to: 'connected to',
    Connect_Wallet: 'Connect Wallet',
    Disconnect: 'Sign out'
  },
  fs3Login: {
    toptip_01: 'Currently supporting ',
    toptip_02: '; additional chains arriving soon!',
    toptip_03: 'Connect to your MetaMask Wallet',
    toptip_04: 'If you would like to visit MCS Polygon Mainnet, please click',
    toptip_04_1: 'Testnet, please click',
    toptip_04_main: 'If you would like to visit MCS',
    toptip_Network: 'or',
    title: 'Connect your wallet.',
    MetaMask: 'MetaMask',
    MetaMask_tip: 'Please connect your wallet to Mumbai Testnet or BSC TestNet.',
    Account: 'Account',
    Connected_MetaMask: 'Connected with MetaMask',
    Connected_Email: 'Connected with Email',
    Connected_Email_Address: 'Connect with Your Email Address',
    View_explorer: 'View on explorer',
    Copy_Address: 'Copy Wallet Address',
    Disconnect_mailbox: 'Disconnect from your mailbox',
    Change_Address: 'Change your Email Address',
    Copied: 'Copied',
    Connect_to_MetaMask: 'Connect to MetaMask',
    Connect_cont_Wallet: 'Connect Wallet',
    Connect_cont_tip: 'Connect your wallet address to create a multichain.storage account.',
    Connect_text: 'Decentralize your Web3 storage across blockchains',
    Connect_text_desc: 'Easily upload individual files or Bucket storage in one spot to both IPFS and Filecoin. Get your CID and then mint to OpenSea in just one click.',
    Connect_text_btn: 'Enjoy your journey with multichain.storage',
    Connected: 'Connected',
    Connect_form_label: 'Please enter your email address here',
    Connect_form_label_change: 'Enter your new email address here',
    Connect_checkbox: 'I have read and agree to the ',
    Connect_checkbox_1: 'privacy policy',
    Connect_checkbox_2: ' and to receive the newsletters.',
    Connect_form_btn: 'Submit',
    Connect_form_btn_change: 'Change',
    Connect_email_title: 'Success!',
    Connect_email_desc: 'Check your email and follow the link to activate your account reward.',
    Connect_email_again: 'Didn’t recieve any email? Try again.',
    Connect_email_jump: 'Click here to jump to the page.',
    Connect_StartFree: 'Get Started',
    Connect_Bucket: 'Back to Bucket Storage',
    Connect_BuyNow: 'Buy Now',
    Connect_TutorialVideo: 'Tutorial Video',
    back_home: 'Back to Home Page',
    Skip: 'Skip',
    Connect_tip: 'Confirm to disconnect from mailbox?'
  },
  uploadFile: {
    title: 'Files List',
    contract_id: 'CONTRACT ID',
    contract_id01: 'Contract ID',
    cid: 'CID',
    status: 'Pin Status',
    status_tooltip: "Reports the status of a file or piece of data stored on Swan's IPFS nodes.",
    status_button_tooltip: "It's an irrevocable option after clicking unpin.",
    file_name: 'File Name',
    file_name_tooltip: 'Click on the file name to view more Deal Details and DAO Signature details.',
    file_name_tip: 'Please choose a file',
    file_size: 'File Size',
    file_status: 'Status',
    upload_title: 'Task information',
    task_name: 'Task Name',
    w3ss_id: 'Storage Provider',
    w3ss_ids: 'Storage Providers',
    w3ss_id_tooltip: 'Storage providers offering storage capacity to the Filecoin network.',
    w3ss_id_nothing: 'Queuing',
    w3ss_id_nothing_tooltip: 'The data from this upload is being aggregated for storage on Filecoin. Filecoin deals will be active within 48 hours of upload.',
    create_time: 'Create Time',
    upload_time: 'Upload Time',
    deal_id: 'Deal ID',
    payment: 'Payment',
    MINT: 'Mint',
    mint_view: 'View',
    Canceled: 'Canceled',
    pay: 'PAY',
    failed: 'FAILED',
    paid: 'PAID',
    refund: 'REFUND',
    download_link: 'Download Link',
    uploadDray: 'Drag&Drop files here',
    uploadDray_or: 'or',
    uploadDray_text: 'Browse Files',
    upload: 'Upload File',
    upload_form_file_tip: 'The uploaded file cannot be empty',
    search_y: 'Search by CID',
    search_input: 'Enter value',
    deal_title: 'Filecoin deals for selected batch (pending deals have deal_id = 0):',
    Live_deal: 'LIVE DEAL',
    live_deal_btn: 'View Deal',
    live_deal_btn_no: 'No link: Error',
    Upload_More_Files: 'Upload File',
    topTip: 'Please click on the storage provider ID to view more Deal and DAO details.',
    search_title: 'Search by File Name',
    data_tooltip: 'The content identifier for a file or a piece of data.',
    payment_tip: 'Please wait until the process of locking funds completed.',
    payment_tip_deal: 'Please wait until the process completed.',
    payment_tip_deal01: 'You had uploaded successfully.',
    payment_tip_deal02: 'Please wait until the payment completed.',
    payment_tip_deal03: 'You had created collection successfully.',
    COMPLETED: 'Completed',
    SUCCESS: 'Your transaction has been submitted successfully. Check more detail in your transaction history. ',
    CLOSE: 'Close',
    Fail: 'Fail',
    FailTIP: 'Your transaction has failed. Check the transaction history for more details.',
    waiting: 'Waiting...',
    waitingTIP: 'Please check the status and payment on My Files later. Check the transaction history for more details.',
    until: 'Please wait until the task is assigned to a storage provider.',
    OK: 'OK',
    Deal_Detail: 'Deal Detail',
    detail_tip01: 'This means your deal has not been accepted by a storage provider yet.',
    no_fund_locked: 'No fund locked.',
    Successfully_unlocked_funds: 'Successfully unlocked funds.',
    Successfully_signed: 'waiting for unlock fund.',
    Successfully_signed_all: 'Successfully signed, ',
    Waiting_for_signature: 'Waiting for signature',
    detail_IPFSDownload: 'IPFS Download',
    detail_Network: 'Network',
    detail_Locked_funds: 'Locked funds',
    detail_Storage_Price: 'Storage Price',
    detail_ProposalCID: 'Proposal CID',
    detail_MessageCID: 'Message CID',
    detail_PieceCID: 'Piece CID',
    detail_Client_Address: 'Client Address',
    detail_Verified_Deal: 'Verified Deal',
    detail_Storage_Price_Per_Epoch: 'Storage Price Per Epoch',
    detail_Signature_Type: 'Signature Type',
    detail_Retrieval_Filecoin: 'Retrieval From Filecoin Network',
    detail_Retrieval_Filecoin_tooltip: "The deal will not be available for retrieval until it is Published on the blockchain. The ‘output-file’ in the end of this command is the name of the file that you'd like to save, you can also add a path to this variable.",
    detail_DAO_Signatures: 'DAO Signatures',
    detail_DAO_Signatures_tooltip: 'The signatures are used to unlock funds to provider.',
    detail_DAO_RKH_Address: 'DAO Address',
    detail_Time: 'Time',
    view_deal_logs: 'View deal logs',
    deal_logs: 'Deal Logs',
    no_logs: 'no logs',
    uploadFile_title: 'Please upload a file, an estimated storage cost will be calculated for you.',
    upload_funds: 'Insufficient funds? Get more Testnet tokens here.',
    Duration: 'Duration',
    Duration_tip: 'Please enter the duration',
    Duration_tooltip: 'Duration refers to the terms in which you want the file to be stored on the Filecoin network.',
    Storage_copy: 'Storage Copy',
    Storage_copy_tooltip: 'The number of copies of files stored in Filecoin network. Your file will be assigned to different storage providers. ',
    Estimated_Storage_Cost: 'Estimated Storage Cost',
    Estimated_Storage_Cost_tooltip: 'The estimated storage cost is calculated according to your file size, the duration you set, and the average provider price.',
    Free_Storage_Capacity: 'Free Storage Capacity',
    Free_Storage_Capacity_tooltip: "You'll be granted 10GB Free Storage Capacity monthly.",
    Select_Lock_Funds_Plan: 'Select Lock Funds Plan',
    Select_Lock_Funds_Plan_tooltip: 'The more funds locked, the sooner your file will be stored on the Filecoin network. The overpaid funds will be returned automatically after the deal is on chain.',
    latest_exchange_rate: 'The latest exchange rate of FIL to USDC is',
    Low: 'Low',
    Average: 'Average',
    High: 'High',
    File_uploading: 'File uploading',
    File_uploading_tooltip: 'Your file is still in the process of uploading to IPFS. Please keep this window open until uploading completes.',
    file_uploaded: 'This file has been uploaded.',
    file_uploaded_tip: 'We currently only support uploading a file once. Please select another file to upload.',
    file_uploaded_tip01: 'Thank you for your comprehension.',
    file_uploaded_tip02: "You don't need to pay this file, it will add to your files. Because we currently only support uploading a file once.",
    update_time: 'Update Time:',
    nft_title: 'Create Your ',
    nft_Name: 'Name',
    nft_Amount: 'Amount',
    nft_IPFSURL: 'IPFS URL',
    nft_Description: 'Description',
    nft_External: 'External Link',
    nft_Seller: 'Seller Fee Basis Points',
    nft_Fee_Recipient: 'Fee Recipient',
    nft_Image: 'Image',
    nft_list_empty: 'This list is empty',
    Payment_Transaction_Hash: 'Payment Transaction Hash',
    Back: 'Back',
    Minting: 'Minting…',
    Mint_NFT: 'Mint',
    Mint_Create_NFT: 'Create a new NFT Collection',
    View_Your_NFT: 'View Your NFT',
    View_Your_NFT_tips: 'Your NFT has been minted!',
    View_Your_NFT_tips01: 'You can view the transaction here:',
    View_Your_NFT_OpenSea: 'Click here to view your NFT on OpenSea',
    View_Your_NFT_Note: 'Note: The NFT will take some time to load on Opensea.',
    goTo: 'Go to',
    goTopage: '',
    xhr_tip: 'Upload file changed, please re-upload.',
    xhr_error500: 'Failed to create a sending deal task, please contact the administrator.',
    filter_minted: 'Minted',
    filter_no_minted: 'Not minted',
    filter_status_Unpaid: 'Unpaid',
    filter_status_Paid: 'Paid',
    filter_status_Refunding: 'Refunding',
    filter_status_Refundable: 'Refundable',
    filter_status_Refunded: 'Refunded',
    filter_status_Unlocked: 'Unlocked',
    filter_status_Success: 'Success',
    filter_status_Pending: 'Pending',
    filter_status_Processing: 'Processing',
    filter_status_Free: 'Free',
    filter_status_Failed: 'Failed',
    bucket_Uploading: 'Uploading...',
    bucket_Upload: 'Upload',
    Start_uploading: 'Start uploading',
    Upload_failed: 'Upload failed:',
    Upload_view: 'view',
    Upload_path: 'path',
    mint_pending: 'Collection information is being scanned from the chain, please wait!'
  },
  billing: {
    search_placeholder: 'Search by ',
    search_option_filename: 'File Name',
    search_option_transaction: 'Transaction',
    search_btn: 'Search',
    clear_btn: 'Clear',
    TRANSACTION: 'Transaction',
    chain_id: 'CHAIN ID',
    AMOUNT: 'Amount',
    UNLOCKAMOUNT: 'Unlock Amount',
    TOKEN: 'Token',
    FILENAME: 'File Name',
    PAYLOADCID: 'Data CID',
    deals_price: 'Price',
    WALLET: 'Wallet',
    NETWORK: 'Network',
    PAYMENTDATE: 'Payment Date',
    UNLOCKDATE: 'Unlock Date',
    Deadline: 'Deadline',
    download_module_title: 'Download Data (',
    download_module_title_kh: ')',
    download_module_title_file: 'File History',
    download_module_title_billing: 'Billing History',
    download_module_title_tooltip: "Downloading data can take a while, depending on the data's size",
    download_module_btn: 'Download',
    verify: 'Verify:',
    verify_result: 'Verification successful',
    verify_tip: 'Drag the slider to the far right',
    start_date: 'Start date',
    end_date: 'End date',
    time_to: 'To',
    data_download: 'Download',
    data_export: 'Export',
    bill_overview: 'Billing Overview',
    bill_btn_pay: 'Upgrade to paid plan',
    bill_btn_paid: 'You have become a paid subscriber',
    bill_cont1_title: 'Billing history',
    bill_cont1_desc: 'View past and current invoices',
    bill_cont2_title: 'Pricing',
    bill_cont2_desc: 'View pricing and FAQs',
    have_faq: 'Have question?',
    bill_tip: 'There may be a delay in transaction processing. Please refresh this page after confirming the success. Please do not make duplicate payments, as it may result in financial loss.'
  },
  network: {
    title: 'All Chains'
  },
  footer: {
    User_Start_Guide: 'User Start Guide',
    Developer_Quick_Start: 'Developer Quick Start',
    FAQ: 'FAQ',
    Get_Help: 'Get Help',
    copy: ' MULTICHAIN.STORAGE. ALL RIGHTS RESERVED.'
  },
  metaSpace: {
    home_title: 'Multichain.Storage',
    home_Introduction: 'Introduction',
    home_Introduction_cont: 'Multichain.Storage is a smart-contract-based cross-chain storage gateway integrated with Oracle technology and the Filecoin networks. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.',
    home_Our_Features: 'Our Features',
    home_Our_Features_cont01: 'Drag and drop files to pin to IPFS and store on the Filecoin network at the same time',
    home_Our_Features_cont02: 'Fast retrieval globally by reliable edge network',
    home_Our_Features_cont03: 'Developer support of JS, Python and Golang SDK',
    empty_prompt: 'You do not have any Bucket storage yet !',
    empty_prompt_detail: 'Start to upload your files.',
    search_bucket: 'Search Bucket Storage',
    bucket_name: 'Bucket Name',
    bucket_subname: 'Sub-Bucket Name',
    folder_name: 'Folder Name',
    add_bucket: 'Add Bucket',
    add_subbucket: 'Add Sub-Bucket',
    create_folder: 'Create Folder',
    create_bucket: 'Create bucket',
    Submit: 'Submit',
    Cancel: 'Cancel',
    Delete: 'Delete',
    Close: 'Close',
    list_bucket: 'Bucket List',
    list_bucket_tip: 'Total active Bucket storage / Total Bucket storage created',
    table_name: 'Name',
    table_type: 'Type',
    table_subdomain: 'SubDomain',
    table_space: 'Space',
    table_size: 'Size',
    table_createon: 'Create on',
    table_cid: 'CID',
    table_cid_tip: 'The content ID for the file or bucket',
    table_status: 'Status',
    table_backup_status: 'Backup Status',
    table_LastModified: 'Last Modified',
    table_action: 'Actions',
    delete_title: 'Delete Bucket',
    delete_title_detail: 'Delete permanently?',
    delete_desc: 'This Bucket item will be permanently deleted. This operation cannot be reversed.',
    Upload_btn: 'Upload',
    Upload_File: 'Upload File',
    Upload_Folder: 'Upload Folder',
    upload_desc: 'Drag&Drop Files Here',
    Browse_Folders: 'Browse Folders',
    Browse_Folders_title: 'Select where the folder will be stored:',
    Browse_Folders_title01: 'Title',
    Browse_Folders_mcs_desc: '*The folder will be stored in the MCS system and later can be edited but not shared. Files in the folder can be viewed and shared.',
    Browse_Folders_desc: "*The folder will be permanently stored in the IPFS system and can be latershared but can't be edited.",
    uploadDray_or: 'or',
    uploadDray_text: 'Browse Files',
    detail_title: 'Bucket Details',
    detail_folder_title: 'Folder Details',
    detail_BucketName: 'Bucket Name:',
    detail_folderName: 'Folder Name:',
    detail_DateCreated: 'Date Created:',
    detail_LastModified: 'Last Modified:',
    detail_CurrentFiles: 'Current Files:',
    detail_CurrentSize: 'Current Size:',
    detail_RemainingServiceDays: 'Remaining Service Days:',
    detail_BackupInfo: 'Backup Info:',
    detail_StorageProvider: 'Storage Provider ',
    detail_PieceCID: 'Piece CID:',
    ob_detail_share: 'Share',
    ob_detail_title: 'Object Details',
    ob_detail_ObjectName: 'Object Name:',
    ob_detail_DateUploaded: 'Date Uploaded:',
    ob_detail_ObjectSize: 'Object Size:',
    ob_detail_ObjectIPFSLink: 'Object IPFS Link:',
    ob_detail_ObjectCID: 'Object CID:',
    Pay: 'Pay',
    pay_title: "You don't have enough bucket credit.",
    pay_tip: '*By purchasing new Bucket storage, you will increase the "avalible Bucket storage" with current wallet address.',
    pay_tip1: '*Maxium number of Bucket storage is 33.',
    pay_body_BucketName: 'Name Your Bucket',
    pay_body_Bucket: 'Bucket',
    pay_body_Bucket_desc: 'Each bucket have 30GB storage',
    pay_body_PriceBucket: 'Price/Bucket',
    pay_body_PriceBucket_desc: 'The service for each bucket starts at the date the bucket is created.',
    pay_body_TotalPayment: 'Total Payment',
    pay_load: 'Please Wait',
    pay_success: 'Successfully Paid!',
    pay_success_desc: 'Thank you for your payment. Bucket credit added.',
    try_again: 'try again',
    Name_bucket: 'Name this bucket',
    active_bucket: 'You have to active your bucket first',
    bucket_card_create: 'Create your first bucket',
    bucket_plan: 'Take look our Pro plan',
    bucket_sdk: 'Try our SDK'
  },
  apiKey: {
    day: 'days',
    label_Expiration: 'Expiration (day)',
    label_tip01: '* Entering 0 means that the API Key is permanently active',
    label_tip02: '* The maximum number is 365',
    Add_bucket: 'Add this bucket',
    apikey_cont: 'APIKey will be permanently deleted.This operation cannot be reversed.'
  },
  comment: {
    Rate_Product: 'Rate our product',
    Tell_Comment: 'Tell us your comment',
    theme: 'Theme'
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
