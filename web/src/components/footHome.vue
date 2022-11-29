<template>
  <div class="footCont">
    <div class="width">
      <img :src="logo" class="imgLogo" alt="logo" />
      <ul>
        <li @click="goLink('https://docs.filswan.com/multichain.storage/mcp-user-guide')">{{$t('footer.User_Start_Guide')}}</li>
        <li @click="goLink('https://docs.filswan.com/multichain.storage/developer-quickstart')">{{$t('footer.Developer_Quick_Start')}}</li>
        <li @click="goLink('https://docs.filswan.com/multichain.storage/faq')">{{$t('footer.FAQ')}}</li>
        <li @click="goLink('https://discord.gg/pGDTxtgB')">{{$t('footer.Get_Help')}}</li>
      </ul>
      <div class="fes-icon-logo">
        <a href="https://filswan.medium.com/" target="_blank"><img :src="share_medium" alt=""></a>
        <a href="https://discord.com/invite/KKGhy8ZqzK" target="_blank"><img :src="share_discord" alt=""></a>
        <a href="https://twitter.com/0xfilswan" target="_blank"><img :src="share_twitter" alt=""></a>
        <a href="https://github.com/filswan" target="_blank"><img :src="share_github" alt=""></a>
        <a href="https://t.me/filswan" target="_blank"><img :src="share_telegram" alt=""></a>
      </div>
      <p class="copy">{{$t('footer.copy')}}</p>
    </div>

    <pop-ups v-if="ComingSoon" :dialogFormVisible="ComingSoon" :typeModule="typeName" :fixed="true" @getPopUps="getPopUps"></pop-ups>
  </div>
</template>
<script>
import popUps from '@/components/popups'
let that
export default {
  data () {
    return {
      logo: require('@/assets/images/MCS_logo_1.png'),
      share_medium: require('@/assets/images/landing/medium.png'),
      share_twitter: require('@/assets/images/landing/twitter.png'),
      share_github: require('@/assets/images/landing/github-fill.png'),
      share_telegram: require('@/assets/images/landing/telegram.png'),
      share_discord: require('@/assets/images/landing/discord.png'),
      ComingSoon: false,
      typeName: 'comingSoon'
    }
  },
  components: {
    popUps
  },
  computed: {},
  watch: {},
  methods: {
    goLink (link) {
      window.open(link)
    },
    getPopUps (dialog) {
      that.ComingSoon = dialog
    }
  },
  mounted () {
    that = this
  }
}
</script>
<style lang="scss" scoped>
@font-face {
  font-family: "gilroy-regular";
  src: url(../assets/font/gilroy-regular-3.otf);
  font-style: normal;
  font-display: block;
}
.footCont {
  width: 100%;
  padding: 0.95rem 0 0.7rem;
  background-color: #23355f;
  font-family: "gilroy-regular";
  font-size: 16px;
  @media screen and (max-width: 999px) {
    min-height: 60px;
  }
  .width {
    display: flex;
    flex-wrap: wrap;
    width: 80%;
    max-width: 1440px;
    min-width: 300px;
    margin: auto;
    @media screen and (max-width: 1150px) {
      width: 90%;
    }
    @media screen and (max-width: 999px) {
      width: 94%;
    }
  }
  .imgLogo {
    display: block;
    width: 80%;
    max-width: 260px;
    margin: auto;
    @media screen and (max-width: 1680px) {
      max-width: 240px;
    }
    @media screen and (max-width: 1440px) {
      max-width: 220px;
    }
    @media screen and (max-width: 768px) {
      max-width: 200px;
    }
    @media screen and (max-width: 441px) {
      max-width: 180px;
    }
  }
  ul {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    width: 100%;
    padding: 0.6rem 0 0.65rem;
    list-style: none;
    @media screen and (max-width: 768px) {
      padding: 0.3rem 0 0.35rem;
    }
    li {
      padding: 0.1rem 0.25rem 0.1rem 0.55rem;
      margin: 0 0.075rem;
      background-color: #fff;
      border-radius: 0.5rem;
      font-size: 0.223rem;
      font-weight: 600;
      line-height: 1.5;
      color: #2e4580;
      cursor: pointer;
      @media screen and (max-width: 768px) {
        width: calc(50% - 1rem);
        margin: 0.15rem auto;
        font-size: 14px;
      }
      @media screen and (max-width: 441px) {
        width: calc(100% - 1rem);
      }
      &:nth-child(1) {
        background: #fff url(../assets/images/space/icon_14.png) no-repeat
          0.15rem center;
        background-size: 0.3rem;
      }
      &:nth-child(2) {
        background: #fff url(../assets/images/space/icon_15.png) no-repeat
          0.15rem center;
        background-size: 0.3rem;
      }
      &:nth-child(3) {
        background: #fff url(../assets/images/space/icon_16.png) no-repeat
          0.15rem center;
        background-size: 0.3rem;
      }
      &:nth-child(4) {
        background: #fff url(../assets/images/space/icon_17.png) no-repeat
          0.15rem center;
        background-size: 0.3rem;
      }
    }
  }
  .fes-icon-logo {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    width: 100%;
    padding: 0.2rem 0;
    img {
      display: block;
      width: 30px;
      margin: 0 0.15rem;
      @media screen and (max-width: 1600px) {
        width: 25px;
      }
      @media screen and (max-width: 1024px) {
        width: 20px;
      }
    }
  }
  .copy {
    width: 100%;
    font-size: 0.14rem;
    font-weight: 100;
    text-align: center;
    line-height: 1.2;
    color: #fff;
    @media screen and (min-width: 1800px) {
      font-size: 14px;
    }
    @media screen and (max-width: 1440px) {
      font-size: 13px;
    }
    @media screen and (max-width: 441px) {
      font-size: 12px;
    }
  }
}
</style>
