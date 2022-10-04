<template>
    <div class="statsCont">
        <div class="stats" v-loading="loading">
            <div class="title">{{generateState('network_overview')}}</div>
            <div class="main">
                <div v-for="(item,key,index) in list" :key="index" class="info">
                    <img src="@/assets/images/icon_shangzhang.png" alt="">
                    <div class="info-up">
                        {{ item.title }}
                    </div>
                    <div class="info-num">{{ item.num | NumFormat }}</div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    // import {getStatsStorage} from "@/api/stats";
    import axios from 'axios'
    import { generateState } from '@/utils/i18n'
    export default {
        name: "stats",
        data() {
            return {
                list: [],
                loading: true
            }
        },
        created() {
            let stats_api = `${process.env.BASE_API}stats/storage?wallet_address=${this.$store.getters.metaAddress}`

            axios.get(stats_api, {
                headers: {
                        'Authorization': "Bearer "+ this.$store.getters.mcsjwtToken
                },
            }).then(res => {
                this.list = this.deepClone(res.data.data)
                let language = localStorage.getItem('languageMcs') ? localStorage.getItem('languageMcs') : 'en'
                for (let item in res.data.data){
                    this.list[item] = {
                        title: '',
                        num: 0
                    }
                    this.list[item].num = res.data.data[item]
                    switch (item) {
                        case 'average_cost_push_message' :
                            switch (language) {
                                case 'cn' :
                                    this.list[item].title = '平均消息推送费用'
                                    break
                                case 'en' :
                                    this.list[item].title = 'Average Push Message Cost'
                            }
                            break
                        case 'average_data_cost_sealing_1TB' :
                            switch (language) {
                                case 'cn' :
                                    this.list[item].title = '每TiB质押费'
                                    break
                                case 'en' :
                                    this.list[item].title = 'Pledge Collateral'
                            }
                            break
                        case 'average_gas_cost_sealing_1TB' :
                            switch (language) {
                                case 'cn' :
                                    this.list[item].title = '平均封装手续费'
                                    break
                                case 'en' :
                                    this.list[item].title = 'Cost of Sealing'
                            }
                            break
                        case 'average_min_piece_size' :
                            switch (language) {
                                case 'cn' :
                                    this.list[item].title = '平均最小文件'
                                    break
                                case 'en' :
                                    this.list[item].title = 'Average Minimum Piece Size'
                            }
                            this.list[item].num = res.data.data[item]?res.data.data[item]:'0 GiB'
                            break
                        case 'average_price_per_GB_per_year' :
                            switch (language) {
                                case 'cn' :
                                    this.list[item].title = '平均数据存储价格'
                                    break
                                case 'en' :
                                    this.list[item].title = 'Average Price'
                            }
                            this.list[item].num = res.data.data[item]?res.data.data[item]:'0 FIL/GiB/year'
                            break
                        case 'average_verified_price_per_GB_per_year' :
                            switch (language) {
                                case 'cn' :
                                    this.list[item].title = '平均真实数据存储价格'
                                    break
                                case 'en' :
                                    this.list[item].title = 'Average Verified Price'
                            }
                            this.list[item].num = res.data.data[item]?res.data.data[item]:'0 FIL/GiB/year'
                            break
                        case 'historical_average_price_regular' :
                            switch (language) {
                                case 'cn' :
                                    this.list[item].title = '历史平均数据存储价格'
                                    break
                                case 'en' :
                                    this.list[item].title = 'Historical Average Regular Price'
                            }
                            this.list[item].num = res.data.data[item]?res.data.data[item]:'0 FIL/GiB/100 deals'
                            break
                        case 'historical_average_price_verified' :
                            switch (language) {
                                case 'cn' :
                                    this.list[item].title = '历史平均真实数据存储价格'
                                    break
                                case 'en' :
                                    this.list[item].title = 'Historical Average Verified Price'
                            }
                            this.list[item].num = res.data.data[item]?res.data.data[item]:'0 FIL/GiB/100 deals'
                            break
                    }
                }
                this.loading = false
            })
        },
        mounted() {
            document.getElementById('content-box').scrollTop = 0
            this.$store.dispatch('setRouterMenu', 4)
            this.$store.dispatch('setHeadertitle', this.$t('navbar.Stats'))
        },
        methods:{
            generateState,
            deepClone(obj) {
                const _obj = JSON.stringify(obj)
                return JSON.parse(_obj)
            }
        },
        filters: {
            NumFormat(value) {
                if (!value) return "0";
                return value;
            },
        }
    }
</script>

<style lang="scss" scoped>
.statsCont{
    padding: 0.3rem 0.2rem;
    .stats {
        .title {
            margin: 0 0.1rem;
            font-size: 0.22rem;
            font-weight: 700;
            color: #333;
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
