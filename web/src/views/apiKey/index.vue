<template>
  <div class="spaceStyle" v-loading="listLoad">
    <div class="slideScroll">
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

      <div v-loading="payLoad">
        <h4>{{$t('billing.bill_overview')}}</h4>
        <h6></h6>
        <div class="form_top">
          <div class="search_file">
            <div class="createTask">
              <a v-if="$root.plan_id === 0" @click="payPlan">
                <span>{{$t('billing.bill_btn_pay')}}</span>
              </a>
              <a v-else class="disabled">
                <span>{{$t('billing.bill_btn_paid')}}</span>
              </a>
            </div>
            <router-link :to="{name: 'home_entrance', query: { id: 'pricing' }}" target="_blank" class="billing_cont">
              {{$t('billing.have_faq')}}
            </router-link>
          </div>
        </div>
        <div class="fes-search">
          <el-table :data="billingData" v-loading="billingLoad" stripe style="width: 100%" max-height="300" :empty-text="$t('deal.formNotData')" class="table_cell" @sort-change="sortChange" :default-sort="{prop: 'date', order: 'descending'}">
            <el-table-column prop="tx_hash" :label="$t('billing.TRANSACTION')" min-width="150">
              <template slot-scope="scope">
                <div class="hot-cold-box">
                  <div class="upload_form_right" @click="networkLink(scope.row.tx_hash, scope.row.chain_id)">
                    <p>{{scope.row.tx_hash}}</p>
                  </div>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="amount" :label="$t('billing.AMOUNT')" min-width="90"></el-table-column>
            <el-table-column prop="chain_id" :label="$t('billing.chain_id')" min-width="90"></el-table-column>
            <el-table-column prop="pay_time" :label="$t('billing.PAYMENTDATE')" min-width="120">
              <template slot-scope="scope">
                <div class="hot-cold-box">
                  {{momentFun(scope.row.pay_time)}}
                </div>
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>

      <div>
        <h4>{{$t('my_profile.apiKey_your_Domain')}}</h4>
        <h6></h6>
        <!-- <div class="form_top">
          <div class="search_file">
            <div class="createTask">
              <a @click="dialogFun('add_domain')">
                <img src="@/assets/images/space/icon_01.png" alt="">
                <span>{{$t('my_profile.apiKey_btn_03')}}</span>
              </a>
            </div>
          </div>
        </div> -->
        <div class="fes-search">
          <el-table :data="domainData" stripe style="width: 100%" max-height="300" :empty-text="$t('deal.formNotData')" class="table_cell">
            <el-table-column prop="api_key" :label="$t('my_profile.apiKey_your_Domain_label')">
              <template slot-scope="scope">
                <div>
                  {{scope.row}}
                </div>
              </template>
            </el-table-column>
            <!-- <el-table-column prop="create_at" :label="$t('my_profile.table_apiKey_th_04')">
              <template slot-scope="scope">
                <div style="">
                  {{momentFun(scope.row.create_at)}}
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="qr_code" label="Actions" max-width="160">
              <template slot-scope="scope">
                <div class="revoke">
                  <el-button type="danger" disabled>{{$t('my_profile.apiKey_btn_04')}}</el-button>
                </div>
              </template>
            </el-table-column> -->
          </el-table>
        </div>
      </div>
    </div>

    <el-dialog :title="$t('my_profile.create_api_title01')" :visible.sync="apiTips" :width="width" custom-class="wrongNet" :modal="true" :close-on-click-modal="false">
      <div class="apiTipCont">
        <b>{{$t('my_profile.create_api_tips08')}}</b>
        <label>{{$t('my_profile.create_api_tips03')}}</label>
        <p>{{apiCont.apiKey}}
          <span class="el-icon-document-copy" @click="copyLink(apiCont.apiKey, 1)"></span>
        </p>

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
// import moment from 'moment'
import popUps from '@/components/popups'
// import Qs from 'qs'
import payAbi from '@/utils/pay.json'
import linkTokenAbi from '@/utils/linkToken.json'
let that
export default {
  data () {
    return {
      width: document.body.clientWidth > 600 ? '450px' : '95%',
      listLoad: true,
      listTableLoad: false,
      createLoad: false,
      billingLoad: false,
      toolData: [],
      domainData: [],
      billingData: [],
      areaBody: {},
      dialogFormVisible: false,
      typeName: 'add_apikey',
      apiTips: false,
      apiCont: {
        apiKey: '',
        access: ''
      },
      changeTitle: '',
      payLoad: false
    }
  },
  components: { popUps },
  computed: {
    languageMcs () {
      return this.$store.state.app.languageMcs
    },
    metaAddress () {
      return this.$store.getters.metaAddress
    },
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
    networkLink (hash, chain_id) {
      window.open(`${this.baseAddressURL}tx/${hash}`)
    },
    async sortChange (column) {
      // this.billingLoad = true
      // this.getListBuckets()
    },
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
    async getPopUps (dialog, rows, day) {
      // console.log('rows', rows)
      switch (rows) {
        case 'delete':
          that.deleteApiKey()
          break
        case 'add_domain':
          that.createDomain(dialog, rows, day)
          break
        case 'refresh':
          that.refreshFun(dialog)
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
      const apiKeyRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/user/generate_api_key?${Qs.stringify(params)}`, 'get')
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
      that.payLoad = false
      document.getElementById('content_client').scrollTop = 120
    },
    async deleteApiKey () {
      that.listTableLoad = true
      const params = {
        'apikey': that.areaBody.api_key
      }
      const deleteRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/user/delete_api_key?${Qs.stringify(params)}`, 'put')
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
        that.$message({
          showClose: true,
          message: directoryRes.message ? directoryRes.message : 'Fail',
          type: 'error'
        })
        that.toolData = []
      } else that.toolData = directoryRes.data.apikey || []

      const domainRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/gateway/get_gateway`, 'get')
      if (!domainRes || domainRes.status !== 'success') {
        that.$message.error(domainRes.message ? domainRes.message : 'Fail')
        that.domainData = []
      } else that.domainData = domainRes.data || []

      const billingRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/payment/get_payment_history`, 'get')
      if (!billingRes || billingRes.status !== 'success') {
        that.$message.error(billingRes.message ? billingRes.message : 'Fail')
        that.billingData = []
      } else that.billingData = billingRes.data || []

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
    },
    async refreshFun (dialog) {
      that.dialogFormVisible = dialog
      that.payLoad = true
      let paymentRes = await that.$commonFun.sendRequest(`${that.baseAPIURL}api/v2/payment/get_payment`, 'get')
      if (!paymentRes || paymentRes.status !== 'success') return false
      else {
        that.$root.plan_id = paymentRes.data.plan_id
        that.$root.max_storage = paymentRes.data.max_storage
      }
      that.payLoad = false
    },
    async getUnit (id) {
      let name = ''
      switch (id) {
        case 97:
          name = 'BSC TestNet'
          break
        case 137:
          name = 'Polygon Mainnet'
          break
        case 80001:
          name = 'Mumbai Testnet'
          break
      }
      return ({
        name
      })
    },
    async payPlan () {
      that.payLoad = true
      const getID = await that.$commonFun.web3Init.eth.net.getId()
      if (that.$root.chain_id !== getID) {
        const { name } = await that.getUnit(Number(that.$root.chain_id))
        that.$message.error(that.languageMcs === 'en' ? `Please switch to the network: ${name}` : `请切换到网络：${name}`)
        that.payLoad = false
        return false
      }
      let tokenFactory = new that.$web3Init.eth.Contract(linkTokenAbi, that.$root.USDC_ADDRESS)
      let payObject = {
        from: that.metaAddress,
        gasPrice: await that.$web3Init.eth.getGasPrice()
      }
      let account = that.$root.pay_account

      tokenFactory.methods.approve(that.$root.PAYMENT_CONTRACT_ADDRESS, account).send(payObject)
        .then(receipt => {
          console.log('approve receipt:', receipt)
          that.contractSend()
        })
        .catch((err) => {
          that.payLoad = false
          if (err && err.message) that.$message.error(err.message)
        })
    },
    async contractSend () {
      let payFactory = new that.$web3Init.eth.Contract(payAbi, that.$root.PAYMENT_CONTRACT_ADDRESS)
      let estimatedGas = await payFactory.methods
        .pay(1)
        .estimateGas({ from: that.metaAddress })
      let gasLimit = Math.floor(estimatedGas * 1.5)

      await payFactory.methods.pay(1)
        .send({ from: that.metaAddress, gasLimit: gasLimit })
        .on('transactionHash', async (hash) => {
          console.log('hash console:', hash)
          that.dialogFun('billing_tip')
        })
        .on('confirmation', function (confirmationNumber, receipt) {
          // console.log('confirmationNumber console:', confirmationNumber, receipt)
        })
        .on('receipt', function (receipt) {
          // receipt example
          console.log('create receipt console:', receipt)
        })
        .on('error', function (error) {
          if (error && error.message) that.$message.error(error.message)
        })
      that.payLoad = false
    }
  },
  mounted () {
    that = this
    document.getElementById('content-box').scrollTop = 0
    that.$store.dispatch('setRouterMenu', 21)
    that.$store.dispatch('setHeadertitle', that.$t('route.ApiKey'))
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
    },
    balanceMweiFilter (value) {
      if (String(value) === '0') return 0
      if (!value) return '-'
      // if (!Number(value)) return 0;
      // if (isNaN(value)) return value;
      // 18 - 单位换算需要 / 1000000000000000000，浮点运算显示有bug
      // value = Number(value)
      if (String(value).length > 9) {
        let v1 = String(value).substring(0, String(value).length - 9)
        let v2 = String(value).substring(String(value).length - 9)
        let v3 = String(v2).replace(/(0+)\b/gi, '')
        if (v3) {
          return v1 + '.' + v3
        } else {
          return v1
        }
      } else {
        let v3 = ''
        for (let i = 0; i < 9 - String(value).length; i++) {
          v3 += '0'
        }
        return '0.' + String(v3 + value).replace(/(0+)\b/gi, '')
      }
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
        b {
          display: block;
          margin: 0 0 0.2rem;
          line-height: 1.2;
          color: #f44336;
          font-size: 14px;
          font-weight: 600;
          @media screen and (max-width: 768px) {
            font-size: 12px;
          }
        }
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
  // width: calc(100% - 0.6rem);
  width: 100%;
  padding: 0 0 0.4rem;
  margin: 0.3rem 0;
  background-color: #fff;
  border-radius: 0.1rem;
  @media screen and (max-width: 600px) {
    width: 100%;
    padding: 0;
    margin: 0.3rem 0;
  }
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
    padding: 0.3rem 0;
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
            &.disabled {
              opacity: 0.5;
              &:hover {
                box-shadow: none;
              }
            }
          }
        }
        .billing_cont {
          margin-left: 0.15rem;
          font-size: 14px;
          color: #333;
          @media screen and (max-width: 600px) {
            font-size: 12px;
          }
          &:hover {
            text-decoration: underline;
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
        font-size: 0.25rem;
        color: #4f87ff;
        line-height: 1.2;
        border: dashed;
        border-radius: 0.1rem;
        @media screen and (max-width: 1600px) {
          font-size: 0.23rem;
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
                .upload_form_right {
                  padding: 0 10px;
                  text-align: left;
                  &:hover {
                    text-decoration: underline;
                  }
                }
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
  .billing_style {
    padding: 0.2rem 0 1rem;
    @media screen and (max-width: 992px) {
      padding: 0.2rem 0 0.8rem;
    }
    @media screen and (max-width: 600px) {
      padding: 0.2rem 0 0.6rem;
    }
    .el-col {
      margin-bottom: 0.2rem;
      .billing_cont {
        display: flex;
        align-items: center;
        cursor: pointer;
        color: #333;
        font-size: 14px;
        .icon {
          position: relative;
          width: 0.7rem;
          height: 0.7rem;
          margin-right: 0.2rem;
          &::before {
            position: absolute;
            left: 0;
            top: 0;
            width: 100%;
            height: 100%;
            content: "";
            background-size: 0.25rem !important;
            border-radius: 0.08rem;
          }
          &.el-icon-s-billing {
            &::before {
              background: #753fdd
                url(../../assets/images/menuIcon/billing@2x-1.png) no-repeat
                center;
            }
          }
          &.el-icon-s-pricing {
            &::before {
              background: #f1a44d url(../../assets/images/menuIcon/billing.png)
                no-repeat center;
            }
          }
        }
        .desc {
          .desc_t {
            font-size: 14px;
            @media screen and (min-width: 1800px) {
              font-size: 16px;
            }
          }
          .desc_d {
            font-size: 12px;
            color: #787889;
            @media screen and (min-width: 1800px) {
              font-size: 14px;
            }
          }
        }
        &:hover {
          background: #f7f7f7;
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
