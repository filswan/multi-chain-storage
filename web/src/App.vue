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
  .el-button, .createTask a, .logged_in .info, .metamaskLogin, .metamask, .sidebar, svg, img,
  .el-popover.el-popper{
    filter: invert(1) hue-rotate(180deg);
  }
  .color, .resno, .sidebar img{
    filter: invert(0) hue-rotate(0deg);
  }
}
</style>