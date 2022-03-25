<template>

    <el-dialog title="Pay the file" :modal="true" :width="widthDia" :visible.sync="payVisible"
        :before-close="closeDia">
        <div class="load" v-if="hashload" v-loading="hashload"></div>
        <div class="upload_form">
            <el-form :model="payRow" status-icon ref="ruleForm" class="demo-ruleForm">
                <el-form-item prop="fileList" :label="$t('uploadFile.upload')">
                    <div>
                        <p>{{payRow.file_name}}</p>
                        <p style="display: flex;align-items: center;">
                            <svg t="1637031488880" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3310" style="width: 0.13rem;height: 0.13rem;margin: 0 7px 0 5px;"><path d="M512 1024a512 512 0 1 1 512-512 32 32 0 0 1-32 32h-448v448a32 32 0 0 1-32 32zM512 64a448 448 0 0 0-32 896V512a32 32 0 0 1 32-32h448A448 448 0 0 0 512 64z" fill="#999999" p-id="3311"></path><path d="M858.88 976a32 32 0 0 1-32-32V640a32 32 0 0 1 32-32 32 32 0 0 1 32 32v304a32 32 0 0 1-32 32z" fill="#999999" p-id="3312"></path><path d="M757.12 773.12a34.56 34.56 0 0 1-22.4-8.96 32 32 0 0 1 0-45.44l101.12-101.12a32 32 0 0 1 45.44 0 30.72 30.72 0 0 1 0 44.8l-101.12 101.76a34.56 34.56 0 0 1-23.04 8.96z" fill="#999999" p-id="3313"></path><path d="M960 773.12a32 32 0 0 1-22.4-8.96l-101.76-101.76a32 32 0 0 1 0-44.8 32 32 0 0 1 45.44 0l101.12 101.12a32 32 0 0 1-22.4 54.4z" fill="#999999" p-id="3314"></path></svg>
                            {{payRow.file_size | sizeChange}}
                        </p>
                    </div>
                </el-form-item>
                <el-form-item prop="duration">
                    <template slot="label">
                        {{$t('uploadFile.Duration')}}  
                        
                        <el-tooltip effect="dark" :content="$t('uploadFile.Duration_tooltip')" placement="top">
                            <img src="@/assets/images/info.png"/>
                        </el-tooltip>
                    </template>
                    {{payRow.duration}} {{$t('components.day')}}
                </el-form-item>
                <el-form-item prop="storage_copy">
                    <template slot="label">
                        {{$t('uploadFile.Storage_copy')}} 
                        
                        <el-tooltip effect="dark" :content="$t('uploadFile.Storage_copy_tooltip')" placement="top">
                            <img src="@/assets/images/info.png"/>
                        </el-tooltip>
                    </template>
                    5
                </el-form-item>
                <el-form-item prop="storage_cost">
                    <template slot="label">
                        {{$t('uploadFile.Estimated_Storage_Cost')}} 
                        
                        <el-tooltip effect="dark" :content="$t('uploadFile.Estimated_Storage_Cost_tooltip')" placement="top">
                            <img src="@/assets/images/info.png"/>
                        </el-tooltip>
                    </template>
                    <span  style="color:#4326ab">{{payRow.storage_cost | NumStorage}} FIL</span> 
                </el-form-item>
            </el-form>
            <div class="upload_plan">
                <div class="title" :style="{'color': pay.lock_plan_tip? '#f67e7e' : '#000'}">
                    {{$t('uploadFile.Select_Lock_Funds_Plan')}}
                    <el-tooltip effect="dark" :content="$t('uploadFile.Select_Lock_Funds_Plan_tooltip')" placement="top">
                        <img src="@/assets/images/info.png"/>
                    </el-tooltip>
                </div>
                <div class="desc">{{$t('uploadFile.latest_exchange_rate')}} {{bilingPrice}}.</div>
                <div class="upload_plan_radio">
                    <el-radio-group v-model="pay.lock_plan" @change="agreeChange">
                        <el-radio label="1" border>
                            <div class="title">{{$t('uploadFile.Low')}}</div>
                            <div class="cont">
                                {{cost.storage_cost_low}} <br/> USDC
                            </div>
                        </el-radio>
                        <el-radio label="2" border>
                            <div class="title">{{$t('uploadFile.Average')}}</div>
                            <div class="cont">
                                {{cost.storage_cost_average}} <br/> USDC
                            </div>
                        </el-radio>
                        <el-radio label="3" border>
                            <div class="title">{{$t('uploadFile.High')}}</div>
                            <div class="cont">
                                {{cost.storage_cost_high}} <br/> USDC
                            </div>
                        </el-radio>
                    </el-radio-group>
                </div>
            </div>
            <div class="upload_bot">
                <el-button type="primary" @click="payCom">{{$t('deal.Submit')}}</el-button>
            </div>
        </div>
    </el-dialog>

</template>

<script>
    export default {
        name: "pay_tip",
        data() {
            return {
                widthDia: document.body.clientWidth<=600?'95%':'600px',
                inputAmount: /^\d+(?:\.\d{0,8})?[\s]{0,5}/,
                inputG: /^[1-9]\d*$/,
                hashload: false,
                pay: {
                    amount: '',
                    lock_plan_tip: false,
                    lock_plan: ''
                }
            };
        },
        props: ['payVisible', 'payRow', 'cost', 'bilingPrice'],
        components: {},
        computed: {
            metaAddress() {
                return this.$store.getters.metaAddress
            }
        },
        methods: {
            closeDia() {
                this.$emit('getDialog', false)
            },
            agreeChange(val){
                this.pay.lock_plan_tip = false
                switch (val) {
                    case '1':
                        this.pay.amount = this.cost.storage_cost_low
                        break;
                    case '2':
                        this.pay.amount = this.cost.storage_cost_average
                        break;
                    case '3':
                        this.pay.amount = this.cost.storage_cost_high
                        break;
                    default:
                        this.pay.amount = this.cost.storage_cost_low
                        return;
                }
            },
            payCom() {
                if(!this.pay.lock_plan || parseFloat(this.pay.amount)<=0){
                    this.pay.lock_plan_tip = true
                    return false
                }else{
                    this.$emit('getDialog', false, this.pay.amount)
                }
            }
        },
        mounted() {
            
        },
        watch: {
            
        },
        filters: {
            NumStorage(value) {
                if (!value) return "-";
                return value.toFixed(10);
            },
            sizeChange(bytes){
                if (bytes === 0) return '0 B';
                if (!bytes) return "-";
                var k = 1024, // or 1000
                    sizes = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
                    i = Math.floor(Math.log(bytes) / Math.log(k));

                if (Math.round((bytes / Math.pow(k, i))).toString().length > 3) {
                    // 判断大小是999999999左右，解决会显示成1.00e+3科学计数法
                    i += 1
                }
                // return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
                return Number(bytes / Math.pow(k, i)) + ' ' + sizes[i];
            }
        }
    };
</script>


<style scoped lang="scss">
.el-dialog__wrapper /deep/{
    .load{
        position: absolute;
        top: 0;
        left: 0;
        right: 0;
        bottom: 0;
        .el-loading-mask{
            .el-loading-spinner{
                height: 100%;
                top: 0;
                margin: 0;
                display: flex;
                justify-content: center;
                align-items: center;
            }
        }
    }
    .el-dialog{
        background: #fff;
        box-shadow: 0 0 13px rgba(128,128,128,0.8);
        .el-dialog__header{
            padding: 0.2rem 0.2rem 0.1rem;
            display: flex;
            .el-dialog__title{
                color: #222222;
                font-size: 0.18rem;
                font-weight: 600;
            }
        }
        .el-dialog__body{
            padding: 0.2rem 0.2rem 0;
            .upload_form{
                // display: flex;
                // align-items: baseline;
                width: 100%; 
                margin: auto; 
                justify-content: flex-start;
                .el-form{
                    width: 96%;
                    margin: 0 2%;
                    .el-form-item{
                        display: flex;
                        align-items: center;
                        width: auto;
                        margin: 0.15rem auto;
                        .el-form-item__label{
                            display: flex;
                            justify-content: flex-end;
                            align-items: center;
                            width: 47%;
                            padding: 0 3% 0 0;
                            // max-width: 2rem;
                            line-height: 1.5;
                            text-align: left;
                            font-size: 0.1372rem;
                            white-space: normal;
                            color: #000;
                            font-weight: 500;
                            text-shadow: 0 0 black;
                            text-align: right;
                            img{
                                width: 0.16rem;
                                height: 0.16rem;
                                margin: 0 0 0 5px;
                                cursor: pointer;
                            }
                            &::before{
                                display: none;
                            }
                        }
                        .el-form-item__content{
                            display: flex;
                            align-items: center;
                            font-size: 0.1372rem;
                            white-space: normal;
                            word-break: break-word;
                            line-height: 1.5;
                            color: #666;
                            h4{
                                width: 100%;
                                font-size: 0.1372rem;
                                font-weight: 500;
                                line-height: 1.7;
                            }
                            h5{
                                width: 90%;
                                margin-top: 5px;
                                font-size: 0.11rem;
                                font-weight: 500;
                                line-height: 1.2;
                                color: #737373;
                            }
                            .el-tag, .el-button--small{
                                margin: 0 5px 5px 0;
                            }
                            .el-input{
                                width: auto;
                                .el-input__inner{
                                    height: 0.32rem;
                                    font-size: 0.1372rem;
                                    line-height: 0.32rem;
                                }
                                .el-input__suffix{
                                    display: none;
                                }
                            }
                            .el-form-item__error {
                                padding-top: 0;
                                margin: 0 0.1rem;
                                position: relative;
                                float: right;
                            }
                            .el-textarea{
                                width: 90% !important;
                            }
                            .upload-demo{
                                display:flex;
                                align-items: center;
                                flex-wrap: wrap;
                                .el-upload-list__item:first-child{
                                    margin-top: 0;
                                }
                                .el-upload--text{
                                    float: left;
                                    width: auto;
                                    height: auto;
                                    text-align: left;
                                    border: 0;
                                    .el-button--primary{
                                        height: 0.32rem;
                                        padding: 0 0.2rem;
                                        margin: 0 5px 0 0;
                                        line-height: 0.32rem;
                                        background-color:transparent;
                                        border: 1px solid #2c4c9e;
                                        border-radius: 0.08rem;
                                        color: #2c4c9e;
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
                                height: 100%;
                                align-items: center;
                                display: flex;
                                margin: 0 0 0 0.1rem;
                                color: #737373;
                                line-height: 1;
                                font-size: 0.12rem;
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
                        }
                    }
                }
            }
            .upload_plan{
                width: 100%;
                margin: auto; 
                justify-content: flex-start;
                .title{
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    margin: 0.3rem 0 0;
                    line-height: 1.5;
                    text-align: center;
                    font-size: 0.15rem;
                    white-space: normal;
                    color: #000;
                    font-weight: 500;
                    text-shadow: 0 0 black;
                    text-indent: 0;
                    img{
                        width: 0.16rem;
                        height: 0.16rem;
                        margin: 0 0 0 5px;
                        cursor: pointer;
                    }
                }
                .desc{
                    margin: 0 0 0.1rem;
                    line-height: 1.5;
                    text-align: center;
                    font-size: 0.13rem;
                    white-space: normal;
                    color: #999;
                    font-weight: normal;
                }
                .upload_plan_radio{
                    .el-radio-group{
                        width: 100%;
                        display: flex;
                        justify-content: center;
                        align-items: center;
                        .el-radio{
                            min-width: 25%;
                            height: auto;
                            padding: 0 0.1rem 0.15rem;
                            margin: auto;
                            // line-height:30px;
                            .el-radio__input{
                                display: none;
                            }
                            .el-radio__label{
                                .title{
                                    margin: 0 0 0.05rem;
                                    font-size: 0.13rem;
                                }
                                .cont{
                                    font-size: 0.145rem;
                                    font-weight: bold;
                                    line-height: 1.5;
                                    text-align: center;
                                }
                            }
                        }
                        .el-radio:nth-child(3n+1){
                            .el-radio__label{
                                .cont{
                                    color: #56c4a6;
                                }
                            }
                        }
                        .el-radio:nth-child(3n+2){
                            .el-radio__label{
                                .cont{
                                    color: #4a92d3;
                                }
                            }
                        }
                        .el-radio:nth-child(3n+3){
                            .el-radio__label{
                                .cont{
                                    color: #922b26;
                                }
                            }
                        }
                        .is-checked{
                            position: relative;
                            &:after {
                                content: "";
                                display: block;
                                height: 25px;
                                width: 25px;
                                background-image: url(../assets/images/plan.png);
                                background-size: 100%;
                                position: absolute;
                                right:0;
                                top:0;
                            }
                        }
                        .el-radio:hover{
                            background-color: rgba(64,158,255,0.1);
                        }

                    }

                }
            }
            .upload_bot{
                display: flex;
                justify-content: center;
                align-items: center;
                width: 100%;
                margin: 0.25rem auto 0.15rem;
                .el-button{
                    height: 0.35rem;
                    padding: 0 0.4rem;
                    margin-left: 0;
                    background-color: #0b318f;
                    line-height: 0.35rem;
                    font-size: 0.1372rem;
                    color: #fff;
                    border: 0;
                }
            }
        }
        .dialog-footer{
            display: flex;
            align-items: center;
            justify-content: flex-end;
        }
    }
}
    @media screen and (max-width: 1024px) {
        
    }
    @media screen and (max-width: 999px) {
        
    }
    @media screen and (max-width: 599px){
    .el-dialog__wrapper /deep/{
        .el-dialog{
            .el-dialog__header{
                .el-dialog__title{
                    font-size: 0.16rem;
                }
            }
            .el-dialog__body{
                padding: 0.15rem;
                .el-form{
                    .el-form-item{
                        display: flex;
                        flex-wrap: wrap;
                        margin: 0 0 0.05rem;
                        .el-form-item__label{
                            width: 100%;
                            margin: 0;
                            text-align: left;
                        }
                        .el-form-item__content{
                            width: 100%;
                            margin: 0;
                        }
                    }
                }
            }
        }
    }
        
    }
</style>
