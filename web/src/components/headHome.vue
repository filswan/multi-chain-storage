<template>
  <div class="headerCont">
    <div class="landing" style="min-height: 0.86rem;">
      <div class="header width">
        <img :src="logo" class="logoImg" alt='FilSwan' @click="header_logo" />
        <!-- 菜单导航 -->
        <el-menu :default-active="activeIndex" class="el-menu-demo pcShow" mode="horizontal" @select="handleSelect" background-color="transparent" text-color="#333" active-text-color="#5580e9">
          <el-menu-item index="about">
            <router-link to="">
              {{ $t('route.About') }}
            </router-link>
          </el-menu-item>
          <el-menu-item index="pricing">
            <router-link to="">
              {{ $t('route.Pricing') }}
            </router-link>
          </el-menu-item>
          <el-menu-item index="resources">
            <router-link to="">
              {{ $t('route.Resources') }}
            </router-link>
          </el-menu-item>
          <el-menu-item index="stats">
            <router-link to="">
              {{ $t('route.Stats') }}
            </router-link>
          </el-menu-item>
          <el-menu-item index="auditReport" @click="goLink('https://github.com/numencyber/AuditReport/blob/main/FILSWAN-Smart-Contract-Audit-Report%20-%20Numen.pdf')">
            <router-link to="">
              {{ $t('route.report') }}
            </router-link>

          </el-menu-item>
          <el-menu-item index="login">
            <a href="javascript:;" v-loading="loginLoad" @click="getLogin" class="target">
              {{ $t('route.Login') }}
            </a>
          </el-menu-item>
        </el-menu>
        <!-- 移动端菜单 -->
        <div class="mobileShow" @click="mobileMenuFun">
          <div>
            <span></span>
            <span></span>
            <span></span>
          </div>
        </div>

        <!-- 移动端导航 -->
        <div class="menuBg" v-if="mobileMenuShow" @click="mobileMenuFun"></div>
        <div :class="{'menuMList': 1==1, 'menuWidth': mobileMenuShow}" v-show="mobileMenuShow">
          <div class="contentMenu" :class="{'contentMenuX': mobileMenuSecord}">
            <div class="first_level_menu">
              <div class="menuMListChild" @click="handleMoSelect('about')">
                {{ $t('route.About') }}
              </div>
              <div class="menuMListChild" @click="handleMoSelect('pricing')">
                {{ $t('route.Pricing') }}
              </div>
              <div class="menuMListChild" @click="handleMoSelect('resources')">
                {{ $t('route.Resources') }}
              </div>
              <div class="menuMListChild" @click="handleMoSelect('stats')">
                {{ $t('route.Stats') }}
              </div>
              <a href="javascript:;" v-loading="loginLoad" @click="getLogin" class="menuMListChild">
                {{ $t('route.Login') }}
              </a>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
let that
export default {
  data () {
    return {
      logo: require('@/assets/images/MCS_logo_2.png'),
      activeIndex: '4',
      mobileMenuShow: false,
      mobileMenuSecord: false
    }
  },
  props: ['loginLoad'],
  computed: {
    routerMenu () {
      return this.$store.getters.routerMenu
    }
  },
  watch: {},
  methods: {
    goLink (link) {
      window.open(link)
    },
    header_logo () {
      that.$router.push({ name: 'home_entrance' })
    },
    handleSelect (key, keyPath) {
      if (key !== 'login') that.$emit('getHome', key)
    },
    handleMoSelect (key) {
      that.$emit('getHome', key)
      that.mobileMenuFun()
    },
    mobileMenuFun () {
      that.mobileMenuShow = !that.mobileMenuShow
    },
    getLogin () {
      that.$emit('getLogin', true)
    }
  },
  mounted () {
    that = this
    window.addEventListener('resize', () => {
      if (document.body.clientWidth > 999) that.mobileMenuShow = false
    })
  }
}
</script>
<style lang="scss" scoped>
.headerCont {
  width: 100%;
  background-color: #fff;
  box-shadow: 0 1px 2px rgba(51, 51, 51, 0.2);
  @media screen and (max-width: 999px) {
    position: fixed;
    min-height: 60px;
    z-index: 999;
  }
  .width {
    display: flex;
    flex-wrap: wrap;
    width: 80%;
    max-width: 1440px;
    min-width: 300px;
    margin: auto;
    @media screen and (max-width: 1150px) {
      width: 90%;
    }
    @media screen and (max-width: 999px) {
      width: 94%;
    }
  }
  .header /deep/ {
    padding: 0.2rem 0;
    display: flex;
    justify-content: space-between;
    align-items: center;
    -webkit-transition: position 0.4s, padding 0s;
    transition: position 0.4s, padding 0s;
    @media screen and (max-width: 1024px) {
      padding: 0.5rem 0;
    }
    @media screen and (max-width: 999px) {
      padding: 0;
      height: 60px;
    }
    .logoImg {
      height: 40px;
      cursor: pointer;
    }
    .menuChildA {
      display: block;
      width: 100%;
    }
    .el-menu.el-menu--horizontal {
      border: 0;
    }
    .el-menu--horizontal > .el-menu-item,
    .el-menu--horizontal > .el-submenu .el-submenu__title {
      height: auto;
      padding: 0;
      margin: 0.05rem 0.05rem 0.05rem 0.6rem;
      line-height: 1;
      font-size: 16px;
      border: 0;
      @media screen and (max-width: 999px) {
        margin: 0.05rem 0.15rem 0.05rem 0.1rem;
      }
      @media screen and (max-width: 600px) {
        font-size: 14px;
      }
      a {
        display: block;
        padding: 0.15rem;
        color: #606060;
        &.target {
          float: left;
          padding: 0.15rem 0.4rem;
          background: #5580e9 !important;
          border: 0;
          border-radius: 0.05rem;
          outline: none;
          color: #fff !important;
          // box-shadow: 0 12px 26px -12px rgba(12, 22, 44, 0.32);
          transition: background-color 0.3s, border-color 0.3s, color 0.3s;
          &:hover {
            opacity: 0.85;
          }
        }
        .el-loading-mask {
          border-radius: 0.05rem;
          z-index: 2;
        }
      }
      &.is-active,
      &:hover {
        background-color: transparent !important;
        color: #5580e9 !important;
        a {
          color: #5580e9;
        }
      }
      .el-submenu__icon-arrow {
        display: none;
      }
    }
  }
}

@media screen and (max-width: 999px) {
  .header {
    .pcShow {
      display: none;
    }
    .mobileShow {
      display: block;
      width: 27px;
      float: right;
      margin-top: 0;
      outline: none;
      cursor: pointer;
      span {
        display: block;
        width: 100%;
        height: 2px;
        margin: 0.07rem 0;
        background-color: #2c367a;
      }
      .close {
        height: 0.35rem;
        display: flex;
        align-items: center;
        justify-content: flex-end;
        i {
          font-size: 20px;
          color: #000;
        }
        .target {
          position: absolute;
          right: 55px;
          height: auto;
          padding: 0 10px;
          font-size: 14px;
          line-height: 1.5;
          border-radius: 15px;
        }
      }
    }
    .menuBg {
      position: fixed;
      right: 0;
      left: 0;
      top: 60px;
      width: 100%;
      height: 100%;
      background: rgba(0, 0, 0, 0.3);
    }
    .menuMList {
      position: fixed;
      right: 0;
      left: auto;
      top: 60px;
      bottom: 0;
      width: 0;
      height: calc(100% - 60px);
      z-index: 99;
      background-color: #fff;
      overflow-x: hidden;
      overflow-y: scroll;
      -webkit-transition: all 0.4s;
      transition: all 0.4s;
      border-top: 1px solid;
      .contentMenu {
        width: 200%;
        -webkit-transition: transform 0.3s;
        transition: transform 0.3s;
        .first_level_menu,
        .secondary_menu {
          float: left;
          width: 50%;
          .menuMListChild {
            display: flex;
            justify-content: space-between;
            align-items: center;
            position: relative;
            width: 90%;
            float: none;
            text-align: left;
            font-size: 16px;
            line-height: 1.5;
            padding: 0.2rem 2%;
            margin: 0 0 0 6%;
            color: #000;
            border-bottom: 1px solid #e4e7ed;
            i {
              float: right;
              color: #101c4f;
              margin: 0 4% 0 0;
              cursor: pointer;
              font-size: 16px;
              padding: 0.1rem 0;
              &:before {
                font-size: 16px;
                font-weight: 500;
              }
            }
            h5 {
              margin: 0 4% 0 0;
              font-size: 0.17rem;
              color: #000;
            }
          }
          .m {
            color: #9baff9;
          }
        }
      }
      .contentMenuX {
        transform: translateX(-50%);
      }
    }
    .menuWidth {
      width: 90%;
    }
    .logoImg {
      height: 30px;
    }
  }
}
@media screen and (max-width: 479px) {
  .header {
    padding: 25px 24px;
  }
}
</style>
