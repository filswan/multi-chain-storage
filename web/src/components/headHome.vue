<template>
  <div class="headerCont">
    <div class="landing">
      <div class="header width">
        <div class="header_left">
          <div class="logoImg" @click="header_logo">
            <img :src="logo" class="img" alt='FilSwan' />
            <img class="beta" src="@/assets/images/landing/beta.png">
          </div>
          <router-link v-if="addrChild" :to="{name: 'Space'}" :class="{'meta-space pcShow':true, 'active': $route.name === 'Space'}">
            {{$t('route.metaSpace')}}
          </router-link>
        </div>
        <el-menu :default-active="activeIndex" class="el-menu-demo " mode="horizontal" @select="handleSelect" background-color="transparent" text-color="#333" active-text-color="#5580e9">
          <el-menu-item index="about" class="pcShow">
            <router-link to="">
              {{ $t('route.About') }}
            </router-link>
          </el-menu-item>
          <el-menu-item index="pricing" class="pcShow">
            <router-link to="">
              {{ $t('route.Pricing') }}
            </router-link>
          </el-menu-item>
          <el-menu-item index="resources" class="pcShow">
            <router-link to="">
              {{ $t('route.Resources') }}
            </router-link>
          </el-menu-item>
          <el-menu-item index="stats" class="pcShow">
            <router-link to="">
              {{ $t('route.Stats') }}
            </router-link>
          </el-menu-item>
          <el-menu-item index="auditReport" class="pcShow" @click="goLink('https://github.com/numencyber/AuditReport/blob/main/FILSWAN-Smart-Contract-Audit-Report%20-%20Numen.pdf')">
            <router-link to="">
              {{ $t('route.report') }}
            </router-link>
          </el-menu-item>
          <el-menu-item index="login" class="pcShow" v-if="!addrChild">
            <a href="javascript:;" v-loading="loginLoad" @click="getLogin" class="target">
              {{ $t('route.Login') }}
            </a>
          </el-menu-item>
          <el-menu-item index="logged" v-else>
            <div class="online feh-metamask">
              <div class="logged_in">
                <div class="info">
                  <h5 @click="wrongVisible = true">{{addrChild | hiddAddress}}</h5>
                  <el-dropdown class="dropdown-style" @command="handleCommand" trigger="click" :hide-on-click="true">
                    <span class="el-dropdown-link">
                      <i class="el-icon-setting"></i>
                    </span>
                    <el-dropdown-menu slot="dropdown" class="dropdown-ul">
                      <el-dropdown-item command="bucket" class="mobileShow">
                        <div class="dropdown-body flex">
                          <div class="dropdown-left flex">
                            <i class="el-icon-s-metaSpace"></i>
                            {{$t('route.metaSpace')}}
                          </div>
                        </div>
                      </el-dropdown-item>
                      <el-dropdown-item command="setting">
                        <div class="dropdown-body flex">
                          <div class="dropdown-left flex">
                            <i class="el-icon-s-myAccount"></i>
                            {{$t('route.myAccount')}}
                          </div>
                        </div>
                      </el-dropdown-item>
                      <el-dropdown-item command="lang">
                        <div class="dropdown-body flex">
                          <div class="dropdown-left flex">
                            <i class="el-icon-language"></i>
                            {{$t('navbar.language')}}
                          </div>
                          <div class="lang_style">
                            <span v-if="languageMcs === 'en'">EN</span>
                            <span v-else>中</span>
                          </div>
                        </div>
                      </el-dropdown-item>
                      <el-dropdown-item command="theme">
                        <div class="dropdown-body flex">
                          <div class="dropdown-left flex">
                            <i class="el-icon-theme"></i>
                            {{$t('comment.theme')}}
                          </div>
                          <div class="switch">
                            <div class="swithUI flex" v-if="reverse">
                              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" class="css-sunny">
                                <path fill-rule="evenodd" clip-rule="evenodd" d="M10.5 2h3v3h-3V2zM16 12a4 4 0 11-8 0 4 4 0 018 0zM5.99 3.869L3.867 5.99 5.99 8.112 8.111 5.99 5.989 3.87zM2 13.5v-3h3v3H2zm1.868 4.51l2.121 2.12 2.122-2.12-2.122-2.122-2.121 2.121zM13.5 19v3h-3v-3h3zm4.51-3.112l-2.121 2.122 2.121 2.121 2.121-2.121-2.121-2.122zM19 10.5h3v3h-3v-3zm-3.11-4.51l2.12 2.121 2.122-2.121-2.121-2.121-2.122 2.121z"
                                  fill="currentColor"></path>
                              </svg>
                            </div>
                            <div class="swithUI flex" v-else>
                              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" class="css-moon">
                                <path d="M20.968 12.768a7 7 0 01-9.735-9.735 9 9 0 109.735 9.735z" fill="currentColor"></path>
                              </svg>
                            </div>
                          </div>
                        </div>
                      </el-dropdown-item>
                      <el-dropdown-item command="doc">
                        <div class="dropdown-body flex" @click="documentLink">
                          <div class="dropdown-left flex">
                            <i class="el-icon-s-documentation"></i>
                            <span slot="title">{{$t('route.documentation')}}</span>
                          </div>
                        </div>
                      </el-dropdown-item>
                      <el-dropdown-item command="disconnect">
                        <div class="dropdown-body flex" @click="signOutFun">
                          <div class="dropdown-left flex">
                            <i class="el-icon-switch-button"></i>
                            <div class="text textTrue">{{$t('fs3.Disconnect')}}</div>
                          </div>
                        </div>
                      </el-dropdown-item>
                    </el-dropdown-menu>
                  </el-dropdown>
                </div>
              </div>
            </div>
          </el-menu-item>
          <el-menu-item index="logged" class="mobileShow">
            <!-- 移动端菜单 -->
            <div class="mobileShow" @click="mobileMenuFun">
              <div>
                <span></span>
                <span></span>
                <span></span>
              </div>
            </div>
          </el-menu-item>
        </el-menu>

        <!-- 移动端导航 -->
        <div class="menuBg" v-if="mobileMenuShow" @click="mobileMenuFun"></div>
        <div :class="{'menuMList': 1==1, 'menuWidth': mobileMenuShow}" v-show="mobileMenuShow">
          <div class="contentMenu" :class="{'contentMenuX': mobileMenuSecord}">
            <div class="first_level_menu">
              <div class="menuMListChild" @click="handleMoSelect('about')">
                {{ $t('route.About') }}
              </div>
              <div class="menuMListChild" @click="handleMoSelect('pricing')">
                {{ $t('route.Pricing') }}
              </div>
              <div class="menuMListChild" @click="handleMoSelect('resources')">
                {{ $t('route.Resources') }}
              </div>
              <div class="menuMListChild" @click="handleMoSelect('stats')">
                {{ $t('route.Stats') }}
              </div>
              <a href="javascript:;" @click="goLink('https://github.com/numencyber/AuditReport/blob/main/FILSWAN-Smart-Contract-Audit-Report%20-%20Numen.pdf')" class="menuMListChild">
                {{ $t('route.report') }}
              </a>
              <a href="javascript:;" v-if="!addrChild" v-loading="loginLoad" @click="getLogin" class="menuMListChild">
                {{ $t('route.Login') }}
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>

    <el-dialog :title="$t('fs3Login.Account')" :visible.sync="wrongVisible" :modal-append-to-body="false" :width="width" custom-class="wrongNet">
      <label>{{$t('fs3Login.Connected_MetaMask')}}</label>
      <div class="address">{{addrChild | hiddAddress}}</div>
      <div class="share">
        <el-button @click="shareTo">
          <svg t="1669800457857" class="icon icon_big" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="6207" width="64" height="64">
            <path d="M923.648 1015.442H100.206a91.648 91.648 0 0 1-91.721-91.72V101.01a91.502 91.502 0 0 1 91.72-91.501H649.29a30.72 30.72 0 0 1 0 61.44H130.487a60.855 60.855 0 0 0-60.928 60.854v762.003a60.855 60.855 0 0 0 60.928 60.928h762.441a60.855 60.855 0 0 0 60.928-60.928V345.088a30.72 30.72 0 1 1 61.44 0v579.291a91.21 91.21 0 0 1-91.648 91.063z m-497.81-403.675a30.574 30.574 0 1 1-43.228-43.228L930.816 17.92a30.574 30.574 0 1 1 43.154 43.3L425.91 611.768z"
              p-id="6208" fill="#0b318f"></path>
            <path d="M923.648 1023.854H100.206A100.206 100.206 0 0 1 0.073 923.72v-822.71C0.22 45.86 44.91 1.096 100.206 1.096h549.083a39.131 39.131 0 1 1 0 78.263H130.414a52.443 52.443 0 0 0-52.444 52.443v762.003c0 28.964 23.48 52.443 52.517 52.516H893a52.368 52.368 0 0 0 37.084-15.36 52.81 52.81 0 0 0 15.36-37.156V345.088a39.131 39.131 0 0 1 78.262 0v579.291a99.913 99.913 0 0 1-100.059 99.475zM100.059 17.92c-45.787 0-82.944 37.23-83.017 83.09v822.784c0.073 46.007 37.303 83.237 83.31 83.31h823.37a83.09 83.09 0 0 0 83.163-82.798V345.015a22.309 22.309 0 0 0-44.544 0v548.864c0 18.359-7.315 35.986-20.188 49.006a68.754 68.754 0 0 1-49.079 20.333H130.487a69.486 69.486 0 0 1-69.34-69.34V131.804a69.266 69.266 0 0 1 69.267-69.339h518.948a22.309 22.309 0 1 0-0.146-44.544h-549.01z m304.202 611.328a39.058 39.058 0 0 1-27.575-66.706L924.818 11.995a38.985 38.985 0 1 1 55.077 55.223l-548.06 550.473c-7.314 7.315-17.261 11.484-27.574 11.557zM952.32 17.335a22.162 22.162 0 0 0-15.58 6.583L388.536 574.39a22.162 22.162 0 1 0 31.378 31.451L968.046 55.296a21.943 21.943 0 0 0 6.583-15.726 22.382 22.382 0 0 0-22.236-22.235z"
              p-id="6209" fill="#0b318f"></path>
          </svg>
          {{$t('fs3Login.View_explorer')}}
        </el-button>

        <el-button v-if="copyClick" @click="copyTextToClipboard(addrChild)">
          <svg t="1640938541398" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4760" width="32" height="32">
            <path d="M746.932 698.108" p-id="4761" fill="#0b318f"></path>
            <path d="M925.731 288.698c-1.261-1.18-3.607-3.272-6.902-6.343-5.486-5.112-11.615-10.758-18.236-16.891-18.921-17.526-38.003-35.028-56.046-51.397-2.038-1.848-2.038-1.835-4.077-3.682-24.075-21.795-44.156-39.556-58.996-52.076-8.682-7.325-15.517-12.807-20.539-16.426-3.333-2.402-6.043-4.13-8.715-5.396-3.365-1.595-6.48-2.566-10.905-2.483C729.478 134.227 720 143.77 720 155.734l0 42.475 0 42.475 0 84.95L720 347l21.205 0L890 347l0 595L358.689 942C323.429 942 295 913.132 295 877.922L295 177l361.205 0c11.736 0 21.25-9.771 21.25-21.5s-9.514-21.5-21.25-21.5l-382.5 0L252 134l0 21.734L252 813l-52.421 0C166.646 813 140 786.928 140 754.678L140 72l566.286 0C739.29 72 766 98.154 766 130.404L766 134l40 0 0-3.596C806 76.596 761.271 33 706.286 33L119.958 33 100 33l0 19.506 0 702.172C100 808.463 144.642 852 199.579 852L252 852l0 25.922C252 936.612 299.979 984 358.689 984l552.515 0L932 984l0-21.237L932 325.635 932 304l0.433 0C932.432 299 930.196 292.878 925.731 288.698zM762 304l0-63.315L762 198.21l0-0.273c14 11.479 30.3 26.369 49.711 43.942 2.022 1.832 2.136 1.832 4.157 3.665 17.923 16.259 36.957 33.492 55.779 50.926 2.878 2.666 5.713 5.531 8.391 7.531L762 304.001z"
              p-id="4762" fill="#0b318f"></path>
            <path d="M816.936 436 407.295 436c-10.996 0-19.91 8.727-19.91 19.5 0 10.77 8.914 19.5 19.91 19.5l409.641 0c11 0 19.914-8.73 19.914-19.5C836.85 444.727 827.936 436 816.936 436z" p-id="4763" fill="#0b318f"></path>
            <path d="M816.936 553 407.295 553c-10.996 0-19.91 8.727-19.91 19.5 0 10.774 8.914 19.5 19.91 19.5l409.641 0c11 0 19.914-8.726 19.914-19.5C836.85 561.727 827.936 553 816.936 553z" p-id="4764" fill="#0b318f"></path>
            <path d="M816.936 689 407.295 689c-10.996 0-19.91 8.729-19.91 19.503 0 10.769 8.914 19.497 19.91 19.497l409.641 0c11 0 19.914-8.729 19.914-19.497C836.85 697.729 827.936 689 816.936 689z" p-id="4765" fill="#0b318f"></path>
          </svg>
          {{$t('fs3Login.Copy_Address')}}
        </el-button>

        <el-button v-else>
          <svg t="1640939105223" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5619" width="32" height="32">
            <path d="M512 998.4c-268.288 0-486.4-218.112-486.4-486.4S243.712 25.6 512 25.6s486.4 218.112 486.4 486.4-218.112 486.4-486.4 486.4z m0-896c-225.792 0-409.6 183.808-409.6 409.6s183.808 409.6 409.6 409.6 409.6-183.808 409.6-409.6-183.808-409.6-409.6-409.6z"
              p-id="5620" fill="#0b318f"></path>
            <path d="M457.728 680.448c-9.728 0-19.456-3.584-27.136-11.264l-150.528-150.016c-14.848-14.848-14.848-39.424 0-54.272 14.848-14.848 39.424-14.848 54.272 0l123.392 122.88 231.936-230.912c14.848-14.848 39.424-14.848 54.272 0 14.848 14.848 14.848 39.424 0 54.272l-259.072 258.048c-7.68 7.168-17.408 11.264-27.136 11.264z"
              p-id="5621" fill="#0b318f"></path>
          </svg>
          {{$t('fs3Login.Copied')}}
        </el-button>

      </div>
      <div class="loadStyle" v-show="loadAccount" v-loading="loadAccount"></div>
    </el-dialog>
    <div class="loadIndexStyle" v-show="loadIndexing" v-loading="loadIndexing"></div>
  </div>
</template>
<script>
let that
export default {
  data () {
    return {
      logo: require('@/assets/images/MCS_logo_2.png'),
      activeIndex: '4',
      mobileMenuShow: false,
      mobileMenuSecord: false,
      addrChild: '',
      width: document.body.clientWidth > 600 ? '450px' : '95%',
      copyClick: true,
      wrongVisible: false
    }
  },
  props: ['loginLoad'],
  computed: {
    metaAddress () {
      return this.$store.getters.metaAddress
    },
    routerMenu () {
      return this.$store.getters.routerMenu
    },
    languageMcs () {
      return this.$store.getters.languageMcs
    },
    reverse () {
      return String(this.$store.getters.reverse) === '1'
    }
  },
  watch: {
    $route: function (to, from) {
      this.activeName()
    },
    addrChild: function () {
      if (this.addrChild) this.commonParam()
    }
  },
  methods: {
    signOutFun () {
      that.addrChild = ''
      that.$store.dispatch('setMetaAddress', '')
      that.$store.dispatch('setMCSjwtToken', '')
      that.$store.dispatch('setMetaNetworkId', 0)
      sessionStorage.removeItem('login_path')
      sessionStorage.removeItem('mcs_dev_jwtToken')
      // that.$router.push('/home')
      setTimeout(function () { window.location.reload() }, 200)
    },
    async commonParam () {
      let paymentRes = await that.$commonFun.sendRequest(`${that.baseAPIURL}api/v2/payment/get_payment`, 'get')
      if (!paymentRes || paymentRes.status !== 'success') {
        if (paymentRes) that.$message.error(paymentRes.message || 'Fail')
        else that.$message.error('Fail')
      } else {
        that.$root.plan_id = paymentRes.data.plan_id
        that.$root.max_storage = paymentRes.data.max_storage
      }

      let networkRes = await that.$commonFun.sendRequest(`${that.baseAPIURL}api/v2/network/get_network`, 'get')
      if (!networkRes || networkRes.status !== 'success') {
        if (networkRes) that.$message.error(networkRes.message || 'Fail')
        else that.$message.error('Fail')
      } else {
        that.$root.PAYMENT_CONTRACT_ADDRESS = networkRes.data[0].payment_contract_address
        that.$root.pay_name = networkRes.data[0].name
        that.$root.chain_id = networkRes.data[0].chain_id
      }
    },
    shareTo () {
      window.open(`${that.baseAddressURL}address/${that.addrChild}`)
    },
    copyTextToClipboard (text, type) {
      var txtArea = document.createElement('textarea')
      txtArea.id = 'txt'
      txtArea.style.position = 'fixed'
      txtArea.style.top = '0'
      txtArea.style.left = '0'
      txtArea.style.opacity = '0'
      txtArea.value = text
      document.body.appendChild(txtArea)
      txtArea.select()

      try {
        var successful = document.execCommand('copy')
        var msg = successful ? 'successful' : 'unsuccessful'
        console.log('Copying text command was ' + msg)
        if (successful) {
          if (type) that.$message({ message: msg, type: 'success' })
          else {
            that.copyClick = false
            setTimeout(function () { that.copyClick = true }, 600)
          }
          return true
        }
      } catch (err) {
        console.log('Oops, unable to copy')
      } finally {
        document.body.removeChild(txtArea)
      }
      return false
    },
    documentLink () {
      window.open('https://docs.filswan.com/multi-chain-storage/overview', '_blank')
    },
    handleCommand (command) {
      // console.log(command)
      sessionStorage.removeItem('dealsPaginationIndexDev')
      switch (command) {
        case 'setting':
          that.$router.push({ name: 'ApiKey' })
          break
        case 'bucket':
          that.$router.push({ name: 'Space' })
          break
        case 'lang':
          const language = that.languageMcs === 'en' ? 'cn' : 'en'
          that.handleSetLanguage(language)
          break
        case 'theme':
          that.reverseChange(!that.reverse)
          break
        default:
      }
    },
    reverseChange (val) {
      let reverse = val ? 1 : 0
      that.$store.dispatch('setReverse', reverse)
    },
    handleSetLanguage (lang) {
      that.loadIndexing = true

      document.body.style.height = '100vh'
      document.body.style['overflow-y'] = 'hidden'
      that.$i18n.locale = lang
      that.$store.dispatch('setLanguage', lang)

      window.location.reload()
    },
    goLink (link) {
      window.open(link)
    },
    header_logo () {
      that.$router.push({ name: 'home_entrance' })
    },
    handleSelect (key, keyPath) {
      // if (key !== 'login') that.$emit('getHome', key)
      if (key !== 'login' && key !== 'auditReport' && key !== 'logged') that.$router.push({ name: 'home_entrance', query: { id: key } })
    },
    handleMoSelect (key) {
      // that.$emit('getHome', key)
      that.$router.push({ name: 'home_entrance', query: { id: key } })
      that.mobileMenuFun()
    },
    mobileMenuFun () {
      that.mobileMenuShow = !that.mobileMenuShow
    },
    getLogin () {
      that.$emit('getLogin', true)
    },
    activeName () {
      let pathAdress = that.$route.query.id
      if (!pathAdress) that.activeIndex = '4'
      else {
        if (pathAdress.indexOf('about') > -1) {
          that.activeIndex = 'about'
        } else if (pathAdress.indexOf('pricing') > -1) {
          that.activeIndex = 'pricing'
        } else if (pathAdress.indexOf('resources') > -1) {
          that.activeIndex = 'resources'
        } else if (pathAdress.indexOf('stats') > -1) {
          that.activeIndex = 'stats'
        } else {
          that.activeIndex = '4'
        }
      }
    }
  },
  mounted () {
    that = this
    that.activeName()
    that.addrChild = that.metaAddress || ''
    window.addEventListener('resize', () => {
      if (document.body.clientWidth > 999) that.mobileMenuShow = false
    })
  },
  filters: {
    number (value) {
      let realVal = ''
      if (!isNaN(value) && value !== '' && value !== null) {
        // realVal = 0.000000000000000001 * value
        realVal = String(value).replace(/^(.*\..{18}).*$/, '$1')
      } else {
        realVal = '--'
      }
      return realVal
    },
    hiddEmail (val) {
      if (!val) return '-'
      if (val.indexOf('@') !== -1) {
        var a = val.indexOf('@')
        val = '****' + val.substring(a - 2)
        return val
      } else {
        return `${val.substring(4, 4)}****${val.substring(val.length - 4)}`
      }
    },
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
<style lang="scss" scoped>
@font-face {
  font-family: "gilroy-regular";
  src: url(../assets/font/gilroy-regular-3.otf);
  font-style: normal;
  font-display: block;
}
.headerCont {
  position: relative;
  width: 100%;
  background-color: #fff;
  font-family: "gilroy-regular";
  box-shadow: 0 1px 2px rgba(51, 51, 51, 0.2);
  z-index: 9;
  @media screen and (max-width: 999px) {
    // position: fixed;
    min-height: 60px;
    z-index: 999;
  }
  .width {
    display: flex;
    flex-wrap: wrap;
    padding: 0 16px;
    max-width: 2560px;
    min-width: 300px;
    margin: auto;
    @media screen and (min-width: 1200px) {
      padding: 0 32px;
    }
    @media screen and (min-width: 1600px) {
      width: calc(100% - 128px);
      padding: 0 64px;
    }
  }
  .header /deep/ {
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 76px;
    -webkit-transition: position 0.4s, padding 0s;
    transition: position 0.4s, padding 0s;
    @media screen and (max-width: 999px) {
      height: 60px;
    }
    .logoImg {
      height: 40px;
      cursor: pointer;
    }
    .header_left {
      float: left;
      display: flex;
      align-items: center;
      width: auto;
      margin: 0;
      font-size: 0.2772rem;
      font-weight: 500;
      color: #0b318f;
      line-height: 0.42rem;
      text-transform: uppercase;
      .meta-space {
        padding-left: 0.25rem;
        margin-left: 0.45rem;
        font-size: 17px;
        border-left: 1px solid #eee;
        color: #606060;
        text-transform: capitalize;
        @media screen and (max-width: 1680px) {
          font-size: 16px;
        }
        @media screen and (max-width: 1600px) {
          font-size: 15px;
        }
        @media screen and (max-width: 1440px) {
          font-size: 14px;
        }
        &:hover,
        &.active {
          // text-decoration: underline;
          color: #5580e9;
        }
      }
      .logoImg {
        position: relative;
        height: 40px;
        cursor: pointer;
        @media screen and (max-width: 992px) {
          height: 30px;
        }
        .img {
          height: 100%;
        }
        .beta {
          position: absolute;
          width: 40px;
          right: -0.2rem;
          bottom: 0.05rem;
          @media screen and (min-width: 1800px) {
            width: 45px;
            bottom: 0.03rem;
          }
          @media screen and (max-width: 1366px) {
            right: -0.25rem;
          }
          @media screen and (max-width: 1280px) {
            right: -0.3rem;
          }
          @media screen and (max-width: 600px) {
            bottom: 0.02rem;
          }
        }
      }
    }
    .menuChildA {
      display: block;
      width: 100%;
    }
    .el-menu.el-menu--horizontal {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      border: 0;
    }
    .el-menu--horizontal > .el-menu-item,
    .el-menu--horizontal > .el-submenu .el-submenu__title {
      height: auto;
      padding: 0;
      margin: 0.05rem 0.05rem 0.05rem 0.2rem;
      line-height: 1;
      font-size: 17px;
      border: 0;
      @media screen and (max-width: 1680px) {
        font-size: 16px;
      }
      @media screen and (max-width: 1600px) {
        font-size: 15px;
      }
      @media screen and (max-width: 1440px) {
        font-size: 14px;
      }
      @media screen and (max-width: 999px) {
        margin: 0.05rem 0.15rem 0.05rem 0.1rem;
      }
      a {
        display: block;
        padding: 0.15rem;
        color: #606060;
        &.target {
          float: left;
          padding: 0.15rem 0.4rem;
          background: #5580e9 !important;
          border: 0;
          border-radius: 0.05rem;
          outline: none;
          color: #fff !important;
          // box-shadow: 0 12px 26px -12px rgba(12, 22, 44, 0.32);
          transition: background-color 0.3s, border-color 0.3s, color 0.3s;
          &:hover {
            opacity: 0.85;
          }
        }
        .el-loading-mask {
          border-radius: 0.05rem;
          z-index: 2;
        }
      }
      &.is-active,
      &:hover {
        background-color: transparent !important;
        color: #5580e9 !important;
        a {
          color: #5580e9;
        }
      }
      .el-submenu__icon-arrow {
        display: none;
      }
    }
    .feh-metamask {
      display: flex;
      align-items: center;
      position: relative;
      width: auto;
      padding: 0.15rem 0.3rem;
      margin: 0;
      cursor: pointer;
      border-radius: 0.14rem;
      background: linear-gradient(45deg, #4f8aff, #4b5eff);
      @media screen and (max-width: 991px) {
        margin: 0;
      }
      img {
        display: block;
        width: auto;
        height: 0.3rem;
        cursor: pointer;
        margin: 0 0.1rem 0 0;
      }
      // &:before{
      //     position: absolute;
      //     left: 30px;
      //     top: -4px;
      //     content: "";
      //     width: 5px;
      //     height: 5px;
      //     border-radius: 100%;
      //     background: #d7d6d6;
      // }
      .logged_in {
        display: flex;
        align-items: center;
        font-size: 0.2rem;
        color: #333;
        h3,
        h4,
        h5 {
          font-size: inherit;
          font-weight: normal;
          cursor: pointer;
        }
        span {
          font-size: inherit;
        }
        .text,
        .el-button {
          padding: 0;
          font-size: inherit;
          color: #fff;
          line-height: 1.36;
          cursor: pointer;
          @media screen and (max-width: 600px) {
            font-size: 12px;
          }
        }
        .el-button {
          padding: 0.08rem 0.15rem;
          margin: 0 0 0 0.2rem;
          line-height: 2;
          color: #fff;
          cursor: pointer;
          border: 0;
          font-weight: normal;
          font-family: inherit;
          background: linear-gradient(45deg, #4f8aff, #4b5eff);
          border-radius: 0.14rem;
          @media screen and (max-width: 600px) {
            padding: 0.08rem 0.15rem;
            margin: 0 0 0 10px;
          }
          &:hover {
            opacity: 0.9;
          }
        }
        .info {
          display: flex;
          align-items: center;
          margin: 0;
          line-height: 1.5;
          color: #fff;
          font-size: inherit;
          cursor: text;
          border-radius: 0.14rem;
          background: linear-gradient(45deg, #4f8aff, #4b5eff);
          h5 {
            padding: 0.08rem 0.1rem;
          }
          h4 {
            padding: 0 0.12rem;
            margin-left: 0.1rem;
            background: #2d43e7;
            font-size: 18px;
            line-height: 2;
            border-radius: 0.1rem;
            cursor: pointer;
            @media screen and (max-width: 600px) {
              margin: 0;
            }
            &:hover {
              opacity: 0.9;
            }
          }
          .dropdown-style {
            padding: 0.15rem;
            // background: #2d43e7;
            font-size: 18px;
            line-height: 2;
            // border-radius: 0.1rem;
            border-left: 1px solid #2d43e7;
            cursor: pointer;
            @media screen and (min-width: 1800px) {
              font-size: 20px;
            }
            @media screen and (max-width: 768px) {
              line-height: 1.5;
            }
            &:hover {
              opacity: 0.9;
            }
            .el-dropdown-link {
              color: #fff;
              i {
                display: flex;
                width: auto;
                padding: 0;
                margin: 0;
                font-size: inherit;
                color: #fff;
              }
            }
          }
          @media screen and (max-width: 600px) {
            font-size: 12px;
          }
        }
      }
    }
    .online {
      // padding: 0.08rem 0.1rem;
      padding: 0;
      border-radius: 0;
      background: transparent;
      &:before {
        // background: #0fce7c;
        display: none;
      }
    }
  }
}
.loadIndexStyle {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  z-index: 2000;
  background: rgba(255, 255, 255, 1);
}

@media screen and (max-width: 999px) {
  .header {
    .pcShow {
      display: none !important;
    }
    .mobileShow {
      display: block;
      width: 27px;
      float: right;
      margin-top: 0;
      outline: none;
      cursor: pointer;
      span {
        display: block;
        width: 100%;
        height: 2px;
        margin: 0.07rem 0;
        background-color: #2c367a;
      }
      .close {
        height: 0.35rem;
        display: flex;
        align-items: center;
        justify-content: flex-end;
        i {
          font-size: 20px;
          color: #000;
        }
        .target {
          position: absolute;
          right: 55px;
          height: auto;
          padding: 0 10px;
          font-size: 14px;
          line-height: 1.5;
          border-radius: 15px;
        }
      }
    }
    .menuBg {
      position: fixed;
      right: 0;
      left: 0;
      top: 60px;
      width: 100%;
      height: 100%;
      background: rgba(0, 0, 0, 0.3);
    }
    .menuMList {
      position: fixed;
      right: 0;
      left: auto;
      top: 60px;
      bottom: 0;
      width: 0;
      height: calc(100% - 60px);
      z-index: 99;
      background-color: #fff;
      overflow-x: hidden;
      overflow-y: scroll;
      -webkit-transition: all 0.4s;
      transition: all 0.4s;
      border-top: 1px solid;
      .contentMenu {
        width: 200%;
        -webkit-transition: transform 0.3s;
        transition: transform 0.3s;
        .first_level_menu,
        .secondary_menu {
          float: left;
          width: 50%;
          .menuMListChild {
            display: flex;
            justify-content: space-between;
            align-items: center;
            position: relative;
            width: 90%;
            float: none;
            text-align: left;
            font-size: 17px;
            line-height: 1.5;
            padding: 0.2rem 2%;
            margin: 0 0 0 6%;
            color: #000;
            border-bottom: 1px solid #e4e7ed;
            @media screen and (max-width: 1680px) {
              font-size: 16px;
            }
            @media screen and (max-width: 1600px) {
              font-size: 15px;
            }
            @media screen and (max-width: 1440px) {
              font-size: 14px;
            }
            i {
              float: right;
              color: #101c4f;
              margin: 0 4% 0 0;
              cursor: pointer;
              font-size: 16px;
              padding: 0.1rem 0;
              &:before {
                font-size: 16px;
                font-weight: 500;
              }
            }
            h5 {
              margin: 0 4% 0 0;
              font-size: 0.17rem;
              color: #000;
            }
          }
          .m {
            color: #9baff9;
          }
        }
      }
      .contentMenuX {
        transform: translateX(-50%);
      }
    }
    .menuWidth {
      width: 90%;
    }
    .logoImg {
      height: 30px;
    }
  }
}
@media screen and (max-width: 479px) {
  .header {
    padding: 25px 24px;
  }
}
.el-dialog__wrapper /deep/ {
  display: flex;
  align-items: center;
  .wrongNet {
    margin: auto !important;
    box-shadow: 0 0 13px rgba(128, 128, 128, 0.8);
    border-radius: 0.2rem;
    .el-dialog__header {
      padding: 0.2rem 0.4rem;
      display: flex;
      align-items: center;
      justify-content: space-between;
      border-bottom: 1px solid #dfdfdf;
      color: #000;
      font-size: 0.22rem;
      font-weight: 500;
      line-height: 1;
      text-transform: capitalize;
      @media screen and (max-width: 479px) {
        padding: 0.3rem 0.2rem;
      }
      .el-dialog__headerbtn {
        position: relative;
        top: auto;
        right: auto;
        font-size: inherit;
        i {
          font-size: inherit;
          &:hover {
            color: #0b318f;
          }
        }
      }
      .el-dialog__title {
        font-size: inherit;
      }
    }
    .el-dialog__body {
      position: relative;
      padding: 0.3rem 0.4rem 0.4rem;
      font-size: 0.2rem;
      @media screen and (max-width: 479px) {
        padding: 0.2rem;
      }
      label {
        word-break: break-word;
        line-height: 1;
        color: #666;
        font-size: inherit;
      }
      .address {
        background: rgba(233, 233, 233, 1);
        padding: 8px;
        margin: 10px 0 22px;
        border-radius: 8px;
        font-size: inherit;
      }
      .address_email {
        margin: 0 0 10px;
        .address_body {
          display: flex;
          align-items: center;
          justify-content: space-between;
          margin: 10px 0 0;
          .address {
            width: 80%;
            margin: 0;
          }
          .address_right {
            position: relative;
            display: inline-block;
            padding: 0.05rem 0.2rem 0.05rem 0.32rem;
            margin: 0 5px;
            background-color: rgba(85, 128, 233, 0.15);
            font-size: 0.148rem;
            border-radius: 0.5rem;
            white-space: nowrap;
            @media screen and (max-width: 1600px) {
              font-size: 13px;
            }
            @media screen and (max-width: 600px) {
              font-size: 12px;
            }
            &::before {
              position: absolute;
              left: 0.16rem;
              top: 50%;
              content: "";
              width: 0.08rem;
              height: 0.08rem;
              margin-top: -0.04rem;
              background-color: #606266;
              border-radius: 0.5rem;
            }
          }
          .bg-primary {
            &::before {
              background-color: #4d73ff;
            }
          }
        }
        .share {
          .el-button {
            width: 100%;
            margin: 3px 0 0;
            font-size: 13px;
            @media screen and (min-width: 1800px) {
              font-size: 14px;
            }
            @media screen and (max-width: 600px) {
              font-size: 12px;
            }
          }
        }
        .el-loading-mask {
          .el-loading-spinner {
            top: 50%;
          }
        }
      }
      .share {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        justify-content: flex-start;
        font-size: inherit;
        .el-button {
          min-width: 50%;
          padding: 0;
          margin: 8px 0 0;
          background: transparent !important;
          border: 0;
          color: #4f7bf5;
          font-size: inherit;
          font-weight: normal;
          font-family: inherit;
          opacity: 0.8;
          span {
            display: flex;
            align-items: center;
            svg {
              width: 15px;
              height: 15px;
              margin: 0 3px 0 0;
            }
            .icon_big {
              width: 13px;
              height: 13px;
            }
          }
          &:hover {
            background: transparent;
            opacity: 1;
          }
        }
      }
      .loadStyle {
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        z-index: 2000;
        background: rgba(255, 255, 255, 1);
        border-radius: 0.2rem;
      }
      .apiTipCont {
        p {
          display: flex;
          align-items: center;
          text-indent: 0.1rem;
          margin: 0.1rem;
          color: #7e7e7e;
          font-size: 0.18rem;
          .el-icon-document-copy {
            display: block;
            font-size: 17px;
            cursor: pointer;
          }
        }
      }
    }
  }
}
</style>

<style lang="scss">
.dropdown-ul {
  min-width: 200px;
  padding: 0.1rem;
  font-family: "gilroy-regular";
  border: 0.01rem solid #eee;
  border-bottom: 0;
  box-shadow: 0 0 0.1rem rgba(187, 180, 180, 0.51);
  .el-dropdown-menu__item {
    color: #333;
    padding: 0.1rem 0.15rem;
    margin: 0.05rem 0;
    font-size: 14px;
    line-height: 2;
    border-radius: 5px;
    @media screen and (min-width: 1800px) {
      font-size: 17px;
    }
    @media screen and (max-width: 1680px) {
      font-size: 16px;
    }
    @media screen and (max-width: 1600px) {
      font-size: 15px;
    }
    @media screen and (max-width: 1440px) {
      font-size: 14px;
    }
    &:focus,
    &:hover {
      color: #333;
    }
    .dropdown-body {
      justify-content: space-between;
      .dropdown-left {
        min-width: 100px;
        margin-right: 0.2rem;
        i,
        svg {
          font-size: inherit;
          position: relative;
          width: 16px;
          height: 16px;
          margin-right: 5px;
          @media screen and (min-width: 1800px) {
            width: 20px;
            height: 20px;
          }
          &::before {
            position: absolute;
            left: 0;
            top: 0;
            width: 100%;
            height: 100%;
            content: "";
            background-size: 16px !important;
            @media screen and (min-width: 1800px) {
              background-size: 20px !important;
            }
          }
        }
        .el-icon-switch-button {
          font-size: 17px;
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
            background: url(../assets/images/menuIcon/browse.png) no-repeat
              center;
          }
        }
        .el-icon-s-deal {
          &::before {
            background: url(../assets/images/menuIcon/icon_Files@2x.png)
              no-repeat center;
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
            background: url(../assets/images/menuIcon/icon_documentation@2x-2.png)
              no-repeat center;
          }
        }
        .el-icon-s-metaSpace {
          &::before {
            background: url(../assets/images/menuIcon/metaSpace-3.png) no-repeat
              center;
          }
        }
        .el-icon-language {
          &::before {
            background: url(../assets/images/menuIcon/language-2.png) no-repeat
              center;
            background-size: 100% !important;
          }
        }
        .el-icon-theme {
          &::before {
            background: url(../assets/images/menuIcon/theme.png) no-repeat
              center;
            background-size: 100% !important;
          }
        }
        .el-icon-s-dataset {
          &::before {
            background: url(../assets/images/menuIcon/dataset.png) no-repeat
              center;
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
            background: url(../assets/images/menuIcon/myAccount-2.png) no-repeat
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
      .lang_style {
        min-width: 0.26rem;
        margin: 0;
        line-height: 0.26rem;
        font-size: 0.14rem;
        font-weight: 500;
        color: #000;
        text-align: center;
        @media screen and (min-width: 1800px) {
          font-size: 17px;
        }
        @media screen and (max-width: 1680px) {
          font-size: 16px;
        }
        @media screen and (max-width: 1600px) {
          font-size: 15px;
        }
        @media screen and (max-width: 1440px) {
          font-size: 14px;
        }
        @media screen and (max-width: 991px) {
          font-size: 13px;
          width: 25px;
          height: 25px;
          line-height: 25px;
        }
        span {
          cursor: pointer;
          color: inherit;
          &:hover {
            color: #4b5eff;
          }
        }
      }
      .el-switch {
        // margin-left: 0.1rem;
        height: 20px;
        .el-switch__core {
          height: 20px;
        }
        .el-switch__core:after {
          background-color: #fff;
          top: 0;
          width: 18px;
          height: 18px;
          z-index: 10;
        }
        .el-switch__label--left {
          position: absolute;
          top: 0;
          left: 5px;
          margin: 0;
          color: #fff;
          font-size: 15px;
          z-index: 9;
          i {
            font-size: inherit;
          }
        }
        .el-switch__label--right {
          position: absolute;
          top: 0;
          right: 5px;
          margin: 0;
          color: #fff;
          font-size: 15px;
          z-index: 9;
          i {
            font-size: inherit;
          }
        }
      }
      .el-switch.is-checked {
        .el-switch__core::after {
          margin-left: -20px;
        }
      }
      .swithUI {
        cursor: pointer;
        font-size: 20px;
        svg {
          box-sizing: border-box;
          margin: 0px;
          min-width: 0px;
          font-size: 20px;
          cursor: pointer;
          width: 20px;
          height: 20px;
        }
        .css-sunny {
          color: #fff;
          &:hover {
            color: #4f8aff;
          }
        }
        .css-moon {
          color: #333;
          &:hover {
            color: #4f8aff;
          }
        }
      }
      .switch {
        position: relative;
        display: flex;
        margin: 0;
      }
    }
  }
  .mobileShow {
    display: none;
    @media screen and (max-width: 999px) {
      display: block;
    }
  }
}
</style>
