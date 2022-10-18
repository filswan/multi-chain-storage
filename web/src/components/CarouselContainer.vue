<template>
  <div class="CarouselContainer" @mouseenter="stopAutoPlay" @mouseleave="startAutoPlay">
    <div ref="mySwiper" class="swiper-container" :id="currentIndex"  >
      <div class="swiper-wrapper">
        <div class="swiper-slide my-swiper-slide" v-for="(item,index) of slideList" :key="index">
            <a :href="item.link" target="_blank"><img :src="item.img" :class="{'height': index == 4 || index == 6 || index == 7}" alt="logo" /></a>
        </div>
      </div>
      <!-- 分页器 -->
      <div class="swiper-pagination"></div>
      <!--导航器-->
      <div class="swiper-button-prev"></div>
      <div class="swiper-button-next"></div>
    </div>
  </div>
</template>
<script>
import Swiper from 'swiper'
import "swiper/css/swiper.css";
export default {
  name: 'CarouselContainer',
  props: ['slideList','currentIndex'],
  data(){
    return {
      currentSwiper:null,
    }
  },
  watch:{
    //slide数据发生变化时,更新swiper
    slideList:{
      deep:true,
      // eslint-disable-next-line
      handler(nv,ov){
        // console.log("数据更新了")
        this.updateSwiper()
      }
    }
  },
  mounted() {
    this.initSwiper()
  },
  methods:{
    //鼠标移入暂停自动播放
    stopAutoPlay() {
       this.currentSwiper.autoplay.stop()
    },
    //鼠标移出开始自动播放
    startAutoPlay() {
      this.currentSwiper.autoplay.start()
    },
    initSwiper(){
      // eslint-disable-next-line
      let vueComponent=this
      this.currentSwiper = new Swiper('#'+this.currentIndex, {
        loop: true, 
        autoHeight:'true',
        pagination: {
          el: '.swiper-pagination',
          clickable:true,
        },
        navigation: {
          nextEl: '.swiper-button-next',
          prevEl: '.swiper-button-prev',
        },
        autoplay: {
          delay: 3000,
          stopOnLastSlide: false,
          disableOnInteraction: false
        },
        breakpoints: {
            1600: {
                slidesPerView: 7,
                spaceBetween: 0
            },
            1440: {
                slidesPerView: 6,
                spaceBetween: 0
            },
            1280: {
                slidesPerView: 5,
                spaceBetween: 0
            },
            992: {
                slidesPerView: 4,
                spaceBetween: 0
            },
            768: {
                slidesPerView: 3,
                spaceBetween: 0
            },
            640: {
                slidesPerView: 2,
                spaceBetween: 0
            },
            375: {
                slidesPerView: 1,
                spaceBetween: 0
            }
        }
      })
    },
    //销毁swiper
    destroySwiper(){
        try {
          this.currentSwiper.destroy(true,false)
        }catch (e) {
          console.log("删除轮播")
        }
    },
    //更新swiper
    updateSwiper(){
      this.destroySwiper()
      this.$nextTick(()=>{
        this.initSwiper()
      })
    },
  },
  destroyed() {
    this.destroySwiper()
  }
}
</script>
<style lang="scss" scoped>
  .CarouselContainer{
    width: 97%;
    margin: auto;
    .swiper-container  /deep/{
        padding: 20px 0;
        .my-swiper-slide{
            height: 100px;
            @media screen and (max-width: 1600px) {
                height: 85px;
            }
            @media screen and (max-width: 1440px) {
                height: 75px;
            }
            a{
                img{
                    display: block;
                    height: 60px;
                    margin: 20px auto;
                    cursor: pointer;
                    @media screen and (max-width: 1600px) {
                        height: 45px;
                    }
                    @media screen and (max-width: 1440px) {
                        margin: 15px auto;
                    }
                }
                .height{
                    height: 80px;
                    margin: 10px auto;
                    @media screen and (max-width: 1600px) {
                        height: 65px;
                    }
                    @media screen and (max-width: 1440px) {
                        margin: 5px auto;
                    }
                }
            }
        }
        .swiper-pagination{
            display: flex;
            justify-content: center;
            bottom: 5px;
            .swiper-pagination-bullet{
                background: #fff;
            }
            .swiper-pagination-bullet-active{
                background: #007aff;
            }
        }
        .swiper-button-prev{
            left: 0;
        }
        .swiper-button-next{
            right: 0;
        }
        .swiper-button-next, .swiper-button-prev{
            &::after{
                font-size: 20px;
            }
        }
    }
  }
</style>