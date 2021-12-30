<template>
  <div class="activation-page">
    <!-- 发送激活邮件 -->
    <div class="activation-wrapper" v-loading="activationLoad">
      <div class="center_flex">
        <!-- 标题 -->
        <div class="activation-title">{{generateAccountActivation('accountActivation_title')}}</div>

        <!-- 发送提示 -->
        <div class="activation-content">{{generateAccountActivation('accountActivation_tips_1')}}<span>{{mail}}</span>，{{generateAccountActivation('accountActivation_tips_2')}}</div>

        <!-- 重发按钮 -->
        <div class="form-btn" @click="sendAgain" :class="isGetVerifyCode === true ? 'active' : ''">{{getVerifyCodeWord}}</div>
      </div>
    </div>

    <div id="captcha"></div>

  </div>
</template>

<script>
import * as myAjax from '@/api/accountActivation'
import { generateAccountActivation } from '@/utils/i18n'
export default {
  name: 'accountActivation',
  data() {
    return {
      // 滑块验证弹框是否显示
      dialogVerifyShow: false,
      // 激活框加载遮罩
      activationLoad: false,
      // 激活邮件地址
      mail: '',
      // 上次发送激活邮件的时间戳
      lastTime: -1,
      // 是否在发送激活邮件
      isGetVerifyCode: false,
      // 发送激活邮件按钮文字
      getVerifyCodeWord: this.generateAccountActivation('accountActivation_resend'),
    }
  },
  methods: {
    generateAccountActivation,
    // 判断是否有邮箱对象
    haveMail() {
      var _this = this
      if (!sessionStorage.oaxRegisterMail) {
        _this.$router.push({ path: '/register' })
      } else {
        _this.mail = sessionStorage.oaxRegisterMail
      }
    },
    // 重新发送邮件
    sendAgain() {
      var _this = this
      if (_this.isGetVerifyCode === true) {
        return false
      }
      _this.activationLoad = true
      var newData = {
        email: _this.mail,
        source: 2
      }
      myAjax
        .sendActivateLink(newData)
        .then(response => {
          console.log(response)
          _this.activationLoad = false
          if (response.status == 'success') {
            _this.$message({
              message: response.message,
              type: 'success'
            })
            _this.lastTime = new Date().getTime()
            _this.reSendTimer()
          } else {
            _this.$message.error(response.message)
          }
        })
        .catch(error => {
          console.log(error)
          _this.activationLoad = false
        })
    },
    // 重新发送邮件按钮倒计时
    reSendTimer() {
      var _this = this
      _this.isGetVerifyCode = true
      var get_code_time = 30 - Math.floor((new Date().getTime() - _this.lastTime) / 1000)
      _this.getVerifyCodeWord = get_code_time + _this.generateAccountActivation('getVerifyCodeWord_time')
      var code_times = setInterval(function() {
        if (get_code_time <= 0) {
          clearInterval(code_times)
          _this.isGetVerifyCode = false
          _this.getVerifyCodeWord = _this.generateAccountActivation('accountActivation_resend')
        } else {
          get_code_time -= 1
          _this.getVerifyCodeWord = get_code_time + _this.generateAccountActivation('getVerifyCodeWord_time')
        }
      }, 1000)
    }
  },

  mounted() {
    var _this = this
    _this.haveMail()
    if (sessionStorage.oaxRegisterMailTime && (new Date().getTime() - parseInt(sessionStorage.oaxRegisterMailTime)  >= 0) && (new Date().getTime() - parseInt(sessionStorage.oaxRegisterMailTime)  <= 30000)) {
      _this.lastTime = parseInt(sessionStorage.oaxRegisterMailTime)
      _this.sendAgain()
    } else if ((new Date().getTime() - parseInt(sessionStorage.oaxRegisterMailTime)  < 0) || (new Date().getTime() - parseInt(sessionStorage.oaxRegisterMailTime)  > 30000)) {
      sessionStorage.removeItem('oaxRegisterMailTime')
      _this.lastTime = -1
    }
  },
  components: {
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scopte>
.activation-page{
  $color: #ffb933;
  $color_hover: #e0a32d;
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: #e6e6e6;
  display: flex;
  align-items: center;
  justify-content: center;
  .activation-wrapper{
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
    .activation-title{
      width: 100%;
      text-align: center;
      font-size: 0.2743rem;
      color: #000;
      line-height: 28px;
    }
    .activation-content{
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
@media screen and (max-width:999px){
  .activation-page .activation-wrapper .activation-content{
    width: 100%
  }
}
</style>
