// 处理图形数据(历史评分)
export function conversionTime(data){
    console.log(data)
    let datum = [],timeArr = []
    // 排序
    data.sort((itema,itemb) => {
        return itema.timestamp-itemb.timestamp
    })
    // 循环处理数组
    data.forEach((item,index) => {
        let time = new Date(parseInt(item.timestamp) * 1000)
        let time_end = addZero(time.getFullYear())+'-'+addZero(time.getMonth()+1)+'-'+addZero(time.getDate())
        // 如果数组里没有这组数据则添加
        if(timeArr.indexOf(time_end)===-1){
            timeArr.push(time_end)
            datum.push(item.score)
            // 如果数组里有则相加
        }else{
            datum[timeArr.indexOf(time_end)] = datum[timeArr.indexOf(time_end)]+item.score
        }
    })
    return {
        datum: datum,
        timeArr: timeArr
    }
}
// 处理图形数据(历史算力)
export function conversionTime_hashrate(data){
    let datum = [],timeArr = [],datum_assemb = [],einheit=einheit = data[0].adjusted_power.slice(-3)
    // 排序
    data.sort((itema,itemb) => {
        return itema.timestamp-itemb.timestamp
    })
    // 循环处理数组
    data.forEach((item,index) => {
        let time = new Date(parseInt(item.timestamp) * 1000)
        let time_end = addZero(time.getFullYear())+'-'+addZero(time.getMonth()+1)+'-'+addZero(time.getDate())
        let adjusted_power = item.adjusted_power.replace(einheit,'')
        let adjusted_power_delta = item.adjusted_power_delta.replace(einheit,'')
        // 如果数组里没有这组数据则添加
        if(timeArr.indexOf(time_end)===-1){
            timeArr.push(time_end)
            datum.push(adjusted_power.toString())
            datum_assemb.push(adjusted_power_delta.toString())
            // 如果数组里有则相加
        }else{
            datum[timeArr.indexOf(time_end)] = (datum[timeArr.indexOf(time_end)] + Number(adjusted_power)).toString()
        }
    })
    return {
        datum: datum,
        datum_assemb: datum_assemb,
        timeArr: timeArr,
        einheit:einheit
    }
}
function addZero(n){
    if(n<=9){
        return `0${n}`
    }
    return n
}