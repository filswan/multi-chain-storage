<template>
  <div class="spaceStyle" v-loading="listLoad">
    <div class="slideScroll" v-if="listBucketIndex">
      <div class="form_top">
        <h4>{{$t('route.metaSpace')}}</h4>
        <div class="search_file">
          <div class="search_right">
            <el-input v-if="false" :placeholder="$t('metaSpace.search_bucket')" prefix-icon="el-icon-search" v-model="search" clearable @input="getListBuckets">
            </el-input>
            <div class="left_progress">
              <div class="need" v-if="languageMcs === 'en'">Need help? Join our
                <a :href="discord_link" target="_blank">Discord</a>
                or send an
                <a href="mailto:team@filswan.com">Email</a> to us.
              </div>
              <div class="need" v-else>需要帮助吗？加入我们的
                <a :href="discord_link" target="_blank">Discord</a>
                或发送
                <a href="mailto:team@filswan.com">电子邮件</a> 给我们。
              </div>
              <div class="progress">
                <el-progress :percentage="(free_bucket/free_bucketAll)*100 || 0"></el-progress>
                <span v-if="languageMcs === 'en'" class="tip">{{free_bucket | byteStorage}}GB of {{ $root.max_storage | byteStorage}}GB for Bucket storage</span>
                <span v-else class="tip">目前使用量：{{free_bucket | byteStorage}}GB（Bucket储存空间配额：{{ $root.max_storage | byteStorage}}GB）</span>
              </div>
            </div>
            <div class="createTask">
              <router-link :to="{name: 'ApiKey'}">
                <span>{{$t('billing.bill_btn_pay')}}</span>
              </router-link>
            </div>
          </div>
          <div class="createTask">
            <a @click="dialogFun('addNewBucket')">
              <img src="@/assets/images/space/icon_01.png" alt="">
              <span>{{$t('metaSpace.add_bucket')}}</span>
            </a>
          </div>
        </div>
      </div>
      <div class="fes-search">
        <div class="title">
          {{$t('metaSpace.list_bucket')}}
        </div>
        <el-table :data="listBuckets" stripe style="width: 100%" max-height="580" :empty-text="$t('deal.formNotData')">
          <el-table-column type="index" label="No." width="50"></el-table-column>
          <el-table-column prop="bucket_name" :label="$t('metaSpace.table_name')">
            <template slot-scope="scope">
              <div class="" v-if="!scope.row.is_active" @click="dialogFun('payActive')">
                <span style="opacity: 0.7;">{{ scope.row.bucket_name }}</span>
              </div>
              <div class="hot-cold-box" v-else @click="getListBucketDetail(scope.row)">
                <span style="text-decoration: underline;">{{ scope.row.bucket_name }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="subdomain" :label="$t('metaSpace.table_subdomain')">
            <template slot-scope="scope">
              <div class="hot-cold-box">
                <el-select v-model="domain.value" placeholder=" ">
                  <el-option v-for="item in domain.data" :key="item" :label="item" :value="item">
                  </el-option>
                </el-select>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="size" :label="$t('metaSpace.table_size')">
            <template slot-scope="scope">
              <div class="hot-cold-box">
                <span>{{ scope.row.size | formatbytes }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="created_at" :label="$t('metaSpace.table_createon')">
            <template slot-scope="scope">
              <div class="hot-cold-box">
                <span>{{ momentFun(scope.row.created_at) }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="" :label="$t('metaSpace.table_action')" min-width="120">
            <template slot-scope="scope">
              <div class="hot-cold-box">
                <!-- <i class="icon icon_pay " v-if="!scope.row.is_active" @click="dialogFun('pay', scope.row)"></i> -->
                <!-- <i class="icon icon_pay icon_pay_disable"></i> -->
                <i class="icon icon_rename" @click="dialogFun('rename', scope.row)"></i>
                <i class="icon icon_details" @click="getDetail(scope.row)"></i>
                <i class="icon icon_delete" @click="dialogFun('delete', scope.row)"></i>
              </div>
            </template>
          </el-table-column>
        </el-table>
        <p class="detailMore" v-if="languageMcs === 'en'">
          Want more Bucket storage? Click
          <a href="https://docs.filswan.com/multichain.storage/faq" target="_blank">here</a> for more details.
        </p>
        <p class="detailMore" v-else>
          想要更多的Bucket存储空间？单击
          <a href="https://docs.filswan.com/multichain.storage/faq" target="_blank">此处</a> 了解更多详细信息。
        </p>
      </div>
    </div>

    <div class="slideScroll" v-else>
      <div class="form_top">
        <h4>{{$t('route.metaSpace')}}</h4>
        <div class="search_file">
          <div class="search_right">
            <div class="left_progress">
              <div class="need" v-if="languageMcs === 'en'">Need help? Join our
                <a :href="discord_link" target="_blank">Discord</a>
                or send an
                <a href="mailto:team@filswan.com">Email</a> to us.
              </div>
              <div class="need" v-else>需要帮助吗？加入我们的
                <a :href="discord_link" target="_blank">Discord</a>
                或发送
                <a href="mailto:team@filswan.com">电子邮件</a> 给我们。
              </div>
              <div class="progress">
                <el-progress :percentage="(free_bucket/free_bucketAll)*100 || 0"></el-progress>
                <span v-if="languageMcs === 'en'" class="tip">{{free_bucket | byteStorage}}GB of {{ $root.max_storage | byteStorage}}GB for Bucket storage</span>
                <span v-else class="tip">目前使用量：{{free_bucket | byteStorage}}GB（Bucket储存空间配额：{{ $root.max_storage | byteStorage}}GB）</span>
              </div>
            </div>
            <div class="createTask">
              <router-link :to="{name: 'ApiKey'}">
                <span>{{$t('billing.bill_btn_pay')}}</span>
              </router-link>
            </div>
          </div>
          <div class="createTask">
            <a @click="dialogFun('add')">
              <img src="@/assets/images/space/icon_01.png" alt="">
              <span>{{$t('metaSpace.create_bucket')}}</span>
            </a>
          </div>
        </div>
      </div>
      <div class="fe-none">
        <p class="p_label">{{$t('metaSpace.empty_prompt')}}</p>
      </div>
    </div>

    <pop-ups v-if="dialogFormVisible" :dialogFormVisible="dialogFormVisible" :typeModule="typeName" :areaBody="areaBody" :createLoad="createLoad" :listTableLoad="listTableLoad" :backupLoad="backupLoad" :payLoad="payLoad" @getPopUps="getPopUps"></pop-ups>
  </div>
</template>
<script>
// import moment from 'moment'
import popUps from '@/components/popups'
// import Qs from 'qs'
let that
export default {
  data () {
    return {
      width: document.body.clientWidth > 600 ? '450px' : '95%',
      listLoad: true,
      createLoad: false,
      payLoad: false,
      listTableLoad: false,
      backupLoad: false,
      search: '',
      listBuckets: [],
      listBucketsAll: [],
      listBucketActive: 0,
      listBucketIndex: false,
      dialogFormVisible: false,
      areaBody: {},
      typeName: 'add',
      domain: {
        value: '',
        data: []
      },
      discord_link: process.env.DISCORD_LINK
    }
  },
  components: { popUps },
  methods: {
    async getDetail (row) {
      that.backupLoad = true
      that.dialogFun('detail', row)
      let bucketDetail = row
      that.dialogFun('detail', bucketDetail)
      that.backupLoad = false
    },
    async getPopUps (dialog, rows, bucketName) {
      // console.log('rows', rows)
      switch (rows) {
        case 'delete':
          that.deleteFun()
          break
        case 'rename':
          that.renameFun(bucketName)
          break
        case 'payClose':
          that.dialogFormVisible = dialog
          that.getListBuckets()
          break
        default:
          if (rows) that.getDialogClose(rows, bucketName)
          else that.dialogFormVisible = dialog
          break
      }
    },
    async dialogFun (name, row) {
      that.typeName = name
      that.areaBody = row || {}
      that.dialogFormVisible = true
    },
    async renameFun (newName) {
      that.createLoad = true
      const params =
        {
          'bucket_uid': that.areaBody.bucket_uid,
          'bucket_name': newName
        }
      const renameRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/bucket/rename`, 'post', params)
      if (!renameRes || renameRes.status !== 'success') {
        that.$message.error(renameRes.message || 'Fail')
      }
      await that.$commonFun.timeout(500)
      that.createLoad = false
      that.dialogFormVisible = false
      that.getListBuckets()
    },
    async deleteFun () {
      that.listTableLoad = true
      const params = {
        'bucket_uid': that.areaBody.bucket_uid
      }
      const deleteRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/bucket/delete?${Qs.stringify(params)}`, 'get')
      if (!deleteRes || deleteRes.status !== 'success') {
        that.$message.error(deleteRes.status || 'Fail')
      }
      await that.$commonFun.timeout(500)
      that.listTableLoad = false
      that.dialogFormVisible = false
      that.getListBuckets()
    },
    getListBucketDetail (row) {
      that.$router.push({ name: 'Space_detail', query: { folder: encodeURIComponent(row.bucket_name), bucket_uuid: row.bucket_uid, domain: that.domain.value } })
    },
    async getDialogClose (formName, name) {
      if (formName === 'pay') that.payLoad = true
      else that.createLoad = true
      const params = {
        'bucket_name': `${name.trim()}`
      }
      const directoryRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/bucket/create`, 'post', params)
      if (!directoryRes || directoryRes.status !== 'success') {
        that.$message.error(directoryRes.message || 'Fail')
      }
      await that.$commonFun.timeout(500)
      that.dialogFormVisible = false
      if (formName === 'pay') {
        that.payLoad = false
        that.dialogFun('success')
      } else {
        that.createLoad = false
        that.getListBuckets()
      }
    },
    async getListBuckets (name) {
      let size = 0
      let maxSize = 0
      that.listBucketActive = 0
      that.listLoad = true
      // const params = {
      //   file_name: `${that.search.trim()}` || ''
      // }
      const directoryRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/bucket/get_bucket_list`, 'get')
      if (!directoryRes || directoryRes.status !== 'success') {
        that.$message.error(directoryRes.message || 'Fail')
        that.listBuckets = []
      } else {
        that.listBuckets = directoryRes.data || []
        that.listBucketIndex = directoryRes.data.length > 0 || !!(that.search)
        that.listBuckets.forEach(element => {
          size += element.size
          if (element.is_active) {
            maxSize += element.max_size
            that.listBucketActive += 1
          }
        })
        that.$store.dispatch('setFreeBucket', size || 0)
        that.$store.dispatch('setFreeBucketAll', maxSize || 0)
        if (!that.listBucketIndex) that.$store.dispatch('setFreeBucket', 0)
      }

      const domainRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/gateway/get_gateway`, 'get')
      if (!domainRes || domainRes.status !== 'success') {
        that.$message.error(domainRes.message || 'Fail')
        that.domain = {
          value: '',
          data: []
        }
      } else {
        that.domain.value = domainRes.data ? domainRes.data[0] : ''
        that.domain.data = domainRes.data || []
      }

      that.listLoad = false
    },
    momentFun (dateItem) {
      let dateNew = new Date(dateItem).getTime()
      let dataUnit = ''
      let dataTime = new Date(dateNew) + ''
      let dataUnitIndex = dataTime.indexOf('GMT')
      let dataUnitArray = dataTime.substring(dataUnitIndex, dataUnitIndex + 8)
      switch (dataUnitArray) {
        case 'GMT+1000':
          dataUnit = 'UTC+10'
          break
        case 'GMT-1000':
          dataUnit = 'UTC-10'
          break
        case 'GMT+0000':
          dataUnit = 'UTC+0'
          break
        default:
          dataUnit = dataUnitArray ? dataUnitArray.replace(/0/g, '').replace('GMT', 'UTC') : '-'
          break
      }
      dateNew = dateNew
        ? moment(new Date(parseInt(dateNew))).format('YYYY-MM-DD HH:mm:ss') + ` (${dataUnit})`
        : '-'
      return dateNew
    },
    copyLink (text) {
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
        var msg = successful ? 'Copied to clipboard!' : 'copy failed!'
        // console.log('Copying text command was ' + msg)
        if (successful) {
          that.$message({
            message: msg,
            type: 'success'
          })
        }
      } catch (err) {
        console.log('Oops, unable to copy')
      } finally {
        document.body.removeChild(txtArea)
      }
    }
  },
  mounted () {
    that = this
    document.getElementById('content-box').scrollTop = 0
    that.$store.dispatch('setRouterMenu', 20)
    that.$store.dispatch('setHeadertitle', that.$t('route.metaSpace'))
    that.getListBuckets()
  },
  filters: {
    NumFormat (value) {
      if (!value) return '-'
      return value
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
      return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
      // return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i]
    },
    byteStorage (limit) {
      // 只转换成GB
      if (limit <= 0) {
        return '0'
      } else if (limit <= 10737419) { // 10737419 ~= 0.01GB
        return '0.01'
      } else {
        // return (limit/( 1024 * 1024 * 1024)).toPrecision(2)  //or 1000
        let value = limit / (1024 * 1024 * 1024)
        let v1 = String(value).split('.')
        let v2 = v1[1] || ''
        let v3 = String(v2).replace(/(0+)\b/gi, '')
        if (v3) {
          return v1[0] + '.' + v3.slice(0, 2)
        } else {
          return v1[0]
        }
      }
    }
  },
  computed: {
    languageMcs () {
      return this.$store.getters.languageMcs
    },
    free_bucket () {
      let freeBucket = this.$store.getters.free_bucket < this.$store.getters.free_bucketAll ? this.$store.getters.free_bucket : this.$store.getters.free_bucketAll
      return freeBucket
    },
    free_bucketAll () {
      return this.$store.getters.free_bucketAll === 0 ? 34359738368 : this.$store.getters.free_bucketAll
    }
  }
}
</script>
<style lang="scss" scoped>
.el-dialog__wrapper /deep/ {
  display: flex;
  align-items: center;
  .customStyle {
    margin: auto !important;
    box-shadow: 0 0 13px rgba(128, 128, 128, 0.8);
    border-radius: 0.2rem;
    .el-dialog__header {
      padding: 0.3rem 0.4rem;
      display: flex;
      align-items: center;
      justify-content: space-between;
      border-bottom: 1px solid #dfdfdf;
      color: #000;
      font-size: 0.22rem;
      font-weight: 500;
      line-height: 1;
      text-transform: capitalize;
      @media screen and (max-width: 479px) {
        padding: 0.3rem 0.2rem;
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
      padding: 0.3rem 0.4rem;
      font-size: 0.2rem;
      @media screen and (max-width: 479px) {
        padding: 0.2rem;
      }
      label {
        word-break: break-word;
        line-height: 1;
        color: #666;
        font-size: inherit;
      }
      .address {
        background: rgba(233, 233, 233, 1);
        padding: 8px;
        margin: 10px 0 22px;
        border-radius: 8px;
        font-size: inherit;
      }
      .share {
        display: flex;
        align-items: center;
        font-size: inherit;
        .el-button {
          padding: 0 20px 0 0;
          background: transparent !important;
          border: 0;
          color: #4f7bf5;
          font-size: inherit;
          font-weight: normal;
          font-family: inherit;
          opacity: 0.8;
          span {
            display: flex;
            align-items: center;
            svg {
              width: 15px;
              height: 15px;
              margin: 0 3px 0 0;
            }
          }
          &:hover {
            background: transparent;
            opacity: 1;
          }
        }
      }
    }
    .dialog-footer {
      display: flex;
      justify-content: flex-end;
    }
  }
}
.spaceStyle /deep/ {
  position: relative;
  width: calc(100% - 0.6rem);
  padding: 0 0 0.4rem;
  margin: 0.3rem;
  background-color: #fff;
  border-radius: 0.1rem;
  @media screen and (max-width: 600px) {
    width: 100%;
    padding: 0;
    margin: 0.3rem 0;
  }
  .el-loading-mask {
    z-index: 5;
  }
  .slideScroll {
    height: calc(100% - 0.6rem);
    min-height: 350px;
    padding: 0.3rem;
    .form_top {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      margin: 0 auto 0.3rem;
      h4 {
        width: 100%;
        margin: 0 0 0.3rem;
        font-size: 0.22rem;
        font-weight: normal;
        color: #000;
        line-height: 1.5;
      }
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

        .el-button {
          height: 0.34rem;
          padding: 0 0.4rem;
          margin: 0 0.1rem;
          color: #fff;
          line-height: 0.34rem;
          font-size: 0.1372rem;
          border: 0;
          border-radius: 0.08rem;
        }

        .el-input {
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
        flex-wrap: wrap;
        width: 100%;
        margin: 0;
        p {
          font-size: 0.13rem;
          color: #222;
        }
        .createTask {
          border-radius: 0.1rem;
          cursor: pointer;
          a {
            display: flex;
            align-items: center;
            justify-content: center;
            min-width: 1.6rem;
            padding: 0.13rem;
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
            white-space: nowrap;
            img {
              display: inline-block;
              height: 0.25rem;
              margin: 0 0.1rem 0 0;
              @media screen and (max-width: 1260px) {
                height: 16px;
              }
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
          @media screen and (max-width: 600px) {
            flex-wrap: wrap;
            width: 100%;
            margin-bottom: 0.2rem;
          }
          .el-button {
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
          .left_progress {
            padding: 0 0.28rem 0 0;
            @media screen and (max-width: 600px) {
              width: 100%;
              padding: 0 0 0.2rem;
            }
            .need {
              font-size: 12px;
              color: #000;
              text-align: center;
              line-height: 1.1;
              @media screen and (min-width: 1800px) {
                font-size: 14px;
              }
              @media screen and (max-width: 600px) {
                text-align: left;
              }
              a {
                color: inherit;
                text-decoration: underline;
              }
            }
            .progress {
              margin: 0.1rem 0 0;
              .el-progress {
                font-size: 12px;
                .el-progress-bar {
                  padding-right: 0;
                  margin: 0 0 5px 0;
                }
                .el-progress__text {
                  display: none;
                }
              }
              .tip {
                display: block;
                font-size: 12px;
                line-height: 1.2;
                color: rgb(179, 192, 231);
                @media screen and (min-width: 1800px) {
                  font-size: 14px;
                }
              }
            }
          }
        }
        .el-input {
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
    .fe-none {
      display: flex;
      justify-content: center;
      align-items: center;
      height: auto;
      padding: 1.5rem 5% 0.5rem;
      @media screen and (max-width: 600px) {
        padding: 0.5rem 5%;
      }
      .p_label {
        padding: 0.15rem 0.3rem;
        font-size: 0.25rem;
        color: #4f87ff;
        line-height: 1.2;
        border: dashed;
        border-radius: 0.1rem;
        @media screen and (max-width: 1600px) {
          font-size: 0.23rem;
        }
      }
    }
    .fes-header {
      display: flex;
      justify-content: space-between;
      align-content: center;
      width: calc(100% - 0.4rem);
      padding: 0.25rem 0.2rem;
      img {
        width: auto;
        max-width: 100%;
        height: 0.35rem;
      }
      svg {
        font-size: 0.25rem;
        width: 0.25rem;
        height: 0.25rem;
        margin: 0.05rem 0 0;
        color: #333;
        cursor: pointer;
      }
      h2 {
        margin: 10px 0 0 13px;
        font-weight: 400;
        color: #333;
        font-size: 0.2rem;
      }
    }
    .fs3_backup {
      margin: 0;
      .introduce {
        margin: 0.1rem 0 0.05rem;
        text-indent: 0.2rem;
        background: #002a39;
        // font-family: 'm-semibold';
        font-weight: bold;
        a {
          display: block;
          line-height: 3;
          font-size: 0.14rem;
          color: #333;
          &:hover {
            color: #2f85e5;
          }
          @media screen and (max-width: 999px) {
            font-size: 13px;
            line-height: 3.5;
          }
        }
        .active {
          color: #2f85e5;
        }
      }
      .introRouter {
        font-size: 0.14rem;
        @media screen and (max-width: 999px) {
          font-size: 13px;
        }
        a {
          display: block;
          padding: 0.07rem 0.2rem;
          color: #333;
          font-size: inherit;
          @media screen and (max-width: 999px) {
            padding: 8px 0.2rem;
          }
          &:hover {
            color: #7ecef4;
            background-color: rgba(0, 0, 0, 0.1);
          }
        }
        .active {
          color: #7ecef4;
        }
      }
      .el-tree {
        padding: 0 0.25rem;
        background: transparent;
        color: #333;
        .el-tree-node {
          .el-tree-node__content {
            height: auto;
            background: transparent !important;
            margin: 0 0 0.08rem;
            .el-tree-node__expand-icon {
              padding: 0 0.05rem;
              &:before {
                font-size: 0.2rem;
              }
            }
            .el-tree-node__label {
              font-size: 0.15rem;
              @media screen and (max-width: 999px) {
                font-size: 14px;
              }
            }
            &:hover {
              color: #5f9dcc;
            }
          }
          .el-tree-node__children {
            .el-tree-node__content {
              .el-tree-node__label {
                font-size: 0.14rem;
                line-height: 1.5;
                @media screen and (max-width: 999px) {
                  font-size: 13px;
                }
              }
            }
          }
          .is-current,
          .is-checked {
            color: #5f9dcc;
          }
        }
      }
    }
    .fes-search {
      // height: calc(100% - 1.7rem);
      .title {
        display: flex;
        align-items: center;
        justify-content: flex-start;
        margin: 0 0.3rem 0.25rem;
        font-size: 16px;
        color: #000;
        line-height: 1;
        @media screen and (max-width: 1600px) {
          font-size: 14px;
        }
        img {
          display: block;
          width: 22px;
          height: 22px;
          margin: 0 0 0 5px;
          cursor: pointer;
          @media screen and (max-width: 1680px) {
            width: 20px;
            height: 20px;
          }
          @media screen and (max-width: 1260px) {
            width: 17px;
            height: 17px;
          }
        }
      }
      .el-table {
        overflow: visible;
        font-size: 0.18rem;
        .el-loading-mask {
          .el-loading-spinner {
            top: 50%;
          }
        }
        .el-table__body-wrapper,
        .el-table__header-wrapper {
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

          th.is-leaf {
            border-bottom: 0;
          }

          td {
            padding: 0.25rem 0.05rem;
            border-bottom: 1px solid #dfdfdf;

            .cell {
              padding: 0;
              font-size: 0.18rem;
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
                display: flex;
                align-items: center;
                justify-content: center;
                flex-wrap: wrap;
                .icon {
                  display: block;
                  width: 22px;
                  height: 22px;
                  margin: 0 0.15rem;
                  font-size: 0.24rem;
                  @media screen and (max-width: 1600px) {
                    width: 18px;
                    height: 18px;
                  }
                  @media screen and (max-width: 768px) {
                    width: 15px;
                    height: 15px;
                    margin: 0 5px;
                  }
                  &:hover {
                    opacity: 0.7;
                  }
                }
                .icon_pay {
                  width: 25px;
                  height: 25px;
                  background: url(../../assets/images/space/pay.png) no-repeat
                    center;
                  background-size: 100% 100%;
                  @media screen and (max-width: 1600px) {
                    width: 22px;
                    height: 22px;
                  }
                  @media screen and (max-width: 768px) {
                    width: 20px;
                    height: 20px;
                  }
                  &.icon_pay_disable {
                    opacity: 0.2;
                  }
                }
                .icon_rename {
                  background: url(../../assets/images/space/icon_02.png)
                    no-repeat center;
                  background-size: 100% 100%;
                }
                .icon_details {
                  background: url(../../assets/images/space/icon_03.png)
                    no-repeat center;
                  background-size: 100% 100%;
                }
                .icon_delete {
                  background: url(../../assets/images/space/icon_04.png)
                    no-repeat center;
                  background-size: 100% 100%;
                }
                .el-select {
                  width: 100%;
                  .el-input {
                    .el-input__inner {
                      height: auto;
                      line-height: 2;
                      font-size: inherit;
                    }
                    .el-icon-arrow-up {
                      display: flex;
                      align-items: center;
                    }
                  }
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
      }
      .el-table::before {
        display: none;
      }
      .detailMore {
        padding: 15px 0;
        font-size: 12px;
        color: #888;
        line-height: 1.5;
        text-align: center;
        @media screen and (min-width: 1800px) {
          font-size: 14px;
        }
        a {
          color: inherit;
          text-decoration: underline;
        }
      }
    }
    .fes-host {
      display: flex;
      justify-content: space-between;
      position: absolute;
      left: 0;
      bottom: 0;
      z-index: 21;
      background-color: rgba(0, 0, 0, 0.1);
      font-size: 15px;
      font-weight: 400;
      // width: calc(3.2rem - 0.4rem);
      width: calc(100% - 0.4rem);
      padding: 0.2rem;
      color: hsla(0, 0%, 100%, 0.75);
      transition: all;
      transition-duration: 0.3s;
      i {
        float: left;
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 0.2rem;
        margin-right: 10px;
        width: 20px;
        height: 20px;
        color: #888b83;
        // background: url(../assets/images/icon_01.jpg) no-repeat center;
        // background-size: 100%;
        @media screen and (max-width: 999px) {
          font-size: 18px;
        }
      }
      a {
        color: hsla(0, 5%, 12%, 0.75);
        font-size: 15px;
        font-weight: 400;
      }
      .fesHostLink {
        display: flex;
        width: calc(100% - 40px);
        a {
          display: block;
          width: calc(100% - 30px);
          overflow: hidden;
          text-overflow: ellipsis;
          white-space: nowrap;
        }
      }
      .fesHostLogout {
        cursor: pointer;
        i {
          margin: 0;
          color: #333;
          &:hover {
            color: #7ecef4;
          }
        }
      }
    }
  }
  .el-dialog__wrapper {
    .policyStyle {
      width: 90%;
      max-width: 600px;
      .el-dialog__header {
        display: flex;
        .el-dialog__title {
          font-size: 0.15rem;
          color: #333;
        }
      }
      .el-dialog__body {
        padding: 0.3rem 0;
        .el-form-item {
          padding: 0.1rem 0.3rem;
          margin-bottom: 0.1rem;
          &:nth-child(2n + 2) {
            background-color: #f7f7f7;
          }
        }
        .el-form-item__content {
          display: flex;
          line-height: 0.3rem;
          .el-input {
            .el-input__inner {
              height: 0.3rem;
              padding-left: 0;
              line-height: 0.3rem;
              border: 0;
              border-bottom: 1px solid #f7f7f7;
              background-color: transparent;
              border-radius: 0;
              font-size: 0.13rem;
              color: #32393f;
              text-align: left;
            }
            .el-input__icon {
              display: flex;
              align-items: center;
            }
          }
          .el-input.is-disabled {
            .el-input__inner {
              background-color: transparent;
            }
          }
          .el-select {
            margin: 0 5%;
          }
          .el-button {
            width: 130px;
            height: 0.3rem;
            padding: 0;
            line-height: 0.3rem;
            color: #333;
            font-size: 12px;
            font-family: inherit;
            border: 0;
            border-radius: 0.02rem;
            text-align: center;
          }
        }
      }
    }
  }
}

@media screen and (max-width: 1024px) {
  .spaceStyle {
    .slideScroll {
      .form_top {
        .search {
          flex-wrap: wrap;
          height: auto;

          .el-input {
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
            .el-button {
              padding: 0 0.2rem;
              font-size: 0.1372rem;
            }
          }
        }
      }
    }
  }
}
@media screen and (max-width: 470px) {
  .spaceStyle {
    .slideScroll {
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
</style>
