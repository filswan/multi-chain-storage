<template>
    <div id="dealManagement" v-loading="loading">
        <div class="form">
            <div class="form_top">
                <div class="search">
                    <el-input placeholder="Search by full Data CID" prefix-icon="el-icon-search" v-model="searchValue">
                    </el-input>
                    <div class="search_right" :style="{'opacity': !searchValue?'0.8':'1'}">
                        <el-button @click="search" style="background-color: #ffb822" :disabled="!searchValue">Search</el-button>
                        <el-button type="primary" style="background-color: #0b318f" @click="clearAll" :disabled="!searchValue">
                            {{ $t("deal.Clear_all") }}
                        </el-button>
                    </div>
                </div>
            </div>
            <div class="form_table">
                <el-table :data="tableData" style="width: 100%" :empty-text="$t('deal.formNotData')" header-click="aaa">
                    <el-table-column prop="deal_cid" :label="$t('deal.detailTable01')">
                        <template slot-scope="scope">
                            <div class="hot-cold-box">
                                <el-popover placement="top" trigger="hover" v-model="scope.row.visible">
                                    <div class="upload_form_right">
                                        <p>{{scope.row.deal_cid}}</p>
                                    </div>
                                    <el-button slot="reference" @click="copyTextToClipboard(scope.row.deal_cid)">
                                        <img src="@/assets/images/copy.png" alt=""> {{scope.row.deal_cid}}
                                    </el-button>
                                </el-popover>
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column prop="payload_cid" :label="$t('deal.detailTable09')">
                        <template slot-scope="scope">
                            <div class="hot-cold-box">
                                <el-popover placement="top" trigger="hover" v-model="scope.row.payload" v-if="scope.row.payload_cid">
                                    <div class="upload_form_right">
                                        <p>{{scope.row.payload_cid | NumFormat}}</p>
                                    </div>
                                    <el-button slot="reference" @click="copyTextToClipboard(scope.row.payload_cid)">
                                        <img src="@/assets/images/copy.png" alt=""> {{scope.row.payload_cid | NumFormat}}
                                    </el-button>
                                </el-popover>
                                <div v-else>{{scope.row.payload_cid | NumFormat}}</div>
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column prop="file_source_url" :label="$t('deal.detailTable03')">
                        <template slot-scope="scope">
                            <div class="hot-cold-box">
                                <el-popover placement="top" trigger="hover" v-model="scope.row.visibleurl">
                                    <div class="upload_form_right">
                                        <p>{{scope.row.file_source_url | NumFormat}}</p>
                                    </div>
                                    <el-button slot="reference" @click="copyTextToClipboard(scope.row.file_source_url)">
                                        <img src="@/assets/images/copy.png" alt=""> {{scope.row.file_source_url | NumFormat}}
                                    </el-button>
                                </el-popover>
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column prop="status" width="150" :label="$t('deal.detailTable04')">
                        <template slot-scope="scope">
                            <div class="statusStyle" v-if="scope.row.status == 'Created'" :style="$statusColor.DealColor('Created')">
                                {{ languageMcs == 'en' ? 'Created' : '已创建' }}
                            </div>
                            <div class="statusStyle" v-if="scope.row.status == 'DealActive'" :style="$statusColor.DealColor('DealActive')">
                                {{ languageMcs == 'en' ? 'DealActive' : '有效交易' }}
                            </div>
                            <div class="statusStyle" v-else-if="scope.row.status == 'Waiting'" :style="$statusColor.DealColor('Waiting')">
                                {{ languageMcs == 'en' ? 'Waiting' : '等待中' }}
                            </div>
                            <div class="statusStyle" v-else-if="scope.row.status == 'ReadyForImport'" :style="$statusColor.DealColor('ReadyForImport')">
                                {{ languageMcs == 'en' ? 'ReadyForImport' : '准备导入' }}
                            </div>
                            <div class="statusStyle" v-else-if="scope.row.status == 'FileImporting'" :style="$statusColor.DealColor('FileImporting')">
                                {{ languageMcs == 'en' ? 'FileImporting' : '文件导入中' }}
                            </div>
                            <div class="statusStyle" v-else-if="scope.row.status == 'FileImported'" :style="$statusColor.DealColor('FileImported')">
                                {{ languageMcs == 'en' ? 'FileImported' : '文件已导入' }}
                            </div>
                            <div class="statusStyle" v-else-if="scope.row.status == 'ImportFailed'" :style="$statusColor.DealColor('ImportFailed')">
                                {{ languageMcs == 'en' ? 'ImportFailed' : '导入失败' }}
                            </div>
                            <div class="statusStyle" v-else-if="scope.row.status == 'Downloading'" :style="$statusColor.DealColor('Downloading')">
                                {{ languageMcs == 'en' ? 'Downloading' : '下载中' }}
                            </div>
                            <div class="statusStyle" v-else-if="scope.row.status == 'Downloaded'" :style="$statusColor.DealColor('Downloaded')">
                                {{ languageMcs == 'en' ? 'Downloaded' : '已下载' }}
                            </div>
                            <div class="statusStyle" v-else-if="scope.row.status == 'DownloadFailed'" :style="$statusColor.DealColor('DownloadFailed')">
                                {{ languageMcs == 'en' ? 'DownloadFailed' : '下载失败' }}
                            </div>
                            <div class="statusStyle" v-else-if="scope.row.status == 'Completed'" :style="$statusColor.DealColor('Completed')">
                                {{ languageMcs == 'en' ? 'Completed' : '已完成' }}
                            </div>
                            <div class="statusStyle" v-else-if="scope.row.status == 'Failed'" :style="$statusColor.DealColor('Failed')">
                                {{ languageMcs == 'en' ? 'Failed' : '已失败' }}
                            </div>
                            <div class="statusStyle" v-else style="display:none;color: rgb(255, 184, 34)">
                                {{ scope.row.status }}
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column prop="md5_origin" width="110" :label="$t('deal.table_th_08')"></el-table-column>
                    <el-table-column prop="note" :label="$t('deal.detailTable08')"></el-table-column>
                    <el-table-column prop="start_epoch" width="100" :label="$t('deal.detailTable05')">
                        <template slot-scope="scope">
                            {{scope.row.start_epoch}} <br> ({{scope.row.start_epoch_time}})
                        </template>
                    </el-table-column>
                    <el-table-column prop="updated_at" width="100" :label="$t('deal.detailTable06')"></el-table-column>
                    <el-table-column prop="created_at" width="100" :label="$t('deal.detailTable07')"></el-table-column>
                </el-table>
            </div>

            <div class="form_pagination" v-if="1==2">
                <!-- <div class="pagination">
                    <div class="pagination_left" @click="prev">&lt;</div>
                    {{ parma.offset }}
                    <div class="pagination_left" @click="next">&gt;</div>
                </div> -->
                <div class="pagination">
                    <el-pagination :total="parma.total" :page-size="parma.limit" :current-page="parma.offset" :pager-count="bodyWidth?5:7" style="padding: 30px 40px 32px 20px;" background :layout="bodyWidth?'prev, pager, next':'total, prev, pager, next, jumper'"
                        @current-change="handleCurrentChange" />
                </div>
            </div>
        </div>

        <!-- 回到顶部 -->
        <el-backtop target=".content-box" :bottom="40" :right="20"></el-backtop>
    </div>
</template>

<script>
// import bus from '@/components/bus'
import * as myAjax from '@/api/uploadFile'
import moment from 'moment'

export default {
  name: 'SearchFile',
  data () {
    return {
      tableData: [],
      tableDataAll: [],
      searchValue: '',
      parma: {
        task_id: '',
        limit: 20,
        offset: 0,
        total: 0
      },
      loading: false,
      downCsv: localStorage.getItem('addressYM') ? localStorage.getItem('addressYM') : '',
      bodyWidth: document.documentElement.clientWidth < 1024,
      max_time_str: null
    }
  },
  computed: {
    languageMcs () {
      return this.$store.state.app.languageMcs
    }
  },
  components: {},
  methods: {
    handleCurrentChange (val) {
      this.parma.offset = val
      this.getData()
    },
    getData () {
      let _this = this
      _this.loading = true

      myAjax.getTasksSearch(_this.searchValue).then(response => {
        if (response.status === 'success' || response.status === 'Success') {
          // _this.tableData = Array.from(new Set(response.data.task));

          _this.tableData = []
          _this.tableData = JSON.parse(JSON.stringify(response.data.deal))
          _this.parma.total = response.data.total_deal_count

          _this.max_time_str = null
          _this.tableData.map(item => {
            item.visible = false
            item.visibleurl = false
            item.payload = false
            _this.max_time_str = Math.max(_this.max_time_str, item.updated_at)
            item.start_epoch_time =
                            item.start_epoch
                              ? moment(new Date(parseInt((parseInt(item.start_epoch) * 30 + parseInt(1598306471)) * 1000))).format('YYYY-MM-DD HH:mm:ss')
                              : '-'
            item.created_at =
                            item.created_at
                              ? item.created_at.length < 13
                                ? moment(new Date(parseInt(item.created_at * 1000))).format('YYYY-MM-DD HH:mm:ss')
                                : moment(new Date(parseInt(item.created_at))).format('YYYY-MM-DD HH:mm:ss')
                              : '-'
            item.updated_at =
                            item.updated_at
                              ? item.updated_at.length < 13
                                ? moment(new Date(parseInt(item.updated_at * 1000))).format('YYYY-MM-DD HH:mm:ss')
                                : moment(new Date(parseInt(item.updated_at))).format('YYYY-MM-DD HH:mm:ss')
                              : '-'
            item.act = false
          })
          _this.tableDataAll = JSON.parse(JSON.stringify(_this.tableData))

          document.getElementById('content-box').scrollTop = 0
          _this.loading = false
        } else {
          _this.$message.error(response.message)
          _this.loading = false
        }
      })
        .catch(error => {
          console.log(error)
          _this.loading = false
        })
    },
    number (data, n) {
      var numbers = ''
      // 保留几位小数后面添加几个0
      for (var i = 0; i < n; i++) {
        numbers += '0'
      }
      var s = 1 + numbers
      // 如果是整数需要添加后面的0
      var spot = '.' + numbers
      // Math.round四舍五入
      //  parseFloat() 函数可解析一个字符串，并返回一个浮点数。
      var value = Math.round(parseFloat(data) * s) / s
      // 从小数点后面进行分割
      if (!value) {
        return value
      } else if (value.toString().indexOf('.') < 0) {
        return value.toFixed(18)
      }
      var d = value.toString().split('.')
      if (d.length === 1) {
        value = value.toString() + spot
        return value
      }
      if (d.length > 1) {
        if (d[1].length < n) {
          for (var tj = d[1].length; tj < n; tj++) {
            value = value.toString() + '0'
          }
        }
        return value
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
    // 查询
    search () {
      let _this = this
      _this.parma.limit = 4
      _this.parma.offset = 0
      _this.getData()
    },
    clearAll () {
      let _this = this
      _this.searchValue = ''
      _this.parma.limit = 4
      _this.parma.offset = 0
      _this.tableData = []
    }
  },
  mounted () {
    let _this = this
    document.getElementById('content-box').scrollTop = 0
    _this.$store.dispatch('setRouterMenu', 2)
    _this.$store.dispatch('setHeadertitle', _this.$t('route.Search_file'))
    // _this.getData()

    document.onkeydown = function (e) {
      if (e.keyCode === 13) {

      }
    }
  },
  filters: {
    priceFilter (value) {
      let realVal = ''
      if (!isNaN(value) && value !== '') {
        let tempVal = parseFloat(value).toFixed(19)
        realVal = tempVal.substring(0, tempVal.length - 1)
      } else {
        realVal = '-'
      }
      return realVal
    },
    NumFormat (value) {
      if (!value) return '-'
      return value
    }
  }
}
</script>

<style scoped lang="scss">
#dealManagement {
  padding: 0.25rem 0.2rem 0.5rem;
  .backTo {
    display: flex;
    align-items: center;
    font-size: 0.24rem;
    line-height: 1;
    margin: 0 0 0.2rem;
    cursor: pointer;
  }
  .upload {
    padding: 0.1rem 0.3rem 0.1rem 0.4rem;
    margin: 0 0 0.2rem;
    background-color: #fff;
    border-radius: 0.1rem;
    overflow: hidden;
    .title {
      display: flex;
      align-items: center;
      justify-content: space-between;
      width: 100%;
      margin: 0.1rem 0;
      text-align: left;
      font-size: 0.2999rem;
      font-weight: 500;
      color: #000;
      line-height: 0.42rem;
      // text-indent: 0.08rem;
      .statusStyle {
        display: inline-block;
        border: 1px solid;
        padding: 0 0.1rem;
        border-radius: 0.05rem;
        white-space: nowrap;
      }
      .el-button /deep/ {
        height: 0.42rem;
        padding: 0 0.17rem;
        line-height: 0.42rem;
        cursor: text;
        color: #fff;
        border: 1px solid;
        border-radius: 0.08rem;
        font-size: 0.1971rem;
        font-weight: normal;
      }
      .editStyle {
        display: inline-block;
        width: 15px;
        height: 15px;
        margin: 0 0.07rem;
        background: url(../../assets/images/edit.png) no-repeat;
        background-size: 100% 100%;
        font-size: 14px;
        cursor: pointer;
      }
    }
    .upload_form_detail {
      position: relative;
      display: flex;
      justify-content: center;
      // align-items: center;
      flex-wrap: wrap;
      margin: 0 0 0.1rem;
      .el-form /deep/ {
        // float: left;
        width: calc(65% - 0.45rem);
        margin: 0 0.3rem 0 0.15rem;
        .el-form-item {
          display: flex;
          align-items: center;
          // flex-wrap: wrap;
          width: auto;
          min-height: 32px;
          margin: 0 auto 0.08rem;
          .el-form-item__label {
            width: auto;
            max-width: 2rem;
            line-height: 1.5;
            text-align: left;
            white-space: nowrap;
            color: #404040;
          }
          .el-form-item__content {
            display: flex;
            align-items: center;
            flex-wrap: wrap;
            width: 100%;
            line-height: 1.5;
            .el-input {
              width: auto;
              .el-input__inner {
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
            .upload-demo {
              .el-upload--text {
                float: left;
                width: auto;
                height: auto;
                text-align: left;
                border: 0;
                .el-button--primary {
                  height: 0.32rem;
                  padding: 0 0.3rem;
                  line-height: 0.32rem;
                  background-color: transparent;
                  border: 1px solid #0b318f;
                  color: #0b318f;
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
              font-size: 0.12rem;
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
            .el-tag,
            .el-button--small {
              margin: 0 5px 0 0;
            }
          }
        }
      }
      .minerStyle {
        display: flex;
        align-items: center;
        width: 100%;
        margin: 0;
        overflow: hidden;
        span {
          float: left;
          width: auto;
          max-width: 2rem;
          line-height: 1.5;
          text-align: left;
          white-space: nowrap;
          color: #000;
          vertical-align: middle;
          font-size: 14px;
          line-height: 32px;
          padding: 0 0.12rem 0 0;
          -webkit-box-sizing: border-box;
          box-sizing: border-box;
        }
        .Leave_tip {
          float: left;
          line-height: 1.5;
          text-align: left;
          color: #000;
          vertical-align: middle;
          font-size: 14px;
          padding: 0;
          -webkit-box-sizing: border-box;
          box-sizing: border-box;
          display: flex;
          align-items: center;
          color: #0b318f;
          word-break: break-word;
          cursor: pointer;
          img {
            display: block;
            height: 0.17rem;
            margin: 0 0.05rem 0 0.01rem;
          }
          p {
            cursor: pointer;
          }
        }
        .el-button /deep/ {
          border: 0;
          line-height: 32px;
          padding: 0;
          color: #0b318f;
          font-size: 14px;
        }
        .el-button:hover {
          background: #fff;
        }
      }
      .upload_form_right {
        float: right;
        width: calc(35% - 0.3rem);
        margin: 0 0.3rem 0 0;
        color: #000;
        line-height: 2;
        h1 {
          font-size: 0.1972rem;
          font-weight: normal;
        }
        p {
          font-size: 14px;
        }
      }
      .el-tabs /deep/ {
        width: 100%;
        margin: 0.1rem 0 0;
        .el-tabs__header {
          padding: 0 0 0.15rem;
          margin: 0;
          border-bottom: 0.02rem solid #060d9d;
        }
        .el-tabs__nav {
          .el-tabs__active-bar {
            display: none;
          }
          .el-tabs__item {
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
          .el-tabs__item.is-active {
            background-color: #0a318e;
            color: #fff;
          }
          #tab-bids::after {
            position: absolute;
            content: "";
            left: -0.15rem;
            top: 50%;
            width: 0.02rem;
            height: 0.2rem;
            margin-top: -0.1rem;
            background-color: #060d9d;
          }
        }
        .el-tabs__nav-wrap::after {
          display: none;
        }
        .tabs_title {
          position: relative;
          display: flex;
          align-items: flex-end;
          justify-content: space-between;
          flex-wrap: wrap;
          margin-bottom: 0.15rem;
          font-size: 0.2571rem;
          font-weight: 500;
          line-height: 0.75rem;
          border-bottom: 1px solid #e0e0e0;
          color: #000;
          text-indent: 0.15rem;
          .tabs_title_left {
            display: flex;
            align-items: center;
          }
          .editStyle {
            display: inline-block;
            width: 15px;
            height: 15px;
            margin: 0 0.07rem;
            background: url(../../assets/images/edit.png) no-repeat;
            background-size: 100% 100%;
            font-size: 14px;
            cursor: pointer;
          }
          .el-button {
            float: right;
            padding: 0 0.11rem;
            font-size: 0.12rem;
            line-height: 18px;
            border-color: #e6eaf4;
            border-radius: 0.04rem;
            background-color: #e6eaf4;
            color: #0b318f;
          }
          .price {
            display: flex;
            align-items: center;
            font-size: 0.1714rem;
            color: #000;
            line-height: 0.25rem;
            .el-input {
              // width: 80px;
              .el-input__inner {
                width: calc(100% - 20px);
                padding: 0 10px;
                height: 30px;
                line-height: 30px;
              }
            }
          }
        }
      }
      .el-tag + .el-tag {
        margin-left: 10px;
      }
      .button-new-tag {
        margin-left: 10px;
        height: 32px;
        line-height: 30px;
        padding-top: 0;
        padding-bottom: 0;
      }
      .input-new-tag {
        width: 90px;
        margin-left: 10px;
        vertical-align: bottom;
      }
      .dateTime {
        position: absolute;
        right: 0;
        bottom: 0;
        text-align: right;
        p {
          font-size: 14px;
          color: #000;
        }
      }
    }
    .upload_bot {
      display: flex;
      justify-content: flex-end;
      align-items: center;
      margin: 0 0.3rem 0.15rem;
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
    .detail_content {
      padding: 0;
      margin: 0 0 0.1rem;
      background-color: #fff;
      border-radius: 0.1rem;
      .title_1 {
        font-size: 0.2571rem;
        font-weight: 500;
        line-height: 0.75rem;
        border-bottom: 1px solid #e0e0e0;
        color: #000;
        text-indent: 0.15rem;
      }
      .cont {
        padding: 0.2rem 0.15rem 0;
        // border-bottom: 1px solid #e0e0e0;
        overflow: hidden;
        .descrip {
          padding: 0 0 0.15rem;
          margin: 0 0 0.05rem;
          font-size: 0.1371rem;
          line-height: 0.3rem;
          color: #404040;
        }
        h3 {
          font-size: 0.1714rem;
          font-weight: 600;
          color: #000;
          line-height: 0.3rem;
        }
        .price_range {
          display: flex;
          // justify-content: space-between;
          align-items: center;
          flex-wrap: wrap;
          margin: 0 0 0.1rem;
          font-size: 0.1371rem;
          color: #737373;
          .range_content {
            margin-right: 0.15rem;
            .el-input /deep/ {
              float: left;
              width: 100%;
              .el-input__inner,
              .el-input-group__append {
                height: 0.28rem;
                border-color: #e3e3e3;
                font-size: 0.1372rem;
                color: #737373;
                line-height: 0.28rem;
              }
              .el-input__inner {
                border-top-left-radius: 0.04rem;
                border-bottom-left-radius: 0.04rem;
              }
              .el-input-group__append {
                height: 0.26rem;
                padding: 0 0.1rem;
                margin-top: -1px;
                border-color: #e3e3e3;
                background-color: transparent;
                border-top-right-radius: 0.04rem;
                border-bottom-right-radius: 0.04rem;
                // cursor: pointer;
                color: #737373;
                font-size: 0.1371rem;
                line-height: 1;
              }
            }
          }
          .sealed_btn {
            display: inline-block;
            padding: 0 0.18rem;
            border: 1px solid #0b318f;
            color: #0b318f;
            font-size: 0.1371rem;
            line-height: 0.2rem;
            border-radius: 0.04rem;
            cursor: pointer;
          }
          .sealed_content {
            width: 50%;
            font-size: 0.12rem;
            color: #737373;
            line-height: 1.1;
          }
        }
        .el-textarea {
          float: left;
          margin: 0 0 0.15rem;
        }
        .placeBid {
          display: inline-block;
          height: 0.42rem;
          margin: 0.15rem 0 0;
          padding: 0 0.25rem;
          background-color: #0b318f;
          line-height: 0.42rem;
          border-radius: 0.08rem;
          color: #fff;
          font-size: 0.1971rem;
          cursor: pointer;
        }
      }
      .cont_flex {
        display: flex;
        align-items: center;
        .el-input /deep/ {
          width: 60px;
          margin: 0 0.1rem;
          .el-input__inner {
            height: 0.28rem;
            border-color: #e3e3e3;
            font-size: 0.1372rem;
            color: #737373;
            line-height: 0.28rem;
          }
        }
      }
      .upload_bot {
        display: flex;
        justify-content: flex-end;
        align-items: center;
        margin: 0 0.3rem 0.15rem;
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
  }
  .readme_cont {
    height: 4.1rem;
    padding: 0.1rem 0.1rem;
    margin: 0.2rem auto;
    background-color: #fff;
    border-radius: 0.1rem;
    #content {
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
  .form {
    margin: 0 0 0.2rem;
    padding: 0.2rem 0.17rem;
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
    .form_table {
      margin: 0.24rem 0 0.1rem;
      .statusStyle {
        display: inline-block;
        border: 1px solid;
        padding: 0 0.1rem;
        border-radius: 0.05rem;
        white-space: nowrap;
      }
      .el-table /deep/ {
        overflow: visible;
        .el-table__body-wrapper,
        .el-table__header-wrapper {
          overflow: visible;
        }
        tr {
          border-radius: 0.08rem;
          // cursor: pointer;
          th {
            height: 0.5rem;
            padding: 0;
            background-color: #f2f2f2 !important;
            text-align: center;
            .cell {
              word-break: break-word;
              font-weight: 500;
              color: #737373;
            }
          }
          th.is-leaf {
            border-bottom: 0;
          }
          th:first-child {
            border-top-left-radius: 0.08rem;
            border-bottom-left-radius: 0.08rem;
          }
          th:last-child {
            border-top-right-radius: 0.08rem;
            border-bottom-right-radius: 0.08rem;
          }
          td {
            border-bottom: 1px solid #f2f2f2;
            .cell {
              padding: 0 0.05rem;
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
              .hot-cold-box {
                .el-button {
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
          left: -10px;
          right: 0px;
          top: 100%;
          background: #fff;
          border: 1px solid #f7f7f7;
          border-radius: 0.08rem;
          z-index: 1002;
          width: 120px;
          p {
            padding: 0.1rem 0;
            border-top: 1px solid #f7f7f7;
            font-size: 0.1372rem;
            color: #ffb822;
            line-height: 1;
          }
          p:nth-child(1) {
            border: 0;
          }
          p:nth-child(2) {
            color: #fd397a;
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
        padding: 0;
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
}
@media screen and (max-width: 1250px) {
  .price {
    width: 100%;
    span {
      display: block;
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
        }
      }
    }
  }
}
@media screen and (max-width: 999px) {
  #dealManagement {
    padding: 0.15rem 0.1rem 0.3rem;
    .upload {
      padding: 0.1rem;
      .title {
        word-break: break-word;
      }
      .upload_form {
        flex-wrap: wrap;
        .el-form /deep/ {
          width: 95%;
        }
      }
      .upload_form_detail {
        flex-wrap: wrap;
        .el-form /deep/ {
          width: calc(100% - 0.3rem);
          margin: auto 0.15rem;
        }
        .el-tabs /deep/ {
          .el-tabs__nav {
            .el-tabs__item {
              width: 1.5rem;
              margin-right: 0.1rem;
            }
            #tab-bids::after {
              display: none;
            }
          }
          .search {
            .el-select {
              width: 48%;
              margin: 0 1% 0.05rem;
            }
          }
        }
      }
    }
  }
}
@media screen and (max-width: 441px) {
  #dealManagement {
    .upload {
      .upload_form_detail {
        .minerStyle {
          flex-wrap: wrap;
          .Leave_tip {
            width: 100%;
          }
        }
      }
    }
  }
}
@media screen and (max-width: 321px) {
  #dealManagement .upload .title {
    font-size: 0.2rem;
    .el-button /deep/ {
      height: 0.35rem;
      line-height: 0.35rem;
      font-size: 0.16rem;
    }
  }
  #dealManagement .upload .upload_form_detail .el-tabs /deep/ {
    .el-tabs__nav {
      .el-tabs__item {
        width: 1.2rem;
        font-size: 14px;
      }
    }
  }
  #dealManagement .form .form_top .title p {
    float: none;
    clear: both;
    font-size: 0.16rem;
    line-height: 2;
    span {
      font-size: 0.16rem;
    }
  }
}
</style>
