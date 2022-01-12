<template>
  <div id="dealManagement">
    <div class="tabTaskStyle">
      <div class="createTask">
        <!-- name: 'upload_file' -->
        <router-link :to="{name: 'upload_file'}">
            {{$t('uploadFile.Upload_More_Files')}}
        </router-link>
      </div>
    </div>
    <div class="form">
      <div class="form_top">
        <div class="upload_title">{{$t('uploadFile.topTip')}}</div>
        <div class="search_file">
          <p>{{$t('uploadFile.search_title')}}</p>

          <div class="search_right">
            <el-input
              :placeholder="$t('uploadFile.search_input')"
              prefix-icon="el-icon-search"
              v-model="searchValue"
            >
            </el-input>
            <el-button type="primary" style="background-color: #0b318f" @click="clearAll">
                        {{ $t("deal.Clear_all") }}
            </el-button>
          </div>
        </div>
      </div>

      <div class="form_table">
        <!-- @row-click="tableTrClick" highlight-current-row  -->
        <el-table
          :data="tableData" ref="singleTable"  stripe
          style="width: 100%"
          :empty-text="$t('deal.formNotData')"
           v-loading="loading"
        >
          <el-table-column prop="file_name" :label="$t('uploadFile.file_name')" min-width="120">
            <template slot-scope="scope">
              <div class="hot-cold-box" @click="toDetail(scope.row)" style="text-decoration: underline;">
                {{ scope.row.file_name }}
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="file_size" :label="$t('uploadFile.file_size')" width="90">
            <template slot-scope="scope">
              <div class="hot-cold-box">
                {{ scope.row.file_size | formatbytes }}
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="status" :label="$t('uploadFile.file_status')" width="140">
            <template slot-scope="scope">
              <el-button type="danger" class="statusStyle" v-if="scope.row.status&&scope.row.status.toLowerCase()=='failed'">
                  {{ languageMcp == "en" ? "Fail" : '失败'}}
              </el-button>
              <el-button type="warning" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='pending'">
                  {{ languageMcp == "en" ? "Pending" : '待支付'}}
              </el-button>
              <el-button type="primary" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='processing'">
                  {{ languageMcp == "en" ? "Processing" : '处理中'}}
              </el-button>
              <el-button type="success" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='active'">
                  {{ languageMcp == "en" ? "Active" : '完成'}}
              </el-button>
              <el-button type="info" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='refunded'">
                  {{ languageMcp == "en" ? "Refunded" : '已退款'}}
              </el-button>
              <el-button type="info" plain class="statusStyle" v-else>
                  {{scope.row.status}}
              </el-button>
            </template>
          </el-table-column>
          <el-table-column prop="pin_status" width="140">
            <template slot="header" slot-scope="scope">
              <div class="tips">
                {{$t('uploadFile.status')}}
                    
                <el-tooltip effect="dark" :content="$t('uploadFile.status_tooltip')" placement="top">
                    <img src="@/assets/images/info.png"/>
                </el-tooltip>
              </div>
            </template>
            <template slot-scope="scope">
              <div class="statusStyle" style="color: #6c757d" v-if="scope.row.pin_status&&scope.row.pin_status.toLowerCase()=='unpinned'">
                  {{scope.row.pin_status}}
              </div>
              <div class="statusStyle" style="color: #ff9900" v-if="scope.row.pin_status&&scope.row.pin_status.toLowerCase()=='pinned'">
                  {{scope.row.pin_status}}
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="payload_cid" min-width="120">
            <template slot="header" slot-scope="scope">
              <div class="tips">
                {{$t('billing.PAYLOADCID')}}
                    
                <el-tooltip effect="dark" :content="$t('uploadFile.data_tooltip')" placement="top">
                    <img src="@/assets/images/info.png"/>
                </el-tooltip>
              </div>
            </template>
            <template slot-scope="scope">
              <div class="hot-cold-box">
                <el-popover
                    v-if="scope.row.payload_cid"
                    placement="top"
                    trigger="hover" width="300"
                    v-model="scope.row.payloadAct">
                    <div class="upload_form_right">
                        <p>{{scope.row.payload_cid}}</p>
                    </div>
                    <el-button slot="reference" @click="copyTextToClipboard(scope.row.payload_cid)">
                        <img src="@/assets/images/copy.png" alt="">
                        {{scope.row.payload_cid}}
                    </el-button>
                </el-popover>
                <span v-else>-</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="miner_fid" width="120">
            <template slot="header" slot-scope="scope">
              <div class="tips">
                {{$t('uploadFile.w3ss_id')}}
                    
                <el-tooltip effect="dark" :content="$t('uploadFile.w3ss_id_tooltip')" placement="top">
                    <img src="@/assets/images/info.png"/>
                </el-tooltip>
              </div>
            </template>
            <template slot-scope="scope">
              <div class="hot-cold-box">
                <div class="elTips">
                  <el-popover
                      placement="top"
                      trigger="hover" popper-class="elPopMiner"
                      v-model="scope.row.payloadAct">
                      <div class="upload_form_right">
                          <div
                            class="statusStyle"
                            v-if="scope.row.deal_status == 'Created'"
                            :style="$status_color.Task_color('Created')">
                            {{ languageMcp == "en" ? "Created" : "已创建" }}
                          </div>
                          <div
                            class="statusStyle"
                            v-else-if="scope.row.deal_status == 'Assigned'"
                            :style="$status_color.Task_color('Assigned')"
                          >
                            {{ languageMcp == "en" ? "Assigned" : "已分配" }}
                          </div>
                          <div
                            class="statusStyle"
                            v-else-if="scope.row.deal_status == 'Accepted'"
                            :style="$status_color.Task_color('Accepted')"
                          >
                            {{ languageMcp == "en" ? "Accepted" : "已接受" }}
                          </div>
                          <div
                            class="statusStyle"
                            v-else-if="scope.row.deal_status == 'Completed'"
                            :style="$status_color.Task_color('Completed')"
                          >
                            {{ languageMcp == "en" ? "Completed" : "已完成" }}
                          </div>
                          <div
                            class="statusStyle"
                            v-else-if="scope.row.deal_status == 'Failed'"
                            :style="$status_color.Task_color('Failed')"
                          >
                            {{ languageMcp == "en" ? "Failed" : "已失败" }}
                          </div>
                          <div
                            class="statusStyle"
                            v-else-if="scope.row.deal_status == 'Cancelled'"
                            :style="$status_color.Task_color('Cancelled')"
                          >
                            {{ languageMcp == "en" ? "Cancelled" : "已取消" }}
                          </div>
                          <div
                            class="statusStyle"
                            v-else-if="scope.row.deal_status == 'Closed'"
                            :style="$status_color.Task_color('Closed')"
                          >
                            {{ languageMcp == "en" ? "Closed" : "已关闭" }}
                          </div>
                          <div
                            class="statusStyle"
                            v-else-if="scope.row.deal_status == 'Expired'"
                            :style="$status_color.Task_color('Expired')"
                          >
                            {{ languageMcp == "en" ? "Expired" : "已过期" }}
                          </div>
                          <div
                              class="statusStyle"
                              v-else-if="scope.row.deal_status == 'ActionRequired'"
                              :style="$status_color.Task_color('ActionRequired')">
                              {{ languageMcp == 'en' ? 'ActionRequired' : '需要操作' }}
                          </div>
                          <div
                              class="statusStyle"
                              v-else-if="scope.row.deal_status == 'DealSent'"
                              :style="$status_color.Task_color('DealSent')">
                              {{ languageMcp == 'en' ? 'DealSent' : '交易已发送' }}
                          </div>
                          <div class="statusStyle"
                                v-else-if="scope.row.deal_status == 'FileImporting'"
                                :style="$status_color.Task_color('FileImporting')">
                              {{ languageMcp == 'en' ? 'FileImporting' : '文件导入中' }}
                          </div>
                          <div class="statusStyle"
                                v-else-if="scope.row.deal_status == 'FileImported'"
                                :style="$status_color.Task_color('FileImported')">
                              {{ languageMcp == 'en' ? 'FileImported' : '文件已导入' }}
                          </div>
                          <div class="statusStyle"
                                v-else-if="scope.row.deal_status == 'ImportFailed'"
                                :style="$status_color.Task_color('ImportFailed')">
                              {{ languageMcp == 'en' ? 'ImportFailed' : '导入失败' }}
                          </div>
                          <div class="statusStyle"
                                v-else-if="scope.row.deal_status == 'Downloading'"
                                :style="$status_color.Task_color('Downloading')">
                              {{ languageMcp == 'en' ? 'Downloading' : '下载中' }}
                          </div>
                          <div class="statusStyle"
                                v-else-if="scope.row.deal_status == 'DownloadFailed'"
                                :style="$status_color.Task_color('DownloadFailed')">
                              {{ languageMcp == 'en' ? 'DownloadFailed' : '下载失败' }}
                          </div>
                          <div class="statusStyle"
                                v-else-if="scope.row.deal_status == 'DealActive'"
                                :style="$status_color.Task_color('DealActive')">
                              {{ languageMcp == 'en' ? 'DealActive' : '有效交易' }}
                          </div>
                          <div class="statusStyle"
                                v-else-if="scope.row.deal_status == 'Waiting'"
                                :style="$status_color.Task_color('Waiting')">
                              {{ languageMcp == 'en' ? 'Waiting' : '等待中' }}
                          </div>
                          <div class="statusStyle"
                                v-else-if="scope.row.deal_status == 'ReadyForImport'"
                                :style="$status_color.Task_color('ReadyForImport')">
                              {{ languageMcp == 'en' ? 'ReadyForImport' : '准备导入' }}
                          </div>
                          <div
                              class="statusStyle"
                              v-else-if="scope.row.deal_status == ''">
                              -
                          </div>
                          <div
                              class="statusStyle"
                              v-else>
                              {{ scope.row.deal_status }}
                          </div>
                      </div>
                      <el-button slot="reference" v-if="scope.row.miner_fid" @click="minerIdLink(scope.row.miner_fid)">
                          {{scope.row.miner_fid}}
                      </el-button>
                      <div v-else slot="reference">
                        {{$t('uploadFile.w3ss_id_nothing')}}
                      </div>
                  </el-popover>
                        
                  <el-tooltip v-if="!scope.row.miner_fid" effect="dark" :content="$t('uploadFile.w3ss_id_nothing_tooltip')" placement="top">
                      <img src="@/assets/images/info.png"/>
                  </el-tooltip>
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="create_at" :label="$t('uploadFile.create_time')" min-width="90">
            <template slot-scope="scope">
              {{ scope.row.create_at }}
            </template>
          </el-table-column>
          <el-table-column prop="active" width="120" :label="$t('uploadFile.payment')">
            <template slot-scope="scope">
              <div class="hot-cold-box">
                <el-button class="uploadBtn blue" type="primary"
                  v-if="tableData[scope.$index].status.toLowerCase()=='pending'"
                  @click.stop="payClick(scope.row)">
                  {{$t('uploadFile.pay')}}
                </el-button>
                <el-button class="uploadBtn blue" type="primary"
                  v-else-if="tableData[scope.$index].status.toLowerCase()=='refunded'"
                  @click.stop="refundClick(scope.row)">
                  {{$t('uploadFile.refund')}}
                </el-button>
                <el-button 
                  v-else-if="tableData[scope.$index].status.toLowerCase()=='failed'"
                  :disabled="true"
                  class="uploadBtn grey opacity">{{$t('uploadFile.failed')}}</el-button>
                <el-button 
                  v-else
                  :disabled="true"
                  class="uploadBtn grey opacity">{{$t('uploadFile.paid')}}</el-button>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="MINT" width="100" :label="$t('uploadFile.MINT')">
            <template slot-scope="scope">
              <div class="hot-cold-box">
                <el-button class="uploadBtn blue" type="primary"
                  v-if="tableData[scope.$index].status.toLowerCase()!='pending'"
                  @click.stop="mintFunction(scope.row)">
                  {{$t('uploadFile.MINT')}}
                </el-button>
                <el-button 
                  v-else
                  :disabled="true"
                  class="uploadBtn grey opacity">{{$t('uploadFile.MINT')}}</el-button>
              </div>
            </template>
          </el-table-column>
        </el-table>

        <div class="form_pagination">
          <div class="pagination">
            <el-pagination
              :total="parma.total"
              :page-size="parma.limit"
              :current-page="parma.offset"
              :pager-count="bodyWidth ? 5 : 7"
              background
              :layout="
                bodyWidth
                  ? 'prev, pager, next'
                  : 'total, prev, pager, next, jumper'
              "
              @current-change="handleCurrentChange"
            />
          </div>
        </div>
      </div>

      <div class="loadMetamaskPay" v-if="loadMetamaskPay">
          <div>
              <div class="el-loading-spinner"><svg viewBox="25 25 50 50" class="circular"><circle cx="50" cy="50" r="20" fill="none" class="path"></circle></svg><!----></div>
              <p>{{$t('uploadFile.payment_tip')}}</p>
          </div>
      </div>
    </div>
 <!-- @getPay="getPay" -->
    <pay-tip v-if="payVisible" :payVisible="payVisible" 
        :payRow="payRow" :cost="cost" :bilingPrice="biling_price"
        @getDialog="getDialog"></pay-tip>

        <el-dialog title="" :visible.sync="finishTransaction" :width="width"
            :before-close="finishClose"
            custom-class="completeDia">
            <img src="@/assets/images/alert-icon.png" />
            <h1>{{$t('uploadFile.COMPLETED')}}!</h1>
            <h3>{{$t('uploadFile.SUCCESS')}}</h3>
            <a :href="'https://mumbai.polygonscan.com/tx/'+txHash" target="_blank">{{txHash}}</a>
            <a class="a-close" @click="finishClose">{{$t('uploadFile.CLOSE')}}</a>
        </el-dialog>

        <el-dialog title="" :visible.sync="failTransaction" :width="width"
            custom-class="completeDia">
            <img src="@/assets/images/error.png" />
            <h1>{{$t('uploadFile.Fail')}}!</h1>
            <h3>{{$t('uploadFile.FailTIP')}}</h3>
            <a :href="'https://mumbai.polygonscan.com/tx/'+txHash" target="_blank">{{txHash}}</a>
            <a class="a-close" @click="failTransaction=false">{{$t('uploadFile.CLOSE')}}</a>
        </el-dialog>

        <el-dialog
          title="Tips"
          :visible.sync="wrongVisible" :show-close="false"
          :width="width" custom-class="wrongNet">
          <span>{{$t('uploadFile.until')}}</span>
        </el-dialog>

        <el-dialog title="" :visible.sync="metamaskLoginTip" :width="width"
            custom-class="completeDia">
            <img src="@/assets/images/box-important.png" />
            <h4>{{$t('fs3Login.toptip_01')}} {{network.name}} {{$t('fs3Login.toptip_02')}} <b>{{$t('fs3Login.toptip_Network')}}</b>.</h4>
            <a class="a-close" @click="metamaskLoginTip=false">{{$t('uploadFile.OK')}}</a>
        </el-dialog>
        
        <el-dialog title="" :visible.sync="mintTransaction" :width="width"
            :before-close="finishClose"
            custom-class="completeDia">
            <img src="@/assets/images/alert-icon.png" />
            <h1>View Your NFT</h1>
            <h3>Your NFT has been minted! You can view the transaction here:</h3>
            <a :href="'https://mumbai.polygonscan.com/tx/'+txHash" target="_blank">{{txHash}}</a>
            <br />
            <a :href="'https://testnets.opensea.io/assets/mumbai/'+mintContractAddress+'/'+tokenId" target="_blank">Note: The NFT will take some time to load on Opensea.</a>
            <a class="a-close" @click="finishClose">{{$t('uploadFile.CLOSE')}}</a>
        </el-dialog>
    <mint-tip v-if="mineVisible" :mineVisible="mineVisible" :cid="mintCID" :dealID="dealID" :fileSize="fileSize"
                @getMintDialog="getMintDialog"></mint-tip>
      
    <!-- 回到顶部 -->
    <el-backtop target=".content-box" :bottom="40" :right="20"></el-backtop>
  </div>
</template>

<script>
import bus from "@/components/bus";
import axios from 'axios'
import QS from 'qs';
import * as myAjax from "@/api/uploadFile";
import moment from "moment";
import payTip from "@/components/payTip"
import mintTip from "@/components/mintTip"
import NCWeb3 from "@/utils/web3";
import first_contract_json from "@/utils/swanPayment.json";
import erc20_contract_json from "@/utils/ERC20.json";
let contract_erc20
let that

export default {
  name: "uploadFiles",
  data() {
    return {
      tableData: [],
      tableDataChild: [],
      tableDataAll: [],
      parma: {
        limit: 10,
        offset: 1,
        locationValue: "",
        total: 0,
      },
      parmaChild: {
        limit: 10,
        offset: 1,
        locationValue: "",
        total: 0,
      },
      searchValue: "",
      loading: false,
      loadingChild: false,
      bodyWidth: document.documentElement.clientWidth < 1024 ? true : false,
      expands: [],
      getRowKeys: (row) => {
        return row.uuid
      },
      exChangeList: [],
      payVisible: false,
      mineVisible: false,
      mintCID: '',
      dealID: '',
      fileSize: '',
      mintContractAddress: process.env.NEXT_PUBLIC_MINT_CONTRACT,
      tokenId: '',
      toAddress: '',
      storage: 0,
      centerDialogVisible: false,
      center_fail: false,
      width: document.body.clientWidth>600?'400px':'95%',
      payment: {
        cid: '',
        amount: '',
        biling_price: 0
      },
      wrongVisible: false,
      paySuccess: false,
      timer: '',
      timeIndex: 0,
      modelClose: false,
      network: {
          name: '',
          unit: '',
          text: false
      },
      gatewayContractAddress: this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS,
      recipientAddress: this.$root.RECIPIENT,
      usdcAddress: this.$root.USDC_ADDRESS,
      txHash: '',
      finishTransaction: false,
      mintTransaction: false,
      failTransaction: false,
      loadMetamaskPay: false,
      payRow: {},
      storage: 0,
      biling_price: 0,
      cost: {
        storage_cost_low: 0,
        storage_cost_average: 0,
        storage_cost_high: 0,
      },
      metamaskLoginTip: false
    };
  },
  computed: {
    languageMcp() {
      return this.$store.state.app.languageMcp;
    },
    metaAddress() {
        return this.$store.getters.metaAddress
    },
    networkID() {
        return this.$store.getters.networkID
    }
  },
  components: {
      payTip,
      mintTip
  },
  watch: {
    'searchValue': function(){
      let _this = this
      _this.parma.limit = 10
      _this.parma.offset = 1
      _this.parmaChild.limit = 10
      _this.parmaChild.offset = 1
      _this.getData()
    }
  },
  created() {
    if (this.$route.query.status !== undefined) {
      this.parma.limit = this.$route.query.limit;
      this.parma.searchValue = this.$route.query.task_name;
    }
  },
  methods: {
    minerIdLink(id){
      window.open(`https://calibration.filscout.com/en/miner/${id}`)
    },
    toDetail(row){
      this.$router.push({name: 'my_files_filename', params: {file_name: row.file_name}})
      localStorage.setItem('files_deal_id', row.deal_id)
      localStorage.setItem('files_payload_cid', row.payload_cid)
      // this.$router.push({name: 'my_files_detail', params: {id: id, cid: cid}})
    },
    clickRowHandle(row, column, event) {
      if (this.expands.includes(row.uuid)) {
        // this.expands = this.expands.filter(val => val !== row.uuid);
        this.expands = []
      } else {
        this.expands = []
        if (row) {
          this.expands.push(row.uuid)
        }
        this.tableTrClick(row)
      }
    },
    payClick(row){
      let _this = this
      if(_this.metaAddress&&_this.networkID!=80001) {
          _this.metamaskLoginTip = true
          return false
      }
      // console.log(row)
      _this.payRow = row
      _this.payRow.storage_cost = row.file_size_byte * row.duration * _this.storage / 365
      _this.payRow.amount_minprice = Number(_this.payRow.storage_cost * _this.biling_price).toFixed(9)
      _this.cost.storage_cost_low = Number(_this.payRow.storage_cost * _this.biling_price * 2).toFixed(9)
      _this.cost.storage_cost_average = Number(_this.payRow.storage_cost * _this.biling_price * 3).toFixed(9)
      _this.cost.storage_cost_high = Number(_this.payRow.storage_cost * _this.biling_price * 5).toFixed(9)

      _this.payVisible = true
      return false
    },
    mintFunction(row){
      this.mintCID = row.payload_cid
      this.dealID = row.deal_id
      this.fileSize = row.file_size
      this.mineVisible=true
    },
    refundClick(row){
        let _this = this
        let contract_instance = new web3.eth.Contract( first_contract_json );
        contract_instance.options.address = _this.gatewayContractAddress
        _this.loading = true

        let payObject = {
            from: _this.metaAddress,
            gas: web3.utils.toHex(_this.$root.PAY_GAS_LIMIT),
        };
        
        let lockObj = {
            id: row.payload_cid,
            orderId: "",
            dealId: row.deal_id,
            amount: row.locked_fee,
            recipient: _this.recipientAddress, //todo:
        }
        console.log(lockObj)
        try{
          contract_instance.methods.unlockTokenPayment(lockObj)
          .send(payObject)
          .on('transactionHash', function(hash){
              // console.log('unlock hash console:', hash);
              _this.txHash = hash
          })
          .on('confirmation', function(confirmationNumber, receipt){
              // console.log('confirmationNumber console:', confirmationNumber, receipt);
          })
          .on('receipt', function(receipt){
              // console.log('receipt console:', receipt);
              _this.checkTransaction(receipt.transactionHash, cid)
              _this.txHash = receipt.transactionHash
          })
          .on('error', function(error){
              // console.log('error console:', error)
              _this.loading = false
              _this.failTransaction = true
          }); 
        }catch (err){
          console.log(err)
          _this.loading = false
        }
    },
    payStartClick(rowAmount){
      let _this = this
      if(_this.metaAddress){
        _this.loading = true
        // 发起请求
        axios.get(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/billing/deal/lockpayment/info?payload_cid=${_this.payRow.payload_cid}&wallet_address=${_this.metaAddress}`,{
            headers: {
            // 'Authorization': "Bearer "
            },
        })
        .then((res) => {
            if (res.data.status == "success") {
                if(res.data.data.tx_hash){
                    _this.$message.error('This file has been paid.')
                    _this.loading = false
                    _this.getData()
                    return false
                }else{
                  // 授权代币
                  contract_erc20 = new web3.eth.Contract( erc20_contract_json );
                  contract_erc20.options.address = _this.usdcAddress
                  // 查询剩余代币余额为：
                  contract_erc20.methods.balanceOf(_this.metaAddress).call()
                  .then(balance => {
                      let usdcAvailable = web3.utils.fromWei(balance, 'ether');
                      console.log('Available:', usdcAvailable, rowAmount)
                      // 判断支付金额是否大于代币余额
                      if(Number(rowAmount) > Number(usdcAvailable)){
                          _this.$message.error('Insufficient balance')
                          return false
                      }else{
                        contract_erc20.methods.allowance(_this.gatewayContractAddress, _this.metaAddress).call()
                        .then(resultUSDC => {
                            console.log('allowance：'+ resultUSDC);
                            if(resultUSDC < web3.utils.toWei(rowAmount, 'ether')){
                                contract_erc20.methods.approve(_this.gatewayContractAddress, web3.utils.toWei(rowAmount, 'ether')).send({from:  _this.metaAddress})
                                .then(receipt => {
                                    // console.log(receipt)
                                })
                                .catch(error => {
                                  // console.log('errorerrorerror', error)
                                })
                            }
                            _this.contractSend(res.data.data.payload_cid, web3.utils.toWei(rowAmount, 'ether'))
                        })
                      }
                  })
                }
            } else {
                _this.$message.error('Fail')
                _this.loading = false
            }
        }).catch(error => {
            console.log(error)
            _this.loading = false
        })

        return false
      }
    },
    contractSend(cid, payAmount){
        let _this = this
        // 合约转账
        let contract_instance = new web3.eth.Contract( first_contract_json );
        contract_instance.options.address = _this.gatewayContractAddress

        let payObject = {
            from: _this.metaAddress,
            gas: web3.utils.toHex(_this.$root.PAY_GAS_LIMIT),
            // gasPrice: web3.utils.toHex(web3.utils.toWei(_this.ruleForm.gasprice + '', 'gwei')),
            // value: web3.utils.toHex(web3.utils.toWei(_this.ruleForm.amount, 'ether')),
        };
        
        let lockObj = {
            id: cid,
            minPayment: web3.utils.toWei(String(_this.payRow.amount_minprice), 'ether'),
            amount: payAmount,
            lockTime: 86400 * Number(_this.$root.LOCK_TIME), // one day
            recipient: _this.recipientAddress, //todo:
        }
        
        contract_instance.methods.lockTokenPayment(lockObj)
        .send(payObject)
        .on('transactionHash', function(hash){
            // console.log('hash console:', hash);
            _this.loadMetamaskPay = true
            _this.loading = false
            _this.txHash = hash
        })
        .on('confirmation', function(confirmationNumber, receipt){
            // console.log('confirmationNumber console:', confirmationNumber, receipt);
        })
        .on('receipt', function(receipt){
            // receipt example
            // console.log('receipt console:', receipt);
            _this.checkTransaction(receipt.transactionHash, cid)
            _this.txHash = receipt.transactionHash
        })
        .on('error', function(error){
            // console.log('error console:', error)
            // console.error
            _this.loading = false
            _this.loadMetamaskPay = false
            _this.failTransaction = true
        }); 
    },
    checkTransaction(txHash, cid) {
        let _this = this
        web3.eth.getTransactionReceipt(txHash).then(
            res => {
                console.log('checking ... ');
                if (!res) { return _this.timer = setTimeout(() => { _this.checkTransaction(txHash, cid); }, 2000); }
                else {
                    clearTimeout(_this.timer)
                    setTimeout(function(){
                      _this.loading = false
                      _this.loadMetamaskPay = false
                      _this.finishTransaction = true
                    },2000)
                }
            },
            err => { console.error(err); }
        );
    },
    finishClose(){
        this.finishTransaction = false
        this.mintTransaction = false
        this.getData()
    },
    getDialog(dialog, rows){
        this.payVisible = dialog
        if(rows) this.payStartClick(rows)
    },
    getMintDialog(dialog, tokenId, nftHash){
        let _this = this
        _this.mineVisible = dialog
        if(nftHash) {
          _this.tokenId = tokenId
          _this.txHash = nftHash
          _this.mintTransaction = true
        }
    },
    exChange(row, rowList) {
      var that = this
      if (rowList.length) {
        that.expands = []
        if (row) {
          that.expands.push(row.uuid)
          //open
        }
        that.tableTrClick(row)
      } else {
        that.expands = []
        //retract
      }
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
    tableTrClick(row, type) {
      let _this = this
      _this.loadingChild = true
      _this.tableDataChild = []
      myAjax
        .getTasksListDetails(row.uuid)
        .then((response) => {
          if (response.status == "success" || response.status == 'Success') {
            const data = response.data;
            _this.tableDataChild = response.data.deal;
            _this.tableDataChild.map((item, index) => {
              item.file_size_gb = _this.byteChange(item.file_size)
              item.storage_cost = item.file_size_gb * _this.storage
              item.dealVisible = false
              item.act = false
              item.contVisible = false
              item.updated_at = item.updated_at
                ? item.updated_at.length < 13
                  ? moment(new Date(parseInt(item.updated_at * 1000))).format(
                      "YYYY-MM-DD HH:mm:ss"
                    )
                  : moment(new Date(parseInt(item.updated_at))).format(
                      "YYYY-MM-DD HH:mm:ss"
                    )
                : "-";

              if(index == 0){
                _this.payFun(item.payload_cid)
                _this.toAddress = item.contract_id
                _this.payment.cid = item.payload_cid
                if(item.storage_cost && _this.payment.biling_price) _this.payment.amount = String(item.storage_cost * _this.payment.biling_price * _this.$root.PAY_WITH_MULTIPLY_FACTOR)
                if(type) _this.walletInfo()
              } 
            });
            // console.log(_this.tableDataChild)
          } else {
            // _this.$message.error(response.message);
          }
          _this.loadingChild = false;
        })
        .catch((error) => {
          console.log(error);
          _this.loadingChild = false;
        });
    
    },
    cancelClick(uuid){
      let _this = this;
      const h = _this.$createElement;
      _this
        .$msgbox({
          title: 'Tips',
          message: h("p", null, [
            h("span", null, 'Are you sure you want to modify the status to '),
            h("i", { style: "color: teal;font-style: normal" }, 'Cancel'),
          ]),
          showCancelButton: true,
          beforeClose: (action, instance, done) => {
            if (action === "confirm") {
              instance.confirmButtonLoading = true;
              instance.confirmButtonText = 'Executing...';

              let cancelDeal_api = `${process.env.BASE_API}tasks/${uuid}/paymentgateway/cancel`
              axios.put(cancelDeal_api, {'wallet_address': _this.metaAddress}, {
                  headers: {
                      // 'Authorization': "Bearer "+ _this.$store.getters.accessToken
                  },
              }).then(res => {
                  if(res.data.status == 'success'){
                      _this.$message({
                          message: 'Cancel Success',
                          type: 'success'
                      });
                      done();
                      setTimeout(() => {
                        instance.confirmButtonLoading = false;
                      }, 300);
                      _this.getData()
                  }else{
                      _this.$message.error(res.data.message);
                  }
              }).catch(error => {
                  console.log(error)
                  done();
                  instance.confirmButtonLoading = false;
              })

            } else {
              done();
              instance.confirmButtonLoading = false;
            }
          },
        })
        .then((action) => {})
        .catch((action) => {});
    },
    payFun(cid, type){
      let _this = this
      let pay_api = `${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/events/logs/lock/${cid}?wallet_address=${_this.metaAddress}`

      axios.get(pay_api, {
          headers: {
            // 'Authorization': "Bearer "+ _this.$store.getters.accessToken
          },
      })
      .then((json) => {
          if(json.data.status == 'success'){
            if(type){
                if(json.data.data.polygon.length>0 || json.data.data.goerli.length>0) {
                  clearTimeout(_this.timer)
                  _this.sendDeal()
                  _this.timeIndex = 0
                }
            }
            _this.paySuccess = json.data.data.polygon.length>0 || json.data.data.goerli.length>0 ? true : false
          }else{
            _this.paySuccess = true
          }
      }).catch(error => {
          console.log(error)
      })
    },
    getPay(status, cid){
      if(!status) return false
      let _this = this
      _this.timeIndex = 0
      _this.timer = setInterval(() => { 
        _this.timeIndex += 1
        _this.payFun(cid, 1); 
        if(_this.timeIndex > 9) clearTimeout(_this.timer)
      }, 5000);
    },
    sendDeal(){
        let _this = this
        let sendDeal_api = `${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/lotus/deal/${_this.expands[0]}?wallet_address=${_this.metaAddress}`
        axios.get(sendDeal_api, {
            headers: {
                // 'Authorization': "Bearer "+ _this.$store.getters.accessToken
            },
        }).then(res => {
            if(res.data.status == 'success'){
                _this.$message({
                    message: localStorage.getItem('languageMcp') == 'en'?'Send Deal Success': '发送交易成功',
                    type: 'success'
                });
                _this.getData()
            }else{
                _this.$message.error(localStorage.getItem('languageMcp') == 'en'?'Send Deal Fail': '发送交易失败');
            }
        }).catch(error => {
            console.log(error)
        })

    },
    //查询
    search() {
      let _this = this;
      _this.parma.limit = 10;
      _this.paginationShow = true;
      _this.parma.offset = 1;
      _this.getData();
    },
    clearAll() {
      let _this = this;
      _this.searchValue = "";
      _this.parma.limit = 10;
      _this.parma.offset = 1;
      _this.getData();
    },
    handleCurrentChange(val) {
      this.parma.offset = val;
      this.getData();
    },
    stats(){
        let _this = this
        _this.loading = true
        if(_this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS){
            _this.gatewayContractAddress = _this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS
            _this.usdcAddress = _this.$root.USDC_ADDRESS
            _this.recipientAddress = _this.$root.RECIPIENT
            
            let stats_api = `${process.env.BASE_API}stats/storage?wallet_address=${_this.metaAddress}`
            axios.get(stats_api, {
                headers: {
                    // 'Authorization': "Bearer "+ _this.$store.getters.accessToken
                },
            }).then(res => {
                if(res.data.data){
                    let cost = res.data.data.average_price_per_GB_per_year.split(" ")
                    if(cost[0]) _this.storage = cost[0]
                }
            }).catch(error => {
                console.log(error)
            })
            
            let billing_api = `${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/billing/price/filecoin?wallet_address=${_this.metaAddress}`
            axios.get(billing_api, {
                headers: {
                    // 'Authorization': "Bearer "+ _this.$store.getters.accessToken
                },
            }).then(res => {
                if(res.data.data){
                    _this.biling_price = res.data.data
                }
            }).catch(error => {
                console.log(error)
            })
            _this.getData()
        }else {
            setTimeout(function(){
                _this.stats()
            }, 1000)
        }
    },
    getData() {
      let _this = this;
      _this.loading = true;
      let offset =
        _this.parma.offset > 0 ? _this.parma.offset - 1 : _this.parma.offset;
      let parma = {
        page_size: _this.parma.limit,
        page_number: _this.parma.offset,
        file_name: _this.searchValue,
        source_id: 4,
        wallet_address: _this.metaAddress
      };

      _this.tableData = []

      let storage_api = `${process.env.BASE_PAYMENT_GATEWAY_API}api/v1/storage/tasks/deals?${QS.stringify(parma)}`
      // let storage_api = `./static/pay-status-response.json?${QS.stringify(parma)}`

      axios.get(storage_api, {
          headers: {
              // 'Authorization': "Bearer "+ _this.$store.getters.accessToken
          },
      })
      .then((response) => {
          if(response.data.status == 'success'){
            const data = response.data.data;
            _this.parma.total = Number(response.data.page_info.total_record_count);
            _this.tableData = response.data.data;
            _this.tableData.map((item,s) => {
              item.payloadAct = false
              item.duration = item.duration/2880
              item.file_size_byte = _this.byteChange(item.file_size)
              item.create_at = item.create_at
                ? item.create_at.length < 13
                  ? moment(new Date(parseInt(item.create_at * 1000))).format(
                      "YYYY-MM-DD HH:mm:ss"
                    )
                  : moment(new Date(parseInt(item.create_at))).format(
                      "YYYY-MM-DD HH:mm:ss"
                    )
                : "-";
            });
            setTimeout(function(){
              _this.loading = false
            }, 2000)
          } else {
            _this.$message.error(response.message);
            _this.loading = false
          }
      }).catch(error => {
          console.log(error)
          _this.loading = false;
      })
      return false

      // myAjax
      //   .getPaymentDeals(parma)
      //   .then((response) => {
      //     if (response.status == "success") {
      //       const data = response.data;
      //       _this.expands = []
      //       // _this.tableData = Array.from(new Set(response.data.deals));
      //       _this.parma.total = response.paging_info.total_items;
      //       _this.tableData = response.data.deals;
      //       _this.tableData.map((item,s) => {
      //         item.payloadAct = false
      //         item.create_at = item.create_at
      //           ? item.create_at.length < 13
      //             ? moment(new Date(parseInt(item.create_at * 1000))).format(
      //                 "YYYY-MM-DD HH:mm:ss"
      //               )
      //             : moment(new Date(parseInt(item.create_at))).format(
      //                 "YYYY-MM-DD HH:mm:ss"
      //               )
      //           : "-";
      //       });

      //     } else {
      //       _this.$message.error(response.message);
      //     }
      //     _this.loading = false
      //   })
      //   .catch((error) => {
      //     console.log(error);
      //     _this.loading = false;
      //   });
    },
    byteChange(limit){
        var size = "";
        // 只转换成GB
        if(limit <= 0){
            return '-'
        }else{
            size = limit/( 1000 * 1000 * 1000)  //or 1024
        }
        return size
        // return Number(size).toFixed(3);
    },
    unique(arr) {
      const res = new Map();
      return arr.filter((arr) => !res.has(arr.value) && res.set(arr.value, 1));
    },
  },
  mounted() {
    let _this = this;
    document.getElementById("content-box").scrollTop = 0;
    _this.$store.dispatch("setRouterMenu", 1);
    _this.$store.dispatch("setHeadertitle", _this.$t('navbar.deal'));
    _this.stats()
    document.onkeydown = function (e) {
      if (e.keyCode === 13) {
      }
    };

  },
  filters: {
    NumFormat(value) {
      if (!value) return "-";
      return value;
    },
    NumStorage(value) {
      if (!value) return "- FIL";
      return value.toFixed(8) + ' FIL';
    },
    formatbytes: function (bytes) {
      if (!bytes) return "-";
      if (bytes == 0) return '0 B';
      var k = 1000, // or 1024
          sizes = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'],
          i = Math.floor(Math.log(bytes) / Math.log(k));
      return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i];
    },
  },
};
</script>


<style scoped lang="scss">
#dealManagement {
  position: relative;
  padding: 0.25rem 0.2rem 0.2rem;
  .el-alert /deep/{
      position: absolute;
      left: 0;
      top: 0;
      .el-alert__content{
          display: flex;
          align-items: center;
      }
  }
  .tabTaskStyle {
    display: flex;
    justify-content: flex-end;
    align-items: center;
    flex-wrap: wrap;
    margin-bottom: 0.2rem;
    overflow: hidden;
    .createTask {
      padding: 0.14rem 0;
      background-color: #fff;
      border-radius: 0.1rem;
      a {
        display: block;
        height: 0.35rem;
        padding: 0 0.1rem;
        margin: 0 0.27rem;
        background-color: #ffb822;
        line-height: 0.35rem;
        border-radius: 0.08rem;
        text-align: center;
        color: #fff;
        font-size: 0.1972rem;
        border: 0;
        outline: none;
      }
    }
  }
  .upload {
    padding: 0 0.17rem;
    margin: 0.2rem auto;
    background-color: #fff;
    border-radius: 0.1rem;
    overflow: hidden;

    .title {
      width: 100%;
      text-align: left;
      font-size: 0.1972rem;
      color: #000;
      line-height: 0.42rem;
      text-indent: 0.08rem;
    }

    .upload_form {
      display: flex;
      justify-content: center;
      // align-items: center;
      align-items: baseline;

      .el-form /deep/ {
        width: 42%;
        margin: 0;

        .el-form-item {
          display: flex;
          align-items: center;
          width: auto;
          margin: 0.08rem auto;

          .el-form-item__label {
            width: auto;
            max-width: 2rem;
            line-height: 1.5;
            text-align: left;
          }

          .el-form-item__content {
            display: flex;
            align-items: center;
            flex-wrap: wrap;

            h4 {
              width: 100%;
              font-size: 0.1372rem;
              font-weight: 500;
              line-height: 1.7;
            }

            h5 {
              width: 90%;
              margin-top: 5px;
              font-size: 10px;
              font-weight: 500;
              line-height: 1.2;
              color: #737373;
            }

            .el-tag,
            .el-button--small {
              margin: 0 5px 5px 0;
            }

            .el-input {
              width: auto;

              .el-input__inner {
                height: 0.32rem;
                font-size: 0.1372rem;
                line-height: 0.32rem;
              }
            }

            .el-form-item__error {
              padding-top: 0;
              margin: 0 0.1rem;
              position: relative;
              float: right;
            }

            .upload-demo {
              .el-upload--text {
                float: left;
                width: auto;
                height: auto;
                text-align: left;
                border: 0;

                .el-button--primary {
                  height: 0.32rem;
                  padding: 0 0.2rem;
                  line-height: 0.32rem;
                  background-color: transparent;
                  border: 1px solid #2c4c9e;
                  border-radius: 0.08rem;
                  color: #2c4c9e;
                  font-size: 0.1372rem;
                }
              }

              .el-upload-list {
                width: 100%;
                float: none;
                clear: both;
              }
            }

            .el-upload__tip {
              // float: left;
              margin: 0 0 0 0.1rem;
              color: #737373;
              line-height: 1;
            }

            .el-radio {
              .el-radio__inner {
                border-color: #d9d9d9;
                background-color: #d9d9d9;
              }
            }

            .el-radio.is-checked {
              .el-radio__inner {
                border-color: #0b318f;
                background-color: #0b318f;
              }

              .el-radio__inner::after {
                width: 6px;
                height: 6px;
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
      margin: 0.25rem 0.3rem 0.15rem;

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
  }

  .form {
    position: relative;
    padding: 0.1rem 0.17rem 0.2rem;
    background-color: #fff;
    border-radius: 0.1rem;

    .loadMetamaskPay{
        position: absolute;
        left: 0;
        top: 0;
        right: 0;
        bottom: 0;
        z-index: 99;
        display: flex;
        align-items: center;
        justify-content: center;
        background-color: rgba(255,255,255,.9);
        .el-loading-spinner{
            top: 0;
            position: relative;
            margin: 0 0 0.2rem;
        }
        p{
            font-size: 14px;
            font-weight: 600;
            color: #666;
        }
    }
    .form_top {
      display: flex;
      align-items: center;
      flex-wrap: wrap;

      .title {
        width: 100%;
        margin: 0;
        text-align: left;
        font-size: 0.12rem;
        font-weight: 600;
        color: #000;
        line-height: 0.42rem;
        text-indent: 0;
      }

      .upload_title{
        width: 100%;
        margin: 0 0 0.05rem;
        text-align: left;
        font-weight: 600;
        line-height: 1.5;
        text-indent: 0;
        font-size: 0.13rem;
        color: #222;
      }

      .search {
        display: flex;
        align-items: center;
        justify-content: space-between;
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
          width: auto;

          .el-input__inner {
            width: 100%;
            max-width: 3rem;
            border-radius: 0.08rem;
            border: 1px solid #f8f8f8;
            color: #737373;
            font-size: 0.12rem;
            height: 0.24rem;
            line-height: 0.24rem;
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
      .search_file{
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 100%;
        height: 0.42rem;
        p{
          font-size: 0.13rem;
          color: #222;
        }
        .search_right {
          display: flex;
          align-items: center;
          width: 50%;
          // margin-left: 0.3rem;
          .el-button /deep/ {
              height: 0.3rem;
              padding: 0 0.15rem;
              margin: 0;
              color: #fff;
              line-height: 0.3rem;
              font-size: 0.1372rem;
              border: 0;
              border-top-left-radius: 0;
              border-bottom-left-radius: 0;
          }
        }
        .el-input /deep/ {
          float: left;

          .el-input__inner {
            width: 100%;
            color: #737373;
            font-size: 0.12rem;
            height: 0.3rem;
            line-height: 0.3rem;
            padding: 0 0.27rem;
            border-top-right-radius: 0;
            border-bottom-right-radius: 0;
          }

          .el-input__icon {
            line-height: 0.3rem;
          }
        }
      }
    }

    .datasetStyle{
      padding: 0 0 0 20px;
      margin: 0.1rem 0 0;
      background: url(../../../assets/images/download.png) no-repeat left center;
      background-size: 15px;
      font-size: 0.14rem;
      span{
        color: rgb(10, 49, 142);
        cursor: pointer;
      }
    }

    .form_table {
      position: relative;
      margin: 0.1rem 0 0.1rem;
      border: 1px solid #e6e6e6;
      border-radius: 4px;
      overflow: hidden;
      .statusStyle /deep/{
        display: inline-block;
        border: 1px solid;
        padding: 0.04rem 0.05rem;
        border-radius: 0.05rem;
        line-height: 1.5;
        // color: inherit !important;
        span{
          white-space: normal;
        }
      }

      .el-table /deep/ {
        overflow: visible;
        overflow-x: scroll;
        .el-table__body-wrapper,
        .el-table__header-wrapper {
          overflow: visible;
        }

        tr {
          cursor: pointer;
          th {
            height: 0.5rem;
            padding: 0;
            background-color: #f2f2f2 !important;
            text-align: center;

            .cell {
              word-break: break-word;
              font-weight: 500;
              color: #737373;
              text-transform: uppercase;
              .caret-wrapper{
                display: none;
              }
              .tips{
                display: flex;
                align-items: center;    
                justify-content: center;
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
          }

          .descending{
              position: relative;
              &::after{
                content: '';
                position: absolute;
                bottom: 0;
                left: 0;
                right: 0;
                height: 3px;
                background-color: #0a318e;
              }
          }

          .ascending{
              position: relative;
              &::after{
                content: '';
                position: absolute;
                top: 0;
                left: 0;
                right: 0;
                height: 3px;
                background-color: #0a318e;
              }
          }

          th.is-leaf {
            border-bottom: 0;
          }

          // th:first-child {
          //   border-top-left-radius: 4px;
          //   border-bottom-left-radius: 4px;
          // }

          // th:last-child {
          //   border-top-right-radius: 4px;
          //   border-bottom-right-radius: 4px;
          // }

          td {
            padding: 0.15rem 0.05rem;
            border-bottom: 1px solid #f2f2f2;

            .cell {
              padding: 0;
              font-size: 0.1372rem;
              word-break: break-word;
              color: #000;
              text-align: center;
              line-height: 0.25rem;
              overflow: visible;

              .el-rate__icon {
                font-size: 0.16rem;
                margin-right: 0;
              }
              .el-icon-arrow-right{
                font-weight: bold;
                font-size: 0.13rem;
              }
              .rightClick {
                color: #0c3090;
                cursor: pointer;
              }

              .hot-cold-box {
                position: relative;
                .elTips{
                    display: flex;
                    align-items: center;
                    justify-content: center;
                    padding: 0 0.1rem;
                    line-height: 1.5;
                    text-align: center;
                    white-space: normal;
                    color: inherit;
                    text-shadow: 0 0 black;
                    text-indent: 0;
                    font-size: inherit;
                    font-weight: normal;
                    text-decoration: underline;
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
                    .el-button{
                      
                        img{
                            display: inline-block;
                            float: right;
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
                .cidLink{
                    line-height: 0.25rem;
                    overflow: hidden;
                    white-space: normal;
                    display: -webkit-box;
                    -webkit-line-clamp: 1;
                    -webkit-box-orient: vertical;
                    color: inherit;
                }
                // .el-button {
                //   width: 100%;
                //   border: 0;
                //   padding: 0;
                //   background-color: transparent;
                //   font-size: 0.1372rem;
                //   word-break: break-word;
                //   color: #000;
                //   text-align: center;
                //   line-height: 0.25rem;
                //   overflow: hidden;
                //   text-overflow: ellipsis;
                //   white-space: normal;
                //   display: -webkit-box;
                //   -webkit-line-clamp: 1;
                //   -webkit-box-orient: vertical;
                //   span, a{
                //     line-height: 0.25rem;
                //     overflow: hidden;
                //     text-overflow: ellipsis;
                //     white-space: normal;
                //     display: -webkit-box;
                //     -webkit-line-clamp: 1;
                //     -webkit-box-orient: vertical;
                //     color: inherit;
                //   }
                // }
                .cidMore{
                  position: absolute;
                  top: 0;
                  right: 0;
                  bottom: 0;
                  background-color: #fff;
                  width: 0.22rem;
                  padding: 0;
                    &::before{
                      content: '···';
                      font-weight: bold;
                      color: #0a318e;
                    }
                }
                .el-button{
                    width: 100%;
                    border: 0;
                    padding: 0;
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
                        text-transform: uppercase;
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
                .uploadBtn{
                  width: auto;
                  padding: 0 0.1rem;
                  margin: auto;
                  color: #5a5a5a;
                  box-shadow: 0 0 0.06rem rgba(191, 191, 191, 0.32);
                  border-radius: 4px;
                  border: 1px solid #DCDFE6;
                  white-space: nowrap;
                  display: inline-block;
                  &:hover{
                    background-color: #fff;
                  }
                }
                .opacity{
                  box-shadow: none;
                  color: #b8b8b8;
                }
              }

              .el-button.el-icon-upload {
                padding: 0 0.03rem;
                line-height: 0.25rem;
                font-size: 0.1372rem;
              }

              .scoreStyle {
                width: 100%;
                clear: both;
                text-align: center;
                color: #ffb822;
                cursor: pointer;
              }

              .scoreStyle:hover {
                text-decoration: underline;
              }
            }
          }

          // td.el-table_1_column_1 {
          //   .cell {
          //     color: #0c3090;
          //   }
          // }
          .expandArea{
            display: flex;
            justify-content: space-between;
            align-items: flex-start;
            margin: 0.2rem auto;
            width: 100%;
            max-width: 1200px;
            .demo-table-expand {
              font-size: 0;
            }
            .demo-table-expand label {
              width: 140px;
              font-size: 0.13rem;
              font-weight: 600;
              color: #737373;
              text-align: right;
            }
            .demo-table-expand .el-form-item {
              display: flex;
              align-items: center;
              margin-right: 0;
              margin-bottom: 0;
              width: 100%;
              .hot-cold-box {
                position: relative;
                float: left;
                margin-right: 0.1rem;
                .cidLink{
                    line-height: 0.25rem;
                    overflow: hidden;
                    white-space: normal;
                    display: -webkit-box;
                    -webkit-line-clamp: 1;
                    -webkit-box-orient: vertical;
                    color: inherit;
                }
                .cidMore{
                  position: absolute;
                  top: 0;
                  right: 0;
                  bottom: 0;
                  background-color: #fff;
                  width: 0.22rem;
                  padding: 0;
                    &::before{
                      content: '···';
                      font-weight: bold;
                      color: #0a318e;
                    }
                }
                .uploadBtn{
                  width: auto;
                  padding: 0.06rem 0.1rem;
                  margin: auto;
                  color: #5a5a5a;
                  box-shadow: 0 0 0.06rem rgba(191, 191, 191, 0.32);
                  border-radius: 4px;
                  border: 1px solid #DCDFE6;
                  white-space: nowrap;
                  &:hover{
                    background-color: #fff;
                  }
                }
                .el-button{
                    width: 100%;
                    border: 0;
                    padding: 0;
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
                .el-button--primary {
                    color: #FFF !important;
                    background-color: #409EFF !important;
                    border-color: #409EFF !important;
                }
                .el-button:hover{
                    img{
                        display: inline-block;
                    }
                }
              }
            }
            .el-table__expanded-cell{
              // padding: 0;
              .el-descriptions{
                overflow: auto;
                line-height: 0.25rem;
                z-index: 999;
                .el-descriptions__body{
                  padding: 0.2rem;
                  .el-descriptions__table{
                    tr{
                      background: transparent !important;
                      th, td{
                        height: auto;
                        padding: 0.1rem;
                        text-align: left;
                        border: 0;
                      }
                      th{
                        text-transform: uppercase;
                        background: #fafafa !important;
                        &::after{
                          display: none;
                        }
                      }
                      
                      .hot-cold-box {
                        position: relative;
                        .cidLink{
                            line-height: 0.25rem;
                            overflow: hidden;
                            white-space: normal;
                            display: -webkit-box;
                            -webkit-line-clamp: 1;
                            -webkit-box-orient: vertical;
                            color: inherit;
                        }
                        .cidMore{
                          position: absolute;
                          top: 0;
                          right: 0;
                          bottom: 0;
                          background-color: #fff;
                          width: 0.22rem;
                          padding: 0;
                            &::before{
                              content: '···';
                              font-weight: bold;
                              color: #0a318e;
                            }
                        }
                        .uploadBtn{
                          width: auto;
                          padding: 0.06rem 0.1rem;
                          margin: auto;
                          color: #5a5a5a;
                          box-shadow: 0 0 0.06rem rgba(191, 191, 191, 0.32);
                          border-radius: 4px;
                          border: 1px solid #DCDFE6;
                          white-space: nowrap;
                          &:hover{
                            background-color: #fff;
                          }
                        }
                        .el-button{
                            width: 100%;
                            border: 0;
                            padding: 0;
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
            }
            .storageStyle{
              display: flex;
              flex-wrap: wrap;
              align-items: center;
              width: 40%;
              max-width: 350px;
              border: 2px solid #cecece;
              border-radius: 0.1rem;
              padding: 0.2rem;
              font-size: 0.16rem;
              text-align: center;
              div{
                width: 100%;
                line-height: 2;
                margin: 0 auto;
              }
              .costTitle{
                font-size: 0.18rem;
                font-weight: 600;
                color: #000;
              }
              .costPrice{
                margin: 0.1rem 0;
                font-size: 0.18rem;
                font-weight: 600;
                color: #e00000;
              }
              .costDesc{
                width: 80%;
                font-size: 0.14rem;
                color: #000;
                line-height: 1;
              }
            }
          }
          &:hover{
            td{
              .cell {
                .hot-cold-box {
                  .cidMore{
                    background-color: #F5F7FA;
                  }
                }
              }
            }
          }
        }
        .current-row{
          td{
            background-color: #F5F7FA;
            .cell {
              .hot-cold-box {
                .cidMore{
                  background-color: #F5F7FA;
                }
              }
            }
          }
        }
      }

      .el-select /deep/ {
        position: relative;
        cursor: pointer;
        color: transparent;

        .el-input {
          background: transparent;
          color: transparent;

          .el-input__inner {
            height: 0.27rem;
            line-height: 0.27rem;
            padding: 0 0.07rem;
            background: transparent;
            color: transparent;
            border-color: #f7f7f7;
            border-radius: 0.08rem;
          }

          .el-input__suffix {
            display: none;
          }

          input::-webkit-input-placeholder {
            color: transparent;
          }

          input::-moz-placeholder {
            /* Mozilla Firefox 19+ */
            color: transparent;
          }

          input:-moz-placeholder {
            /* Mozilla Firefox 4 to 18 */
            color: transparent;
          }

          input:-ms-input-placeholder {
            /* Internet Explorer 10-11 */
            color: transparent;
          }
        }
      }

      .actionStyle {
        position: relative;
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;

        .actionBtn {
          display: flex;
          align-items: center;
          justify-content: center;
          width: 0.4rem;
          padding: 0.1rem 0;
          margin: auto;
          border-radius: 0.08rem;
          background-color: #d5d7de;
          border: 1px solid #f7f7f7;

          span {
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

        .actStatus {
          // display: none;
          position: absolute;
          left: 0;
          right: 0;
          top: 100%;
          background: #fff;
          border: 1px solid #f7f7f7;
          border-radius: 0.08rem;
          z-index: 1002;

          p {
            padding: 0.1rem 0;
            border-top: 1px solid #f7f7f7;
            font-size: 0.1372rem;
            color: #ffb822;
            line-height: 1;
          }

          p:nth-child(1) {
            border: 0;
            color: #1dc9b7;
          }

          p:nth-child(2) {
            // color: #fd397a;
            color: #404040;
          }

          p:nth-child(3) {
            color: #1dc9b7;
          }

          // p:nth-child(4){
          //     color: #404040;
          // }
          p:hover {
            font-size: 0.14rem;
          }
        }
      }

      .actionStatus {
        display: inline-block;
        padding: 0 0.05rem;
        background-color: #1dc9b7;
        font-size: 0.135rem;
        color: #fff;
        border-radius: 0.08rem;
        line-height: 2.2;
        white-space: nowrap;
      }
    }

    .form_pagination {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 0.35rem;
      text-align: center;
      margin: 0.05rem 0;
      .pagination {
        display: flex;
        align-items: center;
        font-size: 0.1372rem;
        color: #000;

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
    }
  }

  .createStyle {
    position: absolute;
    top: -0.5rem;
    right: 0.2rem;
    padding: 0.08rem 0.2rem;
    background-color: #ffb822;
    border: 0;
    font-size: 0.17rem;
    z-index: 1003;
  }

  .el-dialog__wrapper /deep/ {    
    display: flex;
    align-items: center;
    .scoreCont {
      text-align: center;
      outline: none;

      .el-rate {
        outline: none;
      }
    }

    .el-dialog {
      .el-dialog__header {
        display: flex;
      }

      .el-dialog__footer {
        padding: 0.2rem 0.2rem;
        overflow: hidden;

        .el-button {
          float: right;
          margin-left: 0.2rem;
          padding: 0.07rem 0.2rem;
          border-radius: 0.08rem;
          font-size: 0.1372rem;
        }

        .el-button--primary {
          background-color: #0b318f;
          border-color: #0b318f;
        }

        .el-button--default {
          background-color: #fff;
          border-color: #0b318f;
          color: #0b318f;
        }

        .is-disabled {
          opacity: 0.5;
        }
      }
    }

    
    .metaM{
      .el-dialog__body{
        padding: 0.25rem 0.25rem 0.2rem;
        .el-row{
          border-radius: 0.08rem;
          padding: 0.16rem;
          margin: 0.12rem 0px;
          border: 1px solid rgb(240, 185, 11);
          text-align: center;
          display: flex;
          -webkit-box-pack: justify;
          justify-content: space-between;
          -webkit-box-align: center;
          align-items: center;
          transition: all 0.3s ease 0s;
          background: rgba(240, 185, 11, 0.1);
          position: static;
          .el-col{
            text-align: left;
            font-size: 0.14rem;
            img{
              float: right;
              height: 0.24rem;
            }
          }
        }
      }
      .el-dialog__footer{
        padding: 0 0.25rem 0.3rem;
        .dialog-footer{
          .el-button{
            width: 100%;
            font-size: 0.14rem;
            height: 0.4rem;
            padding: 0;
            background: #5c3cd3;
            color: #fff;
            border-radius: 0.08rem;
            &:hover{
              background: #4326ab;
            }
          }
          p{
            font-size: 0.13rem;
            line-height: 1.5;
            color: red;
            margin: 0.1rem 0 0;
          }
        }
      }
    }
    .completeDia{
        text-align: center;
        .el-dialog__header{
        display: none;
        }
        img{
            display: block;
            max-width: 100px;
            margin: auto;
        }
        h1{
            margin: 0.32rem auto 0.1rem;
            font-size: 0.32rem;
            font-weight: 500;
            line-height: 1.2;
            color: #191919;
            word-break: break-word;
        }
        h2{
            margin: 0.1rem auto 0.1rem;
            font-size: 19px;
            font-weight: 600;
            line-height: 1.2;
            color: #191919;
            word-break: break-word;
            text-align: center;
        }
        h4{
            font-size: 14px;
            font-weight: 500;
            line-height: 1.2;
            color: #191919;
            word-break: break-word;
            text-align: center;
        }
        h3, a{
            font-size: 0.16rem;
            font-weight: 500;
            line-height: 1.2;
            color: #191919;
            word-break: break-word;
        }
        a{
            text-decoration: underline;
            color: #007bff;
        }
        a.a-close{
            padding: 5px 45px;
            background: #5c3cd3;
            color: #fff;
            border-radius: 10px;
            cursor: pointer;
            margin: 0.2rem auto 0;
            display: block;
            width: max-content;
            text-decoration: unset;
        }
    }

    .wrongNet{
        margin: auto !important;
        background: #fff url(../../../assets/images/tip_bg.png) no-repeat;
        background-size: 1.45rem;
        background-position: -0.2rem -0.3rem;
        border-radius: 4px;
        .el-dialog__header{
            display: flex;
            color: #000;
            font-size: 18px;
            padding: 0.15rem 0.15rem 0.1rem;
            .el-dialog__headerbtn{
                i{
                    color: #000;
                    font-weight: 600;
                }
            }
        }
        .el-dialog__body {
            padding: 0.2rem 0.2rem 0.3rem;
            span{
              word-break: break-word;
              line-height: 1.5;    
              color: #606266;
              font-size: 14px;
            }
        }
    }

  }
}
@media screen and (max-width: 1024px) {
  #dealManagement {

    .form {
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

      .form_table {
        overflow-x: scroll;

        .el-table /deep/ {
          width: 1024px !important;

          tr {
            td {
              .cell {
                .cellFileName {
                  display: -webkit-box;
                  -webkit-line-clamp: 3;
                  -webkit-box-orient: vertical;
                  overflow: hidden;
                }
              }
            }
          }
        }
      }
    }
  }
}
@media screen and (max-width: 999px) {
  #dealManagement {
    padding: 0.15rem 0.1rem 0.2rem;

    .tabTaskStyle {
      .createTask {
        a {
          font-size: 0.17rem;
        }
      }
    }
    .upload {
      padding: 0.1rem;

      .upload_form {
        flex-wrap: wrap;

        .el-form /deep/ {
          width: 95%;
        }
      }
    }
  }
}
@media screen and (max-width: 470px) {
  #dealManagement{
    .tabTaskStyle {
      .createTask {
        a {
          font-size: 0.15rem;
        }
      }
    }
    .form{
      .form_top{
        .search_file{
          flex-wrap: wrap;
          .search_right{
            width: 100%;
            margin: 0.05rem 0 0.2rem;
          }
        }
      }
    }
  }
}
@media screen and (max-width: 341px) { 
}
</style>
