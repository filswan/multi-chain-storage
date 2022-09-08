<template>
        <el-dialog :modal="true" :width="widthDia" :visible.sync="networkC" :before-close="closeDia">
            <div slot="title" class="dialog-title">
                {{$t('network.title')}}
            </div>
            <div class="upload_form" v-loading="loading">
              <button class="listCont" :class="{'active': networkID == netItem.id}" v-for="(netItem, index) in listData" :key="index" @click="mainnetChange(netItem.id)">
                <div class="ePHqjy">
                  <div class="jaQBhE">
                    <div class="left">
                      <div class="fHDDin">
                        <img :alt="netItem.name" :src="netItem.img">
                      </div>
                      <div id="chain_list_name_1" class="cMfoPw">
                        <div class="dXIEwo"> 
                          <div class="feVFfA" v-if="networkID == netItem.id">
                            <div class="bmDlFW"><div></div></div>
                          </div>
                          {{netItem.name}}
                        </div>
                      </div>
                      <div id="chain_list_url_1" class="dqwsCT">
                        <input id="chain_list_input_1" class="bZSSTA" value="https://ethmainnet.anyswap.exchange"></div>
                      </div>
                      <!-- <div class="right">
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" id="chain_list_set_1" class="sc-dymIpo lhTXwO">
                          <circle cx="12" cy="12" r="3"></circle>
                          <path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"></path>
                        </svg>
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" id="chain_list_tick_1" class="sc-gFaPwZ lmpCoU">
                          <polyline points="9 11 12 14 22 4"></polyline>
                          <path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11"></path>
                        </svg>
                        <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="sc-bnXvFD dJvmUI">
                          <polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"></polygon>
                        </svg>
                      </div> -->
                    </div>
                  </div>
              </button>
            </div>
        </el-dialog>
</template>

<script>
let that;
export default {
  name: "networkChange",
  data() {
    return {
      widthDia: document.body.clientWidth <= 600 ? "95%" : document.body.clientWidth <= 1600 ? "520px" : "640px",
      loading: false,
      listData: [
        {
          img: require('@/assets/images/network_logo/polygon.png'),
          name: 'Mumbai Testnet',
          id: 80001
        },
        {
          img: require('@/assets/images/network_logo/bsc.png'),
          name: 'BSC TestNet',
          id: 97
        },
        {
          img: require('@/assets/images/network_logo/polygon.png'),
          name: 'Polygon Mainnet',
          id: 137
        }
      ]
    };
  },
  props: ["networkC"],
  computed: {
      metaAddress() {
          return this.$store.getters.metaAddress
      },
      networkID() {
          return this.$store.getters.networkID
      }
  },
  methods: {
    closeDia() {
      that.$emit("getNetworkC", false);
    },
    mainnetChange(id) {
      that.$emit("getNetworkC", false, id);
    }
  },
  mounted() {
    that = this
  }
};
</script>


<style scoped lang="scss">
.el-dialog__wrapper /deep/ {
  display: flex;
  align-items: center;
  justify-content: center;
  .el-dialog {
    background: #fff;
    margin: auto !important;
    box-shadow: 0 0 13px rgba(128, 128, 128, 0.8);
    border-radius: 0.2rem;
    overflow: hidden;
    .el-dialog__header {
      padding: 0.3rem 0.4rem;
      display: flex;
      border-bottom: 1px solid #dfdfdf;
      .dialog-title{
          display: flex;
          align-items: center;
          color: #333;
          font-size: 0.22rem;
          font-weight: 500;
          line-height: 1;
          text-transform: capitalize;
          .el-tooltip{
              width: 20px;
              height: 20px;
              margin: 0 0 0 5px;
              @media screen and (min-width:1800px){
                  width: 22px;
                  height: 22px;
              }
              @media screen and (max-width:1440px){
                  width: 17px;
                  height: 17px;
              }
              @media screen and (max-width: 1280px){
                  width: 16px;
                  height: 16px;
              }
          }
      }
      .el-dialog__headerbtn {
        display: none;
      }
    }
    .el-dialog__body {
      padding: 0;
      .upload_form{
        .listCont{
          background-color: #f6f6f6;
          outline: none;
          border-width: 0.01rem;
          border-style: solid;
          border-color: transparent transparent #dfdfdf;
          border-image: initial;
          cursor: pointer;
          display: flex;
          flex-direction: row;
          -webkit-box-align: center;
          align-items: center;
          -webkit-box-pack: justify;
          justify-content: space-between;
          padding: 0px;
          margin-top: 0px;
          opacity: 1;
          width: 100% !important;
          &:hover, &.active{
            background-color: #e5eeff;
          }
          .ePHqjy {
            width: 100%;
            display: flex;
            -webkit-box-pack: justify;
            justify-content: space-between;
            -webkit-box-align: center;
            align-items: center;
            .jaQBhE {
              width: 100%;
              display: flex;
              -webkit-box-pack: justify;
              justify-content: space-between;
              -webkit-box-align: center;
              align-items: center;
              .left {
                display: flex;
                -webkit-box-pack: start;
                justify-content: flex-start;
                -webkit-box-align: center;
                align-items: center;
                width: 80%;
                padding: 0.15rem 0rem 0.15rem 0.4rem;
                .fHDDin {
                  display: flex;
                  flex-flow: column nowrap;
                  -webkit-box-align: center;
                  align-items: center;
                  -webkit-box-pack: center;
                  justify-content: center;
                  margin-right: 0.2rem;
                  border: 2px solid rgba(0, 0, 0, 0.1);
                  background: rgb(255, 255, 255);
                  width: 36px;
                  min-width: 36px;
                  height: 36px;
                  border-radius: 100%;
                  img {
                    width: 36px;
                    height: 36px;
                    min-width: 36px;
                    min-height: 36px;
                    max-width: 100%;
                    max-height: 100%;
                    background-color: white;
                    border-radius: 36px;
                  }
                }
                .cMfoPw {
                  display: flex;
                  flex-flow: column nowrap;
                  -webkit-box-pack: center;
                  justify-content: center;
                  height: 100%;
                  width: 100%;
                  .dXIEwo {
                    display: flex;
                    flex-flow: row nowrap;
                    color: #333;
                    font-size: 0.2rem;
                    font-weight: 500;
                    .feVFfA {
                      color: rgb(39, 174, 96);
                      display: flex;
                      -webkit-box-pack: center;
                      justify-content: center;
                      -webkit-box-align: center;
                      align-items: center;
                      .bmDlFW {
                        display: flex;
                        flex-flow: row nowrap;
                        -webkit-box-pack: center;
                        justify-content: center;
                        -webkit-box-align: center;
                        align-items: center;
                        &:first-child {
                          height: 8px;
                          width: 8px;
                          margin-right: 8px;
                          background-color: rgb(39, 174, 96);
                          border-radius: 50%;
                        }
                      }
                    }
                  }
                }
                .dqwsCT {
                  flex-flow: column nowrap;
                  -webkit-box-pack: center;
                  justify-content: center;
                  height: 100%;
                  width: 100%;
                  display: none;
                  .bZSSTA {
                    outline: none;
                    border-top: none;
                    border-right: none;
                    border-left: none;
                    border-image: initial;
                    flex: 1 1 auto;
                    height: 45px;
                    width: 100%;
                    background-color: transparent;
                    border-bottom: 0.0625rem solid rgb(92, 103, 125);
                    color: rgb(255, 255, 255);
                    overflow: hidden;
                    text-overflow: ellipsis;
                    font-size: 14px;
                    font-weight: 500;
                    font-stretch: normal;
                    font-style: normal;
                    line-height: 1;
                    letter-spacing: -0.0625rem;
                    padding: 8px 0.75rem;
                  }
                }
              }
            }
          }
        }
      }
    }
    .dialog-footer {
      display: flex;
      align-items: center;
      justify-content: flex-end;
    }
  }
}
@media screen and (max-width: 599px) {
  .el-dialog__wrapper /deep/ {
    .el-dialog {
      .el-dialog__header {
        .el-dialog__title {
          font-size: 0.16rem;
        }
      }
    }
  }
}
</style>
