<template>
    <div id="dealManagement">
        <div class="backTo" @click="back">
            <span class="el-icon-back"></span>
            <span style="font-size:0.22rem;margin-left:0.05rem">{{$t('deal.backto')}}</span>
        </div>
        <div class="detailStyle">
            <div v-loading="loading">
                <div class="files_title">
                    <div class="flex_left">
                        {{$t('uploadFile.Deal_Detail')}} 
                        <b v-if="dealCont.source_file_upload_deal.deal_id || dealCont.source_file_upload_deal.deal_id==0" @click="mainnetLink(dealId)" class="golink">#{{dealId}}</b>
                        <b v-else style="margin: 0 0 0 5px;">#</b>

                        <span class="title" v-if="dealId == 0">
                            <el-tooltip effect="dark" :content="$t('uploadFile.detail_tip01')" placement="top">
                                <img src="@/assets/images/info.png"/>
                            </el-tooltip>
                        </span>
                        <span v-if="!dealCont.source_file_upload_deal.locked_fee">
                            <img src="@/assets/images/error.png" />
                            <span>{{$t('uploadFile.no_fund_locked')}}</span>
                        </span>
                        <span v-else-if="dealCont.source_file_upload_deal.unlocked">
                            <img src="@/assets/images/dao_success.png" />
                            <span style="color: #3db39e;">{{$t('uploadFile.Successfully_unlocked_funds')}}</span>
                        </span>
                        <span v-else-if="dealCont.dao_signature.length >= dealCont.dao_threshold">
                            <img src="@/assets/images/dao_waiting.png" />
                            <span>{{dealCont.dao_signatureAll == dealCont.dao_signature.length?$t('uploadFile.Successfully_signed_all')+$t('uploadFile.Successfully_signed'):$t('uploadFile.Successfully_signed')}} {{dealCont.dao_signatureAll}}/{{dealCont.dao_signature.length}} </span>
                        </span>
                        <span v-else>
                            <img src="@/assets/images/dao_waiting.png" />
                            <span>{{$t('uploadFile.Waiting_for_signature')}} {{dealCont.dao_signatureAll}}/{{dealCont.dao_signature.length}} </span>
                        </span>
                    </div>
                    <el-button type="primary" size="small" @click="getDealLogsData">{{$t('uploadFile.view_deal_logs')}}</el-button>
                </div>

                <el-tabs v-model="activeName" :tab-position="tabPosition" type="card" @tab-click="handleClick">
                    <el-tab-pane v-for="(item, i) in offline_deals_data" :key="i" :name="''+i+''">
                        <span slot="label">
                            <img v-if="!dealCont.source_file_upload_deal.locked_fee" src="@/assets/images/error.png" />
                            <img v-else-if="dealCont.source_file_upload_deal.unlocked" src="@/assets/images/dao_success.png" />
                            <img v-else-if="dealCont.dao_signature.length >= dealCont.dao_threshold" src="@/assets/images/dao_waiting.png" />
                            <img v-else src="@/assets/images/dao_waiting.png" />

                            {{item.miner_fid}}
                        </span>
                    </el-tab-pane>
                </el-tabs>

                <div class="upload">
                    <el-row :class="{'elColLeftEn': languageMcs === 'en', 'elColLeftZh': languageMcs === 'cn'}">
                        <el-col :span="8">{{$t('uploadFile.file_name')}}:</el-col>
                        <el-col :span="16">{{dealCont.source_file_upload_deal.file_name | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_IPFSDownload')}}:</el-col>
                        <el-col :span="16">
                            <a :href="dealCont.source_file_upload_deal.ipfs_url" target="_blank" v-if="dealCont.source_file_upload_deal.ipfs_url" class="linkTo">{{dealCont.source_file_upload_deal.ipfs_url}}</a>
                            <span v-else>-</span>
                        </el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Network')}}:</el-col>
                        <el-col :span="16">{{dealCont.source_file_upload_deal.network_name | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Locked_funds')}}:</el-col>
                        <el-col :span="16">{{dealCont.source_file_upload_deal.locked_fee | NumFormatPrice}} USDC</el-col>
                        <el-col :span="8">{{$t('uploadFile.w3ss_id')}}:</el-col>
                        <el-col :span="16" v-if="dealId == 0">-</el-col>
                        <el-col :span="16" v-else>{{dealCont.source_file_upload_deal.provider | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Storage_Price')}}:</el-col>
                        <el-col :span="16" v-if="dealId == 0">-</el-col>
                        <el-col :span="16" v-else>{{dealCont.source_file_upload_deal.storage_price | NumFormatPrice}} FIL</el-col>
                        <el-col :span="8">{{$t('billing.PAYLOADCID')}}:</el-col>
                        <el-col :span="16">{{dealCont.source_file_upload_deal.car_file_payload_cid | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_ProposalCID')}}:</el-col>
                        <el-col :span="16" v-if="dealId == 0">-</el-col>
                        <el-col :span="16" v-else>{{dealCont.source_file_upload_deal.deal_cid | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.create_time')}}:</el-col>
                        <el-col :span="16" v-if="dealId == 0">-</el-col>
                        <el-col :span="16" v-else>{{dealCont.source_file_upload_deal.created_at | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_MessageCID')}}:</el-col>
                        <el-col :span="16">{{dealCont.source_file_upload_deal.message_cid | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_PieceCID')}}:</el-col>
                        <el-col :span="16" v-if="dealId == 0">-</el-col>
                        <el-col :span="16" v-else>{{dealCont.source_file_upload_deal.piece_cid | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Client_Address')}}:</el-col>
                        <el-col :span="16" v-if="dealId == 0">-</el-col>
                        <el-col :span="16" v-else>{{dealCont.source_file_upload_deal.client | NumFormat}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Verified_Deal')}}:</el-col>
                        <el-col :span="16" v-if="dealCont.source_file_upload_deal.verified_deal==null">{{dealCont.source_file_upload_deal.verified_deal | NumFormat}}</el-col>
                        <el-col :span="16" v-else>{{dealCont.source_file_upload_deal.verified_deal?'True':'False'}}</el-col>
                        <el-col :span="8">{{$t('uploadFile.detail_Storage_Price_Per_Epoch')}}:</el-col>
                        <el-col :span="16" v-if="dealId == 0">-</el-col>
                        <el-col :span="16" v-else>{{dealCont.source_file_upload_deal.storage_price_per_epoch | NumFormatPrice}} FIL</el-col>
                        <el-col :span="24" class="lotupTitle">
                            <el-col :span="8">
                                {{$t('uploadFile.detail_Retrieval_Filecoin')}}
                                <el-tooltip effect="dark" :content="$t('uploadFile.detail_Retrieval_Filecoin_tooltip')" placement="top">
                                    <img src="@/assets/images/info.png"/>
                                </el-tooltip>：
                            </el-col>
                            <el-col :span="16">
                                <img class="img" src="@/assets/images/copy.png" @click="copyTextToClipboard(copy_filename)" alt="">
                            </el-col>
                        </el-col>
                        <el-col :span="24" class="lotupContent" :class="{'color': !dealCont.source_file_upload_deal.provider && !dealCont.source_file_upload_deal.car_file_payload_cid}" @click="copyTextToClipboard(copy_filename)">
                            {{copy_filename}}
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
                        <el-table-column prop="wallet_signer" :label="$t('uploadFile.detail_DAO_RKH_Address')" min-width="220">
                            <template slot-scope="scope">
                                <div class="hot-cold-box">
                                    <el-popover
                                        placement="top"
                                        trigger="hover"
                                        v-model="scope.row.daoAddressVis">
                                        <div class="upload_form_right">
                                            <p>{{scope.row.wallet_signer}}</p>
                                        </div>
                                        <el-button slot="reference" @click="copyTextToClipboard(scope.row.wallet_signer)">
                                            <img src="@/assets/images/copy.png" alt="">
                                            {{scope.row.wallet_signer}}
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
                                        <el-button slot="reference" @click="networkLink('https://mumbai.polygonscan.com/tx/'+scope.row.tx_hash)" class="color">
                                            <!-- <img src="@/assets/images/copy.png" alt=""> -->
                                            {{scope.row.tx_hash}}
                                        </el-button>
                                    </el-popover>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column prop="create_at" :label="$t('uploadFile.detail_Time')"></el-table-column>
                        <el-table-column prop="status" :label="$t('uploadFile.file_status')">
                            <template slot-scope="scope">
                                <img src="@/assets/images/dao_success.png" v-if="scope.row.status&&scope.row.status.toLowerCase() == 'success'" />
                                <img src="@/assets/images/dao_waiting.png" v-else />
                            </template>
                        </el-table-column>
                    </el-table>
                </div>
            </div>
        </div>
        
        <el-dialog
            :visible.sync="dialogVisible"
            :width="width" custom-class="dealLogs">
            <div slot="title" class="dialog-title">
                {{$t('uploadFile.deal_logs')}}
                <el-tooltip effect="dark" placement="top">
                    <div slot="content">
                        <a :href="webLink" target="_blank" class="weblinkStyle">{{webLink}}</a>
                    </div>
                    <img src="@/assets/images/info.png"/>
                </el-tooltip>
            </div>
            <div class="block" v-loading="loadlogs">
                <el-timeline v-if="dealLogsData.length>0">
                    <el-timeline-item v-for="(item, n) in dealLogsData" :key="n" :timestamp="item.create_at" placement="top">
                        <el-card>
                            <h4>{{item.on_chain_status}}</h4>
                            <p>{{item.on_chain_message}}</p>
                        </el-card>
                    </el-timeline-item>
                </el-timeline>
                <p v-else class="noLogs">{{$t('uploadFile.no_logs')}}</p>
            </div>
            <!-- <span slot="footer" class="dialog-footer">
                <el-button type="primary" size="small" @click="dialogVisible = false">Close</el-button>
            </span> -->
        </el-dialog>
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
            tabPosition: 'top',
            bodyWidth: document.documentElement.clientWidth<1024?true:false,
            dealId: '',
            dealCont: {
                dao_signatureAll: 0,
                source_file_upload_deal: {}
            },
            daoCont: [],
            copy_filename: '',
            activeName: localStorage.getItem('offlineDealsIndex')?localStorage.getItem('offlineDealsIndex'):'0',
            offline_deals_data: [],
            dialogVisible: false,
            width: document.body.clientWidth>600?'550px':'95%',
            loadlogs: true,
            dealLogsData: [],
            webLink: 'https://docs.filecoin.io/get-started/store-and-retrieve/store-data/#deal-states'
      };
    },
    computed: {
        languageMcs() {
            return this.$store.getters.languageMcs
        }
    },
    watch: {
        $route: function (to, from) {
            this.dealId = to.params.deal_id
            this.getData()
        },
    },
    methods: {
        handleClick(tab, event) {
            localStorage.setItem('offlineDealsIndex', tab.index)
            this.$router.push({
                name: 'my_files_detail', 
                params: {
                    id: this.offline_deals_data[tab.index].id,
                    deal_id: this.offline_deals_data[tab.index].deal_id,
                    // cid: this.$route.params.cid,
                    source_file_upload_id: this.$route.params.source_file_upload_id
                }
            })
        },
        networkLink(link) {
            window.open(link)
        },
        mainnetLink(dealId) {
            const network_name = this.dealCont.source_file_upload_deal.network_name
            switch(network_name) {
                case 'filecoin_calibration':
                    window.open(`${process.env.BASE_CALIBRATION_ADDRESS}?dealid=${dealId}`)
                    break;
                case 'filecoin_mainnet':
                    window.open(`${process.env.BASE_MAINNET_ADDRESS}?dealid=${dealId}`)
                    break;
                default:
                    if(network_name && network_name.indexOf('calibration') > -1) window.open(`${process.env.BASE_CALIBRATION_ADDRESS}?dealid=${dealId}`)
                    else window.open(`${process.env.BASE_MAINNET_ADDRESS}?dealid=${dealId}`)
                    break;
            }
        },
        copyTextToClipboard(text) {
            let _this = this
            let saveLang = localStorage.getItem('languageMcs') == 'cn'?"复制成功":"success";
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
                source_file_upload_id: _this.$route.params.source_file_upload_id,
                // payload_cid: _this.$route.params.cid,
                wallet_address: _this.$store.getters.metaAddress
            }
            axios.get(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/deal/detail/${_this.dealId}?${QS.stringify(dataCid)}`, {headers: {
            // axios.get(`./static/detail_page_response.json`, {headers: {
                    // 'Authorization':"Bearer "
            }}).then((response) => {
                let json = response.data
                let dao_signatureAll = 0
                _this.loading = false
                if (json.status == 'success') {
                    if(!json.data) return false
                    if(json.data.dao_signature){
                        _this.daoCont = json.data.dao_signature
                        _this.daoCont.map(item => {
                            if(item.create_at) dao_signatureAll += 1
                            item.daoAddressVis = false
                            item.txHashVis = false
                            item.create_at =  
                                item.create_at? 
                                    moment(new Date(parseInt(item.create_at * 1000))).format(
                                        "YYYY-MM-DD HH:mm:ss"
                                    )
                                    : "-";
                        })
                    }
                    
                    _this.dealCont = json.data
                    _this.dealCont.dao_signatureAll = dao_signatureAll
                    _this.dealCont.source_file_upload_deal.created_at = 
                        _this.dealCont.source_file_upload_deal.created_at && _this.dealCont.source_file_upload_deal.created_at != 0? 
                            moment(new Date(parseInt(String(_this.dealCont.source_file_upload_deal.created_at).length<13?_this.dealCont.source_file_upload_deal.created_at*1000:_this.dealCont.source_file_upload_deal.created_at))).format(
                                "YYYY-MM-DD HH:mm:ss"
                            )
                            : "-";

                    if(json.data.source_file_upload_deal.provider && json.data.source_file_upload_deal.car_file_payload_cid && json.data.source_file_upload_deal.deal_id != 0){
                        _this.copy_filename = 'lotus client retrieve --miner '+json.data.source_file_upload_deal.provider+' '+json.data.source_file_upload_deal.car_file_payload_cid+' ~/output-file';
                    }else{
                        _this.copy_filename = localStorage.getItem('languageMcs') == 'cn'?"还不可用。":"It's not available yet.";
                    }
                }else{
                    _this.$message.error(json.message);
                    return false
                }
            }).catch(function (error) {
                console.log(error);
                _this.loading = false
            });


        },
        getDealLogsData() {
            let _this = this
            _this.dialogVisible = true
            _this.loadlogs = true
            let obj = {
                wallet_address: _this.$store.getters.metaAddress
            }
            axios.get(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/deal/log/${_this.$route.params.id}?${QS.stringify(obj)}`, {headers: {
            // axios.get(`./static/deal_logs.json`, {headers: {
                    // 'Authorization':"Bearer "
            }}).then((response) => {
                let json = response.data
                _this.loadlogs = false
                if (json.status == 'success') {
                    if(!json.data) return false
                    if(json.data.offline_deal_log){
                        _this.dealLogsData = json.data.offline_deal_log
                        _this.dealLogsData.map(item => {
                            item.create_at =  
                                item.create_at? 
                                    moment(new Date(parseInt(item.create_at*1000))).format(
                                        "YYYY-MM-DD HH:mm:ss"
                                    )
                                    : "-";
                        })
                    }
                }else{
                    _this.$message.error(json.message);
                    return false
                }
            }).catch(function (error) {
                console.log(error);
                _this.loadlogs = false
            });
        }
    },
    mounted() {
        let _this = this
        _this.dealId = _this.$route.params.deal_id
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
.weblinkStyle{
    color: #fff; 
    &:hover{
        text-decoration: underline;
    }
}
.el-dialog__wrapper /deep/{
    display: flex;
    align-items: center;
    justify-content: center;
    .dealLogs{
        height: 80%;
        background: #fff;
        margin: auto !important;
        box-shadow: 0 0 13px rgba(128,128,128,0.8);
        border-radius: 0.2rem;
        overflow: hidden;
        .el-dialog__header{
            display: flex;
            align-items: center;
            justify-content: space-between;
            padding: 0.3rem 0.4rem;
            display: flex;
            border-bottom: 1px solid #dfdfdf;
            color: #000;
            font-size: 20px;
            .dialog-title{
                display: flex;
                align-items: center;
                color: #000;
                font-size: 0.22rem;
                font-weight: 500;
                line-height: 1;
                text-transform: capitalize;
                .el-tooltip{
                    width: 20px;
                    height: 20px;
                    margin: 0 0 0 5px;
                    @media screen and (min-width:1800px){
                        width: 22px;
                        height: 22px;
                    }
                    @media screen and (max-width:1440px){
                        width: 17px;
                        height: 17px;
                    }
                    @media screen and (max-width: 1280px){
                        width: 16px;
                        height: 16px;
                    }
                }
            }
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
            height: calc(100% - 1.35rem);
            padding: 0.2rem 0.4rem;
            font-size: 14px;
            overflow-y: scroll;
            .block{
                position: relative;
                min-height: 100px;
                .el-timeline{
                    .el-timeline-item{
                        .el-card__body{
                            h4{
                                font-size: 0.2rem;
                                color: #000;
                            }
                            p{
                                word-break: break-word;
                                font-size: 0.18rem;
                                color: #555;
                            }
                        }
                        .el-timeline-item__timestamp{
                            font-size: 0.18rem;
                            color: #555;
                        }
                    }
                }  
                .noLogs{
                    font-size: 16px;
                    text-align: center;
                }
                .el-loading-mask{
                    .el-loading-spinner{
                        top: 50%;
                    }
                }
            }
        }
        .el-dialog__footer{
            padding: 0 0.2rem 0.1rem;
            .dialog-footer{
                display: flex;
                align-items: center;
                justify-content: flex-end;
                .el-button{
                    font-size: 13px;
                }
            }
        }
    }

}
#dealManagement{
    padding: 0.3rem;
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
            .el-tabs__header{
                margin: 0;
                background: transparent;
                .el-tabs__nav-next, .el-tabs__nav-prev{
                    font-size: 18px;
                    i{
                        margin-top: 0.3rem;
                        font-weight: 600;
                    }
                }
                .el-tabs__nav{
                    display: flex;
                    border: 0;
                    @media screen and (max-width: 1024px) {
                        display: flex;
                    }
                    .el-tabs__item{
                        position: relative;
                        min-width: 120px;
                        height: auto;
                        padding: 0.1rem 0.5rem 0.03rem 0.3rem;
                        background: transparent;
                        border-radius: 0;
                        border: 0;
                        border-top: 1px solid #eee;
                        border-right: 1px solid #eee;
                        text-align: right;
                        color: #2D43E7;
                        font-weight: 600;
                        outline: none;
                        @media screen and (max-width: 1024px) {
                            min-width: auto;
                        }
                        // &:first-child{
                        //     border-top: 0;
                        //     border-top-left-radius: 0.1rem;
                        //     border-bottom-left-radius: 0.1rem;
                        // }
                        // &:last-child{
                        //     border-top: 0;
                        //     border-top-right-radius: 0.1rem;
                        //     border-bottom-right-radius: 0.1rem;
                        // }
                        span{
                            width: 100%;
                            display: flex;
                            justify-content: space-between;
                            align-items: center;
                            font-size: 0.18rem;
                        }
                        i, img{
                            width: 20px;
                            height: 20px;
                            margin: 0 15px 0 0;
                            opacity: 0;
                            font-size: 16px;
                            color: #67c23a;
                        }
                        &:before{    
                            content: '';
                            position: absolute;
                            top: 0;
                            left: 0;
                            right: 0;
                            bottom: 0;
                            z-index: -1;
                            transform: perspective(0.5em) rotateX(1.6deg);
                            transform-origin: bottom left;
                            background: #fff;
                            border-top-left-radius: 0.1rem;
                            border-top-right-radius: 0.1rem;
                        }
                    }
                    .el-tabs__item:hover {
                        text-decoration: underline;
                    }
                    .is-active{
                        position: relative;
                        border-right: 0;
                        color: #fff;
                        &:before{   
                            background: #2D43E7;
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
        justify-content: space-between;
        margin: 0.4rem 0 0.1rem;
        font-size: 0.22rem;
        font-weight: bold;
        line-height: 2;
        @media screen and (max-width:600px){
            width: 100%;
            font-size: 16px;
            flex-wrap: wrap;
        }
        .flex_left{
            display: flex;
            align-items: center;
            @media screen and (max-width:410px){
                width: 100%;
                font-size: 16px;
                flex-wrap: wrap;
            }
            .golink{
                margin: 0 0 0 5px;
                cursor: pointer;
                &:hover{
                    color: #2d43e7;
                    text-decoration: underline;
                }
            }
            img{
                width: 20px;
                height: 20px;
                margin: 0 0 0 15px;
                cursor: pointer;
                @media screen and (min-width:1800px){
                    width: 22px;
                    height: 22px;
                }
                @media screen and (max-width:1440px){
                    width: 17px;
                    height: 17px;
                }
                @media screen and (max-width:600px){
                    margin: 0;
                }
            }
            span{
                display: flex;
                align-items: center;
                padding-left: 5px;
                color: red;
                font-size: 0.18rem;
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
                    width: 20px;
                    height: 20px;
                    margin: 0 0 0 5px;
                    cursor: pointer;
                    @media screen and (min-width:1800px){
                        width: 22px;
                        height: 22px;
                    }
                    @media screen and (max-width:1440px){
                        width: 17px;
                        height: 17px;
                    }
                    @media screen and (max-width:600px){
                        width: 15px;
                        height: 15px;
                    }
                }
            }
        }
        .el-button {    
            padding: 0.1rem 0.2rem;
            font-size: 0.2rem;
            font-family: inherit;
            border-radius: 0.07rem;
            background: #4e82ff;
            background: linear-gradient(45deg,#4f8aff, #4b5eff);
            border-radius: 0.08rem;
            border: 0;
            &:hover{
                opacity: .9;
            }
            @media screen and (max-width:600px){
                margin: 0 0 5px;
            }
        }
    }
    .upload{
        padding: 0.1rem 0.3rem 0.2rem;
        background-color: #fff;
        border-bottom-left-radius: 0.1rem;
        border-bottom-right-radius: 0.1rem;
        overflow: hidden;
        border-top: 3px solid #0b318f;
        .el-row /deep/{
            width: 100%;
            padding: 0 0 0.15rem;
            .el-col{
                padding: 0.1rem 0;
                font-size: 0.2rem;
                font-weight: normal;
                line-height: 1.3;
                color: #555;
                word-break: break-word;
                @media screen and (max-width:600px){
                    width: 100%;
                    font-size: 14px;
                }
                .linkTo{
                    color: #2d43e7;
                    text-decoration: underline;
                }
            }
            .lotupTitle{
                display: flex;
                align-items: center;
                width: 100%;
                margin: 0;
                font-size: inherit;
                color: inherit;
                @media screen and (max-width:600px){
                    flex-wrap: wrap;
                }
                .el-col{
                    display: flex;
                    align-items: center;
                    @media screen and (max-width:600px){
                        width: auto;
                    }
                }
                img { 
                    width: 18px;
                    height: 18px;
                    margin: 0;
                    cursor: pointer;
                    @media screen and (min-width:1800px){
                        width: 22px;
                        height: 22px;
                    }
                    @media screen and (max-width:1280px){
                        width: 16px;
                        height: 16px;
                    }
                }
                span{
                    display: block;
                    margin: 0 0 0 0.1rem;
                    color: #696262;
                }
            }
            .lotupContent {
                margin: 0 0 0.1rem;
                border-radius: 5px;
                background-color: rgba(0, 0, 0, 0.04);
                color: #696262;
                line-height: 1.3;
                padding: 0.128rem 0.16rem;
                word-break: break-all;
                cursor: pointer;
                &:hover{
                    color: #333;
                }
            }
            .color{
                color: #F63D3D;
                cursor: text;
                &:hover{
                    color: #F63D3D;
                }
            }
            .el-col:nth-child(2n+1){
                color: #000;
            }
        }
        .elColLeftZh{
            .el-col-8{
                width: 2.5rem;
                min-width: 160px;
                @media screen and (max-width:600px){
                    width: 100%;
                }
            }
            .el-col-16{
                width: calc(100% - 3rem);
                max-width: calc(100% - 160px);
                @media screen and (max-width:600px){
                    width: 100%;
                    max-width: 100%;
                }
            }
        }
        .elColLeftEn{
            .el-col-8{
                width: 3.8rem;
                min-width: 240px;
                @media screen and (max-width:600px){
                    width: 100%;
                }
            }
            .el-col-16{
                width: calc(100% - 4.3rem);
                max-width: calc(100% - 240px);
                @media screen and (max-width:600px){
                    width: 100%;
                    max-width: 100%;
                }
            }
        }
        .title{
            display: flex;
            align-items: center;
            justify-content: flex-start;
            padding: 0.2rem 0;
            line-height: 1.5;
            text-align: center;
            font-size: 0.22rem;
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
                width: 18px;
                height: 18px;
                margin: 0 0 0 3px;
                cursor: pointer;
                @media screen and (min-width:1800px){
                    width: 22px;
                    height: 22px;
                }
                @media screen and (max-width:1280px){
                    width: 16px;
                    height: 16px;
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
                    font-size: 0.2rem;
                    color: #555;
                    .hot-cold-box{
                        .el-button{
                            width: 100%;
                            border: 0;
                            padding: 0;
                            margin: 0;
                            background-color: transparent;
                            font-size: 0.2rem;
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
            td{
                padding: 0.15rem 0;
            }
            .el-table__header-wrapper{
                border-radius: 0.2rem;
                th{
                    height: 0.7rem;
                    background-color: #e5eeff !important;
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
