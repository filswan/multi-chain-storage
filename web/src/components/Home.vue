<template>
    <div class="wrapper">
        <el-col :xs="8" :sm="8" :md="5" :lg="4" class="side" :class="{'slidarShow': collapseP&&bodyWidth}"><v-sidebar></v-sidebar></el-col>
        <el-col :xs="24" :sm="24" :md="19" :lg="20" class="content-box" id="content-box">
            <v-head :meta="meta" @getMetamaskLogin="getMetamaskLogin" 
                    :netId="netId" @getNetId="changeNet"></v-head>
            <div id="headerMb" v-if="bodyWidth">
                <div class="headerMb" v-if="email">
                    {{headertitle}}
                </div>
            </div>
            <div class="content">
                <div class="content_body" :class="{'stats': $route.name == 'Stats'}">
                    <el-alert type="warning" effect="dark" center show-icon v-if="metaAddress&&!(networkID==137)">
                        <div slot="title">
                            {{$t('fs3Login.toptip_01')}} {{metaNetworkInfo.name}} {{$t('fs3Login.toptip_02')}} 
                            <span @click="changeNet(137)">{{$t('fs3Login.toptip_Network')}}</span>.
                            <p v-if="networkID == 80001 || networkID == 97">
                                {{$t('fs3Login.toptip_04')}} 
                                {{metaNetworkInfo.name}}
                                {{$t('fs3Login.toptip_04_1')}} 
                                <a href="https://calibration-mcs.filswan.com/#/metamask_login" target="_blank">calibration-mcs.filswan.com</a>.
                            </p>
                        </div>
                    </el-alert>
                    <transition name="move" mode="out-in">
                        <keep-alive :include="tagsList">
                            <router-view @getMetamaskLogin="getMetamaskLogin"></router-view>
                        </keep-alive>
                    </transition>
                </div>
                <div class="fes-icon">
                    <div class="fes-icon-copy">
                        <span>© 2022 FilSwan</span>
                        <el-divider direction="vertical"></el-divider>
                        <a href="https://www.filswan.com/" target="_block">filswan.com</a>

                    </div>
                </div>
                <el-backtop target=".content"></el-backtop>
            </div>
        </el-col>
    </div>
</template>

<script>
import vHead from './Header.vue';
import vSidebar from './Sidebar.vue';
// import bus from './bus';
export default {
    data() {
        return {
            tagsList: [],
            collapse: false,
            collapseP: this.$store.getters.collapseL == 'true'||this.$store.getters.collapseL==true?true: false,
            bodyWidth: document.body.clientWidth<992?true:false,
            tabshow: localStorage.getItem('tabTask_name') == 'User_Profile'?true:false,
            share_img1: require('@/assets/images/landing/medium.png'),
            share_img2: require('@/assets/images/landing/twitter.png'),
            share_img3: require('@/assets/images/landing/github-fill.png'),
            share_img5: require('@/assets/images/landing/facebook-fill.png'),
            share_img7: require('@/assets/images/landing/slack.png'),
            share_img8: require('@/assets/images/landing/youtube.png'),
            share_img9: require('@/assets/images/landing/telegram.png'),
            share_img10: require('@/assets/images/landing/discord.png'),
            share_logo: require('@/assets/images/landing/logo_small.png'),
            meta: false,
            netId: 0
        };
    },
    components: {
        vHead,
        vSidebar
    },
    computed: {
        headertitle() {
            return this.$store.getters.headertitle
        },
        routerMenu() {
            return this.$store.getters.routerMenu
        },
        email() {
            return this.$store.state.user.email
        },
        collapseL() {
            return this.$store.getters.collapseL
        },
        metaAddress() {
            return this.$store.getters.metaAddress
        },
        networkID() {
            return this.$store.getters.networkID
        },
        metaNetworkInfo() {
            return this.$store.getters.metaNetworkInfo?JSON.parse(JSON.stringify(this.$store.getters.metaNetworkInfo)):{}
        },
        reverse() {
            return this.$store.getters.reverse == '1' ? true : false
        }
    },
    watch: {
        'collapseL': function(){
            this.collapseP = this.$store.getters.collapseL == 'true'||this.$store.getters.collapseL==true?true: false
            this.bodyWidth = document.body.clientWidth<992?true:false
        },
        reverse: function() {
            this.init()
        }
    },
    methods: {
        changeNet(id) {
            this.netId = id
        },
        getMetamaskLogin(meta) {
            this.meta = meta
        },
        init() {
            if(this.reverse) document.body.classList.add('reverse_phase');
            else document.body.classList.remove('reverse_phase'); 
        }
    },
    mounted() {
        let bo = this.bodyWidth
        // window.onresize = function () {
        //     if((!bo && document.body.clientWidth<1024) || (bo && document.body.clientWidth>=1024)){
        //         window.location.reload();
        //     }
        // }
        this.init()
        console.log('update time: 2022-10-28') 
    }
};
</script>


<style lang="scss" scoped>
.wrapper{
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
    @media screen and (max-width:1152px){
        width: 21rem;
    }
    @media screen and (max-width:1024px){
        width: 20.48rem;
    }
    .side{
        height: 100%;
        // transition: all 0.3s ease;
        @media screen and (max-width:991px){
            position: fixed;
            left: -33.33333%;
            z-index: 9999;
        }
        @media screen and (max-width:415px){
            left: -60%;
            width: 60%;
        }
        &.slidarShow{
            left: 0;
        }
    }
    .content-box{
        height: 100%;
    }
    .content{
        position: relative;
        height: calc(100% - 1.1rem);
        overflow-y: scroll;
        // transition: all;
        // transition-duration: .3s;
        &::-webkit-scrollbar{
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

        &::-webkit-scrollbar-thumb{
            border-radius: 10px;
            box-shadow: none;
            -webkit-box-shadow: none;
            background-color: #eee;
        }
        .content_body{
            // position: relative;
            min-height: calc(100% - 0.84rem);
            .el-alert /deep/{
                position: absolute;
                left: 0;
                top: 0;
                z-index: 9;
                .el-alert__icon{
                    @media screen and (min-width: 1600px){
                        font-size: 20px;
                        width: 20px;
                    }
                }
                .el-alert__content{
                    display: flex;
                    align-items: center;
                    .el-alert__title{
                        @media screen and (min-width: 1600px){
                            font-size: 14px;
                            line-height: 1.3;
                        }
                        @media screen and (min-width: 1680px){
                            font-size: 16px;
                            line-height: 1.3;
                        }
                        @media screen and (min-width: 1800px){
                            font-size: 18px;
                            line-height: 1.3;
                        }
                        span{
                            text-decoration: underline;
                            cursor: pointer;
                        }
                        a{
                            text-decoration: underline; 
                            color: #fff;
                        }
                    }
                    .el-icon-close{
                        @media screen and (min-width: 1600px){
                            font-size: 16px;
                            line-height: 1.3;
                        }
                        @media screen and (min-width: 1800px){
                            font-size: 18px;
                            line-height: 1.3;
                        }
                    }
                }
            }
            &.stats{
                display: flex;
                flex-wrap: wrap;
            }
        }
        .fes-icon{
            background-color: #f0f0f0;
            z-index: 8;
            .fes-icon-logo{
                display: flex;
                justify-content: center;
                align-items: center;
                padding: 10px 0;
                img{
                    display: block;
                    height: 20px;
                    margin: 0 0.05rem;
                    @media screen and (max-width: 991px) {
                        height: 20px;
                    }
                }
            }
            .fes-icon-copy{
                display: flex;
                justify-content: center;
                align-items: center;
                padding: 0.3rem 0;
                span, a{
                    font-size: 0.16rem;
                    color: #888;
                    line-height: 1.5;
                }
                a{
                    &:hover{
                        color: #409eff;
                    }
                }
                .el-divider--vertical /deep/{
                    height: 15px;
                }
            }
        }
    }
}
</style>

