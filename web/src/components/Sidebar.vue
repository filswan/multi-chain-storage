<template>
    <div class="sidebar" :class="{'slidarShow': !collapseLocal&&bodyWidth}">
        <div class="sidebar_close el-icon-close" v-if="!collapseLocal&&bodyWidth" @click="collapseChage"></div>
        <el-menu
            class="sidebar-el-menu"
            :default-active="routerMenu"
            :collapse="collapseLocal"
            background-color="#0b318f"
            text-color="#fff"
            active-text-color="#fff"
            unique-opened
            router
        >
            <template>
                <template>
                    <!-- 折叠按钮 -->
                    <div class="header_logo pcShow" :class="{'header_left_hidd': collapseLocal}">
                        <div class="logo" v-if="!collapseLocal"><img src="@/assets/images/MCP_logo_1.png"></div>
                        <div class="logo_small" v-else><img src="@/assets/images/MCP_logo.png"></div>
                        <div class="collapse-btn-cont" @click="collapseChage">
                            <div class="header_btn">
                                <span></span>
                            </div>
                        </div>
                    </div>
                    <div class="menu_list">
                        <el-menu-item v-for="(item, i) in items" :key="i" :index="item.index" @click="sidebarLiIndex(item.name, item.index, item.type)">
                            <i :class="item.icon" style="font-size:15px"></i>
                            <span slot="title">{{ item.title }}</span>
                        </el-menu-item>
                        <li @click="logout" v-if="bodyWidth&&email" class="el-menu-item" style="padding-left: 20px; color: rgb(255, 255, 255); background-color: rgb(11, 49, 143);">
                            <svg style="width: 19px;margin: 0 8px 0 2px;" t="1629365939475" class="icon" viewBox="0 0 1126 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2181" width="200" height="200"><path d="M255.999872 460.79977h409.599795v102.399948H255.999872v153.599924l-255.999872-204.799898 255.999872-204.799898v153.599924z m-51.199974 358.39982h138.64953a409.599795 409.599795 0 1 0 0-614.399692H204.799898a511.231744 511.231744 0 0 1 409.599795-204.799898c282.777459 0 511.999744 229.222285 511.999744 511.999744s-229.222285 511.999744-511.999744 511.999744a511.231744 511.231744 0 0 1-409.599795-204.799898z" p-id="2182" fill="#ffffff"></path></svg>
                            <span>{{$t('navbar.logout')}}</span>
                        </li>
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
    data() {
        return {
            git_version:null,
            collapseLocal: this.$store.getters.collapseL == 'true'||this.$store.getters.collapseL==true?true: false,
            lanShow: false,
            bodyWidth: document.body.clientWidth<999?true:false,
            items: [
                {
                    icon: 'el-icon-s-upload',
                    index: '0',
                    title: this.$t('route.Upload_files'),
                    name: 'upload_file',
                    type: ''
                },
                {
                    icon: 'el-icon-s-deal',
                    index: '1',
                    title: this.$t('route.Deal'),
                    name: 'my_files',
                    type: ''
                },
                // {
                //     icon: 'el-icon-search',
                //     index: '2',
                //     title: this.$t('route.Search_file'),
                //     name: 'Search_file',
                //     type: ''
                // },
                {
                    icon: 'el-icon-s-billing',
                    index: '5',
                    title: this.$t('navbar.BillingHistory'),
                    name: 'billing',
                    type: ''
                },
                // {
                //     icon: 'el-icon-s-myAccount',
                //     index: '3',
                //     title: this.$t('route.myAccount'),
                //     name: 'settings',
                //     type: ''
                // },
                {
                    icon: 'el-icon-s-Stats',
                    index: '4',
                    title: this.$t('route.Stats'),
                    name: 'Stats',
                    type: ''
                },
                {
                    icon: 'el-icon-s-documentation',
                    index: '',
                    title: this.$t('route.documentation'),
                    name: '',
                    type: ''
                },
            ],
        };
    },
    computed: {
        routerMenu() {
          return this.$store.getters.routerMenu.toString()
        },
        languageMcp() {
            return this.$store.getters.languageMcp
        },
        email() {
            return this.$store.state.user.email
        },
        collapseL() {
            return this.$store.getters.collapseL
        }
    },
    watch: {
        'collapseL': function(){
            this.collapseLocal = this.$store.getters.collapseL == 'true'||this.$store.getters.collapseL==true?true: false
        }
    },
    created() {
        if(process.env.COMMITHASH){
          this.git_version = process.env.COMMITHASH.slice(0,8)
        }
        // 通过 Event Bus 进行组件间通信，来折叠侧边栏
        // bus.$on('collapse', msg => {
        //     this.collapseChild = msg;
        //     bus.$emit('collapse-content', msg);
        //     this.$emit('collapseVisit', this.collapseChild);
        //     this.$store.dispatch('setCollapse', this.collapseChild)
        // });
    },
    methods: {
        sidebarLiIndex(nameNow, index, typeNow) {
            let _this = this
            let head_title = ''
            let indexNow = Number(index)
            switch (indexNow) {
                case 2:
                    localStorage.removeItem('tabTask_name')
                    localStorage.removeItem('tabTask_search')
                    localStorage.removeItem('tabTaskMiner_search')
                    break;
                case 4:
                    localStorage.removeItem('myProfileActive')
                    break;
                default:
            }
            if(!nameNow && !indexNow){
                window.open('https://docs.filswan.com/multi-chain-payment/overview', "_blank")
                window.location.reload();
                return false
            }
            _this.$store.dispatch("setRouterMenu", Number(index));
            if(typeNow){
                _this.$router.push({ name: nameNow, params:{type:typeNow} })
                return false
            }else{
                _this.$router.push({ name: nameNow })
            }
        },
        collapseChage() {
            this.collapseLocal = !this.collapseLocal;
            this.$store.dispatch('setCollapse', this.collapseLocal)
            // bus.$emit('collapse', this.collapseLocal);
            // this.$emit('collapseVisit', this.collapseLocal);
        },
        logout() {
            var _this = this;

            let params = {};
            axios.post(_this.data_api+'auth/logout', params,{
                headers: {
                  'Authorization': "Bearer "+_this.$store.getters.accessToken
                },
            }).then((response) => {
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
                  if (error.response) {
                    console.log(error.response.headers);
                  } 
                  else if (error.request) {
                      console.log(error.request);
                  } 
                  else {
                    console.log(error.message);
                  }
                  console.log(error.config);
                
                    _this.$store.dispatch("FedLogOut").then(() => {
                        _this.$router.push("/login");
                        _this.loginShow = localStorage.getItem("mcpLoginAccessToken") ? false : true
                    });
            });

        },
        handleSetLanguage(lang){
            let _this = this
            document.body.style.height = '100vh'
            document.body.style['overflow-y'] = 'hidden'
            window.location.reload();

            _this.$i18n.locale = lang;
            _this.$store.dispatch("setLanguage", lang);

        },
    }
};
</script>

<style scoped lang="scss">
.sidebar {
    display: block;
    position: absolute;
    left: 0;
    top: 0;
    bottom: 0;
    z-index: 9999;
    transition: all .3s;
    background: #0b318f;
    .sidebar_close{
        position: absolute;
        right: 0;
        top: 0;
        font-size: 0.18rem;
        font-weight: 600;
        color: #000000;
        cursor: pointer;
        z-index: 999999;
        display: block;
        background: #f2f3f8;
        padding: 0.05rem;
    }
}
.sidebar-el-menu{
    width: 0.65rem;
    // height: calc(100% - 0.31rem);
    height: 100%;
    padding: 0;
    border-right: 0;
    // transition: width .3s;
    .menu_list{
        height: calc(100% - 1.2rem);
        padding: 0;
        overflow-y: scroll;
        // 火狐浏览器滚动条样式设置
        scrollbar-color: #ccc transparent; //滚动条轨道颜色   滚动条滑块的颜色
        scrollbar-width: none;
        scrollbar-width: thin;     //thin模式下滚动条两端的三角按钮会消失
        scrollbar-width: 0;    //默认大小
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
    .header_logo{
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: calc(100% - 0.3rem);
        height: 0.51rem;
        padding: 0.3rem 0.15rem 0.39rem;
        background-color: #0b318f;
        // transition: width .3s;
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
        width: 100%;
        height: 0.9rem;
        padding: 0.2rem 0 0.1rem;
        flex-wrap: wrap;
        justify-content: center;
        .logo_small{
            display: block;
            width: 100%;
            height: 0.34rem;
            margin: auto;
            img{
                display: block;
                height: 100%;
                max-width: 100%;
                margin: auto;
            }
        }
        .collapse-btn-cont{
            transform: rotate(180deg);
            .header_btn{
                margin: auto;
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
    li{
        align-items: center;
        display: flex;
        font-size: 0.14rem;
        font-weight: 400;
        height: 0.48rem;
        line-height: 0.48rem;
        div{
            display: flex;
            align-items: center;
            justify-content: center;
        }
        i{
            position: relative;
            height: 24px;
            &::before{
                position: absolute;
                left: 0;
                top: 0;
                width: 100%;
                height: 100%;
                content: '';
                background-size: 17px !important;
            }
        }
        .el-icon-s-home{
            &::before{
                background: url(../assets/images/menuIcon/Dashboard.png) no-repeat center;
            }
        }
        .el-icon-s-upload{
            &::before{
                background: url(../assets/images/menuIcon/uploadFile.png) no-repeat center;
            }
        }
        .el-icon-s-browse{
            &::before{
                background: url(../assets/images/menuIcon/browse.png) no-repeat center;
            }
        }
        .el-icon-s-deal{
            &::before{
                background: url(../assets/images/menuIcon/myTask.png) no-repeat center;
            }
        }
        .el-icon-s-myFs3{
            &::before{
                background: url(../assets/images/menuIcon/S3.png) no-repeat center;
            }
        }
        .el-icon-s-myProfile{
            &::before{
                background: url(../assets/images/menuIcon/myProfile.png) no-repeat center;
            }
        }
        .el-icon-s-tools{
            &::before{
                background: url(../assets/images/menuIcon/tool.png) no-repeat center;
            }
        }
        .el-icon-s-Stats{
            &::before{
                background: url(../assets/images/menuIcon/Statistics.png) no-repeat center;
            }
        }
        .el-icon-s-documentation{
            &::before{
                background: url(../assets/images/menuIcon/documentation.png) no-repeat center;
            }
        }
        .el-icon-language{
            &::before{
                background: url(../assets/images/menuIcon/language.png) no-repeat center;
                background-size: 100% !important;
            }
        }
        .el-icon-s-dataset{
            &::before{
                background: url(../assets/images/menuIcon/myAccount.png) no-repeat center;
            }
        }
        .el-icon-search{
            &::before{
                background: url(../assets/images/menuIcon/Search-File.png) no-repeat center;
            }
        }
        .el-icon-s-myAccount{
            &::before{
                background: url(../assets/images/menuIcon/myAccount.png) no-repeat center;
            }
        }
        .el-icon-s-billing{
            &::before{
                background: url(../assets/images/menuIcon/billing.png) no-repeat center;
            }
        }
    }

    .el-menu-item:hover, .is-active {
        background-color: #f2f3f8 !important;
        color: #0b318f !important;
        .el-icon-s-home{
            &::before{
                background: url(../assets/images/menuIcon/Dashboard-1.png) no-repeat center;
            }
        }
        .el-icon-s-upload{
            &::before{
                background: url(../assets/images/menuIcon/uploadFile-1.png) no-repeat center;
            }
        }
        .el-icon-s-browse{
            &::before{
                background: url(../assets/images/menuIcon/browse-1.png) no-repeat center;
            }
        }
        .el-icon-s-deal{
            &::before{
                background: url(../assets/images/menuIcon/myTask-1.png) no-repeat center;
            }
        }
        .el-icon-s-myFs3{
            &::before{
                background: url(../assets/images/menuIcon/S3-1.png) no-repeat center;
            }
        }
        .el-icon-s-myProfile{
            &::before{
                background: url(../assets/images/menuIcon/myProfile-1.png) no-repeat center;
            }
        }
        .el-icon-s-tools{
            &::before{
                background: url(../assets/images/menuIcon/tool-1.png) no-repeat center;
            }
        }
        .el-icon-s-Stats{
            &::before{
                background: url(../assets/images/menuIcon/Statistics-1.png) no-repeat center;
            }
        }
        .el-icon-s-documentation{
            &::before{
                background: url(../assets/images/menuIcon/documentation-1.png) no-repeat center;
            }
        }
        .el-icon-language{
            &::before{
                background: url(../assets/images/menuIcon/language-1.png) no-repeat center;
                background-size: 100% !important;
            }
        }
        .el-icon-s-dataset{
            &::before{
                background: url(../assets/images/menuIcon/dataset-1.png) no-repeat center;
            }
        }
        .el-icon-search{
            &::before{
                background: url(../assets/images/menuIcon/Search-File-1.png) no-repeat center;
            }
        }
        .el-icon-s-myAccount{
            &::before{
                background: url(../assets/images/menuIcon/myAccount-1.png) no-repeat center;
            }
        }
        .el-icon-s-billing{
            &::before{
                background: url(../assets/images/menuIcon/billing-1.png) no-repeat center;
            }
        }
    }
    .el-menu-item span{
        position: relative;
        /* transition: all 0.4s ease; */
    }
    .el-menu-item span::after{
        content: '';
        position: absolute;
        left: 50%;
        width: 0;
        bottom: 7px;
        margin: auto;
        height: 2px;
        background-color: #0b318f;
        transition: all 0.4s ease;
    }
    .el-menu-item span:hover::after{
        left: 0;
        width: 100%;
    }
    &:not(.el-menu--collapse) {
        width: 2.37rem;
        li{
          width: auto;
        }
    }
}
.copyStyle {
  padding: 0 0 0 0.1rem;
  background: #0b318f;
  font-size: 0.12rem;
  line-height: 0.3rem;
  color: #9c9c9c;
  border-top: 0.01rem solid rgba(255, 255, 255, 0.2);
  z-index: 999;
}

@media screen and (max-width:999px){
  .sidebar{
    top: 0;
    left: -2.37rem;
    -webkit-transition: all 0.3s ease;
    -moz-transition: all 0.3s ease;
    -ms-transition: all 0.3s ease;
    -o-transition: all 0.3s ease;
    transition: all 0.3s ease;
    .sidebar-el-menu {
        width: 2.37rem;
      .menu_list{
        height: calc(100% - 0.5rem);
        padding: 0.5rem 0 0;
      }

      &:not(.el-menu--collapse) {
          width: 2.37rem;
      }
    }
    .pcShow{
      display: none;
    }
  }
  .slidarShow{
      left: 0;
  }
}
</style>
