<template>
  <div id="dealManagement">
    <div class="form">
      <div class="form_top">
        <div class="search_file">
          <div class="search_right">
            <el-input :placeholder="$t('uploadFile.search_title')" prefix-icon="el-icon-search" :disabled="loading && firstIndex === 0" v-model="searchValue" clearable>
            </el-input>
          </div>
          <div class="createTask" v-loading="loading && firstIndex === 0">
            <a @click="uploadDigShow = true">
              <img src="@/assets/images/my_file/icon_Upload@2x.png" alt="">
              <span>{{$t('uploadFile.Upload_More_Files')}}</span>
            </a>
          </div>
        </div>
      </div>

      <div class="form_table">
        <el-table :data="tableData" ref="singleTable" stripe style="width: 100%" max-height="580" :empty-text="$t('deal.formNotData')" v-loading="loading" @sort-change="sortChange" :default-sort="{prop: 'date', order: 'descending'}" @filter-change="filterChange">
          <el-table-column prop="file_name" width="200" sortable="custom">
            <template slot="header" slot-scope="scope">
              <div class="tips" style="white-space: nowrap;">
                {{$t('uploadFile.file_name')}}

                <el-popover placement="top" popper-class="elPopTitle" width="200" trigger="hover" :content="$t('uploadFile.file_name_tooltip')">
                  <img slot="reference" src="@/assets/images/info.png" />
                </el-popover>
              </div>
            </template>
            <template slot-scope="scope">
              <div class="hot-cold-box" style="text-decoration: underline;font-size:0.17rem;" @click="toDetail(scope.row)">
                <span target="_blank" style="color: inherit;word-break: break-all;">{{ scope.row.file_name }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="file_size" :label="$t('uploadFile.file_size')" min-width="90" sortable="custom">
            <template slot-scope="scope">
              <div class="hot-cold-box">
                {{ scope.row.file_size | formatbytes }}
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="status" :label="$t('uploadFile.file_status')" min-width="110" :filters="[{text: $t('uploadFile.filter_status_Pending'), value: 'Pending'}, {text: $t('uploadFile.filter_status_Processing'), value: 'Processing'},
                       {text: $t('uploadFile.filter_status_Refundable'), value: 'Refundable'}, {text: $t('uploadFile.filter_status_Refunded'), value: 'Refunded'},
                       {text: $t('uploadFile.filter_status_Success'), value: 'Success'}, {text: $t('uploadFile.filter_status_Failed'), value: 'Failed'}]" :filter-multiple="false" :column-key="'payment'">
            <template slot-scope="scope">
              <el-popover :disabled="!scope.row.note" placement="top" popper-class="elPopTitle" class="item" width="200" trigger="hover" :content="scope.row.note">
                <el-button slot="reference" plain type="pending" class="statusStyle" v-if="scope.row.status&&scope.row.status.toLowerCase()=='pending'">
                  {{ languageMcs == "en" ? "Pending" : '待支付'}}
                </el-button>
                <el-button slot="reference" plain type="pending" class="statusStyle" v-else-if="(scope.row.status&&(scope.row.status.toLowerCase()=='completed' || scope.row.status.toLowerCase()=='refundable')&&scope.row.status_failed_file) || (scope.row.status&&scope.row.status.toLowerCase()=='processing'&&scope.row.status_failed_file&&scope.row.offline_deal.length>0)">
                  {{ languageMcs == "en" ? "Failed" : '失败'}}
                </el-button>
                <el-button slot="reference" plain type="primary" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='processing'&&scope.row.status_failed_file&&scope.row.offline_deal.length<=0">
                  {{ languageMcs == "en" ? "Processing" : '处理中'}}
                </el-button>
                <el-button slot="reference" plain type="success" class="statusStyle" v-else-if="scope.row.status&&(scope.row.status.toLowerCase()=='completed' || scope.row.status.toLowerCase()=='refundable' || scope.row.status.toLowerCase()=='processing' || scope.row.status.toLowerCase()=='success')&&scope.row.status_success_file">
                  {{ languageMcs == "en" ? "Success" : '完成'}}
                </el-button>
                <el-button slot="reference" type="successPart" class="statusStyle" v-else-if="scope.row.status_file">
                  {{ languageMcs == "en" ? "Success" : '完成'}}
                </el-button>
                <el-button slot="reference" type="danger" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='failed'">
                  {{ languageMcs == "en" ? "Fail" : '失败'}}
                </el-button>
                <el-button slot="reference" plain type="primary" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='processing'&&!scope.row.status_file">
                  {{ languageMcs == "en" ? "Processing" : '处理中'}}
                </el-button>
                <el-button slot="reference" plain type="success" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='active'&&!scope.row.status_file">
                  {{ languageMcs == "en" ? "Active" : '完成'}}
                </el-button>
                <el-button slot="reference" plain type="info" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='refunded'&&!scope.row.status_file">
                  {{ languageMcs == "en" ? "Refunded" : '已退款'}}
                </el-button>
                <el-button slot="reference" plain type="refunding" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='refunding'&&!scope.row.status_file">
                  {{ languageMcs == "en" ? "Refunding" : '可退款'}}
                </el-button>
                <el-button slot="reference" plain type="refunding" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='refundable'&&!scope.row.status_file">
                  {{ languageMcs == "en" ? "Refundable" : '可退款'}}
                </el-button>
                <el-button slot="reference" plain type="refunding" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='free'&&!scope.row.status_file">
                  {{ languageMcs == "en" ? "Free" : '免费'}}
                </el-button>
                <el-button slot="reference" plain type="info" class="statusStyle" v-else-if="scope.row.status&&scope.row.status.toLowerCase()=='completed'&&!scope.row.status_file">
                  {{ languageMcs == "en" ? "Completed" : '完成'}}
                </el-button>
                <el-button slot="reference" plain type="info" class="statusStyle" v-else>
                  {{scope.row.status}}
                </el-button>
              </el-popover>
            </template>
          </el-table-column>
          <el-table-column prop="pin_status" min-width="100">
            <template slot="header" slot-scope="scope">
              <div class="tips">
                {{$t('uploadFile.status')}}

                <el-popover placement="top" popper-class="elPopTitle" width="200" trigger="hover" :content="$t('uploadFile.status_tooltip')">
                  <img slot="reference" src="@/assets/images/info.png" />
                </el-popover>
              </div>
            </template>
            <template slot-scope="scope">
              <div class="hot-cold-box">
                <el-popover v-if="scope.row.pin_status&&scope.row.pin_status.toLowerCase()=='pinned'" placement="top" popper-class="elPopTitle" class="item" width="200" trigger="hover" :content="$t('uploadFile.status_button_tooltip')">
                  <el-button slot="reference" class="uploadBtn blue" type="primary" @click.stop="pinClick(scope.row)">
                    Unpin
                  </el-button>
                </el-popover>
                <el-button class="uploadBtn grey opacity" v-else :disabled="true">
                  Unpinned
                </el-button>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="pay_amount" min-width="120">
            <template slot="header" slot-scope="scope">
              <div class="tips">
                {{$t('billing.deals_price')}}
              </div>
            </template>
            <template slot-scope="scope">
              <div class="hot-cold-box">
                {{scope.row.pay_amount | balanceMweiFilter}} USDC
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="miner_fid" min-width="120">
            <template slot="header" slot-scope="scope">
              <div class="tips tipsWidth">
                {{$t('uploadFile.w3ss_ids')}}

                <el-popover placement="top" popper-class="elPopTitle" width="200" trigger="hover" :content="$t('uploadFile.w3ss_id_tooltip')">
                  <img slot="reference" src="@/assets/images/info.png" />
                </el-popover>
              </div>
            </template>
            <template slot-scope="scope">
              <div class="hot-cold-box hot-miner">
                <div class="elTips">
                  <el-popover v-for="(miner, i) in scope.row.offline_deal" :key="i" placement="top" trigger="hover" popper-class="elPopMiner" v-model="scope.row.payloadAct">
                    <div class="upload_form_right">
                      <div class="statusStyle" v-if="miner.on_chain_status == 'Created'" :style="$statusColor.TaskColor('Created')">
                        {{ languageMcs == "en" ? "Created" : "已创建" }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'Assigned'" :style="$statusColor.TaskColor('Assigned')">
                        {{ languageMcs == "en" ? "Assigned" : "已分配" }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'Accepted'" :style="$statusColor.TaskColor('Accepted')">
                        {{ languageMcs == "en" ? "Accepted" : "已接受" }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'Completed'" :style="$statusColor.TaskColor('Completed')">
                        {{ languageMcs == "en" ? "Completed" : "已完成" }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'Failed'" :style="$statusColor.TaskColor('Failed')">
                        {{ languageMcs == "en" ? "Failed" : "已失败" }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'Cancelled'" :style="$statusColor.TaskColor('Cancelled')">
                        {{ languageMcs == "en" ? "Cancelled" : "已取消" }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'Closed'" :style="$statusColor.TaskColor('Closed')">
                        {{ languageMcs == "en" ? "Closed" : "已关闭" }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'Expired'" :style="$statusColor.TaskColor('Expired')">
                        {{ languageMcs == "en" ? "Expired" : "已过期" }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'ActionRequired'" :style="$statusColor.TaskColor('ActionRequired')">
                        {{ languageMcs == 'en' ? 'ActionRequired' : '需要操作' }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'DealSent'" :style="$statusColor.TaskColor('DealSent')">
                        {{ languageMcs == 'en' ? 'DealSent' : '交易已发送' }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'FileImporting'" :style="$statusColor.TaskColor('FileImporting')">
                        {{ languageMcs == 'en' ? 'FileImporting' : '文件导入中' }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'FileImported'" :style="$statusColor.TaskColor('FileImported')">
                        {{ languageMcs == 'en' ? 'FileImported' : '文件已导入' }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'ImportFailed'" :style="$statusColor.TaskColor('ImportFailed')">
                        {{ languageMcs == 'en' ? 'ImportFailed' : '导入失败' }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'Downloading'" :style="$statusColor.TaskColor('Downloading')">
                        {{ languageMcs == 'en' ? 'Downloading' : '下载中' }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'DownloadFailed'" :style="$statusColor.TaskColor('DownloadFailed')">
                        {{ languageMcs == 'en' ? 'DownloadFailed' : '下载失败' }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'DealActive'" :style="$statusColor.TaskColor('DealActive')">
                        {{ languageMcs == 'en' ? 'DealActive' : '有效交易' }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'Waiting'" :style="$statusColor.TaskColor('Waiting')">
                        {{ languageMcs == 'en' ? 'Waiting' : '等待中' }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == 'ReadyForImport'" :style="$statusColor.TaskColor('ReadyForImport')">
                        {{ languageMcs == 'en' ? 'ReadyForImport' : '准备导入' }}
                      </div>
                      <div class="statusStyle" v-else-if="miner.on_chain_status == ''">
                        -
                      </div>
                      <div class="statusStyle" v-else>
                        {{ miner.on_chain_status }}
                      </div>
                    </div>
                    <el-button slot="reference" class="resno" @click="minerIdLink(miner.miner_fid)">
                      {{miner.miner_fid}}
                      <small v-if="i<scope.row.offline_deal.length-1">,&nbsp;</small>
                    </el-button>
                  </el-popover>

                  <div class="elPopMiner" v-if="!(scope.row.offline_deal&&scope.row.offline_deal.length>0)" style="display: flex;align-items: center;">
                    {{$t('uploadFile.w3ss_id_nothing')}}

                    <el-popover placement="top" popper-class="elPopTitle" width="200" trigger="hover" :content="$t('uploadFile.w3ss_id_nothing_tooltip')">
                      <img slot="reference" style="display:block" src="@/assets/images/info.png" />
                    </el-popover>
                  </div>
                </div>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="upload_at" :label="$t('uploadFile.upload_time')" width="120" sortable="custom">
            <template slot-scope="scope">
              {{ scope.row.upload_at }}
              <small style="display: block;">({{scope.row.dataUnit}})</small>
            </template>
          </el-table-column>
          <el-table-column prop="active" min-width="100" :label="$t('uploadFile.payment')">
            <template slot-scope="scope">
              <div class="hot-cold-box">
                <el-button class="uploadBtn grey opacity" v-if="tableData[scope.$index].status.toLowerCase()=='pending'&&tableData[scope.$index].pin_status.toLowerCase()=='unpinned'" :disabled="true">
                  {{$t('uploadFile.Canceled')}}
                </el-button>
                <el-button class="uploadBtn blue" type="primary" v-else-if="tableData[scope.$index].status.toLowerCase()=='pending'" @click.stop="payClick(scope.row)">
                  {{$t('uploadFile.pay')}}
                </el-button>
                <el-button v-else-if="tableData[scope.$index].is_free" :disabled="true" class="uploadBtn grey opacity">{{$t('uploadFile.filter_status_Free')}}</el-button>
                <el-button class="uploadBtn grey opacity" type="primary" v-else-if="tableData[scope.$index].status.toLowerCase()=='completed' && !tableData[scope.$index].is_free && tableData[scope.$index].refunded_by_self" :disabled="true">
                  {{$t('uploadFile.refund')}}
                </el-button>
                <el-button v-else-if="tableData[scope.$index].status.toLowerCase()=='completed' && !tableData[scope.$index].is_free && !tableData[scope.$index].refunded_by_self" :disabled="true" class="uploadBtn grey opacity">{{$t('uploadFile.paid')}}</el-button>
                <el-button class="uploadBtn blue" type="primary" v-else-if="tableData[scope.$index].status.toLowerCase()=='refundable'" @click.stop="refundClick(scope.row)">
                  {{$t('uploadFile.refund')}}
                </el-button>
                <el-button v-else-if="tableData[scope.$index].status.toLowerCase()=='failed'" :disabled="true" class="uploadBtn grey opacity">{{$t('uploadFile.failed')}}</el-button>
                <el-button v-else-if="tableData[scope.$index].status.toLowerCase()=='processing'" :disabled="true" class="uploadBtn grey opacity">{{$t('uploadFile.paid')}}</el-button>
                <el-button v-else :disabled="true" class="uploadBtn grey opacity">{{$t('uploadFile.paid')}}</el-button>
              </div>
            </template>
          </el-table-column>
          <el-table-column v-if="networkID != 97" prop="MINT" min-width="100" :label="$t('uploadFile.MINT')" :filters="[{text: $t('uploadFile.filter_no_minted'), value: 'n'}, {text: $t('uploadFile.filter_minted'), value: 'y'}]" :filter-multiple="false"
            :column-key="'minted'">
            <template slot-scope="scope">
              <div class="hot-cold-box">
                <el-button class="uploadBtn blue" type="primary" @click.stop="mintFunction(scope.row, 'mint')">{{$t('uploadFile.MINT')}}</el-button>
                <!-- <el-button class="uploadBtn blue" type="primary" @click.stop="mintFunction(scope.row, 'view')">{{$t('uploadFile.mint_view')}}</el-button> -->
              </div>
            </template>
          </el-table-column>
        </el-table>

        <div class="form_pagination">
          <div class="pagination">
            <el-pagination v-if="parma.total != 0" :total="parma.total" :page-size="parma.limit" :current-page="parma.offset" :layout="'prev, pager, next'" @current-change="handleCurrentChange" @size-change="handleSizeChange" />
            <div class="span" v-if="!bodyWidth">
              <span>{{$t('uploadFile.goTo')}}</span>
              <el-input class="paginaInput" @change="pageSizeChange" v-model.number="parma.jumperOffset" onkeyup="this.value=this.value.replace(/\D/g,'')" onafterpaste="this.value=this.value.replace(/\D/g,'')" autocomplete="off"></el-input>
              <span>{{$t('uploadFile.goTopage')}}</span>
            </div>
          </div>

          <div class="down" @click="downVisible=true">
            [
            <span>{{$t('billing.data_download')}} CSV {{$t('billing.data_export')}}</span> ]
          </div>
        </div>
      </div>

      <div class="loadMetamaskPay" v-if="loadMetamaskPay">
        <div>
          <div class="el-loading-spinner">
            <svg viewBox="25 25 50 50" class="circular">
              <circle cx="50" cy="50" r="20" fill="none" class="path"></circle>
            </svg>
            <!---->
          </div>
          <p>{{$t('uploadFile.payment_tip')}}</p>
        </div>
      </div>
    </div>
    <!-- @getPay="getPay" -->
    <pay-tips v-if="uploadDigShow" :uploadDigShow="uploadDigShow" @getUploadDialog="getUploadDialog"></pay-tips>

    <pay-tip v-if="payVisible" :payVisible="payVisible" :payRow="payRow" @getDialog="getDialog"></pay-tip>

    <el-dialog title="" :visible.sync="finishTransaction" :width="width" :before-close="finishClose" custom-class="completeDia">
      <img src="@/assets/images/alert-icon.png" class="resno" />
      <h1>{{$t('uploadFile.COMPLETED')}}!</h1>
      <h3>{{$t('uploadFile.SUCCESS')}}</h3>
      <a :href="baseAddressURL+'tx/'+txHash" target="_blank">{{txHash}}</a>
      <a class="a-close" @click="finishClose">{{$t('uploadFile.CLOSE')}}</a>
    </el-dialog>

    <el-dialog title="" :visible.sync="failTransaction" :width="width" :before-close="failClose" custom-class="completeDia">
      <img src="@/assets/images/error.png" class="resno" />
      <h1>{{$t('uploadFile.Fail')}}!</h1>
      <h3>{{$t('uploadFile.FailTIP')}}</h3>
      <a :href="baseAddressURL+'tx/'+txHash" target="_blank">{{txHash}}</a>
      <a class="a-close" @click="failClose">{{$t('uploadFile.CLOSE')}}</a>
    </el-dialog>

    <el-dialog title="" :visible.sync="waitTransaction" :width="width" :before-close="failClose" custom-class="completeDia">
      <img src="@/assets/images/waiting.png" class="resno" />
      <h1>{{$t('uploadFile.waiting')}}</h1>
      <h3>{{$t('uploadFile.waitingTIP')}}</h3>
      <a :href="baseAddressURL+'tx/'+txHash" target="_blank">{{txHash}}</a>
      <a class="a-close" @click="failClose">{{$t('uploadFile.CLOSE')}}</a>
    </el-dialog>

    <el-dialog title="Tips" :visible.sync="wrongVisible" :show-close="false" :width="width" custom-class="wrongNet">
      <span>{{$t('uploadFile.until')}}</span>
    </el-dialog>

    <el-dialog title="" :visible.sync="metamaskLoginTip" :width="width" custom-class="completeDia">
      <img src="@/assets/images/box-important.png" class="resno" />
      <h4 v-if="baseNetwork">
        {{$t('fs3Login.toptip_01')}} {{metaNetworkInfo.name}} {{$t('fs3Login.toptip_02')}}
        <b>Polygon Mainnet</b>.
      </h4>
      <h4 v-else>
        {{$t('fs3Login.toptip_01')}} {{metaNetworkInfo.name}} {{$t('fs3Login.toptip_02')}}
        <b>Mumbai Testnet</b>
        {{$t('fs3Login.toptip_Network')}}
        <b>BSC TestNet</b>.
      </h4>
      <a class="a-close" @click="metamaskLoginTip=false">{{$t('uploadFile.OK')}}</a>
    </el-dialog>

    <el-dialog title="" :visible.sync="mintTransaction" :width="width" custom-class="completeDia">
      <img src="@/assets/images/alert-icon.png" class="resno" />
      <h1>{{$t('uploadFile.View_Your_NFT')}}</h1>
      <h3>{{$t('uploadFile.View_Your_NFT_tips')}}</h3>
      <a v-if="networkID === 137" :href="'https://opensea.io/assets/matic/'+mint_address+'/'+tokenId" target="_blank">{{$t('uploadFile.View_Your_NFT_OpenSea')}}</a>
      <a v-else :href="'https://testnets.opensea.io/assets/mumbai/'+mint_address+'/'+tokenId" target="_blank">{{$t('uploadFile.View_Your_NFT_OpenSea')}}</a>
      <h3>{{$t('uploadFile.View_Your_NFT_tips01')}}</h3>
      <a :href="baseAddressURL+'tx/'+txHash" target="_blank">{{txHash}}</a>
      <br />
      <h3>{{$t('uploadFile.View_Your_NFT_Note')}}</h3>
      <a class="a-close" @click="failClose">{{$t('uploadFile.CLOSE')}}</a>
    </el-dialog>

    <mint-tip v-if="mineVisible" :mineVisible="mineVisible" :mintRow="mintRow" @getMintDialog="getMintDialog"></mint-tip>
    <download v-if="downVisible" :downVisible="downVisible" :titlePage="$t('billing.download_module_title_file')" @getDownload="getDownload"></download>
    <!-- 回到顶部 -->
    <el-backtop target=".content-box" :bottom="40" :right="20"></el-backtop>
  </div>
</template>

<script>
import axios from 'axios'
import QS from 'qs'
import moment from 'moment'
import payTip from '@/components/payTip'
import payTips from '@/components/uploadFiles'
import mintTip from '@/components/mintTip'
import download from '@/components/download'
import firstContractJson from '@/utils/swanPayment.json'
import erc20ContractJson from '@/utils/ERC20.json'
let contractErc20
export default {
  name: 'uploadFiles',
  data () {
    return {
      tableData: [],
      tableDataChild: [],
      tableDataAll: [],
      parma: {
        limit: 10,
        offset: 1,
        locationValue: '',
        total: 0,
        jumperOffset: 1,
        order_by: '',
        is_ascend: '',
        status: '',
        is_minted: ''
      },
      parmaChild: {
        limit: 10,
        offset: 1,
        locationValue: '',
        total: 0
      },
      searchValue: '',
      loading: false,
      loadingChild: false,
      bodyWidth: document.documentElement.clientWidth < 1024,
      expands: [],
      getRowKeys: (row) => {
        return row.uuid
      },
      payVisible: false,
      uploadDigShow: false,
      mineVisible: false,
      mintRow: {},
      mintContractAddress: this.$root.MINT_CONTRACT,
      tokenId: '',
      mint_address: '',
      toAddress: '',
      storage: 0,
      centerDialogVisible: false,
      center_fail: false,
      width: document.body.clientWidth > 1600 ? '500px' : document.body.clientWidth > 600 ? '400px' : '95%',
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
      waitTransaction: false,
      loadMetamaskPay: false,
      payRow: {},
      refundPow: {},
      resData: {},
      biling_price: 0,
      metamaskLoginTip: false,
      searchNowCurrent: '',
      downVisible: false,
      firstIndex: 0
    }
  },
  computed: {
    languageMcs () {
      return this.$store.state.app.languageMcs
    },
    metaAddress () {
      return this.$store.getters.metaAddress
    },
    networkID () {
      return Number(this.$store.getters.networkID)
    },
    metaNetworkInfo () {
      return this.$store.getters.metaNetworkInfo ? JSON.parse(JSON.stringify(this.$store.getters.metaNetworkInfo)) : {}
    }
  },
  components: {
    payTip,
    mintTip,
    payTips,
    download
  },
  watch: {
    'searchValue': function () {
      let _this = this
      _this.parma.limit = 10
      _this.parma.offset = 1
      _this.parma.jumperOffset = 1
      _this.parma.order_by = ''
      _this.parma.is_ascend = ''
      _this.parma.status = ''
      _this.parma.is_minted = ''
      _this.parmaChild.limit = 10
      _this.parmaChild.offset = 1
      _this.$refs.singleTable.clearFilter()

      _this.searchNowCurrent = (new Date()).getTime()
      _this.getData(_this.searchNowCurrent)
    }
  },
  created () {
    if (this.$route.query.status !== undefined) {
      this.parma.limit = this.$route.query.limit
      this.parma.searchValue = this.$route.query.task_name
    }
  },
  methods: {
    minerIdLink (id) {
      if (!id && String(id) !== '0') return false
      let link = id.slice(0, 1)
      switch (link) {
        case 'f':
          window.open(`https://filfox.info/en/address/${id}`)
          return false
        default:
          window.open(`https://calibration.filscan.io/address/miner?address=${id}`)
          break
      }
    },
    toDetail (row) {
      if (row.offline_deal && row.offline_deal.length > 0) {
        this.$router.push({ name: 'my_files_detail', params: { id: row.offline_deal[0].id, deal_id: row.offline_deal[0].deal_id, cid: row.payload_cid, source_file_upload_id: row.source_file_upload_id, isFree: row.is_free ? 1 : 0 } })
      } else {
        this.$router.push({ name: 'my_files_detail', params: { id: 0, deal_id: 0, cid: row.payload_cid, source_file_upload_id: row.source_file_upload_id, isFree: row.is_free ? 1 : 0 } })
      }
      localStorage.setItem('offlineDeals', row.offline_deal ? JSON.stringify(row.offline_deal) : [])
      localStorage.setItem('offlineDealsIndex', '0')
      sessionStorage.setItem('dealsPaginationIndexDev', this.parma.offset)
    },
    async payClick (row) {
      let _this = this
      let status = await _this.$metaLogin.netStatus(_this.networkID)
      if (_this.metaAddress && !status) {
        _this.metamaskLoginTip = true
        return false
      }
      // console.log(row)
      _this.payRow = row
      _this.payRow.file_size = row.file_size
      _this.payVisible = true
      return false
    },
    async pinClick (row) {
      let _this = this
      _this.loading = true
      const dataRes = await _this.sendPostRequest(`${_this.baseAPIURL}api/v1/storage/unpin_source_file/${row.source_file_upload_id}`)
      if (!dataRes || dataRes.status !== 'success') {
        _this.loading = false
        _this.$message.error(dataRes ? dataRes.message : 'Fail')
        return false
      }
      await _this.timeout(1000)
      _this.getData()
    },
    timeout (delay) {
      return new Promise((resolve) => setTimeout(resolve, delay))
    },
    mintFunction (row, type) {
      row.mintType = type
      this.mintRow = row
      this.mineVisible = true
    },
    mintViewFunction (row) {
      this.tokenId = row.token_id
      this.mint_address = row.mint_address
      this.txHash = row.nft_tx_hash
      this.mintTransaction = true
    },
    refundClick (row) {
      let _this = this
      let contractInstance = new _this.$web3Init.eth.Contract(firstContractJson)
      contractInstance.options.address = _this.gatewayContractAddress
      _this.loading = true

      let payObject = {
        from: _this.metaAddress,
        gas: _this.$web3Init.utils.toHex(_this.$root.PAY_GAS_LIMIT)
      }

      let wcidApi = `${_this.baseAPIURL}api/v1/storage/source_file_upload/${row.source_file_upload_id}?wallet_address=${_this.metaAddress}`
      axios.get(wcidApi, {
        headers: {
          'Authorization': 'Bearer ' + _this.$store.getters.mcsjwtToken
        }
      })
        .then((json) => {
          if (json.data.status === 'success') {
            let cidArray = []
            cidArray.push(json.data.data.source_file_upload.w_cid)
            try {
              _this.refundPow = row
              contractInstance.methods.refund(cidArray)
                .send(payObject)
                .on('transactionHash', function (hash) {
                  console.log('unlock hash console:', hash)
                  _this.txHash = hash
                })
                .on('confirmation', function (confirmationNumber, receipt) {
                  // console.log('confirmationNumber console:', confirmationNumber, receipt);
                })
                .on('receipt', function (receipt) {
                  // console.log('receipt console:', receipt);
                  _this.checkTransaction(receipt.transactionHash)
                  _this.txHash = receipt.transactionHash
                })
                .on('error', function (error) {
                  console.log('refund error console:', error)
                  _this.loading = false
                  if (_this.finishTransaction) return false
                  if (_this.txHash) _this.waitTransaction = true
                  else _this.failTransaction = true
                })
            } catch (err) {
              console.log(err)
              _this.loading = false
            }
          }
        }).catch(error => {
          console.log(error)
        })
    },
    payStartClick (rowAmount) {
      let _this = this
      if (_this.metaAddress) {
        _this.loading = true
        // 发起请求
        axios.get(`${_this.baseAPIURL}api/v1/billing/deal/lockpayment/info?payload_cid=${_this.payRow.payload_cid}&source_file_upload_id=${_this.payRow.source_file_upload_id}&wallet_address=${_this.metaAddress}`, {
          headers: {
            'Authorization': 'Bearer ' + _this.$store.getters.mcsjwtToken
          }
        })
          .then((res) => {
            if (res.data.status === 'success') {
              if (res.data.data.pay_tx_hash) {
                _this.$message.error(_this.$t('deal.file_error'))
                _this.loading = false
                _this.firstIndex = 0
                _this.getData()
                return false
              } else {
                // 授权代币
                contractErc20 = new _this.$web3Init.eth.Contract(erc20ContractJson)
                contractErc20.options.address = _this.usdcAddress
                // 查询剩余代币余额为：
                contractErc20.methods.balanceOf(_this.metaAddress).call()
                  .then(balance => {
                    let usdcAvailable = _this.$web3Init.utils.fromWei(balance, 'mwei')
                    console.log('Available pay:', usdcAvailable, rowAmount)
                    // 判断支付金额是否大于代币余额
                    if (Number(rowAmount) > Number(usdcAvailable)) {
                      _this.$message.error('Insufficient balance')
                      return false
                    } else {
                      contractErc20.methods.allowance(_this.gatewayContractAddress, _this.metaAddress).call()
                        .then(async resultUSDC => {
                          let amountPay = _this.$web3Init.utils.toWei(rowAmount, 'mwei')
                          console.log('allowance：' + resultUSDC, amountPay)
                          if (resultUSDC < amountPay) {
                            let payObject = {
                              from: _this.metaAddress,
                              gas: _this.$web3Init.utils.toHex(_this.$root.PAY_GAS_LIMIT),
                              gasPrice: await _this.$web3Init.eth.getGasPrice()
                            }
                            contractErc20.methods.approve(_this.gatewayContractAddress, amountPay).send(payObject)
                              .then(receipt => {
                                // console.log('approve receipt:', receipt)
                                _this.contractSend(res.data.data.w_cid, amountPay)
                              })
                              .catch(() => {
                                // console.log('errorerrorerror', error)
                                _this.finishClose()
                              })
                          } else {
                            _this.contractSend(res.data.data.w_cid, amountPay)
                          }
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
    async contractSend (cid, payAmount) {
      let _this = this
      // 合约转账
      let contractInstance = new _this.$web3Init.eth.Contract(firstContractJson)
      contractInstance.options.address = _this.gatewayContractAddress

      let payObject = {
        from: _this.metaAddress,
        gas: _this.$web3Init.utils.toHex(_this.$root.PAY_GAS_LIMIT),
        gasPrice: await _this.$web3Init.eth.getGasPrice()
      }

      let lockObj = {
        id: cid,
        minPayment: String(Math.floor(payAmount / 2)),
        amount: payAmount,
        lockTime: 86400 * Number(_this.$root.LOCK_TIME), // one day
        recipient: _this.recipientAddress, // todo:
        size: _this.payRow.file_size,
        copyLimit: 5
      }
      console.log(lockObj)
      contractInstance.methods.lockTokenPayment(lockObj)
        .send(payObject)
        .on('transactionHash', function (hash) {
          console.log('hash console:', hash)
          _this.loadMetamaskPay = true
          _this.loading = false
          _this.txHash = hash
        })
        .on('confirmation', function (confirmationNumber, receipt) {
          // console.log('confirmationNumber console:', confirmationNumber, receipt);
        })
        .on('receipt', function (receipt) {
          // receipt example
          // console.log('receipt console:', receipt);
          _this.checkTransaction(receipt.transactionHash, cid, _this.payRow, lockObj)
          _this.txHash = receipt.transactionHash
        })
        .on('error', function (error) {
          console.log('lockTokenPayment error console:', error)
          // console.error
          _this.loading = false
          _this.loadMetamaskPay = false
          if (_this.finishTransaction) return false
          if (_this.txHash) _this.waitTransaction = true
          else _this.failTransaction = true
        })
    },
    checkTransaction (txHash, cid, resData, lockObj) {
      let _this = this
      _this.$web3Init.eth.getTransactionReceipt(txHash).then(
        async res => {
          console.log('checking ... ')
          if (!res) {
            _this.timer = setTimeout(() => { _this.checkTransaction(txHash, cid, resData, lockObj) }, 2000)
            return false
          } else {
            clearTimeout(_this.timer)
            setTimeout(function () {
              _this.loading = false
              _this.loadMetamaskPay = false
              _this.finishTransaction = true
            }, 2000)
          }
        },
        err => { console.error(err) }
      )
    },
    async sendPostRequest (apilink, jsonObject) {
      let _this = this
      try {
        const response = await axios.post(apilink, jsonObject, {
          headers: {
            'Authorization': 'Bearer ' + _this.$store.getters.mcsjwtToken
          }
        })
        return response.data
      } catch (err) {
        console.error(err)
      }
    },
    finishClose () {
      this.finishTransaction = false
      this.mintTransaction = false
      this.firstIndex = 0
      this.txHash = ''
      this.getData()
    },
    failClose () {
      this.failTransaction = false
      this.waitTransaction = false
      this.mintTransaction = false
      this.txHash = ''
    },
    getDialog (dialog, rows) {
      this.payVisible = dialog
      if (rows) this.payStartClick(rows)
    },
    getUploadDialog (dialog, rows) {
      this.uploadDigShow = dialog
      if (rows) {
        this.firstIndex = 0
        this.parma.offset = 1
        this.getData()
      }
    },
    getDownload (dialog, rows) {
      this.downVisible = dialog
    },
    getMintDialog (dialog, mintInfoJson) {
      let _this = this
      _this.mineVisible = dialog
      if (mintInfoJson) {
        _this.firstIndex = 0
        _this.getData()
        _this.tokenId = mintInfoJson.token_id
        _this.txHash = mintInfoJson.tx_hash || mintInfoJson.nft_tx_hash
        _this.mint_address = mintInfoJson.mint_address || mintInfoJson.address || mintInfoJson.nft_collection_address
        _this.mintTransaction = true
      }
    },
    copyTextToClipboard (text) {
      let _this = this
      let saveLang = localStorage.getItem('languageMcs') === 'cn' ? '复制成功' : 'success'
      var txtArea = document.createElement('textarea')
      txtArea.id = 'txt'
      txtArea.style.position = 'fixed'
      txtArea.style.top = '0'
      txtArea.style.left = '0'
      txtArea.style.opacity = '0'
      txtArea.value = text
      document.body.appendChild(txtArea)
      txtArea.select()

      try {
        var successful = document.execCommand('copy')
        var msg = successful ? 'successful' : 'unsuccessful'
        console.log('Copying text command was ' + msg)
        if (successful) {
          _this.$message({
            message: saveLang,
            type: 'success'
          })
          return true
        }
      } catch (err) {
        console.log('Oops, unable to copy')
      } finally {
        document.body.removeChild(txtArea)
      }
      return false
    },
    payFun (cid, type) {
      let _this = this
      let payApi = `${_this.baseAPIURL}api/v1/events/logs/lock/${cid}?wallet_address=${_this.metaAddress}`

      axios.get(payApi, {
        headers: {
          'Authorization': 'Bearer ' + _this.$store.getters.mcsjwtToken
        }
      })
        .then((json) => {
          if (json.data.status === 'success') {
            if (type) {
              if (json.data.data.polygon.length > 0 || json.data.data.goerli.length > 0) {
                clearTimeout(_this.timer)
                _this.sendDeal()
                _this.timeIndex = 0
              }
            }
            _this.paySuccess = !!(json.data.data.polygon.length > 0 || json.data.data.goerli.length > 0)
          } else {
            _this.paySuccess = true
          }
        }).catch(error => {
          console.log(error)
        })
    },
    getPay (status, cid) {
      if (!status) return false
      let _this = this
      _this.timeIndex = 0
      _this.timer = setInterval(() => {
        _this.timeIndex += 1
        _this.payFun(cid, 1)
        if (_this.timeIndex > 9) clearTimeout(_this.timer)
      }, 5000)
    },
    sendDeal () {
      let _this = this
      let sendDealApi = `${_this.baseAPIURL}api/v1/storage/lotus/deal/${_this.expands[0]}?wallet_address=${_this.metaAddress}`
      axios.get(sendDealApi, {
        headers: {
          'Authorization': 'Bearer ' + _this.$store.getters.mcsjwtToken
        }
      }).then(res => {
        if (res.data.status === 'success') {
          _this.$message({
            message: localStorage.getItem('languageMcs') === 'en' ? 'Send Deal Success' : '发送交易成功',
            type: 'success'
          })
          _this.getData()
        } else {
          _this.$message.error(localStorage.getItem('languageMcs') === 'en' ? 'Send Deal Fail' : '发送交易失败')
        }
      }).catch(error => {
        console.log(error)
      })
    },
    // 查询
    search () {
      let _this = this
      _this.parma.limit = 10
      _this.paginationShow = true
      _this.parma.offset = 1
      _this.parma.jumperOffset = 1
      _this.parma.order_by = ''
      _this.parma.is_ascend = ''
      _this.parma.status = ''
      _this.parma.is_minted = ''
      _this.$refs.singleTable.clearSort()
      _this.$refs.singleTable.clearFilter()
      _this.getData()
    },
    clearAll () {
      let _this = this
      _this.searchValue = ''
      _this.parma.limit = 10
      _this.parma.offset = 1
      _this.parma.jumperOffset = 1
      _this.parma.order_by = ''
      _this.parma.is_ascend = ''
      _this.parma.status = ''
      _this.parma.is_minted = ''
      _this.$refs.singleTable.clearSort()
      _this.$refs.singleTable.clearFilter()
      _this.getData()
    },
    handleCurrentChange (val) {
      this.parma.offset = Number(val)
      this.parma.jumperOffset = String(val)
      this.getData()
    },
    handleSizeChange (val) {
      this.parma.limit = Number(val)
      this.parma.offset = 1
      this.parma.jumperOffset = 1
      this.getData()
    },
    pageSizeChange (recordPage = parseInt(this.parma.jumperOffset), MaxPagenumber = Math.ceil(this.parma.total / this.parma.limit)) {
      if ((recordPage > MaxPagenumber) && (MaxPagenumber > 0)) {
        recordPage = MaxPagenumber
      } else if (MaxPagenumber <= 0) {
        recordPage = 1
      } else if (recordPage < 1) {
        recordPage = 1
      } else if (isNaN(this.parma.jumperOffset) || this.parma.jumperOffset === '') {
        recordPage = 1
      }
      this.parma.offset = Number(recordPage)
      this.parma.jumperOffset = recordPage
      this.getData()
    },
    async stats () {
      let _this = this
      _this.loading = true
      if (_this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS) {
        _this.gatewayContractAddress = _this.$root.SWAN_PAYMENT_CONTRACT_ADDRESS
        _this.usdcAddress = _this.$root.USDC_ADDRESS
        _this.recipientAddress = _this.$root.RECIPIENT
        _this.mintContractAddress = _this.$root.MINT_CONTRACT

        _this.getData()
      } else {
        setTimeout(function () {
          _this.stats()
        }, 1000)
      }
    },
    async sortOrderBy (sort) {
      switch (sort) {
        case 'file_name':
          return 2
        case 'file_size':
          return 3
        case 'status':
          return 13
        case 'pin_status':
          return 4
        case 'payload_cid':
          return 6
        case 'upload_at':
          return 5
        default:
      }
    },
    async sortChange (column) {
      // console.log(column);
      await this.filterChange({})
      // this.parma.order_by = await this.sortOrderBy(column.prop)
      this.parma.order_by = column.prop
      this.parma.is_ascend = column.order === 'ascending' ? 'y' : column.order === 'descending' ? 'n' : ''
      this.loading = true
      this.getData()
    },
    filterChange (filters) {
      if (('payment' in filters)) {
        let data = filters.payment[0] || ''
        if (data === this.parma.status) return false
        this.parma.status = data
        this.$refs.singleTable.clearFilter('minted')
        this.parma.is_minted = ''
      } else if (('minted' in filters)) {
        let data = filters.minted[0] || 'all'
        if (data === this.parma.is_minted) return false
        this.parma.is_minted = data
        this.$refs.singleTable.clearFilter('payment')
        this.parma.status = ''
      } else {
        this.$refs.singleTable.clearFilter()
        this.parma.status = ''
        this.parma.is_minted = ''
        return false
      }
      this.$refs.singleTable.clearSort()
      this.parma.is_ascend = ''
      this.getData()
    },
    getData (currentData) {
      let _this = this
      _this.loading = true
      let parma = {
        page_size: _this.parma.limit,
        page_number: _this.parma.offset,
        file_name: _this.searchValue,
        wallet_address: _this.metaAddress,
        order_by: _this.parma.is_ascend ? _this.parma.order_by : '',
        is_ascend: _this.parma.is_ascend,
        status: _this.parma.status,
        is_minted: _this.parma.is_minted
      }

      _this.tableData = []

      let uploadRes = new Promise((resolve, reject) => {
        let storageApi = `${_this.baseAPIURL}api/v1/storage/tasks/deals?${QS.stringify(parma)}`
        // let storageApi = `./static/pay-status-response.json?${QS.stringify(parma)}`

        axios.get(storageApi, {
          headers: {
            'Authorization': 'Bearer ' + _this.$store.getters.mcsjwtToken
          }
        })
          .then((response) => {
            if (response.data.status === 'success') {
              if (currentData && _this.searchNowCurrent !== currentData) return false
              _this.$store.dispatch('setFreeUsage', response.data.data.free_usage || 0)
              _this.$store.dispatch('setFreeQuote', response.data.data.free_quota_per_month || 0)
              _this.parma.total = Number(response.data.data.total_record_count)
              _this.tableData = response.data.data.source_file_upload ? response.data.data.source_file_upload : []
              _this.tableData.map((item, s) => {
                item.payloadAct = false
                item.status_file = false
                item.status_failed_file = true
                item.status_success_file = true
                item.file_size_byte = _this.byteChange(item.file_size)

                let dataTime = new Date(item.upload_at * 1000) + '' // 将时间格式转为字符串
                let dataUnitIndex = dataTime.indexOf('GMT')
                let dataUnitArray = dataTime.substring(dataUnitIndex, dataUnitIndex + 8)
                switch (dataUnitArray) {
                  case 'GMT+1000':
                    item.dataUnit = 'UTC+10'
                    break
                  case 'GMT-1000':
                    item.dataUnit = 'UTC-10'
                    break
                  case 'GMT+0000':
                    item.dataUnit = 'UTC+0'
                    break
                  default:
                    item.dataUnit = dataUnitArray ? dataUnitArray.replace(/0/g, '').replace('GMT', 'UTC') : '-'
                    break
                }

                item.upload_at = item.upload_at
                  ? moment(new Date(parseInt(item.upload_at * 1000))).format('YYYY-MM-DD HH:mm:ss')
                  : '-'

                if (item.offline_deal) {
                  item.offline_deal.map(child => {
                    if (child.status !== 'Failed') item.status_failed_file = false
                    if (child.status !== 'Success') item.status_success_file = false
                    if (child.status === 'Success') item.status_file = true
                  })
                } else {
                  item.status_success_file = false
                }
              })
              setTimeout(function () {
                _this.loading = false
                _this.firstIndex = 1
                // resolve('')
              }, 2000)
            } else {
              _this.$message.error(response.data.message)
              _this.loading = false
              _this.firstIndex = 1
              reject(response.data.message)
            }
          }).catch(error => {
            console.log(error)
            _this.loading = false
            _this.firstIndex = 1
            reject(error)
          })
      })
      Promise.all([uploadRes]).then((result) => {
        console.log(result)
      }).catch((error) => {
        console.log(error)
      })
      return false
    },
    byteChange (limit) {
      var size = ''
      // 只转换成GB
      if (limit <= 0) {
        return '-'
      } else {
        size = limit / (1024 * 1024 * 1024) // or 1000
      }
      return size
      // return Number(size).toFixed(3);
    },
    unique (arr) {
      const res = new Map()
      return arr.filter((arr) => !res.has(arr.value) && res.set(arr.value, 1))
    }
  },
  mounted () {
    let _this = this
    document.getElementById('content-box').scrollTop = 0
    _this.$store.dispatch('setRouterMenu', 1)
    _this.$store.dispatch('setHeadertitle', _this.$t('route.Deal'))
    if (sessionStorage.getItem('dealsPaginationIndexDev')) _this.parma.offset = Number(sessionStorage.getItem('dealsPaginationIndexDev'))

    _this.stats()
  },
  filters: {
    NumFormat (value) {
      if (!value) return '-'
      return value
    },
    NumStorage (value) {
      if (!value) return '- FIL'
      return value.toFixed(8) + ' FIL'
    },
    formatbytes: function (bytes) {
      if (bytes === 0) return '0 B'
      if (!bytes) return '-'
      var k = 1024 // or 1000
      var sizes = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
      var i = Math.floor(Math.log(bytes) / Math.log(k))

      if (Math.round((bytes / Math.pow(k, i))).toString().length > 3) {
        // 判断大小是999999999左右，解决会显示成1.00e+3科学计数法
        i += 1
      }
      return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i]
    },
    balanceFilter (value) {
      if (String(value) === '0') return 0
      if (!value) return '-'
      // if (!Number(value)) return 0;
      // if (isNaN(value)) return value;
      // 18 - 单位换算需要 / 1000000000000000000，浮点运算显示有bug
      // value = Number(value)
      if (String(value).length > 18) {
        let v1 = String(value).substring(0, String(value).length - 18)
        let v2 = String(value).substring(String(value).length - 18)
        let v3 = String(v2).replace(/(0+)\b/gi, '')
        if (v3) {
          return v1 + '.' + v3
        } else {
          return v1
        }
      } else {
        let v3 = ''
        for (let i = 0; i < 18 - String(value).length; i++) {
          v3 += '0'
        }
        return '0.' + String(v3 + value).replace(/(0+)\b/gi, '')
      }
    },
    balanceMweiFilter (value) {
      if (String(value) === '0') return 0
      if (!value) return '-'
      if (String(value).length > 6) {
        let v1 = String(value).substring(0, String(value).length - 6)
        let v2 = String(value).substring(String(value).length - 6)
        let v3 = String(v2).replace(/(0+)\b/gi, '')
        if (v3) {
          return v1 + '.' + v3
        } else {
          return v1
        }
      } else {
        let v3 = ''
        for (let i = 0; i < 6 - String(value).length; i++) {
          v3 += '0'
        }
        return '0.' + String(v3 + value).replace(/(0+)\b/gi, '')
      }
    }
  }
}
</script>

<style scoped lang="scss">
#dealManagement {
  position: relative;
  padding: 0.3rem;
  .el-alert /deep/ {
    position: absolute;
    left: 0;
    top: 0;
    .el-alert__content {
      display: flex;
      align-items: center;
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
    padding: 0.3rem;
    background-color: #fff;
    border-radius: 0.1rem;

    .loadMetamaskPay {
      position: absolute;
      left: 0;
      top: 0;
      right: 0;
      bottom: 0;
      z-index: 99;
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: rgba(255, 255, 255, 0.9);
      .el-loading-spinner {
        top: 0;
        position: relative;
        margin: 0 0 0.2rem;
      }
      p {
        font-size: 14px;
        font-weight: 600;
        color: #666;
      }
    }
    .form_top {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      margin: 0 auto 0.3rem;
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

      .upload_title {
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
      }
      .search_file {
        display: flex;
        align-items: center;
        justify-content: space-between;
        width: 100%;
        margin: 0;
        p {
          font-size: 0.13rem;
          color: #222;
        }
        .createTask {
          background-color: #fff;
          border-radius: 0.1rem;
          cursor: pointer;
          a {
            display: flex;
            align-items: center;
            justify-content: center;
            width: 2rem;
            padding: 0.13rem 0;
            margin: 0;
            background: linear-gradient(45deg, #4e88ff, #4b5fff);
            border-radius: 0.14rem;
            line-height: 1.5;
            text-align: center;
            color: #fff;
            font-size: 0.19rem;
            border: 0;
            outline: none;
            transition: background-color 0.3s, border-color 0.3s, color 0.3s,
              box-shadow 0.3s;
            cursor: pointer;
            img {
              display: inline-block;
              height: 0.25rem;
              margin: 0 0.1rem 0 0;
            }
            &:hover {
              opacity: 0.9;
              box-shadow: 0 12px 12px -12px rgba(12, 22, 44, 0.32);
            }
          }
        }
        .search_right {
          display: flex;
          align-items: center;
          width: 100%;
          margin-right: 0.9rem;
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
          padding: 0 0.5rem 0 0.7rem;
          background: #f7f7f7;
          border-radius: 0.2rem;
          .el-input__inner {
            width: 100%;
            color: #555;
            font-size: 0.19rem;
            font-weight: 500;
            height: 0.54rem;
            line-height: 0.3rem;
            padding: 0;
            background: transparent;
            border: 0;
            border-radius: 0.2rem;
          }
          .el-input__prefix {
            left: 0.3rem;
            i {
              display: flex;
              align-items: center;
              font-size: 0.25rem;
              color: #000;
            }
          }
          .el-input__icon {
            line-height: 0.3rem;
          }
          .el-input__suffix {
            right: 0.2rem;
            .el-input__clear {
              display: flex;
              align-items: center;
              font-size: 0.25rem;
              color: #666;
            }
          }
          ::-webkit-input-placeholder {
            color: #555;
          } /* 使用webkit内核的浏览器 */
          :-moz-placeholder {
            color: #555;
          } /* Firefox版本4-18 */
          ::-moz-placeholder {
            color: #555;
          } /* Firefox版本19+ */
          :-ms-input-placeholder {
            color: #555;
          } /* IE浏览器 */
        }
        .is-disabled {
          opacity: 0.2;
        }
      }
    }

    .datasetStyle {
      padding: 0 0 0 20px;
      margin: 0.1rem 0 0;
      background: url(../../../assets/images/download.png) no-repeat left center;
      background-size: 15px;
      font-size: 0.14rem;
      span {
        color: rgb(10, 49, 142);
        cursor: pointer;
      }
    }

    .form_table {
      position: relative;
      margin: 0;
      overflow: hidden;
      .statusStyle /deep/ {
        display: inline-block;
        background: transparent;
        border: 1px solid;
        padding: 0.07rem 0.13rem;
        border-radius: 0.14rem;
        line-height: 1.5;
        font-size: inherit;
        white-space: nowrap;
        // color: inherit !important;
        span {
          white-space: nowrap;
        }
      }
      .el-button--primary,
      .el-button--primary:hover {
        color: #409eff;
      }
      .el-button--success,
      .el-button--success:hover {
        color: #67c23a;
      }
      .el-button--info,
      .el-button--info:hover {
        color: #909399;
      }
      .el-button--warning,
      .el-button--warning:hover {
        color: #e6a23c;
      }
      .el-button--danger,
      .el-button--danger:hover {
        color: #f78989;
      }
      .el-button--refunding {
        color: #f9d54b;
        border-color: #f9d54b;
      }
      .el-button--successPart {
        color: #67c23a;
        border-color: #f9d54b;
      }
      .el-button--pending {
        color: #f2a942;
        border-color: #f2a942;
      }

      .el-table /deep/ {
        overflow: visible;
        font-size: 0.18rem;
        // overflow-x: scroll;
        .el-loading-mask {
          .el-loading-spinner {
            top: 50%;
          }
        }
        .el-table__body-wrapper,
        .el-table__header-wrapper {
          // overflow: visible;
          border-radius: 0.2rem;
        }

        tr {
          cursor: pointer;
          th {
            height: 0.7rem;
            padding: 0;
            background-color: #e5eeff !important;
            text-align: center;

            .cell {
              display: flex;
              align-items: center;
              justify-content: center;
              word-break: break-word;
              font-size: 0.19rem;
              font-weight: 500;
              color: #555;
              text-transform: capitalize;
              line-height: 1.3;
              .caret-wrapper {
                // display: none;
                width: 10px;
                margin-left: 5px;
                .sort-caret {
                  left: 0;
                }
              }
              .tips {
                display: flex;
                align-items: center;
                justify-content: center;
                img {
                  display: block;
                  width: 20px;
                  height: 20px;
                  margin: 0 0 0 5px;
                  cursor: pointer;
                  @media screen and (max-width: 1440px) {
                    width: 17px;
                    height: 17px;
                  }
                  @media screen and (max-width: 600px) {
                    width: 15px;
                    height: 15px;
                  }
                }
              }
              .tipsWidth {
                width: 110px;
                @media screen and (max-width: 1600px) {
                  width: 95px;
                }
                @media screen and (max-width: 1440px) {
                  width: 90px;
                }
                @media screen and (max-width: 1280px) {
                  width: 80px;
                }
              }
              .el-table__column-filter-trigger {
                i {
                  font-size: 13px;
                  font-weight: 600;
                  margin-left: 4px;
                  transform: scale(1);
                }
              }
            }
          }

          // .descending{
          //     position: relative;
          //     &::after{
          //       content: '';
          //       position: absolute;
          //       bottom: 0;
          //       left: 0;
          //       right: 0;
          //       height: 3px;
          //       background-color: #0a318e;
          //     }
          // }

          // .ascending{
          //     position: relative;
          //     &::after{
          //       content: '';
          //       position: absolute;
          //       top: 0;
          //       left: 0;
          //       right: 0;
          //       height: 3px;
          //       background-color: #0a318e;
          //     }
          // }

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
            border-bottom: 1px solid #dfdfdf;

            .cell {
              padding: 0;
              font-size: 0.17rem;
              word-break: break-word;
              color: #000;
              text-align: center;
              line-height: 0.25rem;
              overflow: visible;

              .el-rate__icon {
                font-size: 0.16rem;
                margin-right: 0;
              }
              .el-icon-arrow-right {
                font-weight: bold;
                font-size: 0.13rem;
              }
              .rightClick {
                color: #0c3090;
                cursor: pointer;
              }

              .hot-cold-box {
                position: relative;
                .elTips {
                  display: flex;
                  align-items: center;
                  justify-content: center;
                  flex-wrap: wrap;
                  padding: 0 0.1rem;
                  line-height: 1.5;
                  text-align: center;
                  // white-space: normal;
                  color: inherit;
                  text-shadow: 0 0 black;
                  text-indent: 0;
                  font-size: inherit;
                  font-weight: normal;
                  text-decoration: underline;
                  img {
                    width: 20px;
                    height: 20px;
                    margin: 0 0 0 5px;
                    cursor: pointer;
                    @media screen and (max-width: 1440px) {
                      width: 17px;
                      height: 17px;
                    }
                    @media screen and (max-width: 600px) {
                      width: 15px;
                      height: 15px;
                    }
                  }
                  .el-button {
                    img {
                      display: inline-block;
                      // float: right;
                      width: 0.17rem;
                      margin-top: 0.03rem;
                    }
                  }
                  .el-button:hover {
                    img {
                      display: inline-block;
                    }
                  }
                }
                .cidLink {
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
                .cidMore {
                  position: absolute;
                  top: 0;
                  right: 0;
                  bottom: 0;
                  background-color: #fff;
                  width: 0.22rem;
                  padding: 0;
                  &::before {
                    content: "···";
                    font-weight: bold;
                    color: #0a318e;
                  }
                }
                .el-button {
                  width: 100%;
                  border: 0;
                  padding: 0;
                  background-color: transparent;
                  font-size: 0.165rem;
                  font-family: inherit;
                  word-break: break-word;
                  color: #000;
                  text-align: center;
                  line-height: 1.5;
                  overflow: hidden;
                  text-overflow: ellipsis;
                  white-space: normal;
                  display: -webkit-box;
                  -webkit-line-clamp: 2;
                  -webkit-box-orient: vertical;
                  span {
                    line-height: 1.5;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: normal;
                    display: -webkit-box;
                    -webkit-line-clamp: 2;
                    -webkit-box-orient: vertical;
                  }
                  img {
                    display: none;
                    float: left;
                    width: 0.17rem;
                    margin-top: 0.03rem;
                  }
                }
                .el-button:hover {
                  img {
                    display: inline-block;
                  }
                }
                .uploadBtn {
                  width: auto;
                  min-width: 0.65rem;
                  box-sizing: content-box;
                  padding: 0.05rem 0.07rem;
                  margin: auto;
                  color: #555555;
                  // box-shadow: 0 0 0.06rem rgba(191, 191, 191, 0.32);
                  box-shadow: none;
                  border-radius: 0.11rem;
                  border: 1px solid #dcdfe6;
                  white-space: nowrap;
                  display: inline-block;
                  &:hover {
                    opacity: 0.9;
                  }
                }
                .blue {
                  background: #4f87ff;
                  color: #fff;
                }
                .opacity {
                  box-shadow: none;
                  color: #b8b8b8;
                }
              }
              .hot-miner {
                .el-button {
                  span {
                    white-space: nowrap;
                  }
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
          .expandArea {
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
                .cidLink {
                  line-height: 0.25rem;
                  overflow: hidden;
                  white-space: normal;
                  display: -webkit-box;
                  -webkit-line-clamp: 1;
                  -webkit-box-orient: vertical;
                  color: inherit;
                }
                .cidMore {
                  position: absolute;
                  top: 0;
                  right: 0;
                  bottom: 0;
                  background-color: #fff;
                  width: 0.22rem;
                  padding: 0;
                  &::before {
                    content: "···";
                    font-weight: bold;
                    color: #0a318e;
                  }
                }
                .uploadBtn {
                  width: auto;
                  padding: 0.06rem 0.1rem;
                  margin: auto;
                  color: #5a5a5a;
                  box-shadow: 0 0 0.06rem rgba(191, 191, 191, 0.32);
                  border-radius: 4px;
                  border: 1px solid #dcdfe6;
                  white-space: nowrap;
                  &:hover {
                    background-color: #fff;
                  }
                }
                .el-button {
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
                  span {
                    line-height: 0.25rem;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: normal;
                    display: -webkit-box;
                    -webkit-line-clamp: 2;
                    -webkit-box-orient: vertical;
                  }
                  img {
                    display: none;
                    float: left;
                    width: 0.17rem;
                    margin-top: 0.03rem;
                  }
                }
                .el-button--primary {
                  color: #fff !important;
                  background-color: #409eff !important;
                  border-color: #409eff !important;
                }
                .el-button:hover {
                  img {
                    display: inline-block;
                  }
                }
              }
            }
            .el-table__expanded-cell {
              // padding: 0;
              .el-descriptions {
                overflow: auto;
                line-height: 0.25rem;
                z-index: 999;
                .el-descriptions__body {
                  padding: 0.2rem;
                  .el-descriptions__table {
                    tr {
                      background: transparent !important;
                      th,
                      td {
                        height: auto;
                        padding: 0.1rem;
                        text-align: left;
                        border: 0;
                      }
                      th {
                        text-transform: uppercase;
                        background: #fafafa !important;
                        &::after {
                          display: none;
                        }
                      }

                      .hot-cold-box {
                        position: relative;
                        .cidLink {
                          line-height: 0.25rem;
                          overflow: hidden;
                          white-space: normal;
                          display: -webkit-box;
                          -webkit-line-clamp: 1;
                          -webkit-box-orient: vertical;
                          color: inherit;
                        }
                        .cidMore {
                          position: absolute;
                          top: 0;
                          right: 0;
                          bottom: 0;
                          background-color: #fff;
                          width: 0.22rem;
                          padding: 0;
                          &::before {
                            content: "···";
                            font-weight: bold;
                            color: #0a318e;
                          }
                        }
                        .uploadBtn {
                          width: auto;
                          padding: 0.06rem 0.1rem;
                          margin: auto;
                          color: #5a5a5a;
                          box-shadow: 0 0 0.06rem rgba(191, 191, 191, 0.32);
                          border-radius: 4px;
                          border: 1px solid #dcdfe6;
                          white-space: nowrap;
                          &:hover {
                            background-color: #fff;
                          }
                        }
                        .el-button {
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
                          span {
                            line-height: 0.25rem;
                            overflow: hidden;
                            text-overflow: ellipsis;
                            white-space: normal;
                            display: -webkit-box;
                            -webkit-line-clamp: 2;
                            -webkit-box-orient: vertical;
                          }
                          img {
                            display: none;
                            float: left;
                            width: 0.17rem;
                            margin-top: 0.03rem;
                          }
                        }
                        .el-button:hover {
                          img {
                            display: inline-block;
                          }
                        }
                      }
                    }
                  }
                }
              }
            }
            .storageStyle {
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
              div {
                width: 100%;
                line-height: 2;
                margin: 0 auto;
              }
              .costTitle {
                font-size: 0.18rem;
                font-weight: 600;
                color: #000;
              }
              .costPrice {
                margin: 0.1rem 0;
                font-size: 0.18rem;
                font-weight: 600;
                color: #e00000;
              }
              .costDesc {
                width: 80%;
                font-size: 0.14rem;
                color: #000;
                line-height: 1;
              }
            }
          }
          &:hover {
            td {
              .cell {
                .hot-cold-box {
                  .cidMore {
                    background-color: #f5f7fa;
                  }
                }
              }
            }
          }
        }
        .current-row {
          td {
            background-color: #f5f7fa;
            .cell {
              .hot-cold-box {
                .cidMore {
                  background-color: #f5f7fa;
                }
              }
            }
          }
        }
      }
      .el-table::before {
        display: none;
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
      position: relative;
      display: flex;
      justify-content: center;
      align-items: center;
      text-align: center;
      margin: 0.5rem 0 0.06rem;
      padding: 0;
      @media screen and (max-width: 1024px) {
        padding: 0;
      }
      @media screen and (max-width: 600px) {
        flex-wrap: wrap;
      }
      .pagination {
        display: flex;
        align-items: center;
        margin: 0;
        font-size: 0.18rem;
        color: #000;
        .el-pagination /deep/ {
          .el-icon {
            font-size: 0.18rem;
          }
          .el-pager {
            li {
              min-width: 30px;
              height: 30px;
              font-size: 0.18rem;
              font-weight: normal;
              line-height: 30px;
            }
            .active {
              border: 1px solid #6798f5;
              border-radius: 5px;
              color: #000;
            }
          }
        }
        .el-select /deep/ {
          max-width: 100px;
          margin-right: 0.15rem;
          .el-input__inner,
          .el-input__icon {
            height: 30px;
            line-height: 30px;
          }
        }
        .span {
          margin: 0 0 0 10px;
          font-size: 13px;
          font-weight: 400;
          color: #606266;
          white-space: nowrap;
        }
        .paginaInput /deep/ {
          max-width: 50px;
          .el-input__inner,
          .el-input__icon {
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
      .down {
        position: absolute;
        right: 0;
        font-size: 0.18rem;
        color: #888;
        cursor: pointer;
        @media screen and (max-width: 600px) {
          position: relative;
          width: 100%;
          text-align: center;
        }
        span {
          color: #2c7ff8;
        }
        &:hover {
          text-decoration: underline;
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

    .metaM {
      .el-dialog__body {
        padding: 0.25rem 0.25rem 0.2rem;
        .el-row {
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
          .el-col {
            text-align: left;
            font-size: 0.14rem;
            img {
              float: right;
              height: 0.24rem;
            }
          }
        }
      }
      .el-dialog__footer {
        padding: 0 0.25rem 0.3rem;
        .dialog-footer {
          .el-button {
            width: 100%;
            font-size: 0.14rem;
            height: 0.4rem;
            padding: 0;
            background: #5c3cd3;
            color: #fff;
            border-radius: 0.08rem;
            &:hover {
              background: #4326ab;
            }
          }
          p {
            font-size: 0.13rem;
            line-height: 1.5;
            color: red;
            margin: 0.1rem 0 0;
          }
        }
      }
    }
    .completeDia {
      text-align: center;
      box-shadow: 0 0 13px rgba(128, 128, 128, 0.8);
      border-radius: 0.2rem;
      .el-dialog__header {
        display: none;
      }
      img {
        display: block;
        max-width: 100px;
        margin: 0 auto 0.3rem;
      }
      h1 {
        margin: 0.22rem auto 0.1rem;
        font-size: 0.32rem;
        font-weight: 500;
        line-height: 1.2;
        color: #191919;
        word-break: break-word;
      }
      h2 {
        margin: 0 auto 0.1rem;
        font-size: 0.22rem;
        font-weight: 600;
        line-height: 1.4;
        color: #555;
        word-break: break-word;
        text-align: center;
      }
      h4 {
        font-size: 0.2rem;
        font-weight: 500;
        line-height: 1.4;
        color: #555;
        word-break: break-word;
        text-align: center;
      }
      h3,
      a {
        font-size: 0.22rem;
        font-weight: 500;
        line-height: 1.4;
        color: #191919;
        word-break: break-word;
      }
      a {
        text-decoration: underline;
        color: #007bff;
      }
      a.a-close {
        height: 0.6rem;
        line-height: 0.6rem;
        padding: 0 45px;
        background: linear-gradient(45deg, #4f8aff, #4b5eff);
        font-size: 0.22rem;
        color: #fff;
        border-radius: 0.14rem;
        cursor: pointer;
        margin: 0.4rem auto 0;
        display: block;
        width: max-content;
        text-decoration: unset;
        &:hover {
          opacity: 0.9;
        }
      }
    }

    .wrongNet {
      margin: auto !important;
      background: #fff url(../../../assets/images/tip_bg.png) no-repeat;
      background-size: 1.45rem;
      background-position: -0.2rem -0.3rem;
      border-radius: 4px;
      .el-dialog__header {
        display: flex;
        color: #000;
        font-size: 18px;
        padding: 0.15rem 0.15rem 0.1rem;
        .el-dialog__headerbtn {
          i {
            color: #000;
            font-weight: 600;
          }
        }
      }
      .el-dialog__body {
        padding: 0.2rem 0.2rem 0.3rem;
        span {
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
  #dealManagement {
    .form {
      .form_top {
        .search_file {
          flex-wrap: wrap;
          height: auto;
          .search_right {
            width: 100%;
            margin: 0.05rem 0 0.1rem;
          }
        }
      }
    }
  }
}
@media screen and (max-width: 341px) {
}
</style>
