<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  beforeCreate () {
    if (localStorage.getItem('headertitle') !== 'Details') sessionStorage.removeItem('dealsPaginationIndexDev')
    if (sessionStorage.getItem('mcs_dev_jwtToken') === '') {
      this.$store.dispatch('setMetaAddress', '')
      this.$store.dispatch('setMCSjwtToken', '')
      this.$store.dispatch('setMetaNetworkId', 0)
      const network = {
        name: '',
        unit: '',
        center_fail: false
      }
      this.$store.dispatch('setMetaNetworkInfo', JSON.stringify(network))
      this.$router.push('/home')
    }
  }
}
</script>

<style lang="scss">

.reverse_phase {
  filter: invert(1) hue-rotate(180deg);
  .el-button, .createTask a, .logged_in .info, .metamaskLogin, .metamask, .sidebar, svg, img.resno,
  .el-popover.el-popper, #billing .form .form_top .search .el-input .el-input__icon,
  #dealManagement .form .form_table .el-table tr th .cell i, .filter_status,
  .header-right .progress, .swithUI svg, .network_mainnet img, .el-alert, .el-backtop,
  .statsCont .info, .statsCont .collaborators img, .metamaskLogin, .metamaskHome{
    filter: invert(1) hue-rotate(180deg);
  }
  .metamaskLogin, .metamaskHome{
    *, img{
      filter: invert(0) hue-rotate(0deg);
    }
  }
  .color, .resno, .sidebar img, .statsCont svg{
    filter: invert(0) hue-rotate(0deg);
  }
  .header-right .progress .tip{
    color: #fff;
  }
}
</style>
