<template>
    <div class="statsCont" v-loading="loading_ecosystem">
        <div :class="{'opacity': loading_ecosystem, 'stats': true}">
            <div class="title">
                Collaborators
            </div>
            <el-row class="collaborators">
                <el-col v-for="(item, index) in collaboratorsData" :key="index">
                    <a :href="item.link" target="_blank"><img :src="item.img" :class="{'height': index == 4 || index == 6 || index == 7}" alt="logo" /></a>
                </el-col>
            </el-row>
            <!-- <CarouselContainer :slide-list="collaboratorsData" currentIndex="1"></CarouselContainer> -->

            <div class="title">
                Multichain Storage Overview
            </div>
            <div class="main">
                <div v-for="(item, index) in MCS_Dataset" :key="index+8" class="info">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
                        <path v-if="index == 1" stroke-linecap="round" stroke-linejoin="round" d="M19.5 5.25l-7.5 7.5-7.5-7.5m15 6l-7.5 7.5-7.5-7.5" />
                        <path v-else-if="index == 2" stroke-linecap="round" stroke-linejoin="round" d="M20.25 6.375c0 2.278-3.694 4.125-8.25 4.125S3.75 8.653 3.75 6.375m16.5 0c0-2.278-3.694-4.125-8.25-4.125S3.75 4.097 3.75 6.375m16.5 0v11.25c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125V6.375m16.5 0v3.75m-16.5-3.75v3.75m16.5 0v3.75C20.25 16.153 16.556 18 12 18s-8.25-1.847-8.25-4.125v-3.75m16.5 0c0 2.278-3.694 4.125-8.25 4.125s-8.25-1.847-8.25-4.125" />
                        <path v-else-if="index == 3" stroke-linecap="round" stroke-linejoin="round" d="M19 7.5v3m0 0v3m0-3h3m-3 0h-3m-2.25-4.125a3.375 3.375 0 11-6.75 0 3.375 3.375 0 016.75 0zM4 19.235v-.11a6.375 6.375 0 0112.75 0v.109A12.318 12.318 0 0110.374 21c-2.331 0-4.512-.645-6.374-1.766z" />
                        <path v-else-if="index == 4" stroke-linecap="round" stroke-linejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 00-3.375-3.375h-1.5A1.125 1.125 0 0113.5 7.125v-1.5a3.375 3.375 0 00-3.375-3.375H8.25m6.75 12l-3-3m0 0l-3 3m3-3v6m-1.5-15H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 00-9-9z" />
                        <path v-else-if="index == 5" stroke-linecap="round" stroke-linejoin="round" d="M20.25 14.15v4.25c0 1.094-.787 2.036-1.872 2.18-2.087.277-4.216.42-6.378.42s-4.291-.143-6.378-.42c-1.085-.144-1.872-1.086-1.872-2.18v-4.25m16.5 0a2.18 2.18 0 00.75-1.661V8.706c0-1.081-.768-2.015-1.837-2.175a48.114 48.114 0 00-3.413-.387m4.5 8.006c-.194.165-.42.295-.673.38A23.978 23.978 0 0112 15.75c-2.648 0-5.195-.429-7.577-1.22a2.016 2.016 0 01-.673-.38m0 0A2.18 2.18 0 013 12.489V8.706c0-1.081.768-2.015 1.837-2.175a48.111 48.111 0 013.413-.387m7.5 0V5.25A2.25 2.25 0 0013.5 3h-3a2.25 2.25 0 00-2.25 2.25v.894m7.5 0a48.667 48.667 0 00-7.5 0M12 12.75h.008v.008H12v-.008z" />
                        <path v-else-if="index == 6" stroke-linecap="round" stroke-linejoin="round" d="M3.375 19.5h17.25m-17.25 0a1.125 1.125 0 01-1.125-1.125M3.375 19.5h7.5c.621 0 1.125-.504 1.125-1.125m-9.75 0V5.625m0 12.75v-1.5c0-.621.504-1.125 1.125-1.125m18.375 2.625V5.625m0 12.75c0 .621-.504 1.125-1.125 1.125m1.125-1.125v-1.5c0-.621-.504-1.125-1.125-1.125m0 3.75h-7.5A1.125 1.125 0 0112 18.375m9.75-12.75c0-.621-.504-1.125-1.125-1.125H3.375c-.621 0-1.125.504-1.125 1.125m19.5 0v1.5c0 .621-.504 1.125-1.125 1.125M2.25 5.625v1.5c0 .621.504 1.125 1.125 1.125m0 0h17.25m-17.25 0h7.5c.621 0 1.125.504 1.125 1.125M3.375 8.25c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125m17.25-3.75h-7.5c-.621 0-1.125.504-1.125 1.125m8.625-1.125c.621 0 1.125.504 1.125 1.125v1.5c0 .621-.504 1.125-1.125 1.125m-17.25 0h7.5m-7.5 0c-.621 0-1.125.504-1.125 1.125v1.5c0 .621.504 1.125 1.125 1.125M12 10.875v-1.5m0 1.5c0 .621-.504 1.125-1.125 1.125M12 10.875c0 .621.504 1.125 1.125 1.125m-2.25 0c.621 0 1.125.504 1.125 1.125M13.125 12h7.5m-7.5 0c-.621 0-1.125.504-1.125 1.125M20.625 12c.621 0 1.125.504 1.125 1.125v1.5c0 .621-.504 1.125-1.125 1.125m-17.25 0h7.5M12 14.625v-1.5m0 1.5c0 .621-.504 1.125-1.125 1.125M12 14.625c0 .621.504 1.125 1.125 1.125m-2.25 0c.621 0 1.125.504 1.125 1.125m0 1.5v-1.5m0 0c0-.621.504-1.125 1.125-1.125m0 0h7.5" />
                        <path v-else stroke-linecap="round" stroke-linejoin="round" d="M3 16.5v2.25A2.25 2.25 0 005.25 21h13.5A2.25 2.25 0 0021 18.75V16.5m-13.5-9L12 3m0 0l4.5 4.5M12 3v13.5" />
                    </svg>

                    <div class="info-up">
                        {{item.desc}}
                        <el-popover
                            placement="top" popper-class="elPopTitle"
                            width="200"
                            trigger="hover"
                            :content="item.popover">
                            <img slot="reference" src="@/assets/images/info.png"/>
                        </el-popover>
                    </div>
                    <div class="info-num">{{item.data}}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    let that
    // import {getStatsStorage} from "@/api/stats";
    import axios from 'axios'
    import { generateState } from '@/utils/i18n'
    export default {
        name: "Stats",
        components: { },
        data() {
            return {
                list: [],
                loading: true,
                loading_ecosystem: true,
                logo_img_black: require("@/assets/images/icons/MCSLOGO.png"),
                collaboratorsData: [
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-01.png"),
                        link: 'https://protocol.ai/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-02.png"),
                        link: 'https://ipfs.tech/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-03.png"),
                        link: 'https://filecoin.io/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-04.png"),
                        link: 'https://polygon.technology/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-05.png"),
                        link: 'https://chainlinklabs.com/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-06.png"),
                        link: 'https://labs.binance.com/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-07.png"),
                        link: 'https://opensea.io/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-08.png"),
                        link: 'https://arbitrum.io/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-09.png"),
                        link: 'https://akash.network/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-10.png"),
                        link: 'https://aptos.dev/'
                    },
                    {
                        img: require("@/assets/images/dashboard/MULTI-CHAIN-11.png"),
                        link: 'https://sui.io/'
                    }
                ],
                MCS_Dataset: [
                    {
                        data: '-',
                        desc: 'Total CIDs Quantity',
                        popover: 'The quantity of sourcefile CIDs on Multichain.storage'
                    },
                    {
                        data: '-',
                        desc: 'Total Sealed Storage ',
                        popover: 'Total sealed storage contributed to Filecoin networks.'
                    },
                    {
                        data: '-',
                        desc: 'Total Pinned IPFS Storage',
                        popover: 'Total pinned ipfs storage for hot retrieval from any ipfs gateway. These data are not stored on Filecoin networks.'
                    },
                    {
                        data: '-',
                        desc: 'Total Registered Users',
                        popover: 'Amount of wallet addresses registered on mutlichain.storage'
                    },
                    {
                        data: '-',
                        desc: 'Total Storage Deal Submitted',
                        popover: 'Total storage orders submitted on mutlichain.storage'
                    },
                    {
                        data: '-',
                        desc: 'Active Successful Storage Deals',
                        popover: 'The quantity of deals stored successfully by miners to Filecoin networks on Multichain.storage.'
                    },
                    {
                        data: '-',
                        desc: 'Total Storage Providers',
                        popover: 'Numbers of miners offering storage capacity to Filecoin networks on mutlichain.storage.'
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
            // that.getStats()
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
                    return false
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
                                // switch (language) {
                                //     case 'cn' :
                                //         that.list[item].title = '平均消息推送费用'
                                //         break
                                //     case 'en' :
                                        that.list[item].title = 'Average Push Message Cost'
                                // }
                                break
                            case 'average_data_cost_sealing_1TB' :
                                // switch (language) {
                                //     case 'cn' :
                                //         that.list[item].title = '每TiB质押费'
                                //         break
                                //     case 'en' :
                                        that.list[item].title = 'Pledge Collateral'
                                // }
                                break
                            case 'average_gas_cost_sealing_1TB' :
                                // switch (language) {
                                //     case 'cn' :
                                //         that.list[item].title = '平均封装手续费'
                                //         break
                                //     case 'en' :
                                        that.list[item].title = 'Cost of Sealing'
                                // }
                                break
                            case 'average_min_piece_size' :
                                // switch (language) {
                                //     case 'cn' :
                                //         that.list[item].title = '平均最小文件'
                                //         break
                                //     case 'en' :
                                        that.list[item].title = 'Average Minimum Piece Size'
                                // }
                                that.list[item].num = res.data.data[item]?res.data.data[item]:'0 GiB'
                                break
                            case 'average_price_per_GB_per_year' :
                                // switch (language) {
                                //     case 'cn' :
                                //         that.list[item].title = '平均数据存储价格'
                                //         break
                                //     case 'en' :
                                        that.list[item].title = 'Average Price'
                                // }
                                that.list[item].num = res.data.data[item]?res.data.data[item]:'0 FIL/GiB/year'
                                break
                            case 'average_verified_price_per_GB_per_year' :
                                // switch (language) {
                                //     case 'cn' :
                                //         that.list[item].title = '平均真实数据存储价格'
                                //         break
                                //     case 'en' :
                                        that.list[item].title = 'Average Verified Price'
                                // }
                                that.list[item].num = res.data.data[item]?res.data.data[item]:'0 FIL/GiB/year'
                                break
                            case 'historical_average_price_regular' :
                                // switch (language) {
                                //     case 'cn' :
                                //         that.list[item].title = '历史平均数据存储价格'
                                //         break
                                //     case 'en' :
                                        that.list[item].title = 'Historical Average Regular Price'
                                // }
                                that.list[item].num = res.data.data[item]?res.data.data[item]:'0 FIL/GiB/100 deals'
                                break
                            case 'historical_average_price_verified' :
                                // switch (language) {
                                //     case 'cn' :
                                //         that.list[item].title = '历史平均真实数据存储价格'
                                //         break
                                //     case 'en' :
                                        that.list[item].title = 'Historical Average Verified Price'
                                // }
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
                        return that.NumFormat(data.source_file_cids)
                    case 1:
                        return that.byteChange(data.sealed_storage)
                    case 2:
                        return that.byteChange(data.pinned_ipfs_size)
                    case 3:
                        return that.NumFormat(data.wallet_count)
                    case 4:
                        return that.NumFormat(data.user_uploads)
                    case 5:
                        return that.NumFormat(data.active_deals)
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
    width: calc(100% - 0.4rem);
    padding: 0.3rem 0.2rem;
    background-color: #000;
    background-image: 
        url(../../assets/images/dashboard/bg_top.png), 
        url(../../assets/images/dashboard/bg_bottom.png);
    background-position: left top, right bottom;
    background-repeat: no-repeat, no-repeat;
    background-size: 27%, 22%;
    .el-loading-mask /deep/{
        background-color: rgba(0, 0, 0, 0.95);
    }
    .stats {
        .title {
            margin: 10px 0.2rem;
            font-size: 0.22rem;
            font-weight: 700;
            color: #fff;
            line-height: 1.5;
        }
        .main {
            font-size: 18px;
            width: 100%;
            padding: 0.1rem 0 0.2rem;
            display: flex;
            flex-wrap: wrap;
            .info-num{
                width: 100%;
                text-align: center;
                font-size: 0.22rem;
                font-weight: 500;
                color: #000;
            }
            img, svg{
                display: block;
                width: 0.35rem;
                margin: 0 auto 0.2rem;
            }
        }
        .info {
            width: calc(25% - 0.2rem);
            padding: 0.2rem 0;
            min-height: 100px;
            display: flex;
            justify-content: center;
            /*align-items: center;*/
            flex-direction: column;
            background: #ffffff;
            border-radius: 0.2rem;
            margin: 0.1rem;
            line-height: 1;
            @media screen and (max-width: 1260px) {
                width: calc(25% - 0.2rem);
            }
            @media screen and (max-width: 1024px) {
                width: calc(33.33% - 0.2rem);
            }
            @media screen and (max-width: 768px) {
                width: calc(50% - 0.2rem);
            }
            @media screen and (max-width: 441px) {
                width: calc(100% - 0.2rem);
            }
        }
        .info-up {
            font-size: 0.2rem;
            color: #2C7FF8;
            width: 96%;
            padding: 0 2%;
            text-align: center;
            font-weight: 500;
            margin-bottom: 0.2rem;
            display: flex;
            align-items: center;    
            justify-content: center;
            img{
                display: block;
                width: 20px;
                height: 20px;
                margin: 0 0 0 5px;
                cursor: pointer;
                @media screen and (max-width:1440px){
                    width: 17px;
                    height: 17px;
                }
            }
        }
        .el-icon-dog{
            background-size: contain;
        }
        .el-icon-dog:before{
            content: "dog";
            font-size: 12px;
            visibility: hidden;
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
                    .height{
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
    .opacity{
        opacity: 0;
    }
    .Eco /deep/{
        padding: 0 0.2rem;
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
            padding: 10px 0 0;
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