<template>
    <div class="header" :class="{'content-collapse': collapseLocal}">
        <div class="header_arera">
            <div class="header_left">
                <span v-if="!bodyWidth">{{headertitle}}</span>
                <span v-else>{{$t('navbar.sidebar_header')}}</span>
            </div>
            <div class="header-right">
                <div :class="{'online': addrChild, 'feh-metamask': 1==1}">
                    <el-tooltip class="item" effect="dark" content="Connect to your MetaMask Wallet" placement="bottom" v-if="!addrChild">
                        <img src="@/assets/images/metamask.png" @click="metamaskLogin" />
                    </el-tooltip>


                    <el-popover v-else
                        placement="bottom-end"
                        width="160"
                        trigger="click"
                        popper-class="addressInfo"
                        @show="walletInfo">
                        <h6>{{$t('fs3.connected_to')}}:</h6>
                        <h5 v-if="metaNetworkInfo.name">{{ metaNetworkInfo.name }}</h5>
                        <h4>{{addrChild | hiddAddress}}</h4>
                            <el-divider></el-divider>
                        <h4>{{addrChild | hiddAddress}}</h4>
                        <h5>{{priceAccound}} {{ metaNetworkInfo.unit}}</h5>
                            <el-divider></el-divider>
                        <h3 @click="signOutFun">{{$t('fs3.Disconnect')}}</h3>

                        <img src="@/assets/images/metamask.png" slot="reference" />
                    </el-popover>

                    <!-- <span class="text" v-if="addrChild && !network.text" @click="wrongVisible=true">Wrong Network</span> -->
                    <span class="text" v-if="!addrChild" @click="metamaskLogin">{{$t('fs3.Connect_Wallet')}}</span>
                    <span class="text textTrue" v-else>{{metaNetworkInfo.name}}</span>
                </div>
                <!-- {{$t('navbar.hi')}}
                <router-link :to="{name: 'upload_file'}">
                    <b>&nbsp;{{email}}</b>
                </router-link> -->

                <!-- 国际化 -->
                <el-dropdown
                    trigger="click"
                    class="language leftBoth"
                    @command="handleSetLanguage"
                    v-if="langShow && !bodyWidth"
                >
                    <div>
                    <span v-if="languageMcp === 'cn'" style="cursor: pointer;">CN</span>
                    <span v-if="languageMcp === 'en'" style="cursor: pointer;">EN</span>
                    </div>
                    <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item
                        command="cn"
                        :disabled="languageMcp === 'cn'"
                        v-if="languageMcp === 'cn'"
                    >
                        <img src="../assets/images/cn.jpg" class="elImg" />简体中文
                    </el-dropdown-item>
                    <el-dropdown-item
                        command="cn"
                        :disabled="languageMcp === 'cn'"
                        v-if="languageMcp === 'en'"
                    >
                        <img src="../assets/images/cn.jpg" class="elImg" />Chinese
                    </el-dropdown-item>
                    <el-dropdown-item command="en" :disabled="languageMcp === 'en'">
                        <img src="../assets/images/en.jpg" class="elImg" />English
                    </el-dropdown-item>
                    </el-dropdown-menu>
                </el-dropdown>
                <!-- <div class="sighChild" v-if="!email">
                    <span @click="pageJump(1)">
                        {{$t('navbar.log_in')}}
                    </span>
                    /
                    <span @click="pageJump(2)">
                        {{$t('navbar.sign_up')}}
                    </span>
                </div> -->
                <!-- <a href="javascript:;" @click="logout" class="sighChild" v-if="email && !bodyWidth">{{$t('navbar.logout')}}</a> -->
                <!-- mobile显示 -->
                <div class="mobileShow">
                    <div class="collapse-btn-cont" @click="collapseChage">
                        <div class="header_btn">
                            <span></span>
                        </div>
                    </div>
                </div>
            </div>
        </div>    
        

        <el-dialog
        title="Tips"
        :visible.sync="wrongVisible" :show-close="false"
        :width="width" custom-class="wrongNet">
        <span>Please connect to the appropriate Ethereum network.</span>
        </el-dialog>

        <div class="loadIndexStyle" v-show="loadIndexing" v-loading="loadIndexing"></div>
    </div>
</template>
<script>
// import bus from './bus';
import * as myAjax from "@/api/login";
import axios from 'axios'
import NCWeb3 from "@/utils/web3";
export default {
    data() {
        return {
            collapseLocal: this.$store.getters.collapseL == 'true'||this.$store.getters.collapseL==true?true: false,
            collapse: document.body.clientWidth<1024?true:false,
            fullscreen: false,
            name: 'linxin',
            message: 2,
            langNew: '',
            userShow: false,
            loginShow: localStorage.getItem("mcpLoginAccessToken") ? false : true,
            bodyWidth: document.body.clientWidth<1024?true:false,
            langShow: true,
            loadIndexing: false,
            // 控制是否在路由栈中清理当前页面的数据
            replaceData:null,
            tools: '1',
            tabOaxLogin: localStorage.getItem("mcpLoginAccessToken"),
            tabOaxNew: localStorage.getItem("mcpLoginAccessToken"),
            priceAccound: 0,
            network: {
                name: '',
                unit: '',
                center_fail: false
            },
            addrChild: '',
            wrongVisible: false,
            width: document.body.clientWidth>600?'400px':'95%',
        };
    },
    props: ["meta"],
    computed: {
        email() {
            return this.$store.state.user.email
        },
        languageMcp() {
            return this.$store.getters.languageMcp
        },
        routerMenu() {
            return this.$store.getters.routerMenu
        },
        headertitle() {
            return this.$store.getters.headertitle
        },
        avater() {
            return this.$store.getters.avater
        },
        collapseL() {
            return this.$store.getters.collapseL
        },
        metaAddress() {
            return this.$store.getters.metaAddress
        },
        metaNetworkInfo() {
            return this.$store.getters.metaNetworkInfo?JSON.parse(JSON.stringify(this.$store.getters.metaNetworkInfo)):{}
        }
    },
    watch: {
        $route: function (to, from) {
            if((to.name === 'register' && from.name === 'login') || (to.name === 'login' && from.name === 'register')){
                this.replaceData = 'replace'
            }else{
                this.replaceData = ''
            }
            if(this.bodyWidth){
                this.collapseLocal = false
                this.collapseChage();
            }
        },
        // 'tabOaxLogin': function() {
        //     let _this = this;
        //     let name = _this.$route.name;
        //     if(!localStorage.getItem("mcpLoginAccessToken")){
        //         if(name == 'my_files' || name == 'my_files_detail' || name == 'upload_file' || name == 'Search_file' ||
        //            name == 'billing' || name == 'settings'){
        //             _this.$store.dispatch("FedLogOut").then(() => {
        //                 _this.$router.push("/login");
        //             });
        //         }
        //     }
        //     if(!_this.email && localStorage.getItem("mcpLoginAccessToken")){
        //         location. reload()
        //     }
        // },
        'collapseL': function(){
            this.collapseLocal = this.$store.getters.collapseL == 'true'||this.$store.getters.collapseL==true?true: false
        },
        metaAddress: function() {
            this.addrChild = this.metaAddress
            this.commonParam()
            this.walletInfo()
        },
        meta: function() {
            this.walletInfo()
        }
    },
    methods: {
        handleSelect(key, keyPath) {
            // console.log(key, keyPath);
        },
        pageJump(data){
            let _this = this
            let name = _this.$router.history.current.name
            if(name === 'login' || name === 'register'){
                if(data === 1){
                    _this.$router.replace('/login')
                }else if(data === 2){
                    _this.$router.replace('/register')
                }
            }else{
                if(data === 1){
                    _this.$router.push('/login')
                }else if(data === 2){
                    _this.$router.push('/register')
                }
            }
            
        },
        // 用户名下拉菜单选择事件
        logout() {
            var _this = this;

            let params = {};
            axios.post(_this.data_api+'auth/logout', params,{
                headers: {
                  'Authorization': "Bearer "+_this.$store.getters.accessToken
                },
            }).then((response) => {
                //   console.log('logout', response.data)
                  if(response.data.status == 'success'){
                    _this.$store.dispatch("FedLogOut").then(() => {
                        _this.$router.push("/supplierAllBack");
                        _this.loginShow = localStorage.getItem("mcpLoginAccessToken") ? false : true
                    });
                  }else{
                    console.log(response.data.message);
                    _this.$message.error(response.data.message)
                  }
            }).catch(function (error) {
                console.log(error.config);
                _this.$store.dispatch("FedLogOut").then(() => {
                    _this.$router.push("/login");
                    _this.loginShow = localStorage.getItem("mcpLoginAccessToken") ? false : true
                });
            });

            // myAjax.logout(params).then(res => {
            //   console.log('_RequestUploads_', res)
            //   if (res.status == "success") {
            //     _this.$store.dispatch("FedLogOut").then(() => {
            //         _this.$router.push("/supplierAllBack");
            //         _this.loginShow = localStorage.getItem("mcpLoginAccessToken") ? false : true
            //     });
            //   } else {
            //       _this.$message.error(res.message)
            //   }
            // }).catch(error => {
            //     console.log(error)
            // })

        },
        // 侧边栏折叠
        collapseChage() {
            this.collapseLocal = !this.collapseLocal;
            this.$store.dispatch('setCollapse', this.collapseLocal)
            // bus.$emit('collapse', this.collapse);
        },
        handleSetLanguage(lang){
            let _this = this
            _this.loadIndexing=true;

            document.body.style.height = '100vh'
            document.body.style['overflow-y'] = 'hidden'
            _this.$i18n.locale = lang;
            _this.$store.dispatch("setLanguage", lang);

            window.location.reload();
        },
        getHiddenProp() {
            var prefixes = ['webkit','moz','ms','o'];
            
            // if 'hidden' is natively supported just return it
            if ('hidden' in document) return 'hidden';
            
            // otherwise loop over all the known prefixes until we find one
            for (var i = 0; i < prefixes.length; i++)
            {
                if ((prefixes[i] + 'Hidden') in document) 
                    return prefixes[i] + 'Hidden';
            }
        
            // otherwise it's not supported
            return null;
        },
        getVisibilityState() {
            var prefixes = ['webkit', 'moz', 'ms', 'o'];
            if ('visibilityState' in document) return 'visibilityState';
            for (var i = 0; i < prefixes.length; i++) 
            {
                if ((prefixes[i] + 'VisibilityState') in document)
                    return prefixes[i] + 'VisibilityState';
            }
            // otherwise it's not supported
            return null;
        },
        metamaskLogin() {
            this.$router.push({ path: '/metamask_login' })
        },
        // Wallet address
        signFun(redirect){
            let _this = this
            if(!_this.addrChild){
                NCWeb3.Init(addr=>{
                    //Get the corresponding wallet address
                    // console.log('Wallet address:', addr)
                    _this.$nextTick(() => {
                        _this.addrChild = addr
                        _this.walletInfo()
                    })
                })
                return false
            }
        },
        walletInfo() {
            let _this = this
            if(!_this.addrChild || _this.addrChild == 'undefined'){
                return false
            }

            ethereum
            .request(
                {
                    "jsonrpc":"2.0",
                    "method":"eth_getBalance",
                    "params":[_this.addrChild, "latest"],
                    "id":19999
                }
            )
            .then((balance) => {
                let balanceAll = web3.utils.fromWei(balance, 'ether')
                _this.priceAccound = Number(balanceAll).toFixed(4)
            })
            .catch((error) => {
                console.error(`Error fetching getBalance: ${error.code}: ${error.message}`);
            });
            
            ethereum
            .request({ method: 'eth_chainId' })
            .then((chainId) => {
                let netId = parseInt(chainId, 16)
                // console.log('network ID:', netId)
                // console.log(`decimal number: ${parseInt(chainId, 16)}`);
                _this.$store.dispatch('setMetaNetworkId', netId)
                switch (netId) {
                case 1:
                    _this.network.name = 'mainnet';
                    _this.network.unit = 'ETH';
                    _this.network.center_fail = true
                    _this.$store.dispatch('setMetaNetworkInfo', _this.network)
                    return;
                case 3:
                    _this.network.name = 'ropsten';
                    _this.network.unit = 'ETH';
                    _this.network.center_fail = true
                    _this.$store.dispatch('setMetaNetworkInfo', _this.network)
                    break;
                case 4:
                    _this.network.name = 'rinkeby';
                    _this.network.unit = 'ETH';
                    _this.network.center_fail = true
                    _this.$store.dispatch('setMetaNetworkInfo', _this.network)
                    return;
                case 5:
                    _this.network.name = 'goerli';
                    _this.network.unit = 'ETH';
                    _this.network.center_fail = true
                    _this.$store.dispatch('setMetaNetworkInfo', _this.network)
                    return;
                case 42:
                    _this.network.name = 'kovan';
                    _this.network.unit = 'ETH';
                    _this.network.center_fail = true
                    _this.$store.dispatch('setMetaNetworkInfo', _this.network)
                    return;
                case 56:
                    _this.network.name = 'BSC';
                    _this.network.unit = 'BNB';
                    _this.network.center_fail = true
                    _this.$store.dispatch('setMetaNetworkInfo', _this.network)
                    return;
                case 97:
                    _this.network.name = 'BSC';
                    _this.network.unit = 'BNB';
                    _this.network.center_fail = true
                    _this.$store.dispatch('setMetaNetworkInfo', _this.network)
                    return;
                case 999:
                    _this.network.name = 'NBAI';
                    _this.network.unit = 'NBAI';
                    _this.network.center_fail = true
                    _this.$store.dispatch('setMetaNetworkInfo', _this.network)
                    return;
                case 80001:
                    _this.network.name = 'mumbai';
                    _this.network.unit = 'MATIC';
                    _this.network.center_fail = false
                    _this.$store.dispatch('setMetaNetworkInfo', _this.network)
                    if(_this.meta) {
                        if(_this.$route.query.redirect && _this.$route.query.redirect != '/supplierAllBack'){
                            // 防止登录后需要跳转到指定页面
                            _this.$router.push({ path: _this.$route.query.redirect })
                        }else{
                            _this.$router.push({ path: '/upload_file' })
                        }
                        _this.$emit("getMetamaskLogin", false)
                    }
                    return;
                default:
                    _this.network.name = 'Custom';
                    _this.network.unit = '';
                    _this.network.center_fail = true
                    _this.$store.dispatch('setMetaNetworkInfo', _this.network)
                    return;
                }
            })
            .catch((error) => {
                console.error(`Error fetching chainId: ${error.code}: ${error.message}`);
            });
        },
        fn() {
            let _this = this
            ethereum.on("accountsChanged", function(accounts) {
                if(_this.metaAddress){
                    web3.eth.getAccounts().then(accounts => {
                        _this.addrChild = accounts[0]
                        _this.walletInfo()
                        _this.$store.dispatch('setMetaAddress', accounts[0])
                        _this.$router.go(0)
                    })
                    // console.log('account header:', accounts[0]);  //Once the account is switched, it will be executed here
                }
            });
            // networkChanged
            ethereum.on("chainChanged", function(accounts) {
                _this.walletInfo()
            });
            // 监听metamask网络断开
            ethereum.on('disconnect', (code, reason) => {
                // console.log(`Ethereum Provider connection closed: ${reason}. Code: ${code}`);
            });
        },
        signOutFun() {
            this.addrChild = ''
            this.$store.dispatch('setMetaAddress', '')
            this.$store.dispatch('setMetaNetworkId', 0)
            this.network.name = '';
            this.network.unit = '';
            this.network.center_fail = false
            this.$store.dispatch('setMetaNetworkInfo', JSON.stringify(this.network))
            this.$router.push("/supplierAllBack");
        },
        commonParam(){
            let _this = this
            let common_api = `${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/common/system/params?limit=20&wallet_address=${_this.metaAddress}`

            axios.get(common_api, {
                headers: {
                    // 'Authorization': "Bearer "+ _this.$store.getters.accessToken
                },
            })
            .then((json) => {
                if(json.data.status == 'success'){
                    _this.$root.LOCK_TIME = json.data.data.LOCK_TIME
                    _this.$root.PAY_GAS_LIMIT = json.data.data.PAY_GAS_LIMIT
                    _this.$root.PAY_WITH_MULTIPLY_FACTOR = json.data.data.PAY_WITH_MULTIPLY_FACTOR
                    _this.$root.RECIPIENT = json.data.data.RECIPIENT
                    _this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS = json.data.data.SWAN_PAYMENT_CONTRACT_ADDRESS
                    _this.$root.USDC_ADDRESS = json.data.data.USDC_ADDRESS
                }
            }).catch(error => {
                console.log(error)
            })
        }
    },
    mounted() {
        let _this = this
        if(_this.bodyWidth){
            _this.collapseLocal = false
            _this.collapseChage();
        }

        var visProp = _this.getHiddenProp();
        if (visProp) 
        {
            var evtname = visProp.replace(/[H|h]idden/, '') + 'visibilitychange';
            document.addEventListener(evtname, function () {
                _this.tabOaxLogin = document[_this.getVisibilityState()]
            }, false);
        }

        if(_this.metaAddress){
            _this.commonParam()
            _this.signFun()
        }
        _this.fn()
    },
    filters: {
        number (value) {
            let realVal = ''
            if (!isNaN(value) && value!== '' && value!==null) {
                // realVal = 0.000000000000000001 * value
                realVal = String(value).replace(/^(.*\..{18}).*$/, "$1")
            } else {
                realVal = '--'
            }
            return realVal
        },
        hiddAddress: function (val) {
            if(val){
                return `${val.substring(0, 6)}...${val.substring(val.length - 4)}`
            }else{
                return '-'
            }
        }
    }
};
</script>
<style  lang="scss" scoped>
.el-dialog__wrapper /deep/{
    display: flex;
    align-items: center;
    .wrongNet{
        margin: auto !important;
        border-radius: 4px;
        background: #fff url(../assets/images/tip_bg.png) no-repeat;
        background-size: 1.45rem;
        background-position: -0.2rem -0.3rem;
        .el-dialog__header{
            display: flex;
            color: #000;
            font-size: 18px;
            padding: 0.15rem 0.15rem 0.1rem;
            .el-dialog__headerbtn{
                i{
                    color: #000;
                    font-weight: 600;
                }
            }
        }
        .el-dialog__body {
            padding: 0.2rem 0.2rem 0.3rem;
            span{
              word-break: break-word;
              line-height: 1.5;    
              color: #606266;
              font-size: 14px;
            }
        }
    }

}
.addressInfo{
  padding: 0.2rem;
  h6{
    margin: 0.14rem 0 0;
    font-size: 0.13rem;
    font-weight: normal;
    padding: 0 0.07rem;
  }
  h5{
    font-size: 0.13rem;
    font-weight: normal;
    padding: 0 0.07rem;
  }
  h4{
    font-size: 0.14rem;
    font-weight: normal;
    padding: 0 0.07rem;
  }
  h3{
    margin: 0 0 0.05rem;
    font-size: 0.14rem;
    font-weight: normal;
    padding: 0 0.07rem;
    cursor: pointer;
    &:hover{
      color: rgba(11, 49, 143, 1);
    }
  }
  .el-divider /deep/{
    margin: 0.14rem 0;
  }
}
.loadIndexStyle{
  position: fixed;
  top:0;
  left:0;
  right:0;
  bottom:0;
  z-index: 2000;
  background: rgba(255, 255, 255, 1);
}
.header {
    position: absolute;
    top: 0;
    right: 0;
    left: 2.37rem;
    box-sizing: border-box;
    height: 0.86rem;
    font-size: 0.22rem;
    color: #fff;
    background-color: #fff;
    -webkit-transition: left .3s ease-in-out;
    transition: left .3s ease-in-out;
    .menuMb{
        display: flex;
        justify-content: center;
        padding: 0.2rem 0;
        background: #0b318f;

    }
    .header_logo{
        display: flex;
        align-items: center;
        .logo{
            height: 0.6rem;
            background: #0b318f;
            img{
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
            .header_btn{
                display: flex;
                align-items: center;
                flex-wrap: wrap;
                width: 0.26rem;
                height: 0.26rem;
                margin: 0 0 0 0.06rem;
                transition: all 0.4s ease;
                outline: none;
            }
            .header_btn span{
                position: relative;
                display: block;
                width: 100%;
                height: 1px;
                margin: auto;
                background-color:#fff;
                transition: all 0.4s ease;
            }
            .header_btn span::after{
                content: '';
                position: absolute;
                top: -7px;
                right: 0;
                width: 100%;
                height: 1px;
                background-color: #fff;
                transition: all 0.4s ease;
            }
            .header_btn span::before{
                content: '';
                position: absolute;
                bottom: -7px;
                right: 0;
                width: 100%;
                height: 1px;
                background-color: #fff;
                transition: all 0.4s ease;
            }
            .header_btn.header_btn_left span::before,.header_btn.header_btn_left span::after{
                right: auto;
                left: 0;
            }
            .header_btn:hover span,.header_btn:hover span::before,.header_btn:hover span::after{
                width: 100%;
                background-color: #fff;
            }
        }
        
    }
    .header_left_logo{
        float: left;
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 2.07rem;
        height: 100%;
        padding: 0 0.15rem;
        background-color: #0b318f;
        transition: width .3s;
        .logo {
            float: left;
            width: calc(90% - 0.32rem);
            img{
                display: block;
                width: 100%;
            }
        }


        .header_btn{
            display: flex;
            align-items: center;
            flex-wrap: wrap;
            width: 0.26rem;
            height: 0.26rem;
            margin: 0 0 0 0.06rem;
            transition: all 0.4s ease;
            outline: none;
        }
        .header_btn span{
            position: relative;
            display: block;
            width: 100%;
            height: 1px;
            margin: auto;
            background-color:#fff;
            transition: all 0.4s ease;
        }
        .header_btn span::after{
            content: '';
            position: absolute;
            top: -7px;
            right: 0;
            width: 50%;
            height: 1px;
            background-color: #fff;
            transition: all 0.4s ease;
        }
        .header_btn span::before{
            content: '';
            position: absolute;
            bottom: -7px;
            right: 0;
            width: 75%;
            height: 1px;
            background-color: #fff;
            transition: all 0.4s ease;
        }
        .header_btn.header_btn_left span::before,.header_btn.header_btn_left span::after{
            right: auto;
            left: 0;
        }
        .header_btn:hover span,.header_btn:hover span::before,.header_btn:hover span::after{
            width: 100%;
            background-color: #fff;
        }
    }
    .header_left_hidd{
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
.header_arera{
    display: flex;
    align-items: center;
    justify-content: space-between;
    // width: calc(100% - 3.87rem);
    height: 0.42rem;
    padding: 0 0.2rem;
    margin: 0.35rem 0.55rem 0 0.48rem;
    border-bottom: 1px solid #e6e6e6;
}
.header_arera_hidd{
    width: calc(100% - 1rem);
}
.header_left{
    float: left;
    width: auto;
    margin: 0;
    font-size: 0.2772rem;
    font-weight: 500;
    color: #0b318f;
    line-height: 0.42rem;
    text-transform: uppercase;
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
            .header_btn{
                display: flex;
                align-items: center;
                flex-wrap: wrap;
                width: 0.22rem;
                height: 0.22rem;
                margin: 0 0 0 0.06rem;
                transition: all 0.4s ease;
                outline: none;
            }
            .header_btn span{
                position: relative;
                display: block;
                width: 100%;
                height: 1px;
                margin: auto;
                background-color:#fff;
                transition: all 0.4s ease;
            }
            .header_btn span::after{
                content: '';
                position: absolute;
                top: -7px;
                right: 0;
                width: 100%;
                height: 1px;
                background-color: #fff;
                transition: all 0.4s ease;
            }
            .header_btn span::before{
                content: '';
                position: absolute;
                bottom: -7px;
                right: 0;
                width: 100%;
                height: 1px;
                background-color: #fff;
                transition: all 0.4s ease;
            }
            .header_btn.header_btn_left span::before,.header_btn.header_btn_left span::after{
                right: auto;
                left: 0;
            }
            .header_btn:hover span,.header_btn:hover span::before,.header_btn:hover span::after{
                width: 100%;
                background-color: #fff;
            }
        }
.header-right{
    display: flex;
    justify-content: right;
    align-items: center;
    font-size: 0.1372rem;
    color: #959595;
    .feh-metamask{
        display: flex;
        align-items: center;
        position: relative;
        width: auto;
        height: 30px;
        margin: 0 10px 0 0;
        cursor: pointer;
        img{
            display: block;
            width: auto;
            height: 30px;
            cursor: pointer;
            margin: 0 5px 0 0;
        }
        &:before{
            position: absolute;
            left: 30px;
            top: -4px;
            content: "";
            width: 5px;
            height: 5px;
            border-radius: 100%;
            background: #d7d6d6;
        }
        .text{
            padding: 0.05rem 0.1rem;
            font-size: 0.12rem;
            background: #f56c6c;
            color: #fff;
            border-radius: 0.51rem;
            @media screen and (max-width:441px) {
                display: none;
            }
        }
        .textTrue{
            background: #4326ab;
        }
    }
    .online{
        &:before{
            background: #0fce7c;
        }
    }
    b{
        max-width: 30vw;
        display: inline-block;
        margin-right: 0.15rem;
        color: #000;
        overflow: hidden;
        white-space: nowrap;
        text-overflow: ellipsis;
        vertical-align: text-bottom;
    }
    img{
        display: block;
        width: 0.25rem;
        height: 0.25rem;
        margin: auto 0.13rem;
    }
    .el-menu--horizontal /deep/{
        border: 0;
        >.el-submenu {
            .el-submenu__title{
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
                i{
                    display: none;
                }
                span{
                    color: #ffffff;
                    margin: 0;
                    font-size: 0.12rem;
                    display: inline;
                }
            }
        }
    }
}
.sighChild{
    float: right;
    height: 0.23rem;
    padding: 0 0.1rem;
    margin: 0 0 0 0.1rem;
    font-size: 0.12rem;
    background-color: #c9f7f5;
    border-radius: 0.05rem;
    color: #13c1b8;
    line-height: 0.23rem;
    span{
        color: #13c1b8;
        cursor:pointer
    }
    a{
        color: #13c1b8;
    }
}
.mobileShow{
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
@media screen and (max-width:999px){
    .header{
        left: 0 !important;
        height: 0.6rem;
        background: #0b318f;
    }
    .header_arera{
        margin: 0.1rem;
        height: auto;
        padding: 0;
        border: 0;
    }
    .header_left{
        font-size: 0.2rem;
        color: #fff;
    }
    .sighChild{
        padding: 0 0.08rem;
        margin: 0 0 0 0.05rem;
    }
    .header-right{
        margin: 0.1rem 0;
        color: #fff;
        img{
            width: 0.2rem;
            height: 0.2rem;
            margin: auto 0.05rem 0 0;
        }
        b{
            margin-right: 0;
            color: #fff;
        }
    }
    .pcShow{
        display: none;
    }
    .mobileShow{
        display: block;
    }
    .language{
        margin: 0 0.1rem 0 0.15rem;
        color: #fff;
    }
}
@media screen and (max-width:640px){
    .language{
        margin: 0 0.05rem;
    }
}
@media screen and (max-width:470px){
    .header_arera{
        margin: 0.1rem 0.05rem;
    }
    .header_left{
        font-size: 0.16rem;
    }
}
</style>
