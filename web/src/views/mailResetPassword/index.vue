<template>
  <div class="login">
    <div class="loginArea">

      <!-- <div class="forget-title">{{generateForgetPassword('forgetPassword_title')}}</div> -->

    <div class="loginAreaLeft">
        <img src="@/assets/images/login/login_logo_en.png" />
    </div>
    <!-- 邮箱密码重置 -->
    <div class="loginAreaRight" v-loading="mailresetLoad">

      <h1>{{generateMailResetPassword('mailResetPassword_title')}}</h1>

      <div class="form-box mail-wrapper">

        <div class="loginFormDiv" :class="verify.mailPassword.tipsbox ? 'border-red' : ''">
          <input type="password" class="form-input" :placeholder="generateMailResetPassword('mailResetPassword_new_pw')" v-model="formData.mail.password"/>
        </div>
        <div class="form-tips" v-show="verify.mailPassword.tipsbox">{{verify.mailPassword.tips}}</div>

        <div class="loginFormDiv" :class="verify.mailPasswordVerify.tipsbox ? 'border-red' : ''">
          <input type="password" class="form-input" :placeholder="generateMailResetPassword('mailResetPassword_pwconfirm')" v-model="formData.mail.passwordComfirm"/>
        </div>
        <div class="form-tips" v-show="verify.mailPasswordVerify.tipsbox">{{verify.mailPasswordVerify.tips}}</div>

        <div class="loginFormDiv" style="margin-top: 0.4rem;height:0.37rem">
          <el-button type="primary" @click="mailReset">{{generateMailResetPassword('mailResetPassword_btn')}}</el-button>
        </div>

      </div>

    </div>

    <div id="captcha"></div>

  </div>
  </div>
</template>

<script>
import * as myAjax from '@/api/forgetPassword'
import { generateMailResetPassword } from '@/utils/i18n'
export default {
  name: 'mailResetPassword',
  data () {
    return {
      // 密码验证正则
      // passwordRegular: /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$/,
      passwordRegular: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,16}$/,
      // 找回框加载遮罩
      mailresetLoad: false,
      // 区号弹框是否显示
      dialogZonenumShow: false,
      // 滑块验证弹框是否显示
      dialogVerifyShow: false,
      // 表单验证
      verify: {
        mailPassword: {
          tipsbox: false,
          tips: this.generateMailResetPassword('mailResetPassword_verify_pw_tips_true')
        },
        mailPasswordVerify: {
          tipsbox: false,
          tips: this.generateMailResetPassword('mailResetPassword_verify_pwconfirm_tips_true')
        }
      },
      // 表单数据
      formData: {
        mail: {
          password: '',
          passwordComfirm: ''
        }
      },
      // 重置密码链接参数
      linkData: {
        email: '',
        checkCode: ''
      }
    }
  },
  methods: {
    generateMailResetPassword,
    // 判断是否有邮箱对象
    haveMail () {
      var _this = this
      if (_this.getUrlVars().email && _this.getUrlVars().checkCode) {
        _this.linkData = {
          email: decodeURIComponent(_this.getUrlVars().email),
          checkCode: decodeURIComponent(_this.getUrlVars().checkCode)
        }
        console.log(_this.linkData)
      } else {
        _this.$message.error(_this.generateMailResetPassword('mailResetPassword_link_tips'))
        _this.$router.push({ path: '/forget' })
      }
    },
    // 邮箱找回
    mailReset () {
      var _this = this
      // 邮箱密码
      if (!_this.formData.mail.password) {
        _this.verify.mailPassword.tips = _this.generateMailResetPassword('mailResetPassword_verify_pw_tips_true')
        _this.verify.mailPassword.tipsbox = true
        return false
      } else if (!_this.passwordRegular.test(_this.formData.mail.password)) {
        _this.verify.mailPassword.tips = _this.generateMailResetPassword('mailResetPassword_verify_pw_tips_rule')
        _this.verify.mailPassword.tipsbox = true
        return false
      } else {
        _this.verify.mailPassword.tipsbox = false
      }
      // 邮箱密码确认
      if (!_this.formData.mail.passwordComfirm) {
        _this.verify.mailPasswordVerify.tips = _this.generateMailResetPassword('mailResetPassword_verify_pw_tips_true')
        _this.verify.mailPasswordVerify.tipsbox = true
        return false
      } else if (_this.formData.mail.password !== _this.formData.mail.passwordComfirm) {
        _this.verify.mailPasswordVerify.tips = _this.generateMailResetPassword('mailResetPassword_verify_pw_tips_same')
        _this.verify.mailPasswordVerify.tipsbox = true
        return false
      } else {
        _this.verify.mailPasswordVerify.tipsbox = false
      }

      var reqData = {
        'password': _this.formData.mail.password,
        'repeatPassword': _this.formData.mail.passwordComfirm,
        'email': _this.linkData.email,
        'emailCode': _this.linkData.checkCode,
        source: 2
      }
      _this.mailresetLoad = true
      myAjax
        .forgetPasswordEmailUserUrl(reqData)
        .then(response => {
          console.log(response)
          _this.mailresetLoad = false
          if (response.status === 'success') {
            sessionStorage.oaxResetMail = reqData.email
            _this.$router.push({ path: '/mail_forget_success' })
          } else {
            _this.$message.error(response.message)
          }
        })
        .catch(error => {
          console.log(error)
          _this.mailresetLoad = false
        })
    },
    // 获取地址栏参数
    getUrlVars () {
      var vars = []
      var hash
      var hashes = window.location.href.slice(window.location.href.indexOf('?') + 1).split('&')
      for (var i = 0; i < hashes.length; i++) {
        hash = hashes[i].split('=')
        vars.push(hash[0])
        vars[hash[0]] = hash[1]
      }
      return vars
    }
  },

  mounted () {
    this.haveMail()
    var _this = this
    document.onkeydown = function (e) {
      if (e.keyCode === 13) {
        _this.mailReset()
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
  components: {
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scopte>
.loginAreaRight{
  .form-box{
    width: 100%;
  }
}
</style>
