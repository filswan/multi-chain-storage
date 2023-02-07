<template>
  <div class="sidebar">
    <div class="sidebar_close el-icon-close" v-if="collapseLocal&&bodyWidth" @click="collapseChage"></div>
    <el-menu class="sidebar-el-menu" :default-active="routerMenu" :collapse="false" background-color="#080B29" text-color="#B3C0E7" active-text-color="#fff" unique-opened router>
      <template>
        <template>
          <div class="fes-menu">
            <div class="header_logo pcShow">
              <div class="logo"><img src="@/assets/images/LOGO_MCS@2x.png"></div>
              <img class="beta" src="@/assets/images/landing/beta.png">
<!--              <div class="beta"><h2>FEVM</h2></div>-->
            </div>
            <div class="menu_list">
              <el-menu-item v-for="(item, i) in items" :key="i" :index="item.index" @click="sidebarLiIndex(item.name, item.index, item.type)">
                <i :class="item.icon" style="font-size:15px"></i>
                <span slot="title">{{ item.title }}</span>
              </el-menu-item>
<!--              <el-menu-item @click="documentLink">-->
<!--                <i class="el-icon-s-documentation" style="font-size:15px; "></i>-->
<!--                <span slot="title">{{$t('route.documentation')}}</span>-->
<!--              </el-menu-item>-->
<!--              <el-menu-item class="mobileShow" v-if="metaAddress">-->
<!--                <i class="el-icon-switch-button" style="font-size:15px; "></i>-->
<!--                <span slot="title" @click="signOutFun">{{$t('fs3.Disconnect')}}</span>-->
<!--              </el-menu-item>-->
            </div>
          </div>
          <div class="fes-icon">
            <div class="comment" v-if="false">
              <div class="title">
                <i class="icon"></i> {{$t('comment.Rate_Product')}}
              </div>
              <div class="rate">
                <img src="@/assets/images/space/icon_20.png" class="img" />
                <el-rate v-model="rateValue" :icon-classes="iconClasses" void-icon-class="icon-rate-face-off" void-color="#fff" :colors="colors"></el-rate>
              </div>
              <el-button round>{{$t('comment.Tell_Comment')}}</el-button>
            </div>
            <div class="need">Need help? Join our
              <a href="https://discord.com/invite/KKGhy8ZqzK" target="_blank">Discord</a>
              or send an
              <a href="mailto:team@filswan.com">Email</a> to us.
            </div>
            <div class="progress">
              <el-progress :percentage="(free_bucket/free_bucketAll)*100 || 0"></el-progress>
              <span v-if="languageMcs === 'en'" class="tip">{{free_bucket | byteStorage}}GB of {{free_bucketAll | byteStorage}}GB for Bucket storage</span>
              <span v-else class="tip">目前使用量：{{free_bucket | byteStorage}}GB（Bucket储存空间配额：{{free_bucketAll | byteStorage}}GB）</span>
            </div>
            <!-- <div class="progress">
              <el-progress :percentage="(free_usage/free_quota_per_month)*100 || 0"></el-progress>
              <span v-if="languageMcs === 'en'" class="tip">{{free_usage | byteStorage}}GB of {{free_quota_per_month | byteStorage}}GB for Onchain Storage free storage</span>
              <span v-else class="tip">目前使用量：{{free_usage | byteStorage}}GB（Onchain Storage免费储存空间配额：{{free_quota_per_month | byteStorage}}GB）</span>
            </div> -->
            <div class="fes-icon-logo">
              <a href="https://filswan.medium.com/" target="_blank"><img :src="share_medium" alt=""></a>
              <a href="https://discord.com/invite/KKGhy8ZqzK" target="_blank"><img :src="share_discord" alt=""></a>
              <a href="https://twitter.com/0xfilswan" target="_blank"><img :src="share_twitter" alt=""></a>
              <a href="https://github.com/filswan" target="_blank"><img :src="share_github" alt=""></a>
              <a href="https://t.me/filswan" target="_blank"><img :src="share_telegram" alt=""></a>
            </div>
            <!-- market data white badge -->
            <div class="fes-market">
              <a href="https://chn.lk/3DTWSjE" target="_blank"><img src="https://chain.link/badge-market-data-white" alt="market data secured with chainlink"></a>
            </div>
          </div>
        </template>
      </template>
    </el-menu>

    <!-- <div class="copyStyle" v-if="!collapseLocal">{{ $t('navbar.copy') }} {{git_version}}</div> -->
  </div>
</template>

<script>
// import bus from './bus';
import axios from 'axios'
export default {
  data () {
    return {
      git_version: null,
      collapseLocal: !!(this.$store.getters.collapseL === 'true' || this.$store.getters.collapseL === true),
      lanShow: false,
      bodyWidth: document.body.clientWidth < 999,
      items: [
        // {
        //   icon: 'el-icon-s-deal',
        //   index: '1',
        //   title: this.$t('route.Deal'),
        //   name: 'my_files',
        //   type: ''
        // },
        {
          icon: 'el-icon-s-metaSpace',
          index: '20',
          title: this.$t('route.metaSpace'),
          name: 'Space',
          type: ''
        }
        // {
        //   icon: 'el-icon-s-billing',
        //   index: '5',
        //   title: this.$t('navbar.BillingHistory'),
        //   name: 'billing',
        //   type: ''
        // },
        // {
        //   icon: 'el-icon-s-myAccount',
        //   index: '21',
        //   title: this.$t('route.myAccount'),
        //   name: 'ApiKey',
        //   type: ''
        // }
      ],
      share_medium: require('@/assets/images/landing/medium.png'),
      share_twitter: require('@/assets/images/landing/twitter.png'),
      share_github: require('@/assets/images/landing/github-fill.png'),
      share_telegram: require('@/assets/images/landing/telegram.png'),
      share_discord: require('@/assets/images/landing/discord.png'),
      rateValue: null,
      colors: { 5: '#e92721' },
      iconClasses: ['icon-rate-face']
    }
  },
  computed: {
    routerMenu () {
      return this.$store.getters.routerMenu.toString()
    },
    languageMcs () {
      return this.$store.getters.languageMcs
    },
    collapseL () {
      return this.$store.getters.collapseL
    },
    metaAddress () {
      return this.$store.getters.metaAddress
    },
    free_usage () {
      let freeUsage = this.$store.getters.free_usage < this.$store.getters.free_quota_per_month ? this.$store.getters.free_usage : this.$store.getters.free_quota_per_month
      return freeUsage
    },
    free_quota_per_month () {
      return this.$store.getters.free_quota_per_month
    },
    free_bucket () {
      let freeBucket = this.$store.getters.free_bucket < this.$store.getters.free_bucketAll ? this.$store.getters.free_bucket : this.$store.getters.free_bucketAll
      return freeBucket
    },
    free_bucketAll () {
      return this.$store.getters.free_bucketAll === 0 ? 34359738368 : this.$store.getters.free_bucketAll
    }
  },
  watch: {
    'collapseL': function () {
      this.collapseLocal = !!(this.$store.getters.collapseL === 'true' || this.$store.getters.collapseL === true)
      this.bodyWidth = document.body.clientWidth < 992
    }
  },
  created () {
    if (process.env.COMMITHASH) {
      this.git_version = process.env.COMMITHASH.slice(0, 8)
    }
    this.getListBuckets()
  },
  methods: {
    documentLink () {
      window.open('https://docs.filswan.com/multi-chain-storage/overview', '_blank')
    },
    sidebarLiIndex (nameNow, index, typeNow) {
      let _this = this
      let indexNow = Number(index)
      sessionStorage.removeItem('dealsPaginationIndexDev')
      switch (indexNow) {
        case 2:
          localStorage.removeItem('tabTask_name')
          localStorage.removeItem('tabTask_search')
          localStorage.removeItem('tabTaskMiner_search')
          break
        case 4:
          localStorage.removeItem('myProfileActive')
          break
        default:
      }
      _this.$store.dispatch('setRouterMenu', Number(index))
      if (typeNow) {
        _this.$router.push({ name: nameNow, params: { type: typeNow } })
        return false
      } else {
        _this.$router.push({ name: nameNow })
      }
    },
    collapseChage () {
      this.collapseLocal = !this.collapseLocal
      this.$store.dispatch('setCollapse', this.collapseLocal)
      // bus.$emit('collapse', this.collapseLocal);
      // this.$emit('collapseVisit', this.collapseLocal);
    },
    handleSetLanguage (lang) {
      let _this = this
      _this.$i18n.locale = lang
      _this.$store.dispatch('setLanguage', lang)
      window.location.reload()
    },
    signOutFun () {
      let _this = this
      let params = {}
      axios.post(`${_this.baseAPIURL}api/v1/user/logout_for_metamask_signature`, params, {
        headers: {
          'Authorization': 'Bearer ' + _this.$store.getters.mcsjwtToken
        }
      }).then((response) => {
        if (response.data.status !== 'success') _this.$message.error(response.data.message || 'Fail')
        _this.$store.dispatch('setMetaAddress', '')
        _this.$store.dispatch('setMCSjwtToken', '')
        _this.$store.dispatch('setMetaNetworkId', 0)
        _this.$store.dispatch('setMetaNetworkInfo', JSON.stringify({}))
        sessionStorage.removeItem('login_path')
        // _this.$router.push('/home')
        setTimeout(function () { window.location.reload() }, 200)
      }).catch(function (error) {
        console.log(error.config)
      })
    },
    async getListBuckets (name) {
      let _this = this
      let size = 0
      let maxSize = 0
      const directoryRes = await _this.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/bucket/get_bucket_list`, 'get')
      if (!directoryRes || directoryRes.status !== 'success') return false
      directoryRes.data.forEach(element => {
        size += element.size
        if (element.is_active) maxSize += element.max_size
      })
      await _this.$commonFun.timeout(500)
      _this.$store.dispatch('setFreeBucket', size || 0)
      _this.$store.dispatch('setFreeBucketAll', maxSize || 0)
    }
  },
  filters: {
    byteStorage (limit) {
      // 只转换成GB
      if (limit <= 0) {
        return '0'
      } else if (limit <= 10737419) { // 10737419 ~= 0.01GB
        return '0.01'
      } else {
        // return (limit/( 1024 * 1024 * 1024)).toPrecision(2)  //or 1000
        let value = limit / (1024 * 1024 * 1024)
        let v1 = String(value).split('.')
        let v2 = v1[1] || ''
        let v3 = String(v2).replace(/(0+)\b/gi, '')
        if (v3) {
          return v1[0] + '.' + v3.slice(0, 2)
        } else {
          return v1[0]
        }
      }
    }
  }
}
</script>

<style scoped lang="scss">
.sidebar {
  display: block;
  height: 100%;
  z-index: 9999;
  transition: all 0.3s;
  background: #080b29;
  .mobileShow {
    display: none;
  }
  .sidebar_close {
    position: absolute;
    right: 5px;
    top: 5px;
    font-size: 20px;
    font-weight: 600;
    color: #f2f3f8;
    cursor: pointer;
    z-index: 999999;
    display: block;
    padding: 0.05rem;
  }
}
.sidebar-el-menu {
  // width: 0.65rem;
  // height: calc(100% - 0.31rem);
  height: 100%;
  padding: 0;
  border-right: 0;
  // transition: width .3s;
  .menu_list {
    // height: calc(100% - 1.2rem);
    margin: 0 0.2rem;
    padding: 0;
    overflow-x: hidden;
    overflow-y: scroll;
    // 谷歌滚动条
    &::-webkit-scrollbar-track {
      background: transparent;
    }
    &::-webkit-scrollbar {
      width: 2px;
      background: transparent;
    }
    &::-webkit-scrollbar-thumb {
      background: #ccc;
    }
  }
  .header_logo {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    width: calc(100% - 0.3rem);
    padding: 0.41rem 0 0.62rem;
    background-color: #080b29;
    // transition: width .3s;
    .logo {
      width: 2rem;
      img {
        display: block;
        width: 100%;
        height: auto;
      }
    }
    .beta {
      position: absolute;
      width: 60px;
      right: 0;
      top: 0.75rem;
    }
    .header_btn {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      width: 0.26rem;
      height: 0.26rem;
      margin: 0 0 0 0.06rem;
      transition: all 0.4s ease;
      outline: none;
    }
    .header_btn span {
      position: relative;
      display: block;
      width: 100%;
      height: 1px;
      margin: auto;
      background-color: #fff;
      transition: all 0.4s ease;
    }
    .header_btn span::after {
      content: "";
      position: absolute;
      top: -7px;
      right: 0;
      width: 50%;
      height: 1px;
      background-color: #fff;
      transition: all 0.4s ease;
    }
    .header_btn span::before {
      content: "";
      position: absolute;
      bottom: -7px;
      right: 0;
      width: 75%;
      height: 1px;
      background-color: #fff;
      transition: all 0.4s ease;
    }
    .header_btn.header_btn_left span::before,
    .header_btn.header_btn_left span::after {
      right: auto;
      left: 0;
    }
    .header_btn:hover span,
    .header_btn:hover span::before,
    .header_btn:hover span::after {
      width: 100%;
      background-color: #fff;
    }
  }
  .header_left_hidd {
    width: 100%;
    height: 0.9rem;
    padding: 0.2rem 0 0.1rem;
    flex-wrap: wrap;
    justify-content: center;
    .logo_small {
      display: block;
      width: 100%;
      height: 0.34rem;
      margin: auto;
      img {
        display: block;
        height: 100%;
        max-width: 100%;
        margin: auto;
      }
    }
    .collapse-btn-cont {
      transform: rotate(180deg);
      .header_btn {
        margin: auto;
      }
    }
  }
  .fes-menu {
    // rate 150px
    height: calc(100% - 160px);
    overflow-y: scroll;
    @media screen and (min-width: 1800px) {
      // height: calc(100% - 285px); // rate 140px
      height: calc(100% - 175px);
    }
    @media screen and (max-width: 1366px) {
      // height: calc(100% - 290px); // rate 140px
      height: calc(100% - 180px);
    }
    &::-webkit-scrollbar-track {
      background: transparent;
    }
    &::-webkit-scrollbar {
      width: 6px;
      background: transparent;
    }
    &::-webkit-scrollbar-thumb {
      background: #080b29;
    }
  }
  .fes-icon {
    position: absolute;
    left: 0;
    right: 0;
    bottom: 0;
    padding: 5px 0;
    .comment {
      padding: 10px 17px;
      margin: 0 0.28rem 10px 0.2rem;
      background: linear-gradient(45deg, #4f8aff, #4b5eff);
      border-radius: 0.06rem;
      color: #fff;
      .title {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        width: 100%;
        font-size: 0.15rem;
        font-weight: normal;
        color: #fff;
        line-height: 15px;
        @media screen and (min-width: 1800px) {
          font-size: 15px;
        }
        @media screen and (max-width: 441px) {
          font-size: 13px;
        }
        .icon {
          width: 13px;
          height: 13px;
          margin: 0 5px 0 0;
          background: url(../assets/images/space/icon_19.png) no-repeat center;
          background-size: 100%;
        }
      }
      .rate {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        align-items: center;
        padding: 10px 0;
        .img {
          display: block;
          height: 40px;
          margin: 0 15px 0 0;
          @media screen and (max-width: 1600px) {
            margin: 0 5px;
          }
        }
        .el-rate /deep/ {
          display: flex;
          flex-wrap: wrap;
          .el-rate__item {
            .el-rate__icon,
            .icon-rate-face,
            .hover {
              width: 22px;
              height: 22px;
              margin: 0 5px;
              background: url(../assets/images/space/start_on.png) no-repeat
                center;
              background-size: 100%;
              @media screen and (max-width: 1600px) {
                width: 18px;
                height: 18px;
                margin: 0 3px;
              }
            }
            .icon-rate-face-off {
              width: 22px;
              height: 22px;
              background: url(../assets/images/space/start.png) no-repeat center;
              background-size: 100%;
              @media screen and (max-width: 1600px) {
                width: 18px;
                height: 18px;
                margin: 0 3px;
              }
            }
          }
        }
      }
      .el-button {
        display: flex;
        align-items: center;
        justify-content: center;
        width: 100%;
        height: 35px;
        padding: 0;
        background-color: #fff;
        color: #4e80ff;
        font-family: inherit;
        font-size: 0.15rem;
        font-weight: normal;
        line-height: 1.2;
        @media screen and (min-width: 1800px) {
          height: 40px;
        }
        @media screen and (max-width: 441px) {
          height: 30px;
          font-size: 13px;
        }
      }
    }
    .need {
      min-height: 24px;
      padding: 0 0.28rem;
      font-size: 12px;
      color: #fff;
      text-align: center;
      line-height: 1.1;
      @media screen and (min-width: 1800px) {
        font-size: 14px;
      }
      a {
        color: inherit;
        text-decoration: underline;
      }
    }
    .progress {
      margin: 2px 0.28rem 10px;
      .el-progress /deep/ {
        font-size: 12px;
        .el-progress-bar {
          padding-right: 0;
          margin: 0 0 5px 0;
        }
        .el-progress__text {
          display: none;
        }
      }
      .tip {
        display: block;
        font-size: 12px;
        line-height: 1.2;
        color: rgb(179, 192, 231);
        @media screen and (min-width: 1800px) {
          font-size: 14px;
        }
      }
    }
    .fes-icon-logo {
      display: flex;
      justify-content: center;
      align-items: center;
      img {
        display: block;
        height: 22px;
        margin: 0 0.15rem;
        @media screen and (max-width: 1024px) {
          height: 20px;
        }
      }
    }
    .fes-market {
      display: flex;
      justify-content: center;
      img {
        display: block;
        height: 35px;
        margin: 10px auto 5px;
        @media screen and (min-width: 1800px) {
          height: 45px;
        }
      }
    }
  }
  .collapse-btn-cont {
    float: left;
    padding: 0;
    cursor: pointer;
    align-items: center;
    display: flex;
  }
  li {
    display: flex;
    height: 0.5rem;
    padding: 0 15px;
    margin: 0.12rem 0;
    font-size: 0.2rem;
    font-weight: 500;
    -webkit-box-align: center;
    -ms-flex-align: center;
    align-items: center;
    line-height: 0.5rem;
    @media screen and (max-width: 1200px) {
      padding: 0 10px;
    }
    div {
      display: flex;
      align-items: center;
      justify-content: center;
    }
    i {
      position: relative;
      height: 0.2rem;
      margin-right: 0.22rem;
      &::before {
        position: absolute;
        left: 0;
        top: 0;
        width: 100%;
        height: 100%;
        content: "";
        background-size: 0.2rem !important;
      }
    }
    .el-icon-switch-button {
      &::before {
        content: "\E71B";
      }
    }
    .el-icon-s-home {
      &::before {
        background: url(../assets/images/menuIcon/Dashboard.png) no-repeat
          center;
      }
    }
    .el-icon-s-upload {
      &::before {
        background: url(../assets/images/menuIcon/uploadFile.png) no-repeat
          center;
      }
    }
    .el-icon-s-browse {
      &::before {
        background: url(../assets/images/menuIcon/browse.png) no-repeat center;
      }
    }
    .el-icon-s-deal {
      &::before {
        background: url(../assets/images/menuIcon/icon_Files@2x.png) no-repeat
          center;
      }
    }
    .el-icon-s-myFs3 {
      &::before {
        background: url(../assets/images/menuIcon/S3.png) no-repeat center;
      }
    }
    .el-icon-s-myProfile {
      &::before {
        background: url(../assets/images/menuIcon/myProfile.png) no-repeat
          center;
      }
    }
    .el-icon-s-tools {
      &::before {
        background: url(../assets/images/menuIcon/tool.png) no-repeat center;
      }
    }
    .el-icon-s-Stats {
      &::before {
        background: url(../assets/images/menuIcon/icon_Statistics@2x.png)
          no-repeat center;
      }
    }
    .el-icon-s-ApiKey {
      &::before {
        background: url(../assets/images/menuIcon/icon_ApiKey.png) no-repeat
          center;
      }
    }
    .el-icon-s-documentation {
      &::before {
        background: url(../assets/images/menuIcon/icon_documentation@2x.png)
          no-repeat center;
      }
    }
    .el-icon-s-metaSpace {
      &::before {
        background: url(../assets/images/menuIcon/metaSpace.png) no-repeat
          center;
      }
    }
    .el-icon-language {
      &::before {
        background: url(../assets/images/menuIcon/language.png) no-repeat center;
        background-size: 100% !important;
      }
    }
    .el-icon-s-dataset {
      &::before {
        background: url(../assets/images/menuIcon/dataset.png) no-repeat center;
      }
    }
    .el-icon-search {
      &::before {
        background: url(../assets/images/menuIcon/Search-File.png) no-repeat
          center;
      }
    }
    .el-icon-s-myAccount {
      &::before {
        background: url(../assets/images/menuIcon/myAccount.png) no-repeat
          center;
      }
    }
    .el-icon-s-billing {
      &::before {
        background: url(../assets/images/menuIcon/billing@2x.png) no-repeat
          center;
      }
    }
  }

  .el-menu-item:hover,
  .is-active {
    background: linear-gradient(45deg, #4f8aff, #4b5eff);
    border-radius: 0.06rem;
    color: #fff !important;
    .el-icon-s-home {
      &::before {
        background: url(../assets/images/menuIcon/Dashboard-1.png) no-repeat
          center;
      }
    }
    .el-icon-s-upload {
      &::before {
        background: url(../assets/images/menuIcon/uploadFile-1.png) no-repeat
          center;
      }
    }
    .el-icon-s-browse {
      &::before {
        background: url(../assets/images/menuIcon/browse-1.png) no-repeat center;
      }
    }
    .el-icon-s-deal {
      &::before {
        background: url(../assets/images/menuIcon/icon_Files@2x-1.png) no-repeat
          center;
      }
    }
    .el-icon-s-myFs3 {
      &::before {
        background: url(../assets/images/menuIcon/S3-1.png) no-repeat center;
      }
    }
    .el-icon-s-myProfile {
      &::before {
        background: url(../assets/images/menuIcon/myProfile-1.png) no-repeat
          center;
      }
    }
    .el-icon-s-tools {
      &::before {
        background: url(../assets/images/menuIcon/tool-1.png) no-repeat center;
      }
    }
    .el-icon-s-Stats {
      &::before {
        background: url(../assets/images/menuIcon/icon_Statistics@2x-1.png)
          no-repeat center;
      }
    }
    .el-icon-s-ApiKey {
      &::before {
        background: url(../assets/images/menuIcon/icon_ApiKey-1.png) no-repeat
          center;
      }
    }
    .el-icon-s-documentation {
      &::before {
        background: url(../assets/images/menuIcon/icon_documentation@2x-1.png)
          no-repeat center;
      }
    }
    .el-icon-s-metaSpace {
      &::before {
        background: url(../assets/images/menuIcon/metaSpace-1.png) no-repeat
          center;
      }
    }
    .el-icon-language {
      &::before {
        background: url(../assets/images/menuIcon/language-1.png) no-repeat
          center;
        background-size: 100% !important;
      }
    }
    .el-icon-s-dataset {
      &::before {
        background: url(../assets/images/menuIcon/dataset-1.png) no-repeat
          center;
      }
    }
    .el-icon-search {
      &::before {
        background: url(../assets/images/menuIcon/Search-File-1.png) no-repeat
          center;
      }
    }
    .el-icon-s-myAccount {
      &::before {
        background: url(../assets/images/menuIcon/myAccount-1.png) no-repeat
          center;
      }
    }
    .el-icon-s-billing {
      &::before {
        background: url(../assets/images/menuIcon/billing@2x-1.png) no-repeat
          center;
      }
    }
  }
  .el-menu-item span {
    position: relative;
    /* transition: all 0.4s ease; */
  }
  .el-menu-item span::after {
    content: "";
    position: absolute;
    left: 50%;
    width: 0;
    bottom: 4px;
    margin: auto;
    height: 2px;
    background-color: #080b29;
    transition: all 0.4s ease;
  }
  .el-menu-item span:hover::after {
    left: 0;
    width: 100%;
  }
  &:not(.el-menu--collapse) {
    li {
      width: auto;
    }
  }
}
.copyStyle {
  padding: 0 0 0 20px;
  background: #080b29;
  font-size: 0.12rem;
  line-height: 0.3rem;
  color: #9c9c9c;
  border-top: 0.01rem solid rgba(255, 255, 255, 0.2);
  z-index: 999;
  @media screen and (max-width: 1200px) {
    padding: 0 0 0 10px;
  }
}

@media screen and (max-width: 991px) {
  .sidebar {
    transition: all 0.3s ease;
    .mobileShow {
      display: flex;
      &:hover {
        i {
          color: #fff;
        }
      }
    }
    .sidebar-el-menu {
      width: 100%;
      .menu_list {
        // height: calc(100% - 0.5rem);
        padding: 50px 0 0;
      }
    }
    .pcShow {
      display: none;
    }
  }
}
</style>
