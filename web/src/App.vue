<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App'
}
</script>

<style lang="scss">


.reverse_phase {
  filter: invert(1) hue-rotate(180deg);
  .el-button, .createTask a, .logged_in .info, .metamaskLogin, .metamask, .sidebar, svg, img.resno,
  .el-popover.el-popper, #billing .form .form_top .search .el-input .el-input__icon,
  #dealManagement .form .form_table .el-table tr th .cell i, .filter_status,
  .header-right .progress{
    filter: invert(1) hue-rotate(180deg);
  }
  .color, .resno, .sidebar img{
    filter: invert(0) hue-rotate(0deg);
  }
  .header-right .progress .tip{
    color: #fff;
  }
}
</style>