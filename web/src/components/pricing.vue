<template>
  <div class="pricing_body">
    <h1 class="text_h1">{{$t('pricing.title_1')}}</h1>
    <h2 class="text_h2">{{$t('pricing.title_2')}}</h2>
    <h4 class="text_h4">{{$t('pricing.title_3')}}</h4>
    <a :href="discord_link" target="_blank" class="contact">{{$t('pricing.btn')}}.</a>
    <el-row class="bid">
      <el-col :xs="24" :sm="12" :md="12" :lg="12">
        <div class="grid-content">
          <h3>{{$t('pricing.bid_title_1')}}</h3>
          <p>{{$t('pricing.bid_desc_1')}}</p>
          <div class="price">
            <b>$0</b>USD/forever
            <div class="small"></div>
          </div>
          <el-button type="primary" @click="getLogin">
            {{$t('fs3Login.Connect_StartFree')}}
          </el-button>
          <div class="introduce" v-for="(o, oIndex) in introduce.free" :key="oIndex">
            <div class="cont">
              <i class="el-icon-success"></i>
              <div class="right">
                <b>{{o.title}}</b>
                <span>{{o.desc}}</span>
              </div>
            </div>
          </div>
        </div>
      </el-col>
      <el-col :xs="24" :sm="12" :md="12" :lg="12">
        <div class="grid-content">
          <h3>{{$t('pricing.bid_title_2')}}</h3>
          <p>{{$t('pricing.bid_desc_2')}}</p>
          <div class="price">
            <b>$96</b>USD/annually
            <div class="small"></div>
          </div>
          <el-button type="primary" @click="goLink(discord_link)">
            {{$t('pricing.btn')}}
          </el-button>
          <div class="introduce" v-for="(o, pIndex) in introduce.pro" :key="pIndex">
            <div class="cont">
              <i class="el-icon-success"></i>
              <div class="right">
                <b>{{o.title}}</b>
                <span>{{o.desc}}</span>
              </div>
            </div>
          </div>
        </div>
      </el-col>
    </el-row>
    <el-row :gutter="35" class="sustain">
      <el-col :span="12">
        <div class="grid-content">
          {{$t('pricing.supported_wallet')}}
        </div>
      </el-col>
      <el-col :span="12">
        <div class="grid-content">
          <img src="@/assets/images/pricing/metamask-logo-horizontal.png" class="meta_img" />
        </div>
      </el-col>
      <el-col :span="12">
        <div class="grid-content">
          {{$t('pricing.prepay_using')}}
        </div>
      </el-col>
      <el-col :span="12">
        <div class="grid-content">
          <img src="@/assets/images/pricing/usd-coin.jpg" class="usd_img" />
        </div>
      </el-col>
    </el-row>
    <div class="volume">
      <div class="volume_body">
        <h2>{{$t('pricing.volume_title')}}</h2>
        <p>{{$t('pricing.volume_p')}}</p>
        <el-button type="primary" @click="goLink(discord_link)">
          {{$t('pricing.btn')}}
        </el-button>
        <h3>{{$t('pricing.volume_title_1')}}</h3>
        <img class="img" src="@/assets/images/pricing/GIF3.gif" />
      </div>
    </div>
    <div class="questions">
      <div class="title">{{$t('pricing.question_title')}}</div>
      <el-collapse v-model="activeNames">
        <el-collapse-item name="1">
          <template slot="title">
            {{$t('pricing.collapse_1_title')}}
          </template>
          <div>{{$t('pricing.collapse_1_desc')}}</div>
        </el-collapse-item>
        <el-collapse-item name="2">
          <template slot="title">
            {{$t('pricing.collapse_2_title')}}
          </template>
          <div>{{$t('pricing.collapse_2_desc')}}</div>
        </el-collapse-item>
        <el-collapse-item name="3">
          <template slot="title">
            {{$t('pricing.collapse_3_title')}}
          </template>
          <div>Yes! We provide SDKs for Python and JavaScript. You can check them
            <a @click="goLink('https://docs.filswan.com/multichain.storage/developer-quickstart/sdk')">here</a>
          </div>
        </el-collapse-item>
        <el-collapse-item name="4">
          <template slot="title">
            {{$t('pricing.collapse_4_title')}}
          </template>
          <div>No worries! Our team always open to communicate and customize a special plan for your needs.
            <a @click="goLink(discord_link)">Create a ticket</a> in our Discord community and we will get back to you shortly after.</div>
        </el-collapse-item>
      </el-collapse>
    </div>
  </div>
</template>

<script>
let that
export default {
  name: 'Pricing',
  data () {
    return {
      introduce: {
        free: [
          {
            title: 'Storage - 32GB',
            desc: ''
          },
          {
            title: 'Total Volume - 100 GB',
            desc: 'Total Volume is the total amount of data transferred when someone loads a file through your gateway.'
          },
          {
            title: 'SDK Access',
            desc: ''
          },
          {
            title: 'Unlimited buckets and folders',
            desc: ''
          },
          {
            title: 'Discord ticket support ',
            desc: ''
          }
        ],
        pro: [
          {
            title: 'Storage - 320 GB',
            desc: ''
          },
          {
            title: 'Total Volume - 320 GB',
            desc: 'Total Volume is the total amount of data transferred when someone loads a file through your gateway.'
          },
          {
            title: 'SDK Access',
            desc: ''
          },
          {
            title: 'Unlimited buckets and folders',
            desc: ''
          },
          {
            title: 'Discord ticket support',
            desc: ''
          },
          {
            title: 'Subdomains ',
            desc: 'Subdomains allows you to set up your own name for the link shared'
          }
        ]
      },
      questions: [
        {
          title: this.$t('pricing.collapse_1_title'),
          desc: this.$t('pricing.collapse_1_desc')
        },
        {
          title: this.$t('pricing.collapse_2_title'),
          desc: this.$t('pricing.collapse_2_desc')
        }
      ],
      activeNames: ['1'],
      discord_link: process.env.DISCORD_LINK
    }
  },
  components: {},
  watch: {},
  methods: {
    goLink (link) {
      window.open(link)
    },
    getLogin () {
      that.$emit('getLogin', true)
    }
  },
  mounted () {
    that = this
  },
  filters: {},
  computed: {}
}
</script>

<style scoped lang="scss">
@font-face {
  font-family: "Gilroy-Medium";
  src: url(../assets/font/Gilroy-Medium-2.otf);
  font-style: normal;
  font-display: block;
}
.pricing_body {
  position: relative;
  width: 90%;
  max-width: 1440px;
  min-width: 300px;
  margin: 0.1rem auto;
  padding: 0.8rem 0;
  background-color: #fff;
  font-family: "Gilroy-Medium";
  font-size: 16px;
  word-break: break-word;
  color: #010102;
  line-height: 1.3;
  @media screen and (max-width: 992px) {
    padding: 1.8rem 0 0.25rem;
  }
  @media screen and (max-width: 768px) {
    font-size: 14px;
  }
  .text_h1 {
    font-size: 0.5572rem;
    font-weight: bold;
    text-align: center;
    color: #5580e9;
    @media screen and (max-width: 768px) {
      padding-bottom: 0.1rem;
      line-height: 1;
    }
  }
  .text_h2 {
    font-size: 0.3308rem;
    font-weight: 500;
    text-align: center;
  }
  .text_h4,
  .contact {
    display: block;
    padding: 0.2rem 0 0;
    font-size: 0.194rem;
    font-weight: 500;
    text-align: center;
    @media screen and (max-width: 768px) {
      font-size: 15px;
    }
    &.contact {
      padding: 0;
    }
  }
  .bid /deep/ {
    display: flex;
    align-items: stretch;
    flex-wrap: wrap;
    max-width: 1140px;
    margin: 0.5rem auto 0;
    .el-col {
      padding: 0.5rem;
      @media screen and (max-width: 600px) {
        padding: 0.2rem;
      }
      .grid-content {
        height: calc(100% - 0.7rem);
        padding: 0.35rem;
        margin: auto;
        border: 1px solid #94aff1;
        border-radius: 0.1rem;
        h3 {
          font-size: 0.294rem;
          font-weight: 500;
          line-height: 1.5;
          @media screen and (max-width: 768px) {
            font-size: 17px;
          }
        }
        p {
          font-size: 0.145rem;
          line-height: 1.2;
          @media screen and (min-width: 1800px) {
            font-size: 14.5px;
          }
          @media screen and (max-width: 768px) {
            font-size: 13px;
          }
        }
        .price {
          padding: 0.4rem 0 0;
          font-size: 0.145rem;
          @media screen and (min-width: 1800px) {
            font-size: 14.5px;
          }
          @media screen and (max-width: 768px) {
            font-size: 13px;
          }
          b {
            padding-right: 0.08rem;
            font-size: 0.294rem;
            color: #5580e9;
            @media screen and (max-width: 768px) {
              font-size: 17px;
            }
          }
          .small {
            display: block;
            width: 100%;
            height: 22px;
            font-size: 12px;
            color: #afafaf;
          }
        }
        .el-button {
          font-family: inherit;
          font-size: 0.1724rem;
          font-weight: normal;
          height: auto;
          padding: 0.14rem 0.25rem;
          margin: 0.25rem 0 0.3rem;
          color: #fff;
          border-radius: 0.5rem;
          line-height: 1;
          white-space: normal;
          border: 0;
          @media screen and (max-width: 768px) {
            font-size: 14px;
          }
          &.el-button--primary {
            background: #5580e9;
          }
          &.el-button--info {
            background: #d9d9d9;
          }
          &:hover {
            opacity: 0.8;
          }
        }
        .introduce {
          padding: 0.1rem 0;
          .cont {
            display: flex;
            align-items: flex-start;
            font-size: 0.163rem;
            line-height: 20px;
            @media screen and (max-width: 441px) {
              font-size: 12px;
            }
            i {
              display: block;
              width: 20px;
              height: 20px;
              margin-right: 5px;
              font-size: 20px;
              color: #0da34a;
            }
            .right {
              b {
                line-height: 20px;
              }
              span {
                display: block;
                padding: 0.05rem 0.3rem 0 0;
                font-size: 12px;
                color: #afafaf;
                line-height: 1.2;
                @media screen and (min-width: 1800px) {
                  font-size: 14px;
                }
              }
            }
          }
        }
      }
    }
  }
  .sustain /deep/ {
    margin: 0.55rem auto;
    .el-col {
      display: flex;
      align-items: center;
      min-height: 40px;
      padding: 0.05rem;
      font-size: 0.2229rem;
      line-height: 1.5;
      @media screen and (max-width: 768px) {
        font-size: 15px;
      }
      .grid-content {
        display: flex;
        align-items: center;
        height: 100%;
      }
      .meta_img {
        height: 22px;
        @media screen and (min-width: 1800px) {
          height: 25px;
        }
        @media screen and (max-width: 768px) {
          height: 20px;
        }
        @media screen and (max-width: 600px) {
          height: 18px;
        }
      }
      .usd_img {
        height: 24px;
        @media screen and (min-width: 1800px) {
          height: 27px;
        }
        @media screen and (max-width: 768px) {
          height: 22px;
        }
        @media screen and (max-width: 600px) {
          height: 20px;
        }
      }
      &:nth-child(2n + 1) {
        justify-content: flex-end;
      }
    }
  }
  .volume {
    max-width: 840px;
    padding: 0.4rem 0 0.5rem;
    margin: 0.8rem auto 0.9rem;
    border: 1px solid #94aff1;
    border-radius: 0.1rem;
    .volume_body {
      width: 80%;
      max-width: 650px;
      margin: auto;
      h2 {
        font-size: 0.294rem;
        text-align: center;
        color: #5580e9;
        line-height: 1.5;
        @media screen and (max-width: 768px) {
          font-size: 17px;
        }
      }
      h3 {
        font-size: 0.2rem;
        font-weight: normal;
        line-height: 1.25;
        @media screen and (max-width: 768px) {
          font-size: 15px;
        }
      }
      p {
        margin: 0.15rem auto;
        font-size: 0.2rem;
        text-align: center;
        line-height: 1.3;
        @media screen and (max-width: 768px) {
          font-size: 15px;
        }
      }
      .img {
        display: block;
        width: 100%;
        margin: 0.1rem 0;
        border-radius: 0.15rem;
      }
      .el-button {
        display: inherit;
        font-family: inherit;
        font-size: 0.1724rem;
        font-weight: normal;
        height: auto;
        padding: 0.14rem 0.25rem;
        margin: 0.25rem auto 0.3rem;
        background: #5580e9;
        color: #fff;
        border-radius: 0.5rem;
        line-height: 1;
        white-space: normal;
        border: 0;
        @media screen and (max-width: 768px) {
          font-size: 14px;
        }
        &:hover {
          opacity: 0.8;
        }
      }
    }
  }
  .questions {
    margin: 0 auto 1rem;
    .title {
      padding: 0.05rem 0 0.6rem;
      font-size: 0.3951rem;
      font-weight: 700;
      text-align: center;
      line-height: 1.3;
    }
    .el-collapse /deep/ {
      max-width: 1040px;
      margin: auto;
      border: 1px solid #94aff1;
      border-radius: 0.1rem;
      overflow: hidden;
      .el-collapse-item {
        padding: 0.4rem 7%;
        background-color: #fafafa;
        border-top: 1px solid #94aff1;
        transition: all 0.3s;
        .el-collapse-item__header {
          background-color: inherit;
          font-size: 0.3rem;
          font-weight: bolder;
          border: 0;
          line-height: 1.33;
          color: #010102 !important;
          // @media screen and (max-width: 1440px) {
          //   font-size: 18px;
          // }
          // @media screen and (max-width: 1024px) {
          //   font-size: 16px;
          // }
          // @media screen and (max-width: 768px) {
          //   font-size: 14px;
          // }
        }
        .el-collapse-item__wrap {
          background-color: inherit;
          border: 0;
          .el-collapse-item__content {
            max-width: 800px;
            padding: 0.2rem 0 0;
            font-size: 0.2rem;
            color: #7c7c7c;
            @media screen and (max-width: 768px) {
              font-size: 15px;
            }
            a {
              color: #5580e9;
              text-decoration: underline;
              cursor: pointer;
            }
          }
        }
        &.is-active {
          background-color: #fff;
        }
        .el-icon-arrow-right:before {
          content: "\e6d9";
        }
        .el-collapse-item__arrow.is-active {
          transform: none;
          &::before {
            content: "\e6d8";
          }
        }
        &:nth-child(1) {
          border-top: 0;
        }
      }
    }
  }
}
</style>
