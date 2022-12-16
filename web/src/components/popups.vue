<template>
  <div :class="{'slideScroll slideAdd': true, 'slideAddBg': dialogFormVisible&&(typeName !== 'add'), 'slideComSoon': fixed, slideEmail: dialogFormVisible&&(typeName === 'emailLogin'||typeName === 'emailCheck')}" v-if="dialogFormVisible">
    <div class="fe-none" v-if="typeName === 'add' || typeName === 'rename' || typeName === 'addSub'">
      <div class="addBucket" v-loading="createLoad">
        <i class="el-icon-circle-close close" @click="closeDia()"></i>
        <div class="title">{{typeName === 'addSub'?$t('metaSpace.folder_name'):$t('metaSpace.bucket_name')}}</div>
        <el-form :model="form" status-icon :rules="rules" ref="form">
          <el-form-item prop="name">
            <el-input v-model="form.name" maxlength="256" :placeholder="typeName === 'addSub'?'Folder Name':'Bucket Name'" ref="bucketNameRef"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="getDialogClose('form')">{{$t('metaSpace.Submit')}}</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'delete'">
      <div class="addBucket" v-loading="listTableLoad">
        <div class="title">
          <i class="el-icon-warning-outline"></i>
          {{$route.name == 'Space'?$t('metaSpace.delete_title'):$t('metaSpace.delete_title_detail')}}
        </div>
        <div class="cont">{{$t('metaSpace.delete_desc')}}</div>
        <el-form ref="form">
          <el-form-item>
            <el-button type="info" @click="closeDia()">{{$t('metaSpace.Cancel')}}</el-button>
            <el-button type="primary" @click="closeDia('delete')">{{$t('metaSpace.Delete')}}</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'detail'">
      <div class="addBucket" v-loading="backupLoad">
        <div class="head">
          {{$t('metaSpace.detail_title')}}
        </div>
        <el-form ref="form" class="demo-ruleForm">
          <el-form-item :label="$t('metaSpace.detail_BucketName')">
            {{areaBody.BucketName}}
          </el-form-item>
          <el-form-item :label="$t('metaSpace.detail_DateCreated')">
            <span class="color">{{momentFun(areaBody.CreatedAt)}}</span>
          </el-form-item>
          <el-form-item :label="$t('metaSpace.detail_LastModified')">
            <span class="color">{{momentFun(areaBody.UpdatedAt)}}</span>
          </el-form-item>
          <el-form-item :label="$t('metaSpace.detail_CurrentFiles')">
            {{areaBody.FileNumber}}
          </el-form-item>
          <el-form-item :label="$t('metaSpace.detail_CurrentSize')">
            {{areaBody.Size | formatbytes}}
          </el-form-item>
          <el-form-item></el-form-item>
          <el-form-item :label="$t('metaSpace.detail_BackupInfo')">
            <div class="tip">
              {{$t('metaSpace.detail_StorageProvider')}}({{areaBody.miner_count}})

              <el-popover placement="top" popper-class="elPopTitle" width="200" trigger="hover" v-if="areaBody.miner_count">
                <div>
                  <a v-for="(minerFid,s) in areaBody.miner_list" :key="s" :href="`${areaBody.miner_url_prefix}${minerFid}`" target="_blank" style="color: #474747;">
                    {{minerFid}}
                    <span v-if="s<areaBody.miner_list.length-1">,&nbsp;</span>
                  </a>
                </div>
                <img slot="reference" src="@/assets/images/info.png" />
              </el-popover>
            </div>
          </el-form-item>
          <el-form-item :label="$t('metaSpace.detail_PieceCID')">
            <div class="tip">
              {{areaBody.piece_cid}}
              <img slot="reference" src="@/assets/images/space/icon_10.png" class="copy" @click="copyLink(areaBody.piece_cid)" />
            </div>
          </el-form-item>
          <el-form-item :label="$t('metaSpace.detail_RemainingServiceDays')">
            <span class="color">{{areaBody.remaining_service_days}}</span>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="closeDia()">{{$t('metaSpace.Close')}}</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'detail_folder'">
      <div class="addBucket" v-loading="backupLoad">
        <div class="head">
          {{$t('metaSpace.detail_folder_title')}}
        </div>
        <el-form ref="form" class="demo-ruleForm">
          <el-form-item :label="$t('metaSpace.detail_folderName')">
            {{areaBody.Name}}
          </el-form-item>
          <el-form-item :label="$t('metaSpace.detail_DateCreated')">
            <span class="color">{{momentFun(areaBody.CreatedAt)}}</span>
          </el-form-item>
          <el-form-item :label="$t('metaSpace.detail_LastModified')">
            <span class="color">{{momentFun(areaBody.UpdatedAt)}}</span>
          </el-form-item>
          <el-form-item :label="$t('metaSpace.detail_CurrentSize')">
            {{areaBody.Size | formatbytes}}
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="closeDia()">{{$t('metaSpace.Close')}}</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'detail_file'">
      <div class="addBucket" v-loading="backupLoad">
        <i class="el-icon-circle-close closePop" v-if="ipfsUploadLoad" @click="controllerSignal()"></i>
        <div class="head">
          {{$t('metaSpace.ob_detail_title')}}
        </div>
        <el-form ref="form" class="demo-ruleForm">
          <el-form-item :label="$t('metaSpace.ob_detail_ObjectName')">
            <span class="color">{{areaBody.Name}}</span>
          </el-form-item>
          <el-form-item :label="$t('metaSpace.ob_detail_DateUploaded')">
            <span class="color">{{momentFun(areaBody.CreatedAt)}}</span>
          </el-form-item>
          <el-form-item :label="$t('metaSpace.ob_detail_ObjectSize')">
            <span class="color">{{areaBody.Size | formatbytes}}</span>
          </el-form-item>
          <el-form-item :label="$t('metaSpace.ob_detail_ObjectIPFSLink')">
            <a class="color ipfsStyle" @click="xhrequest(areaBody.IpfsUrl, areaBody.Name)">
              {{areaBody.IpfsUrl}}
            </a>
          </el-form-item>
          <el-form-item :label="$t('metaSpace.ob_detail_ObjectCID')">
            <span class="color">{{areaBody.PayloadCid}}</span>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="closeDia()">{{$t('metaSpace.Close')}}</el-button>
          </el-form-item>
        </el-form>
        <div class="loadTryAgain" v-if="ipfsUploadLoad">
          <div style="width:100%;">
            <div class="load_svg">
              <svg viewBox="25 25 50 50" class="circular">
                <circle cx="50" cy="50" r="20" fill="none" class="path"></circle>
              </svg>
              <p>
                {{$t('uploadFile.payment_tip_deal')}}
                <span @click="controllerSignal('try_again', areaBody.IpfsUrl, areaBody.Name)">{{$t('metaSpace.try_again')}}</span>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'upload'">
      <div class="uploadDig" v-loading="loading">
        <i class="el-icon-circle-close close" @click="closeDia()"></i>
        <div class="upload_form">
          <el-upload class="upload-demo" :style="{'border': ruleForm.fileList_tip?'2px dashed #f56c6c':'0'}" drag ref="uploadFileRef" action="customize" :http-request="uploadFile" :file-list="ruleForm.fileList" :on-change="handleChange" :on-remove="handleRemove">
            <img src="@/assets/images/space/icon_11.png" alt="">
            <div class="el-upload__text">{{$t('metaSpace.upload_desc')}}</div>
            <div class="el-upload__text">
              <small>{{$t('metaSpace.uploadDray_or')}}</small>
            </div>
            <el-button type="primary">
              {{$t('metaSpace.uploadDray_text')}}
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
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'upload_folder'">
      <div class="uploadDig" v-loading="loading">
        <i class="el-icon-circle-close close" @click="closeDia()"></i>
        <div class="upload_form">
          <el-upload class="upload-demo" :multiple="true" :style="{'border': ruleForm.fileList_tip?'2px dashed #f56c6c':'0'}" drag ref="uploadFolderRef" action="customize" :http-request="uploadFile" :file-list="ruleForm.fileList" :on-change="handleChange"
            :on-remove="handleRemove">
            <img src="@/assets/images/space/icon_11.png" alt="">
            <div class="el-upload__text">{{$t('metaSpace.upload_desc')}}</div>
            <div class="el-upload__text">
              <small>{{$t('metaSpace.uploadDray_or')}}</small>
            </div>
            <el-button type="primary">
              {{$t('metaSpace.uploadDray_text')}}
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
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'upload_progress'">
      <div class="uploadDig">
        <div class="upload_progress">
          <img src="@/assets/images/space/load_sunny.gif" class="load" />
          <div class="progress">
            <p>{{uploadPrecent}}%
              <small>{{speedChange(uploadPrecentSpeed)}}</small>
            </p>
          </div>
        </div>
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'pay'">
      <div class="addBucket" v-loading="payLoad" :element-loading-text="$t('metaSpace.pay_load')">
        <i class="el-icon-circle-close close" @click="closeDia()"></i>
        <div class="head">
          {{$t('metaSpace.pay_title')}}
        </div>
        <el-row class="pay_body">
          <el-col :span="24">
            <div class="pay_body_top">
              <label>{{$t('metaSpace.pay_body_BucketName')}}</label>
              <el-form :model="form" status-icon :rules="rules" ref="payForm">
                <el-form-item prop="name">
                  <el-input v-model="form.name" maxlength="256" :placeholder="'Bucket Name'" ref="bucketNameRef"></el-input>
                </el-form-item>
              </el-form>
            </div>
          </el-col>
          <el-col :span="24">
            <div class="pay_body_top">
              <label>{{$t('metaSpace.pay_body_Bucket')}}</label>
              <!-- <el-input-number v-model="pay.num" controls-position="right" @change="payhandChange" :min="1"></el-input-number> -->
              <el-input v-model="pay.num" disabled controls-position="right" @change="payhandChange" :min="1"></el-input>
            </div>
            <p>{{$t('metaSpace.pay_body_Bucket_desc')}}</p>
          </el-col>
          <el-col :span="24">
            <div class="pay_body_top">
              <label>{{$t('metaSpace.pay_body_PriceBucket')}}</label>
              {{pay.price}} USDC/Year
            </div>
            <p>{{$t('metaSpace.pay_body_PriceBucket_desc')}}</p>
          </el-col>
          <el-col :span="24">
            <div class="pay_body_top">
              <label>{{$t('metaSpace.pay_body_TotalPayment')}}</label>
              {{pay.total}} USDC
            </div>
          </el-col>
          <el-col :span="24">
            <div class="pay_tip">
              {{$t('metaSpace.pay_tip')}} <br />{{$t('metaSpace.pay_tip1')}}
            </div>
          </el-col>
        </el-row>
        <el-form ref="form">
          <el-form-item>
            <el-button type="primary" @click="getDialogClose('payForm')">{{$t('metaSpace.Pay')}}</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'success'">
      <div class="addBucket success">
        <div class="title">
          {{$t('metaSpace.pay_success')}}
        </div>
        <div class="cont">{{$t('metaSpace.pay_success_desc')}}</div>
        <el-form ref="form">
          <el-form-item>
            <el-button type="primary" @click="closeDia('payClose')">{{$t('metaSpace.Close')}}</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'comingSoon'">
      <div class="addBucket comingSoon">
        <i class="el-icon-circle-close close" @click="closeDia()"></i>
        <div class="soonImg"></div>
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'emailLogin' && false">
      <div class="addBucket loginEmail">
        <div class="titleCont">
          <div class="address">
            <div class="address_left">
              <img src="@/assets/images/metamask.png" class="resno" alt=""> {{ metaAddress | hiddAddress}}
            </div>
            <div class="address_right">
              <div class="flex-shrink-0 w-2 h-2 rounded-full bg-primary"></div>
              <div>{{$t('fs3Login.Connected')}}</div>
            </div>
          </div>
        </div>
        <div v-loading="emailLoad" class="ruleForm">
          <div class="form_title">{{changeTitle?$t('fs3Login.Connect_form_label_change'):$t('fs3Login.Connect_form_label')}}</div>
          <el-form :model="form" status-icon :rules="rulesEmail" ref="form">
            <el-form-item prop="email">
              <el-input v-model="form.email" placeholder="you@domain.com" ref="bucketEmailRef"></el-input>
            </el-form-item>
            <el-form-item prop="checkType" class="type">
              <el-checkbox-group v-model="form.checkType">
                <el-checkbox label="agreement" name="checkType">
                  {{$t('fs3Login.Connect_checkbox')}}
                  <a>{{$t('fs3Login.Connect_checkbox_1')}}</a>{{$t('fs3Login.Connect_checkbox_2')}}
                </el-checkbox>
              </el-checkbox-group>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitEmail('form')">{{changeTitle?$t('fs3Login.Connect_form_btn_change'):$t('fs3Login.Connect_form_btn')}}</el-button>
            </el-form-item>
          </el-form>
          <a href="javascript:;" @click="signInSkip" class="skip">{{$t('fs3Login.Skip')}}</a>
        </div>
      </div>
    </div>
    <div class="fe-none" v-else-if="typeName === 'emailCheck'">
      <div class="addBucket checkEmail">
        <div class="check_email">
          <div class="check_left">
            <h3>{{$t('fs3Login.Connect_email_title')}}
              <i class="el-icon-circle-check"></i>
            </h3>
            <h4>{{$t('fs3Login.Connect_email_desc')}}</h4>
            <el-button type="primary" @click="closeDia()">{{$t('fs3Login.Connect_email_jump')}}</el-button>
          </div>
          <div class="check_right">
            <img src="@/assets/images/login/icon_01.png" class="resno" alt="">
            <u @click="typeName = 'emailLogin'">{{$t('fs3Login.Connect_email_again')}}</u>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import moment from 'moment'
let that
export default {
  data () {
    let validateName = (rule, value, callback) => {
      let regexp = /[/:*?"<>'|\\]/gi
      if (regexp.test(value)) {
        callback(new Error("The folder name cannot contain any of the following characters /:*?\"<>'|\\"))
      }
      if (value.trim() === '') {
        callback(new Error('Folder name cannot be empty'))
      } else {
        callback()
      }
    }
    return {
      form: {
        name: '',
        email: '',
        checkType: ['agreement']
      },
      rules: {
        name: [{ validator: validateName, trigger: 'blur' }]
      },
      rulesEmail: {
        email: [
          { required: true, message: 'Please enter the email address.', trigger: 'blur' },
          { type: 'email', message: 'Please enter the correct email address.', trigger: ['blur', 'change'] }],
        checkType: [
          { type: 'array', required: true, message: ' ', trigger: 'change' }
        ]
      },
      loading: false,
      ruleForm: {
        fileList: [],
        fileList_tip: false,
        fileList_tip_text: '',
        file_size: '',
        file_size_byte: ''
      },
      bodyWidth: document.documentElement.clientWidth < 1024,
      fileListTip: false,
      width: document.body.clientWidth > 600 ? '400px' : '95%',
      widthUpload: document.body.clientWidth > 600 ? '450px' : '95%',
      widthDia: document.body.clientWidth <= 600 ? '95%' : document.body.clientWidth <= 1440 ? '7rem' : '6.6rem',
      pay: {
        num: 1,
        price: 7.2,
        total: 7.2
      },
      ipfsUploadLoad: false,
      typeName: this.typeModule,
      uploadPrecent: 0,
      emailLoad: false,
      controller: new AbortController(),
      lastUploadTime: 0,
      lastUploadSize: 0,
      uploadPrecentSpeed: 0,
      progressEvent: {}
    }
  },
  props: ['dialogFormVisible', 'typeModule', 'areaBody', 'createLoad', 'listTableLoad', 'currentBucket', 'fixed', 'backupLoad', 'changeTitle', 'payLoad'],
  watch: {
    dialogFormVisible: function () {
      let _this = this
      if (_this.dialogFormVisible && (_this.typeName === 'add' || _this.typeName === 'rename')) {
        _this.$nextTick(() => {
          _this.form.name = ''
          _this.$refs.bucketNameRef.$el.querySelector('input').focus()
        })
      }
    },
    typeModule: function () {
      that.typeName = that.typeModule
    }
  },
  methods: {
    controllerSignal (type, link, name) {
      that.controller.abort()
      if (type) that.xhrequest(link, name)
      else that.closeDia()
    },
    downloadBlob (blob, fileName) {
      try {
        const href = window.URL.createObjectURL(blob)
        if (window.navigator.msSaveBlob) {
          window.navigator.msSaveBlob(blob, fileName)
        } else {
          const downloadElement = document.createElement('a')
          downloadElement.href = href
          downloadElement.target = '_blank'
          downloadElement.download = fileName
          document.body.appendChild(downloadElement)
          downloadElement.click()
          document.body.removeChild(downloadElement)
          window.URL.revokeObjectURL(href)
        }
      } catch (e) {
        console.log('ipfs upload err:', e)
      }
    },
    async xhrequest (link, name) {
      that.ipfsUploadLoad = true
      that.controller = new AbortController()
      let data = await fetch(link, { signal: that.controller.signal })
        .then((response) => response.blob())
        .then((res) => {
          console.log(res)
          let blod = new Blob([res])
          if (res.type.indexOf('javascript') > -1 || res.type.indexOf('image') > -1 || res.type.indexOf('text') > -1 || res.type.indexOf('json') > -1) that.downloadBlob(blod, name)
          else window.open(link)
          that.ipfsUploadLoad = false
        })
        .catch((e) => {
          console.log('Download error: ' + e.message)
          that.$message.error(e.message)
        })
      return data
    },
    payhandChange (value) {
      // console.log(value);
      that.pay.total = value * that.pay.price
    },
    closeDia (type) {
      that.$emit('getPopUps', false, type || '')
    },
    getDialogClose (formName, type) {
      that.$refs[formName].validate(async valid => {
        if (valid) {
          that.$emit('getPopUps', false, that.typeName, that.form.name)
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    momentFun (dateItem) {
      let dateNew = new Date(dateItem).getTime()
      let dataUnit = ''
      let dataTime = new Date(dateNew) + ''
      let dataUnitIndex = dataTime.indexOf('GMT')
      let dataUnitArray = dataTime.substring(dataUnitIndex, dataUnitIndex + 8)
      switch (dataUnitArray) {
        case 'GMT+1000':
          dataUnit = 'GMT+10'
          break
        case 'GMT-1000':
          dataUnit = 'GMT-10'
          break
        case 'GMT+0000':
          dataUnit = 'GMT+0'
          break
        default:
          dataUnit = dataUnitArray ? dataUnitArray.replace(/0/g, '') : '-'
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
    // 文件上传
    uploadFile (params) {
      that._file = params.file
      const isLt2M = that._file.size / 1024 / 1024 / 1024 <= 100 // or 1000
      that.ruleForm.file_size = that.sizeChange(that._file.size)
      that.ruleForm.file_size_byte = that.byteChange(that._file.size)
      if (!isLt2M) {
        that.ruleForm.fileList_tip = true
        that.ruleForm.fileList_tip_text = that.$t('deal.upload_form_file_tip')
        return false
      } else {
        that.ruleForm.fileList_tip = false
      }
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
    speedChange (bytes) {
      if (String(bytes) === '0') return '0 byte/s'
      if (!bytes) return '-'
      var k = 1024 // or 1000
      var sizes = ['byte/s', 'kb/s', 'mb/s', 'gb/s']
      var i = Math.floor(Math.log(bytes) / Math.log(k))
      if (Math.round((bytes / Math.pow(k, i))).toString().length > 3) i += 1
      return (bytes / Math.pow(k, i)).toFixed(1) + ' ' + sizes[i]
    },
    byteChange (limit) {
      var size = ''
      // 只转换成GB
      if (limit <= 0) return '-'
      else size = limit / (1024 * 1024 * 1024) // or 1000
      return size
      // return Number(size).toFixed(3);
    },
    async handleChange (file, fileList) {
      console.log(file)
      let regexp = /[#\\?]/
      let reg = new RegExp(' ', 'g')
      if (file.size <= 0) {
        that.$message.error('Error: Upload file size cannot be 0')
        that.ruleForm.fileList = []
        return false
      } else if (regexp.test(file.name)) {
        that.$message.error('The filename cannot contain any of the following characters # ? \\')
        that.ruleForm.fileList = []
        return false
      } else if (file.name.indexOf(' ') > -1) {
        file.name = file.name.replace(reg, '_')
        file.raw = new File([file.raw], file.name)
      }

      if (fileList.length > 0) {
        that.ruleForm.fileList = [fileList[fileList.length - 1]]

        that.loading = true
        try {
          let alreadyUploadChunks = []
          let { hash } = await that.fileMd5(file)
          // console.log(hash, suffix)
          let fileCheckRes = await that.fileCheck(hash, file)
          if (!fileCheckRes) {
            that.$emit('getUploadDialog', false, 0)
            that.loading = false
            return false
          }

          let max = 0.5 * 1024 * 1024
          let count = Math.ceil(file.size / max)
          let index = 0
          let chunks = []
          let concurrent = 3 // 并发数
          while (index < count) {
            chunks.push({
              file: file.raw.slice(index * max, (index + 1) * max),
              filename: `${index + 1}_${file.name}`
            })
            index++
          }

          that.concurrentExecution(chunks, concurrent, (chunk) => {
            return new Promise(async (resolve, reject) => {
              let inde = chunk + 1
              if (alreadyUploadChunks.indexOf(inde + '') === -1) {
                let uploadData = new FormData()
                const fileBlob = new Blob([chunk.file], {
                  type: 'application/json'
                })
                uploadData.append('file', fileBlob, chunk.filename)
                uploadData.append('hash', hash)
                const config = {
                  onUploadProgress: progressEvent => {
                    that.progressEvent = progressEvent
                    // that.progressHandle(progressEvent)
                  }
                }
                let data = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/oss_file/upload`, 'post', uploadData, config)
                // console.log('data', data.data)
                if (!data) {
                  that.$emit('getUploadDialog', false, 0)
                  that.loading = false
                  return false
                }
                alreadyUploadChunks = data.data
                if (alreadyUploadChunks.length > 0 && alreadyUploadChunks.includes(chunk.filename)) {
                  that.manageProgress(hash, file, data.data.length, count)
                  resolve(data.data)
                  return
                }
                const uploadRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/oss_file/upload`, 'post', uploadData, config)
                if (!uploadRes || uploadRes.status !== 'success') reject(uploadRes ? uploadRes.message : 'Fail')
              } else {
                resolve()
              }
            })
          }).then(res => {
            console.log('finish', res)
          })
        } catch (e) {
          console.log(e)
          await that.$commonFun.timeout(500)
          that.$emit('getUploadDialog', false, 0)
          that.loading = false
        }
      }
    },
    async concurrentExecution (list, limit, asyncHandle) {
      let recursion = (arr) => {
        return asyncHandle(arr.shift()).then(() => {
          if (arr.length !== 0) {
            return recursion(arr)
          } else {
            return 'finish'
          }
        })
      }
      let listCopy = [].concat(list)
      let asyncList = []
      limit = limit > listCopy.length ? listCopy.length : limit
      while (limit--) {
        asyncList.push(recursion(listCopy))
      }
      return Promise.all(asyncList)
    },
    async fileCheck (hash, file) {
      const current = that.currentBucket.split('/').slice(1).join('/')
      let paramCheck = {
        'file_hash': hash,
        'file_name': file.name,
        'prefix': that.areaBody.Prefix || current || '',
        'bucket_uid': that.areaBody
      }
      let checkRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/oss_file/check`, 'post', paramCheck)
      if (!checkRes || checkRes.status !== 'success') {
        that.$message.error(checkRes ? checkRes.message : 'Fail')
        return false
      }
      if (checkRes.data.file_is_exist) {
        that.$message.error('You have uploaded this file!')
        return false
      }
      if (checkRes.data.ipfs_is_exist) {
        await that.$commonFun.timeout(500)
        that.$emit('getUploadDialog', false, 1)
        that.loading = false
        return false
      }
      return true
    },
    async fileMerge (hash, file) {
      const current = that.currentBucket.split('/').slice(1).join('/')
      let paramMerge = {
        'file_hash': hash,
        'file_name': file.name,
        'prefix': that.areaBody.Prefix || current || '',
        'bucket_uid': that.areaBody
      }
      let mergeRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/oss_file/merge`, 'post', paramMerge)
      if (!mergeRes || mergeRes.status !== 'success') {
        that.$message.error(mergeRes ? mergeRes.message : 'Fail')
      }
      await that.$commonFun.timeout(500)
      that.$emit('getUploadDialog', false, 1)
      that.loading = false
    },
    async manageProgress (hash, file, index, count, progressEvent) {
      // console.log(index, count)
      // console.log(that.progressEvent)
      // that.typeName = 'upload_progress'
      that.uploadPrecent = `${index / count * 100}`
      if (index < count) return
      that.uploadPrecent = `100`
      await that.fileMerge(hash, file)
    },
    async fileMd5 (file) {
      return new Promise(resolve => {
        let fileReader = new FileReader()
        fileReader.readAsArrayBuffer(file.raw)
        fileReader.onload = ev => {
          let buffer = ev.target.result
          let spark = new that.$SparkMD5.ArrayBuffer()
          let hash
          let suffix
          spark.append(buffer)
          hash = spark.end()
          let reg = /\.([a-zA-Z0-9]+)$/
          suffix = reg.exec(file.name) ? reg.exec(file.name)[1] : ''
          resolve({
            buffer,
            hash,
            suffix,
            filename: `${hash}.${suffix}`
          })
        }
      })
    },
    handleRemove (file, fileList) {
      // console.log(file, fileList);
      that.ruleForm.fileList = []
      that.ruleForm.file_size = ''
      that.ruleForm.file_size_byte = ''
    },
    onPross (e) {
      const { loaded, total } = e
      const uploadPrecent = ((loaded / total) * 100) | 0
      that.uploadPrecent = loaded < total && uploadPrecent === 100 ? 99 : uploadPrecent
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
      // const uploadPrecent = ((loaded / total) * 100) | 0 // Calculation progress
      // that.uploadPrecent = loaded < total && uploadPrecent === 100 ? 99 : uploadPrecent

      console.log('loaded: ' + loaded, 'total: ', total, 'speed: ', that.uploadPrecentSpeed)
    },
    submitEmail (formName) {
      that.$refs[formName].validate(async valid => {
        if (valid) {
          that.emailLoad = true
          try {
            const params = {
              'email': that.form.email
            }
            const emailRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v3/email`, 'post', params)
            if (!emailRes || emailRes.status !== 'success') {
              that.$message({
                showClose: true,
                message: emailRes ? emailRes.message : 'Fail',
                type: 'error',
                duration: 10000
              })
            } else {
              that.$message({
                showClose: true,
                message: emailRes ? emailRes.data : 'Success',
                type: 'success',
                duration: 10000
              })
              that.typeName = 'emailCheck'
              await that.$metaLogin.emailSign()
            }
          } catch (e) {
            console.log(e)
          }
          that.emailLoad = false
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    async signInSkip () {
      const data = JSON.parse(that.$store.getters.mcsEmail)
      data.apiStatus = false
      that.$store.dispatch('setMCSEmail', JSON.stringify(data))
      that.closeDia()
    }
  },
  mounted () {
    that = this
    console.log(that.typeName)
    if (that.typeName === 'upload_folder') that.$refs.uploadFolderRef.$children[0].$refs.input.webkitdirectory = true
    document.onkeydown = function (e) {
      if (e.keyCode === 13) {
        if (that.typeName === 'add' || that.typeName === 'rename' || that.typeName === 'addSub') that.getDialogClose('form')
        // if (that.typeName === 'emailLogin') that.submitEmail('form')
      }
    }
  },
  computed: {
    metaAddress () {
      return this.$store.getters.metaAddress
    }
  },
  filters: {
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
    hiddAddress: function (val) {
      if (val) {
        return `${val.substring(0, 6)}...${val.substring(val.length - 4)}`
      } else {
        return '-'
      }
    }
  }
}
</script>

<style lang="scss" scoped>
@font-face {
  font-family: "gilroy-regular";
  src: url(../assets/font/gilroy-regular-3.otf);
  font-style: normal;
  font-display: block;
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

      .el-button /deep/ {
        height: 0.34rem;
        padding: 0 0.4rem;
        margin: 0 0.1rem;
        color: #fff;
        line-height: 0.34rem;
        font-size: 0.1372rem;
        border: 0;
        border-radius: 0.08rem;
      }

      .el-input /deep/ {
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
          width: 2rem;
          padding: 0.13rem 0;
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
          img {
            display: inline-block;
            height: 0.25rem;
            margin: 0 0.1rem 0 0;
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
        .el-button /deep/ {
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
      .el-input /deep/ {
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
  .fe-none {
    position: relative;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
    padding: 0 5%;
    .p {
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
    .addBucket {
      position: relative;
      max-width: 750px;
      padding: 0.35rem 0.5rem;
      background-color: #fff;
      border-radius: 0.2rem;
      text-align: left;
      @media screen and (max-width: 441px) {
        max-width: 300px;
        padding: 15px;
      }
      .closePop {
        position: absolute;
        right: -0.1rem;
        top: -0.1rem;
        width: 0.37rem;
        height: 0.37rem;
        background: #fff;
        border-radius: 100%;
        font-size: 0.37rem;
        color: #000;
        cursor: pointer;
        z-index: 2001;
        @media screen and (max-width: 768px) {
          right: -15px;
          top: -15px;
          width: 30px;
          height: 30px;
          font-size: 30px;
        }
      }
      .title {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 0.15rem;
        text-align: center;
        font-size: 0.26rem;
        color: #000;
        line-height: 1;
        i {
          display: block;
          width: 0.37rem;
          height: 0.37rem;
          margin: 0 0.2rem 0 0;
          font-size: 0.37rem;
          color: #f40000;
        }
      }
      .cont,
      .pay_tip {
        padding: 0 20%;
        margin: 0.2rem auto;
        font-size: 0.17rem;
        text-align: center;
        word-break: break-word;
        line-height: 1.5;
        @media screen and (max-width: 1440px) {
          padding: 0 15%;
        }
        @media screen and (max-width: 768px) {
          padding: 0 5%;
          font-size: 13px;
        }
      }
      .pay_tip {
        padding: 0;
        margin: 0;
        font-size: 0.14rem;
        text-align: left;
        @media screen and (min-width: 1800px) {
          font-size: 15px;
        }
        @media screen and (max-width: 1440px) {
          font-size: 12px;
        }
        @media screen and (max-width: 441px) {
          margin: 0;
        }
      }
      .pay_body /deep/ {
        max-width: 750px;
        padding: 0.15rem 0.2rem;
        margin: 0.3rem auto 0;
        background-color: #f7f7f7;
        border-radius: 0.1rem;
        @media screen and (max-width: 441px) {
          padding: 0 0.2rem;
        }
        .el-col {
          padding: 0.15rem 0;
          font-size: 0.17rem;
          line-height: 1.5;
          color: #000;
          @media screen and (max-width: 768px) {
            font-size: 13px;
          }
          .el-form {
            padding: 0;
            .el-form-item {
              margin: 0;
              .el-form-item__content {
                .el-input {
                  margin: 0;
                  .el-input__inner {
                    font-size: 0.17rem;
                    border: 1px solid #999;
                    @media screen and (max-width: 768px) {
                      font-size: 13px;
                    }
                  }
                }
                .el-form-item__error {
                  top: 90%;
                }
              }
            }
          }
          .pay_body_top {
            display: flex;
            align-items: center;
            color: #4e81ff;
            label {
              padding-right: 0.1rem;
              width: 120px;
              text-align: left;
              color: #000;
              font-weight: 600;
              text-transform: capitalize;
            }
            .el-input {
              .el-input__inner {
                height: auto;
                background-color: transparent;
                border: 0;
                line-height: 1;
                color: #333;
              }
            }
            .el-input-number {
              width: auto;
              line-height: inherit;
              .el-input {
                .el-input__inner {
                  width: auto;
                  max-width: 60px;
                  height: auto;
                  padding-left: 0;
                  padding-right: 10px;
                  background-color: transparent;
                  text-align: left;
                  border: 0;
                  line-height: inherit;
                  color: inherit;
                  font-family: inherit;
                  font-size: 0.17rem;
                  @media screen and (max-width: 768px) {
                    font-size: 13px;
                  }
                }
              }
              .el-input-number__increase,
              .el-input-number__decrease {
                display: flex;
                justify-content: center;
                align-items: center;
                width: 8px;
                height: 8px;
                background-color: transparent;
                border: 0;
                line-height: 1;
                color: #999999;
                font-size: 17px;
                .el-icon-arrow-up {
                  &::before {
                    content: "\E78F";
                  }
                }
                .el-icon-arrow-down {
                  &::before {
                    content: "\E790";
                  }
                }
                &:hover {
                  color: #4e81ff;
                }
              }
            }
          }
          p {
            font-size: 0.14rem;
            text-align: left;
            color: #999999;
            padding: 0;
            border: 0;
            @media screen and (min-width: 1800px) {
              font-size: 14px;
            }
            @media screen and (max-width: 1440px) {
              font-size: 12px;
            }
          }
        }
      }
      .el-form /deep/ {
        min-width: 5rem;
        padding: 0.3rem 0 0;
        @media screen and (max-width: 1440px) {
          min-width: 4.5rem;
        }
        @media screen and (max-width: 1200px) {
          min-width: 4rem;
        }
        @media screen and (max-width: 992px) {
          min-width: 3.5rem;
        }
        @media screen and (max-width: 768px) {
          min-width: 3rem;
        }
        .el-form-item {
          margin: 0;
          .el-form-item__content {
            position: relative;
            display: flex;
            align-items: center;
            flex-wrap: wrap;
            text-align: center;
            .el-input {
              margin: 0 auto 0.5rem;
              .el-input__inner {
                height: auto;
                padding: 0.05rem 0.15rem;
                background-color: #f7f7f7;
                border: 0;
                border-radius: 0.1rem;
                font-size: 16px;
                @media screen and (max-width: 1440px) {
                  font-size: 14px;
                }
              }
            }
            .el-form-item__error {
              position: absolute;
              top: 0.65rem;
              text-align: left;
            }
            .el-button {
              width: 45%;
              max-width: 250px;
              min-width: 150px;
              padding: 0.17rem;
              margin: 0 auto;
              font-family: inherit;
              font-size: 16px;
              border: 0;
              border-radius: 0.1rem;
              line-height: 1;
              @media screen and (max-width: 1600px) {
                font-size: 14px;
              }
              @media screen and (max-width: 441px) {
                min-width: auto;
              }
            }
            .el-button--primary {
              background: linear-gradient(45deg, #4e88ff, #4b5fff);
            }
            .el-button--info {
              background: #dadada;
            }
          }
        }
      }
      .demo-ruleForm /deep/ {
        min-width: 7.5rem;
        padding: 0.1rem 0 0;
        @media screen and (max-width: 1440px) {
          min-width: 6.5rem;
        }
        @media screen and (max-width: 1200px) {
          min-width: 6rem;
        }
        @media screen and (max-width: 992px) {
          min-width: 5.5rem;
        }
        @media screen and (max-width: 768px) {
          min-width: 5rem;
        }
        .el-form-item {
          .el-form-item__content,
          label {
            justify-content: flex-start;
            padding: 0.1rem 0;
            font-size: 0.17rem;
            color: #000002;
            line-height: 1.5;
            text-align: left;
            word-break: break-word;
            @media screen and (max-width: 768px) {
              font-size: 14px;
            }
            @media screen and (max-width: 441px) {
              font-size: 13px;
              padding: 2px 0.1rem 2px 0;
            }
            .color {
              color: #4e81ff;
              word-break: break-word;
            }
            .ipfsStyle {
              cursor: pointer;
              word-break: break-all;
              &:hover {
                text-decoration: underline;
              }
            }
            .tip {
              display: flex;
              align-items: center;
              justify-content: flex-start;
              margin: 0;
              font-size: inherit;
              @media screen and (max-width: 441px) {
                width: 100%;
              }
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
              }
              .copy {
                width: 15px !important;
                height: 15px !important;
                @media screen and (min-width: 1800px) {
                  width: 18px !important;
                  height: 18px !important;
                }
              }
            }
          }
          .el-form-item__label {
            padding-right: 0.1rem;
            min-width: 120px;
            text-align: left;
          }
          &:last-child {
            .el-form-item__content {
              justify-content: flex-end;
              .el-button--primary {
                margin: 0;
              }
            }
          }
        }
      }
      .close {
        position: absolute;
        right: 0.2rem;
        top: 0.2rem;
        width: 0.37rem;
        height: 0.37rem;
        font-size: 0.37rem;
        color: #000;
        cursor: pointer;
        @media screen and (max-width: 768px) {
          width: 25px;
          height: 25px;
          font-size: 25px;
        }
      }
      .head {
        margin-top: -0.1rem;
        padding: 0 0 0.2rem;
        font-size: 0.2rem;
        border-bottom: 1px solid #e3e3e3;
        @media screen and (max-width: 768px) {
          font-size: 16px;
        }
        @media screen and (max-width: 441px) {
          font-size: 14px;
          line-height: 1.2;
        }
      }
      .titleCont {
        padding: 0.2rem 0;
        font-weight: 600;
        font-size: 0.286rem;
        font-weight: 600;
        color: #fff;
        line-height: 1.3;
        border-bottom: 1px solid #dde0e2;
        @media screen and (max-width: 600px) {
          font-size: 16px;
        }
        p {
          text-align: center;
        }
        .address {
          display: flex;
          align-items: center;
          justify-content: space-between;
          font-size: 0.223rem;
          font-weight: normal;
          color: #fff;
          line-height: 1.2;
          @media screen and (max-width: 600px) {
            font-size: 13px;
          }
          .address_left {
            display: flex;
            align-items: center;
            img {
              height: 30px;
              margin: 0 0.25rem 0 0;
              @media screen and (max-width: 600px) {
                height: 24px;
              }
            }
          }
          .address_right {
            position: relative;
            display: inline-block;
            padding: 0.1rem 0.2rem 0.1rem 0.32rem;
            background-color: rgba(85, 128, 233, 0.3);
            font-size: 0.148rem;
            border-radius: 0.5rem;
            @media screen and (max-width: 1600px) {
              font-size: 14px;
            }
            @media screen and (max-width: 1440px) {
              font-size: 13px;
            }
            &::before {
              position: absolute;
              left: 0.16rem;
              top: 50%;
              content: "";
              width: 0.08rem;
              height: 0.08rem;
              margin-top: -0.04rem;
              background-color: #4d73ff;
              border-radius: 0.5rem;
            }
          }
        }
      }
      .ruleForm /deep/ {
        padding: 0.3rem 0;
        .form_title {
          font-size: 0.223rem;
          font-weight: normal;
          color: #fff;
          line-height: 1.2;
          @media screen and (max-width: 600px) {
            font-size: 13px;
          }
        }
        .el-form {
          padding: 0.1rem 0 0;
          .el-form-item {
            display: block;
            height: auto;
            max-width: 4.5rem;
            text-align: left;
            .el-form-item__content {
              height: auto;
              width: 100%;
              margin: 0 0 0.1rem;
              font-size: 0.223rem;
              font-weight: normal;
              color: #fff;
              line-height: 1.2;
              @media screen and (max-width: 600px) {
                font-size: 13px;
              }
              .el-checkbox-group {
                .el-checkbox {
                  display: flex;
                  align-items: center;
                  .el-checkbox__label {
                    word-break: break-word;
                    white-space: normal;
                    text-align: left;
                    color: #fff;
                    a {
                      color: inherit;
                      text-decoration: underline;
                    }
                  }
                }
              }
              .el-input {
                margin: 0 auto;
                .el-input__inner {
                  height: auto;
                  padding: 0.2rem 0.15rem;
                  border-radius: 0.1rem;
                  border-color: #3d6ddb;
                  color: #041417;
                  line-height: 1.2;
                  font-size: inherit;
                }
              }
              .el-form-item__error {
                display: flex;
              }
              .el-button {
                display: block;
                width: 80%;
                max-width: 4rem;
                font-family: inherit;
                font-size: 0.25rem;
                font-weight: 600;
                height: auto;
                padding: 0.2rem;
                margin: 0;
                background: linear-gradient(45deg, #0353fe, #7cebfe);
                color: #fff;
                border: 0;
                border-radius: 0.5rem;
                line-height: 1;
                letter-spacing: 1px;
                @media screen and (max-width: 600px) {
                  font-size: 14px;
                }
                span {
                  display: flex;
                  align-items: center;
                  justify-content: center;
                }
                &:hover {
                  background: linear-gradient(45deg, #021cad, #3e70fe);
                }
              }
            }
            &.is-error {
              .el-form-item__content {
                margin: 0 0 0.3rem;
                .el-checkbox-group {
                  .el-checkbox {
                    .el-checkbox__label {
                      color: #f56c6c;
                    }
                  }
                }
              }
            }
            &.type {
              .el-form-item__content {
                margin: 0 0 0.1rem !important;
              }
            }
          }
        }
        .skip {
          float: right;
          color: #fff;
          text-decoration: underline;
          font-size: 14px;
          @media screen and (max-width: 1600px) {
            font-size: 12px;
          }
        }
        .el-loading-mask {
          background-color: rgba(46, 77, 91, 0.75);
        }
      }
      .loadTryAgain {
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        display: flex;
        justify-content: center;
        align-items: center;
        background: rgba(255, 255, 255, 0.8);
        .load_svg {
          display: flex;
          justify-content: center;
          align-items: center;
          flex-wrap: wrap;
          width: 100%;
          svg {
            height: 42px;
            width: 42px;
            -webkit-animation: loading-rotate 2s linear infinite;
            animation: loading-rotate 2s linear infinite;
            .path {
              -webkit-animation: loading-dash 1.5s ease-in-out infinite;
              animation: loading-dash 1.5s ease-in-out infinite;
              stroke-dasharray: 90, 150;
              stroke-dashoffset: 0;
              stroke-width: 2;
              stroke: #409eff;
              stroke-linecap: round;
            }
          }
          p {
            width: 100%;
            text-align: center;
            font-size: 16px;
            color: #333;
            @media screen and (max-width: 1600px) {
              font-size: 14px;
            }
            @media screen and (max-width: 768px) {
              font-size: 13px;
            }
            @media screen and (max-width: 441px) {
              font-size: 12px;
            }
            span {
              color: #4b83fb;
              cursor: pointer;
            }
          }
        }
      }
    }
    .success {
      .title {
        color: #4e7fff;
      }
      .cont,
      .pay_tip {
        font-size: 0.22rem;
        margin: 0.2rem auto 0.1rem;
        @media screen and (max-width: 1440px) {
          font-size: 16px;
        }
        @media screen and (max-width: 768px) {
          font-size: 14px;
        }
      }
    }
    .comingSoon {
      padding: 0;
      background-color: transparent;
      .soonImg {
        display: block;
        width: 450px;
        height: 300px;
        background: url(../assets/images/space/coming-soon.jpg) no-repeat center;
        background-size: cover;
        border-radius: 15px;
        @media screen and (max-width: 768px) {
          width: 350px;
          height: 235px;
        }
        @media screen and (max-width: 441px) {
          width: 300px;
          height: 200px;
        }
      }
      .close {
        position: absolute;
        right: -0.1rem;
        top: -0.1rem;
        width: 0.37rem;
        height: 0.37rem;
        background: #fff;
        border-radius: 100%;
        font-size: 0.37rem;
        color: #000;
        cursor: pointer;
        z-index: 3;
        @media screen and (max-width: 768px) {
          right: -15px;
          top: -15px;
          width: 30px;
          height: 30px;
          font-size: 30px;
        }
      }
    }
    .uploadDig {
      position: relative;
      background: #fff;
      margin: auto !important;
      box-shadow: 0 0 13px rgba(128, 128, 128, 0.8);
      border-radius: 0.2rem;
      width: 6.6rem;
      color: #606266;
      font-size: 14px;
      word-break: break-all;
      @media screen and (max-width: 1440px) {
        width: 7rem;
      }
      @media screen and (max-width: 600px) {
        width: 95%;
      }
      .close {
        position: absolute;
        right: -0.1rem;
        top: -0.1rem;
        width: 0.37rem;
        height: 0.37rem;
        background: #fff;
        border-radius: 100%;
        font-size: 0.37rem;
        color: #000;
        cursor: pointer;
        z-index: 3;
        @media screen and (max-width: 768px) {
          right: -15px;
          top: -15px;
          width: 30px;
          height: 30px;
          font-size: 30px;
        }
      }
      .upload_form {
        // display: flex;
        // align-items: baseline;
        width: calc(100% - 0.5rem);
        padding: 0.25rem;
        margin: auto;
        justify-content: flex-start;
        .upload-demo /deep/ {
          width: 100%;
          margin: 0;
          .el-upload {
            width: 100%;
            height: auto;
            border: 0;
            .el-upload-dragger {
              width: 100%;
              height: auto;
              padding: 0.3rem 0;
              color: #97a8be;
              border-color: #898989;
              svg,
              img {
                width: 0.45rem;
                height: 0.45rem;
                margin: 0;
              }
              .el-upload__text {
                margin: 10px 0;
                font-size: 0.22rem;
                font-weight: normal;
                color: #606060;
                line-height: 1;
                small {
                  color: #969696;
                }
              }
              .el-button {
                width: 45%;
                max-width: 250px;
                min-width: 150px;
                padding: 0.17rem;
                margin: 0 auto;
                background: linear-gradient(45deg, #4e88ff, #4b5fff);
                font-family: inherit;
                font-size: 16px;
                border: 0;
                border-radius: 0.14rem;
                line-height: 1;
                @media screen and (max-width: 1600px) {
                  font-size: 14px;
                }
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
      .upload_progress {
        position: relative;
        width: 100%;
        min-height: 150px;
        padding: 0;
        margin: auto;
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        align-items: center;
        .load {
          display: block;
          width: 100%;
          max-width: 500px;
          border-radius: 0.5rem;
          @media screen and (max-width: 441px) {
            max-width: 300px;
          }
        }
        .moon {
          display: none;
        }
        .sunny {
          display: block;
        }
        .progress {
          position: absolute;
          left: 0;
          top: 0;
          right: 0;
          bottom: 0;
          z-index: 2;
          display: flex;
          justify-content: center;
          align-items: center;
          flex-wrap: wrap;
          font-size: 30px;
          color: #4b83fb;
          line-height: 1.2;
          @media screen and (max-width: 768px) {
            font-size: 28px;
          }
          @media screen and (max-width: 600px) {
            font-size: 23px;
          }
          @media screen and (max-width: 441px) {
            font-size: 18px;
          }
          small {
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
        .load_svg {
          display: flex;
          justify-content: center;
          align-items: center;
          flex-wrap: wrap;
          width: 100%;
          svg {
            height: 42px;
            width: 42px;
            -webkit-animation: loading-rotate 2s linear infinite;
            animation: loading-rotate 2s linear infinite;
            .path {
              -webkit-animation: loading-dash 1.5s ease-in-out infinite;
              animation: loading-dash 1.5s ease-in-out infinite;
              stroke-dasharray: 90, 150;
              stroke-dashoffset: 0;
              stroke-width: 2;
              stroke: #409eff;
              stroke-linecap: round;
            }
          }
          p {
            width: 100%;
            text-align: center;
          }
        }
        .el-progress {
          width: 90%;
          margin: 0.3rem auto 0;
        }
      }
    }
    .loginEmail {
      max-width: 10rem;
      padding: 0.2rem 0.5rem 0.35rem;
      margin: auto;
      background: rgba(4, 20, 23, 1);
      border: 1px solid #4e7cff;
      border-radius: 0.2rem;
      font-family: "gilroy-regular";
    }
    .checkEmail {
      max-width: 5.5rem;
      padding: 0.5rem;
      background: rgba(4, 20, 23, 1);
      border: 1px solid #4e7cff;
      border-radius: 0.2rem;
      font-family: "gilroy-regular";
      @media screen and (min-width: 768px) {
        min-width: 400px;
      }
      .check_email {
        display: flex;
        justify-content: center;
        align-items: center;
        @media screen and (max-width: 441px) {
          flex-wrap: wrap;
        }
        .check_left {
          padding-right: 0.1rem;
          color: #fff;
          text-align: left;
          h3 {
            font-size: 0.298rem;
            line-height: 1.5;
            @media screen and (max-width: 1600px) {
              font-size: 22px;
            }
            @media screen and (max-width: 600px) {
              font-size: 16px;
            }
          }
          h4,
          u,
          a {
            padding: 0.1rem 0;
            font-size: 0.18rem;
            font-weight: normal;
            line-height: 1.5;
            opacity: 0.9;
            color: inherit;
            @media screen and (max-width: 1600px) {
              font-size: 15px;
            }
            @media screen and (max-width: 1440px) {
              font-size: 14px;
            }
            @media screen and (max-width: 600px) {
              font-size: 13px;
            }
            u {
              opacity: 1;
            }
          }
          u {
            cursor: pointer;
            display: block;
            padding: 0;
          }
          .el-button--primary {
            display: flex;
            align-items: center;
            justify-content: center;
            height: 0.5rem;
            padding: 0 0.2rem;
            margin: 0.15rem 0 0;
            background: linear-gradient(45deg, #4e88ff, #4b5fff);
            border-radius: 0.14rem;
            line-height: 0.5rem;
            text-align: center;
            color: #fff;
            font-size: 0.18rem;
            font-family: inherit;
            border: 0;
            outline: none;
            transition: background-color 0.3s, border-color 0.3s, color 0.3s,
              box-shadow 0.3s;
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
        .check_right {
          display: flex;
          justify-content: flex-end;
          align-items: center;
          flex-wrap: wrap;
          width: 50%;
          @media screen and (max-width: 441px) {
            width: 100%;
          }
          img {
            width: 95%;
            max-width: 1.8rem;
            margin: 0 auto 5px;
          }
          u {
            cursor: pointer;
            display: block;
            padding: 0;
            padding: 0.1rem 0;
            font-size: 0.18rem;
            font-weight: normal;
            line-height: 1.1;
            opacity: 0.9;
            color: #fff;
            @media screen and (max-width: 1600px) {
              font-size: 15px;
            }
            @media screen and (max-width: 1440px) {
              font-size: 14px;
            }
            @media screen and (max-width: 600px) {
              font-size: 13px;
            }
            u {
              opacity: 1;
            }
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
    .el-tree /deep/ {
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
    .el-table /deep/ {
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
            text-transform: capitalize;
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

            .hot-cold-box {
              position: relative;
              display: flex;
              align-items: center;
              justify-content: center;
              flex-wrap: wrap;
              .icon {
                display: block;
                width: 24px;
                height: 24px;
                margin: 0 0.15rem;
                font-size: 0.24rem;
                @media screen and (max-width: 1600px) {
                  width: 18px;
                  height: 18px;
                }
                &:hover {
                  opacity: 0.7;
                }
              }
              .icon_rename {
                background: url(../assets/images/space/icon_02.png) no-repeat
                  center;
                background-size: 100% 100%;
              }
              .icon_details {
                background: url(../assets/images/space/icon_03.png) no-repeat
                  center;
                background-size: 100% 100%;
              }
              .icon_delete {
                background: url(../assets/images/space/icon_04.png) no-repeat
                  center;
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
    z-index: 2;
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
      color: hsla(0, 0%, 100%, 0.75);
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
.slideAdd {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: #cecece;
  border-radius: 0.1rem;
  z-index: 1;
}
.slideAddBg {
  background-color: rgba(96, 96, 96, 0.5);
}
.slideComSoon {
  position: fixed;
}
.slideEmail {
  z-index: 2999;
}
.reverse_phase {
  .slideScroll {
    .fe-none {
      .uploadDig {
        .upload_progress {
          .moon {
            display: block;
          }
          .sunny {
            display: none;
          }
        }
      }
    }
  }
}
</style>
