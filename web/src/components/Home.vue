<template>
  <div class="wrapper">
    <!-- <el-col :xs="8" :sm="8" :md="5" :lg="4" class="side" :class="{'slidarShow': collapseP&&bodyWidth}">
      <v-sidebar></v-sidebar>
    </el-col> -->
    <el-col :xs="24" :sm="24" :md="24" :lg="24" class="content-box" id="content-box">
      <v-head :meta="meta" @getMetamaskLogin="getMetamaskLogin" :netId="netId" :networkTip="networkTip" @getNetwork="getNetwork" @getNetId="changeNet" @getPopUps="getPopUps"></v-head>
      <div id="headerMb" v-if="bodyWidth">
        <div class="headerMb" v-if="email">
          {{headertitle}}
        </div>
      </div>
      <div class="content" id="content_client">
        <div class="content_body mWidth" :class="{'stats': $route.name == 'Stats' || $route.name == 'Space' || $route.name == 'Space_detail'}">
          <network-alert v-if="metaAddress&&networkTip" @changeNet="changeNet" @getNetwork="getNetwork"></network-alert>
          <transition name="move" mode="out-in">
            <keep-alive :include="tagsList">
              <router-view @getMetamaskLogin="getMetamaskLogin"></router-view>
            </keep-alive>
          </transition>
        </div>
        <div class="fes-icon mWidth">
          <div class="fes-icon-logo">
            <a :href="medium_link" target="_blank"><img :src="share_medium" alt=""></a>
            <a :href="discord_link" target="_blank"><img :src="share_discord" alt=""></a>
            <a :href="twitter_link" target="_blank"><img :src="share_twitter" alt=""></a>
            <a :href="github_link" target="_blank"><img :src="share_github" alt=""></a>
            <a :href="telegram_link" target="_blank"><img :src="share_telegram" alt=""></a>
          </div>
          <div class="fes-icon-copy">
            <span>© {{fullYear}} FilSwan</span>
            <el-divider direction="vertical"></el-divider>
            <a href="https://www.filswan.com/" target="_block">filswan.com</a>
          </div>
          <div class="fes-market">
            <a href="https://chn.lk/3DTWSjE" target="_blank"><img src="@/assets/images/landing/chainlink-badge-market-data.svg" alt="market data secured with chainlink"></a>
          </div>
        </div>
        <el-backtop target=".content"></el-backtop>
      </div>
    </el-col>
    <pop-ups v-if="dialogFormVisible" :dialogFormVisible="dialogFormVisible" :typeModule="typeName" :changeTitle="changeTitle" @getPopUps="getPopUps"></pop-ups>
  </div>
</template>

<script>
import vHead from './Header.vue'
import vSidebar from './Sidebar.vue'
import popUps from './popups.vue'
import networkAlert from '@/components/networkAlert.vue'
let that
export default {
  data () {
    return {
      tagsList: [],
      collapse: false,
      collapseP: !!(this.$store.getters.collapseL === 'true' || this.$store.getters.collapseL === true),
      bodyWidth: document.body.clientWidth < 992,
      tabshow: localStorage.getItem('tabTask_name') === 'User_Profile',
      share_medium: require('@/assets/images/landing/medium.png'),
      share_twitter: require('@/assets/images/landing/twitter.png'),
      share_github: require('@/assets/images/landing/github-fill.png'),
      share_telegram: require('@/assets/images/landing/telegram.png'),
      share_discord: require('@/assets/images/landing/discord.png'),
      medium_link: process.env.MEDIUM_LINK,
      discord_link: process.env.DISCORD_LINK,
      twitter_link: process.env.TWITTER_LINK,
      github_link: process.env.GITHUB_LINK,
      telegram_link: process.env.TELEGRAM_LINK,
      meta: false,
      netId: 0,
      networkTip: false,
      dialogFormVisible: false,
      typeName: 'emailLogin',
      changeTitle: '',
      fullYear: new Date().getFullYear()
    }
  },
  components: {
    vHead,
    vSidebar,
    popUps,
    networkAlert
  },
  computed: {
    headertitle () {
      return this.$store.getters.headertitle
    },
    routerMenu () {
      return this.$store.getters.routerMenu
    },
    email () {
      return this.$store.state.user.email
    },
    collapseL () {
      return this.$store.getters.collapseL
    },
    metaAddress () {
      return this.$store.getters.metaAddress
    },
    networkID () {
      return Number(this.$store.getters.networkID)
    },
    metaNetworkInfo () {
      return this.$store.getters.metaNetworkInfo ? JSON.parse(JSON.stringify(this.$store.getters.metaNetworkInfo)) : {}
    },
    reverse () {
      return String(this.$store.getters.reverse) === '1'
    },
    mcsEmail () {
      const data = this.$store.getters.mcsEmail
      return data === '{}' ? '' : JSON.parse(data).email
    },
    apiStatus () {
      const data = this.$store.getters.mcsEmail
      return data === '{}' ? false : JSON.parse(data).apiStatus
    }
  },
  watch: {
    'collapseL': function () {
      this.collapseP = !!(this.$store.getters.collapseL === 'true' || this.$store.getters.collapseL === true)
      this.bodyWidth = document.body.clientWidth < 992
    },
    reverse: function () {
      this.init()
    }
  },
  methods: {
    async getPopUps (dialog, rows, bucketName) {
      that.dialogFormVisible = dialog
      that.changeTitle = rows || ''
    },
    getNetwork (dis) {
      that.networkTip = dis
    },
    changeNet (id) {
      that.netId = id
    },
    getMetamaskLogin (meta) {
      that.meta = meta
    },
    async init () {
      that.dialogFormVisible = !!that.metaAddress && !that.mcsEmail && that.apiStatus
      let status = await that.$metaLogin.netStatus(that.networkID)
      that.networkTip = !status
      if (that.reverse) document.body.classList.add('reverse_phase')
      else document.body.classList.remove('reverse_phase')
    }
  },
  mounted () {
    that = this
    that.init()
  }
}
</script>

<style lang="scss" scoped>
.wrapper {
  position: relative;
  display: flex;
  flex-wrap: wrap;
  width: 21.4rem;
  max-width: 100%;
  margin: auto;
  @media screen and (max-width: 1440px) {
    width: 20.6rem;
  }
  @media screen and (max-width: 1366px) {
    width: 21.1rem;
  }
  @media screen and (max-width: 1280px) {
    width: 21.4rem;
  }
  @media screen and (max-width: 1152px) {
    width: 21rem;
  }
  @media screen and (max-width: 1024px) {
    width: 20.48rem;
  }
  .side {
    height: 100%;
    // transition: all 0.3s ease;
    @media screen and (max-width: 991px) {
      position: fixed;
      left: -33.33333%;
      z-index: 9999;
    }
    @media screen and (max-width: 415px) {
      left: -60%;
      width: 60%;
    }
    &.slidarShow {
      left: 0;
    }
  }
  .content-box {
    height: 100%;
  }
  .content {
    position: relative;
    height: calc(100% - 1.1rem);
    overflow-y: scroll;
    // transition: all;
    // transition-duration: .3s;
    &::-webkit-scrollbar {
      width: 1px;
      height: 1px;
      background-color: #eee;
    }

    &::-webkit-scrollbar-track {
      box-shadow: none;
      -webkit-box-shadow: none;
      border-radius: 10px;
      background-color: #eee;
    }

    &::-webkit-scrollbar-thumb {
      border-radius: 10px;
      box-shadow: none;
      -webkit-box-shadow: none;
      background-color: #eee;
    }
    .content_body {
      // position: relative;
      min-height: calc(100% - 0.84rem);
      margin: auto;
      @media screen and (max-width: 441px) {
        min-height: calc(100% - 1.76rem);
      }
      .el-alert /deep/ {
        position: absolute;
        left: 0;
        top: 0;
        z-index: 9;
        .el-alert__icon {
          @media screen and (min-width: 1600px) {
            font-size: 20px;
            width: 20px;
          }
        }
        .el-alert__content {
          display: flex;
          align-items: center;
          .el-alert__title {
            @media screen and (min-width: 1600px) {
              font-size: 14px;
              line-height: 1.3;
            }
            @media screen and (min-width: 1680px) {
              font-size: 16px;
              line-height: 1.3;
            }
            @media screen and (min-width: 1800px) {
              font-size: 18px;
              line-height: 1.3;
            }
            span {
              text-decoration: underline;
              cursor: pointer;
            }
            a {
              text-decoration: underline;
              color: #fff;
            }
          }
          .el-icon-close {
            @media screen and (min-width: 1600px) {
              font-size: 16px;
              line-height: 1.3;
            }
            @media screen and (min-width: 1800px) {
              font-size: 18px;
              line-height: 1.3;
            }
          }
        }
      }
      &.stats {
        display: flex;
        flex-wrap: wrap;
      }
    }
    .fes-icon {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 0 0.6rem;
      margin: auto;
      background-color: #f0f0f0;
      background-color: #23355f;
      background-color: #409eff;
      z-index: 8;
      @media screen and (max-width: 600px) {
        padding: 0 0.3rem;
        flex-wrap: wrap;
        justify-content: center;
      }
      @media screen and (max-width: 441px) {
        padding: 0.15rem 0;
      }
      .fes-icon-logo {
        display: flex;
        justify-content: center;
        align-items: center;
        img {
          display: block;
          height: 20px;
          margin: 0 0.05rem;
          @media screen and (max-width: 991px) {
            height: 20px;
          }
        }
      }
      .fes-icon-copy {
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 0.3rem 0;
        @media screen and (max-width: 441px) {
          width: 100%;
          padding: 0.1rem 0;
        }
        span,
        a {
          font-size: 0.16rem;
          color: #fff;
          line-height: 1.5;
        }
        a {
          &:hover {
            text-decoration: underline;
          }
        }
        .el-divider--vertical /deep/ {
          height: 15px;
        }
      }
      .fes-market {
        display: flex;
        justify-content: center;
        img {
          display: block;
          height: 35px;
          margin: 0;
          @media screen and (min-width: 1800px) {
            height: 45px;
          }
        }
      }
    }
  }
}
</style>
