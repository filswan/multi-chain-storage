<template>
  <div class="statsHome">
    <div class="statsCont" v-loading="loading_ecosystem">
      <div :class="{'opacity': loading_ecosystem, 'stats': true}">
        <div class="title">
          Multichain Storage Overview
        </div>
        <div class="main">
          <div v-for="(item, index) in MCS_Dataset" :key="index" class="info">
            <img v-if="index === 7" src="@/assets/images/menuIcon/metaSpace-2.png">
            <svg v-else xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
              <path v-if="index == 1" stroke-linecap="round" stroke-linejoin="round" d="M20.25 6.375c0 2.278-3.694 4.125-8.25 4.125S3.75 8.653 3.75 6.375m16.5 0c0-2.278-3.694-4.125-8.25-4.125S3.75 4.097 3.75 6.375m16.5 0v11.25c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125V6.375m16.5 0v3.75m-16.5-3.75v3.75m16.5 0v3.75C20.25 16.153 16.556 18 12 18s-8.25-1.847-8.25-4.125v-3.75m16.5 0c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125"
              />
              <path v-else-if="index == 2" stroke-linecap="round" stroke-linejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m6.75 12l-3-3m0 0l-3 3m3-3v6m-1.5-15H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z"
              />
              <path v-else-if="index == 3" stroke-linecap="round" stroke-linejoin="round" d="M19 7.5v3m0 0v3m0-3h3m-3 0h-3m-2.25-4.125a3.375 3.375 0 11-6.75 0 3.375 3.375 0 016.75 0zM4 19.235v-.11a6.375 6.375 0 0112.75 0v.109A12.318 12.318 0 0110.374 21c-2.331 0-4.512-.645-6.374-1.766z"
              />
              <path v-else-if="index == 4" stroke-linecap="round" stroke-linejoin="round" d="M20.25 14.15v4.25c0 1.094-.787 2.036-1.872 2.18-2.087.277-4.216.42-6.378.42s-4.291-.143-6.378-.42c-1.085-.144-1.872-1.086-1.872-2.18v-4.25m16.5 0a2.18 2.18 0 00.75-1.661V8.706c0-1.081-.768-2.015-1.837-2.175a48.114 48.114 0 00-3.413-.387m4.5 8.006c-.194.165-.42.295-.673.38A23.978 23.978 0 0112 15.75c-2.648 0-5.195-.429-7.577-1.22a2.016 2.016 0 01-.673-.38m0 0A2.18 2.18 0 013 12.489V8.706c0-1.081.768-2.015 1.837-2.175a48.111 48.111 0 013.413-.387m7.5 0V5.25A2.25 2.25 0 0013.5 3h-3a2.25 2.25 0 00-2.25 2.25v.894m7.5 0a48.667 48.667 0 00-7.5 0M12 12.75h.008v.008H12v-.008z"
              />
              <path v-else-if="index == 5" stroke-linecap="round" stroke-linejoin="round" d="M7.875 14.25l1.214 1.942a2.25 2.25 0 001.908 1.058h2.006c.776 0 1.497-.4 1.908-1.058l1.214-1.942M2.41 9h4.636a2.25 2.25 0 011.872 1.002l.164.246a2.25 2.25 0 001.872 1.002h2.092a2.25 2.25 0 001.872-1.002l.164-.246A2.25 2.25 0 0116.954 9h4.636M2.41 9a2.25 2.25 0 00-.16.832V12a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 12V9.832c0-.287-.055-.57-.16-.832M2.41 9a2.25 2.25 0 01.382-.632l3.285-3.832a2.25 2.25 0 011.708-.786h8.43c.657 0 1.281.287 1.709.786l3.284 3.832c.163.19.291.404.382.632M4.5 20.25h15A2.25 2.25 0 0021.75 18v-2.625c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125V18a2.25 2.25 0 002.25 2.25z"
              />
              <path v-else-if="index == 6" stroke-linecap="round" stroke-linejoin="round" d="M9 3.75H6.912a2.25 2.25 0 00-2.15 1.588L2.35 13.177a2.25 2.25 0 00-.1.661V18a2.25 2.25 0 002.25 2.25h15A2.25 2.25 0 0021.75 18v-4.162c0-.224-.034-.447-.1-.661L19.24 5.338a2.25 2.25 0 00-2.15-1.588H15M2.25 13.5h3.86a2.25 2.25 0 012.012 1.244l.256.512a2.25 2.25 0 002.013 1.244h3.218a2.25 2.25 0 002.013-1.244l.256-.512a2.25 2.25 0 012.013-1.244h3.859M12 3v8.25m0 0l-3-3m3 3l3-3"
              />
              <path v-else stroke-linecap="round" stroke-linejoin="round" d="M7.864 4.243A7.5 7.5 0 0119.5 10.5c0 2.92-.556 5.709-1.568 8.268M5.742 6.364A7.465 7.465 0 004.5 10.5a7.464 7.464 0 01-1.15 3.993m1.989 3.559A11.209 11.209 0 008.25 10.5a3.75 3.75 0 117.5 0c0 .527-.021 1.049-.064 1.565M12 10.5a14.94 14.94 0 01-3.6 9.75m6.633-4.596a18.666 18.666 0 01-2.485 5.33"
              />
            </svg>

            <div class="info-up">
              {{item.desc}}
              <el-popover placement="top" popper-class="elPopTitle" width="200" trigger="hover" :content="item.popover">
                <svg slot="reference" t="1666579260240" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2746" width="128" height="128">
                  <path d="M512 958.016C266.08 958.016 65.984 757.952 65.984 512 65.984 266.08 266.08 65.984 512 65.984c245.952 0 446.016 200.064 446.016 446.016C958.016 757.952 757.952 958.016 512 958.016zM512 129.984C301.344 129.984 129.984 301.344 129.984 512c0 210.624 171.36 382.016 382.016 382.016 210.624 0 382.016-171.36 382.016-382.016C894.016 301.344 722.624 129.984 512 129.984z"
                    p-id="2747" fill="#6f747c"></path>
                  <path d="M512 304m-48 0a1.5 1.5 0 1 0 96 0 1.5 1.5 0 1 0-96 0Z" p-id="2748" fill="#6f747c"></path>
                  <path d="M512 768c-17.664 0-32-14.304-32-32l0-288c0-17.664 14.336-32 32-32s32 14.336 32 32l0 288C544 753.696 529.664 768 512 768z" p-id="2749" fill="#6f747c"></path>
                </svg>
              </el-popover>
            </div>
            <div class="info-num">
              <span>{{item.data}}</span>
            </div>
          </div>
        </div>
        <div class="roseChart_all">
          <!--          <div class="roseChart" id="roseChart"></div>-->
          <div class="roseChart" id="roseChart_1"></div>
        </div>

        <div class="title">
          Collaborators
        </div>
        <el-row class="collaborators">
          <el-col v-for="(item, i) in collaboratorsData" :key="i">
            <a class="moon" :href="item.link" target="_blank"><img :src="item.img" alt="logo" /></a>
            <a class="sunny" :href="item.link" target="_blank"><img :src="item.img_sunny" alt="logo" /></a>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
// import axios from 'axios'
// import * as echarts from 'echarts'
import { generateState } from '@/utils/i18n'
// import moment from 'moment'
let that
export default {
  name: 'Stats',
  components: {},
  data () {
    return {
      list: [],
      loading: true,
      loading_ecosystem: true,
      collaboratorsData: [
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-01.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-01-sunny.png'),
          link: 'https://protocol.ai/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-02.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-02-sunny.png'),
          link: 'https://ipfs.tech/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-03.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-03-sunny.png'),
          link: 'https://filecoin.io/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-05.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-05-sunny.png'),
          link: 'https://chainlinklabs.com/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-06.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-06-sunny.png'),
          link: 'https://labs.binance.com/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-08.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-08-sunny.png'),
          link: 'https://arbitrum.io/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-07.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-07-sunny.png'),
          link: 'https://opensea.io/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-09.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-09-sunny.png'),
          link: 'https://akash.network/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-04.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-04-sunny.png'),
          link: 'https://polygon.technology/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-10.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-10-sunny.png'),
          link: 'https://aptos.dev/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-11.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-11-sunny.png'),
          link: 'https://sui.io/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-14.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-14-sunny.png'),
          link: 'https://www.web3cloud.tech/'
        },
        // {
        //   img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-15.png'),
        //   img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-15-sunny.png'),
        //   link: 'https://github.com/srblabotw69/Afianthack'
        // },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-16.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-16-sunny.png'),
          link: 'https://github.com/IKalonji/CooperDB'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-17.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-17-sunny.png'),
          link: 'https://sao.network/#/'
        },
        {
          img: require('@/assets/images/dashboard/moon/Oortech-moon.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/Oortech-sunny.png'),
          link: 'https://www.oortech.com/'
        },
        {
          img: require('@/assets/images/dashboard/moon/Telnyx-moon.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/Telnyx-sunny.png'),
          link: 'https://telnyx.com/'
        },
        {
          img: require('@/assets/images/dashboard/moon/Apus-moon.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/Apus-sunny.png'),
          link: 'https://www.apus.network/'
        },
        {
          img: require('@/assets/images/dashboard/moon/PPIO-moon.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/PPIO-sunny.png'),
          link: 'https://www.ppio.cn/'
        },
        {
          img: require('@/assets/images/dashboard/moon/Filmine-moon.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/Filmine-sunny.png'),
          link: 'https://filmine.io/'
        },
        {
          img: require('@/assets/images/dashboard/moon/Lagrange-moon.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/Lagrange-sunny.png'),
          link: 'https://lagrangedao.org/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-13.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-13-sunny.png'),
          link: 'https://en.fogmeta.com/'
        },
        {
          img: require('@/assets/images/dashboard/moon/MULTI-CHAIN-12.png'),
          img_sunny: require('@/assets/images/dashboard/sunny/MULTI-CHAIN-12-sunny.png'),
          link: 'https://www.nebulablock.com/'
        }
      ],
      MCS_Dataset: [
        {
          data: '-',
          desc: 'Total CIDs',
          popover: 'The total amount of pinned IPFS CIDs'
        },
        {
          data: '- GB',
          desc: 'IPFS Storage',
          popover: 'Total data pinned on IPFS'
        },
        {
          data: '-',
          desc: 'Files Uploaded',
          popover: 'The total amount of files uploaded'
        },
        {
          data: '-',
          desc: 'Registered Users',
          popover: 'The number of wallet addresses registered'
        },
        {
          data: '-',
          desc: 'Total Files Archived',
          popover: 'The number of files archived on the Filecoin network'
        },
        {
          data: '-',
          desc: 'Storage Archived',
          popover: 'Total storage archived on the Filecoin network'
        },
        {
          data: '-',
          desc: 'Storage Providers',
          popover: 'The number of storage providers offering decentralized storage capacity'
        },
        {
          data: '-',
          desc: 'Buckets',
          popover: 'Total number of Buckets created'
        }
      ],
      echartData: {
        time: [],
        upload: [],
        archived: [],
        bucket_count: [],
        pinned_size: []
      }
    }
  },
  created () { },
  mounted () {
    that = this
    document.getElementById('app').scrollTop = 0
    that.$store.dispatch('setRouterMenu', 4)
    that.$store.dispatch('setHeadertitle', that.$t('navbar.Stats'))
    that.getData()
  },
  methods: {
    generateState,
    deepClone (obj) {
      const _obj = JSON.stringify(obj)
      return JSON.parse(_obj)
    },
    async getRequest (apilink) {
      try {
        const response = await axios.get(apilink, {
          headers: {
            'Authorization': 'Bearer ' + that.$store.getters.mcsjwtToken
          }
        })
        return response.data
      } catch (err) {
        console.error(err, err.response)
        return false
      }
    },
    // E chart
    async getPie () {
      // Init E chart
      // MyChart for IPFSStorage&StorageArchived
      // const myChart = echarts.init(document.getElementById('roseChart'))
      // MyChart1 for FilesUploaded&Buckets
      const myChart1 = echarts.init(document.getElementById('roseChart_1'))
      // Setting Chart for 'IPFS Storage (GiB)', 'Storage Archived (GiB)'
      // const option = {
      //   toolbox: {
      //     show: true,
      //     feature: {
      //       mark: {
      //         show: true
      //       },
      //       // dataZoom: {
      //       //   yAxisIndex: false,
      //       // },
      //       // dataView: {
      //       //     show: true,
      //       //     readOnly: false,
      //       //     title:'数据视图'
      //       // },
      //       magicType: {
      //         show: true,
      //         type: ['line', 'bar'],
      //         title: {
      //           line: 'Switch to line graph',
      //           bar: 'Switch to bar graph'
      //         }
      //       },
      //       restore: {
      //         show: true,
      //         title: 'Restore'
      //       },
      //       saveAsImage: {
      //         show: true,
      //         title: 'Save'
      //       }
      //     }
      //   },
      //   tooltip: {
      //     trigger: 'axis',
      //     axisPointer: {
      //       type: 'cross',
      //       label: {
      //         backgroundColor: '#6a7985'
      //       }
      //     }
      //   },
      //   legend: {
      //     bottom: 0,
      //     left: 'center',
      //     textStyle: {
      //       color: '#000',
      //       fontSize: 13
      //     },
      //     data: ['IPFS Storage (GiB)', 'Storage Archived (GiB)']
      //   },
      //   grid: {
      //     top: '70',
      //     left: '3%',
      //     right: '3%',
      //     bottom: '50',
      //     containLabel: true
      //   },
      //   // color: ['#4472c4', '#ed7d31', 'green'],
      //   xAxis: {
      //     type: 'category',
      //     boundaryGap: false,
      //     data: that.echartData.time.map(function (str) {
      //       return str.replace(' ', '\n')
      //     })
      //   },
      //   yAxis: [
      //     // IPFS Storage (GiB)
      //     {
      //       // boundaryGap: [0, '30%'],
      //       name: 'GiB',
      //       type: 'value',
      //       // min: that.echartData.min_file,
      //       // max: that.echartData.max_file,
      //       scale: true,
      //       splitNumber: 5,
      //       alignTicks: true
      //     },
      //     // Storage Archived (GiB)
      //     {
      //       // boundaryGap: [0, '50%'],
      //       name: 'GiB',
      //       type: 'value',
      //       position: 'right',
      //       // min: that.echartData.min_val,
      //       // max: that.echartData.max_val,
      //       scale: true,
      //       splitNumber: 5
      //     }
      //   ],
      //   // loading data

      //   series: [
      //     {
      //       name: 'IPFS Storage (GiB)',
      //       type: 'line',
      //       data: that.echartData.pinned_size,
      //       yAxisIndex: 0
      //     },
      //     {
      //       name: 'Storage Archived (GiB)',
      //       type: 'line',
      //       data: that.echartData.archived,
      //       yAxisIndex: 1
      //       // showSymbol: false
      //     }
      //   ]
      // }
      // Setting chart for 'Files uploaded', 'Buckets'
      const option1 = {
        toolbox: {
          show: true,
          feature: {
            mark: {
              show: true
            },
            // dataZoom: {
            //   yAxisIndex: false,
            // },
            // dataView: {
            //     show: true,
            //     readOnly: false,
            //     title:'数据视图'
            // },
            magicType: {
              show: true,
              type: ['line', 'bar'],
              title: {
                line: 'Switch to line graph',
                bar: 'Switch to bar graph'
              }
            },
            restore: {
              show: true,
              title: 'Restore'
            },
            saveAsImage: {
              show: true,
              title: 'Save'
            }
          }
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross',
            label: {
              backgroundColor: '#6a7985'
            }
          }
        },
        legend: {
          bottom: 0,
          left: 'center',
          textStyle: {
            color: '#000',
            fontSize: 13
          },
          data: ['Files Uploaded']
        },
        grid: {
          top: '70',
          left: '3%',
          right: '3%',
          bottom: '50',
          containLabel: true
        },
        // color: ['#4472c4', '#ed7d31', 'green'],
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: that.echartData.time.map(function (str) {
            return str.replace(' ', '\n')
          })
        },
        yAxis: [
          // 'Files uploaded'
          {
            // boundaryGap: [0, '30%'],
            name: 'File Count',
            type: 'value',
            // min: that.echartData.min_file,
            // max: that.echartData.max_file,
            scale: true,
            splitNumber: 5,
            alignTicks: true
          }
        ],
        // loading data
        series: [
          {
            name: 'Files Uploaded',
            type: 'line',
            data: that.echartData.upload,
            yAxisIndex: 0

          }
        ]
      }
      // put data to chart
      // myChart.setOption(option)
      myChart1.setOption(option1)
      window.onresize = function () {
        // myChart.resize()
        myChart1.resize()
      }
    },
    async getData () {
      const ecosysRes = await that.getRequest(`${process.env.BASE_ECO_API}ecosys/data`)
      if (!ecosysRes || ecosysRes.status !== 'success') {
        that.MCS_Dataset = []
        that.loading_ecosystem = false
        return false
      }
      that.MCS_Dataset.forEach((element, i) => {
        element.data = that.dataset(ecosysRes.data, i)
      })
      if (ecosysRes.data.graph) {
        // that.echartData.max_file = ecosysRes.data.graph.max_file || 0
        // that.echartData.max_val = ecosysRes.data.graph.max_val || 0
        // that.echartData.min_file = ecosysRes.data.graph.min_file || 0
        // that.echartData.min_val = ecosysRes.data.graph.min_val || 0
        that.echartData.time = await that.dataFilter(ecosysRes.data.graph.time, 1)
        that.echartData.archived = await that.dataFilter(ecosysRes.data.graph.sealed_storage)
        that.echartData.upload = ecosysRes.data.graph.user_uploads || []
        that.echartData.pinned_size = await that.dataFilter(ecosysRes.data.graph.pinned_size)
        // that.echartData.bucket_count = ecosysRes.data.graph.user_bucket_count || []
      }
      await that.timeout(500)
      that.loading_ecosystem = false
      that.getPie()
    },
    dataFilter (data, type) {
      let timeData = data || []
      let num = 0
      let list = []
      timeData.forEach((item, index) => {
        if (type) {
          item = item
            ? moment(new Date(parseInt(item * 1000))).format('YYYY-MM-DD')
            : '-'
          list.push(item)
        } else {
          if (item <= 0) list.push(0)
          else {
            num = (item / (1024 * 1024 * 1024)) // or 1000
            num = num.toFixed(2)
            list.push(num)
          }
        }
      })
      return list
    },
    // Format data
    dataset (data, i) {
      switch (i) {
        case 0:
          return that.NumFormat(data.source_file_cids)
        case 1:
          return that.changeToGiB(data.pinned_ipfs_size)
        case 2:
          return that.NumFormat(data.user_uploads)
        case 3:
          return that.NumFormat(data.wallet_count)
        case 4:
          return that.NumFormat(data.active_deals)
        case 5:
          return that.byteChange(data.sealed_storage)
        case 6:
          return that.NumFormat(data.miner_count)
        case 7:
          return that.NumFormat(data.bucket_folder_count)
        default:
          return '-'
      }
    },
    timeout (delay) {
      return new Promise((resolve) => setTimeout(resolve, delay))
    },
    byteChange (bytes) {
      if (String(bytes) === '0') return '0 B'
      if (!bytes) return '-'
      var k = 1024 // or 1000
      var sizes = ['B', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB']
      var i = Math.floor(Math.log(bytes) / Math.log(k))
      if (Math.round((bytes / Math.pow(k, i))).toString().length > 3) {
        i += 1
      }
      return (bytes / Math.pow(k, i)).toFixed(2) + ' ' + sizes[i]
    },
    changeToGiB (bytes) {
      var size = ''
      // 只转换成GB
      if (bytes <= 0) {
        return '-'
      } else {
        size = bytes / (1024 * 1024 * 1024) // or 1000
        size = size.toFixed(2)
      }
      return size + ' GiB'
      // return Number(size).toFixed(3);
    },
    NumFormat (value) {
      if (String(value) === '0') return '0'
      else if (!value) return '-'
      var intPartArr = String(value).split('.')
      var intPartFormat = intPartArr[0]
        .toString()
        .replace(/(\d)(?=(?:\d{3})+$)/g, '$1,')
      return intPartArr[1] ? `${intPartFormat}.${intPartArr[1]}` : intPartFormat
    }
  },
  filters: {
    NumStatsFormat (value) {
      if (!value) return '0'
      return value
    }
  }
}
</script>

<style lang="scss" scoped>
.statsCont {
  position: relative;
  width: 80%;
  max-width: 1440px;
  min-width: 300px;
  margin: 0.1rem auto;
  padding: 0.5rem 0;
  background-color: #fff;
  @media screen and (max-width: 992px) {
    width: 90%;
    padding: 0.25rem 0;
  }
  @media screen and (max-width: 441px) {
    width: 94%;
  }
  border-radius: 0.1rem;
  .moon {
    display: none;
  }
  .sunny {
    display: block;
  }
  .el-loading-mask /deep/ {
    background-color: rgba(0, 0, 0, 0.95);
  }
  .stats {
    .title {
      margin: 10px 0.2rem;
      font-size: 0.22rem;
      font-weight: 700;
      color: #02003b;
      line-height: 1.5;
      @media screen and (max-width: 992px) {
        font-size: 18px;
      }
    }
    .main {
      font-size: 18px;
      width: 100%;
      padding: 0.1rem 0 0.2rem;
      display: flex;
      flex-wrap: wrap;
      .info-num {
        width: 100%;
        text-align: center;
        font-size: 0.22rem;
        font-weight: 500;
        color: #000;
      }
      img,
      svg {
        display: block;
        width: 0.35rem;
        margin: 0 auto 0.2rem;
      }
    }
    .roseChart_all {
      //display: flex;
      //flex-wrap: wrap;
      //justify-content: center;
      padding: 0;
      margin: 0.15rem auto 1rem;
      @media screen and (max-width: 600px) {
        display: block;
      }
      .roseChart {
        width: 100%;
        height: 460px;
        margin: 0.2rem 0 0.35rem;
        @media screen and (max-width: 1600px) {
          height: 440px;
        }
        @media screen and (max-width: 1366px) {
          height: 400px;
        }
        @media screen and (max-width: 999px) {
          height: 360px;
        }
        @media screen and (max-width: 768px) {
          width: 100%;
          height: 420px;
        }
        @media screen and (max-width: 441px) {
          height: 390px;
        }
      }
    }
    .info {
      width: calc(25% - 0.4rem);
      padding: 0.2rem 0;
      min-height: 100px;
      display: flex;
      justify-content: center;
      /*align-items: center;*/
      flex-direction: column;
      background: #e6efff;
      border-radius: 0.2rem;
      margin: 0.2rem;
      line-height: 1;
      @media screen and (max-width: 1260px) {
        width: calc(25% - 0.4rem);
      }
      @media screen and (max-width: 1024px) {
        width: calc(33.33% - 0.4rem);
      }
      @media screen and (max-width: 768px) {
        width: calc(50% - 0.4rem);
      }
      @media screen and (max-width: 441px) {
        width: calc(100% - 0.4rem);
      }
    }
    .info-up {
      font-size: 0.2rem;
      color: #02003b;
      width: 96%;
      padding: 0 2%;
      text-align: center;
      font-weight: 500;
      margin-bottom: 0.2rem;
      display: flex;
      align-items: center;
      justify-content: center;
      img,
      svg {
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
    .el-icon-dog {
      background-size: contain;
    }
    .el-icon-dog:before {
      content: "dog";
      font-size: 12px;
      visibility: hidden;
    }
    .collaborators {
      display: flex;
      justify-content: flex-start;
      align-items: center;
      flex-wrap: wrap;
      width: 97%;
      margin: auto;
      .el-col {
        width: auto;
        flex: 0 0 auto;
        margin: 0 2%;
        a {
          img {
            display: block;
            height: 55px;
            margin: 20px auto;
            cursor: pointer;
            @media screen and (max-width: 1600px) {
              height: 50px;
            }
            @media screen and (max-width: 992px) {
              height: 45px;
            }
            @media screen and (max-width: 441px) {
              height: 40px;
            }
            @media screen and (max-width: 1440px) {
              margin: 15px auto;
            }
          }
          .height {
            height: 80px;
            margin: 10px auto;
            @media screen and (max-width: 1600px) {
              height: 65px;
            }
            @media screen and (max-width: 1440px) {
              margin: 5px auto;
            }
          }
        }
      }
    }
  }
  .opacity {
    opacity: 0;
  }
}
.reverse_phase {
  .statsCont {
    .moon {
      display: block;
    }
    .sunny {
      display: none;
    }
  }
}
@media screen and (max-width: 800px) {
  .statsCont {
    .stats {
      .main {
        font-size: 18px;
        width: calc(100% - 40px);
        padding: 20px;
        display: flex;
        flex-wrap: wrap;
        .info-num {
          width: 100%;
          text-align: center;
          font-weight: bold;
        }
      }
      .info {
        display: flex;
        justify-content: center;
        /*align-items: center;*/
        flex-direction: column;
        margin: 5px;
      }
      .info-up {
        font-size: 14px;
        color: #0b318f;
        width: 100%;
        text-align: center;
        font-weight: bold;
        margin-bottom: 10px;
        height: 30px;
      }
      .el-icon-dog {
        background-size: contain;
      }
      .el-icon-dog:before {
        content: "dog";
        font-size: 12px;
        visibility: hidden;
      }
    }
  }
}
</style>
