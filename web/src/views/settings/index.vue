<template>
    <div id="dealManagement" v-loading="loading">
        <div class="upload">
            <div class="upload_form_detail">
                <el-tabs v-model="activeName">
                    <el-tab-pane :label="$t('miner.tab_left')" name="User_Profile">
                        <div class="resetPass">
                            <h2>{{$t('my_profile.Change_password')}}</h2>

                            <div class="formStyle" :class="verify.mailPassword.tipsbox ? 'border-red' : ''">
                                <h4>{{$t('my_profile.pass_current')}}</h4>
                                <el-input v-model="formData.mail.password" type="password" :placeholder="generateMailResetPassword('mailResetPassword_pw')"></el-input>
                                <h6 v-if="passwordWrong">{{$t('my_profile.wrong_password')}}</h6>
                            </div>
                            <div class="form-tips" v-show="verify.mailPassword.tipsbox">{{verify.mailPassword.tips}}</div>

                            <div class="formStyle" :class="verify.newPasswordVerify.tipsbox ? 'border-red' : ''">
                                <h4>{{$t('my_profile.pass_new')}}</h4>
                                <el-input v-model="formData.mail.passwordNew" type="password" :placeholder="generateMailResetPassword('mailResetPassword_new_pw')"></el-input>
                            </div>
                            <div class="form-tips" v-show="verify.newPasswordVerify.tipsbox">{{verify.newPasswordVerify.tips}}</div>

                            <div class="formStyle" :class="verify.mailPasswordVerify.tipsbox ? 'border-red' : ''">
                                <h4>{{$t('my_profile.Reenter_new_password')}}</h4>
                                <el-input v-model="formData.mail.passwordComfirm" type="password" :placeholder="generateMailResetPassword('mailResetPassword_pwconfirm')"></el-input>
                            </div>
                            <div class="form-tips" v-show="verify.mailPasswordVerify.tipsbox">{{verify.mailPasswordVerify.tips}}</div>

                            <el-button type="primary" @click="mailResetSub">{{$t('deal.saveSetting')}}</el-button>
                        </div>
                    </el-tab-pane>
                </el-tabs>
            </div>
        </div>

        <!-- 回到顶部 -->
        <el-backtop target=".content-box" :bottom="40" :right="20"></el-backtop>
    </div>
</template>

<script>
import { generateMailResetPassword } from '@/utils/i18n'
import axios from 'axios'

export default {
    name: 'Settings',
    data() {
      return {
            loading: false,
            activeName: 'User_Profile',
            bodyWidth: document.documentElement.clientWidth<1024?true:false,
            resetPass: {
                Old_password: '',
                New_password: '',
                Confirm_password: ''
            },
            menuTabShow: 0,
            verify: {
                mailPassword: {
                tipsbox: false,
                    tips: this.generateMailResetPassword('mailResetPassword_verify_pw_tips_true')
                },
                mailPasswordVerify: {
                    tipsbox: false,
                    tips: this.generateMailResetPassword('mailResetPassword_verify_pwconfirm_tips_true')
                },
                newPasswordVerify: {
                    tipsbox: false,
                    tips: this.generateMailResetPassword('mailResetPassword_verify_pwconfirm_tips_true')
                }
            },
            // 表单数据
            formData: {
                mail: {
                    password: '', //sessionStorage.getItem("oaxLoginpassword")
                    passwordNew: '',
                    passwordComfirm: ''
                }
            },
            passwordRegular: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[^]{8,16}$/,
            passwordWrong: false,
      };
    },
    components: {},
    computed: {
        email() {
            return this.$store.state.user.email
        },
        assetNow() {
            return this.$store.state.app.assetNow
        },
        languageMcs() {
            return this.$store.getters.languageMcs
        },
        avater() {
            return this.$store.getters.avater
        },
        metaAddress() {
            return this.$store.getters.metaAddress
        }
    },
    watch: {
        'formData.mail.password': function(){
            this.passwordWrong = false
        },
    },
    methods: {
        generateMailResetPassword,
        menuToggle(data){
            this.menuTabShow = data
            localStorage.setItem('myProfile_type',data)
        },
        // 邮箱找回
        mailResetSub() {
            var _this = this
            // 邮箱密码确认
            if (!_this.formData.mail.passwordNew) {
                _this.verify.newPasswordVerify.tips = _this.generateMailResetPassword('mailResetPassword_verify_pw_tips_true')
                _this.verify.newPasswordVerify.tipsbox = true
                return false
            } else if (!_this.passwordRegular.test(_this.formData.mail.passwordNew)) {
                _this.verify.newPasswordVerify.tips = _this.generateMailResetPassword('mailResetPassword_verify_pw_tips_rule')
                _this.verify.newPasswordVerify.tipsbox = true
                return false
            } else if (_this.formData.mail.passwordNew !== _this.formData.mail.passwordComfirm) {
                _this.verify.mailPasswordVerify.tips = _this.generateMailResetPassword('mailResetPassword_verify_pw_tips_same')
                _this.verify.mailPasswordVerify.tipsbox = true
                _this.verify.newPasswordVerify.tips = _this.generateMailResetPassword('mailResetPassword_verify_pw_tips_same')
                _this.verify.newPasswordVerify.tipsbox = true
                return false
            } else {
                _this.verify.mailPasswordVerify.tipsbox = false
                _this.verify.newPasswordVerify.tipsbox = false
            }
            var reqData = {
                email: _this.$store.state.user.email,
                currentPassword: _this.formData.mail.password,
                password: _this.formData.mail.passwordNew,
                repeatPassword: _this.formData.mail.passwordComfirm,
                wallet_address: _this.metaAddress
            }
            _this.loading = true
            axios.post(process.env.BASE_API+'user/update_login_password', reqData,{
                headers: {
                //   'Authorization': "Bearer "+_this.$store.getters.accessToken
                },
            }).then((response) => {
                _this.loading = false
                if (response.data.status == 'success') {
                    _this.$store.dispatch("FedLogOut").then(() => {
                        _this.$message({
                            message: response.data.status,
                            type: 'success'
                        });
                        _this.$router.push("/login");
                    });
                } else {
                    _this.$message.error(response.data.message);
                }
            }).catch(function (error) {
                console.log(error);
                _this.passwordWrong = true
                _this.loading = false
            });
        },
    },
    mounted() {
        let _this = this
        document.getElementById('content-box').scrollTop = 0
        _this.$store.dispatch('setRouterMenu', 3)
        _this.$store.dispatch('setHeadertitle', _this.$t('route.myAccount'))
    },
    filters: {
        priceFilter(value) {
            let realVal = "";
            if (!isNaN(value) && value !== "") {
                let tempVal = parseFloat(value).toFixed(19);
                realVal = tempVal.substring(0, tempVal.length - 1) + " FIL";
            } else {
                realVal = "-";
            }
            return realVal;
        },
        NumFormat (value) {
            if(!value) return '-';
            return value
        }
    },
};
</script>


<style scoped lang="scss">
#dealManagement{
    padding: 0.25rem 0.2rem 0.5rem;
    .upload{
        padding: 0.1rem 0.35rem 0.2rem 0.2rem;
        margin-bottom: 0.2rem;
        background-color: #fff;
        border-radius: 0.1rem;
        overflow: hidden;
        .upload_form_detail{
            display: flex;
            justify-content: center;
            // align-items: center;
            flex-wrap: wrap;
            margin: 0 0 0.1rem;
            .el-form /deep/{
                // float: left;
                width: calc(65% - 0.45rem);
                margin: 0 0.3rem 0 0.15rem;
                .el-form-item{
                    display: flex;
                    align-items: center;
                    // flex-wrap: wrap;
                    width: auto;
                    min-height: 32px;
                    margin: 0 auto .08rem;
                    .el-form-item__label{
                        width: auto;
                        max-width: 2rem;
                        line-height: 1.5;
                        text-align: left;
                        white-space: nowrap;
                        color: #404040;
                    }
                    .el-form-item__content{
                        display: flex;
                        align-items: center;
                        flex-wrap: wrap;
                        width: 100%;
                        .el-input{
                            width: auto;
                            .el-input__inner{
                                height: 0.32rem;
                                line-height: 0.32rem;
                            }
                        }
                        .el-form-item__error {
                            padding-top: 0;
                            margin: 0 0.1rem;
                            position: relative;
                            float: right;
                        }
                        .upload-demo{
                            .el-upload--text{
                                float: left;
                                width: auto;
                                height: auto;
                                text-align: left;
                                border: 0;
                                .el-button--primary{
                                    height: 0.32rem;
                                    padding: 0 0.3rem;
                                    line-height: 0.32rem;
                                    background-color:transparent;
                                    border: 1px solid #0b318f;
                                    color: #0b318f;
                                    font-size: 0.1372rem;
                                }
                            }
                            .el-upload-list{
                                width: 100%;
                                float: none;
                                clear: both;
                            }
                        }
                        .el-upload__tip{
                            // float: left;
                            margin: 0 0 0 0.1rem;
                            color: #737373;
                            line-height: 1;
                        }
                        .el-radio{
                            .el-radio__inner{
                                border-color: #d9d9d9;
                                background-color: #d9d9d9;
                            }
                        }
                        .el-radio.is-checked{
                            .el-radio__inner{
                                border-color: #0b318f;
                                background-color: #0b318f;
                            }
                            .el-radio__inner::after{
                                width: 6px;
                                height: 6px;
                            }
                        }
                        .el-tag, .el-button--small{
                            margin: 0 5px 0 0;
                        }
                    }
                }
            }
            .el-tabs /deep/{
                width: 100%;
                margin: 0.1rem 0 0;
                .el-tabs__header{
                    display: none;
                    padding: 0 0 0.15rem;
                    margin: 0 0 0.15rem;
                    border-bottom: 0.02rem solid #060d9d;
                }
                .el-tabs__nav{
                    .el-tabs__active-bar{
                        display: none;
                    }
                    .el-tabs__item{
                        position: relative;
                        float: left;
                        width: 1.7rem;
                        height: 35px;
                        padding: 0 !important;
                        margin: 0 0.3rem 0 0;
                        background-color: #e6eaf5;
                        line-height: 35px;
                        border-radius: 0.05rem;
                        text-align: center;
                        color: #0a318e;
                        font-size: 0.1372rem;
                    }
                    .el-tabs__item.is-active{
                        background-color: #0a318e;
                        color: #fff;
                    }
                    #tab-View_Miner_Profile::after{
                        position: absolute;
                        content: '';
                        left: -0.15rem;
                        top: 50%;
                        width: 0.02rem;
                        height: 0.2rem;
                        margin-top: -0.1rem;
                        background-color: #060d9d;
                    }
                }
                .el-tabs__nav-wrap::after{
                    display: none;
                }
                .resetPass{
                    float: left;
                    width: 100%;
                    max-width: 3rem;
                    margin: 0 0 0 0.3rem;
                    h2{
                        margin: 0;
                        font-size: 0.1972rem;
                        font-weight: 500;
                        color: #000;
                        line-height: 0.3rem;
                    }
                    .formStyle{
                        margin: 0.15rem 0 0;
                    }
                    .el-button{
                        float: left;
                        margin: 0.2rem 0 0;
                        padding: 0.1rem 0.15rem;
                        border-radius: 0.08rem;
                        font-size: 0.14rem;
                    }
                    .form-tips{
                        font-size: 0.1371rem;
                        color: #ff5656;
                    }
                }
                .user_contRight{
                    float: left;
                    width: calc(100% - 2rem);
                    h4{
                        margin: 0 0 0.2rem;
                        font-size: 0.15rem;
                        font-weight: normal;
                        color: #000;
                        line-height: 1.5;
                    }
                    h6{
                        font-size: 0.12rem;
                        font-weight: normal;
                        color: #222;
                        line-height: 1.3;
                        a{
                            color: #04a2a2;
                            text-decoration: underline;
                        }
                    }
                    .checkDataCap{
                        float: left;
                        margin: 0 0 0.2rem;
                        padding: 0.1rem 0.15rem;
                        border-radius: 0.05rem;
                        font-size: 0.1372rem;
                        font-weight: 600;
                        color: #409EFF;
                        background-color: #fff;
                        border: 0.015rem solid #409EFF;
                    }
                    .form_table{
                        margin: 0.24rem 0 0.1rem;
                        .table_cell{
                            .cell{
                                padding: 0;
                                font-size: 0.13rem;
                                font-weight: 500;
                                word-break: break-word;
                                color: #000;
                                text-align: center;
                                line-height: 0.23rem;
                            }

                        }
                        .el-table /deep/{
                            .el-table__header-wrapper{
                                .has-gutter{
                                    tr{
                                        background-color: #f2f2f2;
                                    }
                                }
                            }
                            tr{
                                border-radius: 0.08rem;
                                th{
                                    height: 0.5rem;
                                    padding: 0;
                                    background-color: #f2f2f2 !important;
                                    text-align: center;
                                    .cell{
                                        word-break: break-word;
                                        font-weight: 500;
                                        color: #000;
                                    }
                                }
                                th.is-leaf{
                                    border-bottom: 0;
                                }
                                th:first-child{
                                    border-top-left-radius: 0.08rem;
                                    border-bottom-left-radius: 0.08rem;
                                }
                                th:last-child{
                                    border-top-right-radius: 0.08rem;
                                    border-bottom-right-radius: 0.08rem;
                                }
                                td{
                                    border-bottom: 1px solid #f2f2f2;
                                    .cell{
                                        padding: 0;
                                        font-size: 0.1372rem;
                                        word-break: break-word;
                                        color: #000;
                                        text-align: center;
                                        line-height: 0.25rem;
                                        .el-rate__icon{
                                            font-size: 0.16rem;
                                            margin-right: 0;
                                        }
                                        .hot-cold-box{
                                            .el-button{
                                                width: auto;
                                                min-width: 0.5rem;
                                                padding: 0;
                                                background: #7ecdf4;
                                                border: 0;
                                                color: #fff;
                                                line-height: 0.3rem;
                                                cursor: pointer;
                                            }
                                        }
                                        .el-button.el-icon-upload{
                                            padding: 0 0.1rem;
                                            line-height: 0.25rem;
                                            font-size: 0.1372rem;
                                        }
                                        .bot{
                                        display: flex;
                                        justify-content: center;
                                        align-items: center;
                                        p{
                                            font-size: 0.1372rem;
                                            padding: 0 0.08rem;
                                                margin: 0 0.05rem;
                                                border: 1px solid #0b318f;
                                                border-radius: 0.05rem;
                                                cursor: pointer;
                                        }
                                        p.color{
                                            background: #0b318f;
                                            color: #fff;
                                        }
                                        .el-radio{
                                            margin: 0;
                                            .el-radio__input{
                                                display: none;
                                            }
                                            .el-radio__label{
                                                display: block;
                                                font-size: 0.1372rem;
                                                padding: 0 0.04rem;
                                                margin: 0 0.01rem;
                                                border: 1px solid #0b318f;
                                                border-radius: 0.05rem;
                                                cursor: pointer;
                                                line-height: 1.8;
                                            }
                                            .el-radio__input.is-checked+.el-radio__label{
                                                background: #0b318f;
                                                color: #fff;
                                            }
                                        }
                                        }
                                    }
                                }
                                td.el-table_1_column_1{
                                    .cell{
                                        // color:#0c3090
                                    }
                                }
                            }
                        }
                        .el-button{
                            float: left;
                            margin: 0 0 0.2rem;
                            padding: 0.1rem 0.15rem;
                            border-radius: 0.08rem;
                            font-size: 0.1372rem;
                        }
                        .revoke{
                            .el-button{
                                float: right;
                                padding: 0.05rem 0.1rem;
                                margin: 0 auto;
                                border-radius: 4px;
                                font-size: 0.12rem;
                            }
                        }
                    }
                    .el-alert{
                        .el-alert__icon{
                            font-size: 0.18rem;
                            width: 0.18rem;
                        }
                        .el-alert__content{
                            span{
                                display: block;
                                font-size: 0.16rem;
                            }
                            p{
                                font-size: 0.14rem;
                            }
                        }
                    }
                }
                .user_menu{
                    float: left;
                    width: 1.8rem;
                    overflow: hidden;
                    .menu_tab{
                        width: 1.5rem;
                        height: 0.35rem;
                        margin: 0 0 0.15rem 0;
                        border-radius: 0.08rem;
                        color: #0b318f;
                        font-size: 0.1372rem;
                        line-height: 0.35rem;
                        text-indent: 0.1rem;
                        cursor: pointer;
                    }
                    .menu_tab_active{
                        font-size: 0.15rem;
                        font-weight: bolder;
                    }
                }
                .user_left{
                    float: left;
                    width: 1.8rem;
                    padding: 0 0.15rem 0;
                    overflow: hidden;
                    img{
                        display: block;
                        width: auto;
                        max-width: 100%;
                        margin: auto;
                    }
                }
                .user_right{
                    float: left;
                    width: 30%;
                    margin-left: 5%;
                    overflow: hidden;
                    .formStyle{
                        margin: 0 0 0.1rem;
                    }
                }
                .formStyle{
                    margin: 0.05rem 0 0;
                    overflow: hidden;
                    h4{
                        margin: 0;
                        font-size: 0.1372rem;
                        font-weight: 500;
                        color: #000;
                        line-height: 0.3rem;
                        // text-shadow: 0 0 1px rgba(0,0,0,0.52);
                    }
                    p{
                        font-size: 0.1372rem;
                        color: #777;
                        line-height: 1.3;
                        word-break: break-word;
                    }
                    h6{
                        font-weight: normal;
                        font-size: 0.12rem;
                        line-height: 1;
                        color: red;
                    }
                    .el-input ,.el-textarea {
                        float: left;
                        .el-input__inner{
                            height: 0.24rem;
                            font-size: 0.1372rem;
                            color: #999;
                            line-height: 0.24rem;
                        }
                    }
                }
                .upload_bot{
                    display: flex;
                    justify-content: flex-end;
                    align-items: center;
                    margin: 0;
                    clear: both;
                    .el-button{
                        height: 0.35rem;
                        padding: 0 0.25rem;
                        margin-left: 0.17rem;
                        background-color: #0b318f;
                        line-height: 0.35rem;
                        font-size: 0.1372rem;
                        color: #fff;
                        border: 0;
                        border-radius: 0.08rem;
                    }
                    .white{
                        height: 0.33rem;
                        line-height: 0.33rem;
                        border: 1px solid #0b318f;
                        background:#fff;
                        color:#0b318f
                    }
                }
            }
        }
    }
    .readme_cont{
        height: 4.1rem;
        padding: 0.1rem 0.1rem;
        background-color: #fff;
        border-radius: 0.1rem;
        #content{
            width: calc(100% - 0.34rem);
            height: calc(100% - 0.2rem);
            padding: 0.1rem 0.17rem;
            font-size: 0.1372rem;
            overflow-y: auto;
        }
        #content::-webkit-scrollbar-track {
            background: #fff;
        }
        #content::-webkit-scrollbar {
            width: 6px;
            background: #fff;
        }
        #content::-webkit-scrollbar-thumb {
            background: #f2f2f2;
            border-radius: 0.08rem;
        }
    }
    .form{
        padding: 0 0.17rem 0.2rem;
        background-color: #fff;
        border-radius: 0.1rem;
        .form_top{
            display: flex;
            align-items: center;
            flex-wrap: wrap;
            .title{
                width: 100%;
                text-align: left;
                font-size: 0.1972rem;
                color: #000;
                line-height: 0.42rem;
                text-indent: 0.08rem;
            }
            .search{
                display: flex;
                align-items: center;
                justify-content: space-between;
                flex-wrap: wrap;
                width: 100%;
                overflow: hidden;
                .search_left, .search_right{
                    display: flex;
                    justify-content: flex-end;
                    // width: calc(100% - 3rem);
                }
                .upload_bot {
                    display: flex;
                    justify-content: flex-start;
                    align-items: center;
                    // width: 3rem;
                    .el-button /deep/ {
                        height: 0.34rem;
                        padding: 0 0.27rem;
                        margin: 0 0.3rem 0 0;
                        background-color: #0b318f;
                        line-height: 0.34rem;
                        font-size: 0.1372rem;
                        color: #fff;
                        border: 0;
                        border-radius: 0.08rem;
                    }
                }
                .search_left{
                    width: 50%;
                    justify-content: flex-start;
                    .el-button /deep/{
                        // width: 25px;
                        height: 0.3rem;
                        padding: 0 0.15rem;
                        background-color: #0b318f;
                        border: 1px solid #0b318f;
                        color: #fff;
                        // padding: 0 0.27rem;
                        // background-color: #fff;
                        // border: 1px solid #DCDFE6;
                        // color: #000;
                        display: flex;
                        align-items: center;
                        justify-content: center;
                        font-size: 0.1371rem;
                        border-radius: 0.08rem;
                    }
                    .el-input /deep/{
                        width: 70%;
                        .el-input__inner {
                            width: 100%;
                            border-radius: 0.05rem;
                            // border: 1px solid #f8f8f8;
                            color: #737373;
                            font-size: 0.12rem;
                            height: 0.34rem;
                            line-height: 0.34rem;
                            padding: 0 0.1rem;
                        }
                    }
                }
                span{
                    margin: auto 0.05rem auto 0.35rem;
                    font-size: 0.1372rem;
                    color: #000;
                    white-space: nowrap;
                }
                .el-button /deep/{
                    height: 0.34rem;
                    padding: 0 0.27rem;
                    margin: 0 0 0 0.1rem;
                    background-color: #0b318f;
                    color: #fff;
                    line-height: 0.34rem;
                    font-size: 0.1372rem;
                    border: 1px solid;
                    border-radius: 0.08rem;
                }
                .white{
                    height: 0.33rem;
                    line-height: 0.33rem;
                    border: 1px solid #0b318f;
                    background:#fff;
                    color:#0b318f
                }
                .el-input /deep/{
                    float: left;
                    width: auto;
                    .el-input__inner {
                        width: 3rem;
                        border-radius: 0.08rem;
                        border: 1px solid #f8f8f8;
                        color: #737373;
                        font-size: 0.12rem;
                        height: 0.24rem;
                        line-height: 0.24rem;
                        padding: 0 0.27rem;
                    }
                    .el-input__icon{
                        line-height: 0.24rem;
                    }
                }
                .el-select /deep/{
                    float: right;
                    // width: 30%;
                    .el-input__inner {
                        border-radius: 0.08rem;
                        border: 1px solid #f8f8f8;
                        color: #737373;
                        font-size: 0.12rem;
                        height: 0.24rem;
                        line-height: 0.24rem;
                        padding: 0 0.1rem;
                    }
                    .el-input__icon{
                        line-height: 0.24rem;
                    }
                }
            }
        }
        .name_all {
            display: flex;
            width: 100%;
            margin: 0.2rem 0 0;
            .el-input /deep/ {
                width: 70%;
                .el-input__inner {
                width: 100%;
                border-radius: 0.05rem;
                // border: 1px solid #f8f8f8;
                color: #737373;
                font-size: 0.12rem;
                height: 0.34rem;
                line-height: 0.34rem;
                padding: 0 0.1rem;
                }
            }
            .el-button /deep/ {
                padding: 0 0.2rem;
                margin: 0 0.15rem;
                color: #000;
                font-size: 0.12rem;
                height: 0.34rem;
                line-height: 0.34rem;
            }
            .name_contact {
                label {
                margin-right: 10px;
                font-size: 0.1371rem;
                color: #000;
                }
                .el-input /deep/ {
                width: 2rem;
                }
            }
        }
        .form_table{
            margin: 0.24rem 0 0.1rem;
            .statusStyle{
                display: inline-block;
                border: 1px solid;
                padding: 0 0.1rem;
                border-radius: 0.05rem;
                // color: inherit !important;
            }
            .el-form /deep/ {
                .el-form-item {
                margin: 0;
                .el-form-item__content {
                    display: flex;
                    align-items: center;
                    flex-wrap: wrap;
                    .minerSing{
                        width: 100%;
                        display: flex;
                        align-items: center;
                        flex-wrap: wrap;
                        margin-top: 0.1rem;
                        font-size: 0.14rem;
                        color: #000;
                        line-height: 0.35rem;
                        h1{
                            width: 100%;
                            font-size: 0.16rem;
                            font-weight: 600;
                        }
                        .left{
                            width: 1.5rem;
                        }
                        .x{
                            color: #e53e3e;
                            margin: 0 0.1rem 0 0.05rem;
                        }
                        .el-input{
                            width: calc(80% - 1.5rem);
                            .el-input__inner{
                                width: 100%;
                                border-color: inherit;
                            }
                        }
                        .minerMessage{
                            display: block;
                            border-radius: .08rem;
                            padding: .05rem;
                            white-space: pre;
                            overflow-x: auto;
                            --text-opacity: 1;
                            color: #2d3748;
                            color: rgba(45,55,72,var(--text-opacity));
                            --bg-opacity: 1;
                            background-color: #f7fafc;
                            background-color: rgba(247,250,252,var(--bg-opacity));
                        }
                        .minerSignCode{
                            display: flex;
                            align-items: center;
                            width: calc(98% - 1.5rem);
                            .minerSignCode_copy{
                                border: 1px solid #dcdfe6;
                                width: 0.5rem;
                                height: 0.3rem;
                                display: flex;
                                align-items: center;
                                justify-content: center;
                                margin-left: 0.1rem;
                                border-radius: 0.05rem;
                                font-size: 0.12rem;
                                cursor: pointer;
                            }
                        }
                    }
                    .el-input {
                    width: auto;
                    .el-input__inner {
                        width: 1.5rem;
                        border-radius: 0.05rem;
                        // border: 1px solid #f8f8f8;
                        color: #737373;
                        font-size: 0.12rem;
                        height: 0.34rem;
                        line-height: 0.34rem;
                        padding: 0 0.1rem;
                    }
                    }
                    .el-button {
                    padding: 0 0.2rem;
                    margin: 0 0.15rem;
                    color: #000;
                    font-size: 0.12rem;
                    height: 0.34rem;
                    line-height: 0.34rem;
                    }
                    .switch {
                    position: relative;
                    width: auto;
                    margin: auto 0.15rem;
                    .on {
                        position: absolute;
                        top: 0;
                        left: 5px;
                        z-index: 9;
                        font-size: 12px;
                        line-height: 24px;
                        color: #fff;
                    }
                    .off {
                        position: absolute;
                        top: 0;
                        left: 25px;
                        z-index: 9;
                        font-size: 12px;
                        line-height: 24px;
                        color: #fff;
                    }
                    .el-switch /deep/ {
                        // margin-left: 0.3rem;
                        .el-switch__core {
                        height: 0.19rem;
                        }
                        .el-switch__core:after {
                        background-color: #fff;
                        top: 0;
                        width: 0.18rem;
                        height: 0.17rem;
                        z-index: 10;
                        }
                    }
                    .el-switch.is-checked /deep/ {
                        .el-switch__core:after {
                        background-color: #0b318f;
                        }
                        .el-switch__core::after {
                        margin-left: -0.19rem;
                        }
                    }
                    }
                    .el-select {
                    .el-input {
                        .el-input__inner {
                        border-color: #dcdfe6 !important;
                        }
                    }
                    }
                }
                }
            }
            .upload_bot {
                display: flex;
                justify-content: flex-end;
                align-items: center;
                margin: 0.3rem 0.3rem 0.15rem;
                .el-button /deep/ {
                height: 0.35rem;
                padding: 0 0.4rem;
                margin-left: 0.3rem;
                background-color: #0b318f;
                line-height: 0.35rem;
                font-size: 0.1372rem;
                color: #fff;
                border: 0;
                }
            }
            .el-table /deep/{
                tr{
                    border-radius: 0.08rem;
                    th{
                        height: 0.5rem;
                        padding: 0;
                        background-color: #f2f2f2 !important;
                        text-align: center;
                        .cell{
                            word-break: break-word;
                            font-weight: 500;
                            color: #000;
                        }
                    }
                    th.is-leaf{
                        border-bottom: 0;
                    }
                    th:first-child{
                        border-top-left-radius: 0.08rem;
                        border-bottom-left-radius: 0.08rem;
                    }
                    th:last-child{
                        border-top-right-radius: 0.08rem;
                        border-bottom-right-radius: 0.08rem;
                    }
                    td{
                        border-bottom: 1px solid #f2f2f2;
                        .cell{
                            padding: 0;
                            font-size: 0.1372rem;
                            word-break: break-word;
                            color: #000;
                            text-align: center;
                            line-height: 0.25rem;
                            .miner_border{
                                width: 0.15rem;
                                height: 0.15rem;
                                margin: auto;
                                background-color: #f2f2f2;
                                border-radius: 0.05rem;
                            }
                            // .el-checkbox {
                            //     .el-checkbox__inner{
                            //         border: 0;
                            //         background-color: #f2f2f2;
                            //     }
                            // }
                            // .el-checkbox.is-checked{
                            //     .el-checkbox__inner{
                            //         background-color: #0b318f;
                            //     }
                            //     .el-checkbox__inner::after{
                            //         width: 5px;
                            //     }
                            // }
                            .el-rate__icon{
                                font-size: 0.16rem;
                                margin-right: 0;
                            }
                            .hot-cold-box{
                                .el-button{
                                    width: auto;
                                    min-width: 0.5rem;
                                    padding: 0;
                                    background: #7ecdf4;
                                    border: 0;
                                    color: #fff;
                                    line-height: 0.3rem;
                                    cursor: pointer;
                                }
                            }
                            .el-button.el-icon-upload{
                                padding: 0 0.1rem;
                                line-height: 0.25rem;
                                font-size: 0.1372rem;
                            }
                            .bot{
                              display: flex;
                              justify-content: center;
                              align-items: center;
                              p{
                                  font-size: 0.1372rem;
                                  padding: 0 0.08rem;
                                    margin: 0 0.05rem;
                                    border: 1px solid #0b318f;
                                    border-radius: 0.05rem;
                                    cursor: pointer;
                              }
                              p.color{
                                  background: #0b318f;
                                  color: #fff;
                              }
                              .el-radio{
                                  margin: 0;
                                  .el-radio__input{
                                      display: none;
                                  }
                                  .el-radio__label{
                                    display: block;
                                    font-size: 0.1372rem;
                                    padding: 0 0.04rem;
                                    margin: 0 0.01rem;
                                    border: 1px solid #0b318f;
                                    border-radius: 0.05rem;
                                    cursor: pointer;
                                    line-height: 1.8;
                                  }
                                  .el-radio__input.is-checked+.el-radio__label{
                                    background: #0b318f;
                                    color: #fff;
                                  }
                              }
                              .el-select{
                                  .el-input__inner{
                                      height: 30px;
                                      padding: 0 0.1rem !important;
                                      line-height: 30px;
                                      font-size: 0.12rem;
                                  }
                                  .el-input__suffix{
                                        display: flex;
                                        align-items: center;
                                        right: 0;
                                  }
                              }
                            }
                        }
                    }
                    td.el-table_1_column_1{
                        .cell{
                            // color:#0c3090
                        }
                    }
                }
                .current-row{
                    td{
                        border-bottom: 1px solid #f2f2f2;
                        .cell{
                            .miner_border{
                                background: #0b318f url(../../assets/images/icon_01.png) no-repeat center;
                                background-size: 80%;
                            }
                        }
                    }

                }
            }
            .el-input /deep/{
                display: block;
                width: 70%;
                clear: both;
                .el-input__inner {
                    width: 100%;
                    border-radius: 0.05rem;
                    // border: 1px solid #f8f8f8;
                    color: #737373;
                    font-size: 0.12rem;
                    height: 0.34rem;
                    line-height: 0.34rem;
                    padding: 0 0.1rem;
                }
            }
        }
        .form_pagination{
            display: flex;
            justify-content: center;
            align-items: center;
            height: 0.35rem;
            text-align: center;
            .pagination{
                display: flex;
                align-items: center;
                font-size: 0.1372rem;
                color: #000;
                .pagination_left{
                    width: 0.24rem;
                    height: 0.24rem;
                    margin: 0 0.2rem;
                    border: 1px solid #f8f8f8;
                    border-radius: 0.04rem;
                    text-align: center;
                    line-height: 0.24rem;
                    font-size: 0.16rem;
                    color: #959494;
                    cursor: pointer;
                }
            }
        }
    }
    .el-dialog__wrapper /deep/{
        .el-dialog__header{
            display: flex;
            font-weight: 600;
        }
        .el-dialog__body{
            padding: 30px 20px 0;
        }
                    p{
                        display: flex;
                        align-items: center;
                        input{
                            margin-left: 3px;
                            background: transparent;
                            border: 0;
                            box-shadow: none;
                            outline: none;
                            cursor: auto;
                            font-size: 0.14rem;
                            color: #000;
                        }
                        span{
                            margin-left: 3px;
                            cursor: pointer;
                        }
                    }
        .el-dialog__footer{
            .dialog-footer{
                display: flex;
                justify-content: flex-end;
                flex-wrap: wrap;
                .footRight{
                    width: calc(98% - 60px);
                    font-size: 0.14rem;
                    text-align: left;
                    p{
                        display: flex;
                        align-items: center;
                        input{
                            margin-left: 3px;
                            background: transparent;
                            border: 0;
                            box-shadow: none;
                            outline: none;
                            cursor: auto;
                            font-size: 0.14rem;
                            color: #000;
                        }
                        span{
                            margin-left: 3px;
                            cursor: pointer;
                        }
                    }
                }
                .el-button{
                    width: 70px;
                }
            }
        }
        .apiTipCont{
            width: 90%;
            margin: auto;
            p{
                margin: 0 0 0.1rem;
                font-size: 0.1372rem;
                color: #000;
                line-height: 1.3;
                word-break: break-word;
                // text-indent: 0.15rem;
            }
        }
    }
}


@media screen and (max-width:999px){
    #dealManagement{
        padding: 0.15rem 0.1rem 0.3rem;
        .backTo{
            margin: 0.2rem 0;
        }
        .upload{
            padding: 0.1rem;
            .upload_form{
                flex-wrap: wrap;
                .el-form /deep/{
                    width: 95%;
                }
            }
            .upload_form_detail{
                .el-tabs /deep/{
                    .el-tabs__nav{
                        .el-tabs__item{
                            width: 1.5rem;
                            margin-right: 0.1rem;
                        }
                        #tab-View_Miner_Profile::after{
                            display: none;
                        }
                    }
                    .formStyle{
                        .el-input, .el-textarea{
                            .el-input__inner, .el-textarea__inner{
                                width: 130px;
                            }
                        }
                    }
                    .user_right{
                        width: 50%;
                        .form_table {
                            .el-form /deep/{
                                .el-form-item{
                                    .el-form-item__content{
                                        display: block;
                                        .switch, .el-button{
                                            margin: 0;
                                        }
                                    }
                                }
                            }
                        }
                    }
                    .user_left{
                        width: 40%;
                        padding: 0 0 0;
                        img{
                            width: 80px;
                        }
                    }
                }
            }
        }
        .form{
            padding: 0 0.1rem 0.2rem;
            .form_top{
                .search{
                    flex-wrap: wrap;
                    height: auto;
                    .search_left{
                        width: 100%;
                        .el-input /deep/{
                            margin: 0;
                        }
                    }
                    .el-input /deep/{
                        width: 100%;
                        margin: 0.1rem 0;
                        .el-input__inner{
                            width: 100%;
                        }
                    }
                    span{
                        margin-left: 0;
                    }
                    .search_right{
                        margin: 0.2rem 0 0;
                    }
                }
            }
            .form_table {
                .el-form /deep/{
                    .el-form-item{
                        .el-form-item__content{
                            display: block;
                            .switch, .el-button{
                                margin: 0;
                            }
                        }
                    }
                }
            }
        }
    }
}
@media screen and (max-width:640px){
    #dealManagement{
        .upload{
            .upload_form_detail {
                .el-tabs /deep/{
                    .user_menu{
                        display: flex;
                        float: none;
                        clear: both;
                        width: 100%;
                        margin: 0 0 0.15rem 0;
                        border-bottom: 1px solid #c7c7c7;
                        overflow-x: scroll;
                        white-space: nowrap;
                        .menu_tab{
                            float: left;
                            width: auto;
                            padding: 0 0.1rem;
                            margin: 0 auto;
                            text-align: center;
                            text-indent: 0;
                        }
                    }
                    .user_contRight, .resetPass{
                        float: none;
                        clear: both;
                        width: calc(100% - 0.3rem);
                        margin: auto;
                    }
                    .resetPass{
                        .formStyle{
                             .el-input{
                                .el-input__inner{
                                    width: 70%;
                                }
                            }
                        }
                    }
                    .user_right{
                        width: 100%;
                        margin: 0;
                        .form_table {
                            .el-form /deep/{
                                .el-form-item{
                                    .el-form-item__content{
                                        display: block;
                                        .switch, .el-button{
                                            margin: 0;
                                        }
                                    }
                                }
                            }
                        }
                    }
                    .user_left{
                        width: 100%;
                        padding: 0 0 0;
                        img{
                            width: 80px;
                        }
                    }
                }
            }
        }
        .el-dialog__wrapper /deep/{
            .el-dialog{
                width: 90%;
                .el-dialog__body{
                    padding: 0.2rem 4% 0 0;
                    .el-form{
                        .el-form-item{
                            margin-bottom: 0.15rem;
                        }
                    }
                }
            }
            .el-dialog__footer{
                .dialog-footer{
                    .footRight{
                        width: 100%;
                        margin-top: -10px;
                        margin-bottom: 10px;
                        clear: both;
                        font-size: 0.14rem;
                        text-align: left;
                        p{
                            span{
                                cursor: pointer;
                            }
                        }
                    }
                }
            }
        }
        .form{
            .form_table {
                .el-form /deep/{
                    .el-form-item {
                        .el-form-item__content {
                            .minerSing {
                                .left, .minerSignCode, .el-input{
                                    width: 100%;
                                    .minerSignCode_copy{
                                        width: 1.5rem;
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
@media screen and (max-width:321px){
    #dealManagement{
        .upload{
            .upload_form_detail{
                .el-tabs /deep/{
                    .el-tabs__nav{
                        .el-tabs__item{
                            width: 1.2rem;
                        }
                    }
                    .user_left{
                        width: auto;
                        margin: 0 5% 0.2rem;
                        img{
                            margin: 0;
                        }
                    }
                }
            }
        }
    }
}

</style>
