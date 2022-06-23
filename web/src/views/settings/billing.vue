<template>
    <div id="dealManagement">
        <div class="upload">
            <div v-if="metaAddress" style="font-size:12px">
                <div id="billing">
                    <div class="form">
                        <div class="form_top">
                            <div class="search">
                                <div style="white-space: nowrap;">{{$t('billing.search_placeholder')}} &nbsp;</div>
                                <el-input placeholder="" v-model="searchValue" class="input-with-select" @input="searchValueChange">
                                    <el-select v-model="selectInput" slot="prepend">
                                        <el-option :label="$t('billing.search_option_filename')" value="1"></el-option>
                                        <el-option :label="$t('billing.search_option_transaction')" value="2"></el-option>
                                    </el-select>
                                </el-input>
                                <div class="search_right" :style="{'opacity': !searchValue?'0.8':'1'}">
                                    <el-button @click="search" style="background-color: #ffb822"
                                    :disabled="!searchValue">{{$t('billing.search_btn')}}</el-button>
                                    <el-button
                                    type="primary"
                                    style="background-color: #2D43E7"
                                    @click="clearAll"
                                    >
                                    {{$t('billing.clear_btn')}}
                                    </el-button>
                                </div>
                            </div>
                        </div>
                        <div class="form_table">
                            <el-table 
                            v-loading="loading" :data="tableData" style="width: 100%" 
                            :empty-text="$t('deal.formNotData')" max-height="580" @sort-change="sortChange"
                            :default-sort = "{prop: 'date', order: 'descending'}">
                                <el-table-column prop="pay_tx_hash" :label="$t('billing.TRANSACTION')" min-width="190">
                                    <template slot-scope="scope">
                                        <div class="hot-cold-box">
                                            <el-popover
                                                placement="top"
                                                trigger="hover" width="200"
                                                v-model="scope.row.txHashVis">
                                                <div class="upload_form_right">
                                                    <p>{{scope.row.pay_tx_hash}}</p>
                                                </div>
                                                <el-button slot="reference" class="resno" @click="networkLink(scope.row.pay_tx_hash, scope.row.network_name)" :class="{'color': scope.row.network_name&&scope.row.network_name.toLowerCase() == 'polygon'}">
                                                    <!-- <img src="@/assets/images/copy.png" alt=""> -->
                                                    {{scope.row.pay_tx_hash}}
                                                </el-button>
                                            </el-popover>
                                        </div>
                                    </template>                    
                                </el-table-column>
                                <el-table-column prop="pay_amount" :label="$t('billing.AMOUNT')" min-width="150" sortable="custom">
                                    <template slot-scope="scope">{{scope.row.pay_amount | balanceFilter}}</template>
                                </el-table-column>
                                <el-table-column prop="unlock_amount" :label="$t('billing.UNLOCKAMOUNT')" min-width="150" sortable="custom">
                                    <template slot-scope="scope">{{scope.row.unlock_amount | balanceFilter}}</template>
                                </el-table-column>
                                <el-table-column prop="token_name" :label="$t('billing.TOKEN')" min-width="120"></el-table-column>
                                <el-table-column prop="file_name" :label="$t('billing.FILENAME')" min-width="180" sortable="custom"></el-table-column>
                                <el-table-column prop="payload_cid" :label="$t('billing.PAYLOADCID')" min-width="140">
                                    <template slot-scope="scope">
                                        <div class="hot-cold-box">
                                            <el-popover
                                                placement="top"
                                                trigger="hover" width="300"
                                                v-model="scope.row.payloadVis">
                                                <div class="upload_form_right">
                                                    <p>{{scope.row.payload_cid}}</p>
                                                </div>
                                                <el-button slot="reference" class="resno" @click="copyTextToClipboard(scope.row.payload_cid)">
                                                    <img src="@/assets/images/copy.png" alt="">
                                                    {{scope.row.payload_cid}}
                                                </el-button>
                                            </el-popover>
                                        </div>
                                    </template>                    
                                </el-table-column>
                                <!-- <el-table-column prop="address_from" :label="$t('billing.WALLET')" min-width="140">
                                    <template slot-scope="scope">
                                        <div class="hot-cold-box">
                                            <el-popover
                                                placement="top"
                                                trigger="hover" width="300"
                                                v-model="scope.row.walletVis">
                                                <div class="upload_form_right">
                                                    <p>{{metaAddress}}</p>
                                                </div>
                                                <el-button slot="reference" @click="copyTextToClipboard(metaAddress)">
                                                    <img src="@/assets/images/copy.png" alt="">
                                                    {{metaAddress}}
                                                </el-button>
                                            </el-popover>
                                        </div>
                                    </template>                    
                                </el-table-column> -->
                                <el-table-column prop="network_name" :label="$t('billing.NETWORK')" min-width="120"></el-table-column>
                                <el-table-column prop="pay_at" :label="$t('billing.PAYMENTDATE')" min-width="140" sortable="custom"></el-table-column>
                                <el-table-column prop="unlock_at" :label="$t('billing.UNLOCKDATE')" min-width="140" sortable="custom"></el-table-column>
                                <el-table-column prop="deadline" :label="$t('billing.Deadline')" min-width="140" sortable="custom"></el-table-column>
                            </el-table>
                        </div>

                        <div class="form_pagination">
                            <div class="pagination">
                                <el-pagination
                                    :total="parma.total"
                                    :page-size="parma.limit"
                                    :current-page="parma.offset"
                                    :layout="'prev, pager, next'"
                                    @current-change="handleCurrentChange"
                                    @size-change="handleSizeChange"
                                />
                            </div>

                            <!-- <div class="down" @click="downVisible=true">
                                [ Download <span>xxxx</span> Export ]
                            </div> -->
                        </div>
                    </div>
                </div>
            </div>
            <div v-else>
                <el-alert
                    title=""
                    type="warning"
                    :closable="false"
                    description="Please log in to metamask"
                    show-icon>
                </el-alert>
            </div>
        </div>

        <download v-if="downVisible" :downVisible="downVisible" :titlePage="$t('billing.download_module_title_billing')" @getDownload="getDownload"></download>
        <!-- 回到顶部 -->
        <el-backtop target=".content-box" :bottom="40" :right="20"></el-backtop>
    </div>
</template>

<script>
    // import bus from '@/components/bus';
    import axios from 'axios'
    import QS from 'qs';
    import moment from "moment"
    import NCWeb3 from "@/utils/web3";
    import download from "@/components/download"
    export default {
        name: 'Billing',
        data() {
            return {
                tableData: [],
                searchValue: '',
                selectInput: '1',
                parma: {
                    limit: 10,
                    offset: 1,
                    total: 0,
                    jumperOffset: 1,
                    order_by: '',
                    is_ascend: ''
                },
                loading: false,
                downCsv: localStorage.getItem("addressYM")?localStorage.getItem("addressYM"):'',
                bodyWidth: document.documentElement.clientWidth<1024?true:false,
                width: document.body.clientWidth>600?'400px':'95%',
                center_fail: false,
                network: {
                    name: '',
                    unit: ''
                },
                modelClose: false,
                centerDialogVisible: false,
                downVisible: false,
            };
        },
        components: {
            download
        },
        computed: {
            languageMcs() {
                return this.$store.state.app.languageMcs
            },
            metaAddress() {
                return this.$store.getters.metaAddress
            },
            networkID() {
                return this.$store.getters.networkID
            }
        },
        watch: {
            'metaAddress': function(){
                if(!this.metaAddress) {
                    this.centerDialogVisible = true
                    this.center_fail = false
                    this.modelClose = false
                    return false
                }
                this.getData()
            }
        },
        methods: {
            networkLink(hash, network) {
                if(network && network.toLowerCase() == 'polygon'){
                    window.open('https://mumbai.polygonscan.com/tx/'+hash)
                }
            },
            getDownload(dialog, rows){
                this.downVisible = dialog
            },
            handleCurrentChange(val) {
                this.parma.offset = Number(val);
                this.parma.jumperOffset = String(val)
                this.getData();
            },
            handleSizeChange (val){
                this.parma.limit = Number(val);
                this.parma.offset = 1;
                this.parma.jumperOffset = 1;
                this.getData();
            },
            pageSizeChange(recordPage=parseInt(this.parma.jumperOffset), MaxPagenumber=Math.ceil(this.parma.total/this.parma.limit)) {
                if((recordPage > MaxPagenumber) && (MaxPagenumber > 0)){ 
                    recordPage = MaxPagenumber; 
                }else if(MaxPagenumber<=0){
                    recordPage = 1;  
                }else if(recordPage < 1){ 
                    recordPage = 1;           
                }else if(this.parma.jumperOffset==NaN || this.parma.jumperOffset==""){  
                    recordPage = 1;
                }
                this.parma.offset = Number(recordPage);
                this.parma.jumperOffset = recordPage;
                this.getData(); 
            },
            async sortOrderBy(sort) {
                switch(sort) {
                    case 'pay_tx_hash':
                        return 1;
                    case 'pay_amount':
                        return 2;
                    case 'unlock_amount':
                        return 7;
                    case 'token_name':
                        return 3;
                    case 'file_name':
                        return 4;
                    case 'payload_cid':
                        return 5;
                    case 'address_from':
                        return 6;
                    case 'network_name':
                        return 9;
                    case 'pay_at':
                        return 10;
                    case 'unlock_at':
                        return 8;
                    case 'deadline':
                        return 11;
                    default:
                        return ;
                }
            },
            async sortChange(column) {
                // console.log(column);
                // this.parma.order_by = await this.sortOrderBy(column.prop)
                this.parma.order_by = column.prop
                this.parma.is_ascend = column.order == "ascending" ? 'y' : column.order == "descending" ? 'n' : ''
                this.loading = true
                this.getData()
            },
            getData(){
                let _this = this
                _this.loading = true
                let obj = {
                    "file_name": _this.selectInput === '1'?_this.searchValue:'', 
                    "tx_hash": _this.selectInput === '2'?_this.searchValue:'', 
                    "wallet_address": _this.metaAddress,    
                    "page_number": _this.parma.offset,    
                    "page_size": _this.parma.limit,
                    "order_by": _this.parma.is_ascend?_this.parma.order_by:'',
                    "is_ascend": _this.parma.is_ascend
                }
                let upload_api = `${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/billing?${QS.stringify(obj)}`
                // let upload_api = `./static/response-billing.json?${QS.stringify(obj)}`;

                axios.get(upload_api)
                .then((json) => {
                    if(json.data.status == 'success'){
                        _this.tableData = []
                        _this.tableData = json.data.data.billing;
                        _this.tableData.map(item => {
                            item.txHashVis = false
                            item.payloadVis = false
                            item.walletVis = false
                            item.pay_at =
                                item.pay_at?
                                    String(item.pay_at).length<13?
                                        moment(new Date(parseInt(item.pay_at * 1000))).format("YYYY-MM-DD HH:mm:ss")
                                        :
                                        moment(new Date(parseInt(item.pay_at))).format("YYYY-MM-DD HH:mm:ss")
                                    :
                                    '-'
                            item.unlock_at =
                                item.unlock_at?
                                    String(item.unlock_at).length<13?
                                        moment(new Date(parseInt(item.unlock_at * 1000))).format("YYYY-MM-DD HH:mm:ss")
                                        :
                                        moment(new Date(parseInt(item.unlock_at))).format("YYYY-MM-DD HH:mm:ss")
                                    :
                                    '-'
                            item.deadline =
                                item.deadline?
                                    String(item.deadline).length<13?
                                        moment(new Date(parseInt(item.deadline * 1000))).format("YYYY-MM-DD HH:mm:ss")
                                        :
                                        moment(new Date(parseInt(item.deadline))).format("YYYY-MM-DD HH:mm:ss")
                                    :
                                    '-'
                            // item.locked_fee = web3.utils.fromWei(item.locked_fee, 'ether')
                            // item.locked_fee = 0.000000000000000001 * item.locked_fee
                        })
                        _this.parma.total = Number(json.data.data.total_record_count)
                    }else{
                        _this.$message.error(json.data.message)
                    }
                    _this.loading = false
                }).catch(error => {
                    _this.loading = false
                    console.log(error)
                })
            },
            number(data, n){
                var numbers = '';
                // 保留几位小数后面添加几个0
                for (var i = 0; i < n; i++) {
                    numbers += '0';
                }
                var s = 1 + numbers;
                // 如果是整数需要添加后面的0
                var spot = "." + numbers;
                // Math.round四舍五入  
                //  parseFloat() 函数可解析一个字符串，并返回一个浮点数。
                var value = Math.round(parseFloat(data) * s) / s;
                // 从小数点后面进行分割
                if(!value){
                    return value
                }else if(value.toString().indexOf('.') < 0){
                    return value.toFixed(18)
                }
                var d = value.toString().split(".");
                if (d.length == 1) {
                    value = value.toString() + spot;
                    return value;
                }
                if (d.length > 1) {
                    if (d[1].length < n) {
                        for (var tj = d[1].length; tj < n; tj++) {
                            value = value.toString() + "0";
                        }
                    }
                    return value;
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
            //查询
            search() {
                this.parma.limit = 10
                this.parma.offset = 1
                this.parma.jumperOffset = 1
                this.parma.order_by = ''
                this.parma.is_ascend = ''
                this.getData();
            },
            clearAll() {
                this.searchValue = ""
                this.search();
            },
            searchValueChange() {
                if(this.searchValue == '') this.clearAll()
            }
        },
        mounted() {
            let _this = this
            _this.getData()
            this.$store.dispatch('setRouterMenu', 5)
            this.$store.dispatch('setHeadertitle', this.$t('navbar.BillingHistory'))

            document.onkeydown = function(e) {
                if (e.keyCode === 13) {
                    
                }
            }
        },
        filters: {
            priceFilter(value) {
                let realVal = "";
                if (!isNaN(value) && value !== "") {
                    let tempVal = parseFloat(value).toFixed(19);
                    realVal = tempVal.substring(0, tempVal.length - 1);
                } else {
                    realVal = "-";
                }
                return realVal;
            },
            NumFormat (value) {
                if(!value) return '-';
                return value
            },
            balanceFilter (value) {
                if (String(value) === '0') return 0;
                if (!value) return '-';
                // if (!Number(value)) return 0;
                // if (isNaN(value)) return value;
                // 18 - 单位换算需要 / 1000000000000000000，浮点运算显示有bug
                value = Number(value)
                if(String(value).length > 18){
                    let v1 = String(value).substring(0, String(value).length - 18)
                    let v2 = String(value).substring(String(value).length - 18)
                    let v3 = String(v2).replace(/(0+)\b/gi,"")
                    if(v3){
                        return v1 + '.' + v3
                    }else{
                        return v1
                    }
                    return parseFloat(v1.replace(/(\d)(?=(?:\d{3})+$)/g, "$1,") + '.' + v2)
                }else{
                    let v3 = ''
                    for(let i = 0; i < 18 - String(value).length; i++){
                        v3 += '0'
                    }
                    return '0.' + String(v3 + value).replace(/(0+)\b/gi,"")
                }
            },
        },
    };
</script>


<style scoped lang="scss">
#dealManagement{
    padding: 0.3rem;
    .metatips /deep/{
        position: absolute;
        left: 0;
        top: 0;
        .el-alert__content{
            display: flex;
            align-items: center;
        }
    }
    .upload{
        padding: 0.3rem;
        background-color: #fff;
        border-radius: 0.1rem;
        overflow: hidden;
        #billing{
            padding: 0;
            .backTo{
                display: flex;
                align-items: center;
                font-size: 0.24rem;
                line-height: 1;
                margin: 0 0 0.2rem;
                cursor: pointer;
            }
            .form{
                margin: 0;
                padding: 0;
                background-color: #fff;
                border-radius: 0.1rem;
                .form_top {
                    display: flex;
                    align-items: center;
                    flex-wrap: wrap;

                    .title {
                        width: 100%;
                        text-align: left;
                        font-size: 0.1972rem;
                        color: #000;
                        line-height: 0.42rem;
                        text-indent: 0.08rem;
                    }

                    .search {
                        display: flex;
                        align-items: center;
                        justify-content: flex-start;
                        width: 100%;
                        font-size: 0.18rem;
                        color: #737373;
                        @media screen and (max-width: 600px){
                            font-size: 12px;
                        }
                        .search_right {
                            display: flex;
                            align-items: center;
                            // margin-left: 0.3rem;
                        }

                        span {
                            margin: auto 0.05rem auto 0.35rem;
                            font-size: 0.1372rem;
                            color: #000;
                            white-space: nowrap;
                        }

                        .el-button /deep/ {
                            height: 0.6rem;
                            padding: 0 0.4rem;
                            margin: 0 0.1rem 0 0;
                            color: #fff;
                            line-height: 0.6rem;
                            font-size: 0.2rem;
                            font-family: inherit;
                            border: 0;
                            border-radius: 0.14rem;
                        }
                        .el-input /deep/ {
                            float: left;
                            // width: 35%;
                            margin-right: 0.1rem;
                            padding: 0 0.3rem 0 0;
                            background: #f7f7f7;
                            border-radius: 0.2rem;
                            .el-input__inner {
                                width: 100%;
                                color: #555;
                                font-size: 0.2rem;
                                font-weight: 500;
                                height: 0.6rem;
                                line-height: 0.3rem;
                                padding: 0 0.2rem;
                                background: transparent;
                                border: 0;
                                border-radius: 0.2rem;
                                @media screen and (max-width: 600px){
                                    font-size: 12px;
                                }
                            }

                            .el-input__icon {
                                line-height: 0.3rem;
                            }

                            .el-input-group__prepend {
                                position: relative;
                                width: 130px;
                                padding: 0;
                                background: transparent;
                                border: 0;
                                &::before{
                                    content: '';
                                    position: absolute;
                                    right: 0;
                                    top: 15%;
                                    bottom: 15%;
                                    width: 0.01rem;
                                    background: #DCDFE6;
                                }
                                .el-select{
                                    width: 100%;
                                    margin: 0;
                                    .el-input{
                                        .el-input__inner{
                                            height: 0.32rem;
                                            padding: 0 0.1rem;
                                            line-height: 0.32rem;
                                            border: 0;
                                        }
                                    }
                                }
                            }
                            ::-webkit-input-placeholder{
                                color: #555;
                            }    /* 使用webkit内核的浏览器 */
                            :-moz-placeholder{
                                color: #555;
                            }                  /* Firefox版本4-18 */
                            ::-moz-placeholder{
                                color: #555;
                            }                  /* Firefox版本19+ */
                            :-ms-input-placeholder{
                                color: #555;
                            }           /* IE浏览器 */
                        }

                        // .el-select /deep/ {
                        //     float: right;
                        //     // width: 30%;
                        //     .el-input__inner {
                        //         border-radius: 0.08rem;
                        //         border: 1px solid #f8f8f8;
                        //         color: #737373;
                        //         font-size: 0.12rem;
                        //         height: 0.24rem;
                        //         line-height: 0.24rem;
                        //         padding: 0 0.1rem;
                        //     }

                        //     .el-input__icon {
                        //         line-height: 0.24rem;
                        //     }
                        // }
                    }
                }
                .form_table{
                    margin: 0.24rem 0 0.1rem;
                    .statusStyle{
                        display: inline-block;
                        border: 1px solid;
                        padding: 0 0.1rem;
                        border-radius: 0.05rem;
                        white-space: nowrap;
                    }
                    .el-table /deep/{
                        overflow: visible;
                        // .el-table__body-wrapper,.el-table__header-wrapper{
                        //     overflow: visible;
                        // }
                        tr{
                            border-radius: 0.08rem;
                            // cursor: pointer;
                            th{
                                height: 0.7rem;
                                padding: 0;
                                background-color: #e5eeff !important;
                                text-align: center;
                                .cell{
                                    display: flex;
                                    align-items: center;
                                    justify-content: center;
                                    // padding-right: 0;
                                    word-break: break-word;
                                    font-size: 0.2rem;
                                    font-weight: 500;
                                    color: #555;
                                    text-transform: capitalize;
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
                                padding: 0.15rem 0.05rem;
                                border-bottom: 1px solid #dfdfdf;
                                .cell{
                                    padding: 0 0.05rem;
                                    font-size: 0.18rem;
                                    word-break: break-word;
                                    color: #000;
                                    text-align: center;
                                    line-height: 0.25rem;
                                    overflow: visible;
                                    .el-rate__icon{
                                        font-size: 0.16rem;
                                        margin-right: 0;
                                    }
                                    .hot-cold-box{
                                        .el-button{
                                            width: 100%;
                                            border: 0;
                                            padding: 0;
                                            margin: 0;
                                            background-color: transparent;
                                            font-size: 0.18rem;
                                            font-family: inherit;
                                            word-break: break-word;
                                            color: #000;
                                            text-align: center;
                                            line-height: 0.25rem;
                                            overflow: hidden;
                                            text-overflow: ellipsis;
                                            white-space: normal;
                                            display: -webkit-box;
                                            -webkit-line-clamp: 2;
                                            -webkit-box-orient: vertical;
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
                                            color: #4e86ff;
                                            text-decoration: underline;
                                        }
                                        .el-button:hover{
                                            img{
                                                display: inline-block;
                                            }
                                        }
                                    }
                                }
                            }
                        }
                    }
                    .actionStyle{
                        position: relative;
                        display: flex;
                        align-items: center;
                        justify-content: center;
                        cursor: pointer;
                        .actionBtn{
                            display: flex;
                            align-items: center;
                            justify-content: center;
                            width: 0.4rem;
                            padding: 0.1rem 0;
                            margin: auto;
                            border-radius: 0.08rem;
                            background-color: #d5d7de;
                            border: 1px solid #f7f7f7;
                            span{
                                display: block;
                                // width: 9px;
                                // height: 9px;
                                // margin: 0 0.04rem;
                                width: 3px;
                                height: 3px;
                                margin: 0 0.01rem;
                                background-color: #767c92;
                                border-radius: 100%;
                            }
                            // span:nth-child(2){
                            //     background-color: #fd397a;
                            // }
                            // span:nth-child(3){
                            //     background-color: #1dc9b7;
                            // }
                            // span:nth-child(4){
                            //     background-color: #404040;
                            // }
                        }
                        .actStatus{
                            // display: none;
                            position: absolute;
                            left: -10px;
                            right: 0px;
                            top: 100%;
                            background: #fff;
                            border: 1px solid #f7f7f7;
                            border-radius: 0.08rem;
                            z-index: 1002;
                            width: 120px;
                            p{
                                padding: 0.1rem 0;
                                border-top: 1px solid #f7f7f7;
                                font-size: 0.1372rem;
                                color: #ffb822;
                                line-height: 1;
                            }
                            p:nth-child(1){
                                border: 0;
                            }
                            p:nth-child(2){
                                color: #fd397a;
                            }
                            p:nth-child(3){
                                color: #1dc9b7;
                            }
                            // p:nth-child(4){
                            //     color: #404040;
                            // }
                            p:hover{
                                font-size: 0.14rem;
                            }
                        }
                    }
                    .actionStatus{
                        padding: 0;
                        background-color: #1dc9b7;
                        font-size: 0.135rem;
                        color: #fff;
                        border-radius: 0.08rem;
                        line-height: 2.2;
                        white-space: nowrap
                    }
                }
                .form_pagination {
                    position: relative;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    text-align: center;
                    margin: 0.5rem 0 0.06rem;
                    padding: 0;
                    @media screen and (max-width: 1024px){
                        padding: 0;
                    }
                    @media screen and (max-width: 600px){
                        flex-wrap: wrap;
                    }
                    .pagination {
                        display: flex;
                        align-items: center;
                        font-size: 0.18rem;
                        color: #000;
                        .el-pagination /deep/{
                        .el-icon{
                            font-size: 0.18rem;
                        }
                        .el-pager{
                            li{
                            min-width: 30px;
                            height: 30px;
                            font-size: 0.18rem;
                            font-weight: normal;
                            line-height: 30px;
                            }
                            .active{
                            border: 1px solid #6798f5;
                            border-radius: 5px;
                            color: #000;
                            }
                        }
                        }
                        .el-select /deep/{
                        max-width: 100px;
                        margin-right: 0.15rem;
                        .el-input__inner, .el-input__icon{
                            height: 30px;
                            line-height: 30px;
                        }
                        }
                        .span{
                        margin: 0 0 0 10px;
                        font-size: 13px;
                        font-weight: 400;
                        color: #606266;
                        white-space: nowrap;
                        }
                        .paginaInput /deep/{
                        max-width: 50px;
                        .el-input__inner, .el-input__icon{
                            padding: 0 3px;
                            height: 30px;
                            line-height: 30px;
                            text-align: center;
                        }
                        }

                        .pagination_left {
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
                    .down{
                        position: absolute;
                        right: 0;
                        font-size: 0.2rem;
                        color: #888;
                        cursor: pointer;
                        @media screen and (max-width: 600px){
                            position: relative;
                            width: 100%;
                            text-align: center;
                        }
                        span{
                            color: #2C7FF8;
                        }
                        &:hover{
                            text-decoration: underline;
                        }
                    }
                }
            }
        }
    }
}
    @media screen and (max-width:1250px){
        .price{
            width: 100%;
            span{
                display: block;
            }
        }
    }
    @media screen and (max-width:1024px){
    #dealManagement{
        .upload{
            #billing{
                .form{
                    .form_table{
                        overflow-x: scroll;
                        .el-table /deep/ {
                            width: 1024px !important;
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
            .upload{
                padding: 0.1rem;
                #billing{
                    padding: 0.15rem 0 0.3rem;
                    .upload{
                        padding: 0.1rem;
                        .title{
                            word-break: break-word;
                        }
                        .upload_form{
                            flex-wrap: wrap;
                            .el-form /deep/{
                                width: 95%;
                            }
                        }
                        .upload_form_detail{
                            flex-wrap: wrap;
                            .el-form /deep/{
                                width: calc(100% - 0.3rem);
                                margin: auto 0.15rem;
                            }
                            .el-tabs /deep/{
                                .el-tabs__nav{
                                    .el-tabs__item{
                                        width: 1.5rem;
                                        margin-right: 0.1rem;
                                    }
                                    #tab-bids::after{
                                        display: none;
                                    }
                                }
                                .search{
                                    .el-select{
                                        width: 48%;
                                        margin: 0 1% 0.05rem;
                                    }
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
            #billing{
                .form{
                    .form_top {
                        .search {
                            flex-wrap: wrap;
                            span {
                                margin-left: 0;
                            }
                            .el-input{
                                margin: 0.1rem 0;
                            }
                        }
                    }
                }
                //     .form{
                //         .form_table{
                //             .el-table /deep/ {
                //                 overflow: hidden;
                //                 overflow-y: visible;
                //                 .el-table__body-wrapper,.el-table__header-wrapper{
                //                     overflow: hidden;
                //                     overflow-y: visible;
                //                     overflow-x: auto;
                //                 }
                //             }
                //         }
                //     }
            }
        }
    }
    }
    @media screen and (max-width:441px){
    #dealManagement{
        .upload{
            #billing{
                .upload{
                    .upload_form_detail{
                        .minerStyle{
                            flex-wrap: wrap;
                            .Leave_tip{
                                width: 100%;
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
            #billing .upload .title{
                font-size: 0.2rem;
                .el-button /deep/{
                    height: 0.35rem;
                    line-height: 0.35rem;
                    font-size: 0.16rem;
                }
            }
            #billing .upload .upload_form_detail .el-tabs /deep/ {
                .el-tabs__nav{
                    .el-tabs__item{
                        width: 1.2rem;
                        font-size: 14px;
                    }
                }
            }
            #billing .form .form_top .title p{
                float: none;
                clear: both;
                font-size: 0.16rem;
                line-height: 2;
                span{
                    font-size: 0.16rem;
                }
            }
        }
    }
    }
</style>
