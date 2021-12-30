export function Task_color(type){
    switch(type){
        // 已创建
        case 'Created':
            return 'color: #0a318e'
        break;
        // 已分配
        case 'Assigned':
            return 'color: #ffb822'
        break;
        // 已接受
        case 'Accepted':
            return 'color: #ffb822'
        break;
        // 已完成
        case 'Completed':
            return 'color: #1dc9b7'
        break;
        // 已失败
        case 'Failed':
            return 'color: #dc3545'
        break;
        // 已取消
        case 'Cancelled':
            return 'color: #6c757d'
        break;
        // 已关闭
        case 'Closed':
            return 'color: #6c757d'
        break;
        // 已过期
        case 'Expired':
            return 'color: #6c757d'
        break;
        // 需要操作
        case 'ActionRequired':
            return 'color: #ffb822'
        break;
        // 交易已发送
        case 'DealSent':
            return 'color: #67c23a'
        break;
        // 文件已导入
        case 'FileImporting':
            return 'color: #ffb822'
        break;
        // 导入失败
        case 'FileImported':
            return 'color: #ffb822'
        break;
        // 下载中
        case 'ImportFailed':
            return 'color: #dc3545'
        break;
        // 已下载
        case 'Downloading':
            return 'color: #ffb822'
        break;
        // 下载失败
        case 'DownloadFailed':
            return 'color: #dc3545'
        break;
        // 有效交易
        case 'DealActive':
            return 'color: #28a745'
        break;
        // 等待中
        case 'Waiting':
            return 'color: #ffb822'
        break;
        // 准备导入
        case 'ReadyForImport':
            return 'color: #ffb822'
        break;
    }
}
export function Deal_color(type){
    switch(type){
        // 已创建
        case 'Created':
            return 'color: #0a318e'
        break;
        // 有效交易
        case 'DealActive':
            return 'color: #28a745'
        break;
        // 等待中
        case 'Waiting':
            return 'color: #ffb822'
        break;
        // 准备导入
        case 'ReadyForImport':
            return 'color: #ffb822'
        break;
        // 文件导入中
        case 'Failed':
            return 'color: #dc3545'
        break;
        // 文件已导入
        case 'FileImporting':
            return 'color: #ffb822'
        break;
        // 导入失败
        case 'FileImported':
            return 'color: #ffb822'
        break;
        // 下载中
        case 'ImportFailed':
            return 'color: #dc3545'
        break;
        // 已下载
        case 'Downloading':
            return 'color: #ffb822'
        break;
        // 下载失败
        case 'DownloadFailed':
            return 'color: #dc3545'
        break;
        // 已完成
        case 'Completed':
            return 'color: #1dc9b7'
        break;
        // 已失败
        case 'Failed':
            return 'color: #6c757d'
        break;
    }
}