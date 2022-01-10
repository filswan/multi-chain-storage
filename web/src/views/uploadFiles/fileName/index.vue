<template>
    <div id="dealManagement" v-loading="loading">
        <div class="backTo" @click="back">
            <span class="el-icon-back"></span>
            <span style="font-size:0.18rem;margin-left:0.05rem">{{$t('deal.backto')}}</span>
        </div>
        <div class="upload">
            <div class="title">{{file_name}}</div>
            <el-row :gutter="30">
                <el-col :span="6" v-for="o in 4" :key="o">
                    <el-card class="box-card">
                        <div slot="header" class="clearfix">
                            <span>f00000</span>
                        </div>
                        <div class="text item">
                            <label>{{$t('uploadFile.deal_id')}}:</label>
                            <p><span @click="toDetail">{{dealID}}</span></p>
                        </div>
                        <div class="text item">
                            <label>{{$t('deal.form_select_title01')}}</label>
                            <p>deal status</p>
                        </div>
                        <div class="text item">
                            <label>{{$t('uploadFile.update_time')}}</label>
                            <p>update time</p>
                        </div>
                    </el-card>
                </el-col>
            </el-row>
        </div>

        <!-- 回到顶部 -->
        <el-backtop target=".content-box" :bottom="40" :right="20"></el-backtop>
    </div>
</template>

<script>
import axios from 'axios'
import QS from 'qs';
import moment from "moment";
export default {
    name: 'my_files',
    data() {
      return {
            loading: false,
            bodyWidth: document.documentElement.clientWidth<1024?true:false,
            file_name: '',
            dealCont: {
                deal: {},
                found: {}
            },
            dealID: localStorage.getItem('files_deal_id'),
            payloadCID: localStorage.getItem('files_payload_cid')
      };
    },
    computed: {},
    watch: {
        $route: function (to, from) {
            this.file_name = to.params.file_name
            // this.getData()
        },
    },
    methods: {
        toDetail(){
            this.$router.push({name: 'my_files_detail', params: {id: this.dealID, cid: this.payloadCID}})
        },
        back(){
            this.$router.go(-1);//返回上一层
        },
        getData() {
            let _this = this
            _this.loading = true

            let dataCid = {
                payload_cid: _this.$route.params.cid,
                wallet_address: _this.$store.getters.metaAddress
            }
            axios.get(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/deal/detail/${_this.dealID}?${QS.stringify(dataCid)}`, {headers: {
            // axios.get(`./static/detail_page_response.json`, {headers: {
                    // 'Authorization':"Bearer "
            }}).then((response) => {
                let json = response.data
                _this.loading = false
                if (json.status == 'success') {
                    if(!json.data) return false
                }else{
                    _this.$message.error(json.message);
                    return false
                }
            }).catch(function (error) {
                console.log(error);
                _this.loading = false
            });

        }
    },
    mounted() {
        let _this = this
        _this.file_name = _this.$route.params.file_name
        // _this.getData()
        document.getElementById("content-box").scrollTop = 0;
        _this.$store.dispatch("setRouterMenu", 1);
        _this.$store.dispatch("setHeadertitle", _this.$t('navbar.swan_share'));
    },
    filters: {
        NumFormat (value) {
            if (String(value) === '0') return 0;
            if (!value) return '-';
            return value
        },
        NumFormatPrice (value) {
            if (String(value) === '0') return 0;
            if (!value) return '-';
            // 18 - 单位换算需要 / 1000000000000000000，浮点运算显示有bug
            let valueNum = String(value)
            if(value.length > 18){
                let v1 = valueNum.substring(0, valueNum.length - 18)
                let v2 = valueNum.substring(valueNum.length - 18)
                let v3 = String(v2).replace(/(0+)\b/gi,"")
                if(v3){
                    return v1 + '.' + v3
                }else{
                    return v1
                }
                return parseFloat(v1.replace(/(\d)(?=(?:\d{3})+$)/g, "$1,") + '.' + v2)
            }else{
                let v3 = ''
                for(let i = 0; i < 18 - valueNum.length; i++){
                    v3 += '0'
                }
                return '0.' + String(v3 + valueNum).replace(/(0+)\b/gi,"")
            }
        }
    },
};
</script>


<style scoped lang="scss">
#dealManagement{
    padding: 0.25rem 0.2rem 0.2rem;
    .backTo{
        display: flex;
        align-items: center;
        font-size: 0.24rem;
        line-height: 1;
        max-width: 200px;
        margin: 0 0 0.2rem;
        cursor: pointer;
    }
    .upload{
        padding: 0.1rem 0.3rem 0.2rem;
        // background-color: #fff;
        border-radius: 5px;
        overflow: hidden;
        // border-top: 3px solid #0b318f;
        .title{
            position: relative;
            padding: 0 0 0.08rem;
            margin: 0 auto 0.15rem;
            font-size: 0.2rem;
            text-align: center;
            word-break: break-all;
            &::before{
                content: '';
                position: absolute;
                left: 50%;
                bottom: 0;
                width: 50px;
                height: 3px;
                margin-left: -25px;
                background: #0b318f;
            }
            @media screen and (max-width:600px){
                font-size: 16px;
            }
        }
        .el-row /deep/{
            padding: 0.15rem 0;
            .el-col{
                padding: 0.05rem 0;
                margin: 0 auto 10px;
                font-size: 0.18rem;
                line-height: 1.3;
                color: #333;
                word-break: break-word;
                @media screen and (max-width:1260px){
                    width: 33%;
                }
                @media screen and (max-width:600px){
                    width: 50%;
                    font-size: 16px;
                }
                @media screen and (max-width:441px){
                    width: 100%;
                }
                .el-card{
                    border-radius: 0.1rem;
                    .el-card__header{
                        padding: 0.1rem;
                        text-align: center;
                        font-weight: 600;
                        border-bottom: 2px solid #999;
                    }
                    .text{
                        label{
                            display: block;
                            line-height: 1.5;
                            color: #000;
                            font-size: 0.15rem;
                            font-weight: 600;
                            @media screen and (max-width:600px){
                                font-size: 14px;
                            }
                        }
                        p{
                            margin: 0.1rem 0;
                            text-align: center;
                            color: #8f8f8f;
                            font-size: 0.14rem;
                            span{
                                font-weight: 600;
                                cursor: pointer;
                                color: #0b318f;
                                &:hover{
                                    text-decoration: underline;
                                }
                            }
                            @media screen and (max-width:600px){
                                font-size: 14px;
                            }
                        }
                    }
                }
            }
        }
    }
}


@media screen and (max-width:999px){
    #dealManagement{
        padding: 0.15rem 0.1rem 0.2rem;
        .backTo{
            margin: 0.2rem 0;
        }
        .upload{
            padding: 0.1rem;
        }
        
    }
}

</style>
