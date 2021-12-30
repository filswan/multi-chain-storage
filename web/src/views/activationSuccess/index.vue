<template>
  <div class="activation-success-page">
    <div class="activation-success-wrapper" v-loading="activationSuccessLoad">
      <!-- 激活验证成功 -->
      <div v-if="activationResult === 1">
        <img class="activation-success-icon" src="../../assets/images/validation_icon_successful.png"/>
        <div class="activation-success-content">{{generateActivationSuccess('activationSuccess_success_title')}}<span class="activation-success-content-tips">（<span class="activation-success-content-tipstime">{{jumpTime}}</span>{{generateActivationSuccess('activationSuccess_success_tips')}}）</span></div>
      </div>

      <!-- 验证失败 -->
      <div v-else-if="activationResult === 2">
        <img class="activation-success-icon" src="../../assets/images/icon_failure.png"/>
        <div class="activation-success-content">{{generateActivationSuccess('activationSuccess_fail_title')}}</div>        
        <div class="activation-success-tips" @click="againCode">{{generateActivationSuccess('activationSuccess_fail_tips01')}}</div>
      </div>

    </div>
  </div>
</template>

<script>
import * as myAjax from '@/api/accountActivation'
import { generateActivationSuccess } from '@/utils/i18n'
export default {
  name: 'activationSuccess',
  data() {
    return {
      // 激活框加载遮罩
      activationSuccessLoad: true,
      // 激活邮件地址
      mail: '',
      // 倒计时
      jumpTime: '',
      // 验证结果
      activationResult: 0
    }
  },
  methods: {
    generateActivationSuccess,
    // 判断是否有地址栏验证参数
    haveCode() {
      var _this = this
      if (_this.getUrlVars().activateCode && _this.getUrlVars().email) {
        _this.checkCode(_this.getUrlVars().activateCode, _this.getUrlVars().email)
      } else {
        _this.$router.push({ path: '/register' })
      }
    },
    // 验证邮箱验证码
    checkCode(code, email) {
      var _this = this
      let param = {
        activateCode: code,
        email: email
      }

      myAjax
        .checkCode(param)
        .then(response => {
          console.log(response)
          if (response.status == "success") {
            _this.activationSuccessLoad = false
            _this.activationResult = 1
            var jump_time = 3
            _this.jumpTime = jump_time + _this.generateActivationSuccess('activationSuccess_success_jumptime_unit')
            var jump_time_inter = setInterval(function() {
              if (jump_time > 0) {
                jump_time -= 1
                _this.jumpTime = jump_time + _this.generateActivationSuccess('activationSuccess_success_jumptime_unit')
              } else {
                clearInterval(jump_time_inter)
                sessionStorage.oaxLoginType = 'mail'
                _this.$router.push({ path: '/login' })
              }
            }, 1000)
          } else {
            _this.activationSuccessLoad = false
            _this.activationResult = 2
          }
        })
        .catch(error => {
          console.log(error)
        })
    },
    // 获取地址栏参数
    getUrlVars() {
      var vars = []
      var hash
      var hashes = window.location.href.slice(window.location.href.indexOf('?') + 1).split('&')
      for (var i = 0; i < hashes.length; i++) {
        hash = hashes[i].split('=')
        vars.push(hash[0])
        vars[hash[0]] = hash[1]
      }
      return vars
    },
    // 再次获取激活链接
    againCode() {
      sessionStorage.oaxRegisterMail = this.getUrlVars().email
      sessionStorage.oaxRegisterMailTime = new Date().getTime()
      this.$router.push('/account_activation');
    }
  },

  mounted() {
    this.haveCode()
  },
  components: {
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scopte>
.activation-success-page{
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
  .activation-success-wrapper{
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
    .activation-success-icon.loading,
    .activation-success-content.loading{
      opacity: 0;
    }
    .activation-success-icon,
    .activation-success-content{
      transition: opacity .5s;
    }
    .activation-success-icon{
      display: block;
      width: 100px;
      height: 100px;
      margin: 0 auto;
    }
    .activation-success-content{
      width: 100%;
      color: #000;
      font-size: 0.1972rem;
      line-height: 1.2;
      margin: 0.6rem auto 0;
      span{
        margin:  0 5px;
        color: #0080F7;
      }
    }
    .activation-success-tips{
      width: 100%;
      text-align: center;
      color: #31a0ff;
      font-size: 12px;
      line-height: 15px;
      margin-top: 15px;
      text-decoration: underline;
      cursor: pointer;
    }
    .form-btn{
      line-height: 0.4rem;
      width: auto;
      height: 0.4rem;
      background-color: #0b318f;
      color: #fff;
      font-size: 0.14rem;
      font-weight: normal;
      text-align: center;
      margin: 0.1rem auto 0;
      padding: 0 0.1rem;
      cursor: pointer;
      border-radius: 0.08rem;
      opacity: 0.8;
      &:hover{
        opacity: 1;
      }
    }
  }
}
</style>
