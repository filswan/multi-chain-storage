<template>
  <div class="metamaskHome">
    <v-head @getHome="getHome" @getLogin="signFun" :loginLoad="loginLoad"></v-head>
    <div class="loginBody">
      <div class="width">
        <el-row>
          <el-col :xs="24" :sm="12" :md="11" :lg="11" :xl="11" class="left">
            <h1>{{$t('fs3Login.Connect_text')}}</h1>
            <h4>{{$t('fs3Login.Connect_text_desc')}}</h4>
            <el-button type="primary" v-loading="loginLoad" @click="signFun">
              {{$t('fs3Login.Connect_StartFree')}}
            </el-button>
            <el-button type="primary" @click="goLink('https://www.youtube.com/watch?v=rgEP4_dhzoI')">
              {{$t('fs3Login.Connect_TutorialVideo')}}
            </el-button>
            <h3>First
              <b>30 GB</b> is
              <b>FREE</b>!</h3>
          </el-col>
          <el-col :xs="24" :sm="12" :md="13" :lg="13" :xl="13" class="right">
            <img :src="img" />
          </el-col>
        </el-row>

        <CarouselContainer :slide-list="collaboratorsData" currentIndex="1"></CarouselContainer>
      </div>
    </div>
    <div class="msCont">
      <div class="width" id="about">
        <div class="title">
          <i class="icon"></i> Multichain.Storage
        </div>
        <el-row type="flex" class="row-bg" justify="space-between">
          <el-col :xs="24" :sm="11" :md="11" :lg="11" :xl="11" class="left">
            <div class="title">
              <i class="icon icon_Introduction"></i>
              Introduction
            </div>
            <p>Multichain.Storage is a smart-contract-based cross-chain storage gateway integrated with Oracle technology and the Filecoin networks. It accelerates the mass adoption of decentralized storage by bridging multiple blockchain networks.</p>
          </el-col>
          <el-col :xs="24" :sm="11" :md="11" :lg="11" :xl="11" class="left">
            <div class="title">
              <i class="icon icon_Features"></i>
              Our Features
            </div>
            <p class="p">Drag and drop files to pin to IPFS and store on the Filecoin network at the same time</p>
            <p class="p">Fast retrieval globally by reliable edge network</p>
            <p class="p">Developer support of JS, Python and Golang SDK</p>
          </el-col>
        </el-row>
      </div>
    </div>
    <div class="pricing" v-if="false">
      <div class="width" id="pricing">
        <div class="title">
          <i class="icon"></i> Pricing for Multichain.Storage
          <el-button type="primary" v-loading="loginLoad" @click="signFun">
            {{$t('fs3Login.Connect_StartFree')}}
          </el-button>
        </div>
        <el-row type="flex" class="row-bg" justify="space-around">
          <el-col :xs="24" :sm="11" :md="11" :lg="11" :xl="11" class="left">
            <h1>Free</h1>
            <p>One bucket (with 30GB storage) free</p>
            <p>Pin to IPFS</p>
            <p>Weekly backup to decentralize storage network (5 copies)</p>
            <p>SDK/API technical support</p>
          </el-col>
          <el-col :xs="24" :sm="11" :md="11" :lg="11" :xl="11" class="left">
            <h1>More Storage</h1>
            <p>USDC 7.2 for each added bucket each year
              <span>(30GB storage for each bucket)</span>
            </p>
            <p>33 buckets for each wallet</p>
            <p>Pin to IPFS</p>
            <p>Weekly backup to decentralize storage network (5 copies)</p>
            <p>SDK/API technical support</p>
          </el-col>
        </el-row>
      </div>
    </div>
    <v-foot id="resources"></v-foot>
    <el-backtop target=".metamaskHome"></el-backtop>
    <network-alert v-if="metaAddress&&networkTip" @changeNet="changeNet" @getNetwork="getNetwork"></network-alert>
  </div>
</template>

<script>
import vHead from '@/components/headHome.vue'
import vFoot from '@/components/footHome.vue'
import networkAlert from '@/components/networkAlert.vue'
import CarouselContainer from '@/components/CarouselContainer.vue'
import NCWeb3 from '@/utils/web3'
import metaLogin from '@/utils/login'
let that
const ethereum = window.ethereum
export default {
  name: 'home',
  data () {
    return {
      collaboratorsData: [
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-01.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-01-sunny.png'),
          link: 'https://protocol.ai/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-02.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-02-sunny.png'),
          link: 'https://ipfs.tech/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-03.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-03-sunny.png'),
          link: 'https://filecoin.io/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-05.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-05-sunny.png'),
          link: 'https://chainlinklabs.com/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-06.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-06-sunny.png'),
          link: 'https://labs.binance.com/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-08.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-08-sunny.png'),
          link: 'https://arbitrum.io/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-07.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-07-sunny.png'),
          link: 'https://opensea.io/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-09.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-09-sunny.png'),
          link: 'https://akash.network/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-04.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-04-sunny.png'),
          link: 'https://polygon.technology/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-10.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-10-sunny.png'),
          link: 'https://aptos.dev/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-11.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-11-sunny.png'),
          link: 'https://sui.io/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-13.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-13-sunny.png'),
          link: 'https://en.fogmeta.com/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-14.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-14-sunny.png'),
          link: 'https://www.web3cloud.tech/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-15.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-15-sunny.png'),
          link: 'https://github.com/srblabotw69/Afianthack'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-16.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-16-sunny.png'),
          link: 'https://github.com/IKalonji/CooperDB'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-17.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-17-sunny.png'),
          link: 'https://sao.network/#/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-12.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-12-sunny.png'),
          link: 'https://www.nebulablock.com/'
        }
      ],
      img: require('@/assets/images/space/1.png'),
      loginLoad: false,
      prevType: true,
      networkTip: false
    }
  },
  components: {
    vHead, vFoot, CarouselContainer, networkAlert
  },
  methods: {
    goLink (link) {
      window.open(link)
    },
    getHome (key) {
      var PageId = document.querySelector('#' + key)
      document.querySelector('.metamaskHome').scrollTo({
        top: PageId.offsetTop,
        behavior: 'smooth'
      })
    },
    // 是否已登录
    async isLogin () {
      let status = await metaLogin.netStatus(that.networkID)
      if (that.mcsjwtToken && that.metaAddress && status) {
        that.$router.push({ path: '/my_files' })
      } else that.$store.dispatch('setMetaAddress', '')
    },
    async changeNet (rows) {
      let text = {}
      switch (rows) {
        case 80001:
          text = {
            chainId: '0x13881',
            chainName: 'Mumbai Testnet',
            nativeCurrency: {
              name: 'MATIC',
              symbol: 'MATIC', // 2-6 characters long
              decimals: 18
            },
            rpcUrls: ['https://rpc-mumbai.maticvigil.com'],
            blockExplorerUrls: ['https://mumbai.polygonscan.com/']
          }
          break
        case 97:
          text = {
            chainId: '0x61',
            chainName: 'BSC TestNet',
            nativeCurrency: {
              name: 'tBNB',
              symbol: 'tBNB', // 2-6 characters long
              decimals: 18
            },
            rpcUrls: ['https://data-seed-prebsc-1-s1.binance.org:8545'],
            blockExplorerUrls: ['https://testnet.bscscan.com']
          }
          break
        case 137:
          text = {
            chainId: '0x89',
            chainName: 'Polygon Mainnet',
            nativeCurrency: {
              name: 'tBNB',
              symbol: 'tBNB', // 2-6 characters long
              decimals: 18
            },
            rpcUrls: ['https://polygon-rpc.com'],
            blockExplorerUrls: ['https://polygonscan.com/']
          }
          break
      }
      try {
        await ethereum.request({
          method: 'wallet_addEthereumChain',
          params: [text]
        })
        that.$store.dispatch('setMetaNetworkId', rows)
        that.signIn()
      } catch (error) {
        console.log(error)
      }
    },
    signFun () {
      if (!that.metaAddress || that.metaAddress === undefined) {
        NCWeb3.Init(async addr => {
          that.$store.dispatch('setMetaAddress', addr)
          sessionStorage.setItem('login_path', addr)
          that.loginLoad = true
          that.signIn()
        })
      }
      const networkCont = {
        name: that.networkID === 137 ? 'Polygon' : that.networkID === 80001 ? 'Mumbai' : that.networkID === 97 ? 'BSC' : 'Custom',
        unit: 'USDC',
        center_fail: false
      }
      that.$store.dispatch('setMetaNetworkInfo', networkCont)
    },
    async signIn () {
      let status = await metaLogin.netStatus(that.networkID)
      that.networkTip = !status
      if (!status) return false
      const lStatus = await metaLogin.login()
      if (lStatus) setTimeout(function () { window.location.reload() }, 200)
      else that.loginLoad = false
    },
    getNetwork (dis) {
      that.networkTip = dis
    },
    fn () {
      document.addEventListener('visibilitychange', function () {
        that.prevType = !document.hidden
      })
      // networkChanged
      ethereum.on('chainChanged', async (accounts) => {
        if (!that.prevType || !that.metaAddress) return false
        let chainID = parseInt(accounts, 16)
        that.$store.dispatch('setMetaNetworkId', chainID)
        that.signFun()
        let status = await metaLogin.netStatus(chainID)
        if (status) that.changeNet(chainID)
      })
    }
  },
  mounted () {
    that = this
    if (that.$route.query.id) that.getHome(that.$route.query.id)
    that.isLogin()
    that.fn()
  },
  computed: {
    metaAddress () {
      return this.$store.getters.metaAddress
    },
    networkID () {
      return Number(this.$store.getters.networkID)
    },
    metaNetworkInfo () {
      return this.$store.getters.metaNetworkInfo ? JSON.parse(JSON.stringify(this.$store.getters.metaNetworkInfo)) : {}
    },
    mcsjwtToken () {
      return this.$store.getters.mcsjwtToken
    }
  },
  filters: {
    hiddAddress: function (val) {
      if (val) {
        return `${val.substring(0, 6)}...${val.substring(val.length - 4)}`
      } else {
        return '-'
      }
    }
  }
}
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
@font-face {
  font-family: "gilroy-regular";
  src: url(../../assets/font/gilroy-regular-3.otf);
  font-style: normal;
  font-display: block;
}
.metamaskHome {
  display: block;
  flex-wrap: wrap;
  position: relative;
  width: 100%;
  height: 100%;
  background-color: #fff;
  font-family: "gilroy-regular";
  font-size: 16px;
  overflow-y: scroll;
  &::-webkit-scrollbar-track {
    background: transparent;
  }
  &::-webkit-scrollbar {
    width: 6px;
    background: transparent;
  }
  &::-webkit-scrollbar-thumb {
    background: #ccc;
  }
  .el-alert {
    position: fixed;
    left: 0;
    top: 0;
    z-index: 99;
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
  .loginBody {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    padding: 1.3rem 0 1rem;
    background: url(../../assets/images/network_logo/login_bg.jpg) no-repeat
      center;
    background-size: cover;
    .el-row /deep/ {
      display: flex;
      flex-wrap: wrap;
      align-items: center;
      width: 100%;
      .right {
        max-width: 10rem;
        padding: 0;
        margin: auto;
        img {
          display: block;
          width: 100%;
          margin: auto;
        }
      }
      .left {
        padding: 0 0.3rem 0 0;
        color: #fff;
        line-height: 1.2;
        @media screen and (max-width: 992px) {
          padding: 0.5rem 0;
          text-align: center;
        }
        h1 {
          font-size: 0.776rem;
          text-transform: capitalize;
          @media screen and (max-width: 768px) {
            font-size: 22px;
          }
        }
        h3 {
          padding: 0.25rem 0 0.15rem;
          font-size: 0.3rem;
          font-weight: normal;
          @media screen and (max-width: 768px) {
            font-size: 16px;
          }
        }
        h4 {
          padding: 0.25rem 0 0.35rem;
          font-size: 0.223rem;
          font-weight: normal;
          @media screen and (max-width: 768px) {
            font-size: 14px;
          }
        }
        .el-button {
          display: inline-block;
          max-width: 3.4rem;
          min-width: 1.55rem;
          font-family: inherit;
          font-size: 0.2rem;
          font-weight: normal;
          height: auto;
          padding: 0.15rem 0.35rem;
          margin: 0 0.2rem 0 0;
          background: #5580e9;
          color: #fff;
          border-radius: 0.5rem;
          line-height: 1.4;
          white-space: normal;
          @media screen and (max-width: 600px) {
            font-size: 14px;
          }
          .el-loading-mask {
            border-radius: 0.5rem;
            z-index: 2;
          }
          &:hover {
            background: #4326ab;
          }
        }
      }
    }
    .collaborators {
      display: flex;
      justify-content: flex-start;
      align-items: center;
      flex-wrap: wrap;
      width: 96%;
      margin: 0.7rem auto 0;
      @media screen and (max-width: 992px) {
        margin: 0.5rem auto 0;
      }
      .el-col {
        margin: 0 auto;
        a {
          img {
            display: block;
            height: 45px;
            margin: 20px auto;
            cursor: pointer;
            @media screen and (max-width: 1600px) {
              height: 40px;
            }
            @media screen and (max-width: 1440px) {
              margin: 15px auto;
            }
          }
          .height {
            height: 80px;
            margin: 10px auto;
            @media screen and (max-width: 1600px) {
              height: 65px;
            }
            @media screen and (max-width: 1440px) {
              margin: 5px auto;
            }
          }
        }
      }
    }
  }
  .msCont {
    padding: 1.7rem 0;
    @media screen and (max-width: 992px) {
      padding: 1rem 0;
    }
    .title {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      width: 100%;
      font-size: 0.55rem;
      font-weight: 900;
      color: #010102;
      line-height: 1.2;
      @media screen and (max-width: 441px) {
        font-size: 22px;
      }
      .icon {
        width: 0.54rem;
        height: 0.54rem;
        margin: 0 0.35rem 0 0;
        background: url(../../assets/images/space/icon_18.png) no-repeat center;
        background-size: 100%;
      }
      .icon_Introduction {
        background: url(../../assets/images/space/icon_12.png) no-repeat center;
        background-size: 100%;
      }
      .icon_Features {
        background: url(../../assets/images/space/icon_13.png) no-repeat center;
        background-size: 100%;
      }
    }
    .el-row /deep/ {
      width: 100%;
      padding: 1rem 0 0;
      flex-wrap: wrap;
      @media screen and (max-width: 1260px) {
        padding: 0.7rem 0 0;
      }
      @media screen and (max-width: 441px) {
        padding: 0.3rem 0 0;
      }
      .el-col {
        margin: 0.2rem 0;
        padding: 0.4rem 8% 0.3rem;
        background-color: #dde6fb;
        border-radius: 0.15rem;
        @media screen and (max-width: 1260px) {
          padding: 0.4rem 5% 0.3rem;
        }
        .title {
          padding: 0 0 0.4rem;
          font-size: 0.31rem;
          color: #5580e9;
        }
        p {
          font-size: 14px;
          color: #010102;
          line-height: 2;
          word-break: break-word;
          @media screen and (min-width: 1800px) {
            font-size: 16px;
          }
        }
        .p {
          position: relative;
          padding: 0 0 0 15px;
          &::before {
            position: absolute;
            left: 0;
            top: 10px;
            content: "";
            width: 6px;
            height: 6px;
            background-color: #5580e9;
            border-radius: 100%;
          }
        }
      }
    }
  }
  .pricing {
    padding: 1rem 0 0.7rem;
    margin: 0 0 1.4rem;
    background-color: #f3f6fd;
    @media screen and (max-width: 992px) {
      padding: 0.6rem 0;
      margin: 0 0 0.7rem;
    }
    .title /deep/ {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      width: 100%;
      font-size: 0.55rem;
      font-weight: 900;
      color: #010102;
      line-height: 1.2;
      @media screen and (max-width: 441px) {
        font-size: 22px;
      }
      .icon {
        width: 0.54rem;
        height: 0.54rem;
        margin: 0 0.35rem 0 0;
        background: url(../../assets/images/space/icon_18.png) no-repeat center;
        background-size: 100%;
      }
      .el-button {
        display: inline-block;
        max-width: 3.4rem;
        min-width: 2.8rem;
        font-family: inherit;
        font-size: 0.316rem;
        font-weight: normal;
        height: auto;
        padding: 0.15rem 0.35rem;
        margin: 0.15rem 0 0.15rem 0.8rem;
        background: #5580e9;
        color: #fff;
        border-radius: 0.5rem;
        line-height: 1.4;
        white-space: normal;
        @media screen and (max-width: 600px) {
          font-size: 16px;
        }
        .el-loading-mask {
          border-radius: 0.5rem;
          z-index: 2;
        }
        &:hover {
          background: #4326ab;
        }
      }
    }
    .el-row /deep/ {
      width: 100%;
      padding: 1rem 0 0;
      flex-wrap: wrap;
      @media screen and (max-width: 1260px) {
        padding: 0.7rem 0 0;
      }
      @media screen and (max-width: 441px) {
        padding: 0.3rem 0 0;
      }
      .el-col {
        margin: 0.2rem 0;
        padding: 0 4%;
        @media screen and (max-width: 1260px) {
          padding: 0 2%;
        }
        h1 {
          display: inline-block;
          font-family: inherit;
          font-size: 0.4rem;
          font-weight: 900;
          height: auto;
          padding: 0.15rem 0;
          margin: 0 0 0.35rem;
          color: #5580e9;
          line-height: 1.4;
          white-space: normal;
          text-align: center;
          @media screen and (max-width: 600px) {
            font-size: 22px;
          }
          @media screen and (max-width: 441px) {
            font-size: 20px;
            margin: 0 0 0.15rem;
          }
        }
        p {
          position: relative;
          padding: 5px 5px 5px 15px;
          font-size: 14px;
          color: #010102;
          line-height: 1.2;
          word-break: break-word;
          @media screen and (min-width: 1800px) {
            font-size: 16px;
          }
          span {
            display: block;
            width: 100%;
            font-size: 11px;
            color: #4f4f4f;
          }
          &::before {
            position: absolute;
            left: 0;
            top: 10px;
            content: "";
            width: 6px;
            height: 6px;
            background-color: #5580e9;
            border-radius: 100%;
          }
        }
      }
    }
  }
}
</style>
