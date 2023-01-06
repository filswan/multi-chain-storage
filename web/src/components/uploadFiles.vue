<template>
  <div id="Create">
    <el-dialog :modal="true" :close-on-click-modal="false" :width="widthDia" :visible.sync="uploadDigShow" custom-class="uploadDig" id="uploadDigBody" :before-close="closeDia">
      <div class="uploadDigID el-icon-upload" :style="{'z-index': bgStyle ? '99' : '-1'}"></div>
      <div class="loadMetamaskPay" v-if="loading">
        <div>
          <div class="el-loading-spinner">
            <svg viewBox="25 25 50 50" class="circular">
              <circle cx="50" cy="50" r="20" fill="none" class="path"></circle>
            </svg>
            <!---->
          </div>
          <p v-if="loadMetamaskPay">{{$t('uploadFile.payment_tip')}}</p>
        </div>
      </div>
      <template slot="title">
        {{$t('uploadFile.upload')}}

        <el-tooltip effect="dark" :content="$t('uploadFile.uploadFile_title')" placement="top">
          <img src="@/assets/images/info.png" />
        </el-tooltip>
      </template>
      <div class="upload_form">
        <div>
          <el-upload class="upload-demo" :style="{'border': ruleForm.fileList_tip?'2px dashed #f56c6c':'0'}" drag ref="uploadFileRef" action="customize" :http-request="uploadFile" :file-list="ruleForm.fileList" :on-change="handleChange" :on-remove="handleRemove">
            <svg t="1666583167832" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="8421" width="128" height="128">
              <path d="M795.694545 395.636364a289.047273 289.047273 0 0 0-567.38909 0 197.585455 197.585455 0 0 0 18.385454 395.636363h30.952727a23.272727 23.272727 0 0 0 0-46.545454H246.690909a151.272727 151.272727 0 1 1-2.327273-302.545455l23.272728 5.352727 3.025454-25.6a242.269091 242.269091 0 0 1 480.814546 0l4.654545 25.134546 23.272727-4.887273a151.272727 151.272727 0 1 1-2.327272 302.545455h-34.909091a23.272727 23.272727 0 0 0 0 46.545454h35.141818a197.585455 197.585455 0 0 0 18.385454-395.636363z"
                p-id="8422" fill="#8a8a8a"></path>
              <path d="M528.523636 480.349091a23.272727 23.272727 0 0 0-33.047272 0l-131.490909 131.490909a23.272727 23.272727 0 0 0 0 33.047273 23.272727 23.272727 0 0 0 32.814545 0L488.727273 552.96V837.818182a23.272727 23.272727 0 0 0 46.545454 0V552.96l93.090909 91.927273a23.272727 23.272727 0 0 0 16.523637 6.749091 23.272727 23.272727 0 0 0 16.290909-39.796364z"
                p-id="8423" fill="#8a8a8a"></path>
            </svg>
            <div class="el-upload__text">
              <small>{{$t('uploadFile.uploadDray')}}</small>
            </div>
            <div class="el-upload__text">
              <small>{{$t('uploadFile.uploadDray_or')}}</small>
            </div>
            <el-button size="small">
              {{$t('uploadFile.uploadDray_text')}}
            </el-button>
          </el-upload>
          <p v-if="ruleForm.fileList.length>0" style="display: flex;align-items: center;">
            <svg t="1637031488880" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3310" style="width: 14px;height: 14px;margin: 0 6px 0 5px;">
              <path d="M512 1024a512 512 0 1 1 512-512 32 32 0 0 1-32 32h-448v448a32 32 0 0 1-32 32zM512 64a448 448 0 0 0-32 896V512a32 32 0 0 1 32-32h448A448 448 0 0 0 512 64z" fill="#999999" p-id="3311"></path>
              <path d="M858.88 976a32 32 0 0 1-32-32V640a32 32 0 0 1 32-32 32 32 0 0 1 32 32v304a32 32 0 0 1-32 32z" fill="#999999" p-id="3312"></path>
              <path d="M757.12 773.12a34.56 34.56 0 0 1-22.4-8.96 32 32 0 0 1 0-45.44l101.12-101.12a32 32 0 0 1 45.44 0 30.72 30.72 0 0 1 0 44.8l-101.12 101.76a34.56 34.56 0 0 1-23.04 8.96z" fill="#999999" p-id="3313"></path>
              <path d="M960 773.12a32 32 0 0 1-22.4-8.96l-101.76-101.76a32 32 0 0 1 0-44.8 32 32 0 0 1 45.44 0l101.12 101.12a32 32 0 0 1-22.4 54.4z" fill="#999999" p-id="3314"></path>
            </svg>
            {{ruleForm.file_size}}
          </p>
          <p v-if="ruleForm.fileList_tip" style="color: #F56C6C;font-size: 12px;line-height: 1;">{{ruleForm.fileList_tip_text}}</p>
        </div>
        <el-form :model="ruleForm" status-icon :rules="rules" ref="ruleForm" class="demo-ruleForm">
          <!-- <el-form-item prop="fileList" :label="$t('uploadFile.upload')" :style="{'align-items': ruleForm.fileList_tip?'flex-start':'center'}"></el-form-item> -->
          <el-form-item prop="duration">
            <template slot="label">
              {{$t('uploadFile.Duration')}}

              <el-tooltip effect="dark" :content="$t('uploadFile.Duration_tooltip')" placement="top">
                <img src="@/assets/images/info.png" />
              </el-tooltip>
            </template>
            {{ruleForm.duration}} {{$t('components.day')}}
          </el-form-item>
          <el-form-item prop="storage_copy">
            <template slot="label">
              {{$t('uploadFile.Storage_copy')}}

              <el-tooltip effect="dark" :content="$t('uploadFile.Storage_copy_tooltip')" placement="top">
                <img src="@/assets/images/info.png" />
              </el-tooltip>
            </template>
            {{ruleForm.storage_copy}}
          </el-form-item>
          <el-form-item prop="Free_Storage_Capacity">
            <template slot="label">
              {{$t('uploadFile.Free_Storage_Capacity')}}

              <el-tooltip effect="dark" :content="$t('uploadFile.Free_Storage_Capacity_tooltip')" placement="top">
                <img src="@/assets/images/info.png" />
              </el-tooltip>
            </template>
            <span style="color:#2C7FF8">{{free_quota_per_month-free_usage | byteStorage}} GB</span>
          </el-form-item>
          <el-form-item prop="storage_cost">
            <template slot="label">
              {{$t('uploadFile.Estimated_Storage_Cost')}}

              <el-tooltip effect="dark" :content="$t('uploadFile.Estimated_Storage_Cost_tooltip')" placement="top">
                <img src="@/assets/images/info.png" />
              </el-tooltip>
            </template>
            <span v-if="free" style="color:#2C7FF8;text-transform: uppercase;">{{$t('uploadFile.filter_status_Free')}}</span>
            <span v-else style="color:#2C7FF8">{{ruleForm.storage_cost | NumStorage}} FIL</span>
          </el-form-item>
        </el-form>
        <div class="upload_plan">
          <div class="title" :style="{'color': ruleForm.lock_plan_tip? '#f67e7e' : '#000'}">
            {{$t('uploadFile.Select_Lock_Funds_Plan')}}
            <el-tooltip effect="dark" :content="$t('uploadFile.Select_Lock_Funds_Plan_tooltip')" placement="top">
              <img src="@/assets/images/info.png" />
            </el-tooltip>
          </div>
          <div class="desc">{{$t('uploadFile.latest_exchange_rate')}} {{biling_price}}.</div>
          <div class="upload_plan_radio" :class="{'upload_plan_radio_free': free}">
            <el-radio-group :disabled="free?true:false" v-model="ruleForm.lock_plan" @change="agreeChange">
              <el-radio label="1" border>
                <div class="title">{{$t('uploadFile.Low')}}</div>
                <div class="cont">
                  {{free?0:storage_cost_low}} USDC
                </div>
              </el-radio>
              <el-radio label="2" border>
                <div class="title">{{$t('uploadFile.Average')}}</div>
                <div class="cont">
                  {{free?0:storage_cost_average}} USDC
                </div>
              </el-radio>
              <el-radio label="3" border>
                <div class="title">{{$t('uploadFile.High')}}</div>
                <div class="cont">
                  {{free?0:storage_cost_high}} USDC
                </div>
              </el-radio>
            </el-radio-group>
          </div>
        </div>
        <div class="upload_bot">
          <div class="found">
            <el-button type="primary" class="cancel" @click="closeDia">{{$t('deal.Cancel')}}</el-button>
            <el-button type="primary" @click="submitForm('ruleForm')">{{$t('deal.Submit')}}</el-button>
          </div>
          <a :href="found_link" target="_blank">{{$t('uploadFile.upload_funds')}}</a>
        </div>
      </div>
    </el-dialog>

    <el-dialog title="" :visible.sync="finishTransaction" :close-on-click-modal="false" :width="width" custom-class="completeDia">
      <img src="@/assets/images/alert-icon.png" class="resno" />
      <h1>{{$t('uploadFile.COMPLETED')}}!</h1>
      <h3>{{$t('uploadFile.SUCCESS')}}</h3>
      <a :href="baseAddressURL+'tx/'+txHash" target="_blank">{{txHash}}</a>
      <a class="a-close" @click="finishClose">{{$t('uploadFile.CLOSE')}}</a>
    </el-dialog>

    <el-dialog title="" :visible.sync="failTransaction" :width="width" custom-class="completeDia">
      <img src="@/assets/images/error.png" class="resno" />
      <h1>{{$t('uploadFile.Fail')}}!</h1>
      <h3>{{$t('uploadFile.FailTIP')}}</h3>
      <a :href="baseAddressURL+'tx/'+txHash" target="_blank">{{txHash}}</a>
      <a class="a-close" @click="failClose">{{$t('uploadFile.CLOSE')}}</a>
    </el-dialog>

    <el-dialog title="" :visible.sync="waitTransaction" :width="width" :before-close="failClose" custom-class="completeDia">
      <img src="@/assets/images/waiting.png" class="resno" />
      <h1>{{$t('uploadFile.waiting')}}</h1>
      <h3>{{$t('uploadFile.waitingTIP')}}</h3>
      <a :href="baseAddressURL+'tx/'+txHash" target="_blank">{{txHash}}</a>
      <a class="a-close" @click="failClose">{{$t('uploadFile.CLOSE')}}</a>
    </el-dialog>

    <el-dialog :visible.sync="fileUploadVisible" :show-close="false" :close-on-click-modal="false" :width="widthUpload" custom-class="fileUpload">
      <span slot="title">{{$t('uploadFile.File_uploading')}}... {{percentIn?'('+percentIn+'%)':''}}</span>
      <h3>{{$t('uploadFile.File_uploading_tooltip')}}</h3>
      <img src="@/assets/images/upload.gif" class="gif_img" alt="">
      <small>Current upload speed: {{speedChange(uploadPrecentSpeed)}}</small>
    </el-dialog>

    <el-dialog title="" :visible.sync="paymentPopup" :width="width" custom-class="completeDia">
      <img src="@/assets/images/box-important.png" class="resno" />
      <h2>{{$t('uploadFile.file_uploaded')}}</h2>
      <h4>{{$t('uploadFile.file_uploaded_tip')}}</h4>
      <h4>{{$t('uploadFile.file_uploaded_tip01')}}</h4>
      <a class="a-close" @click="paymentPopup=false">{{$t('uploadFile.OK')}}</a>
    </el-dialog>

    <el-dialog title="" :visible.sync="paymentPopup01" :width="width" custom-class="completeDia">
      <img src="@/assets/images/box-important.png" class="resno" />
      <h2>{{$t('uploadFile.file_uploaded')}}</h2>
      <h4>{{$t('uploadFile.file_uploaded_tip02')}}</h4>
      <h4>{{$t('uploadFile.file_uploaded_tip01')}}</h4>
      <a class="a-close" @click="paymentPopup01=false">{{$t('uploadFile.OK')}}</a>
    </el-dialog>

    <el-dialog title="" :visible.sync="metamaskLoginTip" :width="width" custom-class="completeDia">
      <img src="@/assets/images/box-important.png" class="resno" />
      <h4 v-if="baseNetwork">
        {{$t('fs3Login.toptip_01')}} {{metaNetworkInfo.name}} {{$t('fs3Login.toptip_02')}}
        <b>Polygon Mainnet</b>.
      </h4>
      <h4 v-else>
        {{$t('fs3Login.toptip_01')}} {{metaNetworkInfo.name}} {{$t('fs3Login.toptip_02')}}
        <b>Mumbai Testnet</b>
        {{$t('fs3Login.toptip_Network')}}
        <b>BSC TestNet</b>.
      </h4>
      <a class="a-close" @click="metamaskLoginTip=false">{{$t('uploadFile.OK')}}</a>
    </el-dialog>
  </div>
</template>

<script>
// import bus from '@/components/bus'
// import * as myAjax from '@/api/uploadFile'
import axios from 'axios'
import firstContractJson from '@/utils/swanPayment.json'
import erc20ContractJson from '@/utils/ERC20.json'
let contractErc20
let that
export default {
  name: 'uploadFiles',
  data () {
    var validateDuration = (rule, value, callback) => {
      if (!value) {
        return callback(new Error(that.$t('uploadFile.Duration_tip')))
      } else {
        callback()
      }
    }
    return {
      tableData: {
        file_name: '',
        task_name: '',
        payload_cid: ''
      },
      ruleForm: {
        duration: '525',
        storage_copy: '5',
        fileList: [],
        fileList_tip: false,
        fileList_tip_text: '',
        file_size: '',
        file_size_byte: '',
        storage_cost: '',
        lock_plan: '2',
        lock_plan_tip: false,
        amount: '',
        amount_minprice: 0,
        gaslimit: this.$root.PAY_GAS_LIMIT
      },
      rules: {
        duration: [
          { validator: validateDuration, trigger: 'blur' }
        ],
        fileList: [
          { type: 'array', required: true, message: this.$t('uploadFile.file_name_tip'), trigger: 'change' }
        ]
      },
      loading: false,
      bodyWidth: document.documentElement.clientWidth < 1024,
      fileListTip: false,
      biling_price: 0,
      storage: 0,
      storage_cost_low: 0,
      storage_cost_average: 0,
      storage_cost_high: 0,
      center_fail: false,
      centerDialogVisible: false,
      modelClose: true,
      width: document.body.clientWidth > 600 ? '400px' : '95%',
      widthUpload: document.body.clientWidth > 600 ? '450px' : '95%',
      widthDia: document.body.clientWidth <= 600 ? '95%' : document.body.clientWidth <= 1440 ? '7rem' : '6.6rem',
      gatewayContractAddress: this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS,
      recipientAddress: this.$root.RECIPIENT,
      usdcAddress: this.$root.USDC_ADDRESS,
      hashload: false,
      timer: '',
      usdcAvailable: '',
      network: {
        name: '',
        unit: '',
        text: false
      },
      finishTransaction: false,
      failTransaction: false,
      waitTransaction: false,
      txHash: '',
      fileUploadVisible: false,
      paymentPopup: false,
      paymentPopup01: false,
      percentIn: '',
      loadMetamaskPay: false,
      metamaskLoginTip: false,
      lastTime: 0,
      found_link: process.env.NODE_ENV === 'production' ? 'https://calibration-faucet.filswan.com/' : 'http://192.168.88.216:8080/faucet/#/dashboard',
      free: false,
      bgStyle: false,
      lastUploadTime: 0,
      lastUploadSize: 0,
      uploadPrecentSpeed: 0
    }
  },
  props: ['uploadDigShow'],
  components: {},
  watch: {
    'ruleForm.fileList': function () {
      this.ruleForm.lock_plan = '2'
      this.$refs['ruleForm'].validate((valid) => { })
      if (this.ruleForm.file_size_byte > 0) {
        this.calculation()
      }
    }
  },
  methods: {
    calculation (type) {
      that.ruleForm.storage_cost = that.ruleForm.file_size_byte * that.ruleForm.duration * that.storage
      let _price = that.ruleForm.storage_cost * that.biling_price
      let numberPrice = Number(_price).toFixed(6)
      that.ruleForm.amount_minprice = numberPrice > 0.000001 ? numberPrice : '0.0000005'
      that.storage_cost_low = numberPrice > 0 ? Number(_price * 2).toFixed(6) : '0.000002'
      that.storage_cost_average = numberPrice > 0 ? Number(_price * 3).toFixed(6) : '0.000003'
      that.storage_cost_high = numberPrice > 0 ? Number(_price * 5).toFixed(6) : '0.000004'
      that.ruleForm.amount = that.free ? 0 : that.storage_cost_average
    },
    agreeChange (val) {
      that.ruleForm.lock_plan_tip = false
      switch (val) {
        case '1':
          that.ruleForm.amount = that.storage_cost_low
          break
        case '2':
          that.ruleForm.amount = that.storage_cost_average
          break
        case '3':
          that.ruleForm.amount = that.storage_cost_high
          break
        default:
          that.ruleForm.amount = that.storage_cost_low
      }
    },
    closeDia () {
      that.$emit('getUploadDialog', false)
    },
    submitForm (formName) {
      that.$refs[formName].validate(async (valid) => {
        if (valid) {
          if (!that._file) {
            that.ruleForm.fileList_tip = true
            return false
          }
          let status = await that.$metaLogin.netStatus(that.networkID)
          if (that.metaAddress && !status) {
            that.metamaskLoginTip = true
            return false
          }
          if (!that.ruleForm.lock_plan || (that.ruleForm.amount <= 0 && !that.free)) {
            that.ruleForm.lock_plan_tip = true
            return false
          }
          if (that.ruleForm.fileList_tip || isNaN(that.ruleForm.amount)) return false

          let now = new Date().valueOf()
          if (that.lastTime === 0) {
            that.lastTime = now
          } else {
            if ((now - that.lastTime) > 2000) {
              // 重置上一次点击时间，设置的2秒间隔
              that.lastTime = now
            } else {
              return false
            }
          }

          if (that.metaAddress) {
            // 授权代币
            contractErc20 = new that.$web3Init.eth.Contract(erc20ContractJson)
            contractErc20.options.address = that.usdcAddress
            // 查询剩余代币余额为：
            contractErc20.methods.balanceOf(that.metaAddress).call()
              .then(resultUSDC => {
                that.usdcAvailable = that.$web3Init.utils.fromWei(resultUSDC, 'mwei')
                console.log('Available upload:', resultUSDC, that.usdcAvailable, that.ruleForm.amount)

                // 判断支付金额是否大于代币余额
                if (Number(that.ruleForm.amount) > Number(that.usdcAvailable)) {
                  that.$message.error('Insufficient balance')
                  return false
                }

                // 通过 FormData 对象上传文件
                var formData = new FormData()
                formData.append('file', that._file)
                formData.append('duration', that.ruleForm.duration)
                formData.append('storage_copy', that.ruleForm.storage_copy)
                formData.append('wallet_address', that.metaAddress)
                // formData.append('task_name', that.ruleForm.task_name)
                that.loading = true
                that.fileUploadVisible = true
                that.lastUploadTime = 0
                that.lastUploadSize = 0
                let xhr = new XMLHttpRequest()
                xhr.open('POST', `${that.baseAPIURL}api/v1/storage/ipfs/upload`, true) // 设置xhr得请求方式和url。
                xhr.withCredentials = false
                const token = that.$store.getters.mcsjwtToken
                if (token) {
                  xhr.setRequestHeader(
                    'Authorization',
                    'Bearer ' + that.$store.getters.mcsjwtToken
                  )
                }
                let i = 0

                xhr.onreadystatechange = function () { // 等待ajax请求完成。
                  if (xhr.status === 200) {
                    if (xhr.responseText) {
                      that.fileUploadVisible = false
                      let res = JSON.parse(xhr.responseText)
                      i += 1
                      if (i <= 1) {
                        that.percentIn = xhr.readyState === 4 ? '100%' : that.percentIn
                        if (res.status === 'success') {
                          if (res.data.status === 'Free') {
                            that.finishClose()
                            return false
                          } else {
                            contractErc20.methods.allowance(that.gatewayContractAddress, that.metaAddress).call()
                              .then(async resultUSDC => {
                                let amountPay = that.$web3Init.utils.toWei(that.ruleForm.amount, 'mwei')
                                console.log('allowance：' + resultUSDC, amountPay)
                                if (resultUSDC < amountPay) {
                                  let payObject = {
                                    from: that.metaAddress,
                                    gas: that.$web3Init.utils.toHex(that.ruleForm.gaslimit),
                                    gasPrice: await that.$web3Init.eth.getGasPrice()
                                  }
                                  contractErc20.methods.approve(that.gatewayContractAddress, amountPay).send(payObject)
                                    .then(receipt => {
                                      // console.log('approve receipt:', receipt)
                                      that.contractSend(res.data, amountPay)
                                    })
                                    .catch(() => {
                                      // console.log('errorerrorerror', error)
                                      that.finishClose()
                                    })
                                } else {
                                  that.contractSend(res.data, amountPay)
                                }
                              })
                          }
                        } else {
                          that.$message.error(res.message || that.$t('uploadFile.xhr_tip'))
                          that.finishClose()
                        }
                      }
                    }
                  } else {
                    i += 1
                    if (i <= 1) {
                      if (xhr.status === 500) that.$message.error(that.$t('uploadFile.xhr_error500'))
                    }
                    that.loading = false
                    that.fileUploadVisible = false
                  }
                  xhr.upload.addEventListener('error', event => {
                    that.$message.error(that.$t('uploadFile.xhr_tip'))
                  })

                  xhr.upload.addEventListener('progress', event => {
                    if (event.lengthComputable) {
                      // let loaded = event.loaded
                      // let total = event.total
                      // console.log('total-loaded', total, loaded)
                      let percentIn = Math.floor(event.loaded / event.total * 100)
                      // 设置进度显示
                      that.percentIn = percentIn + '%'
                      // console.log(percentIn+'%-')
                    }
                  })
                }
                // 获取上传进度
                xhr.upload.onprogress = function (event) {
                  // console.log('event.onprogress', event)
                  if (event.lengthComputable) that.progressHandle(event)
                }
                xhr.send(formData)
                return false
              })
          }
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    async progressHandle (e) {
      const { loaded, total } = e
      if (that.lastUploadTime === 0) {
        that.lastUploadTime = new Date().getTime()
        that.lastUploadSize = loaded
        return
      }
      let nowTime = new Date().getTime()
      let intervalTime = (nowTime - that.lastUploadTime) / 1000
      let intervalSize = loaded - that.lastUploadSize
      that.lastUploadTime = nowTime
      that.lastUploadSize = loaded

      that.uploadPrecentSpeed = intervalSize / intervalTime // Calculation speed
      // const leftTime = ((total - loaded) / that.uploadPrecentSpeed) // Calculate remaining time
      const uploadPrecent = ((loaded / total) * 100) | 0 // Calculation progress
      that.percentIn = loaded < total && uploadPrecent === 100 ? 99 : uploadPrecent

      // console.log('当前已上传文件大小: ' + loaded, '总文件大小: ', total, 'progress: ', uploadPrecent + '%')
    },
    speedChange (bytes) {
      if (String(bytes) === '0') return '0 byte/s'
      if (!bytes) return '-'
      var k = 1024 // or 1000
      var sizes = ['byte/s', 'kb/s', 'mb/s', 'gb/s']
      var i = Math.floor(Math.log(bytes) / Math.log(k))
      if (Math.round((bytes / Math.pow(k, i))).toString().length > 3) i += 1
      return (bytes / Math.pow(k, i)).toFixed(1) + ' ' + sizes[i]
    },
    async contractSend (resData, amountPay) {
      // 合约转账
      let contractInstance = new that.$web3Init.eth.Contract(firstContractJson)
      contractInstance.options.address = that.gatewayContractAddress
      // console.log( 'contractInstance合约实例：', contractInstance );
      // console.log(contractInstance.options.jsonInterface)

      let payObject = {
        from: that.metaAddress,
        gas: that.$web3Init.utils.toHex(that.ruleForm.gaslimit),
        gasPrice: await that.$web3Init.eth.getGasPrice()
      }
      let lockObj = {
        id: resData.w_cid,
        minPayment: String(Math.floor(amountPay / 2)),
        amount: amountPay,
        lockTime: 86400 * Number(that.$root.LOCK_TIME), // one day
        recipient: that.recipientAddress, // todo:
        size: resData.file_size,
        copyLimit: Number(that.ruleForm.storage_copy)
      }
      console.log(lockObj)
      contractInstance.methods.lockTokenPayment(lockObj)
        .send(payObject)
        .on('transactionHash', function (hash) {
          // console.log('hash console:', hash);
          that.loadMetamaskPay = true
          that.txHash = hash
        })
        .on('confirmation', function (confirmationNumber, receipt) {
          // console.log('confirmationNumber console:', confirmationNumber, receipt);
        })
        .on('receipt', function (receipt) {
          // receipt example
          console.log('receipt console:', receipt)
          that.checkTransaction(receipt.transactionHash, resData, lockObj)
          that.txHash = receipt.transactionHash
        })
        .on('error', function (error) {
          console.log('lockTokenPayment error console:', error)
          // console.error
          that.loading = false
          that.loadMetamaskPay = false
          if (that.finishTransaction) return false
          if (that.txHash) that.waitTransaction = true
          else that.failTransaction = true
        })
    },
    failClose () {
      that.failTransaction = false
      that.waitTransaction = false
      that.txHash = ''
    },
    checkTransaction (txHash, resData, lockObj) {
      that.$web3Init.eth.getTransactionReceipt(txHash).then(
        async res => {
          console.log('checking ... ')
          if (!res) {
            that.timer = setTimeout(() => { that.checkTransaction(txHash, resData, lockObj) }, 2000)
            return false
          } else {
            setTimeout(function () {
              that.loading = false
              that.loadMetamaskPay = false
              clearTimeout(that.timer)
              that.finishTransaction = true
            }, 2000)
          }
        },
        err => { console.error(err) }
      )
    },
    finishClose () {
      that.finishTransaction = false
      that.txHash = ''
      that.$emit('getUploadDialog', false, true)
    },
    // 文件上传
    uploadFile (params) {
      let reg = new RegExp(' ', 'g')
      that._file = params.file
      if (that._file.name.indexOf(' ') > -1) {
        let name = that._file.name.replace(reg, '_')
        that._file = new File([params.file], name, params.file)
      }
      const isLt2M = params.file.size / 1024 / 1024 / 1024 <= 25 // or 1000
      that.ruleForm.file_size = that.sizeChange(params.file.size)
      that.ruleForm.file_size_byte = that.byteChange(params.file.size)
      that.free = (that.free_quota_per_month - that.free_usage) >= params.file.size
      if (!isLt2M) {
        that.ruleForm.fileList_tip = true
        that.ruleForm.fileList_tip_text = that.$t('deal.upload_form_file_tip')
        return false
      } else that.ruleForm.fileList_tip = false
    },
    sizeChange (bytes) {
      if (bytes === 0) return '0 B'
      if (!bytes) return '-'
      var k = 1024 // or 1000
      var sizes = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
      var i = Math.floor(Math.log(bytes) / Math.log(k))

      if (Math.round((bytes / Math.pow(k, i))).toString().length > 3) {
        // 判断大小是999999999左右，解决会显示成1.00e+3科学计数法
        i += 1
      }

      // if(i == 2) return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
      return Number(bytes / Math.pow(k, i)) + ' ' + sizes[i]
    },
    byteChange (limit) {
      var size = ''
      // 只转换成GB
      if (limit <= 0) {
        return '-'
      } else {
        size = limit / (1024 * 1024 * 1024) // or 1000
      }
      return size
      // return Number(size).toFixed(3);
    },
    handleChange (file, fileList) {
      if (fileList.length > 0) that.ruleForm.fileList = [fileList[fileList.length - 1]] // 这一步，是 展示最后一次选择的csv文件
    },
    handleRemove (file, fileList) {
      // console.log(file, fileList);
      that.ruleForm.fileList = []
      that.ruleForm.file_size = ''
      that.ruleForm.file_size_byte = ''
    },
    async stats () {
      that.loading = true
      if (that.$root.SWAN_PAYMENT_CONTRACT_ADDRESS) {
        that.gatewayContractAddress = that.$root.SWAN_PAYMENT_CONTRACT_ADDRESS
        that.usdcAddress = that.$root.USDC_ADDRESS
        that.recipientAddress = that.$root.RECIPIENT

        const storageRes = await that.sendRequest(`${process.env.BASE_API}stats/storage?wallet_address=${that.metaAddress}`)
        let cost = storageRes.data.historical_average_price_verified ? storageRes.data.historical_average_price_verified.split(' ') : []
        if (cost[0]) that.storage = cost[0]

        that.biling_price = that.$root.filecoin_price

        that.loading = false
        // that.addEvent()
      } else {
        setTimeout(function () {
          that.stats()
        }, 1000)
      }
    },
    async sendRequest (apilink) {
      try {
        const response = await axios.get(apilink, {
          headers: {
            'Authorization': 'Bearer ' + that.$store.getters.mcsjwtToken
          }
        })
        return response.data
      } catch (err) {
        console.error(err)
      }
    },
    dropHandler (e) {
      e.preventDefault()
      const fileList = e.dataTransfer.files
      if (fileList.length === 0) {
        return
      }
      const uploadFile = {
        file: fileList[0]
      }
      that.uploadFile(uploadFile)
      that.ruleForm.fileList = [fileList[0]]
      that.bgStyle = false
    },
    addEvent () {
      const oDragBody = document.querySelector('.uploadDig')
      oDragBody.addEventListener('dragenter', (e) => {
        e.preventDefault() // 拖进
        that.bgStyle = true
      })
      oDragBody.addEventListener('dragover', (e) => {
        e.preventDefault() // 清除默认事件
      })
      const oDragWrap = document.querySelector('.uploadDigID')
      oDragWrap.addEventListener('dragenter', (e) => {
        e.preventDefault() // 拖进
      })
      oDragWrap.addEventListener('dragleave', (e) => {
        e.preventDefault() // 拖离
        that.bgStyle = false
      })
      oDragWrap.addEventListener('dragover', (e) => {
        e.preventDefault() // 清除默认事件
      })
      oDragWrap.addEventListener('drop', (e) => {
        that.dropHandler(e) // 抛下
      })
    }
  },
  mounted () {
    that = this
    that.stats()
  },
  filters: {
    NumStorage (value) {
      if (String(value) === '0') return 0
      if (!value) return '-'
      return value.toFixed(10)
    },
    NumStoragePlan (value) {
      if (!value) return '-'
      return value.toFixed(9) > 0 ? value.toFixed(9) : '0.000000001'
    },
    byteStorage (limit) {
      // 只转换成GB
      if (limit <= 0) {
        return '0'
      } else {
        return limit / (1024 * 1024 * 1024) // or 1000
      }
    }
  },
  computed: {
    metaAddress () {
      return this.$store.getters.metaAddress
    },
    networkID () {
      return Number(this.$store.getters.networkID)
    },
    free_usage () {
      return this.$store.getters.free_usage
    },
    free_quota_per_month () {
      return this.$store.getters.free_quota_per_month
    },
    metaNetworkInfo () {
      return this.$store.getters.metaNetworkInfo ? JSON.parse(JSON.stringify(this.$store.getters.metaNetworkInfo)) : {}
    }
  }
}
</script>

<style scoped lang="scss">
.el-dialog__wrapper /deep/ {
  display: flex;
  align-items: center;
  .metaM {
    margin: auto !important;
    .el-dialog__body {
      padding: 0.25rem 0.25rem 0.2rem;
      .el-row {
        border-radius: 0.08rem;
        padding: 0.16rem;
        margin: 0.12rem 0px;
        border: 1px solid rgb(240, 185, 11);
        text-align: center;
        display: flex;
        -webkit-box-pack: justify;
        justify-content: space-between;
        -webkit-box-align: center;
        align-items: center;
        transition: all 0.3s ease 0s;
        background: rgba(240, 185, 11, 0.1);
        position: static;
        .el-col {
          text-align: left;
          font-size: 0.18rem;
          img {
            float: right;
            height: 0.24rem;
          }
        }
      }
    }
    .el-dialog__footer {
      padding: 0 0.25rem 0.3rem;
      .dialog-footer {
        .el-button {
          width: 100%;
          font-size: 0.18rem;
          height: 0.4rem;
          padding: 0;
          background: #5c3cd3;
          color: #fff;
          border-radius: 0.08rem;
          &:hover {
            background: #4326ab;
          }
        }
        p {
          font-size: 0.13rem;
          line-height: 1.5;
          color: red;
          margin: 0.1rem 0 0;
        }
      }
    }
  }
  .completeDia {
    margin: auto !important;
    text-align: center;
    box-shadow: 0 0 13px rgba(128, 128, 128, 0.8);
    border-radius: 0.2rem;
    .el-dialog__header {
      display: none;
    }
    img {
      display: block;
      max-width: 100px;
      margin: 0 auto 0.3rem;
    }
    h1 {
      margin: 0.22rem auto 0.1rem;
      font-size: 0.32rem;
      font-weight: 500;
      line-height: 1.2;
      color: #191919;
      word-break: break-word;
    }
    h2 {
      margin: 0 auto 0.1rem;
      font-size: 0.22rem;
      font-weight: 600;
      line-height: 1.4;
      color: #555;
      word-break: break-word;
      text-align: center;
    }
    h4 {
      font-size: 0.2rem;
      font-weight: 500;
      line-height: 1.4;
      color: #555;
      word-break: break-word;
      text-align: center;
    }
    h3,
    a {
      font-size: 0.22rem;
      font-weight: 500;
      line-height: 1.4;
      color: #191919;
      word-break: break-word;
    }
    a {
      text-decoration: underline;
      color: #007bff;
    }
    a.a-close {
      height: 0.6rem;
      line-height: 0.6rem;
      padding: 0 45px;
      background: linear-gradient(45deg, #4f8aff, #4b5eff);
      font-size: 0.22rem;
      color: #fff;
      border-radius: 0.14rem;
      cursor: pointer;
      margin: 0.4rem auto 0;
      display: block;
      width: max-content;
      text-decoration: unset;
    }
  }
  .fileUpload {
    margin: auto !important;
    box-shadow: 0 0 13px rgba(128, 128, 128, 0.8);
    border-radius: 0.2rem;
    .el-dialog__header {
      padding: 0.3rem 0.4rem;
      color: #000;
      font-size: 0.22rem;
      font-weight: 500;
      line-height: 1;
      text-transform: capitalize;
    }
    .el-dialog__body {
      position: relative;
      padding: 0 0.4rem 0.1rem;
      h3 {
        margin: 0 0 0.1rem;
        font-size: 0.2rem;
        font-weight: normal;
        line-height: 1.4;
        color: #555;
        word-break: break-word;
      }
      .gif_img {
        max-width: 200px;
        width: 100%;
        margin: auto;
        display: block;
      }
      small {
        position: absolute;
        bottom: 0.25rem;
        left: 0;
        right: 0;
        display: block;
        width: 100%;
        font-size: 18px;
        text-align: center;
        @media screen and (max-width: 1600px) {
          font-size: 16px;
        }
        @media screen and (max-width: 1440px) {
          font-size: 15px;
        }
        @media screen and (max-width: 1024px) {
          font-size: 14px;
        }
        @media screen and (max-width: 768px) {
          font-size: 13px;
        }
        @media screen and (max-width: 441px) {
          font-size: 12px;
        }
      }
    }
  }
  .uploadDig {
    background: #fff;
    margin: auto !important;
    box-shadow: 0 0 13px rgba(128, 128, 128, 0.8);
    border-radius: 0.2rem;
    .el-dialog__header {
      padding: 0.3rem 0.4rem;
      display: flex;
      align-items: center;
      border-bottom: 1px solid #dfdfdf;
      color: #000;
      font-size: 0.22rem;
      font-weight: 500;
      line-height: 1;
      text-transform: capitalize;
      @media screen and (max-width: 479px) {
        padding: 0.3rem 0.2rem;
      }
      img {
        width: 20px;
        height: 20px;
        margin: 0 0 0 5px;
        cursor: pointer;
        @media screen and (max-width: 1440px) {
          width: 17px;
          height: 17px;
        }
        @media screen and (max-width: 1280px) {
          width: 16px;
          height: 16px;
        }
      }
      .el-dialog__title {
        color: #000;
        font-size: 0.22rem;
        font-weight: 500;
        line-height: 1;
        text-transform: capitalize;
      }
      .el-dialog__headerbtn {
        display: none;
      }
    }
    .el-dialog__body {
      padding: 0;
      .upload_form {
        // display: flex;
        // align-items: baseline;
        width: calc(100% - 0.8rem);
        padding: 0 0.4rem;
        margin: auto;
        justify-content: flex-start;
        @media screen and (max-width: 1200px) {
          width: calc(100% - 0.6rem);
          padding: 0 0.3rem;
        }
        @media screen and (max-width: 479px) {
          width: calc(100% - 0.4rem);
          padding: 0 0.2rem;
        }
        .upload-demo {
          width: 100%;
          margin: 0.25rem 0 0;
          .el-upload {
            width: 100%;
            height: auto;
            border: 0;
            .el-upload-dragger {
              width: 100%;
              height: auto;
              padding: 0.2rem 0;
              color: #97a8be;
              svg {
                width: 0.45rem;
                height: 0.45rem;
                margin: 0;
              }
              .el-upload__text {
                margin: 0 0 6px;
                font-weight: normal;
                color: #9c9c9c;
                line-height: 1;
              }
              .el-button {
                border-color: rgb(44, 127, 248);
                color: rgb(44, 127, 248);
              }
            }
          }
          .el-upload-list__item:first-child {
            margin-top: 0;
          }
          .el-upload-list {
            width: 100%;
            float: none;
            clear: both;
          }
        }
        .el-form {
          width: 100%;
          margin: 0;
          .el-form-item::after,
          .el-form-item::before {
            display: none;
          }
          .el-form-item {
            display: flex;
            align-items: center;
            justify-content: space-between;
            width: 100%;
            margin: 0.2rem auto;
            .el-form-item__label {
              display: flex;
              justify-content: flex-start;
              align-items: center;
              width: 47%;
              padding: 0 2% 0 0;
              // max-width: 2rem;
              line-height: 1.5;
              text-align: left;
              font-size: 0.2rem;
              white-space: normal;
              color: #000;
              font-weight: 500;
              text-shadow: 0 0 black;
              text-align: right;
              @media screen and (max-width: 479px) {
                padding: 0;
              }
              img {
                width: 20px;
                height: 20px;
                margin: 0 0 0 5px;
                cursor: pointer;
                @media screen and (max-width: 1440px) {
                  width: 17px;
                  height: 17px;
                }
                @media screen and (max-width: 1280px) {
                  width: 16px;
                  height: 16px;
                }
              }
              &::before {
                display: none;
              }
            }
            .el-form-item__content {
              width: 50%;
              display: flex;
              align-items: center;
              justify-content: flex-end;
              font-size: 0.2rem;
              white-space: normal;
              word-break: break-word;
              line-height: 1.5;
              color: #555;
              h4 {
                width: 100%;
                font-size: 0.1372rem;
                font-weight: 500;
                line-height: 1.7;
              }
              h5 {
                width: 90%;
                margin-top: 5px;
                font-size: 0.11rem;
                font-weight: 500;
                line-height: 1.2;
                color: #737373;
              }
              .el-tag,
              .el-button--small {
                margin: 0 5px 5px 0;
              }
              .el-input {
                width: auto;
                .el-input__inner {
                  height: 0.32rem;
                  font-size: 0.1372rem;
                  line-height: 0.32rem;
                }
                .el-input__suffix {
                  display: none;
                }
              }
              .el-form-item__error {
                padding-top: 0;
                margin: 0 0.1rem;
                position: relative;
                float: right;
              }
              .el-textarea {
                width: 90% !important;
              }
              .upload-demo {
                display: flex;
                align-items: center;
                flex-wrap: wrap;
                .el-upload-list__item:first-child {
                  margin-top: 0;
                }
                .el-upload--text {
                  float: left;
                  width: auto;
                  height: auto;
                  text-align: left;
                  border: 0;
                  .el-button--primary {
                    // height: 0.32rem;
                    // padding: 0 0.2rem;
                    // margin: 0 5px 0 0;
                    // line-height: 0.32rem;
                    // background-color:transparent;
                    // border: 1px solid #2c4c9e;
                    // border-radius: 0.08rem;
                    // color: #2c4c9e;
                    // font-size: 0.1372rem;
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    height: 0.5rem;
                    padding: 0 0.2rem;
                    margin: 0;
                    background: linear-gradient(45deg, #4e88ff, #4b5fff);
                    border-radius: 0.14rem;
                    line-height: 0.5rem;
                    text-align: center;
                    color: #fff;
                    font-size: 0.18rem;
                    font-family: inherit;
                    border: 0;
                    outline: none;
                    transition: background-color 0.3s, border-color 0.3s,
                      color 0.3s, box-shadow 0.3s;
                    cursor: pointer;
                    span {
                      display: flex;
                      align-items: center;
                    }
                    img {
                      display: inline-block;
                      height: 20px;
                      margin: 0 0.1rem 0 0;
                      @media screen and (max-width: 1280px) {
                        height: 16px;
                      }
                    }
                    &:hover {
                      opacity: 0.9;
                      box-shadow: 0 12px 12px -12px rgba(12, 22, 44, 0.32);
                    }
                  }
                }
                .el-upload-list {
                  width: 100%;
                  max-width: 300px;
                  float: none;
                  clear: both;
                  @media screen and (max-width: 1440px) {
                    max-width: 250px;
                  }
                }
              }
              .el-upload__tip {
                // float: left;
                height: 100%;
                align-items: center;
                display: flex;
                margin: 0 0 0 0.1rem;
                color: #737373;
                line-height: 1;
                font-size: 0.12rem;
              }
              .el-radio {
                .el-radio__inner {
                  border-color: #d9d9d9;
                  background-color: #d9d9d9;
                }
              }
              .el-radio.is-checked {
                .el-radio__inner {
                  border-color: #0b318f;
                  background-color: #0b318f;
                }
                .el-radio__inner::after {
                  width: 6px;
                  height: 6px;
                }
              }
            }
          }
        }
      }
      .upload_plan {
        width: 100%;
        margin: auto;
        justify-content: flex-start;
        .title {
          display: flex;
          align-items: center;
          justify-content: flex-start;
          margin: 0;
          line-height: 1.5;
          text-align: center;
          font-size: 0.22rem;
          white-space: normal;
          color: #000;
          font-weight: 500;
          // text-shadow: 0 0 black;
          text-indent: 0;
          img {
            width: 20px;
            height: 20px;
            margin: 0 0 0 5px;
            cursor: pointer;
            @media screen and (max-width: 1440px) {
              width: 17px;
              height: 17px;
            }
            @media screen and (max-width: 1280px) {
              width: 16px;
              height: 16px;
            }
          }
        }
        .desc {
          margin: 0 0 0.1rem;
          line-height: 1.5;
          font-size: 0.16rem;
          white-space: normal;
          color: #999;
          font-weight: normal;
        }
        .upload_plan_radio {
          .el-radio-group {
            width: 100%;
            background: #f7f7f7;
            border-radius: 0.2rem;
            .el-radio {
              display: flex;
              align-items: center;
              justify-content: space-between;
              width: 100%;
              height: auto;
              padding: 0.2rem 0.3rem;
              margin: auto;
              border: 0;
              // line-height:30px;
              .el-radio__input {
                width: 20px;
                display: flex;
                align-items: center;
                .el-radio__inner {
                  border-color: #555;
                }
              }
              .el-radio__input.is-checked {
                .el-radio__inner {
                  position: relative;
                  width: 16px;
                  height: 16px;
                  border-color: transparent;
                  background: transparent;
                  &:after {
                    content: "";
                    display: block;
                    height: 16px;
                    width: 16px;
                    background-image: url(../assets/images/icon_xuanzhong@2x.png);
                    background-size: 100%;
                    position: absolute;
                    left: 0;
                    top: 0;
                    transform: translate(0, 0) scale(1);
                    transition: all 0.15s;
                  }
                }
              }
              .el-radio__label {
                display: flex;
                justify-content: space-between;
                width: calc(100% - 30px);
                .title {
                  font-size: 0.2rem;
                  line-height: 1;
                }
                .cont {
                  font-size: 0.2rem;
                  font-weight: 500;
                  line-height: 1;
                  text-align: center;
                }
              }
            }
            .el-radio:nth-child(3n + 1) {
              .el-radio__label {
                .cont {
                  color: #35ad92;
                }
              }
            }
            .el-radio:nth-child(3n + 2) {
              border-top: 1px solid #dfdfdf;
              border-bottom: 1px solid #dfdfdf;
              .el-radio__label {
                .cont {
                  color: #2c7ff8;
                }
              }
            }
            .el-radio:nth-child(3n + 3) {
              .el-radio__label {
                .cont {
                  color: #f63d3d;
                }
              }
            }
            .el-radio:hover {
              background-color: rgba(64, 158, 255, 0.1);
            }
          }
        }
        .upload_plan_radio_free {
          opacity: 0.5;
        }
      }
      .upload_bot {
        width: 100%;
        margin: 0.25rem auto 0.2rem;
        text-align: center;
        .found {
          display: flex;
          justify-content: space-between;
          align-items: center;
          width: 100%;
          text-align: center;
          .el-button {
            height: 0.6rem;
            padding: 0;
            margin-left: 0;
            line-height: 0.6rem;
            font-size: 0.22rem;
            font-family: inherit;
            color: #fff;
            border: 0;
            background: linear-gradient(45deg, #4f8aff, #4b5eff);
            border-radius: 14px;
            width: calc(50% - 0.15rem);
            &:hover {
              opacity: 0.9;
            }
          }
          .cancel {
            background: #dadada;
            transition: background-color 0.3s;
            &:hover {
              background: linear-gradient(45deg, #4f8aff, #4b5eff);
            }
          }
        }
        a {
          margin: auto;
          text-decoration: underline;
          font-size: 0.18rem;
          color: rgb(11, 49, 143);
          cursor: pointer;
        }
      }
    }
    .dialog-footer {
      display: flex;
      align-items: center;
      justify-content: flex-end;
    }
  }
}

#Create {
  position: relative;
  height: calc(100% - 0.6rem);
  padding: 0.4rem 0.2rem 0;
  font-size: 0.24rem;
  .uploadDigID {
    position: absolute;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    z-index: 99;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: rgba(255, 255, 255, 0.9);
    border-radius: 0.2rem;
    font-size: 18px;
    @media screen and (max-width: 1600px) {
      font-size: 16px;
    }
    .elUpload {
      i {
        display: flex;
        justify-content: center;
        font-size: 45px;
        color: #979797;
      }
    }
    &:before {
      display: flex;
      justify-content: center;
      font-size: 55px;
      color: #979797;
      @media screen and (max-width: 1600px) {
        font-size: 50px;
      }
      @media screen and (max-width: 1200px) {
        font-size: 45px;
      }
    }
  }
  .loadMetamaskPay {
    position: absolute;
    left: 0;
    top: 0;
    right: 0;
    bottom: 0;
    z-index: 99;
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: rgba(255, 255, 255, 0.9);
    border-radius: 0.2rem;
    .el-loading-spinner {
      top: 0;
      position: relative;
      margin: 0 0 0.2rem;
    }
    p {
      font-size: 14px;
      color: #555;
    }
  }
  .el-alert /deep/ {
    position: absolute;
    left: 0;
    top: 0;
    .el-alert__content {
      display: flex;
      align-items: center;
    }
  }
  .upload {
    padding: 0 0.17rem 0.4rem;
    margin: 0 auto 0.2rem;
    background-color: #fff;
    border-radius: 0.1rem;
    overflow: hidden;
    .title {
      width: 100%;
      margin: 0 0 0.1rem;
      text-align: left;
      font-size: 0.1972rem;
      color: #000;
      line-height: 0.42rem;
      text-indent: 0.08rem;
    }
    .upload_title {
      width: 100%;
      margin: 0.1rem 0 0.3rem;
      text-align: left;
      font-size: 0.2rem;
      font-weight: 600;
      color: #000;
      line-height: 1.5;
      text-indent: 0;
    }
    .upload_form {
      // display: flex;
      // align-items: baseline;
      width: 80%;
      margin: auto;
      justify-content: flex-start;
      .el-form /deep/ {
        width: 96%;
        margin: 0 2%;
        .el-form-item {
          display: flex;
          align-items: center;
          width: auto;
          margin: 0.15rem auto;
          .el-form-item__label {
            display: flex;
            justify-content: flex-end;
            align-items: center;
            width: 47%;
            padding: 0 2% 0 0;
            // max-width: 2rem;
            line-height: 1.5;
            text-align: left;
            font-size: 0.18rem;
            white-space: normal;
            color: #000;
            font-weight: 500;
            text-shadow: 0 0 black;
            text-align: right;
            img {
              width: 0.16rem;
              height: 0.16rem;
              margin: 0 0 0 5px;
              cursor: pointer;
            }
            &::before {
              display: none;
            }
          }
          .el-form-item__content {
            display: flex;
            align-items: center;
            font-size: 0.18rem;
            white-space: normal;
            word-break: break-word;
            line-height: 1.5;
            color: #666;
            @media screen and (max-width: 600px) {
              font-size: 14px;
            }
            small {
              margin: 2px 5px 0;
              font-size: 12px;
            }
            h4 {
              width: 100%;
              font-size: inherit;
              font-weight: 500;
              line-height: 1.7;
            }
            h5 {
              width: 90%;
              margin-top: 5px;
              font-size: 0.11rem;
              font-weight: 500;
              line-height: 1.2;
              color: #737373;
              @media screen and (max-width: 600px) {
                font-size: 12px;
              }
            }
            .el-tag,
            .el-button--small {
              margin: 0 5px 5px 0;
            }
            .el-input {
              width: auto;
              .el-input__inner {
                height: 0.32rem;
                font-size: inherit;
                line-height: 0.32rem;
              }
              .el-input__suffix {
                display: none;
              }
            }
            .el-form-item__error {
              padding-top: 0;
              margin: 0 0.1rem;
              position: relative;
              float: right;
            }
            .el-textarea {
              width: 90% !important;
            }
            .upload-demo {
              display: flex;
              align-items: center;
              flex-wrap: wrap;
              .el-upload-list__item:first-child {
                margin-top: 0;
              }
              .el-upload--text {
                float: left;
                width: auto;
                height: auto;
                text-align: left;
                border: 0;
                .el-button--primary {
                  height: 0.32rem;
                  padding: 0 0.2rem;
                  margin: 0 5px 0 0;
                  line-height: 0.32rem;
                  background-color: transparent;
                  border: 1px solid #2c4c9e;
                  border-radius: 0.08rem;
                  color: #2c4c9e;
                  font-size: 0.18rem;
                }
              }
              .el-upload-list {
                width: 100%;
                float: none;
                clear: both;
              }
            }
            .el-upload__tip {
              // float: left;
              height: 100%;
              align-items: center;
              display: flex;
              margin: 0 0 0 0.1rem;
              color: #737373;
              line-height: 1;
              font-size: 0.12rem;
              @media screen and (max-width: 600px) {
                font-size: 14px;
              }
            }
            .el-radio {
              .el-radio__inner {
                border-color: #d9d9d9;
                background-color: #d9d9d9;
              }
            }
            .el-radio.is-checked {
              .el-radio__inner {
                border-color: #0b318f;
                background-color: #0b318f;
              }
              .el-radio__inner::after {
                width: 6px;
                height: 6px;
              }
            }
          }
        }
      }
    }
    .upload_plan {
      width: 80%;
      margin: auto;
      justify-content: flex-start;
      .title {
        display: flex;
        align-items: center;
        justify-content: center;
        margin: 0.3rem 0 0;
        line-height: 1.5;
        text-align: center;
        font-size: 0.2rem;
        white-space: normal;
        color: #000;
        font-weight: 500;
        text-shadow: 0 0 black;
        text-indent: 0;
        img {
          width: 0.16rem;
          height: 0.16rem;
          margin: 0 0 0 5px;
          cursor: pointer;
        }
      }
      .desc {
        margin: 0 0 0.1rem;
        line-height: 1.5;
        text-align: center;
        font-size: 0.18rem;
        white-space: normal;
        color: #999;
        font-weight: normal;
      }
      .upload_plan_radio {
        .el-radio-group /deep/ {
          width: 100%;
          display: flex;
          justify-content: center;
          align-items: center;
          .el-radio {
            min-width: 20%;
            height: auto;
            padding: 0 0.1rem 0.15rem;
            // line-height:30px;
            .el-radio__input {
              display: none;
            }
            .el-radio__label {
              .title {
                margin: 0 0 0.05rem;
                font-size: 0.18rem;
              }
              .cont {
                font-size: 0.18rem;
                font-weight: bold;
                line-height: 1.5;
                text-align: center;
              }
            }
          }
          .el-radio:nth-child(3n + 1) {
            .el-radio__label {
              .cont {
                color: #56c4a6;
              }
            }
          }
          .el-radio:nth-child(3n + 2) {
            .el-radio__label {
              .cont {
                color: #4a92d3;
              }
            }
          }
          .el-radio:nth-child(3n + 3) {
            .el-radio__label {
              .cont {
                color: #922b26;
              }
            }
          }
          .is-checked {
            position: relative;
            &:after {
              content: "";
              display: block;
              height: 25px;
              width: 25px;
              background-image: url(../assets/images/plan.png);
              background-size: 100%;
              position: absolute;
              right: 0;
              top: 0;
            }
          }
          .el-radio:hover {
            background-color: rgba(64, 158, 255, 0.1);
          }
        }
      }
      .upload_plan_radio_free {
        opacity: 0.5;
      }
    }
    .upload_bot {
      display: flex;
      justify-content: center;
      align-items: center;
      flex-wrap: wrap;
      width: 100%;
      margin: 0.25rem auto 0.15rem;
      .el-button /deep/ {
        height: 0.45rem;
        padding: 0 0.4rem;
        margin-left: 0;
        background-color: #0b318f;
        line-height: 0.45rem;
        font-size: 0.18rem;
        color: #fff;
        border: 0;
      }
      .no_login {
        background: #f56c6c;
      }
      .found {
        width: 100%;
        text-align: center;
        a {
          text-decoration: underline;
          font-size: 0.18rem;
          color: rgb(11, 49, 143);
          cursor: pointer;
        }
      }
    }
    .upload_result /deep/ {
      width: 75%;
      margin: 0.6rem auto 0.2rem;
      .el-col {
        display: flex;
        align-items: flex-start;
        font-size: 0.12rem;
        color: #222;
        margin: 0.15rem 0 0;
        h5 {
          width: 100%;
          font-size: 0.18rem;
          font-weight: 500;
          line-height: 1.3;
          color: #000;
        }
        label {
          width: 100px;
          margin: 0 0 0 0.2rem;
        }
        p {
          word-break: break-word;
        }
      }
    }
  }

  /deep/ .el-list-enter-active,
  /deep/ .el-list-leave-active {
    transition: none;
  }

  /deep/ .el-list-enter,
  /deep/ .el-list-leave-active {
    opacity: 0;
  }
}
// #Create::-webkit-scrollbar-track {
//     background: #ccc;
// }
// #Create::-webkit-scrollbar {
//     width: 6px;
//     background: #ccc;
// }
// #Create::-webkit-scrollbar-thumb {
//     background: #f2f2f2;
//     border-radius: 0.08rem;
// }
@media screen and (max-width: 1024px) {
  #Create {
    .upload {
      .upload_form {
        .el-form {
          width: 48%;
        }
      }
    }
  }
}
@media screen and (max-width: 999px) {
  #Create {
    padding: 0.1rem;
    .upload {
      padding: 0.1rem;
      .upload_form {
        width: 90%;
        flex-wrap: wrap;
        .el-form /deep/ {
          width: 95%;
        }
      }
    }
  }
}
@media screen and (max-width: 600px) {
  #Create {
    .upload {
      .upload_plan {
        .upload_plan_radio {
          .el-radio-group /deep/ {
            .el-radio {
              width: 32%;
              margin: auto;
            }
          }
        }
      }
    }
  }
}
@media screen and (max-width: 441px) {
  #Create {
    .upload {
      .title {
        text-indent: 0;
        font-size: 0.19rem;
        line-height: 1.5;
      }
      .upload_form {
        width: 100%;
        .el-form /deep/ {
          .el-form-item {
            flex-wrap: wrap;
            .el-form-item__label {
              width: 100%;
              padding-bottom: 0.05rem;
              justify-content: center;
            }
            .el-form-item__content {
              flex-wrap: wrap;
              width: 100%;
              justify-content: center;
            }
          }
        }
      }
      .upload_plan {
        width: 95%;
      }
      .upload_result {
        width: 90%;
        margin: 0.2rem auto 0;
        .el-col {
          flex-wrap: wrap;
          label {
            margin: 0 0 5px;
            width: 100%;
            font-weight: 600;
          }
        }
      }
    }
  }
}
</style>
