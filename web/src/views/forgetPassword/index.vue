<template>
  <div class="login register">
    <div class="loginArea">
    <div class="loginAreaLeft">
        <img src="@/assets/images/login/login_logo_en.png" />
    </div>
    <div class="loginAreaRight" v-loading="forgetLoad">
      <a
        href="javascript:void(0);"
        @click="menuIndexFun('/login',7)"
        class="login_register1"
      >{{generateForgetPassword('forgetPassword_h1')}} ></a>
      <h1>{{generateForgetPassword('forgetPassword_mail')}}</h1>

      <transition name="forgetTypeFade" mode="out-in">

        <!-- 邮箱找回 -->
        <div class="form-box-forget mail-wrapper" key="mail">

          <div class="loginFormDiv" :class="verify.mail.tipsbox ? 'border-red' : ''" style="margin-top: 0.4rem">
            <input class="form-input" :placeholder="generateForgetPassword('forgetPassword_mailnum')" v-model="formData.mail.email" @blur="checkAccount(2)"/>
          </div>
          <div class="form-tips" v-show="verify.mail.tipsbox">{{verify.mail.tips}}</div>

          <div class="loginFormDiv" style="margin-top: 0.2rem;height: 0.37rem;">
            <el-button type="primary" @click="mailSend">{{generateForgetPassword('forgetPassword_mailbtn')}}</el-button>
          </div>

        </div>

      </transition>

    </div>

    <div id="captcha"></div>

  </div>
  </div>
</template>

<script>
import * as myAjax from '@/api/forgetPassword'
import { generateForgetPassword, generateLogin } from '@/utils/i18n'
export default {
  name: 'forgetPassword',
  data () {
    return {
      // 找回框加载遮罩
      forgetLoad: false,
      // 找回类型，手机：phone，邮箱：mail
      formType: 'phone',
      // 区号弹框是否显示
      dialogZonenumShow: false,
      // 区号弹框加载遮罩
      zonenumLoad: false,
      // 滑块验证弹框是否显示
      dialogVerifyShow: false,
      // 表单验证
      verify: {
        phone: {
          tipsbox: false,
          tips: this.generateForgetPassword('forgetPassword_verify_phone_tips_true'),
          checkAccount: true
        },
        verifyCode: {
          tipsbox: false,
          tips: this.generateForgetPassword('forgetPassword_verify_verifyCode_tips_true')
        },
        phonePassword: {
          tipsbox: false,
          tips: this.generateForgetPassword('forgetPassword_verify_phonePassword_tips_true')
        },
        phonePasswordVerify: {
          tipsbox: false,
          tips: this.generateForgetPassword('forgetPassword_verify_phonePasswordVerify_tips_true')
        },
        mail: {
          tipsbox: false,
          tips: this.generateForgetPassword('forgetPassword_verify_mail_tips_true'),
          checkAccount: true
        }
      },
      // 表单数据
      formData: {
        phone: {
          phone: '',
          verifyCode: '',
          password: '',
          passwordComfirm: '',
          code: '86'
        },
        mail: {
          email: ''
        }
      },
      // 区号列表
      zonenumList: [],
      // 是否在获取验证码
      isGetVerifyCode: false,
      // 获取验证码按钮文字
      getVerifyCodeWord: this.generateForgetPassword('forgetPassword_getVerifyCodeWord'),
      // 手机号验证正则
      phoneRegular: /^\s*\+?\s*(\(\s*\d+\s*\)|\d+)(\s*-?\s*(\(\s*\d+\s*\)|\s*\d+\s*))*\s*$/,
      // 邮箱验证正则
      mailRegular: /^\w+([-.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/,
      // 密码验证正则
      // passwordRegular: /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$/,
      passwordRegular: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,16}$/,
      // 失焦判断账号变化
      oldAccount: {
        phone: '',
        mail: ''
      }
    }
  },
  methods: {
    generateForgetPassword,
    generateLogin,
    // 校验账号
    checkAccount (type) {
      let _this = this
      if (_this.formData.mail.email === _this.oldAccount.mail) {
        return false
      }
      _this.oldAccount.mail = _this.formData.mail.email
      if (!_this.formData.mail.email) {
        _this.verify.mail.tips = _this.generateLogin('login_verify_mail_tips_empty')
        _this.verify.mail.tipsbox = true
        return false
      } else if (!_this.mailRegular.test(_this.formData.mail.email)) {
        _this.verify.mail.tips = _this.generateLogin('login_verify_mail_tips_true')
        _this.verify.mail.tipsbox = true
        return false
      } else {
        _this.verify.mail.tipsbox = false
      }
    },

    // 发送找回邮件
    mailSend () {
      var _this = this
      // 邮箱
      if (!_this.formData.mail.email) {
        _this.verify.mail.tips = _this.generateForgetPassword('forgetPassword_verify_mail_tips_empty')
        _this.verify.mail.tipsbox = true
        return false
      } else if (!_this.mailRegular.test(_this.formData.mail.email)) {
        _this.verify.mail.tips = _this.generateForgetPassword('forgetPassword_verify_mail_tips_true')
        _this.verify.mail.tipsbox = true
        return false
      } else {
        _this.verify.mail.tipsbox = false
      }
      _this.forgetLoad = true
      var checkData = {
        email: _this.formData.mail.email,
        source: 2
      }
      myAjax
        .sendForgetPasswordUrl(checkData)
        .then(response => {
          // _this.forgetLoad = false
          // console.log(response)
          if (response.status === 'success') {
            // 用户已存在
            _this.verify.mail.tipsbox = false
            _this.verify.mail.checkAccount = false
            _this.forgetLoad = true

            // 发送邮箱找回邮件
            sessionStorage.oaxForgetMail = _this.formData.mail.email
            // console.log(sessionStorage.oaxForgetMail, _this.formData.mail.email)
            _this.$router.push({path: '/mail_forget'})
          }
          // else {
          //   // 用户不存在
          //   _this.verify.mail.tips = _this.generateForgetPassword('forgetPassword_verify_account_tips_no')
          //   _this.verify.mail.tipsbox = true
          //   _this.verify.mail.checkAccount = true
          //   _this.forgetLoad = false
          // }
        })
        .catch(error => {
          console.log(error)
          _this.forgetLoad = false
        })
    },
    menuIndexFun (to, index) {
      // this.$router.replace(to).catch(err => err);
      this.$router.push(to)
      document.documentElement.scrollTop = 0
    }
  },
  computed: {
    languageMcs () {
      return this.$store.getters.languageMcs
    }
  },
  mounted () {
    var _this = this
    document.onkeydown = function (e) {
      if (e.keyCode === 13) {
        _this.mailSend()
      }
    }
  },
  destroyed () {
    document.onkeydown = function (e) {
      if (e.keyCode === 13) {
        e.returnValue = false
        return false
      }
    }
  },
  components: {}
}
</script>

<style rel="stylesheet/scss" lang="scss" scopte>
.forget-page{
  $color: #ffb933;
  $color_hover: #e0a32d;
  position: relative;
  width: 100%;
  height: 100%;
  .forget-wrapper{
    margin: 100px auto;
    width: 540px;
    background-color: #fff;
    padding: 56px 60px;
    box-sizing: border-box;
    .forget-title{
      width: 100%;
      line-height: 28px;
      text-align: center;
      color: #000;
      font-size: 24px;
    }
    .form-box-forget{
      width: 100%;
      margin-top: 45px;
      .form-tips{
        font-size: 12px;
        line-height: 18px;
        color: #FF5656;
        text-align: right;
        position: absolute;
        bottom: -18px;
        left: 0;
        right: 0;
      }
      .form-input-box:not(:first-child){
        margin-top: 30px;
      }
      .form-input-box{
        border: 1px solid #eee;
        box-sizing: border-box;
        width: 100%;
        height: 48px;
        position: relative;
        .form-icon{
          display: block;
          float: left;
          width: 24px;
          height: 24px;
          margin: 11px 0 11px 18px;
        }
        .form-input{
          display: block;
          float: left;
          height: 100%;
          border: none;
          outline: none;
          box-sizing: border-box;
          padding: 0 15px;
          width: 376px;
          font-size: 14px;
          color: #333;
        }
      }
      .form-inputbtn-box{
        border: none;
        .form-inputbtn-bar{
          float: left;
          border: 1px solid #eee;
          box-sizing: border-box;
          width: 280px;
          height: 100%;
          .form-input{
            width: 236px;
          }
        }
        .form-inputbtn-box-btn.active:hover{
          background-color: #d8d8d8;
        }
        .form-inputbtn-box-btn.active{
          background-color: #d8d8d8;
        }
        .form-inputbtn-box-btn:hover{
          background-color: $color_hover;
        }
        .form-inputbtn-box-btn{
          float: left;
          margin-left: 18px;
          width: 122px;
          background-color: $color;
          color: #333;
          text-align: center;
          line-height: 48px;
          cursor: pointer;
        }
      }
      .phone-box{
        .form-input{
          width: 304px;
        }
        .form-zonenum-btn{
          float: left;
          width: 72px;
          height: 100%;
          background-color: #eee;
          line-height: 46px;
          cursor: pointer;
          .form-zonenum-num{
            float: left;
            color: #333;
            margin-left: 12px;
          }
          .form-zonenum-icon{
            display: block;
            float: left;
            width: 12px;
            height: 12px;
            margin: 16px 0 0 10px;
          }
        }
      }
      .form-btn:hover{
        background-color: $color_hover;
      }
      .form-btn{
        line-height: 48px;
        height: 48px;
        background-color: $color;
        color: #333;
        font-size: 16px;
        text-align: center;
        margin-top: 30px;
        cursor: pointer;
      }
      .form-other-btnbox{
        width: 100%;
        margin-top: 30px;
        height: 20px;
        .form-other-btn:hover{
          color: $color;
        }
        .form-other-btn{
          line-height: 20px;
          font-size: 14px;
          color: #999;
          width: 100%;
          text-align: center;
          cursor: pointer;
        }
        .form-other-btn-left{
          float: left;
        }
        .form-other-btn-right{
          float: right;
        }
      }
      .form-term-box{
        width: 100%;
        line-height: 20px;
        color: #999;
        font-size: 14px;
        margin-top: 30px;
        position: relative;
        cursor: pointer;
        .form-term-checkbox{
          display: block;
          float: left;
          width: 12px;
          height: 12px;
          margin-top: 4px;
        }
        .form-term-word{
          float: left;
          margin-left: 5px;
          .form-term-word-link{
            color: $color;
          }
        }
      }
    }
  }
  .dialog-box{
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 99;
    .dialog-mask{
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background-color: rgba($color: #000000, $alpha: .3);
    }
    .dialog-verify-captcha{
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translateX(-50%) translateY(-50%);
      z-index: 9;
    }
    .dialog-board{
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translateX(-50%) translateY(-50%);
      z-index: 9;
      background-color: #fff;
      width: 480px;
      .dialog-header{
        padding: 20px 18px;
        box-sizing: border-box;
        line-height: 20px;
        border-bottom: 1px solid #eee;
        .dialog-title{
          float: left;
          line-height: 20px;
          color: #333;
          font-size: 16px;
        }
        .dialog-close{
          display: block;
          float: right;
          cursor: pointer;
          width: 14px;
          height: 14px;
          margin-top: 3px;
        }
      }
      .dialog-listbox::-webkit-scrollbar{
        width: 5px;
      }
      .dialog-listbox::-webkit-scrollbar-thumb{
        background-color: #BFBFBF;
      }
      .dialog-listbox{
        padding: 0 18px;
        box-sizing: border-box;
        max-height: 429px;
        overflow-y: scroll;
        -webkit-overflow-scrolling: touch;
        .dialog-list.active{
          color: $color;
        }
        .dialog-list:hover{
          color: $color_hover;
        }
        .dialog-list{
          color: #333;
          cursor: pointer;
          font-size: 14px;
          line-height: 20px;
          padding: 12px 0;
          box-sizing: border-box;
          .dialog-list-left{
            float: left;
          }
          .dialog-list-right{
            float: right;
          }
        }
      }
    }
  }
  .border-red{
    border-color: #FF5656 !important;
  }
  .forgetTypeFade-enter-active,
  .forgetTypeFade-leave-active{
    transition: opacity .3s;
  }
  .forgetTypeFade-enter,
  .forgetTypeFade-leave-to{
    opacity: 0;
  }
  .dialogVerifyFade-enter-active,
  .dialogVerifyFade-leave-active{
    transition: opacity .3s;
  }
  .dialogVerifyFade-enter,
  .dialogVerifyFade-leave-to{
    opacity: 0;
  }
}
</style>
