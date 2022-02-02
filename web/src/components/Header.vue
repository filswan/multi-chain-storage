<template>
    <div class="header" :class="{'content-collapse': collapseLocal}">
        <div class="header_arera">
            <div class="header_left">
                <div v-if="!bodyWidth" class="createTask">
                    {{headertitle}}
                    <router-link :to="{name: 'upload_file'}" v-if="$route.path === '/my_files'">
                        <span>{{$t('uploadFile.Upload_More_Files')}}</span>
                        <i class="el-icon-s-upload"></i>
                    </router-link>
                </div>
                <img v-else src="@/assets/images/MCP_logo.png" alt="">
                <!-- <span v-else>{{$t('navbar.sidebar_header')}}</span> -->
            </div>
            <div class="header-right">
                <div :class="{'online': addrChild, 'feh-metamask': 1==1}">
                    <div v-if="!addrChild" class="logged_in">
                        <el-tooltip class="item" effect="dark" :content="$t('fs3Login.toptip_03')" placement="bottom">
                            <img src="@/assets/images/metamask.png" @click="metamaskLogin" />
                        </el-tooltip>
                        <span class="text" @click="metamaskLogin">{{$t('fs3.Connect_Wallet')}}</span>
                    </div>
                    
                    <div v-else class="logged_in">
                        <span class="text textTrue">{{metaNetworkInfo.name}}</span>
                        <div class="info">
                            <h5>{{priceAccound}} {{ metaNetworkInfo.unit}}</h5>
                            <h4 @click="wrongVisible=true">{{addrChild | hiddAddress}} <i class="el-icon-user-solid"></i></h4>
                        </div>
                        <el-button v-if="!bodyWidth" class="text textTrue" @click="signOutFun">{{$t('fs3.Disconnect')}}</el-button>
                    </div>
                </div>

                <!-- 国际化 -->
                <el-dropdown
                    trigger="click"
                    class="language leftBoth"
                    @command="handleSetLanguage" v-if="!bodyWidth"
                >
                    <div class="background">
                    <span v-if="languageMcp === 'cn'" style="cursor: pointer;">CN <i class="el-icon-arrow-down"></i></span>
                    <span v-if="languageMcp === 'en'" style="cursor: pointer;">EN <i class="el-icon-arrow-down"></i></span>
                    </div>
                    <el-dropdown-menu slot="dropdown">
                    <el-dropdown-item command="cn" :disabled="languageMcp === 'cn'">
                        <img src="../assets/images/cn.jpg" class="elImg" />简体中文
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
        :title="$t('fs3Login.Account')"
        :visible.sync="wrongVisible"
        :width="width" custom-class="wrongNet">
            <label>{{$t('fs3Login.Connected_MetaMask')}}</label>
            <div class="address">{{addrChild | hiddAddress}}</div>
            <div class="share">
                <el-button @click="shareTo">
                    <svg t="1640937862402" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3723" width="32" height="32"><path d="M852.77 889.05H171.23A36.27 36.27 0 0 1 135 852.78V171.22A36.27 36.27 0 0 1 171.23 135H375a36.28 36.28 0 0 1 0 72.55H207.5v609h609V649a36.28 36.28 0 1 1 72.55 0v203.78a36.27 36.27 0 0 1-36.28 36.27z" fill="#0b318f" p-id="3724"></path><path d="M407.15 653.13a36.28 36.28 0 0 1-25.66-61.93L747.66 225A36.28 36.28 0 1 1 799 276.34L432.8 642.51a36.17 36.17 0 0 1-25.65 10.62z" fill="#0b318f" p-id="3725"></path><path d="M852.78 414.94V171.23H609.06l140.42 103.29 103.3 140.42z" fill="#0b318f" p-id="3726"></path><path d="M852.78 451.22a36.29 36.29 0 0 1-29.23-14.78l-100-136-136-100a36.28 36.28 0 0 1 21.5-65.5h243.72a36.28 36.28 0 0 1 36.28 36.28V415a36.28 36.28 0 0 1-36.27 36.27zM719.6 207.5l51.4 37.8a36.23 36.23 0 0 1 7.7 7.7l37.8 51.38V207.5z" fill="#0b318f" p-id="3727"></path></svg>
                    {{$t('fs3Login.View_explorer')}}
                </el-button>

                <el-button v-if="copyClick" @click="copyTextToClipboard(addrChild)">
                    <svg t="1640938541398" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4760" width="32" height="32"><path d="M746.932 698.108" p-id="4761" fill="#0b318f"></path><path d="M925.731 288.698c-1.261-1.18-3.607-3.272-6.902-6.343-5.486-5.112-11.615-10.758-18.236-16.891-18.921-17.526-38.003-35.028-56.046-51.397-2.038-1.848-2.038-1.835-4.077-3.682-24.075-21.795-44.156-39.556-58.996-52.076-8.682-7.325-15.517-12.807-20.539-16.426-3.333-2.402-6.043-4.13-8.715-5.396-3.365-1.595-6.48-2.566-10.905-2.483C729.478 134.227 720 143.77 720 155.734l0 42.475 0 42.475 0 84.95L720 347l21.205 0L890 347l0 595L358.689 942C323.429 942 295 913.132 295 877.922L295 177l361.205 0c11.736 0 21.25-9.771 21.25-21.5s-9.514-21.5-21.25-21.5l-382.5 0L252 134l0 21.734L252 813l-52.421 0C166.646 813 140 786.928 140 754.678L140 72l566.286 0C739.29 72 766 98.154 766 130.404L766 134l40 0 0-3.596C806 76.596 761.271 33 706.286 33L119.958 33 100 33l0 19.506 0 702.172C100 808.463 144.642 852 199.579 852L252 852l0 25.922C252 936.612 299.979 984 358.689 984l552.515 0L932 984l0-21.237L932 325.635 932 304l0.433 0C932.432 299 930.196 292.878 925.731 288.698zM762 304l0-63.315L762 198.21l0-0.273c14 11.479 30.3 26.369 49.711 43.942 2.022 1.832 2.136 1.832 4.157 3.665 17.923 16.259 36.957 33.492 55.779 50.926 2.878 2.666 5.713 5.531 8.391 7.531L762 304.001z" p-id="4762" fill="#0b318f"></path><path d="M816.936 436 407.295 436c-10.996 0-19.91 8.727-19.91 19.5 0 10.77 8.914 19.5 19.91 19.5l409.641 0c11 0 19.914-8.73 19.914-19.5C836.85 444.727 827.936 436 816.936 436z" p-id="4763" fill="#0b318f"></path><path d="M816.936 553 407.295 553c-10.996 0-19.91 8.727-19.91 19.5 0 10.774 8.914 19.5 19.91 19.5l409.641 0c11 0 19.914-8.726 19.914-19.5C836.85 561.727 827.936 553 816.936 553z" p-id="4764" fill="#0b318f"></path><path d="M816.936 689 407.295 689c-10.996 0-19.91 8.729-19.91 19.503 0 10.769 8.914 19.497 19.91 19.497l409.641 0c11 0 19.914-8.729 19.914-19.497C836.85 697.729 827.936 689 816.936 689z" p-id="4765" fill="#0b318f"></path></svg>
                    {{$t('fs3Login.Copy_Address')}}
                </el-button>

                <el-button v-else>
                    <svg t="1640939105223" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5619" width="32" height="32"><path d="M512 998.4c-268.288 0-486.4-218.112-486.4-486.4S243.712 25.6 512 25.6s486.4 218.112 486.4 486.4-218.112 486.4-486.4 486.4z m0-896c-225.792 0-409.6 183.808-409.6 409.6s183.808 409.6 409.6 409.6 409.6-183.808 409.6-409.6-183.808-409.6-409.6-409.6z" p-id="5620" fill="#0b318f"></path><path d="M457.728 680.448c-9.728 0-19.456-3.584-27.136-11.264l-150.528-150.016c-14.848-14.848-14.848-39.424 0-54.272 14.848-14.848 39.424-14.848 54.272 0l123.392 122.88 231.936-230.912c14.848-14.848 39.424-14.848 54.272 0 14.848 14.848 14.848 39.424 0 54.272l-259.072 258.048c-7.68 7.168-17.408 11.264-27.136 11.264z" p-id="5621" fill="#0b318f"></path></svg>
                    {{$t('fs3Login.Copied')}}
                </el-button>

            </div>
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
            bodyWidth: document.body.clientWidth<999?true:false,
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
            width: document.body.clientWidth>600?'450px':'95%',
            copyClick: true
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
        shareTo(){
            window.open('https://mumbai.polygonscan.com/address/'+this.addrChild)
        },
        copyTextToClipboard(text) {
            let _this = this
            var txtArea = document.createElement("textarea");
            txtArea.id = 'txt';
            txtArea.style.position = 'fixed';
            txtArea.style.top = '0';
            txtArea.style.left = '0';
            txtArea.style.opacity = '0';
            txtArea.value = text;
            document.body.appendChild(txtArea);
            txtArea.select();

            try {
                var successful = document.execCommand('copy');
                var msg = successful ? 'successful' : 'unsuccessful';
                console.log('Copying text command was ' + msg);
                if (successful) {
                    _this.copyClick = false
                    setTimeout(function(){_this.copyClick = true}, 600)
                    return true;
                }
            } catch (err) {
                console.log('Oops, unable to copy');
            } finally {
                document.body.removeChild(txtArea);
            }
            return false;
        },
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
                    _this.$root.MINT_CONTRACT = json.data.data.MINT_CONTRACT
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
        .el-dialog__header{
            display: flex;
            align-items: center;
            justify-content: space-between;
            color: #000;
            font-size: 20px;
            padding: 0.15rem 0.15rem 0.1rem;
            .el-dialog__headerbtn{
                position: relative;
                top: auto;
                right: auto;
                font-size: inherit;
                i{
                    font-size: inherit;
                    &:hover{
                        color: #0b318f;
                    }
                }
            }
            .el-dialog__title{
                font-size: inherit;
            }
        }
        .el-dialog__body {
            padding: 0.1rem 0.15rem 0.3rem;
            font-size: 14px;
            label{
                word-break: break-word;
                line-height: 1;    
                color: #666;
                font-size: inherit;
            }
            .address{
                background: rgba(233, 233, 233, 1);
                padding: 8px;
                margin: 10px 0 12px;
                border-radius: 8px;
                font-size: inherit;
            }
            .share{
                display: flex;
                align-items: center;
                font-size: inherit;
                .el-button{
                    padding: 0 20px 0 0;
                    background: transparent !important;
                    border: 0;
                    color: #0b318f;
                    font-size: inherit;
                    font-weight: normal;
                    font-family: inherit;
                    opacity: .8;
                    span{
                        display: flex;
                        align-items: center;
                        svg{
                            width: 15px;
                            height: 15px;
                            margin: 0 3px 0 0;
                        }
                    }
                    &:hover{
                        background: transparent;
                        opacity: 1;
                    }
                }
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
    @media screen and (max-width: 1260px) {
        margin: 0.35rem 0.15rem 0 0.18rem;
    }
}
.header_arera_hidd{
    width: calc(100% - 1rem);
}
.header_left{
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
        a{
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
            .el-icon-s-upload{
                display: none;
                width: 24px;
                height: 24px;
                position: relative;
                &::before{
                    position: absolute;
                    left: 0;
                    top: 0;
                    width: 100%;
                    height: 100%;
                    content: '';
                    background-size: 17px !important;
                    background: url(../assets/images/menuIcon/uploadFile.png) no-repeat center;
                }
            }
            @media screen and (max-width: 1024px) {
                span{
                    display: none;
                }
                .el-icon-s-upload{
                    display: block;
                }
            }
        }
    }
    img{
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
        @media screen and (max-width:999px) {
            margin: 0;
        }
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
        .logged_in{
            display: flex;
            align-items: center;
            font-size: 14px;
            color: #333;
            h3, h4, h5{
                font-size: inherit;
                font-weight: normal;
            }
            span{
                font-size: inherit;
            }
            .text, .el-button{
                padding: 0.05rem 0.1rem;
                font-size: inherit;
                background: #f56c6c;
                color: #fff;
                border-radius: 0.08rem;
                line-height: 1.5;
                cursor: text;
                @media screen and (max-width:600px) {
                    font-size: 12px;
                }
            }
            .textTrue{
                background: #4326ab;
            }
            .el-button{
                background: #c9f7f5;
                color: #13c1b8;
                cursor: pointer;
                border: 0;
                font-weight: normal;
                font-family: inherit;
                &:hover{
                    opacity: .9;
                }
            }
            .info{
                display: flex;
                align-items: center;
                padding: 2px 3px 2px 8px;
                margin: 0 8px;
                border-radius: 0.08rem;
                background: rgba(67, 38, 171, 0.85);
                line-height: 1.5;
                color: #fff;
                font-size: inherit;
                cursor: text;
                h4{
                    padding: 0 5px 0 8px;
                    margin-left: 8px;
                    background: #4326ab;
                    line-height: 2;
                    border-radius: 0.05rem;
                    cursor: pointer;
                    @media screen and (max-width:600px) {
                        margin: 0;
                    }
                }
                &:hover{
                    background: rgba(67, 38, 171, 0.9);
                }
                @media screen and (max-width:600px) {
                    font-size: 12px;
                    padding: 2px 2px 2px 5px;
                    margin: 0 2px 0 5px;
                }
            }
        }
    }
    .online{
        &:before{
            // background: #0fce7c;
            display: none;
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
    .language{
    //     .background{
    //         padding: 0.05rem 0.1rem;
    //         font-size: 14px;
    //         background: #4326ab;
    //         color: #fff;
    //         border-radius: 0.08rem;
    //         line-height: 1.5;
    //         cursor: pointer;
    //         @media screen and (max-width:600px) {
    //             font-size: 12px;
    //         }
    //     }
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
        margin: 0.09rem 0.1rem;
        // height: auto;
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
    .header_left{
        font-size: 0.16rem;
        img{
            height: 30px;
        }
    }
}
</style>
