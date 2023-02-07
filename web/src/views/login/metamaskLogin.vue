<template>
  <div class="metamaskLogin">
    <el-alert type="warning" effect="dark" center show-icon v-if="metaAddress&&!(networkID==3141)">
      <div slot="title">
        {{$t('fs3Login.toptip_01')}} {{metaNetworkInfo.name}} {{$t('fs3Login.toptip_02')}}
        <span @click="changeNet(3141)">Hyperspace Testnet</span>.
        <p v-if="networkID == 137">{{$t('fs3Login.toptip_04')}}
          <a href="https://www.multichain.storage/#/home" target="_blank">multichain.storage</a>.</p>
      </div>
    </el-alert>
    <div class="loginBody width">
      <el-row>
        <el-col :xs="24" :sm="12" :md="13" :lg="13" :xl="13" class="left">
          <h1>{{$t('fs3Login.Connect_text')}}</h1>
          <h4>{{$t('fs3Login.Connect_text_desc')}}</h4>
          <h3>First
            <b>30 GB</b> is
            <b>FREE</b>!</h3>
          <el-button type="primary" v-if="false" @click="goLink('https://www.multichain.storage/')">
            {{$t('fs3Login.Connect_text_btn')}}
          </el-button>
        </el-col>
        <el-col :xs="24" :sm="12" :md="11" :lg="11" :xl="11" class="metamask">
          <div class="titleCont" v-if="!activeIndex">
            <p>{{$t('fs3Login.Connect_to_MetaMask')}}</p>
          </div>
          <div class="titleCont" v-else-if="activeIndex == 'connect'">
            <div class="address">
              <div class="address_left">
                <img src="@/assets/images/metamask.png" class="resno" alt=""> {{ metaAddress | hiddAddress}}
              </div>
              <div class="address_right">
                <div class="flex-shrink-0 w-2 h-2 rounded-full bg-primary"></div>
                <div>{{$t('fs3Login.Connected')}}</div>
              </div>
            </div>
          </div>
          <div v-loading="loginLoad">
            <div class="cont_p">{{$t('fs3Login.Connect_cont_tip')}}</div>
            <div class="login_footer">
              <el-button type="primary" @click="signFun">
                <img src="@/assets/images/metamask.png" class="resno" alt=""> {{$t('fs3Login.Connect_cont_Wallet')}}
              </el-button>
              <p v-if="metaNetworkInfo.center_fail">{{$t('fs3Login.MetaMask_tip')}}</p>
            </div>
            <router-link :to="{name: 'home_entrance'}" class="back_home">{{$t('fs3Login.back_home')}}</router-link>
          </div>
        </el-col>

        <el-col :xs="24" :sm="12" :md="11" :lg="11" :xl="11" class="metamask" v-if="false">
          <div class="titleCont">
            <div class="address">
              <div class="address_left">
                <img src="@/assets/images/metamask.png" class="resno" alt=""> {{ metaAddress | hiddAddress}}
              </div>
              <div class="address_right">
                <div class="flex-shrink-0 w-2 h-2 rounded-full bg-primary"></div>
                <div>{{$t('fs3Login.Connected')}}</div>
              </div>
            </div>
          </div>
          <div v-loading="loginLoad" class="demo-ruleForm">
            <div class="form_title">{{$t('fs3Login.Connect_form_label')}}</div>
            <el-form :model="form" status-icon :rules="rules" ref="form">
              <el-form-item prop="email">
                <el-input v-model="form.email" placeholder="you@domain.com" ref="bucketEmailRef"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="submitEmail('form')">{{$t('fs3Login.Connect_form_btn')}}</el-button>
              </el-form-item>
            </el-form>
          </div>
        </el-col>
        <el-col :xs="24" :sm="12" :md="11" :lg="11" :xl="11" class="checkEmail" v-if="false">
          <div class="check_email">
            <div class="check_left">
              <h3>{{$t('fs3Login.Connect_email_title')}}
                <i class="el-icon-circle-check"></i>
              </h3>
              <h4>{{$t('fs3Login.Connect_email_desc')}}</h4>
              <u>{{$t('fs3Login.Connect_email_again')}}</u>
            </div>
            <div class="check_right">
              <img src="@/assets/images/login/icon_01.png" class="resno" alt="">
            </div>
          </div>
        </el-col>
      </el-row>

      <CarouselContainer :slide-list="collaboratorsData" currentIndex="1"></CarouselContainer>
    </div>
  </div>
</template>

<script>
import CarouselContainer from '@/components/CarouselContainer.vue'
let that
const ethereum = window.ethereum
export default {
  name: 'login',
  data () {
    return {
      fromEnter: '',
      account: null,
      token: null,
      address: null,
      welcome: null,
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
      activeIndex: '',
      loginLoad: false,
      form: {
        email: ''
      },
      rules: {
        email: [{ required: true, message: '', trigger: 'blur' }]
      }
    }
  },
  components: {
    CarouselContainer
  },
  methods: {
    signFun () {
      if (!that.metaAddress || that.metaAddress === undefined) {
        that.$commonFun.Init(async addr => {
          that.$store.dispatch('setMetaAddress', addr)
          sessionStorage.setItem('login_path', addr)
          //  || that.networkID !== 97
          if (that.networkID !== 3141) {
            const networkCont = {
              name: 'Hyperspace Testnet',
              symbol: 'tFIL',
              center_fail: false
            }
            that.$store.dispatch('setMetaNetworkInfo', networkCont)
          }
          that.activeIndex = 'connect'
          that.loginLoad = true
          that.signIn()
        })
        return false
      }
    },
    async signIn () {
      // that.networkID === 97 ||
      if (that.networkID === 3141) {
        const lStatus = await that.$metaLogin.login()
        if (lStatus) {
          that.$router.push({ path: '/my_files' })
          setTimeout(function () { window.location.reload() }, 200)
        } else {
          that.loginLoad = false
          that.activeIndex = ''
        }
      }
    },
    walletInfo () {
      if (!that.metaAddress || that.metaAddress === 'undefined') {
        return false
      }
    },
    // 是否已登录
    isLogin () {
      // that.networkID === 97 ||
      if (that.mcsjwtToken && that.metaAddress && (that.networkID === 3141)) {
        that.$router.push({ path: '/my_files' })
        // that.activeIndex = 'connect'
      } else that.$store.dispatch('setMetaAddress', '')
    },
    goLink (link) {
      window.open(link)
    },
    submitEmail (formName) {
      that.$refs[formName].validate(async valid => {
        if (valid) {

        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    getHome (key) {
      that.$router.push({ name: 'home_entrance', query: { id: key } })
    },
    changeNet (rows) {
      let text = {}
      switch (rows) {
        case 3141:
          text = {
            chainId: that.$web3Init.utils.numberToHex(3141),
            chainName: 'Hyperspace Testnet',
            nativeCurrency: {
              name: 'tFIL',
              symbol: 'tFIL', // 2-6 characters long
              decimals: 18
            },
            rpcUrls: ['https://filecoin-hyperspace.chainstacklabs.com/rpc/v1'],
            blockExplorerUrls: ['https://beryx.zondax.ch/']
          }
          break
      }
      ethereum.request({
        method: 'wallet_addEthereumChain',
        params: [
          text
        ]
      }).then((res) => {
        that.$store.dispatch('setMetaNetworkId', rows)
        that.signIn()
      }).catch((err) => {
        console.log(err)
      })
    }
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
  mounted () {
    that = this
    that.isLogin()
    that.fromEnter = that.$route.query.redirect
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

<style rel="stylesheet/scss" lang="scss" scopte>
@font-face {
  font-family: "gilroy-regular";
  src: url(../../assets/font/gilroy-regular-3.otf);
  font-style: normal;
  font-display: block;
}
.metamaskLogin {
  display: block;
  flex-wrap: wrap;
  position: relative;
  width: 100%;
  height: 100%;
  background: url(../../assets/images/network_logo/login_bg.jpg) no-repeat
    center;
  background-size: cover;
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
    height: calc(100% - 0.85rem);
    padding: 0.6rem 0 0.25rem;
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    .el-row {
      display: flex;
      flex-wrap: wrap;
      align-items: center;
      width: 100%;
      .metamask {
        max-width: 10rem;
        padding: 0.2rem 5% 0.35rem;
        margin: auto;
        background: rgba(4, 20, 23, 0.5);
        border: 1px solid #4e7cff;
        border-radius: 0.2rem;
        .titleCont {
          padding: 0.2rem 0;
          font-weight: 600;
          font-size: 0.286rem;
          font-weight: 600;
          color: #fff;
          line-height: 1.3;
          border-bottom: 1px solid #dde0e2;
          @media screen and (max-width: 600px) {
            font-size: 16px;
          }
          p {
            text-align: center;
          }
          .address {
            display: flex;
            align-items: center;
            justify-content: space-between;
            font-size: 0.223rem;
            font-weight: normal;
            color: #fff;
            line-height: 1.2;
            @media screen and (max-width: 600px) {
              font-size: 13px;
            }
            .address_left {
              display: flex;
              align-items: center;
              img {
                height: 30px;
                margin: 0 0.25rem 0 0;
                @media screen and (max-width: 600px) {
                  height: 24px;
                }
              }
            }
            .address_right {
              position: relative;
              display: inline-block;
              padding: 0.1rem 0.2rem 0.1rem 0.32rem;
              background-color: rgba(85, 128, 233, 0.3);
              font-size: 0.148rem;
              border-radius: 0.5rem;
              @media screen and (max-width: 1600px) {
                font-size: 14px;
              }
              @media screen and (max-width: 1440px) {
                font-size: 13px;
              }
              &::before {
                position: absolute;
                left: 0.16rem;
                top: 50%;
                content: "";
                width: 0.08rem;
                height: 0.08rem;
                margin-top: -0.04rem;
                background-color: #4d73ff;
                border-radius: 0.5rem;
              }
            }
          }
        }
        .el-row {
          border-radius: 0.08rem;
          margin: 0.25rem 0px;
          border: 1px solid rgb(229, 232, 235);
          text-align: center;
          position: static;
          transition: all 0.3s ease 0s;
          overflow: hidden;
          &:hover {
            // background: rgba(240, 185, 11, 0.1);
            // border: 1px solid rgb(240, 185, 11);
            box-shadow: 0 0 6px rgba(0, 0, 0, 0.1);
          }
          .el-col {
            display: flex;
            -webkit-box-pack: justify;
            justify-content: flex-start;
            -webkit-box-align: center;
            align-items: center;
            background: #fff;
            padding: 0.16rem;
            transition: all 0.3s ease 0s;
            cursor: pointer;
            img {
              display: block;
              height: 0.3rem;
              margin: 0 0.15rem 0 0;
              order: 1;
              @media screen and (max-width: 600px) {
                height: 24px;
              }
            }
            span {
              order: 2;
              font-size: 0.18rem;
              @media screen and (max-width: 600px) {
                font-size: 14px;
              }
            }
          }
        }
        .cont_p {
          padding: 0.4rem 5%;
          font-size: 0.223rem;
          font-weight: 100;
          line-height: 1.5;
          text-align: center;
          color: #fff;
          @media screen and (max-width: 600px) {
            font-size: 14px;
          }
        }
        .login_footer {
          padding: 0.25rem 5%;
          .el-button {
            display: block;
            width: 100%;
            max-width: 4rem;
            font-family: inherit;
            font-size: 0.25rem;
            font-weight: 600;
            height: auto;
            padding: 0.2rem;
            margin: auto;
            background: linear-gradient(45deg, #0353fe, #7cebfe);
            color: #fff;
            border: 0;
            border-radius: 0.5rem;
            line-height: 30px;
            letter-spacing: 1px;
            @media screen and (max-width: 600px) {
              font-size: 14px;
              line-height: 24px;
            }
            span {
              display: flex;
              align-items: center;
              justify-content: center;
            }
            img {
              height: 30px;
              margin: 0 0.1rem 0 0;
              @media screen and (max-width: 600px) {
                height: 24px;
              }
            }
            &:hover {
              background: linear-gradient(45deg, #021cad, #3e70fe);
            }
          }
          p {
            font-size: 0.13rem;
            line-height: 1.5;
            color: red;
            margin: 0.1rem 0 0;
            @media screen and (min-width: 1600px) {
              font-size: 14px;
            }
            @media screen and (max-width: 600px) {
              font-size: 12px;
            }
          }
        }
        .back_home {
          float: right;
          color: #fff;
          text-decoration: underline;
          font-size: 14px;
          @media screen and (max-width: 1600px) {
            font-size: 12px;
          }
        }
        .el-loading-mask {
          background-color: rgba(46, 77, 91, 0.75);
        }
        .demo-ruleForm {
          padding: 0.3rem 0;
          .form_title {
            font-size: 0.223rem;
            font-weight: normal;
            color: #fff;
            line-height: 1.2;
            @media screen and (max-width: 600px) {
              font-size: 13px;
            }
          }
          .el-form {
            padding: 0.1rem 0 0;
            .el-form-item {
              display: block;
              height: auto;
              max-width: 4.5rem;
              text-align: left;
              .el-form-item__content {
                height: auto;
                width: 100%;
                font-size: 0.223rem;
                font-weight: normal;
                color: #fff;
                line-height: 1.2;
                @media screen and (max-width: 600px) {
                  font-size: 13px;
                }
                .el-input {
                  .el-input__inner {
                    height: auto;
                    padding: 0.2rem 0.15rem;
                    border-radius: 0.1rem;
                    border-color: #3d6ddb;
                    color: #041417;
                    line-height: 1.2;
                    font-size: inherit;
                  }
                }
                .el-button {
                  display: block;
                  width: 80%;
                  max-width: 4rem;
                  font-family: inherit;
                  font-size: 0.25rem;
                  font-weight: 600;
                  height: auto;
                  padding: 0.2rem;
                  margin: 0;
                  background: linear-gradient(45deg, #0353fe, #7cebfe);
                  color: #fff;
                  border: 0;
                  border-radius: 0.5rem;
                  line-height: 1;
                  letter-spacing: 1px;
                  @media screen and (max-width: 600px) {
                    font-size: 14px;
                  }
                  span {
                    display: flex;
                    align-items: center;
                    justify-content: center;
                  }
                  &:hover {
                    background: linear-gradient(45deg, #021cad, #3e70fe);
                  }
                }
              }
            }
          }
        }
      }
      .checkEmail {
        padding: 5%;
        background: rgba(4, 20, 23, 0.5);
        border: 1px solid #4e7cff;
        border-radius: 0.2rem;
        .check_email {
          display: flex;
          justify-content: center;
          align-items: center;
          .check_left {
            padding-right: 0.1rem;
            color: #fff;
            text-align: left;
            h3 {
              font-size: 0.298rem;
              line-height: 1.5;
              @media screen and (max-width: 1600px) {
                font-size: 22px;
              }
              @media screen and (max-width: 600px) {
                font-size: 16px;
              }
            }
            h4,
            u {
              padding: 0.1rem 0;
              font-size: 0.18rem;
              font-weight: normal;
              line-height: 1.5;
              opacity: 0.9;
              @media screen and (max-width: 1600px) {
                font-size: 16px;
              }
              @media screen and (max-width: 600px) {
                font-size: 14px;
              }
            }
            u {
              cursor: pointer;
            }
          }
          .check_right {
            display: flex;
            justify-content: flex-end;
            align-items: center;
            width: 50%;
            img {
              width: 95%;
              max-width: 1.8rem;
              margin: auto;
            }
          }
        }
      }
      .left {
        padding: 0 5% 0 0;
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
          font-size: 0.3rem;
          font-weight: normal;
          @media screen and (max-width: 768px) {
            font-size: 16px;
          }
        }
        h4 {
          max-width: 6.5rem;
          padding: 0.25rem 0 0.15rem;
          font-size: 0.223rem;
          font-weight: normal;
          @media screen and (max-width: 768px) {
            font-size: 14px;
          }
        }
        .el-button {
          display: inline-block;
          max-width: 3.4rem;
          font-family: inherit;
          font-size: 0.2rem;
          font-weight: 600;
          height: auto;
          padding: 0.2rem;
          margin: 0;
          background: #5580e9;
          color: #fff;
          border-radius: 0.5rem;
          line-height: 1.4;
          white-space: normal;
          @media screen and (max-width: 600px) {
            font-size: 14px;
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
      margin: auto;
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
}
</style>
