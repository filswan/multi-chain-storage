<template>
    <div id="dealManagement">
        <div class="backTo" @click="back">
            <span class="el-icon-back"></span>
            <span style="font-size:0.18rem;margin-left:0.05rem">{{$t('deal.backto')}}</span>
        </div>
        <div class="detailStyle">
            <el-tabs v-model="activeName" :tab-position="tabPosition" type="card" @tab-click="handleClick">
                <el-tab-pane v-for="(item, i) in offline_deals_data" :key="i" :name="''+i+''">
                    <span slot="label">
                        <!-- <i class="el-icon-success"></i>  -->
                        <img v-if="!dealCont.found.locked_fee" src="@/assets/images/error.png" />
                        <img v-else-if="dealCont.signed_dao_count >= dealCont.dao_thresh_hold && dealCont.unlock_status" src="@/assets/images/dao_success.png" />
                        <img v-else-if="dealCont.signed_dao_count >= dealCont.dao_thresh_hold && !dealCont.unlock_status" src="@/assets/images/dao_waiting.png" />
                        <img v-else src="@/assets/images/dao_waiting.png" />

                        {{item.miner_fid}}
                    </span>
                </el-tab-pane>
            </el-tabs>
            <div v-loading="loading">
                <div class="files_title">
                    {{$t('uploadFile.Deal_Detail')}} #{{dealId}}
                    <span class="title" v-if="dealId == 0">
                        <el-tooltip effect="dark" :content="$t('uploadFile.detail_tip01')" placement="top">
                            <img src="@/assets/images/info.png"/>
                        </el-tooltip>
                    </span>
                    <span v-if="!dealCont.found.locked_fee">
                        <img src="@/assets/images/error.png" />
                        <span>{{$t('uploadFile.no_fund_locked')}}</span>
                    </span>
                    <span v-else-if="dealCont.signed_dao_count >= dealCont.dao_thresh_hold && dealCont.unlock_status">
                        <img src="@/assets/images/dao_success.png" />
                        <span style="color: #3db39e;">{{$t('uploadFile.Successfully_unlocked_funds')}}</span>
                    </span>
                    <span v-else-if="dealCont.signed_dao_count >= dealCont.dao_thresh_hold && !dealCont.unlock_status">
                        <img src="@/assets/images/dao_waiting.png" />
                        <span>{{$t('uploadFile.Successfully_signed')}} {{dealCont.signed_dao_count}}/{{dealCont.dao_total_count}} </span>
                    </span>
                    <span v-else>
                        <img src="@/assets/images/dao_waiting.png" />
                        <span>{{$t('uploadFile.Waiting_for_signature')}} {{dealCont.signed_dao_count}}/{{dealCont.dao_total_count}} </span>
                    </span>
                </div>
                <div class="upload">
                    <el-row>
                        <el-col :span="8">{{$t('uploadFile.file_name')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.file_name | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_IPFSDownload')}}:</el-col>
                        <el-col :span="16">
                            <a :href="dealCont.deal.ipfs_url" target="_blank" v-if="dealCont.deal.ipfs_url" class="linkTo">{{dealCont.deal.ipfs_url}}</a>
                            <span v-else>-</span>
                        </el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Network')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.network_name | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Locked_funds')}}:</el-col>
                        <el-col :span="16">{{dealCont.found.locked_fee | NumFormatPrice}} USDC</el-col>
                        <el-col :span="8">{{$t('uploadFile.w3ss_id')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.provider | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Storage_Price')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.storage_price | NumFormatPrice}} FIL</el-col>
                        <el-col :span="8">{{$t('billing.PAYLOADCID')}}:</el-col>
                        <el-col :span="16">{{dealCont.found.payload_cid | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_ProposalCID')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.deal_cid | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.create_time')}}:</el-col>
                        <el-col :span="16">{{dealCont.found.create_at | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_MessageCID')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.message_cid | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_PieceCID')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.piece_cid | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Client_Address')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.client | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Verified_Deal')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.verified_deal?'True':'False'}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Storage_Price_Per_Epoch')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.storage_price_per_epoch | NumFormatPrice}} FIL</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Signature_Type')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.signature_type | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('my_profile.miner_add_Signature')}}:</el-col>
                        <el-col :span="16">{{dealCont.deal.signature | NumFormat}}</el-col>
                        <el-col :span="24">
                            <div class="lotupTitle">
                                {{$t('uploadFile.detail_Retrieval_Filecoin')}}
                                <el-tooltip effect="dark" :content="$t('uploadFile.detail_Retrieval_Filecoin_tooltip')" placement="top">
                                    <img src="@/assets/images/info.png"/>
                                </el-tooltip>：
                                
                                <img class="img" src="@/assets/images/copy.png" @click="copyTextToClipboard(copy_filename)" alt="">
                            </div>
                            <div class="lotupContent" :class="{'color': !dealCont.deal.provider && !dealCont.found.payload_cid}" @click="copyTextToClipboard(copy_filename)">{{copy_filename}}</div>
                        </el-col>
                    </el-row>
                        
                    <div class="title">
                        {{$t('uploadFile.detail_DAO_Signatures')}}
                        <el-tooltip effect="dark" :content="$t('uploadFile.detail_DAO_Signatures_tooltip')" placement="top">
                            <img src="@/assets/images/info.png"/>
                        </el-tooltip>
                    </div>
                    <el-table :data="daoCont" stripe style="width: 100%">
                        <el-table-column type="index" width="180">
                            <template slot-scope="scope">
                                <!-- Signature {{scope.$index+1}} -->
                                {{$t('my_profile.miner_add_Signature')}} {{scope.row.order_index?scope.row.order_index:scope.$index+1}}
                            </template>
                        </el-table-column>
                        <el-table-column prop="dao_address" :label="$t('uploadFile.detail_DAO_RKH_Address')" min-width="220">
                            <template slot-scope="scope">
                                <div class="hot-cold-box">
                                    <el-popover
                                        placement="top"
                                        trigger="hover"
                                        v-model="scope.row.daoAddressVis">
                                        <div class="upload_form_right">
                                            <p>{{scope.row.dao_address}}</p>
                                        </div>
                                        <el-button slot="reference" @click="copyTextToClipboard(scope.row.dao_address)">
                                            <img src="@/assets/images/copy.png" alt="">
                                            {{scope.row.dao_address}}
                                        </el-button>
                                    </el-popover>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column prop="tx_hash" :label="$t('billing.TRANSACTION')" min-width="220">
                            <template slot-scope="scope">
                                <div class="hot-cold-box">
                                    <el-popover
                                        placement="top"
                                        trigger="hover" width="200"
                                        v-model="scope.row.txHashVis">
                                        <div class="upload_form_right">
                                            <p>{{scope.row.tx_hash}}</p>
                                        </div>
                                        <!-- :class="{'color': dealCont.network&&dealCont.network.toLowerCase() == 'polygon'}" -->
                                        <el-button slot="reference" @click="networkLink(scope.row.tx_hash)" class="color">
                                            <!-- <img src="@/assets/images/copy.png" alt=""> -->
                                            {{scope.row.tx_hash}}
                                        </el-button>
                                    </el-popover>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column prop="dao_pass_time" :label="$t('uploadFile.detail_Time')"></el-table-column>
                        <el-table-column prop="status" :label="$t('uploadFile.file_status')">
                            <template slot-scope="scope">
                                <img src="@/assets/images/dao_success.png" v-if="scope.row.status == 1" />
                                <img src="@/assets/images/dao_waiting.png" v-else />
                            </template>
                        </el-table-column>
                    </el-table>
                </div>
            </div>
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
            tabPosition: document.documentElement.clientWidth<1024?'top':'left',
            bodyWidth: document.documentElement.clientWidth<1024?true:false,
            dealId: '',
            dealCont: {
                deal: {},
                found: {}
            },
            daoCont: [],
            copy_filename: '',
            activeName: localStorage.getItem('offlineDealsIndex')?localStorage.getItem('offlineDealsIndex'):'0',
            offline_deals_data: []
      };
    },
    computed: {},
    watch: {
        $route: function (to, from) {
            this.dealId = to.params.id
            this.getData()
        },
    },
    methods: {
        handleClick(tab, event) {
            localStorage.setItem('offlineDealsIndex', tab.index)
            this.$router.push({
                name: 'my_files_detail', 
                params: {
                    id: this.offline_deals_data[tab.index].deal_id, 
                    cid: this.$route.params.cid
                }
            })
        },
        networkLink(hash, network) {
            window.open('https://mumbai.polygonscan.com/tx/'+hash)
            // if(network && network.toLowerCase() == 'polygon'){
            //     window.open('https://mumbai.polygonscan.com/tx/'+hash)
            // }
        },
        copyTextToClipboard(text) {
            let _this = this
            let saveLang = localStorage.getItem('languageMcp') == 'cn'?"复制成功":"success";
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
                    _this.$message({
                        message: saveLang,
                        type: 'success'
                    });
                    return true;
                }
            } catch (err) {
                console.log('Oops, unable to copy');
            } finally {
                document.body.removeChild(txtArea);
            }
            return false;
        },
        back(){
            // this.$router.go(-1);//返回上一层
            this.$router.push({name: 'my_files'})
        },
        getData() {
            let _this = this
            _this.loading = true

            let dataCid = {
                payload_cid: _this.$route.params.cid,
                wallet_address: _this.$store.getters.metaAddress
            }
            axios.get(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/deal/detail/${_this.dealId}?${QS.stringify(dataCid)}`, {headers: {
            // axios.get(`./static/detail_page_response.json`, {headers: {
                    // 'Authorization':"Bearer "
            }}).then((response) => {
                let json = response.data
                _this.loading = false
                if (json.status == 'success') {
                    if(!json.data) return false
                    if(json.data.dao){
                        _this.daoCont = json.data.dao
                        _this.daoCont.map(item => {
                            item.daoAddressVis = false
                            item.txHashVis = false
                            item.dao_pass_time =  
                                item.dao_pass_time? 
                                    moment(new Date(parseInt(item.dao_pass_time * 1000))).format(
                                        "YYYY-MM-DD HH:mm:ss"
                                    )
                                    : "-";
                        })
                    }

                    _this.dealCont = json.data

                    if(json.data.deal.provider && json.data.found.payload_cid){
                        _this.copy_filename = 'lotus client retrieve --miner '+json.data.deal.provider+' '+json.data.found.payload_cid+' output-file';
                    }else{
                        _this.copy_filename = localStorage.getItem('languageMcp') == 'cn'?"还不可用。":"It's not available yet.";
                    }

                    if(!json.data.found){
                        _this.dealCont.found = {}
                    }else{
                        _this.dealCont.found.create_at = 
                            _this.dealCont.found.create_at? 
                                moment(new Date(parseInt(_this.dealCont.found.create_at))).format(
                                    "YYYY-MM-DD HH:mm:ss"
                                )
                                : "-";
                    }
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
        _this.dealId = _this.$route.params.id
        if(localStorage.getItem('offlineDeals')) _this.offline_deals_data = JSON.parse(localStorage.getItem('offlineDeals'))
        _this.getData()
        document.getElementById("content-box").scrollTop = 0;
        _this.$store.dispatch("setRouterMenu", 1);
        _this.$store.dispatch("setHeadertitle", _this.$t('navbar.deal'));
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
    .detailStyle{
        .el-tabs /deep/{
            float: left;
            height: auto;
            @media screen and (min-width: 1024px) {
                overflow: inherit;
            }
            @media screen and (max-width: 1024px) {
                float: none;
            }
            .el-tabs__header{
                margin: 0.34rem 1px 0 0;
                background: transparent;
                @media screen and (max-width: 1024px) {
                    margin: 0;
                }
                .el-tabs__nav-wrap, .el-tabs__nav-scroll{
                    @media screen and (min-width: 1024px) {
                        overflow: inherit;
                    }
                }
                .el-tabs__nav{
                    border: 0;
                    @media screen and (max-width: 1024px) {
                        display: flex;
                    }
                    .el-tabs__item{
                        min-width: 150px;
                        height: auto;
                        padding: 0.03rem 0.15rem;
                        background: #fff;
                        border-radius: 0;
                        border: 0;
                        border-top: 1px solid #eee;
                        border-right: 1px solid #eee;
                        text-align: right;
                        color: #4326ab;
                        font-weight: 600;
                        outline: none;
                        @media screen and (max-width: 1024px) {
                            min-width: auto;
                        }
                        &:first-child{
                            border-top: 0;
                        }
                        span{
                            width: 100%;
                            display: flex;
                            justify-content: space-between;
                            align-items: center;
                        }
                        i, img{
                            width: 20px;
                            height: 20px;
                            margin: 0 5px 0 0;
                            opacity: 0;
                            font-size: 16px;
                            color: #67c23a;
                        }
                    }
                    .el-tabs__item:hover {
                        text-decoration: underline;
                    }
                    .is-active{
                        position: relative;
                        border-right: 0;
                        background: #4326ab;
                        color: #fff;
                        &::before{
                            content: " ";
                            position: absolute;
                            display: block;
                            width: 0;
                            height: 0;
                            right: -10px;
                            top: calc(50% - 6px);
                            border-top: 6px solid transparent;
                            border-left: 10px solid #4326ab;
                            border-bottom: 6px solid transparent;
                        }
                        i, img{
                            opacity: 1;
                        }
                    }
                }
            }
        }
    }
    .files_title{
        display: flex;
        align-items: center;
        font-size: 0.16rem;
        font-weight: bold;
        line-height: 2;
        @media screen and (max-width:600px){
            width: 100%;
            font-size: 16px;
            flex-wrap: wrap;
        }
        img{
            width: 20px;
            height: 20px;
            margin: 0 0 0 15px;
            cursor: pointer;
        }
        span{
            display: flex;
            align-items: center;
            padding-left: 5px;
            color: red;
            font-size: 0.145rem;
            @media screen and (max-width:600px){
                font-size: 14px;
            }
        }
        .title{
            display: flex;
            align-items: center;
            justify-content: flex-start;
            padding: 0;
            line-height: 1.5;
            text-align: center;
            white-space: normal;
            color: red;
            text-shadow: 0 0 black;
            text-indent: 0;
            font-size: 13px;
            font-weight: normal;
            img{
                width: 0.16rem;
                height: 0.16rem;
                margin: 0 0 0 5px;
                cursor: pointer;
                @media screen and (max-width:600px){
                    width: 15px;
                    height: 15px;
                }
            }
        }
    }
    .upload{
        padding: 0.1rem 0.3rem 0.2rem;
        background-color: #fff;
        border-radius: 5px;
        overflow: hidden;
        border-top: 3px solid #0b318f;
        .el-row /deep/{
            width: 100%;
            max-width: 800px;
            padding: 0 0 0.15rem;
            .el-col{
                padding: 0.05rem 0;
                font-size: 0.1375rem;
                font-weight: 600;
                line-height: 1.3;
                color: #333;
                word-break: break-word;
                @media screen and (max-width:600px){
                    width: 100%;
                    font-size: 14px;
                }
                .linkTo{
                    color: #4326ab;
                    text-decoration: underline;
                }
                        .lotupTitle{
                            display: flex;
                            align-items: center;
                            margin: 0 0 0.1rem;
                            font-size: inherit;
                            color: inherit;
                            
                            img { 
                                width: 0.16rem;
                                height: 0.16rem;
                                margin: 0 0 0 3px;
                                cursor: pointer;
                                @media screen and (max-width:600px){
                                    width: 14px;
                                    height: 14px;
                                }
                            }
                            span{
                                display: block;
                                margin: 0 0 0 0.1rem;
                                color: #696262;
                            }
                        }
                        .lotupContent {
                            margin: 0;
                            border-radius: 5px;
                            background-color: rgba(0, 0, 0, 0.04);
                            color: #696262;
                            line-height: 1.3;
                            padding: 0.128rem 0.16rem;
                            font-size: inherit;
                            word-break: break-all;
                            cursor: pointer;
                            &:hover{
                                color: #333;
                            }
                        }
                        .color{
                            color: red;
                            cursor: text;
                            &:hover{
                                color: red;
                            }
                        }
            }
            .el-col:nth-child(2n+1){
                color: #4b4b4b;
                font-weight: normal;
            }
        }
        .title{
            display: flex;
            align-items: center;
            justify-content: flex-start;
            padding: 0.1rem 0;
            line-height: 1.5;
            text-align: center;
            font-size: 0.145rem;
            font-weight: 600;
            white-space: normal;
            color: #333;
            text-shadow: 0 0 black;
            text-indent: 0;
            border-top: 1px solid #e4e4e4;
            @media screen and (max-width:600px){
                font-size: 16px;
            }
            img{
                width: 0.16rem;
                height: 0.16rem;
                margin: 0 0 0 5px;
                cursor: pointer;
                @media screen and (max-width:600px){
                    width: 15px;
                    height: 15px;
                }
            }
        }
        .el-table /deep/{
            td, th{
                .cell{
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    word-break: break-word;
                    text-align: center;
                    .hot-cold-box{
                        .el-button{
                            width: 100%;
                            border: 0;
                            padding: 0;
                            margin: 0;
                            background-color: transparent;
                            font-size: 0.1372rem;
                            word-break: break-word;
                            color: #000;
                            text-align: left;
                            line-height: 0.25rem;
                            overflow: hidden;
                            text-overflow: ellipsis;
                            white-space: normal;
                            display: -webkit-box;
                            -webkit-line-clamp: 2;
                            -webkit-box-orient: vertical;
                            @media screen and (max-width:600px){
                                font-size: 13px;
                            }
                            span{
                                line-height: 0.25rem;
                                overflow: hidden;
                                text-overflow: ellipsis;
                                white-space: normal;
                                display: -webkit-box;
                                -webkit-line-clamp: 2;
                                -webkit-box-orient: vertical;
                            }
                            img{
                                display: none;
                                float: left;
                                width: 0.17rem;
                                margin-top: 0.03rem;
                            }
                        }
                        .color{
                            color: #4326ab;
                            text-decoration: underline;
                        }
                        .el-button:hover{
                            img{
                                display: inline-block;
                            }
                        }
                    }
                    a{
                        color: #4326ab;
                        &:hover{
                            text-decoration: underline;
                        }
                    }
                    img{
                        width: 30px;
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
            .el-row /deep/{
                
            }
        }
        
    }
}
@media screen and (max-width:640px){
    #dealManagement{
        .upload{
            .el-row /deep/ {
                
            }
        }
    }
}

</style>
