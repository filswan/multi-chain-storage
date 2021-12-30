// translate router.meta.title, be used in breadcrumb sidebar tagsview
export function generateTitle(title) {
  const hasKey = this.$te('route.' + title)
  const translatedTitle = this.$t('route.' + title) // $t :this method from vue-i18n, inject in @/lang/index.js

  if (hasKey) {
    return translatedTitle
  }
  return title
}

export function generateNavbar(title) {
  const hasKey = this.$te('navbar.' + title)
  const translatedTitle = this.$t('navbar.' + title) // $t :this method from vue-i18n, inject in @/lang/index.js

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 底部
export function generateFooter(title) {
  const hasKey = this.$te('footer.' + title)
  const translatedTitle = this.$t('footer.' + title) // $t :this method from vue-i18n, inject in @/lang/index.js

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 首页
export function generateDashboard(title) {
  const hasKey = this.$te('dashboard.' + title)
  const translatedTitle = this.$t('dashboard.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 矿机
export function generateMiner(title) {
  const hasKey = this.$te('miner.' + title)
  const translatedTitle = this.$t('miner.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 公告中心
export function generateNoticeCenter(title) {
  const hasKey = this.$te('noticeCenter.' + title)
  const translatedTitle = this.$t('noticeCenter.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 更多公告
export function generateNoticeMore(title) {
  const hasKey = this.$te('noticeMore.' + title)
  const translatedTitle = this.$t('noticeMore.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 公告详情
export function generateNoticeDetail(title) {
  const hasKey = this.$te('noticeDetail.' + title)
  const translatedTitle = this.$t('noticeDetail.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 帮助中心
export function generateHelpCenter(title) {
  const hasKey = this.$te('helpCenter.' + title)
  const translatedTitle = this.$t('helpCenter.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 更多帮助
export function generateHelpMore(title) {
  const hasKey = this.$te('helpMore.' + title)
  const translatedTitle = this.$t('helpMore.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 帮助详情
export function generateHelpDetail(title) {
  const hasKey = this.$te('helpDetail.' + title)
  const translatedTitle = this.$t('helpDetail.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 注册
export function generateRegister(title) {
  const hasKey = this.$te('register.' + title)
  const translatedTitle = this.$t('register.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 注册（账号激活）
export function generateAccountActivation(title) {
  const hasKey = this.$te('accountActivation.' + title)
  const translatedTitle = this.$t('accountActivation.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 注册（激活验证）
export function generateActivationSuccess(title) {
  const hasKey = this.$te('activationSuccess.' + title)
  const translatedTitle = this.$t('activationSuccess.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 登录
export function generateLogin(title) {
  const hasKey = this.$te('login.' + title)
  const translatedTitle = this.$t('login.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 找回密码
export function generateForgetPassword(title) {
  const hasKey = this.$te('forgetPassword.' + title)
  const translatedTitle = this.$t('forgetPassword.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 找回密码（重置邮件）
export function generateMailForget(title) {
  const hasKey = this.$te('mailForget.' + title)
  const translatedTitle = this.$t('mailForget.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 找回密码（重置密码）
export function generateMailResetPassword(title) {
  const hasKey = this.$te('mailResetPassword.' + title)
  const translatedTitle = this.$t('mailResetPassword.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 找回密码（重置密码成功）
export function generateMailForgetSuccess(title) {
  const hasKey = this.$te('mailForgetSuccess.' + title)
  const translatedTitle = this.$t('mailForgetSuccess.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// APP下载
export function generateAppDownload(title) {
  const hasKey = this.$te('appDownload.' + title)
  const translatedTitle = this.$t('appDownload.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 分页组件
export function generatePages(title) {
  const hasKey = this.$te('Pages.' + title)
  const translatedTitle = this.$t('Pages.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

// 区号选择组件
export function generateZoneDialog(title) {
  const hasKey = this.$te('ZoneDialog.' + title)
  const translatedTitle = this.$t('ZoneDialog.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

export function generateTrade(title) {
  const hasKey = this.$te('trade.' + title)
  const translatedTitle = this.$t('trade.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
export function generateUser(title) {
  const hasKey = this.$te('userCenter.' + title)
  const translatedTitle = this.$t('userCenter.' + title)
  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 我的资产
export function generateAssetManagement(title) {
  const hasKey = this.$te('assetManagement.' + title)
  const translatedTitle = this.$t('assetManagement.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
export function generateMyPower(title) {
  const hasKey = this.$te('myPower.' + title)
  const translatedTitle = this.$t('myPower.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
export function generateUserChild(title) {
  const hasKey = this.$te('userCenterChild.' + title)
  const translatedTitle = this.$t('userCenterChild.' + title)
  if (hasKey) {
    return translatedTitle
  }
  return title
}

export function generateCurrententrust(title) {
  const hasKey = this.$te('currententrust.' + title)
  const translatedTitle = this.$t('currententrust.' + title)
  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 充提记录
export function generateRechargeRecord(title) {
  const hasKey = this.$te('rechargeRecord.' + title)
  const translatedTitle = this.$t('rechargeRecord.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

export function generateHistoryentrust(title) {
  const hasKey = this.$te('historyentrust.' + title)
  const translatedTitle = this.$t('historyentrust.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}

export function generateTransacation(title) {
  const hasKey = this.$te('transacation.' + title)
  const translatedTitle = this.$t('transacation.' + title)
  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 添加地址
export function generatepresentAddr(title) {
  const hasKey = this.$te('presentAddr.' + title)
  const translatedTitle = this.$t('presentAddr.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 平台分红
export function generateAllocationPlatform(title) {
  const hasKey = this.$te('allocationPlatform.' + title)
  const translatedTitle = this.$t('allocationPlatform.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 回馈累积
export function generateAllocationFeedback(title) {
  const hasKey = this.$te('allocationFeedback.' + title)
  const translatedTitle = this.$t('allocationFeedback.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 分红累积
export function generateAllocationShare(title) {
  const hasKey = this.$te('allocationShare.' + title)
  const translatedTitle = this.$t('allocationShare.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 红人榜
export function generateAllocationRanking(title) {
  const hasKey = this.$te('allocationRanking.' + title)
  const translatedTitle = this.$t('allocationRanking.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 余额宝
export function generateRemainderArea(title) {
  const hasKey = this.$te('remainderArea.' + title)
  const translatedTitle = this.$t('remainderArea.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 锁仓
export function generateLockedPosition(title) {
  const hasKey = this.$te('lockedPosition.' + title)
  const translatedTitle = this.$t('lockedPosition.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 锁仓记录
export function generateLockRecord(title) {
  const hasKey = this.$te('lockRecord.' + title)
  const translatedTitle = this.$t('lockRecord.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 法币交易
export function generateFaitTrade(title) {
  const hasKey = this.$te('faitTrade.' + title)
  const translatedTitle = this.$t('faitTrade.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 个人流水
export function generatePersonalFlow(title) {
  const hasKey = this.$te('personalFlowPage.' + title)
  const translatedTitle = this.$t('personalFlowPage.' + title)

  if (hasKey) {
    return translatedTitle
  }
  return title
}
// 统计
export function generateState(title) {
    const hasKey = this.$te('stats.' + title)
    const translatedTitle = this.$t('stats.' + title)

    if (hasKey) {
        return translatedTitle
    }
    return title
}
//
