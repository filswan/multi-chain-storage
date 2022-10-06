<template>
    <div class="statsCont" v-loading="loading">
        <div :class="{'opacity': loading, 'stats': true}">
            <div class="title">{{generateState('network_overview')}}</div>
            <div class="main">
                <div v-for="(item,key,index) in list" :key="index" class="info">
                    <img src="@/assets/images/icon_shangzhang.png" alt="">
                    <div class="info-up">
                        {{ item.title }}
                    </div>
                    <div class="info-num">{{ item.num | NumStatsFormat }}</div>
                </div>
            </div>
        </div>

        
        <el-row :class="{'opacity': loading_ecosystem, 'Eco': true}">
            <el-col :span="24">
                <div class="gradient"></div>
                <div class="subtitle">Multichain Storage Dataset</div>
                <el-row :gutter="20" class="mcs_dataset">
                    <el-col :xs="24" :sm="12" :md="8" :lg="6" v-for="(item, index) in MCS_Dataset" :key="index">
                        <div class="space">
                            <div class="tit">
                                <svg preserveAspectRatio="none">
                                    <defs>
                                        <linearGradient id="grad" x1="0%" y1="0%" x2="100%" y2="0%">
                                            <stop offset="0%" style="stop-color:#1c73ff; stop-opacity:1" />
                                            <stop offset="100%" style="stop-color:#78e4ff; stop-opacity:1" />
                                        </linearGradient>
                                    </defs>
                                    <text x="0" y="36" fill="url(#grad)">{{item.data}}</text>
                                </svg>
                            </div>
                            <div class="sub">{{item.desc}}</div>
                        </div>
                    </el-col>
                </el-row>
                <div class="subtitle">Collaborators</div>
                <el-row class="collaborators">
                    <el-col v-for="(item, index) in collaboratorsData" :key="index">
                        <a :href="item.link" target="_blank"><img :src="item.img" alt="logo" /></a>
                    </el-col>
                </el-row>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    let that
    // import {getStatsStorage} from "@/api/stats";
    import axios from 'axios'
    import { generateState } from '@/utils/i18n'
    export default {
        name: "stats",
        data() {
            return {
                list: [],
                loading: true,
                loading_ecosystem: true,
                logo_img_black: require("@/assets/images/icons/MCSLOGO.png"),
                collaboratorsData: [
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-02.png"),
                        link: 'https://filecoin.io/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-03.png"),
                        link: 'https://ipfs.tech/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-04.png"),
                        link: 'https://www.binance.com/en'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-05.png"),
                        link: 'https://polygon.technology/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-06.png"),
                        link: 'https://chain.link/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-07.png"),
                        link: 'https://protocol.ai/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-08.png"),
                        link: 'http://waterdrip.io/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-09.png"),
                        link: 'http://capital-chain.com/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-10.png"),
                        link: 'https://www.cidgravity.com/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-11.png"),
                        link: 'https://www.fbg.capital/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-12.png"),
                        link: 'https://ldcap.com/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-13.png"),
                        link: 'https://ti-capital.co/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-14.png"),
                        link: 'https://www.mcgill.ca/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-15.png"),
                        link: ''
                    },
                ],
                MCS_Dataset: [
                    {
                        data: '-',
                        desc: 'Total root CIDs uploaded to MCS. This value does not include sub objects references.'
                    },
                    {
                        data: '-',
                        desc: 'Active successful storage deals on the Filecoin Network'
                    },
                    {
                        data: '-',
                        desc: 'Total sealed storage contributed to Filecoin including a 5x replication'
                    },
                    {
                        data: '-',
                        desc: 'Total registered users'
                    },
                    {
                        data: '-',
                        desc: 'Total number of object references provided by every root CID in the network.'
                    },
                    {
                        data: '-',
                        desc: 'Total pinned IPFS storage for hot retrieval from any IPFS gateway. This data is not stored on Filecoin'
                    },
                    {
                        data: '-',
                        desc: 'Total storage providers receiving deals from our MCS node'
                    }
                ]
            }
        },
        created() {},
        mounted() {
            that = this
            document.getElementById('content-box').scrollTop = 0
            that.$store.dispatch('setRouterMenu', 4)
            that.$store.dispatch('setHeadertitle', that.$t('navbar.Stats'))
            that.getStats()
            that.getData()
        },
        methods:{
            generateState,
            deepClone(obj) {
                const _obj = JSON.stringify(obj)
                return JSON.parse(_obj)
            },
            async getRequest(apilink) {
                try {
                    const response = await axios.get(apilink, {
                        headers: {
                                'Authorization': "Bearer "+ that.$store.getters.mcsjwtToken
                        }
                    })
                    return response.data
                } catch (err) {
                    console.error(err, err.response)
                    return err.response.data
                }
            },
            async getStats() {
                let stats_api = `${process.env.BASE_API}stats/storage?wallet_address=${that.$store.getters.metaAddress}`

                axios.get(stats_api, {
                    headers: {
                            'Authorization': "Bearer "+ that.$store.getters.mcsjwtToken
                    },
                }).then(res => {
                    that.list = that.deepClone(res.data.data)
                    let language = localStorage.getItem('languageMcs') ? localStorage.getItem('languageMcs') : 'en'
                    for (let item in res.data.data){
                        that.list[item] = {
                            title: '',
                            num: 0
                        }
                        that.list[item].num = res.data.data[item]
                        switch (item) {
                            case 'average_cost_push_message' :
                                switch (language) {
                                    case 'cn' :
                                        that.list[item].title = '平均消息推送费用'
                                        break
                                    case 'en' :
                                        that.list[item].title = 'Average Push Message Cost'
                                }
                                break
                            case 'average_data_cost_sealing_1TB' :
                                switch (language) {
                                    case 'cn' :
                                        that.list[item].title = '每TiB质押费'
                                        break
                                    case 'en' :
                                        that.list[item].title = 'Pledge Collateral'
                                }
                                break
                            case 'average_gas_cost_sealing_1TB' :
                                switch (language) {
                                    case 'cn' :
                                        that.list[item].title = '平均封装手续费'
                                        break
                                    case 'en' :
                                        that.list[item].title = 'Cost of Sealing'
                                }
                                break
                            case 'average_min_piece_size' :
                                switch (language) {
                                    case 'cn' :
                                        that.list[item].title = '平均最小文件'
                                        break
                                    case 'en' :
                                        that.list[item].title = 'Average Minimum Piece Size'
                                }
                                that.list[item].num = res.data.data[item]?res.data.data[item]:'0 GiB'
                                break
                            case 'average_price_per_GB_per_year' :
                                switch (language) {
                                    case 'cn' :
                                        that.list[item].title = '平均数据存储价格'
                                        break
                                    case 'en' :
                                        that.list[item].title = 'Average Price'
                                }
                                that.list[item].num = res.data.data[item]?res.data.data[item]:'0 FIL/GiB/year'
                                break
                            case 'average_verified_price_per_GB_per_year' :
                                switch (language) {
                                    case 'cn' :
                                        that.list[item].title = '平均真实数据存储价格'
                                        break
                                    case 'en' :
                                        that.list[item].title = 'Average Verified Price'
                                }
                                that.list[item].num = res.data.data[item]?res.data.data[item]:'0 FIL/GiB/year'
                                break
                            case 'historical_average_price_regular' :
                                switch (language) {
                                    case 'cn' :
                                        that.list[item].title = '历史平均数据存储价格'
                                        break
                                    case 'en' :
                                        that.list[item].title = 'Historical Average Regular Price'
                                }
                                that.list[item].num = res.data.data[item]?res.data.data[item]:'0 FIL/GiB/100 deals'
                                break
                            case 'historical_average_price_verified' :
                                switch (language) {
                                    case 'cn' :
                                        that.list[item].title = '历史平均真实数据存储价格'
                                        break
                                    case 'en' :
                                        that.list[item].title = 'Historical Average Verified Price'
                                }
                                that.list[item].num = res.data.data[item]?res.data.data[item]:'0 FIL/GiB/100 deals'
                                break
                        }
                    }
                    that.loading = false
                })
            },
            async getData() {
                const ecosysRes = await that.getRequest(`${process.env.BASE_ECO_API}ecosys/data`);
                if(!ecosysRes || ecosysRes.status != 'success') {
                    that.MCS_Dataset = []
                    that.loading_ecosystem = false
                    return false
                }
                that.MCS_Dataset.forEach((element, i) => {
                    element.data = that.dataset(ecosysRes.data, i)
                });
                await that.timeout(500)
                that.loading_ecosystem = false
            },
            dataset(data, i){
                switch(i){
                    case 0:
                        return that.NumFormat(data.cid_count)
                    case 1:
                        return that.NumFormat(data.active_deal)
                    case 2:
                        return that.byteChange(data.sealed_storage)
                    case 3:
                        return that.NumFormat(data.wallet_count)
                    case 4:
                        return that.NumFormat(data.cid_object_reference)
                    case 5:
                        return that.byteChange(data.pinned_ipfs_size)
                    case 6:
                        return that.NumFormat(data.miner_count)
                    default:
                        return '-';
                }
            },
            timeout (delay) {
                return new Promise((res) => setTimeout(res, delay))
            },
            byteChange(bytes){
                if (String(bytes) == '0') return '0 B';
                if (!bytes) return "-";
                var k = 1024, // or 1000
                    sizes = ['B', 'KiB', 'MiB', 'GiB', 'TiB', 'PiB', 'EiB', 'ZiB', 'YiB'],
                    i = Math.floor(Math.log(bytes) / Math.log(k));

                if (Math.round((bytes / Math.pow(k, i))).toString().length > 3) {
                    i += 1
                }
                return (bytes / Math.pow(k, i)).toFixed(2) + ' ' + sizes[i];
            },
            NumFormat(value) {
                if (!value) return "-";
                var intPart = Number(value).toFixed(0); 
                var intPartFormat = intPart
                    .toString()
                    .replace(/(\d)(?=(?:\d{3})+$)/g, "$1,"); //将整数部分逢三一断
                return intPartFormat;
            }
        },
        filters: {
            NumStatsFormat(value) {
                if (!value) return "0";
                return value;
            },
        }
    }
</script>

<style lang="scss" scoped>
.statsCont{
    position: relative;
    padding: 0.3rem 0.2rem;
    background-color: #000;
    background-image: url(../../assets/images/dashboard/bg_top.png), url(../../assets/images/dashboard/bg_bottom.png);
    background-position: left top, right bottom;
    background-repeat: no-repeat, no-repeat;
    background-size: 27%, 22%;
    .el-loading-mask /deep/{
        background-color: rgba(0, 0, 0, 0.95);
    }
    .stats {
        .title {
            margin: 0 0.1rem;
            font-size: 0.22rem;
            font-weight: 700;
            color: #fff;
            line-height: 1.5;
        }

        .main {
            font-size: 18px;
            width: 100%;
            padding: 0.1rem 0 0;
            display: flex;
            flex-wrap: wrap;
            .info-num{
                width: 100%;
                text-align: center;
                font-size: 0.22rem;
                font-weight: 500;
                color: #000;
            }
            img{
                display: block;
                width: 0.44rem;
                margin: 0 auto 0.3rem;
            }
        }

        .info {
            width: calc(33.33% - 0.2rem);
            padding: 0.5rem 0;
            min-height: 100px;
            display: flex;
            justify-content: center;
            /*align-items: center;*/
            flex-direction: column;
            background: #ffffff;
            border-radius: 0.2rem;
            margin: 0.1rem;
            line-height: 1;
        }

        .info-up {
            font-size: 0.2rem;
            color: #2C7FF8;
            width: 100%;
            text-align: center;
            font-weight: 500;
            margin-bottom: 0.2rem;
        }

        .el-icon-dog{

            background-size: contain;
        }
        .el-icon-dog:before{
            content: "dog";
            font-size: 12px;
            visibility: hidden;
        }
    }
    .opacity{
        opacity: 0;
    }
    .Eco /deep/{
        padding: 0.1rem;
        font-size: 0.22rem;
        font-weight: 700;
        color: #fff;
        line-height: 1.5;
        .logoImg {
            display: block;
            height: 80px;
            @media screen and (max-width: 1600px) {
                height: 60px;
            }
        }
        .title {
            padding: 30px 0 20px;
            font-size: 22px;
            @media screen and (max-width: 1600px) {
                font-size: 18px;
            }
            @media screen and (max-width: 1440px) {
                font-size: 16px;
            }
        }
        .subtitle {
            padding: 20px 0 10px;
            font-size: 22px;
            font-weight: 600;
            @media screen and (max-width: 1600px) {
                font-size: 16px;
            }
        }
        .gradient{
            width: 100%;
            height: 3px;
            margin: 50px 0 0;
            background-image: linear-gradient(90deg, #0036ab, #7fedff);
            @media screen and (max-width: 1440px) {
                margin: 30px 0 0;
            }
        }
        .collaborators{
            display: flex;
            justify-content: flex-start;
            align-items: center;
            flex-wrap: wrap;
            width: 97%;
            margin: auto;
            .el-col{
                flex: 0 0 14%;
                margin: 0 auto;
                @media screen and (max-width: 1200px) {
                    flex: 0 0 20%;
                }
                @media screen and (max-width: 1024px) {
                    flex: 0 0 25%;
                }
                @media screen and (max-width: 768px) {
                    flex: 0 0 50%;
                }
                @media screen and (max-width: 414px) {
                    flex: 0 0 100%;
                }
                a{
                    img{
                        display: block;
                        height: 60px;
                        margin: 20px auto;
                        cursor: pointer;
                        @media screen and (max-width: 1600px) {
                            height: 45px;
                        }
                        @media screen and (max-width: 1440px) {
                            margin: 15px auto;
                        }
                    }
                }
            }
        }
        .mcs_dataset{
            display: flex;
            justify-content: flex-start;
            align-items: stretch;
            flex-wrap: wrap;
            padding: 0 2.5%;
            .el-col{
                margin-top: 20px;
                margin-bottom: 20px;
                .space{
                    height: 100%;
                    padding: 3px 25px 5px;
                    background-color: #fff;
                    border-radius: 7px;
                    @media screen and (max-width: 1440px) {
                        padding: 3px 20px 5px;
                    }
                    .tit{
                        position: relative;
                        padding: 0;
                        margin: 0;
                        font-size: 25.7px;
                        font-weight: 600;
                        line-height: 1.5;
                        @media screen and (max-width: 1600px) {
                            font-size: 22px;
                        }
                        @media screen and (max-width: 1440px) {
                            font-size: 20px;
                        }
                        svg{
                            display: flex;
                            width: 100%;
                            height: 50px;
                        }
                    }
                    .sub{
                        font-size: 15px;
                        font-weight: 500;
                        color: #7c7c7c;
                        line-height: 1.6;
                        word-break: break-word;
                        @media screen and (max-width: 1600px) {
                            font-size: 14px;
                        }
                        @media screen and (max-width: 1440px) {
                            font-size: 12px;
                            line-height: 1.3;
                        }
                    }
                }
            }
        }
    }
}
@media screen and (max-width: 800px) {
    .statsCont{
        .stats {
            .main {
                font-size: 18px;
                width: calc(100% - 40px);
                padding: 20px;
                display: flex;
                flex-wrap: wrap;
                .info-num{
                    width: 100%;
                    text-align: center;
                    font-weight: bold;
                }
            }
            .info {
                width: calc(100%);
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

            .el-icon-dog{

                background-size: contain;
            }
            .el-icon-dog:before{
                content: "dog";
                font-size: 12px;
                visibility: hidden;
            }
        }
    }
}
</style>
