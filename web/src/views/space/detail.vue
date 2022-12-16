<template>
  <div class="fs3_body">
    <div class="slideScroll">
      <header class="fe-header">
        <h2>
          <a @click="getListBucketMain()">Buckets</a>
          >
          <span class="main" v-for="(item, index) in currentBucket" :key="index">
            <b v-if="index>0">>&nbsp;</b>
            <a href="javascript:;" style="word-break: break-word;" @click="getListBucketMain(item, true, index)">{{item}}</a>
          </span>
        </h2>
      </header>
      <div class="fes-search" v-loading="listLoad">
        <div class="form_top">
          <div class="search_file">
            <div class="search_left">
              <div class="icon">{{detail.size | formatbytes}}</div>
              <div class="icon">{{detail.object}} Object</div>
            </div>
            <div class="createTask">
              <a @click="dialogFun('addSub')">
                <img src="@/assets/images/space/icon_01.png" alt="">
                <span>{{$t('metaSpace.create_folder')}}</span>
              </a>
              <a @click="dialogFun('upload', $route.query.bucket_uuid)">
                <img src="@/assets/images/space/icon_08.png" alt="">
                <span>{{$t('metaSpace.Upload_File')}}</span>
              </a>
              <a @click="dialogFun('upload_folder', $route.query.bucket_uuid)">
                <img src="@/assets/images/space/icon_08.png" alt="">
                <span>Upload Folder</span>
              </a>
            </div>
          </div>
        </div>
        <el-row :gutter="15" class="table_body">
          <el-col :xs="24" :sm="24" :md="24" :lg="14" class="left">
            <el-table :data="listData.listBucketFile" ref="tableList" stripe style="width: 100%" max-height="400" empty-text=" ">
              <el-table-column type="index" label="No." width="50"></el-table-column>
              <el-table-column prop="Name" :label="$t('metaSpace.table_name')">
                <template slot-scope="scope">
                  <div class="hot-cold-box">
                    <span @click="dialogFun('detail_file', scope.row)">{{ scope.row.Name }}</span>
                  </div>
                </template>
              </el-table-column>
              <el-table-column prop="PayloadCid" min-width="120">
                <template slot="header" slot-scope="scope">
                  <div class="tips" style="white-space: nowrap;">
                    {{$t('metaSpace.table_cid')}}

                    <el-popover placement="top" popper-class="elPopTitle" width="200" trigger="hover" :content="$t('metaSpace.table_cid_tip')">
                      <img slot="reference" src="@/assets/images/info.png" />
                    </el-popover>
                  </div>
                </template>
                <template slot-scope="scope">
                  <div class="hot-cold-box">
                    <div class="copyText">
                      <span>{{ scope.row.PayloadCid||'-' }}</span>
                      <img class="imgCopy" src="@/assets/images/space/icon_10.png" @click="copyLink(scope.row.PayloadCid)" v-if="scope.row.PayloadCid" alt="">
                    </div>
                  </div>
                </template>
              </el-table-column>
              <el-table-column prop="PinStatus" :label="$t('metaSpace.table_status')">
                <template slot-scope="scope">
                  <div class="hot-cold-box">
                    <span>{{ scope.row.PinStatus||'-' }}</span>
                  </div>
                </template>
              </el-table-column>
              <el-table-column prop="UpdatedAt" :label="$t('metaSpace.table_LastModified')">
                <template slot-scope="scope">
                  <div class="hot-cold-box">
                    <p>{{ momentFun(scope.row.UpdatedAt) }}</p>
                  </div>
                </template>
              </el-table-column>
              <el-table-column prop="Size" :label="$t('metaSpace.table_size')">
                <template slot-scope="scope">
                  <div class="hot-cold-box">
                    <p>{{ scope.row.Size | formatbytes }}</p>
                  </div>
                </template>
              </el-table-column>
              <el-table-column prop="" :label="$t('metaSpace.table_action')" min-width="120">
                <template slot-scope="scope">
                  <div class="hot-cold-box">
                    <i class="icon icon_share" @click="copyLink(`${scope.row.IpfsUrl}?filename=${scope.row.Name}`, 1)"></i>
                    <i class="icon icon_details" @click="getDetail('detail_file', scope.row)"></i>
                    <i class="icon icon_delete" @click="dialogFun('delete', scope.row)"></i>
                  </div>
                </template>
              </el-table-column>
            </el-table>
          </el-col>
          <el-col :xs="24" :sm="24" :md="24" :lg="10" class="left">
            <el-table :data="listData.listBucketFolder" ref="tableList" stripe style="width: 100%" max-height="400" empty-text=" ">
              <el-table-column type="index" label="No." width="50"></el-table-column>
              <el-table-column prop="Name" :label="$t('metaSpace.table_name')">
                <template slot-scope="scope">
                  <div class="hot-cold-box">
                    <span style="text-decoration: underline;" @click="getListBucketMain(scope.row.Name)">{{ scope.row.Name }}</span>
                  </div>
                </template>
              </el-table-column>
              <el-table-column prop="UpdatedAt" :label="$t('metaSpace.table_LastModified')" width="100">
                <template slot-scope="scope">
                  <div class="hot-cold-box">
                    <p>{{ momentFun(scope.row.UpdatedAt) }}</p>
                  </div>
                </template>
              </el-table-column>
              <el-table-column prop="Size" :label="$t('metaSpace.table_size')">
                <template slot-scope="scope">
                  <div class="hot-cold-box">
                    <p>{{ scope.row.Size | formatbytes }}</p>
                  </div>
                </template>
              </el-table-column>
              <el-table-column prop="" :label="$t('metaSpace.table_action')" width="120">
                <template slot-scope="scope">
                  <div class="hot-cold-box">
                    <i class="icon icon_rename" @click="dialogFun('rename', scope.row)"></i>
                    <i class="icon icon_details" @click="getDetail('detail_folder', scope.row)"></i>
                    <i class="icon icon_delete" @click="dialogFun('delete', scope.row)"></i>
                  </div>
                </template>
              </el-table-column>
            </el-table>
          </el-col>
        </el-row>
        <div class="fe-none" v-if="listData.listBuckets&&listData.listBuckets.length<=0">
          <p class="p_label">{{$t('metaSpace.empty_prompt_detail')}}</p>
        </div>
      </div>
    </div>

    <pop-ups v-if="dialogFormVisible" :dialogFormVisible="dialogFormVisible" :typeModule="typeName" :areaBody="areaBody" :createLoad="createLoad" :listTableLoad="listTableLoad" :currentBucket="url" :backupLoad="backupLoad" @getPopUps="getPopUps"
      @getUploadDialog="getUploadDialog"></pop-ups>
  </div>

</template>

<script>
import moment from 'moment'
import QS from 'qs'
import popUps from '@/components/popups'
let that
export default {
  name: 'Space',
  data () {
    return {
      listLoad: true,
      listTableLoad: false,
      backupLoad: false,
      url: '',
      currentBucket: [],
      dialogFormVisible: false,
      createLoad: false,
      listData: {
        listBuckets: [],
        listBucketFile: [],
        listBucketFolder: []
      },
      areaBody: {},
      typeName: 'delete',
      detail: {
        object: 0,
        size: 0
      }
    }
  },
  components: {
    popUps
  },
  methods: {
    async getDetail (type, row) {
      that.backupLoad = true
      that.dialogFun(type, row)
      let bucketDetail = row
      let params = {
        file_id: row.ID
      }
      const infoRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/oss_file/get_file_info?${QS.stringify(params)}`, 'get')
      if (!infoRes || infoRes.status !== 'success') that.$message.error(infoRes.message || 'Fail')
      else {
        bucketDetail.Name = infoRes.data.Name
        bucketDetail.CreatedAt = infoRes.data.CreatedAt
        bucketDetail.Size = infoRes.data.Size
        bucketDetail.IpfsUrl = `${infoRes.data.IpfsUrl}?filename=${infoRes.data.Name}`
        bucketDetail.PayloadCid = infoRes.data.PayloadCid
      }
      that.dialogFun(type, bucketDetail)
      that.backupLoad = false
    },
    getListBucketMain (fileName, type, index) {
      let fold = ''
      if (type && index) fold = (that.currentBucket.slice(0, index + 1)).join('/')
      else if (type && index === 0) fold = that.currentBucket[0]
      else fold = `${that.url}/${fileName}`

      if (fileName) that.$router.push({ name: 'Space_detail', query: { folder: encodeURIComponent(fold), bucket_uuid: that.$route.query.bucket_uuid } })
      else that.$router.push({ name: 'Space' })
    },
    async getPopUps (dialog, rows, bucketName) {
      switch (rows) {
        case 'delete':
          that.deleteFun()
          break
        case 'rename':
          that.renameFun(bucketName)
          break
        default:
          if (rows) that.getDialogClose(rows, bucketName)
          else that.dialogFormVisible = dialog
          break
      }
    },
    async getDialogClose (formName, name) {
      that.createLoad = true
      const current = that.currentBucket.slice(1).join('/')
      const params = {
        'file_name': `${name.trim()}`,
        'prefix': current || '',
        'bucket_uid': that.$route.query.bucket_uuid
      }
      const directoryRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/oss_file/create_folder`, 'post', params)
      if (!directoryRes || directoryRes.status !== 'success') {
        that.$message.error(directoryRes ? directoryRes.message : 'Fail')
      }
      await that.$commonFun.timeout(500)
      that.createLoad = false
      that.dialogFormVisible = false
      that.getListObjects()
    },
    async dialogFun (name, row) {
      that.typeName = name
      if (row) that.areaBody = row
      that.dialogFormVisible = true
    },
    getUploadDialog (dialog, rows) {
      that.dialogFormVisible = dialog
      if (rows) {
        that.getListObjects(1)
      }
    },
    async renameFun (newName) {
      that.createLoad = true
      const params =
        {
          'action': 'rename',
          'src': {
            'dirs': [that.areaBody.id],
            'items': []
          },
          'new_name': newName
        }
      const renameRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v3/object/rename`, 'post', params)
      if (!renameRes || renameRes.status !== 'success') {
        that.$message.error(renameRes ? renameRes.message : 'Fail')
      }
      await that.$commonFun.timeout(500)
      that.createLoad = false
      that.dialogFormVisible = false
      that.getListObjects()
    },
    async deleteFun () {
      that.listTableLoad = true
      const params = {
        'file_id': that.areaBody.ID
      }
      const deleteRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/oss_file/delete?${QS.stringify(params)}`, 'get')
      if (!deleteRes || deleteRes.status !== 'success') {
        that.$message.error(deleteRes ? deleteRes.status : 'Fail')
      }
      await that.$commonFun.timeout(500)
      that.listTableLoad = false
      that.dialogFormVisible = false
      that.getListObjects()
    },
    copyLink (text, tip) {
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
        var msg = successful ? tip ? 'IPFS link is copied on the clip board!' : 'Copied to clipboard!' : 'copy failed!'
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
    },
    async getListObjects (type) {
      that.listLoad = true
      that.detail.size = 0
      const current = that.currentBucket.slice(1).join('/')
      let params = {
        prefix: current || '',
        bucket_uid: that.$route.query.bucket_uuid,
        limit: 10,
        offset: 0
      }
      const directoryRes = await that.$commonFun.sendRequest(`${process.env.BASE_PAYMENT_GATEWAY_API}api/v2/oss_file/get_file_list?${QS.stringify(params)}`, 'get')
      if (!directoryRes || directoryRes.status !== 'success') {
        that.$message.error(directoryRes ? directoryRes.message : 'Fail')
        return false
      }
      that.listData = {
        listBuckets: directoryRes.data || [],
        listBucketFile: [],
        listBucketFolder: []
      }
      that.listData.listBuckets.forEach((element, i) => {
        element.index = i
        that.detail.size += element.Size
        if (element.IsFolder) that.listData.listBucketFolder.push(element)
        else that.listData.listBucketFile.push(element)
      })
      that.detail.object = directoryRes.data.length || 0
      await that.$commonFun.timeout(500)
      that.listLoad = false
      that.$store.dispatch('setFreeBucket', that.detail.size || 0)
      that.$refs.tableList.bodyWrapper.scrollTop = 0
      // if (type === 1) that.$refs.tableList.bodyWrapper.scrollTop = that.$refs.tableList.bodyWrapper.scrollHeight
    },
    momentFun (dateItem) {
      let dateNew = new Date(dateItem).getTime()
      let dataUnit = ''
      let dataTime = new Date(dateNew) + ''
      let dataUnitIndex = dataTime.indexOf('GMT')
      let dataUnitArray = dataTime.substring(dataUnitIndex, dataUnitIndex + 8)
      switch (dataUnitArray) {
        case 'GMT+1000':
          dataUnit = 'GMT+10'
          break
        case 'GMT-1000':
          dataUnit = 'GMT-10'
          break
        case 'GMT+0000':
          dataUnit = 'GMT+0'
          break
        default:
          dataUnit = dataUnitArray ? dataUnitArray.replace(/0/g, '') : '-'
          break
      }
      dateNew = dateNew
        ? moment(new Date(parseInt(dateNew))).format('YYYY-MM-DD HH:mm:ss') + ` (${dataUnit})`
        : '-'
      return dateNew
    },
    init () {
      if (!that.$route.query.folder) {
        that.$router.push({ name: 'Space' })
        return false
      }
      that.url = decodeURIComponent(that.$route.query.folder)
      that.currentBucket = that.url.split('/')
      that.getListObjects()
    },
    async sort (data) {
      return data.sort((a, b) => { return b.index - a.index })
    }
  },
  watch: {
    $route: function (to, from) {
      // console.log(to, from)
      that.init()
    }
  },
  filters: {
    formatbytes: function (bytes) {
      if (bytes === 0) return '0 B'
      if (!bytes) return '-'
      let k = 1024 // or 1000
      let sizes = ['bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB']
      let i = Math.floor(Math.log(bytes) / Math.log(k))

      if (Math.round((bytes / Math.pow(k, i))).toString().length > 3) {
        i += 1
      }
      return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
      // return (bytes / Math.pow(k, i)).toPrecision(3) + ' ' + sizes[i]
    },
    slideName: function (name) {
      if (!name) return '-'
      let retName = that.prefixName ? name.replace(that.prefixName + '/', '') : name
      return retName
    },
    hiddAddress: function (val) {
      if (val) {
        return `${val.substring(0, 6)}...${val.substring(val.length - 4)}`
      } else {
        return '-'
      }
    }
  },
  mounted () {
    that = this
    document.getElementById('content-box').scrollTop = 0
    that.$store.dispatch('setRouterMenu', 20)
    that.$store.dispatch('setHeadertitle', that.$t('route.metaSpace'))
    that.init()
  }
}
</script>

<style lang="scss" scoped>
.fs3_body {
  position: relative;
  width: calc(100% - 1.2rem);
  padding: 0.3rem;
  margin: 0.3rem;
  background-color: #fff;
  border-radius: 0.1rem;
  .slideScroll {
    .fe-header {
      position: relative;
      padding: 0.2rem 0 0;
      @media screen and (max-width: 441px) {
        padding: 0.2rem 0 0;
      }
      h2 {
        width: calc(100% - 0.6rem);
        font-size: 16px;
        font-weight: 400;
        margin: 0 auto 0.35rem;
        display: flex;
        align-items: center;
        @media screen and (max-width: 1440px) {
          font-size: 14px;
        }
        a {
          margin-right: 5px;
          cursor: pointer;
          &:hover {
            text-decoration: underline;
          }
        }
        span {
          display: flex;
          align-items: center;
          margin-left: 5px;
          color: #4f84ff;
          font-size: inherit;
        }
        .fe-edit {
          font-size: 0.2rem;
          color: #46a5e0;
          margin-left: 4px;
          i {
            font-size: 0.2rem;
          }
        }
        .el-input /deep/ {
          .el-input__inner {
            padding: 0;
            color: #46a5e0;
            font-size: 0.15rem;
            border: 0;
            border-bottom: 1px solid #dcdfe6;
            border-radius: 0;
          }
        }
      }
    }
    .fes-search {
      // height: calc(100% - 1.7rem);
      .form_top {
        display: flex;
        align-items: center;
        flex-wrap: wrap;
        margin: 0 auto 0.7rem;
        @media screen and (max-width: 441px) {
          margin: 0 auto 0.5rem;
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
            display: flex;
            align-items: center;
            cursor: pointer;
            a {
              display: flex;
              align-items: center;
              justify-content: center;
              min-width: 1.6rem;
              padding: 0.13rem;
              margin: 0 0 0 0.2rem;
              background: linear-gradient(45deg, #4e88ff, #4b5fff);
              border-radius: 0.14rem;
              line-height: 1.5;
              text-align: center;
              color: #fff;
              font-size: 0.18rem;
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
          .search_left {
            display: flex;
            align-items: center;
            .icon {
              height: 40px;
              padding-left: 0.35rem;
              margin: 0 0 0 0.3rem;
              font-size: 15px;
              color: #000;
              line-height: 40px;
              @media screen and (max-width: 1440px) {
                font-size: 14px;
              }
              @media screen and (max-width: 441px) {
                height: 35px;
                line-height: 35px;
              }
              &:nth-child(1) {
                background: url(../../assets/images/space/icon_06.png) no-repeat
                  left center;
                background-size: 0.27rem auto;
                @media screen and (max-width: 1600px) {
                  background-size: 0.24rem auto;
                }
                @media screen and (max-width: 441px) {
                  background-size: 13px auto;
                }
              }
              &:nth-child(2) {
                background: url(../../assets/images/space/icon_07.png) no-repeat
                  left center;
                background-size: 0.27rem auto;
                @media screen and (max-width: 1600px) {
                  background-size: 0.24rem auto;
                }
                @media screen and (max-width: 441px) {
                  background-size: 13px auto;
                }
              }
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
              font-size: 0.18rem;
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
          width: 20px;
          height: 20px;
          margin: 0 0 0 5px;
          cursor: pointer;
          @media screen and (max-width: 1440px) {
            width: 17px;
            height: 17px;
          }
        }
      }
      .table_body /deep/ {
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
                font-size: 0.17rem;
                font-weight: 500;
                color: #555;
                text-transform: capitalize;
                line-height: 1.3;
                @media screen and (max-width: 768px) {
                  font-size: 13px;
                }
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
                font-size: 0.165rem;
                word-break: break-word;
                color: #000;
                text-align: center;
                line-height: 0.25rem;
                overflow: visible;
                @media screen and (max-width: 768px) {
                  font-size: 13px;
                }
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
                    width: 20px;
                    height: 20px;
                    margin: 0 5px;
                    font-size: 0.24rem;
                    @media screen and (max-width: 1600px) {
                      width: 18px;
                      height: 18px;
                    }
                    @media screen and (max-width: 768px) {
                      width: 15px;
                      height: 15px;
                    }
                    &:hover {
                      opacity: 0.7;
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
                  .icon_share {
                    background: url(../../assets/images/space/icon_09.png)
                      no-repeat center;
                    background-size: 100% 100%;
                  }
                  .copyText {
                    display: flex;
                    align-items: center;
                    .imgCopy {
                      width: 15px;
                      height: 15px;
                      margin: 0 0 0 5px;
                      cursor: pointer;
                      @media screen and (min-width: 1800px) {
                        width: 18px;
                        height: 18px;
                      }
                    }
                  }
                  span {
                    max-height: 0.5rem;
                    overflow: hidden;
                    text-overflow: ellipsis;
                    white-space: normal;
                    display: -webkit-box;
                    -webkit-line-clamp: 2;
                    -webkit-box-orient: vertical;
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
      }
    }
    .fe-none {
      display: flex;
      justify-content: center;
      align-items: center;
      padding: 0 5%;
      .p_label {
        padding: 0.15rem 0.3rem;
        font-size: 0.27rem;
        color: #4f87ff;
        line-height: 1.2;
        border: dashed;
        border-radius: 0.1rem;
        @media screen and (max-width: 1600px) {
          font-size: 0.25rem;
        }
      }
    }
  }
}
</style>
