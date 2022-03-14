<template>
    <div class="statsCont">
        <div class="stats" v-loading="loading">
            <div class="title">{{generateState('network_overview')}}</div>
            <div class="main">
                <div v-for="(item,key,index) in list" :key="index" class="info">
                    <div class="info-up">
                        {{ item.title }}
                    </div>
                    <div class="info-num">{{ item.num }}</div>
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
                //   'Authorization': "Bearer "+ this.$store.getters.accessToken
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
                            break
                        case 'average_price_per_GB_per_year' :
                            switch (language) {
                                case 'cn' :
                                    this.list[item].title = '平均数据存储价格'
                                    break
                                case 'en' :
                                    this.list[item].title = 'Average Price'
                            }
                            break
                        case 'average_verified_price_per_GB_per_year' :
                            switch (language) {
                                case 'cn' :
                                    this.list[item].title = '平均真实数据存储价格'
                                    break
                                case 'en' :
                                    this.list[item].title = 'Average Verified Price'
                            }
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
        }
    }
</script>

<style lang="scss" scoped>
.statsCont{
    padding: 0.25rem 0.2rem 0.5rem;
    .stats {
        /*width: 100%;*/
        padding: 10px;
        background: white;
        border-radius: 10px;
        .title {
            font-size: 20px;
            border-bottom: 1px solid #ededed;
            padding-left: 20px;
        }

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
            width: calc(33.33% - 30px);
            padding: 10px;
            height: 100px;
            display: flex;
            justify-content: center;
            /*align-items: center;*/
            flex-direction: column;
            background: #f2f3f8;
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
@media screen and (max-width: 800px) {
    .statsCont{
        padding: 0.3rem 4%;
        .stats {
            /*width: 100%;*/
            padding: 10px;
            background: white;
            border-radius: 10px;
            .title {
                font-size: 20px;
                border-bottom: 1px solid #ededed;
                padding-left: 20px;
            }

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
                padding: 10px;
                height: 100px;
                display: flex;
                justify-content: center;
                /*align-items: center;*/
                flex-direction: column;
                background: #f2f3f8;
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
