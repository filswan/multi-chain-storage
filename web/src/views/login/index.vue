<template>
  <div class="login-page login">

    <div class="loginArea">
      <transition name="wrapperFade" mode="out-in">

        <!-- 登录 -->
        <div class="loginAreaRight" v-loading="loginLoad">
          <a href="javascript:void(0);" @click="menuIndexFun('/register',6)" class="login_register">{{generateLogin('login_h1')}} ></a>
          <h1>{{generateLogin('login_h2')}}</h1>

          <transition name="phoneTypeFade" mode="out-in">

            <!-- 邮箱登录 -->
            <div class="form-box mail-wrapper" key="mail">

              <div class="loginFormDiv">
                <input class="form-input" :placeholder="generateLogin('login_mailnum')" v-model="formData.mail.email" @blur="checkAccount(2)" />
              </div>
              <div class="form-tips" v-show="verify.mail.tipsbox">{{verify.mail.tips}}</div>

              <div class="loginFormDiv">
                <input type="password" class="form-input" :placeholder="generateLogin('login_mailpw')" v-model="formData.mail.password" />
              </div>
              <div class="form-tips" v-show="verify.mailPassword.tipsbox">{{verify.mailPassword.tips}}</div>

              <a href="javascript:void(0);" @click="menuIndexFun('/forget',7)" class="login_forgetPass">{{generateLogin('login_mailforget')}}</a>

              <div class="loginFormDiv" style="margin-bottom: 0;height:0.37rem">
                <el-button type="primary" @click="mailLogin">{{generateLogin('login_maillogin')}}</el-button>
              </div>

            </div>

          </transition>

        </div>

      </transition>

    </div>

  </div>
</template>

<script>
// import bus from '@/components/bus'
import * as myAjax from '@/api/login'
import { generateLogin } from '@/utils/i18n'

export default {
  name: 'login',
  data () {
    return {
      // 登录框加载遮罩
      loginLoad: false,
      // 登录类型，手机：phone，邮箱：mail
      formType: 'mail',
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
          tips: this.generateLogin('login_verify_phone_tips_true'),
          checkAccount: true,
          checked: false
        },
        phonePassword: {
          tipsbox: false,
          tips: this.generateLogin('login_verify_phonePassword_tips_true')
        },
        mail: {
          tipsbox: false,
          tips: this.generateLogin('login_verify_mail_tips_true'),
          checkAccount: true,
          checked: false
        },
        mailPassword: {
          tipsbox: false,
          tips: this.generateLogin('login_verify_mailPassword_tips_true')
        }
      },
      // 表单数据
      formData: {
        mail: {
          email: '',
          password: ''
        }
      },
      // 手机号验证正则
      phoneRegular: /^\s*\+?\s*(\(\s*\d+\s*\)|\d+)(\s*-?\s*(\(\s*\d+\s*\)|\s*\d+\s*))*\s*$/,
      // 邮箱验证正则
      mailRegular: /^\w+([-.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/,
      // 密码验证正则
      // passwordRegular: /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$/,
      passwordRegular: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,16}$/,
      // 区号列表
      zonenumList: [],
      // 验证框加载遮罩
      validateLoad: false,
      // 安全验证类型，手机：phone，，谷歌：google
      validateType: 'phone',
      // 安全验证账号
      validateData: {
        phone: '',
        secretPhone: '',
        loginType: 1,
        secretMail: '',
        mail: ''
      },
      // 安全验证验证码
      validateCode: {
        phoneCode: {
          code: '',
          tipsbox: false,
          tipsboxError: false
        },
        googleCode: {
          code: '',
          tipsbox: false,
          tipsboxError: false
        },
        mailCode: {
          code: '',
          tipsbox: false,
          tipsboxError: false
        }
      },
      // 是否在获取手机验证码
      isGetPhoneVerifyCode: false,
      // 是否在获取谷歌验证码
      isGetGoogleVerifyCode: false,
      // 是否在获取邮箱验证码
      isGetMailVerifyCode: false,
      // 获取手机验证码按钮文字
      getPhoneVerifyCodeWord: this.generateLogin('login_getVerifyCodeWord'),
      // 获取谷歌验证码按钮文字
      getGoogleVerifyCodeWord: this.generateLogin('login_getVerifyCodeWord'),
      // 获取邮箱验证码按钮文字
      getMailVerifyCodeWord: this.generateLogin('login_getVerifyCodeWord'),
      // 验证完成登录请求数据
      verifyLoginRegData: {},
      // 失焦判断账号变化
      oldAccount: {
        phone: '',
        mail: ''
      },
      // 記錄從那個頁面進入
      fromEnter: ''
    }
  },
  methods: {
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
    // 邮箱登录
    mailLogin (data) {
      var _this = this
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
      _this.loginLoad = true

      myAjax
        .login(_this.formData.mail)
        .then(response => {
          // console.log(response)
          if (response.status === 'success') {
            localStorage.setItem('mcsLoginAccessToken', response.auth_token)
            localStorage.setItem('mcsLoginEmail', _this.formData.mail.email)
            sessionStorage.oaxLoginpassword = _this.formData.mail.password
            _this.$store.state.user.accessToken = response.auth_token
            _this.$store.state.user.email = _this.formData.mail.email

            if (_this.fromEnter && _this.fromEnter !== '/supplierAllBack') {
              // 防止登录后需要跳转到指定页面
              _this.$router.push({ path: _this.fromEnter })
            } else {
              // this.$router.go(-1)
              _this.$router.push({ path: '/my_buckets' })
            }
            _this.loginLoad = false
          } else {
            _this.$message.error(response.message)
            _this.loginLoad = false
          }
        })
        .catch(error => {
          console.log(error)
          _this.loginLoad = false
        })
    },
    // 是否已登录
    isLogin () {
      var _this = this
      if (localStorage.getItem('mcsLoginAccessToken')) {
        _this.$router.push({ path: '/my_files' })
      }
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
    // console.log(this.$route.query.redirect)
    this.isLogin()
    var _this = this
    _this.fromEnter = this.$route.query.redirect
    if (sessionStorage.oaxLoginType === 'mail') {
      _this.formType = 'mail'
      sessionStorage.removeItem('oaxLoginType')
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scopte>
.login-page {
  $color: #0071ff;
  $color_hover: #0066e2;
  width: 100%;
  height: 100%;

  .loginAreaRight {
    margin: 0 auto;
    width: 81%;
    padding: 0;
    box-sizing: border-box;

    .title {
      margin: 0;
    }

    .el-select {
      position: relative;
      float: left;
      width: 29.5%;
      .form-input {
        height: 0.34rem;
        line-height: 0.34rem;
      }
      .el-input__icon {
        width: 0.25rem;
        display: -webkit-box;
        display: -ms-flexbox;
        display: flex;
        -webkit-box-pack: center;
        -ms-flex-pack: center;
        justify-content: center;
        -webkit-box-align: center;
        -ms-flex-align: center;
        align-items: center;
      }
    }
    .el-select::after {
      content: "";
      position: absolute;
      right: 0;
      top: 50%;
      width: 0.01rem;
      height: 0.17rem;
      margin-top: -0.085rem;
      background-color: #e0e0e0;
    }
    .noborder {
      float: left;
      width: calc(76% - 0.2rem);
      border: 0;
    }

    .validate-type-box {
      width: 81%;
      margin: 0 auto 2%;
      font-size: 0.15rem;
    }

    .login-tabbox {
      .login-tab.active {
        color: #000;
      }

      .login-tab:not(:first-child) {
        margin-left: 76px;

        .login-tab-word {
          float: left;
        }
      }

      .login-tab:not(:first-child)::before {
        content: "";
        position: absolute;
        display: block;
        background-color: #eee;
        width: 2px;
        height: 20px;
        left: -37px;
        top: 4px;
      }

      .login-tab {
        position: relative;
        float: left;
        width: 172px;
        font-size: 24px;
        color: #999;
        line-height: 28px;

        .login-tab-word {
          float: right;
          cursor: pointer;
        }
      }
    }

    .form-box {
      width: 100%;
      margin-top: 0.15rem;

      // .form-tips {
      //   font-size: 12px;
      //   line-height: 18px;
      //   color: #FF5656;
      //   text-align: right;
      //   position: absolute;
      //   bottom: -18px;
      //   left: 0;
      //   right: 0;
      // }

      .form-input-box:not(:first-child) {
        margin-top: 30px;
      }

      .form-input-box {
        border: 1px solid #eee;
        box-sizing: border-box;
        width: 100%;
        height: 48px;
        position: relative;

        .form-icon {
          display: block;
          float: left;
          width: 24px;
          height: 24px;
          margin: 11px 0 11px 18px;
        }

        .form-input {
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

      .phone-box {
        .form-input {
          width: 304px;
        }

        .form-zonenum-btn {
          float: left;
          width: 72px;
          height: 100%;
          background-color: #eee;
          line-height: 46px;
          cursor: pointer;

          .form-zonenum-num {
            float: left;
            color: #333;
            margin-left: 12px;
          }

          .form-zonenum-icon {
            display: block;
            float: left;
            width: 12px;
            height: 12px;
            margin: 16px 0 0 10px;
          }
        }
      }

      .form-btn:hover {
        background-color: $color_hover;
      }

      .form-btn {
        line-height: 48px;
        height: 48px;
        background-color: $color;
        color: #333;
        font-size: 16px;
        text-align: center;
        margin-top: 30px;
        cursor: pointer;
      }

      .form-other-btnbox {
        width: 100%;
        margin-top: 30px;
        height: 20px;

        .form-other-btn:hover {
          color: $color;
        }

        .form-other-btn {
          line-height: 20px;
          font-size: 14px;
          color: #999;
        }

        .form-other-btn-left {
          float: left;
        }

        .form-other-btn-right {
          float: right;
        }
      }
    }
  }

  .validate-wrapper {
    margin: 100px auto;
    width: 540px;
    background-color: #fff;
    padding: 56px 60px;
    box-sizing: border-box;

    .validate-title {
      width: 100%;
      text-align: center;
      font-size: 24px;
      color: #000;
      line-height: 28px;
    }

    .form-box {
      width: 100%;
      margin-top: 45px;

      .validate-box:not(:first-child) {
        margin-top: 30px;
      }

      .validate-box {
        width: 100%;

        .validate-type-box {
          width: 81%;
          margin: 0 auto 2%;
          font-size: 0.15rem;
        }

        // .form-tips {
        //   font-size: 12px;
        //   line-height: 18px;
        //   color: #FF5656;
        //   text-align: right;
        //   position: absolute;
        //   bottom: -18px;
        //   left: 0;
        //   right: 0;
        // }

        .form-input-box {
          border: 1px solid #eee;
          box-sizing: border-box;
          width: 100%;
          height: 48px;
          margin-top: 15px;
          position: relative;

          .form-icon {
            display: block;
            float: left;
            width: 24px;
            height: 24px;
            margin: 11px 0 11px 18px;
          }

          .form-input {
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
      }

      .validate-box-verifycode {
        .form-input-box {
          .form-input {
            width: 270px;
          }

          .validate-getverifycode-btn.active {
            color: #d8d8d8;
          }

          .validate-getverifycode-btn {
            float: left;
            width: 106px;
            line-height: 46px;
            color: $color;
            cursor: pointer;
            text-align: center;
          }
        }
      }

      .form-btn:hover {
        background-color: $color_hover;
      }

      .form-btn {
        line-height: 48px;
        height: 48px;
        background-color: $color;
        color: #333;
        font-size: 16px;
        text-align: center;
        margin-top: 30px;
        cursor: pointer;
      }
    }
  }

  .dialog-box {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 99;

    .dialog-mask {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background-color: rgba($color: #000000, $alpha: 0.3);
    }

    .dialog-verify-captcha {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translateX(-50%) translateY(-50%);
      z-index: 9;
    }

    .dialog-board {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translateX(-50%) translateY(-50%);
      z-index: 9;
      background-color: #fff;
      width: 480px;

      .dialog-header {
        padding: 20px 18px;
        box-sizing: border-box;
        line-height: 20px;
        border-bottom: 1px solid #eee;

        .dialog-title {
          float: left;
          line-height: 20px;
          color: #333;
          font-size: 16px;
        }

        .dialog-close {
          display: block;
          float: right;
          cursor: pointer;
          width: 14px;
          height: 14px;
          margin-top: 3px;
        }
      }

      .dialog-listbox::-webkit-scrollbar {
        width: 5px;
      }

      .dialog-listbox::-webkit-scrollbar-thumb {
        background-color: #bfbfbf;
      }

      .dialog-listbox {
        padding: 0 18px;
        box-sizing: border-box;
        max-height: 429px;
        overflow-y: scroll;
        -webkit-overflow-scrolling: touch;

        .dialog-list.active {
          color: $color;
        }

        .dialog-list:hover {
          color: $color;
        }

        .dialog-list {
          color: #333;
          cursor: pointer;
          font-size: 14px;
          line-height: 20px;
          padding: 12px 0;
          box-sizing: border-box;

          .dialog-list-left {
            float: left;
          }

          .dialog-list-right {
            float: right;
          }
        }
      }
    }
  }

  .border-red {
    // border-color: #FF5656 !important;
  }

  .wrapperFade-enter-active,
  .wrapperFade-leave-active {
    transition: opacity 0.3s;
  }

  .wrapperFade-enter,
  .wrapperFade-leave-to {
    opacity: 0;
  }

  .phoneTypeFade-enter-active,
  .phoneTypeFade-leave-active {
    transition: opacity 0.3s;
  }

  .phoneTypeFade-enter,
  .phoneTypeFade-leave-to {
    opacity: 0;
  }

  .dialogVerifyFade-enter-active,
  .dialogVerifyFade-leave-active {
    transition: opacity 0.3s;
  }

  .dialogVerifyFade-enter,
  .dialogVerifyFade-leave-to {
    opacity: 0;
  }
}
</style>
