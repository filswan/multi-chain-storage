<template>
  <div class="header" :class="{'content-collapse': collapseLocal}">
    <div class="header_arera">
      <div class="header-right">
        <div class="network_mainnet" v-if="addrChild" :class="{'error': networkTip}" @click="networkC=true">
          <!-- <div class="BSC_mainnet" v-if="networkID == 97" title="BSC TestNet mainnet">
            <img src="@/assets/images/network_logo/bsc.png" /> {{bodyWidth?'BSC':'BSC TestNet'}}
          </div> -->
          <div class="Polygon_mainnet" v-if="networkID == 137" title="Polygon mainnet">
            <img src="@/assets/images/network_logo/polygon.png" /> {{bodyWidth?'Polygon':'Polygon Mainnet'}}
          </div>
          <div class="Mumbai_mainnet" v-else-if="networkID == 80001" title="Mumbai Testnet mainnet">
            <img src="@/assets/images/network_logo/polygon.png" /> {{bodyWidth?'Mumbai':'Mumbai Testnet'}}
          </div>
          <div class="Mumbai_mainnet" v-else :title="metaNetworkInfo.name+' mainnet'">
            <span></span>
            {{metaNetworkInfo.name}}
          </div>
        </div>
        <div :class="{'online': addrChild, 'feh-metamask': 1==1}">
          <div v-if="!addrChild" class="logged_in filter_status">
            <el-tooltip class="item" effect="dark" :content="$t('fs3Login.toptip_03')" placement="bottom">
              <img src="@/assets/images/metamask.png" @click="metamaskLogin" />
            </el-tooltip>
            <span class="text" @click="metamaskLogin">{{$t('fs3.Connect_Wallet')}}</span>
          </div>

          <div v-else class="logged_in">
            <!-- <span class="text textTrue">{{metaNetworkInfo.name}}</span> -->
            <div class="info">
              <h5>{{priceAccound}} {{ metaNetworkInfo.unit}}</h5>
              <h4 @click="wrongVisible = true">{{addrChild | hiddAddress}}</h4>
            </div>
            <el-button class="text textTrue pcShow" @click="signOutFun">{{$t('fs3.Disconnect')}}</el-button>
          </div>
        </div>
        <div class="lang_style">
          <span v-if="languageMcs === 'en'" @click="handleSetLanguage('cn')">EN</span>
          <span v-else @click="handleSetLanguage('en')">中</span>
        </div>
        <div class="switch">
          <div class="swithUI" v-if="reverse" @click="reverseChange(false)">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" class="css-sunny">
              <path fill-rule="evenodd" clip-rule="evenodd" d="M10.5 2h3v3h-3V2zM16 12a4 4 0 11-8 0 4 4 0 018 0zM5.99 3.869L3.867 5.99 5.99 8.112 8.111 5.99 5.989 3.87zM2 13.5v-3h3v3H2zm1.868 4.51l2.121 2.12 2.122-2.12-2.122-2.122-2.121 2.121zM13.5 19v3h-3v-3h3zm4.51-3.112l-2.121 2.122 2.121 2.121 2.121-2.121-2.121-2.122zM19 10.5h3v3h-3v-3zm-3.11-4.51l2.12 2.121 2.122-2.121-2.121-2.121-2.122 2.121z"
                fill="currentColor"></path>
            </svg>
          </div>
          <div class="swithUI" v-else @click="reverseChange(true)">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" class="css-moon">
              <path d="M20.968 12.768a7 7 0 01-9.735-9.735 9 9 0 109.735 9.735z" fill="currentColor"></path>
            </svg>
          </div>
        </div>
        <!-- mobile显示 -->
        <div class="mobileShow">
          <div class="collapse-btn-cont" @click="collapseChage">
            <div class="header_btn" :class="{'el-icon-s-unfold': !collapseLocal, 'el-icon-s-fold': collapseLocal}">
              <!-- <span></span> -->
            </div>
          </div>
        </div>
      </div>
    </div>

    <el-dialog :title="$t('fs3Login.Account')" :visible.sync="wrongVisible" :width="width" custom-class="wrongNet">
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

    <network-change v-if="networkC" :networkC="networkC" @getNetworkC="getNetworkC"></network-change>
  </div>
</template>
<script>
import networkChange from '@/components/networkChange'
// import axios from 'axios'
import erc20ContractJson from '@/utils/ERC20.json'
const ethereum = window.ethereum
let contractErc20
let that
export default {
  components: {
    networkChange
  },
  data () {
    return {
      collapseLocal: !!(this.$store.getters.collapseL === 'true' || this.$store.getters.collapseL === true),
      collapse: document.body.clientWidth < 1024,
      fullscreen: false,
      name: 'linxin',
      message: 2,
      langNew: '',
      userShow: false,
      loginShow: !localStorage.getItem('mcsLoginAccessToken'),
      bodyWidth: document.body.clientWidth < 992,
      langShow: true,
      loadIndexing: false,
      // 控制是否在路由栈中清理当前页面的数据
      replaceData: null,
      tools: '1',
      tabOaxLogin: localStorage.getItem('mcsLoginAccessToken'),
      tabOaxNew: localStorage.getItem('mcsLoginAccessToken'),
      priceAccound: 0,
      network: {
        name: '',
        unit: '',
        center_fail: false
      },
      addrChild: '',
      wrongVisible: false,
      width: document.body.clientWidth > 600 ? '450px' : '95%',
      copyClick: true,
      switchWidth: 56,
      reverseSwitch: false,
      networkC: false,
      prevType: true,
      loadAccount: false
    }
  },
  props: ['meta', 'netId', 'networkTip'],
  computed: {
    languageMcs () {
      return this.$store.getters.languageMcs
    },
    collapseL () {
      return this.$store.getters.collapseL
    },
    metaAddress () {
      return this.$store.getters.metaAddress
    },
    metaNetworkInfo () {
      return this.$store.getters.metaNetworkInfo ? JSON.parse(JSON.stringify(this.$store.getters.metaNetworkInfo)) : {}
    },
    reverse () {
      return String(this.$store.getters.reverse) === '1'
    },
    free_usage () {
      return this.$store.getters.free_usage
    },
    free_quota_per_month () {
      return this.$store.getters.free_quota_per_month
    },
    networkID () {
      return Number(this.$store.getters.networkID)
    }
  },
  watch: {
    $route: function (to, from) {
      if ((to.name === 'register' && from.name === 'login') || (to.name === 'login' && from.name === 'register')) {
        this.replaceData = 'replace'
      } else {
        this.replaceData = ''
      }
    },
    'collapseL': function () {
      this.collapseLocal = !!(this.$store.getters.collapseL === 'true' || this.$store.getters.collapseL === true)
    },
    metaAddress: function () {
      this.addrChild = this.metaAddress
      this.commonParam()
      this.walletInfo()
    },
    meta: function () {
      this.walletInfo()
    },
    netId: function () {
      //  this.netId === 97 ||
      if (this.netId === 137 || this.netId === 80001) this.getNetworkC(false, this.netId)
    }
  },
  methods: {
    getNetworkC (dialog, rows) {
      that.networkC = dialog
      if (rows) {
        let text = {}
        switch (rows) {
          case 80001:
            text = {
              chainId: that.$web3Init.utils.numberToHex(80001),
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
          // case 97:
          //   text = {
          //     chainId: that.$web3Init.utils.numberToHex(97),
          //     chainName: 'BSC TestNet',
          //     nativeCurrency: {
          //       name: 'tBNB',
          //       symbol: 'tBNB', // 2-6 characters long
          //       decimals: 18
          //     },
          //     rpcUrls: ['https://data-seed-prebsc-1-s1.binance.org:8545'],
          //     blockExplorerUrls: ['https://testnet.bscscan.com']
          //   }
          //   break
          case 137:
            text = {
              chainId: that.$web3Init.utils.numberToHex(137),
              chainName: 'Polygon Mainnet',
              nativeCurrency: {
                name: 'MATIC',
                symbol: 'MATIC', // 2-6 characters long
                decimals: 18
              },
              rpcUrls: ['https://polygon-rpc.com'],
              blockExplorerUrls: ['https://polygonscan.com/']
            }
            break
        }
        ethereum.request({
          method: 'wallet_addEthereumChain',
          params: [
            text
          ]
        }).then((res) => {
          // 添加成功
          // that.changeChaid(rows)
          that.$emit('getNetId', 0)
        }).catch((err) => {
          console.log(err)
          that.$emit('getNetId', 0)
        })
      }
    },
    async changeChaid (rows) {
      that.$store.dispatch('setMetaNetworkId', rows)
      let status = await that.$metaLogin.netStatus(rows)
      if (!status) {
        let link = that.baseNetwork ? 'https://calibration-mcs.filswan.com' : 'https://www.multichain.storage'
        if (that.networkID === 80001 || that.networkID === 137) window.open(link)
        that.signOutFun()
        return false
      } else {
        const lStatus = await that.$metaLogin.login()
        if (lStatus) setTimeout(function () { window.location.reload() }, 200)
      }
      if (that.$route.name === 'my_files_detail') {
        that.$router.push({ path: '/my_files' })
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
    // 侧边栏折叠
    collapseChage () {
      that.collapseLocal = !that.collapseLocal
      that.$store.dispatch('setCollapse', that.collapseLocal)
      // bus.$emit('collapse', that.collapse);
    },
    handleSetLanguage (lang) {
      that.loadIndexing = true

      document.body.style.height = '100vh'
      document.body.style['overflow-y'] = 'hidden'
      that.$i18n.locale = lang
      that.$store.dispatch('setLanguage', lang)

      window.location.reload()
    },
    getHiddenProp () {
      var prefixes = ['webkit', 'moz', 'ms', 'o']

      // if 'hidden' is natively supported just return it
      if ('hidden' in document) return 'hidden'

      // otherwise loop over all the known prefixes until we find one
      for (var i = 0; i < prefixes.length; i++) {
        if ((prefixes[i] + 'Hidden') in document) { return prefixes[i] + 'Hidden' }
      }

      // otherwise it's not supported
      return null
    },
    getVisibilityState () {
      var prefixes = ['webkit', 'moz', 'ms', 'o']
      if ('visibilityState' in document) return 'visibilityState'
      for (var i = 0; i < prefixes.length; i++) {
        if ((prefixes[i] + 'VisibilityState') in document) { return prefixes[i] + 'VisibilityState' }
      }
      // otherwise it's not supported
      return null
    },
    metamaskLogin () {
      if (!that.metaAddress || that.metaAddress === undefined) {
        that.$commonFun.Init(addr => {
          that.$nextTick(async () => {
            that.$store.dispatch('setMetaAddress', addr)
            let status = await that.$metaLogin.netStatus(that.networkID)
            that.$emit('getNetwork', !status)
            if (!status) return false
            const lStatus = await that.$metaLogin.login()
            if (lStatus) that.$emit('getMetamaskLogin', true)
          })
        })
        return false
      }
    },
    // Wallet address
    signFun (redirect) {
      if (!that.addrChild) {
        ethereum
          .request(
            {
              'jsonrpc': '2.0',
              'method': 'eth_accounts',
              'params': [that.metaAddress, 'latest'],
              'id': 19998
            }
          )
          .then((accounts) => {
            if (accounts.length <= 0) {
              that.signOutFun()
              return false
            }
            that.addrChild = that.metaAddress
            that.walletInfo()
          })
          .catch((error) => {
            console.log(error)
            that.signOutFun()
          })
        return false
      }
    },
    walletInfo () {
      if (!that.addrChild || that.addrChild === 'undefined') {
        return false
      }

      ethereum
        .request({ method: 'eth_chainId' })
        .then(async (chainId) => {
          let netId = parseInt(chainId, 16)
          // console.log('network ID:', netId)
          // console.log(`decimal number: ${parseInt(chainId, 16)}`);
          that.$store.dispatch('setMetaNetworkId', netId)

          await that.contractPrice(netId)

          switch (netId) {
            case 1:
              that.network.name = 'mainnet'
              that.network.unit = 'ETH'
              that.network.center_fail = true
              that.$store.dispatch('setMetaNetworkInfo', that.network)
              break
            case 3:
              that.network.name = 'ropsten'
              that.network.unit = 'ETH'
              that.network.center_fail = true
              that.$store.dispatch('setMetaNetworkInfo', that.network)
              break
            case 4:
              that.network.name = 'rinkeby'
              that.network.unit = 'ETH'
              that.network.center_fail = true
              that.$store.dispatch('setMetaNetworkInfo', that.network)
              break
            case 5:
              that.network.name = 'goerli'
              that.network.unit = 'ETH'
              that.network.center_fail = true
              that.$store.dispatch('setMetaNetworkInfo', that.network)
              break
            case 42:
              that.network.name = 'kovan'
              that.network.unit = 'ETH'
              that.network.center_fail = true
              that.$store.dispatch('setMetaNetworkInfo', that.network)
              break
            case 56:
              that.network.name = 'BSC'
              that.network.unit = 'BNB'
              that.network.center_fail = true
              that.$store.dispatch('setMetaNetworkInfo', that.network)
              break
            case 97:
              that.network.name = 'BSC'
              that.network.unit = 'USDC'
              that.network.center_fail = false
              that.$store.dispatch('setMetaNetworkInfo', that.network)
              if (that.meta) {
                if (that.$route.query.redirect && that.$route.query.redirect !== '/supplierAllBack') {
                  // 防止登录后需要跳转到指定页面
                  that.$router.push({ path: that.$route.query.redirect })
                } else {
                  that.$router.push({ path: '/my_buckets' })
                }
                window.location.reload()
                that.$emit('getMetamaskLogin', false)
              }
              break
            case 137:
              that.network.name = 'Polygon'
              that.network.unit = 'USDC'
              that.network.center_fail = true
              that.$store.dispatch('setMetaNetworkInfo', that.network)
              if (that.meta) {
                if (that.$route.query.redirect && that.$route.query.redirect !== '/supplierAllBack') {
                  // 防止登录后需要跳转到指定页面
                  that.$router.push({ path: that.$route.query.redirect })
                } else {
                  that.$router.push({ path: '/my_files' })
                }
                window.location.reload()
                that.$emit('getMetamaskLogin', false)
              }
              break
            case 999:
              that.network.name = 'NBAI'
              that.network.unit = 'NBAI'
              that.network.center_fail = true
              that.$store.dispatch('setMetaNetworkInfo', that.network)
              break
            case 80001:
              that.network.name = 'Mumbai'
              that.network.unit = 'USDC'
              that.network.center_fail = false
              that.$store.dispatch('setMetaNetworkInfo', that.network)
              if (that.meta) {
                if (that.$route.query.redirect && that.$route.query.redirect !== '/supplierAllBack') {
                  // 防止登录后需要跳转到指定页面
                  that.$router.push({ path: that.$route.query.redirect })
                } else {
                  that.$router.push({ path: '/my_files' })
                }
                window.location.reload()
                that.$emit('getMetamaskLogin', false)
              }
              break
            default:
              that.network.name = 'Custom'
              that.network.unit = ''
              that.network.center_fail = true
              that.$store.dispatch('setMetaNetworkInfo', that.network)
              break
          }

          let status = await that.$metaLogin.netStatus(that.networkID)
          that.$emit('getNetwork', !status)
        })
        .catch((error) => {
          console.error(`Error fetching chainId: ${error.code}: ${error.message}`)
        })
    },
    fn () {
      ethereum.on('accountsChanged', function (account) {
        // console.log('account header:', account[0]);  //Once the account is switched, it will be executed here
        if (that.prevType) that.signOutFun()
      })
      // networkChanged
      ethereum.on('chainChanged', function (accounts) {
        if (!that.prevType) return false
        that.$store.dispatch('setMCSjwtToken', '')
        that.$store.dispatch('setMetaNetworkId', parseInt(accounts, 16))
        that.walletInfo()
        that.changeChaid(parseInt(accounts, 16))
        if (that.$route.name === 'my_files_detail') that.$router.push({ path: '/my_files' })
      })
      // 监听metamask网络断开
      ethereum.on('disconnect', (code, reason) => {
        // console.log(`Ethereum Provider connection closed: ${reason}. Code: ${code}`);
      })
    },
    signOutFun () {
      let params = {}
      axios.post(`${that.baseAPIURL}api/v1/user/logout_for_metamask_signature`, params, {
        headers: {
          'Authorization': 'Bearer ' + that.$store.getters.mcsjwtToken
        }
      }).then((response) => {
        if (response.data.status !== 'success') that.$message.error(response.data.message || 'Fail')
        that.addrChild = ''
        that.$store.dispatch('setMetaAddress', '')
        that.$store.dispatch('setMCSjwtToken', '')
        that.$store.dispatch('setMetaNetworkId', 0)
        that.network.name = ''
        that.network.unit = ''
        that.network.center_fail = false
        that.$store.dispatch('setMetaNetworkInfo', JSON.stringify(that.network))
        sessionStorage.removeItem('login_path')
        sessionStorage.removeItem('mcs_dev_jwtToken')
        // that.$router.push('/home')
        setTimeout(function () { window.location.reload() }, 200)
      }).catch(function (error) {
        console.log(error.config)
      })
    },
    commonParam () {
      let commonApi = `${that.baseAPIURL}api/v1/common/system/params?limit=20&wallet_address=${that.metaAddress}`

      axios.get(commonApi, {
        headers: {
          'Authorization': 'Bearer ' + that.$store.getters.mcsjwtToken
        }
      })
        .then((json) => {
          if (json.data.status === 'success') {
            that.$root.LOCK_TIME = json.data.data.lock_time
            that.$root.PAY_GAS_LIMIT = json.data.data.gas_limit
            that.$root.PAY_WITH_MULTIPLY_FACTOR = json.data.data.pay_multiply_factor
            that.$root.RECIPIENT = json.data.data.payment_recipient_address
            that.$root.SWAN_PAYMENT_CONTRACT_ADDRESS = json.data.data.payment_contract_address
            that.$root.USDC_ADDRESS = json.data.data.usdc_address
            that.$root.MINT_CONTRACT = json.data.data.default_nft_collection_address
            that.$root.dao_threshold = json.data.data.dao_threshold
            that.$root.filecoin_price = json.data.data.filecoin_price
            that.$root.COLLECTION_FACTORY_ADDRESS = json.data.data.nft_collection_factory_address
          }
        }).catch(error => {
          console.log(error)
        })
    },
    contractPrice (netId) {
      try {
        if (netId !== 80001 && netId !== 97 && netId !== 137) {
          ethereum
            .request(
              {
                'jsonrpc': '2.0',
                'method': 'eth_getBalance',
                'params': [that.addrChild, 'latest'],
                'id': 19999
              }
            )
            .then((balance) => {
              let balanceAll = that.$web3Init.utils.fromWei(balance, 'ether')
              // console.log('balance', balanceAll)
              that.priceAccound = Number(balanceAll).toFixed(0)
            })
            .catch((error) => {
              console.error(`Error fetching getBalance: ${error.code}: ${error.message}`)
              that.priceAccound = 0
            })
        } else {
          if (that.$root.SWAN_PAYMENT_CONTRACT_ADDRESS) {
            // 授权代币
            contractErc20 = new that.$web3Init.eth.Contract(erc20ContractJson)
            contractErc20.options.address = that.$root.USDC_ADDRESS
            // 查询剩余代币余额为：
            contractErc20.methods.balanceOf(that.metaAddress).call()
              .then(balance => {
                let usdcAvailable = that.$web3Init.utils.fromWei(balance, 'mwei')
                // console.log('Available balance:', usdcAvailable, balance)
                // that.priceAccound = that.formatDecimal(usdcAvailable, 3)
                // that.priceAccound = Number(usdcAvailable).toFixed(0)
                that.priceAccound = parseInt(usdcAvailable)
              })
          } else {
            setTimeout(function () {
              that.contractPrice(netId)
            }, 1000)
          }
        }
      } catch (err) {
        console.error(err)
        that.priceAccound = 0
      }
    },
    formatDecimal (num, decimal) {
      num = num.toString()
      let index = num.indexOf('.')
      if (index !== -1) {
        num = num.substring(0, decimal + index + 1)
      } else {
        num = num.substring(0)
      }
      return parseFloat(parseFloat(num).toFixed(decimal))
    },
    reverseChange (val) {
      let reverse = val ? 1 : 0
      that.$store.dispatch('setReverse', reverse)
    }
  },
  mounted () {
    that = this
    that.reverseSwitch = that.reverse

    var visProp = that.getHiddenProp()
    if (visProp) {
      var evtname = visProp.replace(/[H|h]idden/, '') + 'visibilitychange'
      document.addEventListener(evtname, function () {
        that.tabOaxLogin = document[that.getVisibilityState()]
      }, false)
    }

    that.commonParam()
    if (that.metaAddress) {
      that.addrChild = that.metaAddress
      that.walletInfo()
    }
    that.fn()

    document.addEventListener('visibilitychange', function () {
      that.prevType = !document.hidden
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
<style  lang="scss" scoped>
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
.addressInfo {
  padding: 0.2rem;
  h6 {
    margin: 0.14rem 0 0;
    font-size: 0.13rem;
    font-weight: normal;
    padding: 0 0.07rem;
  }
  h5 {
    font-size: 0.13rem;
    font-weight: normal;
    padding: 0 0.07rem;
  }
  h4 {
    font-size: 0.14rem;
    font-weight: normal;
    padding: 0 0.07rem;
  }
  h3 {
    margin: 0 0 0.05rem;
    font-size: 0.14rem;
    font-weight: normal;
    padding: 0 0.07rem;
    cursor: pointer;
    &:hover {
      color: rgba(11, 49, 143, 1);
    }
  }
  .el-divider /deep/ {
    margin: 0.14rem 0;
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
.header {
  box-sizing: border-box;
  height: 1.1rem;
  font-size: 0.2rem;
  color: #fff;
  background-color: #fff;
  -webkit-transition: left 0.3s ease-in-out;
  transition: left 0.3s ease-in-out;
  .menuMb {
    display: flex;
    justify-content: center;
    padding: 0.2rem 0;
    background: #0b318f;
  }
  .header_logo {
    display: flex;
    align-items: center;
    .logo {
      height: 0.6rem;
      background: #0b318f;
      img {
        height: 0.4rem;
        margin: 0.1rem 0.05rem;
      }
    }
    .collapse-btn-cont {
      float: left;
      padding: 0;
      cursor: pointer;
      align-items: center;
      display: flex;
      .header_btn {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        width: 0.26rem;
        height: 0.26rem;
        margin: 0 0.03rem 0 0.06rem;
        transition: all 0.4s ease;
        outline: none;
        color: #000;
        font-size: 22px;
      }
      .header_btn span {
        position: relative;
        display: block;
        width: 100%;
        height: 1px;
        margin: auto;
        background-color: #000;
        transition: all 0.4s ease;
      }
      .header_btn span::after {
        content: "";
        position: absolute;
        top: -7px;
        right: 0;
        width: 100%;
        height: 1px;
        background-color: #000;
        transition: all 0.4s ease;
      }
      .header_btn span::before {
        content: "";
        position: absolute;
        bottom: -7px;
        right: 0;
        width: 100%;
        height: 1px;
        background-color: #000;
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
        background-color: #555;
      }
    }
  }
  .header_left_logo {
    float: left;
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: 2.07rem;
    height: 100%;
    padding: 0 0.15rem;
    background-color: #0b318f;
    transition: width 0.3s;
    .logo {
      float: left;
      width: calc(90% - 0.32rem);
      img {
        display: block;
        width: 100%;
      }
    }

    .header_btn {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      width: 0.26rem;
      height: 0.26rem;
      margin: 0 0.03rem 0 0.06rem;
      transition: all 0.4s ease;
      outline: none;
      color: #000;
      font-size: 22px;
    }
    .header_btn span {
      position: relative;
      display: block;
      width: 100%;
      height: 1px;
      margin: auto;
      background-color: #000;
      transition: all 0.4s ease;
    }
    .header_btn span::after {
      content: "";
      position: absolute;
      top: -7px;
      right: 0;
      width: 50%;
      height: 1px;
      background-color: #000;
      transition: all 0.4s ease;
    }
    .header_btn span::before {
      content: "";
      position: absolute;
      bottom: -7px;
      right: 0;
      width: 75%;
      height: 1px;
      background-color: #000;
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
      background-color: #555;
    }
  }
  .header_left_hidd {
    width: 0.74rem;
    padding: 0;
    justify-content: center;
  }
  .collapse-btn-cont {
    float: left;
    padding: 0;
    cursor: pointer;
    align-items: center;
    display: flex;
  }
}
.header_arera {
  display: flex;
  align-items: center;
  justify-content: flex-end;
  height: 100%;
  padding: 0;
  margin: 0 0.6rem;
  @media screen and (max-width: 1260px) {
  }
}
.header_arera_hidd {
  width: calc(100% - 1rem);
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
  .createTask {
    display: flex;
    align-items: center;
    a {
      display: block;
      padding: 0.05rem 0.1rem;
      margin: 0 0.15rem;
      background-color: #4326ab;
      line-height: 1;
      border-radius: 0.08rem;
      text-align: center;
      color: #fff;
      font-size: 0.18rem;
      border: 0;
      outline: none;
      .el-icon-s-upload {
        display: none;
        width: 24px;
        height: 24px;
        position: relative;
        &::before {
          position: absolute;
          left: 0;
          top: 0;
          width: 100%;
          height: 100%;
          content: "";
          background-size: 17px !important;
          background: url(../assets/images/menuIcon/uploadFile.png) no-repeat
            center;
        }
      }
      @media screen and (max-width: 1024px) {
        span {
          display: none;
        }
        .el-icon-s-upload {
          display: block;
        }
      }
    }
  }
  img {
    height: 0.35rem;
  }
}
// .collapse-btn-cont {
//     float: left;
//     height: 100%;
//     padding: 0;
//     cursor: pointer;
// }

.collapse-btn-cont {
  float: left;
  padding: 0;
  cursor: pointer;
  align-items: center;
  display: flex;
  .header_btn {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    width: 22px;
    height: 22px;
    margin: 0 5px 0 15px;
    transition: all 0.4s ease;
    outline: none;
    color: #000;
    font-size: 22px;
    @media screen and (max-width: 479px) {
      margin: 0 0 0 10px;
    }
    &:hover {
      color: #4b5fff;
    }
  }
  .header_btn span {
    position: relative;
    display: block;
    width: 100%;
    height: 1px;
    margin: auto;
    background-color: #000;
    transition: all 0.4s ease;
  }
  .header_btn span::after {
    content: "";
    position: absolute;
    top: -7px;
    right: 0;
    width: 100%;
    height: 1px;
    background-color: #000;
    transition: all 0.4s ease;
  }
  .header_btn span::before {
    content: "";
    position: absolute;
    bottom: -7px;
    right: 0;
    width: 100%;
    height: 1px;
    background-color: #000;
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
    background-color: #555;
  }
}
.header-right {
  display: flex;
  justify-content: right;
  align-items: center;
  font-size: 0.1372rem;
  color: #959595;
  .progress {
    width: 200px;
    margin: 0 0.4rem 0 0;
    .el-progress /deep/ {
      .el-progress-bar {
        padding-right: 20px;
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
    }
  }
  .lang_style {
    min-width: 0.26rem;
    margin: 0 0.2rem;
    line-height: 0.26rem;
    font-size: 0.14rem;
    font-weight: 500;
    color: #000;
    border: 2px solid;
    border-radius: 6px;
    text-align: center;
    @media screen and (min-width: 1800px) {
      font-size: 16px;
    }
    @media screen and (max-width: 991px) {
      font-size: 13px;
      width: 25px;
      height: 25px;
      margin: 0 15px;
      line-height: 25px;
    }
    @media screen and (max-width: 479px) {
      margin: 0 10px;
    }
    span {
      cursor: pointer;
      color: inherit;
      &:hover {
        color: #4b5eff;
      }
    }
  }
  .network_mainnet {
    box-sizing: border-box;
    padding: 0.08rem 0.13rem 0.08rem 0;
    margin: 0 0.2rem 0 0;
    // background: linear-gradient(45deg, #4f8aff, #4b5eff);
    background: rgba(79, 138, 255, 0.05);
    line-height: 2;
    color: #333;
    font-size: 0.2rem;
    font-weight: 500;
    border-radius: 0.14rem;
    white-space: nowrap;
    cursor: pointer;
    @media screen and (max-width: 415px) {
      max-width: 100px;
      white-space: normal;
      line-height: 1.2;
    }
    &:hover {
      color: #000;
    }
    div {
      display: flex;
      align-items: center;
    }
    img,
    span {
      width: 20px;
      height: 20px;
      min-width: 20px;
      min-height: 20px;
      max-width: 100%;
      max-height: 100%;
      margin: auto 0.1rem;
      background-color: white;
      border-radius: 20px;
    }
  }
  .error {
    background: rgba(255, 104, 113, 0.6);
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
        padding: 0.08rem 0.1rem;
        margin: 0;
        line-height: 1.5;
        color: #fff;
        font-size: inherit;
        cursor: text;
        border-radius: 0.14rem;
        background: linear-gradient(45deg, #4f8aff, #4b5eff);
        h4 {
          padding: 0 0.17rem 0 0.1rem;
          margin-left: 0.1rem;
          background: #2d43e7;
          line-height: 2;
          border-radius: 0.14rem;
          cursor: pointer;
          @media screen and (max-width: 600px) {
            margin: 0;
          }
          &:hover {
            opacity: 0.9;
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
  .switch {
    position: relative;
    display: flex;
    margin: 0;
    .on {
      position: absolute;
      top: 0;
      left: 5px;
      z-index: 9;
      font-size: 12px;
      line-height: 2;
      color: #fff;
      svg {
        width: 18px;
        height: 18px;
        margin: 2px 0 0;
      }
    }
    .off {
      position: absolute;
      top: 0;
      right: 5px;
      z-index: 9;
      font-size: 12px;
      line-height: 2;
      color: #fff;
      svg {
        width: 14px;
        height: 14px;
        margin: 4px 0 0;
      }
    }
    .el-switch /deep/ {
      // margin-left: 0.1rem;
      height: 22px;
      .el-switch__core {
        height: 22px;
      }
      .el-switch__core:after {
        background-color: #fff;
        top: 0;
        width: 20px;
        height: 20px;
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
    .el-switch.is-checked /deep/ {
      .el-switch__core::after {
        margin-left: -22px;
      }
    }
    .swithUI {
      display: flex;
      -webkit-box-pack: center;
      justify-content: center;
      -webkit-box-align: center;
      align-items: center;
      cursor: pointer;
      font-size: 22px;
      svg {
        box-sizing: border-box;
        margin: 0px;
        min-width: 0px;
        font-size: 22px;
        cursor: pointer;
        width: 22px;
        height: 22px;
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
  }
  b {
    max-width: 30vw;
    display: inline-block;
    margin-right: 0.15rem;
    color: #000;
    overflow: hidden;
    white-space: nowrap;
    text-overflow: ellipsis;
    vertical-align: text-bottom;
  }
  img {
    display: block;
    width: 0.25rem;
    height: 0.25rem;
    margin: auto 0.13rem;
  }
  .el-menu--horizontal /deep/ {
    border: 0;
    > .el-submenu {
      .el-submenu__title {
        // height: auto;
        padding: 0 0.05rem 0 0.1rem;
        // background: url(../assets/images/up.png) no-repeat right center;
        // background-size: 0.1rem;
        line-height: 1;
        // color: #7db8ff;
        border: 0;
        height: 0.23rem;
        margin: 0 0 0 0.1rem;
        font-size: 0.12rem;
        border-radius: 0.05rem;
        line-height: 0.23rem;
        background-color: #0b318f;
        color: #ffffff;
        i {
          display: none;
        }
        span {
          color: #ffffff;
          margin: 0;
          font-size: 0.12rem;
          display: inline;
        }
      }
    }
  }
}
.sighChild {
  float: right;
  height: 0.23rem;
  padding: 0 0.1rem;
  margin: 0 0 0 0.1rem;
  font-size: 0.12rem;
  background-color: #c9f7f5;
  border-radius: 0.05rem;
  color: #13c1b8;
  line-height: 0.23rem;
  span {
    color: #13c1b8;
    cursor: pointer;
  }
  a {
    color: #13c1b8;
  }
}
.mobileShow {
  display: none;
}

.elImg {
  float: left;
  width: 0.25rem;
  margin: 0.17rem 0.05rem 0 0;
}
.el-popper {
  padding: 0 0.1rem;
  border: 0.01rem solid #eee;
  border-bottom: 0;
  box-shadow: 0 0 0.1rem rgba(187, 180, 180, 0.51);
}
.el-dropdown-menu__item {
  padding: 0 0.2rem;
  font-size: 0.14rem;
  border-bottom: 0.01rem solid #eee;
  line-height: 0.5rem;
}
.el-dropdown-menu__item.is-disabled {
  /* color: #ffb933; */
  pointer-events: none;
  /* background: #f9f9f9; */
  color: #606266;
  background: transparent;
}
.el-dropdown-menu__item:focus,
.el-dropdown-menu__item:not(.is-disabled):hover {
  background-color: transparent;
  color: #ffb933;
}
@media screen and (max-width: 991px) {
  .header_arera {
    margin: 0 0.1rem;
    // height: auto;
    padding: 0;
    border: 0;
  }
  .header_left {
    font-size: 0.2rem;
    color: #fff;
  }
  .sighChild {
    padding: 0 0.08rem;
    margin: 0 0 0 0.05rem;
  }
  .header-right {
    margin: 0.1rem 0;
    color: #fff;
    img {
      width: 0.2rem;
      height: 0.2rem;
      margin: auto 0.05rem 0 0;
    }
    b {
      margin-right: 0;
      color: #fff;
    }
    .pcShow {
      display: none;
    }
  }
  .mobileShow {
    display: block;
  }
  .language {
    margin: 0 0.1rem 0 0.15rem;
    color: #fff;
  }
}
@media screen and (max-width: 640px) {
  .language {
    margin: 0 0.05rem;
  }
}
@media screen and (max-width: 470px) {
  .header_left {
    font-size: 0.16rem;
    img {
      height: 30px;
    }
  }
}
</style>
