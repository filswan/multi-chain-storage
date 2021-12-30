<template>
    <div id="billing">
        <div class="form">
            <div class="form_top">
                <div class="search">
                    <el-input
                        placeholder="Search by TX HASH"
                        prefix-icon="el-icon-search"
                        v-model="searchValue"
                    >
                    </el-input>
                    <div class="search_right" :style="{'opacity': !searchValue?'0.8':'1'}">
                        <el-button @click="search" style="background-color: #ffb822"
                        :disabled="!searchValue">Search</el-button>
                        <el-button
                        type="primary"
                        style="background-color: #0b318f"
                        @click="clearAll"
                        :disabled="!searchValue"
                        >
                        Clear
                        </el-button>
                    </div>
                </div>
            </div>
            <div class="form_table">
                <el-table v-loading="loading" :data="tableData" style="width: 100%" :empty-text="$t('deal.formNotData')" header-click="aaa">
                    <el-table-column prop="tx_hash" :label="$t('billing.TRANSACTION')" min-width="140">
                        <template slot-scope="scope">
                            <div class="hot-cold-box">
                                <el-popover
                                    placement="top"
                                    trigger="hover" width="300"
                                    v-model="scope.row.txHashVis">
                                    <div class="upload_form_right">
                                        <p>{{scope.row.tx_hash}}</p>
                                    </div>
                                    <el-button slot="reference" @click="copyTextToClipboard(scope.row.tx_hash)">
                                        <img src="@/assets/images/copy.png" alt="">
                                        {{scope.row.tx_hash}}
                                    </el-button>
                                </el-popover>
                            </div>
                        </template>                    
                    </el-table-column>
                    <el-table-column prop="locked_fee" :label="$t('billing.AMOUNT')" min-width="150">
                        <template slot-scope="scope">{{scope.row.locked_fee | balanceFilter}}</template>
                    </el-table-column>
                    <el-table-column prop="coin_type" label="COIN" min-width="120"></el-table-column>
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
                                    <el-button slot="reference" @click="copyTextToClipboard(scope.row.payload_cid)">
                                        <img src="@/assets/images/copy.png" alt="">
                                        {{scope.row.payload_cid}}
                                    </el-button>
                                </el-popover>
                            </div>
                        </template>                    
                    </el-table-column>
                    <el-table-column prop="address_from" :label="$t('billing.WALLET')" min-width="140">
                        <template slot-scope="scope">
                            <div class="hot-cold-box">
                                <el-popover
                                    placement="top"
                                    trigger="hover" width="300"
                                    v-model="scope.row.walletVis">
                                    <div class="upload_form_right">
                                        <p>{{scope.row.address_from}}</p>
                                    </div>
                                    <el-button slot="reference" @click="copyTextToClipboard(scope.row.address_from)">
                                        <img src="@/assets/images/copy.png" alt="">
                                        {{scope.row.address_from}}
                                    </el-button>
                                </el-popover>
                            </div>
                        </template>                    
                    </el-table-column>
                    <el-table-column prop="network" :label="$t('billing.NETWORK')" min-width="120"></el-table-column>
                    <el-table-column prop="lock_payment_time" :label="$t('billing.PAYMENTDATE')" min-width="140"></el-table-column>
                    <el-table-column prop="deadline" :label="$t('billing.Deadline')" min-width="140"></el-table-column>
                </el-table>
            </div>



            <div class="form_pagination">
                <div class="pagination">
                    <el-pagination
                        :total="parma.total"
                        :page-size="parma.limit"
                        :current-page="parma.offset"
                        :pager-count="bodyWidth?5:7"
                        style="padding: 30px 40px 32px 20px;"
                        background
                        :layout="bodyWidth?'prev, pager, next':'total, prev, pager, next, jumper'"
                        @current-change="handleCurrentChange" />
                </div>
            </div>
        </div>

        <!-- 回到顶部 -->
        <el-backtop target=".content-box" :bottom="40" :right="20"></el-backtop>
    </div>
</template>

<script>
    // import bus from '@/components/bus';
    import axios from 'axios'
    import QS from 'qs';
    import moment from "moment"

    export default {
        name: 'SearchFile',
        data() {
            return {
                tableData: [],
                searchValue: '',
                parma: {
                    limit: 10,
                    offset: 1,
                    total: 0
                },
                loading: false,
                downCsv: localStorage.getItem("addressYM")?localStorage.getItem("addressYM"):'',
                bodyWidth: document.documentElement.clientWidth<1024?true:false,
            };
        },
        computed: {
            language() {
                return this.$store.state.app.language
            },
            metaAddress() {
                return this.$store.getters.metaAddress
            }
        },
        watch: {
            'metaAddress': function(){
                this.getData()
            }
        },
        components: {},
        methods: {
            handleCurrentChange(val) {
                this.parma.offset = val
                this.getData()
            },
            getData(){
                let _this = this
                _this.loading = true
                let obj = {
                    "tx_hash": _this.searchValue, 
                    "wallet_address": _this.metaAddress,    
                    "page_number": _this.parma.offset,    
                    "page_size": _this.parma.limit
                }
                let upload_api = `${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/billing?${QS.stringify(obj)}`
                // let upload_api = `./static/response-billing.json?${QS.stringify(obj)}`;

                axios.get(upload_api)
                .then((json) => {
                    if(json.data.status == 'success'){
                        _this.tableData = []
                        _this.tableData = json.data.data;
                        _this.tableData.map(item => {
                            item.txHashVis = false
                            item.payloadVis = false
                            item.walletVis = false
                            item.lock_payment_time =
                                item.lock_payment_time?
                                    String(item.lock_payment_time).length<13?
                                        moment(new Date(parseInt(item.lock_payment_time * 1000))).format("YYYY-MM-DD HH:mm:ss")
                                        :
                                        moment(new Date(parseInt(item.lock_payment_time))).format("YYYY-MM-DD HH:mm:ss")
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
                        _this.parma.total = Number(json.data.page_info.total_record_count)
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
                let saveLang = localStorage.getItem('lang') == 'cn'?"复制成功":"success";
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
                this.getData();
            },
            clearAll() {
                this.searchValue = ""
                this.search();
            },
        },
        mounted() {
            let _this = this
            _this.getData()

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
                if (!Number(value)) return 0;
                if (isNaN(value)) return value;
                // 18 - 单位换算需要 / 1000000000000000000，浮点运算显示有bug
                if(String(value).length > 18){
                    let v1 = String(value).substring(0, String(value).length - 18)
                    let v2 = String(value).substring(String(value).length - 18)
                    return v1.replace(/(\d)(?=(?:\d{3})+$)/g, "$1,") + '.' + v2
                }else if(String(value).length == 18){
                    return String(value).replace(/(\d)(?=(?:\d{3})+$)/g, "$1,")
                }else{
                    let v3 = ''
                    for(let i = 0; i < 18 - String(value).length; i++){
                        v3 += '0'
                    }
                    return '0.' + v3 + value
                }
            },
        },
    };
</script>


<style scoped lang="scss">
    #billing{
        padding: 0 0 0.25rem;
        .backTo{
            display: flex;
            align-items: center;
            font-size: 0.24rem;
            line-height: 1;
            margin: 0 0 0.2rem;
            cursor: pointer;
        }
        .form{
            margin: 0 0 0.2rem;
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
                    height: 0.42rem;

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
                        height: 0.34rem;
                        padding: 0 0.4rem;
                        margin: 0 0.1rem;
                        color: #fff;
                        line-height: 0.34rem;
                        font-size: 0.1372rem;
                        border: 0;
                        border-radius: 0.08rem;
                    }

                    .el-input /deep/ {
                        float: left;
                        width: 35%;

                        .el-input__inner {
                            width: 100%;
                            color: #737373;
                            font-size: 0.12rem;
                            height: 0.34rem;
                            line-height: 0.34rem;
                            padding: 0 0.27rem;
                        }

                        .el-input__icon {
                            line-height: 0.24rem;
                        }
                    }

                    .el-select /deep/ {
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

                    .el-input__icon {
                        line-height: 0.24rem;
                    }
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
                            height: 0.5rem;
                            padding: 0;
                            background-color: #f2f2f2 !important;
                            text-align: center;
                            .cell{
                                word-break: break-word;
                                font-weight: 500;
                                color: #737373;
                                text-transform: uppercase;
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
                                padding: 0 0.05rem;
                                font-size: 0.1372rem;
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
                                        font-size: 0.1372rem;
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
        #billing{
            .form{
                .form_top {
                    .search {
                    flex-wrap: wrap;
                    height: auto;

                    .el-input /deep/ {
                        width: 100%;
                        margin: 0.1rem 0;

                        .el-input__inner {
                        width: 100%;
                        font-size: 0.1372rem;
                        }
                    }

                    span {
                        margin-left: 0;
                    }

                    .search_right {

                        .el-select /deep/ {
                        .el-input__inner {
                            font-size: 0.1372rem;
                        }
                        }

                        .el-button /deep/ {
                        padding: 0 0.2rem;
                        font-size: 0.1372rem;
                        }
                    }
                    }
                }
                .form_table{
                    overflow-x: scroll;
                    .el-table /deep/ {
                        width: 1024px !important;
                    }
                }
            }
        }
    }
    @media screen and (max-width:999px){
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
    @media screen and (max-width:640px){
        #billing{
            .form{
                .form_top{
                    .search{
                        .search_right{
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
    @media screen and (max-width:441px){
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
    @media screen and (max-width:321px){
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
</style>
