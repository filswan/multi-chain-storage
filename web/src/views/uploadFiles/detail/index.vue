<template>
  <div id="dealManagement">
    <div class="backTo" @click="back" v-if="false">
      <span class="el-icon-back"></span>
      <span style="font-size:0.22rem;margin-left:0.05rem">{{$t('deal.backto')}}</span>
    </div>
    <div class="detailStyle">
      <div v-loading="loading">
        <div class="files_title">
          <div class="flex_left">
            {{$t('uploadFile.Deal_Detail')}}

            <b v-if="dealCont.source_file_upload_deal.deal_id || dealCont.source_file_upload_deal.deal_id==0 || (dealId&&dealId!=0)" @click="mainnetLink(dealId)" class="golink">#{{dealId}}</b>
            <b v-else style="margin: 0 0 0 5px;">#</b>

            <span class="title" v-if="dealId == 0">
              <el-tooltip effect="dark" :content="$t('uploadFile.detail_tip01')" placement="top">
                <img src="@/assets/images/info.png" class="resno" />
              </el-tooltip>
            </span>
            <span v-if="isFree == 1"></span>
            <span v-else-if="!dealCont.source_file_upload_deal.locked_fee">
              <img src="@/assets/images/error.png" class="resno" />
              <span>{{$t('uploadFile.no_fund_locked')}}</span>
            </span>
            <span v-else-if="dealCont.source_file_upload_deal.unlocked">
              <img src="@/assets/images/dao_success.png" class="resno" />
              <span style="color: #3db39e;">{{$t('uploadFile.Successfully_unlocked_funds')}}</span>
            </span>
            <span v-else-if="dealCont.dao_signature.length >= dealCont.dao_threshold">
              <img src="@/assets/images/dao_waiting.png" class="resno" />
              <span>{{dealCont.daoSignatureAll == dealCont.dao_signature.length?$t('uploadFile.Successfully_signed_all')+$t('uploadFile.Successfully_signed'):$t('uploadFile.Successfully_signed')}} {{dealCont.daoSignatureAll}}/{{dealCont.dao_signature.length}}
              </span>
            </span>
            <span v-else>
              <img src="@/assets/images/dao_waiting.png" class="resno" />
              <span>{{$t('uploadFile.Waiting_for_signature')}} {{dealCont.daoSignatureAll}}/{{dealCont.dao_signature.length}} </span>
            </span>
          </div>
          <el-button type="primary" size="small" @click="getDealLogsData">{{$t('uploadFile.view_deal_logs')}}</el-button>
        </div>

        <div class="descMain">
          <i class="el-icon-circle-close closePop" v-if="ipfsUploadLoad" @click="controllerSignal()"></i>
          <div class="loadTryAgain" v-if="ipfsUploadLoad">
            <div style="width:100%;">
              <div class="load_svg">
                <svg viewBox="25 25 50 50" class="circular">
                  <circle cx="50" cy="50" r="20" fill="none" class="path"></circle>
                </svg>
                <p>
                  {{$t('uploadFile.payment_tip_deal')}}
                  <span @click="controllerSignal('try_again', dealCont.source_file_upload_deal.ipfs_url, dealCont.source_file_upload_deal.file_name)">{{$t('metaSpace.try_again')}}</span>
                </p>
              </div>
            </div>
          </div>
          <el-descriptions title="" :column="1">
            <el-descriptions-item :label="$t('uploadFile.file_name')">{{dealCont.source_file_upload_deal.file_name | NumFormat}}</el-descriptions-item>
            <el-descriptions-item :label="$t('uploadFile.detail_IPFSDownload')">
              <div class="module">
                <!-- <a class="linkTo" v-if="dealCont.source_file_upload_deal.ipfs_url" @click="xhrequest(dealCont.source_file_upload_deal.ipfs_url, dealCont.source_file_upload_deal.file_name)"> -->
                <a class="linkTo" v-if="dealCont.source_file_upload_deal.ipfs_url" :href="dealCont.source_file_upload_deal.ipfs_url" target="_blank">
                  {{dealCont.source_file_upload_deal.ipfs_url}}
                </a>
                <span v-else>-</span>
                <img class="imgCopy" src="@/assets/images/copy.png" @click="copyTextToClipboard(dealCont.source_file_upload_deal.ipfs_url)" v-if="dealCont.source_file_upload_deal.ipfs_url" alt="">
              </div>
            </el-descriptions-item>
            <el-descriptions-item :label="$t('uploadFile.detail_Network')">{{dealCont.source_file_upload_deal.network_name | NumFormat}}</el-descriptions-item>
            <el-descriptions-item :label="$t('billing.PAYLOADCID')">{{dealCont.source_file_upload_deal.car_file_payload_cid | NumFormat}}</el-descriptions-item>
          </el-descriptions>
        </div>

        <el-tabs v-model="activeName" :tab-position="tabPosition" type="card" @tab-click="handleClick">
          <el-tab-pane v-for="(item, i) in offline_deals_data" :key="i" :name="''+i+''">
            <span slot="label">
              <i v-if="isFree == 1" style="display: none;"></i>
              <img v-else-if="!dealCont.source_file_upload_deal.locked_fee" src="@/assets/images/error.png" class="resno" />
              <img v-else-if="dealCont.source_file_upload_deal.unlocked" src="@/assets/images/dao_success.png" class="resno" />
              <img v-else-if="dealCont.dao_signature.length >= dealCont.dao_threshold" src="@/assets/images/dao_waiting.png" class="resno" />
              <img v-else src="@/assets/images/dao_waiting.png" class="resno" /> {{item.miner_fid}}
            </span>
          </el-tab-pane>
        </el-tabs>

        <div class="upload">
          <el-row :class="{'elColLeftEn': languageMcs === 'en', 'elColLeftZh': languageMcs === 'cn'}">
            <el-col :span="8">{{$t('uploadFile.detail_Locked_funds')}}:</el-col>
            <el-col :span="16">{{dealCont.source_file_upload_deal.locked_fee | NumFormatPrice}} USDC</el-col>
            <el-col :span="8">{{$t('uploadFile.w3ss_id')}}:</el-col>
            <el-col :span="16" v-if="dealId == 0">-</el-col>
            <el-col :span="16" v-else>{{dealCont.source_file_upload_deal.provider | NumFormat}}</el-col>
            <el-col :span="8">{{$t('uploadFile.detail_Storage_Price')}}:</el-col>
            <el-col :span="16" v-if="dealId == 0">-</el-col>
            <el-col :span="16" v-else>{{dealCont.source_file_upload_deal.storage_price | NumFormatPrice}} FIL</el-col>
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
                  <img src="@/assets/images/info.png" />
                </el-tooltip>：
              </el-col>
              <el-col :span="16">
                <img class="img" src="@/assets/images/copy.png" @click="copyTextToClipboard(copy_filename)" alt="">
              </el-col>
            </el-col>
            <!-- :class="{'color': !dealCont.source_file_upload_deal.provider && !dealCont.source_file_upload_deal.car_file_payload_cid}" -->
            <el-col :span="24" class="lotupContent" @click="copyTextToClipboard(copy_filename)">
              {{copy_filename}}
            </el-col>
          </el-row>

          <div class="title" v-if="isFree != 1">
            {{$t('uploadFile.detail_DAO_Signatures')}}
            <el-tooltip effect="dark" :content="$t('uploadFile.detail_DAO_Signatures_tooltip')" placement="top">
              <img src="@/assets/images/info.png" />
            </el-tooltip>
          </div>
          <el-table v-if="isFree != 1" :data="daoCont" stripe style="width: 100%">
            <el-table-column type="index" width="180">
              <template slot-scope="scope">
                <!-- Signature {{scope.$index+1}} -->
                {{$t('my_profile.miner_add_Signature')}} {{scope.row.order_index?scope.row.order_index:scope.$index+1}}
              </template>
            </el-table-column>
            <el-table-column prop="wallet_signer" :label="$t('uploadFile.detail_DAO_RKH_Address')" min-width="220">
              <template slot-scope="scope">
                <div class="hot-cold-box">
                  <el-popover placement="top" trigger="hover" v-model="scope.row.daoAddressVis">
                    <div class="upload_form_right">
                      <p>{{scope.row.wallet_signer}}</p>
                    </div>
                    <el-button slot="reference" class="resno" @click="copyTextToClipboard(scope.row.wallet_signer)">
                      <img src="@/assets/images/copy.png" alt=""> {{scope.row.wallet_signer}}
                    </el-button>
                  </el-popover>
                </div>
              </template>
            </el-table-column>
            <el-table-column prop="tx_hash" :label="$t('billing.TRANSACTION')" min-width="220">
              <template slot-scope="scope">
                <div class="hot-cold-box">
                  <el-popover placement="top" trigger="hover" width="200" v-model="scope.row.txHashVis">
                    <div class="upload_form_right">
                      <p>{{scope.row.tx_hash}}</p>
                    </div>
                    <!-- :class="{'color': dealCont.network&&dealCont.network.toLowerCase() == 'polygon'}" -->
                    <el-button slot="reference" @click="networkLink(baseAddressURL+'tx/'+scope.row.tx_hash)" class="color">
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
                <img src="@/assets/images/dao_success.png" v-if="scope.row.status&&scope.row.status.toLowerCase() == 'success'" class="resno" />
                <img src="@/assets/images/dao_waiting.png" v-else class="resno" />
              </template>
            </el-table-column>
          </el-table>
        </div>
      </div>
    </div>

    <el-dialog :visible.sync="dialogVisible" :width="width" custom-class="dealLogs">
      <div slot="title" class="dialog-title">
        {{$t('uploadFile.deal_logs')}}
        <el-tooltip effect="dark" placement="top">
          <div slot="content">
            <a :href="webLink" target="_blank" class="weblinkStyle">{{webLink}}</a>
          </div>
          <img src="@/assets/images/info.png" />
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
// import axios from 'axios'
// import Qs from 'qs'
// import moment from 'moment'
let that
export default {
  name: 'my_files',
  data () {
    return {
      loading: false,
      tabPosition: 'top',
      bodyWidth: document.documentElement.clientWidth < 1024,
      dealId: '',
      logId: null,
      isFree: 0,
      dealCont: {
        daoSignatureAll: 0,
        source_file_upload_deal: {}
      },
      daoCont: [],
      copy_filename: '',
      activeName: localStorage.getItem('offlineDealsIndex') ? localStorage.getItem('offlineDealsIndex') : '0',
      offline_deals_data: [],
      dialogVisible: false,
      width: document.body.clientWidth > 600 ? '550px' : '95%',
      loadlogs: true,
      dealLogsData: [],
      webLink: 'https://docs.filecoin.io/get-started/store-and-retrieve/store-data/#deal-states',
      ipfsUploadLoad: false,
      controller: new AbortController()
    }
  },
  computed: {
    languageMcs () {
      return this.$store.getters.languageMcs
    },
    networkID () {
      return Number(this.$store.getters.networkID)
    }
  },
  watch: {
    $route: function (to, from) {
      this.dealId = to.params.deal_id
      this.getData()
    }
  },
  methods: {
    handleClick (tab, event) {
      localStorage.setItem('offlineDealsIndex', tab.index)
      that.logId = that.offline_deals_data[tab.index].id
      that.$router.push({
        name: 'my_files_detail',
        params: {
          id: that.offline_deals_data[tab.index].id,
          deal_id: that.offline_deals_data[tab.index].deal_id,
          // cid: that.$route.params.cid,
          source_file_upload_id: that.$route.params.source_file_upload_id,
          isFree: that.$route.params.isFree
        }
      })
    },
    networkLink (link) {
      window.open(link)
    },
    mainnetLink (dealId) {
      const networkName = that.dealCont.source_file_upload_deal.network_name
      switch (networkName) {
        case 'filecoin_calibration':
          window.open(`${process.env.BASE_CALIBRATION_ADDRESS}?dealid=${dealId}`)
          break
        case 'filecoin_mainnet':
          window.open(`${process.env.BASE_MAINNET_ADDRESS}?dealid=${dealId}`)
          break
        default:
          if (networkName && networkName.indexOf('calibration') > -1) window.open(`${process.env.BASE_CALIBRATION_ADDRESS}?dealid=${dealId}`)
          else window.open(`${process.env.BASE_MAINNET_ADDRESS}?dealid=${dealId}`)
          break
      }
    },
    copyTextToClipboard (text) {
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
          that.$message({
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
    back () {
      // that.$router.go(-1);//返回上一层
      that.$router.push({ name: 'my_files' })
    },
    getData () {
      that.loading = true

      let dataCid = {
        source_file_upload_id: that.$route.params.source_file_upload_id,
        // payload_cid: that.$route.params.cid,
        wallet_address: that.$store.getters.metaAddress
      }
      axios.get(`${that.baseAPIURL}api/v1/storage/deal/detail/${that.dealId}?${Qs.stringify(dataCid)}`, {
        // axios.get(`./static/detail_page_response.json`, {
        headers: {
          'Authorization': 'Bearer ' + that.$store.getters.mcsjwtToken
        }
      }).then((response) => {
        let json = response.data
        let daoSignatureAll = 0
        that.loading = false
        if (json.status === 'success') {
          if (!json.data) return false
          if (json.data.dao_signature) {
            that.daoCont = json.data.dao_signature
            that.daoCont.map(item => {
              if (item.create_at) daoSignatureAll += 1
              item.daoAddressVis = false
              item.txHashVis = false
              item.create_at =
                item.create_at
                  ? moment(new Date(parseInt(item.create_at * 1000))).format(
                    'YYYY-MM-DD HH:mm:ss'
                  )
                  : '-'
            })
          }

          that.dealCont = json.data
          that.dealCont.daoSignatureAll = daoSignatureAll
          that.dealCont.source_file_upload_deal.created_at =
            that.dealCont.source_file_upload_deal.created_at && that.dealCont.source_file_upload_deal.created_at !== 0
              ? moment(new Date(parseInt(String(that.dealCont.source_file_upload_deal.created_at).length < 13 ? that.dealCont.source_file_upload_deal.created_at * 1000 : that.dealCont.source_file_upload_deal.created_at))).format(
                'YYYY-MM-DD HH:mm:ss'
              )
              : '-'

          if (json.data.source_file_upload_deal.provider && json.data.source_file_upload_deal.car_file_payload_cid && json.data.source_file_upload_deal.deal_id !== 0) {
            that.copy_filename = 'lotus client retrieve --miner ' + json.data.source_file_upload_deal.provider + ' ' + json.data.source_file_upload_deal.car_file_payload_cid + ' ~/output-file'
          } else {
            that.copy_filename = localStorage.getItem('languageMcs') === 'cn' ? '还不可用。' : "It's not available yet."
          }
        } else {
          that.$message.error(json.message)
          return false
        }
      }).catch(function (error) {
        console.log(error)
        that.loading = false
      })
    },
    getDealLogsData () {
      that.dialogVisible = true
      that.loadlogs = true
      let obj = {
        wallet_address: that.$store.getters.metaAddress
      }
      axios.get(`${that.baseAPIURL}api/v1/storage/deal/log/${that.logId}?${Qs.stringify(obj)}`, {
        // axios.get(`./static/deal_logs.json`, {
        headers: {
          'Authorization': 'Bearer ' + that.$store.getters.mcsjwtToken
        }
      }).then((response) => {
        let json = response.data
        that.loadlogs = false
        if (json.status === 'success') {
          if (!json.data) return false
          if (json.data.offline_deal_log) {
            that.dealLogsData = json.data.offline_deal_log
            that.dealLogsData.map(item => {
              item.create_at =
                item.create_at
                  ? moment(new Date(parseInt(item.create_at * 1000))).format(
                    'YYYY-MM-DD HH:mm:ss'
                  )
                  : '-'
            })
          }
        } else {
          that.$message.error(json.message)
          return false
        }
      }).catch(function (error) {
        console.log(error)
        that.loadlogs = false
      })
    },
    downloadBlob (blob, fileName) {
      try {
        const href = window.URL.createObjectURL(blob)
        if (window.navigator.msSaveBlob) {
          window.navigator.msSaveBlob(blob, fileName)
        } else {
          const downloadElement = document.createElement('a')
          downloadElement.href = href
          downloadElement.target = '_blank'
          downloadElement.download = fileName
          document.body.appendChild(downloadElement)
          downloadElement.click()
          document.body.removeChild(downloadElement)
          window.URL.revokeObjectURL(href)
        }
      } catch (e) {
        console.log('ipfs upload err:', e)
      }
    },
    controllerSignal (type, link, name) {
      that.controller.abort()
      if (type) that.xhrequest(link, name)
      else that.ipfsUploadLoad = false
    },
    async xhrequest (link, name) {
      that.ipfsUploadLoad = true
      that.controller = new AbortController()
      let data = await fetch(link, { signal: that.controller.signal })
        .then((response) => response.blob())
        .then((res) => {
          console.log(res)
          let blod = new Blob([res])
          if (res.type.indexOf('javascript') > -1 || res.type.indexOf('image') > -1 || res.type.indexOf('text') > -1 || res.type.indexOf('json') > -1) that.downloadBlob(blod, name)
          else window.open(link)
          that.ipfsUploadLoad = false
        })
        .catch((e) => {
          console.log('Download error: ' + e.message)
          that.$message.error(e.message)
        })
      return data
    }
  },
  mounted () {
    that = this
    that.dealId = that.$route.params.deal_id
    that.logId = that.$route.params.id
    that.isFree = that.$route.params.isFree
    if (localStorage.getItem('offlineDeals')) that.offline_deals_data = JSON.parse(localStorage.getItem('offlineDeals'))
    that.getData()
    document.getElementById('content-box').scrollTop = 0
    that.$store.dispatch('setRouterMenu', 1)
    that.$store.dispatch('setHeadertitle', that.$t('navbar.deal'))
  },
  filters: {
    NumFormat (value) {
      if (String(value) === '0') return 0
      if (!value) return '-'
      return value
    },
    NumFormatPrice (value) {
      if (String(value) === '0') return 0
      if (!value) return '-'
      // 18 - 单位换算需要 / 1000000000000000000，浮点运算显示有bug
      let valueNum = String(value)
      if (value.length > 6) {
        let v1 = valueNum.substring(0, valueNum.length - 6)
        let v2 = valueNum.substring(valueNum.length - 6)
        let v3 = String(v2).replace(/(0+)\b/gi, '')
        if (v3) {
          return v1 + '.' + v3
        } else {
          return v1
        }
      } else {
        let v3 = ''
        for (let i = 0; i < 6 - valueNum.length; i++) {
          v3 += '0'
        }
        return '0.' + String(v3 + valueNum).replace(/(0+)\b/gi, '')
      }
    }
  }
}
</script>

<style scoped lang="scss">
.weblinkStyle {
  color: #fff;
  &:hover {
    text-decoration: underline;
  }
}
.el-dialog__wrapper /deep/ {
  display: flex;
  align-items: center;
  justify-content: center;
  .dealLogs {
    height: 80%;
    background: #fff;
    margin: auto !important;
    box-shadow: 0 0 13px rgba(128, 128, 128, 0.8);
    border-radius: 0.2rem;
    overflow: hidden;
    .el-dialog__header {
      display: flex;
      align-items: center;
      justify-content: space-between;
      padding: 0.3rem 0.4rem;
      display: flex;
      border-bottom: 1px solid #dfdfdf;
      color: #000;
      font-size: 20px;
      .dialog-title {
        display: flex;
        align-items: center;
        color: #000;
        font-size: 0.22rem;
        font-weight: 500;
        line-height: 1;
        text-transform: capitalize;
        .el-tooltip {
          width: 20px;
          height: 20px;
          margin: 0 0 0 5px;
          @media screen and (min-width: 1800px) {
            width: 22px;
            height: 22px;
          }
          @media screen and (max-width: 1440px) {
            width: 17px;
            height: 17px;
          }
          @media screen and (max-width: 1280px) {
            width: 16px;
            height: 16px;
          }
        }
      }
      .el-dialog__headerbtn {
        position: relative;
        top: auto;
        right: auto;
        font-size: inherit;
        i {
          font-size: inherit;
          &:hover {
            color: #0b318f;
          }
        }
      }
      .el-dialog__title {
        font-size: inherit;
      }
    }
    .el-dialog__body {
      height: calc(100% - 1.35rem);
      padding: 0.2rem 0.4rem;
      font-size: 14px;
      overflow-y: scroll;
      .block {
        position: relative;
        min-height: 100px;
        .el-timeline {
          .el-timeline-item {
            .el-card__body {
              h4 {
                font-size: 0.2rem;
                color: #000;
              }
              p {
                word-break: break-word;
                font-size: 0.18rem;
                color: #555;
              }
            }
            .el-timeline-item__timestamp {
              font-size: 0.18rem;
              color: #555;
            }
          }
        }
        .noLogs {
          font-size: 16px;
          text-align: center;
        }
        .el-loading-mask {
          .el-loading-spinner {
            top: 50%;
          }
        }
      }
    }
    .el-dialog__footer {
      padding: 0 0.2rem 0.1rem;
      .dialog-footer {
        display: flex;
        align-items: center;
        justify-content: flex-end;
        .el-button {
          font-size: 13px;
        }
      }
    }
  }
}
#dealManagement {
  padding: 0.3rem;
  .backTo {
    display: flex;
    align-items: center;
    font-size: 0.24rem;
    line-height: 1;
    max-width: 200px;
    margin: 0 0 0.2rem;
    cursor: pointer;
  }
  .detailStyle {
    .el-tabs /deep/ {
      .el-tabs__header {
        margin: 0;
        background: transparent;
        .el-tabs__nav-next,
        .el-tabs__nav-prev {
          font-size: 18px;
          i {
            margin-top: 0.3rem;
            font-weight: 600;
          }
        }
        .el-tabs__nav {
          display: flex;
          border: 0;
          @media screen and (max-width: 1024px) {
            display: flex;
          }
          .el-tabs__item {
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
            color: #2d43e7;
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
            span {
              width: 100%;
              display: flex;
              justify-content: space-between;
              align-items: center;
              font-size: 0.18rem;
            }
            i,
            img {
              width: 20px;
              height: 20px;
              margin: 0 15px 0 0;
              opacity: 0;
              font-size: 16px;
              color: #67c23a;
            }
            &:before {
              content: "";
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
          .is-active {
            position: relative;
            border-right: 0;
            color: #fff;
            &:before {
              background: #2d43e7;
            }
            i,
            img {
              opacity: 1;
            }
          }
        }
      }
    }
  }
  .files_title {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin: 0.4rem 0 0.1rem;
    font-size: 0.22rem;
    font-weight: bold;
    line-height: 2;
    @media screen and (max-width: 600px) {
      width: 100%;
      font-size: 16px;
      flex-wrap: wrap;
    }
    .flex_left {
      display: flex;
      align-items: center;
      @media screen and (max-width: 410px) {
        width: 100%;
        font-size: 16px;
        flex-wrap: wrap;
      }
      .golink {
        margin: 0 0 0 5px;
        cursor: pointer;
        &:hover {
          color: #2d43e7;
          text-decoration: underline;
        }
      }
      img {
        width: 20px;
        height: 20px;
        margin: 0 0 0 15px;
        cursor: pointer;
        @media screen and (min-width: 1800px) {
          width: 22px;
          height: 22px;
        }
        @media screen and (max-width: 1440px) {
          width: 17px;
          height: 17px;
        }
        @media screen and (max-width: 600px) {
          margin: 0;
        }
      }
      span {
        display: flex;
        align-items: center;
        padding-left: 5px;
        color: red;
        font-size: 0.18rem;
        @media screen and (max-width: 600px) {
          font-size: 14px;
        }
      }
      .title {
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
        img {
          width: 20px;
          height: 20px;
          margin: 0 0 0 5px;
          cursor: pointer;
          @media screen and (min-width: 1800px) {
            width: 22px;
            height: 22px;
          }
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
    }
    .el-button {
      padding: 0.1rem 0.2rem;
      font-size: 0.18rem;
      font-family: inherit;
      border-radius: 0.07rem;
      background: #4e82ff;
      background: linear-gradient(45deg, #4f8aff, #4b5eff);
      border-radius: 0.08rem;
      border: 0;
      &:hover {
        opacity: 0.9;
      }
      @media screen and (max-width: 600px) {
        margin: 0 0 5px;
      }
    }
  }
  .upload {
    padding: 0.1rem 0.3rem 0.2rem;
    background-color: #fff;
    border-bottom-left-radius: 0.1rem;
    border-bottom-right-radius: 0.1rem;
    overflow: hidden;
    border-top: 3px solid #0b318f;
    .el-row /deep/ {
      width: 100%;
      padding: 0 0 0.15rem;
      .el-col {
        padding: 0.1rem 0;
        font-size: 0.18rem;
        font-weight: normal;
        line-height: 1.3;
        color: #555;
        word-break: break-word;
        @media screen and (max-width: 600px) {
          width: 100%;
          font-size: 14px;
        }
        .linkTo {
          color: #2d43e7;
          text-decoration: underline;
        }
        .imgCopy {
          width: 18px;
          height: 18px;
          margin: 0 0 0 5px;
          cursor: pointer;
          @media screen and (min-width: 1800px) {
            width: 22px;
            height: 22px;
          }
          @media screen and (max-width: 1280px) {
            width: 16px;
            height: 16px;
          }
        }
      }
      .lotupTitle {
        display: flex;
        align-items: center;
        width: 100%;
        margin: 0;
        font-size: inherit;
        color: inherit;
        @media screen and (max-width: 600px) {
          flex-wrap: wrap;
        }
        .el-col {
          display: flex;
          align-items: center;
          @media screen and (max-width: 600px) {
            width: auto;
          }
        }
        img {
          width: 18px;
          height: 18px;
          margin: 0;
          cursor: pointer;
          @media screen and (min-width: 1800px) {
            width: 22px;
            height: 22px;
          }
          @media screen and (max-width: 1280px) {
            width: 16px;
            height: 16px;
          }
        }
        span {
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
        &:hover {
          color: #333;
        }
      }
      .color {
        color: #f63d3d;
        cursor: text;
        &:hover {
          color: #f63d3d;
        }
      }
      .el-col:nth-child(2n + 1) {
        color: #000;
      }
    }
    .elColLeftZh {
      .el-col-8 {
        width: 2.5rem;
        min-width: 160px;
        @media screen and (max-width: 600px) {
          width: 100%;
        }
      }
      .el-col-16 {
        width: calc(100% - 3rem);
        max-width: calc(100% - 160px);
        @media screen and (max-width: 600px) {
          width: 100%;
          max-width: 100%;
        }
      }
    }
    .elColLeftEn {
      .el-col-8 {
        width: 3.8rem;
        min-width: 240px;
        @media screen and (max-width: 600px) {
          width: 100%;
        }
      }
      .el-col-16 {
        width: calc(100% - 4.3rem);
        max-width: calc(100% - 240px);
        @media screen and (max-width: 600px) {
          width: 100%;
          max-width: 100%;
        }
      }
    }
    .title {
      display: flex;
      align-items: center;
      justify-content: flex-start;
      padding: 0.2rem 0;
      line-height: 1.5;
      text-align: center;
      font-size: 0.2rem;
      font-weight: 600;
      white-space: normal;
      color: #333;
      text-shadow: 0 0 black;
      text-indent: 0;
      border-top: 1px solid #e4e4e4;
      @media screen and (max-width: 600px) {
        font-size: 16px;
      }
      img {
        width: 18px;
        height: 18px;
        margin: 0 0 0 3px;
        cursor: pointer;
        @media screen and (min-width: 1800px) {
          width: 22px;
          height: 22px;
        }
        @media screen and (max-width: 1280px) {
          width: 16px;
          height: 16px;
        }
      }
    }
    .el-table /deep/ {
      td,
      th {
        .cell {
          display: flex;
          justify-content: center;
          align-items: center;
          word-break: break-word;
          text-align: center;
          font-size: 0.18rem;
          color: #555;
          .hot-cold-box {
            .el-button {
              width: 100%;
              border: 0;
              padding: 0;
              margin: 0;
              background-color: transparent;
              font-size: 0.18rem;
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
              @media screen and (max-width: 600px) {
                font-size: 13px;
              }
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
            .color {
              color: #4326ab;
              text-decoration: underline;
            }
            .el-button:hover {
              img {
                display: inline-block;
              }
            }
          }
          a {
            color: #4326ab;
            &:hover {
              text-decoration: underline;
            }
          }
          img {
            width: 30px;
          }
        }
      }
      td {
        padding: 0.15rem 0;
      }
      .el-table__header-wrapper {
        border-radius: 0.2rem;
        th {
          height: 0.6rem;
          background-color: #e5eeff !important;
        }
      }
    }
  }
  .descMain /deep/ {
    position: relative;
    padding: 0.15rem 0.3rem;
    margin: 0 0 0.1rem;
    background-color: #fff;
    border-radius: 0.1rem;
    .el-descriptions {
      .el-descriptions__body {
        .el-descriptions__table {
          .el-descriptions-item__cell {
            padding: 0.05rem 0;
            font-size: 0.18rem;
            font-weight: normal;
            color: #555;
            word-break: break-word;
            .el-descriptions-item__container {
              .el-descriptions-item__label {
                color: #000;
              }
              @media screen and (max-width: 600px) {
                flex-wrap: wrap;
                .el-descriptions-item__label {
                  width: 100%;
                }
              }
            }
            @media screen and (max-width: 600px) {
              font-size: 14px;
            }
          }
        }
      }
    }
    .module {
      display: flex;
      align-items: center;
      .linkTo {
        color: #2d43e7;
        text-decoration: underline;
        cursor: pointer;
      }
      .imgCopy {
        width: 18px;
        height: 18px;
        margin: 0 0 0 5px;
        cursor: pointer;
        @media screen and (min-width: 1800px) {
          width: 22px;
          height: 22px;
        }
        @media screen and (max-width: 1280px) {
          width: 16px;
          height: 16px;
        }
        @media screen and (max-width: 600px) {
          width: 20px;
          height: 20px;
        }
      }
    }
    .el-loading-mask {
      .el-loading-spinner {
        top: 0;
        margin: 0;
      }
    }
    .closePop {
      position: absolute;
      right: -0.1rem;
      top: -0.1rem;
      width: 0.37rem;
      height: 0.37rem;
      background: #fff;
      border-radius: 100%;
      font-size: 0.37rem;
      color: #000;
      cursor: pointer;
      z-index: 2001;
      @media screen and (max-width: 768px) {
        right: -15px;
        top: -15px;
        width: 30px;
        height: 30px;
        font-size: 30px;
      }
    }
    .loadTryAgain {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      display: flex;
      justify-content: center;
      align-items: center;
      background: rgba(255, 255, 255, 0.8);
      .load_svg {
        display: flex;
        justify-content: center;
        align-items: center;
        flex-wrap: wrap;
        width: 100%;
        svg {
          height: 42px;
          width: 42px;
          -webkit-animation: loading-rotate 2s linear infinite;
          animation: loading-rotate 2s linear infinite;
          .path {
            -webkit-animation: loading-dash 1.5s ease-in-out infinite;
            animation: loading-dash 1.5s ease-in-out infinite;
            stroke-dasharray: 90, 150;
            stroke-dashoffset: 0;
            stroke-width: 2;
            stroke: #409eff;
            stroke-linecap: round;
          }
        }
        p {
          width: 100%;
          text-align: center;
          font-size: 16px;
          color: #333;
          @media screen and (max-width: 1600px) {
            font-size: 14px;
          }
          @media screen and (max-width: 768px) {
            font-size: 13px;
          }
          @media screen and (max-width: 441px) {
            font-size: 12px;
          }
          span {
            color: #4b83fb;
            cursor: pointer;
          }
        }
      }
    }
  }
}

@media screen and (max-width: 999px) {
  #dealManagement {
    padding: 0.15rem 0.1rem 0.2rem;
    .backTo {
      margin: 0.2rem 0;
    }
    .upload {
      padding: 0.1rem;
    }
  }
}
</style>
