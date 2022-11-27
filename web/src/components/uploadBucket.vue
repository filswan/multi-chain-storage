<template>
    <div id="Create">
        <el-dialog :modal="true" :width="widthDia" :visible.sync="uploadDigShow"
            custom-class="uploadDig" id="uploadDigBody"
            :before-close="closeDia">
            <div class="loadMetamaskPay" v-if="loading">
                <div>
                    <div class="el-loading-spinner"><svg viewBox="25 25 50 50" class="circular"><circle cx="50" cy="50" r="20" fill="none" class="path"></circle></svg><!----></div>
                </div>
            </div>
            <template slot="title"></template>
            <div class="upload_form">
                <div>
                    <el-upload
                        class="upload-demo" :style="{'border': ruleForm.fileList_tip?'2px dashed #f56c6c':'0'}"
                        drag
                        ref="uploadFileRef"
                        action="customize"
                        :http-request="uploadFile"
                        :file-list="ruleForm.fileList"
                        :on-change="handleChange"
                        :on-remove="handleRemove">
                        <img src="@/assets/images/space/icon_11.png" alt="">
                        <div class="el-upload__text">{{$t('metaSpace.upload_desc')}}</div>
                        <div class="el-upload__text"><small>{{$t('metaSpace.uploadDray_or')}}</small></div>
                        <el-button type="primary">
                            {{$t('metaSpace.uploadDray_text')}}
                        </el-button>
                    </el-upload>
                    <p v-if="ruleForm.fileList.length>0" style="display: flex;align-items: center;">
                        <svg t="1637031488880" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3310" style="width: 14px;height: 14px;margin: 0 6px 0 5px;"><path d="M512 1024a512 512 0 1 1 512-512 32 32 0 0 1-32 32h-448v448a32 32 0 0 1-32 32zM512 64a448 448 0 0 0-32 896V512a32 32 0 0 1 32-32h448A448 448 0 0 0 512 64z" fill="#999999" p-id="3311"></path><path d="M858.88 976a32 32 0 0 1-32-32V640a32 32 0 0 1 32-32 32 32 0 0 1 32 32v304a32 32 0 0 1-32 32z" fill="#999999" p-id="3312"></path><path d="M757.12 773.12a34.56 34.56 0 0 1-22.4-8.96 32 32 0 0 1 0-45.44l101.12-101.12a32 32 0 0 1 45.44 0 30.72 30.72 0 0 1 0 44.8l-101.12 101.76a34.56 34.56 0 0 1-23.04 8.96z" fill="#999999" p-id="3313"></path><path d="M960 773.12a32 32 0 0 1-22.4-8.96l-101.76-101.76a32 32 0 0 1 0-44.8 32 32 0 0 1 45.44 0l101.12 101.12a32 32 0 0 1-22.4 54.4z" fill="#999999" p-id="3314"></path></svg>
                        {{ruleForm.file_size}}
                    </p>
                    <p v-if="ruleForm.fileList_tip" style="color: #F56C6C;font-size: 12px;line-height: 1;">{{ruleForm.fileList_tip_text}}</p>
                </div>
            </div>
        </el-dialog>
    </div>
</template>

<script>
import commonFun from '@/utils/common';
let that
export default {
  name: 'uploadFiles',
  data () {
    return {
      loading: false,
      ruleForm: {
        fileList: [],
        fileList_tip: false,
        fileList_tip_text: '',
        file_size: '',
        file_size_byte: ''
      },
      rules: {
        fileList: [
          { type: 'array', required: true, message: this.$t('uploadFile.file_name_tip'), trigger: 'change' }
        ]
      },
      loading: false,
      bodyWidth: document.documentElement.clientWidth < 1024,
      fileListTip: false,
      width: document.body.clientWidth > 600 ? '400px' : '95%',
      widthUpload: document.body.clientWidth > 600 ? '450px' : '95%',
      widthDia: document.body.clientWidth <= 600 ? '95%' : document.body.clientWidth <= 1440 ? '7rem' : '6.6rem'
    }
  },
  props: ['uploadDigShow', 'dataCont', 'currentBucket'],
  components: {},
  watch: {},
  methods: {
    closeDia () {
      that.$emit('getUploadDialog', false)
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
      var k = 1024, // or 1000
        sizes = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
        i = Math.floor(Math.log(bytes) / Math.log(k))

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
    async handleChange (file, fileList) {
      if (fileList.length > 0) {
        that.ruleForm.fileList = [fileList[fileList.length - 1]] // 这一步，是 展示最后一次选择的csv文件

        that.loading = true
        const params = {
          'path': that.currentBucket,
          'size': that.ruleForm.fileList[0].size,
          'name': that.ruleForm.fileList[0].name,
          'policy_id': that.dataCont.policy.id,
          'last_modified': new Date().getTime()
        }
        const directoryRes = await commonFun.sendRequest(`${process.env.BASE_METASPACE}api/v3/file/upload`, 'put', params)
        if (!directoryRes || directoryRes.status != 'success') {
          that.$message.error(directoryRes ? directoryRes.status : 'Fail')
        }
        const session = directoryRes.data.sessionID

        const uploadRes = await commonFun.sendRequest(`${process.env.BASE_METASPACE}api/v3/file/upload/${session}/0`, 'post', that._file)
        if (!uploadRes || uploadRes.status != 'success') {
          that.$message.error(uploadRes ? uploadRes.status : 'Fail')
        }
        that.$message({
          message: uploadRes.status ? uploadRes.status : 'Success',
          type: 'success'
        })
        await commonFun.timeout(500)
        that.$emit('getUploadDialog', false, 1)
        that.loading = false
      }
    },
    handleRemove (file, fileList) {
      // console.log(file, fileList);
      that.ruleForm.fileList = []
      that.ruleForm.file_size = ''
      that.ruleForm.file_size_byte = ''
    }
  },
  mounted () {
    that = this
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
    }
  }
}
</script>

<style scoped lang="scss">
    .el-dialog__wrapper /deep/ {
        display: flex;
        align-items: center;
        .metaM{
            margin: auto !important;
            .el-dialog__body{
                padding: 0.25rem 0.25rem 0.2rem;
                .el-row{
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
                .el-col{
                    text-align: left;
                    font-size: 0.18rem;
                    img{
                    float: right;
                    height: 0.24rem;
                    }
                }
                }
            }
            .el-dialog__footer{
                padding: 0 0.25rem 0.3rem;
                .dialog-footer{
                .el-button{
                    width: 100%;
                    font-size: 0.18rem;
                    height: 0.4rem;
                    padding: 0;
                    background: #5c3cd3;
                    color: #fff;
                    border-radius: 0.08rem;
                    &:hover{
                    background: #4326ab;
                    }
                }
                p{
                    font-size: 0.13rem;
                    line-height: 1.5;
                    color: red;
                    margin: 0.1rem 0 0;
                }
                }
            }
        }
        .completeDia{
            margin: auto !important;
            text-align: center;
            box-shadow: 0 0 13px rgba(128,128,128,0.8);
            border-radius: 0.2rem;
            .el-dialog__header{
            display: none;
            }
            img{
                display: block;
                max-width: 100px;
                margin: 0 auto 0.3rem;
            }
            h1{
                margin: 0.22rem auto 0.1rem;
                font-size: 0.32rem;
                font-weight: 500;
                line-height: 1.2;
                color: #191919;
                word-break: break-word;
            }
            h2{
                margin: 0 auto 0.1rem;
                font-size: 0.22rem;
                font-weight: 600;
                line-height: 1.4;
                color: #555;
                word-break: break-word;
                text-align: center;
            }
            h4{
                font-size: 0.2rem;
                font-weight: 500;
                line-height: 1.4;
                color: #555;
                word-break: break-word;
                text-align: center;
            }
            h3, a{
                font-size: 0.22rem;
                font-weight: 500;
                line-height: 1.4;
                color: #191919;
                word-break: break-word;
            }
            a{
                text-decoration: underline;
                color: #007bff;
            }
            a.a-close{
                height: 0.6rem;
                line-height: 0.6rem;
                padding: 0 45px;
                background: linear-gradient(45deg,#4f8aff, #4b5eff);
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
        .fileUpload{
            margin: auto !important;
            box-shadow: 0 0 13px rgba(128,128,128,0.8);
            border-radius: 0.2rem;
            .el-dialog__header{
                padding: 0.3rem 0.4rem;
                color: #000;
                font-size: 0.22rem;
                font-weight: 500;
                line-height: 1;
                text-transform: capitalize;
            }
            .el-dialog__body{
                padding: 0 0.4rem;
                h3{
                    margin: 0 0 0.1rem;
                    font-size: 0.2rem;
                    font-weight: normal;
                    line-height: 1.4;
                    color: #555;
                    word-break: break-word;
                }
                .gif_img{
                    max-width: 200px;
                    width: 100%;
                    margin: auto;
                    display: block;
                }
            }
        }
        .uploadDig{
            background: #fff;
            margin: auto !important;
            box-shadow: 0 0 13px rgba(128,128,128,0.8);
            border-radius: 0.2rem;
            .el-dialog__header{
                padding: 0.3rem 0.4rem;
                display: flex;
                align-items: center;
                border-bottom: 1px solid #dfdfdf;
                color: #000;
                font-size: 0.22rem;
                font-weight: 500;
                line-height: 1;
                text-transform: capitalize;
                display: none;
                @media screen and (max-width: 479px){
                    padding: 0.3rem 0.2rem;
                }
                img{
                    width: 20px;
                    height: 20px;
                    margin: 0 0 0 5px;
                    cursor: pointer;
                    @media screen and (max-width:1440px){
                        width: 17px;
                        height: 17px;
                    }
                    @media screen and (max-width: 1280px){
                        width: 16px;
                        height: 16px;
                    }
                }
                .el-dialog__title{
                    color: #000;
                    font-size: 0.22rem;
                    font-weight: 500;
                    line-height: 1;
                    text-transform: capitalize;
                }
                .el-dialog__headerbtn{
                    display: none;
                }
            }
            .el-dialog__body{
                padding: 0;
                .upload_form{
                    // display: flex;
                    // align-items: baseline;
                    width: calc(100% - 0.5rem);
                    padding: 0.25rem;
                    margin: auto;
                    justify-content: flex-start;
                    .upload-demo{
                        width: 100%;
                        margin: 0;
                        .el-upload{
                            width: 100%;
                            height: auto;
                            border: 0;
                            .el-upload-dragger{
                                width: 100%;
                                height: auto;
                                padding: 0.3rem 0;
                                color: #97a8be;
                                border-color: #898989;
                                svg, img{
                                    width: 0.45rem;
                                    height: 0.45rem;
                                    margin: 0;
                                }
                                .el-upload__text{
                                    margin: 10px 0;
                                    font-size: 0.22rem;
                                    font-weight: normal;
                                    color: #606060;
                                    line-height: 1;
                                    small{
                                        color: #969696;
                                    }
                                }
                                .el-button{
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
                        .el-upload-list__item:first-child{
                            margin-top: 0;
                        }
                        .el-upload-list{
                            width: 100%;
                            float: none;
                            clear: both;
                        }
                    }
                    .el-form{
                        width: 100%;
                        margin: 0;
                        .el-form-item::after, .el-form-item::before{
                            display: none;
                        }
                        .el-form-item{
                            display: flex;
                            align-items: center;
                            justify-content: space-between;
                            width: 100%;
                            margin: 0.2rem auto;
                            .el-form-item__label{
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
                                @media screen and (max-width: 479px){
                                    padding: 0;
                                }
                                img{
                                    width: 20px;
                                    height: 20px;
                                    margin: 0 0 0 5px;
                                    cursor: pointer;
                                    @media screen and (max-width:1440px){
                                        width: 17px;
                                        height: 17px;
                                    }
                                    @media screen and (max-width: 1280px){
                                        width: 16px;
                                        height: 16px;
                                    }
                                }
                                &::before{
                                    display: none;
                                }
                            }
                            .el-form-item__content{
                                width: 50%;
                                display: flex;
                                align-items: center;
                                justify-content: flex-end;
                                font-size: 0.2rem;
                                white-space: normal;
                                word-break: break-word;
                                line-height: 1.5;
                                color: #555;
                                h4{
                                    width: 100%;
                                    font-size: 0.1372rem;
                                    font-weight: 500;
                                    line-height: 1.7;
                                }
                                h5{
                                    width: 90%;
                                    margin-top: 5px;
                                    font-size: 0.11rem;
                                    font-weight: 500;
                                    line-height: 1.2;
                                    color: #737373;
                                }
                                .el-tag, .el-button--small{
                                    margin: 0 5px 5px 0;
                                }
                                .el-input{
                                    width: auto;
                                    .el-input__inner{
                                        height: 0.32rem;
                                        font-size: 0.1372rem;
                                        line-height: 0.32rem;
                                    }
                                    .el-input__suffix{
                                        display: none;
                                    }
                                }
                                .el-form-item__error {
                                    padding-top: 0;
                                    margin: 0 0.1rem;
                                    position: relative;
                                    float: right;
                                }
                                .el-textarea{
                                    width: 90% !important;
                                }
                                .upload-demo{
                                    display:flex;
                                    align-items: center;
                                    flex-wrap: wrap;
                                    .el-upload-list__item:first-child{
                                        margin-top: 0;
                                    }
                                    .el-upload--text{
                                        float: left;
                                        width: auto;
                                        height: auto;
                                        text-align: left;
                                        border: 0;
                                        .el-button--primary{
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
                                            background: linear-gradient(45deg,#4e88ff, #4b5fff);
                                            border-radius: 0.14rem;
                                            line-height: 0.5rem;
                                            text-align: center;
                                            color: #fff;
                                            font-size: 0.18rem;
                                            font-family: inherit;
                                            border: 0;
                                            outline: none;
                                            transition: background-color .3s, border-color .3s, color .3s, box-shadow .3s;
                                            cursor: pointer;
                                            span{
                                                display: flex;
                                                align-items: center;
                                            }
                                            img{
                                                display: inline-block;
                                                height: 20px;
                                                margin: 0 0.1rem 0 0;
                                                @media screen and (max-width: 1280px){
                                                    height: 16px;
                                                }
                                            }
                                            &:hover{
                                                opacity: .9;
                                                box-shadow: 0 12px 12px -12px rgba(12, 22, 44, 0.32);
                                            }
                                        }
                                    }
                                    .el-upload-list{
                                        width: 100%;
                                        max-width: 300px;
                                        float: none;
                                        clear: both;
                                        @media screen and (max-width: 1440px){
                                            max-width: 250px;
                                        }
                                    }
                                }
                                .el-upload__tip{
                                    // float: left;
                                    height: 100%;
                                    align-items: center;
                                    display: flex;
                                    margin: 0 0 0 0.1rem;
                                    color: #737373;
                                    line-height: 1;
                                    font-size: 0.12rem;
                                }
                                .el-radio{
                                    .el-radio__inner{
                                        border-color: #d9d9d9;
                                        background-color: #d9d9d9;
                                    }
                                }
                                .el-radio.is-checked{
                                    .el-radio__inner{
                                        border-color: #0b318f;
                                        background-color: #0b318f;
                                    }
                                    .el-radio__inner::after{
                                        width: 6px;
                                        height: 6px;
                                    }
                                }
                            }
                        }
                    }
                }
                .upload_plan{
                    width: 100%;
                    margin: auto;
                    justify-content: flex-start;
                    .title{
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
                        img{
                            width: 20px;
                            height: 20px;
                            margin: 0 0 0 5px;
                            cursor: pointer;
                            @media screen and (max-width:1440px){
                                width: 17px;
                                height: 17px;
                            }
                            @media screen and (max-width: 1280px){
                                width: 16px;
                                height: 16px;
                            }
                        }
                    }
                    .desc{
                        margin: 0 0 0.1rem;
                        line-height: 1.5;
                        font-size: 0.16rem;
                        white-space: normal;
                        color: #999;
                        font-weight: normal;
                    }
                    .upload_plan_radio{
                        .el-radio-group{
                            width: 100%;
                            background: #f7f7f7;
                            border-radius: 0.2rem;
                            .el-radio{
                                display: flex;
                                align-items: center;
                                justify-content: space-between;
                                width: 100%;
                                height: auto;
                                padding: 0.2rem 0.3rem;
                                margin: auto;
                                border: 0;
                                // line-height:30px;
                                .el-radio__input{
                                    width: 20px;
                                    display: flex;
                                    align-items: center;
                                    .el-radio__inner{
                                        border-color: #555;
                                    }
                                }
                                .el-radio__input.is-checked{
                                    .el-radio__inner{
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
                                            left:0;
                                            top:0;
                                            transform: translate(0, 0) scale(1);
                                            transition: all 0.15s;
                                        }
                                    }
                                }
                                .el-radio__label{
                                    display: flex;
                                    justify-content: space-between;
                                    width: calc(100% - 30px);
                                    .title{
                                        font-size: 0.2rem;
                                        line-height: 1;
                                    }
                                    .cont{
                                        font-size: 0.2rem;
                                        font-weight: 500;
                                        line-height: 1;
                                        text-align: center;
                                    }
                                }
                            }
                            .el-radio:nth-child(3n+1){
                                .el-radio__label{
                                    .cont{
                                        color: #35AD92;
                                    }
                                }
                            }
                            .el-radio:nth-child(3n+2){
                                border-top: 1px solid #dfdfdf;
                                border-bottom: 1px solid #dfdfdf;
                                .el-radio__label{
                                    .cont{
                                        color: #2C7FF8;
                                    }
                                }
                            }
                            .el-radio:nth-child(3n+3){
                                .el-radio__label{
                                    .cont{
                                        color: #F63D3D;
                                    }
                                }
                            }
                            .el-radio:hover{
                                background-color: rgba(64,158,255,0.1);
                            }

                        }

                    }
                    .upload_plan_radio_free{
                        opacity: .5;
                    }
                }
                .upload_bot{
                    width: 100%;
                    margin: 0.25rem auto 0.2rem;
                    text-align: center;
                    .found{
                        display: flex;
                        justify-content: space-between;
                        align-items: center;
                        width: 100%;
                        text-align: center;
                        .el-button{
                            height: 0.6rem;
                            padding: 0;
                            margin-left: 0;
                            line-height: 0.6rem;
                            font-size: 0.22rem;
                            font-family: inherit;
                            color: #fff;
                            border: 0;
                            background: linear-gradient(45deg,#4f8aff, #4b5eff);
                            border-radius: 14px;
                            width: calc(50% - 0.15rem);
                            &:hover{
                                opacity: .9;
                            }
                        }
                        .cancel{
                            background: #dadada;
                            transition: background-color .3s;
                            &:hover{
                                background: linear-gradient(45deg,#4f8aff, #4b5eff);
                            }
                        }
                    }
                    a{
                        margin: auto;
                        text-decoration: underline;
                        font-size: 0.18rem;
                        color: rgb(11, 49, 143);
                        cursor: pointer;
                    }
                }
            }
            .dialog-footer{
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
        .uploadDigID{
            position: absolute;
            left: 0;
            top: 0;
            right: 0;
            bottom: 0;
            z-index: 99;
            display: flex;
            align-items: center;
            justify-content: center;
            background-color: rgba(255,255,255,.9);
            border-radius: 0.2rem;
            font-size: 18px;
            @media screen and (max-width: 1600px){
                font-size: 16px;
            }
            .elUpload{
                i{
                    display: flex;
                    justify-content: center;
                    font-size: 45px;
                    color: #979797;
                }
            }
            &:before{
                display: flex;
                justify-content: center;
                font-size: 55px;
                color: #979797;
                @media screen and (max-width: 1600px){
                    font-size: 50px;
                }
                @media screen and (max-width: 1200px){
                    font-size: 45px;
                }
            }
        }
        .loadMetamaskPay{
            position: absolute;
            left: 0;
            top: 0;
            right: 0;
            bottom: 0;
            z-index: 99;
            display: flex;
            align-items: center;
            justify-content: center;
            background-color: rgba(255,255,255,.9);
            border-radius: 0.2rem;
            .el-loading-spinner{
                top: 0;
                position: relative;
                margin: 0 0 0.2rem;
            }
            p{
                font-size: 14px;
                color: #555;
            }
        }
        .el-alert /deep/{
            position: absolute;
            left: 0;
            top: 0;
            .el-alert__content{
                display: flex;
                align-items: center;
            }
        }
        .upload{
            padding: 0 0.17rem 0.4rem;
            margin: 0 auto 0.2rem;
            background-color: #fff;
            border-radius: 0.1rem;
            overflow: hidden;
            .title{
                width: 100%;
                margin: 0 0 0.1rem;
                text-align: left;
                font-size: 0.1972rem;
                color: #000;
                line-height: 0.42rem;
                text-indent: 0.08rem;
            }
            .upload_title{
                width: 100%;
                margin: 0.1rem 0 0.3rem;
                text-align: left;
                font-size: 0.2rem;
                font-weight: 600;
                color: #000;
                line-height: 1.5;
                text-indent: 0;
            }
            .upload_form{
                // display: flex;
                // align-items: baseline;
                width: 80%;
                margin: auto;
                justify-content: flex-start;
                .el-form /deep/{
                    width: 96%;
                    margin: 0 2%;
                    .el-form-item{
                        display: flex;
                        align-items: center;
                        width: auto;
                        margin: 0.15rem auto;
                        .el-form-item__label{
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
                            img{
                                width: 0.16rem;
                                height: 0.16rem;
                                margin: 0 0 0 5px;
                                cursor: pointer;
                            }
                            &::before{
                                display: none;
                            }
                        }
                        .el-form-item__content{
                            display: flex;
                            align-items: center;
                            font-size: 0.18rem;
                            white-space: normal;
                            word-break: break-word;
                            line-height: 1.5;
                            color: #666;
                            @media screen and (max-width:600px){
                                font-size: 14px;
                            }
                            small{
                                margin: 2px 5px 0;
                                font-size: 12px;
                            }
                            h4{
                                width: 100%;
                                font-size: inherit;
                                font-weight: 500;
                                line-height: 1.7;
                            }
                            h5{
                                width: 90%;
                                margin-top: 5px;
                                font-size: 0.11rem;
                                font-weight: 500;
                                line-height: 1.2;
                                color: #737373;
                                @media screen and (max-width:600px){
                                    font-size: 12px;
                                }
                            }
                            .el-tag, .el-button--small{
                                margin: 0 5px 5px 0;
                            }
                            .el-input{
                                width: auto;
                                .el-input__inner{
                                    height: 0.32rem;
                                    font-size: inherit;
                                    line-height: 0.32rem;
                                }
                                .el-input__suffix{
                                    display: none;
                                }
                            }
                            .el-form-item__error {
                                padding-top: 0;
                                margin: 0 0.1rem;
                                position: relative;
                                float: right;
                            }
                            .el-textarea{
                                width: 90% !important;
                            }
                            .upload-demo{
                                display:flex;
                                align-items: center;
                                flex-wrap: wrap;
                                .el-upload-list__item:first-child{
                                    margin-top: 0;
                                }
                                .el-upload--text{
                                    float: left;
                                    width: auto;
                                    height: auto;
                                    text-align: left;
                                    border: 0;
                                    .el-button--primary{
                                        height: 0.32rem;
                                        padding: 0 0.2rem;
                                        margin: 0 5px 0 0;
                                        line-height: 0.32rem;
                                        background-color:transparent;
                                        border: 1px solid #2c4c9e;
                                        border-radius: 0.08rem;
                                        color: #2c4c9e;
                                        font-size: 0.18rem;
                                    }
                                }
                                .el-upload-list{
                                    width: 100%;
                                    float: none;
                                    clear: both;
                                }
                            }
                            .el-upload__tip{
                                // float: left;
                                height: 100%;
                                align-items: center;
                                display: flex;
                                margin: 0 0 0 0.1rem;
                                color: #737373;
                                line-height: 1;
                                font-size: 0.12rem;
                                @media screen and (max-width:600px){
                                    font-size: 14px;
                                }
                            }
                            .el-radio{
                                .el-radio__inner{
                                    border-color: #d9d9d9;
                                    background-color: #d9d9d9;
                                }
                            }
                            .el-radio.is-checked{
                                .el-radio__inner{
                                    border-color: #0b318f;
                                    background-color: #0b318f;
                                }
                                .el-radio__inner::after{
                                    width: 6px;
                                    height: 6px;
                                }
                            }
                        }
                    }
                }
            }
            .upload_plan{
                width: 80%;
                margin: auto;
                justify-content: flex-start;
                .title{
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
                    img{
                        width: 0.16rem;
                        height: 0.16rem;
                        margin: 0 0 0 5px;
                        cursor: pointer;
                    }
                }
                .desc{
                    margin: 0 0 0.1rem;
                    line-height: 1.5;
                    text-align: center;
                    font-size: 0.18rem;
                    white-space: normal;
                    color: #999;
                    font-weight: normal;
                }
                .upload_plan_radio{
                    .el-radio-group /deep/{
                        width: 100%;
                        display: flex;
                        justify-content: center;
                        align-items: center;
                        .el-radio{
                            min-width: 20%;
                            height: auto;
                            padding: 0 0.1rem 0.15rem;
                            // line-height:30px;
                            .el-radio__input{
                                display: none;
                            }
                            .el-radio__label{
                                .title{
                                    margin: 0 0 0.05rem;
                                    font-size: 0.18rem;
                                }
                                .cont{
                                    font-size: 0.18rem;
                                    font-weight: bold;
                                    line-height: 1.5;
                                    text-align: center;
                                }
                            }
                        }
                        .el-radio:nth-child(3n+1){
                            .el-radio__label{
                                .cont{
                                    color: #56c4a6;
                                }
                            }
                        }
                        .el-radio:nth-child(3n+2){
                            .el-radio__label{
                                .cont{
                                    color: #4a92d3;
                                }
                            }
                        }
                        .el-radio:nth-child(3n+3){
                            .el-radio__label{
                                .cont{
                                    color: #922b26;
                                }
                            }
                        }
                        .is-checked{
                            position: relative;
                            &:after {
                                content: "";
                                display: block;
                                height: 25px;
                                width: 25px;
                                background-image: url(../assets/images/plan.png);
                                background-size: 100%;
                                position: absolute;
                                right:0;
                                top:0;
                            }
                        }
                        .el-radio:hover{
                            background-color: rgba(64,158,255,0.1);
                        }

                    }

                }
                .upload_plan_radio_free{
                    opacity: .5;
                }
            }
            .upload_bot{
                display: flex;
                justify-content: center;
                align-items: center;
                flex-wrap: wrap;
                width: 100%;
                margin: 0.25rem auto 0.15rem;
                .el-button /deep/{
                    height: 0.45rem;
                    padding: 0 0.4rem;
                    margin-left: 0;
                    background-color: #0b318f;
                    line-height: 0.45rem;
                    font-size: 0.18rem;
                    color: #fff;
                    border: 0;
                }
                .no_login{
                    background: #f56c6c;
                }
                .found{
                    width: 100%;
                    text-align: center;
                    a{
                        text-decoration: underline;
                        font-size: 0.18rem;
                        color: rgb(11, 49, 143);
                        cursor: pointer;
                    }
                }
            }
            .upload_result /deep/{
                width: 75%;
                margin: 0.6rem auto 0.2rem;
                .el-col{
                    display: flex;
                    align-items: flex-start;
                    font-size: 0.12rem;
                    color: #222;
                    margin: 0.15rem 0 0;
                    h5{
                        width: 100%;
                        font-size: 0.18rem;
                        font-weight: 500;
                        line-height: 1.3;
                        color: #000;
                    }
                    label{
                        width: 100px;
                        margin: 0 0 0 0.2rem;
                    }
                    p{
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
        #Create{
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
            .upload{
                padding: 0.1rem;
                .upload_form{
                    width: 90%;
                    flex-wrap: wrap;
                    .el-form /deep/{
                        width: 95%;
                    }
                }
            }
        }
    }
    @media screen and (max-width: 600px){
        #Create {
            .upload{
                 .upload_plan {
                     .upload_plan_radio {
                         .el-radio-group /deep/{
                             .el-radio{
                                 width: 32%;
                                 margin: auto;
                             }
                         }
                     }
                 }
            }
        }
    }
    @media screen and (max-width: 441px){
        #Create {
            .upload{
                 .title{
                    text-indent: 0;
                    font-size: 0.19rem;
                    line-height: 1.5;
                 }
                 .upload_form {
                     width: 100%;
                     .el-form /deep/{
                         .el-form-item{
                             flex-wrap: wrap;
                             .el-form-item__label{
                                width: 100%;
                                padding-bottom: 0.05rem;
                                justify-content: center;
                             }
                             .el-form-item__content{
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
                 .upload_result{
                     width: 90%;
                     margin: 0.2rem auto 0;
                     .el-col{
                         flex-wrap: wrap;
                         label{
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
