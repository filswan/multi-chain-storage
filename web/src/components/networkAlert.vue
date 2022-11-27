<template>
  <el-alert type="warning" effect="dark" center show-icon @close="closeDia">
    <div slot="title" v-if="baseNetwork">
      {{$t('fs3Login.toptip_01')}} {{metaNetworkInfo.name}} {{$t('fs3Login.toptip_02')}}
      <span @click="changeAlert(137)">Polygon Mainnet</span>.
      <p v-if="networkID == 80001 || networkID == 97">
        {{$t('fs3Login.toptip_04_main')}} {{metaNetworkInfo.name}} {{$t('fs3Login.toptip_04_1')}}
        <a href="https://calibration-mcs.filswan.com/#/home" target="_blank">calibration-mcs.filswan.com</a>.
      </p>
    </div>
    <div slot="title" v-else>
      {{$t('fs3Login.toptip_01')}} {{metaNetworkInfo.name}} {{$t('fs3Login.toptip_02')}}
      <span @click="changeAlert(80001)">Mumbai Testnet</span>
      {{$t('fs3Login.toptip_Network')}}
      <span @click="changeAlert(97)">BSC TestNet</span>.
      <p v-if="networkID == 137">{{$t('fs3Login.toptip_04')}}
        <a href="https://www.multichain.storage/#/home" target="_blank">multichain.storage</a>.</p>
    </div>
  </el-alert>
</template>

<script>
let that
export default {
  name: 'networkAlert',
  data () {
    return {

    }
  },
  props: [],
  computed: {
    metaAddress () {
      return this.$store.getters.metaAddress
    },
    networkID () {
      return Number(this.$store.getters.networkID)
    },
    metaNetworkInfo () {
      return this.$store.getters.metaNetworkInfo ? JSON.parse(JSON.stringify(this.$store.getters.metaNetworkInfo)) : {}
    },
    mcsjwtToken () {
      return this.$store.getters.mcsjwtToken
    }
  },
  methods: {
    closeDia () {
      that.$emit('getNetwork', false)
    },
    changeAlert (rows) {
      that.$emit('changeNet', rows)
    }
  },
  mounted () {
    that = this
  }
}
</script>

<style scoped lang="scss">
.el-alert /deep/ {
  position: fixed;
  left: 0;
  top: 0;
  z-index: 99;
  .el-alert__icon {
    @media screen and (min-width: 1600px) {
      font-size: 20px;
      width: 20px;
    }
  }
  .el-alert__content {
    display: flex;
    align-items: center;
    .el-alert__title {
      @media screen and (min-width: 1600px) {
        font-size: 14px;
        line-height: 1.3;
      }
      @media screen and (min-width: 1680px) {
        font-size: 16px;
        line-height: 1.3;
      }
      @media screen and (min-width: 1800px) {
        font-size: 18px;
        line-height: 1.3;
      }
      span {
        text-decoration: underline;
        cursor: pointer;
      }
      a {
        text-decoration: underline;
        color: #fff;
      }
    }
    .el-icon-close {
      @media screen and (min-width: 1600px) {
        font-size: 16px;
        line-height: 1.3;
      }
      @media screen and (min-width: 1800px) {
        font-size: 18px;
        line-height: 1.3;
      }
    }
  }
}
</style>
