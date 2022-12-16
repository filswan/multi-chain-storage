export function TaskColor (type) {
  switch (type) {
    // 已创建
    case 'Created':
      return 'color: #0a318e'
      // 已分配
    case 'Assigned':
      return 'color: #ffb822'
      // 已接受
    case 'Accepted':
      return 'color: #ffb822'
      // 已完成
    case 'Completed':
      return 'color: #1dc9b7'
      // 已失败
    case 'Failed':
      return 'color: #dc3545'
      // 已取消
    case 'Cancelled':
      return 'color: #6c757d'
      // 已关闭
    case 'Closed':
      return 'color: #6c757d'
      // 已过期
    case 'Expired':
      return 'color: #6c757d'
      // 需要操作
    case 'ActionRequired':
      return 'color: #ffb822'
      // 交易已发送
    case 'DealSent':
      return 'color: #67c23a'
      // 文件已导入
    case 'FileImporting':
      return 'color: #ffb822'
      // 导入失败
    case 'FileImported':
      return 'color: #ffb822'
      // 下载中
    case 'ImportFailed':
      return 'color: #dc3545'
      // 已下载
    case 'Downloading':
      return 'color: #ffb822'
      // 下载失败
    case 'DownloadFailed':
      return 'color: #dc3545'
      // 有效交易
    case 'DealActive':
      return 'color: #28a745'
      // 等待中
    case 'Waiting':
      return 'color: #ffb822'
      // 准备导入
    case 'ReadyForImport':
      return 'color: #ffb822'
  }
}
export function DealColor (type) {
  switch (type) {
    // 已创建
    case 'Created':
      return 'color: #0a318e'
      // 有效交易
    case 'DealActive':
      return 'color: #28a745'
      // 等待中
    case 'Waiting':
      return 'color: #ffb822'
      // 准备导入
    case 'ReadyForImport':
      return 'color: #ffb822'
      // 文件已导入
    case 'FileImporting':
      return 'color: #ffb822'
      // 导入失败
    case 'FileImported':
      return 'color: #ffb822'
      // 下载中
    case 'ImportFailed':
      return 'color: #dc3545'
      // 已下载
    case 'Downloading':
      return 'color: #ffb822'
      // 下载失败
    case 'DownloadFailed':
      return 'color: #dc3545'
      // 已完成
    case 'Completed':
      return 'color: #1dc9b7'
      // 已失败
    case 'Failed':
      return 'color: #6c757d'
  }
}
