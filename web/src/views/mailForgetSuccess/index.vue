<template>
  <div class="mailforget-success-page">
    <!-- 邮箱重置密码邮件验证成功 -->
    <div class="mailforget-success-wrapper" v-loading="mailforgetSuccessLoad">

      <img :class="mailforgetSuccessLoad === true ? 'loading' : ''" class="mailforget-success-icon" src="../../assets/images/validation_icon_successful.png"/>
      <div :class="mailforgetSuccessLoad === true ? 'loading' : ''" class="mailforget-success-content">{{generateMailForgetSuccess('mailForgetSuccess_title')}}<span class="mailforget-success-content-tips">（<span class="mailforget-success-content-tipstime">{{jumpTime}}</span>{{generateMailForgetSuccess('mailForgetSuccess_tips')}}）</span></div>

    </div>
  </div>
</template>

<script>
// import * as myAjax from '@/api/mailForgetSuccess'
import { generateMailForgetSuccess } from '@/utils/i18n'
export default {
  name: 'mailForgetSuccess',
  data() {
    return {
      // 激活框加载遮罩
      mailforgetSuccessLoad: true,
      // 激活邮件地址
      mail: '',
      // 倒计时
      jumpTime: ''
    }
  },
  methods: {
    generateMailForgetSuccess,
    // 判断是否有重置邮箱
    haveCode() {
      var _this = this
      if (sessionStorage.oaxResetMail) {
        _this.checkCode()
      } else {
        _this.$router.push({ path: '/forget' })
      }
    },
    // 验证邮箱验证码
    checkCode(data) {
      var _this = this
      _this.mailforgetSuccessLoad = false
      var jump_time = 3
      _this.jumpTime = jump_time + _this.generateMailForgetSuccess('mailForgetSuccess_jumptime_unit')
      var jump_time_inter = setInterval(function() {
        if (jump_time > 0) {
          jump_time -= 1
          _this.jumpTime = jump_time + _this.generateMailForgetSuccess('mailForgetSuccess_jumptime_unit')
        } else {
          clearInterval(jump_time_inter)
          _this.$router.push({ path: '/login' })
        }
      }, 1000)
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
.mailforget-success-page{
  $color: #ffb933;
  $color_hover: #e0a32d;
  position: absolute;
  width: 100%;
  height: 100%;
  background-color: #e6e6e6;
  display: flex;
  align-items: center;
  justify-content: center;
  .mailforget-success-wrapper{
  
    margin: 0;
    width: 90%;
    height: 90%;    
    border-radius: 0.08rem;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;

    // max-width: 540px;
    background-color: #fff;
    padding: 56px 60px;
    box-sizing: border-box;
    .mailforget-success-icon.loading,
    .mailforget-success-content.loading{
      opacity: 0;
    }
    .mailforget-success-icon,
    .mailforget-success-content{
      transition: opacity .5s;
    }
    .mailforget-success-icon{
      display: block;
      width: 100px;
      height: 100px;
      margin: 0 auto;
    }
    .mailforget-success-content{
      width: 100%;
      text-align: center;
      color: #000;
      font-size: 24px;
      margin-top: 30px;
      .mailforget-success-content-tips{
        color: #999;
        .mailforget-success-content-tipstime{
          color: #0080F6;
        }
      }
    }
  }
}
</style>
