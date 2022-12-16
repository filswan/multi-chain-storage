<template>
  <div class="mailforget-page">
    <!-- 邮箱重置密码邮件发送 -->
    <div class="mailforget-wrapper" v-loading="mailforgetLoad">
      <div class="center_flex">
        <div class="mailforget-title">{{generateMailForget('mailForget_title')}}</div>

        <div class="mailforget-content">{{generateMailForget('mailForget_tips_1')}}<span>{{mail}}</span>{{generateMailForget('mailForget_tips_2')}}</div>

        <div class="form-btn" @click="sendAgain" :class="isGetPhoneVerifyCode === true ? 'active' : ''">{{getPhoneVerifyCodeWord}}</div>
      </div>
    </div>

    <div id="captcha"></div>

  </div>
</template>

<script>
import * as myAjax from '@/api/forgetPassword'
import { generateMailForget } from '@/utils/i18n'
export default {
  name: 'mailForget',
  data () {
    return {
      // 滑块验证弹框是否显示
      dialogVerifyShow: false,
      // 激活框加载遮罩
      mailforgetLoad: false,
      // 激活邮件地址
      mail: '',
      isGetPhoneVerifyCode: true,
      // 获取按钮文字
      getPhoneVerifyCodeWord: this.generateMailForget('getVerifyCodeWord')
    }
  },
  methods: {
    generateMailForget,
    // 判断是否有邮箱对象
    haveMail () {
      let _this = this
      if (!sessionStorage.oaxForgetMail) {
        _this.$router.push({ path: '/forget' })
      } else {
        _this.mail = sessionStorage.oaxForgetMail

        _this.isGetPhoneVerifyCode = true
        let getCodeTime = 60
        _this.getPhoneVerifyCodeWord = getCodeTime + _this.generateMailForget('getVerifyCodeWord_time')
        let codeTimes = setInterval(function () {
          if (getCodeTime <= 0) {
            clearInterval(codeTimes)
            _this.isGetPhoneVerifyCode = false
            _this.getPhoneVerifyCodeWord = _this.generateMailForget('getVerifyCodeWord')
          } else {
            getCodeTime -= 1
            _this.getPhoneVerifyCodeWord = getCodeTime + _this.generateMailForget('getVerifyCodeWord_time')
          }
        }, 1000)
      }
    },
    // 重新发送邮件
    sendAgain () {
      var _this = this
      if (_this.isGetPhoneVerifyCode) {
        return false
      }

      _this.mailforgetLoad = true
      var checkData = {
        email: _this.mail,
        source: 2
      }

      myAjax
        .sendForgetPasswordUrl(checkData)
        .then(response => {
          console.log(response)
          _this.mailforgetLoad = false
          if (response.status === 'success') {
            _this.$message({
              message: response.message,
              type: 'success'
            })

            _this.isGetPhoneVerifyCode = true
            let getCodeTime = 60
            _this.getPhoneVerifyCodeWord = getCodeTime + _this.generateMailForget('getVerifyCodeWord_time')
            let codeTimes = setInterval(function () {
              if (getCodeTime <= 0) {
                clearInterval(codeTimes)
                _this.isGetPhoneVerifyCode = false
                _this.getPhoneVerifyCodeWord = _this.generateMailForget('getVerifyCodeWord')
              } else {
                getCodeTime -= 1
                _this.getPhoneVerifyCodeWord = getCodeTime + _this.generateMailForget('getVerifyCodeWord_time')
              }
            }, 1000)
          } else {
            _this.$message.error(response.message)
          }
        })
        .catch(error => {
          console.log(error)
          _this.mailforgetLoad = false
        })
    }
  },

  mounted () {
    this.haveMail()
  },
  components: {
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scopte>
.mailforget-page{
  $color: #ffb933;
  $color_hover: #e0a32d;
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: #e6e6e6;
  display: flex;
  align-items: center;
  justify-content: center;
  .active{
    opacity: 0.7;
  }
  .mailforget-wrapper{
    margin: 0;
    width: 90%;
    height: 90%;
    border-radius: 0.08rem;
  display: flex;
  align-items: center;
  justify-content: center;

    // max-width: 540px;
    background-color: #fff;
    padding: 56px 60px;
    box-sizing: border-box;
    .center_flex{

    }
    .mailforget-title{
      width: 100%;
      text-align: center;
      font-size: 0.2743rem;
      color: #000;
      line-height: 28px;
    }
    .mailforget-content{
      width: 80%;
      color: #000;
      font-size: 0.1972rem;
      line-height: 1.2;
      margin: 0.6rem auto 0;
      span{
        margin:  0 5px;
        color: #0080F7;
      }
    }
    // .form-btn:hover{
    //   background-color: $color_hover;
    // }
    .form-btn{
      line-height: 48px;
      width: 3rem;
      height: 48px;
      background-color: #0b318f;
      color: #fff;
      font-size: 0.2229rem;
      font-weight: normal;
      text-align: center;
      margin: 0.6rem auto 0;
      cursor: pointer;
      border-radius: 0.08rem;
    }
  }
}
@media screen and (max-width: 640px){
  .mailforget-page .mailforget-wrapper .mailforget-content{
      width: 100%;
  }
}
</style>
