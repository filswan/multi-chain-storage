<template>
  <div class="spaceStyle" v-loading="listLoad">
    <div class="slideScroll">
      <div v-loading="wrongLoad" v-if="mcsEmail" class="address_email">
        <h4>{{$t('fs3Login.Connected_Email')}}</h4>
        <div class="address_body">
          <h6 class="address">{{mcsEmail | hiddEmail}}</h6>
          <div class="address_right" :class="{'bg-primary':mcsEmailStatus === 1}">
            <div>{{mcsEmailStatus === 1?'Active':'Not Active'}}</div>
          </div>
        </div>
        <div class="form_top">
          <div class="search_file">
            <div class="createTask">
              <el-popconfirm @confirm="wrongInfo('disconnect')" :confirm-button-text="$t('uploadFile.OK')" :cancel-button-text="$t('metaSpace.Cancel')" icon="el-icon-info" icon-color="red" title="Confirm to disconnect from mailbox?">
                <a slot="reference">
                  <svg t="1669800414838" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5038" width="64" height="64">
                    <path d="M832.6 191.4c-84.6-84.6-221.5-84.6-306 0l-96.9 96.9 51 51 96.9-96.9c53.8-53.8 144.6-59.5 204 0 59.5 59.5 53.8 150.2 0 204l-96.9 96.9 51.1 51.1 96.9-96.9c84.4-84.6 84.4-221.5-0.1-306.1zM446.5 781.6c-53.8 53.8-144.6 59.5-204 0-59.5-59.5-53.8-150.2 0-204l96.9-96.9-51.1-51.1-96.9 96.9c-84.6 84.6-84.6 221.5 0 306s221.5 84.6 306 0l96.9-96.9-51-51-96.8 97zM260.3 209.4c-3.1-3.1-8.2-3.1-11.3 0L209.4 249c-3.1 3.1-3.1 8.2 0 11.3l554.4 554.4c3.1 3.1 8.2 3.1 11.3 0l39.6-39.6c3.1-3.1 3.1-8.2 0-11.3L260.3 209.4z"
                      p-id="5039" fill="#fff"></path>
                  </svg>
                  <span>{{$t('fs3Login.Disconnect_mailbox')}}</span>
                </a>
              </el-popconfirm>
              <a @click="dialogFun('emailLogin', '', 'change')">
                <svg t="1640937862402" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3723" width="32" height="32">
                  <path d="M852.77 889.05H171.23A36.27 36.27 0 0 1 135 852.78V171.22A36.27 36.27 0 0 1 171.23 135H375a36.28 36.28 0 0 1 0 72.55H207.5v609h609V649a36.28 36.28 0 1 1 72.55 0v203.78a36.27 36.27 0 0 1-36.28 36.27z" fill="#fff" p-id="3724"></path>
                  <path d="M407.15 653.13a36.28 36.28 0 0 1-25.66-61.93L747.66 225A36.28 36.28 0 1 1 799 276.34L432.8 642.51a36.17 36.17 0 0 1-25.65 10.62z" fill="#fff" p-id="3725"></path>
                  <path d="M852.78 414.94V171.23H609.06l140.42 103.29 103.3 140.42z" fill="#fff" p-id="3726"></path>
                  <path d="M852.78 451.22a36.29 36.29 0 0 1-29.23-14.78l-100-136-136-100a36.28 36.28 0 0 1 21.5-65.5h243.72a36.28 36.28 0 0 1 36.28 36.28V415a36.28 36.28 0 0 1-36.27 36.27zM719.6 207.5l51.4 37.8a36.23 36.23 0 0 1 7.7 7.7l37.8 51.38V207.5z" fill="#fff"
                    p-id="3727"></path>
                </svg>
                <span>{{$t('fs3Login.Change_Address')}}</span>
              </a>
            </div>
          </div>
        </div>
      </div>
      <div v-else class="address_email">
        <h4>{{$t('fs3Login.Connected_Email')}}</h4>
        <div class="form_top">
          <div class="search_file">
            <div class="createTask">
              <a @click="dialogFun('emailLogin')">
                <svg t="1669803088505" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2608" width="64" height="64">
                  <path d="M1024.00144 220.64v574.32c0 5.76-0.736 11.456-1.824 17.072v0.064l-0.224 0.944A84.336 84.336 0 0 1 939.29744 880H101.08944a102.72 102.72 0 0 1-82.512-40.832 95.904 95.904 0 0 1-18.56-57.248V236.256a74.08 74.08 0 0 1 52.064-70.784l0.736-0.208 0.208-0.08A127.648 127.648 0 0 1 89.66544 160h883.424c1.328 0.512 2.64 1.232 4.032 1.44 25.52 4.32 40.592 19.152 45.344 44.496v0.08c0.88 4.816 1.536 9.712 1.536 14.608zM512.00144 591.232l451.152-379.472c-66.272-3.888-888.48-1.728-897.84 2.368L512.00144 591.2v0.064z m133.856-45.152l-114.112 96.144c-13.76 11.52-25.6 11.664-39.2 0.208l-48.272-40.752-63.84-53.936L81.02544 825.28l1.024 2.304h873.328L645.79344 546.08h0.064z m-304.864-31.68L51.79344 270.032v512.48l289.28-268.096h-0.08z m344.352-1.648l287.088 260.96V270.944L685.34544 512.768z"
                    p-id="2609" fill="#fff"></path>
                </svg>
                <span>{{$t('fs3Login.Connected_Email_Address')}}</span>
              </a>
            </div>
          </div>
        </div>
      </div>

      <div>
        <h4>{{$t('my_profile.apiKey_your_title')}}</h4>
        <h6>{{$t('my_profile.apiKey_tips_01')}}</h6>
        <div class="form_top">
          <div class="search_file">
            <div class="createTask">
              <a @click="dialogFun('add_apikey')">
                <img src="@/assets/images/space/icon_01.png" alt="">
                <span>{{$t('my_profile.apiKey_btn_01')}}</span>
              </a>
            </div>
          </div>
        </div>
        <div class="fes-search">
          <el-table :data="toolData" stripe style="width: 100%" max-height="300" :empty-text="$t('deal.formNotData')" class="table_cell">
            <el-table-column prop="api_key" :label="$t('my_profile.table_apiKey_th_02')"></el-table-column>
            <el-table-column prop="token" :label="$t('my_profile.table_apiKey_th_03')" max-width="150">
              <template>*******</template>
            </el-table-column>
            <el-table-column prop="valid_days" :label="'Expiration (day)'">
              <template slot-scope="scope">
                <div style="">
                  {{calculateDiffTime(scope.row.valid_days, scope.row.create_at)}}
                  <!-- ({{momentFun(scope.row.create_at,scope.row.valid_days)}}) -->
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="create_at" :label="$t('my_profile.table_apiKey_th_04')">
              <template slot-scope="scope">
                <div style="">
                  {{momentFun(scope.row.create_at)}}
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="qr_code" label="" width="120">
              <template slot-scope="scope">
                <div class="revoke">
                  <el-button type="danger" :disabled="scope.row.status == 'Deleted'?true:false" @click="dialogFun('delete', scope.row)">{{$t('my_profile.apiKey_btn_02')}}</el-button>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>

      <div>
        <h4>{{$t('my_profile.apiKey_your_Domain')}}</h4>
        <h6></h6>
        <div class="form_top">
          <div class="search_file">
            <div class="createTask">
              <a @click="dialogFun('add_domain')">
                <img src="@/assets/images/space/icon_01.png" alt="">
                <span>{{$t('my_profile.apiKey_btn_03')}}</span>
              </a>
            </div>
          </div>
        </div>
        <div class="fes-search">
          <el-table :data="[]" stripe style="width: 100%" max-height="300" :empty-text="$t('deal.formNotData')" class="table_cell">
            <el-table-column prop="api_key" :label="'Domain address'"></el-table-column>
            <el-table-column prop="create_at" :label="$t('my_profile.table_apiKey_th_04')">
              <template slot-scope="scope">
                <div style="">
                  {{momentFun(scope.row.create_at)}}
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="qr_code" label="Actions" max-width="160">
              <template slot-scope="scope">
                <div class="revoke">
                  <el-button type="danger" @click="dialogFun('delete', scope.row)">{{$t('my_profile.apiKey_btn_04')}}</el-button>
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </div>

    <el-dialog :title="$t('my_profile.create_api_title01')" :visible.sync="apiTips" :width="width" custom-class="wrongNet">
      <div class="apiTipCont">
        <label>{{$t('my_profile.create_api_tips03')}}</label>
        <p>{{apiCont.apiKey}}</p>

        <label>{{$t('my_profile.create_api_tips04')}}</label>
        <p>
          {{apiCont.access}}
          <span class="el-icon-document-copy" @click="copyLink(apiCont.access, 1)"></span>
        </p>
      </div>
    </el-dialog>

    <pop-ups v-if="dialogFormVisible" :dialogFormVisible="dialogFormVisible" :typeModule="typeName" :createLoad="createLoad" :listTableLoad="listTableLoad" :changeTitle="changeTitle" @getPopUps="getPopUps"></pop-ups>
  </div>
</template>
<script>
import moment from 'moment'
import popUps from '@/components/popups'
import QS from 'qs'
let that
export default {
  data () {
    return {
      width: document.body.clientWidth > 600 ? '450px' : '95%',
      listLoad: true,
      listTableLoad: false,
      createLoad: false,
      toolData: [],
      areaBody: {},
      dialogFormVisible: false,
      typeName: 'add_apikey',
      apiTips: false,
      apiCont: {
        apiKey: '',
        access: ''
      },
      wrongLoad: false,
      changeTitle: ''
    }
  },
  components: { popUps },
  computed: {
    mcsEmail () {
      const data = this.$store.getters.mcsEmail
      return data === '{}' ? '-' : JSON.parse(data).email
    },
    mcsEmailStatus () {
      const data = this.$store.getters.mcsEmail
      return data === '{}' ? 0 : JSON.parse(data).email_status
    }
  },
  methods: {
    calculateDiffTime (validDays, startTime) {
      if (String(validDays) === '0') return 'Forever'
      if (!parseInt(validDays)) return '-'
      var endTime = Math.round(new Date() / 1000)
      var timeDiff = (startTime + 24 * 60 * 60 * validDays) - endTime
      if (timeDiff <= 0) return 'Expired'
      var day = parseInt(timeDiff / 86400)
      var hour = parseInt((timeDiff % 86400) / 3600)
      var minute = parseInt((timeDiff % 3600) / 60)
      var m = parseInt((timeDiff % 60))
      if (day > 0) return day + ' days'
      else if (hour > 0) return hour + ' hours'
      else if (minute > 0) return minute + ' minutes'
      else if (m > 0) return m + ' seconds'
      else return '-'
      // day = day ? (day + 'Days') : ''
      // hour = hour ? (hour + 'hours') : ''
      // minute = minute ? (minute + 'minutes') : ''
      // m = m ? (m + 'second') : ''
      // return day + hour + minute + m || '-'
    },
    async wrongInfo (status) {
      that.wrongLoad = true
      if (status === 'disconnect') await that.$metaLogin.Disconnect()
      await that.$metaLogin.emailSign('', 'detail')
      that.wrongLoad = false
    },
    async getPopUps (dialog, rows, day) {
      // console.log('rows', rows)
      switch (rows) {
        case 'delete':
          that.deleteApiKey()
          break
        case 'add_domain':
          that.createDomain(dialog, rows, day)
          break
        default:
          if (rows) that.createKey(dialog, rows, day)
          else that.dialogFormVisible = dialog
          break
      }
    },
    async createDomain (dialog, rows, day) {
      that.dialogFormVisible = dialog
    },
    async createKey (dialog, formName, day) {
      that.createLoad = true
      const params = {
        'valid_days': day
      }
      const apiKeyRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/user/generate_api_key?${QS.stringify(params)}`, 'get')
      if (!apiKeyRes || apiKeyRes.status !== 'success') {
        that.$message.error(apiKeyRes.message ? apiKeyRes.message : 'Fail')
        that.createLoad = false
        return false
      }
      await that.$commonFun.timeout(500)
      that.createLoad = false
      that.dialogFormVisible = dialog
      that.apiCont = {
        apiKey: apiKeyRes.data.apikey || '',
        access: apiKeyRes.data.access_token || ''
      }
      that.apiTips = true
      that.getListBuckets()
    },
    async dialogFun (name, row, title) {
      that.typeName = name
      that.areaBody = row || {}
      that.changeTitle = title || ''
      that.dialogFormVisible = true
      if (name === 'add_domain' || name === 'add_apikey') document.getElementById('content_client').scrollTop = 120
    },
    async deleteApiKey () {
      that.listTableLoad = true
      const params = {
        'apikey': that.areaBody.api_key
      }
      const deleteRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/user/delete_api_key?${QS.stringify(params)}`, 'put')
      if (!deleteRes || deleteRes.status !== 'success') {
        that.$message.error(deleteRes.status || 'Fail')
      }
      await that.$commonFun.timeout(500)
      that.listTableLoad = false
      that.dialogFormVisible = false
      that.getListBuckets()
    },
    async getListBuckets (name) {
      that.listLoad = true
      const directoryRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/user/apikey`, 'get')
      if (!directoryRes || directoryRes.status !== 'success') {
        that.$message.error(directoryRes.message || 'Fail')
        return false
      }
      that.toolData = directoryRes.data.apikey || []
      that.listLoad = false
    },
    momentFun (dateItem, type) {
      let dateNew = type ? new Date(dateItem * 1000 + 24 * 60 * 60 * type * 1000).getTime() : new Date(dateItem * 1000).getTime()
      let dataUnit = ''
      let dataTime = new Date(dateNew) + ''
      let dataUnitIndex = dataTime.indexOf('GMT')
      let dataUnitArray = dataTime.substring(dataUnitIndex, dataUnitIndex + 8)
      switch (dataUnitArray) {
        case 'GMT+1000':
          dataUnit = 'UTC+10'
          break
        case 'GMT-1000':
          dataUnit = 'UTC-10'
          break
        case 'GMT+0000':
          dataUnit = 'UTC+0'
          break
        default:
          dataUnit = dataUnitArray ? dataUnitArray.replace(/0/g, '').replace('GMT', 'UTC') : '-'
          break
      }
      dateNew = dateNew
        ? moment(new Date(parseInt(dateNew))).format('YYYY-MM-DD HH:mm:ss') + ` (${dataUnit})`
        : '-'
      return dateNew
    },
    copyLink (text) {
      var txtArea = document.createElement('textarea')
      txtArea.id = 'txt'
      txtArea.style.position = 'fixed'
      txtArea.style.top = '0'
      txtArea.style.left = '0'
      txtArea.style.opacity = '0'
      txtArea.value = text
      document.body.appendChild(txtArea)
      txtArea.select()

      try {
        var successful = document.execCommand('copy')
        var msg = successful ? 'Copied to clipboard!' : 'copy failed!'
        // console.log('Copying text command was ' + msg)
        if (successful) {
          that.$message({
            message: msg,
            type: 'success'
          })
        }
      } catch (err) {
        console.log('Oops, unable to copy')
      } finally {
        document.body.removeChild(txtArea)
      }
    }
  },
  mounted () {
    that = this
    document.getElementById('content-box').scrollTop = 0
    that.$store.dispatch('setRouterMenu', 21)
    that.$store.dispatch('setHeadertitle', that.$t('route.ApiKey'))
    that.wrongInfo()
    that.getListBuckets()
  },
  filters: {
    NumFormat (value) {
      if (!value) return '-'
      return value
    },
    hiddEmail (val) {
      if (!val) return '-'
      if (val.indexOf('@') !== -1) {
        var a = val.indexOf('@')
        val = '****' + val.substring(a - 2)
        return val
      } else {
        return `${val.substring(4, 4)}****${val.substring(val.length - 4)}`
      }
    },
    formatbytes: function (bytes) {
      if (bytes === 0) return '0 B'
      if (!bytes) return '-'
      var k = 1024 // or 1000
      var sizes = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
      var i = Math.floor(Math.log(bytes) / Math.log(k))

      if (Math.round((bytes / Math.pow(k, i))).toString().length > 3) {
        // 判断大小是999999999左右，解决会显示成1.00e+3科学计数法
        i += 1
      }
      return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
      // return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i]
    }
  }
}
</script>
<style lang="scss" scoped>
.el-dialog__wrapper /deep/ {
  display: flex;
  align-items: center;
  .wrongNet {
    margin: auto !important;
    box-shadow: 0 0 13px rgba(128, 128, 128, 0.8);
    border-radius: 0.2rem;
    .el-dialog__header {
      padding: 0.2rem 0.4rem;
      display: flex;
      align-items: center;
      justify-content: space-between;
      border-bottom: 1px solid #dfdfdf;
      color: #000;
      font-size: 0.22rem;
      font-weight: 500;
      line-height: 1;
      text-transform: capitalize;
      @media screen and (max-width: 479px) {
        padding: 0.3rem 0.2rem;
      }
      .el-dialog__headerbtn {
        position: relative;
        top: auto;
        right: auto;
        font-size: inherit;
        i {
          font-size: inherit;
          &:hover {
            color: #0b318f;
          }
        }
      }
      .el-dialog__title {
        font-size: inherit;
      }
    }
    .el-dialog__body {
      position: relative;
      padding: 0.3rem 0.4rem 0.4rem;
      font-size: 0.2rem;
      @media screen and (max-width: 479px) {
        padding: 0.2rem;
      }
      label {
        word-break: break-word;
        line-height: 1;
        color: #666;
        font-size: inherit;
      }
      .address {
        background: rgba(233, 233, 233, 1);
        padding: 8px;
        margin: 10px 0 22px;
        border-radius: 8px;
        font-size: inherit;
      }
      .address_email {
        margin: 0 0 10px;
        .address_body {
          display: flex;
          align-items: center;
          justify-content: space-between;
          margin: 10px 0 0;
          .address {
            width: 80%;
            margin: 0;
          }
          .address_right {
            position: relative;
            display: inline-block;
            padding: 0.05rem 0.2rem 0.05rem 0.32rem;
            margin: 0 5px;
            background-color: rgba(85, 128, 233, 0.15);
            font-size: 0.148rem;
            border-radius: 0.5rem;
            white-space: nowrap;
            @media screen and (max-width: 1600px) {
              font-size: 13px;
            }
            @media screen and (max-width: 600px) {
              font-size: 12px;
            }
            &::before {
              position: absolute;
              left: 0.16rem;
              top: 50%;
              content: "";
              width: 0.08rem;
              height: 0.08rem;
              margin-top: -0.04rem;
              background-color: #606266;
              border-radius: 0.5rem;
            }
          }
          .bg-primary {
            &::before {
              background-color: #4d73ff;
            }
          }
        }
        .share {
          .el-button {
            width: 100%;
            margin: 3px 0 0;
            font-size: 13px;
            @media screen and (min-width: 1800px) {
              font-size: 14px;
            }
            @media screen and (max-width: 600px) {
              font-size: 12px;
            }
          }
        }
        .el-loading-mask {
          .el-loading-spinner {
            top: 50%;
          }
        }
      }
      .share {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        justify-content: flex-start;
        font-size: inherit;
        .el-button {
          min-width: 50%;
          padding: 0;
          margin: 8px 0 0;
          background: transparent !important;
          border: 0;
          color: #4f7bf5;
          font-size: inherit;
          font-weight: normal;
          font-family: inherit;
          opacity: 0.8;
          span {
            display: flex;
            align-items: center;
            svg {
              width: 15px;
              height: 15px;
              margin: 0 3px 0 0;
            }
            .icon_big {
              width: 13px;
              height: 13px;
            }
          }
          &:hover {
            background: transparent;
            opacity: 1;
          }
        }
      }
      .loadStyle {
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        z-index: 2000;
        background: rgba(255, 255, 255, 1);
        border-radius: 0.2rem;
      }
      .apiTipCont {
        p {
          display: flex;
          align-items: center;
          text-indent: 0.1rem;
          margin: 0.1rem;
          color: #7e7e7e;
          font-size: 0.18rem;
          .el-icon-document-copy {
            display: block;
            font-size: 17px;
            cursor: pointer;
          }
        }
      }
    }
  }
}
.spaceStyle /deep/ {
  position: relative;
  width: calc(100% - 0.6rem);
  padding: 0 0 0.4rem;
  margin: 0.3rem;
  background-color: #fff;
  border-radius: 0.1rem;
  h4 {
    font-size: 0.22rem;
    font-weight: normal;
    color: #000;
    line-height: 1.5;
  }
  h6 {
    margin: 0.1rem 0 0.35rem;
    font-size: 16px;
    font-weight: normal;
    color: #666;
    line-height: 1.3;
    @media screen and (max-width: 1600px) {
      font-size: 14px;
    }
    a {
      color: #04a2a2;
      text-decoration: underline;
    }
  }
  .el-loading-mask {
    z-index: 5;
  }
  .slideScroll {
    height: calc(100% - 0.6rem);
    min-height: 350px;
    padding: 0.3rem;
    .form_top {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      margin: 0 auto 0.3rem;
      .title {
        width: 100%;
        margin: 0;
        text-align: left;
        font-size: 0.12rem;
        font-weight: 600;
        color: #000;
        line-height: 0.42rem;
        text-indent: 0;
      }

      .upload_title {
        width: 100%;
        margin: 0 0 0.05rem;
        text-align: left;
        font-weight: 600;
        line-height: 1.5;
        text-indent: 0;
        font-size: 0.13rem;
        color: #222;
      }
      .search {
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 100%;
        height: 0.42rem;
        .search_right {
          display: flex;
          align-items: center;
          // margin-left: 0.3rem;
        }

        span {
          margin: auto 0.05rem auto 0.35rem;
          font-size: 0.1372rem;
          color: #000;
          white-space: nowrap;
        }

        .el-button {
          height: 0.34rem;
          padding: 0 0.4rem;
          margin: 0 0.1rem;
          color: #fff;
          line-height: 0.34rem;
          font-size: 0.1372rem;
          border: 0;
          border-radius: 0.08rem;
        }

        .el-input {
          float: left;
          width: auto;

          .el-input__inner {
            width: 100%;
            max-width: 3rem;
            border-radius: 0.08rem;
            border: 1px solid #f8f8f8;
            color: #737373;
            font-size: 0.12rem;
            height: 0.24rem;
            line-height: 0.24rem;
            padding: 0 0.27rem;
          }

          .el-input__icon {
            line-height: 0.24rem;
          }
        }
      }
      .search_file {
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 100%;
        margin: 0;
        p {
          font-size: 0.13rem;
          color: #222;
        }
        .createTask {
          border-radius: 0.1rem;
          cursor: pointer;
          a {
            display: flex;
            align-items: center;
            justify-content: center;
            min-width: 1.6rem;
            padding: 0.13rem;
            margin: 0;
            background: linear-gradient(45deg, #4e88ff, #4b5fff);
            border-radius: 0.14rem;
            line-height: 1.5;
            text-align: center;
            color: #fff;
            font-size: 0.19rem;
            border: 0;
            outline: none;
            transition: background-color 0.3s, border-color 0.3s, color 0.3s,
              box-shadow 0.3s;
            cursor: pointer;
            white-space: nowrap;
            img,
            svg {
              display: inline-block;
              height: 0.25rem;
              margin: 0 0.1rem 0 0;
              @media screen and (max-width: 1260px) {
                height: 16px;
              }
            }
            svg {
              width: 0.25rem;
              @media screen and (max-width: 1260px) {
                width: 16px;
              }
            }
            &:hover {
              opacity: 0.9;
              box-shadow: 0 12px 12px -12px rgba(12, 22, 44, 0.32);
            }
          }
        }
        .search_right {
          display: flex;
          align-items: center;
          width: 100%;
          margin-right: 0.9rem;
          .el-button {
            height: 0.3rem;
            padding: 0 0.15rem;
            margin: 0;
            color: #fff;
            line-height: 0.3rem;
            font-size: 0.1372rem;
            border: 0;
            border-top-left-radius: 0;
            border-bottom-left-radius: 0;
          }
        }
        .el-input {
          float: left;
          padding: 0 0.5rem 0 0.7rem;
          background: #f7f7f7;
          border-radius: 0.2rem;
          .el-input__inner {
            width: 100%;
            color: #555;
            font-size: 0.19rem;
            font-weight: 500;
            height: 0.54rem;
            line-height: 0.3rem;
            padding: 0;
            background: transparent;
            border: 0;
            border-radius: 0.2rem;
          }
          .el-input__prefix {
            left: 0.3rem;
            i {
              display: flex;
              align-items: center;
              font-size: 0.25rem;
              color: #000;
            }
          }
          .el-input__icon {
            line-height: 0.3rem;
          }
          .el-input__suffix {
            right: 0.2rem;
            .el-input__clear {
              display: flex;
              align-items: center;
              font-size: 0.25rem;
              color: #666;
            }
          }
          ::-webkit-input-placeholder {
            color: #555;
          } /* 使用webkit内核的浏览器 */
          :-moz-placeholder {
            color: #555;
          } /* Firefox版本4-18 */
          ::-moz-placeholder {
            color: #555;
          } /* Firefox版本19+ */
          :-ms-input-placeholder {
            color: #555;
          } /* IE浏览器 */
        }
        .is-disabled {
          opacity: 0.2;
        }
      }
    }
    .address_email {
      margin: 0 0 10px;
      .address_body {
        display: flex;
        align-items: center;
        justify-content: flex-start;
        margin: 10px 0 0;
        .address {
          margin: 0;
        }
        .address_right {
          position: relative;
          display: inline-block;
          padding: 0.02rem 0.2rem 0.02rem 0.32rem;
          margin: 0 5px;
          background-color: rgba(85, 128, 233, 0.15);
          font-size: 0.148rem;
          border-radius: 0.5rem;
          white-space: nowrap;
          @media screen and (max-width: 1600px) {
            font-size: 13px;
          }
          @media screen and (max-width: 600px) {
            font-size: 12px;
          }
          &::before {
            position: absolute;
            left: 0.16rem;
            top: 50%;
            content: "";
            width: 0.08rem;
            height: 0.08rem;
            margin-top: -0.04rem;
            background-color: #606266;
            border-radius: 0.5rem;
          }
        }
        .bg-primary {
          &::before {
            background-color: #4d73ff;
          }
        }
      }
      .share {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        justify-content: flex-start;
        font-size: inherit;
        .el-button {
          display: block;
          width: 100%;
          padding: 0;
          margin: 8px 0 0;
          background: transparent !important;
          border: 0;
          color: #4f7bf5;
          font-weight: normal;
          font-family: inherit;
          opacity: 0.8;
          font-size: 13px;
          @media screen and (min-width: 1800px) {
            font-size: 14px;
          }
          @media screen and (max-width: 600px) {
            font-size: 12px;
          }
          span {
            display: flex;
            align-items: center;
            svg {
              width: 15px;
              height: 15px;
              margin: 0 3px 0 0;
            }
            .icon_big {
              width: 13px;
              height: 13px;
            }
          }
          &:hover {
            background: transparent;
            opacity: 1;
          }
        }
      }
      .form_top {
        .search_file {
          .createTask {
            display: flex;
            align-items: center;
            flex-wrap: wrap;
            a {
              margin: 0.35rem 0.1rem 0.35rem 0;
              @media screen and (max-width: 600px) {
                margin: 0.15rem 0.1rem 0 0;
              }
            }
          }
        }
      }
      .el-loading-mask {
        .el-loading-spinner {
          top: 0%;
        }
      }
    }
    .fe-none {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100%;
      padding: 0 5%;
      .p_label {
        padding: 0.15rem 0.3rem;
        font-size: 0.27rem;
        color: #4f87ff;
        line-height: 1.2;
        border: dashed;
        border-radius: 0.1rem;
        @media screen and (max-width: 1600px) {
          font-size: 0.25rem;
        }
      }
    }
    .fes-header {
      display: flex;
      justify-content: space-between;
      align-content: center;
      width: calc(100% - 0.4rem);
      padding: 0.25rem 0.2rem;
      img {
        width: auto;
        max-width: 100%;
        height: 0.35rem;
      }
      svg {
        font-size: 0.25rem;
        width: 0.25rem;
        height: 0.25rem;
        margin: 0.05rem 0 0;
        color: #333;
        cursor: pointer;
      }
      h2 {
        margin: 10px 0 0 13px;
        font-weight: 400;
        color: #333;
        font-size: 0.2rem;
      }
    }
    .fs3_backup {
      margin: 0;
      .introduce {
        margin: 0.1rem 0 0.05rem;
        text-indent: 0.2rem;
        background: #002a39;
        // font-family: 'm-semibold';
        font-weight: bold;
        a {
          display: block;
          line-height: 3;
          font-size: 0.14rem;
          color: #333;
          &:hover {
            color: #2f85e5;
          }
          @media screen and (max-width: 999px) {
            font-size: 13px;
            line-height: 3.5;
          }
        }
        .active {
          color: #2f85e5;
        }
      }
      .introRouter {
        font-size: 0.14rem;
        @media screen and (max-width: 999px) {
          font-size: 13px;
        }
        a {
          display: block;
          padding: 0.07rem 0.2rem;
          color: #333;
          font-size: inherit;
          @media screen and (max-width: 999px) {
            padding: 8px 0.2rem;
          }
          &:hover {
            color: #7ecef4;
            background-color: rgba(0, 0, 0, 0.1);
          }
        }
        .active {
          color: #7ecef4;
        }
      }
      .el-tree {
        padding: 0 0.25rem;
        background: transparent;
        color: #333;
        .el-tree-node {
          .el-tree-node__content {
            height: auto;
            background: transparent !important;
            margin: 0 0 0.08rem;
            .el-tree-node__expand-icon {
              padding: 0 0.05rem;
              &:before {
                font-size: 0.2rem;
              }
            }
            .el-tree-node__label {
              font-size: 0.15rem;
              @media screen and (max-width: 999px) {
                font-size: 14px;
              }
            }
            &:hover {
              color: #5f9dcc;
            }
          }
          .el-tree-node__children {
            .el-tree-node__content {
              .el-tree-node__label {
                font-size: 0.14rem;
                line-height: 1.5;
                @media screen and (max-width: 999px) {
                  font-size: 13px;
                }
              }
            }
          }
          .is-current,
          .is-checked {
            color: #5f9dcc;
          }
        }
      }
    }
    .fes-search {
      // height: calc(100% - 1.7rem);
      margin: 0 auto 0.65rem;
      .title {
        display: flex;
        align-items: center;
        justify-content: flex-start;
        margin: 0 0.3rem 0.25rem;
        font-size: 16px;
        color: #000;
        line-height: 1;
        @media screen and (max-width: 1600px) {
          font-size: 14px;
        }
        img {
          display: block;
          width: 22px;
          height: 22px;
          margin: 0 0 0 5px;
          cursor: pointer;
          @media screen and (max-width: 1680px) {
            width: 20px;
            height: 20px;
          }
          @media screen and (max-width: 1260px) {
            width: 17px;
            height: 17px;
          }
        }
      }
      .el-table {
        overflow: visible;
        font-size: 0.18rem;
        .el-loading-mask {
          .el-loading-spinner {
            top: 50%;
          }
        }
        .el-table__body-wrapper,
        .el-table__header-wrapper {
          border-radius: 0.2rem;
        }

        tr {
          cursor: pointer;
          th {
            height: 0.7rem;
            padding: 0;
            background-color: #e5eeff !important;
            text-align: center;

            .cell {
              display: flex;
              align-items: center;
              justify-content: center;
              word-break: break-word;
              font-size: 0.19rem;
              font-weight: 500;
              color: #555;
              text-transform: uppercase;
              line-height: 1.3;
              .caret-wrapper {
                // display: none;
                width: 10px;
                margin-left: 5px;
                .sort-caret {
                  left: 0;
                }
              }
              .tips {
                display: flex;
                align-items: center;
                justify-content: center;
                img {
                  display: block;
                  width: 20px;
                  height: 20px;
                  margin: 0 0 0 5px;
                  cursor: pointer;
                  @media screen and (max-width: 1440px) {
                    width: 17px;
                    height: 17px;
                  }
                  @media screen and (max-width: 600px) {
                    width: 15px;
                    height: 15px;
                  }
                }
              }
              .tipsWidth {
                width: 110px;
                @media screen and (max-width: 1600px) {
                  width: 95px;
                }
                @media screen and (max-width: 1440px) {
                  width: 90px;
                }
                @media screen and (max-width: 1280px) {
                  width: 80px;
                }
              }
              .el-table__column-filter-trigger {
                i {
                  font-size: 13px;
                  font-weight: 600;
                  margin-left: 4px;
                  transform: scale(1);
                }
              }
            }
          }

          th.is-leaf {
            border-bottom: 0;
          }

          td {
            padding: 0.25rem 0.05rem;
            border-bottom: 1px solid #dfdfdf;

            .cell {
              padding: 0;
              font-size: 0.18rem;
              word-break: break-word;
              color: #000;
              text-align: center;
              line-height: 0.25rem;
              overflow: visible;

              .el-rate__icon {
                font-size: 0.16rem;
                margin-right: 0;
              }
              .el-icon-arrow-right {
                font-weight: bold;
                font-size: 0.13rem;
              }
              .rightClick {
                color: #0c3090;
                cursor: pointer;
              }

              .revoke {
                .el-button {
                  padding: 0.1rem 0.12rem;
                  margin: 0 auto;
                  border-radius: 4px;
                  font-size: inherit;
                }
              }
              .hot-cold-box {
                position: relative;
                display: flex;
                align-items: center;
                justify-content: center;
                flex-wrap: wrap;
                .icon {
                  display: block;
                  width: 22px;
                  height: 22px;
                  margin: 0 0.15rem;
                  font-size: 0.24rem;
                  @media screen and (max-width: 1600px) {
                    width: 18px;
                    height: 18px;
                  }
                  @media screen and (max-width: 768px) {
                    width: 15px;
                    height: 15px;
                  }
                  &:hover {
                    opacity: 0.7;
                  }
                }
                .icon_pay {
                  width: 25px;
                  height: 25px;
                  background: url(../../assets/images/space/pay.png) no-repeat
                    center;
                  background-size: 100% 100%;
                  &.icon_pay_disable {
                    opacity: 0.2;
                  }
                }
                .icon_rename {
                  background: url(../../assets/images/space/icon_02.png)
                    no-repeat center;
                  background-size: 100% 100%;
                }
                .icon_details {
                  background: url(../../assets/images/space/icon_03.png)
                    no-repeat center;
                  background-size: 100% 100%;
                }
                .icon_delete {
                  background: url(../../assets/images/space/icon_04.png)
                    no-repeat center;
                  background-size: 100% 100%;
                }
              }
              .hot-miner {
                .el-button {
                  span {
                    white-space: nowrap;
                  }
                }
              }

              .el-button.el-icon-upload {
                padding: 0 0.03rem;
                line-height: 0.25rem;
                font-size: 0.1372rem;
              }

              .scoreStyle {
                width: 100%;
                clear: both;
                text-align: center;
                color: #ffb822;
                cursor: pointer;
              }

              .scoreStyle:hover {
                text-decoration: underline;
              }
            }
          }

          &:hover {
            td {
              .cell {
                .hot-cold-box {
                  .cidMore {
                    background-color: #f5f7fa;
                  }
                }
              }
            }
          }
        }
      }
      .el-table::before {
        display: none;
      }
      .detailMore {
        padding: 15px 0;
        font-size: 12px;
        color: #888;
        line-height: 1.5;
        text-align: center;
        @media screen and (min-width: 1800px) {
          font-size: 14px;
        }
        a {
          color: inherit;
          text-decoration: underline;
        }
      }
    }
    .fes-host {
      display: flex;
      justify-content: space-between;
      position: absolute;
      left: 0;
      bottom: 0;
      z-index: 21;
      background-color: rgba(0, 0, 0, 0.1);
      font-size: 15px;
      font-weight: 400;
      // width: calc(3.2rem - 0.4rem);
      width: calc(100% - 0.4rem);
      padding: 0.2rem;
      color: hsla(0, 0%, 100%, 0.75);
      transition: all;
      transition-duration: 0.3s;
      i {
        float: left;
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 0.2rem;
        margin-right: 10px;
        width: 20px;
        height: 20px;
        color: #888b83;
        // background: url(../assets/images/icon_01.jpg) no-repeat center;
        // background-size: 100%;
        @media screen and (max-width: 999px) {
          font-size: 18px;
        }
      }
      a {
        color: hsla(0, 5%, 12%, 0.75);
        font-size: 15px;
        font-weight: 400;
      }
      .fesHostLink {
        display: flex;
        width: calc(100% - 40px);
        a {
          display: block;
          width: calc(100% - 30px);
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }
      }
      .fesHostLogout {
        cursor: pointer;
        i {
          margin: 0;
          color: #333;
          &:hover {
            color: #7ecef4;
          }
        }
      }
    }
  }
}

@media screen and (max-width: 1024px) {
  .spaceStyle {
    .slideScroll {
      .form_top {
        .search {
          flex-wrap: wrap;
          height: auto;

          .el-input {
            width: 100%;
            margin: 0.1rem 0;

            .el-input__inner {
              width: 100%;
              font-size: 0.1372rem;
            }
          }

          span {
            margin-left: 0;
          }

          .search_right {
            .el-button {
              padding: 0 0.2rem;
              font-size: 0.1372rem;
            }
          }
        }
      }
    }
  }
}
@media screen and (max-width: 470px) {
  .spaceStyle {
    .slideScroll {
      .form_top {
        .search_file {
          flex-wrap: wrap;
          height: auto;
          .search_right {
            width: 100%;
            margin: 0.05rem 0 0.1rem;
          }
        }
      }
    }
  }
}
</style>
