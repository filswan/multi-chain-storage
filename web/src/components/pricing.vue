<template>
    <div class="pricing_body">
        <h1 class="text_h1">{{$t('pricing.title_1')}}</h1>
        <h2 class="text_h2">{{$t('pricing.title_2')}}</h2>
        <h4 class="text_h4">{{$t('pricing.title_3')}}</h4>
        <h4 class="text_h4">{{$t('pricing.title_4')}}</h4>
        <a class="contact" href="mailto:team@filswan.com">{{$t('pricing.btn')}}.</a>
        <el-row :gutter="20" class="bid">
            <el-col :span="12">
                <div class="grid-content">
                    <h3>{{$t('pricing.bid_title_1')}}</h3>
                    <p>{{$t('pricing.bid_desc_1')}}</p>
                    <div class="price">
                        <b>$0</b> USD/forever
                    </div>
                    <el-button type="primary">
                        {{$t('fs3Login.Connect_StartFree')}}
                    </el-button>
                    <div class="introduce" v-for="o in 4" :key="o">
                        <div class="cont">
                            <i class="el-icon-success"></i>
                            <b>xxxxx</b>
                        </div>
                        <div class="cont">
                            <i></i>
                            <span>xxxxx</span>
                        </div>
                    </div>
                </div>
            </el-col>
            <el-col :span="12">
                <div class="grid-content">
                    <h3>{{$t('pricing.bid_title_2')}}</h3>
                    <p>{{$t('pricing.bid_desc_2')}}</p>
                    <div class="price">
                        <b>$0</b> USD/forever
                    </div>
                    <el-button type="info">
                        {{$t('pricing.coming_soon')}}
                    </el-button>
                    <div class="introduce" v-for="o in 4" :key="o">
                        <div class="cont">
                            <i class="el-icon-success"></i>
                            <b>xxxxx</b>
                        </div>
                        <div class="cont">
                            <i></i>
                            <span>xxxxx</span>
                        </div>
                    </div>
                </div>
            </el-col>
        </el-row>
        <el-row :gutter="20" class="sustain">
            <el-col :span="12">
                <div class="grid-content">
                    {{$t('pricing.supported_wallet')}}
                </div>
            </el-col>
            <el-col :span="12">
                <div class="grid-content">
                    {{$t('pricing.supported_wallet')}}
                </div>
            </el-col>
            <el-col :span="12">
                <div class="grid-content">
                    {{$t('pricing.prepay_using')}}
                </div>
            </el-col>
            <el-col :span="12">
                <div class="grid-content">
                    {{$t('pricing.prepay_using')}}
                </div>
            </el-col>
        </el-row>
        <div class="volume">
            <h2>{{$t('pricing.volume_title')}}</h2>
            <p>{{$t('pricing.volume_p')}}</p>
            <el-button type="primary">
                {{$t('pricing.btn')}}
            </el-button>
            <h3>{{$t('pricing.volume_title_1')}}</h3>
            <a class="create">{{$t('pricing.volume_create')}}</a>
        </div>
        <div class="questions">
            <div class="title">{{$t('pricing.question_title')}}</div>
            <el-collapse accordion>
                <el-collapse-item>
                    <template slot="title">
                        {{$t('pricing.collapse_1_title')}}
                    </template>
                    <div>与现实生活一致：与现实生活的流程、逻辑保持一致，遵循用户习惯的语言和概念；</div>
                    <div>在界面中一致：所有的元素和结构需保持一致，比如：设计样式、图标和文本、元素的位置等。</div>
                </el-collapse-item>
                <el-collapse-item title="反馈 Feedback">
                    <div>控制反馈：通过界面样式和交互动效让用户可以清晰的感知自己的操作；</div>
                    <div>页面反馈：操作后，通过页面元素的变化清晰地展现当前状态。</div>
                </el-collapse-item>
                <el-collapse-item title="效率 Efficiency">
                    <div>简化流程：设计简洁直观的操作流程；</div>
                    <div>清晰明确：语言表达清晰且表意明确，让用户快速理解进而作出决策；</div>
                    <div>帮助用户识别：界面简单直白，让用户快速识别而非回忆，减少用户记忆负担。</div>
                </el-collapse-item>
                <el-collapse-item title="可控 Controllability">
                    <div>用户决策：根据场景可给予用户操作建议或安全提示，但不能代替用户进行决策；</div>
                    <div>结果可控：用户可以自由的进行操作，包括撤销、回退和终止当前操作等。</div>
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
    }
  },
  components: {},
  watch: {},
  methods: {},
  mounted () {
    that = this
  },
  filters: {},
  computed: {}
}
</script>

<style scoped lang="scss">
.pricing_body {
  position: relative;
  width: 80%;
  max-width: 1440px;
  min-width: 300px;
  margin: 0.1rem auto;
  padding: 0.5rem 0;
  background-color: #fff;
  color: #000;
  font-size: 16px;
  word-break: break-word;
  @media screen and (max-width: 992px) {
    padding: 0.25rem 0;
  }
  @media screen and (max-width: 768px) {
    font-size: 14px;
  }
  .text_h1 {
    font-size: 25px;
    text-align: center;
    color: #547fe8;
    line-height: 1.5;
  }
  .text_h2 {
    font-size: 20px;
    text-align: center;
    line-height: 1.5;
  }
  .text_h4,
  .contact {
    display: block;
    padding: 0.2rem 0 0;
    font-weight: 500;
    text-align: center;
    line-height: 1.3;
    color: inherit;
    &.contact {
      padding: 0;
    }
  }
  .bid /deep/ {
    margin: 0.3rem auto;
    .el-col {
      padding: 0.25rem;
      .grid-content {
        width: 90%;
        padding: 0.25rem;
        margin: auto;
        border: 1px solid #94aff1;
        border-radius: 0.1rem;
        h3 {
          line-height: 1.5;
        }
        p {
          line-height: 1.2;
        }
        .price {
          padding: 0.25rem 0;
          font-size: 14px;
          b {
            font-size: 24px;
            color: #84a3ef;
          }
        }
        .el-button {
          font-family: inherit;
          font-size: 0.2rem;
          font-weight: normal;
          height: auto;
          padding: 0.1rem 0.15rem;
          margin: 0.3rem 0;
          color: #fff;
          border-radius: 0.5rem;
          line-height: 1.4;
          white-space: normal;
          border: 0;
          @media screen and (max-width: 600px) {
            font-size: 16px;
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
          padding: 0.05rem 0;
          .cont {
            display: flex;
            align-items: flex-start;
            font-size: 16px;
            @media screen and (min-width: 1800px) {
              font-size: 18px;
            }
            i {
              display: block;
              width: 20px;
              height: 20px;
              margin-right: 5px;
              font-size: 20px;
              color: #0da34a;
            }
            span {
              font-size: 14px;
              color: #b1b1b1;
              @media screen and (min-width: 1800px) {
                font-size: 16px;
              }
            }
          }
        }
      }
    }
  }
  .sustain /deep/ {
    margin: 0.3rem auto;
    .el-col {
      padding: 0.05rem;
      line-height: 1.5;
      &:nth-child(2n + 1) {
        text-align: right;
      }
    }
  }
  .volume {
    max-width: 800px;
    padding: 0.4rem 5%;
    margin: 0.8rem auto;
    border: 1px solid #94aff1;
    border-radius: 0.1rem;
    h2 {
      text-align: center;
      color: #5d86ea;
      line-height: 1.5;
    }
    h3 {
      font-weight: normal;
      line-height: 1.5;
    }
    p {
      margin: 0.15rem auto;
      text-align: center;
      line-height: 1.3;
    }
    .create {
      display: block;
      padding: 0.45rem 0;
      margin: 0.1rem 0;
      background-color: #e5e5e5;
      border-radius: 0.1rem;
      text-align: center;
      color: inherit;
    }
    .el-button {
      display: inherit;
      font-family: inherit;
      font-size: 0.2rem;
      font-weight: normal;
      height: auto;
      padding: 0.07rem 0.2rem;
      margin: 0.3rem auto;
      background: #5580e9;
      color: #fff;
      border-radius: 0.5rem;
      line-height: 1.4;
      white-space: normal;
      border: 0;
      @media screen and (max-width: 600px) {
        font-size: 16px;
      }
      &:hover {
        opacity: 0.8;
      }
    }
  }
  .questions {
    margin: 0.3rem auto;
    .title {
      padding: 0 0 0.5rem;
      font-size: 24px;
      font-weight: bolder;
      text-align: center;
    }
    .el-collapse /deep/ {
      border: 1px solid #94aff1;
      border-radius: 0.1rem;
      overflow: hidden;
      .el-collapse-item {
        padding: 0.2rem 3%;
        background-color: #fafafa;
        border-top: 1px solid #94aff1;
        .el-collapse-item__header {
          background-color: inherit;
          font-size: 20px;
          font-weight: bolder;
          border: 0;
        }
        .el-collapse-item__wrap {
          background-color: inherit;
          border: 0;
          .el-collapse-item__content {
            padding: 0;
            font-size: 16px;
            color: #7f7f7f;
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
