<template>
<div class="login register">
    <div class="loginArea">
      <div class="loginAreaLeft">
        <img src="@/assets/images/login/login_logo_en.png" />
      </div>
      <div class="loginAreaRight" v-loading="regLoad">
        <a
          href="javascript:void(0);"
          @click="menuIndexFun('/login',7)"
          class="login_register1"
        >{{generateRegister('register_h1')}} ></a>
        <h1 style="margin:0.65rem 0px 0.4rem">{{generateRegister('register_registermail')}}</h1>

        <transition name="phoneTypeFade" mode="out-in">

          <!-- 邮箱注册 -->
          <div key="mail">

            <div class="loginFormDiv" :class="verify.mail.tipsbox ? 'border-red' : ''">
              <input class="form-input" :placeholder="generateRegister('register_mailnum')" v-model="formData.mail.email" @blur="checkAccount(2)"/>
            </div>
            <div class="form-tips" v-show="verify.mail.tipsbox">{{verify.mail.tips}}</div>

            <div class="loginFormDiv" :class="verify.mailPassword.tipsbox ? 'border-red' : ''">
              <input type="password" class="form-input" :placeholder="generateRegister('register_verify_placehold_tips_rule')" v-model="formData.mail.password" @blur="checkPass(2)"/>
            </div>
            <div class="form-tips" v-show="verify.mailPassword.tipsbox">{{verify.mailPassword.tips}}</div>

            <div class="loginFormDiv" :class="verify.mailPasswordVerify.tipsbox ? 'border-red' : ''">
              <input type="password" class="form-input" :placeholder="generateRegister('register_verify_pw_tips_reenter')" v-model="formData.mail.passwordComfirm" @blur="checkAccountPass(2)"/>
            </div>
            <div class="form-tips" v-show="verify.mailPasswordVerify.tipsbox">{{verify.mailPasswordVerify.tips}}</div>

            <div class="loginFormDiv" style="height:0.37rem">
              <el-button type="primary" @click="mailReg">{{generateRegister('register_mailreg')}}</el-button>
            </div>
          </div>

        </transition>

      </div>

      <div id="captcha"></div>

      <!-- 手机注册成功 -->
      <transition name="phoneRegSuccessFade">
        <div class="dialog-box" v-show="dialogRegSuccessShow">
          <div class="dialog-mask"></div>
          <div class="dialog-RegSuccess-box dialog-board">
            <!-- <img class="dialog-RegSuccess-icon" src="../../assets/images/cn.jpg"/> -->
            <div class="dialog-RegSuccess-content">{{generateRegister('register_regsuccess')}}<span class="dialog-RegSuccess-content-tips">（<span class="dialog-RegSuccess-content-tipstime">{{jumpTime}}</span>{{generateRegister('register_regsuccess_tips')}}）</span></div>
          </div>
        </div>
      </transition>

    </div>
</div>
</template>

<script>
import * as myAjax from '@/api/register'
import { generateRegister } from '@/utils/i18n'
export default {
  name: 'register',
  data () {
    return {
      // 注册框加载遮罩
      regLoad: false,
      // 注册类型，手机：phone，邮箱：mail
      formType: 'mail',
      // 区号弹框是否显示
      dialogZonenumShow: false,
      // 区号弹框加载遮罩
      zonenumLoad: false,
      // 滑块验证弹框是否显示
      dialogVerifyShow: false,
      // 手机注册成功弹窗是否显示
      dialogRegSuccessShow: false,
      // 倒计时
      jumpTime: '',
      // 表单验证
      verify: {
        phone: {
          tipsbox: false,
          tips: this.generateRegister('register_verify_phone_tips_true'),
          checkAccount: true
        },
        verifyCode: {
          tipsbox: false,
          tips: this.generateRegister('register_verify_verifyCode_tips_true')
        },
        phonePassword: {
          tipsbox: false,
          tips: this.generateRegister('register_verify_phonePassword_tips_true')
        },
        phonePasswordVerify: {
          tipsbox: false,
          tips: this.generateRegister('register_verify_phonePasswordVerify_tips_true')
        },
        term: {
          tipsbox: false,
          tips: this.generateRegister('register_verify_term_tips_true')
        },
        mail: {
          tipsbox: false,
          tips: this.generateRegister('register_verify_mail_tips_true'),
          checkAccount: true
        },
        mailPassword: {
          tipsbox: false,
          tips: this.generateRegister('register_verify_mailPassword_tips_true')
        },
        mailPasswordVerify: {
          tipsbox: false,
          tips: this.generateRegister('register_verify_mailPasswordVerify_tips_true')
        }
      },
      // 表单数据
      formData: {
        phone: {
          phone: '',
          verifyCode: '',
          password: '',
          passwordComfirm: '',
          code: '86',
          inviteCode: ''
        },
        mail: {
          email: '',
          password: '',
          passwordComfirm: ''
        }
      },
      // 区号列表
      zonenumList: [],
      // 是否在获取验证码
      isGetVerifyCode: false,
      // 获取验证码按钮文字
      getVerifyCodeWord: this.generateRegister('register_getVerifyCodeWord'),
      // 是否勾选用户服务条款
      isSelectTerm: false,
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
    generateRegister,
    // 邮箱注册
    mailReg () {
      var _this = this
      // 邮箱
      if (!_this.formData.mail.email) {
        _this.verify.mail.tips = _this.generateRegister('register_verify_mail_tips_empty')
        _this.verify.mail.tipsbox = true
        return false
      } else if (!_this.mailRegular.test(_this.formData.mail.email)) {
        _this.verify.mail.tips = _this.generateRegister('register_verify_mail_tips_true')
        _this.verify.mail.tipsbox = true
        return false
      } else {
        _this.verify.mail.tipsbox = false
      }
      _this.checkPass()
      _this.checkAccountPass()

      if (_this.verify.mailPassword.tipsbox === false && _this.verify.mailPasswordVerify.tipsbox === false) {
        _this.regLoad = true
        // 注册
        _this.formData.mail.source = 2
        myAjax
          .emailRegister(_this.formData.mail)
          .then(response => {
            // console.log(response)
            if (response.status === 'success') {
              sessionStorage.oaxRegisterMail = _this.formData.mail.email
              sessionStorage.oaxRegisterMailTime = new Date().getTime()
              _this.$router.push('/account_activation')
              _this.regLoad = false
            } else {
              _this.$message.error(response.message)
              console.log(response.message)
              _this.regLoad = false
            }
          })
          .catch(error => {
            console.log(error)
            // _this.$message.error(error)
          })
      }
    },
    // 校验密码格式
    checkPass (type) {
      var _this = this
      if (!_this.formData.mail.password) {
        _this.verify.mailPassword.tips = _this.generateRegister('register_verify_pw_tips_empty')
        _this.verify.mailPassword.tipsbox = true
        return false
      } else if (!_this.passwordRegular.test(_this.formData.mail.password)) {
        _this.verify.mailPassword.tips = _this.generateRegister('register_verify_pw_tips_rule')
        _this.verify.mailPassword.tipsbox = true
        return false
      } else {
        _this.verify.mailPassword.tipsbox = false
      }
    },
    // 校验账号
    checkAccount (type) {
      let _this = this
      if (_this.formData.mail.email === _this.oldAccount.mail) {
        return false
      }
      _this.oldAccount.mail = _this.formData.mail.email
      if (!_this.formData.mail.email) {
        _this.verify.mail.tips = _this.generateRegister('register_verify_mail_tips_empty')
        _this.verify.mail.tipsbox = true
        return false
      } else if (!_this.mailRegular.test(_this.formData.mail.email)) {
        _this.verify.mail.tips = _this.generateRegister('register_verify_mail_tips_true')
        _this.verify.mail.tipsbox = true
        return false
      } else {
        _this.verify.mail.tipsbox = false
      }
    },
    // 校验密码是否一致
    checkAccountPass (type) {
      var _this = this
      if (_this.formData.mail.passwordComfirm !== _this.formData.mail.password) {
        _this.verify.mailPasswordVerify.tipsbox = true
        return false
      } else {
        _this.verify.mailPasswordVerify.tipsbox = false
      }
    },
    menuIndexFun (to, index) {
      // this.$router.replace(to).catch(err => err);
      this.$router.push(to)
      document.documentElement.scrollTop = 0
    }
  },

  mounted () {
    var _this = this
    document.onkeydown = function (e) {
      if (e.keyCode === 13) {
        if (_this.formType === 'mail') {
          _this.mailReg()
        }
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
  components: {},
  computed: {
    languageMcs () {
      return this.$store.getters.languageMcs
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scopte>
  .border-red{
    border-color: #FF5656 !important;
  }
  .phoneTypeFade-enter-active,
  .phoneTypeFade-leave-active{
    transition: opacity .3s;
  }
  .phoneTypeFade-enter,
  .phoneTypeFade-leave-to{
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
  .phoneRegSuccessFade-enter-active,
  .phoneRegSuccessFade-leave-active{
    transition: opacity .3s;
  }
  .phoneRegSuccessFade-enter,
  .phoneRegSuccessFade-leave-to{
    opacity: 0;
  }
  .form-term-checkbox{
    float: left;
    margin: 0.07rem 0.05rem 0 0;
  }
</style>
